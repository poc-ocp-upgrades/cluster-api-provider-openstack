package machine

import (
	"bytes"
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const InstanceStatusAnnotationKey = "instance-status"

type instanceStatus *machinev1.Machine

func (oc *OpenstackClient) instanceStatus(machine *machinev1.Machine) (instanceStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	currentMachine, err := GetMachineIfExists(oc.client, machine.Namespace, machine.Name)
	if err != nil {
		return nil, err
	}
	if currentMachine == nil {
		return nil, nil
	}
	return oc.machineInstanceStatus(currentMachine)
}
func GetMachineIfExists(c client.Client, namespace, name string) (*machinev1.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c == nil {
		return nil, nil
	}
	machine := &machinev1.Machine{}
	key := client.ObjectKey{Namespace: namespace, Name: name}
	err := c.Get(context.Background(), key, machine)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	return machine, nil
}
func (oc *OpenstackClient) updateInstanceStatus(machine *machinev1.Machine) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status := instanceStatus(machine)
	currentMachine, err := GetMachineIfExists(oc.client, machine.Namespace, machine.Name)
	if err != nil {
		return err
	}
	if currentMachine == nil {
		return fmt.Errorf("Machine has already been deleted. Cannot update current instance status for machine %v", machine.ObjectMeta.Name)
	}
	m, err := oc.setMachineInstanceStatus(currentMachine, status)
	if err != nil {
		return err
	}
	return oc.client.Update(nil, m)
}
func (oc *OpenstackClient) machineInstanceStatus(machine *machinev1.Machine) (instanceStatus, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if machine.ObjectMeta.Annotations == nil {
		return nil, nil
	}
	a := machine.ObjectMeta.Annotations[InstanceStatusAnnotationKey]
	if a == "" {
		return nil, nil
	}
	serializer := json.NewSerializer(json.DefaultMetaFactory, oc.scheme, oc.scheme, false)
	var status machinev1.Machine
	_, _, err := serializer.Decode([]byte(a), &schema.GroupVersionKind{Group: "cluster.k8s.io", Version: "v1alpha1", Kind: "Machine"}, &status)
	if err != nil {
		return nil, fmt.Errorf("decoding failure: %v", err)
	}
	return instanceStatus(&status), nil
}
func (oc *OpenstackClient) setMachineInstanceStatus(machine *machinev1.Machine, status instanceStatus) (*machinev1.Machine, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	status.ObjectMeta.Annotations[InstanceStatusAnnotationKey] = ""
	serializer := json.NewSerializer(json.DefaultMetaFactory, oc.scheme, oc.scheme, false)
	b := []byte{}
	buff := bytes.NewBuffer(b)
	err := serializer.Encode((*machinev1.Machine)(status), buff)
	if err != nil {
		return nil, fmt.Errorf("encoding failure: %v", err)
	}
	if machine.ObjectMeta.Annotations == nil {
		machine.ObjectMeta.Annotations = make(map[string]string)
	}
	machine.ObjectMeta.Annotations[InstanceStatusAnnotationKey] = buff.String()
	return machine, nil
}
