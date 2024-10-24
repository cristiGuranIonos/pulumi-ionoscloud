// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAutoCertificate(ctx *pulumi.Context, args *LookupAutoCertificateArgs, opts ...pulumi.InvokeOption) (*LookupAutoCertificateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoCertificateResult
	err := ctx.Invoke("ionoscloud:index/getAutoCertificate:getAutoCertificate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoCertificate.
type LookupAutoCertificateArgs struct {
	Id       *string `pulumi:"id"`
	Location string  `pulumi:"location"`
	Name     *string `pulumi:"name"`
}

// A collection of values returned by getAutoCertificate.
type LookupAutoCertificateResult struct {
	CommonName              string   `pulumi:"commonName"`
	Id                      *string  `pulumi:"id"`
	KeyAlgorithm            string   `pulumi:"keyAlgorithm"`
	LastIssuedCertificateId string   `pulumi:"lastIssuedCertificateId"`
	Location                string   `pulumi:"location"`
	Name                    *string  `pulumi:"name"`
	ProviderId              string   `pulumi:"providerId"`
	SubjectAlternativeNames []string `pulumi:"subjectAlternativeNames"`
}

func LookupAutoCertificateOutput(ctx *pulumi.Context, args LookupAutoCertificateOutputArgs, opts ...pulumi.InvokeOption) LookupAutoCertificateResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoCertificateResult, error) {
			args := v.(LookupAutoCertificateArgs)
			r, err := LookupAutoCertificate(ctx, &args, opts...)
			var s LookupAutoCertificateResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoCertificateResultOutput)
}

// A collection of arguments for invoking getAutoCertificate.
type LookupAutoCertificateOutputArgs struct {
	Id       pulumi.StringPtrInput `pulumi:"id"`
	Location pulumi.StringInput    `pulumi:"location"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAutoCertificateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoCertificateArgs)(nil)).Elem()
}

// A collection of values returned by getAutoCertificate.
type LookupAutoCertificateResultOutput struct{ *pulumi.OutputState }

func (LookupAutoCertificateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoCertificateResult)(nil)).Elem()
}

func (o LookupAutoCertificateResultOutput) ToLookupAutoCertificateResultOutput() LookupAutoCertificateResultOutput {
	return o
}

func (o LookupAutoCertificateResultOutput) ToLookupAutoCertificateResultOutputWithContext(ctx context.Context) LookupAutoCertificateResultOutput {
	return o
}

func (o LookupAutoCertificateResultOutput) CommonName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) string { return v.CommonName }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupAutoCertificateResultOutput) KeyAlgorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) string { return v.KeyAlgorithm }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateResultOutput) LastIssuedCertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) string { return v.LastIssuedCertificateId }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupAutoCertificateResultOutput) ProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) string { return v.ProviderId }).(pulumi.StringOutput)
}

func (o LookupAutoCertificateResultOutput) SubjectAlternativeNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAutoCertificateResult) []string { return v.SubjectAlternativeNames }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoCertificateResultOutput{})
}
