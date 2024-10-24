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

type MariadbCluster struct {
	pulumi.CustomResourceState

	// The network connection for your cluster. Only one connection is allowed.
	Connections MariadbClusterConnectionsOutput `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores pulumi.IntOutput `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials MariadbClusterCredentialsOutput `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The DNS name pointing to your cluster.
	DnsName pulumi.StringOutput `pulumi:"dnsName"`
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances pulumi.IntOutput `pulumi:"instances"`
	// The cluster location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow MariadbClusterMaintenanceWindowOutput `pulumi:"maintenanceWindow"`
	// The MariaDB version of your cluster.
	MariadbVersion pulumi.StringOutput `pulumi:"mariadbVersion"`
	// The amount of memory per instance in gigabytes (GB).
	Ram pulumi.IntOutput `pulumi:"ram"`
	// The amount of storage per instance in gigabytes (GB).
	StorageSize pulumi.IntOutput `pulumi:"storageSize"`
}

// NewMariadbCluster registers a new resource with the given unique name, arguments, and options.
func NewMariadbCluster(ctx *pulumi.Context,
	name string, args *MariadbClusterArgs, opts ...pulumi.ResourceOption) (*MariadbCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connections == nil {
		return nil, errors.New("invalid value for required argument 'Connections'")
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
	if args.MariadbVersion == nil {
		return nil, errors.New("invalid value for required argument 'MariadbVersion'")
	}
	if args.Ram == nil {
		return nil, errors.New("invalid value for required argument 'Ram'")
	}
	if args.StorageSize == nil {
		return nil, errors.New("invalid value for required argument 'StorageSize'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MariadbCluster
	err := ctx.RegisterResource("ionoscloud:index/mariadbCluster:MariadbCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMariadbCluster gets an existing MariadbCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMariadbCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MariadbClusterState, opts ...pulumi.ResourceOption) (*MariadbCluster, error) {
	var resource MariadbCluster
	err := ctx.ReadResource("ionoscloud:index/mariadbCluster:MariadbCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MariadbCluster resources.
type mariadbClusterState struct {
	// The network connection for your cluster. Only one connection is allowed.
	Connections *MariadbClusterConnections `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores *int `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials *MariadbClusterCredentials `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName *string `pulumi:"displayName"`
	// The DNS name pointing to your cluster.
	DnsName *string `pulumi:"dnsName"`
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances *int `pulumi:"instances"`
	// The cluster location
	Location *string `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow *MariadbClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The MariaDB version of your cluster.
	MariadbVersion *string `pulumi:"mariadbVersion"`
	// The amount of memory per instance in gigabytes (GB).
	Ram *int `pulumi:"ram"`
	// The amount of storage per instance in gigabytes (GB).
	StorageSize *int `pulumi:"storageSize"`
}

type MariadbClusterState struct {
	// The network connection for your cluster. Only one connection is allowed.
	Connections MariadbClusterConnectionsPtrInput
	// The number of CPU cores per instance.
	Cores pulumi.IntPtrInput
	// Credentials for the database user to be created.
	Credentials MariadbClusterCredentialsPtrInput
	// The friendly name of your cluster.
	DisplayName pulumi.StringPtrInput
	// The DNS name pointing to your cluster.
	DnsName pulumi.StringPtrInput
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances pulumi.IntPtrInput
	// The cluster location
	Location pulumi.StringPtrInput
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow MariadbClusterMaintenanceWindowPtrInput
	// The MariaDB version of your cluster.
	MariadbVersion pulumi.StringPtrInput
	// The amount of memory per instance in gigabytes (GB).
	Ram pulumi.IntPtrInput
	// The amount of storage per instance in gigabytes (GB).
	StorageSize pulumi.IntPtrInput
}

func (MariadbClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*mariadbClusterState)(nil)).Elem()
}

type mariadbClusterArgs struct {
	// The network connection for your cluster. Only one connection is allowed.
	Connections MariadbClusterConnections `pulumi:"connections"`
	// The number of CPU cores per instance.
	Cores int `pulumi:"cores"`
	// Credentials for the database user to be created.
	Credentials MariadbClusterCredentials `pulumi:"credentials"`
	// The friendly name of your cluster.
	DisplayName string `pulumi:"displayName"`
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances int `pulumi:"instances"`
	// The cluster location
	Location *string `pulumi:"location"`
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow *MariadbClusterMaintenanceWindow `pulumi:"maintenanceWindow"`
	// The MariaDB version of your cluster.
	MariadbVersion string `pulumi:"mariadbVersion"`
	// The amount of memory per instance in gigabytes (GB).
	Ram int `pulumi:"ram"`
	// The amount of storage per instance in gigabytes (GB).
	StorageSize int `pulumi:"storageSize"`
}

// The set of arguments for constructing a MariadbCluster resource.
type MariadbClusterArgs struct {
	// The network connection for your cluster. Only one connection is allowed.
	Connections MariadbClusterConnectionsInput
	// The number of CPU cores per instance.
	Cores pulumi.IntInput
	// Credentials for the database user to be created.
	Credentials MariadbClusterCredentialsInput
	// The friendly name of your cluster.
	DisplayName pulumi.StringInput
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances pulumi.IntInput
	// The cluster location
	Location pulumi.StringPtrInput
	// A weekly 4 hour-long window, during which maintenance might occur.
	MaintenanceWindow MariadbClusterMaintenanceWindowPtrInput
	// The MariaDB version of your cluster.
	MariadbVersion pulumi.StringInput
	// The amount of memory per instance in gigabytes (GB).
	Ram pulumi.IntInput
	// The amount of storage per instance in gigabytes (GB).
	StorageSize pulumi.IntInput
}

func (MariadbClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mariadbClusterArgs)(nil)).Elem()
}

type MariadbClusterInput interface {
	pulumi.Input

	ToMariadbClusterOutput() MariadbClusterOutput
	ToMariadbClusterOutputWithContext(ctx context.Context) MariadbClusterOutput
}

func (*MariadbCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**MariadbCluster)(nil)).Elem()
}

func (i *MariadbCluster) ToMariadbClusterOutput() MariadbClusterOutput {
	return i.ToMariadbClusterOutputWithContext(context.Background())
}

func (i *MariadbCluster) ToMariadbClusterOutputWithContext(ctx context.Context) MariadbClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MariadbClusterOutput)
}

// MariadbClusterArrayInput is an input type that accepts MariadbClusterArray and MariadbClusterArrayOutput values.
// You can construct a concrete instance of `MariadbClusterArrayInput` via:
//
//	MariadbClusterArray{ MariadbClusterArgs{...} }
type MariadbClusterArrayInput interface {
	pulumi.Input

	ToMariadbClusterArrayOutput() MariadbClusterArrayOutput
	ToMariadbClusterArrayOutputWithContext(context.Context) MariadbClusterArrayOutput
}

type MariadbClusterArray []MariadbClusterInput

func (MariadbClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MariadbCluster)(nil)).Elem()
}

func (i MariadbClusterArray) ToMariadbClusterArrayOutput() MariadbClusterArrayOutput {
	return i.ToMariadbClusterArrayOutputWithContext(context.Background())
}

func (i MariadbClusterArray) ToMariadbClusterArrayOutputWithContext(ctx context.Context) MariadbClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MariadbClusterArrayOutput)
}

// MariadbClusterMapInput is an input type that accepts MariadbClusterMap and MariadbClusterMapOutput values.
// You can construct a concrete instance of `MariadbClusterMapInput` via:
//
//	MariadbClusterMap{ "key": MariadbClusterArgs{...} }
type MariadbClusterMapInput interface {
	pulumi.Input

	ToMariadbClusterMapOutput() MariadbClusterMapOutput
	ToMariadbClusterMapOutputWithContext(context.Context) MariadbClusterMapOutput
}

type MariadbClusterMap map[string]MariadbClusterInput

func (MariadbClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MariadbCluster)(nil)).Elem()
}

func (i MariadbClusterMap) ToMariadbClusterMapOutput() MariadbClusterMapOutput {
	return i.ToMariadbClusterMapOutputWithContext(context.Background())
}

func (i MariadbClusterMap) ToMariadbClusterMapOutputWithContext(ctx context.Context) MariadbClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MariadbClusterMapOutput)
}

type MariadbClusterOutput struct{ *pulumi.OutputState }

func (MariadbClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MariadbCluster)(nil)).Elem()
}

func (o MariadbClusterOutput) ToMariadbClusterOutput() MariadbClusterOutput {
	return o
}

func (o MariadbClusterOutput) ToMariadbClusterOutputWithContext(ctx context.Context) MariadbClusterOutput {
	return o
}

// The network connection for your cluster. Only one connection is allowed.
func (o MariadbClusterOutput) Connections() MariadbClusterConnectionsOutput {
	return o.ApplyT(func(v *MariadbCluster) MariadbClusterConnectionsOutput { return v.Connections }).(MariadbClusterConnectionsOutput)
}

// The number of CPU cores per instance.
func (o MariadbClusterOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.IntOutput { return v.Cores }).(pulumi.IntOutput)
}

// Credentials for the database user to be created.
func (o MariadbClusterOutput) Credentials() MariadbClusterCredentialsOutput {
	return o.ApplyT(func(v *MariadbCluster) MariadbClusterCredentialsOutput { return v.Credentials }).(MariadbClusterCredentialsOutput)
}

// The friendly name of your cluster.
func (o MariadbClusterOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The DNS name pointing to your cluster.
func (o MariadbClusterOutput) DnsName() pulumi.StringOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.StringOutput { return v.DnsName }).(pulumi.StringOutput)
}

// The total number of instances in the cluster (one primary and n-1 secondary).
func (o MariadbClusterOutput) Instances() pulumi.IntOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.IntOutput { return v.Instances }).(pulumi.IntOutput)
}

// The cluster location
func (o MariadbClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// A weekly 4 hour-long window, during which maintenance might occur.
func (o MariadbClusterOutput) MaintenanceWindow() MariadbClusterMaintenanceWindowOutput {
	return o.ApplyT(func(v *MariadbCluster) MariadbClusterMaintenanceWindowOutput { return v.MaintenanceWindow }).(MariadbClusterMaintenanceWindowOutput)
}

// The MariaDB version of your cluster.
func (o MariadbClusterOutput) MariadbVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.StringOutput { return v.MariadbVersion }).(pulumi.StringOutput)
}

// The amount of memory per instance in gigabytes (GB).
func (o MariadbClusterOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.IntOutput { return v.Ram }).(pulumi.IntOutput)
}

// The amount of storage per instance in gigabytes (GB).
func (o MariadbClusterOutput) StorageSize() pulumi.IntOutput {
	return o.ApplyT(func(v *MariadbCluster) pulumi.IntOutput { return v.StorageSize }).(pulumi.IntOutput)
}

type MariadbClusterArrayOutput struct{ *pulumi.OutputState }

func (MariadbClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MariadbCluster)(nil)).Elem()
}

func (o MariadbClusterArrayOutput) ToMariadbClusterArrayOutput() MariadbClusterArrayOutput {
	return o
}

func (o MariadbClusterArrayOutput) ToMariadbClusterArrayOutputWithContext(ctx context.Context) MariadbClusterArrayOutput {
	return o
}

func (o MariadbClusterArrayOutput) Index(i pulumi.IntInput) MariadbClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MariadbCluster {
		return vs[0].([]*MariadbCluster)[vs[1].(int)]
	}).(MariadbClusterOutput)
}

type MariadbClusterMapOutput struct{ *pulumi.OutputState }

func (MariadbClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MariadbCluster)(nil)).Elem()
}

func (o MariadbClusterMapOutput) ToMariadbClusterMapOutput() MariadbClusterMapOutput {
	return o
}

func (o MariadbClusterMapOutput) ToMariadbClusterMapOutputWithContext(ctx context.Context) MariadbClusterMapOutput {
	return o
}

func (o MariadbClusterMapOutput) MapIndex(k pulumi.StringInput) MariadbClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MariadbCluster {
		return vs[0].(map[string]*MariadbCluster)[vs[1].(string)]
	}).(MariadbClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MariadbClusterInput)(nil)).Elem(), &MariadbCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*MariadbClusterArrayInput)(nil)).Elem(), MariadbClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MariadbClusterMapInput)(nil)).Elem(), MariadbClusterMap{})
	pulumi.RegisterOutputType(MariadbClusterOutput{})
	pulumi.RegisterOutputType(MariadbClusterArrayOutput{})
	pulumi.RegisterOutputType(MariadbClusterMapOutput{})
}