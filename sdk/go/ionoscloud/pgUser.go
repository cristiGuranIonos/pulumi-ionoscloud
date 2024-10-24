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

type PgUser struct {
	pulumi.CustomResourceState

	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Describes whether this user is a system user or not. A system user cannot be updated or deleted.
	IsSystemUser pulumi.BoolOutput   `pulumi:"isSystemUser"`
	Password     pulumi.StringOutput `pulumi:"password"`
	Username     pulumi.StringOutput `pulumi:"username"`
}

// NewPgUser registers a new resource with the given unique name, arguments, and options.
func NewPgUser(ctx *pulumi.Context,
	name string, args *PgUserArgs, opts ...pulumi.ResourceOption) (*PgUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource PgUser
	err := ctx.RegisterResource("ionoscloud:index/pgUser:PgUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPgUser gets an existing PgUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPgUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PgUserState, opts ...pulumi.ResourceOption) (*PgUser, error) {
	var resource PgUser
	err := ctx.ReadResource("ionoscloud:index/pgUser:PgUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PgUser resources.
type pgUserState struct {
	ClusterId *string `pulumi:"clusterId"`
	// Describes whether this user is a system user or not. A system user cannot be updated or deleted.
	IsSystemUser *bool   `pulumi:"isSystemUser"`
	Password     *string `pulumi:"password"`
	Username     *string `pulumi:"username"`
}

type PgUserState struct {
	ClusterId pulumi.StringPtrInput
	// Describes whether this user is a system user or not. A system user cannot be updated or deleted.
	IsSystemUser pulumi.BoolPtrInput
	Password     pulumi.StringPtrInput
	Username     pulumi.StringPtrInput
}

func (PgUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*pgUserState)(nil)).Elem()
}

type pgUserArgs struct {
	ClusterId string `pulumi:"clusterId"`
	Password  string `pulumi:"password"`
	Username  string `pulumi:"username"`
}

// The set of arguments for constructing a PgUser resource.
type PgUserArgs struct {
	ClusterId pulumi.StringInput
	Password  pulumi.StringInput
	Username  pulumi.StringInput
}

func (PgUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pgUserArgs)(nil)).Elem()
}

type PgUserInput interface {
	pulumi.Input

	ToPgUserOutput() PgUserOutput
	ToPgUserOutputWithContext(ctx context.Context) PgUserOutput
}

func (*PgUser) ElementType() reflect.Type {
	return reflect.TypeOf((**PgUser)(nil)).Elem()
}

func (i *PgUser) ToPgUserOutput() PgUserOutput {
	return i.ToPgUserOutputWithContext(context.Background())
}

func (i *PgUser) ToPgUserOutputWithContext(ctx context.Context) PgUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgUserOutput)
}

// PgUserArrayInput is an input type that accepts PgUserArray and PgUserArrayOutput values.
// You can construct a concrete instance of `PgUserArrayInput` via:
//
//	PgUserArray{ PgUserArgs{...} }
type PgUserArrayInput interface {
	pulumi.Input

	ToPgUserArrayOutput() PgUserArrayOutput
	ToPgUserArrayOutputWithContext(context.Context) PgUserArrayOutput
}

type PgUserArray []PgUserInput

func (PgUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgUser)(nil)).Elem()
}

func (i PgUserArray) ToPgUserArrayOutput() PgUserArrayOutput {
	return i.ToPgUserArrayOutputWithContext(context.Background())
}

func (i PgUserArray) ToPgUserArrayOutputWithContext(ctx context.Context) PgUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgUserArrayOutput)
}

// PgUserMapInput is an input type that accepts PgUserMap and PgUserMapOutput values.
// You can construct a concrete instance of `PgUserMapInput` via:
//
//	PgUserMap{ "key": PgUserArgs{...} }
type PgUserMapInput interface {
	pulumi.Input

	ToPgUserMapOutput() PgUserMapOutput
	ToPgUserMapOutputWithContext(context.Context) PgUserMapOutput
}

type PgUserMap map[string]PgUserInput

func (PgUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgUser)(nil)).Elem()
}

func (i PgUserMap) ToPgUserMapOutput() PgUserMapOutput {
	return i.ToPgUserMapOutputWithContext(context.Background())
}

func (i PgUserMap) ToPgUserMapOutputWithContext(ctx context.Context) PgUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PgUserMapOutput)
}

type PgUserOutput struct{ *pulumi.OutputState }

func (PgUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PgUser)(nil)).Elem()
}

func (o PgUserOutput) ToPgUserOutput() PgUserOutput {
	return o
}

func (o PgUserOutput) ToPgUserOutputWithContext(ctx context.Context) PgUserOutput {
	return o
}

func (o PgUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *PgUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Describes whether this user is a system user or not. A system user cannot be updated or deleted.
func (o PgUserOutput) IsSystemUser() pulumi.BoolOutput {
	return o.ApplyT(func(v *PgUser) pulumi.BoolOutput { return v.IsSystemUser }).(pulumi.BoolOutput)
}

func (o PgUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *PgUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

func (o PgUserOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *PgUser) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

type PgUserArrayOutput struct{ *pulumi.OutputState }

func (PgUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PgUser)(nil)).Elem()
}

func (o PgUserArrayOutput) ToPgUserArrayOutput() PgUserArrayOutput {
	return o
}

func (o PgUserArrayOutput) ToPgUserArrayOutputWithContext(ctx context.Context) PgUserArrayOutput {
	return o
}

func (o PgUserArrayOutput) Index(i pulumi.IntInput) PgUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PgUser {
		return vs[0].([]*PgUser)[vs[1].(int)]
	}).(PgUserOutput)
}

type PgUserMapOutput struct{ *pulumi.OutputState }

func (PgUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PgUser)(nil)).Elem()
}

func (o PgUserMapOutput) ToPgUserMapOutput() PgUserMapOutput {
	return o
}

func (o PgUserMapOutput) ToPgUserMapOutputWithContext(ctx context.Context) PgUserMapOutput {
	return o
}

func (o PgUserMapOutput) MapIndex(k pulumi.StringInput) PgUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PgUser {
		return vs[0].(map[string]*PgUser)[vs[1].(string)]
	}).(PgUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PgUserInput)(nil)).Elem(), &PgUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgUserArrayInput)(nil)).Elem(), PgUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PgUserMapInput)(nil)).Elem(), PgUserMap{})
	pulumi.RegisterOutputType(PgUserOutput{})
	pulumi.RegisterOutputType(PgUserArrayOutput{})
	pulumi.RegisterOutputType(PgUserMapOutput{})
}
