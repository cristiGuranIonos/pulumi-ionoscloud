// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Datacenter data source** can be used to search for and return an existing Virtual Data Center.
// You can provide a string for the name and location parameters which will be compared with provisioned Virtual Data Centers.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
//
// ## Example Usage
//
// ### By Name & Location
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
//			_, err := ionoscloud.LookupDatacenter(ctx, &ionoscloud.LookupDatacenterArgs{
//				Location: pulumi.StringRef("us/las"),
//				Name:     pulumi.StringRef("Datacenter Example"),
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
func LookupDatacenter(ctx *pulumi.Context, args *LookupDatacenterArgs, opts ...pulumi.InvokeOption) (*LookupDatacenterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDatacenterResult
	err := ctx.Invoke("ionoscloud:index/getDatacenter:getDatacenter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatacenter.
type LookupDatacenterArgs struct {
	// Id of an existing Virtual Data Center that you want to search for.
	Id *string `pulumi:"id"`
	// Id of the existing Virtual Data Center's location.
	//
	// Either `name`, `location` or `id` must be provided. If none, the datasource will return an error.
	Location *string `pulumi:"location"`
	// Name of an existing Virtual Data Center that you want to search for.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getDatacenter.
type LookupDatacenterResult struct {
	// Array of features and CPU families available in a location
	CpuArchitectures []GetDatacenterCpuArchitecture `pulumi:"cpuArchitectures"`
	// Description for the Virtual Data Center
	Description string `pulumi:"description"`
	// List of features supported by the location this data center is part of
	Features []string `pulumi:"features"`
	// UUID of the Virtual Data Center
	Id            *string `pulumi:"id"`
	Ipv6CidrBlock string  `pulumi:"ipv6CidrBlock"`
	// The regional location where the Virtual Data Center will be created
	Location *string `pulumi:"location"`
	// The name of the Virtual Data Center
	Name *string `pulumi:"name"`
	// Boolean value representing if the data center requires extra protection e.g. two factor protection
	SecAuthProtection bool `pulumi:"secAuthProtection"`
	// The version of that Data Center. Gets incremented with every change
	Version int `pulumi:"version"`
}

func LookupDatacenterOutput(ctx *pulumi.Context, args LookupDatacenterOutputArgs, opts ...pulumi.InvokeOption) LookupDatacenterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDatacenterResult, error) {
			args := v.(LookupDatacenterArgs)
			r, err := LookupDatacenter(ctx, &args, opts...)
			var s LookupDatacenterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDatacenterResultOutput)
}

// A collection of arguments for invoking getDatacenter.
type LookupDatacenterOutputArgs struct {
	// Id of an existing Virtual Data Center that you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Id of the existing Virtual Data Center's location.
	//
	// Either `name`, `location` or `id` must be provided. If none, the datasource will return an error.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// Name of an existing Virtual Data Center that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupDatacenterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterArgs)(nil)).Elem()
}

// A collection of values returned by getDatacenter.
type LookupDatacenterResultOutput struct{ *pulumi.OutputState }

func (LookupDatacenterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDatacenterResult)(nil)).Elem()
}

func (o LookupDatacenterResultOutput) ToLookupDatacenterResultOutput() LookupDatacenterResultOutput {
	return o
}

func (o LookupDatacenterResultOutput) ToLookupDatacenterResultOutputWithContext(ctx context.Context) LookupDatacenterResultOutput {
	return o
}

// Array of features and CPU families available in a location
func (o LookupDatacenterResultOutput) CpuArchitectures() GetDatacenterCpuArchitectureArrayOutput {
	return o.ApplyT(func(v LookupDatacenterResult) []GetDatacenterCpuArchitecture { return v.CpuArchitectures }).(GetDatacenterCpuArchitectureArrayOutput)
}

// Description for the Virtual Data Center
func (o LookupDatacenterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterResult) string { return v.Description }).(pulumi.StringOutput)
}

// List of features supported by the location this data center is part of
func (o LookupDatacenterResultOutput) Features() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDatacenterResult) []string { return v.Features }).(pulumi.StringArrayOutput)
}

// UUID of the Virtual Data Center
func (o LookupDatacenterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatacenterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupDatacenterResultOutput) Ipv6CidrBlock() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDatacenterResult) string { return v.Ipv6CidrBlock }).(pulumi.StringOutput)
}

// The regional location where the Virtual Data Center will be created
func (o LookupDatacenterResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatacenterResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the Virtual Data Center
func (o LookupDatacenterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDatacenterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Boolean value representing if the data center requires extra protection e.g. two factor protection
func (o LookupDatacenterResultOutput) SecAuthProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDatacenterResult) bool { return v.SecAuthProtection }).(pulumi.BoolOutput)
}

// The version of that Data Center. Gets incremented with every change
func (o LookupDatacenterResultOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDatacenterResult) int { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDatacenterResultOutput{})
}