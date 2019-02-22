// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"

	"github.com/cpmech/gosl/chk"
)

var name = flag.String("name", "", "Example name. Use empty to run all examples")

var examples = map[string]func(){
	"ang01":    ang01,
	"ang02":    ang02,
	"kmeans01": kmeans01,
	"mclass01": mclass01,
	"simple01": simple01,
}

func main() {
	defer chk.Recover()
	flag.Parse()
	if *name == "" {
		for _, f := range examples {
			f()
		}
		return
	}
	if f, ok := examples[*name]; ok {
		f()
	}
}
