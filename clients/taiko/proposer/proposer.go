package proposer

import (
	"context"
	"fmt"
	"github.com/taikoxyz/hive-taiko-clients/clients"
	"github.com/taikoxyz/hive-taiko-clients/clients/utils"
)

const (
// PortProposerTCP    = 9000
// PortProposerUDP    = 9000
// PortProposerAPI    = 4000
// PortProposerGRPC   = 4001
// PortProposerGRPC   = 4001
// PortMetrics      = 8080
// PortValidatorAPI = 5000
)

type ProposerClientConfig struct {
	ClientIndex int
	Subnet      string
}

type ProposerClient struct {
	clients.Client
	Logger  utils.Logging
	Config  ProposerClientConfig
	Builder interface{}
}

func (pn *ProposerClient) Logf(format string, values ...interface{}) {
	if l := pn.Logger; l != nil {
		l.Logf(format, values...)
	}
}

func (pn *ProposerClient) Start() error {
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

func (pn *ProposerClient) Init(ctx context.Context) error {
	//if pn.api == nil {
	//	port := pn.Config.ProposerAPIPort
	//	if port == 0 {
	//		port = PortProposerAPI
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
	//			if cfg, err := pn.ProposerConfig(ctx); err == nil && cfg != nil {
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

func (pn *ProposerClient) Shutdown() error {
	if managedClient, ok := pn.Client.(clients.ManagedClient); !ok {
		return fmt.Errorf("attempted to shutdown an unmanaged client")
	} else {
		return managedClient.Shutdown()
	}
}

type ProposerClients []*ProposerClient

// Return subset of clients that are currently running
func (all ProposerClients) Running() ProposerClients {
	res := make(ProposerClients, 0)
	for _, bc := range all {
		if bc.IsRunning() {
			res = append(res, bc)
		}
	}
	return res
}

// Return subset of clients that are part of an specific subnet
func (all ProposerClients) Subnet(subnet string) ProposerClients {
	if subnet == "" {
		return all
	}
	res := make(ProposerClients, 0)
	for _, pn := range all {
		if pn.Config.Subnet == subnet {
			res = append(res, pn)
		}
	}
	return res
}
