package test

import (
	"github.com/hultan/softteam-tools/pkg/resources"
	"testing"
)

func TestResource(t *testing.T) {
	res := resources.NewResources()
	path := res.GetResourcePath("test.txt")
	if path != "/tmp/test.txt" {
		t.Error("failed to find resource")
	}
}
