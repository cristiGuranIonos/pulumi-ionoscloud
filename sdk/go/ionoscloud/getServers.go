// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Servers data source** can be used to search for and return existing servers based on filters used.
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
//			_, err := ionoscloud.GetServers(ctx, &ionoscloud.GetServersArgs{
//				DatacenterId: ionoscloud_datacenter.Example.Id,
//				Filters: []ionoscloud.GetServersFilter{
//					{
//						Name:  "name",
//						Value: "server_name_to_look_here",
//					},
//				},
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
// ### By CPU Family
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
//			_, err := ionoscloud.GetServers(ctx, &ionoscloud.GetServersArgs{
//				DatacenterId: ionoscloud_datacenter.Example.Id,
//				Filters: []ionoscloud.GetServersFilter{
//					{
//						Name:  "cpu_family",
//						Value: "INTEL_XEON",
//					},
//				},
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
// ### By Name and Cores
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
//			_, err := ionoscloud.GetServers(ctx, &ionoscloud.GetServersArgs{
//				DatacenterId: ionoscloud_datacenter.Example.Id,
//				Filters: []ionoscloud.GetServersFilter{
//					{
//						Name:  "name",
//						Value: "test",
//					},
//					{
//						Name:  "cores",
//						Value: "1",
//					},
//				},
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
func GetServers(ctx *pulumi.Context, args *GetServersArgs, opts ...pulumi.InvokeOption) (*GetServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServersResult
	err := ctx.Invoke("ionoscloud:index/getServers:getServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServers.
type GetServersArgs struct {
	// Name of an existing datacenter that the servers are a part of
	DatacenterId string `pulumi:"datacenterId"`
	// One or more name/value pairs to filter off of. You can use most base fields in the server resource. These do **NOT** include nested fields in nics or volume nested fields.
	//
	// `datacenterId` must be provided. If `datacenterId` is missing , the datasource will return an error.
	//
	// **NOTE:** Lookup by filter is partial. Searching for a server using filter name and value `test`, will find all servers that have `test` in the name.
	// For example, it will find servers named `test`, `test1`, `testsomething`.
	//
	// **NOTE:** You cannot search by `imageName` by providing an alias like `ubuntu`.
	Filters []GetServersFilter `pulumi:"filters"`
}

// A collection of values returned by getServers.
type GetServersResult struct {
	DatacenterId string             `pulumi:"datacenterId"`
	Filters      []GetServersFilter `pulumi:"filters"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// list of servers that matches the filters provided.
	// For a full reference of all attributes returned, check out documentation
	Servers []GetServersServer `pulumi:"servers"`
}

func GetServersOutput(ctx *pulumi.Context, args GetServersOutputArgs, opts ...pulumi.InvokeOption) GetServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetServersResult, error) {
			args := v.(GetServersArgs)
			r, err := GetServers(ctx, &args, opts...)
			var s GetServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetServersResultOutput)
}

// A collection of arguments for invoking getServers.
type GetServersOutputArgs struct {
	// Name of an existing datacenter that the servers are a part of
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// One or more name/value pairs to filter off of. You can use most base fields in the server resource. These do **NOT** include nested fields in nics or volume nested fields.
	//
	// `datacenterId` must be provided. If `datacenterId` is missing , the datasource will return an error.
	//
	// **NOTE:** Lookup by filter is partial. Searching for a server using filter name and value `test`, will find all servers that have `test` in the name.
	// For example, it will find servers named `test`, `test1`, `testsomething`.
	//
	// **NOTE:** You cannot search by `imageName` by providing an alias like `ubuntu`.
	Filters GetServersFilterArrayInput `pulumi:"filters"`
}

func (GetServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServersArgs)(nil)).Elem()
}

// A collection of values returned by getServers.
type GetServersResultOutput struct{ *pulumi.OutputState }

func (GetServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServersResult)(nil)).Elem()
}

func (o GetServersResultOutput) ToGetServersResultOutput() GetServersResultOutput {
	return o
}

func (o GetServersResultOutput) ToGetServersResultOutputWithContext(ctx context.Context) GetServersResultOutput {
	return o
}

func (o GetServersResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServersResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

func (o GetServersResultOutput) Filters() GetServersFilterArrayOutput {
	return o.ApplyT(func(v GetServersResult) []GetServersFilter { return v.Filters }).(GetServersFilterArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServersResult) string { return v.Id }).(pulumi.StringOutput)
}

// list of servers that matches the filters provided.
// For a full reference of all attributes returned, check out documentation
func (o GetServersResultOutput) Servers() GetServersServerArrayOutput {
	return o.ApplyT(func(v GetServersResult) []GetServersServer { return v.Servers }).(GetServersServerArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServersResultOutput{})
}