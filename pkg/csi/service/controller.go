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
    "github.com/arturoguerra/xcpng-csi/pkg/errs"
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
        return nil, status.Error(codes.InvalidArgument, "")
    }

    // Calculates disk size in bytes and sets a min size of 5Gi
    volSizeBytes := int64(minSize)
    if req.GetCapacityRange() != nil && req.GetCapacityRange().RequiredBytes != 0 {
        if int64(req.GetCapacityRange().GetRequiredBytes()) > volSizeBytes {
            log.Info("Setting custom disk size")
            volSizeBytes = int64(req.GetCapacityRange().GetRequiredBytes())
        }
    }

    VolId, err := s.XClient.CreateVolume(name, params.SR, params.FSType, int(volSizeBytes))
    if err != nil {
        log.Error(err)
        return nil, status.Error(codes.Internal, "")
    }

    return &csi.CreateVolumeResponse{
        Volume: &csi.Volume{
            VolumeId: VolId,
            CapacityBytes: volSizeBytes,
            VolumeContext: req.GetParameters(),
        },
    }, nil
}

func (s *service) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {

    if err := s.XClient.DeleteVolume(req.GetVolumeId()); err != nil {
        log.Error(err)
        return nil, status.Error(codes.Internal, "")
    }

    return &csi.DeleteVolumeResponse{}, nil
}

func (s *service) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
    params, err := s.ParseParams(req.GetVolumeContext())
    if err != nil {
        log.Error(err)
        return nil, status.Error(codes.InvalidArgument, "")
    }

    device, err := s.XClient.Attach(req.GetVolumeId(), req.GetNodeId(), "rw", params.FSType)
    if err != nil {
        log.Error(err)
        switch err.Error() {
        case errs.InvalidVolume:
            return nil, status.Error(codes.NotFound, "")
        case errs.InvalidNode:
            return nil, status.Error(codes.NotFound, "")
        case errs.AlreadyExists:
            return nil, status.Error(codes.AlreadyExists, "")
        default:
            return nil, status.Error(codes.Internal, "")
        }
    }

    log.Infof("VM Device: %s", device)

    context := req.GetVolumeContext()
    context["device"] = device

    return &csi.ControllerPublishVolumeResponse{
        PublishContext: context,
    }, nil
}

func (s *service) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
    if err := s.XClient.Detach(req.GetVolumeId(), req.GetNodeId()); err != nil {
        log.Error(err)
        return nil, status.Error(codes.NotFound, "")
    }

    return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (s *service) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
    return nil, status.Error(codes.Unimplemented, "")
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
