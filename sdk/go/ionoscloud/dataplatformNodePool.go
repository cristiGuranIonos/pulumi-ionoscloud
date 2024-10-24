// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"errors"
	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud"
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDatacenter, err := compute.NewDatacenter(ctx, "exampleDatacenter", &compute.DatacenterArgs{
//				Location:    pulumi.String("de/txl"),
//				Description: pulumi.String("Datacenter for testing Dataplatform Cluster"),
//			})
//			if err != nil {
//				return err
//			}
//			exampleDataplatformCluster, err := ionoscloud.NewDataplatformCluster(ctx, "exampleDataplatformCluster", &ionoscloud.DataplatformClusterArgs{
//				DatacenterId: exampleDatacenter.ID(),
//				MaintenanceWindows: ionoscloud.DataplatformClusterMaintenanceWindowArray{
//					&ionoscloud.DataplatformClusterMaintenanceWindowArgs{
//						DayOfTheWeek: pulumi.String("Sunday"),
//						Time:         pulumi.String("09:00:00"),
//					},
//				},
//				Version: pulumi.String("23.7"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewDataplatformNodePool(ctx, "exampleDataplatformNodePool", &ionoscloud.DataplatformNodePoolArgs{
//				ClusterId:        exampleDataplatformCluster.ID(),
//				NodeCount:        pulumi.Int(1),
//				CpuFamily:        pulumi.String("INTEL_SKYLAKE"),
//				CoresCount:       pulumi.Int(1),
//				RamSize:          pulumi.Int(2048),
//				AvailabilityZone: pulumi.String("AUTO"),
//				StorageType:      pulumi.String("HDD"),
//				StorageSize:      pulumi.Int(10),
//				MaintenanceWindows: ionoscloud.DataplatformNodePoolMaintenanceWindowArray{
//					&ionoscloud.DataplatformNodePoolMaintenanceWindowArgs{
//						DayOfTheWeek: pulumi.String("Monday"),
//						Time:         pulumi.String("09:00:00"),
//					},
//				},
//				Labels: pulumi.StringMap{
//					"foo":   pulumi.String("bar"),
//					"color": pulumi.String("green"),
//				},
//				Annotations: pulumi.StringMap{
//					"ann1": pulumi.String("value1"),
//					"ann2": pulumi.String("value2"),
//				},
//			})
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
// ## Import
//
// A Dataplatform Node Pool resource can be imported using its cluster's UUID as well as its own UUID, e.g.:
//
// ```sh
// $ pulumi import ionoscloud:index/dataplatformNodePool:DataplatformNodePool mynodepool {dataplatform_cluster_uuid}/{dataplatform_nodepool_id}
// ```
type DataplatformNodePool struct {
	pulumi.CustomResourceState

	// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
	AvailabilityZone pulumi.StringOutput `pulumi:"availabilityZone"`
	// [string] The UUID of an existing Dataplatform cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
	CoresCount pulumi.IntOutput `pulumi:"coresCount"`
	// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
	CpuFamily pulumi.StringOutput `pulumi:"cpuFamily"`
	// The UUID of the virtual data center (VDC) in which the nodepool is provisioned
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformNodePoolMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringOutput `pulumi:"name"`
	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
	RamSize pulumi.IntOutput `pulumi:"ramSize"`
	// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
	StorageSize pulumi.IntOutput `pulumi:"storageSize"`
	// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// The version of the Data Platform.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDataplatformNodePool registers a new resource with the given unique name, arguments, and options.
func NewDataplatformNodePool(ctx *pulumi.Context,
	name string, args *DataplatformNodePoolArgs, opts ...pulumi.ResourceOption) (*DataplatformNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.NodeCount == nil {
		return nil, errors.New("invalid value for required argument 'NodeCount'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataplatformNodePool
	err := ctx.RegisterResource("ionoscloud:index/dataplatformNodePool:DataplatformNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplatformNodePool gets an existing DataplatformNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplatformNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplatformNodePoolState, opts ...pulumi.ResourceOption) (*DataplatformNodePool, error) {
	var resource DataplatformNodePool
	err := ctx.ReadResource("ionoscloud:index/dataplatformNodePool:DataplatformNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplatformNodePool resources.
type dataplatformNodePoolState struct {
	// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations map[string]string `pulumi:"annotations"`
	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// [string] The UUID of an existing Dataplatform cluster.
	ClusterId *string `pulumi:"clusterId"`
	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
	CoresCount *int `pulumi:"coresCount"`
	// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
	CpuFamily *string `pulumi:"cpuFamily"`
	// The UUID of the virtual data center (VDC) in which the nodepool is provisioned
	DatacenterId *string `pulumi:"datacenterId"`
	// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels map[string]string `pulumi:"labels"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []DataplatformNodePoolMaintenanceWindow `pulumi:"maintenanceWindows"`
	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `pulumi:"name"`
	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	NodeCount *int `pulumi:"nodeCount"`
	// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
	RamSize *int `pulumi:"ramSize"`
	// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
	StorageSize *int `pulumi:"storageSize"`
	// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
	StorageType *string `pulumi:"storageType"`
	// The version of the Data Platform.
	Version *string `pulumi:"version"`
}

type DataplatformNodePoolState struct {
	// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations pulumi.StringMapInput
	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
	AvailabilityZone pulumi.StringPtrInput
	// [string] The UUID of an existing Dataplatform cluster.
	ClusterId pulumi.StringPtrInput
	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
	CoresCount pulumi.IntPtrInput
	// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
	CpuFamily pulumi.StringPtrInput
	// The UUID of the virtual data center (VDC) in which the nodepool is provisioned
	DatacenterId pulumi.StringPtrInput
	// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels pulumi.StringMapInput
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformNodePoolMaintenanceWindowArrayInput
	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringPtrInput
	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	NodeCount pulumi.IntPtrInput
	// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
	RamSize pulumi.IntPtrInput
	// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
	StorageSize pulumi.IntPtrInput
	// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
	StorageType pulumi.StringPtrInput
	// The version of the Data Platform.
	Version pulumi.StringPtrInput
}

func (DataplatformNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplatformNodePoolState)(nil)).Elem()
}

type dataplatformNodePoolArgs struct {
	// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations map[string]string `pulumi:"annotations"`
	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
	AvailabilityZone *string `pulumi:"availabilityZone"`
	// [string] The UUID of an existing Dataplatform cluster.
	ClusterId string `pulumi:"clusterId"`
	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
	CoresCount *int `pulumi:"coresCount"`
	// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
	CpuFamily *string `pulumi:"cpuFamily"`
	// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels map[string]string `pulumi:"labels"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []DataplatformNodePoolMaintenanceWindow `pulumi:"maintenanceWindows"`
	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `pulumi:"name"`
	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	NodeCount int `pulumi:"nodeCount"`
	// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
	RamSize *int `pulumi:"ramSize"`
	// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
	StorageSize *int `pulumi:"storageSize"`
	// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
	StorageType *string `pulumi:"storageType"`
}

// The set of arguments for constructing a DataplatformNodePool resource.
type DataplatformNodePoolArgs struct {
	// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
	Annotations pulumi.StringMapInput
	// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
	AvailabilityZone pulumi.StringPtrInput
	// [string] The UUID of an existing Dataplatform cluster.
	ClusterId pulumi.StringInput
	// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
	CoresCount pulumi.IntPtrInput
	// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
	CpuFamily pulumi.StringPtrInput
	// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
	Labels pulumi.StringMapInput
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformNodePoolMaintenanceWindowArrayInput
	// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringPtrInput
	// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
	NodeCount pulumi.IntInput
	// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
	RamSize pulumi.IntPtrInput
	// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
	StorageSize pulumi.IntPtrInput
	// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
	StorageType pulumi.StringPtrInput
}

func (DataplatformNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplatformNodePoolArgs)(nil)).Elem()
}

type DataplatformNodePoolInput interface {
	pulumi.Input

	ToDataplatformNodePoolOutput() DataplatformNodePoolOutput
	ToDataplatformNodePoolOutputWithContext(ctx context.Context) DataplatformNodePoolOutput
}

func (*DataplatformNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplatformNodePool)(nil)).Elem()
}

func (i *DataplatformNodePool) ToDataplatformNodePoolOutput() DataplatformNodePoolOutput {
	return i.ToDataplatformNodePoolOutputWithContext(context.Background())
}

func (i *DataplatformNodePool) ToDataplatformNodePoolOutputWithContext(ctx context.Context) DataplatformNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformNodePoolOutput)
}

// DataplatformNodePoolArrayInput is an input type that accepts DataplatformNodePoolArray and DataplatformNodePoolArrayOutput values.
// You can construct a concrete instance of `DataplatformNodePoolArrayInput` via:
//
//	DataplatformNodePoolArray{ DataplatformNodePoolArgs{...} }
type DataplatformNodePoolArrayInput interface {
	pulumi.Input

	ToDataplatformNodePoolArrayOutput() DataplatformNodePoolArrayOutput
	ToDataplatformNodePoolArrayOutputWithContext(context.Context) DataplatformNodePoolArrayOutput
}

type DataplatformNodePoolArray []DataplatformNodePoolInput

func (DataplatformNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataplatformNodePool)(nil)).Elem()
}

func (i DataplatformNodePoolArray) ToDataplatformNodePoolArrayOutput() DataplatformNodePoolArrayOutput {
	return i.ToDataplatformNodePoolArrayOutputWithContext(context.Background())
}

func (i DataplatformNodePoolArray) ToDataplatformNodePoolArrayOutputWithContext(ctx context.Context) DataplatformNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformNodePoolArrayOutput)
}

// DataplatformNodePoolMapInput is an input type that accepts DataplatformNodePoolMap and DataplatformNodePoolMapOutput values.
// You can construct a concrete instance of `DataplatformNodePoolMapInput` via:
//
//	DataplatformNodePoolMap{ "key": DataplatformNodePoolArgs{...} }
type DataplatformNodePoolMapInput interface {
	pulumi.Input

	ToDataplatformNodePoolMapOutput() DataplatformNodePoolMapOutput
	ToDataplatformNodePoolMapOutputWithContext(context.Context) DataplatformNodePoolMapOutput
}

type DataplatformNodePoolMap map[string]DataplatformNodePoolInput

func (DataplatformNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataplatformNodePool)(nil)).Elem()
}

func (i DataplatformNodePoolMap) ToDataplatformNodePoolMapOutput() DataplatformNodePoolMapOutput {
	return i.ToDataplatformNodePoolMapOutputWithContext(context.Background())
}

func (i DataplatformNodePoolMap) ToDataplatformNodePoolMapOutputWithContext(ctx context.Context) DataplatformNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformNodePoolMapOutput)
}

type DataplatformNodePoolOutput struct{ *pulumi.OutputState }

func (DataplatformNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplatformNodePool)(nil)).Elem()
}

func (o DataplatformNodePoolOutput) ToDataplatformNodePoolOutput() DataplatformNodePoolOutput {
	return o
}

func (o DataplatformNodePoolOutput) ToDataplatformNodePoolOutputWithContext(ctx context.Context) DataplatformNodePoolOutput {
	return o
}

// [map] Key-value pairs attached to node pool resource as [Kubernetes annotations](https://kubernetes.io/docs/concepts/overview/working-with-objects/annotations/).
func (o DataplatformNodePoolOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// [string] The availability zone of the virtual datacenter region where the node pool resources should be provisioned. Must be set with one of the values `AUTO`, `ZONE_1` or `ZONE_2`. The default value is `AUTO`.
func (o DataplatformNodePoolOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.AvailabilityZone }).(pulumi.StringOutput)
}

// [string] The UUID of an existing Dataplatform cluster.
func (o DataplatformNodePoolOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// [int] The number of CPU cores per node. Must be set with a minimum value of 1. The default value is `4`.
func (o DataplatformNodePoolOutput) CoresCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.IntOutput { return v.CoresCount }).(pulumi.IntOutput)
}

// [string] A valid CPU family name or `AUTO` if the platform shall choose the best fitting option. Available CPU architectures can be retrieved from the datacenter resource. The default value is `AUTO`.
func (o DataplatformNodePoolOutput) CpuFamily() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.CpuFamily }).(pulumi.StringOutput)
}

// The UUID of the virtual data center (VDC) in which the nodepool is provisioned
func (o DataplatformNodePoolOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.DatacenterId }).(pulumi.StringOutput)
}

// [map] Key-value pairs attached to the node pool resource as [Kubernetes labels](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/).
func (o DataplatformNodePoolOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
func (o DataplatformNodePoolOutput) MaintenanceWindows() DataplatformNodePoolMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v *DataplatformNodePool) DataplatformNodePoolMaintenanceWindowArrayOutput {
		return v.MaintenanceWindows
	}).(DataplatformNodePoolMaintenanceWindowArrayOutput)
}

// [string] The name of your node pool. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
func (o DataplatformNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [int] The number of nodes that make up the node pool. Must be set with a minimum value of 1.
func (o DataplatformNodePoolOutput) NodeCount() pulumi.IntOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.IntOutput { return v.NodeCount }).(pulumi.IntOutput)
}

// [int] The RAM size for one node in MB. Must be set in multiples of `1024`MB, with a minimum size is of `2048`MB. The default value is `4096`.
func (o DataplatformNodePoolOutput) RamSize() pulumi.IntOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.IntOutput { return v.RamSize }).(pulumi.IntOutput)
}

// [int] The size of the volume in GB. The size must be greater than `10`GB. The default value is `20`.
func (o DataplatformNodePoolOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.IntOutput { return v.StorageSize }).(pulumi.IntOutput)
}

// [int] The type of hardware for the volume. Must be set with one of the values `HDD` or `SSD`. The default value is `SSD`.
func (o DataplatformNodePoolOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// The version of the Data Platform.
func (o DataplatformNodePoolOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformNodePool) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type DataplatformNodePoolArrayOutput struct{ *pulumi.OutputState }

func (DataplatformNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataplatformNodePool)(nil)).Elem()
}

func (o DataplatformNodePoolArrayOutput) ToDataplatformNodePoolArrayOutput() DataplatformNodePoolArrayOutput {
	return o
}

func (o DataplatformNodePoolArrayOutput) ToDataplatformNodePoolArrayOutputWithContext(ctx context.Context) DataplatformNodePoolArrayOutput {
	return o
}

func (o DataplatformNodePoolArrayOutput) Index(i pulumi.IntInput) DataplatformNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataplatformNodePool {
		return vs[0].([]*DataplatformNodePool)[vs[1].(int)]
	}).(DataplatformNodePoolOutput)
}

type DataplatformNodePoolMapOutput struct{ *pulumi.OutputState }

func (DataplatformNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataplatformNodePool)(nil)).Elem()
}

func (o DataplatformNodePoolMapOutput) ToDataplatformNodePoolMapOutput() DataplatformNodePoolMapOutput {
	return o
}

func (o DataplatformNodePoolMapOutput) ToDataplatformNodePoolMapOutputWithContext(ctx context.Context) DataplatformNodePoolMapOutput {
	return o
}

func (o DataplatformNodePoolMapOutput) MapIndex(k pulumi.StringInput) DataplatformNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataplatformNodePool {
		return vs[0].(map[string]*DataplatformNodePool)[vs[1].(string)]
	}).(DataplatformNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformNodePoolInput)(nil)).Elem(), &DataplatformNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformNodePoolArrayInput)(nil)).Elem(), DataplatformNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformNodePoolMapInput)(nil)).Elem(), DataplatformNodePoolMap{})
	pulumi.RegisterOutputType(DataplatformNodePoolOutput{})
	pulumi.RegisterOutputType(DataplatformNodePoolArrayOutput{})
	pulumi.RegisterOutputType(DataplatformNodePoolMapOutput{})
}