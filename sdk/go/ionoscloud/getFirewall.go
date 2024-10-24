// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Firewall data source** can be used to search for and return an existing FirewallRules.
// You can provide a string for either id or name parameters which will be compared with provisioned Firewall Rules.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
//
// ## Example Usage
func LookupFirewall(ctx *pulumi.Context, args *LookupFirewallArgs, opts ...pulumi.InvokeOption) (*LookupFirewallResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupFirewallResult
	err := ctx.Invoke("ionoscloud:index/getFirewall:getFirewall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFirewall.
type LookupFirewallArgs struct {
	// The Virtual Data Center ID.
	DatacenterId string `pulumi:"datacenterId"`
	// ID of the firewall rule you want to search for.
	Id *string `pulumi:"id"`
	// Name of an existing firewall rule that you want to search for.
	Name *string `pulumi:"name"`
	// The NIC ID.
	//
	// Either `name` or   `id` must be provided. If none, or both are provided, the datasource will return an error.
	NicId string `pulumi:"nicId"`
	// The Server ID.
	ServerId string `pulumi:"serverId"`
}

// A collection of values returned by getFirewall.
type LookupFirewallResult struct {
	DatacenterId string `pulumi:"datacenterId"`
	// Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
	IcmpCode string `pulumi:"icmpCode"`
	// Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen.
	IcmpType string `pulumi:"icmpType"`
	// The id of the firewall rule.
	Id *string `pulumi:"id"`
	// The name of the firewall rule.
	Name  *string `pulumi:"name"`
	NicId string  `pulumi:"nicId"`
	// Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen.
	PortRangeEnd int `pulumi:"portRangeEnd"`
	// Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen.
	PortRangeStart int `pulumi:"portRangeStart"`
	// The protocol for the rule: TCP, UDP, ICMP, ANY. This property is immutable.
	Protocol string `pulumi:"protocol"`
	ServerId string `pulumi:"serverId"`
	// Only traffic originating from the respective IPv4 address is allowed.
	SourceIp string `pulumi:"sourceIp"`
	// Only traffic originating from the respective MAC address is allowed.
	SourceMac string `pulumi:"sourceMac"`
	// Only traffic directed to the respective IP address of the NIC is allowed.
	TargetIp string `pulumi:"targetIp"`
	Type     string `pulumi:"type"`
}

func LookupFirewallOutput(ctx *pulumi.Context, args LookupFirewallOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallResult, error) {
			args := v.(LookupFirewallArgs)
			r, err := LookupFirewall(ctx, &args, opts...)
			var s LookupFirewallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallResultOutput)
}

// A collection of arguments for invoking getFirewall.
type LookupFirewallOutputArgs struct {
	// The Virtual Data Center ID.
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// ID of the firewall rule you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing firewall rule that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The NIC ID.
	//
	// Either `name` or   `id` must be provided. If none, or both are provided, the datasource will return an error.
	NicId pulumi.StringInput `pulumi:"nicId"`
	// The Server ID.
	ServerId pulumi.StringInput `pulumi:"serverId"`
}

func (LookupFirewallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallArgs)(nil)).Elem()
}

// A collection of values returned by getFirewall.
type LookupFirewallResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallResult)(nil)).Elem()
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutput() LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) ToLookupFirewallResultOutputWithContext(ctx context.Context) LookupFirewallResultOutput {
	return o
}

func (o LookupFirewallResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// Defines the allowed code (from 0 to 254) if protocol ICMP is chosen.
func (o LookupFirewallResultOutput) IcmpCode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.IcmpCode }).(pulumi.StringOutput)
}

// Defines the allowed type (from 0 to 254) if the protocol ICMP is chosen.
func (o LookupFirewallResultOutput) IcmpType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.IcmpType }).(pulumi.StringOutput)
}

// The id of the firewall rule.
func (o LookupFirewallResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the firewall rule.
func (o LookupFirewallResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallResultOutput) NicId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.NicId }).(pulumi.StringOutput)
}

// Defines the end range of the allowed port (from 1 to 65534) if the protocol TCP or UDP is chosen.
func (o LookupFirewallResultOutput) PortRangeEnd() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallResult) int { return v.PortRangeEnd }).(pulumi.IntOutput)
}

// Defines the start range of the allowed port (from 1 to 65534) if protocol TCP or UDP is chosen.
func (o LookupFirewallResultOutput) PortRangeStart() pulumi.IntOutput {
	return o.ApplyT(func(v LookupFirewallResult) int { return v.PortRangeStart }).(pulumi.IntOutput)
}

// The protocol for the rule: TCP, UDP, ICMP, ANY. This property is immutable.
func (o LookupFirewallResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupFirewallResultOutput) ServerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.ServerId }).(pulumi.StringOutput)
}

// Only traffic originating from the respective IPv4 address is allowed.
func (o LookupFirewallResultOutput) SourceIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.SourceIp }).(pulumi.StringOutput)
}

// Only traffic originating from the respective MAC address is allowed.
func (o LookupFirewallResultOutput) SourceMac() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.SourceMac }).(pulumi.StringOutput)
}

// Only traffic directed to the respective IP address of the NIC is allowed.
func (o LookupFirewallResultOutput) TargetIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.TargetIp }).(pulumi.StringOutput)
}

func (o LookupFirewallResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallResultOutput{})
}
