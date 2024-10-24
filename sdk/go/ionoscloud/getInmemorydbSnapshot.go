// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInmemorydbSnapshot(ctx *pulumi.Context, args *GetInmemorydbSnapshotArgs, opts ...pulumi.InvokeOption) (*GetInmemorydbSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInmemorydbSnapshotResult
	err := ctx.Invoke("ionoscloud:index/getInmemorydbSnapshot:getInmemorydbSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInmemorydbSnapshot.
type GetInmemorydbSnapshotArgs struct {
	Id       string `pulumi:"id"`
	Location string `pulumi:"location"`
}

// A collection of values returned by getInmemorydbSnapshot.
type GetInmemorydbSnapshotResult struct {
	Id        string                          `pulumi:"id"`
	Location  string                          `pulumi:"location"`
	Metadatas []GetInmemorydbSnapshotMetadata `pulumi:"metadatas"`
}

func GetInmemorydbSnapshotOutput(ctx *pulumi.Context, args GetInmemorydbSnapshotOutputArgs, opts ...pulumi.InvokeOption) GetInmemorydbSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetInmemorydbSnapshotResult, error) {
			args := v.(GetInmemorydbSnapshotArgs)
			r, err := GetInmemorydbSnapshot(ctx, &args, opts...)
			var s GetInmemorydbSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetInmemorydbSnapshotResultOutput)
}

// A collection of arguments for invoking getInmemorydbSnapshot.
type GetInmemorydbSnapshotOutputArgs struct {
	Id       pulumi.StringInput `pulumi:"id"`
	Location pulumi.StringInput `pulumi:"location"`
}

func (GetInmemorydbSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInmemorydbSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getInmemorydbSnapshot.
type GetInmemorydbSnapshotResultOutput struct{ *pulumi.OutputState }

func (GetInmemorydbSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInmemorydbSnapshotResult)(nil)).Elem()
}

func (o GetInmemorydbSnapshotResultOutput) ToGetInmemorydbSnapshotResultOutput() GetInmemorydbSnapshotResultOutput {
	return o
}

func (o GetInmemorydbSnapshotResultOutput) ToGetInmemorydbSnapshotResultOutputWithContext(ctx context.Context) GetInmemorydbSnapshotResultOutput {
	return o
}

func (o GetInmemorydbSnapshotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInmemorydbSnapshotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInmemorydbSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v GetInmemorydbSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o GetInmemorydbSnapshotResultOutput) Metadatas() GetInmemorydbSnapshotMetadataArrayOutput {
	return o.ApplyT(func(v GetInmemorydbSnapshotResult) []GetInmemorydbSnapshotMetadata { return v.Metadatas }).(GetInmemorydbSnapshotMetadataArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInmemorydbSnapshotResultOutput{})
}
