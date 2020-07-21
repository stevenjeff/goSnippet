// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that it's OK to have C code that does nothing other than
// initialize a global variable.  This used to fail with gccgo.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"strings"
	"unicode/utf8"
)

func main() {
	//testEnum2()
	//aliasTest()
	structTest()
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

func flagTest() {
	var model = flag.String("model", "", "process model")
	flag.Parse()
	fmt.Printf(*model)
}

func newPoint() {
	str := new(string)
	*str = "ninja的打法"
	fmt.Println(str)
	fmt.Println(len(*str))
	fmt.Println(utf8.RuneCountInString(*str))
	fmt.Println(*str)
}

func changeStr() {
	angel := "fdfsdfsdfffgfdgdfgdfgdfgdfgfdfg"
	angleBytes := []byte(angel)
	for i := 5; i <= 13; i++ {
		angleBytes[i] = ' '
	}
	fmt.Println(string(angleBytes))
}

func stringBuilder() {
	var stringBuilder bytes.Buffer
	stringBuilder.WriteString("3333")
	stringBuilder.WriteString("2fsdfsd")
	fmt.Println(stringBuilder.String())
}

func fileTest() {
	file, err := os.Open("C:\\Windows\\INF\\.NET CLR Data\\_DataPerfCounters.ini")
	if err != nil {
		return
	}
	reader := bufio.NewReader(file)
	for {
		linestr, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		linestr = strings.TrimSpace(linestr)
		if linestr == "" {
			continue
		}
		if linestr[0] == ';' {
			continue
		}
		if linestr[1] == '[' && linestr[len(linestr)-4] == ']' {
			sectionName := linestr[1 : len(linestr)-1]
			fmt.Println(sectionName)
		}
	}
	defer file.Close()
}

func enumTest() {
	type Weapon int
	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)
	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Println(FlagRed, FlagGreen, FlagBlue)
}

type chipType int

const (
	None chipType = iota
	CPU
	GPU
)

func (c chipType) String() string {
	switch c {
	case None:
		return "NONE"
	case CPU:
		return "CPU"
	case GPU:
		return "GPU"
	}
	return "N/A"
}

func testEnum2() {
	fmt.Println("%s %d", GPU, GPU)
	fmt.Println("%s %d", GPU, GPU)
}

func aliasTest() {
	type NewInt int
	type IntAlias = int
	var a1 NewInt
	var a2 IntAlias
	fmt.Printf("%T : %T\n", a1, a2)
}

type Brand struct {
}

func (t Brand) Show() {
	fmt.Print("brand show :")
	fmt.Println(t)
}

type FakeBrand = Brand
type Vehicle struct {
	FakeBrand
	Brand
}

func structTest() {
	var a Vehicle
	a.FakeBrand.Show()
	ta := reflect.TypeOf(a)
	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		fmt.Printf("FieldName: %v ,FieldType : %v\n", f.Name, f.Type.Name())
	}
}
