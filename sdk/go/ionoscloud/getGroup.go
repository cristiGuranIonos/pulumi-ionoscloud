// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Group data source** can be used to search for and return existing groups.
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
//			_, err := ionoscloud.LookupGroup(ctx, &ionoscloud.LookupGroupArgs{
//				Name: pulumi.StringRef("Group Example"),
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
func LookupGroup(ctx *pulumi.Context, args *LookupGroupArgs, opts ...pulumi.InvokeOption) (*LookupGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupGroupResult
	err := ctx.Invoke("ionoscloud:index/getGroup:getGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGroup.
type LookupGroupArgs struct {
	// ID of the group you want to search for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	Id *string `pulumi:"id"`
	// Name of an existing group that you want to search for.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getGroup.
type LookupGroupResult struct {
	// The group will be allowed to access the activity log.
	AccessActivityLog bool `pulumi:"accessActivityLog"`
	// The group will be allowed to access and manage certificates.
	AccessAndManageCertificates bool `pulumi:"accessAndManageCertificates"`
	// The group will be allowed to access and manage monitoring.
	AccessAndManageMonitoring bool `pulumi:"accessAndManageMonitoring"`
	// The group will be allowed to create backup unit privilege.
	CreateBackupUnit bool `pulumi:"createBackupUnit"`
	// The group will be allowed to create virtual data centers.
	CreateDatacenter bool `pulumi:"createDatacenter"`
	// The group will be allowed to create flow log.
	CreateFlowLog bool `pulumi:"createFlowLog"`
	// The group will be allowed to create internet access privilege.
	CreateInternetAccess bool `pulumi:"createInternetAccess"`
	// The group will be allowed to create kubernetes cluster privilege.
	CreateK8sCluster bool `pulumi:"createK8sCluster"`
	// The group will be allowed to create Cross Connects privilege.
	CreatePcc bool `pulumi:"createPcc"`
	// The group will be allowed to create snapshots.
	CreateSnapshot bool `pulumi:"createSnapshot"`
	// The id of the group.
	Id *string `pulumi:"id"`
	// Privilege for a group to manage DBaaS related functionality.
	ManageDbaas bool `pulumi:"manageDbaas"`
	// A name for the group.
	Name *string `pulumi:"name"`
	// The group will be allowed to reserve IP addresses.
	ReserveIp bool `pulumi:"reserveIp"`
	// The group will have S3 privilege.
	S3Privilege bool `pulumi:"s3Privilege"`
	// List of users in group.
	Users []GetGroupUser `pulumi:"users"`
}

func LookupGroupOutput(ctx *pulumi.Context, args LookupGroupOutputArgs, opts ...pulumi.InvokeOption) LookupGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGroupResult, error) {
			args := v.(LookupGroupArgs)
			r, err := LookupGroup(ctx, &args, opts...)
			var s LookupGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGroupResultOutput)
}

// A collection of arguments for invoking getGroup.
type LookupGroupOutputArgs struct {
	// ID of the group you want to search for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing group that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupArgs)(nil)).Elem()
}

// A collection of values returned by getGroup.
type LookupGroupResultOutput struct{ *pulumi.OutputState }

func (LookupGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGroupResult)(nil)).Elem()
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutput() LookupGroupResultOutput {
	return o
}

func (o LookupGroupResultOutput) ToLookupGroupResultOutputWithContext(ctx context.Context) LookupGroupResultOutput {
	return o
}

// The group will be allowed to access the activity log.
func (o LookupGroupResultOutput) AccessActivityLog() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AccessActivityLog }).(pulumi.BoolOutput)
}

// The group will be allowed to access and manage certificates.
func (o LookupGroupResultOutput) AccessAndManageCertificates() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AccessAndManageCertificates }).(pulumi.BoolOutput)
}

// The group will be allowed to access and manage monitoring.
func (o LookupGroupResultOutput) AccessAndManageMonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.AccessAndManageMonitoring }).(pulumi.BoolOutput)
}

// The group will be allowed to create backup unit privilege.
func (o LookupGroupResultOutput) CreateBackupUnit() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateBackupUnit }).(pulumi.BoolOutput)
}

// The group will be allowed to create virtual data centers.
func (o LookupGroupResultOutput) CreateDatacenter() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateDatacenter }).(pulumi.BoolOutput)
}

// The group will be allowed to create flow log.
func (o LookupGroupResultOutput) CreateFlowLog() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateFlowLog }).(pulumi.BoolOutput)
}

// The group will be allowed to create internet access privilege.
func (o LookupGroupResultOutput) CreateInternetAccess() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateInternetAccess }).(pulumi.BoolOutput)
}

// The group will be allowed to create kubernetes cluster privilege.
func (o LookupGroupResultOutput) CreateK8sCluster() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateK8sCluster }).(pulumi.BoolOutput)
}

// The group will be allowed to create Cross Connects privilege.
func (o LookupGroupResultOutput) CreatePcc() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreatePcc }).(pulumi.BoolOutput)
}

// The group will be allowed to create snapshots.
func (o LookupGroupResultOutput) CreateSnapshot() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.CreateSnapshot }).(pulumi.BoolOutput)
}

// The id of the group.
func (o LookupGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Privilege for a group to manage DBaaS related functionality.
func (o LookupGroupResultOutput) ManageDbaas() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.ManageDbaas }).(pulumi.BoolOutput)
}

// A name for the group.
func (o LookupGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The group will be allowed to reserve IP addresses.
func (o LookupGroupResultOutput) ReserveIp() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.ReserveIp }).(pulumi.BoolOutput)
}

// The group will have S3 privilege.
func (o LookupGroupResultOutput) S3Privilege() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupGroupResult) bool { return v.S3Privilege }).(pulumi.BoolOutput)
}

// List of users in group.
func (o LookupGroupResultOutput) Users() GetGroupUserArrayOutput {
	return o.ApplyT(func(v LookupGroupResult) []GetGroupUser { return v.Users }).(GetGroupUserArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGroupResultOutput{})
}