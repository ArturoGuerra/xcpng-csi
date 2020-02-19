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
    "github.com/arturoguerra/xcpng-csi/pkg/mounter"
    "github.com/container-storage-interface/spec/lib/go/csi"
    "context"
    "fmt"
)

func (s *service) NodeGetCapabilities(ctx context.Context, req *csi.NodeGetCapabilitiesRequest) (*csi.NodeGetCapabilitiesResponse, error) {
    log.Info("Getting Node Capabilities")
    return &csi.NodeGetCapabilitiesResponse{}, nil
}

// Mounts to a common directory for pods to bind mount to
func (s *service) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
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

    log.Infof("Mounting Device: [%s], Path: (%s)", device, stagingTargetPath)

    /* Check if path exists */
    notPath, err := mounter.IsNotExist(stagingTargetPath)
    if notPath {
        /* Path doesn't exist so we are creating it and assuming nothing is mounted to it */
        if err = mounter.MakeDir(stagingTargetPath); err != nil {
            log.Error(err)
            return nil, status.Error(codes.Internal, err.Error())
        }
    }

    log.Infof("Formating and/or mounting..")
    opts := ""
    if err := mounter.FormatAndMount(device, stagingTargetPath, params.FSType, opts); err != nil {
        log.Error(err)
        return nil, status.Error(
            codes.Internal,
            fmt.Sprintf("Error when mounting Device: %s, Path: %s, FSType: %s, Error: %v", device, stagingTargetPath, params.FSType, err),
        )
    }

    return &csi.NodeStageVolumeResponse{}, nil
}

func (s *service) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
    stagingTargetPath := req.GetStagingTargetPath()

    if len(stagingTargetPath) == 0 {
        return nil, status.Error(codes.InvalidArgument, "NodePublishVolume Staging Target Path must be provided")
    }

    log.Infof("Unmounting: (%s)", stagingTargetPath)
    if err := mounter.Unmount(stagingTargetPath); err != nil {
        log.Error(err)
    }

    return &csi.NodeUnstageVolumeResponse{}, nil
}

// Bind Mounts from staging to pod mount path
func (s *service) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
    stagingTargetPath := req.GetStagingTargetPath()
    targetPath := req.GetTargetPath()
    params, err := s.ParseParams(req.GetVolumeContext())
    if err != nil {
        return nil, status.Error(codes.InvalidArgument, err.Error())
    }

    if len(stagingTargetPath) == 0 {
        return nil, status.Error(codes.InvalidArgument, "NodePublishVolume Staging Target Path must be provided")
    }

    if len(targetPath) == 0 {
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
    if err = mounter.BindMount(stagingTargetPath, targetPath, params.FSType, opts); err != nil {
        log.Error(err)
        return nil, status.Error(codes.Internal, err.Error())
    }

    log.Infof("Mounted: %s", targetPath)
    return nil, status.Error(codes.NotFound, "")
}

func (s *service) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
    targetPath := req.GetTargetPath()
    if len(targetPath) == 0 {
        return nil, status.Error(codes.InvalidArgument, "NodeUnpublishVolume Target Path must be provided")
    }

    log.Infof("Unmounting: %s", targetPath)
    if err := mounter.Unmount(targetPath); err != nil {
        log.Error(err)
    }

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
