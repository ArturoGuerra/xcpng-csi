package service

/*
Node Service
NodeStageVolume Unimplemented
NodeUnstageVolume Unimplemented
NodePublishVolume
NodeUnpublishVolume
NodeGetInfo
NodeGetCapabilities
NodeGetUsage Unimplemented ATM
*/

import (
    "github.com/container-storage-interface/spec/lib/go/csi"
    "context"
)

func (s *service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
    return &csi.NodeStageVolumeResponse{}, nil
}

func (s *service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
    return &csi.NodeUnstageVolumeResponse{}, nil
}

func (s *service) NodePublish(ctx context.Context, req *csi.NodePublishRequest) (*csi.NodePublishResponse, error) {
    return &csi.NodePublishResponse{}, nil
}

func (s *service) NodeUnpublish(ctx context.Context, req *csi.NodeUnpublishRequest) (*csi.NodeUnpublishResponse, error) {
    return &csi.NodeUnpublishResponse{}, nil
}

func (s *service) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
    return &csi.NodeGetInfoResponse{}, nil
}

func (s *service) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
    return &csi.NodeGetCapabilities{}, nil
}

func (s *service) NodeGetUsage(ctx context.Context, req *csi.NodeGetUsageRequest) (*csi.NodeGetUsageResponse, error) {
    return &csi.NodeGetUsageResponse{}, nil
}
