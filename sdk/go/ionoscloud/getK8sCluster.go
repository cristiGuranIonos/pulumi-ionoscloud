// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ionoscloud

import (
	"context"
	"reflect"

	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The **k8s Cluster data source** can be used to search for and return existing k8s clusters.
// If a single match is found, it will be returned. If your search results in multiple matches, an error will be returned.
// When this happens, please refine your search string so that it is specific enough to return only one result.
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
//			_, err := ionoscloud.LookupK8sCluster(ctx, &ionoscloud.LookupK8sClusterArgs{
//				Name: pulumi.StringRef("K8s Cluster Example"),
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
func LookupK8sCluster(ctx *pulumi.Context, args *LookupK8sClusterArgs, opts ...pulumi.InvokeOption) (*LookupK8sClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupK8sClusterResult
	err := ctx.Invoke("ionoscloud:index/getK8sCluster:getK8sCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getK8sCluster.
type LookupK8sClusterArgs struct {
	// ID of the cluster you want to search for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	Id *string `pulumi:"id"`
	// Name of an existing cluster that you want to search for.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getK8sCluster.
type LookupK8sClusterResult struct {
	// access to the K8s API server is restricted to these CIDRs
	ApiSubnetAllowLists []string `pulumi:"apiSubnetAllowLists"`
	// A list of available versions for upgrading the cluster
	AvailableUpgradeVersions []string `pulumi:"availableUpgradeVersions"`
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
	Configs []GetK8sClusterConfig `pulumi:"configs"`
	// id of the cluster
	Id *string `pulumi:"id"`
	// Kubernetes version
	K8sVersion string `pulumi:"k8sVersion"`
	// Kubernetes configuration
	KubeConfig string `pulumi:"kubeConfig"`
	// this attribute is mandatory if the cluster is private.
	Location string `pulumi:"location"`
	// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
	MaintenanceWindows []GetK8sClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// name of the cluster
	Name *string `pulumi:"name"`
	// the NAT gateway IP of the cluster if the cluster is private.
	NatGatewayIp string `pulumi:"natGatewayIp"`
	// list of the IDs of the node pools in this cluster
	NodePools []string `pulumi:"nodePools"`
	// the node subnet of the cluster, if the cluster is private.
	NodeSubnet string `pulumi:"nodeSubnet"`
	// indicates if the cluster is public or private.
	Public bool `pulumi:"public"`
	// list of S3 bucket configured for K8s usage
	S3Buckets []GetK8sClusterS3Bucket `pulumi:"s3Buckets"`
	// cluster server (same as `config[0].clusters[0].cluster.server` but provided as an attribute for ease of use)
	Server string `pulumi:"server"`
	// one of "AVAILABLE",
	// "INACTIVE",
	// "BUSY",
	// "DEPLOYING",
	// "ACTIVE",
	// "FAILED",
	// "SUSPENDED",
	// "FAILED_SUSPENDED",
	// "UPDATING",
	// "FAILED_UPDATING",
	// "DESTROYING",
	// "FAILED_DESTROYING",
	// "TERMINATED"
	State string `pulumi:"state"`
	// a convenience map to be search the token of a specific user
	// - key - is the user name
	// - value - is the token
	UserTokens map[string]string `pulumi:"userTokens"`
	// A list of versions that may be used for node pools under this cluster
	ViableNodePoolVersions []string `pulumi:"viableNodePoolVersions"`
}

func LookupK8sClusterOutput(ctx *pulumi.Context, args LookupK8sClusterOutputArgs, opts ...pulumi.InvokeOption) LookupK8sClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupK8sClusterResult, error) {
			args := v.(LookupK8sClusterArgs)
			r, err := LookupK8sCluster(ctx, &args, opts...)
			var s LookupK8sClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupK8sClusterResultOutput)
}

// A collection of arguments for invoking getK8sCluster.
type LookupK8sClusterOutputArgs struct {
	// ID of the cluster you want to search for.
	//
	// Either `name` or `id` must be provided. If none, or both are provided, the datasource will return an error.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Name of an existing cluster that you want to search for.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupK8sClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sClusterArgs)(nil)).Elem()
}

// A collection of values returned by getK8sCluster.
type LookupK8sClusterResultOutput struct{ *pulumi.OutputState }

func (LookupK8sClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupK8sClusterResult)(nil)).Elem()
}

func (o LookupK8sClusterResultOutput) ToLookupK8sClusterResultOutput() LookupK8sClusterResultOutput {
	return o
}

func (o LookupK8sClusterResultOutput) ToLookupK8sClusterResultOutputWithContext(ctx context.Context) LookupK8sClusterResultOutput {
	return o
}

// access to the K8s API server is restricted to these CIDRs
func (o LookupK8sClusterResultOutput) ApiSubnetAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.ApiSubnetAllowLists }).(pulumi.StringArrayOutput)
}

// A list of available versions for upgrading the cluster
func (o LookupK8sClusterResultOutput) AvailableUpgradeVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.AvailableUpgradeVersions }).(pulumi.StringArrayOutput)
}

// base64 decoded cluster certificate authority data (provided as an attribute for direct use)
func (o LookupK8sClusterResultOutput) CaCrt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.CaCrt }).(pulumi.StringOutput)
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
func (o LookupK8sClusterResultOutput) Configs() GetK8sClusterConfigArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterConfig { return v.Configs }).(GetK8sClusterConfigArrayOutput)
}

// id of the cluster
func (o LookupK8sClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Kubernetes version
func (o LookupK8sClusterResultOutput) K8sVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.K8sVersion }).(pulumi.StringOutput)
}

// Kubernetes configuration
func (o LookupK8sClusterResultOutput) KubeConfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.KubeConfig }).(pulumi.StringOutput)
}

// this attribute is mandatory if the cluster is private.
func (o LookupK8sClusterResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Location }).(pulumi.StringOutput)
}

// A maintenance window comprise of a day of the week and a time for maintenance to be allowed
func (o LookupK8sClusterResultOutput) MaintenanceWindows() GetK8sClusterMaintenanceWindowArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterMaintenanceWindow { return v.MaintenanceWindows }).(GetK8sClusterMaintenanceWindowArrayOutput)
}

// name of the cluster
func (o LookupK8sClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// the NAT gateway IP of the cluster if the cluster is private.
func (o LookupK8sClusterResultOutput) NatGatewayIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.NatGatewayIp }).(pulumi.StringOutput)
}

// list of the IDs of the node pools in this cluster
func (o LookupK8sClusterResultOutput) NodePools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.NodePools }).(pulumi.StringArrayOutput)
}

// the node subnet of the cluster, if the cluster is private.
func (o LookupK8sClusterResultOutput) NodeSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.NodeSubnet }).(pulumi.StringOutput)
}

// indicates if the cluster is public or private.
func (o LookupK8sClusterResultOutput) Public() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) bool { return v.Public }).(pulumi.BoolOutput)
}

// list of S3 bucket configured for K8s usage
func (o LookupK8sClusterResultOutput) S3Buckets() GetK8sClusterS3BucketArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []GetK8sClusterS3Bucket { return v.S3Buckets }).(GetK8sClusterS3BucketArrayOutput)
}

// cluster server (same as `config[0].clusters[0].cluster.server` but provided as an attribute for ease of use)
func (o LookupK8sClusterResultOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.Server }).(pulumi.StringOutput)
}

// one of "AVAILABLE",
// "INACTIVE",
// "BUSY",
// "DEPLOYING",
// "ACTIVE",
// "FAILED",
// "SUSPENDED",
// "FAILED_SUSPENDED",
// "UPDATING",
// "FAILED_UPDATING",
// "DESTROYING",
// "FAILED_DESTROYING",
// "TERMINATED"
func (o LookupK8sClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) string { return v.State }).(pulumi.StringOutput)
}

// a convenience map to be search the token of a specific user
// - key - is the user name
// - value - is the token
func (o LookupK8sClusterResultOutput) UserTokens() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) map[string]string { return v.UserTokens }).(pulumi.StringMapOutput)
}

// A list of versions that may be used for node pools under this cluster
func (o LookupK8sClusterResultOutput) ViableNodePoolVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupK8sClusterResult) []string { return v.ViableNodePoolVersions }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupK8sClusterResultOutput{})
}