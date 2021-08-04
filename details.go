package main

import (
	"reflect"
)

// PackageSet represents the pacages within a PKGBUILD as a slice of Details
// maps from pkgname to Details
type PackageSet map[string]Details

// Details represents the desired details for a single package within a PKGBUILD,
// or the results of handing flags to `gendesk`.
// Examples of package names within a PKGBUILD file is, for instance, "app" and "app-doc".
type Details struct {
	pkgname     string `pkgname`
	pkgdesc     string `pkgdesc`
	name        string `_name`
	execCommand string `_exec`
	path        string `_path`
	icon        string `_icon`
	genericName string `_genericname`
	comment     string `_comment`
	custom      string `_custom`
	mimeTypes   string `_mimetypes` // _mimetype is an alias for _mimetypes
	categories  string `_categories`
}

// DetailsTags returns all the tags belonging to the Details struct
func DetailsTags() []string {
	var d Details
	v := reflect.ValueOf(d)
	numKeys := v.NumField()
	// Collect all tag names
	keyList := make([]string, numKeys, numKeys)
	for i := 0; i < numKeys; i++ {
		keyList[i] = string(v.Type().Field(i).Tag)
	}
	return keyList
}

// DetailsToVars will store all Detail files to the given Variables map.
// This function uses the struct field tags as variable keys.
func DetailsToVars(d *Details, vars *Variables) {
	v := reflect.ValueOf(*d)
	for i := 0; i < v.NumField(); i++ {
		tag := string(v.Type().Field(i).Tag)
		value := v.Field(i).String()
		if value != "" {
			vars.Set(tag, value)
		}
	}
}

// Given a vars map, return a new Details struct
func NewDetails(vars *Variables) *Details {
	var details Details
	VarsToDetails(vars, &details)
	return &details
}

func (d *Details) Has(key string) bool {
	v := reflect.ValueOf(*d)
	for i := 0; i < v.NumField(); i++ {
		tag := string(v.Type().Field(i).Tag)
		if key == tag {
			// return true if the key has a value
			value := v.Field(i).String()
			if value != "" {
				return true
			}
		}
	}
	return false
}

// VarsToDetails will copy variables to Details fields with the
// corresponding tag names.
func VarsToDetails(vars *Variables, d *Details) {
	v := reflect.ValueOf(*d)
	for i := 0; i < v.NumField(); i++ {
		tag := string(v.Type().Field(i).Tag)
		field := v.Field(i)
		if varValue, ok := (*vars)[tag]; ok && varValue != "" {
			field.SetString(varValue)
		}
	}
}

// ExpandOnce will expand variables in each field, once,
// using the keys and values in the Variables map.
// If any variable expansion was done, return true.
// This function uses the struct field tags as variable keys.
func (d *Details) ExpandOnce(vars *Variables) bool {
	changed := false
	v := reflect.ValueOf(*d)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		original := field.String()
		if original == "" {
			continue // don't try to expand empty strings
		}
		expanded := vars.Expand(original)
		field.SetString(expanded)
		if original != expanded {
			changed = true
		}
	}
	return changed
}

// ExpandUntilNoChangeMaxN is the same as ExpandUntilNoChange,
// but will stop at a maximum of N expansions.
// N can be ie. 10.
// If the maximum expansion was reached, false is returned.
// It's the fields in the Details structs that are expanded,
// using the keys+values in the given Variables map.
func (d *Details) ExpandUntilNoChangeMaxN(vars *Variables, n int) bool {
	for i := 0; i < n; i++ {
		if !d.ExpandOnce(vars) {
			return true
		}
	}
	return false
}
