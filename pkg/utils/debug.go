package utils

import (
    "fmt"
    "os"
)

const debugLogFile = "/tmp/xenserver-driver.log"

func Debug(message string) {
    if _, err := os.Stat(debugLogFile); err == nil {
        f, _ := os.OpenFile(debugLogFile, os.O_APPEND|os.O_WRONLY, 0600)
        defer f.Close()
        f.WriteString(fmt.Sprintln(message))
    }
}
