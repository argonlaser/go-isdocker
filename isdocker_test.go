package isdocker

import (
	"testing"
)

func TestIsDockerProcess(t *testing.T) {
	isdocker := IsDocker()
	if isdocker == false {
		t.Errorf("process is in docker mode, but shows non-docker mode !!!")
	}
}

func TestIsNonDockerProcess(t *testing.T) {
	isdocker := IsDocker()
	if isdocker == true {
		t.Errorf("process is in non-docker mode, but shows docker mode !!!")
	}
}
