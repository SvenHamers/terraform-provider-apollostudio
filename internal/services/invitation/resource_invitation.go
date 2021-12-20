package invitation

import (
	"context"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type graph struct {
	Account struct {
		Invite struct {
			ID string `json:"id"`
		} `json:"invite"`
	} `json:"account"`
}

func Resource() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceRead,
		CreateContext: resourceCreate,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"org_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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

	role := data.Get("role").(string)
	email := data.Get("email").(string)
	orgName := data.Get("org_name").(string)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	var query string

	query = `
	mutation Invite {
		account(id: "` + orgName + `") {
		  invite(email: "` + email + `", role: ` + role + ` ) {
			id
		  }
		}
	  }
		`

	err := appollo.Query(ctx, query,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	data.SetId(result.Account.Invite.ID)
	return diags
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags

}

func resourceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	id := data.Id()
	orgName := data.Get("org_name").(string)

	err := appollo.Query(ctx, `
	mutation Invite {
		account(id: "`+orgName+`") {
	  
		  removeInvitation(id: `+id+`)
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
