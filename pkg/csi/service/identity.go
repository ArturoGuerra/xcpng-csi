package service

import (
    "context"
    "github.com/golang/protobuf/ptypes/wrappers"
    "github.com/container-storage-interface/spec/lib/go/csi"
)

func (s *service) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
    return &csi.ProbeResponse{
        Ready: &wrappers.BoolValue{Value: true},
    }, nil
}

func (s *service) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
    log.Info("Getting plugin info")
    return &csi.GetPluginInfoResponse{
        Name:          Name,
        VendorVersion: VendorVersion,
        Manifest:      Manifest,
    }, nil
}

func (s *service) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
    log.Info("Getting Plugin Capabilities")
    return &csi.GetPluginCapabilitiesResponse{
        Capabilities: []*csi.PluginCapability{
            &csi.PluginCapability{
                Type: &csi.PluginCapability_Service_{
                    Service: &csi.PluginCapability_Service{
                        Type: csi.PluginCapability_Service_CONTROLLER_SERVICE,
                    },
                },
            },
        },
    }, nil
}
