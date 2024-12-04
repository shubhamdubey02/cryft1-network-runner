package local

import (
	"context"
	"crypto"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/MetalBlockchain/metalgo/ids"
	"github.com/MetalBlockchain/metalgo/message"
	"github.com/MetalBlockchain/metalgo/network/peer"
	"github.com/MetalBlockchain/metalgo/network/throttling"
	"github.com/MetalBlockchain/metalgo/snow/networking/router"
	"github.com/MetalBlockchain/metalgo/snow/networking/tracker"
	"github.com/MetalBlockchain/metalgo/snow/validators"
	"github.com/MetalBlockchain/metalgo/staking"
	"github.com/MetalBlockchain/metalgo/utils/constants"
	"github.com/MetalBlockchain/metalgo/utils/crypto/bls"
	"github.com/MetalBlockchain/metalgo/utils/ips"
	"github.com/MetalBlockchain/metalgo/utils/logging"
	"github.com/MetalBlockchain/metalgo/utils/math/meter"
	"github.com/MetalBlockchain/metalgo/utils/resource"
	"github.com/MetalBlockchain/metalgo/utils/set"
	"github.com/MetalBlockchain/metalgo/version"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shubhamdubey02/cryft1-network-runner/api"
	"github.com/shubhamdubey02/cryft1-network-runner/network/node"
	"github.com/shubhamdubey02/cryft1-network-runner/network/node/status"
)

var (
	_ getConnFunc = defaultGetConnFunc
	_ node.Node   = (*localNode)(nil)
)

type getConnFunc func(context.Context, node.Node) (net.Conn, error)

const (
	peerMsgQueueBufferSize      = 1024
	peerResourceTrackerDuration = 10 * time.Second
	peerStartWaitTimeout        = 30 * time.Second
)

// Gives access to basic node info, and to most metalgo apis
type localNode struct {
	// Must be unique across all nodes in this network.
	name string
	// [nodeID] is this node's Avalannche Node ID.
	// Set in network.AddNode
	nodeID ids.NodeID
	// The ID of the network this node exists in
	networkID uint32
	// Allows user to make API calls to this node.
	client api.Client
	// The process running this node.
	process NodeProcess
	// The API port
	apiPort uint16
	// The P2P (staking) port
	p2pPort uint16
	// Returns a connection to this node
	getConnFunc getConnFunc
	// The data dir of the node
	dataDir string
	// The db dir of the node
	dbDir string
	// The logs dir of the node
	logsDir string
	// The plugin dir of the node
	pluginDir string
	// The node config
	config node.Config
	// The node httpHost
	httpHost string
	// maps from peer ID to peer object
	attachedPeers map[string]peer.Peer
	// signals that the process is stopped but the information is valid
	// and can be resumed
	paused bool
}

func defaultGetConnFunc(ctx context.Context, node node.Node) (net.Conn, error) {
	dialer := net.Dialer{}
	return dialer.DialContext(ctx, constants.NetworkType, net.JoinHostPort(node.GetURL(), fmt.Sprintf("%d", node.GetP2PPort())))
}

// AttachPeer: see Network
func (node *localNode) AttachPeer(ctx context.Context, router router.InboundHandler) (peer.Peer, error) {
	tlsCert, err := staking.NewTLSCert()
	if err != nil {
		return nil, err
	}
	tlsConfg := peer.TLSConfig(*tlsCert, nil)
	clientUpgrader := peer.NewTLSClientUpgrader(tlsConfg, prometheus.NewCounter(prometheus.CounterOpts{}))
	conn, err := node.getConnFunc(ctx, node)
	if err != nil {
		return nil, err
	}
	mc, err := message.NewCreator(
		logging.NoLog{},
		prometheus.NewRegistry(),
		"",
		constants.DefaultNetworkCompressionType,
		10*time.Second,
	)
	if err != nil {
		return nil, err
	}

	metrics, err := peer.NewMetrics(
		logging.NoLog{},
		"",
		prometheus.NewRegistry(),
	)
	if err != nil {
		return nil, err
	}
	resourceTracker, err := tracker.NewResourceTracker(
		prometheus.NewRegistry(),
		resource.NoUsage,
		meter.ContinuousFactory{},
		peerResourceTrackerDuration,
	)
	if err != nil {
		return nil, err
	}
	signerIP := ips.NewDynamicIPPort(net.IPv6zero, 0)
	tls := tlsCert.PrivateKey.(crypto.Signer)
	blsKey, err := bls.NewSecretKey()
	if err != nil {
		return nil, err
	}
	config := &peer.Config{
		Metrics:              metrics,
		MessageCreator:       mc,
		Log:                  logging.NoLog{},
		InboundMsgThrottler:  throttling.NewNoInboundThrottler(),
		Network:              peer.TestNetwork,
		Router:               router,
		VersionCompatibility: version.GetCompatibility(node.networkID),
		MySubnets:            set.Set[ids.ID]{},
		Beacons:              validators.NewManager(),
		NetworkID:            node.networkID,
		PingFrequency:        constants.DefaultPingFrequency,
		PongTimeout:          constants.DefaultPingPongTimeout,
		MaxClockDifference:   time.Minute,
		ResourceTracker:      resourceTracker,
		IPSigner:             peer.NewIPSigner(signerIP, tls, blsKey),
	}
	peerID, conn, cert, err := clientUpgrader.Upgrade(conn)
	if err != nil {
		return nil, err
	}

	p := peer.Start(
		config,
		conn,
		cert,
		peerID,
		peer.NewBlockingMessageQueue(
			config.Metrics,
			logging.NoLog{},
			peerMsgQueueBufferSize,
		),
	)
	cctx, cancel := context.WithTimeout(ctx, peerStartWaitTimeout)
	err = p.AwaitReady(cctx)
	cancel()
	if err != nil {
		return nil, err
	}

	node.attachedPeers[p.ID().String()] = p
	return p, nil
}

func (node *localNode) SendOutboundMessage(ctx context.Context, peerID string, content []byte, op uint32) (bool, error) {
	attachedPeer, ok := node.attachedPeers[peerID]
	if !ok {
		return false, fmt.Errorf("peer with ID %s is not attached here", peerID)
	}
	msg := NewTestMsg(message.Op(op), content, false)
	return attachedPeer.Send(ctx, msg), nil
}

// See node.Node
func (node *localNode) GetName() string {
	return node.name
}

// See node.Node
func (node *localNode) GetNodeID() ids.NodeID {
	return node.nodeID
}

// See node.Node
func (node *localNode) GetAPIClient() api.Client {
	return node.client
}

// See node.Node
func (node *localNode) GetURL() string {
	if node.httpHost == "0.0.0.0" || node.httpHost == "." {
		return "0.0.0.0"
	}
	return "127.0.0.1"
}

// See node.Node
func (node *localNode) GetP2PPort() uint16 {
	return node.p2pPort
}

// See node.Node
func (node *localNode) GetAPIPort() uint16 {
	return node.apiPort
}

func (node *localNode) Status() status.Status {
	return node.process.Status()
}

// See node.Node
func (node *localNode) GetBinaryPath() string {
	return node.config.BinaryPath
}

// See node.Node
func (node *localNode) GetPluginDir() string {
	return node.pluginDir
}

// See node.Node
func (node *localNode) GetDataDir() string {
	return node.dataDir
}

// See node.Node
// TODO rename method so linter doesn't complain.
func (node *localNode) GetDbDir() string { //nolint
	return node.dbDir
}

// See node.Node
func (node *localNode) GetLogsDir() string {
	return node.logsDir
}

// See node.Node
func (node *localNode) GetConfigFile() string {
	return node.config.ConfigFile
}

// See node.Node
func (node *localNode) GetConfig() node.Config {
	return node.config
}

// See node.Node
func (node *localNode) GetFlag(k string) (string, error) {
	var v string
	if node.config.ConfigFile != "" {
		var configFileMap map[string]interface{}
		if err := json.Unmarshal([]byte(node.config.ConfigFile), &configFileMap); err != nil {
			return "", err
		}
		vIntf, ok := configFileMap[k]
		if ok {
			v, ok = vIntf.(string)
			if !ok {
				return "", fmt.Errorf("unexpected type for %q expected string got %T", k, vIntf)
			}
		}
	} else if node.config.Flags != nil {
		vIntf, ok := node.config.Flags[k]
		if ok {
			v, ok = vIntf.(string)
			if !ok {
				return "", fmt.Errorf("unexpected type for %q expected string got %T", k, vIntf)
			}
		}
	}
	return v, nil
}

// See node.Node
func (node *localNode) GetPaused() bool {
	return node.paused
}
