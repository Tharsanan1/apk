package main

import (
	"github.com/wso2/apk/gateway/mediation/internal/config"
	"github.com/wso2/apk/gateway/mediation/internal/extproc"
)

func main() {
	cfg := config.GetConfig()
	cfg.Logger.Info("Mediation server started")
	extproc.StartExternalProcessingServer(cfg)
	cfg.Logger.Info("External processing server started")
	// Wait forever
	select {}
}
