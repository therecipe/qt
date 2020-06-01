package interop

import "github.com/therecipe/qt"

func init() {
	qt.ItfMap["interop.PseudoQJSEngine_ITF"] = PseudoQJSEngine{}
	qt.FuncMap["interop.PseudoQJSEngine_qjsEngine"] = PseudoQJSEngine_qjsEngine
	qt.ItfMap["interop.PseudoQJSValue_ITF"] = PseudoQJSValue{}
	qt.FuncMap["interop.NewPseudoQJSValue"] = NewPseudoQJSValue
	qt.FuncMap["interop.NewPseudoQJSValue1"] = NewPseudoQJSValue1
	qt.FuncMap["interop.NewPseudoQJSValue2"] = NewPseudoQJSValue2
	qt.FuncMap["interop.NewPseudoQJSValue8"] = NewPseudoQJSValue8
}
