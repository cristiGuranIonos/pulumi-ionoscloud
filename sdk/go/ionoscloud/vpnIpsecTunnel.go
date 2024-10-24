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

// An IPSec Gateway Tunnel resource manages the creation, management, and deletion of VPN IPSec Gateway Tunnels within the
// IONOS Cloud infrastructure. This resource facilitates the creation of VPN IPSec Gateway Tunnels, enabling secure
// connections between your network resources.
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
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			// Basic example
//			testDatacenter, err := compute.NewDatacenter(ctx, "testDatacenter", &compute.DatacenterArgs{
//				Location: pulumi.String("de/fra"),
//			})
//			if err != nil {
//				return err
//			}
//			testLan, err := ionoscloud.NewLan(ctx, "testLan", &ionoscloud.LanArgs{
//				Public:       pulumi.Bool(false),
//				DatacenterId: testDatacenter.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			testIpblock, err := ionoscloud.NewIpblock(ctx, "testIpblock", &ionoscloud.IpblockArgs{
//				Location: pulumi.String("de/fra"),
//				Size:     pulumi.Int(1),
//			})
//			if err != nil {
//				return err
//			}
//			exampleVpnIpsecGateway, err := ionoscloud.NewVpnIpsecGateway(ctx, "exampleVpnIpsecGateway", &ionoscloud.VpnIpsecGatewayArgs{
//				Location: pulumi.String("de/fra"),
//				GatewayIp: testIpblock.Ips.ApplyT(func(ips []string) (string, error) {
//					return ips[0], nil
//				}).(pulumi.StringOutput),
//				Version:     pulumi.String("IKEv2"),
//				Description: pulumi.String("This gateway connects site A to VDC X."),
//				Connections: ionoscloud.VpnIpsecGatewayConnectionArray{
//					&ionoscloud.VpnIpsecGatewayConnectionArgs{
//						DatacenterId: testDatacenter.ID(),
//						LanId:        testLan.ID(),
//						Ipv4Cidr:     pulumi.String("192.168.100.10/24"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewVpnIpsecTunnel(ctx, "exampleVpnIpsecTunnel", &ionoscloud.VpnIpsecTunnelArgs{
//				Location:    pulumi.String("de/fra"),
//				GatewayId:   exampleVpnIpsecGateway.ID(),
//				RemoteHost:  pulumi.String("vpn.mycompany.com"),
//				Description: pulumi.String("Allows local subnet X to connect to virtual network Y."),
//				Auth: &ionoscloud.VpnIpsecTunnelAuthArgs{
//					Method: pulumi.String("PSK"),
//					PskKey: pulumi.String("X2wosbaw74M8hQGbK3jCCaEusR6CCFRa"),
//				},
//				Ike: &ionoscloud.VpnIpsecTunnelIkeArgs{
//					DiffieHellmanGroup:  pulumi.String("16-MODP4096"),
//					EncryptionAlgorithm: pulumi.String("AES256"),
//					IntegrityAlgorithm:  pulumi.String("SHA256"),
//					Lifetime:            pulumi.Int(86400),
//				},
//				Esps: ionoscloud.VpnIpsecTunnelEspArray{
//					&ionoscloud.VpnIpsecTunnelEspArgs{
//						DiffieHellmanGroup:  pulumi.String("16-MODP4096"),
//						EncryptionAlgorithm: pulumi.String("AES256"),
//						IntegrityAlgorithm:  pulumi.String("SHA256"),
//						Lifetime:            pulumi.Int(3600),
//					},
//				},
//				CloudNetworkCidrs: pulumi.StringArray{
//					pulumi.String("0.0.0.0/0"),
//				},
//				PeerNetworkCidrs: pulumi.StringArray{
//					pulumi.String("1.2.3.4/32"),
//				},
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
// The resource can be imported using the `location`, `gateway_id` and `tunnel_id`, for example:
//
// ```sh
// $ pulumi import ionoscloud:index/vpnIpsecTunnel:VpnIpsecTunnel example {location}:{gateway_id}:{tunnel_id}
// ```
type VpnIpsecTunnel struct {
	pulumi.CustomResourceState

	// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
	// items: 1.
	Auth VpnIpsecTunnelAuthOutput `pulumi:"auth"`
	// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
	// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
	// Maximum items: 20.
	CloudNetworkCidrs pulumi.StringArrayOutput `pulumi:"cloudNetworkCidrs"`
	// [string] The human-readable description of your IPSec Gateway Tunnel.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
	Esps VpnIpsecTunnelEspArrayOutput `pulumi:"esps"`
	// [string] The ID of the IPSec Gateway that the tunnel belongs to.
	GatewayId pulumi.StringOutput `pulumi:"gatewayId"`
	// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
	Ike VpnIpsecTunnelIkeOutput `pulumi:"ike"`
	// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location pulumi.StringOutput `pulumi:"location"`
	// [string] The name of the IPSec Gateway Tunnel.
	Name pulumi.StringOutput `pulumi:"name"`
	// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
	// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
	PeerNetworkCidrs pulumi.StringArrayOutput `pulumi:"peerNetworkCidrs"`
	// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	RemoteHost pulumi.StringOutput `pulumi:"remoteHost"`
}

// NewVpnIpsecTunnel registers a new resource with the given unique name, arguments, and options.
func NewVpnIpsecTunnel(ctx *pulumi.Context,
	name string, args *VpnIpsecTunnelArgs, opts ...pulumi.ResourceOption) (*VpnIpsecTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Auth == nil {
		return nil, errors.New("invalid value for required argument 'Auth'")
	}
	if args.CloudNetworkCidrs == nil {
		return nil, errors.New("invalid value for required argument 'CloudNetworkCidrs'")
	}
	if args.Esps == nil {
		return nil, errors.New("invalid value for required argument 'Esps'")
	}
	if args.GatewayId == nil {
		return nil, errors.New("invalid value for required argument 'GatewayId'")
	}
	if args.Ike == nil {
		return nil, errors.New("invalid value for required argument 'Ike'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.PeerNetworkCidrs == nil {
		return nil, errors.New("invalid value for required argument 'PeerNetworkCidrs'")
	}
	if args.RemoteHost == nil {
		return nil, errors.New("invalid value for required argument 'RemoteHost'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VpnIpsecTunnel
	err := ctx.RegisterResource("ionoscloud:index/vpnIpsecTunnel:VpnIpsecTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnIpsecTunnel gets an existing VpnIpsecTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnIpsecTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnIpsecTunnelState, opts ...pulumi.ResourceOption) (*VpnIpsecTunnel, error) {
	var resource VpnIpsecTunnel
	err := ctx.ReadResource("ionoscloud:index/vpnIpsecTunnel:VpnIpsecTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnIpsecTunnel resources.
type vpnIpsecTunnelState struct {
	// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
	// items: 1.
	Auth *VpnIpsecTunnelAuth `pulumi:"auth"`
	// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
	// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
	// Maximum items: 20.
	CloudNetworkCidrs []string `pulumi:"cloudNetworkCidrs"`
	// [string] The human-readable description of your IPSec Gateway Tunnel.
	Description *string `pulumi:"description"`
	// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
	Esps []VpnIpsecTunnelEsp `pulumi:"esps"`
	// [string] The ID of the IPSec Gateway that the tunnel belongs to.
	GatewayId *string `pulumi:"gatewayId"`
	// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
	Ike *VpnIpsecTunnelIke `pulumi:"ike"`
	// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location *string `pulumi:"location"`
	// [string] The name of the IPSec Gateway Tunnel.
	Name *string `pulumi:"name"`
	// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
	// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
	PeerNetworkCidrs []string `pulumi:"peerNetworkCidrs"`
	// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	RemoteHost *string `pulumi:"remoteHost"`
}

type VpnIpsecTunnelState struct {
	// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
	// items: 1.
	Auth VpnIpsecTunnelAuthPtrInput
	// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
	// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
	// Maximum items: 20.
	CloudNetworkCidrs pulumi.StringArrayInput
	// [string] The human-readable description of your IPSec Gateway Tunnel.
	Description pulumi.StringPtrInput
	// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
	Esps VpnIpsecTunnelEspArrayInput
	// [string] The ID of the IPSec Gateway that the tunnel belongs to.
	GatewayId pulumi.StringPtrInput
	// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
	Ike VpnIpsecTunnelIkePtrInput
	// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location pulumi.StringPtrInput
	// [string] The name of the IPSec Gateway Tunnel.
	Name pulumi.StringPtrInput
	// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
	// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
	PeerNetworkCidrs pulumi.StringArrayInput
	// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	RemoteHost pulumi.StringPtrInput
}

func (VpnIpsecTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecTunnelState)(nil)).Elem()
}

type vpnIpsecTunnelArgs struct {
	// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
	// items: 1.
	Auth VpnIpsecTunnelAuth `pulumi:"auth"`
	// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
	// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
	// Maximum items: 20.
	CloudNetworkCidrs []string `pulumi:"cloudNetworkCidrs"`
	// [string] The human-readable description of your IPSec Gateway Tunnel.
	Description *string `pulumi:"description"`
	// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
	Esps []VpnIpsecTunnelEsp `pulumi:"esps"`
	// [string] The ID of the IPSec Gateway that the tunnel belongs to.
	GatewayId string `pulumi:"gatewayId"`
	// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
	Ike VpnIpsecTunnelIke `pulumi:"ike"`
	// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location string `pulumi:"location"`
	// [string] The name of the IPSec Gateway Tunnel.
	Name *string `pulumi:"name"`
	// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
	// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
	PeerNetworkCidrs []string `pulumi:"peerNetworkCidrs"`
	// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	RemoteHost string `pulumi:"remoteHost"`
}

// The set of arguments for constructing a VpnIpsecTunnel resource.
type VpnIpsecTunnelArgs struct {
	// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
	// items: 1.
	Auth VpnIpsecTunnelAuthInput
	// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
	// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
	// Maximum items: 20.
	CloudNetworkCidrs pulumi.StringArrayInput
	// [string] The human-readable description of your IPSec Gateway Tunnel.
	Description pulumi.StringPtrInput
	// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
	Esps VpnIpsecTunnelEspArrayInput
	// [string] The ID of the IPSec Gateway that the tunnel belongs to.
	GatewayId pulumi.StringInput
	// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
	Ike VpnIpsecTunnelIkeInput
	// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
	// gb/lhr, us/ewr, us/las, us/mci, fr/par
	Location pulumi.StringInput
	// [string] The name of the IPSec Gateway Tunnel.
	Name pulumi.StringPtrInput
	// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
	// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
	PeerNetworkCidrs pulumi.StringArrayInput
	// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
	RemoteHost pulumi.StringInput
}

func (VpnIpsecTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnIpsecTunnelArgs)(nil)).Elem()
}

type VpnIpsecTunnelInput interface {
	pulumi.Input

	ToVpnIpsecTunnelOutput() VpnIpsecTunnelOutput
	ToVpnIpsecTunnelOutputWithContext(ctx context.Context) VpnIpsecTunnelOutput
}

func (*VpnIpsecTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecTunnel)(nil)).Elem()
}

func (i *VpnIpsecTunnel) ToVpnIpsecTunnelOutput() VpnIpsecTunnelOutput {
	return i.ToVpnIpsecTunnelOutputWithContext(context.Background())
}

func (i *VpnIpsecTunnel) ToVpnIpsecTunnelOutputWithContext(ctx context.Context) VpnIpsecTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecTunnelOutput)
}

// VpnIpsecTunnelArrayInput is an input type that accepts VpnIpsecTunnelArray and VpnIpsecTunnelArrayOutput values.
// You can construct a concrete instance of `VpnIpsecTunnelArrayInput` via:
//
//	VpnIpsecTunnelArray{ VpnIpsecTunnelArgs{...} }
type VpnIpsecTunnelArrayInput interface {
	pulumi.Input

	ToVpnIpsecTunnelArrayOutput() VpnIpsecTunnelArrayOutput
	ToVpnIpsecTunnelArrayOutputWithContext(context.Context) VpnIpsecTunnelArrayOutput
}

type VpnIpsecTunnelArray []VpnIpsecTunnelInput

func (VpnIpsecTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecTunnel)(nil)).Elem()
}

func (i VpnIpsecTunnelArray) ToVpnIpsecTunnelArrayOutput() VpnIpsecTunnelArrayOutput {
	return i.ToVpnIpsecTunnelArrayOutputWithContext(context.Background())
}

func (i VpnIpsecTunnelArray) ToVpnIpsecTunnelArrayOutputWithContext(ctx context.Context) VpnIpsecTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecTunnelArrayOutput)
}

// VpnIpsecTunnelMapInput is an input type that accepts VpnIpsecTunnelMap and VpnIpsecTunnelMapOutput values.
// You can construct a concrete instance of `VpnIpsecTunnelMapInput` via:
//
//	VpnIpsecTunnelMap{ "key": VpnIpsecTunnelArgs{...} }
type VpnIpsecTunnelMapInput interface {
	pulumi.Input

	ToVpnIpsecTunnelMapOutput() VpnIpsecTunnelMapOutput
	ToVpnIpsecTunnelMapOutputWithContext(context.Context) VpnIpsecTunnelMapOutput
}

type VpnIpsecTunnelMap map[string]VpnIpsecTunnelInput

func (VpnIpsecTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecTunnel)(nil)).Elem()
}

func (i VpnIpsecTunnelMap) ToVpnIpsecTunnelMapOutput() VpnIpsecTunnelMapOutput {
	return i.ToVpnIpsecTunnelMapOutputWithContext(context.Background())
}

func (i VpnIpsecTunnelMap) ToVpnIpsecTunnelMapOutputWithContext(ctx context.Context) VpnIpsecTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VpnIpsecTunnelMapOutput)
}

type VpnIpsecTunnelOutput struct{ *pulumi.OutputState }

func (VpnIpsecTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VpnIpsecTunnel)(nil)).Elem()
}

func (o VpnIpsecTunnelOutput) ToVpnIpsecTunnelOutput() VpnIpsecTunnelOutput {
	return o
}

func (o VpnIpsecTunnelOutput) ToVpnIpsecTunnelOutputWithContext(ctx context.Context) VpnIpsecTunnelOutput {
	return o
}

// [string] Properties with all data needed to define IPSec Authentication. Minimum items: 1. Maximum
// items: 1.
func (o VpnIpsecTunnelOutput) Auth() VpnIpsecTunnelAuthOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) VpnIpsecTunnelAuthOutput { return v.Auth }).(VpnIpsecTunnelAuthOutput)
}

// [list] The network CIDRs on the "Left" side that are allowed to connect to the IPSec
// tunnel, i.e. the CIDRs within your IONOS Cloud LAN. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1.
// Maximum items: 20.
func (o VpnIpsecTunnelOutput) CloudNetworkCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringArrayOutput { return v.CloudNetworkCidrs }).(pulumi.StringArrayOutput)
}

// [string] The human-readable description of your IPSec Gateway Tunnel.
func (o VpnIpsecTunnelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// [list] Settings for the IPSec SA (ESP) phase. Minimum items: 1. Maximum items: 1.
func (o VpnIpsecTunnelOutput) Esps() VpnIpsecTunnelEspArrayOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) VpnIpsecTunnelEspArrayOutput { return v.Esps }).(VpnIpsecTunnelEspArrayOutput)
}

// [string] The ID of the IPSec Gateway that the tunnel belongs to.
func (o VpnIpsecTunnelOutput) GatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringOutput { return v.GatewayId }).(pulumi.StringOutput)
}

// [list] Settings for the initial security exchange phase. Minimum items: 1. Maximum items: 1.
func (o VpnIpsecTunnelOutput) Ike() VpnIpsecTunnelIkeOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) VpnIpsecTunnelIkeOutput { return v.Ike }).(VpnIpsecTunnelIkeOutput)
}

// [string] The location of the IPSec Gateway Tunnel. Supported locations: de/fra, de/txl, es/vit,
// gb/lhr, us/ewr, us/las, us/mci, fr/par
func (o VpnIpsecTunnelOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// [string] The name of the IPSec Gateway Tunnel.
func (o VpnIpsecTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [list] The network CIDRs on the "Right" side that are allowed to connect to the IPSec
// tunnel. Specify "0.0.0.0/0" or "::/0" for all addresses. Minimum items: 1. Maximum items: 20.
func (o VpnIpsecTunnelOutput) PeerNetworkCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringArrayOutput { return v.PeerNetworkCidrs }).(pulumi.StringArrayOutput)
}

// [string] The remote peer host fully qualified domain name or public IPV4 IP to connect to.
func (o VpnIpsecTunnelOutput) RemoteHost() pulumi.StringOutput {
	return o.ApplyT(func(v *VpnIpsecTunnel) pulumi.StringOutput { return v.RemoteHost }).(pulumi.StringOutput)
}

type VpnIpsecTunnelArrayOutput struct{ *pulumi.OutputState }

func (VpnIpsecTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VpnIpsecTunnel)(nil)).Elem()
}

func (o VpnIpsecTunnelArrayOutput) ToVpnIpsecTunnelArrayOutput() VpnIpsecTunnelArrayOutput {
	return o
}

func (o VpnIpsecTunnelArrayOutput) ToVpnIpsecTunnelArrayOutputWithContext(ctx context.Context) VpnIpsecTunnelArrayOutput {
	return o
}

func (o VpnIpsecTunnelArrayOutput) Index(i pulumi.IntInput) VpnIpsecTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VpnIpsecTunnel {
		return vs[0].([]*VpnIpsecTunnel)[vs[1].(int)]
	}).(VpnIpsecTunnelOutput)
}

type VpnIpsecTunnelMapOutput struct{ *pulumi.OutputState }

func (VpnIpsecTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VpnIpsecTunnel)(nil)).Elem()
}

func (o VpnIpsecTunnelMapOutput) ToVpnIpsecTunnelMapOutput() VpnIpsecTunnelMapOutput {
	return o
}

func (o VpnIpsecTunnelMapOutput) ToVpnIpsecTunnelMapOutputWithContext(ctx context.Context) VpnIpsecTunnelMapOutput {
	return o
}

func (o VpnIpsecTunnelMapOutput) MapIndex(k pulumi.StringInput) VpnIpsecTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VpnIpsecTunnel {
		return vs[0].(map[string]*VpnIpsecTunnel)[vs[1].(string)]
	}).(VpnIpsecTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecTunnelInput)(nil)).Elem(), &VpnIpsecTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecTunnelArrayInput)(nil)).Elem(), VpnIpsecTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VpnIpsecTunnelMapInput)(nil)).Elem(), VpnIpsecTunnelMap{})
	pulumi.RegisterOutputType(VpnIpsecTunnelOutput{})
	pulumi.RegisterOutputType(VpnIpsecTunnelArrayOutput{})
	pulumi.RegisterOutputType(VpnIpsecTunnelMapOutput{})
}