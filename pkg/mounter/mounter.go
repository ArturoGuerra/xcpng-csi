package mounter

/*
Was going to use k8s.io/pkg/util/mount but it's missing in the 1.17 release don't know if that is intended,
Hopefully I understood mount propagation correctly lul
*/

import (
	ctx "context"
	"github.com/arturoguerra/go-logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"

	"github.com/akutz/gofsutil"
	// Filesystem stuff
	"os"
)

var (
	log = logging.New()
)

/* Mounts directory NOTE: Is used inside FormatAndMount */
func BindMount(source, target, fstype string, opts string) error {
	return gofsutil.BindMount(ctx.Background(), source, target, fstype, opts)
}

/* Formats and/or Mounts a device to directory */
func FormatAndMount(source, target, fstype string, opts string) error {
	return gofsutil.FormatAndMount(ctx.Background(), source, target, fstype, opts)
}

/* Unmounts directory */
func Unmount(target string) error {
	return gofsutil.Unmount(ctx.Background(), target)
}

/* Check if directory exists */
func IsNotExist(target string) (bool, error) {
	_, err := os.Stat(target)
	if os.IsNotExist(err) {
		log.Infof("Target: (%s) does not exist", target)
		return true, nil
	} else {
		log.Infof("Target: (%s) exists", target)
		return false, err
	}
}

func MakeDir(target string) error {
	// make sure that all paths exist
	return os.MkdirAll(target, 0775)
}

// cleanly unmount path
func CleanUnmount(target string) error {
	// check that target exists (because mount cannot exist without path)
	notPath, err := IsNotExist(target)

	// if something went wrong when checking, fail
	if err != nil {
		log.Error(err)
		return status.Error(codes.Internal, err.Error())
	}

	// if we don't have path, we don't have mount, so we skip unmount
	if notPath {
		log.Info("Nothing to unmount")
		return nil
	}

	// path exists, let's unmount!
	log.Infof("Unmounting Path: %s", target)
	if err := Unmount(target); err != nil {
		// just ignore not mounted errors, otherwise fail
		if !strings.Contains(err.Error(), " not mounted") {
			log.Error(err)
			return status.Error(codes.Internal, err.Error())
		}
	}

	// remove folder, after successful unmount, so we don't try to unmount again
	log.Infof("Removing dir %s", target)
	if err := os.Remove(target); err != nil {
		log.Error(err)
		// if we can't remove folder, something went wrong
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}
