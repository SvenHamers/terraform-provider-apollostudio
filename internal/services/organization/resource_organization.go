package organization

import (
	"context"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type graph struct {
	NewAccount struct {
		ID string `json:"id"`
	} `json:"newAccount"`
}

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceRead,
		CreateContext: resourceCreate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags
}

func resourceCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	name := data.Get("name").(string)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	err := appollo.Query(ctx, `
		mutation NewAccount {
			newAccount(id: "`+name+`") {
		  		id
			}
	  }
		`,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.NewAccount.ID)

	return diags
}

func resourceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	name := data.Get("name").(string)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	err := appollo.Query(ctx, `
	mutation NewAccount {
	account(id: "`+name+`") {
		hardDelete
	  }
	}
		`,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId("")

	return diags

}
