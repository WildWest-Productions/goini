//
// Created by WestleyR <westleyr@nym.hush.com> on Dec 10, 2021
// Source code: https://github.com/WestleyR/goini
//              https://github.com/WildWest-Productions/goini
//
// Copyright (c) 2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package goini

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

type testStruct struct {
	ConfigStr     string     `ini:"hello"`
	ConfigInt     int        `ini:"bar"`
	Timmer        float64    `ini:"timer"`
	Val           string     `ini:"t"`
	Hello         subStruct  `ini:"sector"`
	SystemEnabled bool       `ini:"sys_enable"`
	Command       subStruct2 `ini:"command"`
}

type pointerStruct struct {
	Hello string     `ini:"hello"`
	Ptr   *subStruct `ini:"ptr"`
}

type subStruct struct {
	Bar           string `ini:"hello"`
	SystemEnabled bool   `ini:"sys_enable"`
}

type subStruct2 struct {
	Command string `ini:"command"`
	Runs    int    `ini:"runs"`
}

var testIni = []byte(`
hello=world
bar = 2
t=hi
sys_enable = true
# test comment 1
timer = 23.31

[command]
; another test comment
command = echo hello
runs = 42

[sector]
hello=world-world
sys_enable = yes
`)

var pointerIni = []byte(`
hello = hello world

[ptr]
hello = bar
sys_enable = y
`)

func assertNil(t *testing.T, n interface{}, m ...string) {
	if n != nil {
		msg := fmt.Sprintf("\nExpected nil; got: %+v\n      Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("          Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func assertNotNil(t *testing.T, n interface{}, m ...string) {
	if n == nil {
		msg := fmt.Sprintf("\nExpected non-nil; got: %+v\n      Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("              Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func assertTrue(t *testing.T, n bool, m ...string) {
	if n != true {
		msg := fmt.Sprintf("\nExpected true; got: %v\n       Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("           Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func assertEqual(t *testing.T, e, g interface{}, m ...string) {
	if !reflect.DeepEqual(e, g) {
		msg := fmt.Sprintf("\nExpected to be equal:\n")
		msg += fmt.Sprintf("  Expected: %+v\n", e)
		msg += fmt.Sprintf("  Got     : %+v\n", g)
		msg += fmt.Sprintf("In test   : %s\n", t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("Message   : %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func assertEqualBytes(t *testing.T, e, g []byte, m ...string) {
	if !bytes.Equal(e, g) {
		msg := fmt.Sprintf("\nExpected to be equal:\n")
		msg += fmt.Sprintf("  Expected: %q\n", string(e))
		msg += fmt.Sprintf("  Got     : %q\n", string(g))
		msg += fmt.Sprintf("In test   : %s\n", t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("Message   : %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func TestUnmarshalPointer(t *testing.T) {
	s := &testStruct{}

	err := Unmarshal(testIni, &s)
	assertNil(t, err, "failed to unmarshal")

	expected := &testStruct{
		ConfigStr: "world",
		ConfigInt: 2,
		Timmer:    23.31,
		Val:       "hi",
		Hello: subStruct{
			Bar:           "world-world",
			SystemEnabled: true,
		},
		SystemEnabled: true,
		Command: subStruct2{
			Command: "echo hello",
			Runs:    42,
		},
	}

	assertEqual(t, expected, s, "did not get expected struct")
}

func TestUnmarshalSubPointer(t *testing.T) {
	t.Skip("skipping for now... (does not work yet...)")
	s := &pointerStruct{}

	err := Unmarshal(pointerIni, &s)
	assertNil(t, err, "failed to unmarshal")

	expected := &pointerStruct{
		Hello: "hello world11",
		Ptr: &subStruct{
			Bar:           "bar",
			SystemEnabled: true,
		},
	}

	assertEqual(t, expected, s, "did not get expected struct")

	t.FailNow()
}

func TestUnmarshalNonPointer(t *testing.T) {
	s := testStruct{}

	err := Unmarshal(testIni, &s)
	assertNil(t, err, "failed to unmarshal")

	expected := testStruct{
		ConfigStr: "world",
		ConfigInt: 2,
		Timmer:    23.31,
		Val:       "hi",
		Hello: subStruct{
			Bar:           "world-world",
			SystemEnabled: true,
		},
		SystemEnabled: true,
		Command: subStruct2{
			Command: "echo hello",
			Runs:    42,
		},
	}

	assertEqual(t, expected, s, "did not get expected struct")
}

func TestUnmarshalNonStruct(t *testing.T) {
	{
		s := "string"
		err := Unmarshal([]byte("foo = bar"), &s)
		assertNotNil(t, err, "expected error")
	}
	{
		s := 42
		err := Unmarshal([]byte("foo = bar"), &s)
		assertNotNil(t, err, "expected error")
	}
	{
		s := true
		err := Unmarshal([]byte("foo = bar"), &s)
		assertNotNil(t, err, "expected error")
	}
	{
		s := make(map[string]string)
		err := Unmarshal([]byte("foo = bar"), &s)
		assertNotNil(t, err, "expected error")
	}
}

func TestMarshalPointer(t *testing.T) {
	testStruct := &testStruct{
		ConfigStr: "world",
		ConfigInt: 2,
		Timmer:    23.31,
		Val:       "hi",
		Hello: subStruct{
			Bar:           "world-world",
			SystemEnabled: true,
		},
		SystemEnabled: true,
		Command: subStruct2{
			Command: "echo hello",
			Runs:    42,
		},
	}

	b, err := Marshal(testStruct)
	assertNil(t, err, "failed to marshal")

	expectedIni := []byte(`hello = world
bar = 2
timer = 23.31
t = hi
sys_enable = true

[sector]
hello = world-world
sys_enable = true

[command]
command = echo hello
runs = 42
`)
	assertEqualBytes(t, expectedIni, b, "did not get expected output")
}

func TestMarshalNonPointer(t *testing.T) {
	testStruct := testStruct{
		ConfigStr: "world",
		ConfigInt: 2,
		Timmer:    23.31,
		Val:       "hi",
		Hello: subStruct{
			Bar:           "world-world",
			SystemEnabled: true,
		},
		SystemEnabled: true,
		Command: subStruct2{
			Command: "echo hello",
			Runs:    42,
		},
	}

	b, err := Marshal(testStruct)
	assertNil(t, err, "failed to marshal")

	expectedIni := []byte(`hello = world
bar = 2
timer = 23.31
t = hi
sys_enable = true

[sector]
hello = world-world
sys_enable = true

[command]
command = echo hello
runs = 42
`)
	assertEqualBytes(t, expectedIni, b, "did not get expected output")
}

func TestMarshalNonStruct(t *testing.T) {
	{
		s := "string"
		b, err := Marshal(s)
		assertNotNil(t, err, "expected error")
		assertEqualBytes(t, []byte{}, b)
	}
	{
		s := "string"
		b, err := Marshal(&s)
		assertNotNil(t, err, "expected error")
		assertEqualBytes(t, []byte{}, b)
	}
	{
		s := 42
		b, err := Marshal(s)
		assertNotNil(t, err, "expected error")
		assertEqualBytes(t, []byte{}, b)
	}
	{
		s := true
		b, err := Marshal(s)
		assertNotNil(t, err, "expected error")
		assertEqualBytes(t, []byte{}, b)
	}
	{
		s := make(map[string]string)
		b, err := Marshal(s)
		assertNotNil(t, err, "expected error")
		assertEqualBytes(t, []byte{}, b)
	}
}
