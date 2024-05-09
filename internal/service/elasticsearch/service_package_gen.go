// Code generated by internal/generate/servicepackages/main.go; DO NOT EDIT.

package elasticsearch

import (
	"context"

	aws_sdkv1 "github.com/aws/aws-sdk-go/aws"
	session_sdkv1 "github.com/aws/aws-sdk-go/aws/session"
	elasticsearchservice_sdkv1 "github.com/aws/aws-sdk-go/service/elasticsearchservice"
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
			Factory:  DataSourceDomain,
			TypeName: "aws_elasticsearch_domain",
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  ResourceDomain,
			TypeName: "aws_elasticsearch_domain",
			Name:     "Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  ResourceDomainPolicy,
			TypeName: "aws_elasticsearch_domain_policy",
		},
		{
			Factory:  ResourceDomainSAMLOptions,
			TypeName: "aws_elasticsearch_domain_saml_options",
		},
		{
			Factory:  ResourceVPCEndpoint,
			TypeName: "aws_elasticsearch_vpc_endpoint",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.Elasticsearch
}

// NewConn returns a new AWS SDK for Go v1 client for this service package's AWS API.
func (p *servicePackage) NewConn(ctx context.Context, config map[string]any) (*elasticsearchservice_sdkv1.ElasticsearchService, error) {
	sess := config["session"].(*session_sdkv1.Session)

	return elasticsearchservice_sdkv1.New(sess.Copy(&aws_sdkv1.Config{Endpoint: aws_sdkv1.String(config[names.AttrEndpoint].(string))})), nil
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
