package csi-driver

func (s *service) Probe(ctx context.Context, req *csi.ProbeRequest) (*csi.ProbeResponse, error) {
}

func (s *service) GetPluginInfo(ctx context.Context, req *csi.GetPluginInfoRequest) (*csi.GetPluginInfoResponse, error) {
}

func (s *service) GetPluginCapabilities(ctx context.Context, req *csi.GetPluginCapabilitiesRequest) (*csi.GetPluginCapabilitiesResponse, error) {
}
