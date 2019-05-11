package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type OpenstackProviderSpec struct {
	metav1.TypeMeta		`json:",inline"`
	metav1.ObjectMeta	`json:"metadata,omitempty"`
	CloudsSecret		*corev1.SecretReference	`json:"cloudsSecret"`
	CloudName			string					`json:"cloudName"`
	Flavor				string					`json:"flavor"`
	Image				string					`json:"image"`
	KeyName				string					`json:"keyName,omitempty"`
	SshUserName			string					`json:"sshUserName,omitempty"`
	Networks			[]NetworkParam			`json:"networks,omitempty"`
	FloatingIP			string					`json:"floatingIP,omitempty"`
	AvailabilityZone	string					`json:"availabilityZone,omitempty"`
	SecurityGroups		[]SecurityGroupParam	`json:"securityGroups,omitempty"`
	UserDataSecret		*corev1.SecretReference	`json:"userDataSecret,omitempty"`
	Trunk				bool					`json:"trunk,omitempty"`
	Tags				[]string				`json:"tags,omitempty"`
	ServerMetadata		map[string]string		`json:"serverMetadata,omitempty"`
	ConfigDrive			*bool					`json:"configDrive,omitempty"`
	RootVolume			*RootVolume				`json:"rootVolume,omitempty"`
}
type SecurityGroupParam struct {
	UUID	string				`json:"uuid,omitempty"`
	Name	string				`json:"name,omitempty"`
	Filter	SecurityGroupFilter	`json:"filter,omitempty"`
}
type SecurityGroupFilter struct {
	ID			string	`json:"id,omitempty"`
	Name		string	`json:"name,omitempty"`
	Description	string	`json:"description,omitempty"`
	TenantID	string	`json:"tenantId,omitempty"`
	ProjectID	string	`json:"projectId,omitempty"`
	Limit		int		`json:"limit,omitempty"`
	Marker		string	`json:"marker,omitempty"`
	SortKey		string	`json:"sortKey,omitempty"`
	SortDir		string	`json:"sortDir,omitempty"`
	Tags		string	`json:"tags,omitempty"`
	TagsAny		string	`json:"tagsAny,omitempty"`
	NotTags		string	`json:"notTags,omitempty"`
	NotTagsAny	string	`json:"notTagsAny,omitempty"`
}
type NetworkParam struct {
	UUID	string			`json:"uuid,omitempty"`
	FixedIp	string			`json:"fixedIp,omitempty"`
	Filter	Filter			`json:"filter,omitempty"`
	Subnets	[]SubnetParam	`json:"subnets,omitempty"`
}
type Filter struct {
	Status			string	`json:"status,omitempty"`
	Name			string	`json:"name,omitempty"`
	Description		string	`json:"description,omitempty"`
	AdminStateUp	*bool	`json:"adminStateUp,omitempty"`
	TenantID		string	`json:"tenantId,omitempty"`
	ProjectID		string	`json:"projectId,omitempty"`
	Shared			*bool	`json:"shared,omitempty"`
	ID				string	`json:"id,omitempty"`
	Marker			string	`json:"marker,omitempty"`
	Limit			int		`json:"limit,omitempty"`
	SortKey			string	`json:"sortKey,omitempty"`
	SortDir			string	`json:"sortDir,omitempty"`
	Tags			string	`json:"tags,omitempty"`
	TagsAny			string	`json:"tagsAny,omitempty"`
	NotTags			string	`json:"notTags,omitempty"`
	NotTagsAny		string	`json:"notTagsAny,omitempty"`
}
type SubnetParam struct {
	UUID	string			`json:"uuid,omitempty"`
	Filter	SubnetFilter	`json:"filter,omitempty"`
}
type SubnetFilter struct {
	Name			string	`json:"name,omitempty"`
	Description		string	`json:"description,omitempty"`
	EnableDHCP		*bool	`json:"enableDhcp,omitempty"`
	NetworkID		string	`json:"networkId,omitempty"`
	TenantID		string	`json:"tenantId,omitempty"`
	ProjectID		string	`json:"projectId,omitempty"`
	IPVersion		int		`json:"ipVersion,omitempty"`
	GatewayIP		string	`json:"gateway_ip,omitempty"`
	CIDR			string	`json:"cidr,omitempty"`
	IPv6AddressMode	string	`json:"ipv6AddressMode,omitempty"`
	IPv6RAMode		string	`json:"ipv6RaMode,omitempty"`
	ID				string	`json:"id,omitempty"`
	SubnetPoolID	string	`json:"subnetpoolId,omitempty"`
	Limit			int		`json:"limit,omitempty"`
	Marker			string	`json:"marker,omitempty"`
	SortKey			string	`json:"sortKey,omitempty"`
	SortDir			string	`json:"sortDir,omitempty"`
	Tags			string	`json:"tags,omitempty"`
	TagsAny			string	`json:"tagsAny,omitempty"`
	NotTags			string	`json:"notTags,omitempty"`
	NotTagsAny		string	`json:"notTagsAny,omitempty"`
}
type RootVolume struct {
	SourceType	string	`json:"sourceType,omitempty"`
	SourceUUID	string	`json:"sourceUUID,omitempty"`
	DeviceType	string	`json:"deviceType"`
	Size		int		`json:"diskSize,omitempty"`
}
type OpenstackClusterProviderSpec struct {
	metav1.TypeMeta			`json:",inline"`
	metav1.ObjectMeta		`json:"metadata,omitempty"`
	NodeCIDR				string		`json:"nodeCidr,omitempty"`
	DNSNameservers			[]string	`json:"dnsNameservers,omitempty"`
	ExternalNetworkID		string		`json:"externalNetworkId,omitempty"`
	ManagedSecurityGroups	bool		`json:"managedSecurityGroups"`
	Tags					[]string	`json:"tags,omitempty"`
	DisableServerTags		bool		`json:"disableServerTags,omitempty"`
}
type OpenstackClusterProviderStatus struct {
	metav1.TypeMeta				`json:",inline"`
	metav1.ObjectMeta			`json:"metadata,omitempty"`
	CACertificate				[]byte
	CAPrivateKey				[]byte
	Network						*Network		`json:"network,omitempty"`
	ControlPlaneSecurityGroup	*SecurityGroup	`json:"controlPlaneSecurityGroup,omitempty"`
	GlobalSecurityGroup			*SecurityGroup	`json:"globalSecurityGroup,omitempty"`
}
type Network struct {
	Name	string	`json:"name"`
	ID		string	`json:"id"`
	Subnet	*Subnet	`json:"subnet,omitempty"`
	Router	*Router	`json:"router,omitempty"`
}
type Subnet struct {
	Name	string	`json:"name"`
	ID		string	`json:"id"`
	CIDR	string	`json:"cidr"`
}
type Router struct {
	Name	string	`json:"name"`
	ID		string	`json:"id"`
}

func init() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	SchemeBuilder.Register(&OpenstackProviderSpec{})
	SchemeBuilder.Register(&OpenstackClusterProviderSpec{})
	SchemeBuilder.Register(&OpenstackClusterProviderStatus{})
}
