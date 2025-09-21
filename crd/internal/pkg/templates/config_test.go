package templates

import (
	"os"
	"path/filepath"
	"testing"

	"gopkg.in/yaml.v3"
)

// helper to load a YAML file from testdata
func loadYAMLConfig(t *testing.T, name string) Config {
	t.Helper()
	path := filepath.Join("testdata", name)
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed reading %s: %v", path, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(b, &cfg); err != nil {
		t.Fatalf("failed to unmarshal %s: %v", path, err)
	}
	return cfg
}

func TestUnmarshalSingle(t *testing.T) {
	cfg := loadYAMLConfig(t, "sdks_single.yaml")

	if len(cfg) != 1 {
		t.Fatalf("expected 1 CRD entry, got %d", len(cfg))
	}

	crd, ok := cfg["cert-manager"]
	if !ok {
		t.Fatalf("expected key 'cert-manager' to exist")
	}
	if crd.Repository != "https://github.com/cert-manager/cert-manager" {
		t.Errorf("repository mismatch: %q", crd.Repository)
	}
	if crd.Version != "1.0.0" {
		t.Errorf("version mismatch: %q", crd.Version)
	}
	if len(crd.CRD) != 1 {
		t.Fatalf("expected 1 crd url, got %d", len(crd.CRD))
	}
	if crd.CRD[0] != "https://github.com/cert-manager/cert-manager/releases/download/${VERSION}/cert-manager.crds.yaml" {
		t.Errorf("crd url mismatch: %q", crd.CRD[0])
	}
}

func TestUnmarshalMultiple(t *testing.T) {
	cfg := loadYAMLConfig(t, "sdks_multiple.yaml")

	if len(cfg) != 2 {
		t.Fatalf("expected 2 CRD entries, got %d", len(cfg))
	}

	cm, ok := cfg["cert-manager"]
	if !ok {
		t.Fatalf("expected key 'cert-manager' to exist")
	}
	if cm.Version != "1.11.2" {
		t.Errorf("cert-manager version mismatch: %q", cm.Version)
	}

	ed, ok := cfg["external-dns"]
	if !ok {
		t.Fatalf("expected key 'external-dns' to exist")
	}
	if ed.Repository != "https://github.com/kubernetes-sigs/external-dns" {
		t.Errorf("external-dns repository mismatch: %q", ed.Repository)
	}
	if ed.Version != "0.14.1" {
		t.Errorf("external-dns version mismatch: %q", ed.Version)
	}
	if len(ed.CRD) != 2 {
		t.Fatalf("expected 2 crd urls for external-dns, got %d", len(ed.CRD))
	}
	if ed.CRD[0] != "https://raw.githubusercontent.com/kubernetes-sigs/external-dns/v${VERSION}/docs/contributing/crd-source/crd-release-1.yaml" {
		t.Errorf("first crd url mismatch: %q", ed.CRD[0])
	}
	if ed.CRD[1] != "https://raw.githubusercontent.com/kubernetes-sigs/external-dns/v${VERSION}/docs/contributing/crd-source/crd-release-2.yaml" {
		t.Errorf("second crd url mismatch: %q", ed.CRD[1])
	}
}
