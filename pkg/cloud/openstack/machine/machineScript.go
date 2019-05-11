package machine

import (
	"bytes"
	"errors"
	"text/template"
	"fmt"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	openstackconfigv1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
)

type setupParams struct {
	Token				string
	Cluster				*machinev1.Cluster
	Machine				*machinev1.Machine
	MachineSpec			*openstackconfigv1.OpenstackProviderSpec
	PodCIDR				string
	ServiceCIDR			string
	GetMasterEndpoint	func() (string, error)
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
}
func masterStartupScript(cluster *machinev1.Cluster, machine *machinev1.Machine, script string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	machineSpec, err := openstackconfigv1.MachineSpecFromProviderSpec(machine.Spec.ProviderSpec)
	if err != nil {
		return "", err
	}
	params := setupParams{Cluster: cluster, Machine: machine, MachineSpec: machineSpec}
	if cluster != nil {
		params.PodCIDR = getSubnet(cluster.Spec.ClusterNetwork.Pods)
		params.ServiceCIDR = getSubnet(cluster.Spec.ClusterNetwork.Services)
	}
	masterStartUpScript := template.Must(template.New("masterStartUp").Parse(script))
	var buf bytes.Buffer
	if err := masterStartUpScript.Execute(&buf, params); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func nodeStartupScript(cluster *machinev1.Cluster, machine *machinev1.Machine, token, script string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	machineSpec, err := openstackconfigv1.MachineSpecFromProviderSpec(machine.Spec.ProviderSpec)
	if err != nil {
		return "", err
	}
	GetMasterEndpoint := func() (string, error) {
		if cluster == nil {
			return "", nil
		} else if len(cluster.Status.APIEndpoints) == 0 {
			return "", errors.New("no cluster status found")
		}
		return getEndpoint(cluster.Status.APIEndpoints[0]), nil
	}
	params := setupParams{Token: token, Cluster: cluster, Machine: machine, MachineSpec: machineSpec, GetMasterEndpoint: GetMasterEndpoint}
	if cluster != nil {
		params.PodCIDR = getSubnet(cluster.Spec.ClusterNetwork.Pods)
		params.ServiceCIDR = getSubnet(cluster.Spec.ClusterNetwork.Services)
	}
	nodeStartUpScript := template.Must(template.New("nodeStartUp").Parse(script))
	var buf bytes.Buffer
	if err := nodeStartUpScript.Execute(&buf, params); err != nil {
		return "", err
	}
	return buf.String(), nil
}
func getEndpoint(apiEndpoint machinev1.APIEndpoint) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fmt.Sprintf("%s:%d", apiEndpoint.Host, apiEndpoint.Port)
}
func getSubnet(netRange machinev1.NetworkRanges) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if len(netRange.CIDRBlocks) == 0 {
		return ""
	}
	return netRange.CIDRBlocks[0]
}
