// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **DNS Zone** can be used to search for and return an existing DNS Zone.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search and make sure that your resources have unique names.
//
// > ⚠️  Only tokens are accepted for authorization in the **ionoscloud_dns_zone** data source. Please ensure you are using tokens as other methods will not be valid.
//
// ## Example Usage
//
// ### By name
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
//			_, err := ionoscloud.LookupDnsZone(ctx, &ionoscloud.LookupDnsZoneArgs{
//				Name: pulumi.StringRef("example.com"),
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
// ### By name with partial match
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
//			_, err := ionoscloud.LookupDnsZone(ctx, &ionoscloud.LookupDnsZoneArgs{
//				Name:         pulumi.StringRef("example"),
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
func LookupDnsZone(ctx *pulumi.Context, args *LookupDnsZoneArgs, opts ...pulumi.InvokeOption) (*LookupDnsZoneResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDnsZoneResult
	err := ctx.Invoke("ionoscloud:index/getDnsZone:getDnsZone", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneArgs struct {
	// [string] The ID of the DNS Zone you want to search for.
	Id *string `pulumi:"id"`
	// [string] The name of the DNS Zone you want to search for.
	Name *string `pulumi:"name"`
	// [bool] Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch *bool `pulumi:"partialMatch"`
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResult struct {
	// The description of the DNS Zone.
	Description string `pulumi:"description"`
	// Indicates if the DNS Zone is activated or not.
	Enabled bool `pulumi:"enabled"`
	// The UUID of the DNS Zone.
	Id *string `pulumi:"id"`
	// The name of the DNS Zone.
	Name *string `pulumi:"name"`
	// A list of available name servers.
	Nameservers  []string `pulumi:"nameservers"`
	PartialMatch *bool    `pulumi:"partialMatch"`
}

func LookupDnsZoneOutput(ctx *pulumi.Context, args LookupDnsZoneOutputArgs, opts ...pulumi.InvokeOption) LookupDnsZoneResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDnsZoneResult, error) {
			args := v.(LookupDnsZoneArgs)
			r, err := LookupDnsZone(ctx, &args, opts...)
			var s LookupDnsZoneResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDnsZoneResultOutput)
}

// A collection of arguments for invoking getDnsZone.
type LookupDnsZoneOutputArgs struct {
	// [string] The ID of the DNS Zone you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// [string] The name of the DNS Zone you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// [bool] Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch pulumi.BoolPtrInput `pulumi:"partialMatch"`
}

func (LookupDnsZoneOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneArgs)(nil)).Elem()
}

// A collection of values returned by getDnsZone.
type LookupDnsZoneResultOutput struct{ *pulumi.OutputState }

func (LookupDnsZoneResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDnsZoneResult)(nil)).Elem()
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutput() LookupDnsZoneResultOutput {
	return o
}

func (o LookupDnsZoneResultOutput) ToLookupDnsZoneResultOutputWithContext(ctx context.Context) LookupDnsZoneResultOutput {
	return o
}

// The description of the DNS Zone.
func (o LookupDnsZoneResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) string { return v.Description }).(pulumi.StringOutput)
}

// Indicates if the DNS Zone is activated or not.
func (o LookupDnsZoneResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The UUID of the DNS Zone.
func (o LookupDnsZoneResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the DNS Zone.
func (o LookupDnsZoneResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// A list of available name servers.
func (o LookupDnsZoneResultOutput) Nameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) []string { return v.Nameservers }).(pulumi.StringArrayOutput)
}

func (o LookupDnsZoneResultOutput) PartialMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDnsZoneResult) *bool { return v.PartialMatch }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDnsZoneResultOutput{})
}
