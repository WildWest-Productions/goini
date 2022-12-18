//
// Created by WestleyR <westleyr@nym.hush.com> on Nov 22, 2021
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
	"bytes"
	"encoding"
	"fmt"
	"reflect"
)

func marshaler(buf *bytes.Buffer, value reflect.Value, sector string) error {
	structFields := getStructTags(value.Interface())

	// Need to loop twice since we want to get all "default" fields first, then
	// sub-structs.

	for i := 0; i < value.NumField(); i++ {
		if impWritter(value.Field(i)) && sector == "" {
			writeExp(buf, structFields.fieldTags[i], value.Field(i))
		}
	}

	for i := 0; i < value.NumField(); i++ {
		if impWritter(value.Field(i)) {
			if sector != "" {
				writeExp(buf, structFields.fieldTags[i], value.Field(i))
			}
		} else {
			sectorT := value.FieldByName(structFields.fieldNames[i])
			if !sectorT.IsValid() {
				return fmt.Errorf("internal error: field not found: %s", structFields.fieldNames[i])
			}

			writeSection(buf, structFields.fieldTags[i])

			err := marshaler(buf, sectorT, structFields.fieldTags[i])
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func impWritter(value reflect.Value) bool {
	if value.Kind() == reflect.Struct {
		v := value.Interface()
		_, ok := v.(encoding.TextMarshaler)
		return ok
	}

	if value.Kind() != reflect.Struct {
		return true
	}

	return false
}

//func writeExp(buf *bytes.Buffer, key string, val interface{}) {
func writeExp(buf *bytes.Buffer, key string, val reflect.Value) {
	v := val.Interface()

	u, ok := v.(encoding.TextMarshaler)
	if ok {
		b, err := u.MarshalText()
		if err != nil {
			panic(fmt.Sprintf("error using text marshaler: %s", err))
		}
		buf.WriteString(fmt.Sprintf("%s = %s", key, string(b)))
	} else {
		buf.WriteString(fmt.Sprintf("%s = %v", key, val.Interface()))
	}

	buf.WriteString("\n")
}

func writeSection(buf *bytes.Buffer, sect string) {
	// TODO: cat exps
	buf.WriteString("\n")
	buf.WriteString(fmt.Sprintf("[%s]", sect))
	buf.WriteString("\n")
}

func MarshalValue(value reflect.Value) ([]byte, error) {
	b := bytes.NewBuffer(nil)

	err := marshaler(b, value, "")

	return b.Bytes(), err
}
