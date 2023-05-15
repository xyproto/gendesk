package main

import (
	"os"
	"strings"

	"github.com/xyproto/env/v2"
	"github.com/xyproto/textoutput"
)

// Return a list of pkgnames for split packages
// or just a list with the pkgname for regular packages
func pkgList(splitpkgname string) []string {
	center := between(splitpkgname, "(", ")")
	if center == "" {
		center = splitpkgname
	}
	if strings.Contains(center, " ") {
		unquoted := strings.Replace(center, "\"", "", -1)
		unquoted = strings.Replace(unquoted, "'", "", -1)
		return strings.Split(unquoted, " ")
	}
	return []string{splitpkgname}
}

// fromEnvIfEmpty will retrieve a value from the environment,
// but only if the given value is empty
func fromEnvIfEmpty(field *string, envVarName string) {
	if *field == "" {
		*field = env.Str(envVarName)
	}
}

func dataFromEnvironment(pkgdesc, execCommand, name, genericname, mimetypes, comment, categories, custom *string) {
	// Environment variables
	fromEnvIfEmpty(pkgdesc, "pkgdesc")
	fromEnvIfEmpty(execCommand, "_exec")
	fromEnvIfEmpty(name, "_name")
	fromEnvIfEmpty(genericname, "_genericname")
	fromEnvIfEmpty(mimetypes, "_mimetypes")
	fromEnvIfEmpty(mimetypes, "_mimetype")
	fromEnvIfEmpty(comment, "_comment")
	fromEnvIfEmpty(categories, "_categories")
	fromEnvIfEmpty(custom, "_custom")
}

// resolve will expand variables within the given string,
// using the supplied map as a source of keys and values.
// Advanced bash variable expensions like ${x:1:2} or ${x%.*} are not handled.
func resolve(vars map[string]string, s *string) {
	// simple variable expansion
	for k, v := range vars {
		if strings.Contains(*s, "$"+k) {
			*s = strings.Replace(*s, "$"+k, v, -1)
		} else if strings.Contains(*s, "${"+k+"}") {
			*s = strings.Replace(*s, "${"+k+"}", v, -1)
		}
	}
}

func parsePKGBUILD(o *textoutput.TextOutput, filename string, iconurl, pkgname *string, pkgnames *[]string, pkgdescMap, execMap, nameMap, genericNameMap, mimeTypesMap, commentMap, categoriesMap, customMap *map[string]string) {
	// Fill in the dictionaries using a PKGBUILD
	filedata, err := os.ReadFile(filename)
	if err != nil {
		o.ErrExit("Could not read " + filename)
	}
	vars := make(map[string]string) // variables found along the way
	for _, line := range strings.Split(string(filedata), "\n") {
		switch {
		case strings.HasPrefix(line, "pkgname"):
			*pkgname = betweenQuotesOrAfterEquals(line)
			*pkgnames = pkgList(*pkgname)
			// Select the first pkgname in the array as the "current" pkgname
			if len(*pkgnames) > 0 {
				*pkgname = (*pkgnames)[0]
			}
			resolve(vars, pkgname)
			vars["pkgname"] = *pkgname
		case strings.HasPrefix(line, "package_"):
			*pkgname = between(line, "_", "(")
			resolve(vars, pkgname)
			vars["pkgname"] = *pkgname
		case strings.HasPrefix(line, "pkgdesc") && *pkgname != "":
			// Description for the package
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*pkgdescMap)[*pkgname] = s
			vars["pkgdesc"] = s
		case strings.HasPrefix(line, "_exec") && *pkgname != "":
			// Custom executable for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)

			(*execMap)[*pkgname] = s
			vars["_exec"] = s
		case strings.HasPrefix(line, "_name") && *pkgname != "":
			// Custom Name for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*nameMap)[*pkgname] = s
			vars["_name"] = s
		case strings.HasPrefix(line, "_genericname") && *pkgname != "":
			// Custom GenericName for the .desktop file per (split) package
			genericName := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if genericName != "" {
				resolve(vars, &genericName)
				(*genericNameMap)[*pkgname] = genericName
				vars["_genericname"] = genericName
			}
		case strings.HasPrefix(line, "_mimetype") && *pkgname != "":
			// Custom MimeType for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*mimeTypesMap)[*pkgname] = s
			vars["_mimetype"] = s
		case strings.HasPrefix(line, "_comment") && *pkgname != "":
			// Custom Comment for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*commentMap)[*pkgname] = s
			vars["_comment"] = s
		case strings.HasPrefix(line, "_custom") && *pkgname != "":
			// Custom string to be added to the end
			// of the .desktop file in question
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*customMap)[*pkgname] = s
			vars["_custom"] = s
		case strings.HasPrefix(line, "_categories") && *pkgname != "":
			// Use the last found pkgname as the key
			s := betweenQuotesOrAfterEquals(line)
			resolve(vars, &s)
			(*categoriesMap)[*pkgname] = s
			vars["_categories"] = s
		case ((strings.Contains(line, "http://") || strings.Contains(line, "https://")) && (strings.Contains(line, ".png") || strings.Contains(line, ".svg"))) && *iconurl == "":
			// Only supports detecting png icon filenames when represented as just the filename or an URL starting with http/https.
			*iconurl = betweenInclusive(line, "h", "g")
			resolve(vars, iconurl)
			vars["_icon"] = *iconurl
		}

		// Strip the "-bin", "-git", "-hg" or "-svn" suffix, if present
		for _, suf := range []string{"bin", "git", "hg", "svn"} {
			*pkgname = strings.TrimSuffix(*pkgname, "-"+suf)
		}
	}
}
