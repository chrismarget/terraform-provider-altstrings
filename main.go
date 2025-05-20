package main

import (
	"context"
	"flag"
	"log"

	"github.com/chrismarget/terraform-provider-altstrings/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	err := providerserver.Serve(context.Background(), provider.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/chrismarget/altstrings",
		Debug:   debug,
	})
	if err != nil {
		log.Fatal(err)
	}
}
