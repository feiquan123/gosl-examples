// Copyright 2016 The Gosl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "flag"

var name = flag.String("name", "simple01", "Example name")
var runall = flag.Bool("all", true, "Run all examples?")

func main() {
	flag.Parse()
	if *runall {
		ang01()
		ang02()
		kmeans01()
		mclass01()
		simple01()
		return
	}
	if *name == "simple01" {
		simple01()
	}
}
