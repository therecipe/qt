package core

//#include "core.h"
import "C"
import (
	"log"
	"strings"
	"unsafe"
)

type QRegularExpressionMatch struct {
	ptr unsafe.Pointer
}

type QRegularExpressionMatch_ITF interface {
	QRegularExpressionMatch_PTR() *QRegularExpressionMatch
}

func (p *QRegularExpressionMatch) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegularExpressionMatch) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegularExpressionMatch(ptr QRegularExpressionMatch_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionMatch_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionMatchFromPointer(ptr unsafe.Pointer) *QRegularExpressionMatch {
	var n = new(QRegularExpressionMatch)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegularExpressionMatch) QRegularExpressionMatch_PTR() *QRegularExpressionMatch {
	return ptr
}

func NewQRegularExpressionMatch() *QRegularExpressionMatch {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::QRegularExpressionMatch")
		}
	}()

	return NewQRegularExpressionMatchFromPointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch())
}

func NewQRegularExpressionMatch2(match QRegularExpressionMatch_ITF) *QRegularExpressionMatch {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::QRegularExpressionMatch")
		}
	}()

	return NewQRegularExpressionMatchFromPointer(C.QRegularExpressionMatch_NewQRegularExpressionMatch2(PointerFromQRegularExpressionMatch(match)))
}

func (ptr *QRegularExpressionMatch) Captured2(name string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::captured")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpressionMatch_Captured2(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QRegularExpressionMatch) Captured(nth int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::captured")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpressionMatch_Captured(ptr.Pointer(), C.int(nth)))
	}
	return ""
}

func (ptr *QRegularExpressionMatch) CapturedEnd2(name string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedEnd2(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedEnd(nth int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedEnd")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedEnd(ptr.Pointer(), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedLength2(name string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedLength2(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedLength(nth int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedLength(ptr.Pointer(), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedRef2(name string) *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedRef")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QRegularExpressionMatch_CapturedRef2(ptr.Pointer(), C.CString(name)))
	}
	return nil
}

func (ptr *QRegularExpressionMatch) CapturedRef(nth int) *QStringRef {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedRef")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QRegularExpressionMatch_CapturedRef(ptr.Pointer(), C.int(nth)))
	}
	return nil
}

func (ptr *QRegularExpressionMatch) CapturedStart2(name string) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedStart")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedStart2(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedStart(nth int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedStart")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_CapturedStart(ptr.Pointer(), C.int(nth)))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) CapturedTexts() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::capturedTexts")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QRegularExpressionMatch_CapturedTexts(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QRegularExpressionMatch) HasMatch() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::hasMatch")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_HasMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) HasPartialMatch() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::hasPartialMatch")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_HasPartialMatch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegularExpressionMatch_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegularExpressionMatch) LastCapturedIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::lastCapturedIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegularExpressionMatch_LastCapturedIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) MatchOptions() QRegularExpression__MatchOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::matchOptions")
		}
	}()

	if ptr.Pointer() != nil {
		return QRegularExpression__MatchOption(C.QRegularExpressionMatch_MatchOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) MatchType() QRegularExpression__MatchType {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::matchType")
		}
	}()

	if ptr.Pointer() != nil {
		return QRegularExpression__MatchType(C.QRegularExpressionMatch_MatchType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpressionMatch) RegularExpression() *QRegularExpression {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::regularExpression")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQRegularExpressionFromPointer(C.QRegularExpressionMatch_RegularExpression(ptr.Pointer()))
	}
	return nil
}

func (ptr *QRegularExpressionMatch) Swap(other QRegularExpressionMatch_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegularExpressionMatch_Swap(ptr.Pointer(), PointerFromQRegularExpressionMatch(other))
	}
}

func (ptr *QRegularExpressionMatch) DestroyQRegularExpressionMatch() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegularExpressionMatch::~QRegularExpressionMatch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegularExpressionMatch_DestroyQRegularExpressionMatch(ptr.Pointer())
	}
}
