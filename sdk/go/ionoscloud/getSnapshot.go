// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Snapshot data source** can be used to search for and return an existing snapshot which can then be used to provision a server. If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned. When this happens, please refine your search string so that it is specific enough to return only one result.
//
// ## Example Usage
//
// ### By Name & Size & Location
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
//			_, err := ionoscloud.LookupSnapshot(ctx, &ionoscloud.LookupSnapshotArgs{
//				Location: pulumi.StringRef("us/las"),
//				Name:     pulumi.StringRef("Snapshot Example"),
//				Size:     pulumi.IntRef(2),
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
// Note: The size argument is in GB
func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupSnapshotResult
	err := ctx.Invoke("ionoscloud:index/getSnapshot:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotArgs struct {
	// UUID of an existing snapshot that you want to search for.
	Id *string `pulumi:"id"`
	// Existing snapshot's location.
	Location *string `pulumi:"location"`
	// Name of an existing snapshot that you want to search for.
	Name *string `pulumi:"name"`
	// The size of the snapshot to look for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	// Additionally, you can add `location` and `size` along with the `name` argument for a more refined search.
	// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
	// When this happens, please refine your search string so that it is specific enough to return only one result.
	Size *int `pulumi:"size"`
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResult struct {
	// Is capable of CPU hot plug (no reboot required)
	CpuHotPlug bool `pulumi:"cpuHotPlug"`
	// Is capable of CPU hot unplug (no reboot required)
	CpuHotUnplug bool `pulumi:"cpuHotUnplug"`
	// Human readable description
	Description string `pulumi:"description"`
	// Is capable of SCSI drive hot plug (no reboot required)
	DiscScsiHotPlug bool `pulumi:"discScsiHotPlug"`
	// Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscScsiHotUnplug bool `pulumi:"discScsiHotUnplug"`
	// Is capable of Virt-IO drive hot plug (no reboot required)
	DiscVirtioHotPlug bool `pulumi:"discVirtioHotPlug"`
	// Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
	DiscVirtioHotUnplug bool `pulumi:"discVirtioHotUnplug"`
	// UUID of the snapshot
	Id *string `pulumi:"id"`
	// OS type of this Snapshot
	LicenceType string `pulumi:"licenceType"`
	// Location of that image/snapshot
	Location string `pulumi:"location"`
	// The name of the snapshot.
	Name string `pulumi:"name"`
	// Is capable of nic hot plug (no reboot required)
	NicHotPlug bool `pulumi:"nicHotPlug"`
	// Is capable of nic hot unplug (no reboot required)
	NicHotUnplug bool `pulumi:"nicHotUnplug"`
	// Is capable of memory hot plug (no reboot required)
	RamHotPlug bool `pulumi:"ramHotPlug"`
	// Is capable of memory hot unplug (no reboot required)
	RamHotUnplug bool `pulumi:"ramHotUnplug"`
	// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
	SecAuthProtection bool `pulumi:"secAuthProtection"`
	// The size of the image in GB
	Size int `pulumi:"size"`
}

func LookupSnapshotOutput(ctx *pulumi.Context, args LookupSnapshotOutputArgs, opts ...pulumi.InvokeOption) LookupSnapshotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSnapshotResult, error) {
			args := v.(LookupSnapshotArgs)
			r, err := LookupSnapshot(ctx, &args, opts...)
			var s LookupSnapshotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSnapshotResultOutput)
}

// A collection of arguments for invoking getSnapshot.
type LookupSnapshotOutputArgs struct {
	// UUID of an existing snapshot that you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Existing snapshot's location.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// Name of an existing snapshot that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The size of the snapshot to look for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	// Additionally, you can add `location` and `size` along with the `name` argument for a more refined search.
	// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
	// When this happens, please refine your search string so that it is specific enough to return only one result.
	Size pulumi.IntPtrInput `pulumi:"size"`
}

func (LookupSnapshotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotArgs)(nil)).Elem()
}

// A collection of values returned by getSnapshot.
type LookupSnapshotResultOutput struct{ *pulumi.OutputState }

func (LookupSnapshotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSnapshotResult)(nil)).Elem()
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutput() LookupSnapshotResultOutput {
	return o
}

func (o LookupSnapshotResultOutput) ToLookupSnapshotResultOutputWithContext(ctx context.Context) LookupSnapshotResultOutput {
	return o
}

// Is capable of CPU hot plug (no reboot required)
func (o LookupSnapshotResultOutput) CpuHotPlug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.CpuHotPlug }).(pulumi.BoolOutput)
}

// Is capable of CPU hot unplug (no reboot required)
func (o LookupSnapshotResultOutput) CpuHotUnplug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.CpuHotUnplug }).(pulumi.BoolOutput)
}

// Human readable description
func (o LookupSnapshotResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Description }).(pulumi.StringOutput)
}

// Is capable of SCSI drive hot plug (no reboot required)
func (o LookupSnapshotResultOutput) DiscScsiHotPlug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.DiscScsiHotPlug }).(pulumi.BoolOutput)
}

// Is capable of SCSI drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
func (o LookupSnapshotResultOutput) DiscScsiHotUnplug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.DiscScsiHotUnplug }).(pulumi.BoolOutput)
}

// Is capable of Virt-IO drive hot plug (no reboot required)
func (o LookupSnapshotResultOutput) DiscVirtioHotPlug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.DiscVirtioHotPlug }).(pulumi.BoolOutput)
}

// Is capable of Virt-IO drive hot unplug (no reboot required). This works only for non-Windows virtual Machines.
func (o LookupSnapshotResultOutput) DiscVirtioHotUnplug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.DiscVirtioHotUnplug }).(pulumi.BoolOutput)
}

// UUID of the snapshot
func (o LookupSnapshotResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSnapshotResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// OS type of this Snapshot
func (o LookupSnapshotResultOutput) LicenceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.LicenceType }).(pulumi.StringOutput)
}

// Location of that image/snapshot
func (o LookupSnapshotResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the snapshot.
func (o LookupSnapshotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSnapshotResult) string { return v.Name }).(pulumi.StringOutput)
}

// Is capable of nic hot plug (no reboot required)
func (o LookupSnapshotResultOutput) NicHotPlug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.NicHotPlug }).(pulumi.BoolOutput)
}

// Is capable of nic hot unplug (no reboot required)
func (o LookupSnapshotResultOutput) NicHotUnplug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.NicHotUnplug }).(pulumi.BoolOutput)
}

// Is capable of memory hot plug (no reboot required)
func (o LookupSnapshotResultOutput) RamHotPlug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.RamHotPlug }).(pulumi.BoolOutput)
}

// Is capable of memory hot unplug (no reboot required)
func (o LookupSnapshotResultOutput) RamHotUnplug() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.RamHotUnplug }).(pulumi.BoolOutput)
}

// Boolean value representing if the snapshot requires extra protection e.g. two factor protection
func (o LookupSnapshotResultOutput) SecAuthProtection() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupSnapshotResult) bool { return v.SecAuthProtection }).(pulumi.BoolOutput)
}

// The size of the image in GB
func (o LookupSnapshotResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSnapshotResult) int { return v.Size }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSnapshotResultOutput{})
}