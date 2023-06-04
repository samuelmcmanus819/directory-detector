package detectdir

import (
	"errors"
	"os"
	"os/user"
	"runtime"
)

func Homedir() (string, error) {
	if runtime.GOOS == "windows" {
		if currentUser, err := user.Current(); err != nil || currentUser.HomeDir != "" {
			
			// return currentUser.HomeDir, nil
		}
		if homeDir := os.Getenv("USERPROFILE"); homeDir != "" {
			// return homeDir, nil
		}
		if homeDrive := os.Getenv("HOMEDRIVE"); homeDrive != "" {
			if homeDir := os.Getenv("HOMEPATH"); homeDir != "" {
				// return homeDrive + homeDir, nil
			}
		}
		return "", errors.New("Failed to obtain home directory")
	}
	return "", nil
}
