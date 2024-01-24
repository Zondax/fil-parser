package types

type NodeInfo struct {
	// NodeFullVersion contains the node version from which this metadata was extracted
	NodeFullVersion string `json:"node_full_version,omitempty"`
	// NodeMajorMinorVersion contains the node major.minor version from which this metadata was extracted
	NodeMajorMinorVersion string `json:"node_major_minor_version,omitempty"`
}
