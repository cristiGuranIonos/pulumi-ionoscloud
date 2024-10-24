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

type PgCluster struct {
	pulumi.CustomResourceState

	// The S3 location where the backups will be stored.
	BackupLocation pulumi.StringOutput `pulumi:"backupLocation"`
	// Configuration options for the connection pooler
	ConnectionPooler PgClusterConnectionPoolerOutput `pulumi:"connectionPooler"`
	// Details about the network connection for your cluster.
	Connections PgClusterConnectionsPtrOutput `pulumi:"connections"`
	// The number of CPU cores per replica.
	Cores pulumi.IntOutput `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials PgClusterCredentialsOutput `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The DNS name pointing to your cluster
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// Creates the cluster based on the existing backup.
	FromBackup PgClusterFromBackupPtrOutput `pulumi:"fromBackup"`
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances pulumi.IntOutput `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests)
	Location pulumi.StringOutput `pulumi:"location"`
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow PgClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// The PostgreSQL version of your cluster.
	PostgresVersion pulumi.StringOutput `pulumi:"postgresVersion"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize pulumi.IntOutput `pulumi:"storageSize"`
	// The storage type used in your cluster.
	StorageType pulumi.StringOutput `pulumi:"storageType"`
	// Represents different modes of replication.
	SynchronizationMode pulumi.StringOutput `pulumi:"synchronizationMode"`
}

// NewPgCluster registers a new resource with the given unique name, arguments, and options.
func NewPgCluster(ctx *pulumi.Context,
	name string, args *PgClusterArgs, opts ...pulumi.ResourceOption) (*PgCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Cores == nil {
		return nil, errors.New("invalid value for required argument 'Cores'")
	}
	if args.Credentials == nil {
		return nil, errors.New("invalid value for required argument 'Credentials'")
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
	if args.PostgresVersion == nil {
		return nil, errors.New("invalid value for required argument 'PostgresVersion'")
	}
	if args.Ram == nil {
		return nil, errors.New("invalid value for required argument 'Ram'")
	}
	if args.StorageSize == nil {
		return nil, errors.New("invalid value for required argument 'StorageSize'")
	}
	if args.StorageType == nil {
		return nil, errors.New("invalid value for required argument 'StorageType'")
	}
	if args.SynchronizationMode == nil {
		return nil, errors.New("invalid value for required argument 'SynchronizationMode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PgCluster
	err := ctx.RegisterResource("ionoscloud:index/pgCluster:PgCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgCluster gets an existing PgCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgClusterState, opts ...pulumi.ResourceOption) (*PgCluster, error) {
	var resource PgCluster
	err := ctx.ReadResource("ionoscloud:index/pgCluster:PgCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgCluster resources.
type pgClusterState struct {
	// The S3 location where the backups will be stored.
	BackupLocation *string `pulumi:"backupLocation"`
	// Configuration options for the connection pooler
	ConnectionPooler *PgClusterConnectionPooler `pulumi:"connectionPooler"`
	// Details about the network connection for your cluster.
	Connections *PgClusterConnections `pulumi:"connections"`
	// The number of CPU cores per replica.
	Cores *int `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials *PgClusterCredentials `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName *string `pulumi:"displayName"`
	// The DNS name pointing to your cluster
	DnsName *string `pulumi:"dnsName"`
	// Creates the cluster based on the existing backup.
	FromBackup *PgClusterFromBackup `pulumi:"fromBackup"`
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances *int `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests)
	Location *string `pulumi:"location"`
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *PgClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The PostgreSQL version of your cluster.
	PostgresVersion *string `pulumi:"postgresVersion"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram *int `pulumi:"ram"`
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize *int `pulumi:"storageSize"`
	// The storage type used in your cluster.
	StorageType *string `pulumi:"storageType"`
	// Represents different modes of replication.
	SynchronizationMode *string `pulumi:"synchronizationMode"`
}

type PgClusterState struct {
	// The S3 location where the backups will be stored.
	BackupLocation pulumi.StringPtrInput
	// Configuration options for the connection pooler
	ConnectionPooler PgClusterConnectionPoolerPtrInput
	// Details about the network connection for your cluster.
	Connections PgClusterConnectionsPtrInput
	// The number of CPU cores per replica.
	Cores pulumi.IntPtrInput
	// Credentials for the database user to be created.
	Credentials PgClusterCredentialsPtrInput
	// The friendly name of your cluster.
	DisplayName pulumi.StringPtrInput
	// The DNS name pointing to your cluster
	DnsName pulumi.StringPtrInput
	// Creates the cluster based on the existing backup.
	FromBackup PgClusterFromBackupPtrInput
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances pulumi.IntPtrInput
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests)
	Location pulumi.StringPtrInput
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow PgClusterMaintenanceWindowPtrInput
	// The PostgreSQL version of your cluster.
	PostgresVersion pulumi.StringPtrInput
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram pulumi.IntPtrInput
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize pulumi.IntPtrInput
	// The storage type used in your cluster.
	StorageType pulumi.StringPtrInput
	// Represents different modes of replication.
	SynchronizationMode pulumi.StringPtrInput
}

func (PgClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgClusterState)(nil)).Elem()
}

type pgClusterArgs struct {
	// The S3 location where the backups will be stored.
	BackupLocation *string `pulumi:"backupLocation"`
	// Configuration options for the connection pooler
	ConnectionPooler *PgClusterConnectionPooler `pulumi:"connectionPooler"`
	// Details about the network connection for your cluster.
	Connections *PgClusterConnections `pulumi:"connections"`
	// The number of CPU cores per replica.
	Cores int `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials PgClusterCredentials `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName string `pulumi:"displayName"`
	// Creates the cluster based on the existing backup.
	FromBackup *PgClusterFromBackup `pulumi:"fromBackup"`
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances int `pulumi:"instances"`
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests)
	Location string `pulumi:"location"`
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow *PgClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The PostgreSQL version of your cluster.
	PostgresVersion string `pulumi:"postgresVersion"`
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram int `pulumi:"ram"`
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize int `pulumi:"storageSize"`
	// The storage type used in your cluster.
	StorageType string `pulumi:"storageType"`
	// Represents different modes of replication.
	SynchronizationMode string `pulumi:"synchronizationMode"`
}

// The set of arguments for constructing a PgCluster resource.
type PgClusterArgs struct {
	// The S3 location where the backups will be stored.
	BackupLocation pulumi.StringPtrInput
	// Configuration options for the connection pooler
	ConnectionPooler PgClusterConnectionPoolerPtrInput
	// Details about the network connection for your cluster.
	Connections PgClusterConnectionsPtrInput
	// The number of CPU cores per replica.
	Cores pulumi.IntInput
	// Credentials for the database user to be created.
	Credentials PgClusterCredentialsInput
	// The friendly name of your cluster.
	DisplayName pulumi.StringInput
	// Creates the cluster based on the existing backup.
	FromBackup PgClusterFromBackupPtrInput
	// The total number of instances in the cluster (one master and n-1 standbys)
	Instances pulumi.IntInput
	// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
	// be modified after datacenter creation (disallowed in update requests)
	Location pulumi.StringInput
	// a weekly 4 hour-long window, during which maintenance might occur
	MaintenanceWindow PgClusterMaintenanceWindowPtrInput
	// The PostgreSQL version of your cluster.
	PostgresVersion pulumi.StringInput
	// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
	Ram pulumi.IntInput
	// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
	StorageSize pulumi.IntInput
	// The storage type used in your cluster.
	StorageType pulumi.StringInput
	// Represents different modes of replication.
	SynchronizationMode pulumi.StringInput
}

func (PgClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgClusterArgs)(nil)).Elem()
}

type PgClusterInput interface {
	pulumi.Input

	ToPgClusterOutput() PgClusterOutput
	ToPgClusterOutputWithContext(ctx context.Context) PgClusterOutput
}

func (*PgCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**PgCluster)(nil)).Elem()
}

func (i *PgCluster) ToPgClusterOutput() PgClusterOutput {
	return i.ToPgClusterOutputWithContext(context.Background())
}

func (i *PgCluster) ToPgClusterOutputWithContext(ctx context.Context) PgClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgClusterOutput)
}

// PgClusterArrayInput is an input type that accepts PgClusterArray and PgClusterArrayOutput values.
// You can construct a concrete instance of `PgClusterArrayInput` via:
//
//	PgClusterArray{ PgClusterArgs{...} }
type PgClusterArrayInput interface {
	pulumi.Input

	ToPgClusterArrayOutput() PgClusterArrayOutput
	ToPgClusterArrayOutputWithContext(context.Context) PgClusterArrayOutput
}

type PgClusterArray []PgClusterInput

func (PgClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgCluster)(nil)).Elem()
}

func (i PgClusterArray) ToPgClusterArrayOutput() PgClusterArrayOutput {
	return i.ToPgClusterArrayOutputWithContext(context.Background())
}

func (i PgClusterArray) ToPgClusterArrayOutputWithContext(ctx context.Context) PgClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgClusterArrayOutput)
}

// PgClusterMapInput is an input type that accepts PgClusterMap and PgClusterMapOutput values.
// You can construct a concrete instance of `PgClusterMapInput` via:
//
//	PgClusterMap{ "key": PgClusterArgs{...} }
type PgClusterMapInput interface {
	pulumi.Input

	ToPgClusterMapOutput() PgClusterMapOutput
	ToPgClusterMapOutputWithContext(context.Context) PgClusterMapOutput
}

type PgClusterMap map[string]PgClusterInput

func (PgClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgCluster)(nil)).Elem()
}

func (i PgClusterMap) ToPgClusterMapOutput() PgClusterMapOutput {
	return i.ToPgClusterMapOutputWithContext(context.Background())
}

func (i PgClusterMap) ToPgClusterMapOutputWithContext(ctx context.Context) PgClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgClusterMapOutput)
}

type PgClusterOutput struct{ *pulumi.OutputState }

func (PgClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgCluster)(nil)).Elem()
}

func (o PgClusterOutput) ToPgClusterOutput() PgClusterOutput {
	return o
}

func (o PgClusterOutput) ToPgClusterOutputWithContext(ctx context.Context) PgClusterOutput {
	return o
}

// The S3 location where the backups will be stored.
func (o PgClusterOutput) BackupLocation() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.BackupLocation }).(pulumi.StringOutput)
}

// Configuration options for the connection pooler
func (o PgClusterOutput) ConnectionPooler() PgClusterConnectionPoolerOutput {
	return o.ApplyT(func(v *PgCluster) PgClusterConnectionPoolerOutput { return v.ConnectionPooler }).(PgClusterConnectionPoolerOutput)
}

// Details about the network connection for your cluster.
func (o PgClusterOutput) Connections() PgClusterConnectionsPtrOutput {
	return o.ApplyT(func(v *PgCluster) PgClusterConnectionsPtrOutput { return v.Connections }).(PgClusterConnectionsPtrOutput)
}

// The number of CPU cores per replica.
func (o PgClusterOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.IntOutput { return v.Cores }).(pulumi.IntOutput)
}

// Credentials for the database user to be created.
func (o PgClusterOutput) Credentials() PgClusterCredentialsOutput {
	return o.ApplyT(func(v *PgCluster) PgClusterCredentialsOutput { return v.Credentials }).(PgClusterCredentialsOutput)
}

// The friendly name of your cluster.
func (o PgClusterOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The DNS name pointing to your cluster
func (o PgClusterOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// Creates the cluster based on the existing backup.
func (o PgClusterOutput) FromBackup() PgClusterFromBackupPtrOutput {
	return o.ApplyT(func(v *PgCluster) PgClusterFromBackupPtrOutput { return v.FromBackup }).(PgClusterFromBackupPtrOutput)
}

// The total number of instances in the cluster (one master and n-1 standbys)
func (o PgClusterOutput) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.IntOutput { return v.Instances }).(pulumi.IntOutput)
}

// The physical location where the cluster will be created. This will be where all of your instances live. Property cannot
// be modified after datacenter creation (disallowed in update requests)
func (o PgClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// a weekly 4 hour-long window, during which maintenance might occur
func (o PgClusterOutput) MaintenanceWindow() PgClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *PgCluster) PgClusterMaintenanceWindowOutput { return v.MaintenanceWindow }).(PgClusterMaintenanceWindowOutput)
}

// The PostgreSQL version of your cluster.
func (o PgClusterOutput) PostgresVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.PostgresVersion }).(pulumi.StringOutput)
}

// The amount of memory per instance in megabytes. Has to be a multiple of 1024.
func (o PgClusterOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// The amount of storage per instance in megabytes. Has to be a multiple of 2048.
func (o PgClusterOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.IntOutput { return v.StorageSize }).(pulumi.IntOutput)
}

// The storage type used in your cluster.
func (o PgClusterOutput) StorageType() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.StorageType }).(pulumi.StringOutput)
}

// Represents different modes of replication.
func (o PgClusterOutput) SynchronizationMode() pulumi.StringOutput {
	return o.ApplyT(func(v *PgCluster) pulumi.StringOutput { return v.SynchronizationMode }).(pulumi.StringOutput)
}

type PgClusterArrayOutput struct{ *pulumi.OutputState }

func (PgClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgCluster)(nil)).Elem()
}

func (o PgClusterArrayOutput) ToPgClusterArrayOutput() PgClusterArrayOutput {
	return o
}

func (o PgClusterArrayOutput) ToPgClusterArrayOutputWithContext(ctx context.Context) PgClusterArrayOutput {
	return o
}

func (o PgClusterArrayOutput) Index(i pulumi.IntInput) PgClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgCluster {
		return vs[0].([]*PgCluster)[vs[1].(int)]
	}).(PgClusterOutput)
}

type PgClusterMapOutput struct{ *pulumi.OutputState }

func (PgClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgCluster)(nil)).Elem()
}

func (o PgClusterMapOutput) ToPgClusterMapOutput() PgClusterMapOutput {
	return o
}

func (o PgClusterMapOutput) ToPgClusterMapOutputWithContext(ctx context.Context) PgClusterMapOutput {
	return o
}

func (o PgClusterMapOutput) MapIndex(k pulumi.StringInput) PgClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgCluster {
		return vs[0].(map[string]*PgCluster)[vs[1].(string)]
	}).(PgClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgClusterInput)(nil)).Elem(), &PgCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgClusterArrayInput)(nil)).Elem(), PgClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgClusterMapInput)(nil)).Elem(), PgClusterMap{})
	pulumi.RegisterOutputType(PgClusterOutput{})
	pulumi.RegisterOutputType(PgClusterArrayOutput{})
	pulumi.RegisterOutputType(PgClusterMapOutput{})
}
