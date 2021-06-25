package main

import (
	"strings"

	"github.com/xyproto/env"
)

// Variables is a collection of variables from either the environment
// PKGBUILD fields or given arguments, that can be expanded. ("$x" and "${x}")
type Variables map[string]string

// EnvToVars will take values from all environment variables with matching names
// and place them as keys+values in the given vars map.
func EnvToVars(keys []string, vars *Variables) {
	for _, k := range keys {
		if v := env.Str(k); v != "" {
			vars.Set(k, v)
		}
	}
}

// NewVars will create a new Variables map, reading the given keys from the environment
func NewVars(keys []string) *Variables {
	var vars Variables
	EnvToVars(keys, &vars)
	return &vars
}

// Expand performs simple variable expansion, using the available keys and values
// in the Variables map.
// Advanced bash variable expensions like ${x:1:2} or ${x%.*} are not handled.
func (vars *Variables) Expand(s string) string {
	if s == "" {
		return s
	}
	// simple variable expansion
	for k, v := range *vars {
		s = strings.Replace(s, "$"+k, v, -1)
		s = strings.Replace(s, "${"+k+"}", v, -1)
	}
	return s
}

// ExpandAll will expand all variables, complexity O(n^2)
func (vars *Variables) ExpandAll() {
	for k, v := range *vars {
		(*vars)[k] = vars.Expand(v)
	}
}

// Set a key+value, initializing the map if needed
func (vars *Variables) Set(k, v string) {
	if *vars == nil {
		*vars = make(Variables)
	}
	(*vars)[k] = v
}

// Has checks if a variables is set and non-empty
func (vars *Variables) Has(k string) bool {
	v, ok := (*vars)[k]
	return ok && v != ""
}

// Get returns the value, or the given defaultValue if it's not available
func (vars *Variables) Get(k, defaultValue string) string {
	if v, ok := (*vars)[k]; ok && v != "" {
		return v
	}
	return defaultValue
}
