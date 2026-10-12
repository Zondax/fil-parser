package v2

import "testing"

// translateParserVersionFromMetadata routes each height's traces by the node major.minor version
// captured in its metadata; a version missing here makes every height captured from that node
// fail with "node version not supported". Lotus v1.38.0 is the recommended NV29 (Solstice) release.
func TestIsNodeVersionSupported(t *testing.T) {
	p := &Parser{}
	tests := []struct {
		version string
		want    bool
	}{
		{"v1.23", true},
		{"v1.36", true},
		{"v1.37", true},
		{"v1.38", true},
		{"V1.38", true},
		{"v1.22", false},
		{"v1.39", false},
	}
	for _, tt := range tests {
		t.Run(tt.version, func(t *testing.T) {
			if got := p.IsNodeVersionSupported(tt.version); got != tt.want {
				t.Fatalf("IsNodeVersionSupported(%q) = %v, want %v", tt.version, got, tt.want)
			}
		})
	}
}
