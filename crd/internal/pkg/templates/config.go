package templates

// Config describes the top-level structure of `sdks.yaml`.
//
// The YAML file is a mapping where each key is the name of a CRD and the value
// contains the metadata needed to fetch and generate SDKs for that CRD.
// Example:
// cert-manager:
//   repository: https://github.com/cert-manager/cert-manager
//   version: 1.0.0
//   crd:
//     - https://.../${VERSION}/cert-manager.crds.yaml
//
// Therefore, Config is represented as a map keyed by CRD name to its definition.
type Config map[string]CRDDefinition

// CRDDefinition represents the per-CRD configuration block.
type CRDDefinition struct {
    // Repository is a URL to the upstream repository of the CRD project.
    Repository string   `yaml:"repository"`
    // Version is a semantic version in the form MAJOR.MINOR.BUILD.
    Version    string   `yaml:"version"`
    // CRD is a list of URLs to the CRD schema files. The placeholder ${VERSION}
    // can be used to be replaced with the Version value.
    CRD        []string `yaml:"crd"`
}
