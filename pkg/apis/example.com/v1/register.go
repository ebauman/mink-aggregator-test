package v1

import (
	example_com "github.com/ebauman/mink-aggregator-test/pkg/apis/example.com"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var SchemeGroupVersion = schema.GroupVersion{Group: example_com.Group, Version: Version}

var (
	SchemeBuilder = runtime.NewSchemeBuilder(addToScheme)
	AddToScheme   = SchemeBuilder.AddToScheme
)

func addToScheme(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&Foo{},
		&FooList{},
	)

	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)

	return nil
}
