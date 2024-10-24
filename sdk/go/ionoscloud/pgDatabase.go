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

type PgDatabase struct {
	pulumi.CustomResourceState

	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// The databasename of a given database.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of the role owning a given database.
	Owner pulumi.StringOutput `pulumi:"owner"`
}

// NewPgDatabase registers a new resource with the given unique name, arguments, and options.
func NewPgDatabase(ctx *pulumi.Context,
	name string, args *PgDatabaseArgs, opts ...pulumi.ResourceOption) (*PgDatabase, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Owner == nil {
		return nil, errors.New("invalid value for required argument 'Owner'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PgDatabase
	err := ctx.RegisterResource("ionoscloud:index/pgDatabase:PgDatabase", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgDatabase gets an existing PgDatabase resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgDatabase(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgDatabaseState, opts ...pulumi.ResourceOption) (*PgDatabase, error) {
	var resource PgDatabase
	err := ctx.ReadResource("ionoscloud:index/pgDatabase:PgDatabase", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgDatabase resources.
type pgDatabaseState struct {
	ClusterId *string `pulumi:"clusterId"`
	// The databasename of a given database.
	Name *string `pulumi:"name"`
	// The name of the role owning a given database.
	Owner *string `pulumi:"owner"`
}

type PgDatabaseState struct {
	ClusterId pulumi.StringPtrInput
	// The databasename of a given database.
	Name pulumi.StringPtrInput
	// The name of the role owning a given database.
	Owner pulumi.StringPtrInput
}

func (PgDatabaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgDatabaseState)(nil)).Elem()
}

type pgDatabaseArgs struct {
	ClusterId string `pulumi:"clusterId"`
	// The databasename of a given database.
	Name *string `pulumi:"name"`
	// The name of the role owning a given database.
	Owner string `pulumi:"owner"`
}

// The set of arguments for constructing a PgDatabase resource.
type PgDatabaseArgs struct {
	ClusterId pulumi.StringInput
	// The databasename of a given database.
	Name pulumi.StringPtrInput
	// The name of the role owning a given database.
	Owner pulumi.StringInput
}

func (PgDatabaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgDatabaseArgs)(nil)).Elem()
}

type PgDatabaseInput interface {
	pulumi.Input

	ToPgDatabaseOutput() PgDatabaseOutput
	ToPgDatabaseOutputWithContext(ctx context.Context) PgDatabaseOutput
}

func (*PgDatabase) ElementType() reflect.Type {
	return reflect.TypeOf((**PgDatabase)(nil)).Elem()
}

func (i *PgDatabase) ToPgDatabaseOutput() PgDatabaseOutput {
	return i.ToPgDatabaseOutputWithContext(context.Background())
}

func (i *PgDatabase) ToPgDatabaseOutputWithContext(ctx context.Context) PgDatabaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgDatabaseOutput)
}

// PgDatabaseArrayInput is an input type that accepts PgDatabaseArray and PgDatabaseArrayOutput values.
// You can construct a concrete instance of `PgDatabaseArrayInput` via:
//
//	PgDatabaseArray{ PgDatabaseArgs{...} }
type PgDatabaseArrayInput interface {
	pulumi.Input

	ToPgDatabaseArrayOutput() PgDatabaseArrayOutput
	ToPgDatabaseArrayOutputWithContext(context.Context) PgDatabaseArrayOutput
}

type PgDatabaseArray []PgDatabaseInput

func (PgDatabaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgDatabase)(nil)).Elem()
}

func (i PgDatabaseArray) ToPgDatabaseArrayOutput() PgDatabaseArrayOutput {
	return i.ToPgDatabaseArrayOutputWithContext(context.Background())
}

func (i PgDatabaseArray) ToPgDatabaseArrayOutputWithContext(ctx context.Context) PgDatabaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgDatabaseArrayOutput)
}

// PgDatabaseMapInput is an input type that accepts PgDatabaseMap and PgDatabaseMapOutput values.
// You can construct a concrete instance of `PgDatabaseMapInput` via:
//
//	PgDatabaseMap{ "key": PgDatabaseArgs{...} }
type PgDatabaseMapInput interface {
	pulumi.Input

	ToPgDatabaseMapOutput() PgDatabaseMapOutput
	ToPgDatabaseMapOutputWithContext(context.Context) PgDatabaseMapOutput
}

type PgDatabaseMap map[string]PgDatabaseInput

func (PgDatabaseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgDatabase)(nil)).Elem()
}

func (i PgDatabaseMap) ToPgDatabaseMapOutput() PgDatabaseMapOutput {
	return i.ToPgDatabaseMapOutputWithContext(context.Background())
}

func (i PgDatabaseMap) ToPgDatabaseMapOutputWithContext(ctx context.Context) PgDatabaseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgDatabaseMapOutput)
}

type PgDatabaseOutput struct{ *pulumi.OutputState }

func (PgDatabaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgDatabase)(nil)).Elem()
}

func (o PgDatabaseOutput) ToPgDatabaseOutput() PgDatabaseOutput {
	return o
}

func (o PgDatabaseOutput) ToPgDatabaseOutputWithContext(ctx context.Context) PgDatabaseOutput {
	return o
}

func (o PgDatabaseOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *PgDatabase) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// The databasename of a given database.
func (o PgDatabaseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PgDatabase) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of the role owning a given database.
func (o PgDatabaseOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v *PgDatabase) pulumi.StringOutput { return v.Owner }).(pulumi.StringOutput)
}

type PgDatabaseArrayOutput struct{ *pulumi.OutputState }

func (PgDatabaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgDatabase)(nil)).Elem()
}

func (o PgDatabaseArrayOutput) ToPgDatabaseArrayOutput() PgDatabaseArrayOutput {
	return o
}

func (o PgDatabaseArrayOutput) ToPgDatabaseArrayOutputWithContext(ctx context.Context) PgDatabaseArrayOutput {
	return o
}

func (o PgDatabaseArrayOutput) Index(i pulumi.IntInput) PgDatabaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgDatabase {
		return vs[0].([]*PgDatabase)[vs[1].(int)]
	}).(PgDatabaseOutput)
}

type PgDatabaseMapOutput struct{ *pulumi.OutputState }

func (PgDatabaseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgDatabase)(nil)).Elem()
}

func (o PgDatabaseMapOutput) ToPgDatabaseMapOutput() PgDatabaseMapOutput {
	return o
}

func (o PgDatabaseMapOutput) ToPgDatabaseMapOutputWithContext(ctx context.Context) PgDatabaseMapOutput {
	return o
}

func (o PgDatabaseMapOutput) MapIndex(k pulumi.StringInput) PgDatabaseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgDatabase {
		return vs[0].(map[string]*PgDatabase)[vs[1].(string)]
	}).(PgDatabaseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgDatabaseInput)(nil)).Elem(), &PgDatabase{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgDatabaseArrayInput)(nil)).Elem(), PgDatabaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgDatabaseMapInput)(nil)).Elem(), PgDatabaseMap{})
	pulumi.RegisterOutputType(PgDatabaseOutput{})
	pulumi.RegisterOutputType(PgDatabaseArrayOutput{})
	pulumi.RegisterOutputType(PgDatabaseMapOutput{})
}
