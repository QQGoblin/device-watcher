package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Nic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NicSpec   `json:"spec"`
	Status            NicStatus `json:"status"`
}

type NicSpec struct {
	DeviceName string   `json:"deviceName,omitempty"`
	MacAddress string   `json:"macAddress,omitempty"`
	Ipaddress  []string `json:"ipaddress,omitempty"`
	Node       string   `json:"node"`
}

type NicStatus struct {
	Up bool `json:"up"`
}

type NicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Nic `json:"items"`
}
