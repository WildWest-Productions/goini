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
	"fmt"
	"reflect"
	"testing"

	"github.com/wildwest-productions/goini/internal/pkg/test"
)

type tstruct struct {
	S   string
	B   bool
	I   int
	F32 float32
	F64 float64
}

func Test_setValue(t *testing.T) {
	ss := &tstruct{}
	v := reflect.ValueOf(ss)
	v = v.Elem()

	type args struct {
		s     reflect.Value
		typeT string
		value string
	}
	tests := []struct {
		name        string
		args        args
		expectedOut bool
	}{
		{
			name: "string",
			args: args{
				s:     v.Field(0),
				typeT: "string",
				value: "my-string",
			},
			expectedOut: true,
		},
		{
			name: "bool",
			args: args{
				s:     v.Field(1),
				typeT: "bool",
				value: "true",
			},
			expectedOut: true,
		},
		{
			name: "bool_false",
			args: args{
				s:     v.Field(1),
				typeT: "bool",
				value: "false",
			},
			expectedOut: true,
		},
		{
			name: "int",
			args: args{
				s:     v.Field(2),
				typeT: "int",
				value: "42",
			},
			expectedOut: true,
		},
		{
			name: "float32",
			args: args{
				s:     v.Field(3),
				typeT: "float32",
				value: "32.21",
			},
			expectedOut: true,
		},
		{
			name: "float64",
			args: args{
				s:     v.Field(4),
				typeT: "float64",
				value: "84.32",
			},
			expectedOut: true,
		},
		{
			name: "wrong_type",
			args: args{
				s:     v.Field(0),
				typeT: "bool",
				value: "true",
			},
			expectedOut: false,
		},
		{
			name: "unknown",
			args: args{
				s:     v.Field(0),
				typeT: "mytype",
				value: "nope",
			},
			expectedOut: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !tt.expectedOut {
				// Only expect a panic if not expecting any output
				defer recoverPanic()
			}
			setValue(tt.args.s, tt.args.typeT, tt.args.value)
			if tt.expectedOut {
				test.AssertEqual(t, tt.args.value, fmt.Sprintf("%v", tt.args.s.Interface()), "did not set value correctly")
			}
		})
	}
}

func recoverPanic() {
	r := recover()
	if r != nil {
		fmt.Printf("expected recover from panic: %s\n", r)
	}
}
