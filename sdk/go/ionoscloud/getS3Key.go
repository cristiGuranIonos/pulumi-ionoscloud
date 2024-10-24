// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **S3 key data source** can be used to search for and return an existing s3 key.
// You can provide a string id which will be compared with provisioned s3 keys.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
func LookupS3Key(ctx *pulumi.Context, args *LookupS3KeyArgs, opts ...pulumi.InvokeOption) (*LookupS3KeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupS3KeyResult
	err := ctx.Invoke("ionoscloud:index/getS3Key:getS3Key", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getS3Key.
type LookupS3KeyArgs struct {
	// The state of the s3 key
	Active *bool `pulumi:"active"`
	// ID of the s3 key you want to search for.
	Id *string `pulumi:"id"`
	// [string] The UUID of the user owning the S3 Key.
	UserId string `pulumi:"userId"`
}

// A collection of values returned by getS3Key.
type LookupS3KeyResult struct {
	// The state of the s3 key
	Active *bool `pulumi:"active"`
	// The id of the s3 key
	Id *string `pulumi:"id"`
	// (Computed)The S3 Secret key.
	SecretKey string `pulumi:"secretKey"`
	// The ID of the user that owns the key
	UserId string `pulumi:"userId"`
}

func LookupS3KeyOutput(ctx *pulumi.Context, args LookupS3KeyOutputArgs, opts ...pulumi.InvokeOption) LookupS3KeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupS3KeyResult, error) {
			args := v.(LookupS3KeyArgs)
			r, err := LookupS3Key(ctx, &args, opts...)
			var s LookupS3KeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupS3KeyResultOutput)
}

// A collection of arguments for invoking getS3Key.
type LookupS3KeyOutputArgs struct {
	// The state of the s3 key
	Active pulumi.BoolPtrInput `pulumi:"active"`
	// ID of the s3 key you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// [string] The UUID of the user owning the S3 Key.
	UserId pulumi.StringInput `pulumi:"userId"`
}

func (LookupS3KeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupS3KeyArgs)(nil)).Elem()
}

// A collection of values returned by getS3Key.
type LookupS3KeyResultOutput struct{ *pulumi.OutputState }

func (LookupS3KeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupS3KeyResult)(nil)).Elem()
}

func (o LookupS3KeyResultOutput) ToLookupS3KeyResultOutput() LookupS3KeyResultOutput {
	return o
}

func (o LookupS3KeyResultOutput) ToLookupS3KeyResultOutputWithContext(ctx context.Context) LookupS3KeyResultOutput {
	return o
}

// The state of the s3 key
func (o LookupS3KeyResultOutput) Active() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupS3KeyResult) *bool { return v.Active }).(pulumi.BoolPtrOutput)
}

// The id of the s3 key
func (o LookupS3KeyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupS3KeyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// (Computed)The S3 Secret key.
func (o LookupS3KeyResultOutput) SecretKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupS3KeyResult) string { return v.SecretKey }).(pulumi.StringOutput)
}

// The ID of the user that owns the key
func (o LookupS3KeyResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupS3KeyResult) string { return v.UserId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupS3KeyResultOutput{})
}