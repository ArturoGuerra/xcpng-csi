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
	"context"
	"fmt"

	"github.com/arturoguerra/xcpng-csi/pkg/mounter"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *service) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
	return &csi.NodeGetCapabilitiesResponse{
		Capabilities: []*csi.NodeServiceCapability{
			{
				Type: &csi.NodeServiceCapability_Rpc{
					Rpc: &csi.NodeServiceCapability_RPC{
						Type: csi.NodeServiceCapability_RPC_STAGE_UNSTAGE_VOLUME,
					},
				},
			},
		},
	}, nil
}

// NodeStateVolume Mounts to a common directory for pods to bind mount to
func (s *service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	log.Info("Running NodeStageVolume")
	stagingTargetPath := req.GetStagingTargetPath()
	publishContext := req.GetPublishContext()
	params, err := s.ParseParams(req.GetVolumeContext())
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "")
	}

	device, ok := publishContext["device"]
	if !ok {
		return nil, status.Error(codes.NotFound, "")
	}

	// we should have device at this time
	if device == "" {
		log.Info("No device specified")
		return nil, status.Error(codes.Internal, "No device specified")
	}

	/* Check if path exists */
	notPath, err := mounter.IsNotExist(stagingTargetPath)
	if notPath {
		/* Path doesn't exist so we are creating it and assuming nothing is mounted to it */
		if err = mounter.MakeDir(stagingTargetPath); err != nil {
			log.Error(err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	opts := ""
	if err := mounter.FormatAndMount(device, stagingTargetPath, params.FSType, opts); err != nil {
		log.Error(err)
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("Error when mounting Device: %s, Path: %s, FSType: %s, Error: %v", device, stagingTargetPath, params.FSType, err),
		)
	}
	log.Infof("Mounted Device: [%s] to General Path: (%s)", device, stagingTargetPath)
	return &csi.NodeStageVolumeResponse{}, nil
}

// NodeUnstageVolume unmounts the common directory
func (s *service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	log.Info("Running NodeUnstageVolume")
	stagingTargetPath := req.GetStagingTargetPath()

	if len(stagingTargetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume Staging Target Path must be provided")
	}

	log.Info("Unmounting GLOBAL path")

	if err := mounter.CleanUnmount(stagingTargetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeUnstageVolume DONE: %s", req.GetVolumeId())

	return &csi.NodeUnstageVolumeResponse{}, nil
}

// NodePublishVolume Bind Mounts from staging to pod mount path
func (s *service) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	log.Info("Running NodePublishVolume")
	stagingTargetPath := req.GetStagingTargetPath()
	targetPath := req.GetTargetPath()

	if len(stagingTargetPath) == 0 {
		log.Infof("Missing staging target path: %s", stagingTargetPath)
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume Staging Target Path must be provided")
	}

	if len(targetPath) == 0 {
		log.Infof("Missing target path: %s", targetPath)
		return nil, status.Error(codes.InvalidArgument, "NodePublishVolume Target Path must be provided")
	}

	/* Check if target is a path and creates it if its not */
	notPath, err := mounter.IsNotExist(targetPath)
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	/* If its not a path then nothing is mounted */
	if notPath {
		if err = mounter.MakeDir(targetPath); err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	opts := ""
	if err = mounter.BindMount(stagingTargetPath, targetPath, "auto", opts); err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("Bind Mounted Volume: (%s) to Path: (%s)", stagingTargetPath, targetPath)
	return &csi.NodePublishVolumeResponse{}, nil
}

// NodeUnpublishVolume unmounts the bind mount between the pod and common directory
func (s *service) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Info("Running NodeUnpublishVolume")
	targetPath := req.GetTargetPath()
	if len(targetPath) == 0 {
		return nil, status.Error(codes.InvalidArgument, "NodeUnpublishVolume Target Path must be provided")
	}

	log.Info("Unmounting BIND path")

	if err := mounter.CleanUnmount(targetPath); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodeUnpublishVolume DONE: %s", req.GetVolumeId())

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (s *service) NodeGetInfo(ctx context.Context, req *csi.NodeGetInfoRequest) (*csi.NodeGetInfoResponse, error) {
	log.Infof("Getting NodeInfo for %s", s.NodeID)
	topology := new(csi.Topology)

	/*
	   failure-domain.beta.kubernetes.io/zone
	   topology.kubernetes.io/zone
	*/

	if nodeInfo := s.XClient.GetNodeInfo(s.NodeID); nodeInfo != nil {
		log.Info("Nodeinfo:")
		log.Infof(" - NodeID: %s", nodeInfo.NodeID)
		log.Infof(" - NodeUUID: %s", nodeInfo.NodeUUID)
		log.Infof(" - ZoneID: %s", nodeInfo.Zone)
		log.Infof(" - ZoneUUID: %s", nodeInfo.ZoneUUID)

		AccessibleTopology := make(map[string]string)
		AccessibleTopology[ZoneLabel] = nodeInfo.Zone
		AccessibleTopology[ZoneUUID] = nodeInfo.ZoneUUID
		AccessibleTopology[NodeUUID] = nodeInfo.NodeUUID
		AccessibleTopology[NodeName] = nodeInfo.NodeID

		topology.Segments = AccessibleTopology
	}

	return &csi.NodeGetInfoResponse{
		NodeId:             s.NodeID,
		AccessibleTopology: topology,
	}, nil
}

func (s *service) NodeGetVolumeStats(ctx context.Context, req *csi.NodeGetVolumeStatsRequest) (*csi.NodeGetVolumeStatsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (s *service) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
