// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/cristiGuranIonos/pulumi-ionoscloud/sdk/go/ionoscloud/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetContractNumber(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:contractNumber")
}

// IonosCloud REST API URL. Usually not necessary to be set, SDKs know internally how to route requests to the API.
func GetEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:endpoint")
}

// IonosCloud password for API operations. If token is provided, token is preferred
func GetPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:password")
}

// Deprecated: Timeout is used instead of this functionality
func GetRetries(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "ionoscloud:retries")
}

// Access key for IONOS S3 operations.
func GetS3AccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:s3AccessKey")
}

// Region for IONOS S3 operations.
func GetS3Region(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:s3Region")
}

// Secret key for IONOS S3 operations.
func GetS3SecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:s3SecretKey")
}

// IonosCloud bearer token for API operations.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:token")
}

// IonosCloud username for API operations. If token is provided, token is preferred
func GetUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "ionoscloud:username")
}