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

type MongoCluster struct {
	pulumi.CustomResourceState

	// Backup related properties.
	Backup MongoClusterBackupPtrOutput `pulumi:"backup"`
	// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
	// analysis.
	BiConnector MongoClusterBiConnectorOutput `pulumi:"biConnector"`
	// The connection string for your cluster.
	ConnectionString pulumi.StringOutput `pulumi:"connectionString"`
	// Details about the network connection for your cluster.
	Connections MongoClusterConnectionsOutput `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores pulumi.IntOutput `pulumi:"cores"`
	// The name of your cluster.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The cluster edition. Must be one of: playground, business, enterprise
	Edition pulumi.StringOutput `pulumi:"edition"`
	// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
	// at least 3.
	Instances pulumi.IntOutput `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
	// Update forces cluster re-creation.
	Location pulumi.StringOutput `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow MongoClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// The MongoDB version of your cluster. Update forces cluster re-creation.
	MongodbVersion pulumi.StringOutput `pulumi:"mongodbVersion"`
	// The amount of memory per instance in megabytes. Multiple of 1024
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The total number of shards in the cluster.
	Shards pulumi.IntPtrOutput `pulumi:"shards"`
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152
	StorageSize pulumi.IntOutput `pulumi:"storageSize"`
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
	// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
	// changes use the /templates endpoint.
	TemplateId pulumi.StringPtrOutput `pulumi:"templateId"`
	// The cluster type, either `replicaset` or `sharded-cluster`
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMongoCluster registers a new resource with the given unique name, arguments, and options.
func NewMongoCluster(ctx *pulumi.Context,
	name string, args *MongoClusterArgs, opts ...pulumi.ResourceOption) (*MongoCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connections == nil {
		return nil, errors.New("invalid value for required argument 'Connections'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Instances == nil {
		return nil, errors.New("invalid value for required argument 'Instances'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MongodbVersion == nil {
		return nil, errors.New("invalid value for required argument 'MongodbVersion'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MongoCluster
	err := ctx.RegisterResource("ionoscloud:index/mongoCluster:MongoCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoCluster gets an existing MongoCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoClusterState, opts ...pulumi.ResourceOption) (*MongoCluster, error) {
	var resource MongoCluster
	err := ctx.ReadResource("ionoscloud:index/mongoCluster:MongoCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoCluster resources.
type mongoClusterState struct {
	// Backup related properties.
	Backup *MongoClusterBackup `pulumi:"backup"`
	// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
	// analysis.
	BiConnector *MongoClusterBiConnector `pulumi:"biConnector"`
	// The connection string for your cluster.
	ConnectionString *string `pulumi:"connectionString"`
	// Details about the network connection for your cluster.
	Connections *MongoClusterConnections `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores *int `pulumi:"cores"`
	// The name of your cluster.
	DisplayName *string `pulumi:"displayName"`
	// The cluster edition. Must be one of: playground, business, enterprise
	Edition *string `pulumi:"edition"`
	// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
	// at least 3.
	Instances *int `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
	// Update forces cluster re-creation.
	Location *string `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MongoClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The MongoDB version of your cluster. Update forces cluster re-creation.
	MongodbVersion *string `pulumi:"mongodbVersion"`
	// The amount of memory per instance in megabytes. Multiple of 1024
	Ram *int `pulumi:"ram"`
	// The total number of shards in the cluster.
	Shards *int `pulumi:"shards"`
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152
	StorageSize *int `pulumi:"storageSize"`
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
	StorageType *string `pulumi:"storageType"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
	// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
	// changes use the /templates endpoint.
	TemplateId *string `pulumi:"templateId"`
	// The cluster type, either `replicaset` or `sharded-cluster`
	Type *string `pulumi:"type"`
}

type MongoClusterState struct {
	// Backup related properties.
	Backup MongoClusterBackupPtrInput
	// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
	// analysis.
	BiConnector MongoClusterBiConnectorPtrInput
	// The connection string for your cluster.
	ConnectionString pulumi.StringPtrInput
	// Details about the network connection for your cluster.
	Connections MongoClusterConnectionsPtrInput
	// The number of CPU cores per instance.
	Cores pulumi.IntPtrInput
	// The name of your cluster.
	DisplayName pulumi.StringPtrInput
	// The cluster edition. Must be one of: playground, business, enterprise
	Edition pulumi.StringPtrInput
	// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
	// at least 3.
	Instances pulumi.IntPtrInput
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
	// Update forces cluster re-creation.
	Location pulumi.StringPtrInput
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow MongoClusterMaintenanceWindowPtrInput
	// The MongoDB version of your cluster. Update forces cluster re-creation.
	MongodbVersion pulumi.StringPtrInput
	// The amount of memory per instance in megabytes. Multiple of 1024
	Ram pulumi.IntPtrInput
	// The total number of shards in the cluster.
	Shards pulumi.IntPtrInput
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152
	StorageSize pulumi.IntPtrInput
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
	StorageType pulumi.StringPtrInput
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
	// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
	// changes use the /templates endpoint.
	TemplateId pulumi.StringPtrInput
	// The cluster type, either `replicaset` or `sharded-cluster`
	Type pulumi.StringPtrInput
}

func (MongoClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterState)(nil)).Elem()
}

type mongoClusterArgs struct {
	// Backup related properties.
	Backup *MongoClusterBackup `pulumi:"backup"`
	// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
	// analysis.
	BiConnector *MongoClusterBiConnector `pulumi:"biConnector"`
	// Details about the network connection for your cluster.
	Connections MongoClusterConnections `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores *int `pulumi:"cores"`
	// The name of your cluster.
	DisplayName string `pulumi:"displayName"`
	// The cluster edition. Must be one of: playground, business, enterprise
	Edition *string `pulumi:"edition"`
	// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
	// at least 3.
	Instances int `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
	// Update forces cluster re-creation.
	Location string `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *MongoClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The MongoDB version of your cluster. Update forces cluster re-creation.
	MongodbVersion string `pulumi:"mongodbVersion"`
	// The amount of memory per instance in megabytes. Multiple of 1024
	Ram *int `pulumi:"ram"`
	// The total number of shards in the cluster.
	Shards *int `pulumi:"shards"`
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152
	StorageSize *int `pulumi:"storageSize"`
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
	StorageType *string `pulumi:"storageType"`
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
	// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
	// changes use the /templates endpoint.
	TemplateId *string `pulumi:"templateId"`
	// The cluster type, either `replicaset` or `sharded-cluster`
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a MongoCluster resource.
type MongoClusterArgs struct {
	// Backup related properties.
	Backup MongoClusterBackupPtrInput
	// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
	// analysis.
	BiConnector MongoClusterBiConnectorPtrInput
	// Details about the network connection for your cluster.
	Connections MongoClusterConnectionsInput
	// The number of CPU cores per instance.
	Cores pulumi.IntPtrInput
	// The name of your cluster.
	DisplayName pulumi.StringInput
	// The cluster edition. Must be one of: playground, business, enterprise
	Edition pulumi.StringPtrInput
	// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
	// at least 3.
	Instances pulumi.IntInput
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
	// Update forces cluster re-creation.
	Location pulumi.StringInput
	// A weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow MongoClusterMaintenanceWindowPtrInput
	// The MongoDB version of your cluster. Update forces cluster re-creation.
	MongodbVersion pulumi.StringInput
	// The amount of memory per instance in megabytes. Multiple of 1024
	Ram pulumi.IntPtrInput
	// The total number of shards in the cluster.
	Shards pulumi.IntPtrInput
	// The amount of storage per instance in megabytes. At least 5120, at most 2097152
	StorageSize pulumi.IntPtrInput
	// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
	StorageType pulumi.StringPtrInput
	// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
	// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
	// changes use the /templates endpoint.
	TemplateId pulumi.StringPtrInput
	// The cluster type, either `replicaset` or `sharded-cluster`
	Type pulumi.StringPtrInput
}

func (MongoClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoClusterArgs)(nil)).Elem()
}

type MongoClusterInput interface {
	pulumi.Input

	ToMongoClusterOutput() MongoClusterOutput
	ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput
}

func (*MongoCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCluster)(nil)).Elem()
}

func (i *MongoCluster) ToMongoClusterOutput() MongoClusterOutput {
	return i.ToMongoClusterOutputWithContext(context.Background())
}

func (i *MongoCluster) ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoClusterOutput)
}

// MongoClusterArrayInput is an input type that accepts MongoClusterArray and MongoClusterArrayOutput values.
// You can construct a concrete instance of `MongoClusterArrayInput` via:
//
//	MongoClusterArray{ MongoClusterArgs{...} }
type MongoClusterArrayInput interface {
	pulumi.Input

	ToMongoClusterArrayOutput() MongoClusterArrayOutput
	ToMongoClusterArrayOutputWithContext(context.Context) MongoClusterArrayOutput
}

type MongoClusterArray []MongoClusterInput

func (MongoClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoCluster)(nil)).Elem()
}

func (i MongoClusterArray) ToMongoClusterArrayOutput() MongoClusterArrayOutput {
	return i.ToMongoClusterArrayOutputWithContext(context.Background())
}

func (i MongoClusterArray) ToMongoClusterArrayOutputWithContext(ctx context.Context) MongoClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoClusterArrayOutput)
}

// MongoClusterMapInput is an input type that accepts MongoClusterMap and MongoClusterMapOutput values.
// You can construct a concrete instance of `MongoClusterMapInput` via:
//
//	MongoClusterMap{ "key": MongoClusterArgs{...} }
type MongoClusterMapInput interface {
	pulumi.Input

	ToMongoClusterMapOutput() MongoClusterMapOutput
	ToMongoClusterMapOutputWithContext(context.Context) MongoClusterMapOutput
}

type MongoClusterMap map[string]MongoClusterInput

func (MongoClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoCluster)(nil)).Elem()
}

func (i MongoClusterMap) ToMongoClusterMapOutput() MongoClusterMapOutput {
	return i.ToMongoClusterMapOutputWithContext(context.Background())
}

func (i MongoClusterMap) ToMongoClusterMapOutputWithContext(ctx context.Context) MongoClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoClusterMapOutput)
}

type MongoClusterOutput struct{ *pulumi.OutputState }

func (MongoClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoCluster)(nil)).Elem()
}

func (o MongoClusterOutput) ToMongoClusterOutput() MongoClusterOutput {
	return o
}

func (o MongoClusterOutput) ToMongoClusterOutputWithContext(ctx context.Context) MongoClusterOutput {
	return o
}

// Backup related properties.
func (o MongoClusterOutput) Backup() MongoClusterBackupPtrOutput {
	return o.ApplyT(func(v *MongoCluster) MongoClusterBackupPtrOutput { return v.Backup }).(MongoClusterBackupPtrOutput)
}

// The MongoDB Connector for Business Intelligence allows you to query a MongoDB database using SQL commands to aid in data
// analysis.
func (o MongoClusterOutput) BiConnector() MongoClusterBiConnectorOutput {
	return o.ApplyT(func(v *MongoCluster) MongoClusterBiConnectorOutput { return v.BiConnector }).(MongoClusterBiConnectorOutput)
}

// The connection string for your cluster.
func (o MongoClusterOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.ConnectionString }).(pulumi.StringOutput)
}

// Details about the network connection for your cluster.
func (o MongoClusterOutput) Connections() MongoClusterConnectionsOutput {
	return o.ApplyT(func(v *MongoCluster) MongoClusterConnectionsOutput { return v.Connections }).(MongoClusterConnectionsOutput)
}

// The number of CPU cores per instance.
func (o MongoClusterOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.IntOutput { return v.Cores }).(pulumi.IntOutput)
}

// The name of your cluster.
func (o MongoClusterOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The cluster edition. Must be one of: playground, business, enterprise
func (o MongoClusterOutput) Edition() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Edition }).(pulumi.StringOutput)
}

// The total number of instances in the cluster (one master and n-1 standbys). Example: 1, 3, 5, 7. For enterprise edition
// at least 3.
func (o MongoClusterOutput) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.IntOutput { return v.Instances }).(pulumi.IntOutput)
}

// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
// be modified after datacenter creation (disallowed in update requests). Available locations: de/txl, gb/lhr, es/vit.
// Update forces cluster re-creation.
func (o MongoClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// A weekly 4 hour-long window, during which maintenance might occur
func (o MongoClusterOutput) MaintenanceWindow() MongoClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *MongoCluster) MongoClusterMaintenanceWindowOutput { return v.MaintenanceWindow }).(MongoClusterMaintenanceWindowOutput)
}

// The MongoDB version of your cluster. Update forces cluster re-creation.
func (o MongoClusterOutput) MongodbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.MongodbVersion }).(pulumi.StringOutput)
}

// The amount of memory per instance in megabytes. Multiple of 1024
func (o MongoClusterOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// The total number of shards in the cluster.
func (o MongoClusterOutput) Shards() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.IntPtrOutput { return v.Shards }).(pulumi.IntPtrOutput)
}

// The amount of storage per instance in megabytes. At least 5120, at most 2097152
func (o MongoClusterOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.IntOutput { return v.StorageSize }).(pulumi.IntOutput)
}

// The storage type. One of : HDD, SSD, SSD Standard, SSD Premium
func (o MongoClusterOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// The unique ID of the template, which specifies the number of cores, storage size, and memory. You cannot downgrade to a
// smaller template or minor edition (e.g. from business to playground). To get a list of all templates to confirm the
// changes use the /templates endpoint.
func (o MongoClusterOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringPtrOutput { return v.TemplateId }).(pulumi.StringPtrOutput)
}

// The cluster type, either `replicaset` or `sharded-cluster`
func (o MongoClusterOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoCluster) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type MongoClusterArrayOutput struct{ *pulumi.OutputState }

func (MongoClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoCluster)(nil)).Elem()
}

func (o MongoClusterArrayOutput) ToMongoClusterArrayOutput() MongoClusterArrayOutput {
	return o
}

func (o MongoClusterArrayOutput) ToMongoClusterArrayOutputWithContext(ctx context.Context) MongoClusterArrayOutput {
	return o
}

func (o MongoClusterArrayOutput) Index(i pulumi.IntInput) MongoClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MongoCluster {
		return vs[0].([]*MongoCluster)[vs[1].(int)]
	}).(MongoClusterOutput)
}

type MongoClusterMapOutput struct{ *pulumi.OutputState }

func (MongoClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoCluster)(nil)).Elem()
}

func (o MongoClusterMapOutput) ToMongoClusterMapOutput() MongoClusterMapOutput {
	return o
}

func (o MongoClusterMapOutput) ToMongoClusterMapOutputWithContext(ctx context.Context) MongoClusterMapOutput {
	return o
}

func (o MongoClusterMapOutput) MapIndex(k pulumi.StringInput) MongoClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MongoCluster {
		return vs[0].(map[string]*MongoCluster)[vs[1].(string)]
	}).(MongoClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MongoClusterInput)(nil)).Elem(), &MongoCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoClusterArrayInput)(nil)).Elem(), MongoClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoClusterMapInput)(nil)).Elem(), MongoClusterMap{})
	pulumi.RegisterOutputType(MongoClusterOutput{})
	pulumi.RegisterOutputType(MongoClusterArrayOutput{})
	pulumi.RegisterOutputType(MongoClusterMapOutput{})
}