package qt

//#include "qvalidator.h"
import "C"

type qvalidator struct {
	qobject
}

type QValidator interface {
	QObject
	ConnectSignalChanged(f func())
	DisconnectSignalChanged()
	SignalChanged() func()
}

func (p *qvalidator) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qvalidator) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

//State
type State int

var (
	INVALID      = State(C.QValidator_Invalid())
	INTERMEDIATE = State(C.QValidator_Intermediate())
	ACCEPTABLE   = State(C.QValidator_Acceptable())
)

func (p *qvalidator) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QValidator_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qvalidator) ConnectSignalChanged(f func()) {
	C.QValidator_ConnectSignalChanged(p.Pointer())
	connectSignal(p.ObjectName(), "changed", f)
}

func (p *qvalidator) DisconnectSignalChanged() {
	C.QValidator_DisconnectSignalChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "changed")
}

func (p *qvalidator) SignalChanged() func() {
	return getSignal(p.ObjectName(), "changed")
}

//export callbackQValidator
func callbackQValidator(ptr C.QtObjectPtr, msg *C.char) {
	var qvalidator = new(qvalidator)
	qvalidator.SetPointer(ptr)
	getSignal(qvalidator.ObjectName(), C.GoString(msg))()
}
