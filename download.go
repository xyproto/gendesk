package main

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"

	"github.com/unknwon/goconfig"
	"github.com/xyproto/textoutput"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
)

const (
	defaultIconSearchURL = "http://www.iconarchive.com/search?q=%s&res=48&page=1&sort=popularity"
)

var (
	errNoDownloadURL = errors.New("no icon download URL")
	errNoIconFound   = errors.New("no icon found")
	errNotPNG        = errors.New("no PNG icon found")
)

// MustDownloadFile takes an URL to a filename and attempts to download it to the given filename
// If force is true, any existing file may be overwritten.
// May exit the program if there are fundamental problems.
func MustDownloadFile(url, filename string, o *textoutput.TextOutput, force bool) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		o.ErrExit("Could not download " + url)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrExit("Could not dump body")
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrExit("no! " + filename + " already exists. Use -f to overwrite.")
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		o.ErrExit("Could not write to " + filename + "!")
	}
}

// Expand ~ to $HOME, or return the given string
func userexpand(path string) string {
	u, err := user.Current()
	if err != nil {
		return path
	}
	return strings.Replace(path, "~", u.HomeDir, -1)
}

// GetIconSearchURL reads configuration from ~/.gendeskrc, ~/.config/gendesk or /etc/gendeskrc
// in order to retrieve an URL containing "%s" that can be used for searching for icons by name.
// May exit the program if there are fundamental problems.
func GetIconSearchURL(o *textoutput.TextOutput) string {
	// Read the configuration file from various locations,
	cfilename := "~/.config/gendesk"
	cfile, err := goconfig.LoadConfigFile(userexpand(cfilename))
	if err != nil {
		cfilename = "~/.gendeskrc"
		cfile, err = goconfig.LoadConfigFile(userexpand(cfilename))
		if err != nil {
			cfilename = "/etc/gendeskrc"
			cfile, err = goconfig.LoadConfigFile(cfilename)
			if err != nil {
				return defaultIconSearchURL
			}
		}
	}

	// Found a configuration file, find the url under the [default] section with the key iconSearchURL
	iconURL, err := cfile.GetValue("default", "icon_url")
	if err != nil {
		o.Err("error!\n")
		o.Fprintln(os.Stderr, o.DarkRed(cfilename+" does not contain icon_url under under a [default] section. Example:"))
		o.Fprintln(os.Stderr, o.LightGreen("[default]"))
		o.Fprintln(os.Stderr, o.LightGreen("icon_url = http://example.iconrepository.com/q=%s.png\n"))
		os.Exit(1)
	}

	if !strings.Contains(iconURL, "%s") {
		o.Err("error!\n")
		o.Fprintln(os.Stderr, o.DarkRed(cfilename+" does not contain an icon search url containing %s under a [default] section. Example:"))
		o.Fprintln(os.Stderr, o.LightGreen("[default]"))
		o.Fprintln(os.Stderr, o.LightGreen("icon_url = http://example.iconrepository.com/q=%s.png\n"))
		os.Exit(1)
	}

	// Found an url in the configuration file, use that instead of the default search url
	return iconURL
}

// findIconURL searches the given iconarchive-compatible URL for a keyword and returns an URL to the PNG image.
// Returns an empty string if no match was found. nmatch is the desired icon if serveral are found.
func findIconURL(searchURL, keyword string, nmatch int) (URL string) {
	// request and parse the front page
	resp, err := http.Get(fmt.Sprintf(searchURL, keyword))
	if err != nil {
		return
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		return
	}

	// Count the number of icons found to be able to pick out the Nth
	counter := 0

	// Find all tags with the class "icondetail"
	matcher := scrape.ByClass("icondetail")
	iconDetails := scrape.FindAll(root, matcher)
	for _, iconDetail := range iconDetails {
		// Find all tags with the class "lastitem" (these are also "a" tags)
		atags := scrape.FindAll(iconDetail, scrape.ByClass("lastitem"))
		for _, atag := range atags {
			// Sanity check: check that the text is "PNG"
			if scrape.Text(atag) == "PNG" {
				// Found one!
				URL = scrape.Attr(atag, "href")
				counter++
				// Use the Nth match, if we get this far (if not, use the last matched URL)
				if counter >= nmatch {
					return
				}
			}
		}
	}

	// No match found
	return
}

// WriteIconFile will search for and download an icon, using the icon search
// URL given in the configuration file, or from iconarchive.com.
// Only supports downloading png icons.
// May exit the program if there are fundamental problems.
func WriteIconFile(name string, o *textoutput.TextOutput, force bool) error {
	var (
		downloadURL   string
		client        http.Client
		iconSearchURL = GetIconSearchURL(o)
		filename      = name + ".png"
	)

	// Use different methods for different icon archives
	if strings.Contains(iconSearchURL, "iconarchive.com/") {
		downloadURL = findIconURL(iconSearchURL, name, 3)
	} else {
		downloadURL = fmt.Sprintf(iconSearchURL, name)
	}

	if downloadURL == "" {
		return errNoDownloadURL
	}

	resp, err := client.Get(downloadURL)
	if err != nil {
		// Could not download the icon, for some reason
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrExit("could not dump body")
	}

	// If the icon is the "No icon found" icon (known hash), return with an error
	h := md5.New()
	h.Write(b)
	if fmt.Sprintf("%x", h.Sum(nil)) == "12928aa3233965175ea30f5acae593bf" {
		return errNoIconFound
	}

	PNGheader := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	if !bytes.HasPrefix(b, PNGheader) {
		return errNotPNG
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrExit(filename + " already exists, use -f to overwrite")
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		o.ErrExit("Could not write icon to: " + filename)
	}
	return nil
}
