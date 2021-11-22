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
	"testing"

	"github.com/wildwest-productions/goini/internal/pkg/test"
)

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

func TestGetIniTags(t *testing.T) {

	gotTags := GetIniTags(testIni)

	expectedTags := make(map[string]string)

	expectedTags["#hello"] = "world"
	expectedTags["#bar"] = "2"
	expectedTags["#t"] = "hi"
	expectedTags["#timer"] = "23.31"
	expectedTags["#sys_enable"] = "true"
	expectedTags["command#command"] = "echo hello"
	expectedTags["command#runs"] = "42"
	expectedTags["sector#hello"] = "world-world"
	expectedTags["sector#sys_enable"] = "yes"

	test.AssertEqual(t, expectedTags, gotTags, "did not get expected tags")
}
