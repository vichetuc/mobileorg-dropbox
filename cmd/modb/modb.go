// Copyright 2014 Lenilson Jose Dias. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	EnvMobileOrgToken  = "MODB_TOKEN"
	EnvMobileOrgLocal  = "MODB_LOCAL_DIR"
	EnvMobileOrgRemote = "MODB_REMOTE_DIR"
)

var Usage = func() {
	usageText := `
Usage:

      modb -upload
      modb -download

modbget credentials can be this environment variables:

      MODB_TOKEN      - Dropbox Oauth2AccessToken for Mobileorg App.
      MODB_LOCAL_DIR  - Local MobileOrg directory.
      MODB_REMOTE_DIR - Remote MobileOrg directory.

Options:

`
	fmt.Fprintf(os.Stderr, usageText)
	flag.PrintDefaults()
}

var (
	mobtoken  = flag.String("t", "", "Dropbox app Oauth2AccessToken for MobileOrg.")
	moblocal  = flag.String("l", "", "Local Mobileorg directory.")
	mobremote = flag.String("r", "", "Remote MobileOrg directory;")
	upload    = flag.Bool("upload", false, "Upload MobileOrg files to Dropbox app folder")
	download  = flag.Bool("download", false, "Download MobileOrg Files from Dropbox app foder")
)

func main() {
	flag.Usage = Usage
	flag.Parse()

	if !ValidFlags() {
		flag.Usage()
		os.Exit(128)
	}

	if *upload {
		fmt.Println("Uploading...")
	} else {
		fmt.Println("Downloading..")
	}
}

func getFlagEnv(name *string, envname string) bool {
	if *name == "" {
		*name = os.Getenv(envname)
		if *name == "" {
			return false
		}
	}
	return true
}

func ValidFlags() bool {
	if !getFlagEnv(mobtoken, EnvMobileOrgToken) {
		return false
	}

	if !getFlagEnv(mobremote, EnvMobileOrgRemote) {
		return false
	}

	if !getFlagEnv(moblocal, EnvMobileOrgLocal) {
		return false
	}

	if (!*upload && !*download) || (*upload && *download) {
		return false
	}

	return true
}
