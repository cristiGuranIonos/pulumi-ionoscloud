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

// This page provides an overview of the `VpnWireguardPeer` resource, which allows you to manage a WireGuard Peer in your cloud infrastructure.
// This resource enables the creation, management, and deletion of a WireGuard VPN Peer, facilitating secure connections between your network resources.
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
//			_, err := ionoscloud.NewVpnWireguardPeer(ctx, "example", &ionoscloud.VpnWireguardPeerArgs{
//				AllowedIps: pulumi.StringArray{
//					pulumi.String("10.0.0.0/8"),
//					pulumi.String("192.168.1.0/24"),
//				},
//				Description: pulumi.String("An example WireGuard peer"),
//				Endpoint: &ionoscloud.VpnWireguardPeerEndpointArgs{
//					Host: pulumi.String("1.2.3.4"),
//					Port: pulumi.Int(51820),
//				},
//				GatewayId: pulumi.String("your gateway id here"),
//				Location:  pulumi.String("de/fra"),
//				PublicKey: pulumi.String("examplePublicKey=="),
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
// WireGuard Peers can be imported using the `gateway_id` and `id`, e.g.,
//
// ```sh
// $ pulumi import ionoscloud:index/vpnWireguardPeer:VpnWireguardPeer example <gateway_id>:<peer_id>
// ```
type VpnWireguardPeer struct {
	pulumi.CustomResourceState

	// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps pulumi.StringArrayOutput `pulumi:"allowedIps"`
	// [string] A description of the WireGuard Gateway.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
	Endpoint VpnWireguardPeerEndpointPtrOutput `pulumi:"endpoint"`
	// [string] The ID of the WireGuard Gateway that the Peer will connect to.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// [string] The location of the WireGuard Gateway.
	Location pulumi.StringOutput `pulumi:"location"`
	// [string] The human-readable name of the WireGuard Gateway.
	Name pulumi.StringOutput `pulumi:"name"`
	// [string] The public key for the WireGuard Gateway.
	PublicKey pulumi.StringOutput `pulumi:"publicKey"`
	// The current status of the WireGuard Gateway Peer.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewVpnWireguardPeer registers a new resource with the given unique name, arguments, and options.
func NewVpnWireguardPeer(ctx *pulumi.Context,
	name string, args *VpnWireguardPeerArgs, opts ...pulumi.ResourceOption) (*VpnWireguardPeer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AllowedIps == nil {
		return nil, errors.New("invalid value for required argument 'AllowedIps'")
	}
	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.PublicKey == nil {
		return nil, errors.New("invalid value for required argument 'PublicKey'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnWireguardPeer
	err := ctx.RegisterResource("ionoscloud:index/vpnWireguardPeer:VpnWireguardPeer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnWireguardPeer gets an existing VpnWireguardPeer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnWireguardPeer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnWireguardPeerState, opts ...pulumi.ResourceOption) (*VpnWireguardPeer, error) {
	var resource VpnWireguardPeer
	err := ctx.ReadResource("ionoscloud:index/vpnWireguardPeer:VpnWireguardPeer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnWireguardPeer resources.
type vpnWireguardPeerState struct {
	// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps []string `pulumi:"allowedIps"`
	// [string] A description of the WireGuard Gateway.
	Description *string `pulumi:"description"`
	// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
	Endpoint *VpnWireguardPeerEndpoint `pulumi:"endpoint"`
	// [string] The ID of the WireGuard Gateway that the Peer will connect to.
	GatewayId *string `pulumi:"gatewayId"`
	// [string] The location of the WireGuard Gateway.
	Location *string `pulumi:"location"`
	// [string] The human-readable name of the WireGuard Gateway.
	Name *string `pulumi:"name"`
	// [string] The public key for the WireGuard Gateway.
	PublicKey *string `pulumi:"publicKey"`
	// The current status of the WireGuard Gateway Peer.
	Status *string `pulumi:"status"`
}

type VpnWireguardPeerState struct {
	// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps pulumi.StringArrayInput
	// [string] A description of the WireGuard Gateway.
	Description pulumi.StringPtrInput
	// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
	Endpoint VpnWireguardPeerEndpointPtrInput
	// [string] The ID of the WireGuard Gateway that the Peer will connect to.
	GatewayId pulumi.StringPtrInput
	// [string] The location of the WireGuard Gateway.
	Location pulumi.StringPtrInput
	// [string] The human-readable name of the WireGuard Gateway.
	Name pulumi.StringPtrInput
	// [string] The public key for the WireGuard Gateway.
	PublicKey pulumi.StringPtrInput
	// The current status of the WireGuard Gateway Peer.
	Status pulumi.StringPtrInput
}

func (VpnWireguardPeerState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnWireguardPeerState)(nil)).Elem()
}

type vpnWireguardPeerArgs struct {
	// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps []string `pulumi:"allowedIps"`
	// [string] A description of the WireGuard Gateway.
	Description *string `pulumi:"description"`
	// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
	Endpoint *VpnWireguardPeerEndpoint `pulumi:"endpoint"`
	// [string] The ID of the WireGuard Gateway that the Peer will connect to.
	GatewayId string `pulumi:"gatewayId"`
	// [string] The location of the WireGuard Gateway.
	Location string `pulumi:"location"`
	// [string] The human-readable name of the WireGuard Gateway.
	Name *string `pulumi:"name"`
	// [string] The public key for the WireGuard Gateway.
	PublicKey string `pulumi:"publicKey"`
}

// The set of arguments for constructing a VpnWireguardPeer resource.
type VpnWireguardPeerArgs struct {
	// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
	AllowedIps pulumi.StringArrayInput
	// [string] A description of the WireGuard Gateway.
	Description pulumi.StringPtrInput
	// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
	Endpoint VpnWireguardPeerEndpointPtrInput
	// [string] The ID of the WireGuard Gateway that the Peer will connect to.
	GatewayId pulumi.StringInput
	// [string] The location of the WireGuard Gateway.
	Location pulumi.StringInput
	// [string] The human-readable name of the WireGuard Gateway.
	Name pulumi.StringPtrInput
	// [string] The public key for the WireGuard Gateway.
	PublicKey pulumi.StringInput
}

func (VpnWireguardPeerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnWireguardPeerArgs)(nil)).Elem()
}

type VpnWireguardPeerInput interface {
	pulumi.Input

	ToVpnWireguardPeerOutput() VpnWireguardPeerOutput
	ToVpnWireguardPeerOutputWithContext(ctx context.Context) VpnWireguardPeerOutput
}

func (*VpnWireguardPeer) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnWireguardPeer)(nil)).Elem()
}

func (i *VpnWireguardPeer) ToVpnWireguardPeerOutput() VpnWireguardPeerOutput {
	return i.ToVpnWireguardPeerOutputWithContext(context.Background())
}

func (i *VpnWireguardPeer) ToVpnWireguardPeerOutputWithContext(ctx context.Context) VpnWireguardPeerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnWireguardPeerOutput)
}

// VpnWireguardPeerArrayInput is an input type that accepts VpnWireguardPeerArray and VpnWireguardPeerArrayOutput values.
// You can construct a concrete instance of `VpnWireguardPeerArrayInput` via:
//
//	VpnWireguardPeerArray{ VpnWireguardPeerArgs{...} }
type VpnWireguardPeerArrayInput interface {
	pulumi.Input

	ToVpnWireguardPeerArrayOutput() VpnWireguardPeerArrayOutput
	ToVpnWireguardPeerArrayOutputWithContext(context.Context) VpnWireguardPeerArrayOutput
}

type VpnWireguardPeerArray []VpnWireguardPeerInput

func (VpnWireguardPeerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnWireguardPeer)(nil)).Elem()
}

func (i VpnWireguardPeerArray) ToVpnWireguardPeerArrayOutput() VpnWireguardPeerArrayOutput {
	return i.ToVpnWireguardPeerArrayOutputWithContext(context.Background())
}

func (i VpnWireguardPeerArray) ToVpnWireguardPeerArrayOutputWithContext(ctx context.Context) VpnWireguardPeerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnWireguardPeerArrayOutput)
}

// VpnWireguardPeerMapInput is an input type that accepts VpnWireguardPeerMap and VpnWireguardPeerMapOutput values.
// You can construct a concrete instance of `VpnWireguardPeerMapInput` via:
//
//	VpnWireguardPeerMap{ "key": VpnWireguardPeerArgs{...} }
type VpnWireguardPeerMapInput interface {
	pulumi.Input

	ToVpnWireguardPeerMapOutput() VpnWireguardPeerMapOutput
	ToVpnWireguardPeerMapOutputWithContext(context.Context) VpnWireguardPeerMapOutput
}

type VpnWireguardPeerMap map[string]VpnWireguardPeerInput

func (VpnWireguardPeerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnWireguardPeer)(nil)).Elem()
}

func (i VpnWireguardPeerMap) ToVpnWireguardPeerMapOutput() VpnWireguardPeerMapOutput {
	return i.ToVpnWireguardPeerMapOutputWithContext(context.Background())
}

func (i VpnWireguardPeerMap) ToVpnWireguardPeerMapOutputWithContext(ctx context.Context) VpnWireguardPeerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnWireguardPeerMapOutput)
}

type VpnWireguardPeerOutput struct{ *pulumi.OutputState }

func (VpnWireguardPeerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnWireguardPeer)(nil)).Elem()
}

func (o VpnWireguardPeerOutput) ToVpnWireguardPeerOutput() VpnWireguardPeerOutput {
	return o
}

func (o VpnWireguardPeerOutput) ToVpnWireguardPeerOutputWithContext(ctx context.Context) VpnWireguardPeerOutput {
	return o
}

// [list, string] A list of subnet CIDRs that are allowed to connect to the WireGuard Gateway.
func (o VpnWireguardPeerOutput) AllowedIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringArrayOutput { return v.AllowedIps }).(pulumi.StringArrayOutput)
}

// [string] A description of the WireGuard Gateway.
func (o VpnWireguardPeerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// [block] An endpoint configuration block for the WireGuard Gateway. The structure of this block is as follows:
func (o VpnWireguardPeerOutput) Endpoint() VpnWireguardPeerEndpointPtrOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) VpnWireguardPeerEndpointPtrOutput { return v.Endpoint }).(VpnWireguardPeerEndpointPtrOutput)
}

// [string] The ID of the WireGuard Gateway that the Peer will connect to.
func (o VpnWireguardPeerOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// [string] The location of the WireGuard Gateway.
func (o VpnWireguardPeerOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// [string] The human-readable name of the WireGuard Gateway.
func (o VpnWireguardPeerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [string] The public key for the WireGuard Gateway.
func (o VpnWireguardPeerOutput) PublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringOutput { return v.PublicKey }).(pulumi.StringOutput)
}

// The current status of the WireGuard Gateway Peer.
func (o VpnWireguardPeerOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnWireguardPeer) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type VpnWireguardPeerArrayOutput struct{ *pulumi.OutputState }

func (VpnWireguardPeerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnWireguardPeer)(nil)).Elem()
}

func (o VpnWireguardPeerArrayOutput) ToVpnWireguardPeerArrayOutput() VpnWireguardPeerArrayOutput {
	return o
}

func (o VpnWireguardPeerArrayOutput) ToVpnWireguardPeerArrayOutputWithContext(ctx context.Context) VpnWireguardPeerArrayOutput {
	return o
}

func (o VpnWireguardPeerArrayOutput) Index(i pulumi.IntInput) VpnWireguardPeerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnWireguardPeer {
		return vs[0].([]*VpnWireguardPeer)[vs[1].(int)]
	}).(VpnWireguardPeerOutput)
}

type VpnWireguardPeerMapOutput struct{ *pulumi.OutputState }

func (VpnWireguardPeerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnWireguardPeer)(nil)).Elem()
}

func (o VpnWireguardPeerMapOutput) ToVpnWireguardPeerMapOutput() VpnWireguardPeerMapOutput {
	return o
}

func (o VpnWireguardPeerMapOutput) ToVpnWireguardPeerMapOutputWithContext(ctx context.Context) VpnWireguardPeerMapOutput {
	return o
}

func (o VpnWireguardPeerMapOutput) MapIndex(k pulumi.StringInput) VpnWireguardPeerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnWireguardPeer {
		return vs[0].(map[string]*VpnWireguardPeer)[vs[1].(string)]
	}).(VpnWireguardPeerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnWireguardPeerInput)(nil)).Elem(), &VpnWireguardPeer{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnWireguardPeerArrayInput)(nil)).Elem(), VpnWireguardPeerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnWireguardPeerMapInput)(nil)).Elem(), VpnWireguardPeerMap{})
	pulumi.RegisterOutputType(VpnWireguardPeerOutput{})
	pulumi.RegisterOutputType(VpnWireguardPeerArrayOutput{})
	pulumi.RegisterOutputType(VpnWireguardPeerMapOutput{})
}
