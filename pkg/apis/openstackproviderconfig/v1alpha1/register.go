package v1alpha1

import (
	"errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"sigs.k8s.io/yaml"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

const GroupName = "openstackproviderconfig"

var (
	SchemeGroupVersion	= schema.GroupVersion{Group: "openstackproviderconfig.k8s.io", Version: "v1alpha1"}
	SchemeBuilder		= &scheme.Builder{GroupVersion: SchemeGroupVersion}
)

func ClusterSpecFromProviderSpec(providerSpec clusterv1.ProviderSpec) (*OpenstackClusterProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}
	var config OpenstackClusterProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
func ClusterStatusFromProviderStatus(extension *runtime.RawExtension) (*OpenstackClusterProviderStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if extension == nil {
		return &OpenstackClusterProviderStatus{}, nil
	}
	status := new(OpenstackClusterProviderStatus)
	if err := yaml.Unmarshal(extension.Raw, status); err != nil {
		return nil, err
	}
	return status, nil
}
func MachineSpecFromProviderSpec(providerSpec machinev1.ProviderSpec) (*OpenstackProviderSpec, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if providerSpec.Value == nil {
		return nil, errors.New("no such providerSpec found in manifest")
	}
	var config OpenstackProviderSpec
	if err := yaml.Unmarshal(providerSpec.Value.Raw, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
func EncodeClusterStatus(status *OpenstackClusterProviderStatus) (*runtime.RawExtension, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if status == nil {
		return &runtime.RawExtension{}, nil
	}
	var rawBytes []byte
	var err error
	if rawBytes, err = json.Marshal(status); err != nil {
		return nil, err
	}
	return &runtime.RawExtension{Raw: rawBytes}, nil
}
