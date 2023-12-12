package registry

import (
	"github.com/acorn-io/mink/pkg/stores"
	"github.com/acorn-io/mink/pkg/strategy"
	v1 "github.com/ebauman/mink-aggregator-test/pkg/apis/badapple.com/v1"
	"k8s.io/apiserver/pkg/registry/rest"
)

func NewBarStorage(barStrategy strategy.CompleteStrategy) (rest.Storage, error) {
	return stores.NewBuilder(barStrategy.Scheme(), &v1.Bar{}).
		WithCompleteCRUD(barStrategy).Build(), nil
}
