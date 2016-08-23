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
			case Minimal:
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
					return "// +build sailfish"
				}

			default:
				{
					return "// +build !minimal"
				}
			}
		}(),

		strings.ToUpper(module),
		strings.ToUpper(module))

	for _, class := range getSortedClassesForModule(module) {
		var implementedVirtuals = make(map[string]bool)

		//all class enums
		for _, enum := range class.Enums {
			for _, value := range enum.Values {
				if converter.EnumNeedsCppGlue(value.Value) {
					fmt.Fprintf(bb, "%v;\n", cppEnumHeader(enum, value))
				}
			}
		}

		if classIsSupported(class) {

			//all class functions
			for _, function := range class.Functions {
				implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)] = true
				if functionIsSupported(class, function) {

					switch {
					case function.Meta == parser.SIGNAL:
						{
							for _, signalMode := range []string{parser.CONNECT, parser.DISCONNECT} {
								var function = *function
								function.SignalMode = signalMode

								fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))
							}

							var function = *function
							function.Meta = parser.PLAIN
							if !converter.IsPrivateSignal(&function) {
								fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))
							}
						}

					case (function.Virtual == parser.IMPURE || function.Virtual == parser.PURE) && !strings.Contains(function.Meta, "structor"):
						{
							var function = *function
							if function.Meta != parser.SLOT {
								function.Meta = parser.PLAIN
							}

							fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))
							if function.Virtual == parser.IMPURE {
								function.Meta = parser.PLAIN
								function.Default = true
								fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))
							}
						}

					case isGeneric(function):
						{
							for _, mode := range converter.CppOutputParametersJNIGenericModes(function) {
								var function = *function
								function.TemplateMode = mode

								fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))
							}
						}

					default:
						{
							if !(function.Meta == parser.CONSTRUCTOR && hasUnimplementedPureVirtualFunctions(class.Name)) {
								fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(function))
							}
						}
					}

				}
			}

			//virtual parent functions
			for _, parentClassName := range class.GetAllBases() {
				var parentClass = parser.ClassMap[parentClassName]
				if classIsSupported(parentClass) {

					for _, function := range parentClass.Functions {
						if _, exists := implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)]; !exists {
							implementedVirtuals[fmt.Sprint(function.Name, function.OverloadNumber)] = true

							if functionIsSupported(parentClass, function) {
								if function.Meta != parser.SIGNAL && (function.Virtual == parser.IMPURE || function.Virtual == parser.PURE || function.Meta == parser.SLOT) && !strings.Contains(function.Meta, "structor") {

									var function = *function
									function.Fullname = fmt.Sprintf("%v::%v", class.Name, function.Name)
									if function.Meta != parser.SLOT {
										function.Meta = parser.PLAIN
									}
									fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))

									function.Meta = parser.PLAIN
									function.Default = true
									fmt.Fprintf(bb, "%v;\n", cppFunctionHeader(&function))

								}
							}

						}
					}

				}
			}

		}

	}

	fmt.Fprint(bb, `
#ifdef __cplusplus
}
#endif

#endif`)

	return bb.Bytes()
}
