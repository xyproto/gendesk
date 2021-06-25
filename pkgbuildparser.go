package main

import (
	"io/ioutil"
	"strings"

	"github.com/xyproto/textoutput"
)

const (
	PKGNAME         = "pkgname"
	PKGDESC         = "pkgdesc"
	EXEC            = "_exec"
	NAME            = "_name"
	GENERICNAME     = "_genericname"
	MIMETYPES       = "_mimetypes"
	MIMETYPES_ALIAS = "_mimetype"
	COMMENT         = "_comment"
	CATEGORIES      = "_categories"
	CUSTOM          = "_custom"
	ICON            = "_icon"
	PATH            = "_path"
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

// ParsePKGBUILD will attempt to parse a PKGBUILD file (bash shell script).
// Advanced variable expansion is not supported, yet.
// A map of all found variable values (with their values expanded, once) is returned at the end.
// The given maps are filled with values.
func (vars *Variables) ParsePKGBUILD(o *textoutput.TextOutput, filename string, iconurl, pkgname *string, pkgnames *[]string, pkgdescMap, execMap, nameMap, genericNameMap, mimeTypesMap, commentMap, categoriesMap, customMap *map[string]string) {
	// Fill in the dictionaries using a PKGBUILD
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		o.ErrExit("Could not read " + filename)
	}
	for _, line := range strings.Split(string(filedata), "\n") {
		switch {
		case strings.HasPrefix(line, PKGNAME):
			*pkgname = vars.Expand(betweenQuotesOrAfterEquals(line))
			*pkgnames = pkgList(*pkgname)
			// Select the first pkgname in the array as the "current" pkgname
			if len(*pkgnames) > 0 {
				*pkgname = (*pkgnames)[0]
			}
			(*vars)[PKGNAME] = *pkgname
		case strings.HasPrefix(line, "package_"):
			*pkgname = vars.Expand(between(line, "_", "("))
			(*vars)[PKGNAME] = *pkgname
		case strings.HasPrefix(line, PKGDESC) && *pkgname != "":
			// Description for the package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*pkgdescMap)[*pkgname] = s
			(*vars)[PKGDESC] = s
		case strings.HasPrefix(line, EXEC) && *pkgname != "":
			// Custom executable for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*execMap)[*pkgname] = s
			(*vars)[EXEC] = s
		case strings.HasPrefix(line, NAME) && *pkgname != "":
			// Custom Name for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*nameMap)[*pkgname] = s
			(*vars)[NAME] = s
		case strings.HasPrefix(line, GENERICNAME) && *pkgname != "":
			// Custom GenericName for the .desktop file per (split) package
			genericName := vars.Expand(betweenQuotesOrAfterEquals(line))
			// Use the last found pkgname as the key
			if genericName != "" {
				(*genericNameMap)[*pkgname] = genericName
				(*vars)[GENERICNAME] = genericName
			}
		case strings.HasPrefix(line, MIMETYPES_ALIAS) && *pkgname != "":
			// Custom MimeType for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*mimeTypesMap)[*pkgname] = s
			(*vars)[MIMETYPES_ALIAS] = s
		case strings.HasPrefix(line, MIMETYPES) && *pkgname != "":
			// Custom MimeType for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*mimeTypesMap)[*pkgname] = s
			(*vars)[MIMETYPES] = s
		case strings.HasPrefix(line, COMMENT) && *pkgname != "":
			// Custom Comment for the .desktop file per (split) package
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*commentMap)[*pkgname] = s
			(*vars)[COMMENT] = s
		case strings.HasPrefix(line, CUSTOM) && *pkgname != "":
			// Custom string to be added to the end
			// of the .desktop file in question
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*customMap)[*pkgname] = s
			(*vars)[CUSTOM] = s
		case strings.HasPrefix(line, CATEGORIES) && *pkgname != "":
			// Use the last found pkgname as the key
			s := vars.Expand(betweenQuotesOrAfterEquals(line))
			(*categoriesMap)[*pkgname] = s
			(*vars)[CATEGORIES] = s
		case ((strings.Contains(line, "http://") || strings.Contains(line, "https://")) && strings.Contains(line, ".png")) && *iconurl == "":
			// Only supports detecting png icon filenames when represented as just the filename or an URL starting with http/https.
			*iconurl = vars.Expand("h" + between(line, "h", "g") + "g")
			(*vars)[ICON] = *iconurl
		}

		// Strip the "-bin", "-git", "-hg" or "-svn" suffix, if present
		for _, suf := range []string{"bin", "git", "hg", "svn"} {
			*pkgname = strings.TrimSuffix(*pkgname, "-"+suf)
			(*vars)[PKGNAME] = *pkgname
		}
	}
}
