// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Template data source** can be used to search for and return existing templates by providing any of template properties (name, cores, ram, storage_size).
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
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
//			_, err := ionoscloud.GetTemplate(ctx, &ionoscloud.GetTemplateArgs{
//				Name: pulumi.StringRef("CUBES S"),
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
// ### By Cores
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
//			_, err := ionoscloud.GetTemplate(ctx, &ionoscloud.GetTemplateArgs{
//				Cores: pulumi.Float64Ref(6),
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
// ### By Ram
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
//			_, err := ionoscloud.GetTemplate(ctx, &ionoscloud.GetTemplateArgs{
//				Ram: pulumi.Float64Ref(49152),
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
// ### By Storage Size
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
//			_, err := ionoscloud.GetTemplate(ctx, &ionoscloud.GetTemplateArgs{
//				StorageSize: pulumi.Float64Ref(80),
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
func GetTemplate(ctx *pulumi.Context, args *GetTemplateArgs, opts ...pulumi.InvokeOption) (*GetTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTemplateResult
	err := ctx.Invoke("ionoscloud:index/getTemplate:getTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTemplate.
type GetTemplateArgs struct {
	// The CPU cores count.
	Cores *float64 `pulumi:"cores"`
	// A name of that resource.
	Name *string `pulumi:"name"`
	// The RAM size in MB.
	Ram *float64 `pulumi:"ram"`
	// The storage size in GB.
	//
	// Any of the arguments ca be provided. If none, the datasource will return an error.
	StorageSize *float64 `pulumi:"storageSize"`
}

// A collection of values returned by getTemplate.
type GetTemplateResult struct {
	// The CPU cores count
	Cores float64 `pulumi:"cores"`
	// Id of template
	Id string `pulumi:"id"`
	// Name of template
	Name string `pulumi:"name"`
	// The RAM size in MB
	Ram float64 `pulumi:"ram"`
	// The storage size in GB
	StorageSize float64 `pulumi:"storageSize"`
}

func GetTemplateOutput(ctx *pulumi.Context, args GetTemplateOutputArgs, opts ...pulumi.InvokeOption) GetTemplateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetTemplateResult, error) {
			args := v.(GetTemplateArgs)
			r, err := GetTemplate(ctx, &args, opts...)
			var s GetTemplateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetTemplateResultOutput)
}

// A collection of arguments for invoking getTemplate.
type GetTemplateOutputArgs struct {
	// The CPU cores count.
	Cores pulumi.Float64PtrInput `pulumi:"cores"`
	// A name of that resource.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The RAM size in MB.
	Ram pulumi.Float64PtrInput `pulumi:"ram"`
	// The storage size in GB.
	//
	// Any of the arguments ca be provided. If none, the datasource will return an error.
	StorageSize pulumi.Float64PtrInput `pulumi:"storageSize"`
}

func (GetTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getTemplate.
type GetTemplateResultOutput struct{ *pulumi.OutputState }

func (GetTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTemplateResult)(nil)).Elem()
}

func (o GetTemplateResultOutput) ToGetTemplateResultOutput() GetTemplateResultOutput {
	return o
}

func (o GetTemplateResultOutput) ToGetTemplateResultOutputWithContext(ctx context.Context) GetTemplateResultOutput {
	return o
}

// The CPU cores count
func (o GetTemplateResultOutput) Cores() pulumi.Float64Output {
	return o.ApplyT(func(v GetTemplateResult) float64 { return v.Cores }).(pulumi.Float64Output)
}

// Id of template
func (o GetTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of template
func (o GetTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

// The RAM size in MB
func (o GetTemplateResultOutput) Ram() pulumi.Float64Output {
	return o.ApplyT(func(v GetTemplateResult) float64 { return v.Ram }).(pulumi.Float64Output)
}

// The storage size in GB
func (o GetTemplateResultOutput) StorageSize() pulumi.Float64Output {
	return o.ApplyT(func(v GetTemplateResult) float64 { return v.StorageSize }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(GetTemplateResultOutput{})
}