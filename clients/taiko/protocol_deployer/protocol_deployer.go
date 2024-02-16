package protocol_deployer

import (
	"context"
	"fmt"
	"github.com/taikoxyz/hive-taiko-clients/clients"
	"github.com/taikoxyz/hive-taiko-clients/clients/utils"
	"net"
	"time"
)

const (
// PortProtocolDeployerTCP    = 9000
// PortProtocolDeployerUDP    = 9000
// PortProtocolDeployerAPI    = 4000
// PortProtocolDeployerGRPC   = 4001
// PortProtocolDeployerGRPC   = 4001
// PortMetrics      = 8080
// PortValidatorAPI = 5000
)

type ProtocolDeployerClientConfig struct {
	ClientIndex int
	Subnet      string
}

type ProtocolDeployerClient struct {
	clients.Client
	Logger          utils.Logging
	Config          ProtocolDeployerClientConfig
	Builder         interface{}
	startupComplete bool
}

type ProtocolDeployerClients []*ProtocolDeployerClient

func (pn *ProtocolDeployerClient) Logf(format string, values ...interface{}) {
	if l := pn.Logger; l != nil {
		l.Logf(format, values...)
	}
}

func (pn *ProtocolDeployerClient) Start() error {
	if !pn.IsRunning() {
		if managedClient, ok := pn.Client.(clients.ManagedClient); !ok {
			return fmt.Errorf("attempted to start an unmanaged client")
		} else {
			if err := managedClient.Start(); err != nil {
				return err
			}
		}
	}

	return pn.Init(context.Background())
}

func (pn *ProtocolDeployerClient) Init(ctx context.Context) error {
	if !pn.startupComplete {
		defer func() {
			pn.startupComplete = true
		}()
		for {
			port := "8545"
			_, err := net.Dial("tcp", ":"+port)
			if err == nil {
				fmt.Printf("Connection on port %s is open\n", port)
				break
			}
			fmt.Printf("Waiting for connection on port %s...\n", port)
			time.Sleep(5 * time.Second)
		}
	}
	//if pn.api == nil {
	//	port := pn.Config.ProtocolDeployerAPIPort
	//	if port == 0 {
	//		port = PortProtocolDeployerAPI
	//	}
	//	pn.api = &eth2api.Eth2HttpClient{
	//		Addr:  pn.GetAddress(),
	//		Cli:   &http.Client{},
	//		Codec: eth2api.JSONCodec{},
	//	}
	//}
	//
	//var wg sync.WaitGroup
	//var errs = make(chan error, 2)
	//if pn.Config.Spec == nil {
	//	// Try to fetch config directly from the client
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for {
	//			if cfg, err := pn.ProtocolDeployerConfig(ctx); err == nil && cfg != nil {
	//				if spec, err := SpecFromConfig(cfg); err != nil {
	//					errs <- err
	//					return
	//				} else {
	//					pn.Config.Spec = spec
	//					return
	//				}
	//			}
	//			select {
	//			case <-ctx.Done():
	//				errs <- ctx.Err()
	//				return
	//			case <-time.After(time.Second):
	//			}
	//		}
	//	}()
	//}
	//
	//if pn.Config.GenesisTime == nil || pn.Config.GenesisValidatorsRoot == nil {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for {
	//			if gen, err := pn.GenesisConfig(ctx); err == nil &&
	//				gen != nil {
	//				pn.Config.GenesisTime = &gen.GenesisTime
	//				pn.Config.GenesisValidatorsRoot = &gen.GenesisValidatorsRoot
	//				return
	//			}
	//			select {
	//			case <-ctx.Done():
	//				errs <- ctx.Err()
	//				return
	//			case <-time.After(time.Second):
	//			}
	//		}
	//	}()
	//}
	//wg.Wait()
	//
	//select {
	//case err := <-errs:
	//	return err
	//default:
	//	return nil
	//}
	return nil
}

func (pn *ProtocolDeployerClient) Shutdown() error {
	if managedClient, ok := pn.Client.(clients.ManagedClient); !ok {
		return fmt.Errorf("attempted to shutdown an unmanaged client")
	} else {
		return managedClient.Shutdown()
	}
}

// Return subset of clients that are currently running
func (all ProtocolDeployerClients) Running() ProtocolDeployerClients {
	res := make(ProtocolDeployerClients, 0)
	for _, bc := range all {
		if bc.IsRunning() {
			res = append(res, bc)
		}
	}
	return res
}

// Return subset of clients that are part of an specific subnet
func (all ProtocolDeployerClients) Subnet(subnet string) ProtocolDeployerClients {
	if subnet == "" {
		return all
	}
	res := make(ProtocolDeployerClients, 0)
	for _, pn := range all {
		if pn.Config.Subnet == subnet {
			res = append(res, pn)
		}
	}
	return res
}
