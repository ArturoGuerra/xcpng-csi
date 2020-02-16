package csi-driver

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

func (s *service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
}

func (s *service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
}

func (s *service) NodePublish(ctx context.Context, req *csi.NodePublishRequest) (*csi.NodePublishResponse, error) {
}

func (s *service) NodeUnpublish(ctx context.Context, req *csi.NodeUnpublishRequest) (*csi.NodeUnpublishResponse, error) {
}

func (s *service) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
}

func (s *service) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
}

func (s *service) NodeGetUsage(ctx context.Context, req *csi.NodeGetUsageRequest) (*csi.NodeGetUsageResponse, error) {
}
