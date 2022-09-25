package mounter

/*
Was going to use k8s.io/pkg/util/mount but it's missing in the 1.17 release don't know if that is intended,
Hopefully I understood mount propagation correctly lul
*/

import (
    ctx "context"
    "github.com/arturoguerra/go-logging"

    // Filesystem stuff
    "os"
    "github.com/akutz/gofsutil"
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
