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

type AutoCertificateProvider struct {
	pulumi.CustomResourceState

	// The email address of the certificate requester
	Email                  pulumi.StringOutput                                    `pulumi:"email"`
	ExternalAccountBinding AutoCertificateProviderExternalAccountBindingPtrOutput `pulumi:"externalAccountBinding"`
	// The location of the certificate provider
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the certificate provider
	Name pulumi.StringOutput `pulumi:"name"`
	// The URL of the certificate provider
	Server pulumi.StringOutput `pulumi:"server"`
}

// NewAutoCertificateProvider registers a new resource with the given unique name, arguments, and options.
func NewAutoCertificateProvider(ctx *pulumi.Context,
	name string, args *AutoCertificateProviderArgs, opts ...pulumi.ResourceOption) (*AutoCertificateProvider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Email == nil {
		return nil, errors.New("invalid value for required argument 'Email'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Server == nil {
		return nil, errors.New("invalid value for required argument 'Server'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AutoCertificateProvider
	err := ctx.RegisterResource("ionoscloud:index/autoCertificateProvider:AutoCertificateProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAutoCertificateProvider gets an existing AutoCertificateProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAutoCertificateProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AutoCertificateProviderState, opts ...pulumi.ResourceOption) (*AutoCertificateProvider, error) {
	var resource AutoCertificateProvider
	err := ctx.ReadResource("ionoscloud:index/autoCertificateProvider:AutoCertificateProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AutoCertificateProvider resources.
type autoCertificateProviderState struct {
	// The email address of the certificate requester
	Email                  *string                                        `pulumi:"email"`
	ExternalAccountBinding *AutoCertificateProviderExternalAccountBinding `pulumi:"externalAccountBinding"`
	// The location of the certificate provider
	Location *string `pulumi:"location"`
	// The name of the certificate provider
	Name *string `pulumi:"name"`
	// The URL of the certificate provider
	Server *string `pulumi:"server"`
}

type AutoCertificateProviderState struct {
	// The email address of the certificate requester
	Email                  pulumi.StringPtrInput
	ExternalAccountBinding AutoCertificateProviderExternalAccountBindingPtrInput
	// The location of the certificate provider
	Location pulumi.StringPtrInput
	// The name of the certificate provider
	Name pulumi.StringPtrInput
	// The URL of the certificate provider
	Server pulumi.StringPtrInput
}

func (AutoCertificateProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*autoCertificateProviderState)(nil)).Elem()
}

type autoCertificateProviderArgs struct {
	// The email address of the certificate requester
	Email                  string                                         `pulumi:"email"`
	ExternalAccountBinding *AutoCertificateProviderExternalAccountBinding `pulumi:"externalAccountBinding"`
	// The location of the certificate provider
	Location string `pulumi:"location"`
	// The name of the certificate provider
	Name *string `pulumi:"name"`
	// The URL of the certificate provider
	Server string `pulumi:"server"`
}

// The set of arguments for constructing a AutoCertificateProvider resource.
type AutoCertificateProviderArgs struct {
	// The email address of the certificate requester
	Email                  pulumi.StringInput
	ExternalAccountBinding AutoCertificateProviderExternalAccountBindingPtrInput
	// The location of the certificate provider
	Location pulumi.StringInput
	// The name of the certificate provider
	Name pulumi.StringPtrInput
	// The URL of the certificate provider
	Server pulumi.StringInput
}

func (AutoCertificateProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*autoCertificateProviderArgs)(nil)).Elem()
}

type AutoCertificateProviderInput interface {
	pulumi.Input

	ToAutoCertificateProviderOutput() AutoCertificateProviderOutput
	ToAutoCertificateProviderOutputWithContext(ctx context.Context) AutoCertificateProviderOutput
}

func (*AutoCertificateProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoCertificateProvider)(nil)).Elem()
}

func (i *AutoCertificateProvider) ToAutoCertificateProviderOutput() AutoCertificateProviderOutput {
	return i.ToAutoCertificateProviderOutputWithContext(context.Background())
}

func (i *AutoCertificateProvider) ToAutoCertificateProviderOutputWithContext(ctx context.Context) AutoCertificateProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoCertificateProviderOutput)
}

// AutoCertificateProviderArrayInput is an input type that accepts AutoCertificateProviderArray and AutoCertificateProviderArrayOutput values.
// You can construct a concrete instance of `AutoCertificateProviderArrayInput` via:
//
//	AutoCertificateProviderArray{ AutoCertificateProviderArgs{...} }
type AutoCertificateProviderArrayInput interface {
	pulumi.Input

	ToAutoCertificateProviderArrayOutput() AutoCertificateProviderArrayOutput
	ToAutoCertificateProviderArrayOutputWithContext(context.Context) AutoCertificateProviderArrayOutput
}

type AutoCertificateProviderArray []AutoCertificateProviderInput

func (AutoCertificateProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoCertificateProvider)(nil)).Elem()
}

func (i AutoCertificateProviderArray) ToAutoCertificateProviderArrayOutput() AutoCertificateProviderArrayOutput {
	return i.ToAutoCertificateProviderArrayOutputWithContext(context.Background())
}

func (i AutoCertificateProviderArray) ToAutoCertificateProviderArrayOutputWithContext(ctx context.Context) AutoCertificateProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoCertificateProviderArrayOutput)
}

// AutoCertificateProviderMapInput is an input type that accepts AutoCertificateProviderMap and AutoCertificateProviderMapOutput values.
// You can construct a concrete instance of `AutoCertificateProviderMapInput` via:
//
//	AutoCertificateProviderMap{ "key": AutoCertificateProviderArgs{...} }
type AutoCertificateProviderMapInput interface {
	pulumi.Input

	ToAutoCertificateProviderMapOutput() AutoCertificateProviderMapOutput
	ToAutoCertificateProviderMapOutputWithContext(context.Context) AutoCertificateProviderMapOutput
}

type AutoCertificateProviderMap map[string]AutoCertificateProviderInput

func (AutoCertificateProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoCertificateProvider)(nil)).Elem()
}

func (i AutoCertificateProviderMap) ToAutoCertificateProviderMapOutput() AutoCertificateProviderMapOutput {
	return i.ToAutoCertificateProviderMapOutputWithContext(context.Background())
}

func (i AutoCertificateProviderMap) ToAutoCertificateProviderMapOutputWithContext(ctx context.Context) AutoCertificateProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoCertificateProviderMapOutput)
}

type AutoCertificateProviderOutput struct{ *pulumi.OutputState }

func (AutoCertificateProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoCertificateProvider)(nil)).Elem()
}

func (o AutoCertificateProviderOutput) ToAutoCertificateProviderOutput() AutoCertificateProviderOutput {
	return o
}

func (o AutoCertificateProviderOutput) ToAutoCertificateProviderOutputWithContext(ctx context.Context) AutoCertificateProviderOutput {
	return o
}

// The email address of the certificate requester
func (o AutoCertificateProviderOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoCertificateProvider) pulumi.StringOutput { return v.Email }).(pulumi.StringOutput)
}

func (o AutoCertificateProviderOutput) ExternalAccountBinding() AutoCertificateProviderExternalAccountBindingPtrOutput {
	return o.ApplyT(func(v *AutoCertificateProvider) AutoCertificateProviderExternalAccountBindingPtrOutput {
		return v.ExternalAccountBinding
	}).(AutoCertificateProviderExternalAccountBindingPtrOutput)
}

// The location of the certificate provider
func (o AutoCertificateProviderOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoCertificateProvider) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of the certificate provider
func (o AutoCertificateProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoCertificateProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The URL of the certificate provider
func (o AutoCertificateProviderOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v *AutoCertificateProvider) pulumi.StringOutput { return v.Server }).(pulumi.StringOutput)
}

type AutoCertificateProviderArrayOutput struct{ *pulumi.OutputState }

func (AutoCertificateProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AutoCertificateProvider)(nil)).Elem()
}

func (o AutoCertificateProviderArrayOutput) ToAutoCertificateProviderArrayOutput() AutoCertificateProviderArrayOutput {
	return o
}

func (o AutoCertificateProviderArrayOutput) ToAutoCertificateProviderArrayOutputWithContext(ctx context.Context) AutoCertificateProviderArrayOutput {
	return o
}

func (o AutoCertificateProviderArrayOutput) Index(i pulumi.IntInput) AutoCertificateProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AutoCertificateProvider {
		return vs[0].([]*AutoCertificateProvider)[vs[1].(int)]
	}).(AutoCertificateProviderOutput)
}

type AutoCertificateProviderMapOutput struct{ *pulumi.OutputState }

func (AutoCertificateProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AutoCertificateProvider)(nil)).Elem()
}

func (o AutoCertificateProviderMapOutput) ToAutoCertificateProviderMapOutput() AutoCertificateProviderMapOutput {
	return o
}

func (o AutoCertificateProviderMapOutput) ToAutoCertificateProviderMapOutputWithContext(ctx context.Context) AutoCertificateProviderMapOutput {
	return o
}

func (o AutoCertificateProviderMapOutput) MapIndex(k pulumi.StringInput) AutoCertificateProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AutoCertificateProvider {
		return vs[0].(map[string]*AutoCertificateProvider)[vs[1].(string)]
	}).(AutoCertificateProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AutoCertificateProviderInput)(nil)).Elem(), &AutoCertificateProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoCertificateProviderArrayInput)(nil)).Elem(), AutoCertificateProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutoCertificateProviderMapInput)(nil)).Elem(), AutoCertificateProviderMap{})
	pulumi.RegisterOutputType(AutoCertificateProviderOutput{})
	pulumi.RegisterOutputType(AutoCertificateProviderArrayOutput{})
	pulumi.RegisterOutputType(AutoCertificateProviderMapOutput{})
}
