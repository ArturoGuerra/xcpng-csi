package main

import (
    "context"
    "github.com/rexray/gocsi"
    "github.com/arturoguerra/go-logging"
    "github.com/arturoguerra/xcpng-csi/pkg/xapi"
    config "github.com/arturoguerra/xcpng-csi/internal/config"
    service "github.com/arturoguerra/xcpng-csi/pkg/csi/service"
    provider "github.com/arturoguerra/xcpng-csi/pkg/csi/provider"
)

func main() {
    log := logging.New()
    cfg := config.Load()

    xclient := xapi.New(cfg.Username, cfg.Password, cfg.Host)

    log.Info("Starting GoCSI for XCP-ng...")
    gocsi.Run(
        context.Background(),
        service.Name,
        "CSI Driver for XCP-ng",
        "",
        provider.New(xclient, cfg.NodeID),
    )
}
