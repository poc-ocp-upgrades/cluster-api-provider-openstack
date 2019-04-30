package openstack

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/record"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ActuatorParams struct {
	KubeClient	kubernetes.Interface
	Client		client.Client
	EventRecorder	record.EventRecorder
	Scheme		*runtime.Scheme
}
