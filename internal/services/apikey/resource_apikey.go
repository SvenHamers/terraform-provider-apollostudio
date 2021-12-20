package apikey

import (
	"context"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type graph struct {
	Service struct {
		NewKey struct {
			Id      string      `json:"id"`
			Token   string      `json:"token"`
			KeyName interface{} `json:"keyName"`
		} `json:"newKey"`
	} `json:"service"`
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
			"graph_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"token": &schema.Schema{
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
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

	graphName := data.Get("graph_name").(string)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	err := appollo.Query(ctx, `
			mutation Service {
				service(id: "`+graphName+`") {
					newKey(keyName: "`+name+`") {
						keyName
						id
						token
					}
				}
			}
		`,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Service.NewKey.Id)
	data.Set("token", result.Service.NewKey.Token)

	return diags
}

func resourceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	name := data.Get("graph_name").(string)
	id := data.Id()

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	err := appollo.Query(ctx, `

	mutation Service {
		service(id: "`+name+`") {
		  removeKey(id: "`+id+`")
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
