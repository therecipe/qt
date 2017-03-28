package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/parser"
	"github.com/therecipe/qt/internal/utils"
)

func HTemplate(m string, mode int) []byte {
	utils.Log.WithField("module", m).Debug("generating h")

	var bb = new(bytes.Buffer)
	defer bb.Reset()

	if m != parser.MOC {
		m = "Qt" + m
	}

	//header
	fmt.Fprintf(bb, "%v\n\n", buildTags(m, false, mode))

	fmt.Fprint(bb, "#pragma once\n\n")

	fmt.Fprintf(bb, "#ifndef GO_%v_H\n", strings.ToUpper(m))
	fmt.Fprintf(bb, "#define GO_%v_H\n\n", strings.ToUpper(m))

	fmt.Fprint(bb, "#include <stdint.h>\n\n")

	fmt.Fprint(bb, "#ifdef __cplusplus\n")
	if m == parser.MOC {
		for _, c := range parser.SortedClassNamesForModule(m, true) {
			fmt.Fprintf(bb, "class %v;\n", c)
		}
	}
	fmt.Fprint(bb, "extern \"C\" {\n#endif\n\n")

	fmt.Fprintf(bb, "struct %v_PackedString { char* data; long long len; };\n", strings.Title(m))
	fmt.Fprintf(bb, "struct %v_PackedList { void* data; long long len; };\n", strings.Title(m))

	//body
	for _, c := range parser.SortedClassesForModule(m, true) {
		cTemplate(bb, c, cppEnumHeader, cppFunctionHeader, ";\n", false)
	}

	//footer
	fmt.Fprint(bb, "\n#ifdef __cplusplus\n}\n#endif\n\n#endif")

	return bb.Bytes()
}
