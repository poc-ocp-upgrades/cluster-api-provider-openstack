package v1alpha1

type SecurityGroup struct {
	Name	string			`json:"name"`
	ID	string			`json:"id"`
	Rules	[]SecurityGroupRule	`json:"rules"`
}
type SecurityGroupRule struct {
	ID		string	`json:"name"`
	Direction	string	`json:"direction"`
	EtherType	string	`json:"etherType"`
	SecurityGroupID	string	`json:"securityGroupID"`
	PortRangeMin	int	`json:"portRangeMin"`
	PortRangeMax	int	`json:"portRangeMax"`
	Protocol	string	`json:"protocol"`
	RemoteGroupID	string	`json:"remoteGroupID"`
	RemoteIPPrefix	string	`json:"remoteIPPrefix"`
}

func (r SecurityGroupRule) Equal(x SecurityGroupRule) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	return (r.Direction == x.Direction && r.EtherType == x.EtherType && r.PortRangeMin == x.PortRangeMin && r.PortRangeMax == x.PortRangeMax && r.Protocol == x.Protocol && r.RemoteGroupID == x.RemoteGroupID && r.RemoteIPPrefix == x.RemoteIPPrefix)
}
