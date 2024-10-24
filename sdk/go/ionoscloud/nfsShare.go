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

// Creates and manages Network File Storage (NFS) Share objects on IonosCloud.
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
//			// Basic example
//			nfsDc, err := compute.NewDatacenter(ctx, "nfsDc", &compute.DatacenterArgs{
//				Location:          pulumi.String("de/txl"),
//				Description:       pulumi.String("Datacenter Description"),
//				SecAuthProtection: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			nfsLan, err := ionoscloud.NewLan(ctx, "nfsLan", &ionoscloud.LanArgs{
//				DatacenterId: nfsDc.ID(),
//				Public:       pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			exampleNfsCluster, err := ionoscloud.NewNfsCluster(ctx, "exampleNfsCluster", &ionoscloud.NfsClusterArgs{
//				Location: pulumi.String("de/txl"),
//				Size:     pulumi.Int(2),
//				Nfs: &ionoscloud.NfsClusterNfsArgs{
//					MinVersion: pulumi.String("4.2"),
//				},
//				Connections: &ionoscloud.NfsClusterConnectionsArgs{
//					DatacenterId: nfsDc.ID(),
//					IpAddress:    pulumi.String("192.168.100.10/24"),
//					Lan:          nfsLan.ID(),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ionoscloud.NewNfsShare(ctx, "exampleNfsShare", &ionoscloud.NfsShareArgs{
//				Location:  pulumi.String("de/txl"),
//				ClusterId: exampleNfsCluster.ID(),
//				Quota:     pulumi.Int(512),
//				Gid:       pulumi.Int(512),
//				Uid:       pulumi.Int(512),
//				ClientGroups: ionoscloud.NfsShareClientGroupArray{
//					&ionoscloud.NfsShareClientGroupArgs{
//						Description: pulumi.String("Client Group 1"),
//						IpNetworks: pulumi.StringArray{
//							pulumi.String("10.234.50.0/24"),
//						},
//						Hosts: pulumi.StringArray{
//							pulumi.String("10.234.62.123"),
//						},
//						Nfs: &ionoscloud.NfsShareClientGroupNfsArgs{
//							Squash: pulumi.String("all-anonymous"),
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
// A Network File Storage Share resource can be imported using its `location`, `cluster_id` and `resource id`:
//
// ```sh
// $ pulumi import ionoscloud:index/nfsShare:NfsShare name location:cluster_id:resource_id
// ```
type NfsShare struct {
	pulumi.CustomResourceState

	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	ClientGroups NfsShareClientGroupArrayOutput `pulumi:"clientGroups"`
	// The ID of the Network File Storage Cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Gid pulumi.IntPtrOutput `pulumi:"gid"`
	// The location of the Network File Storage Cluster.
	Location pulumi.StringOutput `pulumi:"location"`
	// The directory being exported.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to the NFS export. The NFS path is the path to the directory being exported.
	NfsPath pulumi.StringOutput `pulumi:"nfsPath"`
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
	Quota pulumi.IntPtrOutput `pulumi:"quota"`
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Uid pulumi.IntPtrOutput `pulumi:"uid"`
}

// NewNfsShare registers a new resource with the given unique name, arguments, and options.
func NewNfsShare(ctx *pulumi.Context,
	name string, args *NfsShareArgs, opts ...pulumi.ResourceOption) (*NfsShare, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClientGroups == nil {
		return nil, errors.New("invalid value for required argument 'ClientGroups'")
	}
	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NfsShare
	err := ctx.RegisterResource("ionoscloud:index/nfsShare:NfsShare", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNfsShare gets an existing NfsShare resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNfsShare(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NfsShareState, opts ...pulumi.ResourceOption) (*NfsShare, error) {
	var resource NfsShare
	err := ctx.ReadResource("ionoscloud:index/nfsShare:NfsShare", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NfsShare resources.
type nfsShareState struct {
	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	ClientGroups []NfsShareClientGroup `pulumi:"clientGroups"`
	// The ID of the Network File Storage Cluster.
	ClusterId *string `pulumi:"clusterId"`
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Gid *int `pulumi:"gid"`
	// The location of the Network File Storage Cluster.
	Location *string `pulumi:"location"`
	// The directory being exported.
	Name *string `pulumi:"name"`
	// Path to the NFS export. The NFS path is the path to the directory being exported.
	NfsPath *string `pulumi:"nfsPath"`
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
	Quota *int `pulumi:"quota"`
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Uid *int `pulumi:"uid"`
}

type NfsShareState struct {
	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	ClientGroups NfsShareClientGroupArrayInput
	// The ID of the Network File Storage Cluster.
	ClusterId pulumi.StringPtrInput
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Gid pulumi.IntPtrInput
	// The location of the Network File Storage Cluster.
	Location pulumi.StringPtrInput
	// The directory being exported.
	Name pulumi.StringPtrInput
	// Path to the NFS export. The NFS path is the path to the directory being exported.
	NfsPath pulumi.StringPtrInput
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
	Quota pulumi.IntPtrInput
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Uid pulumi.IntPtrInput
}

func (NfsShareState) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsShareState)(nil)).Elem()
}

type nfsShareArgs struct {
	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	ClientGroups []NfsShareClientGroup `pulumi:"clientGroups"`
	// The ID of the Network File Storage Cluster.
	ClusterId string `pulumi:"clusterId"`
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Gid *int `pulumi:"gid"`
	// The location of the Network File Storage Cluster.
	Location string `pulumi:"location"`
	// The directory being exported.
	Name *string `pulumi:"name"`
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
	Quota *int `pulumi:"quota"`
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Uid *int `pulumi:"uid"`
}

// The set of arguments for constructing a NfsShare resource.
type NfsShareArgs struct {
	// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
	ClientGroups NfsShareClientGroupArrayInput
	// The ID of the Network File Storage Cluster.
	ClusterId pulumi.StringInput
	// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Gid pulumi.IntPtrInput
	// The location of the Network File Storage Cluster.
	Location pulumi.StringInput
	// The directory being exported.
	Name pulumi.StringPtrInput
	// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
	Quota pulumi.IntPtrInput
	// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
	Uid pulumi.IntPtrInput
}

func (NfsShareArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsShareArgs)(nil)).Elem()
}

type NfsShareInput interface {
	pulumi.Input

	ToNfsShareOutput() NfsShareOutput
	ToNfsShareOutputWithContext(ctx context.Context) NfsShareOutput
}

func (*NfsShare) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsShare)(nil)).Elem()
}

func (i *NfsShare) ToNfsShareOutput() NfsShareOutput {
	return i.ToNfsShareOutputWithContext(context.Background())
}

func (i *NfsShare) ToNfsShareOutputWithContext(ctx context.Context) NfsShareOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsShareOutput)
}

// NfsShareArrayInput is an input type that accepts NfsShareArray and NfsShareArrayOutput values.
// You can construct a concrete instance of `NfsShareArrayInput` via:
//
//	NfsShareArray{ NfsShareArgs{...} }
type NfsShareArrayInput interface {
	pulumi.Input

	ToNfsShareArrayOutput() NfsShareArrayOutput
	ToNfsShareArrayOutputWithContext(context.Context) NfsShareArrayOutput
}

type NfsShareArray []NfsShareInput

func (NfsShareArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NfsShare)(nil)).Elem()
}

func (i NfsShareArray) ToNfsShareArrayOutput() NfsShareArrayOutput {
	return i.ToNfsShareArrayOutputWithContext(context.Background())
}

func (i NfsShareArray) ToNfsShareArrayOutputWithContext(ctx context.Context) NfsShareArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsShareArrayOutput)
}

// NfsShareMapInput is an input type that accepts NfsShareMap and NfsShareMapOutput values.
// You can construct a concrete instance of `NfsShareMapInput` via:
//
//	NfsShareMap{ "key": NfsShareArgs{...} }
type NfsShareMapInput interface {
	pulumi.Input

	ToNfsShareMapOutput() NfsShareMapOutput
	ToNfsShareMapOutputWithContext(context.Context) NfsShareMapOutput
}

type NfsShareMap map[string]NfsShareInput

func (NfsShareMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NfsShare)(nil)).Elem()
}

func (i NfsShareMap) ToNfsShareMapOutput() NfsShareMapOutput {
	return i.ToNfsShareMapOutputWithContext(context.Background())
}

func (i NfsShareMap) ToNfsShareMapOutputWithContext(ctx context.Context) NfsShareMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsShareMapOutput)
}

type NfsShareOutput struct{ *pulumi.OutputState }

func (NfsShareOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsShare)(nil)).Elem()
}

func (o NfsShareOutput) ToNfsShareOutput() NfsShareOutput {
	return o
}

func (o NfsShareOutput) ToNfsShareOutputWithContext(ctx context.Context) NfsShareOutput {
	return o
}

// The groups of clients are the systems connecting to the Network File Storage cluster. Each group includes:
func (o NfsShareOutput) ClientGroups() NfsShareClientGroupArrayOutput {
	return o.ApplyT(func(v *NfsShare) NfsShareClientGroupArrayOutput { return v.ClientGroups }).(NfsShareClientGroupArrayOutput)
}

// The ID of the Network File Storage Cluster.
func (o NfsShareOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The group ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
func (o NfsShareOutput) Gid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.IntPtrOutput { return v.Gid }).(pulumi.IntPtrOutput)
}

// The location of the Network File Storage Cluster.
func (o NfsShareOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The directory being exported.
func (o NfsShareOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Path to the NFS export. The NFS path is the path to the directory being exported.
func (o NfsShareOutput) NfsPath() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.StringOutput { return v.NfsPath }).(pulumi.StringOutput)
}

// The quota in MiB for the export. The quota can restrict the amount of data that can be stored within the export. The quota can be disabled using `0`. Default is `0`.
func (o NfsShareOutput) Quota() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.IntPtrOutput { return v.Quota }).(pulumi.IntPtrOutput)
}

// The user ID that will own the exported directory. If not set, **anonymous** (`512`) will be used.
func (o NfsShareOutput) Uid() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *NfsShare) pulumi.IntPtrOutput { return v.Uid }).(pulumi.IntPtrOutput)
}

type NfsShareArrayOutput struct{ *pulumi.OutputState }

func (NfsShareArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NfsShare)(nil)).Elem()
}

func (o NfsShareArrayOutput) ToNfsShareArrayOutput() NfsShareArrayOutput {
	return o
}

func (o NfsShareArrayOutput) ToNfsShareArrayOutputWithContext(ctx context.Context) NfsShareArrayOutput {
	return o
}

func (o NfsShareArrayOutput) Index(i pulumi.IntInput) NfsShareOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NfsShare {
		return vs[0].([]*NfsShare)[vs[1].(int)]
	}).(NfsShareOutput)
}

type NfsShareMapOutput struct{ *pulumi.OutputState }

func (NfsShareMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NfsShare)(nil)).Elem()
}

func (o NfsShareMapOutput) ToNfsShareMapOutput() NfsShareMapOutput {
	return o
}

func (o NfsShareMapOutput) ToNfsShareMapOutputWithContext(ctx context.Context) NfsShareMapOutput {
	return o
}

func (o NfsShareMapOutput) MapIndex(k pulumi.StringInput) NfsShareOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NfsShare {
		return vs[0].(map[string]*NfsShare)[vs[1].(string)]
	}).(NfsShareOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NfsShareInput)(nil)).Elem(), &NfsShare{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsShareArrayInput)(nil)).Elem(), NfsShareArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsShareMapInput)(nil)).Elem(), NfsShareMap{})
	pulumi.RegisterOutputType(NfsShareOutput{})
	pulumi.RegisterOutputType(NfsShareArrayOutput{})
	pulumi.RegisterOutputType(NfsShareMapOutput{})
}
