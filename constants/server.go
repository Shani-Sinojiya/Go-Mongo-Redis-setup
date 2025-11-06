package constants

import "accounts.workflecks.com/config"

var (
	// server
	ServerPort = config.GetConfig().GetString("server.port")
	ServerMode = config.GetConfig().GetString("server.mode")
)
