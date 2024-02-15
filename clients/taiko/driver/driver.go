package driver

import (
	"context"
	"fmt"
	"github.com/taikoxyz/hive-taiko-clients/clients"
	"github.com/taikoxyz/hive-taiko-clients/clients/utils"
)

const (
// PortDriverTCP    = 9000
// PortDriverUDP    = 9000
// PortDriverAPI    = 4000
// PortDriverGRPC   = 4001
// PortMetrics      = 8080
// PortValidatorAPI = 5000
)

type DriverClientConfig struct {
	ClientIndex int
	Subnet      string
}

type DriverClient struct {
	clients.Client
	Logger          utils.Logging
	Config          DriverClientConfig
	Builder         interface{}
	startupComplete bool
}

func (dn *DriverClient) Logf(format string, values ...interface{}) {
	if l := dn.Logger; l != nil {
		l.Logf(format, values...)
	}
}

func (dn *DriverClient) Start() error {
	if !dn.IsRunning() {
		if managedClient, ok := dn.Client.(clients.ManagedClient); !ok {
			return fmt.Errorf("attempted to start an unmanaged client")
		} else {
			if err := managedClient.Start(); err != nil {
				return err
			}
		}
	}

	return dn.Init(context.Background())
}

func (dn *DriverClient) Init(ctx context.Context) error {
	if !dn.startupComplete {
		defer func() {
			dn.startupComplete = true
		}()
	}
	//if dn.api == nil {
	//	port := dn.Config.DriverAPIPort
	//	if port == 0 {
	//		port = PortDriverAPI
	//	}
	//	dn.api = &eth2api.Eth2HttpClient{
	//		Addr:  dn.GetAddress(),
	//		Cli:   &http.Client{},
	//		Codec: eth2api.JSONCodec{},
	//	}
	//}
	//
	//var wg sync.WaitGroup
	//var errs = make(chan error, 2)
	//if dn.Config.Spec == nil {
	//	// Try to fetch config directly from the client
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for {
	//			if cfg, err := dn.DriverConfig(ctx); err == nil && cfg != nil {
	//				if spec, err := SpecFromConfig(cfg); err != nil {
	//					errs <- err
	//					return
	//				} else {
	//					dn.Config.Spec = spec
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
	//if dn.Config.GenesisTime == nil || dn.Config.GenesisValidatorsRoot == nil {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		for {
	//			if gen, err := dn.GenesisConfig(ctx); err == nil &&
	//				gen != nil {
	//				dn.Config.GenesisTime = &gen.GenesisTime
	//				dn.Config.GenesisValidatorsRoot = &gen.GenesisValidatorsRoot
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

func (dn *DriverClient) Shutdown() error {
	if managedClient, ok := dn.Client.(clients.ManagedClient); !ok {
		return fmt.Errorf("attempted to shutdown an unmanaged client")
	} else {
		return managedClient.Shutdown()
	}
}

type DriverClients []*DriverClient

// Return subset of clients that are currently running
func (all DriverClients) Running() DriverClients {
	res := make(DriverClients, 0)
	for _, bc := range all {
		if bc.IsRunning() {
			res = append(res, bc)
		}
	}
	return res
}

// Return subset of clients that are part of an specific subnet
func (all DriverClients) Subnet(subnet string) DriverClients {
	if subnet == "" {
		return all
	}
	res := make(DriverClients, 0)
	for _, dn := range all {
		if dn.Config.Subnet == subnet {
			res = append(res, dn)
		}
	}
	return res
}
