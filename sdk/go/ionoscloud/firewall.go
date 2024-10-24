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

// Manages a set of **Firewall Rules** on IonosCloud.
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
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/compute"
//	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDatacenter, err := compute.NewDatacenter(ctx, "exampleDatacenter", &compute.DatacenterArgs{
//				Location:          pulumi.String("us/las"),
//				Description:       pulumi.String("Datacenter Description"),
//				SecAuthProtection: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleIpblock, err := ionoscloud.NewIpblock(ctx, "exampleIpblock", &ionoscloud.IpblockArgs{
//				Location: exampleDatacenter.Location,
//				Size:     pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			serverImagePassword, err := random.NewRandomPassword(ctx, "serverImagePassword", &random.RandomPasswordArgs{
//				Length:  pulumi.Int(16),
//				Special: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleServer, err := ionoscloud.NewServer(ctx, "exampleServer", &ionoscloud.ServerArgs{
//				DatacenterId:     exampleDatacenter.ID(),
//				Cores:            pulumi.Int(1),
//				Ram:              pulumi.Int(1024),
//				AvailabilityZone: pulumi.String("ZONE_1"),
//				CpuFamily:        pulumi.String("INTEL_XEON"),
//				ImageName:        pulumi.String("Ubuntu-20.04"),
//				ImagePassword:    serverImagePassword.Result,
//				Volume: &ionoscloud.ServerVolumeArgs{
//					Name:     pulumi.String("system"),
//					Size:     pulumi.Int(14),
//					DiskType: pulumi.String("SSD"),
//				},
//				Nic: &ionoscloud.ServerNicArgs{
//					Lan:            pulumi.Int(1),
//					Dhcp:           pulumi.Bool(true),
//					FirewallActive: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			exampleNic, err := ionoscloud.NewNic(ctx, "exampleNic", &ionoscloud.NicArgs{
//				DatacenterId:   exampleDatacenter.ID(),
//				ServerId:       exampleServer.ID(),
//				Lan:            pulumi.Int(2),
//				Dhcp:           pulumi.Bool(true),
//				FirewallActive: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewFirewall(ctx, "exampleFirewall", &ionoscloud.FirewallArgs{
//				DatacenterId: exampleDatacenter.ID(),
//				ServerId:     exampleServer.ID(),
//				NicId:        exampleNic.ID(),
//				Protocol:     pulumi.String("ICMP"),
//				SourceMac:    pulumi.String("00:0a:95:9d:68:16"),
//				SourceIp: exampleIpblock.Ips.ApplyT(func(ips []string) (string, error) {
//					return ips[0], nil
//				}).(pulumi.StringOutput),
//				TargetIp: exampleIpblock.Ips.ApplyT(func(ips []string) (string, error) {
//					return ips[1], nil
//				}).(pulumi.StringOutput),
//				IcmpType: pulumi.String("1"),
//				IcmpCode: pulumi.String("8"),
//				Type:     pulumi.String("INGRESS"),
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
// Resource Firewall can be imported using the `resource id`, e.g.
//
// ```sh
// $ pulumi import ionoscloud:index/firewall:Firewall myfwrule {datacenter uuid}/{server uuid}/{nic uuid}/{firewall uuid}
// ```
type Firewall struct {
	pulumi.CustomResourceState

	// [string] The Virtual Data Center ID.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode pulumi.StringPtrOutput `pulumi:"icmpCode"`
	// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpType pulumi.StringPtrOutput `pulumi:"icmpType"`
	// [string] The name of the firewall rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// [string] The NIC ID.
	NicId pulumi.StringOutput `pulumi:"nicId"`
	// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd pulumi.IntPtrOutput `pulumi:"portRangeEnd"`
	// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeStart pulumi.IntPtrOutput `pulumi:"portRangeStart"`
	// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// [string] The Server ID.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
	// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
	SourceIp pulumi.StringOutput `pulumi:"sourceIp"`
	// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
	SourceMac pulumi.StringPtrOutput `pulumi:"sourceMac"`
	// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
	TargetIp pulumi.StringOutput `pulumi:"targetIp"`
	// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewFirewall registers a new resource with the given unique name, arguments, and options.
func NewFirewall(ctx *pulumi.Context,
	name string, args *FirewallArgs, opts ...pulumi.ResourceOption) (*Firewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatacenterId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterId'")
	}
	if args.NicId == nil {
		return nil, errors.New("invalid value for required argument 'NicId'")
	}
	if args.Protocol == nil {
		return nil, errors.New("invalid value for required argument 'Protocol'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Firewall
	err := ctx.RegisterResource("ionoscloud:index/firewall:Firewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewall gets an existing Firewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallState, opts ...pulumi.ResourceOption) (*Firewall, error) {
	var resource Firewall
	err := ctx.ReadResource("ionoscloud:index/firewall:Firewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Firewall resources.
type firewallState struct {
	// [string] The Virtual Data Center ID.
	DatacenterId *string `pulumi:"datacenterId"`
	// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode *string `pulumi:"icmpCode"`
	// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpType *string `pulumi:"icmpType"`
	// [string] The name of the firewall rule.
	Name *string `pulumi:"name"`
	// [string] The NIC ID.
	NicId *string `pulumi:"nicId"`
	// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd *int `pulumi:"portRangeEnd"`
	// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeStart *int `pulumi:"portRangeStart"`
	// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
	Protocol *string `pulumi:"protocol"`
	// [string] The Server ID.
	ServerId *string `pulumi:"serverId"`
	// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
	SourceIp *string `pulumi:"sourceIp"`
	// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
	SourceMac *string `pulumi:"sourceMac"`
	// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
	TargetIp *string `pulumi:"targetIp"`
	// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
	Type *string `pulumi:"type"`
}

type FirewallState struct {
	// [string] The Virtual Data Center ID.
	DatacenterId pulumi.StringPtrInput
	// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode pulumi.StringPtrInput
	// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpType pulumi.StringPtrInput
	// [string] The name of the firewall rule.
	Name pulumi.StringPtrInput
	// [string] The NIC ID.
	NicId pulumi.StringPtrInput
	// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd pulumi.IntPtrInput
	// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeStart pulumi.IntPtrInput
	// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
	Protocol pulumi.StringPtrInput
	// [string] The Server ID.
	ServerId pulumi.StringPtrInput
	// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
	SourceIp pulumi.StringPtrInput
	// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
	SourceMac pulumi.StringPtrInput
	// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
	TargetIp pulumi.StringPtrInput
	// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
	Type pulumi.StringPtrInput
}

func (FirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallState)(nil)).Elem()
}

type firewallArgs struct {
	// [string] The Virtual Data Center ID.
	DatacenterId string `pulumi:"datacenterId"`
	// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode *string `pulumi:"icmpCode"`
	// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpType *string `pulumi:"icmpType"`
	// [string] The name of the firewall rule.
	Name *string `pulumi:"name"`
	// [string] The NIC ID.
	NicId string `pulumi:"nicId"`
	// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd *int `pulumi:"portRangeEnd"`
	// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeStart *int `pulumi:"portRangeStart"`
	// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
	Protocol string `pulumi:"protocol"`
	// [string] The Server ID.
	ServerId string `pulumi:"serverId"`
	// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
	SourceIp *string `pulumi:"sourceIp"`
	// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
	SourceMac *string `pulumi:"sourceMac"`
	// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
	TargetIp *string `pulumi:"targetIp"`
	// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Firewall resource.
type FirewallArgs struct {
	// [string] The Virtual Data Center ID.
	DatacenterId pulumi.StringInput
	// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode pulumi.StringPtrInput
	// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
	IcmpType pulumi.StringPtrInput
	// [string] The name of the firewall rule.
	Name pulumi.StringPtrInput
	// [string] The NIC ID.
	NicId pulumi.StringInput
	// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeEnd pulumi.IntPtrInput
	// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
	PortRangeStart pulumi.IntPtrInput
	// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
	Protocol pulumi.StringInput
	// [string] The Server ID.
	ServerId pulumi.StringInput
	// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
	SourceIp pulumi.StringPtrInput
	// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
	SourceMac pulumi.StringPtrInput
	// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
	TargetIp pulumi.StringPtrInput
	// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
	Type pulumi.StringPtrInput
}

func (FirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallArgs)(nil)).Elem()
}

type FirewallInput interface {
	pulumi.Input

	ToFirewallOutput() FirewallOutput
	ToFirewallOutputWithContext(ctx context.Context) FirewallOutput
}

func (*Firewall) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (i *Firewall) ToFirewallOutput() FirewallOutput {
	return i.ToFirewallOutputWithContext(context.Background())
}

func (i *Firewall) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallOutput)
}

// FirewallArrayInput is an input type that accepts FirewallArray and FirewallArrayOutput values.
// You can construct a concrete instance of `FirewallArrayInput` via:
//
//	FirewallArray{ FirewallArgs{...} }
type FirewallArrayInput interface {
	pulumi.Input

	ToFirewallArrayOutput() FirewallArrayOutput
	ToFirewallArrayOutputWithContext(context.Context) FirewallArrayOutput
}

type FirewallArray []FirewallInput

func (FirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (i FirewallArray) ToFirewallArrayOutput() FirewallArrayOutput {
	return i.ToFirewallArrayOutputWithContext(context.Background())
}

func (i FirewallArray) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallArrayOutput)
}

// FirewallMapInput is an input type that accepts FirewallMap and FirewallMapOutput values.
// You can construct a concrete instance of `FirewallMapInput` via:
//
//	FirewallMap{ "key": FirewallArgs{...} }
type FirewallMapInput interface {
	pulumi.Input

	ToFirewallMapOutput() FirewallMapOutput
	ToFirewallMapOutputWithContext(context.Context) FirewallMapOutput
}

type FirewallMap map[string]FirewallInput

func (FirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (i FirewallMap) ToFirewallMapOutput() FirewallMapOutput {
	return i.ToFirewallMapOutputWithContext(context.Background())
}

func (i FirewallMap) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallMapOutput)
}

type FirewallOutput struct{ *pulumi.OutputState }

func (FirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Firewall)(nil)).Elem()
}

func (o FirewallOutput) ToFirewallOutput() FirewallOutput {
	return o
}

func (o FirewallOutput) ToFirewallOutputWithContext(ctx context.Context) FirewallOutput {
	return o
}

// [string] The Virtual Data Center ID.
func (o FirewallOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.DatacenterId }).(pulumi.StringOutput)
}

// [int] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
func (o FirewallOutput) IcmpCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.IcmpCode }).(pulumi.StringPtrOutput)
}

// [string] Defines the allowed code (from 0 to 254) if protocol ICMP is chosen. Value null allows all codes.
func (o FirewallOutput) IcmpType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.IcmpType }).(pulumi.StringPtrOutput)
}

// [string] The name of the firewall rule.
func (o FirewallOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [string] The NIC ID.
func (o FirewallOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.NicId }).(pulumi.StringOutput)
}

// [int] Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
func (o FirewallOutput) PortRangeEnd() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.IntPtrOutput { return v.PortRangeEnd }).(pulumi.IntPtrOutput)
}

// [int] Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen. Leave portRangeStart and portRangeEnd null to allow all ports.
func (o FirewallOutput) PortRangeStart() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.IntPtrOutput { return v.PortRangeStart }).(pulumi.IntPtrOutput)
}

// [string] The protocol for the rule: TCP, UDP, ICMP, ANY. Property cannot be modified after creation (disallowed in update requests).
func (o FirewallOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Protocol }).(pulumi.StringOutput)
}

// [string] The Server ID.
func (o FirewallOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

// [string] Only traffic originating from the respective IPv4 address is allowed. Value null allows all source IPs.
func (o FirewallOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.SourceIp }).(pulumi.StringOutput)
}

// [string] Only traffic originating from the respective MAC address is allowed. Valid format: aa:bb:cc:dd:ee:ff. Value null allows all source MAC address. Valid format: aa:bb:cc:dd:ee:ff.
func (o FirewallOutput) SourceMac() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringPtrOutput { return v.SourceMac }).(pulumi.StringPtrOutput)
}

// [string] In case the target NIC has multiple IP addresses, only traffic directed to the respective IP address of the NIC is allowed. Value null allows all target IPs.
func (o FirewallOutput) TargetIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.TargetIp }).(pulumi.StringOutput)
}

// [string] The type of firewall rule. If is not specified, it will take the default value INGRESS.
func (o FirewallOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Firewall) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type FirewallArrayOutput struct{ *pulumi.OutputState }

func (FirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Firewall)(nil)).Elem()
}

func (o FirewallArrayOutput) ToFirewallArrayOutput() FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) ToFirewallArrayOutputWithContext(ctx context.Context) FirewallArrayOutput {
	return o
}

func (o FirewallArrayOutput) Index(i pulumi.IntInput) FirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].([]*Firewall)[vs[1].(int)]
	}).(FirewallOutput)
}

type FirewallMapOutput struct{ *pulumi.OutputState }

func (FirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Firewall)(nil)).Elem()
}

func (o FirewallMapOutput) ToFirewallMapOutput() FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) ToFirewallMapOutputWithContext(ctx context.Context) FirewallMapOutput {
	return o
}

func (o FirewallMapOutput) MapIndex(k pulumi.StringInput) FirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Firewall {
		return vs[0].(map[string]*Firewall)[vs[1].(string)]
	}).(FirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallInput)(nil)).Elem(), &Firewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallArrayInput)(nil)).Elem(), FirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallMapInput)(nil)).Elem(), FirewallMap{})
	pulumi.RegisterOutputType(FirewallOutput{})
	pulumi.RegisterOutputType(FirewallArrayOutput{})
	pulumi.RegisterOutputType(FirewallMapOutput{})
}