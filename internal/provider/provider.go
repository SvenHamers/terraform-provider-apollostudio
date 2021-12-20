package provider

import (
	"context"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/client"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/services/apikey"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/services/graph"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/services/invitation"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/services/organization"
	"github.com/SvenHamers/terraform-provider-apollostudio/internal/services/user"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func AppolloStudioProvider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"personal_apikey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"enterprise_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"apollostudio_graph":        graph.Resource(),
			"apollostudio_apikey":       apikey.Resource(),
			"apollostudio_organization": organization.Resource(),
			"apollostudio_user":         user.Resource(),
			"apollostudio_invitation":   invitation.Resource(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: ProviderConfigure,
	}
}

func ProviderConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	var diags diag.Diagnostics

	Appollo := &client.Client{
		ApiKey:            d.Get("personal_apikey").(string),
		EnterPriseEnabled: d.Get("enterprise_enabled").(bool),
	}

	return Appollo, diags

}
