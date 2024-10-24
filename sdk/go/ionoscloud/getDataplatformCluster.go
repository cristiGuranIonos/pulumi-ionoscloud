// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **Dataplatform Cluster Data Source** can be used to search for and return an existing Dataplatform Cluster.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search and make sure that your resources have unique names.
//
// ## Example Usage
//
// ### By Name
//
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
//			_, err := ionoscloud.LookupDataplatformCluster(ctx, &ionoscloud.LookupDataplatformClusterArgs{
//				Name: pulumi.StringRef("Dataplatform_Cluster_Example"),
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
//
// ### By Name with Partial Match
//
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
//			_, err := ionoscloud.LookupDataplatformCluster(ctx, &ionoscloud.LookupDataplatformClusterArgs{
//				Name:         pulumi.StringRef("_Example"),
//				PartialMatch: pulumi.BoolRef(true),
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
func LookupDataplatformCluster(ctx *pulumi.Context, args *LookupDataplatformClusterArgs, opts ...pulumi.InvokeOption) (*LookupDataplatformClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDataplatformClusterResult
	err := ctx.Invoke("ionoscloud:index/getDataplatformCluster:getDataplatformCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDataplatformCluster.
type LookupDataplatformClusterArgs struct {
	// ID of the cluster you want to search for.
	Id *string `pulumi:"id"`
	// Name of an existing cluster that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name *string `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch *bool `pulumi:"partialMatch"`
}

// A collection of values returned by getDataplatformCluster.
type LookupDataplatformClusterResult struct {
	// base64 decoded cluster certificate authority data (provided as an attribute for direct use)
	CaCrt string `pulumi:"caCrt"`
	// structured kubernetes config consisting of a list with 1 item with the following fields:
	// * apiVersion - Kubernetes API Version
	// * kind - "Config"
	// * current-context - string
	// * clusters - list of
	// * name - name of cluster
	// * cluster - map of
	// * certificate-authority-data - **base64 decoded** cluster CA data
	// * server -  server address in the form `https://host:port`
	// * contexts - list of
	// * name - context name
	// * context - map of
	// * cluster - cluster name
	// * user - cluster user
	// * users - list of
	// * name - user name
	// * user - map of
	// * token - user token used for authentication
	Configs []GetDataplatformClusterConfig `pulumi:"configs"`
	// The UUID of the virtual data center (VDC) in which the cluster is provisioned.
	DatacenterId string `pulumi:"datacenterId"`
	// The UUID of the cluster.
	Id *string `pulumi:"id"`
	// Kubernetes configuration
	KubeConfig string `pulumi:"kubeConfig"`
	// A list of LANs you want this node pool to be part of
	Lans []GetDataplatformClusterLan `pulumi:"lans"`
	// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
	MaintenanceWindows []GetDataplatformClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of your cluster.
	Name         *string `pulumi:"name"`
	PartialMatch *bool   `pulumi:"partialMatch"`
	// cluster server (same as `config[0].clusters[0].cluster.server` but provided as an attribute for ease of use)
	Server string `pulumi:"server"`
	// a convenience map to be search the token of a specific user
	// * key - is the user name
	// * value - is the token
	UserTokens map[string]string `pulumi:"userTokens"`
	// The version of the Data Platform.
	Version string `pulumi:"version"`
}

func LookupDataplatformClusterOutput(ctx *pulumi.Context, args LookupDataplatformClusterOutputArgs, opts ...pulumi.InvokeOption) LookupDataplatformClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDataplatformClusterResult, error) {
			args := v.(LookupDataplatformClusterArgs)
			r, err := LookupDataplatformCluster(ctx, &args, opts...)
			var s LookupDataplatformClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDataplatformClusterResultOutput)
}

// A collection of arguments for invoking getDataplatformCluster.
type LookupDataplatformClusterOutputArgs struct {
	// ID of the cluster you want to search for.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing cluster that you want to search for. Search by name is case-insensitive. The whole resource name is required if `partialMatch` parameter is not set to true.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Whether partial matching is allowed or not when using name argument. Default value is false.
	//
	// Either `id` or `name` must be provided. If none, or both are provided, the datasource will return an error.
	PartialMatch pulumi.BoolPtrInput `pulumi:"partialMatch"`
}

func (LookupDataplatformClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataplatformClusterArgs)(nil)).Elem()
}

// A collection of values returned by getDataplatformCluster.
type LookupDataplatformClusterResultOutput struct{ *pulumi.OutputState }

func (LookupDataplatformClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDataplatformClusterResult)(nil)).Elem()
}

func (o LookupDataplatformClusterResultOutput) ToLookupDataplatformClusterResultOutput() LookupDataplatformClusterResultOutput {
	return o
}

func (o LookupDataplatformClusterResultOutput) ToLookupDataplatformClusterResultOutputWithContext(ctx context.Context) LookupDataplatformClusterResultOutput {
	return o
}

// base64 decoded cluster certificate authority data (provided as an attribute for direct use)
func (o LookupDataplatformClusterResultOutput) CaCrt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) string { return v.CaCrt }).(pulumi.StringOutput)
}

// structured kubernetes config consisting of a list with 1 item with the following fields:
// * apiVersion - Kubernetes API Version
// * kind - "Config"
// * current-context - string
// * clusters - list of
// * name - name of cluster
// * cluster - map of
// * certificate-authority-data - **base64 decoded** cluster CA data
// * server -  server address in the form `https://host:port`
// * contexts - list of
// * name - context name
// * context - map of
// * cluster - cluster name
// * user - cluster user
// * users - list of
// * name - user name
// * user - map of
// * token - user token used for authentication
func (o LookupDataplatformClusterResultOutput) Configs() GetDataplatformClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) []GetDataplatformClusterConfig { return v.Configs }).(GetDataplatformClusterConfigArrayOutput)
}

// The UUID of the virtual data center (VDC) in which the cluster is provisioned.
func (o LookupDataplatformClusterResultOutput) DatacenterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) string { return v.DatacenterId }).(pulumi.StringOutput)
}

// The UUID of the cluster.
func (o LookupDataplatformClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kubernetes configuration
func (o LookupDataplatformClusterResultOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) string { return v.KubeConfig }).(pulumi.StringOutput)
}

// A list of LANs you want this node pool to be part of
func (o LookupDataplatformClusterResultOutput) Lans() GetDataplatformClusterLanArrayOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) []GetDataplatformClusterLan { return v.Lans }).(GetDataplatformClusterLanArrayOutput)
}

// Starting time of a weekly 4 hour-long window, during which maintenance might occur in hh:mm:ss format
func (o LookupDataplatformClusterResultOutput) MaintenanceWindows() GetDataplatformClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) []GetDataplatformClusterMaintenanceWindow {
		return v.MaintenanceWindows
	}).(GetDataplatformClusterMaintenanceWindowArrayOutput)
}

// The name of your cluster.
func (o LookupDataplatformClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupDataplatformClusterResultOutput) PartialMatch() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) *bool { return v.PartialMatch }).(pulumi.BoolPtrOutput)
}

// cluster server (same as `config[0].clusters[0].cluster.server` but provided as an attribute for ease of use)
func (o LookupDataplatformClusterResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) string { return v.Server }).(pulumi.StringOutput)
}

// a convenience map to be search the token of a specific user
// * key - is the user name
// * value - is the token
func (o LookupDataplatformClusterResultOutput) UserTokens() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) map[string]string { return v.UserTokens }).(pulumi.StringMapOutput)
}

// The version of the Data Platform.
func (o LookupDataplatformClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDataplatformClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDataplatformClusterResultOutput{})
}
