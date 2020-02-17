package main

import (
    "context"
    "github.com/rexray/gocsi"
    "github.com/arturoguerra/gologging"
    config "github.com/arturoguerra/xcpng-csi/internal/config"
    service "github.com/arturoguerra/xcpng-csi/pkg/csi/service"
    service "github.com/arturoguerra/xcpng-csi/pkg/csi/provider"
)

func main() {
    log := logging.New()
    cfg, err := config.Load()
    if err != nil {
        log.Fatal(err)
    }

    xclient := xapi.New(cfg.Username, cfg.Password, cfg.Host)

    log.Info("Starting GoCSI for XCP-ng...")
    gocsi.Run(
        context.Backgorund(),
        service.Name,
        "CSI Driver for XCP-ng",
        "",
        provider.New(xclient, cfg.NodeID),
    )
}
