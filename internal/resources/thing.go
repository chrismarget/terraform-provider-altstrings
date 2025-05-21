package resources

import (
	"context"
	"fmt"
	"github.com/chrismarget/terraform-provider-altstrings/internal/crayola"
	"github.com/chrismarget/terraform-provider-altstrings/internal/customtype"
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
	Id    types.String              `tfsdk:"id"`
	Color customtype.StringWithAlts `tfsdk:"color"`
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
				Computed:      true,
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"color": schema.StringAttribute{
				Required:   true,
				CustomType: customtype.StringWithAltsType{},
				Validators: []validator.String{stringvalidator.OneOf(crayola.ValidColors()...)},
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

	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}

func (t thingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var m thingModel
	resp.Diagnostics.Append(req.State.Get(ctx, &m)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// look up the hue of the previously configured color
	hue := crayola.Hue(m.Color.ValueString()) // e.g. "red"
	if hue == "" {
		resp.Diagnostics.AddError("color lookup failure", fmt.Sprintf("color %q has unknown hue", m.Color.ValueString()))
	}

	// look up all colors associated with that hue. these are "alt" values with semantic equality
	colors := crayola.HueColors(hue) // e.g. []string{"crimson", "scarlet", "raspberry"}

	// pretend to have read from the API a hue, but use the custom type with alternate values.
	// e.g. the API returned "red", but we know that's semantically equal to "crimson", "scarlet", etc...
	m.Color = customtype.NewStringWithAlts(hue, colors...)

	resp.Diagnostics.Append(resp.State.Set(ctx, &m)...)
}

func (t thingResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	// no-op
}

func (t thingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// no-op
}
