package templater

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/therecipe/qt/internal/binding/converter"
	"github.com/therecipe/qt/internal/binding/parser"
)

func test(bb *bytes.Buffer, class *parser.Class, fu func(*parser.Function) string, del string) {
	var implementedVirtuals = make(map[string]struct{})

	//all class fs
	for _, f := range class.Functions {
		if !functionIsSupported(class, f) {
			continue
		}
		implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}

		switch {
		case f.Meta == parser.SIGNAL:
			{
				for _, signalMode := range []string{parser.CONNECT, parser.DISCONNECT} {
					var f = *f
					f.SignalMode = signalMode
					fmt.Fprintf(bb, "%v%v", fu(&f), del)
				}

				var f = *f
				f.Meta = parser.PLAIN
				if converter.IsPrivateSignal(&f) {
					continue
				}
				fmt.Fprintf(bb, "%v%v", fu(&f), del)
			}

		case (f.Virtual == parser.IMPURE || f.Virtual == parser.PURE) && !strings.Contains(f.Meta, "constructor"):
			{
				var f = *f
				if f.Meta != parser.SLOT {
					f.Meta = parser.PLAIN
				}

				fmt.Fprintf(bb, "%v%v", fu(&f), del)
				if f.Virtual == parser.IMPURE {
					f.Meta = parser.PLAIN
					f.Default = true
					fmt.Fprintf(bb, "%v%v", fu(&f), del)
				}
			}

		case isGeneric(f):
			{
				for _, mode := range converter.CppOutputParametersJNIGenericModes(f) {
					var f = *f
					f.TemplateModeJNI = mode

					fmt.Fprintf(bb, "%v%v", fu(&f), del)
				}
			}

		default:
			{
				if !(f.Meta == parser.CONSTRUCTOR && hasUnimplementedPureVirtualFunctions(class.Name)) {
					fmt.Fprintf(bb, "%v%v", fu(f), del)
				}
			}
		}

	}

	//virtual parent functions
	for _, parentClassName := range class.GetAllBases() {
		var parentClass = parser.CurrentState.ClassMap[parentClassName]
		if classIsSupported(parentClass) {

			for _, f := range parentClass.Functions {
				if _, exists := implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)]; !exists {
					implementedVirtuals[fmt.Sprint(f.Name, f.OverloadNumber)] = struct{}{}

					if functionIsSupported(parentClass, f) {
						if f.Meta != parser.SIGNAL && (f.Virtual == parser.IMPURE || f.Virtual == parser.PURE || f.Meta == parser.SLOT) && !strings.Contains(f.Meta, "structor") {

							var f = *f
							f.Fullname = fmt.Sprintf("%v::%v", class.Name, f.Name)
							if f.Meta != parser.SLOT {
								f.Meta = parser.PLAIN
							}
							fmt.Fprintf(bb, "%v%v", fu(&f), del)

							var class, exist = f.Class()
							if !exist || (class.Module == parser.MOC && f.Virtual == parser.PURE) {
								continue
							}

							f.Meta = parser.PLAIN
							f.Default = true
							fmt.Fprintf(bb, "%v%v", fu(&f), del)

						}
					}

				}
			}

		}
	}
}
