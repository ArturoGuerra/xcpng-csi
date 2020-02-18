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
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
    "github.com/container-storage-interface/spec/lib/go/csi"
    "context"
)

func (s *service) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
    log.Info("Getting Node Capabilities")
    return &csi.NodeGetCapabilitiesResponse{}, nil
}

func (s *service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
    return &csi.NodeStageVolumeResponse{}, nil
}

func (s *service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
    return &csi.NodeUnstageVolumeResponse{}, nil
}

func (s *service) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
    return &csi.NodePublishVolumeResponse{}, nil
}

func (s *service) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
    return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (s *service) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
    log.Infof("Getting Node Info")
    return &csi.NodeGetInfoResponse{
        NodeId: s.NodeID,
    }, nil
}

func (s *service) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}
