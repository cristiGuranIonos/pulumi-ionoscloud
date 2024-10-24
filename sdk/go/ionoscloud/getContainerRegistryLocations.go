// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Container Registry Locations data source** can be used to get a list of Container Registry Locations
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
//			_, err := ionoscloud.GetContainerRegistryLocations(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func GetContainerRegistryLocations(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetContainerRegistryLocationsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetContainerRegistryLocationsResult
	err := ctx.Invoke("ionoscloud:index/getContainerRegistryLocations:getContainerRegistryLocations", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getContainerRegistryLocations.
type GetContainerRegistryLocationsResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of container registry locations
	Locations []string `pulumi:"locations"`
}

func GetContainerRegistryLocationsOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetContainerRegistryLocationsResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetContainerRegistryLocationsResult, error) {
		r, err := GetContainerRegistryLocations(ctx, opts...)
		var s GetContainerRegistryLocationsResult
		if r != nil {
			s = *r
		}
		return s, err
	}).(GetContainerRegistryLocationsResultOutput)
}

// A collection of values returned by getContainerRegistryLocations.
type GetContainerRegistryLocationsResultOutput struct{ *pulumi.OutputState }

func (GetContainerRegistryLocationsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetContainerRegistryLocationsResult)(nil)).Elem()
}

func (o GetContainerRegistryLocationsResultOutput) ToGetContainerRegistryLocationsResultOutput() GetContainerRegistryLocationsResultOutput {
	return o
}

func (o GetContainerRegistryLocationsResultOutput) ToGetContainerRegistryLocationsResultOutputWithContext(ctx context.Context) GetContainerRegistryLocationsResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetContainerRegistryLocationsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetContainerRegistryLocationsResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of container registry locations
func (o GetContainerRegistryLocationsResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetContainerRegistryLocationsResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetContainerRegistryLocationsResultOutput{})
}
