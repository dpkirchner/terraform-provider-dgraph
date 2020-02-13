package dtype

import (
	"context"
	"fmt"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const iDTypeTemplate = "type_%s"

func Create(d *schema.ResourceData, m interface{}) error {

	typeName := d.Get("name").(string)
	typeFields := d.Get("fields").(map[string]interface{})

	fieldList := ""
	for k, v := range typeFields {
		prefix := "\n"
		if fieldList == "" {
			prefix = ""
		}
		fieldList = fmt.Sprintf("%s%s  %s: %s", fieldList, prefix, k, v.(string))
	}
	typeDefinition := fmt.Sprintf("type %s {\n%s\n}", typeName, fieldList)

	err := client.Alter(context.Background(), &api.Operation{
		Schema: typeDefinition,
	})

	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf(iDTypeTemplate, typeName))

	return Read(d, m)
}