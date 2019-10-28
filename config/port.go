package config

import (
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
			CmdHTTPPort: (*confEnv)["CMD_HTTP_PORT"],
			CmdGRPCPort: (*confEnv)["CMD_GRPC_PORT"],
			QryHTTPPort: (*confEnv)["QRY_HTTP_PORT"],
			QryGRPCPort: (*confEnv)["QRY_GRPC_PORT"],
		}
	})
	return confPort
}
