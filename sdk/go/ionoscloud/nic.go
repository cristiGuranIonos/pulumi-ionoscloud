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

// ## Import
//
// Resource **Nic** can be imported using the `resource id`, e.g.
//
// ```sh
// $ pulumi import ionoscloud:index/nic:Nic mynic {datacenter uuid}/{server uuid}/{nic uuid}
// ```
type Nic struct {
	pulumi.CustomResourceState

	// [string] The ID of a Virtual Data Center.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created from CloudAPI and no DCD changes were done on the Datacenter.
	DeviceNumber pulumi.IntOutput `pulumi:"deviceNumber"`
	// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
	Dhcp pulumi.BoolPtrOutput `pulumi:"dhcp"`
	// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
	Dhcpv6 pulumi.BoolPtrOutput `pulumi:"dhcpv6"`
	// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
	FirewallActive pulumi.BoolPtrOutput `pulumi:"firewallActive"`
	// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
	FirewallType pulumi.StringOutput `pulumi:"firewallType"`
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlog NicFlowlogPtrOutput `pulumi:"flowlog"`
	// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips pulumi.StringArrayOutput `pulumi:"ips"`
	// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
	Ipv6CidrBlock pulumi.StringOutput `pulumi:"ipv6CidrBlock"`
	// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
	Ipv6Ips pulumi.StringArrayOutput `pulumi:"ipv6Ips"`
	// [integer] The LAN ID the NIC will sit on.
	Lan pulumi.IntOutput `pulumi:"lan"`
	// The MAC address of the NIC.
	Mac pulumi.StringOutput `pulumi:"mac"`
	// Specifies the name of the flow log.
	//
	// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The PCI slot number of the Nic.
	PciSlot pulumi.IntOutput `pulumi:"pciSlot"`
	// [string] The ID of a server.
	ServerId pulumi.StringOutput `pulumi:"serverId"`
}

// NewNic registers a new resource with the given unique name, arguments, and options.
func NewNic(ctx *pulumi.Context,
	name string, args *NicArgs, opts ...pulumi.ResourceOption) (*Nic, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatacenterId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterId'")
	}
	if args.Lan == nil {
		return nil, errors.New("invalid value for required argument 'Lan'")
	}
	if args.ServerId == nil {
		return nil, errors.New("invalid value for required argument 'ServerId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Nic
	err := ctx.RegisterResource("ionoscloud:index/nic:Nic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNic gets an existing Nic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NicState, opts ...pulumi.ResourceOption) (*Nic, error) {
	var resource Nic
	err := ctx.ReadResource("ionoscloud:index/nic:Nic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Nic resources.
type nicState struct {
	// [string] The ID of a Virtual Data Center.
	DatacenterId *string `pulumi:"datacenterId"`
	// The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created from CloudAPI and no DCD changes were done on the Datacenter.
	DeviceNumber *int `pulumi:"deviceNumber"`
	// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
	Dhcp *bool `pulumi:"dhcp"`
	// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
	Dhcpv6 *bool `pulumi:"dhcpv6"`
	// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
	FirewallActive *bool `pulumi:"firewallActive"`
	// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
	FirewallType *string `pulumi:"firewallType"`
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlog *NicFlowlog `pulumi:"flowlog"`
	// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips []string `pulumi:"ips"`
	// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
	Ipv6Ips []string `pulumi:"ipv6Ips"`
	// [integer] The LAN ID the NIC will sit on.
	Lan *int `pulumi:"lan"`
	// The MAC address of the NIC.
	Mac *string `pulumi:"mac"`
	// Specifies the name of the flow log.
	//
	// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
	Name *string `pulumi:"name"`
	// The PCI slot number of the Nic.
	PciSlot *int `pulumi:"pciSlot"`
	// [string] The ID of a server.
	ServerId *string `pulumi:"serverId"`
}

type NicState struct {
	// [string] The ID of a Virtual Data Center.
	DatacenterId pulumi.StringPtrInput
	// The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created from CloudAPI and no DCD changes were done on the Datacenter.
	DeviceNumber pulumi.IntPtrInput
	// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
	Dhcp pulumi.BoolPtrInput
	// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
	Dhcpv6 pulumi.BoolPtrInput
	// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
	FirewallActive pulumi.BoolPtrInput
	// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
	FirewallType pulumi.StringPtrInput
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlog NicFlowlogPtrInput
	// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips pulumi.StringArrayInput
	// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
	Ipv6CidrBlock pulumi.StringPtrInput
	// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
	Ipv6Ips pulumi.StringArrayInput
	// [integer] The LAN ID the NIC will sit on.
	Lan pulumi.IntPtrInput
	// The MAC address of the NIC.
	Mac pulumi.StringPtrInput
	// Specifies the name of the flow log.
	//
	// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
	Name pulumi.StringPtrInput
	// The PCI slot number of the Nic.
	PciSlot pulumi.IntPtrInput
	// [string] The ID of a server.
	ServerId pulumi.StringPtrInput
}

func (NicState) ElementType() reflect.Type {
	return reflect.TypeOf((*nicState)(nil)).Elem()
}

type nicArgs struct {
	// [string] The ID of a Virtual Data Center.
	DatacenterId string `pulumi:"datacenterId"`
	// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
	Dhcp *bool `pulumi:"dhcp"`
	// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
	Dhcpv6 *bool `pulumi:"dhcpv6"`
	// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
	FirewallActive *bool `pulumi:"firewallActive"`
	// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
	FirewallType *string `pulumi:"firewallType"`
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlog *NicFlowlog `pulumi:"flowlog"`
	// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips []string `pulumi:"ips"`
	// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
	Ipv6CidrBlock *string `pulumi:"ipv6CidrBlock"`
	// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
	Ipv6Ips []string `pulumi:"ipv6Ips"`
	// [integer] The LAN ID the NIC will sit on.
	Lan int `pulumi:"lan"`
	// Specifies the name of the flow log.
	//
	// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
	Name *string `pulumi:"name"`
	// [string] The ID of a server.
	ServerId string `pulumi:"serverId"`
}

// The set of arguments for constructing a Nic resource.
type NicArgs struct {
	// [string] The ID of a Virtual Data Center.
	DatacenterId pulumi.StringInput
	// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
	Dhcp pulumi.BoolPtrInput
	// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
	Dhcpv6 pulumi.BoolPtrInput
	// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
	FirewallActive pulumi.BoolPtrInput
	// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
	FirewallType pulumi.StringPtrInput
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlog NicFlowlogPtrInput
	// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
	Ips pulumi.StringArrayInput
	// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
	Ipv6CidrBlock pulumi.StringPtrInput
	// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
	Ipv6Ips pulumi.StringArrayInput
	// [integer] The LAN ID the NIC will sit on.
	Lan pulumi.IntInput
	// Specifies the name of the flow log.
	//
	// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
	Name pulumi.StringPtrInput
	// [string] The ID of a server.
	ServerId pulumi.StringInput
}

func (NicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nicArgs)(nil)).Elem()
}

type NicInput interface {
	pulumi.Input

	ToNicOutput() NicOutput
	ToNicOutputWithContext(ctx context.Context) NicOutput
}

func (*Nic) ElementType() reflect.Type {
	return reflect.TypeOf((**Nic)(nil)).Elem()
}

func (i *Nic) ToNicOutput() NicOutput {
	return i.ToNicOutputWithContext(context.Background())
}

func (i *Nic) ToNicOutputWithContext(ctx context.Context) NicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicOutput)
}

// NicArrayInput is an input type that accepts NicArray and NicArrayOutput values.
// You can construct a concrete instance of `NicArrayInput` via:
//
//	NicArray{ NicArgs{...} }
type NicArrayInput interface {
	pulumi.Input

	ToNicArrayOutput() NicArrayOutput
	ToNicArrayOutputWithContext(context.Context) NicArrayOutput
}

type NicArray []NicInput

func (NicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nic)(nil)).Elem()
}

func (i NicArray) ToNicArrayOutput() NicArrayOutput {
	return i.ToNicArrayOutputWithContext(context.Background())
}

func (i NicArray) ToNicArrayOutputWithContext(ctx context.Context) NicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicArrayOutput)
}

// NicMapInput is an input type that accepts NicMap and NicMapOutput values.
// You can construct a concrete instance of `NicMapInput` via:
//
//	NicMap{ "key": NicArgs{...} }
type NicMapInput interface {
	pulumi.Input

	ToNicMapOutput() NicMapOutput
	ToNicMapOutputWithContext(context.Context) NicMapOutput
}

type NicMap map[string]NicInput

func (NicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nic)(nil)).Elem()
}

func (i NicMap) ToNicMapOutput() NicMapOutput {
	return i.ToNicMapOutputWithContext(context.Background())
}

func (i NicMap) ToNicMapOutputWithContext(ctx context.Context) NicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NicMapOutput)
}

type NicOutput struct{ *pulumi.OutputState }

func (NicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Nic)(nil)).Elem()
}

func (o NicOutput) ToNicOutput() NicOutput {
	return o
}

func (o NicOutput) ToNicOutputWithContext(ctx context.Context) NicOutput {
	return o
}

// [string] The ID of a Virtual Data Center.
func (o NicOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.DatacenterId }).(pulumi.StringOutput)
}

// The Logical Unit Number (LUN) of the storage volume. Null if this NIC was created from CloudAPI and no DCD changes were done on the Datacenter.
func (o NicOutput) DeviceNumber() pulumi.IntOutput {
	return o.ApplyT(func(v *Nic) pulumi.IntOutput { return v.DeviceNumber }).(pulumi.IntOutput)
}

// [Boolean] Indicates if the NIC should get an IP address using DHCP (true) or not (false).
func (o NicOutput) Dhcp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nic) pulumi.BoolPtrOutput { return v.Dhcp }).(pulumi.BoolPtrOutput)
}

// [Boolean] Indicates if the NIC should get an IPv6 address using DHCP (true) or not (false).
func (o NicOutput) Dhcpv6() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nic) pulumi.BoolPtrOutput { return v.Dhcpv6 }).(pulumi.BoolPtrOutput)
}

// [Boolean] If this resource is set to true and is nested under a server resource firewall, with open SSH port, resource must be nested under the NIC.
func (o NicOutput) FirewallActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Nic) pulumi.BoolPtrOutput { return v.FirewallActive }).(pulumi.BoolPtrOutput)
}

// [String] The type of firewall rules that will be allowed on the NIC. If it is not specified it will take the default value INGRESS
func (o NicOutput) FirewallType() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.FirewallType }).(pulumi.StringOutput)
}

// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
func (o NicOutput) Flowlog() NicFlowlogPtrOutput {
	return o.ApplyT(func(v *Nic) NicFlowlogPtrOutput { return v.Flowlog }).(NicFlowlogPtrOutput)
}

// [list] Collection of IP addresses assigned to a NIC. Explicitly assigned public IPs need to come from reserved IP blocks, Passing value null or empty array will assign an IP address automatically.
func (o NicOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringArrayOutput { return v.Ips }).(pulumi.StringArrayOutput)
}

// Automatically assigned /80 IPv6 CIDR block if the NIC is connected to an IPv6 enabled LAN. You can also specify an /80 IPv6 CIDR block for the NIC on your own, which must be inside the /64 IPv6 CIDR block of the LAN and unique.
func (o NicOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// [list] Collection of IPv6 addresses assigned to a NIC. Explicitly assigned public IPs need to come from the NIC's Ipv6 CIDR block, Passing value null or empty array will assign an IPv6 address automatically from the NIC's CIDR block.
func (o NicOutput) Ipv6Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringArrayOutput { return v.Ipv6Ips }).(pulumi.StringArrayOutput)
}

// [integer] The LAN ID the NIC will sit on.
func (o NicOutput) Lan() pulumi.IntOutput {
	return o.ApplyT(func(v *Nic) pulumi.IntOutput { return v.Lan }).(pulumi.IntOutput)
}

// The MAC address of the NIC.
func (o NicOutput) Mac() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.Mac }).(pulumi.StringOutput)
}

// Specifies the name of the flow log.
//
// ⚠️ **Note:**: Removing the `flowlog` forces re-creation of the NIC resource.
func (o NicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The PCI slot number of the Nic.
func (o NicOutput) PciSlot() pulumi.IntOutput {
	return o.ApplyT(func(v *Nic) pulumi.IntOutput { return v.PciSlot }).(pulumi.IntOutput)
}

// [string] The ID of a server.
func (o NicOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v *Nic) pulumi.StringOutput { return v.ServerId }).(pulumi.StringOutput)
}

type NicArrayOutput struct{ *pulumi.OutputState }

func (NicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Nic)(nil)).Elem()
}

func (o NicArrayOutput) ToNicArrayOutput() NicArrayOutput {
	return o
}

func (o NicArrayOutput) ToNicArrayOutputWithContext(ctx context.Context) NicArrayOutput {
	return o
}

func (o NicArrayOutput) Index(i pulumi.IntInput) NicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Nic {
		return vs[0].([]*Nic)[vs[1].(int)]
	}).(NicOutput)
}

type NicMapOutput struct{ *pulumi.OutputState }

func (NicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Nic)(nil)).Elem()
}

func (o NicMapOutput) ToNicMapOutput() NicMapOutput {
	return o
}

func (o NicMapOutput) ToNicMapOutputWithContext(ctx context.Context) NicMapOutput {
	return o
}

func (o NicMapOutput) MapIndex(k pulumi.StringInput) NicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Nic {
		return vs[0].(map[string]*Nic)[vs[1].(string)]
	}).(NicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NicInput)(nil)).Elem(), &Nic{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicArrayInput)(nil)).Elem(), NicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NicMapInput)(nil)).Elem(), NicMap{})
	pulumi.RegisterOutputType(NicOutput{})
	pulumi.RegisterOutputType(NicArrayOutput{})
	pulumi.RegisterOutputType(NicMapOutput{})
}
