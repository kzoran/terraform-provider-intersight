package intersight

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	models "github.com/CiscoDevNet/terraform-provider-intersight/intersight_gosdk"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHyperflexVmSnapshotInfo() *schema.Resource {
	var subSchema = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster_id_snap_map": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_id": {
						Description: "ClusterId of the snapshot point.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"snapshot_point": {
						Description: "Snapshot point details for this cluster.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"cluster_entity_reference": {
									Description: "Cluster to which this snapshot belongs.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"confignum": {
												Description: "Configuration number for this reference.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"id": {
												Description: "Uuid of the entity for this reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"idtype": {
												Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name": {
												Description: "Name of the entity for this entity reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"type": {
												Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"replication_status": {
									Description: "Replication status information for this particular snapshot.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"bytes_replicated": {
												Description: "Number of bytes currently replicated.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"end_time": {
												Description: "Replication end time for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"error": {
												Description: "Error information in case of failure.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message": {
															Description: "The error message string for this error stack.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message_id": {
															Description: "The error message ID for this error stack.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"pack_entity_reference": {
												Description: "EntityReference for the Replication Pack.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"confignum": {
															Description: "Configuration number for this reference.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"id": {
															Description: "Uuid of the entity for this reference.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"idtype": {
															Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"name": {
															Description: "Name of the entity for this entity reference.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"type": {
															Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"pct_complete": {
												Description: "Completion percentage for the replication job.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"rpo_status": {
												Description: "Status for timeliness of replication job in relation to the schedule interval.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"actual": {
															Description: "Actual end time for the snapshot.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"expected": {
															Description: "Expected end time for the snapshot.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"rpo_exceeded": {
															Description: "A flag to determine if snapshot delivery is delayed.",
															Type:        schema.TypeBool,
															Optional:    true,
														},
													},
												},
											},
											"start_time": {
												Description: "Replication start time for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"status": {
												Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"snapshot_files": {
									Description: "Files this snapshot is comprised of.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name_tracked_files": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"ds_info": {
															Description: "Datastore for which this file resides.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"ds_backend_id": {
																		Description: "This datastore's backend unique id.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"ds_frontend_id": {
																		Description: "This datastore's frontend id.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"relative_file_path": {
															Description: "Relative file path within the datastore.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"uuid_tracked_disks_map": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"tracked_disk": {
															Description: "Tracked Disk for a snapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"disk_files": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"file_path": {
																					Description: "File path and datastore information.",
																					Type:        schema.TypeList,
																					MaxItems:    1,
																					Optional:    true,
																					Elem: &schema.Resource{
																						Schema: map[string]*schema.Schema{
																							"additional_properties": {
																								Type:             schema.TypeString,
																								Optional:         true,
																								DiffSuppressFunc: SuppressDiffAdditionProps,
																							},
																							"class_id": {
																								Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																							"ds_info": {
																								Description: "Datastore for which this file resides.",
																								Type:        schema.TypeList,
																								MaxItems:    1,
																								Optional:    true,
																								Elem: &schema.Resource{
																									Schema: map[string]*schema.Schema{
																										"additional_properties": {
																											Type:             schema.TypeString,
																											Optional:         true,
																											DiffSuppressFunc: SuppressDiffAdditionProps,
																										},
																										"class_id": {
																											Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"ds_backend_id": {
																											Description: "This datastore's backend unique id.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"ds_frontend_id": {
																											Description: "This datastore's frontend id.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"object_type": {
																											Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																									},
																								},
																							},
																							"object_type": {
																								Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																							"relative_file_path": {
																								Description: "Relative file path within the datastore.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																						},
																					},
																				},
																				"file_type": {
																					Description: "File type for the tracked file.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																			},
																		},
																	},
																	"disk_type": {
																		Description: "Disk type for a vm virtual disk.\n* `NONE` - The disk type for this VM is None.\n* `NATIVE` - The disk type for this VM is Native.\n* `NONNATIVE` - The disk type for this VM is Non-Native.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
														},
														"uuid": {
															Description: "Disk unique id for a snapshot.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
										},
									},
								},
								"snapshot_point_entity_reference": {
									Description: "EntityReference for this snapshot point.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"confignum": {
												Description: "Configuration number for this reference.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"id": {
												Description: "Uuid of the entity for this reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"idtype": {
												Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name": {
												Description: "Name of the entity for this entity reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"type": {
												Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"snapshot_status": {
									Description: "Status for this Snapshot Point.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"description": {
												Description: "Description of this Snapshot Point.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"error": {
												Description: "Error information in case of failure.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message": {
															Description: "The error message string for this error stack.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message_id": {
															Description: "The error message ID for this error stack.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"pct_complete": {
												Description: "Completion percentage for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"status": {
												Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.\n* `INIT` - This snapshot status code is initializing.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"timestamp": {
												Description: "Timestamp at which the Snapshot is taken.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"used_space": {
												Description: "Space Used by this Snapshot Point.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
										},
									},
								},
								"vm_runtime_info": {
									Description: "Virtual machine runtime information.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"run_time_info": {
												Description: "Virtual machine runtime details.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"bios_uuid": {
															Description: "BiosUuid of the Protected Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"connection_state": {
															Description: "Connection state of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"cpu_usage": {
															Description: "CPU Usage of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"folder": {
															Description: "Folder which VM belongs to.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_family": {
															Description: "Guest operating system family, if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_full_name": {
															Description: "Guest operating system full name, if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_id": {
															Description: "Guest operating system identifier (short name), if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_state": {
															Description: "VMware guest reset, powercycle, or boot action states.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"host_name": {
															Description: "Hostname of Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"instance_uuid": {
															Description: "InstanceUuid of the Protected Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"memory_mb": {
															Description: "CPU Memory in MB of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"memory_usage": {
															Description: "Memory usage of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"moid": {
															Description: "Virtual Machine unique MOID.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"name": {
															Description: "Name of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"networks": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"num_cpu": {
															Description: "Number of CPUs for the VM.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"power_state": {
															Description: "Power state of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"provisioned_size": {
															Description: "Provisioned Size of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"resource_pool": {
															Description: "Resource pool which VM belongs to.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"used_size": {
															Description: "Used Size of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"nr_version": {
															Description: "Version of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"vmx_path": {
															Description: "Vmx Path in VC datastore format.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"status_code": {
												Description: "Virtual machine status code.\n* `VM_ACCESSIBLE` - This virtual machine is accessible.\n* `VM_INACCESSIBLE` - This virtual machine is not accessible.\n* `VM_NOT_SUPPORTED` - This virtual machine is not supported.\n* `NONE` - This virtual machine does not have a status code.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"uuid": {
												Description: "Virtual machine's current UUID.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"display_status": {
			Description: "Combined status code from replication and snapshot status to display in the Intersight UI.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"error": {
			Description: "List of errors associated with this VmBackupInfo.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"message": {
						Description: "The error message string for this error stack.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"message_id": {
						Description: "The error message ID for this error stack.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"label": {
			Description: "The name of the Virtual Machine and the time stamp of the snapshot.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mode": {
			Description: "Quiesce Mode for snapshot taken on Hyperflex cluster.\n* `NONE` - The snapshot quiesce mode is none.\n* `CRASH` - The snapshot quiesce mode is crash.\n* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.\n* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"parent_snapshot": {
			Description: "Entity reference to the parent snapshot of this VmSnapshotInfo.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"replication_status": {
			Description: "Replication status of the least successful target cluster.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snapshot_error_msg": {
			Description: "Error message from snapshot creation or replcation if any exist.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snapshot_status": {
			Description: "Snapshot status of the source cluster.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.\n* `INIT` - This snapshot status code is initializing.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_timestamp": {
			Description: "Timestamp the snapshot was created on the source cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"src_cluster": {
			Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"src_cluster_name": {
			Description: "Name of the cluster this snapshot resides on.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"target_completion_timestamp": {
			Description: "Timestamp the snapshot was finished replicating on the target cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tgt_cluster": {
			Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tgt_cluster_name": {
			Description: "Name of the cluster this snapshot is replicated to.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_backup_info": {
			Description: "A reference to a hyperflexVmBackupInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_entity_reference": {
			Description: "Entity reference to the virtual machine this snapshot info is attached to.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_snapshot_entity_reference": {
			Description: "Entity reference to the snapshot this snapshot info is attached to.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	var model = map[string]*schema.Schema{"account_moid": {
		Description: "The Account ID for this managed object.",
		Type:        schema.TypeString,
		Optional:    true,
	},
		"additional_properties": {
			Type:             schema.TypeString,
			Optional:         true,
			DiffSuppressFunc: SuppressDiffAdditionProps,
		},
		"ancestors": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"class_id": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"cluster_id_snap_map": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"cluster_id": {
						Description: "ClusterId of the snapshot point.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"snapshot_point": {
						Description: "Snapshot point details for this cluster.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"cluster_entity_reference": {
									Description: "Cluster to which this snapshot belongs.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"confignum": {
												Description: "Configuration number for this reference.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"id": {
												Description: "Uuid of the entity for this reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"idtype": {
												Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name": {
												Description: "Name of the entity for this entity reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"type": {
												Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"object_type": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"replication_status": {
									Description: "Replication status information for this particular snapshot.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"bytes_replicated": {
												Description: "Number of bytes currently replicated.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"end_time": {
												Description: "Replication end time for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"error": {
												Description: "Error information in case of failure.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message": {
															Description: "The error message string for this error stack.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message_id": {
															Description: "The error message ID for this error stack.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"pack_entity_reference": {
												Description: "EntityReference for the Replication Pack.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"confignum": {
															Description: "Configuration number for this reference.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"id": {
															Description: "Uuid of the entity for this reference.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"idtype": {
															Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"name": {
															Description: "Name of the entity for this entity reference.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"type": {
															Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"pct_complete": {
												Description: "Completion percentage for the replication job.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"rpo_status": {
												Description: "Status for timeliness of replication job in relation to the schedule interval.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"actual": {
															Description: "Actual end time for the snapshot.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"expected": {
															Description: "Expected end time for the snapshot.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"rpo_exceeded": {
															Description: "A flag to determine if snapshot delivery is delayed.",
															Type:        schema.TypeBool,
															Optional:    true,
														},
													},
												},
											},
											"start_time": {
												Description: "Replication start time for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"status": {
												Description: "Current replication state for a particular snapshot.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"snapshot_files": {
									Description: "Files this snapshot is comprised of.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name_tracked_files": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"ds_info": {
															Description: "Datastore for which this file resides.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"ds_backend_id": {
																		Description: "This datastore's backend unique id.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"ds_frontend_id": {
																		Description: "This datastore's frontend id.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"relative_file_path": {
															Description: "Relative file path within the datastore.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"uuid_tracked_disks_map": {
												Type:     schema.TypeList,
												Optional: true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"tracked_disk": {
															Description: "Tracked Disk for a snapshot.",
															Type:        schema.TypeList,
															MaxItems:    1,
															Optional:    true,
															Elem: &schema.Resource{
																Schema: map[string]*schema.Schema{
																	"additional_properties": {
																		Type:             schema.TypeString,
																		Optional:         true,
																		DiffSuppressFunc: SuppressDiffAdditionProps,
																	},
																	"class_id": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"disk_files": {
																		Type:     schema.TypeList,
																		Optional: true,
																		Elem: &schema.Resource{
																			Schema: map[string]*schema.Schema{
																				"additional_properties": {
																					Type:             schema.TypeString,
																					Optional:         true,
																					DiffSuppressFunc: SuppressDiffAdditionProps,
																				},
																				"class_id": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"file_path": {
																					Description: "File path and datastore information.",
																					Type:        schema.TypeList,
																					MaxItems:    1,
																					Optional:    true,
																					Elem: &schema.Resource{
																						Schema: map[string]*schema.Schema{
																							"additional_properties": {
																								Type:             schema.TypeString,
																								Optional:         true,
																								DiffSuppressFunc: SuppressDiffAdditionProps,
																							},
																							"class_id": {
																								Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																							"ds_info": {
																								Description: "Datastore for which this file resides.",
																								Type:        schema.TypeList,
																								MaxItems:    1,
																								Optional:    true,
																								Elem: &schema.Resource{
																									Schema: map[string]*schema.Schema{
																										"additional_properties": {
																											Type:             schema.TypeString,
																											Optional:         true,
																											DiffSuppressFunc: SuppressDiffAdditionProps,
																										},
																										"class_id": {
																											Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"ds_backend_id": {
																											Description: "This datastore's backend unique id.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"ds_frontend_id": {
																											Description: "This datastore's frontend id.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																										"object_type": {
																											Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																											Type:        schema.TypeString,
																											Optional:    true,
																										},
																									},
																								},
																							},
																							"object_type": {
																								Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																							"relative_file_path": {
																								Description: "Relative file path within the datastore.",
																								Type:        schema.TypeString,
																								Optional:    true,
																							},
																						},
																					},
																				},
																				"file_type": {
																					Description: "File type for the tracked file.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																				"object_type": {
																					Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																					Type:        schema.TypeString,
																					Optional:    true,
																				},
																			},
																		},
																	},
																	"disk_type": {
																		Description: "Disk type for a vm virtual disk.\n* `NONE` - The disk type for this VM is None.\n* `NATIVE` - The disk type for this VM is Native.\n* `NONNATIVE` - The disk type for this VM is Non-Native.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																	"object_type": {
																		Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
																		Type:        schema.TypeString,
																		Optional:    true,
																	},
																},
															},
														},
														"uuid": {
															Description: "Disk unique id for a snapshot.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
										},
									},
								},
								"snapshot_point_entity_reference": {
									Description: "EntityReference for this snapshot point.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"confignum": {
												Description: "Configuration number for this reference.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"id": {
												Description: "Uuid of the entity for this reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"idtype": {
												Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"name": {
												Description: "Name of the entity for this entity reference.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"type": {
												Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
								"snapshot_status": {
									Description: "Status for this Snapshot Point.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"description": {
												Description: "Description of this Snapshot Point.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"error": {
												Description: "Error information in case of failure.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message": {
															Description: "The error message string for this error stack.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"message_id": {
															Description: "The error message ID for this error stack.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"pct_complete": {
												Description: "Completion percentage for this snapshot.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"status": {
												Description: "Current snapshot state for this snapshot.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.\n* `INIT` - This snapshot status code is initializing.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"timestamp": {
												Description: "Timestamp at which the Snapshot is taken.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
											"used_space": {
												Description: "Space Used by this Snapshot Point.",
												Type:        schema.TypeInt,
												Optional:    true,
											},
										},
									},
								},
								"vm_runtime_info": {
									Description: "Virtual machine runtime information.",
									Type:        schema.TypeList,
									MaxItems:    1,
									Optional:    true,
									Elem: &schema.Resource{
										Schema: map[string]*schema.Schema{
											"additional_properties": {
												Type:             schema.TypeString,
												Optional:         true,
												DiffSuppressFunc: SuppressDiffAdditionProps,
											},
											"class_id": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"object_type": {
												Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"run_time_info": {
												Description: "Virtual machine runtime details.",
												Type:        schema.TypeList,
												MaxItems:    1,
												Optional:    true,
												Elem: &schema.Resource{
													Schema: map[string]*schema.Schema{
														"additional_properties": {
															Type:             schema.TypeString,
															Optional:         true,
															DiffSuppressFunc: SuppressDiffAdditionProps,
														},
														"bios_uuid": {
															Description: "BiosUuid of the Protected Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"class_id": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"connection_state": {
															Description: "Connection state of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"cpu_usage": {
															Description: "CPU Usage of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"folder": {
															Description: "Folder which VM belongs to.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_family": {
															Description: "Guest operating system family, if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_full_name": {
															Description: "Guest operating system full name, if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_id": {
															Description: "Guest operating system identifier (short name), if known.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"guest_state": {
															Description: "VMware guest reset, powercycle, or boot action states.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"host_name": {
															Description: "Hostname of Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"instance_uuid": {
															Description: "InstanceUuid of the Protected Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"memory_mb": {
															Description: "CPU Memory in MB of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"memory_usage": {
															Description: "Memory usage of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"moid": {
															Description: "Virtual Machine unique MOID.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"name": {
															Description: "Name of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"networks": {
															Type:     schema.TypeList,
															Optional: true,
															Elem: &schema.Schema{
																Type: schema.TypeString}},
														"num_cpu": {
															Description: "Number of CPUs for the VM.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"object_type": {
															Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"power_state": {
															Description: "Power state of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"provisioned_size": {
															Description: "Provisioned Size of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"resource_pool": {
															Description: "Resource pool which VM belongs to.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"used_size": {
															Description: "Used Size of Virtual Machine.",
															Type:        schema.TypeInt,
															Optional:    true,
														},
														"nr_version": {
															Description: "Version of the Virtual Machine.",
															Type:        schema.TypeString,
															Optional:    true,
														},
														"vmx_path": {
															Description: "Vmx Path in VC datastore format.",
															Type:        schema.TypeString,
															Optional:    true,
														},
													},
												},
											},
											"status_code": {
												Description: "Virtual machine status code.\n* `VM_ACCESSIBLE` - This virtual machine is accessible.\n* `VM_INACCESSIBLE` - This virtual machine is not accessible.\n* `VM_NOT_SUPPORTED` - This virtual machine is not supported.\n* `NONE` - This virtual machine does not have a status code.",
												Type:        schema.TypeString,
												Optional:    true,
											},
											"uuid": {
												Description: "Virtual machine's current UUID.",
												Type:        schema.TypeString,
												Optional:    true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
		"create_time": {
			Description: "The time when this managed object was created.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"display_status": {
			Description: "Combined status code from replication and snapshot status to display in the Intersight UI.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"domain_group_moid": {
			Description: "The DomainGroup ID for this managed object.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"error": {
			Description: "List of errors associated with this VmBackupInfo.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"message": {
						Description: "The error message string for this error stack.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"message_id": {
						Description: "The error message ID for this error stack.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"label": {
			Description: "The name of the Virtual Machine and the time stamp of the snapshot.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mod_time": {
			Description: "The time when this managed object was last modified.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"mode": {
			Description: "Quiesce Mode for snapshot taken on Hyperflex cluster.\n* `NONE` - The snapshot quiesce mode is none.\n* `CRASH` - The snapshot quiesce mode is crash.\n* `VMTOOLS` - The snapshot quiesce mode is VMTOOLS.\n* `APP_CONSISTENT` - The snapshot quiesce mode is app consistent.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"moid": {
			Description: "The unique identifier of this Managed Object instance.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"object_type": {
			Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"owners": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Schema{
				Type: schema.TypeString}},
		"parent": {
			Description: "A reference to a moBaseMo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"parent_snapshot": {
			Description: "Entity reference to the parent snapshot of this VmSnapshotInfo.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"permission_resources": {
			Description: "An array of relationships to moBaseMo resources.",
			Type:        schema.TypeList,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"replication_status": {
			Description: "Replication status of the least successful target cluster.\n* `NONE` - Snapshot replication state is none.\n* `SUCCESS` - Snapshot completed successfully.\n* `FAILED` - Snapshot failed replication status code.\n* `PAUSED` - Snapshot replication paused status code.\n* `IN_USE` - Snapshot replica in use status code.\n* `STARTING` - Snapshot replication starting.\n* `REPLICATING` - Snapshot replication in progress.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"shared_scope": {
			Description: "Intersight provides pre-built workflows, tasks and policies to end users through global catalogs.\nObjects that are made available through global catalogs are said to have a 'shared' ownership. Shared objects are either made globally available to all end users or restricted to end users based on their license entitlement. Users can use this property to differentiate the scope (global or a specific license tier) to which a shared MO belongs.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snapshot_error_msg": {
			Description: "Error message from snapshot creation or replcation if any exist.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"snapshot_status": {
			Description: "Snapshot status of the source cluster.\n* `SUCCESS` - This snapshot status code is success.\n* `FAILED` - This snapshot status code is failed.\n* `IN_PROGRESS` - This snapshot status code is in progress.\n* `DELETING` - This snapshot status code is deleting.\n* `DELETED` - This snapshot status code is deleted.\n* `NONE` - This snapshot status code is none.\n* `INIT` - This snapshot status code is initializing.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"source_timestamp": {
			Description: "Timestamp the snapshot was created on the source cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"src_cluster": {
			Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"src_cluster_name": {
			Description: "Name of the cluster this snapshot resides on.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"tags": {
			Type:     schema.TypeList,
			Optional: true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"key": {
						Description: "The string representation of a tag key.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"value": {
						Description: "The string representation of a tag value.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"target_completion_timestamp": {
			Description: "Timestamp the snapshot was finished replicating on the target cluster.",
			Type:        schema.TypeInt,
			Optional:    true,
		},
		"tgt_cluster": {
			Description: "A reference to a hyperflexCluster resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"tgt_cluster_name": {
			Description: "Name of the cluster this snapshot is replicated to.",
			Type:        schema.TypeString,
			Optional:    true,
		},
		"version_context": {
			Description: "The versioning info for this managed object.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"interested_mos": {
						Type:     schema.TypeList,
						Optional: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"ref_mo": {
						Description: "A reference to the original Managed Object.",
						Type:        schema.TypeList,
						MaxItems:    1,
						Optional:    true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"additional_properties": {
									Type:             schema.TypeString,
									Optional:         true,
									DiffSuppressFunc: SuppressDiffAdditionProps,
								},
								"class_id": {
									Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"moid": {
									Description: "The Moid of the referenced REST resource.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"object_type": {
									Description: "The fully-qualified name of the remote type referred by this relationship.",
									Type:        schema.TypeString,
									Optional:    true,
								},
								"selector": {
									Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
									Type:        schema.TypeString,
									Optional:    true,
								},
							},
						},
					},
					"timestamp": {
						Description: "The time this versioned Managed Object was created.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"nr_version": {
						Description: "The version of the Managed Object, e.g. an incrementing number or a hash id.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"version_type": {
						Description: "Specifies type of version. Currently the only supported value is \"Configured\"\nthat is used to keep track of snapshots of policies and profiles that are intended\nto be configured to target endpoints.\n* `Modified` - Version created every time an object is modified.\n* `Configured` - Version created every time an object is configured to the service profile.\n* `Deployed` - Version created for objects related to a service profile when it is deployed.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_backup_info": {
			Description: "A reference to a hyperflexVmBackupInfo resource.\nWhen the $expand query parameter is specified, the referenced resource is returned inline.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"moid": {
						Description: "The Moid of the referenced REST resource.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the remote type referred by this relationship.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"selector": {
						Description: "An OData $filter expression which describes the REST resource to be referenced. This field may\nbe set instead of 'moid' by clients.\n1. If 'moid' is set this field is ignored.\n1. If 'selector' is set and 'moid' is empty/absent from the request, Intersight determines the Moid of the\nresource matching the filter expression and populates it in the MoRef that is part of the object\ninstance being inserted/updated to fulfill the REST request.\nAn error is returned if the filter matches zero or more than one REST resource.\nAn example filter string is: Serial eq '3AA8B7T11'.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_entity_reference": {
			Description: "Entity reference to the virtual machine this snapshot info is attached to.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
		"vm_snapshot_entity_reference": {
			Description: "Entity reference to the snapshot this snapshot info is attached to.",
			Type:        schema.TypeList,
			MaxItems:    1,
			Optional:    true,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"additional_properties": {
						Type:             schema.TypeString,
						Optional:         true,
						DiffSuppressFunc: SuppressDiffAdditionProps,
					},
					"class_id": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThis property is used as a discriminator to identify the type of the payload\nwhen marshaling and unmarshaling data.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"confignum": {
						Description: "Configuration number for this reference.",
						Type:        schema.TypeInt,
						Optional:    true,
					},
					"id": {
						Description: "Uuid of the entity for this reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"idtype": {
						Description: "Type of entity id for this reference.\n* `VCMOID` - The entity reference ID type is VC MOID.\n* `VMBIOSUUID` - The entity reference ID type is VM Bios UUID.\n* `VMDSPATH` - The entity reference ID type is VM Datastore Path.\n* `VMINSTANCEUUID` - The entity reference ID type is VM Instance UUID.\n* `VMNAME` - The entity reference ID type is VM Name.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"name": {
						Description: "Name of the entity for this entity reference.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"object_type": {
						Description: "The fully-qualified name of the instantiated, concrete type.\nThe value should be the same as the 'ClassId' property.",
						Type:        schema.TypeString,
						Optional:    true,
					},
					"type": {
						Description: "Type of the entity for this entity reference.\n* `DISK` - This entity type is a disk.\n* `PNODE` - This entity type is a P node.\n* `NODE` - This entity type is a node.\n* `CLUSTER` - This entity type is a cluster.\n* `DATASTORE` - This entity is a datastore.\n* `VIRTNODE` - This entity is a HyperFlex virtual node.\n* `VIRTCLUSTER` - This entity type is a virtual cluster.\n* `VIRTDATASTORE` - This entity type is a virtual data store.\n* `VIRTMACHINE` - This entity type is a virtual machine.\n* `PDISK` - This entity type is a P disk.\n* `PDATASTORE` - This entity type is a P Datastore.\n* `VIRTMACHINESNAPSHOT` - This entity is a virtual machine snapshot.\n* `FOLDER` - This entity type is a folder.\n* `RESOURCEPOOL` - This entity type is a resource pool.\n* `FILE` - This entity type is a file.\n* `VIRTDATACENTER` - This entity type is a virtual data center.\n* `REPLICATION_APPLIANCE` - This entity type is a replication appliance.\n* `REPLICATION_JOB` - This entity type is a replication job.\n* `IP_POOL` - This entity type is an IP Pool.\n* `REPLICATION_INFO` - This entity type is a replication information.\n* `DP_VM_SNAPSHOT` - This entity type is a DP VM Snapshot.\n* `DP_VMGROUP_SNAPSHOT` - This entity type is a DP VM Group Snapshot.\n* `DP_VM_CONFIG` - This entity type is a DP VM Configuration.\n* `DP_VM` - This entity type is a DP VM.\n* `DP_VMGROUP` - This entity type is a DP VM Group.\n* `DP_VM_SNAPSHOT_POINT` - This entity type is a DP VM Snapshot Point.\n* `CLUSTER_PAIR` - This entity is a cluster pair.\n* `HX_TASK` - This entity type is a HyperFlex task.\n* `ZONE` - This entity type is a zone.",
						Type:        schema.TypeString,
						Optional:    true,
					},
				},
			},
		},
	}
	model["results"] = &schema.Schema{
		Type:     schema.TypeList,
		Elem:     &schema.Resource{Schema: subSchema},
		Computed: true,
	}
	return &schema.Resource{
		ReadContext: dataSourceHyperflexVmSnapshotInfoRead,
		Schema:      model}
}

func dataSourceHyperflexVmSnapshotInfoRead(c context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Printf("%v", meta)
	conn := meta.(*Config)
	var de diag.Diagnostics
	var o = &models.HyperflexVmSnapshotInfo{}
	if v, ok := d.GetOk("account_moid"); ok {
		x := (v.(string))
		o.SetAccountMoid(x)
	}

	if v, ok := d.GetOk("additional_properties"); ok {
		x := []byte(v.(string))
		var x1 interface{}
		err := json.Unmarshal(x, &x1)
		if err == nil && x1 != nil {
			o.AdditionalProperties = x1.(map[string]interface{})
		}
	}

	if v, ok := d.GetOk("ancestors"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetAncestors(x)
	}

	if v, ok := d.GetOk("class_id"); ok {
		x := (v.(string))
		o.SetClassId(x)
	}

	if v, ok := d.GetOk("cluster_id_snap_map"); ok {
		x := make([]models.HyperflexMapClusterIdToStSnapshotPoint, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.HyperflexMapClusterIdToStSnapshotPoint{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hyperflex.MapClusterIdToStSnapshotPoint")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			x = append(x, *o)
		}
		o.SetClusterIdSnapMap(x)
	}

	if v, ok := d.GetOk("create_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetCreateTime(x)
	}

	if v, ok := d.GetOk("display_status"); ok {
		x := (v.(string))
		o.SetDisplayStatus(x)
	}

	if v, ok := d.GetOk("domain_group_moid"); ok {
		x := (v.(string))
		o.SetDomainGroupMoid(x)
	}

	if v, ok := d.GetOk("error"); ok {
		p := make([]models.HyperflexErrorStack, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexErrorStack{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hyperflex.ErrorStack")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetError(x)
		}
	}

	if v, ok := d.GetOk("label"); ok {
		x := (v.(string))
		o.SetLabel(x)
	}

	if v, ok := d.GetOk("mod_time"); ok {
		x, _ := time.Parse(time.RFC1123, v.(string))
		o.SetModTime(x)
	}

	if v, ok := d.GetOk("mode"); ok {
		x := (v.(string))
		o.SetMode(x)
	}

	if v, ok := d.GetOk("moid"); ok {
		x := (v.(string))
		o.SetMoid(x)
	}

	if v, ok := d.GetOk("object_type"); ok {
		x := (v.(string))
		o.SetObjectType(x)
	}

	if v, ok := d.GetOk("owners"); ok {
		x := make([]string, 0)
		y := reflect.ValueOf(v)
		for i := 0; i < y.Len(); i++ {
			if y.Index(i).Interface() != nil {
				x = append(x, y.Index(i).Interface().(string))
			}
		}
		o.SetOwners(x)
	}

	if v, ok := d.GetOk("parent"); ok {
		p := make([]models.MoBaseMoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParent(x)
		}
	}

	if v, ok := d.GetOk("parent_snapshot"); ok {
		p := make([]models.HyperflexEntityReference, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexEntityReference{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hyperflex.EntityReference")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetParentSnapshot(x)
		}
	}

	if v, ok := d.GetOk("permission_resources"); ok {
		x := make([]models.MoBaseMoRelationship, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoMoRef{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			x = append(x, models.MoMoRefAsMoBaseMoRelationship(o))
		}
		o.SetPermissionResources(x)
	}

	if v, ok := d.GetOk("replication_status"); ok {
		x := (v.(string))
		o.SetReplicationStatus(x)
	}

	if v, ok := d.GetOk("shared_scope"); ok {
		x := (v.(string))
		o.SetSharedScope(x)
	}

	if v, ok := d.GetOk("snapshot_error_msg"); ok {
		x := (v.(string))
		o.SetSnapshotErrorMsg(x)
	}

	if v, ok := d.GetOk("snapshot_status"); ok {
		x := (v.(string))
		o.SetSnapshotStatus(x)
	}

	if v, ok := d.GetOkExists("source_timestamp"); ok {
		x := int64(v.(int))
		o.SetSourceTimestamp(x)
	}

	if v, ok := d.GetOk("src_cluster"); ok {
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetSrcCluster(x)
		}
	}

	if v, ok := d.GetOk("src_cluster_name"); ok {
		x := (v.(string))
		o.SetSrcClusterName(x)
	}

	if v, ok := d.GetOk("tags"); ok {
		x := make([]models.MoTag, 0)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			o := &models.MoTag{}
			l := s[i].(map[string]interface{})
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			if v, ok := l["key"]; ok {
				{
					x := (v.(string))
					o.SetKey(x)
				}
			}
			if v, ok := l["value"]; ok {
				{
					x := (v.(string))
					o.SetValue(x)
				}
			}
			x = append(x, *o)
		}
		o.SetTags(x)
	}

	if v, ok := d.GetOkExists("target_completion_timestamp"); ok {
		x := int64(v.(int))
		o.SetTargetCompletionTimestamp(x)
	}

	if v, ok := d.GetOk("tgt_cluster"); ok {
		p := make([]models.HyperflexClusterRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexClusterRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetTgtCluster(x)
		}
	}

	if v, ok := d.GetOk("tgt_cluster_name"); ok {
		x := (v.(string))
		o.SetTgtClusterName(x)
	}

	if v, ok := d.GetOk("version_context"); ok {
		p := make([]models.MoVersionContext, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoVersionContext{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.VersionContext")
			if v, ok := l["interested_mos"]; ok {
				{
					x := make([]models.MoMoRef, 0)
					s := v.([]interface{})
					for i := 0; i < len(s); i++ {
						o := models.NewMoMoRefWithDefaults()
						l := s[i].(map[string]interface{})
						if v, ok := l["additional_properties"]; ok {
							{
								x := []byte(v.(string))
								var x1 interface{}
								err := json.Unmarshal(x, &x1)
								if err == nil && x1 != nil {
									o.AdditionalProperties = x1.(map[string]interface{})
								}
							}
						}
						o.SetClassId("mo.MoRef")
						if v, ok := l["moid"]; ok {
							{
								x := (v.(string))
								o.SetMoid(x)
							}
						}
						if v, ok := l["object_type"]; ok {
							{
								x := (v.(string))
								o.SetObjectType(x)
							}
						}
						if v, ok := l["selector"]; ok {
							{
								x := (v.(string))
								o.SetSelector(x)
							}
						}
						x = append(x, *o)
					}
					if len(x) > 0 {
						o.SetInterestedMos(x)
					}
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVersionContext(x)
		}
	}

	if v, ok := d.GetOk("vm_backup_info"); ok {
		p := make([]models.HyperflexVmBackupInfoRelationship, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.MoMoRef{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("mo.MoRef")
			if v, ok := l["moid"]; ok {
				{
					x := (v.(string))
					o.SetMoid(x)
				}
			}
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			if v, ok := l["selector"]; ok {
				{
					x := (v.(string))
					o.SetSelector(x)
				}
			}
			p = append(p, models.MoMoRefAsHyperflexVmBackupInfoRelationship(o))
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmBackupInfo(x)
		}
	}

	if v, ok := d.GetOk("vm_entity_reference"); ok {
		p := make([]models.HyperflexEntityReference, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexEntityReference{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hyperflex.EntityReference")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmEntityReference(x)
		}
	}

	if v, ok := d.GetOk("vm_snapshot_entity_reference"); ok {
		p := make([]models.HyperflexEntityReference, 0, 1)
		s := v.([]interface{})
		for i := 0; i < len(s); i++ {
			l := s[i].(map[string]interface{})
			o := &models.HyperflexEntityReference{}
			if v, ok := l["additional_properties"]; ok {
				{
					x := []byte(v.(string))
					var x1 interface{}
					err := json.Unmarshal(x, &x1)
					if err == nil && x1 != nil {
						o.AdditionalProperties = x1.(map[string]interface{})
					}
				}
			}
			o.SetClassId("hyperflex.EntityReference")
			if v, ok := l["object_type"]; ok {
				{
					x := (v.(string))
					o.SetObjectType(x)
				}
			}
			p = append(p, *o)
		}
		if len(p) > 0 {
			x := p[0]
			o.SetVmSnapshotEntityReference(x)
		}
	}

	data, err := o.MarshalJSON()
	if err != nil {
		return diag.Errorf("json marshal of HyperflexVmSnapshotInfo object failed with error : %s", err.Error())
	}
	countResponse, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmSnapshotInfoList(conn.ctx).Filter(getRequestParams(data)).Inlinecount("allpages").Execute()
	if responseErr != nil {
		errorType := fmt.Sprintf("%T", responseErr)
		if strings.Contains(errorType, "GenericOpenAPIError") {
			responseErr := responseErr.(models.GenericOpenAPIError)
			return diag.Errorf("error occurred while fetching count of HyperflexVmSnapshotInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
		}
		return diag.Errorf("error occurred while fetching count of HyperflexVmSnapshotInfo: %s", responseErr.Error())
	}
	count := countResponse.HyperflexVmSnapshotInfoList.GetCount()
	if count == 0 {
		return diag.Errorf("your query for HyperflexVmSnapshotInfo data source did not return any results. Please change your search criteria and try again")
	}
	var i int32
	var hyperflexVmSnapshotInfoResults = make([]map[string]interface{}, count, count)
	var j = 0
	for i = 0; i < count; i += 100 {
		resMo, _, responseErr := conn.ApiClient.HyperflexApi.GetHyperflexVmSnapshotInfoList(conn.ctx).Filter(getRequestParams(data)).Top(100).Skip(i).Execute()
		if responseErr != nil {
			errorType := fmt.Sprintf("%T", responseErr)
			if strings.Contains(errorType, "GenericOpenAPIError") {
				responseErr := responseErr.(models.GenericOpenAPIError)
				return diag.Errorf("error occurred while fetching HyperflexVmSnapshotInfo: %s Response from endpoint: %s", responseErr.Error(), string(responseErr.Body()))
			}
			return diag.Errorf("error occurred while fetching HyperflexVmSnapshotInfo: %s", responseErr.Error())
		}
		results := resMo.HyperflexVmSnapshotInfoList.GetResults()
		switch reflect.TypeOf(results).Kind() {
		case reflect.Slice:
			for i := 0; i < len(results); i++ {
				var s = results[i]
				var temp = make(map[string]interface{})
				temp["account_moid"] = (s.GetAccountMoid())
				temp["additional_properties"] = flattenAdditionalProperties(s.AdditionalProperties)

				temp["ancestors"] = flattenListMoBaseMoRelationship(s.GetAncestors(), d)
				temp["class_id"] = (s.GetClassId())

				temp["cluster_id_snap_map"] = flattenListHyperflexMapClusterIdToStSnapshotPoint(s.GetClusterIdSnapMap(), d)

				temp["create_time"] = (s.GetCreateTime()).String()
				temp["display_status"] = (s.GetDisplayStatus())
				temp["domain_group_moid"] = (s.GetDomainGroupMoid())

				temp["error"] = flattenMapHyperflexErrorStack(s.GetError(), d)
				temp["label"] = (s.GetLabel())

				temp["mod_time"] = (s.GetModTime()).String()
				temp["mode"] = (s.GetMode())
				temp["moid"] = (s.GetMoid())
				temp["object_type"] = (s.GetObjectType())
				temp["owners"] = (s.GetOwners())

				temp["parent"] = flattenMapMoBaseMoRelationship(s.GetParent(), d)

				temp["parent_snapshot"] = flattenMapHyperflexEntityReference(s.GetParentSnapshot(), d)

				temp["permission_resources"] = flattenListMoBaseMoRelationship(s.GetPermissionResources(), d)
				temp["replication_status"] = (s.GetReplicationStatus())
				temp["shared_scope"] = (s.GetSharedScope())
				temp["snapshot_error_msg"] = (s.GetSnapshotErrorMsg())
				temp["snapshot_status"] = (s.GetSnapshotStatus())
				temp["source_timestamp"] = (s.GetSourceTimestamp())

				temp["src_cluster"] = flattenMapHyperflexClusterRelationship(s.GetSrcCluster(), d)
				temp["src_cluster_name"] = (s.GetSrcClusterName())

				temp["tags"] = flattenListMoTag(s.GetTags(), d)
				temp["target_completion_timestamp"] = (s.GetTargetCompletionTimestamp())

				temp["tgt_cluster"] = flattenMapHyperflexClusterRelationship(s.GetTgtCluster(), d)
				temp["tgt_cluster_name"] = (s.GetTgtClusterName())

				temp["version_context"] = flattenMapMoVersionContext(s.GetVersionContext(), d)

				temp["vm_backup_info"] = flattenMapHyperflexVmBackupInfoRelationship(s.GetVmBackupInfo(), d)

				temp["vm_entity_reference"] = flattenMapHyperflexEntityReference(s.GetVmEntityReference(), d)

				temp["vm_snapshot_entity_reference"] = flattenMapHyperflexEntityReference(s.GetVmSnapshotEntityReference(), d)
				hyperflexVmSnapshotInfoResults[j] = temp
				j += 1
			}
		}
	}
	log.Println("length of results: ", len(hyperflexVmSnapshotInfoResults))
	if err := d.Set("results", hyperflexVmSnapshotInfoResults); err != nil {
		return diag.Errorf("error occurred while setting results: %s", err.Error())
	}
	d.SetId(hyperflexVmSnapshotInfoResults[0]["moid"].(string))
	return de
}
