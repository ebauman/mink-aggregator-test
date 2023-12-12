package server

import (
	"github.com/acorn-io/baaah/pkg/restconfig"
	"github.com/acorn-io/mink/pkg/aggregator"
	"github.com/acorn-io/mink/pkg/serializer"
	"github.com/acorn-io/mink/pkg/server"
	"github.com/acorn-io/mink/pkg/strategy/remote"
	badapplev1 "github.com/ebauman/mink-aggregator-test/pkg/apis/badapple.com/v1"
	examplev1 "github.com/ebauman/mink-aggregator-test/pkg/apis/example.com/v1"
	"github.com/ebauman/mink-aggregator-test/pkg/openapi/badapple_com"
	"github.com/ebauman/mink-aggregator-test/pkg/openapi/example_com"
	"github.com/ebauman/mink-aggregator-test/pkg/scheme/bar"
	"github.com/ebauman/mink-aggregator-test/pkg/scheme/foo"
	"github.com/ebauman/mink-aggregator-test/pkg/server/registry"
	"k8s.io/apiserver/pkg/registry/rest"
	server2 "k8s.io/apiserver/pkg/server"
	rest2 "k8s.io/client-go/rest"
	"net/url"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	exampleUrl  *url.URL
	badappleUrl *url.URL
	restConfig  *rest2.Config
)

func init() {
	var err error
	exampleUrl, err = url.Parse("https://localhost:8443/")
	if err != nil {
		panic(err)
	}
	badappleUrl, err = url.Parse("https://localhost:8444/")
	if err != nil {
		panic(err)
	}

	restConfig, err = restconfig.FromFile("kubeconfig", "k3d-hf")
	if err != nil {
		panic(err)
	}
}

func NewAggregatedServer() (*server.Server, error) {
	return aggregator.NewAggregator(&aggregator.Config{
		Config: server.Config{
			Name:                         "mink-aggregator",
			Version:                      "1",
			HTTPListenPort:               9090,
			HTTPSListenPort:              9443,
			SkipInClusterLookup:          true,
			RemoteKubeConfigFileOptional: true,
		},
		Delegates: []*aggregator.Delegate{
			{
				Config: &rest2.Config{
					TLSClientConfig: rest2.TLSClientConfig{
						Insecure: true,
					},
					Host: exampleUrl.Host,
				},
			},
			{
				Config: &rest2.Config{
					TLSClientConfig: rest2.TLSClientConfig{
						Insecure: true,
					},
					Host: badappleUrl.Host,
				},
			},
			{
				Config: restConfig,
			},
		},
	})
}

//func NewAggregatedServer() (*server.Server, error) {
//	return server.New(&server.Config{
//		Name:                  "mink-aggregator",
//		Version:               "",
//		Authenticator:         nil,
//		Authorization:         nil,
//		HTTPListenPort:        9090,
//		HTTPSListenPort:       9443,
//		LongRunningVerbs:      nil,
//		LongRunningResources:  nil,
//		OpenAPIConfig:         nil,
//		Scheme:                nil,
//		CodecFactory:          nil,
//		APIGroups:             nil,
//		Middleware:            nil,
//		PostStartFunc:         nil,
//		SupportAPIAggregation: false,
//		DefaultOptions:        nil,
//		AuditConfig:           nil,
//		IgnoreStartFailure:    false,
//		ReadinessCheckers:     nil,
//		Aggregates: []*server.AggregateConfig{
//			{
//				TransportConfig: &transport.Config{
//					TLS: transport.TLSConfig{
//						CAFile: "./apiserver.local.config/certificates/ca.crt",
//					},
//				},
//				Endpoint: exampleUrl,
//			},
//			{
//				TransportConfig: &transport.Config{
//					TLS: transport.TLSConfig{
//						CAFile: "./apiserver.local.config/certificates/ca.crt",
//					},
//				},
//				Endpoint: badappleUrl,
//			},
//		},
//	})
//}

func NewFooServer(kclient client.WithWatch) (*server.Server, error) {
	fooRemote := remote.NewRemote(&examplev1.Foo{}, kclient)
	fooStorage, err := registry.NewFooStorage(fooRemote)
	if err != nil {
		return nil, err
	}

	apiGroupInfo := server2.NewDefaultAPIGroupInfo(
		examplev1.SchemeGroupVersion.Group,
		foo.Scheme,
		foo.ParameterCodec,
		foo.Codecs,
	)
	apiGroupInfo.VersionedResourcesStorageMap["v1"] = map[string]rest.Storage{
		"foos": fooStorage,
	}
	apiGroupInfo.NegotiatedSerializer = serializer.NewNoProtobufSerializer(apiGroupInfo.NegotiatedSerializer)

	return server.New(&server.Config{
		Name:                         "fooserver",
		Version:                      "1",
		Authenticator:                nil,
		Authorization:                nil,
		HTTPListenPort:               8080,
		HTTPSListenPort:              8443,
		LongRunningVerbs:             []string{"watch"},
		LongRunningResources:         nil,
		OpenAPIConfig:                example_com.GetOpenAPIDefinitions,
		Scheme:                       foo.Scheme,
		CodecFactory:                 &foo.Codecs,
		APIGroups:                    []*server2.APIGroupInfo{&apiGroupInfo},
		Middleware:                   nil,
		PostStartFunc:                nil,
		SupportAPIAggregation:        false,
		DefaultOptions:               nil,
		AuditConfig:                  nil,
		IgnoreStartFailure:           false,
		ReadinessCheckers:            nil,
		RemoteKubeConfigFileOptional: true,
		SkipInClusterLookup:          true,
	})
}

func NewBarServer(kclient client.WithWatch) (*server.Server, error) {
	barRemote := remote.NewRemote(&badapplev1.Bar{}, kclient)
	barStorage, err := registry.NewBarStorage(barRemote)
	if err != nil {
		return nil, err
	}

	apiGroupInfo := server2.NewDefaultAPIGroupInfo(
		badapplev1.SchemeGroupVersion.Group,
		bar.Scheme,
		bar.ParameterCodec,
		bar.Codecs,
	)
	apiGroupInfo.VersionedResourcesStorageMap["v1"] = map[string]rest.Storage{
		"bars": barStorage,
	}
	apiGroupInfo.NegotiatedSerializer = serializer.NewNoProtobufSerializer(apiGroupInfo.NegotiatedSerializer)

	return server.New(&server.Config{
		Name:                         "barserver",
		Version:                      "1",
		Authenticator:                nil,
		Authorization:                nil,
		HTTPListenPort:               8081,
		HTTPSListenPort:              8444,
		LongRunningVerbs:             []string{"watch"},
		LongRunningResources:         nil,
		OpenAPIConfig:                badapple_com.GetOpenAPIDefinitions,
		Scheme:                       bar.Scheme,
		CodecFactory:                 &bar.Codecs,
		APIGroups:                    []*server2.APIGroupInfo{&apiGroupInfo},
		Middleware:                   nil,
		PostStartFunc:                nil,
		SupportAPIAggregation:        false,
		DefaultOptions:               nil,
		AuditConfig:                  nil,
		IgnoreStartFailure:           false,
		ReadinessCheckers:            nil,
		RemoteKubeConfigFileOptional: true,
		SkipInClusterLookup:          true,
	})
}
