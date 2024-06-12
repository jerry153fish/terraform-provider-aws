// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package servicediscovery

import (
	"context"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	servicediscovery_sdkv2 "github.com/aws/aws-sdk-go-v2/service/servicediscovery"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceDNSNamespace,
			TypeName: "aws_service_discovery_dns_namespace",
			Name:     "DNS Namespace",
		},
		{
			Factory:  dataSourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
			Name:     "HTTP Namespace",
		},
		{
			Factory:  DataSourceService,
			TypeName: "aws_service_discovery_service",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceHTTPNamespace,
			TypeName: "aws_service_discovery_http_namespace",
			Name:     "HTTP Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceInstance,
			TypeName: "aws_service_discovery_instance",
		},
		{
			Factory:  ResourcePrivateDNSNamespace,
			TypeName: "aws_service_discovery_private_dns_namespace",
			Name:     "Private DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourcePublicDNSNamespace,
			TypeName: "aws_service_discovery_public_dns_namespace",
			Name:     "Public DNS Namespace",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
		{
			Factory:  ResourceService,
			TypeName: "aws_service_discovery_service",
			Name:     "Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrARN,
			},
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.ServiceDiscovery
}

// NewClient returns a new AWS SDK for Go v2 client for this service package's AWS API.
func (p *servicePackage) NewClient(ctx context.Context, config map[string]any) (*servicediscovery_sdkv2.Client, error) {
	cfg := *(config["aws_sdkv2_config"].(*aws_sdkv2.Config))

	return servicediscovery_sdkv2.NewFromConfig(cfg, func(o *servicediscovery_sdkv2.Options) {
		if endpoint := config[names.AttrEndpoint].(string); endpoint != "" {
			tflog.Debug(ctx, "setting endpoint", map[string]any{
				"tf_aws.endpoint": endpoint,
			})
			o.BaseEndpoint = aws_sdkv2.String(endpoint)

			if o.EndpointOptions.UseFIPSEndpoint == aws_sdkv2.FIPSEndpointStateEnabled {
				tflog.Debug(ctx, "endpoint set, ignoring UseFIPSEndpoint setting")
				o.EndpointOptions.UseFIPSEndpoint = aws_sdkv2.FIPSEndpointStateDisabled
			}
		}
	}), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
