package interop

import (
	"unsafe"

	"github.com/bluszcz/cutego"

	"github.com/bluszcz/cutego/core"
)

//TODO: this == nil checks

var globalEngine *PseudoQJSEngine

type PseudoQJSEngine struct {
	QObject       core.QObject //TODO: make it possible to subclass without invoking qtmoc
	_globalObject *PseudoQJSValue
}

type PseudoQJSEngine_ITF interface {
	core.QObject_ITF
	PseudoQJSEngine_PTR() *PseudoQJSEngine
}

//TODO: make it possible to subclass without invoking qtmoc
func (this *PseudoQJSEngine) Pointer() unsafe.Pointer {
	return this.QObject.Pointer()
}

func (this *PseudoQJSEngine) SetPointer(ptr unsafe.Pointer) {
	this.QObject.SetPointer(ptr)
}

func (this *PseudoQJSEngine) QObject_PTR() *core.QObject {
	return this.QObject.QObject_PTR()
}

//<<<

//public >>>

func PseudoQJSEngine_qjsEngine(i *core.QObject) *PseudoQJSEngine {
	if globalEngine != nil {
		return globalEngine
	}

	globalEngine = &PseudoQJSEngine{*core.NewQObject(nil), NewPseudoQJSValue1(make(map[string]interface{}))}
	qt.Register(globalEngine.Pointer(), globalEngine)
	return globalEngine
}

func (this *PseudoQJSEngine) NewObject() *PseudoQJSValue {
	return NewPseudoQJSValue1(make(map[string]interface{}))
}

func (this *PseudoQJSEngine) NewArray(l uint) *PseudoQJSValue {
	return NewPseudoQJSValue1(make([]interface{}, l))
}

func (this *PseudoQJSEngine) GlobalObject() *PseudoQJSValue {
	o, _ := qt.Receive(this.Pointer())
	return o.(*PseudoQJSEngine)._globalObject
}

func (this *PseudoQJSEngine) PseudoQJSEngine_PTR() *PseudoQJSEngine {
	return this
}

//<<< public

func (this *PseudoQJSEngine) ToScriptValue(i *core.QVariant) *PseudoQJSValue {
	return NewPseudoQJSValue2(i)
}

func (this *PseudoQJSEngine) NewQObject(i *core.QObject) *PseudoQJSValue {
	o, _ := qt.Receive(this.Pointer())
	i.SetParent(o.(*PseudoQJSEngine))
	return NewPseudoQJSValue1(i)
}

func (this *PseudoQJSEngine) Evaluate(string, string, int) *PseudoQJSValue {
	return NewPseudoQJSValue(1) //TODO: can be used to dynamically generate bindings on the remote side
}

func (this *PseudoQJSEngine) DestroyPseudoQJSEngine() {
	this.QObject.DestroyQObject()
}
