package vmworkflow

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func VirtualMachineOvfDeploySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"local_ovf_path": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The absolute path to the ovf/ova file in the local system.",
			ForceNew:    true,
		},
		"remote_ovf_url": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "URL to the remote ovf/ova file to be deployed.",
			ForceNew:    true,
		},
		"ip_allocation_policy": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The IP allocation policy.",
			ForceNew:    true,
		},
		"ip_protocol": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "The IP protocol.",
			ForceNew:    true,
		},
		"disk_provisioning": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "An optional disk provisioning. If set, all the disks in the deployed ovf will have the same specified disk type (e.g., thin provisioned).",
			ForceNew:    true,
		},
		"deployment": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "A specifc OVA deployment to use if multiple are available.",
			ForceNew:    true,
		},
		"ovf_network_map": {
			Type:        schema.TypeMap,
			Optional:    true,
			Description: "The mapping of name of network identifiers from the ovf descriptor to network UUID in the VI infrastructure.",
			Elem:        &schema.Schema{Type: schema.TypeString},
			ForceNew:    true,
		},
		"allow_unverified_ssl_cert": {
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     true,
			Description: "Allow unverified ssl certificates while deploying ovf/ova from url.",
		},
	}
}
