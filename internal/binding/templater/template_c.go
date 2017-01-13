package templater

import (
	"bytes"
	"fmt"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func cTemplate(bb *bytes.Buffer, c *parser.Class, ef func(*parser.Enum, *parser.Value) string, ff func(*parser.Function) string, del string) {
	cTemplateEnums(bb, c, ef, del)

	if c.IsSupported() {
		cTemplateFunctions(bb, c, ff, del)
	}
}

func cTemplateEnums(bb *bytes.Buffer, c *parser.Class, ef func(*parser.Enum, *parser.Value) string, del string) {
	for _, enum := range c.Enums {
		for _, value := range enum.Values {
			if converter.EnumNeedsCppGlue(value.Value) {
				fmt.Fprintf(bb, "%v%v", ef(enum, value), del)
			}
		}
	}
}

func cTemplateFunctions(bb *bytes.Buffer, c *parser.Class, ff func(*parser.Function) string, del string) {

	var implementedVirtuals = make(map[string]struct{})

	for _, f := range c.Functions {
		if !f.IsSupported() {
			continue
		}

		implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}

		switch {
		case f.Meta == parser.SIGNAL:
			{
				for _, m := range []string{parser.CONNECT, parser.DISCONNECT} {
					var f = *f
					f.SignalMode = m
					fmt.Fprintf(bb, "%v%v", ff(&f), del)
				}

				if !converter.IsPrivateSignal(f) {
					var f = *f
					f.Meta = parser.PLAIN
					fmt.Fprintf(bb, "%v%v", ff(&f), del)
				}
			}

		case (f.Virtual == parser.IMPURE || f.Virtual == parser.PURE):
			{
				var f = *f
				if f.Meta != parser.SLOT {
					f.Meta = parser.PLAIN
				}
				fmt.Fprintf(bb, "%v%v", ff(&f), del)

				if f.Virtual == parser.IMPURE {
					f.Meta = parser.PLAIN
					f.Default = true
					fmt.Fprintf(bb, "%v%v", ff(&f), del)
				}
			}

		case f.IsJNIGeneric():
			{
				for _, m := range converter.CppOutputParametersJNIGenericModes(f) {
					var f = *f
					f.TemplateModeJNI = m
					fmt.Fprintf(bb, "%v%v", ff(&f), del)
				}
			}

		default:
			{
				if !(f.Meta == parser.CONSTRUCTOR && hasUnimplementedPureVirtualFunctions(c.Name)) {
					fmt.Fprintf(bb, "%v%v", ff(f), del)
				}
			}
		}
	}

	for _, bcn := range c.GetAllBases() {
		var bc, exist = parser.State.ClassMap[bcn]
		if !exist || !bc.IsSupported() {
			continue
		}

		for _, f := range bc.Functions {
			var _, exists = implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)]
			if exists || !f.IsSupported() {
				continue
			}

			//TODO: signals and slots may not be needed -> poly problems ?
			if f.Meta != parser.SIGNAL && (f.Virtual == parser.IMPURE || f.Virtual == parser.PURE || f.Meta == parser.SLOT) && f.Meta != parser.DESTRUCTOR {
				implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}

				var f = *f
				f.Fullname = fmt.Sprintf("%v::%v", c.Name, f.Name)
				if f.Meta != parser.SLOT {
					f.Meta = parser.PLAIN
				}
				fmt.Fprintf(bb, "%v%v", ff(&f), del)

				var c, _ = f.Class()
				if c.Module == parser.MOC && f.Virtual == parser.PURE {
					continue
				}

				f.Meta = parser.PLAIN
				f.Default = true
				fmt.Fprintf(bb, "%v%v", ff(&f), del)
			}
		}
	}
}
