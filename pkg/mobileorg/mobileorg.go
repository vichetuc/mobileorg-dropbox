// Copyright 2014 Lenilson Jose Dias. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Package mobileorg implements an interface with MobileOrg
package mobileorg

import (
	"bufio"
	"os"
	"strings"
)

// Get MobileOrg Files from diretory
func GetFiles(dir string) ([]string, error) {
	checksums, err := os.Open(dir + "/checksums.dat")
	if err != nil {
		return nil, err
	}
	defer checksums.Close()

	files := []string{dir + "/checksums.dat"}
	scanner := bufio.NewScanner(checksums)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		f := strings.Split(scanner.Text(), " ")
		files = append(files, dir+"/"+f[2])
	}
	return files, nil
}
