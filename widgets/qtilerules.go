package widgets

//#include "qtilerules.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTileRules struct {
	ptr unsafe.Pointer
}

type QTileRules_ITF interface {
	QTileRules_PTR() *QTileRules
}

func (p *QTileRules) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTileRules) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTileRules(ptr QTileRules_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTileRules_PTR().Pointer()
	}
	return nil
}

func NewQTileRulesFromPointer(ptr unsafe.Pointer) *QTileRules {
	var n = new(QTileRules)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTileRules) QTileRules_PTR() *QTileRules {
	return ptr
}

func NewQTileRules(horizontalRule core.Qt__TileRule, verticalRule core.Qt__TileRule) *QTileRules {
	return NewQTileRulesFromPointer(C.QTileRules_NewQTileRules(C.int(horizontalRule), C.int(verticalRule)))
}

func NewQTileRules2(rule core.Qt__TileRule) *QTileRules {
	return NewQTileRulesFromPointer(C.QTileRules_NewQTileRules2(C.int(rule)))
}
