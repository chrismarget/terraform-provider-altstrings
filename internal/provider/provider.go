package provider

import (
	"context"

	"github.com/chrismarget/terraform-provider-altstrings/internal/resources"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*Provider)(nil)

// New instantiates the provider in main
func New() provider.Provider {
	return new(Provider)
}

type Provider struct{}

func (p *Provider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "altstrings"
}

func (p *Provider) Schema(_ context.Context, _ provider.SchemaRequest, _ *provider.SchemaResponse) {}

func (p *Provider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

// DataSources defines provider data sources
func (p *Provider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines provider resources
func (p *Provider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewThingResource,
	}
}
