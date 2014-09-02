// Copyright 2014 Lenilson Jose Dias. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package dropbox implements the interface with Dropbox API.
//
// This package uses the library: [https://github.com/scottferg/Dropbox-Go]
package dropbox

import (
	"io/ioutil"
	"path"

	"github.com/scottferg/Dropbox-Go/dropbox"
	"os"
)

// Get URI to remote Dropbox file.
func getUri(filename string, path string) dropbox.Uri {
	uriPath := dropbox.Uri{
		Root: dropbox.RootDropbox,
		Path: path + filename,
	}

	return uriPath
}

// Upload a file using dropbox connection
func UploadFile(ds dropbox.Session, local string, remote string) (bool, error) {
	fname := path.Base(local)
	uriPath := getUri(fname, remote)

	if file, err := ioutil.ReadFile(local); err != nil {
		return false, err
	} else {
		_, err := dropbox.UploadFile(ds, file, uriPath, nil)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// Download a file using dropbox
func DownloadFile(ds dropbox.Session, local string, remote string) (bool, error) {
	fname := path.Base(local)
	uriPath := getUri(fname, remote)

	f, _, err := dropbox.GetFile(ds, uriPath, nil)
	if err != nil {
		return false, err
	}
	err = ioutil.WriteFile(local, f, os.ModePerm)
	if err != nil {
		return false, err
	}

	return true, nil
}
