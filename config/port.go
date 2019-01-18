package config

import (
	"os"
	"sync"
)

// Port ..
type Port struct {
	CmdHTTPPort string
	CmdGRPCPort string
	QryHTTPPort string
	QryGRPCPort string
}

var confPort *Port
var portInit sync.Once

// GetPort ..
func GetPort() *Port {
	portInit.Do(func() {
		confPort = &Port{
			CmdHTTPPort: os.Getenv("CMD_HTTP_PORT"),
			CmdGRPCPort: os.Getenv("CMD_GRPC_PORT"),
			QryHTTPPort: os.Getenv("QRY_HTTP_PORT"),
			QryGRPCPort: os.Getenv("QRY_GRPC_PORT"),
		}
	})
	return confPort
}
