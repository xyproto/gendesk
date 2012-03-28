package main

import (
	"fmt"
	"flag"
	"strings"
	"io/ioutil"
	"os"
	"http"
	"bytes"
	"path/filepath"
	"crypto/md5"
	"hash"
)

const (
	version_string  = "Desktop File Generator v.0.2"
	icon_search_url = "https://admin.fedoraproject.org/pkgdb/appicon/show/%s"
)

var (
	multimedia_kw  = []string{"video", "audio", "sound", "graphics", "draw", "demo"}
	programming_kw = []string{"code", "c", "ide", "programming", "develop", "compile"}
	network_kw     = []string{"network", "p2p"}
	game_kw        = []string{"racing", "game", "arcade", "rts", "mmorpg", "rpg", "fps", "nintendo emulator"}
)

func createDesktopContents(name string, genericName string, comment string, exec string, icon string, useTerminal bool, categories []string) *bytes.Buffer {
	var buf []byte
	b := bytes.NewBuffer(buf)
	b.WriteString("[Desktop Entry]\n")
	b.WriteString("Encoding=UTF-8\n")
	b.WriteString("Name=" + name + "\n")
	b.WriteString("GenericName=" + genericName + "\n")
	b.WriteString("Comment=" + comment + "\n")
	b.WriteString("Exec=" + exec + "\n")
	b.WriteString("Icon=" + icon + "\n")
	if useTerminal {
		b.WriteString("Terminal=true\n")
	} else {
		b.WriteString("Terminal=false\n")
	}
	b.WriteString("Type=Application\n")
	b.WriteString("Categories=" + strings.Join(categories, ";") + ";\n")
	return b
}

func capitalize(name string) string {
	return strings.ToTitle(name[0:1]) + name[1:]
}

func writeDesktopFile(name string, comment string, exec string, categories string) {
	if len(name) < 1 {
		fmt.Println("No name, can't download icon")
		os.Exit(1)
	}
	if len(categories) == 0 {
		categories = "Application"
	}
	// Only supports png icons
	buf := createDesktopContents(capitalize(name), capitalize(name), comment, exec, name+".png", false, strings.Split(categories, ";"))
	ioutil.WriteFile(name+".desktop", buf.Bytes(), 0666)
}

func startsWith(line string, word string) bool {
	return 0 == strings.Index(strings.TrimSpace(line), word)
}

// Return what's between two strings, "a" and "b", in another string
func between(orig string, a string, b string) string {
	if strings.Contains(orig, a) && strings.Contains(orig, b) {
		posa := strings.Index(orig, a) + len(a)
		posb := strings.LastIndex(orig, b)
		return orig[posa:posb]
	}
	return ""
}

// Return the contents between "" or '' (or an empty string)
func betweenQuotes(orig string) string {
	var s string
	for _, quote := range []string{"\"", "'"} {
		s = between(orig, quote, quote)
		if s != "" {
			return s
		}
	}
	return ""
}

func betweenQuotesOrAfterEquals(orig string) string {
	s := betweenQuotes(orig)
	// If the string is not between quotes, get the text after "="
	if (s == "") && (strings.Count(orig, "=") == 1) {
		s = strings.TrimSpace(strings.Split(orig, "=")[1])
	}
	return s
}

// Does a keyword exist in a lowercase string?
func has(s string, kw string) bool {
	// Replace "-" with " " when searching for keywords
	return -1 != strings.Index(strings.Replace(strings.ToLower(s), "-", " ", -1), " "+kw+" ")
}

func keywordsInDescription(pkgdesc string, keywords []string) bool {
	for _, keyword := range keywords {
		if has(pkgdesc, keyword) {
			return true
		}
	}
	return false
}

// Download icon from the search url in icon_search_url
func writeIconFile(pkgname string) os.Error {
	// Only supports png icons
	filename := pkgname + ".png"
	var client http.Client
	resp, err := client.Get(fmt.Sprintf(icon_search_url, capitalize(pkgname)))
	if err != nil {
		fmt.Println(darkRedText("Could not download icon"))
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(darkRedText("Could not dump body"))
		os.Exit(1)
	}

	var h hash.Hash = md5.New()
	h.Write(b)
	//fmt.Printf("Icon MD5: %x\n", h.Sum())

	// If the icon is the "No icon found" icon (known hash), return with an error
	if fmt.Sprintf("%x", h.Sum()) == "12928aa3233965175ea30f5acae593bf" {
		return os.NewError("No icon found")
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		fmt.Println(darkRedText("Could not write icon to " + filename + "!"))
		os.Exit(1)
	}
	return nil
}

func writeDefaultIconFile(pkgname string) os.Error {
	defaultIconFilename := "/usr/share/pixmaps/default.png"
	b, err := ioutil.ReadFile(defaultIconFilename)
	if err != nil {
		fmt.Println(darkRedText("could not read " + defaultIconFilename + "!"))
		os.Exit(1)
	}
	filename := pkgname + ".png"
	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		fmt.Println(darkRedText("could not write icon to " + filename + "!"))
		os.Exit(1)
	}
	return nil
}

// Download a file
func downloadFile(url string, filename string) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(darkRedText("Could not download file"))
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(darkRedText("Could not dump body"))
		os.Exit(1)
	}
	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		fmt.Println(darkRedText("Could not write data to " + filename + "!"))
		os.Exit(1)
	}
}

// Use a function for each element in a string list and
// return the modified list
func stringMap(f func(string) string, stringlist []string) []string {
	newlist := make([]string, len(stringlist))
	for i, elem := range stringlist {
		newlist[i] = f(elem)
	}
	return newlist
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
		unquoted = strings.Replace(center, "'", "", -1)
		return strings.Split(unquoted, " ")
	}
	return []string{splitpkgname}
}

func darkRedText(s string) string {
	return "\033[0;31m" + s + "\033[0m"
}

func lightGreenText(s string) string {
	return "\033[1;32m" + s + "\033[0m"
}

func darkGreenText(s string) string {
	return "\033[0;32m" + s + "\033[0m"
}

func lightYellowText(s string) string {
	return "\033[1;33m" + s + "\033[0m"
}

func darkYellowText(s string) string {
	return "\033[0;33m" + s + "\033[0m"
}

func lightBlueText(s string) string {
	return "\033[1;34m" + s + "\033[0m"
}

func darkBlueText(s string) string {
	return "\033[0;34m" + s + "\033[0m"
}

func lightCyanText(s string) string {
	return "\033[1;36m" + s + "\033[0m"
}

func lightPurpleText(s string) string {
	return "\033[1;35m" + s + "\033[0m"
}

func darkPurpleText(s string) string {
	return "\033[0;35m" + s + "\033[0m"
}

func darkGrayText(s string) string {
	return "\033[1;30m" + s + "\033[0m"
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var filename string
	flag.Usage = func() {
		fmt.Println()
		fmt.Println(version_string)
		fmt.Println("generates .desktop files from a PKGBUILD")
		fmt.Println()
		fmt.Println("Syntax: gendesk filename")
		fmt.Println()
		fmt.Println("Note:")
		fmt.Println("  * \"../PKGBUILD\" is the default filename")
		fmt.Println("  * _exec in the PKGBUILD can be used to specific a different executable for the .desktop file")
		fmt.Println("    Example: _exec=('appname-gui')")
		fmt.Println("  * Split packages are supported")
		fmt.Println("  * If a .png icon is not found as a file or in the PKGBUILD, an icon will be downloaded from:")
		fmt.Println("    " + icon_search_url)
		fmt.Println("    This may or may not result in the icon you wished for.")
		fmt.Println("  * Categories are guessed based on keywords in the package description")
		fmt.Println("  * Icons are assumed to be installed to \"/usr/share/pixmaps/$pkgname.png\" by the PKGBUILD")
		fmt.Println()
	}
	var version *bool = flag.Bool("version", false, version_string)
	flag.Parse()
	args := flag.Args()
	if *version {
		fmt.Println(version_string)
		os.Exit(0)
	} else if len(args) == 0 {
		filename = "../PKGBUILD"
	} else if len(args) == 1 {
		filename = args[0]
	} else {
		fmt.Println("too many arguments")
		os.Exit(1)
	}
	filedata, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, darkRedText("Could not read %s\n"), filename)
		os.Exit(1)
	}
	filetext := string(filedata)

	var pkgname string
	var pkgnames []string
	var iconurl string
	pkgdescMap := make(map[string]string)
	execMap := make(map[string]string)

	for _, line := range strings.Split(filetext, "\n") {
		if startsWith(line, "pkgname") {
			pkgname = betweenQuotesOrAfterEquals(line)
			pkgnames = pkgList(pkgname)
			// Select the first pkgname in the array as the "current" pkgname
			if len(pkgnames) > 0 {
				pkgname = pkgnames[0]
			}
		} else if startsWith(line, "package_") {
			pkgname = between(line, "_", "(")
		} else if startsWith(line, "pkgdesc") {
			// Description for the package
			pkgdesc := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if pkgname != "" {
				pkgdescMap[pkgname] = pkgdesc
			}
		} else if startsWith(line, "_exec") {
			// Custom executable for the .desktop file per (split) package
			exec := betweenQuotesOrAfterEquals(line)
			// Use the last found pkgname as the key
			if pkgname != "" {
				execMap[pkgname] = exec
			}
		} else if strings.Contains(line, "http://") && strings.Contains(line, ".png") {
			// Only supports png icons downloaded over http, picks the first fitting url
			if iconurl == "" {
				iconurl = "h" + between(line, "h", "g") + "g"
				if strings.Contains(iconurl, "$pkgname") {
					iconurl = strings.Replace(iconurl, "$pkgname", pkgname, -1)
				}
				if strings.Contains(iconurl, "${pkgname}") {
					iconurl = strings.Replace(iconurl, "${pkgname}", pkgname, -1)
				}
				if strings.Contains(iconurl, "$") {
					// Will only replace pkgname. There are more replacements
					iconurl = ""
				}
			}
		}
	}

	//fmt.Println("pkgnames:", pkgnames)

	// Write .desktop and .png icon for each package
	for _, pkgname := range pkgnames {
		if strings.Contains(pkgname, "-nox") || strings.Contains(pkgname, "-cli") {
			// Don't bother if it's a -nox or -cli package
			continue
		}
		pkgdesc, found := pkgdescMap[pkgname]
		if !found {
			// Fall back on the package name
			pkgdesc = pkgname
		}
		exec, found := execMap[pkgname]
		if !found {
			// Fall back on the package name
			exec = pkgname
		}
		// Approximately identify various categories
		categories := ""
		if keywordsInDescription(pkgdesc, multimedia_kw) {
			categories = "Application;Multimedia"
		} else if keywordsInDescription(pkgdesc, programming_kw) {
			categories = "Application;Development"
		} else if keywordsInDescription(pkgdesc, network_kw) {
			categories = "Application;Network"
		} else if keywordsInDescription(pkgdesc, game_kw) {
			categories = "Application;Game"
		}
		const nSpaces = 32
		spaces := strings.Repeat(" ", nSpaces)[:nSpaces-min(nSpaces, len(pkgname))]
		fmt.Printf("%s%s%s%s%s ", darkGrayText("["), lightBlueText(pkgname), darkGrayText("]"), spaces, darkGrayText("Generating desktop file..."))
		writeDesktopFile(pkgname, pkgdesc, exec, categories)
		fmt.Printf("%s\n", darkGreenText("ok"))

		// Download an icon if it's not downloaded by the PKGBUILD and not there already
		files, _ := filepath.Glob("*.png")
		if (len(files) == 0) && (iconurl == "") {
			fmt.Printf("%s%s%s%s%s ", darkGrayText("["), lightBlueText(pkgname), darkGrayText("]"), spaces, darkGrayText("Downloading icon..."))
			err := writeIconFile(pkgname)
			if err == nil {
				fmt.Printf("%s\n", lightCyanText("ok"))
			} else {
				fmt.Printf("%s\n", darkYellowText("no"))
				fmt.Printf("%s%s%s%s%s ", darkGrayText("["), lightBlueText(pkgname), darkGrayText("]"), spaces, darkGrayText("Using default icon instead..."))
				err := writeDefaultIconFile(pkgname)
				if err == nil {
					fmt.Printf("%s\n", lightPurpleText("yes"))
				}
			}
		}
	}
}
