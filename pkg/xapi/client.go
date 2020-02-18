package xapi

import (
)

type (
    XClient interface {
        Attach(string, string, string, string) (string, error)
        Detach(string, string) error
        IsAttached(string, string) (bool, error)
        CreateVolume(string, string, string, int) (string, error)
        DeleteVolume(string) error
    }

    xClient struct {
       Username string
       Password string
       Host     string
    }
)

func New(username string, password string, host string) XClient {
    return &xClient{
        username,
        password,
        host,
    }
}
