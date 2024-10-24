// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The `VpnWireguardGateway` data source provides information about a specific IonosCloud VPN WireGuard Gateway. You can use this data source to retrieve details of a WireGuard Gateway for use in other resources and configurations.
//
// ## Example Usage
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
//			_, err := ionoscloud.LookupVpnWireguardPeer(ctx, &ionoscloud.LookupVpnWireguardPeerArgs{
//				Location:  "de/fra",
//				GatewayId: "example-gateway",
//				Name:      pulumi.StringRef("example-peer"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("vpnWireguardPeerPublicKey", data.Vpn_wireguard_peer.Example.Public_key)
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupVpnWireguardPeer(ctx *pulumi.Context, args *LookupVpnWireguardPeerArgs, opts ...pulumi.InvokeOption) (*LookupVpnWireguardPeerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpnWireguardPeerResult
	err := ctx.Invoke("ionoscloud:index/getVpnWireguardPeer:getVpnWireguardPeer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVpnWireguardPeer.
type LookupVpnWireguardPeerArgs struct {
	// [String] The ID of the WireGuard Gateway.
	GatewayId string `pulumi:"gatewayId"`
	// [String] The ID of the WireGuard Peer.
	Id *string `pulumi:"id"`
	// [String] The location of the WireGuard Gateway.
	Location string `pulumi:"location"`
	// [String] The name of the WireGuard Peer.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getVpnWireguardPeer.
type LookupVpnWireguardPeerResult struct {
	// The subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps []string `pulumi:"allowedIps"`
	// The description of the WireGuard Peer.
	Description string `pulumi:"description"`
	// The endpoint of the WireGuard Peer.
	Endpoints []GetVpnWireguardPeerEndpoint `pulumi:"endpoints"`
	GatewayId string                        `pulumi:"gatewayId"`
	// The unique ID of the WireGuard Peer.
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	// The name of the WireGuard Peer.
	Name string `pulumi:"name"`
	// WireGuard public key of the connecting peer.
	PublicKey string `pulumi:"publicKey"`
	// The current status of the WireGuard Peer.
	Status string `pulumi:"status"`
}

func LookupVpnWireguardPeerOutput(ctx *pulumi.Context, args LookupVpnWireguardPeerOutputArgs, opts ...pulumi.InvokeOption) LookupVpnWireguardPeerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnWireguardPeerResult, error) {
			args := v.(LookupVpnWireguardPeerArgs)
			r, err := LookupVpnWireguardPeer(ctx, &args, opts...)
			var s LookupVpnWireguardPeerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnWireguardPeerResultOutput)
}

// A collection of arguments for invoking getVpnWireguardPeer.
type LookupVpnWireguardPeerOutputArgs struct {
	// [String] The ID of the WireGuard Gateway.
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// [String] The ID of the WireGuard Peer.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// [String] The location of the WireGuard Gateway.
	Location pulumi.StringInput `pulumi:"location"`
	// [String] The name of the WireGuard Peer.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupVpnWireguardPeerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnWireguardPeerArgs)(nil)).Elem()
}

// A collection of values returned by getVpnWireguardPeer.
type LookupVpnWireguardPeerResultOutput struct{ *pulumi.OutputState }

func (LookupVpnWireguardPeerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnWireguardPeerResult)(nil)).Elem()
}

func (o LookupVpnWireguardPeerResultOutput) ToLookupVpnWireguardPeerResultOutput() LookupVpnWireguardPeerResultOutput {
	return o
}

func (o LookupVpnWireguardPeerResultOutput) ToLookupVpnWireguardPeerResultOutputWithContext(ctx context.Context) LookupVpnWireguardPeerResultOutput {
	return o
}

// The subnet CIDRs that are allowed to connect to the WireGuard Gateway.
func (o LookupVpnWireguardPeerResultOutput) AllowedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) []string { return v.AllowedIps }).(pulumi.StringArrayOutput)
}

// The description of the WireGuard Peer.
func (o LookupVpnWireguardPeerResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.Description }).(pulumi.StringOutput)
}

// The endpoint of the WireGuard Peer.
func (o LookupVpnWireguardPeerResultOutput) Endpoints() GetVpnWireguardPeerEndpointArrayOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) []GetVpnWireguardPeerEndpoint { return v.Endpoints }).(GetVpnWireguardPeerEndpointArrayOutput)
}

func (o LookupVpnWireguardPeerResultOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.GatewayId }).(pulumi.StringOutput)
}

// The unique ID of the WireGuard Peer.
func (o LookupVpnWireguardPeerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpnWireguardPeerResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the WireGuard Peer.
func (o LookupVpnWireguardPeerResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.Name }).(pulumi.StringOutput)
}

// WireGuard public key of the connecting peer.
func (o LookupVpnWireguardPeerResultOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.PublicKey }).(pulumi.StringOutput)
}

// The current status of the WireGuard Peer.
func (o LookupVpnWireguardPeerResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnWireguardPeerResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnWireguardPeerResultOutput{})
}
