package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: "device.lqingcloud.cn", Version: "v1beta1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// AddToScheme 会被ClientSet代码调用
	AddToScheme = SchemeBuilder.AddToScheme
)

// Resource 会被 lister 代码调用
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func init() {
	SchemeBuilder.Register(&Nic{}, &NicList{})
}
