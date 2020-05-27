package bigip

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func bigipAs3DataSource() *schema.Resource {
	bigipAs3Schema := resourceBigipAs3().Schema
	updateFieldToOptional("as3_json", bigipAs3Schema)
	return &schema.Resource{
		Read:   bigipAs3DataSourceRead,
		Schema: bigipAs3Schema,
	}
}

func bigipAs3DataSourceRead(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("tenant_list").(string)
	if name == "" {
		return fmt.Errorf("Error obtaining tenant_list during read")
	}
	d.SetId(name)
	d.Set("tenant_list", name)
	return resourceBigipAs3Read(d, meta)
}
