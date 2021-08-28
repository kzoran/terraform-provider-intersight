/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document.

API version: 1.0.9-4437
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// ApplianceDeviceClaim DeviceClaim managed object represents a user initiated claim request for claiming an endpoint device. There can be many DeviceClaim managed object for a given endpoint device when users claim and unclaim devices repeatedly. Claiming an endpoint device is a multi-step operation. The Intersight Appliance starts a workflow with multiple tasks to process the device claim request. The status of the device claim operation can be obtained from the claim workflow.
type ApplianceDeviceClaim struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Device identifier of the endpoint device.
	DeviceId *string `json:"DeviceId,omitempty"`
	// Hostname or IP address of the endpoint device the user wants to claim.
	Hostname *string `json:"Hostname,omitempty"`
	// Indicates whether the value of the 'password' property has been set.
	IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
	// Tracks if this device is to be claimed or certificate renewal.
	IsRenew *bool `json:"IsRenew,omitempty"`
	// Message set by the device claim process.
	Message *string `json:"Message,omitempty"`
	// Password to be used to login to the endpoint device.
	Password *string `json:"Password,omitempty"`
	// Platform type of the endpoint device. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `CiscoCatalyst` - A Cisco Catalyst networking switch device.
	PlatformType *string `json:"PlatformType,omitempty"`
	// User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set.
	RequestId *string `json:"RequestId,omitempty"`
	// Device security token of the endpoint device.
	SecurityToken *string `json:"SecurityToken,omitempty"`
	// Status of the device claim process. * `started` - Device claim operation has started. * `failed` - Device claim operation has failed. * `completed` - Device claim operation has completed.
	Status *string `json:"Status,omitempty"`
	// Username to log in to the endpoint device.
	Username             *string                 `json:"Username,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ApplianceDeviceClaim ApplianceDeviceClaim

// NewApplianceDeviceClaim instantiates a new ApplianceDeviceClaim object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplianceDeviceClaim(classId string, objectType string) *ApplianceDeviceClaim {
	this := ApplianceDeviceClaim{}
	this.ClassId = classId
	this.ObjectType = objectType
	var isRenew bool = false
	this.IsRenew = &isRenew
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// NewApplianceDeviceClaimWithDefaults instantiates a new ApplianceDeviceClaim object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplianceDeviceClaimWithDefaults() *ApplianceDeviceClaim {
	this := ApplianceDeviceClaim{}
	var classId string = "appliance.DeviceClaim"
	this.ClassId = classId
	var objectType string = "appliance.DeviceClaim"
	this.ObjectType = objectType
	var isRenew bool = false
	this.IsRenew = &isRenew
	var platformType string = ""
	this.PlatformType = &platformType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ApplianceDeviceClaim) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ApplianceDeviceClaim) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ApplianceDeviceClaim) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ApplianceDeviceClaim) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ApplianceDeviceClaim) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetHostname() string {
	if o == nil || o.Hostname == nil {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetHostnameOk() (*string, bool) {
	if o == nil || o.Hostname == nil {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasHostname() bool {
	if o != nil && o.Hostname != nil {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ApplianceDeviceClaim) SetHostname(v string) {
	o.Hostname = &v
}

// GetIsPasswordSet returns the IsPasswordSet field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetIsPasswordSet() bool {
	if o == nil || o.IsPasswordSet == nil {
		var ret bool
		return ret
	}
	return *o.IsPasswordSet
}

// GetIsPasswordSetOk returns a tuple with the IsPasswordSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetIsPasswordSetOk() (*bool, bool) {
	if o == nil || o.IsPasswordSet == nil {
		return nil, false
	}
	return o.IsPasswordSet, true
}

// HasIsPasswordSet returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasIsPasswordSet() bool {
	if o != nil && o.IsPasswordSet != nil {
		return true
	}

	return false
}

// SetIsPasswordSet gets a reference to the given bool and assigns it to the IsPasswordSet field.
func (o *ApplianceDeviceClaim) SetIsPasswordSet(v bool) {
	o.IsPasswordSet = &v
}

// GetIsRenew returns the IsRenew field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetIsRenew() bool {
	if o == nil || o.IsRenew == nil {
		var ret bool
		return ret
	}
	return *o.IsRenew
}

// GetIsRenewOk returns a tuple with the IsRenew field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetIsRenewOk() (*bool, bool) {
	if o == nil || o.IsRenew == nil {
		return nil, false
	}
	return o.IsRenew, true
}

// HasIsRenew returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasIsRenew() bool {
	if o != nil && o.IsRenew != nil {
		return true
	}

	return false
}

// SetIsRenew gets a reference to the given bool and assigns it to the IsRenew field.
func (o *ApplianceDeviceClaim) SetIsRenew(v bool) {
	o.IsRenew = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApplianceDeviceClaim) SetMessage(v string) {
	o.Message = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ApplianceDeviceClaim) SetPassword(v string) {
	o.Password = &v
}

// GetPlatformType returns the PlatformType field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetPlatformType() string {
	if o == nil || o.PlatformType == nil {
		var ret string
		return ret
	}
	return *o.PlatformType
}

// GetPlatformTypeOk returns a tuple with the PlatformType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetPlatformTypeOk() (*string, bool) {
	if o == nil || o.PlatformType == nil {
		return nil, false
	}
	return o.PlatformType, true
}

// HasPlatformType returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasPlatformType() bool {
	if o != nil && o.PlatformType != nil {
		return true
	}

	return false
}

// SetPlatformType gets a reference to the given string and assigns it to the PlatformType field.
func (o *ApplianceDeviceClaim) SetPlatformType(v string) {
	o.PlatformType = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ApplianceDeviceClaim) SetRequestId(v string) {
	o.RequestId = &v
}

// GetSecurityToken returns the SecurityToken field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetSecurityToken() string {
	if o == nil || o.SecurityToken == nil {
		var ret string
		return ret
	}
	return *o.SecurityToken
}

// GetSecurityTokenOk returns a tuple with the SecurityToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetSecurityTokenOk() (*string, bool) {
	if o == nil || o.SecurityToken == nil {
		return nil, false
	}
	return o.SecurityToken, true
}

// HasSecurityToken returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasSecurityToken() bool {
	if o != nil && o.SecurityToken != nil {
		return true
	}

	return false
}

// SetSecurityToken gets a reference to the given string and assigns it to the SecurityToken field.
func (o *ApplianceDeviceClaim) SetSecurityToken(v string) {
	o.SecurityToken = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApplianceDeviceClaim) SetStatus(v string) {
	o.Status = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ApplianceDeviceClaim) SetUsername(v string) {
	o.Username = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ApplianceDeviceClaim) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplianceDeviceClaim) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ApplianceDeviceClaim) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *ApplianceDeviceClaim) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o ApplianceDeviceClaim) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.Hostname != nil {
		toSerialize["Hostname"] = o.Hostname
	}
	if o.IsPasswordSet != nil {
		toSerialize["IsPasswordSet"] = o.IsPasswordSet
	}
	if o.IsRenew != nil {
		toSerialize["IsRenew"] = o.IsRenew
	}
	if o.Message != nil {
		toSerialize["Message"] = o.Message
	}
	if o.Password != nil {
		toSerialize["Password"] = o.Password
	}
	if o.PlatformType != nil {
		toSerialize["PlatformType"] = o.PlatformType
	}
	if o.RequestId != nil {
		toSerialize["RequestId"] = o.RequestId
	}
	if o.SecurityToken != nil {
		toSerialize["SecurityToken"] = o.SecurityToken
	}
	if o.Status != nil {
		toSerialize["Status"] = o.Status
	}
	if o.Username != nil {
		toSerialize["Username"] = o.Username
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ApplianceDeviceClaim) UnmarshalJSON(bytes []byte) (err error) {
	type ApplianceDeviceClaimWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Device identifier of the endpoint device.
		DeviceId *string `json:"DeviceId,omitempty"`
		// Hostname or IP address of the endpoint device the user wants to claim.
		Hostname *string `json:"Hostname,omitempty"`
		// Indicates whether the value of the 'password' property has been set.
		IsPasswordSet *bool `json:"IsPasswordSet,omitempty"`
		// Tracks if this device is to be claimed or certificate renewal.
		IsRenew *bool `json:"IsRenew,omitempty"`
		// Message set by the device claim process.
		Message *string `json:"Message,omitempty"`
		// Password to be used to login to the endpoint device.
		Password *string `json:"Password,omitempty"`
		// Platform type of the endpoint device. * `` - The device reported an empty or unrecognized platform type. * `APIC` - An Application Policy Infrastructure Controller cluster. * `DCNM` - A Data Center Network Manager instance. Data Center Network Manager (DCNM) is the network management platform for all NX-OS-enabled deployments, spanning new fabric architectures, IP Fabric for Media, and storage networking deployments for the Cisco Nexus-powered data center. * `UCSFI` - A UCS Fabric Interconnect in HA or standalone mode, which is being managed by UCS Manager (UCSM). * `UCSFIISM` - A UCS Fabric Interconnect in HA or standalone mode, managed directly by Intersight. * `IMC` - A standalone UCS Server Integrated Management Controller. * `IMCM4` - A standalone UCS M4 Server. * `IMCM5` - A standalone UCS M5 server. * `IMCRack` - A standalone UCS M6 and above server. * `UCSIOM` - An UCS Chassis IO module. * `HX` - A HyperFlex storage controller. * `HyperFlexAP` - A HyperFlex Application Platform. * `UCSD` - A UCS Director virtual appliance. Cisco UCS Director automates, orchestrates, and manages Cisco and third-party hardware. * `IntersightAppliance` - A Cisco Intersight Connected Virtual Appliance. * `IntersightAssist` - A Cisco Intersight Assist. * `PureStorageFlashArray` - A Pure Storage FlashArray device. * `UCSC890` - A standalone Cisco UCSC890 server. * `NetAppOntap` - A NetApp ONTAP storage system. * `NetAppActiveIqUnifiedManager` - A NetApp Active IQ Unified Manager. * `EmcScaleIo` - An EMC ScaleIO storage system. * `EmcVmax` - An EMC VMAX storage system. * `EmcVplex` - An EMC VPLEX storage system. * `EmcXtremIo` - An EMC XtremIO storage system. * `VmwareVcenter` - A VMware vCenter device that manages Virtual Machines. * `MicrosoftHyperV` - A Microsoft Hyper-V system that manages Virtual Machines. * `AppDynamics` - An AppDynamics controller that monitors applications. * `Dynatrace` - A software-intelligence monitoring platform that simplifies enterprise cloud complexity and accelerates digital transformation. * `ReadHatOpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `CloudFoundry` - An open source cloud platform on which developers can build, deploy, run and scale applications. * `MicrosoftAzureApplicationInsights` - A feature of Azure Monitor, is an extensible Application Performance Management service for developers and DevOps professionals to monitor their live applications. * `OpenStack` - An OpenStack target manages Virtual Machines, Physical Machines, Datacenters and Virtual Datacenters using different OpenStack services as administrative endpoints. * `MicrosoftSqlServer` - A Microsoft SQL database server. * `Kubernetes` - A Kubernetes cluster that runs containerized applications. * `AmazonWebService` - A Amazon web service target that discovers and monitors different services like EC2. It discovers entities like VMs, Volumes, regions etc. and monitors attributes like Mem, CPU, cost. * `AmazonWebServiceBilling` - A Amazon web service billing target to retrieve billing information stored in S3 bucket. * `MicrosoftAzureServicePrincipal` - A Microsoft Azure Service Principal target that discovers all the associated Azure subscriptions. * `MicrosoftAzureEnterpriseAgreement` - A Microsoft Azure Enterprise Agreement target that discovers cost, billing and RIs. * `DellCompellent` - A Dell Compellent storage system. * `HPE3Par` - A HPE 3PAR storage system. * `RedHatEnterpriseVirtualization` - A Red Hat Enterprise Virtualization Hypervisor system that manages Virtual Machines. * `NutanixAcropolis` - A Nutanix Acropolis system that combines servers and storage into a distributed infrastructure platform. * `HPEOneView` - A HPE Oneview management system that manages compute, storage, and networking. * `ServiceEngine` - Cisco Application Services Engine. Cisco Application Services Engine is a platform to deploy and manage applications. * `HitachiVirtualStoragePlatform` - A Hitachi Virtual Storage Platform also referred to as Hitachi VSP. It includes various storage systems designed for data centers. * `IMCBlade` - An Intersight managed UCS Blade Server. * `TerraformCloud` - A Terraform Cloud account. * `TerraformAgent` - A Terraform Cloud Agent that Intersight will deploy in datacenter. The agent will execute Terraform plan for Terraform Cloud workspace configured to use the agent. * `CustomTarget` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `HTTPEndpoint` - An external endpoint added as Target that can be accessed through its HTTP API interface in Intersight Orchestrator automation workflow.Standard HTTP authentication scheme supported: Basic. * `CiscoCatalyst` - A Cisco Catalyst networking switch device.
		PlatformType *string `json:"PlatformType,omitempty"`
		// User defined claim request identifier set by the UI. The RequestId field is not a mandatory. The Intersight Appliance will assign a unique value automatically if the field is not set.
		RequestId *string `json:"RequestId,omitempty"`
		// Device security token of the endpoint device.
		SecurityToken *string `json:"SecurityToken,omitempty"`
		// Status of the device claim process. * `started` - Device claim operation has started. * `failed` - Device claim operation has failed. * `completed` - Device claim operation has completed.
		Status *string `json:"Status,omitempty"`
		// Username to log in to the endpoint device.
		Username *string                 `json:"Username,omitempty"`
		Account  *IamAccountRelationship `json:"Account,omitempty"`
	}

	varApplianceDeviceClaimWithoutEmbeddedStruct := ApplianceDeviceClaimWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varApplianceDeviceClaimWithoutEmbeddedStruct)
	if err == nil {
		varApplianceDeviceClaim := _ApplianceDeviceClaim{}
		varApplianceDeviceClaim.ClassId = varApplianceDeviceClaimWithoutEmbeddedStruct.ClassId
		varApplianceDeviceClaim.ObjectType = varApplianceDeviceClaimWithoutEmbeddedStruct.ObjectType
		varApplianceDeviceClaim.DeviceId = varApplianceDeviceClaimWithoutEmbeddedStruct.DeviceId
		varApplianceDeviceClaim.Hostname = varApplianceDeviceClaimWithoutEmbeddedStruct.Hostname
		varApplianceDeviceClaim.IsPasswordSet = varApplianceDeviceClaimWithoutEmbeddedStruct.IsPasswordSet
		varApplianceDeviceClaim.IsRenew = varApplianceDeviceClaimWithoutEmbeddedStruct.IsRenew
		varApplianceDeviceClaim.Message = varApplianceDeviceClaimWithoutEmbeddedStruct.Message
		varApplianceDeviceClaim.Password = varApplianceDeviceClaimWithoutEmbeddedStruct.Password
		varApplianceDeviceClaim.PlatformType = varApplianceDeviceClaimWithoutEmbeddedStruct.PlatformType
		varApplianceDeviceClaim.RequestId = varApplianceDeviceClaimWithoutEmbeddedStruct.RequestId
		varApplianceDeviceClaim.SecurityToken = varApplianceDeviceClaimWithoutEmbeddedStruct.SecurityToken
		varApplianceDeviceClaim.Status = varApplianceDeviceClaimWithoutEmbeddedStruct.Status
		varApplianceDeviceClaim.Username = varApplianceDeviceClaimWithoutEmbeddedStruct.Username
		varApplianceDeviceClaim.Account = varApplianceDeviceClaimWithoutEmbeddedStruct.Account
		*o = ApplianceDeviceClaim(varApplianceDeviceClaim)
	} else {
		return err
	}

	varApplianceDeviceClaim := _ApplianceDeviceClaim{}

	err = json.Unmarshal(bytes, &varApplianceDeviceClaim)
	if err == nil {
		o.MoBaseMo = varApplianceDeviceClaim.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "Hostname")
		delete(additionalProperties, "IsPasswordSet")
		delete(additionalProperties, "IsRenew")
		delete(additionalProperties, "Message")
		delete(additionalProperties, "Password")
		delete(additionalProperties, "PlatformType")
		delete(additionalProperties, "RequestId")
		delete(additionalProperties, "SecurityToken")
		delete(additionalProperties, "Status")
		delete(additionalProperties, "Username")
		delete(additionalProperties, "Account")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApplianceDeviceClaim struct {
	value *ApplianceDeviceClaim
	isSet bool
}

func (v NullableApplianceDeviceClaim) Get() *ApplianceDeviceClaim {
	return v.value
}

func (v *NullableApplianceDeviceClaim) Set(val *ApplianceDeviceClaim) {
	v.value = val
	v.isSet = true
}

func (v NullableApplianceDeviceClaim) IsSet() bool {
	return v.isSet
}

func (v *NullableApplianceDeviceClaim) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplianceDeviceClaim(val *ApplianceDeviceClaim) *NullableApplianceDeviceClaim {
	return &NullableApplianceDeviceClaim{value: val, isSet: true}
}

func (v NullableApplianceDeviceClaim) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplianceDeviceClaim) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
