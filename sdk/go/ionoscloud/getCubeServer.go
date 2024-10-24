// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Cube Server data source** can be used to search for and return existing servers.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
//
// ## Example Usage
func LookupCubeServer(ctx *pulumi.Context, args *LookupCubeServerArgs, opts ...pulumi.InvokeOption) (*LookupCubeServerResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCubeServerResult
	err := ctx.Invoke("ionoscloud:index/getCubeServer:getCubeServer", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCubeServer.
type LookupCubeServerArgs struct {
	// Datacenter's UUID.
	DatacenterId string `pulumi:"datacenterId"`
	// ID of the server you want to search for.
	//
	// `datacenterId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	Id *string `pulumi:"id"`
	// Name of an existing server that you want to search for.
	Name *string `pulumi:"name"`
	// The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource
	TemplateUuid *string `pulumi:"templateUuid"`
}

// A collection of values returned by getCubeServer.
type LookupCubeServerResult struct {
	// The availability zone in which the volume should exist
	AvailabilityZone string `pulumi:"availabilityZone"`
	BootCdrom        string `pulumi:"bootCdrom"`
	BootImage        string `pulumi:"bootImage"`
	BootVolume       string `pulumi:"bootVolume"`
	// list of
	Cdroms    []GetCubeServerCdrom `pulumi:"cdroms"`
	Cores     int                  `pulumi:"cores"`
	CpuFamily string               `pulumi:"cpuFamily"`
	// The id of the datacenter
	DatacenterId string `pulumi:"datacenterId"`
	// Id of the firewall rule
	Id *string `pulumi:"id"`
	// Name of the firewall rule
	Name *string `pulumi:"name"`
	// list of
	Nics []GetCubeServerNic `pulumi:"nics"`
	Ram  int                `pulumi:"ram"`
	// The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource
	TemplateUuid *string `pulumi:"templateUuid"`
	Token        string  `pulumi:"token"`
	// Status of the virtual Machine
	VmState string `pulumi:"vmState"`
	// list of
	Volumes []GetCubeServerVolume `pulumi:"volumes"`
}

func LookupCubeServerOutput(ctx *pulumi.Context, args LookupCubeServerOutputArgs, opts ...pulumi.InvokeOption) LookupCubeServerResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCubeServerResult, error) {
			args := v.(LookupCubeServerArgs)
			r, err := LookupCubeServer(ctx, &args, opts...)
			var s LookupCubeServerResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCubeServerResultOutput)
}

// A collection of arguments for invoking getCubeServer.
type LookupCubeServerOutputArgs struct {
	// Datacenter's UUID.
	DatacenterId pulumi.StringInput `pulumi:"datacenterId"`
	// ID of the server you want to search for.
	//
	// `datacenterId` and either `name` or `id` must be provided. If none, or both of `name` and `id` are provided, the datasource will return an error.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing server that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource
	TemplateUuid pulumi.StringPtrInput `pulumi:"templateUuid"`
}

func (LookupCubeServerOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCubeServerArgs)(nil)).Elem()
}

// A collection of values returned by getCubeServer.
type LookupCubeServerResultOutput struct{ *pulumi.OutputState }

func (LookupCubeServerResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCubeServerResult)(nil)).Elem()
}

func (o LookupCubeServerResultOutput) ToLookupCubeServerResultOutput() LookupCubeServerResultOutput {
	return o
}

func (o LookupCubeServerResultOutput) ToLookupCubeServerResultOutputWithContext(ctx context.Context) LookupCubeServerResultOutput {
	return o
}

// The availability zone in which the volume should exist
func (o LookupCubeServerResultOutput) AvailabilityZone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.AvailabilityZone }).(pulumi.StringOutput)
}

func (o LookupCubeServerResultOutput) BootCdrom() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.BootCdrom }).(pulumi.StringOutput)
}

func (o LookupCubeServerResultOutput) BootImage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.BootImage }).(pulumi.StringOutput)
}

func (o LookupCubeServerResultOutput) BootVolume() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.BootVolume }).(pulumi.StringOutput)
}

// list of
func (o LookupCubeServerResultOutput) Cdroms() GetCubeServerCdromArrayOutput {
	return o.ApplyT(func(v LookupCubeServerResult) []GetCubeServerCdrom { return v.Cdroms }).(GetCubeServerCdromArrayOutput)
}

func (o LookupCubeServerResultOutput) Cores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCubeServerResult) int { return v.Cores }).(pulumi.IntOutput)
}

func (o LookupCubeServerResultOutput) CpuFamily() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.CpuFamily }).(pulumi.StringOutput)
}

// The id of the datacenter
func (o LookupCubeServerResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// Id of the firewall rule
func (o LookupCubeServerResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCubeServerResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Name of the firewall rule
func (o LookupCubeServerResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCubeServerResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// list of
func (o LookupCubeServerResultOutput) Nics() GetCubeServerNicArrayOutput {
	return o.ApplyT(func(v LookupCubeServerResult) []GetCubeServerNic { return v.Nics }).(GetCubeServerNicArrayOutput)
}

func (o LookupCubeServerResultOutput) Ram() pulumi.IntOutput {
	return o.ApplyT(func(v LookupCubeServerResult) int { return v.Ram }).(pulumi.IntOutput)
}

// The UUID of the template for creating a CUBE server; the available templates for CUBE servers can be found on the templates resource
func (o LookupCubeServerResultOutput) TemplateUuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCubeServerResult) *string { return v.TemplateUuid }).(pulumi.StringPtrOutput)
}

func (o LookupCubeServerResultOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.Token }).(pulumi.StringOutput)
}

// Status of the virtual Machine
func (o LookupCubeServerResultOutput) VmState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCubeServerResult) string { return v.VmState }).(pulumi.StringOutput)
}

// list of
func (o LookupCubeServerResultOutput) Volumes() GetCubeServerVolumeArrayOutput {
	return o.ApplyT(func(v LookupCubeServerResult) []GetCubeServerVolume { return v.Volumes }).(GetCubeServerVolumeArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCubeServerResultOutput{})
}