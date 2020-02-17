package utils

import (
    "os"
    "io/ioutil"
    "encoding/json"
)

const fileName = "/opt/xenserver/config.json"

type Config struct {
    Master bool     `json:"master"`
    Username string `json:"username"`
    Password string `json:"password"`
    Host     string `json:"host"`
}

func Init() *Config {
    jsonFile, err := os.Open(fileName)
    if err != nil {
        return &Config{
            Master: false,
        }
    }

    byteValues, _ := ioutil.ReadAll(jsonFile)
    var cfg Config
    json.Unmarshal(byteValues, &cfg)
    return &cfg
}
