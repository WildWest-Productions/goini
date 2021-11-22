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

func TestMarshalValue(t *testing.T) {
	type subStruct struct {
		Hello string `ini:"hello"`
	}
	type testStruct struct {
		S   string    `ini:"string"`
		B   bool      `ini:"bool"`
		Sub subStruct `ini:"sub"`
	}

	s := &testStruct{
		S: "hello-world",
		B: true,
		Sub: subStruct{
			Hello: "world",
		},
	}

	expected := []byte(`string = hello-world
bool = true

[sub]
hello = world
`)

	v := reflect.ValueOf(s)
	v = v.Elem()

	got, err := MarshalValue(v)
	test.AssertNil(t, err, "error while marshaling")
	test.AssertEqualBytes(t, expected, got, "did not get expected output")
}
