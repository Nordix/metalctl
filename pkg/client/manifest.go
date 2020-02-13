package client

import (
	"sigs.k8s.io/kustomize/api/resource"
)

type ManifestFactory struct {
	resource.Resource
}

// Manifest interface
type Manifest interface {
	GetKustomizeResource() resource.Resource
	SetKustomizeResource(*resource.Resource) error
	AsYAML() ([]byte, error)
	MarshalJSON() ([]byte, error)
	GetName() string
	GetKind() string
	GetNamespace() string
	GetString(path string) (string, error)
	GetStringSlice(path string) ([]string, error)
	GetBool(path string) (bool, error)
	GetFloat64(path string) (float64, error)
	GetInt64(path string) (int64, error)
	GetSlice(path string) ([]interface{}, error)
	GetStringMap(path string) (map[string]string, error)
	GetMap(path string) (map[string]interface{}, error)
}

// GetNamespace returns the namespace the resource thinks it's in.
func (d *ManifestFactory) GetNamespace() string {
	r := d.GetKustomizeResource()
	return r.GetNamespace()
}

// GetString returns the string value at path.
func (d *ManifestFactory) GetString(path string) (string, error) {
	r := d.GetKustomizeResource()
	return r.GetString(path)
}

// GetStringSlice returns a string slice at path.
func (d *ManifestFactory) GetStringSlice(path string) ([]string, error) {
	r := d.GetKustomizeResource()
	return r.GetStringSlice(path)
}

// GetBool returns a bool at path.
func (d *ManifestFactory) GetBool(path string) (bool, error) {
	r := d.GetKustomizeResource()
	return r.GetBool(path)
}

// GetFloat64 returns a float64 at path.
func (d *ManifestFactory) GetFloat64(path string) (float64, error) {
	r := d.GetKustomizeResource()
	return r.GetFloat64(path)
}

// GetInt64 returns an int64 at path.
func (d *ManifestFactory) GetInt64(path string) (int64, error) {
	r := d.GetKustomizeResource()
	return r.GetInt64(path)
}

// GetSlice returns a slice at path.
func (d *ManifestFactory) GetSlice(path string) ([]interface{}, error) {
	r := d.GetKustomizeResource()
	return r.GetSlice(path)
}

// GetStringMap returns a string map at path.
func (d *ManifestFactory) GetStringMap(path string) (map[string]string, error) {
	r := d.GetKustomizeResource()
	return r.GetStringMap(path)
}

// GetMap returns a map at path.
func (d *ManifestFactory) GetMap(path string) (map[string]interface{}, error) {
	r := d.GetKustomizeResource()
	return r.GetMap(path)
}

// AsYAML returns the Manifest as a YAML byte stream.
func (d *ManifestFactory) AsYAML() ([]byte, error) {
	r := d.GetKustomizeResource()
	return r.AsYAML()
}

// MarshalJSON returns the Manifest as JSON.
func (d *ManifestFactory) MarshalJSON() ([]byte, error) {
	r := d.GetKustomizeResource()
	return r.MarshalJSON()
}

// GetName returns the name: field from the Manifest.
func (d *ManifestFactory) GetName() string {
	r := d.GetKustomizeResource()
	return r.GetName()
}

// GetKind returns the Kind: field from the Manifest.
func (d *ManifestFactory) GetKind() string {
	r := d.GetKustomizeResource()
	return r.GetKind()
}

// GetKustomizeResource returns a Kustomize Resource object for this Manifest.
func (d *ManifestFactory) GetKustomizeResource() resource.Resource {
	return d.Resource
}

// SetKustomizeResource sets a Kustomize Resource object for this Manifest.
func (d *ManifestFactory) SetKustomizeResource(r *resource.Resource) error {
	d.Resource = *r
	return nil
}

// NewManifest is a convenience method to construct a new Manifest.  Although
// an error is unlikely at this time, this provides some future proofing for
// when we want more strict airship specific validation of Manifests getting
// created as this would be the front door for all Kustomize->Airship
// Manifests - e.g. in the future all Manifests require an airship
// annotation X
func NewManifest(r *resource.Resource) (Manifest, error) {
	var doc Manifest = &ManifestFactory{}
	err := doc.SetKustomizeResource(r)
	return doc, err
}
