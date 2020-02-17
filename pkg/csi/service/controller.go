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
    "github.com/container-storage-interface/spec/lib/go/csi"
)

func (s *service) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
    name := req.GetName()
//    params, err := s.ParseParams(req.GetParameters())
//    if err != nil {
//        log.Error(err)
//        return nil, err
//    }

    log.Info(name)


    return &csi.CreateVolumeRespone{
        Volume: &csi.Volume{
            VolumeId: name,
            CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
            VolumeContext: req.GetParameters(),
        },
    }, nil
}

func (s *service) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
    return &csi.DeleteVolumeResponse{}, nil
}

func (s *service) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
    return &csi.ControllerPublishVolumeResponse{}, nil
}

func (s *service) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
    return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (s *service) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
    return &csi.ValidateVolumeCapabilities{}, nil
}

func (s *service) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
    return &csi.ListVolumes{}, nil
}

func (s *service) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
    return &csi.GetCapacity{}, nil
}
