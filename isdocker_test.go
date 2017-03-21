package isdocker

import (
	"testing"
)

func TestIsDocker(t *testing.T) {
	isdocker := IsDocker()
	if isdocker == false {
		t.Errorf("process is in docker mode, but shows non-docker mode !!!")
	}
}
