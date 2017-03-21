package isdocker

import (
	"testing"
)

func TestIsDocker(t *testing.T) {
	isdocker := IsDocker()
	if isdocker != false {
		t.Errorf("In non-docker mode, but shows docker mode !!!")
	}
}
