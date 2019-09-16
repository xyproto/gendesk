package main

import (
	"github.com/xyproto/textoutput"
	"io/ioutil"
	"os"
	"strings"
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

// Retrieve value from environment if the given value is empty
func fromEnvIfEmpty(field *string, envVarName string) {
	if *field == "" {
		*field = os.Getenv(envVarName)
	}
}

func dataFromEnvironment(pkgdesc, exec, name, genericname, mimetypes, comment, categories, custom *string) {
	// Environment variables
	fromEnvIfEmpty(pkgdesc, "pkgdesc")
	fromEnvIfEmpty(exec, "_exec")
	fromEnvIfEmpty(name, "_name")
	fromEnvIfEmpty(genericname, "_genericname")
	fromEnvIfEmpty(mimetypes, "_mimetypes")
	fromEnvIfEmpty(mimetypes, "_mimetype")
	fromEnvIfEmpty(comment, "_comment")
	fromEnvIfEmpty(categories, "_categories")
	fromEnvIfEmpty(custom, "_custom")
}

func parsePKGBUILD(o *textoutput.TextOutput, filename string, iconurl, pkgname *string, pkgnames *[]string, pkgdescMap, execMap, nameMap, genericNameMap, mimeTypesMap, commentMap, categoriesMap, customMap *map[string]string) {
	// Fill in the dictionaries using a PKGBUILD
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		o.ErrExit("Could not read " + filename)
	}
	for _, line := range strings.Split(string(filedata), "\n") {
		switch {
		case strings.HasPrefix(line, "pkgname"):
			*pkgname = betweenQuotesOrAfterEquals(line)
			*pkgnames = pkgList(*pkgname)
			// Select the first pkgname in the array as the "current" pkgname
			if len(*pkgnames) > 0 {
				*pkgname = (*pkgnames)[0]
			}
		case strings.HasPrefix(line, "package_"):
			*pkgname = between(line, "_", "(")
		case strings.HasPrefix(line, "pkgdesc") && *pkgname != "":
			// Description for the package
			// Use the last found pkgname as the key
			(*pkgdescMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_exec") && *pkgname != "":
			// Custom executable for the .desktop file per (split) package
			// Use the last found pkgname as the key
			(*execMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_name") && *pkgname != "":
			// Custom Name for the .desktop file per (split) package
			// Use the last found pkgname as the key
			(*nameMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_genericname") && *pkgname != "":
			// Custom GenericName for the .desktop file per (split) package
			genericName := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if genericName != "" {
				(*genericNameMap)[*pkgname] = genericName
			}
		case strings.HasPrefix(line, "_mimetype") && *pkgname != "":
			// Custom MimeType for the .desktop file per (split) package
			// Use the last found pkgname as the key
			(*mimeTypesMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_comment") && *pkgname != "":
			// Custom Comment for the .desktop file per (split) package
			// Use the last found pkgname as the key
			(*commentMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_custom") && *pkgname != "":
			// Custom string to be added to the end
			// of the .desktop file in question
			// Use the last found pkgname as the key
			(*customMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case strings.HasPrefix(line, "_categories") && *pkgname != "":
			// Use the last found pkgname as the key
			(*categoriesMap)[*pkgname] = betweenQuotesOrAfterEquals(line)
		case ((strings.Contains(line, "http://") || strings.Contains(line, "https://")) && strings.Contains(line, ".png")) && *iconurl == "":
			// Only supports detecting png icon filenames when represented as just the filename or an URL starting with http/https.
			*iconurl = "h" + between(line, "h", "g") + "g"
			if strings.Contains(*iconurl, "$pkgname") {
				*iconurl = strings.Replace(*iconurl,
					"$pkgname", *pkgname, -1)
			} else if strings.Contains(*iconurl, "${pkgname}") {
				*iconurl = strings.Replace(*iconurl,
					"${pkgname}", *pkgname, -1)
			} else if strings.Contains(*iconurl, "$") {
				// If there are more $variables, don't bother (for now)
				// TODO: replace all defined $variables...
				*iconurl = ""
			}
		}

		// Strip the "-git", "-svn" or "-hg" suffix, if present
		if strings.HasSuffix(*pkgname, "-git") || strings.HasSuffix(*pkgname, "-svn") {
			*pkgname = (*pkgname)[:len(*pkgname)-4]
		} else if strings.HasSuffix(*pkgname, "-hg") {
			*pkgname = (*pkgname)[:len(*pkgname)-3]
		}

	}
}
