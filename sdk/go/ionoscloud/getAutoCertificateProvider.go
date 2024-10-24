// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutoCertificateProvider(ctx *pulumi.Context, args *LookupAutoCertificateProviderArgs, opts ...pulumi.InvokeOption) (*LookupAutoCertificateProviderResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoCertificateProviderResult
	err := ctx.Invoke("ionoscloud:index/getAutoCertificateProvider:getAutoCertificateProvider", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoCertificateProvider.
type LookupAutoCertificateProviderArgs struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     *string `pulumi:"name"`
}

// A collection of values returned by getAutoCertificateProvider.
type LookupAutoCertificateProviderResult struct {
	Email                   string                                             `pulumi:"email"`
	ExternalAccountBindings []GetAutoCertificateProviderExternalAccountBinding `pulumi:"externalAccountBindings"`
	Id                      *string                                            `pulumi:"id"`
	Location                string                                             `pulumi:"location"`
	Name                    *string                                            `pulumi:"name"`
	Server                  string                                             `pulumi:"server"`
}

func LookupAutoCertificateProviderOutput(ctx *pulumi.Context, args LookupAutoCertificateProviderOutputArgs, opts ...pulumi.InvokeOption) LookupAutoCertificateProviderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoCertificateProviderResult, error) {
			args := v.(LookupAutoCertificateProviderArgs)
			r, err := LookupAutoCertificateProvider(ctx, &args, opts...)
			var s LookupAutoCertificateProviderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoCertificateProviderResultOutput)
}

// A collection of arguments for invoking getAutoCertificateProvider.
type LookupAutoCertificateProviderOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAutoCertificateProviderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoCertificateProviderArgs)(nil)).Elem()
}

// A collection of values returned by getAutoCertificateProvider.
type LookupAutoCertificateProviderResultOutput struct{ *pulumi.OutputState }

func (LookupAutoCertificateProviderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoCertificateProviderResult)(nil)).Elem()
}

func (o LookupAutoCertificateProviderResultOutput) ToLookupAutoCertificateProviderResultOutput() LookupAutoCertificateProviderResultOutput {
	return o
}

func (o LookupAutoCertificateProviderResultOutput) ToLookupAutoCertificateProviderResultOutputWithContext(ctx context.Context) LookupAutoCertificateProviderResultOutput {
	return o
}

func (o LookupAutoCertificateProviderResultOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) string { return v.Email }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateProviderResultOutput) ExternalAccountBindings() GetAutoCertificateProviderExternalAccountBindingArrayOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) []GetAutoCertificateProviderExternalAccountBinding {
		return v.ExternalAccountBindings
	}).(GetAutoCertificateProviderExternalAccountBindingArrayOutput)
}

func (o LookupAutoCertificateProviderResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAutoCertificateProviderResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateProviderResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAutoCertificateProviderResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateProviderResult) string { return v.Server }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoCertificateProviderResultOutput{})
}