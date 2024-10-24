// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMariadbBackups(ctx *pulumi.Context, args *GetMariadbBackupsArgs, opts ...pulumi.InvokeOption) (*GetMariadbBackupsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetMariadbBackupsResult
	err := ctx.Invoke("ionoscloud:index/getMariadbBackups:getMariadbBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMariadbBackups.
type GetMariadbBackupsArgs struct {
	BackupId  *string `pulumi:"backupId"`
	ClusterId *string `pulumi:"clusterId"`
	Location  *string `pulumi:"location"`
}

// A collection of values returned by getMariadbBackups.
type GetMariadbBackupsResult struct {
	BackupId  *string                   `pulumi:"backupId"`
	Backups   []GetMariadbBackupsBackup `pulumi:"backups"`
	ClusterId *string                   `pulumi:"clusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id       string  `pulumi:"id"`
	Location *string `pulumi:"location"`
}

func GetMariadbBackupsOutput(ctx *pulumi.Context, args GetMariadbBackupsOutputArgs, opts ...pulumi.InvokeOption) GetMariadbBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMariadbBackupsResult, error) {
			args := v.(GetMariadbBackupsArgs)
			r, err := GetMariadbBackups(ctx, &args, opts...)
			var s GetMariadbBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMariadbBackupsResultOutput)
}

// A collection of arguments for invoking getMariadbBackups.
type GetMariadbBackupsOutputArgs struct {
	BackupId  pulumi.StringPtrInput `pulumi:"backupId"`
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	Location  pulumi.StringPtrInput `pulumi:"location"`
}

func (GetMariadbBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMariadbBackupsArgs)(nil)).Elem()
}

// A collection of values returned by getMariadbBackups.
type GetMariadbBackupsResultOutput struct{ *pulumi.OutputState }

func (GetMariadbBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMariadbBackupsResult)(nil)).Elem()
}

func (o GetMariadbBackupsResultOutput) ToGetMariadbBackupsResultOutput() GetMariadbBackupsResultOutput {
	return o
}

func (o GetMariadbBackupsResultOutput) ToGetMariadbBackupsResultOutputWithContext(ctx context.Context) GetMariadbBackupsResultOutput {
	return o
}

func (o GetMariadbBackupsResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMariadbBackupsResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o GetMariadbBackupsResultOutput) Backups() GetMariadbBackupsBackupArrayOutput {
	return o.ApplyT(func(v GetMariadbBackupsResult) []GetMariadbBackupsBackup { return v.Backups }).(GetMariadbBackupsBackupArrayOutput)
}

func (o GetMariadbBackupsResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMariadbBackupsResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetMariadbBackupsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetMariadbBackupsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetMariadbBackupsResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMariadbBackupsResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMariadbBackupsResultOutput{})
}