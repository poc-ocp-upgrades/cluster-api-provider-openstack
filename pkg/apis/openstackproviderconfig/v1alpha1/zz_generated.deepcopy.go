package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Filter) DeepCopyInto(out *Filter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.AdminStateUp != nil {
		in, out := &in.AdminStateUp, &out.AdminStateUp
		*out = new(bool)
		**out = **in
	}
	if in.Shared != nil {
		in, out := &in.Shared, &out.Shared
		*out = new(bool)
		**out = **in
	}
	return
}
func (in *Filter) DeepCopy() *Filter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Filter)
	in.DeepCopyInto(out)
	return out
}
func (in *Network) DeepCopyInto(out *Network) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(Subnet)
		**out = **in
	}
	if in.Router != nil {
		in, out := &in.Router, &out.Router
		*out = new(Router)
		**out = **in
	}
	return
}
func (in *Network) DeepCopy() *Network {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}
func (in *NetworkParam) DeepCopyInto(out *NetworkParam) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Subnets
		*out = make([]SubnetParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *NetworkParam) DeepCopy() *NetworkParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(NetworkParam)
	in.DeepCopyInto(out)
	return out
}
func (in *OpenstackClusterProviderSpec) DeepCopyInto(out *OpenstackClusterProviderSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.DNSNameservers != nil {
		in, out := &in.DNSNameservers, &out.DNSNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *OpenstackClusterProviderSpec) DeepCopy() *OpenstackClusterProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OpenstackClusterProviderSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *OpenstackClusterProviderSpec) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OpenstackClusterProviderStatus) DeepCopyInto(out *OpenstackClusterProviderStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAPrivateKey != nil {
		in, out := &in.CAPrivateKey, &out.CAPrivateKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network)
		(*in).DeepCopyInto(*out)
	}
	if in.ControlPlaneSecurityGroup != nil {
		in, out := &in.ControlPlaneSecurityGroup, &out.ControlPlaneSecurityGroup
		*out = new(SecurityGroup)
		(*in).DeepCopyInto(*out)
	}
	if in.GlobalSecurityGroup != nil {
		in, out := &in.GlobalSecurityGroup, &out.GlobalSecurityGroup
		*out = new(SecurityGroup)
		(*in).DeepCopyInto(*out)
	}
	return
}
func (in *OpenstackClusterProviderStatus) DeepCopy() *OpenstackClusterProviderStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OpenstackClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}
func (in *OpenstackClusterProviderStatus) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *OpenstackProviderSpec) DeepCopyInto(out *OpenstackProviderSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.CloudsSecret != nil {
		in, out := &in.CloudsSecret, &out.CloudsSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkParam, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]SecurityGroupParam, len(*in))
		copy(*out, *in)
	}
	if in.UserDataSecret != nil {
		in, out := &in.UserDataSecret, &out.UserDataSecret
		*out = new(v1.SecretReference)
		**out = **in
	}
	out.RootVolume = in.RootVolume
	out.Tags = in.Tags
	if in.ServerMetadata != nil {
		in, out := &in.ServerMetadata, &out.ServerMetadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}
func (in *OpenstackProviderSpec) DeepCopy() *OpenstackProviderSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(OpenstackProviderSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *OpenstackProviderSpec) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *RootVolume) DeepCopyInto(out *RootVolume) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *RootVolume) DeepCopy() *RootVolume {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(RootVolume)
	in.DeepCopyInto(out)
	return out
}
func (in *Router) DeepCopyInto(out *Router) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *Router) DeepCopy() *Router {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Router)
	in.DeepCopyInto(out)
	return out
}
func (in *SecurityGroup) DeepCopyInto(out *SecurityGroup) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]SecurityGroupRule, len(*in))
		copy(*out, *in)
	}
	return
}
func (in *SecurityGroup) DeepCopy() *SecurityGroup {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecurityGroup)
	in.DeepCopyInto(out)
	return out
}
func (in *SecurityGroupFilter) DeepCopyInto(out *SecurityGroupFilter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *SecurityGroupFilter) DeepCopy() *SecurityGroupFilter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecurityGroupFilter)
	in.DeepCopyInto(out)
	return out
}
func (in *SecurityGroupParam) DeepCopyInto(out *SecurityGroupParam) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.Filter = in.Filter
	return
}
func (in *SecurityGroupParam) DeepCopy() *SecurityGroupParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecurityGroupParam)
	in.DeepCopyInto(out)
	return out
}
func (in *SecurityGroupRule) DeepCopyInto(out *SecurityGroupRule) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *SecurityGroupRule) DeepCopy() *SecurityGroupRule {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SecurityGroupRule)
	in.DeepCopyInto(out)
	return out
}
func (in *Subnet) DeepCopyInto(out *Subnet) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *Subnet) DeepCopy() *Subnet {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Subnet)
	in.DeepCopyInto(out)
	return out
}
func (in *SubnetFilter) DeepCopyInto(out *SubnetFilter) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.EnableDHCP != nil {
		in, out := &in.EnableDHCP, &out.EnableDHCP
		*out = new(bool)
		**out = **in
	}
	return
}
func (in *SubnetFilter) DeepCopy() *SubnetFilter {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SubnetFilter)
	in.DeepCopyInto(out)
	return out
}
func (in *SubnetParam) DeepCopyInto(out *SubnetParam) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	in.Filter.DeepCopyInto(&out.Filter)
	return
}
func (in *SubnetParam) DeepCopy() *SubnetParam {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(SubnetParam)
	in.DeepCopyInto(out)
	return out
}
