package templater

import (
	"bytes"
	"fmt"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cTemplate(bb *bytes.Buffer, c *parser.Class, ef func(*parser.Enum, *parser.Value) string, ff func(*parser.Function) string, del string, isGo bool) {
	cTemplateEnums(bb, c, ef, del, isGo)

	if c.IsSupported() {
		cTemplateFunctions(bb, c, ff, del, isGo)
	}
}

func cTemplateEnums(bb *bytes.Buffer, c *parser.Class, ef func(*parser.Enum, *parser.Value) string, del string, isGo bool) {
	for _, enum := range c.Enums {
		if isGo {
			fmt.Fprintf(bb, "%v%v", ef(enum, nil), del)
		} else {
			for _, value := range enum.Values {
				if converter.EnumNeedsCppGlue(value.Value) {
					fmt.Fprintf(bb, "%v%v", ef(enum, value), del)
				}
			}
		}
	}
}

func cTemplateFunctions(bb *bytes.Buffer, pc *parser.Class, ff func(*parser.Function) string, del string, isGo bool) {

	var implemented = make(map[string]struct{})

	for i, cn := range append([]string{pc.Name}, pc.GetAllBases()...) {
		var c, e = parser.State.ClassMap[cn]
		if !e || !c.IsSupported() {
			continue
		}

		for _, f := range c.Functions {
			var _, e = implemented[fmt.Sprint(f.Name, f.OverloadNumber)]
			if e || !f.IsSupported() {
				continue
			}

			if i > 0 && (f.Meta == parser.CONSTRUCTOR || f.Meta == parser.DESTRUCTOR || (f.Meta == parser.SIGNAL && !parser.State.Moc)) {
				continue
			}

			if i > 0 && !(f.Meta == parser.SLOT || f.Virtual == parser.PURE || f.Virtual == parser.IMPURE) {
				continue
			}

			if !f.Checked {
				cppFunction(f)
				goFunction(f)
				f.Checked = true
			}

			if !f.IsSupported() {
				continue
			}

			var f = *f

			if i > 0 {
				f.Fullname = fmt.Sprintf("%v::%v", pc.Name, f.Name)
				f.Synthetic = true
			}

			implemented[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}
			switch {
			case f.Meta == parser.SLOT:
				{
					if isGo {
						for _, signalMode := range []string{parser.CALLBACK, parser.CONNECT, parser.DISCONNECT} {
							var f = f
							f.SignalMode = signalMode
							if out := ff(&f); out != "" {
								if signalMode == parser.CALLBACK {
									fmt.Fprintf(bb, "//export %v\n", converter.GoHeaderName(&f))
								}
								fmt.Fprintf(bb, "%v%v", out, del)
							}
						}
					}

					if isGo {
						f.Meta = parser.PLAIN
					}
					fmt.Fprintf(bb, "%v%v", ff(&f), del)

					if f.Virtual != parser.PURE || (f.Synthetic && f.IsDerivedFromImpure()) {
						f.Meta = parser.PLAIN
						f.Default = true
						fmt.Fprintf(bb, "%v%v", ff(&f), del)
					}
				}

			case f.Meta == parser.SIGNAL:
				{
					for _, signalMode := range []string{parser.CALLBACK, parser.CONNECT, parser.DISCONNECT} {
						if !isGo && signalMode == parser.CALLBACK {
							continue
						}
						var f = f
						f.SignalMode = signalMode
						if out := ff(&f); out != "" {
							if signalMode == parser.CALLBACK && isGo {
								fmt.Fprintf(bb, "//export %v\n", converter.GoHeaderName(&f))
							}
							fmt.Fprintf(bb, "%v%v", out, del)
						}
					}

					if !converter.IsPrivateSignal(&f) {
						f.Meta = parser.PLAIN
						fmt.Fprintf(bb, "%v%v", ff(&f), del)
					}
				}

			case f.Virtual == parser.IMPURE, f.Virtual == parser.PURE:
				{
					if isGo {
						for _, signalMode := range []string{parser.CALLBACK, parser.CONNECT, parser.DISCONNECT} {
							var f = f
							f.SignalMode = signalMode
							if out := ff(&f); out != "" {
								if signalMode == parser.CALLBACK {
									fmt.Fprintf(bb, "//export %v\n", converter.GoHeaderName(&f))
								}
								fmt.Fprintf(bb, "%v%v", out, del)
							}
						}
					}

					f.Meta = parser.PLAIN
					fmt.Fprintf(bb, "%v%v", ff(&f), del)

					if f.Virtual != parser.PURE || (f.Synthetic && f.IsDerivedFromImpure()) {
						f.Default = true
						fmt.Fprintf(bb, "%v%v", ff(&f), del)
					}
				}

			case f.IsJNIGeneric():
				{
					for _, m := range converter.CppOutputParametersJNIGenericModes(&f) {
						f.TemplateModeJNI = m
						fmt.Fprintf(bb, "%v%v", ff(&f), del)

						f.Exception = true
						fmt.Fprintf(bb, "%v%v", ff(&f), del)
						f.Exception = false
					}
				}

			default:
				{
					if !(f.Meta == parser.CONSTRUCTOR && hasUnimplementedPureVirtualFunctions(c.Name)) {
						var out = ff(&f)
						if out != "" {
							fmt.Fprintf(bb, "%v%v", out, del)
							if isGo && f.Static {
								fmt.Fprintf(bb, "%v{\n%v\n}\n\n",
									func() string {
										var f = f
										f.Static = false
										return goFunctionHeader(&f)
									}(),
									goFunctionBody(&f),
								)
							}
						} else {
							f.Access = "unsupported_"
						}
					}
				}
			}
		}
	}
}
