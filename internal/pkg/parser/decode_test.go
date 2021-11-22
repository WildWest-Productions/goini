//
// Created by WestleyR <westleyr@nym.hush.com> on 2021-12-11
// Source code: https://github.com/WestleyR/goini
//              https://github.com/WildWest-Productions/goini
//
// Copyright (c) 2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package parser

import (
	"reflect"
	"testing"

	"github.com/wildwest-productions/goini/internal/pkg/test"
)

func Test_getStructTags(t *testing.T) {
	type subS struct {
		Name string `ini:"name"`
	}

	type testStruct struct {
		S   string `ini:"s"`
		B   bool   `ini:"b"`
		Sub subS   `ini:"sub"`
	}

	s := &testStruct{}

	expectedValue := &values{}
	expectedValue.interfaceType = make(map[string]string)
	expectedValue.interfaceType["s"] = "string"
	expectedValue.interfaceType["b"] = "bool"
	expectedValue.interfaceType["sub"] = "parser.subS"
	expectedValue.fieldNames = []string{"S", "B", "Sub"}
	expectedValue.fieldTags = []string{"s", "b", "sub"}

	got := getStructTags(s)
	test.AssertEqual(t, expectedValue, got, "did not get expected struct data")
}

func TestUnmarshal(t *testing.T) {
	type subStruct struct {
		Bar           string `ini:"hello"`
		SystemEnabled bool   `ini:"sys_enable"`
	}
	type testStruct struct {
		Hello string    `ini:"hello"`
		S     subStruct `ini:"s"`
	}

	expected := &testStruct{
		Hello: "default",
		S: subStruct{
			Bar:           "world",
			SystemEnabled: true,
		},
	}

	iniData := []byte(`
hello = default

[s]
hello = world
sys_enable = true
`)

	s := &testStruct{}

	svalue := reflect.ValueOf(s)
	svalue = svalue.Elem()

	iniTags := GetIniTags(iniData)

	err := Unmarshal(s, svalue, iniTags, "")
	test.AssertNil(t, err, "error while unmarshaling ini")
	test.AssertEqual(t, expected, s, "did not get expected struct")
}
