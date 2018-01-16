package main

import (
	"bytes"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/akrennmair/goconf"
	"github.com/xyproto/term"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

const (
	defaultIconSearchURL = "http://www.iconarchive.com/search?q=%s&res=48&page=1&sort=popularity"
)

// Download a file
func DownloadFile(url, filename string, o *term.TextOutput, force bool) error {
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

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrExit("Could not write to " + filename + "!")
	}
	return err
}

// Find the icon search url (must contain %s) from the first found configuration file
func GetIconSearchURL(o *term.TextOutput) string {
	usr, err := user.Current()
	if err != nil {
		return defaultIconSearchURL
	}
	homedir := usr.HomeDir

	// Read the configuration file from various locations,
	cfilename := "~/.gendeskrc"
	cfile, err := conf.ReadConfigFile(strings.Replace(cfilename, "~", homedir, -1))
	if err != nil {
		cfilename = "~/.config/gendesk"
		cfile, err = conf.ReadConfigFile(strings.Replace(cfilename, "~", homedir, -1))
		if err != nil {
			cfilename = "/etc/gendeskrc"
			conf.ReadConfigFile(cfilename)
			if err != nil {
				return defaultIconSearchURL
			}
		}
	}

	// Found a configuration file, find the url under the [default] section with the key iconSearchURL
	icon_url, err := cfile.GetString("default", "icon_url")
	if err != nil {
		o.Err("error!\n")
		o.Println(o.DarkRed(cfilename + " does not contain icon_url under under a [default] section. Example:"))
		o.Println(o.LightGreen("[default]"))
		o.Println(o.LightGreen("icon_url = http://some.iconrepository.com/q=%s.png\n"))
		os.Exit(1)
	}

	if !strings.Contains(icon_url, "%s") {
		o.Err("error!\n")
		o.Println(o.DarkRed(cfilename + " does not contain an icon search url containing %s under a [default] section. Example:"))
		o.Println(o.LightGreen("[default]"))
		o.Println(o.LightGreen("icon_url = http://some.iconrepository.com/q=%s.png\n"))
		os.Exit(1)
	}

	// Found an url in the configuration file, use that instead of the default search url
	return icon_url
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

// Download icon from the search url in iconSearchURL, or from iconarchive.
// Only supports downloading png icons.
func WriteIconFile(name string, o *term.TextOutput, force bool) error {
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

	resp, err := client.Get(downloadURL)
	if err != nil {
		o.ErrExit("Could not download icon")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrExit("Could not dump body")
	}

	// If the icon is the "No icon found" icon (known hash), return with an error
	h := md5.New()
	h.Write(b)
	if fmt.Sprintf("%x", h.Sum(nil)) == "12928aa3233965175ea30f5acae593bf" {
		return errors.New("No icon found")
	}

	PNGheader := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
	if !bytes.HasPrefix(b, PNGheader) {
		return errors.New("No PNG icon found")
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrExit("no! " + filename + " already exists. Use -f to overwrite.")
	}

	err = ioutil.WriteFile(filename, b, 0644)
	if err != nil {
		o.ErrExit("Could not write icon to " + filename + "!")
	}
	return nil
}
