package config

import (
	"encoding/json"
	"fmt"
	"go/types"
	"regexp"
	"strings"
)

type GoType struct {
	Path    string `json:"import" yaml:"import"`
	Package string `json:"package" yaml:"package"`
	Name    string `json:"type" yaml:"type"`
	Pointer bool   `json:"pointer" yaml:"pointer"`
	Spec    string
	BuiltIn bool
}

type ParsedGoType struct {
	ImportPath string
	Package    string
	TypeName   string
	BasicType  bool
	StructTag  string
}

func (o *GoType) UnmarshalJSON(data []byte) error {
	var spec string
	if err := json.Unmarshal(data, &spec); err == nil {
		*o = GoType{Spec: spec}
		return nil
	}
	type alias GoType
	var a alias
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}
	*o = GoType(a)
	return nil
}

func (o *GoType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var spec string
	if err := unmarshal(&spec); err == nil {
		*o = GoType{Spec: spec}
		return nil
	}
	type alias GoType
	var a alias
	if err := unmarshal(&a); err != nil {
		return err
	}
	*o = GoType(a)
	return nil
}

var validIdentifier = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var versionNumber = regexp.MustCompile(`^v[0-9]+$`)
var invalidIdentifier = regexp.MustCompile(`[^a-zA-Z0-9_]`)

func generatePackageID(importPath string) (string, bool) {
	parts := strings.Split(importPath, "/")
	name := parts[len(parts)-1]
	// If the last part of the import path is a valid identifier, assume that's the package name
	if versionNumber.MatchString(name) && len(parts) >= 2 {
		name = parts[len(parts)-2]
		return invalidIdentifier.ReplaceAllString(strings.ToLower(name), "_"), true
	}
	if validIdentifier.MatchString(name) {
		return name, false
	}
	return invalidIdentifier.ReplaceAllString(strings.ToLower(name), "_"), true
}

// validate GoType
func (gt GoType) Parse() (*ParsedGoType, error) {
	var o ParsedGoType

	if gt.Spec == "" {
		// TODO: Validation
		if gt.Path == "" && gt.Package != "" {
			return nil, fmt.Errorf("Package override `go_type`: package name requires an import path")
		}
		var pkg string
		var pkgNeedsAlias bool

		if gt.Package == "" && gt.Path != "" {
			pkg, pkgNeedsAlias = generatePackageID(gt.Path)
			if pkgNeedsAlias {
				o.Package = pkg
			}
		} else {
			pkg = gt.Package
			o.Package = gt.Package
		}

		o.ImportPath = gt.Path
		o.TypeName = gt.Name
		o.Ba