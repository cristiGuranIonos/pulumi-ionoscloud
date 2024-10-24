// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An API gateway consists of the generic rules and configurations.
//
// ## Usage example
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ionoscloud.NewApigateway(ctx, "example", &ionoscloud.ApigatewayArgs{
//				Metrics: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// In order to import an API Gateway, you can define an empty API Gateway resource in the plan:
//
// resource "ionoscloud_apigateway" "example" {
//
// }
//
// The resource can be imported using the `gateway_id`, for example:
//
// ```sh
// $ pulumi import ionoscloud:index/apigateway:Apigateway example {gateway_id}
// ```
type Apigateway struct {
	pulumi.CustomResourceState

	// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
	CustomDomains ApigatewayCustomDomainArrayOutput `pulumi:"customDomains"`
	// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
	Logs pulumi.BoolPtrOutput `pulumi:"logs"`
	// [bool] Enable or disable metrics. Defaults to `false`.
	Metrics pulumi.BoolPtrOutput `pulumi:"metrics"`
	// [string] The domain name. Externally reachable.
	Name pulumi.StringOutput `pulumi:"name"`
	// [string] The public endpoint of the API Gateway.
	PublicEndpoint pulumi.StringOutput `pulumi:"publicEndpoint"`
}

// NewApigateway registers a new resource with the given unique name, arguments, and options.
func NewApigateway(ctx *pulumi.Context,
	name string, args *ApigatewayArgs, opts ...pulumi.ResourceOption) (*Apigateway, error) {
	if args == nil {
		args = &ApigatewayArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Apigateway
	err := ctx.RegisterResource("ionoscloud:index/apigateway:Apigateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigateway gets an existing Apigateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigatewayState, opts ...pulumi.ResourceOption) (*Apigateway, error) {
	var resource Apigateway
	err := ctx.ReadResource("ionoscloud:index/apigateway:Apigateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Apigateway resources.
type apigatewayState struct {
	// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
	CustomDomains []ApigatewayCustomDomain `pulumi:"customDomains"`
	// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
	Logs *bool `pulumi:"logs"`
	// [bool] Enable or disable metrics. Defaults to `false`.
	Metrics *bool `pulumi:"metrics"`
	// [string] The domain name. Externally reachable.
	Name *string `pulumi:"name"`
	// [string] The public endpoint of the API Gateway.
	PublicEndpoint *string `pulumi:"publicEndpoint"`
}

type ApigatewayState struct {
	// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
	CustomDomains ApigatewayCustomDomainArrayInput
	// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
	Logs pulumi.BoolPtrInput
	// [bool] Enable or disable metrics. Defaults to `false`.
	Metrics pulumi.BoolPtrInput
	// [string] The domain name. Externally reachable.
	Name pulumi.StringPtrInput
	// [string] The public endpoint of the API Gateway.
	PublicEndpoint pulumi.StringPtrInput
}

func (ApigatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigatewayState)(nil)).Elem()
}

type apigatewayArgs struct {
	// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
	CustomDomains []ApigatewayCustomDomain `pulumi:"customDomains"`
	// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
	Logs *bool `pulumi:"logs"`
	// [bool] Enable or disable metrics. Defaults to `false`.
	Metrics *bool `pulumi:"metrics"`
	// [string] The domain name. Externally reachable.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Apigateway resource.
type ApigatewayArgs struct {
	// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
	CustomDomains ApigatewayCustomDomainArrayInput
	// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
	Logs pulumi.BoolPtrInput
	// [bool] Enable or disable metrics. Defaults to `false`.
	Metrics pulumi.BoolPtrInput
	// [string] The domain name. Externally reachable.
	Name pulumi.StringPtrInput
}

func (ApigatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigatewayArgs)(nil)).Elem()
}

type ApigatewayInput interface {
	pulumi.Input

	ToApigatewayOutput() ApigatewayOutput
	ToApigatewayOutputWithContext(ctx context.Context) ApigatewayOutput
}

func (*Apigateway) ElementType() reflect.Type {
	return reflect.TypeOf((**Apigateway)(nil)).Elem()
}

func (i *Apigateway) ToApigatewayOutput() ApigatewayOutput {
	return i.ToApigatewayOutputWithContext(context.Background())
}

func (i *Apigateway) ToApigatewayOutputWithContext(ctx context.Context) ApigatewayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayOutput)
}

// ApigatewayArrayInput is an input type that accepts ApigatewayArray and ApigatewayArrayOutput values.
// You can construct a concrete instance of `ApigatewayArrayInput` via:
//
//	ApigatewayArray{ ApigatewayArgs{...} }
type ApigatewayArrayInput interface {
	pulumi.Input

	ToApigatewayArrayOutput() ApigatewayArrayOutput
	ToApigatewayArrayOutputWithContext(context.Context) ApigatewayArrayOutput
}

type ApigatewayArray []ApigatewayInput

func (ApigatewayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apigateway)(nil)).Elem()
}

func (i ApigatewayArray) ToApigatewayArrayOutput() ApigatewayArrayOutput {
	return i.ToApigatewayArrayOutputWithContext(context.Background())
}

func (i ApigatewayArray) ToApigatewayArrayOutputWithContext(ctx context.Context) ApigatewayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayArrayOutput)
}

// ApigatewayMapInput is an input type that accepts ApigatewayMap and ApigatewayMapOutput values.
// You can construct a concrete instance of `ApigatewayMapInput` via:
//
//	ApigatewayMap{ "key": ApigatewayArgs{...} }
type ApigatewayMapInput interface {
	pulumi.Input

	ToApigatewayMapOutput() ApigatewayMapOutput
	ToApigatewayMapOutputWithContext(context.Context) ApigatewayMapOutput
}

type ApigatewayMap map[string]ApigatewayInput

func (ApigatewayMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apigateway)(nil)).Elem()
}

func (i ApigatewayMap) ToApigatewayMapOutput() ApigatewayMapOutput {
	return i.ToApigatewayMapOutputWithContext(context.Background())
}

func (i ApigatewayMap) ToApigatewayMapOutputWithContext(ctx context.Context) ApigatewayMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayMapOutput)
}

type ApigatewayOutput struct{ *pulumi.OutputState }

func (ApigatewayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Apigateway)(nil)).Elem()
}

func (o ApigatewayOutput) ToApigatewayOutput() ApigatewayOutput {
	return o
}

func (o ApigatewayOutput) ToApigatewayOutputWithContext(ctx context.Context) ApigatewayOutput {
	return o
}

// [list] Custom domains for the API Gateway, a list that contains elements with the following structure:
func (o ApigatewayOutput) CustomDomains() ApigatewayCustomDomainArrayOutput {
	return o.ApplyT(func(v *Apigateway) ApigatewayCustomDomainArrayOutput { return v.CustomDomains }).(ApigatewayCustomDomainArrayOutput)
}

// [bool] Enable or disable logging. Defaults to `false`. **NOTE**: Central Logging must be enabled through the Logging API to enable this feature.
func (o ApigatewayOutput) Logs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Apigateway) pulumi.BoolPtrOutput { return v.Logs }).(pulumi.BoolPtrOutput)
}

// [bool] Enable or disable metrics. Defaults to `false`.
func (o ApigatewayOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Apigateway) pulumi.BoolPtrOutput { return v.Metrics }).(pulumi.BoolPtrOutput)
}

// [string] The domain name. Externally reachable.
func (o ApigatewayOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Apigateway) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [string] The public endpoint of the API Gateway.
func (o ApigatewayOutput) PublicEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Apigateway) pulumi.StringOutput { return v.PublicEndpoint }).(pulumi.StringOutput)
}

type ApigatewayArrayOutput struct{ *pulumi.OutputState }

func (ApigatewayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Apigateway)(nil)).Elem()
}

func (o ApigatewayArrayOutput) ToApigatewayArrayOutput() ApigatewayArrayOutput {
	return o
}

func (o ApigatewayArrayOutput) ToApigatewayArrayOutputWithContext(ctx context.Context) ApigatewayArrayOutput {
	return o
}

func (o ApigatewayArrayOutput) Index(i pulumi.IntInput) ApigatewayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Apigateway {
		return vs[0].([]*Apigateway)[vs[1].(int)]
	}).(ApigatewayOutput)
}

type ApigatewayMapOutput struct{ *pulumi.OutputState }

func (ApigatewayMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Apigateway)(nil)).Elem()
}

func (o ApigatewayMapOutput) ToApigatewayMapOutput() ApigatewayMapOutput {
	return o
}

func (o ApigatewayMapOutput) ToApigatewayMapOutputWithContext(ctx context.Context) ApigatewayMapOutput {
	return o
}

func (o ApigatewayMapOutput) MapIndex(k pulumi.StringInput) ApigatewayOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Apigateway {
		return vs[0].(map[string]*Apigateway)[vs[1].(string)]
	}).(ApigatewayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayInput)(nil)).Elem(), &Apigateway{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayArrayInput)(nil)).Elem(), ApigatewayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayMapInput)(nil)).Elem(), ApigatewayMap{})
	pulumi.RegisterOutputType(ApigatewayOutput{})
	pulumi.RegisterOutputType(ApigatewayArrayOutput{})
	pulumi.RegisterOutputType(ApigatewayMapOutput{})
}