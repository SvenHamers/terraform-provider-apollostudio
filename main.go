package main

import (
	"context"
	"flag"
	"log"

	"github.com/SvenHamers/terraform-provider-apollostudio/internal/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	// remove date and time stamp from log output as the plugin SDK already adds its own
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	var debugMode bool

	flag.BoolVar(&debugMode, "debuggable", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	if debugMode {
		err := plugin.Debug(context.Background(), "registry.terraform.io/svenhamers/appollostudio",
			&plugin.ServeOpts{
				ProviderFunc: provider.AppolloStudioProvider,
			})
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		plugin.Serve(&plugin.ServeOpts{
			ProviderFunc: provider.AppolloStudioProvider,
		})
	}
}