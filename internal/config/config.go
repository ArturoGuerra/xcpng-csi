package config

import (
    "os"
    "github.com/arturoguerra/xcpng-csi/internal/structs"
)

func Load() *structs.Config {
    return &structs.Config{
        Username: os.Getenv("XCPNG_USERNAME"),
        Password: os.Getenv("XCPNG_PASSWORD"),
        Host:     os.Getenv("XCPNG_HOST"),
        Zone:     os.Getenv("XCPNG_ZONE"),
        NodeID:   os.Getenv("NODE_ID"),
    }
}
