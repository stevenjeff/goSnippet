// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that it's OK to have C code that does nothing other than
// initialize a global variable.  This used to fail with gccgo.

package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("dfd")
	strPoint()
}

func network() {
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	fmt.Println(conn)
	fmt.Println(err)
}

func strPoint() {
	var house = "fdsfsdf"
	ptr := &house
	fmt.Printf("ptr  type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)
	value := *ptr
	fmt.Printf("value type: %T\n", value)
	fmt.Printf("value: %s\n", value)
}
