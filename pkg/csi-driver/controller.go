package csi-driver

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


func (s *service) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
}

func (s *service) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
}

func (s *service) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
}

func (s *service) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
}

func (s *service) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
}

func (s *service) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
}

func (s *service) GetCapacity(ctx context.Context, req *csi.GetCapacitiyRequest) (*csi.GetCapacityResponse, error) {
}
