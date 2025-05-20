package resources

import (
	"context"

	"github.com/chrismarget-j/terraform-provider-altstrings/crayola"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type thingModel struct {
	Id    types.String `tfsdk:"id"`
	Color types.String `tfsdk:"color"`
}

type thingResource struct{}

var _ resource.Resource = (*thingResource)(nil)

func NewThingResource() resource.Resource {
	return new(thingResource)
}

func (t thingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_thing"
}

func (t thingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      false,
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"color": schema.StringAttribute{
				Required:   true,
				Validators: []validator.String{stringvalidator.OneOf(crayola.BaseColors()...)},
			},
		},
	}
}

func (t thingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var m thingModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set an ID like we did some actual work here.
	m.Id = types.StringValue(uuid.New().String())
}

func (t thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var m thingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// "read" a random color synonym to force state churn
	m.Color = types.StringValue(crayola.Synonym(m.Color.ValueString()))

	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}

func (t thingResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	// no-op
}

func (t thingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// no-op
}
