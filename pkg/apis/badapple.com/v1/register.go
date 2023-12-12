package v1

import (
	badapple_com "github.com/ebauman/mink-aggregator-test/pkg/apis/badapple.com"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Version: Version, Group: badapple_com.Group}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addToScheme)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Bar{},
		&BarList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
