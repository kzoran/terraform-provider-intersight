package intersight

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func compareMaps(expected, ffOutput map[string]interface{}, t *testing.T) error {
	for key, value := range expected {
		v, ok := ffOutput[key]
		if !ok {
			return fmt.Errorf("error: key not found %s", key)
		}
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			value := value.([]map[string]interface{})
			v := v.([]map[string]interface{})
			for i := 0; i < len(value); i++ {
				if err := compareMaps(value[i], v[i], t); err != nil {
					return err
				}
			}
		case reflect.Map:
			return compareMaps(value.(map[string]interface{}), v.(map[string]interface{}), t)
		case reflect.Int32, reflect.Int, reflect.Int64:
			if reflect.ValueOf(value).Int() != reflect.ValueOf(v).Int() {
				return fmt.Errorf("error: mismatch in int type: expectedValue %+v kind %T, gotValue %+v kind %T", value, value, v,
					v)
			}
		case reflect.Float64, reflect.Float32:
			if reflect.ValueOf(value).Float() != reflect.ValueOf(v).Float() {
				return fmt.Errorf("error: mismatch in float type: expectedValue %+v kind %T, gotValue %+v kind %T", value, value, v,
					v)
			}
		default:
			if !reflect.DeepEqual(value, v) {
				return fmt.Errorf("error: mismatch in value %+v %+v", v, value)
			}
		}
	}
	return nil
}

//CheckError Panics when error occurs
func CheckError(t *testing.T, e error) {
	if e != nil {
		t.Errorf("Error while testing flatten function: %+v", e)
	}
}

func TestFlattenListAdapterAdapterConfig(t *testing.T) {
	p := []models.AdapterAdapterConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"adapter.AdapterConfig","ObjectType":"adapter.AdapterConfig","SlotId":"SlotId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterAdapterConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterAdapterConfig{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterAdapterConfig(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "adapter.AdapterConfig", "object_type": "adapter.AdapterConfig", "slot_id": "SlotId 1"}, {"class_id": "adapter.AdapterConfig", "object_type": "adapter.AdapterConfig", "slot_id": "SlotId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterExtEthInterfaceRelationship(t *testing.T) {
	p := []models.AdapterExtEthInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterExtEthInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterExtEthInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterExtEthInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostEthInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostEthInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostEthInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostEthInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostEthInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostFcInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostFcInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostFcInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostFcInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostFcInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterHostIscsiInterfaceRelationship(t *testing.T) {
	p := []models.AdapterHostIscsiInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterHostIscsiInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterHostIscsiInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterHostIscsiInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAdapterUnitRelationship(t *testing.T) {
	p := []models.AdapterUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAdapterUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AdapterUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAdapterUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceApiStatus(t *testing.T) {
	p := []models.ApplianceApiStatus{}
	var d = &schema.ResourceData{}
	c := `{"Response":32,"ObjectType":"appliance.ApiStatus","ClassId":"appliance.ApiStatus","ElapsedTime":32.000000,"ObjectName":"ObjectName %d","Reason":"Reason %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceApiStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceApiStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceApiStatus(p, d)
	expectedOp := []map[string]interface{}{{"response": 32, "object_type": "appliance.ApiStatus", "class_id": "appliance.ApiStatus", "elapsed_time": 32.000000, "object_name": "ObjectName 1", "reason": "Reason 1"}, {"response": 32, "object_type": "appliance.ApiStatus", "class_id": "appliance.ApiStatus", "elapsed_time": 32.000000, "object_name": "ObjectName 2", "reason": "Reason 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceAppStatusRelationship(t *testing.T) {
	p := []models.ApplianceAppStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceAppStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceAppStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceAppStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceCertRenewalPhase(t *testing.T) {
	p := []models.ApplianceCertRenewalPhase{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"appliance.CertRenewalPhase","ClassId":"appliance.CertRenewalPhase","Failed":true,"Message":"Message %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceCertRenewalPhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceCertRenewalPhase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceCertRenewalPhase(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "appliance.CertRenewalPhase", "class_id": "appliance.CertRenewalPhase", "failed": true, "message": "Message 1", "name": "Name 1"}, {"object_type": "appliance.CertRenewalPhase", "class_id": "appliance.CertRenewalPhase", "failed": true, "message": "Message 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceDataExportPolicyRelationship(t *testing.T) {
	p := []models.ApplianceDataExportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceDataExportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceDataExportPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceDataExportPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceFileSystemStatusRelationship(t *testing.T) {
	p := []models.ApplianceFileSystemStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceFileSystemStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceFileSystemStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceFileSystemStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceGroupStatusRelationship(t *testing.T) {
	p := []models.ApplianceGroupStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceGroupStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceGroupStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceGroupStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceKeyValuePair(t *testing.T) {
	p := []models.ApplianceKeyValuePair{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"appliance.KeyValuePair","Key":"Key %d","Value":"Value %d","ClassId":"appliance.KeyValuePair"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceKeyValuePair(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceKeyValuePair{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceKeyValuePair(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "appliance.KeyValuePair", "key": "Key 1", "value": "Value 1", "class_id": "appliance.KeyValuePair"}, {"object_type": "appliance.KeyValuePair", "key": "Key 2", "value": "Value 2", "class_id": "appliance.KeyValuePair"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListApplianceStatusCheck(t *testing.T) {
	p := []models.ApplianceStatusCheck{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"appliance.StatusCheck","Code":"Code %d","Result":"Result %d","ClassId":"appliance.StatusCheck"}`

	//test when the response is empty
	ffOpEmpty := flattenListApplianceStatusCheck(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ApplianceStatusCheck{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListApplianceStatusCheck(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "appliance.StatusCheck", "code": "Code 1", "result": "Result 1", "class_id": "appliance.StatusCheck"}, {"object_type": "appliance.StatusCheck", "code": "Code 2", "result": "Result 2", "class_id": "appliance.StatusCheck"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetClusterMemberRelationship(t *testing.T) {
	p := []models.AssetClusterMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetClusterMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetClusterMemberRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetClusterMemberRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetConnection(t *testing.T) {
	p := []models.AssetConnection{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.Connection","ObjectType":"asset.Connection"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetConnection(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetConnection{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetConnection(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "asset.Connection", "object_type": "asset.Connection"}, {"class_id": "asset.Connection", "object_type": "asset.Connection"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeploymentRelationship(t *testing.T) {
	p := []models.AssetDeploymentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeploymentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeploymentRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeploymentRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeploymentDeviceRelationship(t *testing.T) {
	p := []models.AssetDeploymentDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeploymentDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeploymentDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeploymentDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetDeviceRegistrationRelationship(t *testing.T) {
	p := []models.AssetDeviceRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetDeviceRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetDeviceRegistrationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetDeviceRegistrationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetMeteringType(t *testing.T) {
	p := []models.AssetMeteringType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.MeteringType","ObjectType":"asset.MeteringType","Name":"Name %d","Unit":"Unit %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetMeteringType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetMeteringType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetMeteringType(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "asset.MeteringType", "object_type": "asset.MeteringType", "name": "Name 1", "unit": "Unit 1"}, {"class_id": "asset.MeteringType", "object_type": "asset.MeteringType", "name": "Name 2", "unit": "Unit 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListAssetService(t *testing.T) {
	p := []models.AssetService{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.Service","ObjectType":"asset.Service"}`

	//test when the response is empty
	ffOpEmpty := flattenListAssetService(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.AssetService{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListAssetService(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "asset.Service", "object_type": "asset.Service"}, {"class_id": "asset.Service", "object_type": "asset.Service"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBiosBootDeviceRelationship(t *testing.T) {
	p := []models.BiosBootDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBiosBootDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BiosBootDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBiosBootDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBiosUnitRelationship(t *testing.T) {
	p := []models.BiosUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBiosUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BiosUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBiosUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootCddDeviceRelationship(t *testing.T) {
	p := []models.BootCddDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootCddDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootCddDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootCddDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootDeviceBase(t *testing.T) {
	p := []models.BootDeviceBase{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Enabled":true,"ClassId":"boot.DeviceBase","ObjectType":"boot.DeviceBase"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootDeviceBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootDeviceBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootDeviceBase(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "enabled": true, "class_id": "boot.DeviceBase", "object_type": "boot.DeviceBase"}, {"name": "Name 2", "enabled": true, "class_id": "boot.DeviceBase", "object_type": "boot.DeviceBase"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootHddDeviceRelationship(t *testing.T) {
	p := []models.BootHddDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootHddDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootHddDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootHddDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootIscsiDeviceRelationship(t *testing.T) {
	p := []models.BootIscsiDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootIscsiDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootIscsiDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootIscsiDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootNvmeDeviceRelationship(t *testing.T) {
	p := []models.BootNvmeDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootNvmeDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootNvmeDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootNvmeDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootPchStorageDeviceRelationship(t *testing.T) {
	p := []models.BootPchStorageDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootPchStorageDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootPchStorageDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootPchStorageDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootPxeDeviceRelationship(t *testing.T) {
	p := []models.BootPxeDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootPxeDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootPxeDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootPxeDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootSanDeviceRelationship(t *testing.T) {
	p := []models.BootSanDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootSanDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootSanDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootSanDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootSdDeviceRelationship(t *testing.T) {
	p := []models.BootSdDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootSdDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootSdDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootSdDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootUefiShellDeviceRelationship(t *testing.T) {
	p := []models.BootUefiShellDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootUefiShellDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootUefiShellDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootUefiShellDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootUsbDeviceRelationship(t *testing.T) {
	p := []models.BootUsbDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootUsbDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootUsbDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootUsbDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBootVmediaDeviceRelationship(t *testing.T) {
	p := []models.BootVmediaDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBootVmediaDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BootVmediaDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBootVmediaDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBulkApiResult(t *testing.T) {
	p := []models.BulkApiResult{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"bulk.ApiResult","ObjectType":"bulk.ApiResult","Status":32}`

	//test when the response is empty
	ffOpEmpty := flattenListBulkApiResult(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BulkApiResult{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBulkApiResult(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "bulk.ApiResult", "object_type": "bulk.ApiResult", "status": 32}, {"class_id": "bulk.ApiResult", "object_type": "bulk.ApiResult", "status": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBulkExportedItemRelationship(t *testing.T) {
	p := []models.BulkExportedItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBulkExportedItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BulkExportedItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBulkExportedItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBulkHttpHeader(t *testing.T) {
	p := []models.BulkHttpHeader{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"bulk.HttpHeader","ObjectType":"bulk.HttpHeader","Value":"Value %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBulkHttpHeader(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BulkHttpHeader{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBulkHttpHeader(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "bulk.HttpHeader", "object_type": "bulk.HttpHeader", "value": "Value 1", "name": "Name 1"}, {"class_id": "bulk.HttpHeader", "object_type": "bulk.HttpHeader", "value": "Value 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBulkSubRequest(t *testing.T) {
	p := []models.BulkSubRequest{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"bulk.SubRequest","ObjectType":"bulk.SubRequest","Verb":"Verb %d","Uri":"Uri %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBulkSubRequest(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BulkSubRequest{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBulkSubRequest(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "bulk.SubRequest", "object_type": "bulk.SubRequest", "verb": "Verb 1", "uri": "Uri 1"}, {"class_id": "bulk.SubRequest", "object_type": "bulk.SubRequest", "verb": "Verb 2", "uri": "Uri 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListBulkSubRequestObjRelationship(t *testing.T) {
	p := []models.BulkSubRequestObjRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListBulkSubRequestObjRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.BulkSubRequestObjRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListBulkSubRequestObjRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilityCapabilityRelationship(t *testing.T) {
	p := []models.CapabilityCapabilityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilityCapabilityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilityCapabilityRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilityCapabilityRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilityPortRange(t *testing.T) {
	p := []models.CapabilityPortRange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"capability.PortRange","ObjectType":"capability.PortRange","StartPortId":32,"StartSlotId":32,"EndPortId":32,"EndSlotId":32}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilityPortRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilityPortRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilityPortRange(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "capability.PortRange", "object_type": "capability.PortRange", "start_port_id": 32, "start_slot_id": 32, "end_port_id": 32, "end_slot_id": 32}, {"class_id": "capability.PortRange", "object_type": "capability.PortRange", "start_port_id": 32, "start_slot_id": 32, "end_port_id": 32, "end_slot_id": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCapabilitySwitchingModeCapability(t *testing.T) {
	p := []models.CapabilitySwitchingModeCapability{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"capability.SwitchingModeCapability","ObjectType":"capability.SwitchingModeCapability","SwitchingMode":"SwitchingMode %d","VpCompressionSupported":true}`

	//test when the response is empty
	ffOpEmpty := flattenListCapabilitySwitchingModeCapability(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CapabilitySwitchingModeCapability{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCapabilitySwitchingModeCapability(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "capability.SwitchingModeCapability", "object_type": "capability.SwitchingModeCapability", "switching_mode": "SwitchingMode 1", "vp_compression_supported": true}, {"class_id": "capability.SwitchingModeCapability", "object_type": "capability.SwitchingModeCapability", "switching_mode": "SwitchingMode 2", "vp_compression_supported": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCertificatemanagementCertificateBase(t *testing.T) {
	p := []models.CertificatemanagementCertificateBase{}
	var d = &schema.ResourceData{}
	c := `{"IsPrivatekeySet":true,"ClassId":"certificatemanagement.CertificateBase","ObjectType":"certificatemanagement.CertificateBase","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenListCertificatemanagementCertificateBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CertificatemanagementCertificateBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCertificatemanagementCertificateBase(p, d)
	expectedOp := []map[string]interface{}{{"is_privatekey_set": true, "class_id": "certificatemanagement.CertificateBase", "object_type": "certificatemanagement.CertificateBase", "enabled": true}, {"is_privatekey_set": true, "class_id": "certificatemanagement.CertificateBase", "object_type": "certificatemanagement.CertificateBase", "enabled": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisConfigChangeDetailRelationship(t *testing.T) {
	p := []models.ChassisConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisConfigResultEntryRelationship(t *testing.T) {
	p := []models.ChassisConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListChassisIomProfileRelationship(t *testing.T) {
	p := []models.ChassisIomProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListChassisIomProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ChassisIomProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListChassisIomProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudAwsSecurityGroupRelationship(t *testing.T) {
	p := []models.CloudAwsSecurityGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudAwsSecurityGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudAwsSecurityGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudAwsSecurityGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudCloudTag(t *testing.T) {
	p := []models.CloudCloudTag{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.CloudTag","ObjectType":"cloud.CloudTag","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudCloudTag(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudCloudTag{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudCloudTag(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "cloud.CloudTag", "object_type": "cloud.CloudTag", "key": "Key 1", "value": "Value 1"}, {"class_id": "cloud.CloudTag", "object_type": "cloud.CloudTag", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudCustomAttributes(t *testing.T) {
	p := []models.CloudCustomAttributes{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"cloud.CustomAttributes","AttributeName":"AttributeName %d","AttributeType":"AttributeType %d","AttributeValue":"AttributeValue %d","ClassId":"cloud.CustomAttributes"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudCustomAttributes(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudCustomAttributes{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudCustomAttributes(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "cloud.CustomAttributes", "attribute_name": "AttributeName 1", "attribute_type": "AttributeType 1", "attribute_value": "AttributeValue 1", "class_id": "cloud.CustomAttributes"}, {"object_type": "cloud.CustomAttributes", "attribute_name": "AttributeName 2", "attribute_type": "AttributeType 2", "attribute_value": "AttributeValue 2", "class_id": "cloud.CustomAttributes"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudNetworkInterfaceAttachment(t *testing.T) {
	p := []models.CloudNetworkInterfaceAttachment{}
	var d = &schema.ResourceData{}
	c := `{"SubNetworkId":"SubNetworkId %d","MacAddress":"MacAddress %d","ClassId":"cloud.NetworkInterfaceAttachment","NicIndex":32,"SubNetworkName":"SubNetworkName %d","Identity":"Identity %d","NetworkId":"NetworkId %d","ObjectType":"cloud.NetworkInterfaceAttachment","NetworkName":"NetworkName %d","IpForwardingEnabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudNetworkInterfaceAttachment(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudNetworkInterfaceAttachment{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudNetworkInterfaceAttachment(p, d)
	expectedOp := []map[string]interface{}{{"sub_network_id": "SubNetworkId 1", "mac_address": "MacAddress 1", "class_id": "cloud.NetworkInterfaceAttachment", "nic_index": 32, "sub_network_name": "SubNetworkName 1", "identity": "Identity 1", "network_id": "NetworkId 1", "object_type": "cloud.NetworkInterfaceAttachment", "network_name": "NetworkName 1", "ip_forwarding_enabled": true}, {"sub_network_id": "SubNetworkId 2", "mac_address": "MacAddress 2", "class_id": "cloud.NetworkInterfaceAttachment", "nic_index": 32, "sub_network_name": "SubNetworkName 2", "identity": "Identity 2", "network_id": "NetworkId 2", "object_type": "cloud.NetworkInterfaceAttachment", "network_name": "NetworkName 2", "ip_forwarding_enabled": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudSecurityGroupRule(t *testing.T) {
	p := []models.CloudSecurityGroupRule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.SecurityGroupRule","ObjectType":"cloud.SecurityGroupRule","StartPort":32,"Protocol":"Protocol %d","Index":32,"Action":"Action %d","EndPort":32,"Description":"Description %d","Name":"Name %d","EtherType":"EtherType %d","SourceSecurityGroup":"SourceSecurityGroup %d","Identity":"Identity %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudSecurityGroupRule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudSecurityGroupRule{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudSecurityGroupRule(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "cloud.SecurityGroupRule", "object_type": "cloud.SecurityGroupRule", "start_port": 32, "protocol": "Protocol 1", "index": 32, "action": "Action 1", "end_port": 32, "description": "Description 1", "name": "Name 1", "ether_type": "EtherType 1", "source_security_group": "SourceSecurityGroup 1", "identity": "Identity 1"}, {"class_id": "cloud.SecurityGroupRule", "object_type": "cloud.SecurityGroupRule", "start_port": 32, "protocol": "Protocol 2", "index": 32, "action": "Action 2", "end_port": 32, "description": "Description 2", "name": "Name 2", "ether_type": "EtherType 2", "source_security_group": "SourceSecurityGroup 2", "identity": "Identity 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudTfcWorkspaceVariables(t *testing.T) {
	p := []models.CloudTfcWorkspaceVariables{}
	var d = &schema.ResourceData{}
	c := `{"Value":"Value %d","Identity":"Identity %d","Description":"Description %d","Key":"Key %d","Sensitive":true,"ClassId":"cloud.TfcWorkspaceVariables","ObjectType":"cloud.TfcWorkspaceVariables","Hcl":true,"Category":"Category %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudTfcWorkspaceVariables(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudTfcWorkspaceVariables{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudTfcWorkspaceVariables(p, d)
	expectedOp := []map[string]interface{}{{"value": "Value 1", "identity": "Identity 1", "description": "Description 1", "key": "Key 1", "sensitive": true, "class_id": "cloud.TfcWorkspaceVariables", "object_type": "cloud.TfcWorkspaceVariables", "hcl": true, "category": "Category 1"}, {"value": "Value 2", "identity": "Identity 2", "description": "Description 2", "key": "Key 2", "sensitive": true, "class_id": "cloud.TfcWorkspaceVariables", "object_type": "cloud.TfcWorkspaceVariables", "hcl": true, "category": "Category 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudVolumeAttachment(t *testing.T) {
	p := []models.CloudVolumeAttachment{}
	var d = &schema.ResourceData{}
	c := `{"Index":32,"ObjectType":"cloud.VolumeAttachment","Identity":"Identity %d","DeviceName":"DeviceName %d","ClassId":"cloud.VolumeAttachment","IsRoot":true,"AutoDelete":true}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudVolumeAttachment(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudVolumeAttachment{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudVolumeAttachment(p, d)
	expectedOp := []map[string]interface{}{{"index": 32, "object_type": "cloud.VolumeAttachment", "identity": "Identity 1", "device_name": "DeviceName 1", "class_id": "cloud.VolumeAttachment", "is_root": true, "auto_delete": true}, {"index": 32, "object_type": "cloud.VolumeAttachment", "identity": "Identity 2", "device_name": "DeviceName 2", "class_id": "cloud.VolumeAttachment", "is_root": true, "auto_delete": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCloudVolumeInstanceAttachment(t *testing.T) {
	p := []models.CloudVolumeInstanceAttachment{}
	var d = &schema.ResourceData{}
	c := `{"DeviceName":"DeviceName %d","InstanceId":"InstanceId %d","State":"State %d","ClassId":"cloud.VolumeInstanceAttachment","ObjectType":"cloud.VolumeInstanceAttachment","AutoDelete":true}`

	//test when the response is empty
	ffOpEmpty := flattenListCloudVolumeInstanceAttachment(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CloudVolumeInstanceAttachment{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCloudVolumeInstanceAttachment(p, d)
	expectedOp := []map[string]interface{}{{"device_name": "DeviceName 1", "instance_id": "InstanceId 1", "state": "State 1", "class_id": "cloud.VolumeInstanceAttachment", "object_type": "cloud.VolumeInstanceAttachment", "auto_delete": true}, {"device_name": "DeviceName 2", "instance_id": "InstanceId 2", "state": "State 2", "class_id": "cloud.VolumeInstanceAttachment", "object_type": "cloud.VolumeInstanceAttachment", "auto_delete": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeBladeRelationship(t *testing.T) {
	p := []models.ComputeBladeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeBladeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeBladeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeBladeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeIpAddress(t *testing.T) {
	p := []models.ComputeIpAddress{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.IpAddress","KvmVlan":32,"DefaultGateway":"DefaultGateway %d","HttpPort":32,"Name":"Name %d","Dn":"Dn %d","ObjectType":"compute.IpAddress","Address":"Address %d","Type":"Type %d","HttpsPort":32,"Subnet":"Subnet %d","KvmPort":32,"Category":"Category %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeIpAddress(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeIpAddress{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeIpAddress(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "compute.IpAddress", "kvm_vlan": 32, "default_gateway": "DefaultGateway 1", "http_port": 32, "name": "Name 1", "dn": "Dn 1", "object_type": "compute.IpAddress", "address": "Address 1", "type": "Type 1", "https_port": 32, "subnet": "Subnet 1", "kvm_port": 32, "category": "Category 1"}, {"class_id": "compute.IpAddress", "kvm_vlan": 32, "default_gateway": "DefaultGateway 2", "http_port": 32, "name": "Name 2", "dn": "Dn 2", "object_type": "compute.IpAddress", "address": "Address 2", "type": "Type 2", "https_port": 32, "subnet": "Subnet 2", "kvm_port": 32, "category": "Category 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeMappingRelationship(t *testing.T) {
	p := []models.ComputeMappingRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeMappingRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeMappingRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeMappingRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputePhysicalRelationship(t *testing.T) {
	p := []models.ComputePhysicalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputePhysicalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputePhysicalRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputePhysicalRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListComputeRackUnitRelationship(t *testing.T) {
	p := []models.ComputeRackUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListComputeRackUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ComputeRackUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListComputeRackUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCondHclStatusDetailRelationship(t *testing.T) {
	p := []models.CondHclStatusDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCondHclStatusDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CondHclStatusDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCondHclStatusDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListConnectorpackConnectorPackUpdate(t *testing.T) {
	p := []models.ConnectorpackConnectorPackUpdate{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"connectorpack.ConnectorPackUpdate","ObjectType":"connectorpack.ConnectorPackUpdate","NewVersion":"NewVersion %d","CurrentVersion":"CurrentVersion %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListConnectorpackConnectorPackUpdate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ConnectorpackConnectorPackUpdate{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListConnectorpackConnectorPackUpdate(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "connectorpack.ConnectorPackUpdate", "object_type": "connectorpack.ConnectorPackUpdate", "new_version": "NewVersion 1", "current_version": "CurrentVersion 1", "name": "Name 1"}, {"class_id": "connectorpack.ConnectorPackUpdate", "object_type": "connectorpack.ConnectorPackUpdate", "new_version": "NewVersion 2", "current_version": "CurrentVersion 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListContentComplexType(t *testing.T) {
	p := []models.ContentComplexType{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","ClassId":"content.ComplexType","ObjectType":"content.ComplexType"}`

	//test when the response is empty
	ffOpEmpty := flattenListContentComplexType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ContentComplexType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListContentComplexType(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "class_id": "content.ComplexType", "object_type": "content.ComplexType"}, {"name": "Name 2", "class_id": "content.ComplexType", "object_type": "content.ComplexType"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListContentParameter(t *testing.T) {
	p := []models.ContentParameter{}
	var d = &schema.ResourceData{}
	c := `{"ComplexType":"ComplexType %d","Secure":true,"Name":"Name %d","Type":"Type %d","ClassId":"content.Parameter","AcceptSingleValue":true,"ItemType":"ItemType %d","ObjectType":"content.Parameter","Path":"Path %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListContentParameter(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ContentParameter{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListContentParameter(p, d)
	expectedOp := []map[string]interface{}{{"complex_type": "ComplexType 1", "secure": true, "name": "Name 1", "type": "Type 1", "class_id": "content.Parameter", "accept_single_value": true, "item_type": "ItemType 1", "object_type": "content.Parameter", "path": "Path 1"}, {"complex_type": "ComplexType 2", "secure": true, "name": "Name 2", "type": "Type 2", "class_id": "content.Parameter", "accept_single_value": true, "item_type": "ItemType 2", "object_type": "content.Parameter", "path": "Path 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListCrdCustomResourceConfigProperty(t *testing.T) {
	p := []models.CrdCustomResourceConfigProperty{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"crd.CustomResourceConfigProperty","ObjectType":"crd.CustomResourceConfigProperty","Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListCrdCustomResourceConfigProperty(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.CrdCustomResourceConfigProperty{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListCrdCustomResourceConfigProperty(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "crd.CustomResourceConfigProperty", "object_type": "crd.CustomResourceConfigProperty", "key": "Key 1", "value": "Value 1"}, {"class_id": "crd.CustomResourceConfigProperty", "object_type": "crd.CustomResourceConfigProperty", "key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentExpanderModuleRelationship(t *testing.T) {
	p := []models.EquipmentExpanderModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentExpanderModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentExpanderModuleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentExpanderModuleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentFanRelationship(t *testing.T) {
	p := []models.EquipmentFanRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentFanRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentFanRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentFanRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentFanModuleRelationship(t *testing.T) {
	p := []models.EquipmentFanModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentFanModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentFanModuleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentFanModuleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoCardRelationship(t *testing.T) {
	p := []models.EquipmentIoCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoCardIdentity(t *testing.T) {
	p := []models.EquipmentIoCardIdentity{}
	var d = &schema.ResourceData{}
	c := `{"IoCardMoid":"IoCardMoid %d","ModuleId":32,"NetworkElementMoid":"NetworkElementMoid %d","SwitchId":32,"ObjectType":"equipment.IoCardIdentity","ClassId":"equipment.IoCardIdentity"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoCardIdentity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoCardIdentity{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoCardIdentity(p, d)
	expectedOp := []map[string]interface{}{{"io_card_moid": "IoCardMoid 1", "module_id": 32, "network_element_moid": "NetworkElementMoid 1", "switch_id": 32, "object_type": "equipment.IoCardIdentity", "class_id": "equipment.IoCardIdentity"}, {"io_card_moid": "IoCardMoid 2", "module_id": 32, "network_element_moid": "NetworkElementMoid 2", "switch_id": 32, "object_type": "equipment.IoCardIdentity", "class_id": "equipment.IoCardIdentity"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentIoExpanderRelationship(t *testing.T) {
	p := []models.EquipmentIoExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentIoExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentIoExpanderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentIoExpanderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentPsuRelationship(t *testing.T) {
	p := []models.EquipmentPsuRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentPsuRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentPsuRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentPsuRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentRackEnclosureSlotRelationship(t *testing.T) {
	p := []models.EquipmentRackEnclosureSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentRackEnclosureSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentRackEnclosureSlotRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentRackEnclosureSlotRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentSwitchCardRelationship(t *testing.T) {
	p := []models.EquipmentSwitchCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentSwitchCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentSwitchCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentSwitchCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentSystemIoControllerRelationship(t *testing.T) {
	p := []models.EquipmentSystemIoControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentSystemIoControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentSystemIoControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentSystemIoControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEquipmentTpmRelationship(t *testing.T) {
	p := []models.EquipmentTpmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEquipmentTpmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EquipmentTpmRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEquipmentTpmRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherHostPortRelationship(t *testing.T) {
	p := []models.EtherHostPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherHostPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherHostPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherHostPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherNetworkPortRelationship(t *testing.T) {
	p := []models.EtherNetworkPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherNetworkPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherNetworkPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherNetworkPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherPhysicalPortRelationship(t *testing.T) {
	p := []models.EtherPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherPhysicalPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherPhysicalPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListEtherPortChannelRelationship(t *testing.T) {
	p := []models.EtherPortChannelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListEtherPortChannelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.EtherPortChannelRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListEtherPortChannelRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricConfigChangeDetailRelationship(t *testing.T) {
	p := []models.FabricConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricConfigResultEntryRelationship(t *testing.T) {
	p := []models.FabricConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricEthNetworkGroupPolicyRelationship(t *testing.T) {
	p := []models.FabricEthNetworkGroupPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricEthNetworkGroupPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricEthNetworkGroupPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricEthNetworkGroupPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricPortIdentifier(t *testing.T) {
	p := []models.FabricPortIdentifier{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.PortIdentifier","AggregatePortId":32,"PortId":32,"SlotId":32,"ObjectType":"fabric.PortIdentifier"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricPortIdentifier(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricPortIdentifier{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricPortIdentifier(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "fabric.PortIdentifier", "aggregate_port_id": 32, "port_id": 32, "slot_id": 32, "object_type": "fabric.PortIdentifier"}, {"class_id": "fabric.PortIdentifier", "aggregate_port_id": 32, "port_id": 32, "slot_id": 32, "object_type": "fabric.PortIdentifier"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricQosClass(t *testing.T) {
	p := []models.FabricQosClass{}
	var d = &schema.ResourceData{}
	c := `{"MulticastOptimize":true,"Mtu":32,"Name":"Name %d","ObjectType":"fabric.QosClass","Cos":32,"AdminState":"AdminState %d","BandwidthPercent":32,"ClassId":"fabric.QosClass","PacketDrop":true,"Weight":32}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricQosClass(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricQosClass{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricQosClass(p, d)
	expectedOp := []map[string]interface{}{{"multicast_optimize": true, "mtu": 32, "name": "Name 1", "object_type": "fabric.QosClass", "cos": 32, "admin_state": "AdminState 1", "bandwidth_percent": 32, "class_id": "fabric.QosClass", "packet_drop": true, "weight": 32}, {"multicast_optimize": true, "mtu": 32, "name": "Name 2", "object_type": "fabric.QosClass", "cos": 32, "admin_state": "AdminState 2", "bandwidth_percent": 32, "class_id": "fabric.QosClass", "packet_drop": true, "weight": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFabricSwitchProfileRelationship(t *testing.T) {
	p := []models.FabricSwitchProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFabricSwitchProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FabricSwitchProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFabricSwitchProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcPhysicalPortRelationship(t *testing.T) {
	p := []models.FcPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcPhysicalPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcPhysicalPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcPortChannelRelationship(t *testing.T) {
	p := []models.FcPortChannelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcPortChannelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcPortChannelRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcPortChannelRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcpoolBlock(t *testing.T) {
	p := []models.FcpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ClassId":"fcpool.Block","Size":32,"ObjectType":"fcpool.Block"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcpoolBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcpoolBlock(p, d)
	expectedOp := []map[string]interface{}{{"from": "From 1", "to": "To 1", "class_id": "fcpool.Block", "size": 32, "object_type": "fcpool.Block"}, {"from": "From 2", "to": "To 2", "class_id": "fcpool.Block", "size": 32, "object_type": "fcpool.Block"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFcpoolFcBlockRelationship(t *testing.T) {
	p := []models.FcpoolFcBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFcpoolFcBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FcpoolFcBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFcpoolFcBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareBaseImpact(t *testing.T) {
	p := []models.FirmwareBaseImpact{}
	var d = &schema.ResourceData{}
	c := `{"DomainName":"DomainName %d","ComputationStatusDetail":"ComputationStatusDetail %d","ObjectType":"firmware.BaseImpact","VersionImpact":"VersionImpact %d","ImpactType":"ImpactType %d","FirmwareVersion":"FirmwareVersion %d","Model":"Model %d","EndPoint":"EndPoint %d","TargetFirmwareVersion":"TargetFirmwareVersion %d","ComputationError":"ComputationError %d","ClassId":"firmware.BaseImpact"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareBaseImpact(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareBaseImpact{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareBaseImpact(p, d)
	expectedOp := []map[string]interface{}{{"domain_name": "DomainName 1", "computation_status_detail": "ComputationStatusDetail 1", "object_type": "firmware.BaseImpact", "version_impact": "VersionImpact 1", "impact_type": "ImpactType 1", "firmware_version": "FirmwareVersion 1", "model": "Model 1", "end_point": "EndPoint 1", "target_firmware_version": "TargetFirmwareVersion 1", "computation_error": "ComputationError 1", "class_id": "firmware.BaseImpact"}, {"domain_name": "DomainName 2", "computation_status_detail": "ComputationStatusDetail 2", "object_type": "firmware.BaseImpact", "version_impact": "VersionImpact 2", "impact_type": "ImpactType 2", "firmware_version": "FirmwareVersion 2", "model": "Model 2", "end_point": "EndPoint 2", "target_firmware_version": "TargetFirmwareVersion 2", "computation_error": "ComputationError 2", "class_id": "firmware.BaseImpact"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareComponentMeta(t *testing.T) {
	p := []models.FirmwareComponentMeta{}
	var d = &schema.ResourceData{}
	c := `{"Disruption":"Disruption %d","RedfishUrl":"RedfishUrl %d","Vendor":"Vendor %d","Description":"Description %d","IsOobSupported":true,"PackedVersion":"PackedVersion %d","Model":"Model %d","ClassId":"firmware.ComponentMeta","ObjectType":"firmware.ComponentMeta","ComponentType":"ComponentType %d","ComponentLabel":"ComponentLabel %d","ImagePath":"ImagePath %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareComponentMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareComponentMeta{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareComponentMeta(p, d)
	expectedOp := []map[string]interface{}{{"disruption": "Disruption 1", "redfish_url": "RedfishUrl 1", "vendor": "Vendor 1", "description": "Description 1", "is_oob_supported": true, "packed_version": "PackedVersion 1", "model": "Model 1", "class_id": "firmware.ComponentMeta", "object_type": "firmware.ComponentMeta", "component_type": "ComponentType 1", "component_label": "ComponentLabel 1", "image_path": "ImagePath 1"}, {"disruption": "Disruption 2", "redfish_url": "RedfishUrl 2", "vendor": "Vendor 2", "description": "Description 2", "is_oob_supported": true, "packed_version": "PackedVersion 2", "model": "Model 2", "class_id": "firmware.ComponentMeta", "object_type": "firmware.ComponentMeta", "component_type": "ComponentType 2", "component_label": "ComponentLabel 2", "image_path": "ImagePath 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareDistributableMetaRelationship(t *testing.T) {
	p := []models.FirmwareDistributableMetaRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareDistributableMetaRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareDistributableMetaRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareDistributableMetaRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareFirmwareInventory(t *testing.T) {
	p := []models.FirmwareFirmwareInventory{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"firmware.FirmwareInventory","Label":"Label %d","ObjectType":"firmware.FirmwareInventory","Category":"Category %d","Model":"Model %d","UpdateUri":"UpdateUri %d","Vendor":"Vendor %d","Version":"Version %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareFirmwareInventory(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareFirmwareInventory{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareFirmwareInventory(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "firmware.FirmwareInventory", "label": "Label 1", "object_type": "firmware.FirmwareInventory", "category": "Category 1", "model": "Model 1", "update_uri": "UpdateUri 1", "vendor": "Vendor 1", "nr_version": "Version 1"}, {"class_id": "firmware.FirmwareInventory", "label": "Label 2", "object_type": "firmware.FirmwareInventory", "category": "Category 2", "model": "Model 2", "update_uri": "UpdateUri 2", "vendor": "Vendor 2", "nr_version": "Version 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListFirmwareRunningFirmwareRelationship(t *testing.T) {
	p := []models.FirmwareRunningFirmwareRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListFirmwareRunningFirmwareRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.FirmwareRunningFirmwareRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListFirmwareRunningFirmwareRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListForecastDefinitionRelationship(t *testing.T) {
	p := []models.ForecastDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListForecastDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ForecastDefinitionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListForecastDefinitionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListGraphicsCardRelationship(t *testing.T) {
	p := []models.GraphicsCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListGraphicsCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.GraphicsCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListGraphicsCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListGraphicsControllerRelationship(t *testing.T) {
	p := []models.GraphicsControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListGraphicsControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.GraphicsControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListGraphicsControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclConstraint(t *testing.T) {
	p := []models.HclConstraint{}
	var d = &schema.ResourceData{}
	c := `{"ConstraintName":"ConstraintName %d","ConstraintValue":"ConstraintValue %d","ObjectType":"hcl.Constraint","ClassId":"hcl.Constraint"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclConstraint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclConstraint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclConstraint(p, d)
	expectedOp := []map[string]interface{}{{"constraint_name": "ConstraintName 1", "constraint_value": "ConstraintValue 1", "object_type": "hcl.Constraint", "class_id": "hcl.Constraint"}, {"constraint_name": "ConstraintName 2", "constraint_value": "ConstraintValue 2", "object_type": "hcl.Constraint", "class_id": "hcl.Constraint"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclHyperflexSoftwareCompatibilityInfoRelationship(t *testing.T) {
	p := []models.HclHyperflexSoftwareCompatibilityInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclHyperflexSoftwareCompatibilityInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclHyperflexSoftwareCompatibilityInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHclOperatingSystemRelationship(t *testing.T) {
	p := []models.HclOperatingSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHclOperatingSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HclOperatingSystemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHclOperatingSystemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexAlarmRelationship(t *testing.T) {
	p := []models.HyperflexAlarmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexAlarmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexAlarmRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexAlarmRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexBaseClusterRelationship(t *testing.T) {
	p := []models.HyperflexBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexBaseClusterRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexBaseClusterRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexCapabilityInfoRelationship(t *testing.T) {
	p := []models.HyperflexCapabilityInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexCapabilityInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexCapabilityInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexCapabilityInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexClusterProfileRelationship(t *testing.T) {
	p := []models.HyperflexClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexClusterProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexClusterProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexConfigResultEntryRelationship(t *testing.T) {
	p := []models.HyperflexConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexDriveRelationship(t *testing.T) {
	p := []models.HyperflexDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexFeatureLimitEntry(t *testing.T) {
	p := []models.HyperflexFeatureLimitEntry{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.FeatureLimitEntry","Name":"Name %d","Value":"Value %d","ObjectType":"hyperflex.FeatureLimitEntry"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexFeatureLimitEntry(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexFeatureLimitEntry{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexFeatureLimitEntry(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.FeatureLimitEntry", "name": "Name 1", "value": "Value 1", "object_type": "hyperflex.FeatureLimitEntry"}, {"class_id": "hyperflex.FeatureLimitEntry", "name": "Name 2", "value": "Value 2", "object_type": "hyperflex.FeatureLimitEntry"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHealthCheckScriptInfo(t *testing.T) {
	p := []models.HyperflexHealthCheckScriptInfo{}
	var d = &schema.ResourceData{}
	c := `{"AggregateScriptName":"AggregateScriptName %d","HyperflexVersion":"HyperflexVersion %d","ScriptInput":"ScriptInput %d","ScriptName":"ScriptName %d","ObjectType":"hyperflex.HealthCheckScriptInfo","ScriptExecuteLocation":"ScriptExecuteLocation %d","Version":"Version %d","ClassId":"hyperflex.HealthCheckScriptInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHealthCheckScriptInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHealthCheckScriptInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHealthCheckScriptInfo(p, d)
	expectedOp := []map[string]interface{}{{"aggregate_script_name": "AggregateScriptName 1", "hyperflex_version": "HyperflexVersion 1", "script_input": "ScriptInput 1", "script_name": "ScriptName 1", "object_type": "hyperflex.HealthCheckScriptInfo", "script_execute_location": "ScriptExecuteLocation 1", "nr_version": "Version 1", "class_id": "hyperflex.HealthCheckScriptInfo"}, {"aggregate_script_name": "AggregateScriptName 2", "hyperflex_version": "HyperflexVersion 2", "script_input": "ScriptInput 2", "script_name": "ScriptName 2", "object_type": "hyperflex.HealthCheckScriptInfo", "script_execute_location": "ScriptExecuteLocation 2", "nr_version": "Version 2", "class_id": "hyperflex.HealthCheckScriptInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxHostMountStatusDt(t *testing.T) {
	p := []models.HyperflexHxHostMountStatusDt{}
	var d = &schema.ResourceData{}
	c := `{"Mounted":true,"Reason":"Reason %d","ClassId":"hyperflex.HxHostMountStatusDt","ObjectType":"hyperflex.HxHostMountStatusDt","Accessibility":"Accessibility %d","HostName":"HostName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxHostMountStatusDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxHostMountStatusDt{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxHostMountStatusDt(p, d)
	expectedOp := []map[string]interface{}{{"mounted": true, "reason": "Reason 1", "class_id": "hyperflex.HxHostMountStatusDt", "object_type": "hyperflex.HxHostMountStatusDt", "accessibility": "Accessibility 1", "host_name": "HostName 1"}, {"mounted": true, "reason": "Reason 2", "class_id": "hyperflex.HxHostMountStatusDt", "object_type": "hyperflex.HxHostMountStatusDt", "accessibility": "Accessibility 2", "host_name": "HostName 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxZoneResiliencyInfoDt(t *testing.T) {
	p := []models.HyperflexHxZoneResiliencyInfoDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxZoneResiliencyInfoDt","ObjectType":"hyperflex.HxZoneResiliencyInfoDt","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxZoneResiliencyInfoDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxZoneResiliencyInfoDt{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxZoneResiliencyInfoDt(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.HxZoneResiliencyInfoDt", "object_type": "hyperflex.HxZoneResiliencyInfoDt", "name": "Name 1"}, {"class_id": "hyperflex.HxZoneResiliencyInfoDt", "object_type": "hyperflex.HxZoneResiliencyInfoDt", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexHxdpVersionRelationship(t *testing.T) {
	p := []models.HyperflexHxdpVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexHxdpVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexHxdpVersionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexHxdpVersionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexIpAddrRange(t *testing.T) {
	p := []models.HyperflexIpAddrRange{}
	var d = &schema.ResourceData{}
	c := `{"EndAddr":"EndAddr %d","Gateway":"Gateway %d","Netmask":"Netmask %d","StartAddr":"StartAddr %d","ClassId":"hyperflex.IpAddrRange","ObjectType":"hyperflex.IpAddrRange"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexIpAddrRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexIpAddrRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexIpAddrRange(p, d)
	expectedOp := []map[string]interface{}{{"end_addr": "EndAddr 1", "gateway": "Gateway 1", "netmask": "Netmask 1", "start_addr": "StartAddr 1", "class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange"}, {"end_addr": "EndAddr 2", "gateway": "Gateway 2", "netmask": "Netmask 2", "start_addr": "StartAddr 2", "class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexMapClusterIdToProtectionInfo(t *testing.T) {
	p := []models.HyperflexMapClusterIdToProtectionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClusterId":"ClusterId %d","ClassId":"hyperflex.MapClusterIdToProtectionInfo","ObjectType":"hyperflex.MapClusterIdToProtectionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexMapClusterIdToProtectionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexMapClusterIdToProtectionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexMapClusterIdToProtectionInfo(p, d)
	expectedOp := []map[string]interface{}{{"cluster_id": "ClusterId 1", "class_id": "hyperflex.MapClusterIdToProtectionInfo", "object_type": "hyperflex.MapClusterIdToProtectionInfo"}, {"cluster_id": "ClusterId 2", "class_id": "hyperflex.MapClusterIdToProtectionInfo", "object_type": "hyperflex.MapClusterIdToProtectionInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexMapClusterIdToStSnapshotPoint(t *testing.T) {
	p := []models.HyperflexMapClusterIdToStSnapshotPoint{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.MapClusterIdToStSnapshotPoint","ObjectType":"hyperflex.MapClusterIdToStSnapshotPoint","ClusterId":"ClusterId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexMapClusterIdToStSnapshotPoint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexMapClusterIdToStSnapshotPoint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexMapClusterIdToStSnapshotPoint(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.MapClusterIdToStSnapshotPoint", "object_type": "hyperflex.MapClusterIdToStSnapshotPoint", "cluster_id": "ClusterId 1"}, {"class_id": "hyperflex.MapClusterIdToStSnapshotPoint", "object_type": "hyperflex.MapClusterIdToStSnapshotPoint", "cluster_id": "ClusterId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNamedVlan(t *testing.T) {
	p := []models.HyperflexNamedVlan{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.NamedVlan","ObjectType":"hyperflex.NamedVlan","VlanId":32,"Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNamedVlan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNamedVlan{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNamedVlan(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "vlan_id": 32, "name": "Name 1"}, {"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "vlan_id": 32, "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNodeRelationship(t *testing.T) {
	p := []models.HyperflexNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNodeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNodeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexNodeProfileRelationship(t *testing.T) {
	p := []models.HyperflexNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexNodeProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexNodeProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexReplicationClusterReferenceToSchedule(t *testing.T) {
	p := []models.HyperflexReplicationClusterReferenceToSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.ReplicationClusterReferenceToSchedule","ObjectType":"hyperflex.ReplicationClusterReferenceToSchedule"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexReplicationClusterReferenceToSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexReplicationClusterReferenceToSchedule{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexReplicationClusterReferenceToSchedule(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.ReplicationClusterReferenceToSchedule", "object_type": "hyperflex.ReplicationClusterReferenceToSchedule"}, {"class_id": "hyperflex.ReplicationClusterReferenceToSchedule", "object_type": "hyperflex.ReplicationClusterReferenceToSchedule"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerFirmwareVersionEntryRelationship(t *testing.T) {
	p := []models.HyperflexServerFirmwareVersionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerFirmwareVersionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerFirmwareVersionEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerFirmwareVersionEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerFirmwareVersionInfo(t *testing.T) {
	p := []models.HyperflexServerFirmwareVersionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ServerPlatform":"ServerPlatform %d","Version":"Version %d","ClassId":"hyperflex.ServerFirmwareVersionInfo","ObjectType":"hyperflex.ServerFirmwareVersionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerFirmwareVersionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerFirmwareVersionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerFirmwareVersionInfo(p, d)
	expectedOp := []map[string]interface{}{{"server_platform": "ServerPlatform 1", "nr_version": "Version 1", "class_id": "hyperflex.ServerFirmwareVersionInfo", "object_type": "hyperflex.ServerFirmwareVersionInfo"}, {"server_platform": "ServerPlatform 2", "nr_version": "Version 2", "class_id": "hyperflex.ServerFirmwareVersionInfo", "object_type": "hyperflex.ServerFirmwareVersionInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexServerModelEntry(t *testing.T) {
	p := []models.HyperflexServerModelEntry{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.ServerModelEntry","Name":"Name %d","Value":"Value %d","ObjectType":"hyperflex.ServerModelEntry"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexServerModelEntry(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexServerModelEntry{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexServerModelEntry(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "hyperflex.ServerModelEntry", "name": "Name 1", "value": "Value 1", "object_type": "hyperflex.ServerModelEntry"}, {"class_id": "hyperflex.ServerModelEntry", "name": "Name 2", "value": "Value 2", "object_type": "hyperflex.ServerModelEntry"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionComponentRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionComponentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionComponentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionComponentRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionComponentRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionEntryRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexSoftwareDistributionVersionRelationship(t *testing.T) {
	p := []models.HyperflexSoftwareDistributionVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexSoftwareDistributionVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexSoftwareDistributionVersionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexSoftwareDistributionVersionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListHyperflexVolumeRelationship(t *testing.T) {
	p := []models.HyperflexVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListHyperflexVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.HyperflexVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListHyperflexVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasConnectorPackRelationship(t *testing.T) {
	p := []models.IaasConnectorPackRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasConnectorPackRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasConnectorPackRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasConnectorPackRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasDeviceStatusRelationship(t *testing.T) {
	p := []models.IaasDeviceStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasDeviceStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasDeviceStatusRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasDeviceStatusRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasLicenseKeysInfo(t *testing.T) {
	p := []models.IaasLicenseKeysInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iaas.LicenseKeysInfo","ObjectType":"iaas.LicenseKeysInfo","Count":32,"ExpirationDate":"ExpirationDate %d","LicenseId":"LicenseId %d","Pid":"Pid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasLicenseKeysInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasLicenseKeysInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasLicenseKeysInfo(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iaas.LicenseKeysInfo", "object_type": "iaas.LicenseKeysInfo", "nr_count": 32, "expiration_date": "ExpirationDate 1", "license_id": "LicenseId 1", "pid": "Pid 1"}, {"class_id": "iaas.LicenseKeysInfo", "object_type": "iaas.LicenseKeysInfo", "nr_count": 32, "expiration_date": "ExpirationDate 2", "license_id": "LicenseId 2", "pid": "Pid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasLicenseUtilizationInfo(t *testing.T) {
	p := []models.IaasLicenseUtilizationInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iaas.LicenseUtilizationInfo","ObjectType":"iaas.LicenseUtilizationInfo","LicensedLimit":"LicensedLimit %d","Sku":"Sku %d","ActualUsed":32,"Label":"Label %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasLicenseUtilizationInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasLicenseUtilizationInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasLicenseUtilizationInfo(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "iaas.LicenseUtilizationInfo", "object_type": "iaas.LicenseUtilizationInfo", "licensed_limit": "LicensedLimit 1", "sku": "Sku 1", "actual_used": 32, "label": "Label 1"}, {"class_id": "iaas.LicenseUtilizationInfo", "object_type": "iaas.LicenseUtilizationInfo", "licensed_limit": "LicensedLimit 2", "sku": "Sku 2", "actual_used": 32, "label": "Label 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasMostRunTasksRelationship(t *testing.T) {
	p := []models.IaasMostRunTasksRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasMostRunTasksRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasMostRunTasksRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasMostRunTasksRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIaasWorkflowSteps(t *testing.T) {
	p := []models.IaasWorkflowSteps{}
	var d = &schema.ResourceData{}
	c := `{"CompletedTime":"CompletedTime %d","ClassId":"iaas.WorkflowSteps","ObjectType":"iaas.WorkflowSteps","Name":"Name %d","Status":"Status %d","StatusMessage":"StatusMessage %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIaasWorkflowSteps(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IaasWorkflowSteps{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIaasWorkflowSteps(p, d)
	expectedOp := []map[string]interface{}{{"completed_time": "CompletedTime 1", "class_id": "iaas.WorkflowSteps", "object_type": "iaas.WorkflowSteps", "name": "Name 1", "status": "Status 1", "status_message": "StatusMessage 1"}, {"completed_time": "CompletedTime 2", "class_id": "iaas.WorkflowSteps", "object_type": "iaas.WorkflowSteps", "name": "Name 2", "status": "Status 2", "status_message": "StatusMessage 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamAccountPermissions(t *testing.T) {
	p := []models.IamAccountPermissions{}
	var d = &schema.ResourceData{}
	c := `{"AccountName":"AccountName %d","AccountStatus":"AccountStatus %d","ClassId":"iam.AccountPermissions","ObjectType":"iam.AccountPermissions","AccountIdentifier":"AccountIdentifier %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamAccountPermissions(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamAccountPermissions{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamAccountPermissions(p, d)
	expectedOp := []map[string]interface{}{{"account_name": "AccountName 1", "account_status": "AccountStatus 1", "class_id": "iam.AccountPermissions", "object_type": "iam.AccountPermissions", "account_identifier": "AccountIdentifier 1"}, {"account_name": "AccountName 2", "account_status": "AccountStatus 2", "class_id": "iam.AccountPermissions", "object_type": "iam.AccountPermissions", "account_identifier": "AccountIdentifier 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamApiKeyRelationship(t *testing.T) {
	p := []models.IamApiKeyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamApiKeyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamApiKeyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamApiKeyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamAppRegistrationRelationship(t *testing.T) {
	p := []models.IamAppRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamAppRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamAppRegistrationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamAppRegistrationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamDomainGroupRelationship(t *testing.T) {
	p := []models.IamDomainGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamDomainGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamDomainGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamDomainGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointPrivilegeRelationship(t *testing.T) {
	p := []models.IamEndPointPrivilegeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointPrivilegeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointPrivilegeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointPrivilegeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointRoleRelationship(t *testing.T) {
	p := []models.IamEndPointRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamEndPointUserRoleRelationship(t *testing.T) {
	p := []models.IamEndPointUserRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamEndPointUserRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamEndPointUserRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamEndPointUserRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamFeatureDefinition(t *testing.T) {
	p := []models.IamFeatureDefinition{}
	var d = &schema.ResourceData{}
	c := `{"Feature":"Feature %d","ObjectType":"iam.FeatureDefinition","ClassId":"iam.FeatureDefinition"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamFeatureDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamFeatureDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamFeatureDefinition(p, d)
	expectedOp := []map[string]interface{}{{"feature": "Feature 1", "object_type": "iam.FeatureDefinition", "class_id": "iam.FeatureDefinition"}, {"feature": "Feature 2", "object_type": "iam.FeatureDefinition", "class_id": "iam.FeatureDefinition"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamGroupPermissionToRoles(t *testing.T) {
	p := []models.IamGroupPermissionToRoles{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"iam.GroupPermissionToRoles","ClassId":"iam.GroupPermissionToRoles"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamGroupPermissionToRoles(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamGroupPermissionToRoles{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamGroupPermissionToRoles(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "iam.GroupPermissionToRoles", "class_id": "iam.GroupPermissionToRoles"}, {"object_type": "iam.GroupPermissionToRoles", "class_id": "iam.GroupPermissionToRoles"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIdpRelationship(t *testing.T) {
	p := []models.IamIdpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIdpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIdpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIdpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIdpReferenceRelationship(t *testing.T) {
	p := []models.IamIdpReferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIdpReferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIdpReferenceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIdpReferenceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamIpAddressRelationship(t *testing.T) {
	p := []models.IamIpAddressRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamIpAddressRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamIpAddressRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamIpAddressRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamLdapGroupRelationship(t *testing.T) {
	p := []models.IamLdapGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamLdapGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamLdapGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamLdapGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamLdapProviderRelationship(t *testing.T) {
	p := []models.IamLdapProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamLdapProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamLdapProviderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamLdapProviderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamOAuthTokenRelationship(t *testing.T) {
	p := []models.IamOAuthTokenRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamOAuthTokenRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamOAuthTokenRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamOAuthTokenRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPermissionRelationship(t *testing.T) {
	p := []models.IamPermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPermissionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPermissionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPermissionToRoles(t *testing.T) {
	p := []models.IamPermissionToRoles{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"iam.PermissionToRoles","ClassId":"iam.PermissionToRoles"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPermissionToRoles(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPermissionToRoles{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPermissionToRoles(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "iam.PermissionToRoles", "class_id": "iam.PermissionToRoles"}, {"object_type": "iam.PermissionToRoles", "class_id": "iam.PermissionToRoles"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPrivilegeRelationship(t *testing.T) {
	p := []models.IamPrivilegeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPrivilegeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPrivilegeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPrivilegeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamPrivilegeSetRelationship(t *testing.T) {
	p := []models.IamPrivilegeSetRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamPrivilegeSetRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamPrivilegeSetRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamPrivilegeSetRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamResourcePermissionRelationship(t *testing.T) {
	p := []models.IamResourcePermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamResourcePermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamResourcePermissionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamResourcePermissionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamResourceRolesRelationship(t *testing.T) {
	p := []models.IamResourceRolesRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamResourceRolesRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamResourceRolesRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamResourceRolesRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamRoleRelationship(t *testing.T) {
	p := []models.IamRoleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamRoleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamRoleRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamRoleRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamSessionRelationship(t *testing.T) {
	p := []models.IamSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamSessionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamSessionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserRelationship(t *testing.T) {
	p := []models.IamUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserGroupRelationship(t *testing.T) {
	p := []models.IamUserGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIamUserPreferenceRelationship(t *testing.T) {
	p := []models.IamUserPreferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListIamUserPreferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IamUserPreferenceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIamUserPreferenceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInfraBaseClusterRelationship(t *testing.T) {
	p := []models.InfraBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListInfraBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InfraBaseClusterRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInfraBaseClusterRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInfraMetaData(t *testing.T) {
	p := []models.InfraMetaData{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"infra.MetaData","ObjectType":"infra.MetaData","Value":"Value %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListInfraMetaData(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InfraMetaData{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInfraMetaData(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "infra.MetaData", "object_type": "infra.MetaData", "value": "Value 1", "name": "Name 1"}, {"class_id": "infra.MetaData", "object_type": "infra.MetaData", "value": "Value 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInventoryGenericInventoryRelationship(t *testing.T) {
	p := []models.InventoryGenericInventoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListInventoryGenericInventoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InventoryGenericInventoryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInventoryGenericInventoryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListInventoryGenericInventoryHolderRelationship(t *testing.T) {
	p := []models.InventoryGenericInventoryHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListInventoryGenericInventoryHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.InventoryGenericInventoryHolderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListInventoryGenericInventoryHolderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpLeaseRelationship(t *testing.T) {
	p := []models.IppoolIpLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpLeaseRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpLeaseRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpV4Block(t *testing.T) {
	p := []models.IppoolIpV4Block{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ObjectType":"ippool.IpV4Block","ClassId":"ippool.IpV4Block","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpV4Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpV4Block{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpV4Block(p, d)
	expectedOp := []map[string]interface{}{{"from": "From 1", "to": "To 1", "object_type": "ippool.IpV4Block", "class_id": "ippool.IpV4Block", "size": 32}, {"from": "From 2", "to": "To 2", "object_type": "ippool.IpV4Block", "class_id": "ippool.IpV4Block", "size": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolIpV6Block(t *testing.T) {
	p := []models.IppoolIpV6Block{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ippool.IpV6Block","ClassId":"ippool.IpV6Block","Size":32,"From":"From %d","To":"To %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolIpV6Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolIpV6Block{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolIpV6Block(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "ippool.IpV6Block", "class_id": "ippool.IpV6Block", "size": 32, "from": "From 1", "to": "To 1"}, {"object_type": "ippool.IpV6Block", "class_id": "ippool.IpV6Block", "size": 32, "from": "From 2", "to": "To 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolPoolRelationship(t *testing.T) {
	p := []models.IppoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolPoolRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolPoolRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolShadowBlockRelationship(t *testing.T) {
	p := []models.IppoolShadowBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolShadowBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolShadowBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolShadowBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIppoolShadowPoolRelationship(t *testing.T) {
	p := []models.IppoolShadowPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIppoolShadowPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IppoolShadowPoolRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIppoolShadowPoolRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIqnpoolBlockRelationship(t *testing.T) {
	p := []models.IqnpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListIqnpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IqnpoolBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIqnpoolBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListIqnpoolIqnSuffixBlock(t *testing.T) {
	p := []models.IqnpoolIqnSuffixBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":32,"Suffix":"Suffix %d","To":32,"ObjectType":"iqnpool.IqnSuffixBlock","ClassId":"iqnpool.IqnSuffixBlock","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenListIqnpoolIqnSuffixBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.IqnpoolIqnSuffixBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListIqnpoolIqnSuffixBlock(p, d)
	expectedOp := []map[string]interface{}{{"from": 32, "suffix": "Suffix 1", "to": 32, "object_type": "iqnpool.IqnSuffixBlock", "class_id": "iqnpool.IqnSuffixBlock", "size": 32}, {"from": 32, "suffix": "Suffix 2", "to": 32, "object_type": "iqnpool.IqnSuffixBlock", "class_id": "iqnpool.IqnSuffixBlock", "size": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAciCniProfileRelationship(t *testing.T) {
	p := []models.KubernetesAciCniProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAciCniProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAciCniProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAciCniProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAciCniTenantClusterAllocationRelationship(t *testing.T) {
	p := []models.KubernetesAciCniTenantClusterAllocationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAciCniTenantClusterAllocationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAciCniTenantClusterAllocationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAciCniTenantClusterAllocationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesAddon(t *testing.T) {
	p := []models.KubernetesAddon{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.Addon","Name":"Name %d","ClassId":"kubernetes.Addon"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesAddon(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesAddon{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesAddon(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "kubernetes.Addon", "name": "Name 1", "class_id": "kubernetes.Addon"}, {"object_type": "kubernetes.Addon", "name": "Name 2", "class_id": "kubernetes.Addon"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesClusterProfileRelationship(t *testing.T) {
	p := []models.KubernetesClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesClusterProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesClusterProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesConfigResultEntryRelationship(t *testing.T) {
	p := []models.KubernetesConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesEssentialAddon(t *testing.T) {
	p := []models.KubernetesEssentialAddon{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.EssentialAddon","ObjectType":"kubernetes.EssentialAddon","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesEssentialAddon(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesEssentialAddon{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesEssentialAddon(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.EssentialAddon", "object_type": "kubernetes.EssentialAddon", "name": "Name 1"}, {"class_id": "kubernetes.EssentialAddon", "object_type": "kubernetes.EssentialAddon", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesEthernet(t *testing.T) {
	p := []models.KubernetesEthernet{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.Ethernet","ProviderName":"ProviderName %d","Gateway":"Gateway %d","Name":"Name %d","Mtu":32,"ObjectType":"kubernetes.Ethernet"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesEthernet(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesEthernet{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesEthernet(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.Ethernet", "provider_name": "ProviderName 1", "gateway": "Gateway 1", "name": "Name 1", "mtu": 32, "object_type": "kubernetes.Ethernet"}, {"class_id": "kubernetes.Ethernet", "provider_name": "ProviderName 2", "gateway": "Gateway 2", "name": "Name 2", "mtu": 32, "object_type": "kubernetes.Ethernet"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeAddress(t *testing.T) {
	p := []models.KubernetesNodeAddress{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeAddress","ObjectType":"kubernetes.NodeAddress","Address":"Address %d","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeAddress(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeAddress{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeAddress(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.NodeAddress", "object_type": "kubernetes.NodeAddress", "address": "Address 1", "type": "Type 1"}, {"class_id": "kubernetes.NodeAddress", "object_type": "kubernetes.NodeAddress", "address": "Address 2", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupLabel(t *testing.T) {
	p := []models.KubernetesNodeGroupLabel{}
	var d = &schema.ResourceData{}
	c := `{"Key":"Key %d","Value":"Value %d","ClassId":"kubernetes.NodeGroupLabel","ObjectType":"kubernetes.NodeGroupLabel"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupLabel(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupLabel{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupLabel(p, d)
	expectedOp := []map[string]interface{}{{"key": "Key 1", "value": "Value 1", "class_id": "kubernetes.NodeGroupLabel", "object_type": "kubernetes.NodeGroupLabel"}, {"key": "Key 2", "value": "Value 2", "class_id": "kubernetes.NodeGroupLabel", "object_type": "kubernetes.NodeGroupLabel"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupProfileRelationship(t *testing.T) {
	p := []models.KubernetesNodeGroupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeGroupTaint(t *testing.T) {
	p := []models.KubernetesNodeGroupTaint{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeGroupTaint","ObjectType":"kubernetes.NodeGroupTaint","Value":"Value %d","Effect":"Effect %d","Key":"Key %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeGroupTaint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeGroupTaint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeGroupTaint(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.NodeGroupTaint", "object_type": "kubernetes.NodeGroupTaint", "value": "Value 1", "effect": "Effect 1", "key": "Key 1"}, {"class_id": "kubernetes.NodeGroupTaint", "object_type": "kubernetes.NodeGroupTaint", "value": "Value 2", "effect": "Effect 2", "key": "Key 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeProfileRelationship(t *testing.T) {
	p := []models.KubernetesNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesNodeStatus(t *testing.T) {
	p := []models.KubernetesNodeStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeStatus","ObjectType":"kubernetes.NodeStatus","Status":"Status %d","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesNodeStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesNodeStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesNodeStatus(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "kubernetes.NodeStatus", "object_type": "kubernetes.NodeStatus", "status": "Status 1", "type": "Type 1"}, {"class_id": "kubernetes.NodeStatus", "object_type": "kubernetes.NodeStatus", "status": "Status 2", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesTaint(t *testing.T) {
	p := []models.KubernetesTaint{}
	var d = &schema.ResourceData{}
	c := `{"Value":"Value %d","Effect":"Effect %d","Key":"Key %d","ClassId":"kubernetes.Taint","ObjectType":"kubernetes.Taint"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesTaint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesTaint{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesTaint(p, d)
	expectedOp := []map[string]interface{}{{"value": "Value 1", "effect": "Effect 1", "key": "Key 1", "class_id": "kubernetes.Taint", "object_type": "kubernetes.Taint"}, {"value": "Value 2", "effect": "Effect 2", "key": "Key 2", "class_id": "kubernetes.Taint", "object_type": "kubernetes.Taint"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListKubernetesVirtualMachineInfrastructureProviderRelationship(t *testing.T) {
	p := []models.KubernetesVirtualMachineInfrastructureProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.KubernetesVirtualMachineInfrastructureProviderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListKubernetesVirtualMachineInfrastructureProviderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListLicenseLicenseInfoRelationship(t *testing.T) {
	p := []models.LicenseLicenseInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListLicenseLicenseInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.LicenseLicenseInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListLicenseLicenseInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMacpoolBlock(t *testing.T) {
	p := []models.MacpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"To":"To %d","ClassId":"macpool.Block","ObjectType":"macpool.Block","Size":32,"From":"From %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMacpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MacpoolBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMacpoolBlock(p, d)
	expectedOp := []map[string]interface{}{{"to": "To 1", "class_id": "macpool.Block", "object_type": "macpool.Block", "size": 32, "from": "From 1"}, {"to": "To 2", "class_id": "macpool.Block", "object_type": "macpool.Block", "size": 32, "from": "From 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMacpoolIdBlockRelationship(t *testing.T) {
	p := []models.MacpoolIdBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMacpoolIdBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MacpoolIdBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMacpoolIdBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListManagementInterfaceRelationship(t *testing.T) {
	p := []models.ManagementInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListManagementInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ManagementInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListManagementInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryArrayRelationship(t *testing.T) {
	p := []models.MemoryArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryArrayRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryArrayRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryGoal(t *testing.T) {
	p := []models.MemoryPersistentMemoryGoal{}
	var d = &schema.ResourceData{}
	c := `{"PersistentMemoryType":"PersistentMemoryType %d","SocketId":"SocketId %d","ClassId":"memory.PersistentMemoryGoal","ObjectType":"memory.PersistentMemoryGoal","MemoryModePercentage":32}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryGoal(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryGoal{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryGoal(p, d)
	expectedOp := []map[string]interface{}{{"persistent_memory_type": "PersistentMemoryType 1", "socket_id": "SocketId 1", "class_id": "memory.PersistentMemoryGoal", "object_type": "memory.PersistentMemoryGoal", "memory_mode_percentage": 32}, {"persistent_memory_type": "PersistentMemoryType 2", "socket_id": "SocketId 2", "class_id": "memory.PersistentMemoryGoal", "object_type": "memory.PersistentMemoryGoal", "memory_mode_percentage": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryLogicalNamespace(t *testing.T) {
	p := []models.MemoryPersistentMemoryLogicalNamespace{}
	var d = &schema.ResourceData{}
	c := `{"Capacity":32,"Mode":"Mode %d","Name":"Name %d","SocketId":32,"SocketMemoryId":"SocketMemoryId %d","ClassId":"memory.PersistentMemoryLogicalNamespace","ObjectType":"memory.PersistentMemoryLogicalNamespace"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryLogicalNamespace(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryLogicalNamespace{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryLogicalNamespace(p, d)
	expectedOp := []map[string]interface{}{{"capacity": 32, "mode": "Mode 1", "name": "Name 1", "socket_id": 32, "socket_memory_id": "SocketMemoryId 1", "class_id": "memory.PersistentMemoryLogicalNamespace", "object_type": "memory.PersistentMemoryLogicalNamespace"}, {"capacity": 32, "mode": "Mode 2", "name": "Name 2", "socket_id": 32, "socket_memory_id": "SocketMemoryId 2", "class_id": "memory.PersistentMemoryLogicalNamespace", "object_type": "memory.PersistentMemoryLogicalNamespace"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryNamespaceRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryNamespaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryNamespaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryNamespaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryNamespaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryNamespaceConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryNamespaceConfigResultRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryNamespaceConfigResultRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryRegionRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryRegionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryRegionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryRegionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryRegionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryPersistentMemoryUnitRelationship(t *testing.T) {
	p := []models.MemoryPersistentMemoryUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryPersistentMemoryUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryPersistentMemoryUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryPersistentMemoryUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMemoryUnitRelationship(t *testing.T) {
	p := []models.MemoryUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMemoryUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MemoryUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMemoryUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaAccessPrivilege(t *testing.T) {
	p := []models.MetaAccessPrivilege{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"meta.AccessPrivilege","ObjectType":"meta.AccessPrivilege","Method":"Method %d","Privilege":"Privilege %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaAccessPrivilege(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaAccessPrivilege{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaAccessPrivilege(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "meta.AccessPrivilege", "object_type": "meta.AccessPrivilege", "method": "Method 1", "privilege": "Privilege 1"}, {"class_id": "meta.AccessPrivilege", "object_type": "meta.AccessPrivilege", "method": "Method 2", "privilege": "Privilege 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaDisplayNameDefinition(t *testing.T) {
	p := []models.MetaDisplayNameDefinition{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"meta.DisplayNameDefinition","ObjectType":"meta.DisplayNameDefinition","Format":"Format %d","IncludeAncestor":true,"Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaDisplayNameDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaDisplayNameDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaDisplayNameDefinition(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "meta.DisplayNameDefinition", "object_type": "meta.DisplayNameDefinition", "format": "Format 1", "include_ancestor": true, "name": "Name 1"}, {"class_id": "meta.DisplayNameDefinition", "object_type": "meta.DisplayNameDefinition", "format": "Format 2", "include_ancestor": true, "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaIdentityDefinition(t *testing.T) {
	p := []models.MetaIdentityDefinition{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"meta.IdentityDefinition","ObjectType":"meta.IdentityDefinition"}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaIdentityDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaIdentityDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaIdentityDefinition(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "meta.IdentityDefinition", "object_type": "meta.IdentityDefinition"}, {"class_id": "meta.IdentityDefinition", "object_type": "meta.IdentityDefinition"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaPropDefinition(t *testing.T) {
	p := []models.MetaPropDefinition{}
	var d = &schema.ResourceData{}
	c := `{"OpSecurity":"OpSecurity %d","Name":"Name %d","IsCollection":true,"ClassId":"meta.PropDefinition","Type":"Type %d","ApiAccess":"ApiAccess %d","IsComplexType":true,"ObjectType":"meta.PropDefinition","Kind":"Kind %d","SearchWeight":32.000000}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaPropDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaPropDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaPropDefinition(p, d)
	expectedOp := []map[string]interface{}{{"op_security": "OpSecurity 1", "name": "Name 1", "is_collection": true, "class_id": "meta.PropDefinition", "type": "Type 1", "api_access": "ApiAccess 1", "is_complex_type": true, "object_type": "meta.PropDefinition", "kind": "Kind 1", "search_weight": 32.000000}, {"op_security": "OpSecurity 2", "name": "Name 2", "is_collection": true, "class_id": "meta.PropDefinition", "type": "Type 2", "api_access": "ApiAccess 2", "is_complex_type": true, "object_type": "meta.PropDefinition", "kind": "Kind 2", "search_weight": 32.000000}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMetaRelationshipDefinition(t *testing.T) {
	p := []models.MetaRelationshipDefinition{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","PeerRelName":"PeerRelName %d","ApiAccess":"ApiAccess %d","PeerSync":true,"Type":"Type %d","ExportWithPeer":true,"ClassId":"meta.RelationshipDefinition","Collection":true,"ObjectType":"meta.RelationshipDefinition","Export":true}`

	//test when the response is empty
	ffOpEmpty := flattenListMetaRelationshipDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MetaRelationshipDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMetaRelationshipDefinition(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "peer_rel_name": "PeerRelName 1", "api_access": "ApiAccess 1", "peer_sync": true, "type": "Type 1", "export_with_peer": true, "class_id": "meta.RelationshipDefinition", "collection": true, "object_type": "meta.RelationshipDefinition", "export": true}, {"name": "Name 2", "peer_rel_name": "PeerRelName 2", "api_access": "ApiAccess 2", "peer_sync": true, "type": "Type 2", "export_with_peer": true, "class_id": "meta.RelationshipDefinition", "collection": true, "object_type": "meta.RelationshipDefinition", "export": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMoBaseMoRelationship(t *testing.T) {
	p := []models.MoBaseMoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMoBaseMoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MoBaseMoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMoBaseMoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMoMoRef(t *testing.T) {
	p := []models.MoMoRef{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListMoMoRef(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MoMoRef{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMoMoRef(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListMoTag(t *testing.T) {
	p := []models.MoTag{}
	var d = &schema.ResourceData{}
	c := `{"Key":"Key %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListMoTag(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.MoTag{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListMoTag(p, d)
	expectedOp := []map[string]interface{}{{"key": "Key 1", "value": "Value 1"}, {"key": "Key 2", "value": "Value 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNetworkElementRelationship(t *testing.T) {
	p := []models.NetworkElementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListNetworkElementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NetworkElementRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNetworkElementRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiaapiDetail(t *testing.T) {
	p := []models.NiaapiDetail{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niaapi.Detail","ObjectType":"niaapi.Detail","Chksum":"Chksum %d","Filename":"Filename %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiaapiDetail(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiaapiDetail{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiaapiDetail(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "niaapi.Detail", "object_type": "niaapi.Detail", "chksum": "Chksum 1", "filename": "Filename 1", "name": "Name 1"}, {"class_id": "niaapi.Detail", "object_type": "niaapi.Detail", "chksum": "Chksum 2", "filename": "Filename 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiaapiRevisionInfo(t *testing.T) {
	p := []models.NiaapiRevisionInfo{}
	var d = &schema.ResourceData{}
	c := `{"RevisionComment":"RevisionComment %d","RevisionNo":"RevisionNo %d","ClassId":"niaapi.RevisionInfo","ObjectType":"niaapi.RevisionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiaapiRevisionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiaapiRevisionInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiaapiRevisionInfo(p, d)
	expectedOp := []map[string]interface{}{{"revision_comment": "RevisionComment 1", "revision_no": "RevisionNo 1", "class_id": "niaapi.RevisionInfo", "object_type": "niaapi.RevisionInfo"}, {"revision_comment": "RevisionComment 2", "revision_no": "RevisionNo 2", "class_id": "niaapi.RevisionInfo", "object_type": "niaapi.RevisionInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiatelemetryInterfaceElement(t *testing.T) {
	p := []models.NiatelemetryInterfaceElement{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.InterfaceElement","ObjectType":"niatelemetry.InterfaceElement","XcvrPresent":"XcvrPresent %d","Name":"Name %d","OperState":"OperState %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListNiatelemetryInterfaceElement(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiatelemetryInterfaceElement{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiatelemetryInterfaceElement(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "niatelemetry.InterfaceElement", "object_type": "niatelemetry.InterfaceElement", "xcvr_present": "XcvrPresent 1", "name": "Name 1", "oper_state": "OperState 1"}, {"class_id": "niatelemetry.InterfaceElement", "object_type": "niatelemetry.InterfaceElement", "xcvr_present": "XcvrPresent 2", "name": "Name 2", "oper_state": "OperState 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNiatelemetryLogicalLink(t *testing.T) {
	p := []models.NiatelemetryLogicalLink{}
	var d = &schema.ResourceData{}
	c := `{"LinkAddr2":"LinkAddr2 %d","DbId":32,"ObjectType":"niatelemetry.LogicalLink","LinkType":"LinkType %d","LinkAddr1":"LinkAddr1 %d","LinkState":"LinkState %d","ClassId":"niatelemetry.LogicalLink","Uptime":"Uptime %d","IsPresent":true}`

	//test when the response is empty
	ffOpEmpty := flattenListNiatelemetryLogicalLink(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NiatelemetryLogicalLink{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNiatelemetryLogicalLink(p, d)
	expectedOp := []map[string]interface{}{{"link_addr2": "LinkAddr2 1", "db_id": 32, "object_type": "niatelemetry.LogicalLink", "link_type": "LinkType 1", "link_addr1": "LinkAddr1 1", "link_state": "LinkState 1", "class_id": "niatelemetry.LogicalLink", "uptime": "Uptime 1", "is_present": true}, {"link_addr2": "LinkAddr2 2", "db_id": 32, "object_type": "niatelemetry.LogicalLink", "link_type": "LinkType 2", "link_addr1": "LinkAddr1 2", "link_state": "LinkState 2", "class_id": "niatelemetry.LogicalLink", "uptime": "Uptime 2", "is_present": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNotificationAbstractCondition(t *testing.T) {
	p := []models.NotificationAbstractCondition{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"notification.AbstractCondition","ClassId":"notification.AbstractCondition"}`

	//test when the response is empty
	ffOpEmpty := flattenListNotificationAbstractCondition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NotificationAbstractCondition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNotificationAbstractCondition(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "notification.AbstractCondition", "class_id": "notification.AbstractCondition"}, {"object_type": "notification.AbstractCondition", "class_id": "notification.AbstractCondition"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNotificationAction(t *testing.T) {
	p := []models.NotificationAction{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"notification.Action","ObjectType":"notification.Action"}`

	//test when the response is empty
	ffOpEmpty := flattenListNotificationAction(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NotificationAction{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNotificationAction(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "notification.Action", "object_type": "notification.Action"}, {"class_id": "notification.Action", "object_type": "notification.Action"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListNtpAuthNtpServer(t *testing.T) {
	p := []models.NtpAuthNtpServer{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ntp.AuthNtpServer","SymKeyId":32,"SymKeyValue":"SymKeyValue %d","KeyType":"KeyType %d","ServerName":"ServerName %d","ClassId":"ntp.AuthNtpServer"}`

	//test when the response is empty
	ffOpEmpty := flattenListNtpAuthNtpServer(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.NtpAuthNtpServer{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListNtpAuthNtpServer(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "ntp.AuthNtpServer", "sym_key_id": 32, "sym_key_value": "SymKeyValue 1", "key_type": "KeyType 1", "server_name": "ServerName 1", "class_id": "ntp.AuthNtpServer"}, {"object_type": "ntp.AuthNtpServer", "sym_key_id": 32, "sym_key_value": "SymKeyValue 2", "key_type": "KeyType 2", "server_name": "ServerName 2", "class_id": "ntp.AuthNtpServer"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremImagePackage(t *testing.T) {
	p := []models.OnpremImagePackage{}
	var d = &schema.ResourceData{}
	c := `{"FilePath":"FilePath %d","Filename":"Filename %d","FileSha":"FileSha %d","Name":"Name %d","FileSize":32,"ClassId":"onprem.ImagePackage","ObjectType":"onprem.ImagePackage","Version":"Version %d","PackageType":"PackageType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremImagePackage(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremImagePackage{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremImagePackage(p, d)
	expectedOp := []map[string]interface{}{{"file_path": "FilePath 1", "filename": "Filename 1", "file_sha": "FileSha 1", "name": "Name 1", "file_size": 32, "class_id": "onprem.ImagePackage", "object_type": "onprem.ImagePackage", "nr_version": "Version 1", "package_type": "PackageType 1"}, {"file_path": "FilePath 2", "filename": "Filename 2", "file_sha": "FileSha 2", "name": "Name 2", "file_size": 32, "class_id": "onprem.ImagePackage", "object_type": "onprem.ImagePackage", "nr_version": "Version 2", "package_type": "PackageType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremUpgradeNote(t *testing.T) {
	p := []models.OnpremUpgradeNote{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.UpgradeNote","ObjectType":"onprem.UpgradeNote","Message":"Message %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremUpgradeNote(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremUpgradeNote{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremUpgradeNote(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "onprem.UpgradeNote", "object_type": "onprem.UpgradeNote", "message": "Message 1"}, {"class_id": "onprem.UpgradeNote", "object_type": "onprem.UpgradeNote", "message": "Message 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOnpremUpgradePhase(t *testing.T) {
	p := []models.OnpremUpgradePhase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.UpgradePhase","Failed":true,"Message":"Message %d","Name":"Name %d","ElapsedTime":32,"ObjectType":"onprem.UpgradePhase"}`

	//test when the response is empty
	ffOpEmpty := flattenListOnpremUpgradePhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OnpremUpgradePhase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOnpremUpgradePhase(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "onprem.UpgradePhase", "failed": true, "message": "Message 1", "name": "Name 1", "elapsed_time": 32, "object_type": "onprem.UpgradePhase"}, {"class_id": "onprem.UpgradePhase", "failed": true, "message": "Message 2", "name": "Name 2", "elapsed_time": 32, "object_type": "onprem.UpgradePhase"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOprsKvpair(t *testing.T) {
	p := []models.OprsKvpair{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"oprs.Kvpair","Key":"Key %d","Value":"Value %d","ClassId":"oprs.Kvpair"}`

	//test when the response is empty
	ffOpEmpty := flattenListOprsKvpair(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OprsKvpair{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOprsKvpair(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "oprs.Kvpair", "key": "Key 1", "value": "Value 1", "class_id": "oprs.Kvpair"}, {"object_type": "oprs.Kvpair", "key": "Key 2", "value": "Value 2", "class_id": "oprs.Kvpair"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOrganizationOrganizationRelationship(t *testing.T) {
	p := []models.OrganizationOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListOrganizationOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OrganizationOrganizationRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOrganizationOrganizationRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsConfigurationFileRelationship(t *testing.T) {
	p := []models.OsConfigurationFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsConfigurationFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsConfigurationFileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsConfigurationFileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsDistributionRelationship(t *testing.T) {
	p := []models.OsDistributionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsDistributionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsDistributionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsDistributionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsPlaceHolder(t *testing.T) {
	p := []models.OsPlaceHolder{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.PlaceHolder","ObjectType":"os.PlaceHolder","IsValueSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenListOsPlaceHolder(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsPlaceHolder{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsPlaceHolder(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "os.PlaceHolder", "object_type": "os.PlaceHolder", "is_value_set": true}, {"class_id": "os.PlaceHolder", "object_type": "os.PlaceHolder", "is_value_set": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsServerConfig(t *testing.T) {
	p := []models.OsServerConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.ServerConfig","InstallTarget":"InstallTarget %d","SerialNumber":"SerialNumber %d","ObjectType":"os.ServerConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsServerConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsServerConfig{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsServerConfig(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "os.ServerConfig", "install_target": "InstallTarget 1", "serial_number": "SerialNumber 1", "object_type": "os.ServerConfig"}, {"class_id": "os.ServerConfig", "install_target": "InstallTarget 2", "serial_number": "SerialNumber 2", "object_type": "os.ServerConfig"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListOsValidationInformation(t *testing.T) {
	p := []models.OsValidationInformation{}
	var d = &schema.ResourceData{}
	c := `{"StepName":"StepName %d","ErrorMsg":"ErrorMsg %d","Status":"Status %d","ClassId":"os.ValidationInformation","ObjectType":"os.ValidationInformation"}`

	//test when the response is empty
	ffOpEmpty := flattenListOsValidationInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.OsValidationInformation{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListOsValidationInformation(p, d)
	expectedOp := []map[string]interface{}{{"step_name": "StepName 1", "error_msg": "ErrorMsg 1", "status": "Status 1", "class_id": "os.ValidationInformation", "object_type": "os.ValidationInformation"}, {"step_name": "StepName 2", "error_msg": "ErrorMsg 2", "status": "Status 2", "class_id": "os.ValidationInformation", "object_type": "os.ValidationInformation"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciCoprocessorCardRelationship(t *testing.T) {
	p := []models.PciCoprocessorCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciCoprocessorCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciCoprocessorCardRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciCoprocessorCardRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciDeviceRelationship(t *testing.T) {
	p := []models.PciDeviceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciDeviceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciDeviceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciDeviceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciLinkRelationship(t *testing.T) {
	p := []models.PciLinkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciLinkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciLinkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciLinkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPciSwitchRelationship(t *testing.T) {
	p := []models.PciSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPciSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PciSwitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPciSwitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyAbstractConfigProfileRelationship(t *testing.T) {
	p := []models.PolicyAbstractConfigProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyAbstractConfigProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyAbstractConfigProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyAbstractConfigProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyAbstractPolicyRelationship(t *testing.T) {
	p := []models.PolicyAbstractPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyAbstractPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyAbstractPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyAbstractPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPolicyinventoryJobInfo(t *testing.T) {
	p := []models.PolicyinventoryJobInfo{}
	var d = &schema.ResourceData{}
	c := `{"PolicyName":"PolicyName %d","ExecutionStatus":"ExecutionStatus %d","PolicyId":"PolicyId %d","ClassId":"policyinventory.JobInfo","ObjectType":"policyinventory.JobInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenListPolicyinventoryJobInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PolicyinventoryJobInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPolicyinventoryJobInfo(p, d)
	expectedOp := []map[string]interface{}{{"policy_name": "PolicyName 1", "execution_status": "ExecutionStatus 1", "policy_id": "PolicyId 1", "class_id": "policyinventory.JobInfo", "object_type": "policyinventory.JobInfo"}, {"policy_name": "PolicyName 2", "execution_status": "ExecutionStatus 2", "policy_id": "PolicyId 2", "class_id": "policyinventory.JobInfo", "object_type": "policyinventory.JobInfo"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortGroupRelationship(t *testing.T) {
	p := []models.PortGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortMacBindingRelationship(t *testing.T) {
	p := []models.PortMacBindingRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortMacBindingRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortMacBindingRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortMacBindingRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListPortSubGroupRelationship(t *testing.T) {
	p := []models.PortSubGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListPortSubGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.PortSubGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListPortSubGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListProcessorUnitRelationship(t *testing.T) {
	p := []models.ProcessorUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListProcessorUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ProcessorUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListProcessorUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRackUnitPersonalityRelationship(t *testing.T) {
	p := []models.RackUnitPersonalityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRackUnitPersonalityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RackUnitPersonalityRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRackUnitPersonalityRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecommendationPhysicalItemRelationship(t *testing.T) {
	p := []models.RecommendationPhysicalItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecommendationPhysicalItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecommendationPhysicalItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecommendationPhysicalItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecoveryBackupProfileRelationship(t *testing.T) {
	p := []models.RecoveryBackupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecoveryBackupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecoveryBackupProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecoveryBackupProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListRecoveryConfigResultEntryRelationship(t *testing.T) {
	p := []models.RecoveryConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListRecoveryConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.RecoveryConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListRecoveryConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourceGroupRelationship(t *testing.T) {
	p := []models.ResourceGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourceGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourceGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourceGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourcePerTypeCombinedSelector(t *testing.T) {
	p := []models.ResourcePerTypeCombinedSelector{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"resource.PerTypeCombinedSelector","ObjectType":"resource.PerTypeCombinedSelector","CombinedSelector":"CombinedSelector %d","EmptyFilter":true,"SelectorObjectType":"SelectorObjectType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourcePerTypeCombinedSelector(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourcePerTypeCombinedSelector{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourcePerTypeCombinedSelector(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "resource.PerTypeCombinedSelector", "object_type": "resource.PerTypeCombinedSelector", "combined_selector": "CombinedSelector 1", "empty_filter": true, "selector_object_type": "SelectorObjectType 1"}, {"class_id": "resource.PerTypeCombinedSelector", "object_type": "resource.PerTypeCombinedSelector", "combined_selector": "CombinedSelector 2", "empty_filter": true, "selector_object_type": "SelectorObjectType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListResourceSelector(t *testing.T) {
	p := []models.ResourceSelector{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"resource.Selector","ObjectType":"resource.Selector"}`

	//test when the response is empty
	ffOpEmpty := flattenListResourceSelector(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ResourceSelector{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListResourceSelector(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "resource.Selector", "object_type": "resource.Selector"}, {"selector": "Selector 2", "class_id": "resource.Selector", "object_type": "resource.Selector"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdcardPartition(t *testing.T) {
	p := []models.SdcardPartition{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"sdcard.Partition","ObjectType":"sdcard.Partition","Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdcardPartition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdcardPartition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdcardPartition(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "sdcard.Partition", "object_type": "sdcard.Partition", "type": "Type 1"}, {"class_id": "sdcard.Partition", "object_type": "sdcard.Partition", "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanNetworkConfigurationType(t *testing.T) {
	p := []models.SdwanNetworkConfigurationType{}
	var d = &schema.ResourceData{}
	c := `{"NetworkType":"NetworkType %d","PortGroup":"PortGroup %d","Vlan":32,"ClassId":"sdwan.NetworkConfigurationType","ObjectType":"sdwan.NetworkConfigurationType"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanNetworkConfigurationType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanNetworkConfigurationType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanNetworkConfigurationType(p, d)
	expectedOp := []map[string]interface{}{{"network_type": "NetworkType 1", "port_group": "PortGroup 1", "vlan": 32, "class_id": "sdwan.NetworkConfigurationType", "object_type": "sdwan.NetworkConfigurationType"}, {"network_type": "NetworkType 2", "port_group": "PortGroup 2", "vlan": 32, "class_id": "sdwan.NetworkConfigurationType", "object_type": "sdwan.NetworkConfigurationType"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanProfileRelationship(t *testing.T) {
	p := []models.SdwanProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanProfileRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanProfileRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanRouterNodeRelationship(t *testing.T) {
	p := []models.SdwanRouterNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanRouterNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanRouterNodeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanRouterNodeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSdwanTemplateInputsType(t *testing.T) {
	p := []models.SdwanTemplateInputsType{}
	var d = &schema.ResourceData{}
	c := `{"Template":"Template %d","Type":"Type %d","Editable":true,"Value":"Value %d","ClassId":"sdwan.TemplateInputsType","ObjectType":"sdwan.TemplateInputsType","Key":"Key %d","Required":true,"Title":"Title %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSdwanTemplateInputsType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SdwanTemplateInputsType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSdwanTemplateInputsType(p, d)
	expectedOp := []map[string]interface{}{{"template": "Template 1", "type": "Type 1", "editable": true, "value": "Value 1", "class_id": "sdwan.TemplateInputsType", "object_type": "sdwan.TemplateInputsType", "key": "Key 1", "required": true, "title": "Title 1"}, {"template": "Template 2", "type": "Type 2", "editable": true, "value": "Value 2", "class_id": "sdwan.TemplateInputsType", "object_type": "sdwan.TemplateInputsType", "key": "Key 2", "required": true, "title": "Title 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSecurityUnitRelationship(t *testing.T) {
	p := []models.SecurityUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSecurityUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SecurityUnitRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSecurityUnitRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListServerConfigChangeDetailRelationship(t *testing.T) {
	p := []models.ServerConfigChangeDetailRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListServerConfigChangeDetailRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ServerConfigChangeDetailRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListServerConfigChangeDetailRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListServerConfigResultEntryRelationship(t *testing.T) {
	p := []models.ServerConfigResultEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListServerConfigResultEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.ServerConfigResultEntryRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListServerConfigResultEntryRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSnmpTrap(t *testing.T) {
	p := []models.SnmpTrap{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"Version":"Version %d","Type":"Type %d","User":"User %d","ClassId":"snmp.Trap","Community":"Community %d","Destination":"Destination %d","Port":32,"ObjectType":"snmp.Trap"}`

	//test when the response is empty
	ffOpEmpty := flattenListSnmpTrap(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SnmpTrap{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSnmpTrap(p, d)
	expectedOp := []map[string]interface{}{{"enabled": true, "nr_version": "Version 1", "type": "Type 1", "user": "User 1", "class_id": "snmp.Trap", "community": "Community 1", "destination": "Destination 1", "port": 32, "object_type": "snmp.Trap"}, {"enabled": true, "nr_version": "Version 2", "type": "Type 2", "user": "User 2", "class_id": "snmp.Trap", "community": "Community 2", "destination": "Destination 2", "port": 32, "object_type": "snmp.Trap"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSnmpUser(t *testing.T) {
	p := []models.SnmpUser{}
	var d = &schema.ResourceData{}
	c := `{"IsPrivacyPasswordSet":true,"ObjectType":"snmp.User","Name":"Name %d","IsAuthPasswordSet":true,"PrivacyType":"PrivacyType %d","AuthType":"AuthType %d","ClassId":"snmp.User","SecurityLevel":"SecurityLevel %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSnmpUser(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SnmpUser{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSnmpUser(p, d)
	expectedOp := []map[string]interface{}{{"is_privacy_password_set": true, "object_type": "snmp.User", "name": "Name 1", "is_auth_password_set": true, "privacy_type": "PrivacyType 1", "auth_type": "AuthType 1", "class_id": "snmp.User", "security_level": "SecurityLevel 1"}, {"is_privacy_password_set": true, "object_type": "snmp.User", "name": "Name 2", "is_auth_password_set": true, "privacy_type": "PrivacyType 2", "auth_type": "AuthType 2", "class_id": "snmp.User", "security_level": "SecurityLevel 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwareHyperflexDistributableRelationship(t *testing.T) {
	p := []models.SoftwareHyperflexDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwareHyperflexDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwareHyperflexDistributableRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwareHyperflexDistributableRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwareUcsdDistributableRelationship(t *testing.T) {
	p := []models.SoftwareUcsdDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwareUcsdDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwareUcsdDistributableRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwareUcsdDistributableRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSoftwarerepositoryConstraintModels(t *testing.T) {
	p := []models.SoftwarerepositoryConstraintModels{}
	var d = &schema.ResourceData{}
	c := `{"PlatformRegex":"PlatformRegex %d","ClassId":"softwarerepository.ConstraintModels","ObjectType":"softwarerepository.ConstraintModels","MinVersion":"MinVersion %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSoftwarerepositoryConstraintModels(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SoftwarerepositoryConstraintModels{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSoftwarerepositoryConstraintModels(p, d)
	expectedOp := []map[string]interface{}{{"platform_regex": "PlatformRegex 1", "class_id": "softwarerepository.ConstraintModels", "object_type": "softwarerepository.ConstraintModels", "min_version": "MinVersion 1", "name": "Name 1"}, {"platform_regex": "PlatformRegex 2", "class_id": "softwarerepository.ConstraintModels", "object_type": "softwarerepository.ConstraintModels", "min_version": "MinVersion 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageBaseClusterRelationship(t *testing.T) {
	p := []models.StorageBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageBaseClusterRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageBaseClusterRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageBaseInitiator(t *testing.T) {
	p := []models.StorageBaseInitiator{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"storage.BaseInitiator","ObjectType":"storage.BaseInitiator","Iqn":"Iqn %d","Name":"Name %d","Type":"Type %d","Wwn":"Wwn %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageBaseInitiator(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageBaseInitiator{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageBaseInitiator(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "storage.BaseInitiator", "object_type": "storage.BaseInitiator", "iqn": "Iqn 1", "name": "Name 1", "type": "Type 1", "wwn": "Wwn 1"}, {"class_id": "storage.BaseInitiator", "object_type": "storage.BaseInitiator", "iqn": "Iqn 2", "name": "Name 2", "type": "Type 2", "wwn": "Wwn 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageControllerRelationship(t *testing.T) {
	p := []models.StorageControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageDiskGroupRelationship(t *testing.T) {
	p := []models.StorageDiskGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageDiskGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageDiskGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageDiskGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageDiskSlotRelationship(t *testing.T) {
	p := []models.StorageDiskSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageDiskSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageDiskSlotRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageDiskSlotRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageDriveGroupRelationship(t *testing.T) {
	p := []models.StorageDriveGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageDriveGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageDriveGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageDriveGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureRelationship(t *testing.T) {
	p := []models.StorageEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureDiskRelationship(t *testing.T) {
	p := []models.StorageEnclosureDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureDiskRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureDiskRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageEnclosureDiskSlotEpRelationship(t *testing.T) {
	p := []models.StorageEnclosureDiskSlotEpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageEnclosureDiskSlotEpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageEnclosureDiskSlotEpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageEnclosureDiskSlotEpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashControllerRelationship(t *testing.T) {
	p := []models.StorageFlexFlashControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashControllerPropsRelationship(t *testing.T) {
	p := []models.StorageFlexFlashControllerPropsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashControllerPropsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashControllerPropsRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashControllerPropsRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashPhysicalDriveRelationship(t *testing.T) {
	p := []models.StorageFlexFlashPhysicalDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashPhysicalDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashPhysicalDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashPhysicalDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexFlashVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageFlexFlashVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexFlashVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexFlashVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexFlashVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilControllerRelationship(t *testing.T) {
	p := []models.StorageFlexUtilControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilControllerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilControllerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilPhysicalDriveRelationship(t *testing.T) {
	p := []models.StorageFlexUtilPhysicalDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilPhysicalDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilPhysicalDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilPhysicalDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageFlexUtilVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageFlexUtilVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageFlexUtilVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageFlexUtilVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageFlexUtilVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHitachiParityGroupRelationship(t *testing.T) {
	p := []models.StorageHitachiParityGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHitachiParityGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHitachiParityGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHitachiParityGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHyperFlexStorageContainerRelationship(t *testing.T) {
	p := []models.StorageHyperFlexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHyperFlexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHyperFlexStorageContainerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHyperFlexStorageContainerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageHyperFlexVolumeRelationship(t *testing.T) {
	p := []models.StorageHyperFlexVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageHyperFlexVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageHyperFlexVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageHyperFlexVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageItemRelationship(t *testing.T) {
	p := []models.StorageItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageItemRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageItemRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppAggregateRelationship(t *testing.T) {
	p := []models.StorageNetAppAggregateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppAggregateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppAggregateRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppAggregateRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppExportPolicyRule(t *testing.T) {
	p := []models.StorageNetAppExportPolicyRule{}
	var d = &schema.ResourceData{}
	c := `{"Index":32,"User":"User %d","ClassId":"storage.NetAppExportPolicyRule","ObjectType":"storage.NetAppExportPolicyRule"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppExportPolicyRule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppExportPolicyRule{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppExportPolicyRule(p, d)
	expectedOp := []map[string]interface{}{{"index": 32, "user": "User 1", "class_id": "storage.NetAppExportPolicyRule", "object_type": "storage.NetAppExportPolicyRule"}, {"index": 32, "user": "User 2", "class_id": "storage.NetAppExportPolicyRule", "object_type": "storage.NetAppExportPolicyRule"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageNetAppInitiatorGroupRelationship(t *testing.T) {
	p := []models.StorageNetAppInitiatorGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageNetAppInitiatorGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageNetAppInitiatorGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageNetAppInitiatorGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskExtensionRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskExtensionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskExtensionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePhysicalDiskUsageRelationship(t *testing.T) {
	p := []models.StoragePhysicalDiskUsageRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePhysicalDiskUsageRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePhysicalDiskUsageRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePhysicalDiskUsageRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureHostRelationship(t *testing.T) {
	p := []models.StoragePureHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureHostGroupRelationship(t *testing.T) {
	p := []models.StoragePureHostGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureHostGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureHostGroupRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureHostGroupRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureReplicationBlackout(t *testing.T) {
	p := []models.StoragePureReplicationBlackout{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"storage.PureReplicationBlackout","End":"End %d","Start":"Start %d","ClassId":"storage.PureReplicationBlackout"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureReplicationBlackout(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureReplicationBlackout{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureReplicationBlackout(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "storage.PureReplicationBlackout", "end": "End 1", "start": "Start 1", "class_id": "storage.PureReplicationBlackout"}, {"object_type": "storage.PureReplicationBlackout", "end": "End 2", "start": "Start 2", "class_id": "storage.PureReplicationBlackout"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStoragePureVolumeRelationship(t *testing.T) {
	p := []models.StoragePureVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListStoragePureVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StoragePureVolumeRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStoragePureVolumeRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSasExpanderRelationship(t *testing.T) {
	p := []models.StorageSasExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSasExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSasExpanderRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSasExpanderRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSasPortRelationship(t *testing.T) {
	p := []models.StorageSasPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSasPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSasPortRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSasPortRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageSpanRelationship(t *testing.T) {
	p := []models.StorageSpanRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageSpanRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageSpanRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageSpanRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageStorageContainerHostMountStatus(t *testing.T) {
	p := []models.StorageStorageContainerHostMountStatus{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"storage.StorageContainerHostMountStatus","Mounted":true,"Reason":"Reason %d","Accessibility":"Accessibility %d","HostName":"HostName %d","ClassId":"storage.StorageContainerHostMountStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageStorageContainerHostMountStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageStorageContainerHostMountStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageStorageContainerHostMountStatus(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "storage.StorageContainerHostMountStatus", "mounted": true, "reason": "Reason 1", "accessibility": "Accessibility 1", "host_name": "HostName 1", "class_id": "storage.StorageContainerHostMountStatus"}, {"object_type": "storage.StorageContainerHostMountStatus", "mounted": true, "reason": "Reason 2", "accessibility": "Accessibility 2", "host_name": "HostName 2", "class_id": "storage.StorageContainerHostMountStatus"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVdMemberEpRelationship(t *testing.T) {
	p := []models.StorageVdMemberEpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVdMemberEpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVdMemberEpRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVdMemberEpRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveConfiguration(t *testing.T) {
	p := []models.StorageVirtualDriveConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Size":32,"ClassId":"storage.VirtualDriveConfiguration","ObjectType":"storage.VirtualDriveConfiguration","BootDrive":true,"ExpandToAvailable":true}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveConfiguration{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveConfiguration(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "size": 32, "class_id": "storage.VirtualDriveConfiguration", "object_type": "storage.VirtualDriveConfiguration", "boot_drive": true, "expand_to_available": true}, {"name": "Name 2", "size": 32, "class_id": "storage.VirtualDriveConfiguration", "object_type": "storage.VirtualDriveConfiguration", "boot_drive": true, "expand_to_available": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveContainerRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveContainerRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveContainerRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListStorageVirtualDriveExtensionRelationship(t *testing.T) {
	p := []models.StorageVirtualDriveExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListStorageVirtualDriveExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.StorageVirtualDriveExtensionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListStorageVirtualDriveExtensionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSyslogLocalClientBase(t *testing.T) {
	p := []models.SyslogLocalClientBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"syslog.LocalClientBase","ObjectType":"syslog.LocalClientBase","MinSeverity":"MinSeverity %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListSyslogLocalClientBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SyslogLocalClientBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSyslogLocalClientBase(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "syslog.LocalClientBase", "object_type": "syslog.LocalClientBase", "min_severity": "MinSeverity 1"}, {"class_id": "syslog.LocalClientBase", "object_type": "syslog.LocalClientBase", "min_severity": "MinSeverity 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListSyslogRemoteClientBase(t *testing.T) {
	p := []models.SyslogRemoteClientBase{}
	var d = &schema.ResourceData{}
	c := `{"Hostname":"Hostname %d","MinSeverity":"MinSeverity %d","ClassId":"syslog.RemoteClientBase","ObjectType":"syslog.RemoteClientBase","Port":32,"Protocol":"Protocol %d","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenListSyslogRemoteClientBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.SyslogRemoteClientBase{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListSyslogRemoteClientBase(p, d)
	expectedOp := []map[string]interface{}{{"hostname": "Hostname 1", "min_severity": "MinSeverity 1", "class_id": "syslog.RemoteClientBase", "object_type": "syslog.RemoteClientBase", "port": 32, "protocol": "Protocol 1", "enabled": true}, {"hostname": "Hostname 2", "min_severity": "MinSeverity 2", "class_id": "syslog.RemoteClientBase", "object_type": "syslog.RemoteClientBase", "port": 32, "protocol": "Protocol 2", "enabled": true}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamAction(t *testing.T) {
	p := []models.TamAction{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Type":"Type %d","AlertType":"AlertType %d","ClassId":"tam.Action","ObjectType":"tam.Action","OperationType":"OperationType %d","AffectedObjectType":"AffectedObjectType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamAction(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamAction{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamAction(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "type": "Type 1", "alert_type": "AlertType 1", "class_id": "tam.Action", "object_type": "tam.Action", "operation_type": "OperationType 1", "affected_object_type": "AffectedObjectType 1"}, {"name": "Name 2", "type": "Type 2", "alert_type": "AlertType 2", "class_id": "tam.Action", "object_type": "tam.Action", "operation_type": "OperationType 2", "affected_object_type": "AffectedObjectType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamApiDataSource(t *testing.T) {
	p := []models.TamApiDataSource{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"tam.ApiDataSource","ObjectType":"tam.ApiDataSource","Name":"Name %d","Type":"Type %d","MoType":"MoType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamApiDataSource(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamApiDataSource{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamApiDataSource(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "tam.ApiDataSource", "object_type": "tam.ApiDataSource", "name": "Name 1", "type": "Type 1", "mo_type": "MoType 1"}, {"class_id": "tam.ApiDataSource", "object_type": "tam.ApiDataSource", "name": "Name 2", "type": "Type 2", "mo_type": "MoType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTamS3DataSource(t *testing.T) {
	p := []models.TamS3DataSource{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Type":"Type %d","ClassId":"tam.S3DataSource","S3Path":"S3Path %d","ObjectType":"tam.S3DataSource"}`

	//test when the response is empty
	ffOpEmpty := flattenListTamS3DataSource(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TamS3DataSource{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTamS3DataSource(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "type": "Type 1", "class_id": "tam.S3DataSource", "s3_path": "S3Path 1", "object_type": "tam.S3DataSource"}, {"name": "Name 2", "type": "Type 2", "class_id": "tam.S3DataSource", "s3_path": "S3Path 2", "object_type": "tam.S3DataSource"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTerraformCloudResource(t *testing.T) {
	p := []models.TerraformCloudResource{}
	var d = &schema.ResourceData{}
	c := `{"ResourceId":"ResourceId %d","CurrentStatus":"CurrentStatus %d","ObjectType":"terraform.CloudResource","ClassId":"terraform.CloudResource","DesiredStatus":"DesiredStatus %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListTerraformCloudResource(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TerraformCloudResource{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTerraformCloudResource(p, d)
	expectedOp := []map[string]interface{}{{"resource_id": "ResourceId 1", "current_status": "CurrentStatus 1", "object_type": "terraform.CloudResource", "class_id": "terraform.CloudResource", "desired_status": "DesiredStatus 1"}, {"resource_id": "ResourceId 2", "current_status": "CurrentStatus 2", "object_type": "terraform.CloudResource", "class_id": "terraform.CloudResource", "desired_status": "DesiredStatus 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListTerraformRunstate(t *testing.T) {
	p := []models.TerraformRunstate{}
	var d = &schema.ResourceData{}
	c := `{"RunId":"RunId %d","StateFile":"StateFile %d","ClassId":"terraform.Runstate","ObjectType":"terraform.Runstate"}`

	//test when the response is empty
	ffOpEmpty := flattenListTerraformRunstate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.TerraformRunstate{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListTerraformRunstate(p, d)
	expectedOp := []map[string]interface{}{{"run_id": "RunId 1", "state_file": "StateFile 1", "class_id": "terraform.Runstate", "object_type": "terraform.Runstate"}, {"run_id": "RunId 2", "state_file": "StateFile 2", "class_id": "terraform.Runstate", "object_type": "terraform.Runstate"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUcsdConnectorPack(t *testing.T) {
	p := []models.UcsdConnectorPack{}
	var d = &schema.ResourceData{}
	c := `{"DownloadedVersion":"DownloadedVersion %d","ConnectorFeature":"ConnectorFeature %d","State":"State %d","Name":"Name %d","Version":"Version %d","ClassId":"ucsd.ConnectorPack","ObjectType":"ucsd.ConnectorPack"}`

	//test when the response is empty
	ffOpEmpty := flattenListUcsdConnectorPack(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UcsdConnectorPack{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUcsdConnectorPack(p, d)
	expectedOp := []map[string]interface{}{{"downloaded_version": "DownloadedVersion 1", "connector_feature": "ConnectorFeature 1", "state": "State 1", "name": "Name 1", "nr_version": "Version 1", "class_id": "ucsd.ConnectorPack", "object_type": "ucsd.ConnectorPack"}, {"downloaded_version": "DownloadedVersion 2", "connector_feature": "ConnectorFeature 2", "state": "State 2", "name": "Name 2", "nr_version": "Version 2", "class_id": "ucsd.ConnectorPack", "object_type": "ucsd.ConnectorPack"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUuidpoolBlockRelationship(t *testing.T) {
	p := []models.UuidpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListUuidpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UuidpoolBlockRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUuidpoolBlockRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListUuidpoolUuidBlock(t *testing.T) {
	p := []models.UuidpoolUuidBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ObjectType":"uuidpool.UuidBlock","ClassId":"uuidpool.UuidBlock","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenListUuidpoolUuidBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.UuidpoolUuidBlock{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListUuidpoolUuidBlock(p, d)
	expectedOp := []map[string]interface{}{{"from": "From 1", "to": "To 1", "object_type": "uuidpool.UuidBlock", "class_id": "uuidpool.UuidBlock", "size": 32}, {"from": "From 2", "to": "To 2", "object_type": "uuidpool.UuidBlock", "class_id": "uuidpool.UuidBlock", "size": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationBaseNetworkRelationship(t *testing.T) {
	p := []models.VirtualizationBaseNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationBaseNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationBaseNetworkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationBaseNetworkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationIweHostRelationship(t *testing.T) {
	p := []models.VirtualizationIweHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationIweHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationIweHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationIweHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationIweHostInterfaceRelationship(t *testing.T) {
	p := []models.VirtualizationIweHostInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationIweHostInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationIweHostInterfaceRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationIweHostInterfaceRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationIweHostVswitchRelationship(t *testing.T) {
	p := []models.VirtualizationIweHostVswitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationIweHostVswitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationIweHostVswitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationIweHostVswitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationNetworkInterface(t *testing.T) {
	p := []models.VirtualizationNetworkInterface{}
	var d = &schema.ResourceData{}
	c := `{"Ipv6Address":true,"IpForwardingEnabled":true,"Bridge":"Bridge %d","PublicIpAllocate":true,"DirectPathIo":true,"SubnetId":"SubnetId %d","ClassId":"virtualization.NetworkInterface","MacAddress":"MacAddress %d","PrivateIpAllocationMode":"PrivateIpAllocationMode %d","AdaptorType":"AdaptorType %d","ObjectType":"virtualization.NetworkInterface","ConnectAtPowerOn":true,"Name":"Name %d","Order":32,"NetworkId":"NetworkId %d","NicId":"NicId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationNetworkInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationNetworkInterface{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationNetworkInterface(p, d)
	expectedOp := []map[string]interface{}{{"ipv6_address": true, "ip_forwarding_enabled": true, "bridge": "Bridge 1", "public_ip_allocate": true, "direct_path_io": true, "subnet_id": "SubnetId 1", "class_id": "virtualization.NetworkInterface", "mac_address": "MacAddress 1", "private_ip_allocation_mode": "PrivateIpAllocationMode 1", "adaptor_type": "AdaptorType 1", "object_type": "virtualization.NetworkInterface", "connect_at_power_on": true, "name": "Name 1", "order": 32, "network_id": "NetworkId 1", "nic_id": "NicId 1"}, {"ipv6_address": true, "ip_forwarding_enabled": true, "bridge": "Bridge 2", "public_ip_allocate": true, "direct_path_io": true, "subnet_id": "SubnetId 2", "class_id": "virtualization.NetworkInterface", "mac_address": "MacAddress 2", "private_ip_allocation_mode": "PrivateIpAllocationMode 2", "adaptor_type": "AdaptorType 2", "object_type": "virtualization.NetworkInterface", "connect_at_power_on": true, "name": "Name 2", "order": 32, "network_id": "NetworkId 2", "nic_id": "NicId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationNetworkPort(t *testing.T) {
	p := []models.VirtualizationNetworkPort{}
	var d = &schema.ResourceData{}
	c := `{"Vlans":"Vlans %d","ClassId":"virtualization.NetworkPort","ObjectType":"virtualization.NetworkPort","Name":"Name %d","PortType":"PortType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationNetworkPort(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationNetworkPort{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationNetworkPort(p, d)
	expectedOp := []map[string]interface{}{{"vlans": "Vlans 1", "class_id": "virtualization.NetworkPort", "object_type": "virtualization.NetworkPort", "name": "Name 1", "port_type": "PortType 1"}, {"vlans": "Vlans 2", "class_id": "virtualization.NetworkPort", "object_type": "virtualization.NetworkPort", "name": "Name 2", "port_type": "PortType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVirtualMachineDisk(t *testing.T) {
	p := []models.VirtualizationVirtualMachineDisk{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.VirtualMachineDisk","Name":"Name %d","VirtualDiskReference":"VirtualDiskReference %d","ClassId":"virtualization.VirtualMachineDisk","Bus":"Bus %d","Order":32,"Type":"Type %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVirtualMachineDisk(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVirtualMachineDisk{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVirtualMachineDisk(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "virtualization.VirtualMachineDisk", "name": "Name 1", "virtual_disk_reference": "VirtualDiskReference 1", "class_id": "virtualization.VirtualMachineDisk", "bus": "Bus 1", "order": 32, "type": "Type 1"}, {"object_type": "virtualization.VirtualMachineDisk", "name": "Name 2", "virtual_disk_reference": "VirtualDiskReference 2", "class_id": "virtualization.VirtualMachineDisk", "bus": "Bus 2", "order": 32, "type": "Type 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmDisk(t *testing.T) {
	p := []models.VirtualizationVmDisk{}
	var d = &schema.ResourceData{}
	c := `{"BootOrder":32,"Bus":"Bus %d","Type":"Type %d","VirtualDiskReference":"VirtualDiskReference %d","ClassId":"virtualization.VmDisk","ObjectType":"virtualization.VmDisk","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmDisk(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmDisk{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmDisk(p, d)
	expectedOp := []map[string]interface{}{{"boot_order": 32, "bus": "Bus 1", "type": "Type 1", "virtual_disk_reference": "VirtualDiskReference 1", "class_id": "virtualization.VmDisk", "object_type": "virtualization.VmDisk", "name": "Name 1"}, {"boot_order": 32, "bus": "Bus 2", "type": "Type 2", "virtual_disk_reference": "VirtualDiskReference 2", "class_id": "virtualization.VmDisk", "object_type": "virtualization.VmDisk", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmInterface(t *testing.T) {
	p := []models.VirtualizationVmInterface{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Bridge":"Bridge %d","ClassId":"virtualization.VmInterface","ObjectType":"virtualization.VmInterface","MacAddress":"MacAddress %d","Model":"Model %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmInterface{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmInterface(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "bridge": "Bridge 1", "class_id": "virtualization.VmInterface", "object_type": "virtualization.VmInterface", "mac_address": "MacAddress 1", "model": "Model 1"}, {"name": "Name 2", "bridge": "Bridge 2", "class_id": "virtualization.VmInterface", "object_type": "virtualization.VmInterface", "mac_address": "MacAddress 2", "model": "Model 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDatastoreRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDatastoreRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDatastoreRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDatastoreRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDatastoreRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDistributedNetworkRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDistributedNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDistributedNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDistributedNetworkRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDistributedNetworkRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareDistributedSwitchRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareDistributedSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareDistributedSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareDistributedSwitchRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareDistributedSwitchRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareHostRelationship(t *testing.T) {
	p := []models.VirtualizationVmwareHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareHostRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareHostRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVirtualizationVmwareVlanRange(t *testing.T) {
	p := []models.VirtualizationVmwareVlanRange{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.VmwareVlanRange","VlanRangeEnd":32,"VlanRangeStart":32,"ClassId":"virtualization.VmwareVlanRange"}`

	//test when the response is empty
	ffOpEmpty := flattenListVirtualizationVmwareVlanRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VirtualizationVmwareVlanRange{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVirtualizationVmwareVlanRange(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "virtualization.VmwareVlanRange", "vlan_range_end": 32, "vlan_range_start": 32, "class_id": "virtualization.VmwareVlanRange"}, {"object_type": "virtualization.VmwareVlanRange", "vlan_range_end": 32, "vlan_range_start": 32, "class_id": "virtualization.VmwareVlanRange"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVmediaMapping(t *testing.T) {
	p := []models.VmediaMapping{}
	var d = &schema.ResourceData{}
	c := `{"HostName":"HostName %d","Username":"Username %d","RemoteFile":"RemoteFile %d","MountProtocol":"MountProtocol %d","ClassId":"vmedia.Mapping","ObjectType":"vmedia.Mapping","FileLocation":"FileLocation %d","AuthenticationProtocol":"AuthenticationProtocol %d","MountOptions":"MountOptions %d","RemotePath":"RemotePath %d","SanitizedFileLocation":"SanitizedFileLocation %d","IsPasswordSet":true,"VolumeName":"VolumeName %d","DeviceType":"DeviceType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVmediaMapping(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VmediaMapping{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVmediaMapping(p, d)
	expectedOp := []map[string]interface{}{{"host_name": "HostName 1", "username": "Username 1", "remote_file": "RemoteFile 1", "mount_protocol": "MountProtocol 1", "class_id": "vmedia.Mapping", "object_type": "vmedia.Mapping", "file_location": "FileLocation 1", "authentication_protocol": "AuthenticationProtocol 1", "mount_options": "MountOptions 1", "remote_path": "RemotePath 1", "sanitized_file_location": "SanitizedFileLocation 1", "is_password_set": true, "volume_name": "VolumeName 1", "device_type": "DeviceType 1"}, {"host_name": "HostName 2", "username": "Username 2", "remote_file": "RemoteFile 2", "mount_protocol": "MountProtocol 2", "class_id": "vmedia.Mapping", "object_type": "vmedia.Mapping", "file_location": "FileLocation 2", "authentication_protocol": "AuthenticationProtocol 2", "mount_options": "MountOptions 2", "remote_path": "RemotePath 2", "sanitized_file_location": "SanitizedFileLocation 2", "is_password_set": true, "volume_name": "VolumeName 2", "device_type": "DeviceType 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicEthIfRelationship(t *testing.T) {
	p := []models.VnicEthIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicEthIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicEthIfRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicEthIfRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicEthNetworkPolicyRelationship(t *testing.T) {
	p := []models.VnicEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicEthNetworkPolicyRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicEthNetworkPolicyRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicFcIfRelationship(t *testing.T) {
	p := []models.VnicFcIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicFcIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicFcIfRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicFcIfRelationship(p, d)
	expectedOp := []map[string]interface{}{{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}, {"selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListVnicVifStatus(t *testing.T) {
	p := []models.VnicVifStatus{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Reason":"Reason %d","Status":"Status %d","ClassId":"vnic.VifStatus","ObjectType":"vnic.VifStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenListVnicVifStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.VnicVifStatus{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListVnicVifStatus(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "reason": "Reason 1", "status": "Status 1", "class_id": "vnic.VifStatus", "object_type": "vnic.VifStatus"}, {"name": "Name 2", "reason": "Reason 2", "status": "Status 2", "class_id": "vnic.VifStatus", "object_type": "vnic.VifStatus"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowActionWorkflowDefinition(t *testing.T) {
	p := []models.WorkflowActionWorkflowDefinition{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","WorkflowDefinitionName":"WorkflowDefinitionName %d","ClassId":"workflow.ActionWorkflowDefinition","ObjectType":"workflow.ActionWorkflowDefinition","Version":32,"Description":"Description %d","CatalogMoid":"CatalogMoid %d","Label":"Label %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowActionWorkflowDefinition(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowActionWorkflowDefinition{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowActionWorkflowDefinition(p, d)
	expectedOp := []map[string]interface{}{{"name": "Name 1", "workflow_definition_name": "WorkflowDefinitionName 1", "class_id": "workflow.ActionWorkflowDefinition", "object_type": "workflow.ActionWorkflowDefinition", "nr_version": 32, "description": "Description 1", "catalog_moid": "CatalogMoid 1", "label": "Label 1"}, {"name": "Name 2", "workflow_definition_name": "WorkflowDefinitionName 2", "class_id": "workflow.ActionWorkflowDefinition", "object_type": "workflow.ActionWorkflowDefinition", "nr_version": 32, "description": "Description 2", "catalog_moid": "CatalogMoid 2", "label": "Label 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowApi(t *testing.T) {
	p := []models.WorkflowApi{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.Api","Description":"Description %d","Body":"Body %d","ErrorContentType":"ErrorContentType %d","AssetTargetMoid":"AssetTargetMoid %d","Timeout":32,"Name":"Name %d","ContentType":"ContentType %d","ClassId":"workflow.Api","SkipOnCondition":"SkipOnCondition %d","Label":"Label %d","StartDelay":32}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowApi(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowApi{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowApi(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.Api", "description": "Description 1", "body": "Body 1", "error_content_type": "ErrorContentType 1", "asset_target_moid": "AssetTargetMoid 1", "timeout": 32, "name": "Name 1", "content_type": "ContentType 1", "class_id": "workflow.Api", "skip_on_condition": "SkipOnCondition 1", "label": "Label 1", "start_delay": 32}, {"object_type": "workflow.Api", "description": "Description 2", "body": "Body 2", "error_content_type": "ErrorContentType 2", "asset_target_moid": "AssetTargetMoid 2", "timeout": 32, "name": "Name 2", "content_type": "ContentType 2", "class_id": "workflow.Api", "skip_on_condition": "SkipOnCondition 2", "label": "Label 2", "start_delay": 32}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowBaseDataType(t *testing.T) {
	p := []models.WorkflowBaseDataType{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.BaseDataType","Name":"Name %d","ClassId":"workflow.BaseDataType","Description":"Description %d","Required":true,"Label":"Label %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowBaseDataType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowBaseDataType{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowBaseDataType(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.BaseDataType", "name": "Name 1", "class_id": "workflow.BaseDataType", "description": "Description 1", "required": true, "label": "Label 1"}, {"object_type": "workflow.BaseDataType", "name": "Name 2", "class_id": "workflow.BaseDataType", "description": "Description 2", "required": true, "label": "Label 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowDynamicWorkflowActionTaskList(t *testing.T) {
	p := []models.WorkflowDynamicWorkflowActionTaskList{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.DynamicWorkflowActionTaskList","Action":"Action %d","ClassId":"workflow.DynamicWorkflowActionTaskList"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowDynamicWorkflowActionTaskList(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowDynamicWorkflowActionTaskList{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowDynamicWorkflowActionTaskList(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.DynamicWorkflowActionTaskList", "action": "Action 1", "class_id": "workflow.DynamicWorkflowActionTaskList"}, {"object_type": "workflow.DynamicWorkflowActionTaskList", "action": "Action 2", "class_id": "workflow.DynamicWorkflowActionTaskList"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowMessage(t *testing.T) {
	p := []models.WorkflowMessage{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.Message","ObjectType":"workflow.Message","Message":"Message %d","Severity":"Severity %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowMessage(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowMessage{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowMessage(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.Message", "object_type": "workflow.Message", "message": "Message 1", "severity": "Severity 1"}, {"class_id": "workflow.Message", "object_type": "workflow.Message", "message": "Message 2", "severity": "Severity 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowParameterSet(t *testing.T) {
	p := []models.WorkflowParameterSet{}
	var d = &schema.ResourceData{}
	c := `{"Value":"Value %d","ClassId":"workflow.ParameterSet","ObjectType":"workflow.ParameterSet","Condition":"Condition %d","ControlParameter":"ControlParameter %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowParameterSet(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowParameterSet{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowParameterSet(p, d)
	expectedOp := []map[string]interface{}{{"value": "Value 1", "class_id": "workflow.ParameterSet", "object_type": "workflow.ParameterSet", "condition": "Condition 1", "control_parameter": "ControlParameter 1", "name": "Name 1"}, {"value": "Value 2", "class_id": "workflow.ParameterSet", "object_type": "workflow.ParameterSet", "condition": "Condition 2", "control_parameter": "ControlParameter 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowRollbackTask(t *testing.T) {
	p := []models.WorkflowRollbackTask{}
	var d = &schema.ResourceData{}
	c := `{"CatalogMoid":"CatalogMoid %d","TaskMoid":"TaskMoid %d","Name":"Name %d","SkipCondition":"SkipCondition %d","Version":32,"Description":"Description %d","ObjectType":"workflow.RollbackTask","ClassId":"workflow.RollbackTask"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowRollbackTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowRollbackTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowRollbackTask(p, d)
	expectedOp := []map[string]interface{}{{"catalog_moid": "CatalogMoid 1", "task_moid": "TaskMoid 1", "name": "Name 1", "skip_condition": "SkipCondition 1", "nr_version": 32, "description": "Description 1", "object_type": "workflow.RollbackTask", "class_id": "workflow.RollbackTask"}, {"catalog_moid": "CatalogMoid 2", "task_moid": "TaskMoid 2", "name": "Name 2", "skip_condition": "SkipCondition 2", "nr_version": 32, "description": "Description 2", "object_type": "workflow.RollbackTask", "class_id": "workflow.RollbackTask"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowRollbackWorkflowTask(t *testing.T) {
	p := []models.WorkflowRollbackWorkflowTask{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.RollbackWorkflowTask","Name":"Name %d","RefName":"RefName %d","RollbackCompleted":true,"TaskPath":"TaskPath %d","ObjectType":"workflow.RollbackWorkflowTask","TaskInfoMoid":"TaskInfoMoid %d","Description":"Description %d","Status":"Status %d","RollbackTaskName":"RollbackTaskName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowRollbackWorkflowTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowRollbackWorkflowTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowRollbackWorkflowTask(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.RollbackWorkflowTask", "name": "Name 1", "ref_name": "RefName 1", "rollback_completed": true, "task_path": "TaskPath 1", "object_type": "workflow.RollbackWorkflowTask", "task_info_moid": "TaskInfoMoid 1", "description": "Description 1", "status": "Status 1", "rollback_task_name": "RollbackTaskName 1"}, {"class_id": "workflow.RollbackWorkflowTask", "name": "Name 2", "ref_name": "RefName 2", "rollback_completed": true, "task_path": "TaskPath 2", "object_type": "workflow.RollbackWorkflowTask", "task_info_moid": "TaskInfoMoid 2", "description": "Description 2", "status": "Status 2", "rollback_task_name": "RollbackTaskName 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowSolutionActionDefinitionRelationship(t *testing.T) {
	p := []models.WorkflowSolutionActionDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowSolutionActionDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowSolutionActionDefinitionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowSolutionActionDefinitionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskDefinitionRelationship(t *testing.T) {
	p := []models.WorkflowTaskDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskDefinitionRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskDefinitionRelationship(p, d)
	expectedOp := []map[string]interface{}{{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}, {"moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskInfoRelationship(t *testing.T) {
	p := []models.WorkflowTaskInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}, {"object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2", "class_id": "mo.MoRef"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowTaskRetryInfo(t *testing.T) {
	p := []models.WorkflowTaskRetryInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"workflow.TaskRetryInfo","ClassId":"workflow.TaskRetryInfo","Status":"Status %d","TaskInstId":"TaskInstId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowTaskRetryInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowTaskRetryInfo{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowTaskRetryInfo(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "workflow.TaskRetryInfo", "class_id": "workflow.TaskRetryInfo", "status": "Status 1", "task_inst_id": "TaskInstId 1"}, {"object_type": "workflow.TaskRetryInfo", "class_id": "workflow.TaskRetryInfo", "status": "Status 2", "task_inst_id": "TaskInstId 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowUiInputFilter(t *testing.T) {
	p := []models.WorkflowUiInputFilter{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.UiInputFilter","ObjectType":"workflow.UiInputFilter","Name":"Name %d","UserHelpMessage":"UserHelpMessage %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowUiInputFilter(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowUiInputFilter{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowUiInputFilter(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.UiInputFilter", "object_type": "workflow.UiInputFilter", "name": "Name 1", "user_help_message": "UserHelpMessage 1"}, {"class_id": "workflow.UiInputFilter", "object_type": "workflow.UiInputFilter", "name": "Name 2", "user_help_message": "UserHelpMessage 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowWorkflowInfoRelationship(t *testing.T) {
	p := []models.WorkflowWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowWorkflowInfoRelationship{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowWorkflowInfoRelationship(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}, {"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 2", "selector": "Selector 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListWorkflowWorkflowTask(t *testing.T) {
	p := []models.WorkflowWorkflowTask{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.WorkflowTask","ObjectType":"workflow.WorkflowTask","Description":"Description %d","Label":"Label %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListWorkflowWorkflowTask(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.WorkflowWorkflowTask{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListWorkflowWorkflowTask(p, d)
	expectedOp := []map[string]interface{}{{"class_id": "workflow.WorkflowTask", "object_type": "workflow.WorkflowTask", "description": "Description 1", "label": "Label 1", "name": "Name 1"}, {"class_id": "workflow.WorkflowTask", "object_type": "workflow.WorkflowTask", "description": "Description 2", "label": "Label 2", "name": "Name 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenListX509Certificate(t *testing.T) {
	p := []models.X509Certificate{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"x509.Certificate","PemCertificate":"PemCertificate %d","SignatureAlgorithm":"SignatureAlgorithm %d","ClassId":"x509.Certificate","Sha256Fingerprint":"Sha256Fingerprint %d"}`

	//test when the response is empty
	ffOpEmpty := flattenListX509Certificate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	for i := 1; i < 3; i++ {
		x := models.X509Certificate{}
		err := x.UnmarshalJSON([]byte(strings.Replace(c, "%d", fmt.Sprint(i), -1)))
		CheckError(t, err)
		p = append(p, x)
	}
	ffOp := flattenListX509Certificate(p, d)
	expectedOp := []map[string]interface{}{{"object_type": "x509.Certificate", "pem_certificate": "PemCertificate 1", "signature_algorithm": "SignatureAlgorithm 1", "class_id": "x509.Certificate", "sha256_fingerprint": "Sha256Fingerprint 1"}, {"object_type": "x509.Certificate", "pem_certificate": "PemCertificate 2", "signature_algorithm": "SignatureAlgorithm 2", "class_id": "x509.Certificate", "sha256_fingerprint": "Sha256Fingerprint 2"}}
	for i := 0; i < len(expectedOp); i++ {
		err := compareMaps(expectedOp[i], ffOp[i], t)
		CheckError(t, err)
	}
}
func TestFlattenMapAccessAddressType(t *testing.T) {
	p := models.AccessAddressType{}
	var d = &schema.ResourceData{}
	c := `{"EnableIpV4":true,"EnableIpV6":true,"ClassId":"access.AddressType","ObjectType":"access.AddressType"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAccessAddressType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAccessAddressType(p, d)[0]
	expectedOp := map[string]interface{}{"enable_ip_v4": true, "enable_ip_v6": true, "class_id": "access.AddressType", "object_type": "access.AddressType"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAdapterUnitRelationship(t *testing.T) {
	p := models.AdapterUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAdapterUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAdapterUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAdapterUnitExpanderRelationship(t *testing.T) {
	p := models.AdapterUnitExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAdapterUnitExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAdapterUnitExpanderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceCertRenewalPhase(t *testing.T) {
	p := models.ApplianceCertRenewalPhase{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"appliance.CertRenewalPhase","ClassId":"appliance.CertRenewalPhase","Failed":true,"Message":"Message %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceCertRenewalPhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceCertRenewalPhase(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "appliance.CertRenewalPhase", "class_id": "appliance.CertRenewalPhase", "failed": true, "message": "Message 1", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceDataExportPolicyRelationship(t *testing.T) {
	p := models.ApplianceDataExportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceDataExportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceDataExportPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceGroupStatusRelationship(t *testing.T) {
	p := models.ApplianceGroupStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceGroupStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceGroupStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceImageBundleRelationship(t *testing.T) {
	p := models.ApplianceImageBundleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceImageBundleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceImageBundleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceNodeInfoRelationship(t *testing.T) {
	p := models.ApplianceNodeInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceNodeInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceNodeInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceNodeStatusRelationship(t *testing.T) {
	p := models.ApplianceNodeStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceNodeStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceNodeStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceSystemInfoRelationship(t *testing.T) {
	p := models.ApplianceSystemInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceSystemInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceSystemInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapApplianceSystemStatusRelationship(t *testing.T) {
	p := models.ApplianceSystemStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapApplianceSystemStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapApplianceSystemStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetClaimSignature(t *testing.T) {
	p := models.AssetClaimSignature{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.ClaimSignature","ObjectType":"asset.ClaimSignature"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetClaimSignature(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetClaimSignature(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.ClaimSignature", "object_type": "asset.ClaimSignature"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetClusterMemberRelationship(t *testing.T) {
	p := models.AssetClusterMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetClusterMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetClusterMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetContractInformation(t *testing.T) {
	p := models.AssetContractInformation{}
	var d = &schema.ResourceData{}
	c := `{"ContractNumber":"ContractNumber %d","ClassId":"asset.ContractInformation","ObjectType":"asset.ContractInformation","LineStatus":"LineStatus %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetContractInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetContractInformation(p, d)[0]
	expectedOp := map[string]interface{}{"contract_number": "ContractNumber 1", "class_id": "asset.ContractInformation", "object_type": "asset.ContractInformation", "line_status": "LineStatus 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetCustomerInformation(t *testing.T) {
	p := models.AssetCustomerInformation{}
	var d = &schema.ResourceData{}
	c := `{"Id":"Id %d","Name":"Name %d","ObjectType":"asset.CustomerInformation","ClassId":"asset.CustomerInformation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetCustomerInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetCustomerInformation(p, d)[0]
	expectedOp := map[string]interface{}{"id": "Id 1", "name": "Name 1", "object_type": "asset.CustomerInformation", "class_id": "asset.CustomerInformation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentRelationship(t *testing.T) {
	p := models.AssetDeploymentRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentAlarmInfo(t *testing.T) {
	p := models.AssetDeploymentAlarmInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.DeploymentAlarmInfo","ObjectType":"asset.DeploymentAlarmInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentAlarmInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentAlarmInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.DeploymentAlarmInfo", "object_type": "asset.DeploymentAlarmInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentDeviceAlarmInfo(t *testing.T) {
	p := models.AssetDeploymentDeviceAlarmInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.DeploymentDeviceAlarmInfo","ObjectType":"asset.DeploymentDeviceAlarmInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentDeviceAlarmInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentDeviceAlarmInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.DeploymentDeviceAlarmInfo", "object_type": "asset.DeploymentDeviceAlarmInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeploymentDeviceInformation(t *testing.T) {
	p := models.AssetDeploymentDeviceInformation{}
	var d = &schema.ResourceData{}
	c := `{"InstanceId":"InstanceId %d","Description":"Description %d","OldInstanceId":"OldInstanceId %d","MlbProductId":32,"ApplicationName":"ApplicationName %d","ObjectType":"asset.DeploymentDeviceInformation","OldDeviceId":"OldDeviceId %d","OldDeviceStatusId":32,"ClassId":"asset.DeploymentDeviceInformation","MlbProductName":"MlbProductName %d","OldDeviceStatusDescription":"OldDeviceStatusDescription %d","ItemType":"ItemType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeploymentDeviceInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeploymentDeviceInformation(p, d)[0]
	expectedOp := map[string]interface{}{"instance_id": "InstanceId 1", "description": "Description 1", "old_instance_id": "OldInstanceId 1", "mlb_product_id": 32, "application_name": "ApplicationName 1", "object_type": "asset.DeploymentDeviceInformation", "old_device_id": "OldDeviceId 1", "old_device_status_id": 32, "class_id": "asset.DeploymentDeviceInformation", "mlb_product_name": "MlbProductName 1", "old_device_status_description": "OldDeviceStatusDescription 1", "item_type": "ItemType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceClaimRelationship(t *testing.T) {
	p := models.AssetDeviceClaimRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceClaimRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceClaimRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceConfigurationRelationship(t *testing.T) {
	p := models.AssetDeviceConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceConnectionRelationship(t *testing.T) {
	p := models.AssetDeviceConnectionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceConnectionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceConnectionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceContractInformationRelationship(t *testing.T) {
	p := models.AssetDeviceContractInformationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceContractInformationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceContractInformationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceInformation(t *testing.T) {
	p := models.AssetDeviceInformation{}
	var d = &schema.ResourceData{}
	c := `{"MlbProductName":"MlbProductName %d","UnitOfMeasure":"UnitOfMeasure %d","OldDeviceId":"OldDeviceId %d","ProductType":"ProductType %d","OldDeviceStatusDescription":"OldDeviceStatusDescription %d","ClassId":"asset.DeviceInformation","ObjectType":"asset.DeviceInformation","MlbProductId":32,"ApplicationName":"ApplicationName %d","OldInstanceId":"OldInstanceId %d","ProductFamily":"ProductFamily %d","InstanceId":"InstanceId %d","OldDeviceStatusId":32,"MlbOfferType":"MlbOfferType %d","ItemType":"ItemType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceInformation(p, d)[0]
	expectedOp := map[string]interface{}{"mlb_product_name": "MlbProductName 1", "unit_of_measure": "UnitOfMeasure 1", "old_device_id": "OldDeviceId 1", "product_type": "ProductType 1", "old_device_status_description": "OldDeviceStatusDescription 1", "class_id": "asset.DeviceInformation", "object_type": "asset.DeviceInformation", "mlb_product_id": 32, "application_name": "ApplicationName 1", "old_instance_id": "OldInstanceId 1", "product_family": "ProductFamily 1", "instance_id": "InstanceId 1", "old_device_status_id": 32, "mlb_offer_type": "MlbOfferType 1", "item_type": "ItemType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceRegistrationRelationship(t *testing.T) {
	p := models.AssetDeviceRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceRegistrationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetDeviceStatistics(t *testing.T) {
	p := models.AssetDeviceStatistics{}
	var d = &schema.ResourceData{}
	c := `{"ClusterReplicationFactor":32,"MembershipRatio":32.000000,"ClassId":"asset.DeviceStatistics","ClusterDeploymentType":"ClusterDeploymentType %d","ClusterName":"ClusterName %d","MemoryMirroringFactor":32.000000,"ObjectType":"asset.DeviceStatistics","Connected":32,"ClusterDeviceMoid":"ClusterDeviceMoid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetDeviceStatistics(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetDeviceStatistics(p, d)[0]
	expectedOp := map[string]interface{}{"cluster_replication_factor": 32, "membership_ratio": 32.000000, "class_id": "asset.DeviceStatistics", "cluster_deployment_type": "ClusterDeploymentType 1", "cluster_name": "ClusterName 1", "memory_mirroring_factor": 32.000000, "object_type": "asset.DeviceStatistics", "connected": 32, "cluster_device_moid": "ClusterDeviceMoid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetGlobalUltimate(t *testing.T) {
	p := models.AssetGlobalUltimate{}
	var d = &schema.ResourceData{}
	c := `{"Id":"Id %d","Name":"Name %d","ClassId":"asset.GlobalUltimate","ObjectType":"asset.GlobalUltimate"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetGlobalUltimate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetGlobalUltimate(p, d)[0]
	expectedOp := map[string]interface{}{"id": "Id 1", "name": "Name 1", "class_id": "asset.GlobalUltimate", "object_type": "asset.GlobalUltimate"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetProductInformation(t *testing.T) {
	p := models.AssetProductInformation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"asset.ProductInformation","Group":"Group %d","SubType":"SubType %d","Family":"Family %d","ObjectType":"asset.ProductInformation","Number":"Number %d","Description":"Description %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetProductInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetProductInformation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "asset.ProductInformation", "group": "Group 1", "sub_type": "SubType 1", "family": "Family 1", "object_type": "asset.ProductInformation", "number": "Number 1", "description": "Description 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSubscriptionRelationship(t *testing.T) {
	p := models.AssetSubscriptionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSubscriptionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSubscriptionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSubscriptionAccountRelationship(t *testing.T) {
	p := models.AssetSubscriptionAccountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSubscriptionAccountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSubscriptionAccountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetSudiInfo(t *testing.T) {
	p := models.AssetSudiInfo{}
	var d = &schema.ResourceData{}
	c := `{"Signature":"Signature %d","Status":"Status %d","Pid":"Pid %d","SerialNumber":"SerialNumber %d","ClassId":"asset.SudiInfo","ObjectType":"asset.SudiInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetSudiInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetSudiInfo(p, d)[0]
	expectedOp := map[string]interface{}{"signature": "Signature 1", "status": "Status 1", "pid": "Pid 1", "serial_number": "SerialNumber 1", "class_id": "asset.SudiInfo", "object_type": "asset.SudiInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapAssetTargetRelationship(t *testing.T) {
	p := models.AssetTargetRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapAssetTargetRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapAssetTargetRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosBootModeRelationship(t *testing.T) {
	p := models.BiosBootModeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosBootModeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosBootModeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosSystemBootOrderRelationship(t *testing.T) {
	p := models.BiosSystemBootOrderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosSystemBootOrderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosSystemBootOrderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosTokenSettingsRelationship(t *testing.T) {
	p := models.BiosTokenSettingsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosTokenSettingsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosTokenSettingsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosUnitRelationship(t *testing.T) {
	p := models.BiosUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBiosVfSelectMemoryRasConfigurationRelationship(t *testing.T) {
	p := models.BiosVfSelectMemoryRasConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBiosVfSelectMemoryRasConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBiosVfSelectMemoryRasConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBootDeviceBootModeRelationship(t *testing.T) {
	p := models.BootDeviceBootModeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBootDeviceBootModeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBootDeviceBootModeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBootDeviceBootSecurityRelationship(t *testing.T) {
	p := models.BootDeviceBootSecurityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBootDeviceBootSecurityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBootDeviceBootSecurityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBulkApiResult(t *testing.T) {
	p := models.BulkApiResult{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"bulk.ApiResult","ObjectType":"bulk.ApiResult","Status":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapBulkApiResult(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBulkApiResult(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "bulk.ApiResult", "object_type": "bulk.ApiResult", "status": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBulkExportRelationship(t *testing.T) {
	p := models.BulkExportRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBulkExportRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBulkExportRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBulkExportedItemRelationship(t *testing.T) {
	p := models.BulkExportedItemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBulkExportedItemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBulkExportedItemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapBulkRequestRelationship(t *testing.T) {
	p := models.BulkRequestRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapBulkRequestRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapBulkRequestRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchNetworkLimits(t *testing.T) {
	p := models.CapabilitySwitchNetworkLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumPrimaryVlan":32,"MaximumVlans":32,"MaximumVifs":32,"MaxCompressedPortVlanCount":32,"MinimumActiveFans":32,"MaximumIgmpGroups":32,"MaxUncompressedPortVlanCount":32,"ClassId":"capability.SwitchNetworkLimits","MaximumSecondaryVlan":32,"MaximumPortChannelMembers":32,"MaximumEthernetUplinkPorts":32,"MaximumFcPortChannels":32,"ObjectType":"capability.SwitchNetworkLimits","MaximumEthernetPortChannels":32,"MaximumFcPortChannelMembers":32,"MaximumSecondaryVlanPerPrimary":32,"MaximumActiveTrafficMonitoringSessions":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchNetworkLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchNetworkLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_primary_vlan": 32, "maximum_vlans": 32, "maximum_vifs": 32, "max_compressed_port_vlan_count": 32, "minimum_active_fans": 32, "maximum_igmp_groups": 32, "max_uncompressed_port_vlan_count": 32, "class_id": "capability.SwitchNetworkLimits", "maximum_secondary_vlan": 32, "maximum_port_channel_members": 32, "maximum_ethernet_uplink_ports": 32, "maximum_fc_port_channels": 32, "object_type": "capability.SwitchNetworkLimits", "maximum_ethernet_port_channels": 32, "maximum_fc_port_channel_members": 32, "maximum_secondary_vlan_per_primary": 32, "maximum_active_traffic_monitoring_sessions": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchStorageLimits(t *testing.T) {
	p := models.CapabilitySwitchStorageLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumUserZoneCount":32,"MaximumVirtualFcInterfaces":32,"MaximumVirtualFcInterfacesPerBladeServer":32,"MaximumVsans":32,"MaximumZoneCount":32,"ClassId":"capability.SwitchStorageLimits","ObjectType":"capability.SwitchStorageLimits"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchStorageLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchStorageLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_user_zone_count": 32, "maximum_virtual_fc_interfaces": 32, "maximum_virtual_fc_interfaces_per_blade_server": 32, "maximum_vsans": 32, "maximum_zone_count": 32, "class_id": "capability.SwitchStorageLimits", "object_type": "capability.SwitchStorageLimits"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCapabilitySwitchSystemLimits(t *testing.T) {
	p := models.CapabilitySwitchSystemLimits{}
	var d = &schema.ResourceData{}
	c := `{"MaximumServersPerDomain":32,"ClassId":"capability.SwitchSystemLimits","ObjectType":"capability.SwitchSystemLimits","MaximumChassisCount":32,"MaximumFexPerDomain":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapCapabilitySwitchSystemLimits(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCapabilitySwitchSystemLimits(p, d)[0]
	expectedOp := map[string]interface{}{"maximum_servers_per_domain": 32, "class_id": "capability.SwitchSystemLimits", "object_type": "capability.SwitchSystemLimits", "maximum_chassis_count": 32, "maximum_fex_per_domain": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCertificatemanagementCertificateBase(t *testing.T) {
	p := models.CertificatemanagementCertificateBase{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"IsPrivatekeySet":true,"ClassId":"certificatemanagement.CertificateBase","ObjectType":"certificatemanagement.CertificateBase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCertificatemanagementCertificateBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCertificatemanagementCertificateBase(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "is_privatekey_set": true, "class_id": "certificatemanagement.CertificateBase", "object_type": "certificatemanagement.CertificateBase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapChassisConfigResultRelationship(t *testing.T) {
	p := models.ChassisConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapChassisConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapChassisConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapChassisProfileRelationship(t *testing.T) {
	p := models.ChassisProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapChassisProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapChassisProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAvailabilityZone(t *testing.T) {
	p := models.CloudAvailabilityZone{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.AvailabilityZone","ObjectType":"cloud.AvailabilityZone","ZoneId":"ZoneId %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAvailabilityZone(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAvailabilityZone(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "cloud.AvailabilityZone", "object_type": "cloud.AvailabilityZone", "zone_id": "ZoneId 1", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAwsBillingUnitRelationship(t *testing.T) {
	p := models.CloudAwsBillingUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAwsBillingUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAwsBillingUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAwsKeyPairRelationship(t *testing.T) {
	p := models.CloudAwsKeyPairRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAwsKeyPairRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAwsKeyPairRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAwsOrganizationalUnitRelationship(t *testing.T) {
	p := models.CloudAwsOrganizationalUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAwsOrganizationalUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAwsOrganizationalUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAwsSubnetRelationship(t *testing.T) {
	p := models.CloudAwsSubnetRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAwsSubnetRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAwsSubnetRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudAwsVpcRelationship(t *testing.T) {
	p := models.CloudAwsVpcRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudAwsVpcRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudAwsVpcRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudBaseSkuRelationship(t *testing.T) {
	p := models.CloudBaseSkuRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudBaseSkuRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudBaseSkuRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudBillingUnit(t *testing.T) {
	p := models.CloudBillingUnit{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"cloud.BillingUnit","BillingId":"BillingId %d","Name":"Name %d","ClassId":"cloud.BillingUnit"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudBillingUnit(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudBillingUnit(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "cloud.BillingUnit", "billing_id": "BillingId 1", "name": "Name 1", "class_id": "cloud.BillingUnit"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudCloudRegion(t *testing.T) {
	p := models.CloudCloudRegion{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.CloudRegion","ObjectType":"cloud.CloudRegion","RegionId":"RegionId %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudCloudRegion(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudCloudRegion(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "cloud.CloudRegion", "object_type": "cloud.CloudRegion", "region_id": "RegionId 1", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudImageReference(t *testing.T) {
	p := models.CloudImageReference{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.ImageReference","ObjectType":"cloud.ImageReference","ImageId":"ImageId %d","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudImageReference(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudImageReference(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "cloud.ImageReference", "object_type": "cloud.ImageReference", "image_id": "ImageId 1", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudInstanceType(t *testing.T) {
	p := models.CloudInstanceType{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"cloud.InstanceType","Name":"Name %d","Cpus":32,"Memory":32,"Platform":"Platform %d","ClassId":"cloud.InstanceType","Architecture":"Architecture %d","CpuSpeed":32,"InstanceTypeId":"InstanceTypeId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudInstanceType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudInstanceType(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "cloud.InstanceType", "name": "Name 1", "cpus": 32, "memory": 32, "platform": "Platform 1", "class_id": "cloud.InstanceType", "architecture": "Architecture 1", "cpu_speed": 32, "instance_type_id": "InstanceTypeId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudNetworkInstanceAttachment(t *testing.T) {
	p := models.CloudNetworkInstanceAttachment{}
	var d = &schema.ResourceData{}
	c := `{"AutoDelete":true,"DeviceIndex":32,"InstanceId":"InstanceId %d","State":"State %d","ClassId":"cloud.NetworkInstanceAttachment","ObjectType":"cloud.NetworkInstanceAttachment"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudNetworkInstanceAttachment(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudNetworkInstanceAttachment(p, d)[0]
	expectedOp := map[string]interface{}{"auto_delete": true, "device_index": 32, "instance_id": "InstanceId 1", "state": "State 1", "class_id": "cloud.NetworkInstanceAttachment", "object_type": "cloud.NetworkInstanceAttachment"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudRegionsRelationship(t *testing.T) {
	p := models.CloudRegionsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudRegionsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudRegionsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudTfcOrganizationRelationship(t *testing.T) {
	p := models.CloudTfcOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudTfcOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudTfcOrganizationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudVolumeIopsInfo(t *testing.T) {
	p := models.CloudVolumeIopsInfo{}
	var d = &schema.ResourceData{}
	c := `{"IopsReadLimit":32,"IopsWriteLimit":32,"ThroughputReadLimit":32,"ThroughputWriteLimit":32,"ClassId":"cloud.VolumeIopsInfo","ObjectType":"cloud.VolumeIopsInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudVolumeIopsInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudVolumeIopsInfo(p, d)[0]
	expectedOp := map[string]interface{}{"iops_read_limit": 32, "iops_write_limit": 32, "throughput_read_limit": 32, "throughput_write_limit": 32, "class_id": "cloud.VolumeIopsInfo", "object_type": "cloud.VolumeIopsInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCloudVolumeType(t *testing.T) {
	p := models.CloudVolumeType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"cloud.VolumeType","ObjectType":"cloud.VolumeType","Name":"Name %d","VolumeTypeId":"VolumeTypeId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCloudVolumeType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCloudVolumeType(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "cloud.VolumeType", "object_type": "cloud.VolumeType", "name": "Name 1", "volume_type_id": "VolumeTypeId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommHttpProxyPolicyRelationship(t *testing.T) {
	p := models.CommHttpProxyPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommHttpProxyPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommHttpProxyPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommIpV4Interface(t *testing.T) {
	p := models.CommIpV4Interface{}
	var d = &schema.ResourceData{}
	c := `{"Gateway":"Gateway %d","IpAddress":"IpAddress %d","Netmask":"Netmask %d","ClassId":"comm.IpV4Interface","ObjectType":"comm.IpV4Interface"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommIpV4Interface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommIpV4Interface(p, d)[0]
	expectedOp := map[string]interface{}{"gateway": "Gateway 1", "ip_address": "IpAddress 1", "netmask": "Netmask 1", "class_id": "comm.IpV4Interface", "object_type": "comm.IpV4Interface"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCommIpV6Interface(t *testing.T) {
	p := models.CommIpV6Interface{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"comm.IpV6Interface","Gateway":"Gateway %d","IpAddress":"IpAddress %d","Prefix":"Prefix %d","ClassId":"comm.IpV6Interface"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCommIpV6Interface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCommIpV6Interface(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "comm.IpV6Interface", "gateway": "Gateway 1", "ip_address": "IpAddress 1", "prefix": "Prefix 1", "class_id": "comm.IpV6Interface"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeAlarmSummary(t *testing.T) {
	p := models.ComputeAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.AlarmSummary","ObjectType":"compute.AlarmSummary","Critical":32,"Warning":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.AlarmSummary", "object_type": "compute.AlarmSummary", "critical": 32, "warning": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeBaseClusterRelationship(t *testing.T) {
	p := models.ComputeBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeBaseClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeBladeRelationship(t *testing.T) {
	p := models.ComputeBladeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeBladeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeBladeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeBoardRelationship(t *testing.T) {
	p := models.ComputeBoardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeBoardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeBoardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputePersistentMemoryOperation(t *testing.T) {
	p := models.ComputePersistentMemoryOperation{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"compute.PersistentMemoryOperation","ClassId":"compute.PersistentMemoryOperation","AdminAction":"AdminAction %d","IsSecurePassphraseSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputePersistentMemoryOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputePersistentMemoryOperation(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "compute.PersistentMemoryOperation", "class_id": "compute.PersistentMemoryOperation", "admin_action": "AdminAction 1", "is_secure_passphrase_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputePhysicalRelationship(t *testing.T) {
	p := models.ComputePhysicalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputePhysicalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputePhysicalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputePhysicalSummaryRelationship(t *testing.T) {
	p := models.ComputePhysicalSummaryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputePhysicalSummaryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputePhysicalSummaryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeRackUnitRelationship(t *testing.T) {
	p := models.ComputeRackUnitRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeRackUnitRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeRackUnitRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeServerConfig(t *testing.T) {
	p := models.ComputeServerConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.ServerConfig","ObjectType":"compute.ServerConfig","AssetTag":"AssetTag %d","UserLabel":"UserLabel %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeServerConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeServerConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.ServerConfig", "object_type": "compute.ServerConfig", "asset_tag": "AssetTag 1", "user_label": "UserLabel 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStorageControllerOperation(t *testing.T) {
	p := models.ComputeStorageControllerOperation{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"compute.StorageControllerOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d","ClassId":"compute.StorageControllerOperation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStorageControllerOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStorageControllerOperation(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "compute.StorageControllerOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1", "class_id": "compute.StorageControllerOperation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStoragePhysicalDriveOperation(t *testing.T) {
	p := models.ComputeStoragePhysicalDriveOperation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.StoragePhysicalDriveOperation","ObjectType":"compute.StoragePhysicalDriveOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStoragePhysicalDriveOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStoragePhysicalDriveOperation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.StoragePhysicalDriveOperation", "object_type": "compute.StoragePhysicalDriveOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeStorageVirtualDriveOperation(t *testing.T) {
	p := models.ComputeStorageVirtualDriveOperation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"compute.StorageVirtualDriveOperation","ObjectType":"compute.StorageVirtualDriveOperation","AdminAction":"AdminAction %d","ControllerId":"ControllerId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeStorageVirtualDriveOperation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeStorageVirtualDriveOperation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "compute.StorageVirtualDriveOperation", "object_type": "compute.StorageVirtualDriveOperation", "admin_action": "AdminAction 1", "controller_id": "ControllerId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapComputeVmediaRelationship(t *testing.T) {
	p := models.ComputeVmediaRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapComputeVmediaRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapComputeVmediaRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCondAlarmSummary(t *testing.T) {
	p := models.CondAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"cond.AlarmSummary","ClassId":"cond.AlarmSummary","Warning":32,"Critical":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapCondAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCondAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "cond.AlarmSummary", "class_id": "cond.AlarmSummary", "warning": 32, "critical": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapCondHclStatusRelationship(t *testing.T) {
	p := models.CondHclStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapCondHclStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapCondHclStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConnectorFileChecksum(t *testing.T) {
	p := models.ConnectorFileChecksum{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"connector.FileChecksum","ObjectType":"connector.FileChecksum","HashAlgorithm":"HashAlgorithm %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConnectorFileChecksum(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConnectorFileChecksum(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "connector.FileChecksum", "object_type": "connector.FileChecksum", "hash_algorithm": "HashAlgorithm 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConnectorPlatformParamBase(t *testing.T) {
	p := models.ConnectorPlatformParamBase{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"connector.PlatformParamBase","ObjectType":"connector.PlatformParamBase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConnectorPlatformParamBase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConnectorPlatformParamBase(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "connector.PlatformParamBase", "object_type": "connector.PlatformParamBase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConvergedinfraHealthCheckDefinitionRelationship(t *testing.T) {
	p := models.ConvergedinfraHealthCheckDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapConvergedinfraHealthCheckDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConvergedinfraHealthCheckDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapConvergedinfraPodSummary(t *testing.T) {
	p := models.ConvergedinfraPodSummary{}
	var d = &schema.ResourceData{}
	c := `{"VmCount":32,"StorageAvailable":32,"StorageCapacity":32,"NodeCount":32,"StorageUtilization":32.000000,"ClassId":"convergedinfra.PodSummary","ObjectType":"convergedinfra.PodSummary","ActiveNodes":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapConvergedinfraPodSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapConvergedinfraPodSummary(p, d)[0]
	expectedOp := map[string]interface{}{"vm_count": 32, "storage_available": 32, "storage_capacity": 32, "node_count": 32, "storage_utilization": 32.000000, "class_id": "convergedinfra.PodSummary", "object_type": "convergedinfra.PodSummary", "active_nodes": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentBaseRelationship(t *testing.T) {
	p := models.EquipmentBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentChassisRelationship(t *testing.T) {
	p := models.EquipmentChassisRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentChassisRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentChassisRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentExpanderModuleRelationship(t *testing.T) {
	p := models.EquipmentExpanderModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentExpanderModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentExpanderModuleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFanControlRelationship(t *testing.T) {
	p := models.EquipmentFanControlRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFanControlRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFanControlRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFanModuleRelationship(t *testing.T) {
	p := models.EquipmentFanModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFanModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFanModuleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFexRelationship(t *testing.T) {
	p := models.EquipmentFexRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFexRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFexRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentFruRelationship(t *testing.T) {
	p := models.EquipmentFruRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentFruRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentFruRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentIoCardRelationship(t *testing.T) {
	p := models.EquipmentIoCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentIoCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentIoCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentIoCardBaseRelationship(t *testing.T) {
	p := models.EquipmentIoCardBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentIoCardBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentIoCardBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentLocatorLedRelationship(t *testing.T) {
	p := models.EquipmentLocatorLedRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentLocatorLedRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentLocatorLedRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentPhysicalIdentityRelationship(t *testing.T) {
	p := models.EquipmentPhysicalIdentityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentPhysicalIdentityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentPhysicalIdentityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentPsuControlRelationship(t *testing.T) {
	p := models.EquipmentPsuControlRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentPsuControlRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentPsuControlRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentRackEnclosureRelationship(t *testing.T) {
	p := models.EquipmentRackEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentRackEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentRackEnclosureRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentRackEnclosureSlotRelationship(t *testing.T) {
	p := models.EquipmentRackEnclosureSlotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentRackEnclosureSlotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentRackEnclosureSlotRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSharedIoModuleRelationship(t *testing.T) {
	p := models.EquipmentSharedIoModuleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSharedIoModuleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSharedIoModuleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSwitchCardRelationship(t *testing.T) {
	p := models.EquipmentSwitchCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSwitchCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSwitchCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEquipmentSystemIoControllerRelationship(t *testing.T) {
	p := models.EquipmentSystemIoControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEquipmentSystemIoControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEquipmentSystemIoControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEtherPhysicalPortRelationship(t *testing.T) {
	p := models.EtherPhysicalPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEtherPhysicalPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEtherPhysicalPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapEtherPhysicalPortBaseRelationship(t *testing.T) {
	p := models.EtherPhysicalPortBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapEtherPhysicalPortBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapEtherPhysicalPortBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricConfigResultRelationship(t *testing.T) {
	p := models.FabricConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkControlPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkGroupPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkGroupPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkGroupPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkGroupPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricEthNetworkPolicyRelationship(t *testing.T) {
	p := models.FabricEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricEthNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricFcNetworkPolicyRelationship(t *testing.T) {
	p := models.FabricFcNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricFcNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricFcNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricFlowControlPolicyRelationship(t *testing.T) {
	p := models.FabricFlowControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricFlowControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricFlowControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLinkAggregationPolicyRelationship(t *testing.T) {
	p := models.FabricLinkAggregationPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLinkAggregationPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLinkAggregationPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLinkControlPolicyRelationship(t *testing.T) {
	p := models.FabricLinkControlPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLinkControlPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLinkControlPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricLldpSettings(t *testing.T) {
	p := models.FabricLldpSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.LldpSettings","ObjectType":"fabric.LldpSettings","ReceiveEnabled":true,"TransmitEnabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricLldpSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricLldpSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.LldpSettings", "object_type": "fabric.LldpSettings", "receive_enabled": true, "transmit_enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricMacAgingSettings(t *testing.T) {
	p := models.FabricMacAgingSettings{}
	var d = &schema.ResourceData{}
	c := `{"MacAgingOption":"MacAgingOption %d","MacAgingTime":32,"ClassId":"fabric.MacAgingSettings","ObjectType":"fabric.MacAgingSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricMacAgingSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricMacAgingSettings(p, d)[0]
	expectedOp := map[string]interface{}{"mac_aging_option": "MacAgingOption 1", "mac_aging_time": 32, "class_id": "fabric.MacAgingSettings", "object_type": "fabric.MacAgingSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricMulticastPolicyRelationship(t *testing.T) {
	p := models.FabricMulticastPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricMulticastPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricMulticastPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricPortPolicyRelationship(t *testing.T) {
	p := models.FabricPortPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricPortPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricPortPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricSwitchClusterProfileRelationship(t *testing.T) {
	p := models.FabricSwitchClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricSwitchClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricSwitchClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricSwitchProfileRelationship(t *testing.T) {
	p := models.FabricSwitchProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricSwitchProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricSwitchProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricUdldGlobalSettings(t *testing.T) {
	p := models.FabricUdldGlobalSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.UdldGlobalSettings","ObjectType":"fabric.UdldGlobalSettings","MessageInterval":32,"RecoveryAction":"RecoveryAction %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricUdldGlobalSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricUdldGlobalSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.UdldGlobalSettings", "object_type": "fabric.UdldGlobalSettings", "message_interval": 32, "recovery_action": "RecoveryAction 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricUdldSettings(t *testing.T) {
	p := models.FabricUdldSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.UdldSettings","ObjectType":"fabric.UdldSettings","AdminState":"AdminState %d","Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricUdldSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricUdldSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.UdldSettings", "object_type": "fabric.UdldSettings", "admin_state": "AdminState 1", "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFabricVlanSettings(t *testing.T) {
	p := models.FabricVlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"fabric.VlanSettings","ObjectType":"fabric.VlanSettings","AllowedVlans":"AllowedVlans %d","NativeVlan":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapFabricVlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFabricVlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "fabric.VlanSettings", "object_type": "fabric.VlanSettings", "allowed_vlans": "AllowedVlans 1", "native_vlan": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolBlock(t *testing.T) {
	p := models.FcpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ClassId":"fcpool.Block","Size":32,"ObjectType":"fcpool.Block"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolBlock(p, d)[0]
	expectedOp := map[string]interface{}{"from": "From 1", "to": "To 1", "class_id": "fcpool.Block", "size": 32, "object_type": "fcpool.Block"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolFcBlockRelationship(t *testing.T) {
	p := models.FcpoolFcBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolFcBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolFcBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolLeaseRelationship(t *testing.T) {
	p := models.FcpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolPoolRelationship(t *testing.T) {
	p := models.FcpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolPoolMemberRelationship(t *testing.T) {
	p := models.FcpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFcpoolUniverseRelationship(t *testing.T) {
	p := models.FcpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFcpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFcpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareBaseDistributableRelationship(t *testing.T) {
	p := models.FirmwareBaseDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareBaseDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareBaseDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareDirectDownload(t *testing.T) {
	p := models.FirmwareDirectDownload{}
	var d = &schema.ResourceData{}
	c := `{"Upgradeoption":"Upgradeoption %d","Username":"Username %d","ObjectType":"firmware.DirectDownload","ImageSource":"ImageSource %d","IsPasswordSet":true,"ClassId":"firmware.DirectDownload"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareDirectDownload(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareDirectDownload(p, d)[0]
	expectedOp := map[string]interface{}{"upgradeoption": "Upgradeoption 1", "username": "Username 1", "object_type": "firmware.DirectDownload", "image_source": "ImageSource 1", "is_password_set": true, "class_id": "firmware.DirectDownload"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareDistributableRelationship(t *testing.T) {
	p := models.FirmwareDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareIncludeComponentListType(t *testing.T) {
	p := models.FirmwareIncludeComponentListType{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"firmware.IncludeComponentListType","ObjectType":"firmware.IncludeComponentListType"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareIncludeComponentListType(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareIncludeComponentListType(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "firmware.IncludeComponentListType", "object_type": "firmware.IncludeComponentListType"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareNetworkShare(t *testing.T) {
	p := models.FirmwareNetworkShare{}
	var d = &schema.ResourceData{}
	c := `{"Username":"Username %d","IsPasswordSet":true,"ClassId":"firmware.NetworkShare","Upgradeoption":"Upgradeoption %d","MapType":"MapType %d","ObjectType":"firmware.NetworkShare"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareNetworkShare(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareNetworkShare(p, d)[0]
	expectedOp := map[string]interface{}{"username": "Username 1", "is_password_set": true, "class_id": "firmware.NetworkShare", "upgradeoption": "Upgradeoption 1", "map_type": "MapType 1", "object_type": "firmware.NetworkShare"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareRunningFirmwareRelationship(t *testing.T) {
	p := models.FirmwareRunningFirmwareRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareRunningFirmwareRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareRunningFirmwareRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareServerConfigurationUtilityDistributableRelationship(t *testing.T) {
	p := models.FirmwareServerConfigurationUtilityDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareServerConfigurationUtilityDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeBaseRelationship(t *testing.T) {
	p := models.FirmwareUpgradeBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeImpactStatusRelationship(t *testing.T) {
	p := models.FirmwareUpgradeImpactStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeImpactStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeImpactStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapFirmwareUpgradeStatusRelationship(t *testing.T) {
	p := models.FirmwareUpgradeStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapFirmwareUpgradeStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapFirmwareUpgradeStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastCatalogRelationship(t *testing.T) {
	p := models.ForecastCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastDefinitionRelationship(t *testing.T) {
	p := models.ForecastDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastInstanceRelationship(t *testing.T) {
	p := models.ForecastInstanceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastInstanceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastInstanceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapForecastModel(t *testing.T) {
	p := models.ForecastModel{}
	var d = &schema.ResourceData{}
	c := `{"ModelType":"ModelType %d","Accuracy":32.000000,"ClassId":"forecast.Model","ObjectType":"forecast.Model"}`

	//test when the response is empty
	ffOpEmpty := flattenMapForecastModel(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapForecastModel(p, d)[0]
	expectedOp := map[string]interface{}{"model_type": "ModelType 1", "accuracy": 32.000000, "class_id": "forecast.Model", "object_type": "forecast.Model"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapGraphicsCardRelationship(t *testing.T) {
	p := models.GraphicsCardRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapGraphicsCardRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapGraphicsCardRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHclOperatingSystemRelationship(t *testing.T) {
	p := models.HclOperatingSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHclOperatingSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHclOperatingSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHclOperatingSystemVendorRelationship(t *testing.T) {
	p := models.HclOperatingSystemVendorRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHclOperatingSystemVendorRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHclOperatingSystemVendorRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAlarmSummary(t *testing.T) {
	p := models.HyperflexAlarmSummary{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.AlarmSummary","ObjectType":"hyperflex.AlarmSummary","Critical":32,"Warning":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAlarmSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAlarmSummary(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.AlarmSummary", "object_type": "hyperflex.AlarmSummary", "critical": 32, "warning": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAppCatalogRelationship(t *testing.T) {
	p := models.HyperflexAppCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAppCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAppCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAppSettingConstraint(t *testing.T) {
	p := models.HyperflexAppSettingConstraint{}
	var d = &schema.ResourceData{}
	c := `{"DeploymentType":"DeploymentType %d","HxdpVersion":"HxdpVersion %d","HypervisorType":"HypervisorType %d","MgmtPlatform":"MgmtPlatform %d","ServerModel":"ServerModel %d","ClassId":"hyperflex.AppSettingConstraint","ObjectType":"hyperflex.AppSettingConstraint"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAppSettingConstraint(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAppSettingConstraint(p, d)[0]
	expectedOp := map[string]interface{}{"deployment_type": "DeploymentType 1", "hxdp_version": "HxdpVersion 1", "hypervisor_type": "HypervisorType 1", "mgmt_platform": "MgmtPlatform 1", "server_model": "ServerModel 1", "class_id": "hyperflex.AppSettingConstraint", "object_type": "hyperflex.AppSettingConstraint"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexAutoSupportPolicyRelationship(t *testing.T) {
	p := models.HyperflexAutoSupportPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexAutoSupportPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexAutoSupportPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexBackupClusterRelationship(t *testing.T) {
	p := models.HyperflexBackupClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexBackupClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexBackupClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexBackupPolicySettings(t *testing.T) {
	p := models.HyperflexBackupPolicySettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.BackupPolicySettings","BackupDataStoreSizeUnit":"BackupDataStoreSizeUnit %d","BackupDataStoreName":"BackupDataStoreName %d","ReplicationIntervalInMinutes":32,"SnapshotRetentionCount":32,"ObjectType":"hyperflex.BackupPolicySettings","DataStoreEncryptionEnabled":true,"LocalSnapshotRetentionCount":32,"ReplicationPairNamePrefix":"ReplicationPairNamePrefix %d","BackupDataStoreSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexBackupPolicySettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexBackupPolicySettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.BackupPolicySettings", "backup_data_store_size_unit": "BackupDataStoreSizeUnit 1", "backup_data_store_name": "BackupDataStoreName 1", "replication_interval_in_minutes": 32, "snapshot_retention_count": 32, "object_type": "hyperflex.BackupPolicySettings", "data_store_encryption_enabled": true, "local_snapshot_retention_count": 32, "replication_pair_name_prefix": "ReplicationPairNamePrefix 1", "backup_data_store_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterRelationship(t *testing.T) {
	p := models.HyperflexClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterNetworkPolicyRelationship(t *testing.T) {
	p := models.HyperflexClusterNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterProfileRelationship(t *testing.T) {
	p := models.HyperflexClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexClusterStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexClusterStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexClusterStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexClusterStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexConfigResultRelationship(t *testing.T) {
	p := models.HyperflexConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexDataProtectionPeerRelationship(t *testing.T) {
	p := models.HyperflexDataProtectionPeerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexDataProtectionPeerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexDataProtectionPeerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexDatastoreStatisticRelationship(t *testing.T) {
	p := models.HyperflexDatastoreStatisticRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexDatastoreStatisticRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexDatastoreStatisticRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexEntityReference(t *testing.T) {
	p := models.HyperflexEntityReference{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.EntityReference","Id":"Id %d","Idtype":"Idtype %d","Name":"Name %d","Type":"Type %d","Confignum":32,"ClassId":"hyperflex.EntityReference"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexEntityReference(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexEntityReference(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "hyperflex.EntityReference", "id": "Id 1", "idtype": "Idtype 1", "name": "Name 1", "type": "Type 1", "confignum": 32, "class_id": "hyperflex.EntityReference"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexErrorStack(t *testing.T) {
	p := models.HyperflexErrorStack{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"hyperflex.ErrorStack","ClassId":"hyperflex.ErrorStack","Message":"Message %d","MessageId":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexErrorStack(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexErrorStack(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "hyperflex.ErrorStack", "class_id": "hyperflex.ErrorStack", "message": "Message 1", "message_id": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexExtFcStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexExtFcStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexExtFcStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexExtFcStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexExtIscsiStoragePolicyRelationship(t *testing.T) {
	p := models.HyperflexExtIscsiStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexExtIscsiStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexExtIscsiStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexFeatureLimitExternalRelationship(t *testing.T) {
	p := models.HyperflexFeatureLimitExternalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexFeatureLimitExternalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexFeatureLimitExternalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexFeatureLimitInternalRelationship(t *testing.T) {
	p := models.HyperflexFeatureLimitInternalRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexFeatureLimitInternalRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexFeatureLimitInternalRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthRelationship(t *testing.T) {
	p := models.HyperflexHealthRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthCheckDefinitionRelationship(t *testing.T) {
	p := models.HyperflexHealthCheckDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthCheckDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthCheckDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHealthCheckScriptInfo(t *testing.T) {
	p := models.HyperflexHealthCheckScriptInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HealthCheckScriptInfo","ScriptExecuteLocation":"ScriptExecuteLocation %d","Version":"Version %d","ObjectType":"hyperflex.HealthCheckScriptInfo","AggregateScriptName":"AggregateScriptName %d","HyperflexVersion":"HyperflexVersion %d","ScriptInput":"ScriptInput %d","ScriptName":"ScriptName %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHealthCheckScriptInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHealthCheckScriptInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HealthCheckScriptInfo", "script_execute_location": "ScriptExecuteLocation 1", "nr_version": "Version 1", "object_type": "hyperflex.HealthCheckScriptInfo", "aggregate_script_name": "AggregateScriptName 1", "hyperflex_version": "HyperflexVersion 1", "script_input": "ScriptInput 1", "script_name": "ScriptName 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxLicenseAuthorizationDetailsDt(t *testing.T) {
	p := models.HyperflexHxLicenseAuthorizationDetailsDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxLicenseAuthorizationDetailsDt","Status":"Status %d","CommunicationDeadlineDate":"CommunicationDeadlineDate %d","ObjectType":"hyperflex.HxLicenseAuthorizationDetailsDt","LastCommunicationAttemptDate":"LastCommunicationAttemptDate %d","NextCommunicationAttemptDate":"NextCommunicationAttemptDate %d","EvaluationPeriodExpiredAt":"EvaluationPeriodExpiredAt %d","EvaluationPeriodRemaining":"EvaluationPeriodRemaining %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxLicenseAuthorizationDetailsDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxLicenseAuthorizationDetailsDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxLicenseAuthorizationDetailsDt", "status": "Status 1", "communication_deadline_date": "CommunicationDeadlineDate 1", "object_type": "hyperflex.HxLicenseAuthorizationDetailsDt", "last_communication_attempt_date": "LastCommunicationAttemptDate 1", "next_communication_attempt_date": "NextCommunicationAttemptDate 1", "evaluation_period_expired_at": "EvaluationPeriodExpiredAt 1", "evaluation_period_remaining": "EvaluationPeriodRemaining 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxNetworkAddressDt(t *testing.T) {
	p := models.HyperflexHxNetworkAddressDt{}
	var d = &schema.ResourceData{}
	c := `{"Fqdn":"Fqdn %d","Ip":"Ip %d","ClassId":"hyperflex.HxNetworkAddressDt","ObjectType":"hyperflex.HxNetworkAddressDt","Address":"Address %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxNetworkAddressDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxNetworkAddressDt(p, d)[0]
	expectedOp := map[string]interface{}{"fqdn": "Fqdn 1", "ip": "Ip 1", "class_id": "hyperflex.HxNetworkAddressDt", "object_type": "hyperflex.HxNetworkAddressDt", "address": "Address 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxPlatformDatastoreConfigDt(t *testing.T) {
	p := models.HyperflexHxPlatformDatastoreConfigDt{}
	var d = &schema.ResourceData{}
	c := `{"UsageType":"UsageType %d","DataBlockSize":32,"ProvisionedCapacity":32,"ClassId":"hyperflex.HxPlatformDatastoreConfigDt","ObjectType":"hyperflex.HxPlatformDatastoreConfigDt","SystemDatastore":true,"Name":"Name %d","NumMirrors":32,"NumStripesForLargeFiles":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxPlatformDatastoreConfigDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxPlatformDatastoreConfigDt(p, d)[0]
	expectedOp := map[string]interface{}{"usage_type": "UsageType 1", "data_block_size": 32, "provisioned_capacity": 32, "class_id": "hyperflex.HxPlatformDatastoreConfigDt", "object_type": "hyperflex.HxPlatformDatastoreConfigDt", "system_datastore": true, "name": "Name 1", "num_mirrors": 32, "num_stripes_for_large_files": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxRegistrationDetailsDt(t *testing.T) {
	p := models.HyperflexHxRegistrationDetailsDt{}
	var d = &schema.ResourceData{}
	c := `{"RegistrationFailedReason":"RegistrationFailedReason %d","Status":"Status %d","OutOfComplianceStartDate":"OutOfComplianceStartDate %d","RegistrationExpireDate":"RegistrationExpireDate %d","LastRenewalAttemptDate":"LastRenewalAttemptDate %d","NextRenewalAttemptDate":"NextRenewalAttemptDate %d","ClassId":"hyperflex.HxRegistrationDetailsDt","ObjectType":"hyperflex.HxRegistrationDetailsDt","VirtualAccount":"VirtualAccount %d","SmartAccount":"SmartAccount %d","InitialRegistrationDate":"InitialRegistrationDate %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxRegistrationDetailsDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxRegistrationDetailsDt(p, d)[0]
	expectedOp := map[string]interface{}{"registration_failed_reason": "RegistrationFailedReason 1", "status": "Status 1", "out_of_compliance_start_date": "OutOfComplianceStartDate 1", "registration_expire_date": "RegistrationExpireDate 1", "last_renewal_attempt_date": "LastRenewalAttemptDate 1", "next_renewal_attempt_date": "NextRenewalAttemptDate 1", "class_id": "hyperflex.HxRegistrationDetailsDt", "object_type": "hyperflex.HxRegistrationDetailsDt", "virtual_account": "VirtualAccount 1", "smart_account": "SmartAccount 1", "initial_registration_date": "InitialRegistrationDate 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxResiliencyInfoDt(t *testing.T) {
	p := models.HyperflexHxResiliencyInfoDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxResiliencyInfoDt","SsdFailuresTolerable":32,"DataReplicationFactor":"DataReplicationFactor %d","NodeFailuresTolerable":32,"ObjectType":"hyperflex.HxResiliencyInfoDt","PolicyCompliance":"PolicyCompliance %d","HddFailuresTolerable":32,"ResiliencyState":"ResiliencyState %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxResiliencyInfoDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxResiliencyInfoDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxResiliencyInfoDt", "ssd_failures_tolerable": 32, "data_replication_factor": "DataReplicationFactor 1", "node_failures_tolerable": 32, "object_type": "hyperflex.HxResiliencyInfoDt", "policy_compliance": "PolicyCompliance 1", "hdd_failures_tolerable": 32, "resiliency_state": "ResiliencyState 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxSiteDt(t *testing.T) {
	p := models.HyperflexHxSiteDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxSiteDt","ObjectType":"hyperflex.HxSiteDt","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxSiteDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxSiteDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxSiteDt", "object_type": "hyperflex.HxSiteDt", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexHxUuIdDt(t *testing.T) {
	p := models.HyperflexHxUuIdDt{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.HxUuIdDt","ObjectType":"hyperflex.HxUuIdDt","Uuid":"Uuid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexHxUuIdDt(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexHxUuIdDt(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.HxUuIdDt", "object_type": "hyperflex.HxUuIdDt", "uuid": "Uuid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexIpAddrRange(t *testing.T) {
	p := models.HyperflexIpAddrRange{}
	var d = &schema.ResourceData{}
	c := `{"StartAddr":"StartAddr %d","ClassId":"hyperflex.IpAddrRange","ObjectType":"hyperflex.IpAddrRange","EndAddr":"EndAddr %d","Gateway":"Gateway %d","Netmask":"Netmask %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexIpAddrRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexIpAddrRange(p, d)[0]
	expectedOp := map[string]interface{}{"start_addr": "StartAddr 1", "class_id": "hyperflex.IpAddrRange", "object_type": "hyperflex.IpAddrRange", "end_addr": "EndAddr 1", "gateway": "Gateway 1", "netmask": "Netmask 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLicenseRelationship(t *testing.T) {
	p := models.HyperflexLicenseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLicenseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLicenseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLocalCredentialPolicyRelationship(t *testing.T) {
	p := models.HyperflexLocalCredentialPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLocalCredentialPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLocalCredentialPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexLogicalAvailabilityZone(t *testing.T) {
	p := models.HyperflexLogicalAvailabilityZone{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.LogicalAvailabilityZone","ObjectType":"hyperflex.LogicalAvailabilityZone","AutoConfig":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexLogicalAvailabilityZone(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexLogicalAvailabilityZone(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.LogicalAvailabilityZone", "object_type": "hyperflex.LogicalAvailabilityZone", "auto_config": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexMacAddrPrefixRange(t *testing.T) {
	p := models.HyperflexMacAddrPrefixRange{}
	var d = &schema.ResourceData{}
	c := `{"StartAddr":"StartAddr %d","ObjectType":"hyperflex.MacAddrPrefixRange","ClassId":"hyperflex.MacAddrPrefixRange","EndAddr":"EndAddr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexMacAddrPrefixRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexMacAddrPrefixRange(p, d)[0]
	expectedOp := map[string]interface{}{"start_addr": "StartAddr 1", "object_type": "hyperflex.MacAddrPrefixRange", "class_id": "hyperflex.MacAddrPrefixRange", "end_addr": "EndAddr 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNamedVlan(t *testing.T) {
	p := models.HyperflexNamedVlan{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.NamedVlan","ObjectType":"hyperflex.NamedVlan","VlanId":32,"Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNamedVlan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNamedVlan(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.NamedVlan", "object_type": "hyperflex.NamedVlan", "vlan_id": 32, "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNamedVsan(t *testing.T) {
	p := models.HyperflexNamedVsan{}
	var d = &schema.ResourceData{}
	c := `{"VsanId":32,"ClassId":"hyperflex.NamedVsan","ObjectType":"hyperflex.NamedVsan","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNamedVsan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNamedVsan(p, d)[0]
	expectedOp := map[string]interface{}{"vsan_id": 32, "class_id": "hyperflex.NamedVsan", "object_type": "hyperflex.NamedVsan", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNodeRelationship(t *testing.T) {
	p := models.HyperflexNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNodeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexNodeConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexNodeConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexNodeConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexNodeConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexProxySettingPolicyRelationship(t *testing.T) {
	p := models.HyperflexProxySettingPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexProxySettingPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexProxySettingPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexReplicationPeerInfo(t *testing.T) {
	p := models.HyperflexReplicationPeerInfo{}
	var d = &schema.ResourceData{}
	c := `{"Dcip":"Dcip %d","Mcip":"Mcip %d","ClassId":"hyperflex.ReplicationPeerInfo","ObjectType":"hyperflex.ReplicationPeerInfo","ReplCip":"ReplCip %d","StatusDetails":"StatusDetails %d","Status":"Status %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexReplicationPeerInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexReplicationPeerInfo(p, d)[0]
	expectedOp := map[string]interface{}{"dcip": "Dcip 1", "mcip": "Mcip 1", "class_id": "hyperflex.ReplicationPeerInfo", "object_type": "hyperflex.ReplicationPeerInfo", "repl_cip": "ReplCip 1", "status_details": "StatusDetails 1", "status": "Status 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexReplicationSchedule(t *testing.T) {
	p := models.HyperflexReplicationSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"hyperflex.ReplicationSchedule","ObjectType":"hyperflex.ReplicationSchedule","BackupInterval":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexReplicationSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexReplicationSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "hyperflex.ReplicationSchedule", "object_type": "hyperflex.ReplicationSchedule", "backup_interval": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexServerFirmwareVersionRelationship(t *testing.T) {
	p := models.HyperflexServerFirmwareVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexServerFirmwareVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexServerFirmwareVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexServerModelRelationship(t *testing.T) {
	p := models.HyperflexServerModelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexServerModelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexServerModelRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareDistributionEntryRelationship(t *testing.T) {
	p := models.HyperflexSoftwareDistributionEntryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareDistributionEntryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareDistributionEntryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareDistributionVersionRelationship(t *testing.T) {
	p := models.HyperflexSoftwareDistributionVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareDistributionVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareDistributionVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSoftwareVersionPolicyRelationship(t *testing.T) {
	p := models.HyperflexSoftwareVersionPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSoftwareVersionPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSoftwareVersionPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexStorageContainerRelationship(t *testing.T) {
	p := models.HyperflexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexStorageContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSummary(t *testing.T) {
	p := models.HyperflexSummary{}
	var d = &schema.ResourceData{}
	c := `{"FreeCapacity":32,"ActiveNodes":"ActiveNodes %d","Downtime":"Downtime %d","Uptime":"Uptime %d","ObjectType":"hyperflex.Summary","Name":"Name %d","DeduplicationSavings":32.000000,"ClusterAccessPolicy":"ClusterAccessPolicy %d","State":"State %d","CompressionSavings":32.000000,"TotalSavings":32.000000,"ResiliencyDetailsSize":32,"Boottime":32,"ClassId":"hyperflex.Summary","SpaceStatus":"SpaceStatus %d","Address":"Address %d","DataReplicationFactor":"DataReplicationFactor %d","UsedCapacity":32,"TotalCapacity":32,"DataReplicationCompliance":"DataReplicationCompliance %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSummary(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSummary(p, d)[0]
	expectedOp := map[string]interface{}{"free_capacity": 32, "active_nodes": "ActiveNodes 1", "downtime": "Downtime 1", "uptime": "Uptime 1", "object_type": "hyperflex.Summary", "name": "Name 1", "deduplication_savings": 32.000000, "cluster_access_policy": "ClusterAccessPolicy 1", "state": "State 1", "compression_savings": 32.000000, "total_savings": 32.000000, "resiliency_details_size": 32, "boottime": 32, "class_id": "hyperflex.Summary", "space_status": "SpaceStatus 1", "address": "Address 1", "data_replication_factor": "DataReplicationFactor 1", "used_capacity": 32, "total_capacity": 32, "data_replication_compliance": "DataReplicationCompliance 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexSysConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexSysConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexSysConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexSysConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexUcsmConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexUcsmConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexUcsmConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexUcsmConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVcenterConfigPolicyRelationship(t *testing.T) {
	p := models.HyperflexVcenterConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVcenterConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVcenterConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVirtualMachine(t *testing.T) {
	p := models.HyperflexVirtualMachine{}
	var d = &schema.ResourceData{}
	c := `{"StatusCode":"StatusCode %d","Uuid":"Uuid %d","ClassId":"hyperflex.VirtualMachine","ObjectType":"hyperflex.VirtualMachine"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVirtualMachine(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVirtualMachine(p, d)[0]
	expectedOp := map[string]interface{}{"status_code": "StatusCode 1", "uuid": "Uuid 1", "class_id": "hyperflex.VirtualMachine", "object_type": "hyperflex.VirtualMachine"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVmBackupInfoRelationship(t *testing.T) {
	p := models.HyperflexVmBackupInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVmBackupInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVmBackupInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexVmSnapshotInfoRelationship(t *testing.T) {
	p := models.HyperflexVmSnapshotInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexVmSnapshotInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexVmSnapshotInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapHyperflexWwxnPrefixRange(t *testing.T) {
	p := models.HyperflexWwxnPrefixRange{}
	var d = &schema.ResourceData{}
	c := `{"StartAddr":"StartAddr %d","ClassId":"hyperflex.WwxnPrefixRange","ObjectType":"hyperflex.WwxnPrefixRange","EndAddr":"EndAddr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapHyperflexWwxnPrefixRange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapHyperflexWwxnPrefixRange(p, d)[0]
	expectedOp := map[string]interface{}{"start_addr": "StartAddr 1", "class_id": "hyperflex.WwxnPrefixRange", "object_type": "hyperflex.WwxnPrefixRange", "end_addr": "EndAddr 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasLicenseInfoRelationship(t *testing.T) {
	p := models.IaasLicenseInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasLicenseInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasLicenseInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasUcsdInfoRelationship(t *testing.T) {
	p := models.IaasUcsdInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasUcsdInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasUcsdInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIaasUcsdManagedInfraRelationship(t *testing.T) {
	p := models.IaasUcsdManagedInfraRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIaasUcsdManagedInfraRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIaasUcsdManagedInfraRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamAccountRelationship(t *testing.T) {
	p := models.IamAccountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamAccountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamAccountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamAppRegistrationRelationship(t *testing.T) {
	p := models.IamAppRegistrationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamAppRegistrationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamAppRegistrationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamCertificateRelationship(t *testing.T) {
	p := models.IamCertificateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamCertificateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamCertificateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamCertificateRequestRelationship(t *testing.T) {
	p := models.IamCertificateRequestRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamCertificateRequestRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamCertificateRequestRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamClientMeta(t *testing.T) {
	p := models.IamClientMeta{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"iam.ClientMeta","ObjectType":"iam.ClientMeta","DeviceModel":"DeviceModel %d","UserAgent":"UserAgent %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamClientMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamClientMeta(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "iam.ClientMeta", "object_type": "iam.ClientMeta", "device_model": "DeviceModel 1", "user_agent": "UserAgent 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamDomainGroupRelationship(t *testing.T) {
	p := models.IamDomainGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamDomainGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamDomainGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointPasswordProperties(t *testing.T) {
	p := models.IamEndPointPasswordProperties{}
	var d = &schema.ResourceData{}
	c := `{"PasswordHistory":32,"EnforceStrongPassword":true,"ForceSendPassword":true,"EnablePasswordExpiry":true,"ClassId":"iam.EndPointPasswordProperties","PasswordExpiryDuration":32,"GracePeriod":32,"NotificationPeriod":32,"ObjectType":"iam.EndPointPasswordProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointPasswordProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointPasswordProperties(p, d)[0]
	expectedOp := map[string]interface{}{"password_history": 32, "enforce_strong_password": true, "force_send_password": true, "enable_password_expiry": true, "class_id": "iam.EndPointPasswordProperties", "password_expiry_duration": 32, "grace_period": 32, "notification_period": 32, "object_type": "iam.EndPointPasswordProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointUserRelationship(t *testing.T) {
	p := models.IamEndPointUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointUserRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamEndPointUserPolicyRelationship(t *testing.T) {
	p := models.IamEndPointUserPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamEndPointUserPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamEndPointUserPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIdpRelationship(t *testing.T) {
	p := models.IamIdpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIdpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIdpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIdpReferenceRelationship(t *testing.T) {
	p := models.IamIdpReferenceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIdpReferenceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIdpReferenceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamIpAccessManagementRelationship(t *testing.T) {
	p := models.IamIpAccessManagementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamIpAccessManagementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamIpAccessManagementRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapBaseProperties(t *testing.T) {
	p := models.IamLdapBaseProperties{}
	var d = &schema.ResourceData{}
	c := `{"Domain":"Domain %d","IsPasswordSet":true,"Attribute":"Attribute %d","BindDn":"BindDn %d","Timeout":32,"NestedGroupSearchDepth":32,"EnableGroupAuthorization":true,"EnableEncryption":true,"ObjectType":"iam.LdapBaseProperties","BaseDn":"BaseDn %d","GroupAttribute":"GroupAttribute %d","Filter":"Filter %d","BindMethod":"BindMethod %d","ClassId":"iam.LdapBaseProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapBaseProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapBaseProperties(p, d)[0]
	expectedOp := map[string]interface{}{"domain": "Domain 1", "is_password_set": true, "attribute": "Attribute 1", "bind_dn": "BindDn 1", "timeout": 32, "nested_group_search_depth": 32, "enable_group_authorization": true, "enable_encryption": true, "object_type": "iam.LdapBaseProperties", "base_dn": "BaseDn 1", "group_attribute": "GroupAttribute 1", "filter": "Filter 1", "bind_method": "BindMethod 1", "class_id": "iam.LdapBaseProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapDnsParameters(t *testing.T) {
	p := models.IamLdapDnsParameters{}
	var d = &schema.ResourceData{}
	c := `{"Source":"Source %d","ClassId":"iam.LdapDnsParameters","ObjectType":"iam.LdapDnsParameters","SearchDomain":"SearchDomain %d","SearchForest":"SearchForest %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapDnsParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapDnsParameters(p, d)[0]
	expectedOp := map[string]interface{}{"nr_source": "Source 1", "class_id": "iam.LdapDnsParameters", "object_type": "iam.LdapDnsParameters", "search_domain": "SearchDomain 1", "search_forest": "SearchForest 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLdapPolicyRelationship(t *testing.T) {
	p := models.IamLdapPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLdapPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLdapPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamLocalUserPasswordRelationship(t *testing.T) {
	p := models.IamLocalUserPasswordRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamLocalUserPasswordRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamLocalUserPasswordRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamPermissionRelationship(t *testing.T) {
	p := models.IamPermissionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamPermissionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamPermissionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamPrivateKeySpecRelationship(t *testing.T) {
	p := models.IamPrivateKeySpecRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamPrivateKeySpecRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamPrivateKeySpecRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamQualifierRelationship(t *testing.T) {
	p := models.IamQualifierRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamQualifierRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamQualifierRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamResourceLimitsRelationship(t *testing.T) {
	p := models.IamResourceLimitsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamResourceLimitsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamResourceLimitsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSecurityHolderRelationship(t *testing.T) {
	p := models.IamSecurityHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSecurityHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSecurityHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamServiceProviderRelationship(t *testing.T) {
	p := models.IamServiceProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamServiceProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamServiceProviderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSessionRelationship(t *testing.T) {
	p := models.IamSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSessionLimitsRelationship(t *testing.T) {
	p := models.IamSessionLimitsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSessionLimitsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSessionLimitsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamSystemRelationship(t *testing.T) {
	p := models.IamSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamUserRelationship(t *testing.T) {
	p := models.IamUserRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamUserRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamUserRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIamUserGroupRelationship(t *testing.T) {
	p := models.IamUserGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIamUserGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIamUserGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInfraHardwareInfo(t *testing.T) {
	p := models.InfraHardwareInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"infra.HardwareInfo","ObjectType":"infra.HardwareInfo","CpuCores":32,"CpuSpeed":32,"MemorySize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapInfraHardwareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInfraHardwareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "infra.HardwareInfo", "object_type": "infra.HardwareInfo", "cpu_cores": 32, "cpu_speed": 32, "memory_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryBaseRelationship(t *testing.T) {
	p := models.InventoryBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryDeviceInfoRelationship(t *testing.T) {
	p := models.InventoryDeviceInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryDeviceInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryDeviceInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapInventoryGenericInventoryHolderRelationship(t *testing.T) {
	p := models.InventoryGenericInventoryHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapInventoryGenericInventoryHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapInventoryGenericInventoryHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolBlockLeaseRelationship(t *testing.T) {
	p := models.IppoolBlockLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolBlockLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolBlockLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpLeaseRelationship(t *testing.T) {
	p := models.IppoolIpLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV4Block(t *testing.T) {
	p := models.IppoolIpV4Block{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ippool.IpV4Block","ClassId":"ippool.IpV4Block","Size":32,"From":"From %d","To":"To %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV4Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV4Block(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "ippool.IpV4Block", "class_id": "ippool.IpV4Block", "size": 32, "from": "From 1", "to": "To 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV4Config(t *testing.T) {
	p := models.IppoolIpV4Config{}
	var d = &schema.ResourceData{}
	c := `{"PrimaryDns":"PrimaryDns %d","SecondaryDns":"SecondaryDns %d","Gateway":"Gateway %d","Netmask":"Netmask %d","ClassId":"ippool.IpV4Config","ObjectType":"ippool.IpV4Config"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV4Config(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV4Config(p, d)[0]
	expectedOp := map[string]interface{}{"primary_dns": "PrimaryDns 1", "secondary_dns": "SecondaryDns 1", "gateway": "Gateway 1", "netmask": "Netmask 1", "class_id": "ippool.IpV4Config", "object_type": "ippool.IpV4Config"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV6Block(t *testing.T) {
	p := models.IppoolIpV6Block{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"ippool.IpV6Block","ClassId":"ippool.IpV6Block","Size":32,"From":"From %d","To":"To %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV6Block(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV6Block(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "ippool.IpV6Block", "class_id": "ippool.IpV6Block", "size": 32, "from": "From 1", "to": "To 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolIpV6Config(t *testing.T) {
	p := models.IppoolIpV6Config{}
	var d = &schema.ResourceData{}
	c := `{"Gateway":"Gateway %d","Prefix":32,"PrimaryDns":"PrimaryDns %d","SecondaryDns":"SecondaryDns %d","ClassId":"ippool.IpV6Config","ObjectType":"ippool.IpV6Config"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolIpV6Config(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolIpV6Config(p, d)[0]
	expectedOp := map[string]interface{}{"gateway": "Gateway 1", "prefix": 32, "primary_dns": "PrimaryDns 1", "secondary_dns": "SecondaryDns 1", "class_id": "ippool.IpV6Config", "object_type": "ippool.IpV6Config"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolPoolRelationship(t *testing.T) {
	p := models.IppoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolPoolMemberRelationship(t *testing.T) {
	p := models.IppoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolShadowBlockRelationship(t *testing.T) {
	p := models.IppoolShadowBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolShadowBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolShadowBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolShadowPoolRelationship(t *testing.T) {
	p := models.IppoolShadowPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolShadowPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolShadowPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIppoolUniverseRelationship(t *testing.T) {
	p := models.IppoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIppoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIppoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolBlockRelationship(t *testing.T) {
	p := models.IqnpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolIqnSuffixBlock(t *testing.T) {
	p := models.IqnpoolIqnSuffixBlock{}
	var d = &schema.ResourceData{}
	c := `{"Size":32,"From":32,"Suffix":"Suffix %d","To":32,"ObjectType":"iqnpool.IqnSuffixBlock","ClassId":"iqnpool.IqnSuffixBlock"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolIqnSuffixBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolIqnSuffixBlock(p, d)[0]
	expectedOp := map[string]interface{}{"size": 32, "from": 32, "suffix": "Suffix 1", "to": 32, "object_type": "iqnpool.IqnSuffixBlock", "class_id": "iqnpool.IqnSuffixBlock"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolLeaseRelationship(t *testing.T) {
	p := models.IqnpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolPoolRelationship(t *testing.T) {
	p := models.IqnpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolPoolMemberRelationship(t *testing.T) {
	p := models.IqnpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapIqnpoolUniverseRelationship(t *testing.T) {
	p := models.IqnpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapIqnpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapIqnpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAciCniProfileRelationship(t *testing.T) {
	p := models.KubernetesAciCniProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAciCniProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAciCniProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesActionInfo(t *testing.T) {
	p := models.KubernetesActionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.ActionInfo","ObjectType":"kubernetes.ActionInfo","FailureReason":"FailureReason %d","Name":"Name %d","Status":"Status %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesActionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesActionInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.ActionInfo", "object_type": "kubernetes.ActionInfo", "failure_reason": "FailureReason 1", "name": "Name 1", "status": "Status 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAddonConfiguration(t *testing.T) {
	p := models.KubernetesAddonConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ReleaseName":"ReleaseName %d","UpgradeStrategy":"UpgradeStrategy %d","ClassId":"kubernetes.AddonConfiguration","ObjectType":"kubernetes.AddonConfiguration","InstallStrategy":"InstallStrategy %d","Overrides":"Overrides %d","ReleaseNamespace":"ReleaseNamespace %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAddonConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAddonConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"release_name": "ReleaseName 1", "upgrade_strategy": "UpgradeStrategy 1", "class_id": "kubernetes.AddonConfiguration", "object_type": "kubernetes.AddonConfiguration", "install_strategy": "InstallStrategy 1", "overrides": "Overrides 1", "release_namespace": "ReleaseNamespace 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesAddonDefinitionRelationship(t *testing.T) {
	p := models.KubernetesAddonDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesAddonDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesAddonDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesBaremetalNetworkInfo(t *testing.T) {
	p := models.KubernetesBaremetalNetworkInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.BaremetalNetworkInfo","ObjectType":"kubernetes.BaremetalNetworkInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesBaremetalNetworkInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesBaremetalNetworkInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.BaremetalNetworkInfo", "object_type": "kubernetes.BaremetalNetworkInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesBaseInfrastructureProviderRelationship(t *testing.T) {
	p := models.KubernetesBaseInfrastructureProviderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesBaseInfrastructureProviderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesBaseInfrastructureProviderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesBaseVirtualMachineInfraConfig(t *testing.T) {
	p := models.KubernetesBaseVirtualMachineInfraConfig{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.BaseVirtualMachineInfraConfig","ClassId":"kubernetes.BaseVirtualMachineInfraConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesBaseVirtualMachineInfraConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesBaseVirtualMachineInfraConfig(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "kubernetes.BaseVirtualMachineInfraConfig", "class_id": "kubernetes.BaseVirtualMachineInfraConfig"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesCatalogRelationship(t *testing.T) {
	p := models.KubernetesCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterRelationship(t *testing.T) {
	p := models.KubernetesClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterAddonProfileRelationship(t *testing.T) {
	p := models.KubernetesClusterAddonProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterAddonProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterAddonProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterCertificateConfiguration(t *testing.T) {
	p := models.KubernetesClusterCertificateConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"CaCert":"CaCert %d","SaPublicKey":"SaPublicKey %d","FrontProxyCert":"FrontProxyCert %d","EtcdCert":"EtcdCert %d","ClassId":"kubernetes.ClusterCertificateConfiguration","ObjectType":"kubernetes.ClusterCertificateConfiguration"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterCertificateConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterCertificateConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"ca_cert": "CaCert 1", "sa_public_key": "SaPublicKey 1", "front_proxy_cert": "FrontProxyCert 1", "etcd_cert": "EtcdCert 1", "class_id": "kubernetes.ClusterCertificateConfiguration", "object_type": "kubernetes.ClusterCertificateConfiguration"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterManagementConfig(t *testing.T) {
	p := models.KubernetesClusterManagementConfig{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.ClusterManagementConfig","IsTacPasswdSet":true,"SshUser":"SshUser %d","ObjectType":"kubernetes.ClusterManagementConfig","MasterVip":"MasterVip %d","LoadBalancerCount":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterManagementConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterManagementConfig(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.ClusterManagementConfig", "is_tac_passwd_set": true, "ssh_user": "SshUser 1", "object_type": "kubernetes.ClusterManagementConfig", "master_vip": "MasterVip 1", "load_balancer_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesClusterProfileRelationship(t *testing.T) {
	p := models.KubernetesClusterProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesClusterProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesClusterProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesCniConfig(t *testing.T) {
	p := models.KubernetesCniConfig{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"kubernetes.CniConfig","ClassId":"kubernetes.CniConfig","Registry":"Registry %d","Version":"Version %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesCniConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesCniConfig(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "kubernetes.CniConfig", "class_id": "kubernetes.CniConfig", "registry": "Registry 1", "nr_version": "Version 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesConfigResultRelationship(t *testing.T) {
	p := models.KubernetesConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesConfiguration(t *testing.T) {
	p := models.KubernetesConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.Configuration","ObjectType":"kubernetes.Configuration","KubeConfig":"KubeConfig %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.Configuration", "object_type": "kubernetes.Configuration", "kube_config": "KubeConfig 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesContainerRuntimePolicyRelationship(t *testing.T) {
	p := models.KubernetesContainerRuntimePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesContainerRuntimePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesContainerRuntimePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesDaemonSetStatus(t *testing.T) {
	p := models.KubernetesDaemonSetStatus{}
	var d = &schema.ResourceData{}
	c := `{"CurrentNumberScheduled":32,"NumberMisscheduled":32,"NumberReady":32,"DesiredNumberScheduled":32,"NumberAvailable":"NumberAvailable %d","ObjectType":"kubernetes.DaemonSetStatus","UpdatedNumberScheduled":"UpdatedNumberScheduled %d","ObservedGeneration":32,"ClassId":"kubernetes.DaemonSetStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesDaemonSetStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesDaemonSetStatus(p, d)[0]
	expectedOp := map[string]interface{}{"current_number_scheduled": 32, "number_misscheduled": 32, "number_ready": 32, "desired_number_scheduled": 32, "number_available": "NumberAvailable 1", "object_type": "kubernetes.DaemonSetStatus", "updated_number_scheduled": "UpdatedNumberScheduled 1", "observed_generation": 32, "class_id": "kubernetes.DaemonSetStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesDeploymentStatus(t *testing.T) {
	p := models.KubernetesDeploymentStatus{}
	var d = &schema.ResourceData{}
	c := `{"Replicas":32,"UpdatedReplicas":32,"ClassId":"kubernetes.DeploymentStatus","ObjectType":"kubernetes.DeploymentStatus","AvailableReplicas":32,"ObservedGeneration":32,"ReadyReplicas":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesDeploymentStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesDeploymentStatus(p, d)[0]
	expectedOp := map[string]interface{}{"replicas": 32, "updated_replicas": 32, "class_id": "kubernetes.DeploymentStatus", "object_type": "kubernetes.DeploymentStatus", "available_replicas": 32, "observed_generation": 32, "ready_replicas": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesIngressStatus(t *testing.T) {
	p := models.KubernetesIngressStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.IngressStatus","ObjectType":"kubernetes.IngressStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesIngressStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesIngressStatus(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.IngressStatus", "object_type": "kubernetes.IngressStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNetworkPolicyRelationship(t *testing.T) {
	p := models.KubernetesNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeGroupProfileRelationship(t *testing.T) {
	p := models.KubernetesNodeGroupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeGroupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeGroupProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeInfo(t *testing.T) {
	p := models.KubernetesNodeInfo{}
	var d = &schema.ResourceData{}
	c := `{"BootId":"BootId %d","KernelVersion":"KernelVersion %d","KubeletVersion":"KubeletVersion %d","Architecture":"Architecture %d","OperatingSystem":"OperatingSystem %d","KubeProxyVersion":"KubeProxyVersion %d","OsImage":"OsImage %d","ClassId":"kubernetes.NodeInfo","ObjectType":"kubernetes.NodeInfo","SystemUuid":"SystemUuid %d","MachineId":"MachineId %d","ContainerRuntimeVersion":"ContainerRuntimeVersion %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeInfo(p, d)[0]
	expectedOp := map[string]interface{}{"boot_id": "BootId 1", "kernel_version": "KernelVersion 1", "kubelet_version": "KubeletVersion 1", "architecture": "Architecture 1", "operating_system": "OperatingSystem 1", "kube_proxy_version": "KubeProxyVersion 1", "os_image": "OsImage 1", "class_id": "kubernetes.NodeInfo", "object_type": "kubernetes.NodeInfo", "system_uuid": "SystemUuid 1", "machine_id": "MachineId 1", "container_runtime_version": "ContainerRuntimeVersion 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeProfileRelationship(t *testing.T) {
	p := models.KubernetesNodeProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesNodeSpec(t *testing.T) {
	p := models.KubernetesNodeSpec{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.NodeSpec","ObjectType":"kubernetes.NodeSpec","ProviderId":"ProviderId %d","PodCidr":"PodCidr %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesNodeSpec(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesNodeSpec(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.NodeSpec", "object_type": "kubernetes.NodeSpec", "provider_id": "ProviderId 1", "pod_cidr": "PodCidr 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesObjectMeta(t *testing.T) {
	p := models.KubernetesObjectMeta{}
	var d = &schema.ResourceData{}
	c := `{"Name":"Name %d","Namespace":"Namespace %d","ResourceVersion":"ResourceVersion %d","Uuid":"Uuid %d","CreationTimestamp":"CreationTimestamp %d","ClassId":"kubernetes.ObjectMeta","ObjectType":"kubernetes.ObjectMeta"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesObjectMeta(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesObjectMeta(p, d)[0]
	expectedOp := map[string]interface{}{"name": "Name 1", "namespace": "Namespace 1", "resource_version": "ResourceVersion 1", "uuid": "Uuid 1", "creation_timestamp": "CreationTimestamp 1", "class_id": "kubernetes.ObjectMeta", "object_type": "kubernetes.ObjectMeta"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesPodStatus(t *testing.T) {
	p := models.KubernetesPodStatus{}
	var d = &schema.ResourceData{}
	c := `{"StartTime":"StartTime %d","ClassId":"kubernetes.PodStatus","ObjectType":"kubernetes.PodStatus","HostIp":"HostIp %d","Phase":"Phase %d","PodIp":"PodIp %d","QosClass":"QosClass %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesPodStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesPodStatus(p, d)[0]
	expectedOp := map[string]interface{}{"start_time": "StartTime 1", "class_id": "kubernetes.PodStatus", "object_type": "kubernetes.PodStatus", "host_ip": "HostIp 1", "phase": "Phase 1", "pod_ip": "PodIp 1", "qos_class": "QosClass 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesProxyConfig(t *testing.T) {
	p := models.KubernetesProxyConfig{}
	var d = &schema.ResourceData{}
	c := `{"Port":32,"Username":"Username %d","Hostname":"Hostname %d","ObjectType":"kubernetes.ProxyConfig","Protocol":"Protocol %d","IsPasswordSet":true,"ClassId":"kubernetes.ProxyConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesProxyConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesProxyConfig(p, d)[0]
	expectedOp := map[string]interface{}{"port": 32, "username": "Username 1", "hostname": "Hostname 1", "object_type": "kubernetes.ProxyConfig", "protocol": "Protocol 1", "is_password_set": true, "class_id": "kubernetes.ProxyConfig"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesServiceStatus(t *testing.T) {
	p := models.KubernetesServiceStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.ServiceStatus","ObjectType":"kubernetes.ServiceStatus"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesServiceStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesServiceStatus(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.ServiceStatus", "object_type": "kubernetes.ServiceStatus"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesStatefulSetStatus(t *testing.T) {
	p := models.KubernetesStatefulSetStatus{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"kubernetes.StatefulSetStatus","CollisionCount":32,"ObservedGeneration":32,"UpdateRevision":"UpdateRevision %d","Replicas":32,"UpdatedReplicas":32,"ObjectType":"kubernetes.StatefulSetStatus","CurrentRevision":"CurrentRevision %d","AvailableReplicas":32,"ReadyReplicas":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesStatefulSetStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesStatefulSetStatus(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "kubernetes.StatefulSetStatus", "collision_count": 32, "observed_generation": 32, "update_revision": "UpdateRevision 1", "replicas": 32, "updated_replicas": 32, "object_type": "kubernetes.StatefulSetStatus", "current_revision": "CurrentRevision 1", "available_replicas": 32, "ready_replicas": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesSysConfigPolicyRelationship(t *testing.T) {
	p := models.KubernetesSysConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesSysConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesSysConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesTrustedRegistriesPolicyRelationship(t *testing.T) {
	p := models.KubernetesTrustedRegistriesPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesTrustedRegistriesPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesTrustedRegistriesPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVersionRelationship(t *testing.T) {
	p := models.KubernetesVersionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVersionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVersionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVersionPolicyRelationship(t *testing.T) {
	p := models.KubernetesVersionPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVersionPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVersionPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(t *testing.T) {
	p := models.KubernetesVirtualMachineInfraConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVirtualMachineInfraConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKubernetesVirtualMachineInstanceTypeRelationship(t *testing.T) {
	p := models.KubernetesVirtualMachineInstanceTypeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKubernetesVirtualMachineInstanceTypeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKubernetesVirtualMachineInstanceTypeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKvmSessionRelationship(t *testing.T) {
	p := models.KvmSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKvmSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKvmSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapKvmTunnelRelationship(t *testing.T) {
	p := models.KvmTunnelRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapKvmTunnelRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapKvmTunnelRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseAccountLicenseDataRelationship(t *testing.T) {
	p := models.LicenseAccountLicenseDataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseAccountLicenseDataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseAccountLicenseDataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseCustomerOpRelationship(t *testing.T) {
	p := models.LicenseCustomerOpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseCustomerOpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseCustomerOpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseIwoCustomerOpRelationship(t *testing.T) {
	p := models.LicenseIwoCustomerOpRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseIwoCustomerOpRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseIwoCustomerOpRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseIwoLicenseCountRelationship(t *testing.T) {
	p := models.LicenseIwoLicenseCountRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseIwoLicenseCountRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseIwoLicenseCountRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapLicenseSmartlicenseTokenRelationship(t *testing.T) {
	p := models.LicenseSmartlicenseTokenRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapLicenseSmartlicenseTokenRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapLicenseSmartlicenseTokenRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolBlock(t *testing.T) {
	p := models.MacpoolBlock{}
	var d = &schema.ResourceData{}
	c := `{"From":"From %d","To":"To %d","ClassId":"macpool.Block","ObjectType":"macpool.Block","Size":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolBlock(p, d)[0]
	expectedOp := map[string]interface{}{"from": "From 1", "to": "To 1", "class_id": "macpool.Block", "object_type": "macpool.Block", "size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolIdBlockRelationship(t *testing.T) {
	p := models.MacpoolIdBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolIdBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolIdBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolLeaseRelationship(t *testing.T) {
	p := models.MacpoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolPoolRelationship(t *testing.T) {
	p := models.MacpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolPoolMemberRelationship(t *testing.T) {
	p := models.MacpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMacpoolUniverseRelationship(t *testing.T) {
	p := models.MacpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMacpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMacpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapManagementControllerRelationship(t *testing.T) {
	p := models.ManagementControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapManagementControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapManagementControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapManagementEntityRelationship(t *testing.T) {
	p := models.ManagementEntityRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapManagementEntityRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapManagementEntityRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryArrayRelationship(t *testing.T) {
	p := models.MemoryArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryConfigResultRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryConfigurationRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryConfigurationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryConfigurationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryConfigurationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryLocalSecurity(t *testing.T) {
	p := models.MemoryPersistentMemoryLocalSecurity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"memory.PersistentMemoryLocalSecurity","ObjectType":"memory.PersistentMemoryLocalSecurity","Enabled":true,"IsSecurePassphraseSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryLocalSecurity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryLocalSecurity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "memory.PersistentMemoryLocalSecurity", "object_type": "memory.PersistentMemoryLocalSecurity", "enabled": true, "is_secure_passphrase_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMemoryPersistentMemoryRegionRelationship(t *testing.T) {
	p := models.MemoryPersistentMemoryRegionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMemoryPersistentMemoryRegionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMemoryPersistentMemoryRegionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoBaseMo(t *testing.T) {
	p := models.MoBaseMo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.BaseMo","ClassId":"mo.BaseMo","AccountMoid":"AccountMoid %d","Moid":"Moid %d","SharedScope":"SharedScope %d","DomainGroupMoid":"DomainGroupMoid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoBaseMo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoBaseMo(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.BaseMo", "class_id": "mo.BaseMo", "account_moid": "AccountMoid 1", "moid": "Moid 1", "shared_scope": "SharedScope 1", "domain_group_moid": "DomainGroupMoid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoBaseMoRelationship(t *testing.T) {
	p := models.MoBaseMoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoBaseMoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoBaseMoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoMoRef(t *testing.T) {
	p := models.MoMoRef{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoMoRef(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoMoRef(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapMoVersionContext(t *testing.T) {
	p := models.MoVersionContext{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.VersionContext","ClassId":"mo.VersionContext","Version":"Version %d","VersionType":"VersionType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapMoVersionContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapMoVersionContext(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.VersionContext", "class_id": "mo.VersionContext", "nr_version": "Version 1", "version_type": "VersionType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkElementRelationship(t *testing.T) {
	p := models.NetworkElementRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkElementRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkElementRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkFcZoneInfoRelationship(t *testing.T) {
	p := models.NetworkFcZoneInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkFcZoneInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkFcZoneInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNetworkVlanPortInfoRelationship(t *testing.T) {
	p := models.NetworkVlanPortInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNetworkVlanPortInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNetworkVlanPortInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiaapiNewReleaseDetail(t *testing.T) {
	p := models.NiaapiNewReleaseDetail{}
	var d = &schema.ResourceData{}
	c := `{"SoftwareDownloadLink":"SoftwareDownloadLink %d","Description":"Description %d","ObjectType":"niaapi.NewReleaseDetail","Version":"Version %d","ReleaseNoteLink":"ReleaseNoteLink %d","ReleaseNoteLinkTitle":"ReleaseNoteLinkTitle %d","Link":"Link %d","SoftwareDownloadLinkTitle":"SoftwareDownloadLinkTitle %d","Title":"Title %d","ClassId":"niaapi.NewReleaseDetail"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiaapiNewReleaseDetail(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiaapiNewReleaseDetail(p, d)[0]
	expectedOp := map[string]interface{}{"software_download_link": "SoftwareDownloadLink 1", "description": "Description 1", "object_type": "niaapi.NewReleaseDetail", "nr_version": "Version 1", "release_note_link": "ReleaseNoteLink 1", "release_note_link_title": "ReleaseNoteLinkTitle 1", "link": "Link 1", "software_download_link_title": "SoftwareDownloadLinkTitle 1", "title": "Title 1", "class_id": "niaapi.NewReleaseDetail"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiaapiVersionRegexPlatform(t *testing.T) {
	p := models.NiaapiVersionRegexPlatform{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niaapi.VersionRegexPlatform","ObjectType":"niaapi.VersionRegexPlatform","Anyllregex":"Anyllregex %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiaapiVersionRegexPlatform(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiaapiVersionRegexPlatform(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niaapi.VersionRegexPlatform", "object_type": "niaapi.VersionRegexPlatform", "anyllregex": "Anyllregex 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryBootflashDetails(t *testing.T) {
	p := models.NiatelemetryBootflashDetails{}
	var d = &schema.ResourceData{}
	c := `{"FwRev":"FwRev %d","ModelType":"ModelType %d","Serial":"Serial %d","ClassId":"niatelemetry.BootflashDetails","ObjectType":"niatelemetry.BootflashDetails"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryBootflashDetails(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryBootflashDetails(p, d)[0]
	expectedOp := map[string]interface{}{"fw_rev": "FwRev 1", "model_type": "ModelType 1", "serial": "Serial 1", "class_id": "niatelemetry.BootflashDetails", "object_type": "niatelemetry.BootflashDetails"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryDiskinfo(t *testing.T) {
	p := models.NiatelemetryDiskinfo{}
	var d = &schema.ResourceData{}
	c := `{"Free":32,"Name":"Name %d","Total":32,"Used":32,"ClassId":"niatelemetry.Diskinfo","ObjectType":"niatelemetry.Diskinfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryDiskinfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryDiskinfo(p, d)[0]
	expectedOp := map[string]interface{}{"free": 32, "name": "Name 1", "total": 32, "used": 32, "class_id": "niatelemetry.Diskinfo", "object_type": "niatelemetry.Diskinfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryInterface(t *testing.T) {
	p := models.NiatelemetryInterface{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.Interface","ObjectType":"niatelemetry.Interface","InterfaceDownCount":32,"InterfaceUpCount":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryInterface(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryInterface(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.Interface", "object_type": "niatelemetry.Interface", "interface_down_count": 32, "interface_up_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNexusDashboardsRelationship(t *testing.T) {
	p := models.NiatelemetryNexusDashboardsRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNexusDashboardsRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNexusDashboardsRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNiaInventoryRelationship(t *testing.T) {
	p := models.NiatelemetryNiaInventoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNiaInventoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNiaInventoryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNiaLicenseStateRelationship(t *testing.T) {
	p := models.NiatelemetryNiaLicenseStateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNiaLicenseStateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNiaLicenseStateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNvePacketCounters(t *testing.T) {
	p := models.NiatelemetryNvePacketCounters{}
	var d = &schema.ResourceData{}
	c := `{"McastInpkts":32,"McastOutbytes":32,"UcastInpkts":32,"ClassId":"niatelemetry.NvePacketCounters","ObjectType":"niatelemetry.NvePacketCounters","UcastOutpkts":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNvePacketCounters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNvePacketCounters(p, d)[0]
	expectedOp := map[string]interface{}{"mcast_inpkts": 32, "mcast_outbytes": 32, "ucast_inpkts": 32, "class_id": "niatelemetry.NvePacketCounters", "object_type": "niatelemetry.NvePacketCounters", "ucast_outpkts": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNveVni(t *testing.T) {
	p := models.NiatelemetryNveVni{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.NveVni","DpVniUp":32,"CpVniCount":32,"ObjectType":"niatelemetry.NveVni","CpVniUp":32,"DpVniCount":32,"DpVniDown":32,"CpVniDown":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNveVni(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNveVni(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.NveVni", "dp_vni_up": 32, "cp_vni_count": 32, "object_type": "niatelemetry.NveVni", "cp_vni_up": 32, "dp_vni_count": 32, "dp_vni_down": 32, "cp_vni_down": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNxosBgpMvpn(t *testing.T) {
	p := models.NiatelemetryNxosBgpMvpn{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.NxosBgpMvpn","ConfiguredPeers":32,"NumberOfCommunities":32,"NumberOfClusterLists":32,"CapablePeers":32,"TotalPaths":32,"TableVersion":32,"TotalNetworks":32,"ObjectType":"niatelemetry.NxosBgpMvpn","MemoryUsed":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNxosBgpMvpn(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNxosBgpMvpn(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.NxosBgpMvpn", "configured_peers": 32, "number_of_communities": 32, "number_of_cluster_lists": 32, "capable_peers": 32, "total_paths": 32, "table_version": 32, "total_networks": 32, "object_type": "niatelemetry.NxosBgpMvpn", "memory_used": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetryNxosVtp(t *testing.T) {
	p := models.NiatelemetryNxosVtp{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"niatelemetry.NxosVtp","OperMode":"OperMode %d","PruningMode":"PruningMode %d","RunningVersion":"RunningVersion %d","TrapEnabled":"TrapEnabled %d","V2Mode":"V2Mode %d","Version":32,"ClassId":"niatelemetry.NxosVtp"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetryNxosVtp(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetryNxosVtp(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "niatelemetry.NxosVtp", "oper_mode": "OperMode 1", "pruning_mode": "PruningMode 1", "running_version": "RunningVersion 1", "trap_enabled": "TrapEnabled 1", "v2_mode": "V2Mode 1", "nr_version": 32, "class_id": "niatelemetry.NxosVtp"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapNiatelemetrySmartLicense(t *testing.T) {
	p := models.NiatelemetrySmartLicense{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"niatelemetry.SmartLicense","ObjectType":"niatelemetry.SmartLicense","SmartAccount":"SmartAccount %d","ActiveMode":"ActiveMode %d","AuthStatus":"AuthStatus %d","LicenseUdi":"LicenseUdi %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapNiatelemetrySmartLicense(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapNiatelemetrySmartLicense(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "niatelemetry.SmartLicense", "object_type": "niatelemetry.SmartLicense", "smart_account": "SmartAccount 1", "active_mode": "ActiveMode 1", "auth_status": "AuthStatus 1", "license_udi": "LicenseUdi 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOnpremSchedule(t *testing.T) {
	p := models.OnpremSchedule{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"onprem.Schedule","ObjectType":"onprem.Schedule","DayOfWeek":32,"MonthOfYear":32,"TimeOfDay":32,"RepeatInterval":32,"DayOfMonth":32,"TimeZone":"TimeZone %d","WeekOfMonth":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapOnpremSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOnpremSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "onprem.Schedule", "object_type": "onprem.Schedule", "day_of_week": 32, "month_of_year": 32, "time_of_day": 32, "repeat_interval": 32, "day_of_month": 32, "time_zone": "TimeZone 1", "week_of_month": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOnpremUpgradePhase(t *testing.T) {
	p := models.OnpremUpgradePhase{}
	var d = &schema.ResourceData{}
	c := `{"Failed":true,"Message":"Message %d","Name":"Name %d","ElapsedTime":32,"ClassId":"onprem.UpgradePhase","ObjectType":"onprem.UpgradePhase"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOnpremUpgradePhase(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOnpremUpgradePhase(p, d)[0]
	expectedOp := map[string]interface{}{"failed": true, "message": "Message 1", "name": "Name 1", "elapsed_time": 32, "class_id": "onprem.UpgradePhase", "object_type": "onprem.UpgradePhase"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOrganizationOrganizationRelationship(t *testing.T) {
	p := models.OrganizationOrganizationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOrganizationOrganizationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOrganizationOrganizationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsAnswers(t *testing.T) {
	p := models.OsAnswers{}
	var d = &schema.ResourceData{}
	c := `{"IsRootPasswordSet":true,"IpConfigType":"IpConfigType %d","ObjectType":"os.Answers","Source":"Source %d","NetworkDevice":"NetworkDevice %d","IsRootPasswordCrypted":true,"ClassId":"os.Answers","Hostname":"Hostname %d","Nameserver":"Nameserver %d","ProductKey":"ProductKey %d","IsAnswerFileSet":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsAnswers(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsAnswers(p, d)[0]
	expectedOp := map[string]interface{}{"is_root_password_set": true, "ip_config_type": "IpConfigType 1", "object_type": "os.Answers", "nr_source": "Source 1", "network_device": "NetworkDevice 1", "is_root_password_crypted": true, "class_id": "os.Answers", "hostname": "Hostname 1", "nameserver": "Nameserver 1", "product_key": "ProductKey 1", "is_answer_file_set": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsCatalogRelationship(t *testing.T) {
	p := models.OsCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsConfigurationFileRelationship(t *testing.T) {
	p := models.OsConfigurationFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsConfigurationFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsConfigurationFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsGlobalConfig(t *testing.T) {
	p := models.OsGlobalConfig{}
	var d = &schema.ResourceData{}
	c := `{"InstallMethod":"InstallMethod %d","InstallTargetType":"InstallTargetType %d","OsImageName":"OsImageName %d","WindowsEdition":"WindowsEdition %d","ClassId":"os.GlobalConfig","ObjectType":"os.GlobalConfig","ScuImageName":"ScuImageName %d","ConfigurationFileName":"ConfigurationFileName %d","ConfigurationSource":"ConfigurationSource %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsGlobalConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsGlobalConfig(p, d)[0]
	expectedOp := map[string]interface{}{"install_method": "InstallMethod 1", "install_target_type": "InstallTargetType 1", "os_image_name": "OsImageName 1", "windows_edition": "WindowsEdition 1", "class_id": "os.GlobalConfig", "object_type": "os.GlobalConfig", "scu_image_name": "ScuImageName 1", "configuration_file_name": "ConfigurationFileName 1", "configuration_source": "ConfigurationSource 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsInstallTarget(t *testing.T) {
	p := models.OsInstallTarget{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.InstallTarget","ObjectType":"os.InstallTarget"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsInstallTarget(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsInstallTarget(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "os.InstallTarget", "object_type": "os.InstallTarget"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapOsOperatingSystemParameters(t *testing.T) {
	p := models.OsOperatingSystemParameters{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"os.OperatingSystemParameters","ObjectType":"os.OperatingSystemParameters"}`

	//test when the response is empty
	ffOpEmpty := flattenMapOsOperatingSystemParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapOsOperatingSystemParameters(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "os.OperatingSystemParameters", "object_type": "os.OperatingSystemParameters"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPciSwitchRelationship(t *testing.T) {
	p := models.PciSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPciSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPciSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixDistinguishedName(t *testing.T) {
	p := models.PkixDistinguishedName{}
	var d = &schema.ResourceData{}
	c := `{"CommonName":"CommonName %d","ClassId":"pkix.DistinguishedName","ObjectType":"pkix.DistinguishedName"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixDistinguishedName(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixDistinguishedName(p, d)[0]
	expectedOp := map[string]interface{}{"common_name": "CommonName 1", "class_id": "pkix.DistinguishedName", "object_type": "pkix.DistinguishedName"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixKeyGenerationSpec(t *testing.T) {
	p := models.PkixKeyGenerationSpec{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"pkix.KeyGenerationSpec","ObjectType":"pkix.KeyGenerationSpec","Name":"Name %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixKeyGenerationSpec(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixKeyGenerationSpec(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "pkix.KeyGenerationSpec", "object_type": "pkix.KeyGenerationSpec", "name": "Name 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPkixSubjectAlternateName(t *testing.T) {
	p := models.PkixSubjectAlternateName{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"pkix.SubjectAlternateName","ObjectType":"pkix.SubjectAlternateName"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPkixSubjectAlternateName(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPkixSubjectAlternateName(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "pkix.SubjectAlternateName", "object_type": "pkix.SubjectAlternateName"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyAbstractConfigProfileRelationship(t *testing.T) {
	p := models.PolicyAbstractConfigProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyAbstractConfigProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyAbstractConfigProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyAbstractProfileRelationship(t *testing.T) {
	p := models.PolicyAbstractProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyAbstractProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyAbstractProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigChange(t *testing.T) {
	p := models.PolicyConfigChange{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"policy.ConfigChange","ObjectType":"policy.ConfigChange"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigChange(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigChange(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "policy.ConfigChange", "object_type": "policy.ConfigChange"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigChangeContext(t *testing.T) {
	p := models.PolicyConfigChangeContext{}
	var d = &schema.ResourceData{}
	c := `{"ConfigChangeState":"ConfigChangeState %d","ClassId":"policy.ConfigChangeContext","ObjectType":"policy.ConfigChangeContext","ConfigChangeError":"ConfigChangeError %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigChangeContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigChangeContext(p, d)[0]
	expectedOp := map[string]interface{}{"config_change_state": "ConfigChangeState 1", "class_id": "policy.ConfigChangeContext", "object_type": "policy.ConfigChangeContext", "config_change_error": "ConfigChangeError 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigContext(t *testing.T) {
	p := models.PolicyConfigContext{}
	var d = &schema.ResourceData{}
	c := `{"ErrorState":"ErrorState %d","OperState":"OperState %d","ConfigState":"ConfigState %d","ClassId":"policy.ConfigContext","ObjectType":"policy.ConfigContext","ConfigType":"ConfigType %d","ControlAction":"ControlAction %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigContext(p, d)[0]
	expectedOp := map[string]interface{}{"error_state": "ErrorState 1", "oper_state": "OperState 1", "config_state": "ConfigState 1", "class_id": "policy.ConfigContext", "object_type": "policy.ConfigContext", "config_type": "ConfigType 1", "control_action": "ControlAction 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPolicyConfigResultContext(t *testing.T) {
	p := models.PolicyConfigResultContext{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"policy.ConfigResultContext","ObjectType":"policy.ConfigResultContext","EntityName":"EntityName %d","EntityType":"EntityType %d","EntityMoid":"EntityMoid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPolicyConfigResultContext(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPolicyConfigResultContext(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "policy.ConfigResultContext", "object_type": "policy.ConfigResultContext", "entity_name": "EntityName 1", "entity_type": "EntityType 1", "entity_moid": "EntityMoid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortGroupRelationship(t *testing.T) {
	p := models.PortGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortInterfaceBaseRelationship(t *testing.T) {
	p := models.PortInterfaceBaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortInterfaceBaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortInterfaceBaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPortSubGroupRelationship(t *testing.T) {
	p := models.PortSubGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPortSubGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPortSubGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapPowerControlStateRelationship(t *testing.T) {
	p := models.PowerControlStateRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapPowerControlStateRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapPowerControlStateRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecommendationCapacityRunwayRelationship(t *testing.T) {
	p := models.RecommendationCapacityRunwayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecommendationCapacityRunwayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecommendationCapacityRunwayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryAbstractBackupInfoRelationship(t *testing.T) {
	p := models.RecoveryAbstractBackupInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryAbstractBackupInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryAbstractBackupInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupConfigPolicyRelationship(t *testing.T) {
	p := models.RecoveryBackupConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupProfileRelationship(t *testing.T) {
	p := models.RecoveryBackupProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryBackupSchedule(t *testing.T) {
	p := models.RecoveryBackupSchedule{}
	var d = &schema.ResourceData{}
	c := `{"FrequencyUnit":"FrequencyUnit %d","Hours":32,"ClassId":"recovery.BackupSchedule","ObjectType":"recovery.BackupSchedule"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryBackupSchedule(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryBackupSchedule(p, d)[0]
	expectedOp := map[string]interface{}{"frequency_unit": "FrequencyUnit 1", "hours": 32, "class_id": "recovery.BackupSchedule", "object_type": "recovery.BackupSchedule"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryConfigParams(t *testing.T) {
	p := models.RecoveryConfigParams{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"recovery.ConfigParams","ClassId":"recovery.ConfigParams"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryConfigParams(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryConfigParams(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "recovery.ConfigParams", "class_id": "recovery.ConfigParams"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryConfigResultRelationship(t *testing.T) {
	p := models.RecoveryConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapRecoveryScheduleConfigPolicyRelationship(t *testing.T) {
	p := models.RecoveryScheduleConfigPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapRecoveryScheduleConfigPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapRecoveryScheduleConfigPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourceGroupRelationship(t *testing.T) {
	p := models.ResourceGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourceGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourceGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourceMembershipHolderRelationship(t *testing.T) {
	p := models.ResourceMembershipHolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourceMembershipHolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourceMembershipHolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourceReservationRelationship(t *testing.T) {
	p := models.ResourceReservationRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourceReservationRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourceReservationRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolLeaseRelationship(t *testing.T) {
	p := models.ResourcepoolLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolLeaseParameters(t *testing.T) {
	p := models.ResourcepoolLeaseParameters{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"resourcepool.LeaseParameters","ObjectType":"resourcepool.LeaseParameters"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolLeaseParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolLeaseParameters(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "resourcepool.LeaseParameters", "object_type": "resourcepool.LeaseParameters"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolLeaseResourceRelationship(t *testing.T) {
	p := models.ResourcepoolLeaseResourceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolLeaseResourceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolLeaseResourceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolPoolRelationship(t *testing.T) {
	p := models.ResourcepoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolPoolMemberRelationship(t *testing.T) {
	p := models.ResourcepoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolResourcePoolParameters(t *testing.T) {
	p := models.ResourcepoolResourcePoolParameters{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"resourcepool.ResourcePoolParameters","ObjectType":"resourcepool.ResourcePoolParameters"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolResourcePoolParameters(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolResourcePoolParameters(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "resourcepool.ResourcePoolParameters", "object_type": "resourcepool.ResourcePoolParameters"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapResourcepoolUniverseRelationship(t *testing.T) {
	p := models.ResourcepoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapResourcepoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapResourcepoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanProfileRelationship(t *testing.T) {
	p := models.SdwanProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanRouterPolicyRelationship(t *testing.T) {
	p := models.SdwanRouterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanRouterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanRouterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSdwanVmanageAccountPolicyRelationship(t *testing.T) {
	p := models.SdwanVmanageAccountPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSdwanVmanageAccountPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSdwanVmanageAccountPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapServerConfigResultRelationship(t *testing.T) {
	p := models.ServerConfigResultRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapServerConfigResultRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapServerConfigResultRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapServerProfileRelationship(t *testing.T) {
	p := models.ServerProfileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapServerProfileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapServerProfileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSessionAbstractSessionRelationship(t *testing.T) {
	p := models.SessionAbstractSessionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSessionAbstractSessionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSessionAbstractSessionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwareHyperflexDistributableRelationship(t *testing.T) {
	p := models.SoftwareHyperflexDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwareHyperflexDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwareHyperflexDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwareSolutionDistributableRelationship(t *testing.T) {
	p := models.SoftwareSolutionDistributableRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwareSolutionDistributableRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwareSolutionDistributableRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryCatalogRelationship(t *testing.T) {
	p := models.SoftwarerepositoryCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryFileRelationship(t *testing.T) {
	p := models.SoftwarerepositoryFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryFileServer(t *testing.T) {
	p := models.SoftwarerepositoryFileServer{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"softwarerepository.FileServer","ObjectType":"softwarerepository.FileServer"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryFileServer(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryFileServer(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "softwarerepository.FileServer", "object_type": "softwarerepository.FileServer"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryOperatingSystemFileRelationship(t *testing.T) {
	p := models.SoftwarerepositoryOperatingSystemFileRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryOperatingSystemFileRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapSoftwarerepositoryReleaseRelationship(t *testing.T) {
	p := models.SoftwarerepositoryReleaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapSoftwarerepositoryReleaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapSoftwarerepositoryReleaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageAutomaticDriveGroup(t *testing.T) {
	p := models.StorageAutomaticDriveGroup{}
	var d = &schema.ResourceData{}
	c := `{"DrivesPerSpan":32,"MinimumDriveSize":32,"NumDedicatedHotSpares":"NumDedicatedHotSpares %d","NumberOfSpans":32,"UseRemainingDrives":true,"ClassId":"storage.AutomaticDriveGroup","ObjectType":"storage.AutomaticDriveGroup","DriveType":"DriveType %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageAutomaticDriveGroup(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageAutomaticDriveGroup(p, d)[0]
	expectedOp := map[string]interface{}{"drives_per_span": 32, "minimum_drive_size": 32, "num_dedicated_hot_spares": "NumDedicatedHotSpares 1", "number_of_spans": 32, "use_remaining_drives": true, "class_id": "storage.AutomaticDriveGroup", "object_type": "storage.AutomaticDriveGroup", "drive_type": "DriveType 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageBaseCapacity(t *testing.T) {
	p := models.StorageBaseCapacity{}
	var d = &schema.ResourceData{}
	c := `{"CapacityUtilization":32.000000,"Free":32,"Total":32,"Used":32,"Available":32,"ClassId":"storage.BaseCapacity","ObjectType":"storage.BaseCapacity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageBaseCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageBaseCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"capacity_utilization": 32.000000, "free": 32, "total": 32, "used": 32, "available": 32, "class_id": "storage.BaseCapacity", "object_type": "storage.BaseCapacity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageBaseClusterRelationship(t *testing.T) {
	p := models.StorageBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageBaseClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageControllerRelationship(t *testing.T) {
	p := models.StorageControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageDiskGroupRelationship(t *testing.T) {
	p := models.StorageDiskGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageDiskGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageDiskGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageEnclosureRelationship(t *testing.T) {
	p := models.StorageEnclosureRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageEnclosureRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageEnclosureRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageFlexFlashControllerRelationship(t *testing.T) {
	p := models.StorageFlexFlashControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageFlexFlashControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageFlexFlashControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageFlexUtilControllerRelationship(t *testing.T) {
	p := models.StorageFlexUtilControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageFlexUtilControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageFlexUtilControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiArrayRelationship(t *testing.T) {
	p := models.StorageHitachiArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiHostRelationship(t *testing.T) {
	p := models.StorageHitachiHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiParityGroupRelationship(t *testing.T) {
	p := models.StorageHitachiParityGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiParityGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiParityGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiPoolRelationship(t *testing.T) {
	p := models.StorageHitachiPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHitachiVolumeRelationship(t *testing.T) {
	p := models.StorageHitachiVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHitachiVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHitachiVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageHyperFlexStorageContainerRelationship(t *testing.T) {
	p := models.StorageHyperFlexStorageContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageHyperFlexStorageContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageHyperFlexStorageContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageM2VirtualDriveConfig(t *testing.T) {
	p := models.StorageM2VirtualDriveConfig{}
	var d = &schema.ResourceData{}
	c := `{"ControllerSlot":"ControllerSlot %d","Enable":true,"ClassId":"storage.M2VirtualDriveConfig","ObjectType":"storage.M2VirtualDriveConfig"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageM2VirtualDriveConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageM2VirtualDriveConfig(p, d)[0]
	expectedOp := map[string]interface{}{"controller_slot": "ControllerSlot 1", "enable": true, "class_id": "storage.M2VirtualDriveConfig", "object_type": "storage.M2VirtualDriveConfig"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageManualDriveGroup(t *testing.T) {
	p := models.StorageManualDriveGroup{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"storage.ManualDriveGroup","ObjectType":"storage.ManualDriveGroup","DedicatedHotSpares":"DedicatedHotSpares %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageManualDriveGroup(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageManualDriveGroup(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "storage.ManualDriveGroup", "object_type": "storage.ManualDriveGroup", "dedicated_hot_spares": "DedicatedHotSpares 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppClusterRelationship(t *testing.T) {
	p := models.StorageNetAppClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppEthernetPortRelationship(t *testing.T) {
	p := models.StorageNetAppEthernetPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppEthernetPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppEthernetPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppEthernetPortLag(t *testing.T) {
	p := models.StorageNetAppEthernetPortLag{}
	var d = &schema.ResourceData{}
	c := `{"Mode":"Mode %d","ClassId":"storage.NetAppEthernetPortLag","ObjectType":"storage.NetAppEthernetPortLag","DistributionPolicy":"DistributionPolicy %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppEthernetPortLag(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppEthernetPortLag(p, d)[0]
	expectedOp := map[string]interface{}{"mode": "Mode 1", "class_id": "storage.NetAppEthernetPortLag", "object_type": "storage.NetAppEthernetPortLag", "distribution_policy": "DistributionPolicy 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppEthernetPortVlan(t *testing.T) {
	p := models.StorageNetAppEthernetPortVlan{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"storage.NetAppEthernetPortVlan","ObjectType":"storage.NetAppEthernetPortVlan","Tag":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppEthernetPortVlan(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppEthernetPortVlan(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "storage.NetAppEthernetPortVlan", "object_type": "storage.NetAppEthernetPortVlan", "tag": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppFcPortRelationship(t *testing.T) {
	p := models.StorageNetAppFcPortRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppFcPortRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppFcPortRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppHighAvailability(t *testing.T) {
	p := models.StorageNetAppHighAvailability{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"GivebackState":"GivebackState %d","PartnerName":"PartnerName %d","PartnerUuid":"PartnerUuid %d","TakeoverState":"TakeoverState %d","ClassId":"storage.NetAppHighAvailability","ObjectType":"storage.NetAppHighAvailability"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppHighAvailability(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppHighAvailability(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "giveback_state": "GivebackState 1", "partner_name": "PartnerName 1", "partner_uuid": "PartnerUuid 1", "takeover_state": "TakeoverState 1", "class_id": "storage.NetAppHighAvailability", "object_type": "storage.NetAppHighAvailability"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppLunRelationship(t *testing.T) {
	p := models.StorageNetAppLunRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppLunRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppLunRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppNodeRelationship(t *testing.T) {
	p := models.StorageNetAppNodeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppNodeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppNodeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppPerformanceMetricsAverage(t *testing.T) {
	p := models.StorageNetAppPerformanceMetricsAverage{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"storage.NetAppPerformanceMetricsAverage","Period":32,"Throughput":32.000000,"Iops":32.000000,"Latency":32.000000,"ClassId":"storage.NetAppPerformanceMetricsAverage"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppPerformanceMetricsAverage(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppPerformanceMetricsAverage(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "storage.NetAppPerformanceMetricsAverage", "period": 32, "throughput": 32.000000, "iops": 32.000000, "latency": 32.000000, "class_id": "storage.NetAppPerformanceMetricsAverage"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppStorageClusterEfficiency(t *testing.T) {
	p := models.StorageNetAppStorageClusterEfficiency{}
	var d = &schema.ResourceData{}
	c := `{"LogicalUsed":32,"ClassId":"storage.NetAppStorageClusterEfficiency","ObjectType":"storage.NetAppStorageClusterEfficiency","Ratio":32.000000,"Savings":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppStorageClusterEfficiency(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppStorageClusterEfficiency(p, d)[0]
	expectedOp := map[string]interface{}{"logical_used": 32, "class_id": "storage.NetAppStorageClusterEfficiency", "object_type": "storage.NetAppStorageClusterEfficiency", "ratio": 32.000000, "savings": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppStorageVmRelationship(t *testing.T) {
	p := models.StorageNetAppStorageVmRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppStorageVmRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppStorageVmRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageNetAppVolumeRelationship(t *testing.T) {
	p := models.StorageNetAppVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageNetAppVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageNetAppVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePhysicalDiskRelationship(t *testing.T) {
	p := models.StoragePhysicalDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePhysicalDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePhysicalDiskRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureArrayRelationship(t *testing.T) {
	p := models.StoragePureArrayRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureArrayRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureArrayRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureControllerRelationship(t *testing.T) {
	p := models.StoragePureControllerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureControllerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureControllerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureHostRelationship(t *testing.T) {
	p := models.StoragePureHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureHostGroupRelationship(t *testing.T) {
	p := models.StoragePureHostGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureHostGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureHostGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureProtectionGroupRelationship(t *testing.T) {
	p := models.StoragePureProtectionGroupRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureProtectionGroupRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureProtectionGroupRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureProtectionGroupSnapshotRelationship(t *testing.T) {
	p := models.StoragePureProtectionGroupSnapshotRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureProtectionGroupSnapshotRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureProtectionGroupSnapshotRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStoragePureVolumeRelationship(t *testing.T) {
	p := models.StoragePureVolumeRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStoragePureVolumeRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStoragePureVolumeRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageR0Drive(t *testing.T) {
	p := models.StorageR0Drive{}
	var d = &schema.ResourceData{}
	c := `{"DriveSlots":"DriveSlots %d","ClassId":"storage.R0Drive","ObjectType":"storage.R0Drive","DriveSlotsList":"DriveSlotsList %d","Enable":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageR0Drive(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageR0Drive(p, d)[0]
	expectedOp := map[string]interface{}{"drive_slots": "DriveSlots 1", "class_id": "storage.R0Drive", "object_type": "storage.R0Drive", "drive_slots_list": "DriveSlotsList 1", "enable": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageSasExpanderRelationship(t *testing.T) {
	p := models.StorageSasExpanderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageSasExpanderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageSasExpanderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageStoragePolicyRelationship(t *testing.T) {
	p := models.StorageStoragePolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageStoragePolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageStoragePolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveRelationship(t *testing.T) {
	p := models.StorageVirtualDriveRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveContainerRelationship(t *testing.T) {
	p := models.StorageVirtualDriveContainerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveContainerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveContainerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapStorageVirtualDriveExtensionRelationship(t *testing.T) {
	p := models.StorageVirtualDriveExtensionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapStorageVirtualDriveExtensionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapStorageVirtualDriveExtensionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamBaseAdvisoryRelationship(t *testing.T) {
	p := models.TamBaseAdvisoryRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamBaseAdvisoryRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamBaseAdvisoryRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamBaseAdvisoryDetails(t *testing.T) {
	p := models.TamBaseAdvisoryDetails{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"tam.BaseAdvisoryDetails","Description":"Description %d","ClassId":"tam.BaseAdvisoryDetails"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamBaseAdvisoryDetails(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamBaseAdvisoryDetails(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "tam.BaseAdvisoryDetails", "description": "Description 1", "class_id": "tam.BaseAdvisoryDetails"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTamSeverity(t *testing.T) {
	p := models.TamSeverity{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"tam.Severity","ClassId":"tam.Severity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTamSeverity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTamSeverity(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "tam.Severity", "class_id": "tam.Severity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTechsupportmanagementTechSupportBundleRelationship(t *testing.T) {
	p := models.TechsupportmanagementTechSupportBundleRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTechsupportmanagementTechSupportBundleRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTechsupportmanagementTechSupportBundleRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTechsupportmanagementTechSupportStatusRelationship(t *testing.T) {
	p := models.TechsupportmanagementTechSupportStatusRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTechsupportmanagementTechSupportStatusRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTechsupportmanagementTechSupportStatusRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapTopSystemRelationship(t *testing.T) {
	p := models.TopSystemRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapTopSystemRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapTopSystemRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolBlockRelationship(t *testing.T) {
	p := models.UuidpoolBlockRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolBlockRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolBlockRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolPoolRelationship(t *testing.T) {
	p := models.UuidpoolPoolRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolPoolRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolPoolRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolPoolMemberRelationship(t *testing.T) {
	p := models.UuidpoolPoolMemberRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolPoolMemberRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolPoolMemberRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUniverseRelationship(t *testing.T) {
	p := models.UuidpoolUniverseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUniverseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUniverseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUuidBlock(t *testing.T) {
	p := models.UuidpoolUuidBlock{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"uuidpool.UuidBlock","Size":32,"From":"From %d","To":"To %d","ObjectType":"uuidpool.UuidBlock"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUuidBlock(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUuidBlock(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "uuidpool.UuidBlock", "size": 32, "from": "From 1", "to": "To 1", "object_type": "uuidpool.UuidBlock"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapUuidpoolUuidLeaseRelationship(t *testing.T) {
	p := models.UuidpoolUuidLeaseRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapUuidpoolUuidLeaseRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapUuidpoolUuidLeaseRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationActionInfo(t *testing.T) {
	p := models.VirtualizationActionInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.ActionInfo","Name":"Name %d","Status":"Status %d","FailureReason":"FailureReason %d","ObjectType":"virtualization.ActionInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationActionInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationActionInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.ActionInfo", "name": "Name 1", "status": "Status 1", "failure_reason": "FailureReason 1", "object_type": "virtualization.ActionInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseClusterRelationship(t *testing.T) {
	p := models.VirtualizationBaseClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseHostRelationship(t *testing.T) {
	p := models.VirtualizationBaseHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseHostConfiguration(t *testing.T) {
	p := models.VirtualizationBaseHostConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.BaseHostConfiguration","ClassId":"virtualization.BaseHostConfiguration"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseHostConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseHostConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "virtualization.BaseHostConfiguration", "class_id": "virtualization.BaseHostConfiguration"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseNetworkRelationship(t *testing.T) {
	p := models.VirtualizationBaseNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVirtualDiskRelationship(t *testing.T) {
	p := models.VirtualizationBaseVirtualDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVirtualDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVirtualDiskRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationBaseVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVirtualNetworkRelationship(t *testing.T) {
	p := models.VirtualizationBaseVirtualNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVirtualNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVirtualNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBaseVmConfiguration(t *testing.T) {
	p := models.VirtualizationBaseVmConfiguration{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.BaseVmConfiguration","ObjectType":"virtualization.BaseVmConfiguration"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBaseVmConfiguration(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBaseVmConfiguration(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.BaseVmConfiguration", "object_type": "virtualization.BaseVmConfiguration"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationBondState(t *testing.T) {
	p := models.VirtualizationBondState{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.BondState","ObjectType":"virtualization.BondState","ActiveSlave":"ActiveSlave %d","Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationBondState(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationBondState(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.BondState", "object_type": "virtualization.BondState", "active_slave": "ActiveSlave 1", "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCiscoHypervisorManagerRelationship(t *testing.T) {
	p := models.VirtualizationCiscoHypervisorManagerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCiscoHypervisorManagerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCiscoHypervisorManagerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCloudInitConfig(t *testing.T) {
	p := models.VirtualizationCloudInitConfig{}
	var d = &schema.ResourceData{}
	c := `{"NetworkDataBase64Encoded":true,"UserData":"UserData %d","UserDataBase64Encoded":true,"ConfigType":"ConfigType %d","ClassId":"virtualization.CloudInitConfig","ObjectType":"virtualization.CloudInitConfig","NetworkData":"NetworkData %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCloudInitConfig(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCloudInitConfig(p, d)[0]
	expectedOp := map[string]interface{}{"network_data_base64_encoded": true, "user_data": "UserData 1", "user_data_base64_encoded": true, "config_type": "ConfigType 1", "class_id": "virtualization.CloudInitConfig", "object_type": "virtualization.CloudInitConfig", "network_data": "NetworkData 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationComputeCapacity(t *testing.T) {
	p := models.VirtualizationComputeCapacity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.ComputeCapacity","ObjectType":"virtualization.ComputeCapacity","Capacity":32,"Free":32,"Used":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationComputeCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationComputeCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.ComputeCapacity", "object_type": "virtualization.ComputeCapacity", "capacity": 32, "free": 32, "used": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCpuAllocation(t *testing.T) {
	p := models.VirtualizationCpuAllocation{}
	var d = &schema.ResourceData{}
	c := `{"Reserved":32,"Total":32,"Used":32,"Free":32,"ClassId":"virtualization.CpuAllocation","ObjectType":"virtualization.CpuAllocation"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCpuAllocation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCpuAllocation(p, d)[0]
	expectedOp := map[string]interface{}{"reserved": 32, "total": 32, "used": 32, "free": 32, "class_id": "virtualization.CpuAllocation", "object_type": "virtualization.CpuAllocation"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationCpuInfo(t *testing.T) {
	p := models.VirtualizationCpuInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.CpuInfo","Cores":32,"Description":"Description %d","Sockets":32,"Speed":32,"Vendor":"Vendor %d","ClassId":"virtualization.CpuInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationCpuInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationCpuInfo(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "virtualization.CpuInfo", "cores": 32, "description": "Description 1", "sockets": 32, "speed": 32, "vendor": "Vendor 1", "class_id": "virtualization.CpuInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationDiskStatus(t *testing.T) {
	p := models.VirtualizationDiskStatus{}
	var d = &schema.ResourceData{}
	c := `{"VolumeHandle":"VolumeHandle %d","ObjectType":"virtualization.DiskStatus","ClassId":"virtualization.DiskStatus","VolumeName":"VolumeName %d","DownloadPercentage":"DownloadPercentage %d","State":"State %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationDiskStatus(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationDiskStatus(p, d)[0]
	expectedOp := map[string]interface{}{"volume_handle": "VolumeHandle 1", "object_type": "virtualization.DiskStatus", "class_id": "virtualization.DiskStatus", "volume_name": "VolumeName 1", "download_percentage": "DownloadPercentage 1", "state": "State 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationGuestInfo(t *testing.T) {
	p := models.VirtualizationGuestInfo{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.GuestInfo","ClassId":"virtualization.GuestInfo","Name":"Name %d","OperatingSystem":"OperatingSystem %d","Hostname":"Hostname %d","IpAddress":"IpAddress %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationGuestInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationGuestInfo(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "virtualization.GuestInfo", "class_id": "virtualization.GuestInfo", "name": "Name 1", "operating_system": "OperatingSystem 1", "hostname": "Hostname 1", "ip_address": "IpAddress 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweClusterRelationship(t *testing.T) {
	p := models.VirtualizationIweClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweDvUplinkRelationship(t *testing.T) {
	p := models.VirtualizationIweDvUplinkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweDvUplinkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweDvUplinkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweDvswitchRelationship(t *testing.T) {
	p := models.VirtualizationIweDvswitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweDvswitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweDvswitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweHostRelationship(t *testing.T) {
	p := models.VirtualizationIweHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweNetworkRelationship(t *testing.T) {
	p := models.VirtualizationIweNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweVirtualDiskRelationship(t *testing.T) {
	p := models.VirtualizationIweVirtualDiskRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweVirtualDiskRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweVirtualDiskRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationIweVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationIweVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationIweVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationIweVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationMemoryAllocation(t *testing.T) {
	p := models.VirtualizationMemoryAllocation{}
	var d = &schema.ResourceData{}
	c := `{"Free":32,"ClassId":"virtualization.MemoryAllocation","ObjectType":"virtualization.MemoryAllocation","Reserved":32,"Total":32,"Used":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationMemoryAllocation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationMemoryAllocation(p, d)[0]
	expectedOp := map[string]interface{}{"free": 32, "class_id": "virtualization.MemoryAllocation", "object_type": "virtualization.MemoryAllocation", "reserved": 32, "total": 32, "used": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationMemoryCapacity(t *testing.T) {
	p := models.VirtualizationMemoryCapacity{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.MemoryCapacity","Free":32,"Used":32,"Capacity":32,"ObjectType":"virtualization.MemoryCapacity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationMemoryCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationMemoryCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.MemoryCapacity", "free": 32, "used": 32, "capacity": 32, "object_type": "virtualization.MemoryCapacity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationProductInfo(t *testing.T) {
	p := models.VirtualizationProductInfo{}
	var d = &schema.ResourceData{}
	c := `{"Build":"Build %d","ProductName":"ProductName %d","ProductType":"ProductType %d","ProductVendor":"ProductVendor %d","Version":"Version %d","ClassId":"virtualization.ProductInfo","ObjectType":"virtualization.ProductInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationProductInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationProductInfo(p, d)[0]
	expectedOp := map[string]interface{}{"build": "Build 1", "product_name": "ProductName 1", "product_type": "ProductType 1", "product_vendor": "ProductVendor 1", "nr_version": "Version 1", "class_id": "virtualization.ProductInfo", "object_type": "virtualization.ProductInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationStorageCapacity(t *testing.T) {
	p := models.VirtualizationStorageCapacity{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"virtualization.StorageCapacity","Free":32,"Used":32,"Capacity":32,"ClassId":"virtualization.StorageCapacity"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationStorageCapacity(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationStorageCapacity(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "virtualization.StorageCapacity", "free": 32, "used": 32, "capacity": 32, "class_id": "virtualization.StorageCapacity"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareClusterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDatacenterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDatacenterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDatacenterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDatacenterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDatastoreRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDatastoreRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDatastoreRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDatastoreRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDatastoreClusterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDatastoreClusterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDatastoreClusterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDatastoreClusterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDistributedNetworkRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDistributedNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDistributedNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDistributedNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareDistributedSwitchRelationship(t *testing.T) {
	p := models.VirtualizationVmwareDistributedSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareDistributedSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareDistributedSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareFolderRelationship(t *testing.T) {
	p := models.VirtualizationVmwareFolderRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareFolderRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareFolderRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareHostRelationship(t *testing.T) {
	p := models.VirtualizationVmwareHostRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareHostRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareHostRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareNetworkRelationship(t *testing.T) {
	p := models.VirtualizationVmwareNetworkRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareNetworkRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareNetworkRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(t *testing.T) {
	p := models.VirtualizationVmwarePhysicalNetworkInterfaceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwarePhysicalNetworkInterfaceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareRemoteDisplayInfo(t *testing.T) {
	p := models.VirtualizationVmwareRemoteDisplayInfo{}
	var d = &schema.ResourceData{}
	c := `{"RemoteDisplayPassword":"RemoteDisplayPassword %d","ClassId":"virtualization.VmwareRemoteDisplayInfo","ObjectType":"virtualization.VmwareRemoteDisplayInfo","RemoteDisplayVncKey":"RemoteDisplayVncKey %d","RemoteDisplayVncPort":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareRemoteDisplayInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareRemoteDisplayInfo(p, d)[0]
	expectedOp := map[string]interface{}{"remote_display_password": "RemoteDisplayPassword 1", "class_id": "virtualization.VmwareRemoteDisplayInfo", "object_type": "virtualization.VmwareRemoteDisplayInfo", "remote_display_vnc_key": "RemoteDisplayVncKey 1", "remote_display_vnc_port": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareResourceConsumption(t *testing.T) {
	p := models.VirtualizationVmwareResourceConsumption{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareResourceConsumption","ObjectType":"virtualization.VmwareResourceConsumption","CpuConsumed":32,"MemoryConsumed":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareResourceConsumption(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareResourceConsumption(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareResourceConsumption", "object_type": "virtualization.VmwareResourceConsumption", "cpu_consumed": 32, "memory_consumed": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareSharesInfo(t *testing.T) {
	p := models.VirtualizationVmwareSharesInfo{}
	var d = &schema.ResourceData{}
	c := `{"Level":"Level %d","Shares":32,"ClassId":"virtualization.VmwareSharesInfo","ObjectType":"virtualization.VmwareSharesInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareSharesInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareSharesInfo(p, d)[0]
	expectedOp := map[string]interface{}{"level": "Level 1", "shares": 32, "class_id": "virtualization.VmwareSharesInfo", "object_type": "virtualization.VmwareSharesInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareTeamingAndFailover(t *testing.T) {
	p := models.VirtualizationVmwareTeamingAndFailover{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareTeamingAndFailover","ObjectType":"virtualization.VmwareTeamingAndFailover","LoadBalancing":"LoadBalancing %d","Name":"Name %d","Failback":true,"NotifySwitches":true,"NetworkFailureDetection":"NetworkFailureDetection %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareTeamingAndFailover(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareTeamingAndFailover(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareTeamingAndFailover", "object_type": "virtualization.VmwareTeamingAndFailover", "load_balancing": "LoadBalancing 1", "name": "Name 1", "failback": true, "notify_switches": true, "network_failure_detection": "NetworkFailureDetection 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVcenterRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVcenterRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVcenterRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVcenterRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVirtualMachineRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVirtualMachineRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVirtualMachineRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVirtualMachineRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVirtualSwitchRelationship(t *testing.T) {
	p := models.VirtualizationVmwareVirtualSwitchRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVirtualSwitchRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVirtualSwitchRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmCpuShareInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmCpuShareInfo{}
	var d = &schema.ResourceData{}
	c := `{"CpuLimit":32,"CpuOverheadLimit":32,"CpuReservation":32,"CpuShares":32,"ClassId":"virtualization.VmwareVmCpuShareInfo","ObjectType":"virtualization.VmwareVmCpuShareInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmCpuShareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmCpuShareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"cpu_limit": 32, "cpu_overhead_limit": 32, "cpu_reservation": 32, "cpu_shares": 32, "class_id": "virtualization.VmwareVmCpuShareInfo", "object_type": "virtualization.VmwareVmCpuShareInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmCpuSocketInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmCpuSocketInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareVmCpuSocketInfo","ObjectType":"virtualization.VmwareVmCpuSocketInfo","NumSockets":32,"CoresPerSocket":32,"NumCpus":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmCpuSocketInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmCpuSocketInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareVmCpuSocketInfo", "object_type": "virtualization.VmwareVmCpuSocketInfo", "num_sockets": 32, "cores_per_socket": 32, "num_cpus": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmDiskCommitInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmDiskCommitInfo{}
	var d = &schema.ResourceData{}
	c := `{"UnsharedDisk":32,"CommittedDisk":32,"UnCommittedDisk":32,"ClassId":"virtualization.VmwareVmDiskCommitInfo","ObjectType":"virtualization.VmwareVmDiskCommitInfo"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmDiskCommitInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmDiskCommitInfo(p, d)[0]
	expectedOp := map[string]interface{}{"unshared_disk": 32, "committed_disk": 32, "un_committed_disk": 32, "class_id": "virtualization.VmwareVmDiskCommitInfo", "object_type": "virtualization.VmwareVmDiskCommitInfo"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVirtualizationVmwareVmMemoryShareInfo(t *testing.T) {
	p := models.VirtualizationVmwareVmMemoryShareInfo{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"virtualization.VmwareVmMemoryShareInfo","ObjectType":"virtualization.VmwareVmMemoryShareInfo","MemShares":32,"MemLimit":32,"MemOverheadLimit":32,"MemReservation":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVirtualizationVmwareVmMemoryShareInfo(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVirtualizationVmwareVmMemoryShareInfo(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "virtualization.VmwareVmMemoryShareInfo", "object_type": "virtualization.VmwareVmMemoryShareInfo", "mem_shares": 32, "mem_limit": 32, "mem_overhead_limit": 32, "mem_reservation": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicArfsSettings(t *testing.T) {
	p := models.VnicArfsSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.ArfsSettings","ObjectType":"vnic.ArfsSettings","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicArfsSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicArfsSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.ArfsSettings", "object_type": "vnic.ArfsSettings", "enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicCdn(t *testing.T) {
	p := models.VnicCdn{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.Cdn","ObjectType":"vnic.Cdn","Source":"Source %d","Value":"Value %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicCdn(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicCdn(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.Cdn", "object_type": "vnic.Cdn", "nr_source": "Source 1", "value": "Value 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicCompletionQueueSettings(t *testing.T) {
	p := models.VnicCompletionQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.CompletionQueueSettings","ObjectType":"vnic.CompletionQueueSettings","Count":32,"RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicCompletionQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicCompletionQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.CompletionQueueSettings", "object_type": "vnic.CompletionQueueSettings", "nr_count": 32, "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicEthAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthIfRelationship(t *testing.T) {
	p := models.VnicEthIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthIfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthInterruptSettings(t *testing.T) {
	p := models.VnicEthInterruptSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.EthInterruptSettings","ObjectType":"vnic.EthInterruptSettings","CoalescingTime":32,"CoalescingType":"CoalescingType %d","Count":32,"Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthInterruptSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthInterruptSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.EthInterruptSettings", "object_type": "vnic.EthInterruptSettings", "coalescing_time": 32, "coalescing_type": "CoalescingType 1", "nr_count": 32, "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthNetworkPolicyRelationship(t *testing.T) {
	p := models.VnicEthNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthQosPolicyRelationship(t *testing.T) {
	p := models.VnicEthQosPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthQosPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthQosPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthRxQueueSettings(t *testing.T) {
	p := models.VnicEthRxQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"vnic.EthRxQueueSettings","Count":32,"RingSize":32,"ClassId":"vnic.EthRxQueueSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthRxQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthRxQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "vnic.EthRxQueueSettings", "nr_count": 32, "ring_size": 32, "class_id": "vnic.EthRxQueueSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicEthTxQueueSettings(t *testing.T) {
	p := models.VnicEthTxQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"Count":32,"ClassId":"vnic.EthTxQueueSettings","ObjectType":"vnic.EthTxQueueSettings","RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicEthTxQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicEthTxQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"nr_count": 32, "class_id": "vnic.EthTxQueueSettings", "object_type": "vnic.EthTxQueueSettings", "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicFcAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcErrorRecoverySettings(t *testing.T) {
	p := models.VnicFcErrorRecoverySettings{}
	var d = &schema.ResourceData{}
	c := `{"IoRetryCount":32,"IoRetryTimeout":32,"LinkDownTimeout":32,"PortDownTimeout":32,"ClassId":"vnic.FcErrorRecoverySettings","ObjectType":"vnic.FcErrorRecoverySettings","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcErrorRecoverySettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcErrorRecoverySettings(p, d)[0]
	expectedOp := map[string]interface{}{"io_retry_count": 32, "io_retry_timeout": 32, "link_down_timeout": 32, "port_down_timeout": 32, "class_id": "vnic.FcErrorRecoverySettings", "object_type": "vnic.FcErrorRecoverySettings", "enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcIfRelationship(t *testing.T) {
	p := models.VnicFcIfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcIfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcIfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcInterruptSettings(t *testing.T) {
	p := models.VnicFcInterruptSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FcInterruptSettings","ObjectType":"vnic.FcInterruptSettings","Mode":"Mode %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcInterruptSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcInterruptSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FcInterruptSettings", "object_type": "vnic.FcInterruptSettings", "mode": "Mode 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcNetworkPolicyRelationship(t *testing.T) {
	p := models.VnicFcNetworkPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcNetworkPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcNetworkPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcQosPolicyRelationship(t *testing.T) {
	p := models.VnicFcQosPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcQosPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcQosPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFcQueueSettings(t *testing.T) {
	p := models.VnicFcQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FcQueueSettings","ObjectType":"vnic.FcQueueSettings","Count":32,"RingSize":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFcQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFcQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FcQueueSettings", "object_type": "vnic.FcQueueSettings", "nr_count": 32, "ring_size": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicFlogiSettings(t *testing.T) {
	p := models.VnicFlogiSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.FlogiSettings","ObjectType":"vnic.FlogiSettings","Retries":32,"Timeout":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicFlogiSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicFlogiSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.FlogiSettings", "object_type": "vnic.FlogiSettings", "retries": 32, "timeout": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiAdapterPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiAdapterPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiAdapterPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiAdapterPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiAuthProfile(t *testing.T) {
	p := models.VnicIscsiAuthProfile{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.IscsiAuthProfile","ObjectType":"vnic.IscsiAuthProfile","IsPasswordSet":true,"UserId":"UserId %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiAuthProfile(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiAuthProfile(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.IscsiAuthProfile", "object_type": "vnic.IscsiAuthProfile", "is_password_set": true, "user_id": "UserId 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiBootPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiBootPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiBootPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiBootPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicIscsiStaticTargetPolicyRelationship(t *testing.T) {
	p := models.VnicIscsiStaticTargetPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicIscsiStaticTargetPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicIscsiStaticTargetPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicLanConnectivityPolicyRelationship(t *testing.T) {
	p := models.VnicLanConnectivityPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicLanConnectivityPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicLanConnectivityPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicLun(t *testing.T) {
	p := models.VnicLun{}
	var d = &schema.ResourceData{}
	c := `{"LunId":32,"ClassId":"vnic.Lun","ObjectType":"vnic.Lun","Bootable":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicLun(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicLun(p, d)[0]
	expectedOp := map[string]interface{}{"lun_id": 32, "class_id": "vnic.Lun", "object_type": "vnic.Lun", "bootable": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicNvgreSettings(t *testing.T) {
	p := models.VnicNvgreSettings{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"ClassId":"vnic.NvgreSettings","ObjectType":"vnic.NvgreSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicNvgreSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicNvgreSettings(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "class_id": "vnic.NvgreSettings", "object_type": "vnic.NvgreSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicPlacementSettings(t *testing.T) {
	p := models.VnicPlacementSettings{}
	var d = &schema.ResourceData{}
	c := `{"PciLink":32,"SwitchId":"SwitchId %d","Uplink":32,"ClassId":"vnic.PlacementSettings","ObjectType":"vnic.PlacementSettings","Id":"Id %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicPlacementSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicPlacementSettings(p, d)[0]
	expectedOp := map[string]interface{}{"pci_link": 32, "switch_id": "SwitchId 1", "uplink": 32, "class_id": "vnic.PlacementSettings", "object_type": "vnic.PlacementSettings", "id": "Id 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicPlogiSettings(t *testing.T) {
	p := models.VnicPlogiSettings{}
	var d = &schema.ResourceData{}
	c := `{"Timeout":32,"Retries":32,"ClassId":"vnic.PlogiSettings","ObjectType":"vnic.PlogiSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicPlogiSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicPlogiSettings(p, d)[0]
	expectedOp := map[string]interface{}{"timeout": 32, "retries": 32, "class_id": "vnic.PlogiSettings", "object_type": "vnic.PlogiSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicRoceSettings(t *testing.T) {
	p := models.VnicRoceSettings{}
	var d = &schema.ResourceData{}
	c := `{"Enabled":true,"QueuePairs":32,"ResourceGroups":32,"Version":32,"ClassOfService":32,"ClassId":"vnic.RoceSettings","MemoryRegions":32,"ObjectType":"vnic.RoceSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicRoceSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicRoceSettings(p, d)[0]
	expectedOp := map[string]interface{}{"enabled": true, "queue_pairs": 32, "resource_groups": 32, "nr_version": 32, "class_of_service": 32, "class_id": "vnic.RoceSettings", "memory_regions": 32, "object_type": "vnic.RoceSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicRssHashSettings(t *testing.T) {
	p := models.VnicRssHashSettings{}
	var d = &schema.ResourceData{}
	c := `{"TcpIpv4Hash":true,"TcpIpv6ExtHash":true,"TcpIpv6Hash":true,"Ipv4Hash":true,"UdpIpv4Hash":true,"Ipv6ExtHash":true,"UdpIpv6Hash":true,"ClassId":"vnic.RssHashSettings","ObjectType":"vnic.RssHashSettings","Ipv6Hash":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicRssHashSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicRssHashSettings(p, d)[0]
	expectedOp := map[string]interface{}{"tcp_ipv4_hash": true, "tcp_ipv6_ext_hash": true, "tcp_ipv6_hash": true, "ipv4_hash": true, "udp_ipv4_hash": true, "ipv6_ext_hash": true, "udp_ipv6_hash": true, "class_id": "vnic.RssHashSettings", "object_type": "vnic.RssHashSettings", "ipv6_hash": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicSanConnectivityPolicyRelationship(t *testing.T) {
	p := models.VnicSanConnectivityPolicyRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicSanConnectivityPolicyRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicSanConnectivityPolicyRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicScsiQueueSettings(t *testing.T) {
	p := models.VnicScsiQueueSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.ScsiQueueSettings","ObjectType":"vnic.ScsiQueueSettings","RingSize":32,"Count":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicScsiQueueSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicScsiQueueSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.ScsiQueueSettings", "object_type": "vnic.ScsiQueueSettings", "ring_size": 32, "nr_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicTcpOffloadSettings(t *testing.T) {
	p := models.VnicTcpOffloadSettings{}
	var d = &schema.ResourceData{}
	c := `{"TxChecksum":true,"LargeReceive":true,"LargeSend":true,"RxChecksum":true,"ClassId":"vnic.TcpOffloadSettings","ObjectType":"vnic.TcpOffloadSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicTcpOffloadSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicTcpOffloadSettings(p, d)[0]
	expectedOp := map[string]interface{}{"tx_checksum": true, "large_receive": true, "large_send": true, "rx_checksum": true, "class_id": "vnic.TcpOffloadSettings", "object_type": "vnic.TcpOffloadSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicUsnicSettings(t *testing.T) {
	p := models.VnicUsnicSettings{}
	var d = &schema.ResourceData{}
	c := `{"UsnicAdapterPolicy":"UsnicAdapterPolicy %d","ClassId":"vnic.UsnicSettings","ObjectType":"vnic.UsnicSettings","Cos":32,"Count":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicUsnicSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicUsnicSettings(p, d)[0]
	expectedOp := map[string]interface{}{"usnic_adapter_policy": "UsnicAdapterPolicy 1", "class_id": "vnic.UsnicSettings", "object_type": "vnic.UsnicSettings", "cos": 32, "nr_count": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVlanSettings(t *testing.T) {
	p := models.VnicVlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"AllowedVlans":"AllowedVlans %d","DefaultVlan":32,"Mode":"Mode %d","ClassId":"vnic.VlanSettings","ObjectType":"vnic.VlanSettings"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"allowed_vlans": "AllowedVlans 1", "default_vlan": 32, "mode": "Mode 1", "class_id": "vnic.VlanSettings", "object_type": "vnic.VlanSettings"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVmqSettings(t *testing.T) {
	p := models.VnicVmqSettings{}
	var d = &schema.ResourceData{}
	c := `{"VmmqAdapterPolicy":"VmmqAdapterPolicy %d","NumInterrupts":32,"NumVmqs":32,"NumSubVnics":32,"ClassId":"vnic.VmqSettings","ObjectType":"vnic.VmqSettings","Enabled":true,"MultiQueueSupport":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVmqSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVmqSettings(p, d)[0]
	expectedOp := map[string]interface{}{"vmmq_adapter_policy": "VmmqAdapterPolicy 1", "num_interrupts": 32, "num_vmqs": 32, "num_sub_vnics": 32, "class_id": "vnic.VmqSettings", "object_type": "vnic.VmqSettings", "enabled": true, "multi_queue_support": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVsanSettings(t *testing.T) {
	p := models.VnicVsanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.VsanSettings","ObjectType":"vnic.VsanSettings","DefaultVlanId":32,"Id":32}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVsanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVsanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.VsanSettings", "object_type": "vnic.VsanSettings", "default_vlan_id": 32, "id": 32}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVnicVxlanSettings(t *testing.T) {
	p := models.VnicVxlanSettings{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"vnic.VxlanSettings","ObjectType":"vnic.VxlanSettings","Enabled":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapVnicVxlanSettings(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVnicVxlanSettings(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "vnic.VxlanSettings", "object_type": "vnic.VxlanSettings", "enabled": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapVrfVrfRelationship(t *testing.T) {
	p := models.VrfVrfRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapVrfVrfRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapVrfVrfRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCatalogRelationship(t *testing.T) {
	p := models.WorkflowCatalogRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCatalogRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCatalogRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowComments(t *testing.T) {
	p := models.WorkflowComments{}
	var d = &schema.ResourceData{}
	c := `{"Description":"Description %d","ClassId":"workflow.Comments","ObjectType":"workflow.Comments"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowComments(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowComments(p, d)[0]
	expectedOp := map[string]interface{}{"description": "Description 1", "class_id": "workflow.Comments", "object_type": "workflow.Comments"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCustomDataTypeDefinitionRelationship(t *testing.T) {
	p := models.WorkflowCustomDataTypeDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCustomDataTypeDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCustomDataTypeDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowCustomDataTypeProperties(t *testing.T) {
	p := models.WorkflowCustomDataTypeProperties{}
	var d = &schema.ResourceData{}
	c := `{"Cloneable":true,"ExternalMeta":true,"ClassId":"workflow.CustomDataTypeProperties","ObjectType":"workflow.CustomDataTypeProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowCustomDataTypeProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowCustomDataTypeProperties(p, d)[0]
	expectedOp := map[string]interface{}{"cloneable": true, "external_meta": true, "class_id": "workflow.CustomDataTypeProperties", "object_type": "workflow.CustomDataTypeProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowErrorResponseHandlerRelationship(t *testing.T) {
	p := models.WorkflowErrorResponseHandlerRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowErrorResponseHandlerRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowErrorResponseHandlerRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowInternalProperties(t *testing.T) {
	p := models.WorkflowInternalProperties{}
	var d = &schema.ResourceData{}
	c := `{"BaseTaskType":"BaseTaskType %d","Internal":true,"Owner":"Owner %d","ClassId":"workflow.InternalProperties","ObjectType":"workflow.InternalProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowInternalProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowInternalProperties(p, d)[0]
	expectedOp := map[string]interface{}{"base_task_type": "BaseTaskType 1", "internal": true, "owner": "Owner 1", "class_id": "workflow.InternalProperties", "object_type": "workflow.InternalProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowPendingDynamicWorkflowInfoRelationship(t *testing.T) {
	p := models.WorkflowPendingDynamicWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowPendingDynamicWorkflowInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowProperties(t *testing.T) {
	p := models.WorkflowProperties{}
	var d = &schema.ResourceData{}
	c := `{"TimeoutPolicy":"TimeoutPolicy %d","RetryCount":32,"Cloneable":true,"RetryDelay":32,"SupportStatus":"SupportStatus %d","Timeout":32,"ClassId":"workflow.Properties","StartsWorkflow":true,"ObjectType":"workflow.Properties","RetryPolicy":"RetryPolicy %d","ExternalMeta":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowProperties(p, d)[0]
	expectedOp := map[string]interface{}{"timeout_policy": "TimeoutPolicy 1", "retry_count": 32, "cloneable": true, "retry_delay": 32, "support_status": "SupportStatus 1", "timeout": 32, "class_id": "workflow.Properties", "starts_workflow": true, "object_type": "workflow.Properties", "retry_policy": "RetryPolicy 1", "external_meta": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowSolutionActionDefinitionRelationship(t *testing.T) {
	p := models.WorkflowSolutionActionDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowSolutionActionDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowSolutionActionDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowSolutionDefinitionRelationship(t *testing.T) {
	p := models.WorkflowSolutionDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowSolutionDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowSolutionDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowSolutionInstanceRelationship(t *testing.T) {
	p := models.WorkflowSolutionInstanceRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowSolutionInstanceRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowSolutionInstanceRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskConstraints(t *testing.T) {
	p := models.WorkflowTaskConstraints{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.TaskConstraints","ObjectType":"workflow.TaskConstraints"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskConstraints(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskConstraints(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.TaskConstraints", "object_type": "workflow.TaskConstraints"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskDefinitionRelationship(t *testing.T) {
	p := models.WorkflowTaskDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskInfoRelationship(t *testing.T) {
	p := models.WorkflowTaskInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowTaskMetadataRelationship(t *testing.T) {
	p := models.WorkflowTaskMetadataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"Moid":"Moid %d","Selector":"Selector %d","ClassId":"mo.MoRef","ObjectType":"mo.MoRef"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowTaskMetadataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowTaskMetadataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"moid": "Moid 1", "selector": "Selector 1", "class_id": "mo.MoRef", "object_type": "mo.MoRef"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowValidationInformation(t *testing.T) {
	p := models.WorkflowValidationInformation{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.ValidationInformation","ObjectType":"workflow.ValidationInformation","State":"State %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowValidationInformation(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowValidationInformation(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.ValidationInformation", "object_type": "workflow.ValidationInformation", "state": "State 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowCtx(t *testing.T) {
	p := models.WorkflowWorkflowCtx{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"workflow.WorkflowCtx","ObjectType":"workflow.WorkflowCtx","WorkflowType":"WorkflowType %d","WorkflowMetaName":"WorkflowMetaName %d","WorkflowSubtype":"WorkflowSubtype %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowCtx(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowCtx(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "workflow.WorkflowCtx", "object_type": "workflow.WorkflowCtx", "workflow_type": "WorkflowType 1", "workflow_meta_name": "WorkflowMetaName 1", "workflow_subtype": "WorkflowSubtype 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowDefinitionRelationship(t *testing.T) {
	p := models.WorkflowWorkflowDefinitionRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowDefinitionRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowDefinitionRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowInfoRelationship(t *testing.T) {
	p := models.WorkflowWorkflowInfoRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowInfoRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowInfoRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowInfoProperties(t *testing.T) {
	p := models.WorkflowWorkflowInfoProperties{}
	var d = &schema.ResourceData{}
	c := `{"Retryable":true,"RollbackAction":"RollbackAction %d","ClassId":"workflow.WorkflowInfoProperties","ObjectType":"workflow.WorkflowInfoProperties"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowInfoProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowInfoProperties(p, d)[0]
	expectedOp := map[string]interface{}{"retryable": true, "rollback_action": "RollbackAction 1", "class_id": "workflow.WorkflowInfoProperties", "object_type": "workflow.WorkflowInfoProperties"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowMetadataRelationship(t *testing.T) {
	p := models.WorkflowWorkflowMetadataRelationship{}
	var d = &schema.ResourceData{}
	c := `{"ClassId":"mo.MoRef","ObjectType":"mo.MoRef","Moid":"Moid %d","Selector":"Selector %d"}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowMetadataRelationship(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowMetadataRelationship(p, d)[0]
	expectedOp := map[string]interface{}{"class_id": "mo.MoRef", "object_type": "mo.MoRef", "moid": "Moid 1", "selector": "Selector 1"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapWorkflowWorkflowProperties(t *testing.T) {
	p := models.WorkflowWorkflowProperties{}
	var d = &schema.ResourceData{}
	c := `{"EnableDebug":true,"ExternalMeta":true,"Retryable":true,"ClassId":"workflow.WorkflowProperties","ObjectType":"workflow.WorkflowProperties","SupportStatus":"SupportStatus %d","Cloneable":true}`

	//test when the response is empty
	ffOpEmpty := flattenMapWorkflowWorkflowProperties(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapWorkflowWorkflowProperties(p, d)[0]
	expectedOp := map[string]interface{}{"enable_debug": true, "external_meta": true, "retryable": true, "class_id": "workflow.WorkflowProperties", "object_type": "workflow.WorkflowProperties", "support_status": "SupportStatus 1", "cloneable": true}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
func TestFlattenMapX509Certificate(t *testing.T) {
	p := models.X509Certificate{}
	var d = &schema.ResourceData{}
	c := `{"PemCertificate":"PemCertificate %d","ObjectType":"x509.Certificate","Sha256Fingerprint":"Sha256Fingerprint %d","SignatureAlgorithm":"SignatureAlgorithm %d","ClassId":"x509.Certificate"}`

	//test when the response is empty
	ffOpEmpty := flattenMapX509Certificate(p, d)
	if len(ffOpEmpty) != 0 {
		t.Errorf("error: no elements should be present. Found %d elements", len(ffOpEmpty))
	}
	// test when response is available and resourceData is empty
	err := p.UnmarshalJSON([]byte(strings.Replace(c, "%d", "1", -1)))
	CheckError(t, err)
	ffOp := flattenMapX509Certificate(p, d)[0]
	expectedOp := map[string]interface{}{"pem_certificate": "PemCertificate 1", "object_type": "x509.Certificate", "sha256_fingerprint": "Sha256Fingerprint 1", "signature_algorithm": "SignatureAlgorithm 1", "class_id": "x509.Certificate"}
	err = compareMaps(expectedOp, ffOp, t)
	CheckError(t, err)
}
