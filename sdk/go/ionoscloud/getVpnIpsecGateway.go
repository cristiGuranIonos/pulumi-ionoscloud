// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **VPN IPSec Gateway data source** can be used to search for and return an existing IPSec Gateway.
// You can provide a string for the name parameter which will be compared with provisioned IPSec Gateways.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
//
// ## Example Usage
func LookupVpnIpsecGateway(ctx *pulumi.Context, args *LookupVpnIpsecGatewayArgs, opts ...pulumi.InvokeOption) (*LookupVpnIpsecGatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnIpsecGatewayResult
	err := ctx.Invoke("ionoscloud:index/getVpnIpsecGateway:getVpnIpsecGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpnIpsecGateway.
type LookupVpnIpsecGatewayArgs struct {
	// ID of an existing IPSec Gateway that you want to search for.
	Id *string `pulumi:"id"`
	// The location of the IPSec Gateway.
	Location string `pulumi:"location"`
	// Name of an existing IPSec Gateway that you want to search for.
	Name *string `pulumi:"name"`
	// The IKE version that is permitted for the VPN tunnels.
	Version *string `pulumi:"version"`
}

// A collection of values returned by getVpnIpsecGateway.
type LookupVpnIpsecGatewayResult struct {
	// The network connection for your gateway.
	Connections []GetVpnIpsecGatewayConnection `pulumi:"connections"`
	// (Optional)[string] The human-readable description of the IPSec Gateway.
	Description string `pulumi:"description"`
	// Public IP address to be assigned to the gateway.
	GatewayIp string `pulumi:"gatewayIp"`
	// The unique ID of the IPSec Gateway.
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
	// The name of the IPSec Gateway.
	Name string `pulumi:"name"`
	// The IKE version that is permitted for the VPN tunnels.
	Version string `pulumi:"version"`
}

func LookupVpnIpsecGatewayOutput(ctx *pulumi.Context, args LookupVpnIpsecGatewayOutputArgs, opts ...pulumi.InvokeOption) LookupVpnIpsecGatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnIpsecGatewayResult, error) {
			args := v.(LookupVpnIpsecGatewayArgs)
			r, err := LookupVpnIpsecGateway(ctx, &args, opts...)
			var s LookupVpnIpsecGatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnIpsecGatewayResultOutput)
}

// A collection of arguments for invoking getVpnIpsecGateway.
type LookupVpnIpsecGatewayOutputArgs struct {
	// ID of an existing IPSec Gateway that you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The location of the IPSec Gateway.
	Location pulumi.StringInput `pulumi:"location"`
	// Name of an existing IPSec Gateway that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The IKE version that is permitted for the VPN tunnels.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (LookupVpnIpsecGatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnIpsecGatewayArgs)(nil)).Elem()
}

// A collection of values returned by getVpnIpsecGateway.
type LookupVpnIpsecGatewayResultOutput struct{ *pulumi.OutputState }

func (LookupVpnIpsecGatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnIpsecGatewayResult)(nil)).Elem()
}

func (o LookupVpnIpsecGatewayResultOutput) ToLookupVpnIpsecGatewayResultOutput() LookupVpnIpsecGatewayResultOutput {
	return o
}

func (o LookupVpnIpsecGatewayResultOutput) ToLookupVpnIpsecGatewayResultOutputWithContext(ctx context.Context) LookupVpnIpsecGatewayResultOutput {
	return o
}

// The network connection for your gateway.
func (o LookupVpnIpsecGatewayResultOutput) Connections() GetVpnIpsecGatewayConnectionArrayOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) []GetVpnIpsecGatewayConnection { return v.Connections }).(GetVpnIpsecGatewayConnectionArrayOutput)
}

// (Optional)[string] The human-readable description of the IPSec Gateway.
func (o LookupVpnIpsecGatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

// Public IP address to be assigned to the gateway.
func (o LookupVpnIpsecGatewayResultOutput) GatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.GatewayIp }).(pulumi.StringOutput)
}

// The unique ID of the IPSec Gateway.
func (o LookupVpnIpsecGatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVpnIpsecGatewayResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the IPSec Gateway.
func (o LookupVpnIpsecGatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

// The IKE version that is permitted for the VPN tunnels.
func (o LookupVpnIpsecGatewayResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnIpsecGatewayResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnIpsecGatewayResultOutput{})
}
