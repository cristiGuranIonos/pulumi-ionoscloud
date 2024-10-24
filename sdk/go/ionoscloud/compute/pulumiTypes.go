// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type DatacenterCpuArchitecture struct {
	// A valid CPU family name
	CpuFamily *string `pulumi:"cpuFamily"`
	// The maximum number of cores available
	MaxCores *int `pulumi:"maxCores"`
	// The maximum number of RAM in MB
	MaxRam *int `pulumi:"maxRam"`
	// A valid CPU vendor name
	Vendor *string `pulumi:"vendor"`
}

// DatacenterCpuArchitectureInput is an input type that accepts DatacenterCpuArchitectureArgs and DatacenterCpuArchitectureOutput values.
// You can construct a concrete instance of `DatacenterCpuArchitectureInput` via:
//
//	DatacenterCpuArchitectureArgs{...}
type DatacenterCpuArchitectureInput interface {
	pulumi.Input

	ToDatacenterCpuArchitectureOutput() DatacenterCpuArchitectureOutput
	ToDatacenterCpuArchitectureOutputWithContext(context.Context) DatacenterCpuArchitectureOutput
}

type DatacenterCpuArchitectureArgs struct {
	// A valid CPU family name
	CpuFamily pulumi.StringPtrInput `pulumi:"cpuFamily"`
	// The maximum number of cores available
	MaxCores pulumi.IntPtrInput `pulumi:"maxCores"`
	// The maximum number of RAM in MB
	MaxRam pulumi.IntPtrInput `pulumi:"maxRam"`
	// A valid CPU vendor name
	Vendor pulumi.StringPtrInput `pulumi:"vendor"`
}

func (DatacenterCpuArchitectureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterCpuArchitecture)(nil)).Elem()
}

func (i DatacenterCpuArchitectureArgs) ToDatacenterCpuArchitectureOutput() DatacenterCpuArchitectureOutput {
	return i.ToDatacenterCpuArchitectureOutputWithContext(context.Background())
}

func (i DatacenterCpuArchitectureArgs) ToDatacenterCpuArchitectureOutputWithContext(ctx context.Context) DatacenterCpuArchitectureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterCpuArchitectureOutput)
}

// DatacenterCpuArchitectureArrayInput is an input type that accepts DatacenterCpuArchitectureArray and DatacenterCpuArchitectureArrayOutput values.
// You can construct a concrete instance of `DatacenterCpuArchitectureArrayInput` via:
//
//	DatacenterCpuArchitectureArray{ DatacenterCpuArchitectureArgs{...} }
type DatacenterCpuArchitectureArrayInput interface {
	pulumi.Input

	ToDatacenterCpuArchitectureArrayOutput() DatacenterCpuArchitectureArrayOutput
	ToDatacenterCpuArchitectureArrayOutputWithContext(context.Context) DatacenterCpuArchitectureArrayOutput
}

type DatacenterCpuArchitectureArray []DatacenterCpuArchitectureInput

func (DatacenterCpuArchitectureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatacenterCpuArchitecture)(nil)).Elem()
}

func (i DatacenterCpuArchitectureArray) ToDatacenterCpuArchitectureArrayOutput() DatacenterCpuArchitectureArrayOutput {
	return i.ToDatacenterCpuArchitectureArrayOutputWithContext(context.Background())
}

func (i DatacenterCpuArchitectureArray) ToDatacenterCpuArchitectureArrayOutputWithContext(ctx context.Context) DatacenterCpuArchitectureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatacenterCpuArchitectureArrayOutput)
}

type DatacenterCpuArchitectureOutput struct{ *pulumi.OutputState }

func (DatacenterCpuArchitectureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatacenterCpuArchitecture)(nil)).Elem()
}

func (o DatacenterCpuArchitectureOutput) ToDatacenterCpuArchitectureOutput() DatacenterCpuArchitectureOutput {
	return o
}

func (o DatacenterCpuArchitectureOutput) ToDatacenterCpuArchitectureOutputWithContext(ctx context.Context) DatacenterCpuArchitectureOutput {
	return o
}

// A valid CPU family name
func (o DatacenterCpuArchitectureOutput) CpuFamily() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatacenterCpuArchitecture) *string { return v.CpuFamily }).(pulumi.StringPtrOutput)
}

// The maximum number of cores available
func (o DatacenterCpuArchitectureOutput) MaxCores() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatacenterCpuArchitecture) *int { return v.MaxCores }).(pulumi.IntPtrOutput)
}

// The maximum number of RAM in MB
func (o DatacenterCpuArchitectureOutput) MaxRam() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DatacenterCpuArchitecture) *int { return v.MaxRam }).(pulumi.IntPtrOutput)
}

// A valid CPU vendor name
func (o DatacenterCpuArchitectureOutput) Vendor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatacenterCpuArchitecture) *string { return v.Vendor }).(pulumi.StringPtrOutput)
}

type DatacenterCpuArchitectureArrayOutput struct{ *pulumi.OutputState }

func (DatacenterCpuArchitectureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatacenterCpuArchitecture)(nil)).Elem()
}

func (o DatacenterCpuArchitectureArrayOutput) ToDatacenterCpuArchitectureArrayOutput() DatacenterCpuArchitectureArrayOutput {
	return o
}

func (o DatacenterCpuArchitectureArrayOutput) ToDatacenterCpuArchitectureArrayOutputWithContext(ctx context.Context) DatacenterCpuArchitectureArrayOutput {
	return o
}

func (o DatacenterCpuArchitectureArrayOutput) Index(i pulumi.IntInput) DatacenterCpuArchitectureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatacenterCpuArchitecture {
		return vs[0].([]DatacenterCpuArchitecture)[vs[1].(int)]
	}).(DatacenterCpuArchitectureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatacenterCpuArchitectureInput)(nil)).Elem(), DatacenterCpuArchitectureArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatacenterCpuArchitectureArrayInput)(nil)).Elem(), DatacenterCpuArchitectureArray{})
	pulumi.RegisterOutputType(DatacenterCpuArchitectureOutput{})
	pulumi.RegisterOutputType(DatacenterCpuArchitectureArrayOutput{})
}
