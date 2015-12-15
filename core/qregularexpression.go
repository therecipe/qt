package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QRegularExpression struct {
	ptr unsafe.Pointer
}

type QRegularExpression_ITF interface {
	QRegularExpression_PTR() *QRegularExpression
}

func (p *QRegularExpression) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegularExpression) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegularExpression(ptr QRegularExpression_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpression_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionFromPointer(ptr unsafe.Pointer) *QRegularExpression {
	var n = new(QRegularExpression)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegularExpression) QRegularExpression_PTR() *QRegularExpression {
	return ptr
}

//QRegularExpression::MatchOption
type QRegularExpression__MatchOption int64

const (
	QRegularExpression__NoMatchOption                     = QRegularExpression__MatchOption(0x0000)
	QRegularExpression__AnchoredMatchOption               = QRegularExpression__MatchOption(0x0001)
	QRegularExpression__DontCheckSubjectStringMatchOption = QRegularExpression__MatchOption(0x0002)
)

//QRegularExpression::MatchType
type QRegularExpression__MatchType int64

const (
	QRegularExpression__NormalMatch                = QRegularExpression__MatchType(0)
	QRegularExpression__PartialPreferCompleteMatch = QRegularExpression__MatchType(1)
	QRegularExpression__PartialPreferFirstMatch    = QRegularExpression__MatchType(2)
	QRegularExpression__NoMatch                    = QRegularExpression__MatchType(3)
)

//QRegularExpression::PatternOption
type QRegularExpression__PatternOption int64

const (
	QRegularExpression__NoPatternOption                 = QRegularExpression__PatternOption(0x0000)
	QRegularExpression__CaseInsensitiveOption           = QRegularExpression__PatternOption(0x0001)
	QRegularExpression__DotMatchesEverythingOption      = QRegularExpression__PatternOption(0x0002)
	QRegularExpression__MultilineOption                 = QRegularExpression__PatternOption(0x0004)
	QRegularExpression__ExtendedPatternSyntaxOption     = QRegularExpression__PatternOption(0x0008)
	QRegularExpression__InvertedGreedinessOption        = QRegularExpression__PatternOption(0x0010)
	QRegularExpression__DontCaptureOption               = QRegularExpression__PatternOption(0x0020)
	QRegularExpression__UseUnicodePropertiesOption      = QRegularExpression__PatternOption(0x0040)
	QRegularExpression__OptimizeOnFirstUsageOption      = QRegularExpression__PatternOption(0x0080)
	QRegularExpression__DontAutomaticallyOptimizeOption = QRegularExpression__PatternOption(0x0100)
)

func NewQRegularExpression() *QRegularExpression {
	defer qt.Recovering("QRegularExpression::QRegularExpression")

	return NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression())
}

func NewQRegularExpression3(re QRegularExpression_ITF) *QRegularExpression {
	defer qt.Recovering("QRegularExpression::QRegularExpression")

	return NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression3(PointerFromQRegularExpression(re)))
}

func NewQRegularExpression2(pattern string, options QRegularExpression__PatternOption) *QRegularExpression {
	defer qt.Recovering("QRegularExpression::QRegularExpression")

	return NewQRegularExpressionFromPointer(C.QRegularExpression_NewQRegularExpression2(C.CString(pattern), C.int(options)))
}

func (ptr *QRegularExpression) CaptureCount() int {
	defer qt.Recovering("QRegularExpression::captureCount")

	if ptr.Pointer() != nil {
		return int(C.QRegularExpression_CaptureCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpression) ErrorString() string {
	defer qt.Recovering("QRegularExpression::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpression_ErrorString(ptr.Pointer()))
	}
	return ""
}

func QRegularExpression_Escape(str string) string {
	defer qt.Recovering("QRegularExpression::escape")

	return C.GoString(C.QRegularExpression_QRegularExpression_Escape(C.CString(str)))
}

func (ptr *QRegularExpression) GlobalMatch(subject string, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatchIterator {
	defer qt.Recovering("QRegularExpression::globalMatch")

	if ptr.Pointer() != nil {
		return NewQRegularExpressionMatchIteratorFromPointer(C.QRegularExpression_GlobalMatch(ptr.Pointer(), C.CString(subject), C.int(offset), C.int(matchType), C.int(matchOptions)))
	}
	return nil
}

func (ptr *QRegularExpression) GlobalMatch2(subjectRef QStringRef_ITF, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatchIterator {
	defer qt.Recovering("QRegularExpression::globalMatch")

	if ptr.Pointer() != nil {
		return NewQRegularExpressionMatchIteratorFromPointer(C.QRegularExpression_GlobalMatch2(ptr.Pointer(), PointerFromQStringRef(subjectRef), C.int(offset), C.int(matchType), C.int(matchOptions)))
	}
	return nil
}

func (ptr *QRegularExpression) IsValid() bool {
	defer qt.Recovering("QRegularExpression::isValid")

	if ptr.Pointer() != nil {
		return C.QRegularExpression_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QRegularExpression) Match(subject string, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	defer qt.Recovering("QRegularExpression::match")

	if ptr.Pointer() != nil {
		return NewQRegularExpressionMatchFromPointer(C.QRegularExpression_Match(ptr.Pointer(), C.CString(subject), C.int(offset), C.int(matchType), C.int(matchOptions)))
	}
	return nil
}

func (ptr *QRegularExpression) Match2(subjectRef QStringRef_ITF, offset int, matchType QRegularExpression__MatchType, matchOptions QRegularExpression__MatchOption) *QRegularExpressionMatch {
	defer qt.Recovering("QRegularExpression::match")

	if ptr.Pointer() != nil {
		return NewQRegularExpressionMatchFromPointer(C.QRegularExpression_Match2(ptr.Pointer(), PointerFromQStringRef(subjectRef), C.int(offset), C.int(matchType), C.int(matchOptions)))
	}
	return nil
}

func (ptr *QRegularExpression) NamedCaptureGroups() []string {
	defer qt.Recovering("QRegularExpression::namedCaptureGroups")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QRegularExpression_NamedCaptureGroups(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QRegularExpression) Optimize() {
	defer qt.Recovering("QRegularExpression::optimize")

	if ptr.Pointer() != nil {
		C.QRegularExpression_Optimize(ptr.Pointer())
	}
}

func (ptr *QRegularExpression) Pattern() string {
	defer qt.Recovering("QRegularExpression::pattern")

	if ptr.Pointer() != nil {
		return C.GoString(C.QRegularExpression_Pattern(ptr.Pointer()))
	}
	return ""
}

func (ptr *QRegularExpression) PatternErrorOffset() int {
	defer qt.Recovering("QRegularExpression::patternErrorOffset")

	if ptr.Pointer() != nil {
		return int(C.QRegularExpression_PatternErrorOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpression) PatternOptions() QRegularExpression__PatternOption {
	defer qt.Recovering("QRegularExpression::patternOptions")

	if ptr.Pointer() != nil {
		return QRegularExpression__PatternOption(C.QRegularExpression_PatternOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QRegularExpression) SetPattern(pattern string) {
	defer qt.Recovering("QRegularExpression::setPattern")

	if ptr.Pointer() != nil {
		C.QRegularExpression_SetPattern(ptr.Pointer(), C.CString(pattern))
	}
}

func (ptr *QRegularExpression) SetPatternOptions(options QRegularExpression__PatternOption) {
	defer qt.Recovering("QRegularExpression::setPatternOptions")

	if ptr.Pointer() != nil {
		C.QRegularExpression_SetPatternOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QRegularExpression) Swap(other QRegularExpression_ITF) {
	defer qt.Recovering("QRegularExpression::swap")

	if ptr.Pointer() != nil {
		C.QRegularExpression_Swap(ptr.Pointer(), PointerFromQRegularExpression(other))
	}
}

func (ptr *QRegularExpression) DestroyQRegularExpression() {
	defer qt.Recovering("QRegularExpression::~QRegularExpression")

	if ptr.Pointer() != nil {
		C.QRegularExpression_DestroyQRegularExpression(ptr.Pointer())
	}
}
