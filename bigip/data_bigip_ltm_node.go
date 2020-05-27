package bigip

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func bigipLtmNodeDataSource() *schema.Resource {
	bigipLtmNodeSchema := resourceBigipLtmNode().Schema
	updateFieldToOptional("address", bigipLtmNodeSchema)
	return &schema.Resource{
		Read:   bigipLtmNodeDataSourceRead,
		Schema: bigipLtmNodeSchema,
	}
}

func bigipLtmNodeDataSourceRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	d.SetId(name)
	return resourceBigipLtmNodeRead(d, meta)
}
