package templater

import (
	"bytes"
	"fmt"
	"strings"

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
%vextern "C" {
#endif

`,
		buildTags(module),

		strings.ToUpper(module),
		strings.ToUpper(module),

		func() string {
			if module != parser.MOC {
				return ""
			}
			var bb = new(bytes.Buffer)
			defer bb.Reset()
			for _, c := range parser.CurrentState.ClassMap {
				if c.Module == parser.MOC {
					fmt.Fprintf(bb, "class %v;\n", c.Name)
				}
			}
			return bb.String()
		}(),
	)

	fmt.Fprintf(bb, "struct %v_PackedString { char* data; long long len; };\n", strings.Title(module))
	fmt.Fprintf(bb, "struct %v_PackedList { void* data; long long len; };\n", strings.Title(module))

	//

	for _, class := range getSortedClassesForModule(module) {

		cTemplate(bb, class, cppEnumHeader, cppFunctionHeader, ";\n")

	}

	//

	fmt.Fprint(bb, `
#ifdef __cplusplus
}
#endif

#endif`)

	return bb.Bytes()
}
