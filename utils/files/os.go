package files

import (
	"os"
	"time"
)


type OsUtils interface {
    getHomeDir() (string, error)
    getDateStr() string
}

type osUtils struct {
}

func newOsUtil() OsUtils {
    return osUtils{}
}

// getHomeDir return the home directory path of the system
func (o osUtils) getHomeDir() (string, error) {
    return os.UserHomeDir()
}

// getDateStr() return the time in a string format that can be used to create the file 
// journal. eg: MM-DD-YYYY
func (o osUtils) getDateStr() string {
    return time.Now().Format("01-02-2006")
}
