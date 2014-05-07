package main

import (
	"code.google.com/p/goconf/conf"
	"crypto/md5"
	"errors"
	"fmt"
	. "github.com/xyproto/textgui"
	"hash"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

const (
	default_icon_search_url = "http://openiconlibrary.sourceforge.net/gallery2/open_icon_library-full/icons/png/48x48/apps/%s.png"
)

// Download a file
func DownloadFile(url string, filename string, o *TextOutput, force bool) error {
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
func GetIconSearchURL(o *TextOutput) string {
	usr, err := user.Current()
	if err != nil {
		return default_icon_search_url
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
				return default_icon_search_url
			}
		}
	}

	// Found a configuration file, find the url under the [default] section with the key icon_search_url
	icon_url, err := cfile.GetString("default", "icon_search_url")
	if err != nil {
		return default_icon_search_url
	}

	if strings.Contains(icon_url, "%s") {
		o.Err(cfilename + " does not contain an icon search url containing %s under a [default] section. Example:")
		o.LightGreen("[default]")
		o.LightGreen("icon_search_url = http://some.iconrepository.com/%s.png")
		os.Exit(1)
	}

	// Found an url in the configuration file, use that instead of the default search url
	return icon_url
}

// Download icon from the search url in icon_search_url
func WriteIconFile(name string, o *TextOutput, force bool) error {
	icon_search_url := GetIconSearchURL(o)
	// Only supports png icons
	filename := name + ".png"
	var client http.Client
	resp, err := client.Get(fmt.Sprintf(icon_search_url, name))
	if err != nil {
		o.ErrExit("Could not download icon")
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrExit("Could not dump body")
	}
	var h hash.Hash = md5.New()
	h.Write(b)
	//fmt.Printf("Icon MD5: %x\n", h.Sum())

	// If the icon is the "No icon found" icon (known hash), return with an error
	if fmt.Sprintf("%x", h.Sum(nil)) == "12928aa3233965175ea30f5acae593bf" {
		return errors.New("No icon found")
	}

	if b[0] == 60 && b[1] == 104 && b[2] == 116 {
		// if it starts with "<ht", it's not a png
		return errors.New("No icon found")
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrExit("no! " + filename + " already exists. Use -f to overwrite.")
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrExit("Could not write icon to " + filename + "!")
	}
	return nil
}
