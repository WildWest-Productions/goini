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

package test

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func AssertNil(t *testing.T, n interface{}, m ...string) {
	if n != nil {
		msg := fmt.Sprintf("\nExpected nil; got: %+v\n      Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("          Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func AssertNotNil(t *testing.T, n interface{}, m ...string) {
	if n == nil {
		msg := fmt.Sprintf("\nExpected non-nil; got: %+v\n      Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("              Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func AssertTrue(t *testing.T, n bool, m ...string) {
	if n != true {
		msg := fmt.Sprintf("\nExpected true; got: %v\n       Test failed: %s\n", n, t.Name())
		if len(m) > 0 {
			msg += fmt.Sprintf("           Message: %s\n", m[0])
		}
		t.Fatalf("%s\n", msg)
	}
}

func AssertEqual(t *testing.T, e, g interface{}, m ...string) {
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

func AssertEqualBytes(t *testing.T, e, g []byte, m ...string) {
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
