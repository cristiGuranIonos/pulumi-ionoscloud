// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Network Load Balancer Forwarding Rule data source** can be used to search for and return existing network forwarding rules.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
func LookupNetworkloadbalancerForwardingrule(ctx *pulumi.Context, args *LookupNetworkloadbalancerForwardingruleArgs, opts ...pulumi.InvokeOption) (*LookupNetworkloadbalancerForwardingruleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkloadbalancerForwardingruleResult
	err := ctx.Invoke("ionoscloud:index/getNetworkloadbalancerForwardingrule:getNetworkloadbalancerForwardingrule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkloadbalancerForwardingrule.
type LookupNetworkloadbalancerForwardingruleArgs struct {
	// Datacenter's UUID.
	DatacenterId string `pulumi:"datacenterId"`
	// ID of the network load balancer forwarding rule you want to search for.
	//
	// Both `datacenterId` and `networkloadbalancerId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	Id *string `pulumi:"id"`
	// Name of an existing network load balancer forwarding rule that you want to search for.
	Name *string `pulumi:"name"`
	// Network Load Balancer's UUID.
	NetworkloadbalancerId string `pulumi:"networkloadbalancerId"`
}

// A collection of values returned by getNetworkloadbalancerForwardingrule.
type LookupNetworkloadbalancerForwardingruleResult struct {
	// Algorithm for the balancing.
	Algorithm    string `pulumi:"algorithm"`
	DatacenterId string `pulumi:"datacenterId"`
	// Health check attributes for Network Load Balancer forwarding rule target.
	HealthChecks []GetNetworkloadbalancerForwardingruleHealthCheck `pulumi:"healthChecks"`
	// The id of that Network Load Balancer forwarding rule.
	Id *string `pulumi:"id"`
	// Listening IP. (inbound)
	ListenerIp string `pulumi:"listenerIp"`
	// Listening port number. (inbound) (range: 1 to 65535)
	ListenerPort int `pulumi:"listenerPort"`
	// The name of that Network Load Balancer forwarding rule.
	Name                  *string `pulumi:"name"`
	NetworkloadbalancerId string  `pulumi:"networkloadbalancerId"`
	// Protocol of the balancing.
	Protocol string `pulumi:"protocol"`
	// Array of items in that collection.
	Targets []GetNetworkloadbalancerForwardingruleTarget `pulumi:"targets"`
}

func LookupNetworkloadbalancerForwardingruleOutput(ctx *pulumi.Context, args LookupNetworkloadbalancerForwardingruleOutputArgs, opts ...pulumi.InvokeOption) LookupNetworkloadbalancerForwardingruleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNetworkloadbalancerForwardingruleResult, error) {
			args := v.(LookupNetworkloadbalancerForwardingruleArgs)
			r, err := LookupNetworkloadbalancerForwardingrule(ctx, &args, opts...)
			var s LookupNetworkloadbalancerForwardingruleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNetworkloadbalancerForwardingruleResultOutput)
}

// A collection of arguments for invoking getNetworkloadbalancerForwardingrule.
type LookupNetworkloadbalancerForwardingruleOutputArgs struct {
	// Datacenter's UUID.
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// ID of the network load balancer forwarding rule you want to search for.
	//
	// Both `datacenterId` and `networkloadbalancerId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing network load balancer forwarding rule that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Network Load Balancer's UUID.
	NetworkloadbalancerId pulumi.StringInput `pulumi:"networkloadbalancerId"`
}

func (LookupNetworkloadbalancerForwardingruleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkloadbalancerForwardingruleArgs)(nil)).Elem()
}

// A collection of values returned by getNetworkloadbalancerForwardingrule.
type LookupNetworkloadbalancerForwardingruleResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkloadbalancerForwardingruleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkloadbalancerForwardingruleResult)(nil)).Elem()
}

func (o LookupNetworkloadbalancerForwardingruleResultOutput) ToLookupNetworkloadbalancerForwardingruleResultOutput() LookupNetworkloadbalancerForwardingruleResultOutput {
	return o
}

func (o LookupNetworkloadbalancerForwardingruleResultOutput) ToLookupNetworkloadbalancerForwardingruleResultOutputWithContext(ctx context.Context) LookupNetworkloadbalancerForwardingruleResultOutput {
	return o
}

// Algorithm for the balancing.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o LookupNetworkloadbalancerForwardingruleResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// Health check attributes for Network Load Balancer forwarding rule target.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) HealthChecks() GetNetworkloadbalancerForwardingruleHealthCheckArrayOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) []GetNetworkloadbalancerForwardingruleHealthCheck {
		return v.HealthChecks
	}).(GetNetworkloadbalancerForwardingruleHealthCheckArrayOutput)
}

// The id of that Network Load Balancer forwarding rule.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Listening IP. (inbound)
func (o LookupNetworkloadbalancerForwardingruleResultOutput) ListenerIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) string { return v.ListenerIp }).(pulumi.StringOutput)
}

// Listening port number. (inbound) (range: 1 to 65535)
func (o LookupNetworkloadbalancerForwardingruleResultOutput) ListenerPort() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) int { return v.ListenerPort }).(pulumi.IntOutput)
}

// The name of that Network Load Balancer forwarding rule.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkloadbalancerForwardingruleResultOutput) NetworkloadbalancerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) string { return v.NetworkloadbalancerId }).(pulumi.StringOutput)
}

// Protocol of the balancing.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) string { return v.Protocol }).(pulumi.StringOutput)
}

// Array of items in that collection.
func (o LookupNetworkloadbalancerForwardingruleResultOutput) Targets() GetNetworkloadbalancerForwardingruleTargetArrayOutput {
	return o.ApplyT(func(v LookupNetworkloadbalancerForwardingruleResult) []GetNetworkloadbalancerForwardingruleTarget {
		return v.Targets
	}).(GetNetworkloadbalancerForwardingruleTargetArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkloadbalancerForwardingruleResultOutput{})
}
