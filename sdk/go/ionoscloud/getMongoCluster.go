// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMongoCluster(ctx *pulumi.Context, args *LookupMongoClusterArgs, opts ...pulumi.InvokeOption) (*LookupMongoClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupMongoClusterResult
	err := ctx.Invoke("ionoscloud:index/getMongoCluster:getMongoCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getMongoCluster.
type LookupMongoClusterArgs struct {
	DisplayName *string `pulumi:"displayName"`
	Id          *string `pulumi:"id"`
}

// A collection of values returned by getMongoCluster.
type LookupMongoClusterResult struct {
	Backups            []GetMongoClusterBackup            `pulumi:"backups"`
	BiConnectors       []GetMongoClusterBiConnector       `pulumi:"biConnectors"`
	ConnectionString   string                             `pulumi:"connectionString"`
	Connections        []GetMongoClusterConnection        `pulumi:"connections"`
	Cores              int                                `pulumi:"cores"`
	DisplayName        *string                            `pulumi:"displayName"`
	Edition            string                             `pulumi:"edition"`
	Id                 *string                            `pulumi:"id"`
	Instances          int                                `pulumi:"instances"`
	Location           string                             `pulumi:"location"`
	MaintenanceWindows []GetMongoClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	MongodbVersion     string                             `pulumi:"mongodbVersion"`
	Ram                int                                `pulumi:"ram"`
	Shards             int                                `pulumi:"shards"`
	StorageSize        int                                `pulumi:"storageSize"`
	StorageType        string                             `pulumi:"storageType"`
	TemplateId         string                             `pulumi:"templateId"`
	Type               string                             `pulumi:"type"`
}

func LookupMongoClusterOutput(ctx *pulumi.Context, args LookupMongoClusterOutputArgs, opts ...pulumi.InvokeOption) LookupMongoClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMongoClusterResult, error) {
			args := v.(LookupMongoClusterArgs)
			r, err := LookupMongoCluster(ctx, &args, opts...)
			var s LookupMongoClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMongoClusterResultOutput)
}

// A collection of arguments for invoking getMongoCluster.
type LookupMongoClusterOutputArgs struct {
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	Id          pulumi.StringPtrInput `pulumi:"id"`
}

func (LookupMongoClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterArgs)(nil)).Elem()
}

// A collection of values returned by getMongoCluster.
type LookupMongoClusterResultOutput struct{ *pulumi.OutputState }

func (LookupMongoClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMongoClusterResult)(nil)).Elem()
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutput() LookupMongoClusterResultOutput {
	return o
}

func (o LookupMongoClusterResultOutput) ToLookupMongoClusterResultOutputWithContext(ctx context.Context) LookupMongoClusterResultOutput {
	return o
}

func (o LookupMongoClusterResultOutput) Backups() GetMongoClusterBackupArrayOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) []GetMongoClusterBackup { return v.Backups }).(GetMongoClusterBackupArrayOutput)
}

func (o LookupMongoClusterResultOutput) BiConnectors() GetMongoClusterBiConnectorArrayOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) []GetMongoClusterBiConnector { return v.BiConnectors }).(GetMongoClusterBiConnectorArrayOutput)
}

func (o LookupMongoClusterResultOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) Connections() GetMongoClusterConnectionArrayOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) []GetMongoClusterConnection { return v.Connections }).(GetMongoClusterConnectionArrayOutput)
}

func (o LookupMongoClusterResultOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) int { return v.Cores }).(pulumi.IntOutput)
}

func (o LookupMongoClusterResultOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o LookupMongoClusterResultOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Edition }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupMongoClusterResultOutput) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) int { return v.Instances }).(pulumi.IntOutput)
}

func (o LookupMongoClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) MaintenanceWindows() GetMongoClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) []GetMongoClusterMaintenanceWindow { return v.MaintenanceWindows }).(GetMongoClusterMaintenanceWindowArrayOutput)
}

func (o LookupMongoClusterResultOutput) MongodbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.MongodbVersion }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) int { return v.Ram }).(pulumi.IntOutput)
}

func (o LookupMongoClusterResultOutput) Shards() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) int { return v.Shards }).(pulumi.IntOutput)
}

func (o LookupMongoClusterResultOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) int { return v.StorageSize }).(pulumi.IntOutput)
}

func (o LookupMongoClusterResultOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.StorageType }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.TemplateId }).(pulumi.StringOutput)
}

func (o LookupMongoClusterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMongoClusterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMongoClusterResultOutput{})
}
