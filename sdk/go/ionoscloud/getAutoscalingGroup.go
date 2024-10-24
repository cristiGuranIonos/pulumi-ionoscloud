// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The autoscaling group data source can be used to search for and return an existing Autoscaling Group. You can provide a string for the name or id parameters which will be compared with provisioned Autoscaling Groups. If a single match is found, it will be returned.
//
// ## Example Usage
//
// ### By Name
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ionoscloud.LookupAutoscalingGroup(ctx, &ionoscloud.LookupAutoscalingGroupArgs{
//				Name: pulumi.StringRef("test_ds"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
func LookupAutoscalingGroup(ctx *pulumi.Context, args *LookupAutoscalingGroupArgs, opts ...pulumi.InvokeOption) (*LookupAutoscalingGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoscalingGroupResult
	err := ctx.Invoke("ionoscloud:index/getAutoscalingGroup:getAutoscalingGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAutoscalingGroup.
type LookupAutoscalingGroupArgs struct {
	// Id of an existing Autoscaling Group that you want to search for.
	Id *string `pulumi:"id"`
	// Name of an existing Autoscaling Group that you want to search for.
	//
	// Either `name` or `id` must be provided. If none or both are provided, the datasource will return an error.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAutoscalingGroup.
type LookupAutoscalingGroupResult struct {
	DatacenterId string `pulumi:"datacenterId"`
	// Unique identifier for the resource
	Id *string `pulumi:"id"`
	// Location of the datacenter. This location is the same as the one from the selected template.
	Location string `pulumi:"location"`
	// Maximum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes.
	MaxReplicaCount int `pulumi:"maxReplicaCount"`
	// Minimum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes.
	MinReplicaCount int `pulumi:"minReplicaCount"`
	// The name of the Autoscaling Group.
	Name *string `pulumi:"name"`
	// Specifies the behavior of this Autoscaling Group. A policy consists of Triggers and Actions, whereby an Action is some kind of automated behavior, and a Trigger is defined by the circumstances under which the Action is triggered. Currently, two separate Actions, namely Scaling In and Out are supported, triggered through Thresholds defined on a given Metric.
	Policies              []GetAutoscalingGroupPolicy               `pulumi:"policies"`
	ReplicaConfigurations []GetAutoscalingGroupReplicaConfiguration `pulumi:"replicaConfigurations"`
	TargetReplicaCount    int                                       `pulumi:"targetReplicaCount"`
}

func LookupAutoscalingGroupOutput(ctx *pulumi.Context, args LookupAutoscalingGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAutoscalingGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoscalingGroupResult, error) {
			args := v.(LookupAutoscalingGroupArgs)
			r, err := LookupAutoscalingGroup(ctx, &args, opts...)
			var s LookupAutoscalingGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoscalingGroupResultOutput)
}

// A collection of arguments for invoking getAutoscalingGroup.
type LookupAutoscalingGroupOutputArgs struct {
	// Id of an existing Autoscaling Group that you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing Autoscaling Group that you want to search for.
	//
	// Either `name` or `id` must be provided. If none or both are provided, the datasource will return an error.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAutoscalingGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscalingGroupArgs)(nil)).Elem()
}

// A collection of values returned by getAutoscalingGroup.
type LookupAutoscalingGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAutoscalingGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscalingGroupResult)(nil)).Elem()
}

func (o LookupAutoscalingGroupResultOutput) ToLookupAutoscalingGroupResultOutput() LookupAutoscalingGroupResultOutput {
	return o
}

func (o LookupAutoscalingGroupResultOutput) ToLookupAutoscalingGroupResultOutputWithContext(ctx context.Context) LookupAutoscalingGroupResultOutput {
	return o
}

func (o LookupAutoscalingGroupResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// Unique identifier for the resource
func (o LookupAutoscalingGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Location of the datacenter. This location is the same as the one from the selected template.
func (o LookupAutoscalingGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

// Maximum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes.
func (o LookupAutoscalingGroupResultOutput) MaxReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) int { return v.MaxReplicaCount }).(pulumi.IntOutput)
}

// Minimum replica count value for `targetReplicaCount`. Will be enforced for both automatic and manual changes.
func (o LookupAutoscalingGroupResultOutput) MinReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) int { return v.MinReplicaCount }).(pulumi.IntOutput)
}

// The name of the Autoscaling Group.
func (o LookupAutoscalingGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Specifies the behavior of this Autoscaling Group. A policy consists of Triggers and Actions, whereby an Action is some kind of automated behavior, and a Trigger is defined by the circumstances under which the Action is triggered. Currently, two separate Actions, namely Scaling In and Out are supported, triggered through Thresholds defined on a given Metric.
func (o LookupAutoscalingGroupResultOutput) Policies() GetAutoscalingGroupPolicyArrayOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) []GetAutoscalingGroupPolicy { return v.Policies }).(GetAutoscalingGroupPolicyArrayOutput)
}

func (o LookupAutoscalingGroupResultOutput) ReplicaConfigurations() GetAutoscalingGroupReplicaConfigurationArrayOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) []GetAutoscalingGroupReplicaConfiguration {
		return v.ReplicaConfigurations
	}).(GetAutoscalingGroupReplicaConfigurationArrayOutput)
}

func (o LookupAutoscalingGroupResultOutput) TargetReplicaCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupAutoscalingGroupResult) int { return v.TargetReplicaCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoscalingGroupResultOutput{})
}
