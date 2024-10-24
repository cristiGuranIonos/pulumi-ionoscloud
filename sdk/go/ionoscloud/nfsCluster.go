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

// Create clusters of Network File Storage (NFS) on IonosCloud.
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
//			_, err = ionoscloud.NewNfsCluster(ctx, "example", &ionoscloud.NfsClusterArgs{
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
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// A Network File Storage Cluster resource can be imported using its `location` and `resource id`:
//
// ```sh
// $ pulumi import ionoscloud:index/nfsCluster:NfsCluster name {location}:{uuid}
// ```
type NfsCluster struct {
	pulumi.CustomResourceState

	// The network connections for the Network File Storage Cluster.
	Connections NfsClusterConnectionsOutput `pulumi:"connections"`
	// The location where the Network File Storage cluster is located.
	// - `de/fra` - Frankfurt
	// - `de/txl` - Berlin
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the Network File Storage cluster.
	Name pulumi.StringOutput    `pulumi:"name"`
	Nfs  NfsClusterNfsPtrOutput `pulumi:"nfs"`
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
	Size pulumi.IntOutput `pulumi:"size"`
}

// NewNfsCluster registers a new resource with the given unique name, arguments, and options.
func NewNfsCluster(ctx *pulumi.Context,
	name string, args *NfsClusterArgs, opts ...pulumi.ResourceOption) (*NfsCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Connections == nil {
		return nil, errors.New("invalid value for required argument 'Connections'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NfsCluster
	err := ctx.RegisterResource("ionoscloud:index/nfsCluster:NfsCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNfsCluster gets an existing NfsCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNfsCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NfsClusterState, opts ...pulumi.ResourceOption) (*NfsCluster, error) {
	var resource NfsCluster
	err := ctx.ReadResource("ionoscloud:index/nfsCluster:NfsCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NfsCluster resources.
type nfsClusterState struct {
	// The network connections for the Network File Storage Cluster.
	Connections *NfsClusterConnections `pulumi:"connections"`
	// The location where the Network File Storage cluster is located.
	// - `de/fra` - Frankfurt
	// - `de/txl` - Berlin
	Location *string `pulumi:"location"`
	// The name of the Network File Storage cluster.
	Name *string        `pulumi:"name"`
	Nfs  *NfsClusterNfs `pulumi:"nfs"`
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
	Size *int `pulumi:"size"`
}

type NfsClusterState struct {
	// The network connections for the Network File Storage Cluster.
	Connections NfsClusterConnectionsPtrInput
	// The location where the Network File Storage cluster is located.
	// - `de/fra` - Frankfurt
	// - `de/txl` - Berlin
	Location pulumi.StringPtrInput
	// The name of the Network File Storage cluster.
	Name pulumi.StringPtrInput
	Nfs  NfsClusterNfsPtrInput
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
	Size pulumi.IntPtrInput
}

func (NfsClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsClusterState)(nil)).Elem()
}

type nfsClusterArgs struct {
	// The network connections for the Network File Storage Cluster.
	Connections NfsClusterConnections `pulumi:"connections"`
	// The location where the Network File Storage cluster is located.
	// - `de/fra` - Frankfurt
	// - `de/txl` - Berlin
	Location string `pulumi:"location"`
	// The name of the Network File Storage cluster.
	Name *string        `pulumi:"name"`
	Nfs  *NfsClusterNfs `pulumi:"nfs"`
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
	Size int `pulumi:"size"`
}

// The set of arguments for constructing a NfsCluster resource.
type NfsClusterArgs struct {
	// The network connections for the Network File Storage Cluster.
	Connections NfsClusterConnectionsInput
	// The location where the Network File Storage cluster is located.
	// - `de/fra` - Frankfurt
	// - `de/txl` - Berlin
	Location pulumi.StringInput
	// The name of the Network File Storage cluster.
	Name pulumi.StringPtrInput
	Nfs  NfsClusterNfsPtrInput
	// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
	Size pulumi.IntInput
}

func (NfsClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nfsClusterArgs)(nil)).Elem()
}

type NfsClusterInput interface {
	pulumi.Input

	ToNfsClusterOutput() NfsClusterOutput
	ToNfsClusterOutputWithContext(ctx context.Context) NfsClusterOutput
}

func (*NfsCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsCluster)(nil)).Elem()
}

func (i *NfsCluster) ToNfsClusterOutput() NfsClusterOutput {
	return i.ToNfsClusterOutputWithContext(context.Background())
}

func (i *NfsCluster) ToNfsClusterOutputWithContext(ctx context.Context) NfsClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsClusterOutput)
}

// NfsClusterArrayInput is an input type that accepts NfsClusterArray and NfsClusterArrayOutput values.
// You can construct a concrete instance of `NfsClusterArrayInput` via:
//
//	NfsClusterArray{ NfsClusterArgs{...} }
type NfsClusterArrayInput interface {
	pulumi.Input

	ToNfsClusterArrayOutput() NfsClusterArrayOutput
	ToNfsClusterArrayOutputWithContext(context.Context) NfsClusterArrayOutput
}

type NfsClusterArray []NfsClusterInput

func (NfsClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NfsCluster)(nil)).Elem()
}

func (i NfsClusterArray) ToNfsClusterArrayOutput() NfsClusterArrayOutput {
	return i.ToNfsClusterArrayOutputWithContext(context.Background())
}

func (i NfsClusterArray) ToNfsClusterArrayOutputWithContext(ctx context.Context) NfsClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsClusterArrayOutput)
}

// NfsClusterMapInput is an input type that accepts NfsClusterMap and NfsClusterMapOutput values.
// You can construct a concrete instance of `NfsClusterMapInput` via:
//
//	NfsClusterMap{ "key": NfsClusterArgs{...} }
type NfsClusterMapInput interface {
	pulumi.Input

	ToNfsClusterMapOutput() NfsClusterMapOutput
	ToNfsClusterMapOutputWithContext(context.Context) NfsClusterMapOutput
}

type NfsClusterMap map[string]NfsClusterInput

func (NfsClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NfsCluster)(nil)).Elem()
}

func (i NfsClusterMap) ToNfsClusterMapOutput() NfsClusterMapOutput {
	return i.ToNfsClusterMapOutputWithContext(context.Background())
}

func (i NfsClusterMap) ToNfsClusterMapOutputWithContext(ctx context.Context) NfsClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NfsClusterMapOutput)
}

type NfsClusterOutput struct{ *pulumi.OutputState }

func (NfsClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NfsCluster)(nil)).Elem()
}

func (o NfsClusterOutput) ToNfsClusterOutput() NfsClusterOutput {
	return o
}

func (o NfsClusterOutput) ToNfsClusterOutputWithContext(ctx context.Context) NfsClusterOutput {
	return o
}

// The network connections for the Network File Storage Cluster.
func (o NfsClusterOutput) Connections() NfsClusterConnectionsOutput {
	return o.ApplyT(func(v *NfsCluster) NfsClusterConnectionsOutput { return v.Connections }).(NfsClusterConnectionsOutput)
}

// The location where the Network File Storage cluster is located.
// - `de/fra` - Frankfurt
// - `de/txl` - Berlin
func (o NfsClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the Network File Storage cluster.
func (o NfsClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NfsCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o NfsClusterOutput) Nfs() NfsClusterNfsPtrOutput {
	return o.ApplyT(func(v *NfsCluster) NfsClusterNfsPtrOutput { return v.Nfs }).(NfsClusterNfsPtrOutput)
}

// The size of the Network File Storage cluster in TiB. Note that the cluster size cannot be reduced after provisioning. This value determines the billing fees. Default is `2`. The minimum value is `2` and the maximum value is `42`.
func (o NfsClusterOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v *NfsCluster) pulumi.IntOutput { return v.Size }).(pulumi.IntOutput)
}

type NfsClusterArrayOutput struct{ *pulumi.OutputState }

func (NfsClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NfsCluster)(nil)).Elem()
}

func (o NfsClusterArrayOutput) ToNfsClusterArrayOutput() NfsClusterArrayOutput {
	return o
}

func (o NfsClusterArrayOutput) ToNfsClusterArrayOutputWithContext(ctx context.Context) NfsClusterArrayOutput {
	return o
}

func (o NfsClusterArrayOutput) Index(i pulumi.IntInput) NfsClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NfsCluster {
		return vs[0].([]*NfsCluster)[vs[1].(int)]
	}).(NfsClusterOutput)
}

type NfsClusterMapOutput struct{ *pulumi.OutputState }

func (NfsClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NfsCluster)(nil)).Elem()
}

func (o NfsClusterMapOutput) ToNfsClusterMapOutput() NfsClusterMapOutput {
	return o
}

func (o NfsClusterMapOutput) ToNfsClusterMapOutputWithContext(ctx context.Context) NfsClusterMapOutput {
	return o
}

func (o NfsClusterMapOutput) MapIndex(k pulumi.StringInput) NfsClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NfsCluster {
		return vs[0].(map[string]*NfsCluster)[vs[1].(string)]
	}).(NfsClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NfsClusterInput)(nil)).Elem(), &NfsCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsClusterArrayInput)(nil)).Elem(), NfsClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NfsClusterMapInput)(nil)).Elem(), NfsClusterMap{})
	pulumi.RegisterOutputType(NfsClusterOutput{})
	pulumi.RegisterOutputType(NfsClusterArrayOutput{})
	pulumi.RegisterOutputType(NfsClusterMapOutput{})
}
