package registry

import (
	"github.com/acorn-io/mink/pkg/stores"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/ebauman/mink-aggregator-test/pkg/apis/example.com/v1"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewFooStorage(fooStrategy strategy.CompleteStrategy) (rest.Storage, error) {
	return stores.NewBuilder(fooStrategy.Scheme(), &v1.Foo{}).
		WithCompleteCRUD(fooStrategy).Build(), nil
}
