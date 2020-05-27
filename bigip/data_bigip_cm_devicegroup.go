package bigip

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func bigipCmDevicegroupDataSource() *schema.Resource {
	bigipCmDevicegroupSchema := resourceBigipCmDevicegroup().Schema
	return &schema.Resource{
		Read:   bigipCmDevicegroupDataSourceRead,
		Schema: bigipCmDevicegroupSchema,
	}
}

func bigipCmDevicegroupDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	if name == "" {
		return fmt.Errorf("Error obtaining name during read")
	}
	d.SetId(name)
	return resourceBigipCmDevicegroupRead(d, meta)
}
