package user

import (
	"context"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/helpers"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type graph struct {
	SignUp struct {
		ID string `json:"id"`
	} `json:"signUp"`
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
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password_link_after_creation": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
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
	email := data.Get("email").(string)
	resetPassword := data.Get("password_link_after_creation").(bool)

	appollo := meta.(*client.Client)

	appollo.Init()

	var result graph

	var query string

	query = `
		mutation NewService {
			signUp(email: "` + email + `", fullName: "` + name + `", password: "****` + helpers.RandomNumberString(10) + `") {
			  id
			}
		  }
		`

	err := appollo.Query(ctx, query,
		&result)

	if err != nil {
		return diag.FromErr(err)
	}

	if resetPassword {
		err := appollo.Query(ctx, `

		mutation Service
		{
			resetPassword(email: "`+email+`")
	  }
		`,
			&result)

		if err != nil {
			return diag.FromErr(err)
		}
	}

	data.SetId(result.SignUp.ID)

	return diags
}

func resourceUpdate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	return diags

}

func resourceDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	//needs to be implenmented

	data.SetId("")

	return diags

}
