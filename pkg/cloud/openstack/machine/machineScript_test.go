package machine

import (
	"testing"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"sigs.k8s.io/yaml"
)

const providerSpecYAML = `
value:
  apiVersion: "openstackproviderconfig/v1alpha1"
  kind: "OpenstackProviderSpec"
`

func TestNodeStartupScriptEmpty(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := &machinev1.Cluster{}
	machine := &machinev1.Machine{}
	err := yaml.Unmarshal([]byte(providerSpecYAML), &machine.Spec.ProviderSpec)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	token := ""
	script_template := ""
	script, err := nodeStartupScript(cluster, machine, token, script_template)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if script != "" {
		t.Errorf("Expected script, found %q instead", script)
	}
}
func TestNodeStartupScriptEndpointError(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := &machinev1.Cluster{}
	machine := &machinev1.Machine{}
	err := yaml.Unmarshal([]byte(providerSpecYAML), &machine.Spec.ProviderSpec)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	token := ""
	script_template := "{{ call .GetMasterEndpoint }}"
	script, err := nodeStartupScript(cluster, machine, token, script_template)
	if err == nil {
		t.Errorf("Expected GetMasterEndpoint to fail, but it succeeded. Startup script %q", script)
	}
}
func TestNodeStartupScriptWithEndpoint(t *testing.T) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	cluster := machinev1.Cluster{}
	cluster.Status.APIEndpoints = make([]machinev1.APIEndpoint, 1)
	cluster.Status.APIEndpoints[0] = machinev1.APIEndpoint{Host: "example.com", Port: 8080}
	machine := &machinev1.Machine{}
	err := yaml.Unmarshal([]byte(providerSpecYAML), &machine.Spec.ProviderSpec)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	token := ""
	script_template := "{{ call .GetMasterEndpoint }}"
	script, err := nodeStartupScript(&cluster, machine, token, script_template)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	expected := "example.com:8080"
	if script != expected {
		t.Errorf("Expected %q master endpoint, found %q instead", expected, script)
	}
}
