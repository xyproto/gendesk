package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	icon_search_url = "http://openiconlibrary.sourceforge.net/gallery2/open_icon_library-full/icons/png/48x48/apps/%s.png"
)

// Download a file
func DownloadFile(url string, filename string, o *Output, force bool) error {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		o.ErrText("Could not download " + url)
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrText("Could not dump body")
		os.Exit(1)
	}

	// Check if the file exists (and that force is not enabled)
	if _, err := os.Stat(filename); err == nil && (!force) {
		o.ErrText("no! " + filename + " already exists. Use -f to overwrite.")
		os.Exit(1)
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrText("Could not write to " + filename + "!")
		os.Exit(1)
	}
	return err
}

// Download icon from the search url in icon_search_url
func WriteIconFile(name string, o *Output, force bool) error {
	// Only supports png icons
	filename := name + ".png"
	var client http.Client
	resp, err := client.Get(fmt.Sprintf(icon_search_url, name))
	if err != nil {
		o.ErrText("Could not download icon")
		os.Exit(1)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		o.ErrText("Could not dump body")
		os.Exit(1)
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
		o.ErrText("no! " + filename + " already exists. Use -f to overwrite.")
		os.Exit(1)
	}

	err = ioutil.WriteFile(filename, b, 0666)
	if err != nil {
		o.ErrText("Could not write icon to " + filename + "!")
		os.Exit(1)
	}
	return nil
}
