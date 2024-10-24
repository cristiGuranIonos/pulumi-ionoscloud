// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Application Load Balancer data source** can be used to search for and return an existing Application Load Balancer.
// You can provide a string for the name parameter which will be compared with provisioned Application Load Balancers.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search and make sure that your resources have unique names.
//
// ## Example Usage
//
// ### By Name
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
//			_, err := ionoscloud.LookupApplicationLoadbalancer(ctx, &ionoscloud.LookupApplicationLoadbalancerArgs{
//				DatacenterId: ionoscloud_datacenter.Example.Id,
//				Name:         pulumi.StringRef("ALB name"),
//			}, nil)
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
// ### By Name with Partial Match
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
//			_, err := ionoscloud.LookupApplicationLoadbalancer(ctx, &ionoscloud.LookupApplicationLoadbalancerArgs{
//				DatacenterId: ionoscloud_datacenter.Example.Id,
//				Name:         pulumi.StringRef("name"),
//				PartialMatch: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupApplicationLoadbalancer(ctx *pulumi.Context, args *LookupApplicationLoadbalancerArgs, opts ...pulumi.InvokeOption) (*LookupApplicationLoadbalancerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationLoadbalancerResult
	err := ctx.Invoke("ionoscloud:index/getApplicationLoadbalancer:getApplicationLoadbalancer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getApplicationLoadbalancer.
type LookupApplicationLoadbalancerArgs struct {
	// Datacenter's UUID.
	DatacenterId string `pulumi:"datacenterId"`
	// ID of the application load balancer you want to search for.
	Id *string `pulumi:"id"`
	// Name of an existing application load balancer that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name *string `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// `datacenterId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	PartialMatch *bool `pulumi:"partialMatch"`
}

// A collection of values returned by getApplicationLoadbalancer.
type LookupApplicationLoadbalancerResult struct {
	// Turn logging on and off for this product. Default value is 'false'.
	CentralLogging bool   `pulumi:"centralLogging"`
	DatacenterId   string `pulumi:"datacenterId"`
	// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
	Flowlogs []GetApplicationLoadbalancerFlowlog `pulumi:"flowlogs"`
	// Id of the application load balancer.
	Id *string `pulumi:"id"`
	// Collection of the Application Load Balancer IP addresses. (Inbound and outbound) IPs of the listenerLan are customer-reserved public IPs for the public Load Balancers, and private IPs for the private Load Balancers.
	Ips []string `pulumi:"ips"`
	// Collection of private IP addresses with the subnet mask of the Application Load Balancer. IPs must contain valid a subnet mask. If no IP is provided, the system will generate an IP with /24 subnet.
	LbPrivateIps []string `pulumi:"lbPrivateIps"`
	// ID of the listening (inbound) LAN.
	ListenerLan   int    `pulumi:"listenerLan"`
	LoggingFormat string `pulumi:"loggingFormat"`
	// Specifies the name of the flow log.
	Name         *string `pulumi:"name"`
	PartialMatch *bool   `pulumi:"partialMatch"`
	// ID of the balanced private target LAN (outbound).
	TargetLan int `pulumi:"targetLan"`
}

func LookupApplicationLoadbalancerOutput(ctx *pulumi.Context, args LookupApplicationLoadbalancerOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationLoadbalancerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationLoadbalancerResult, error) {
			args := v.(LookupApplicationLoadbalancerArgs)
			r, err := LookupApplicationLoadbalancer(ctx, &args, opts...)
			var s LookupApplicationLoadbalancerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationLoadbalancerResultOutput)
}

// A collection of arguments for invoking getApplicationLoadbalancer.
type LookupApplicationLoadbalancerOutputArgs struct {
	// Datacenter's UUID.
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// ID of the application load balancer you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing application load balancer that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// `datacenterId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	PartialMatch pulumi.BoolPtrInput `pulumi:"partialMatch"`
}

func (LookupApplicationLoadbalancerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationLoadbalancerArgs)(nil)).Elem()
}

// A collection of values returned by getApplicationLoadbalancer.
type LookupApplicationLoadbalancerResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationLoadbalancerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationLoadbalancerResult)(nil)).Elem()
}

func (o LookupApplicationLoadbalancerResultOutput) ToLookupApplicationLoadbalancerResultOutput() LookupApplicationLoadbalancerResultOutput {
	return o
}

func (o LookupApplicationLoadbalancerResultOutput) ToLookupApplicationLoadbalancerResultOutputWithContext(ctx context.Context) LookupApplicationLoadbalancerResultOutput {
	return o
}

// Turn logging on and off for this product. Default value is 'false'.
func (o LookupApplicationLoadbalancerResultOutput) CentralLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) bool { return v.CentralLogging }).(pulumi.BoolOutput)
}

func (o LookupApplicationLoadbalancerResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// Only 1 flow log can be configured. Only the name field can change as part of an update. Flow logs holistically capture network information such as source and destination IP addresses, source and destination ports, number of packets, amount of bytes, the start and end time of the recording, and the type of protocol – and log the extent to which your instances are being accessed.
func (o LookupApplicationLoadbalancerResultOutput) Flowlogs() GetApplicationLoadbalancerFlowlogArrayOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) []GetApplicationLoadbalancerFlowlog { return v.Flowlogs }).(GetApplicationLoadbalancerFlowlogArrayOutput)
}

// Id of the application load balancer.
func (o LookupApplicationLoadbalancerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Collection of the Application Load Balancer IP addresses. (Inbound and outbound) IPs of the listenerLan are customer-reserved public IPs for the public Load Balancers, and private IPs for the private Load Balancers.
func (o LookupApplicationLoadbalancerResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// Collection of private IP addresses with the subnet mask of the Application Load Balancer. IPs must contain valid a subnet mask. If no IP is provided, the system will generate an IP with /24 subnet.
func (o LookupApplicationLoadbalancerResultOutput) LbPrivateIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) []string { return v.LbPrivateIps }).(pulumi.StringArrayOutput)
}

// ID of the listening (inbound) LAN.
func (o LookupApplicationLoadbalancerResultOutput) ListenerLan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) int { return v.ListenerLan }).(pulumi.IntOutput)
}

func (o LookupApplicationLoadbalancerResultOutput) LoggingFormat() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) string { return v.LoggingFormat }).(pulumi.StringOutput)
}

// Specifies the name of the flow log.
func (o LookupApplicationLoadbalancerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupApplicationLoadbalancerResultOutput) PartialMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) *bool { return v.PartialMatch }).(pulumi.BoolPtrOutput)
}

// ID of the balanced private target LAN (outbound).
func (o LookupApplicationLoadbalancerResultOutput) TargetLan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupApplicationLoadbalancerResult) int { return v.TargetLan }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationLoadbalancerResultOutput{})
}