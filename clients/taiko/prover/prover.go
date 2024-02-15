package prover

import (
	"context"
	"fmt"
	"github.com/taikoxyz/hive-taiko-clients/clients"
	"github.com/taikoxyz/hive-taiko-clients/clients/utils"
)

const (
// PortProverTCP    = 9000
// PortProverUDP    = 9000
// PortProverAPI    = 4000
// PortProverGRPC   = 4001
// PortProverGRPC   = 4001
// PortMetrics      = 8080
// PortValidatorAPI = 5000
)

type ProverClientConfig struct {
	ClientIndex int
	Subnet      string
}

type ProverClient struct {
	clients.Client
	Logger          utils.Logging
	Config          ProverClientConfig
	Builder         interface{}
	startupComplete bool
}

func (pn *ProverClient) Logf(format string, values ...interface{}) {
	if l := pn.Logger; l != nil {
		l.Logf(format, values...)
	}
}

func (pn *ProverClient) Start() error {
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

func (pn *ProverClient) Init(ctx context.Context) error {
	if !pn.startupComplete {
		defer func() {
			pn.startupComplete = true
		}()
	}
	//if pn.api == nil {
	//	port := pn.Config.ProverAPIPort
	//	if port == 0 {
	//		port = PortProverAPI
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
	//			if cfg, err := pn.ProverConfig(ctx); err == nil && cfg != nil {
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

func (pn *ProverClient) Shutdown() error {
	if managedClient, ok := pn.Client.(clients.ManagedClient); !ok {
		return fmt.Errorf("attempted to shutdown an unmanaged client")
	} else {
		return managedClient.Shutdown()
	}
}

type ProverClients []*ProverClient

// Return subset of clients that are currently running
func (all ProverClients) Running() ProverClients {
	res := make(ProverClients, 0)
	for _, bc := range all {
		if bc.IsRunning() {
			res = append(res, bc)
		}
	}
	return res
}

// Return subset of clients that are part of an specific subnet
func (all ProverClients) Subnet(subnet string) ProverClients {
	if subnet == "" {
		return all
	}
	res := make(ProverClients, 0)
	for _, pn := range all {
		if pn.Config.Subnet == subnet {
			res = append(res, pn)
		}
	}
	return res
}
