package srt

import "os"

// check .srt file is valid
func ValidSrtFile(f string) bool {
	_, err := os.Stat(f)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
