package bigip

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func bigipCmDeviceDataSource() *schema.Resource {
	bigipCmDeviceSchema := resourceBigipCmDevice().Schema
	updateFieldToOptional("configsync_ip", bigipCmDeviceSchema)
	return &schema.Resource{
		Read:   bigipCmDeviceDataSourceRead,
		Schema: bigipCmDeviceSchema,
	}
}

func bigipCmDeviceDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	if name == "" {
		return fmt.Errorf("Error obtaining name during read")
	}
	d.SetId(name)
	return resourceBigipCmDeviceRead(d, meta)
}
