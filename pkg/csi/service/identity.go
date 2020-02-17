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
    return &csi.GetPluginInfoResponse{
        Name:          Name,
        VendorVersion: VendorVersion,
        Manifest:      Manifest,
    }, nil
}

func (s *service) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
    return &csi.GetPluginCapabilitiesResponse{
        Capabilities: []*csi.PluginCapability{},
    }, nil
}
