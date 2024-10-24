// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an **API Gateway Route** on IonosCloud.
//
// ## Example Usage
//
// This resource will create an operational API Gateway Route. After this section completes, the provisioner can be called.
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
//			example, err := ionoscloud.NewApigateway(ctx, "example", &ionoscloud.ApigatewayArgs{
//				Metrics: pulumi.Bool(true),
//				CustomDomains: ionoscloud.ApigatewayCustomDomainArray{
//					&ionoscloud.ApigatewayCustomDomainArgs{
//						Name:          pulumi.String("example.com"),
//						CertificateId: pulumi.String("00000000-0000-0000-0000-000000000000"),
//					},
//					&ionoscloud.ApigatewayCustomDomainArgs{
//						Name:          pulumi.String("example.org"),
//						CertificateId: pulumi.String("00000000-0000-0000-0000-000000000000"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewApigatewayRoute(ctx, "apigatewayRoute", &ionoscloud.ApigatewayRouteArgs{
//				Type: pulumi.String("http"),
//				Paths: pulumi.StringArray{
//					pulumi.String("/foo/*"),
//					pulumi.String("/bar"),
//				},
//				Methods: pulumi.StringArray{
//					pulumi.String("GET"),
//					pulumi.String("POST"),
//				},
//				Websocket: pulumi.Bool(false),
//				Upstreams: ionoscloud.ApigatewayRouteUpstreamArray{
//					&ionoscloud.ApigatewayRouteUpstreamArgs{
//						Scheme:       pulumi.String("http"),
//						Loadbalancer: pulumi.String("roundrobin"),
//						Host:         pulumi.String("example.com"),
//						Port:         pulumi.Int(80),
//						Weight:       pulumi.Int(100),
//					},
//				},
//				GatewayId: example.ID(),
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
// API Gateway route can be imported using the `apigateway route id`:
//
// ```sh
// $ pulumi import ionoscloud:index/apigatewayRoute:ApigatewayRoute myroute {apigateway uuid}:{apigateway route uuid}
// ```
type ApigatewayRoute struct {
	pulumi.CustomResourceState

	// [string] The ID of the API Gateway that the route belongs to.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
	// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
	Methods pulumi.StringArrayOutput `pulumi:"methods"`
	// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
	Name pulumi.StringOutput `pulumi:"name"`
	// [list] The paths that the route should match. Minimum items: 1.
	Paths pulumi.StringArrayOutput `pulumi:"paths"`
	// [string] This field specifies the protocol used by the ingress to route traffic to the backend
	// service. Default value: `http`.
	Type pulumi.StringPtrOutput `pulumi:"type"`
	// Upstreams information of the API Gateway Route. Minimum items: 1.
	Upstreams ApigatewayRouteUpstreamArrayOutput `pulumi:"upstreams"`
	// [bool] To enable websocket support. Default value: `false`.
	Websocket pulumi.BoolPtrOutput `pulumi:"websocket"`
}

// NewApigatewayRoute registers a new resource with the given unique name, arguments, and options.
func NewApigatewayRoute(ctx *pulumi.Context,
	name string, args *ApigatewayRouteArgs, opts ...pulumi.ResourceOption) (*ApigatewayRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.Methods == nil {
		return nil, errors.New("invalid value for required argument 'Methods'")
	}
	if args.Paths == nil {
		return nil, errors.New("invalid value for required argument 'Paths'")
	}
	if args.Upstreams == nil {
		return nil, errors.New("invalid value for required argument 'Upstreams'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ApigatewayRoute
	err := ctx.RegisterResource("ionoscloud:index/apigatewayRoute:ApigatewayRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApigatewayRoute gets an existing ApigatewayRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApigatewayRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ApigatewayRouteState, opts ...pulumi.ResourceOption) (*ApigatewayRoute, error) {
	var resource ApigatewayRoute
	err := ctx.ReadResource("ionoscloud:index/apigatewayRoute:ApigatewayRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ApigatewayRoute resources.
type apigatewayRouteState struct {
	// [string] The ID of the API Gateway that the route belongs to.
	GatewayId *string `pulumi:"gatewayId"`
	// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
	// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
	Methods []string `pulumi:"methods"`
	// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
	Name *string `pulumi:"name"`
	// [list] The paths that the route should match. Minimum items: 1.
	Paths []string `pulumi:"paths"`
	// [string] This field specifies the protocol used by the ingress to route traffic to the backend
	// service. Default value: `http`.
	Type *string `pulumi:"type"`
	// Upstreams information of the API Gateway Route. Minimum items: 1.
	Upstreams []ApigatewayRouteUpstream `pulumi:"upstreams"`
	// [bool] To enable websocket support. Default value: `false`.
	Websocket *bool `pulumi:"websocket"`
}

type ApigatewayRouteState struct {
	// [string] The ID of the API Gateway that the route belongs to.
	GatewayId pulumi.StringPtrInput
	// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
	// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
	Methods pulumi.StringArrayInput
	// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
	Name pulumi.StringPtrInput
	// [list] The paths that the route should match. Minimum items: 1.
	Paths pulumi.StringArrayInput
	// [string] This field specifies the protocol used by the ingress to route traffic to the backend
	// service. Default value: `http`.
	Type pulumi.StringPtrInput
	// Upstreams information of the API Gateway Route. Minimum items: 1.
	Upstreams ApigatewayRouteUpstreamArrayInput
	// [bool] To enable websocket support. Default value: `false`.
	Websocket pulumi.BoolPtrInput
}

func (ApigatewayRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*apigatewayRouteState)(nil)).Elem()
}

type apigatewayRouteArgs struct {
	// [string] The ID of the API Gateway that the route belongs to.
	GatewayId string `pulumi:"gatewayId"`
	// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
	// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
	Methods []string `pulumi:"methods"`
	// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
	Name *string `pulumi:"name"`
	// [list] The paths that the route should match. Minimum items: 1.
	Paths []string `pulumi:"paths"`
	// [string] This field specifies the protocol used by the ingress to route traffic to the backend
	// service. Default value: `http`.
	Type *string `pulumi:"type"`
	// Upstreams information of the API Gateway Route. Minimum items: 1.
	Upstreams []ApigatewayRouteUpstream `pulumi:"upstreams"`
	// [bool] To enable websocket support. Default value: `false`.
	Websocket *bool `pulumi:"websocket"`
}

// The set of arguments for constructing a ApigatewayRoute resource.
type ApigatewayRouteArgs struct {
	// [string] The ID of the API Gateway that the route belongs to.
	GatewayId pulumi.StringInput
	// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
	// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
	Methods pulumi.StringArrayInput
	// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
	Name pulumi.StringPtrInput
	// [list] The paths that the route should match. Minimum items: 1.
	Paths pulumi.StringArrayInput
	// [string] This field specifies the protocol used by the ingress to route traffic to the backend
	// service. Default value: `http`.
	Type pulumi.StringPtrInput
	// Upstreams information of the API Gateway Route. Minimum items: 1.
	Upstreams ApigatewayRouteUpstreamArrayInput
	// [bool] To enable websocket support. Default value: `false`.
	Websocket pulumi.BoolPtrInput
}

func (ApigatewayRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apigatewayRouteArgs)(nil)).Elem()
}

type ApigatewayRouteInput interface {
	pulumi.Input

	ToApigatewayRouteOutput() ApigatewayRouteOutput
	ToApigatewayRouteOutputWithContext(ctx context.Context) ApigatewayRouteOutput
}

func (*ApigatewayRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigatewayRoute)(nil)).Elem()
}

func (i *ApigatewayRoute) ToApigatewayRouteOutput() ApigatewayRouteOutput {
	return i.ToApigatewayRouteOutputWithContext(context.Background())
}

func (i *ApigatewayRoute) ToApigatewayRouteOutputWithContext(ctx context.Context) ApigatewayRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayRouteOutput)
}

// ApigatewayRouteArrayInput is an input type that accepts ApigatewayRouteArray and ApigatewayRouteArrayOutput values.
// You can construct a concrete instance of `ApigatewayRouteArrayInput` via:
//
//	ApigatewayRouteArray{ ApigatewayRouteArgs{...} }
type ApigatewayRouteArrayInput interface {
	pulumi.Input

	ToApigatewayRouteArrayOutput() ApigatewayRouteArrayOutput
	ToApigatewayRouteArrayOutputWithContext(context.Context) ApigatewayRouteArrayOutput
}

type ApigatewayRouteArray []ApigatewayRouteInput

func (ApigatewayRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigatewayRoute)(nil)).Elem()
}

func (i ApigatewayRouteArray) ToApigatewayRouteArrayOutput() ApigatewayRouteArrayOutput {
	return i.ToApigatewayRouteArrayOutputWithContext(context.Background())
}

func (i ApigatewayRouteArray) ToApigatewayRouteArrayOutputWithContext(ctx context.Context) ApigatewayRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayRouteArrayOutput)
}

// ApigatewayRouteMapInput is an input type that accepts ApigatewayRouteMap and ApigatewayRouteMapOutput values.
// You can construct a concrete instance of `ApigatewayRouteMapInput` via:
//
//	ApigatewayRouteMap{ "key": ApigatewayRouteArgs{...} }
type ApigatewayRouteMapInput interface {
	pulumi.Input

	ToApigatewayRouteMapOutput() ApigatewayRouteMapOutput
	ToApigatewayRouteMapOutputWithContext(context.Context) ApigatewayRouteMapOutput
}

type ApigatewayRouteMap map[string]ApigatewayRouteInput

func (ApigatewayRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigatewayRoute)(nil)).Elem()
}

func (i ApigatewayRouteMap) ToApigatewayRouteMapOutput() ApigatewayRouteMapOutput {
	return i.ToApigatewayRouteMapOutputWithContext(context.Background())
}

func (i ApigatewayRouteMap) ToApigatewayRouteMapOutputWithContext(ctx context.Context) ApigatewayRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApigatewayRouteMapOutput)
}

type ApigatewayRouteOutput struct{ *pulumi.OutputState }

func (ApigatewayRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApigatewayRoute)(nil)).Elem()
}

func (o ApigatewayRouteOutput) ToApigatewayRouteOutput() ApigatewayRouteOutput {
	return o
}

func (o ApigatewayRouteOutput) ToApigatewayRouteOutputWithContext(ctx context.Context) ApigatewayRouteOutput {
	return o
}

// [string] The ID of the API Gateway that the route belongs to.
func (o ApigatewayRouteOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// [list] The HTTP methods that the route should match. Minimum items: 1. Possible values: `GET`,
// `POST`, `PUT`, `DELETE`, `PATCH`, `OPTIONS`, `HEAD`, `CONNECT`, `TRACE`.
func (o ApigatewayRouteOutput) Methods() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.StringArrayOutput { return v.Methods }).(pulumi.StringArrayOutput)
}

// [string] Name of the API Gateway Route. Only alphanumeric characters are allowed.
func (o ApigatewayRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [list] The paths that the route should match. Minimum items: 1.
func (o ApigatewayRouteOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.StringArrayOutput { return v.Paths }).(pulumi.StringArrayOutput)
}

// [string] This field specifies the protocol used by the ingress to route traffic to the backend
// service. Default value: `http`.
func (o ApigatewayRouteOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.StringPtrOutput { return v.Type }).(pulumi.StringPtrOutput)
}

// Upstreams information of the API Gateway Route. Minimum items: 1.
func (o ApigatewayRouteOutput) Upstreams() ApigatewayRouteUpstreamArrayOutput {
	return o.ApplyT(func(v *ApigatewayRoute) ApigatewayRouteUpstreamArrayOutput { return v.Upstreams }).(ApigatewayRouteUpstreamArrayOutput)
}

// [bool] To enable websocket support. Default value: `false`.
func (o ApigatewayRouteOutput) Websocket() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ApigatewayRoute) pulumi.BoolPtrOutput { return v.Websocket }).(pulumi.BoolPtrOutput)
}

type ApigatewayRouteArrayOutput struct{ *pulumi.OutputState }

func (ApigatewayRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ApigatewayRoute)(nil)).Elem()
}

func (o ApigatewayRouteArrayOutput) ToApigatewayRouteArrayOutput() ApigatewayRouteArrayOutput {
	return o
}

func (o ApigatewayRouteArrayOutput) ToApigatewayRouteArrayOutputWithContext(ctx context.Context) ApigatewayRouteArrayOutput {
	return o
}

func (o ApigatewayRouteArrayOutput) Index(i pulumi.IntInput) ApigatewayRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ApigatewayRoute {
		return vs[0].([]*ApigatewayRoute)[vs[1].(int)]
	}).(ApigatewayRouteOutput)
}

type ApigatewayRouteMapOutput struct{ *pulumi.OutputState }

func (ApigatewayRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ApigatewayRoute)(nil)).Elem()
}

func (o ApigatewayRouteMapOutput) ToApigatewayRouteMapOutput() ApigatewayRouteMapOutput {
	return o
}

func (o ApigatewayRouteMapOutput) ToApigatewayRouteMapOutputWithContext(ctx context.Context) ApigatewayRouteMapOutput {
	return o
}

func (o ApigatewayRouteMapOutput) MapIndex(k pulumi.StringInput) ApigatewayRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ApigatewayRoute {
		return vs[0].(map[string]*ApigatewayRoute)[vs[1].(string)]
	}).(ApigatewayRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayRouteInput)(nil)).Elem(), &ApigatewayRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayRouteArrayInput)(nil)).Elem(), ApigatewayRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ApigatewayRouteMapInput)(nil)).Elem(), ApigatewayRouteMap{})
	pulumi.RegisterOutputType(ApigatewayRouteOutput{})
	pulumi.RegisterOutputType(ApigatewayRouteArrayOutput{})
	pulumi.RegisterOutputType(ApigatewayRouteMapOutput{})
}
