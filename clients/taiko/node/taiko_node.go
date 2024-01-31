package node

import (
	"fmt"
	"github.com/taikoxyz/hive-taiko-clients/clients/execution"
	"github.com/taikoxyz/hive-taiko-clients/clients/taiko/driver"
	"github.com/taikoxyz/hive-taiko-clients/clients/taiko/proposer"
	"github.com/taikoxyz/hive-taiko-clients/clients/taiko/protocol_deployer"
	"github.com/taikoxyz/hive-taiko-clients/clients/taiko/prover"
	"github.com/taikoxyz/hive-taiko-clients/clients/utils"
)

// A node bundles together:
// - Running L1/L2 Execution clients
// - Running Protocol Deployer
// - Running Driver client
// - Running Proposer client
// - Running Prover client
type TaikoNode struct {
	// Logging interface for all the events that happen in the node
	Logging utils.Logging
	// Index of the node in the network/testnet
	Index int
	// Clients that comprise the node
	L1ExecutionClient          *execution.ExecutionClient
	L2ExecutionClient          *execution.ExecutionClient
	L1L2ProtocolDeployerClient *protocol_deployer.ProtocolDeployerClient
	L2DriverClient             *driver.DriverClient
	L2ProposerClient           *proposer.ProposerClient
	L2ProverClient             *prover.ProverClient
}

func (n *TaikoNode) Logf(format string, values ...interface{}) {
	if l := n.Logging; l != nil {
		l.Logf(format, values...)
	}
}

// Starts all clients included in the bundle
func (n *TaikoNode) Start() error {
	n.Logf("Starting L1 execution client in bundle %d", n.Index)
	if n.L1ExecutionClient != nil {
		if err := n.L1ExecutionClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L1 execution client started in bundle %d", n.Index)
	}

	n.Logf("Starting L2 execution client in bundle %d", n.Index)
	if n.L2ExecutionClient != nil {
		if err := n.L2ExecutionClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L2 execution client started in bundle %d", n.Index)
	}

	n.Logf("Starting L1L2 protocol deployer client in bundle %d", n.Index)
	if n.L1L2ProtocolDeployerClient != nil {
		if err := n.L1L2ProtocolDeployerClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L1L2 protocol deployer client started in bundle %d", n.Index)
	}

	n.Logf("Starting L2 driver client in bundle %d", n.Index)
	if n.L2DriverClient != nil {
		if err := n.L2DriverClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L2 driver client started in bundle %d", n.Index)
	}

	n.Logf("Starting L2 proposer client in bundle %d", n.Index)
	if n.L2ProposerClient != nil {
		if err := n.L2ProposerClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L2 proposer client started in bundle %d", n.Index)
	}

	n.Logf("Starting L2 prover client in bundle %d", n.Index)
	if n.L2ProverClient != nil {
		if err := n.L2ProverClient.Start(); err != nil {
			return err
		}
	} else {
		n.Logf("No L2 prover client started in bundle %d", n.Index)
	}
	return nil
}

func (n *TaikoNode) Shutdown() error {
	if err := n.L1ExecutionClient.Shutdown(); err != nil {
		return err
	}
	if err := n.L2ExecutionClient.Shutdown(); err != nil {
		return err
	}
	if err := n.L1L2ProtocolDeployerClient.Shutdown(); err != nil {
		return err
	}
	if err := n.L2DriverClient.Shutdown(); err != nil {
		return err
	}
	if err := n.L2ProposerClient.Shutdown(); err != nil {
		return err
	}
	if err := n.L2ProverClient.Shutdown(); err != nil {
		return err
	}
	return nil
}

func (n *TaikoNode) ClientNames() string {
	var name string
	if n.L1ExecutionClient != nil {
		name = n.L1ExecutionClient.ClientType()
	}
	if n.L2ExecutionClient != nil {
		name = fmt.Sprintf("%s/%s", name, n.L2ExecutionClient.ClientType())
	}
	if n.L2DriverClient != nil {
		name = fmt.Sprintf("%s/%s", name, n.L2DriverClient.ClientType())
	}
	if n.L2ProposerClient != nil {
		name = fmt.Sprintf("%s/%s", name, n.L2ProposerClient.ClientType())
	}
	if n.L2ProverClient != nil {
		name = fmt.Sprintf("%s/%s", name, n.L2ProverClient.ClientType())
	}
	return name
}

func (n *TaikoNode) IsRunning() bool {
	return n.L1ExecutionClient.IsRunning() && n.L2ExecutionClient.IsRunning() && n.L2DriverClient.IsRunning() && n.L2ProposerClient.IsRunning() && n.L2ProverClient.IsRunning()
}

// Node cluster operations
type TaikoNodes []*TaikoNode

// Return all L1 execution clients, even the ones not currently running
func (all TaikoNodes) L1ExecutionClients() execution.ExecutionClients {
	en := make(execution.ExecutionClients, 0)
	for _, n := range all {
		if n.L1ExecutionClient != nil {
			en = append(en, n.L1ExecutionClient)
		}
	}
	return en
}

// Return all L2 execution clients, even the ones not currently running
func (all TaikoNodes) L2ExecutionClients() execution.ExecutionClients {
	en := make(execution.ExecutionClients, 0)
	for _, n := range all {
		if n.L2ExecutionClient != nil {
			en = append(en, n.L2ExecutionClient)
		}
	}
	return en
}

// Return all L2 driver clients, even the ones not currently running
func (all TaikoNodes) L2DriverClients() driver.DriverClients {
	en := make(driver.DriverClients, 0)
	for _, n := range all {
		if n.L2DriverClient != nil {
			en = append(en, n.L2DriverClient)
		}
	}
	return en
}

// Return all L2 proposer clients, even the ones not currently running
func (all TaikoNodes) L2ProposerClients() proposer.ProposerClients {
	en := make(proposer.ProposerClients, 0)
	for _, n := range all {
		if n.L2ProposerClient != nil {
			en = append(en, n.L2ProposerClient)
		}
	}
	return en
}

// Return all L2 prover clients, even the ones not currently running
func (all TaikoNodes) L2ProverClients() prover.ProverClients {
	en := make(prover.ProverClients, 0)
	for _, n := range all {
		if n.L2ProverClient != nil {
			en = append(en, n.L2ProverClient)
		}
	}
	return en
}

// Return all proxy pointers, even the ones not currently running
//func (all Nodes) Proxies() execution.Proxies {
//	ps := make(execution.Proxies, 0)
//	for _, n := range all {
//		if n.ExecutionClient != nil {
//			ps = append(ps, n.ExecutionClient)
//		}
//	}
//	return ps
//}

// Return subset of nodes that are currently running
func (all TaikoNodes) Running() TaikoNodes {
	res := make(TaikoNodes, 0)
	for _, n := range all {
		if n.IsRunning() {
			res = append(res, n)
		}
	}
	return res
}

//func (all Nodes) FilterByEL(filters []string) Nodes {
//	ret := make(Nodes, 0)
//	for _, n := range all {
//		for _, filter := range filters {
//			if strings.Contains(n.ExecutionClient.ClientType(), filter) {
//				ret = append(ret, n)
//				break
//			}
//		}
//	}
//	return ret
//}
