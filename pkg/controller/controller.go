package controller

import (
	"k8s.io/klog"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/cloud/openstack"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var AddToManagerFuncs []func(manager.Manager) error

func AddToManager(m manager.Manager) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, f := range AddToManagerFuncs {
		if err := f(m); err != nil {
			return err
		}
	}
	return nil
}
func getActuatorParams(mgr manager.Manager) openstack.ActuatorParams {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	config := mgr.GetConfig()
	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Could not create kubernetes client to talk to the apiserver: %v", err)
	}
	return openstack.ActuatorParams{Client: mgr.GetClient(), KubeClient: kubeClient, Scheme: mgr.GetScheme(), EventRecorder: mgr.GetRecorder("openstack-controller")}
}
