package isdocker

import (
	"io/ioutil"
	"os"
	"strings"
)

func hasDockerEnvFile() bool {
	if _, err := os.Stat("/.dockerenv"); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func hasCGroupDocker() bool {
	bytes, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return false
	}
	cGroupContent := string(bytes)
	return strings.Contains(cGroupContent, "docker")
}

func IsDocker() bool {
	return hasDockerEnvFile() || hasCGroupDocker()
}
