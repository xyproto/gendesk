package main

import (
	"io/ioutil"
	"os"
	"strings"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

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

func dataFromEnvironment(pkgdesc, exec, name, genericname, mimetypes, comment, categories, custom *string) {
	// Environment variables
	if *pkgdesc == "" {
		// $pkgdesc is either empty or not
		*pkgdesc = os.Getenv("pkgdesc")
	}
	if *exec == "" {
		*exec = os.Getenv("_exec")
	}
	if *name == "" {
		*name = os.Getenv("_name")
	}
	if *genericname == "" {
		*genericname = os.Getenv("_genericname")
	}
	if *mimetypes == "" {
		*mimetypes = os.Getenv("_mimetypes")
	}
	// support "_mimetype" as well (deprecated)
	if *mimetypes == "" {
		*mimetypes = os.Getenv("_mimetype")
	}
	if *comment == "" {
		*comment = os.Getenv("_comment")
	}
	if *categories == "" {
		*categories = os.Getenv("_categories")
	}
	if *custom == "" {
		*custom = os.Getenv("_custom")
	}
}

func parsePKGBUILD(o *Output, filename string, iconurl *string, pkgname *string, pkgnames *[]string, pkgdescMap, execMap, nameMap, genericNameMap, mimeTypesMap, commentMap, categoriesMap, customMap *map[string]string) {
	// Fill in the dictionaries using a PKGBUILD
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		o.ErrText("Could not read " + filename)
		os.Exit(1)
	}
	filetext := string(filedata)
	for _, line := range strings.Split(filetext, "\n") {
		// TODO: Use a loop instead of "if / else if / else if"
		if startsWith(line, "pkgname") {
			*pkgname = betweenQuotesOrAfterEquals(line)
			*pkgnames = pkgList(*pkgname)
			// Select the first pkgname in the array as the "current" pkgname
			if len(*pkgnames) > 0 {
				*pkgname = (*pkgnames)[0]
			}
		} else if startsWith(line, "package_") {
			*pkgname = between(line, "_", "(")
		} else if startsWith(line, "pkgdesc") {
			// Description for the package
			pkgdesc := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*pkgdescMap)[*pkgname] = pkgdesc
			}
		} else if startsWith(line, "_exec") {
			// Custom executable for the .desktop file per (split) package
			exec := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*execMap)[*pkgname] = exec
			}
		} else if startsWith(line, "_name") {
			// Custom Name for the .desktop file per (split) package
			name := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*nameMap)[*pkgname] = name
			}
		} else if startsWith(line, "_genericname") {
			// Custom GenericName for the .desktop file per (split) package
			genericName := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if (*pkgname != "") && (genericName != "") {
				(*genericNameMap)[*pkgname] = genericName
			}
		} else if startsWith(line, "_mimetype") {
			// Custom MimeType for the .desktop file per (split) package
			mimeType := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*mimeTypesMap)[*pkgname] = mimeType
			}
		} else if startsWith(line, "_comment") {
			// Custom Comment for the .desktop file per (split) package
			comment := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*commentMap)[*pkgname] = comment
			}
		} else if startsWith(line, "_custom") {
			// Custom string to be added to the end
			// of the .desktop file in question
			custom := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*customMap)[*pkgname] = custom
			}
		} else if startsWith(line, "_categories") {
			categories := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if *pkgname != "" {
				(*categoriesMap)[*pkgname] = categories
			}
		} else if strings.Contains(line, "http://") && strings.Contains(line, ".png") {
			// Only supports png icons downloaded over http,
			// picks the first fitting url
			if *iconurl == "" {
				*iconurl = "h" + between(line, "h", "g") + "g"
				if strings.Contains(*iconurl, "$pkgname") {
					*iconurl = strings.Replace(*iconurl,
						"$pkgname", *pkgname, -1)
				}
				if strings.Contains(*iconurl, "${pkgname}") {
					*iconurl = strings.Replace(*iconurl,
						"${pkgname}", *pkgname, -1)
				}
				if strings.Contains(*iconurl, "$") {
					// If there are more $variables, don't bother (for now)
					// TODO: replace all defined $variables...
					*iconurl = ""
				}
			}
		}
	}

}
