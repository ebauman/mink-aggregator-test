package foo

import (
	v1 "github.com/ebauman/mink-aggregator-test/pkg/apis/example.com/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var (
	Scheme = runtime.NewScheme()

	Codecs = serializer.NewCodecFactory(Scheme)

	ParameterCodec = runtime.NewParameterCodec(Scheme)
)

func AddToScheme(scheme *runtime.Scheme) error {
	metav1.AddToGroupVersion(scheme, v1.SchemeGroupVersion)

	if err := v1.AddToScheme(scheme); err != nil {
		return err
	}

	if err := corev1.AddToScheme(scheme); err != nil {
		return err
	}

	return nil
}

func init() {
	utilruntime.Must(AddToScheme(Scheme))
}
