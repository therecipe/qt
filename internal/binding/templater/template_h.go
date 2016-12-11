package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func HTemplate(module string) []byte {
	var bb = new(bytes.Buffer)
	defer bb.Reset()

	fmt.Fprintf(bb, `%v

#pragma once

#ifndef GO_%v_H
#define GO_%v_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

`,
		func() string {
			switch {
			case parser.CurrentState.Minimal:
				{
					return "// +build minimal"
				}

			case module == parser.MOC:
				{
					return ""
				}

			case module == "QtAndroidExtras":
				{
					return "// +build android"
				}

			case module == "QtSailfish":
				{
					return "// +build sailfish sailfish_emulator"
				}

			default:
				{
					return "// +build !minimal"
				}
			}
		}(),

		strings.ToUpper(module),
		strings.ToUpper(module))

	fmt.Fprintf(bb, "struct %v_PackedString { char* data; long long len; };\n", strings.Title(module))
	fmt.Fprintf(bb, "struct %v_PackedList { void* data; long long len; };\n", strings.Title(module))

	for _, class := range getSortedClassesForModule(module) {

		//all class enums
		for _, enum := range class.Enums {
			for _, value := range enum.Values {
				if converter.EnumNeedsCppGlue(value.Value) {
					fmt.Fprintf(bb, "%v;\n", cppEnumHeader(enum, value))
				}
			}
		}

		if classIsSupported(class) {

			test(bb, class, cppFunctionHeader, ";\n")

		}

	}

	fmt.Fprint(bb, `
#ifdef __cplusplus
}
#endif

#endif`)

	return bb.Bytes()
}
