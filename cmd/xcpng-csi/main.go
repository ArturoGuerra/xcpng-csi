package main

import (
	"context"

	"github.com/arturoguerra/go-logging"
	config "github.com/arturoguerra/xcpng-csi/internal/config"
	provider "github.com/arturoguerra/xcpng-csi/pkg/csi/provider"
	service "github.com/arturoguerra/xcpng-csi/pkg/csi/service"
	"github.com/arturoguerra/xcpng-csi/pkg/xapi"
	"github.com/rexray/gocsi"
)

/*
NodeID is passed as an env variable though the downwards api.

NOTE: Its important that the node hostname and the xcp-ng vm name are the same for attachment to work, this may be configurable in the future though the config file located in the node
*/

func main() {
	log := logging.New()
	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	// Ensures we always have a node ID
	if cfg.NodeID == "" {
		log.Fatal("Invalid Node ID")
	}

	log.Infof("NodeID: %s", cfg.NodeID)

	xclient := xapi.New(cfg.Regions)

	log.Info("Starting GoCSI for XCP-ng...")
	gocsi.Run(
		context.Background(),
		service.Name,
		"CSI Driver for XCP-ng",
		"",
		provider.New(xclient, cfg.NodeID),
	)
}
