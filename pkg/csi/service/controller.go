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
	"math/rand"
	"time"

	"github.com/arturoguerra/xcpng-csi/pkg/errs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	/* Locks to ensure we don't get double volume creating something that has happened */
	s.CVMux.Lock()
	defer s.CVMux.Unlock()
	log.Info("Running CreateVolume")
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

	// zoneRequirements.Requisite []*Topology
	// zoneRequirements.Preferred []*Topology
	// Topology { Segments: map[string]string }

	zoneRequirements := req.GetAccessibilityRequirements()

	if zoneRequirements != nil {
		var topologies *csi.Topology

		rand.Seed(time.Now().UnixNano())

		if len(zoneRequirements.Preferred) > 0 {
			idx := rand.Intn(len(zoneRequirements.Requisite) - 0)
			topologies = zoneRequirements.Preferred[idx]
		} else if len(zoneRequirements.Requisite) > 0 {
			idx := rand.Intn(len(zoneRequirements.Requisite) - 0)
			topologies = zoneRequirements.Requisite[idx]
		} else {
			log.Error("Missing topology information")
			return nil, status.Error(codes.Unknown, "Missing topology information")
		}

		zoneName, err := s.GetTopologyFromLabels(topologies.Segments)
		if err != nil {
			log.Error(err)
			return nil, status.Error(codes.Unknown, "")
		}

		zone := s.XClient.GetZoneByLabel(zoneName)
		if zone == nil {
			log.Error("Invalid topology passed")
			return nil, status.Error(codes.InvalidArgument, "Invalid Topology passed")
		}

		VolumeID, err := s.XClient.CreateVolume(name, params.FSType, params.Datastore, volSizeBytes, zone)
		if err != nil {
			log.Error(err)
			return nil, status.Error(codes.Internal, err.Error())
		}

		parameters := map[string]string{}
		parameters["FSType"] = params.FSType

		resp := &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				VolumeId:           string(*VolumeID),
				CapacityBytes:      volSizeBytes,
				VolumeContext:      parameters,
				AccessibleTopology: []*csi.Topology{topologies},
			},
		}

		return resp, nil
	}

	return nil, status.Error(codes.Unknown, "Zone are required")

}

// TODO
func (s *service) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("Running DeleteVolume %s", req.GetVolumeId())
	if err := s.XClient.DeleteVolume(req.GetVolumeId()); err != nil {
		log.Error(err)
		return nil, status.Error(codes.Internal, "")
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (s *service) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	s.PVMux.Lock()
	defer s.PVMux.Unlock()
	log.Info("Running ControllerPublishVolume")
	params, err := s.ParseParams(req.GetVolumeContext())
	if err != nil {
		log.Error(err)
		return nil, status.Error(codes.InvalidArgument, "")
	}

	device, err := s.XClient.Attach(req.GetVolumeId(), req.GetNodeId(), params.FSType)
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

	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			"device": device,
		},
	}, nil
}

func (s *service) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Info("Running ControllerUnpublishVolume")
	if err := s.XClient.Detach(req.GetVolumeId(), req.GetNodeId()); err != nil {
		log.Error(err)
		/*
		   Temp fix for an issue where kubernetes calls this twice causing the pv to stay in Terminating
		   TODO: Implement error filtering for when a volume is not found
		*/
		return &csi.ControllerUnpublishVolumeResponse{}, nil
		/*return nil, status.Error(codes.NotFound, "")*/
	}

	log.Infof("Volume %s detached from node %s", req.GetVolumeId(), nodeName)

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (s *service) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	log.Info("Validating VolumeCapabilities")
	return &csi.ValidateVolumeCapabilitiesResponse{
		Confirmed: &csi.ValidateVolumeCapabilitiesResponse_Confirmed{
			VolumeContext:      req.GetVolumeContext(),
			VolumeCapabilities: req.GetVolumeCapabilities(),
			Parameters:         req.GetParameters(),
		},
	}, nil
}

// Unimplemented

func (s *service) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	log.Info("Running ListVolumes")
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
