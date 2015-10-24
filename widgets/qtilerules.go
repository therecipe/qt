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

type QTileRulesITF interface {
	QTileRulesPTR() *QTileRules
}

func (p *QTileRules) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTileRules) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTileRules(ptr QTileRulesITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTileRulesPTR().Pointer()
	}
	return nil
}

func QTileRulesFromPointer(ptr unsafe.Pointer) *QTileRules {
	var n = new(QTileRules)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTileRules) QTileRulesPTR() *QTileRules {
	return ptr
}

func NewQTileRules(horizontalRule core.Qt__TileRule, verticalRule core.Qt__TileRule) *QTileRules {
	return QTileRulesFromPointer(unsafe.Pointer(C.QTileRules_NewQTileRules(C.int(horizontalRule), C.int(verticalRule))))
}

func NewQTileRules2(rule core.Qt__TileRule) *QTileRules {
	return QTileRulesFromPointer(unsafe.Pointer(C.QTileRules_NewQTileRules2(C.int(rule))))
}
