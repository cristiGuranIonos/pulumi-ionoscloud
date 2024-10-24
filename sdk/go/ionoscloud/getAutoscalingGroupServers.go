// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The autoscaling group servers data source can be used to search for and return existing servers that are part of a specific autoscaling group.
func GetAutoscalingGroupServers(ctx *pulumi.Context, args *GetAutoscalingGroupServersArgs, opts ...pulumi.InvokeOption) (*GetAutoscalingGroupServersResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAutoscalingGroupServersResult
	err := ctx.Invoke("ionoscloud:index/getAutoscalingGroupServers:getAutoscalingGroupServers", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoscalingGroupServers.
type GetAutoscalingGroupServersArgs struct {
	// The unique ID of the autoscaling group.
	//
	// `groupId` must be provided. If it is not provided, the datasource will return an error.
	GroupId string `pulumi:"groupId"`
}

// A collection of values returned by getAutoscalingGroupServers.
type GetAutoscalingGroupServersResult struct {
	// Id of the autoscaling group.
	GroupId string `pulumi:"groupId"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// List of servers.
	Servers []GetAutoscalingGroupServersServer `pulumi:"servers"`
}

func GetAutoscalingGroupServersOutput(ctx *pulumi.Context, args GetAutoscalingGroupServersOutputArgs, opts ...pulumi.InvokeOption) GetAutoscalingGroupServersResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAutoscalingGroupServersResult, error) {
			args := v.(GetAutoscalingGroupServersArgs)
			r, err := GetAutoscalingGroupServers(ctx, &args, opts...)
			var s GetAutoscalingGroupServersResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAutoscalingGroupServersResultOutput)
}

// A collection of arguments for invoking getAutoscalingGroupServers.
type GetAutoscalingGroupServersOutputArgs struct {
	// The unique ID of the autoscaling group.
	//
	// `groupId` must be provided. If it is not provided, the datasource will return an error.
	GroupId pulumi.StringInput `pulumi:"groupId"`
}

func (GetAutoscalingGroupServersOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutoscalingGroupServersArgs)(nil)).Elem()
}

// A collection of values returned by getAutoscalingGroupServers.
type GetAutoscalingGroupServersResultOutput struct{ *pulumi.OutputState }

func (GetAutoscalingGroupServersResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAutoscalingGroupServersResult)(nil)).Elem()
}

func (o GetAutoscalingGroupServersResultOutput) ToGetAutoscalingGroupServersResultOutput() GetAutoscalingGroupServersResultOutput {
	return o
}

func (o GetAutoscalingGroupServersResultOutput) ToGetAutoscalingGroupServersResultOutputWithContext(ctx context.Context) GetAutoscalingGroupServersResultOutput {
	return o
}

// Id of the autoscaling group.
func (o GetAutoscalingGroupServersResultOutput) GroupId() pulumi.StringOutput {
	return o.ApplyT(func(v GetAutoscalingGroupServersResult) string { return v.GroupId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetAutoscalingGroupServersResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetAutoscalingGroupServersResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of servers.
func (o GetAutoscalingGroupServersResultOutput) Servers() GetAutoscalingGroupServersServerArrayOutput {
	return o.ApplyT(func(v GetAutoscalingGroupServersResult) []GetAutoscalingGroupServersServer { return v.Servers }).(GetAutoscalingGroupServersServerArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAutoscalingGroupServersResultOutput{})
}