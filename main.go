//
// Created by WestleyR <westleyr@nym.hush.com> on Nov 20, 2021
// Source code: https://github.com/WestleyR/goini
//              https://github.com/WildWest-Productions/goini
//
// Copyright (c) 2021 WestleyR. All rights reserved.
// This software is licensed under a BSD 3-Clause Clear License.
// Consult the LICENSE file that came with this software regarding
// your rights to distribute this software.
//

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Point struct {
	ConfigStr string `ini:"hello"`
	ConfigInt int    `ini:"bar"`
	Val       string `ini:"t"`
	Hello     Foo    `ini:"sector"`
}

type Foo struct {
	Hello string `ini:"hello"`
}

func getStructTags(s interface{}) (map[string]string, []string) {
	//	var tags []string
	tags := make(map[string]string)
	tagArr := []string{}

	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		//		fmt.Printf("INFO: %s -> %s\n", t.Field(i).Name, t.Field(i).Tag.Get("ini"))
		tags[t.Field(i).Tag.Get("ini")] = t.Field(i).Type.String()
		tagArr = append(tagArr, t.Field(i).Tag.Get("ini"))
		//		fmt.Printf("FOOO: %T\n", t.Field(i).Type.String())
	}

	return tags, tagArr
}

func getIniTags(iniData []byte) map[string]string {
	tags := make(map[string]string)

	scanner := bufio.NewScanner(bytes.NewReader(iniData))
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), "=")

		// Remove any whitespaces
		for i := 0; i < len(values); i++ {
			values[i] = strings.ReplaceAll(values[i], " ", "")
		}

		if len(values) == 2 {
			tags[values[0]] = values[1]
		}
	}

	return tags
}

func Unmarshal(iniData []byte, s interface{}) error {
	t := reflect.ValueOf(s)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	tags, tagArr := getStructTags(s)
	fmt.Printf("TAGS: %v\n", tags)
	fmt.Printf("Tags: %s\n", tagArr)

	iniTags := getIniTags(iniData)
	fmt.Printf("TAGS: %v\n", iniTags)

	for i := 0; i < t.NumField(); i++ {
		iniValue, _ := iniTags[tagArr[i]]
		iniValueType, _ := tags[tagArr[i]]

		fmt.Println("")
		fmt.Printf("tag=%v type=%s value=%s\n", tagArr[i], iniValueType, iniValue)

		switch iniValueType {
		case "string":
			t.Field(i).SetString(iniValue)
		case "int":
			val, err := strconv.Atoi(iniValue)
			if err == nil {
				t.Field(i).SetInt(int64(val))
			}
		}
	}

	return nil
}

func main() {
	var reply = &Point{}

	iniData := []byte(
		`
hello = world
bar = 2
t = hi

[sector]
hello = world

`)

	Unmarshal(iniData, reply)

	//	for i := 0; i < t.NumField(); i++ {
	//		fmt.Printf("%+v\n", t.Field(i))
	//	}

	fmt.Printf("\nEND_STRUCT: %+v\n", reply)
}
