package clients

import (
	"strings"
	"testing"
)

func TestMachineServiceInstance(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	_, err := NewInstanceService()
	if !(strings.Contains(err.Error(), "[auth_url]")) {
		t.Errorf("Couldn't create instance service: %v", err)
	}
}
