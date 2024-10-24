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

// Manages a **Dataplatform Cluster**.
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
//			exampleLan, err := ionoscloud.NewLan(ctx, "exampleLan", &ionoscloud.LanArgs{
//				DatacenterId: exampleDatacenter.ID(),
//				Public:       pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewDataplatformCluster(ctx, "exampleDataplatformCluster", &ionoscloud.DataplatformClusterArgs{
//				DatacenterId: exampleDatacenter.ID(),
//				MaintenanceWindows: ionoscloud.DataplatformClusterMaintenanceWindowArray{
//					&ionoscloud.DataplatformClusterMaintenanceWindowArgs{
//						DayOfTheWeek: pulumi.String("Sunday"),
//						Time:         pulumi.String("09:00:00"),
//					},
//				},
//				Version: pulumi.String("23.11"),
//				Lans: ionoscloud.DataplatformClusterLanArray{
//					&ionoscloud.DataplatformClusterLanArgs{
//						LanId: exampleLan.ID(),
//						Dhcp:  pulumi.Bool(false),
//						Routes: ionoscloud.DataplatformClusterLanRouteArray{
//							&ionoscloud.DataplatformClusterLanRouteArgs{
//								Network: pulumi.String("182.168.42.1/24"),
//								Gateway: pulumi.String("192.168.42.1"),
//							},
//						},
//					},
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
// Resource Dataplatform Cluster can be imported using the `cluster_id`, e.g.
//
// ```sh
// $ pulumi import ionoscloud:index/dataplatformCluster:DataplatformCluster mycluser {cluster uuid}
// ```
type DataplatformCluster struct {
	pulumi.CustomResourceState

	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId pulumi.StringOutput `pulumi:"datacenterId"`
	// [list] A list of LANs you want this node pool to be part of.
	Lans DataplatformClusterLanArrayOutput `pulumi:"lans"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformClusterMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringOutput `pulumi:"name"`
	// [int] The version of the Data Platform.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewDataplatformCluster registers a new resource with the given unique name, arguments, and options.
func NewDataplatformCluster(ctx *pulumi.Context,
	name string, args *DataplatformClusterArgs, opts ...pulumi.ResourceOption) (*DataplatformCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatacenterId == nil {
		return nil, errors.New("invalid value for required argument 'DatacenterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DataplatformCluster
	err := ctx.RegisterResource("ionoscloud:index/dataplatformCluster:DataplatformCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDataplatformCluster gets an existing DataplatformCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDataplatformCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataplatformClusterState, opts ...pulumi.ResourceOption) (*DataplatformCluster, error) {
	var resource DataplatformCluster
	err := ctx.ReadResource("ionoscloud:index/dataplatformCluster:DataplatformCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DataplatformCluster resources.
type dataplatformClusterState struct {
	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId *string `pulumi:"datacenterId"`
	// [list] A list of LANs you want this node pool to be part of.
	Lans []DataplatformClusterLan `pulumi:"lans"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []DataplatformClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `pulumi:"name"`
	// [int] The version of the Data Platform.
	Version *string `pulumi:"version"`
}

type DataplatformClusterState struct {
	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId pulumi.StringPtrInput
	// [list] A list of LANs you want this node pool to be part of.
	Lans DataplatformClusterLanArrayInput
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformClusterMaintenanceWindowArrayInput
	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringPtrInput
	// [int] The version of the Data Platform.
	Version pulumi.StringPtrInput
}

func (DataplatformClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplatformClusterState)(nil)).Elem()
}

type dataplatformClusterArgs struct {
	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId string `pulumi:"datacenterId"`
	// [list] A list of LANs you want this node pool to be part of.
	Lans []DataplatformClusterLan `pulumi:"lans"`
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []DataplatformClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name *string `pulumi:"name"`
	// [int] The version of the Data Platform.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a DataplatformCluster resource.
type DataplatformClusterArgs struct {
	// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
	DatacenterId pulumi.StringInput
	// [list] A list of LANs you want this node pool to be part of.
	Lans DataplatformClusterLanArrayInput
	// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows DataplatformClusterMaintenanceWindowArrayInput
	// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
	Name pulumi.StringPtrInput
	// [int] The version of the Data Platform.
	Version pulumi.StringPtrInput
}

func (DataplatformClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataplatformClusterArgs)(nil)).Elem()
}

type DataplatformClusterInput interface {
	pulumi.Input

	ToDataplatformClusterOutput() DataplatformClusterOutput
	ToDataplatformClusterOutputWithContext(ctx context.Context) DataplatformClusterOutput
}

func (*DataplatformCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplatformCluster)(nil)).Elem()
}

func (i *DataplatformCluster) ToDataplatformClusterOutput() DataplatformClusterOutput {
	return i.ToDataplatformClusterOutputWithContext(context.Background())
}

func (i *DataplatformCluster) ToDataplatformClusterOutputWithContext(ctx context.Context) DataplatformClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformClusterOutput)
}

// DataplatformClusterArrayInput is an input type that accepts DataplatformClusterArray and DataplatformClusterArrayOutput values.
// You can construct a concrete instance of `DataplatformClusterArrayInput` via:
//
//	DataplatformClusterArray{ DataplatformClusterArgs{...} }
type DataplatformClusterArrayInput interface {
	pulumi.Input

	ToDataplatformClusterArrayOutput() DataplatformClusterArrayOutput
	ToDataplatformClusterArrayOutputWithContext(context.Context) DataplatformClusterArrayOutput
}

type DataplatformClusterArray []DataplatformClusterInput

func (DataplatformClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataplatformCluster)(nil)).Elem()
}

func (i DataplatformClusterArray) ToDataplatformClusterArrayOutput() DataplatformClusterArrayOutput {
	return i.ToDataplatformClusterArrayOutputWithContext(context.Background())
}

func (i DataplatformClusterArray) ToDataplatformClusterArrayOutputWithContext(ctx context.Context) DataplatformClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformClusterArrayOutput)
}

// DataplatformClusterMapInput is an input type that accepts DataplatformClusterMap and DataplatformClusterMapOutput values.
// You can construct a concrete instance of `DataplatformClusterMapInput` via:
//
//	DataplatformClusterMap{ "key": DataplatformClusterArgs{...} }
type DataplatformClusterMapInput interface {
	pulumi.Input

	ToDataplatformClusterMapOutput() DataplatformClusterMapOutput
	ToDataplatformClusterMapOutputWithContext(context.Context) DataplatformClusterMapOutput
}

type DataplatformClusterMap map[string]DataplatformClusterInput

func (DataplatformClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataplatformCluster)(nil)).Elem()
}

func (i DataplatformClusterMap) ToDataplatformClusterMapOutput() DataplatformClusterMapOutput {
	return i.ToDataplatformClusterMapOutputWithContext(context.Background())
}

func (i DataplatformClusterMap) ToDataplatformClusterMapOutputWithContext(ctx context.Context) DataplatformClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataplatformClusterMapOutput)
}

type DataplatformClusterOutput struct{ *pulumi.OutputState }

func (DataplatformClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataplatformCluster)(nil)).Elem()
}

func (o DataplatformClusterOutput) ToDataplatformClusterOutput() DataplatformClusterOutput {
	return o
}

func (o DataplatformClusterOutput) ToDataplatformClusterOutputWithContext(ctx context.Context) DataplatformClusterOutput {
	return o
}

// [string] The UUID of the virtual data center (VDC) the cluster is provisioned.
func (o DataplatformClusterOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformCluster) pulumi.StringOutput { return v.DatacenterId }).(pulumi.StringOutput)
}

// [list] A list of LANs you want this node pool to be part of.
func (o DataplatformClusterOutput) Lans() DataplatformClusterLanArrayOutput {
	return o.ApplyT(func(v *DataplatformCluster) DataplatformClusterLanArrayOutput { return v.Lans }).(DataplatformClusterLanArrayOutput)
}

// [string] Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
func (o DataplatformClusterOutput) MaintenanceWindows() DataplatformClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v *DataplatformCluster) DataplatformClusterMaintenanceWindowArrayOutput {
		return v.MaintenanceWindows
	}).(DataplatformClusterMaintenanceWindowArrayOutput)
}

// [string] The name of your cluster. Must be 63 characters or less and must be empty or begin and end with an alphanumeric character ([a-z0-9A-Z]). It can contain dashes (-), underscores (_), dots (.), and alphanumerics in-between.
func (o DataplatformClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [int] The version of the Data Platform.
func (o DataplatformClusterOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *DataplatformCluster) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type DataplatformClusterArrayOutput struct{ *pulumi.OutputState }

func (DataplatformClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DataplatformCluster)(nil)).Elem()
}

func (o DataplatformClusterArrayOutput) ToDataplatformClusterArrayOutput() DataplatformClusterArrayOutput {
	return o
}

func (o DataplatformClusterArrayOutput) ToDataplatformClusterArrayOutputWithContext(ctx context.Context) DataplatformClusterArrayOutput {
	return o
}

func (o DataplatformClusterArrayOutput) Index(i pulumi.IntInput) DataplatformClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DataplatformCluster {
		return vs[0].([]*DataplatformCluster)[vs[1].(int)]
	}).(DataplatformClusterOutput)
}

type DataplatformClusterMapOutput struct{ *pulumi.OutputState }

func (DataplatformClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DataplatformCluster)(nil)).Elem()
}

func (o DataplatformClusterMapOutput) ToDataplatformClusterMapOutput() DataplatformClusterMapOutput {
	return o
}

func (o DataplatformClusterMapOutput) ToDataplatformClusterMapOutputWithContext(ctx context.Context) DataplatformClusterMapOutput {
	return o
}

func (o DataplatformClusterMapOutput) MapIndex(k pulumi.StringInput) DataplatformClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DataplatformCluster {
		return vs[0].(map[string]*DataplatformCluster)[vs[1].(string)]
	}).(DataplatformClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformClusterInput)(nil)).Elem(), &DataplatformCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformClusterArrayInput)(nil)).Elem(), DataplatformClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DataplatformClusterMapInput)(nil)).Elem(), DataplatformClusterMap{})
	pulumi.RegisterOutputType(DataplatformClusterOutput{})
	pulumi.RegisterOutputType(DataplatformClusterArrayOutput{})
	pulumi.RegisterOutputType(DataplatformClusterMapOutput{})
}
