// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Dataplatform Node Pool Data Source** can be used to search for and return an existing Dataplatform Node Pool under a Dataplatform Cluster.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search and make sure that your resources have unique names.
//
// ## Example Usage
func LookupDataplatformNodePool(ctx *pulumi.Context, args *LookupDataplatformNodePoolArgs, opts ...pulumi.InvokeOption) (*LookupDataplatformNodePoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataplatformNodePoolResult
	err := ctx.Invoke("ionoscloud:index/getDataplatformNodePool:getDataplatformNodePool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataplatformNodePool.
type LookupDataplatformNodePoolArgs struct {
	// ID of the cluster the searched node pool is part of.
	ClusterId string `pulumi:"clusterId"`
	// ID of the node pool you want to search for.
	Id *string `pulumi:"id"`
	// Name of an existing cluster that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name *string `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch *bool `pulumi:"partialMatch"`
}

// A collection of values returned by getDataplatformNodePool.
type LookupDataplatformNodePoolResult struct {
	// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations map[string]string `pulumi:"annotations"`
	// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
	AvailabilityZone string `pulumi:"availabilityZone"`
	// ID of the cluster the searched node pool is part of.
	ClusterId string `pulumi:"clusterId"`
	// The number of CPU cores per node.
	CoresCount int `pulumi:"coresCount"`
	// A CPU family.
	CpuFamily string `pulumi:"cpuFamily"`
	// The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId string `pulumi:"datacenterId"`
	// ID of your node pool.
	Id *string `pulumi:"id"`
	// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels map[string]string `pulumi:"labels"`
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []GetDataplatformNodePoolMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of your node pool
	Name *string `pulumi:"name"`
	// The number of nodes that make up the node pool.
	NodeCount    int   `pulumi:"nodeCount"`
	PartialMatch *bool `pulumi:"partialMatch"`
	// The RAM size for one node in MB.
	RamSize int `pulumi:"ramSize"`
	// The size of the volume in GB.
	StorageSize int `pulumi:"storageSize"`
	// The type of hardware for the volume.
	StorageType string `pulumi:"storageType"`
	// The version of the Data Platform.
	Version string `pulumi:"version"`
}

func LookupDataplatformNodePoolOutput(ctx *pulumi.Context, args LookupDataplatformNodePoolOutputArgs, opts ...pulumi.InvokeOption) LookupDataplatformNodePoolResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataplatformNodePoolResult, error) {
			args := v.(LookupDataplatformNodePoolArgs)
			r, err := LookupDataplatformNodePool(ctx, &args, opts...)
			var s LookupDataplatformNodePoolResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataplatformNodePoolResultOutput)
}

// A collection of arguments for invoking getDataplatformNodePool.
type LookupDataplatformNodePoolOutputArgs struct {
	// ID of the cluster the searched node pool is part of.
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// ID of the node pool you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing cluster that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch pulumi.BoolPtrInput `pulumi:"partialMatch"`
}

func (LookupDataplatformNodePoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataplatformNodePoolArgs)(nil)).Elem()
}

// A collection of values returned by getDataplatformNodePool.
type LookupDataplatformNodePoolResultOutput struct{ *pulumi.OutputState }

func (LookupDataplatformNodePoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataplatformNodePoolResult)(nil)).Elem()
}

func (o LookupDataplatformNodePoolResultOutput) ToLookupDataplatformNodePoolResultOutput() LookupDataplatformNodePoolResultOutput {
	return o
}

func (o LookupDataplatformNodePoolResultOutput) ToLookupDataplatformNodePoolResultOutputWithContext(ctx context.Context) LookupDataplatformNodePoolResultOutput {
	return o
}

// Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
func (o LookupDataplatformNodePoolResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// The availability zone of the virtual datacenter region where the node pool resources should be provisioned.
func (o LookupDataplatformNodePoolResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// ID of the cluster the searched node pool is part of.
func (o LookupDataplatformNodePoolResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// The number of CPU cores per node.
func (o LookupDataplatformNodePoolResultOutput) CoresCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) int { return v.CoresCount }).(pulumi.IntOutput)
}

// A CPU family.
func (o LookupDataplatformNodePoolResultOutput) CpuFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.CpuFamily }).(pulumi.StringOutput)
}

// The UUID of the virtual data center (VDC) the cluster is provisioned.
func (o LookupDataplatformNodePoolResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// ID of your node pool.
func (o LookupDataplatformNodePoolResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
func (o LookupDataplatformNodePoolResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
func (o LookupDataplatformNodePoolResultOutput) MaintenanceWindows() GetDataplatformNodePoolMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) []GetDataplatformNodePoolMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetDataplatformNodePoolMaintenanceWindowArrayOutput)
}

// The name of your node pool
func (o LookupDataplatformNodePoolResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The number of nodes that make up the node pool.
func (o LookupDataplatformNodePoolResultOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) int { return v.NodeCount }).(pulumi.IntOutput)
}

func (o LookupDataplatformNodePoolResultOutput) PartialMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) *bool { return v.PartialMatch }).(pulumi.BoolPtrOutput)
}

// The RAM size for one node in MB.
func (o LookupDataplatformNodePoolResultOutput) RamSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) int { return v.RamSize }).(pulumi.IntOutput)
}

// The size of the volume in GB.
func (o LookupDataplatformNodePoolResultOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) int { return v.StorageSize }).(pulumi.IntOutput)
}

// The type of hardware for the volume.
func (o LookupDataplatformNodePoolResultOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.StorageType }).(pulumi.StringOutput)
}

// The version of the Data Platform.
func (o LookupDataplatformNodePoolResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformNodePoolResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataplatformNodePoolResultOutput{})
}
