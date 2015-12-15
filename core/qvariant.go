package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"strings"
	"unsafe"
)

type QVariant struct {
	ptr unsafe.Pointer
}

type QVariant_ITF interface {
	QVariant_PTR() *QVariant
}

func (p *QVariant) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QVariant) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQVariant(ptr QVariant_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVariant_PTR().Pointer()
	}
	return nil
}

func NewQVariantFromPointer(ptr unsafe.Pointer) *QVariant {
	var n = new(QVariant)
	n.SetPointer(ptr)
	return n
}

func (ptr *QVariant) QVariant_PTR() *QVariant {
	return ptr
}

func NewQVariant20(c QChar_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant20(PointerFromQChar(c)))
}

func NewQVariant18(val QLatin1String_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant18(PointerFromQLatin1String(val)))
}

func NewQVariant11(val bool) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant11(C.int(qt.GoBoolToInt(val))))
}

func NewQVariant16(val QBitArray_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant16(PointerFromQBitArray(val)))
}

func NewQVariant15(val QByteArray_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant15(PointerFromQByteArray(val)))
}

func NewQVariant21(val QDate_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant21(PointerFromQDate(val)))
}

func NewQVariant23(val QDateTime_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant23(PointerFromQDateTime(val)))
}

func NewQVariant39(val QEasingCurve_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant39(PointerFromQEasingCurve(val)))
}

func NewQVariant45(val QJsonArray_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant45(PointerFromQJsonArray(val)))
}

func NewQVariant46(val QJsonDocument_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant46(PointerFromQJsonDocument(val)))
}

func NewQVariant44(val QJsonObject_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant44(PointerFromQJsonObject(val)))
}

func NewQVariant43(val QJsonValue_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant43(PointerFromQJsonValue(val)))
}

func NewQVariant31(val QLine_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant31(PointerFromQLine(val)))
}

func NewQVariant32(val QLineF_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant32(PointerFromQLineF(val)))
}

func NewQVariant35(l QLocale_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant35(PointerFromQLocale(l)))
}

func NewQVariant41(val QModelIndex_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant41(PointerFromQModelIndex(val)))
}

func NewQVariant42(val QPersistentModelIndex_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant42(PointerFromQPersistentModelIndex(val)))
}

func NewQVariant29(val QPoint_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant29(PointerFromQPoint(val)))
}

func NewQVariant30(val QPointF_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant30(PointerFromQPointF(val)))
}

func NewQVariant33(val QRect_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant33(PointerFromQRect(val)))
}

func NewQVariant34(val QRectF_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant34(PointerFromQRectF(val)))
}

func NewQVariant36(regExp QRegExp_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant36(PointerFromQRegExp(regExp)))
}

func NewQVariant37(re QRegularExpression_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant37(PointerFromQRegularExpression(re)))
}

func NewQVariant27(val QSize_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant27(PointerFromQSize(val)))
}

func NewQVariant28(val QSizeF_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant28(PointerFromQSizeF(val)))
}

func NewQVariant17(val string) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant17(C.CString(val)))
}

func NewQVariant19(val []string) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant19(C.CString(strings.Join(val, ",,,"))))
}

func NewQVariant22(val QTime_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant22(PointerFromQTime(val)))
}

func NewQVariant38(val QUrl_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant38(PointerFromQUrl(val)))
}

func NewQVariant40(val QUuid_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant40(PointerFromQUuid(val)))
}

func NewQVariant5(p QVariant_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant5(PointerFromQVariant(p)))
}

func NewQVariant14(val string) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant14(C.CString(val)))
}

func NewQVariant7(val int) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant7(C.int(val)))
}

func (ptr *QVariant) ToByteArray() *QByteArray {
	defer qt.Recovering("QVariant::toByteArray")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QVariant_ToByteArray(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToDateTime() *QDateTime {
	defer qt.Recovering("QVariant::toDateTime")

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QVariant_ToDateTime(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToEasingCurve() *QEasingCurve {
	defer qt.Recovering("QVariant::toEasingCurve")

	if ptr.Pointer() != nil {
		return NewQEasingCurveFromPointer(C.QVariant_ToEasingCurve(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToPoint() *QPoint {
	defer qt.Recovering("QVariant::toPoint")

	if ptr.Pointer() != nil {
		return NewQPointFromPointer(C.QVariant_ToPoint(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToRegExp() *QRegExp {
	defer qt.Recovering("QVariant::toRegExp")

	if ptr.Pointer() != nil {
		return NewQRegExpFromPointer(C.QVariant_ToRegExp(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToRegularExpression() *QRegularExpression {
	defer qt.Recovering("QVariant::toRegularExpression")

	if ptr.Pointer() != nil {
		return NewQRegularExpressionFromPointer(C.QVariant_ToRegularExpression(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToStringList() []string {
	defer qt.Recovering("QVariant::toStringList")

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QVariant_ToStringList(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QVariant) DestroyQVariant() {
	defer qt.Recovering("QVariant::~QVariant")

	if ptr.Pointer() != nil {
		C.QVariant_DestroyQVariant(ptr.Pointer())
	}
}

func NewQVariant() *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant())
}

func NewQVariant6(s QDataStream_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant6(PointerFromQDataStream(s)))
}

func NewQVariant47(other QVariant_ITF) *QVariant {
	defer qt.Recovering("QVariant::QVariant")

	return NewQVariantFromPointer(C.QVariant_NewQVariant47(PointerFromQVariant(other)))
}

func (ptr *QVariant) Clear() {
	defer qt.Recovering("QVariant::clear")

	if ptr.Pointer() != nil {
		C.QVariant_Clear(ptr.Pointer())
	}
}

func (ptr *QVariant) Convert(targetTypeId int) bool {
	defer qt.Recovering("QVariant::convert")

	if ptr.Pointer() != nil {
		return C.QVariant_Convert(ptr.Pointer(), C.int(targetTypeId)) != 0
	}
	return false
}

func (ptr *QVariant) IsNull() bool {
	defer qt.Recovering("QVariant::isNull")

	if ptr.Pointer() != nil {
		return C.QVariant_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVariant) IsValid() bool {
	defer qt.Recovering("QVariant::isValid")

	if ptr.Pointer() != nil {
		return C.QVariant_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVariant) Swap(other QVariant_ITF) {
	defer qt.Recovering("QVariant::swap")

	if ptr.Pointer() != nil {
		C.QVariant_Swap(ptr.Pointer(), PointerFromQVariant(other))
	}
}

func (ptr *QVariant) ToBool() bool {
	defer qt.Recovering("QVariant::toBool")

	if ptr.Pointer() != nil {
		return C.QVariant_ToBool(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QVariant) ToInt(ok bool) int {
	defer qt.Recovering("QVariant::toInt")

	if ptr.Pointer() != nil {
		return int(C.QVariant_ToInt(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QVariant) ToJsonArray() *QJsonArray {
	defer qt.Recovering("QVariant::toJsonArray")

	if ptr.Pointer() != nil {
		return NewQJsonArrayFromPointer(C.QVariant_ToJsonArray(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToJsonDocument() *QJsonDocument {
	defer qt.Recovering("QVariant::toJsonDocument")

	if ptr.Pointer() != nil {
		return NewQJsonDocumentFromPointer(C.QVariant_ToJsonDocument(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToJsonObject() *QJsonObject {
	defer qt.Recovering("QVariant::toJsonObject")

	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QVariant_ToJsonObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToModelIndex() *QModelIndex {
	defer qt.Recovering("QVariant::toModelIndex")

	if ptr.Pointer() != nil {
		return NewQModelIndexFromPointer(C.QVariant_ToModelIndex(ptr.Pointer()))
	}
	return nil
}

func (ptr *QVariant) ToReal(ok bool) float64 {
	defer qt.Recovering("QVariant::toReal")

	if ptr.Pointer() != nil {
		return float64(C.QVariant_ToReal(ptr.Pointer(), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QVariant) ToString() string {
	defer qt.Recovering("QVariant::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVariant_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QVariant) UserType() int {
	defer qt.Recovering("QVariant::userType")

	if ptr.Pointer() != nil {
		return int(C.QVariant_UserType(ptr.Pointer()))
	}
	return 0
}
