package v1alpha1

type InterfaceConfig struct {
	Name   string   `json:"name"`
	IpAddr []string `json:"ipAddr"`
	GwAddr []string `json:"gwAddr"`
}

type UpfCapacity struct {
	UplinkThroughput   string `json:"uplinkThroughput"`
	DownlinkThroughput string `json:"downlinkThroughput"`
}

type N6InterfaceConfig struct {
	Dnn        string          `json:"dnn"`
	Interface  InterfaceConfig `json:"interface"`
	IpAddrPool string          `json:"ipAddrPool"`
}

// UpfDeploySpec specifies config parameters for UPF
type UpfDeploySpec struct {
	ImagePaths   map[string]string   `json:"imagePaths,omitempty"`
	Capacity     UpfCapacity         `json:"capacity,omitempty"`
	N3Interfaces []InterfaceConfig   `json:"n3Interfaces,omitempty"`
	N4Interfaces []InterfaceConfig   `json:"n4Interfaces,omitempty"`
	N6Interfaces []N6InterfaceConfig `json:"n6Interfaces,omitempty"`
	// +optional
	N9Interfaces []InterfaceConfig `json:"n9Interfaces,omitempty"`
}
