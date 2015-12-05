package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QRegExp struct {
	ptr unsafe.Pointer
}

type QRegExp_ITF interface {
	QRegExp_PTR() *QRegExp
}

func (p *QRegExp) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegExp) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegExp(ptr QRegExp_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegExp_PTR().Pointer()
	}
	return nil
}

func NewQRegExpFromPointer(ptr unsafe.Pointer) *QRegExp {
	var n = new(QRegExp)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegExp) QRegExp_PTR() *QRegExp {
	return ptr
}

//QRegExp::CaretMode
type QRegExp__CaretMode int64

const (
	QRegExp__CaretAtZero    = QRegExp__CaretMode(0)
	QRegExp__CaretAtOffset  = QRegExp__CaretMode(1)
	QRegExp__CaretWontMatch = QRegExp__CaretMode(2)
)

//QRegExp::PatternSyntax
type QRegExp__PatternSyntax int64

const (
	QRegExp__RegExp         = QRegExp__PatternSyntax(0)
	QRegExp__Wildcard       = QRegExp__PatternSyntax(1)
	QRegExp__FixedString    = QRegExp__PatternSyntax(2)
	QRegExp__RegExp2        = QRegExp__PatternSyntax(3)
	QRegExp__WildcardUnix   = QRegExp__PatternSyntax(4)
	QRegExp__W3CXmlSchema11 = QRegExp__PatternSyntax(5)
)

func NewQRegExp() *QRegExp {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::QRegExp")
		}
	}()

	return NewQRegExpFromPointer(C.QRegExp_NewQRegExp())
}

func NewQRegExp3(rx QRegExp_ITF) *QRegExp {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::QRegExp")
		}
	}()

	return NewQRegExpFromPointer(C.QRegExp_NewQRegExp3(PointerFromQRegExp(rx)))
}

func NewQRegExp2(pattern string, cs Qt__CaseSensitivity, syntax QRegExp__PatternSyntax) *QRegExp {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::QRegExp")
		}
	}()

	return NewQRegExpFromPointer(C.QRegExp_NewQRegExp2(C.CString(pattern), C.int(cs), C.int(syntax)))
}

func (ptr *QRegExp) Cap(nth int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::cap")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegExp_Cap(ptr.Pointer(), C.int(nth)))
	}
	return ""
}

func (ptr *QRegExp) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegExp_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRegExp) CaptureCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::captureCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegExp_CaptureCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegExp) CapturedTexts() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::capturedTexts")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QRegExp_CapturedTexts(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QRegExp) CaseSensitivity() Qt__CaseSensitivity {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::caseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__CaseSensitivity(C.QRegExp_CaseSensitivity(ptr.Pointer()))
	}
	return 0
}

func QRegExp_Escape(str string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::escape")
		}
	}()

	return C.GoString(C.QRegExp_QRegExp_Escape(C.CString(str)))
}

func (ptr *QRegExp) ExactMatch(str string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::exactMatch")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegExp_ExactMatch(ptr.Pointer(), C.CString(str)) != 0
	}
	return false
}

func (ptr *QRegExp) IndexIn(str string, offset int, caretMode QRegExp__CaretMode) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::indexIn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegExp_IndexIn(ptr.Pointer(), C.CString(str), C.int(offset), C.int(caretMode)))
	}
	return 0
}

func (ptr *QRegExp) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegExp_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegExp) IsMinimal() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::isMinimal")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegExp_IsMinimal(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegExp) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QRegExp_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegExp) LastIndexIn(str string, offset int, caretMode QRegExp__CaretMode) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::lastIndexIn")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegExp_LastIndexIn(ptr.Pointer(), C.CString(str), C.int(offset), C.int(caretMode)))
	}
	return 0
}

func (ptr *QRegExp) MatchedLength() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::matchedLength")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegExp_MatchedLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegExp) Pattern() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::pattern")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegExp_Pattern(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRegExp) PatternSyntax() QRegExp__PatternSyntax {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::patternSyntax")
		}
	}()

	if ptr.Pointer() != nil {
		return QRegExp__PatternSyntax(C.QRegExp_PatternSyntax(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegExp) Pos(nth int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::pos")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QRegExp_Pos(ptr.Pointer(), C.int(nth)))
	}
	return 0
}

func (ptr *QRegExp) SetCaseSensitivity(cs Qt__CaseSensitivity) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::setCaseSensitivity")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_SetCaseSensitivity(ptr.Pointer(), C.int(cs))
	}
}

func (ptr *QRegExp) SetMinimal(minimal bool) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::setMinimal")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_SetMinimal(ptr.Pointer(), C.int(qt.GoBoolToInt(minimal)))
	}
}

func (ptr *QRegExp) SetPattern(pattern string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::setPattern")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_SetPattern(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QRegExp) SetPatternSyntax(syntax QRegExp__PatternSyntax) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::setPatternSyntax")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_SetPatternSyntax(ptr.Pointer(), C.int(syntax))
	}
}

func (ptr *QRegExp) Swap(other QRegExp_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_Swap(ptr.Pointer(), PointerFromQRegExp(other))
	}
}

func (ptr *QRegExp) DestroyQRegExp() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QRegExp::~QRegExp")
		}
	}()

	if ptr.Pointer() != nil {
		C.QRegExp_DestroyQRegExp(ptr.Pointer())
	}
}
