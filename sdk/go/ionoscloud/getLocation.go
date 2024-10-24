// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Location data source** can be used to search for and return an existing location which can then be used elsewhere in the configuration.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
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
//			_, err := ionoscloud.GetLocation(ctx, &ionoscloud.GetLocationArgs{
//				Feature: pulumi.StringRef("SSD"),
//				Name:    pulumi.StringRef("karlsruhe"),
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
func GetLocation(ctx *pulumi.Context, args *GetLocationArgs, opts ...pulumi.InvokeOption) (*GetLocationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLocationResult
	err := ctx.Invoke("ionoscloud:index/getLocation:getLocation", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLocation.
type GetLocationArgs struct {
	// A desired feature that the location must be able to provide.
	Feature *string `pulumi:"feature"`
	// Name of the location to search for.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getLocation.
type GetLocationResult struct {
	// Array of features and CPU families available in a location
	CpuArchitectures []GetLocationCpuArchitecture `pulumi:"cpuArchitectures"`
	Feature          *string                      `pulumi:"feature"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of image aliases available for the location
	ImageAliases []string `pulumi:"imageAliases"`
	Name         *string  `pulumi:"name"`
}

func GetLocationOutput(ctx *pulumi.Context, args GetLocationOutputArgs, opts ...pulumi.InvokeOption) GetLocationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLocationResult, error) {
			args := v.(GetLocationArgs)
			r, err := GetLocation(ctx, &args, opts...)
			var s GetLocationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLocationResultOutput)
}

// A collection of arguments for invoking getLocation.
type GetLocationOutputArgs struct {
	// A desired feature that the location must be able to provide.
	Feature pulumi.StringPtrInput `pulumi:"feature"`
	// Name of the location to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (GetLocationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocationArgs)(nil)).Elem()
}

// A collection of values returned by getLocation.
type GetLocationResultOutput struct{ *pulumi.OutputState }

func (GetLocationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLocationResult)(nil)).Elem()
}

func (o GetLocationResultOutput) ToGetLocationResultOutput() GetLocationResultOutput {
	return o
}

func (o GetLocationResultOutput) ToGetLocationResultOutputWithContext(ctx context.Context) GetLocationResultOutput {
	return o
}

// Array of features and CPU families available in a location
func (o GetLocationResultOutput) CpuArchitectures() GetLocationCpuArchitectureArrayOutput {
	return o.ApplyT(func(v GetLocationResult) []GetLocationCpuArchitecture { return v.CpuArchitectures }).(GetLocationCpuArchitectureArrayOutput)
}

func (o GetLocationResultOutput) Feature() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocationResult) *string { return v.Feature }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetLocationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLocationResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of image aliases available for the location
func (o GetLocationResultOutput) ImageAliases() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetLocationResult) []string { return v.ImageAliases }).(pulumi.StringArrayOutput)
}

func (o GetLocationResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetLocationResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLocationResultOutput{})
}