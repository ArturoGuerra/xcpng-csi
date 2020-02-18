package service

/*
Controller Implements
CreateVolume
DeleteVolume
ControllerPublishVolume
ControllerUnpublishVolume
ValidateVolumeCapabilities
ListVolumes
GetCapacity
*/

import (
    "context"
    "google.golang.org/grpc/status"
    "google.golang.org/grpc/codes"
    "github.com/container-storage-interface/spec/lib/go/csi"
)


func (s *service) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
    return &csi.ControllerGetCapabilitiesResponse{
        Capabilities: []*csi.ControllerServiceCapability{
            {
                Type: &csi.ControllerServiceCapability_Rpc{
                    Rpc: &csi.ControllerServiceCapability_RPC{
                        Type: csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
                    },
                },
            },
            {
                Type: &csi.ControllerServiceCapability_Rpc{
                    Rpc: &csi.ControllerServiceCapability_RPC{
                        Type: csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
                    },
                },
            },
        },
    }, nil
}

func (s *service) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
    name := req.GetName()
    params, err := s.ParseParams(req.GetParameters())
    if err != nil {
        log.Error(err)
        return nil, err
    }

    // Calculates disk size in bytes and sets a min size of 5Gi
    volSizeBytes := int64(minSize)
    if req.GetCapacityRange() != nil && req.GetCapacityRange().RequiredBytes != 0 {
        if int64(req.GetCapacityRange().GetRequiredBytes()) > volSizeBytes {
            log.Info("Setting custom disk size")
            volSizeBytes = int64(req.GetCapacityRange().GetRequiredBytes())
        }
    }

    if err = s.XClient.CreateVolume(name, params.SR, params.FSType, int(volSizeBytes)); err != nil {
        log.Error(err)
        return nil, err
    }

    return &csi.CreateVolumeResponse{
        Volume: &csi.Volume{
            VolumeId: name,
            CapacityBytes: volSizeBytes,
            VolumeContext: req.GetParameters(),
        },
    }, nil
}

func (s *service) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {

    if err := s.XClient.DeleteVolume(req.GetVolumeId()); err != nil {
        log.Error(err)
        return nil, err
    }

    return &csi.DeleteVolumeResponse{}, nil
}

func (s *service) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
    return &csi.ControllerPublishVolumeResponse{}, nil
}

func (s *service) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
    return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (s *service) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
    return &csi.ValidateVolumeCapabilitiesResponse{}, nil
}


// Unimplemented

func (s *service) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}


func (s *service) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}



func (s *service) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
}
