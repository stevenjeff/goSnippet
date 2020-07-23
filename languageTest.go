// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that it's OK to have C code that does nothing other than
// initialize a global variable.  This used to fail with gccgo.

package main

import (
	"bufio"
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"net"
	"os"
	"reflect"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

func main() {
	//testEnum2()
	//aliasTest()
	//structTest()
	//arrayTest()
	//sliceTest()
	//mapTest()
	//syncMap()
	//listTest()
	//forTest()
	switchTest()
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

func arrayTest() {
	var team [3]string
	team[0] = "a"
	team[1] = "b"
	team[2] = "c"
	fmt.Println(team)
	var teamA = [3]string{"e", "f", "g"}
	fmt.Println(teamA)
	var teamB = [...]string{"m", "n", "o"}
	fmt.Println(teamB)
	for k, v := range team {
		fmt.Println(k, v)
	}
}

func sliceTest() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:3])
	m := make([]int, 2)
	n := make([]int, 2, 10)
	fmt.Println(m, n)
	fmt.Println(len(m), len(n))
	m = append(m, 4)
	fmt.Println(m)
	k := make([]int, 3)
	copy(k, m)
	fmt.Println(k)
}

func mapTest() {
	scene := make(map[string]int)
	scene["route"] = 66
	scene["china"] = 223
	scene["brazil"] = 3323
	fmt.Println(scene["route"])
	m := map[string]string{
		"aaa": "bbb",
		"cc":  "ddd",
	}
	fmt.Println(m)
	for k, v := range scene {
		fmt.Println(k, v)
	}
	var sceneList []string
	for k := range scene {
		sceneList = append(sceneList, k)
	}
	sort.Strings(sceneList)
	fmt.Println(sceneList)
	delete(scene, "route")
	fmt.Println(scene)
}

func syncMap() {
	var scene sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)
	fmt.Println(scene.Load("greece"))
	scene.Delete("london")
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})
}

func listTest() {
	l := list.New()
	l.PushBack("first")
	l.PushFront(333)
	fmt.Println(l)
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func forTest() {
	var ten int = 11
	if ten > 10 {
		fmt.Println(ten)
	} else if ten < 10 {
		fmt.Println(ten)
	} else {
		fmt.Println(ten)
	}

	if t := 4; t > 1 {
		fmt.Println(t)
	}
	for key, value := range []int{1, 2, 3, 4, 5, 6} {
		fmt.Println(key, value)
	}
	c := make(chan int)
	go func() {
		c <- 5
		c <- 2
		c <- 3
		close(c)
	}()
	for v := range c {
		fmt.Println(v)
	}
	m := map[string]int{
		"hello": 100,
		"world": 200,
	}
	for _, value := range m {
		fmt.Println(value)
	}
}

func switchTest() {
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world", "tide":
		fmt.Println(2)
	default:
		fmt.Println(0)
	}

	switch {
	case a != "f" && a != "ttt":
		fmt.Println(2)
	}

	switch {
	case a == "hello":
		fmt.Println("hello")
		fallthrough
	case a != "world":
		fmt.Println("world")
	}
}
