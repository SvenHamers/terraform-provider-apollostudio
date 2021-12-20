package graph

import (
	"context"
	"strconv"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type graph struct {
	NewService struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Title string `json:"title"`
	} `json:"newService"`
}

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceRead,
		CreateContext: resourceCreate,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization_id": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_developer": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
			},
			"admin_only": &schema.Schema{
				Type:        schema.TypeBool,
				Description: "Enterprise users only",
				Default:     false,
				Optional:    true,
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
	orgId := data.Get("organization_id").(string)

	id := name + helpers.RandomNumberString(6)

	adminOnly := strconv.FormatBool(!data.Get("admin_only").(bool))

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	var query string

	if appollo.EnterPriseEnabled {
		query = `
		mutation Service {
			newService(accountId: "` + orgId + `", id:  "` + id + `", name: "` + name + `", title: "` + name + `", hiddenFromUninvitedNonAdminAccountMembers: ` + adminOnly + `) {
				id
				name
				title
			}
	  	}
		`
	} else {
		query = `
		mutation Service {
			newService(accountId: "` + orgId + `", id:  "` + id + `", name: "` + name + `", title: "` + name + `") {
				id
				name
				title
			}
	  	}
		`
	}

	err := appollo.Query(ctx, query,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.NewService.ID)

	return diags
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags

}

func resourceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	name := data.Get("name").(string)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	err := appollo.Query(ctx, `

		mutation Service
		{
			service( id:  "`+name+`") {
		 		delete
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
