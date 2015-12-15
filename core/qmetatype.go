package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMetaType struct {
	ptr unsafe.Pointer
}

type QMetaType_ITF interface {
	QMetaType_PTR() *QMetaType
}

func (p *QMetaType) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaType) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaType(ptr QMetaType_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaType_PTR().Pointer()
	}
	return nil
}

func NewQMetaTypeFromPointer(ptr unsafe.Pointer) *QMetaType {
	var n = new(QMetaType)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaType) QMetaType_PTR() *QMetaType {
	return ptr
}

//QMetaType::Type
type QMetaType__Type int64

const (
	QMetaType__UnknownType           = QMetaType__Type(0)
	QMetaType__Bool                  = QMetaType__Type(1)
	QMetaType__Int                   = QMetaType__Type(2)
	QMetaType__UInt                  = QMetaType__Type(3)
	QMetaType__LongLong              = QMetaType__Type(4)
	QMetaType__ULongLong             = QMetaType__Type(5)
	QMetaType__Double                = QMetaType__Type(6)
	QMetaType__Long                  = QMetaType__Type(32)
	QMetaType__Short                 = QMetaType__Type(33)
	QMetaType__Char                  = QMetaType__Type(34)
	QMetaType__ULong                 = QMetaType__Type(35)
	QMetaType__UShort                = QMetaType__Type(36)
	QMetaType__UChar                 = QMetaType__Type(37)
	QMetaType__Float                 = QMetaType__Type(38)
	QMetaType__VoidStar              = QMetaType__Type(31)
	QMetaType__QChar                 = QMetaType__Type(7)
	QMetaType__QString               = QMetaType__Type(10)
	QMetaType__QStringList           = QMetaType__Type(11)
	QMetaType__QByteArray            = QMetaType__Type(12)
	QMetaType__QBitArray             = QMetaType__Type(13)
	QMetaType__QDate                 = QMetaType__Type(14)
	QMetaType__QTime                 = QMetaType__Type(15)
	QMetaType__QDateTime             = QMetaType__Type(16)
	QMetaType__QUrl                  = QMetaType__Type(17)
	QMetaType__QLocale               = QMetaType__Type(18)
	QMetaType__QRect                 = QMetaType__Type(19)
	QMetaType__QRectF                = QMetaType__Type(20)
	QMetaType__QSize                 = QMetaType__Type(21)
	QMetaType__QSizeF                = QMetaType__Type(22)
	QMetaType__QLine                 = QMetaType__Type(23)
	QMetaType__QLineF                = QMetaType__Type(24)
	QMetaType__QPoint                = QMetaType__Type(25)
	QMetaType__QPointF               = QMetaType__Type(26)
	QMetaType__QRegExp               = QMetaType__Type(27)
	QMetaType__QEasingCurve          = QMetaType__Type(29)
	QMetaType__QUuid                 = QMetaType__Type(30)
	QMetaType__QVariant              = QMetaType__Type(41)
	QMetaType__QModelIndex           = QMetaType__Type(42)
	QMetaType__QPersistentModelIndex = QMetaType__Type(50)
	QMetaType__QRegularExpression    = QMetaType__Type(44)
	QMetaType__QJsonValue            = QMetaType__Type(45)
	QMetaType__QJsonObject           = QMetaType__Type(46)
	QMetaType__QJsonArray            = QMetaType__Type(47)
	QMetaType__QJsonDocument         = QMetaType__Type(48)
	QMetaType__QByteArrayList        = QMetaType__Type(49)
	QMetaType__QObjectStar           = QMetaType__Type(39)
	QMetaType__SChar                 = QMetaType__Type(40)
	QMetaType__Void                  = QMetaType__Type(43)
	QMetaType__QVariantMap           = QMetaType__Type(8)
	QMetaType__QVariantList          = QMetaType__Type(9)
	QMetaType__QVariantHash          = QMetaType__Type(28)
	QMetaType__QFont                 = QMetaType__Type(64)
	QMetaType__QPixmap               = QMetaType__Type(65)
	QMetaType__QBrush                = QMetaType__Type(66)
	QMetaType__QColor                = QMetaType__Type(67)
	QMetaType__QPalette              = QMetaType__Type(68)
	QMetaType__QIcon                 = QMetaType__Type(69)
	QMetaType__QImage                = QMetaType__Type(70)
	QMetaType__QPolygon              = QMetaType__Type(71)
	QMetaType__QRegion               = QMetaType__Type(72)
	QMetaType__QBitmap               = QMetaType__Type(73)
	QMetaType__QCursor               = QMetaType__Type(74)
	QMetaType__QKeySequence          = QMetaType__Type(75)
	QMetaType__QPen                  = QMetaType__Type(76)
	QMetaType__QTextLength           = QMetaType__Type(77)
	QMetaType__QTextFormat           = QMetaType__Type(78)
	QMetaType__QMatrix               = QMetaType__Type(79)
	QMetaType__QTransform            = QMetaType__Type(80)
	QMetaType__QMatrix4x4            = QMetaType__Type(81)
	QMetaType__QVector2D             = QMetaType__Type(82)
	QMetaType__QVector3D             = QMetaType__Type(83)
	QMetaType__QVector4D             = QMetaType__Type(84)
	QMetaType__QQuaternion           = QMetaType__Type(85)
	QMetaType__QPolygonF             = QMetaType__Type(86)
	QMetaType__QSizePolicy           = QMetaType__Type(121)
	QMetaType__User                  = QMetaType__Type(1024)
)

//QMetaType::TypeFlag
type QMetaType__TypeFlag int64

const (
	QMetaType__NeedsConstruction        = QMetaType__TypeFlag(0x1)
	QMetaType__NeedsDestruction         = QMetaType__TypeFlag(0x2)
	QMetaType__MovableType              = QMetaType__TypeFlag(0x4)
	QMetaType__PointerToQObject         = QMetaType__TypeFlag(0x8)
	QMetaType__IsEnumeration            = QMetaType__TypeFlag(0x10)
	QMetaType__SharedPointerToQObject   = QMetaType__TypeFlag(0x20)
	QMetaType__WeakPointerToQObject     = QMetaType__TypeFlag(0x40)
	QMetaType__TrackingPointerToQObject = QMetaType__TypeFlag(0x80)
	QMetaType__WasDeclaredAsMetaType    = QMetaType__TypeFlag(0x100)
	QMetaType__IsGadget                 = QMetaType__TypeFlag(0x200)
)

func NewQMetaType(typeId int) *QMetaType {
	defer qt.Recovering("QMetaType::QMetaType")

	return NewQMetaTypeFromPointer(C.QMetaType_NewQMetaType(C.int(typeId)))
}

func (ptr *QMetaType) Flags() QMetaType__TypeFlag {
	defer qt.Recovering("QMetaType::flags")

	if ptr.Pointer() != nil {
		return QMetaType__TypeFlag(C.QMetaType_Flags(ptr.Pointer()))
	}
	return 0
}

func QMetaType_IsRegistered(ty int) bool {
	defer qt.Recovering("QMetaType::isRegistered")

	return C.QMetaType_QMetaType_IsRegistered(C.int(ty)) != 0
}

func (ptr *QMetaType) IsRegistered2() bool {
	defer qt.Recovering("QMetaType::isRegistered")

	if ptr.Pointer() != nil {
		return C.QMetaType_IsRegistered2(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaType) IsValid() bool {
	defer qt.Recovering("QMetaType::isValid")

	if ptr.Pointer() != nil {
		return C.QMetaType_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaType) MetaObject() *QMetaObject {
	defer qt.Recovering("QMetaType::metaObject")

	if ptr.Pointer() != nil {
		return NewQMetaObjectFromPointer(C.QMetaType_MetaObject(ptr.Pointer()))
	}
	return nil
}

func QMetaType_MetaObjectForType(ty int) *QMetaObject {
	defer qt.Recovering("QMetaType::metaObjectForType")

	return NewQMetaObjectFromPointer(C.QMetaType_QMetaType_MetaObjectForType(C.int(ty)))
}

func QMetaType_SizeOf(ty int) int {
	defer qt.Recovering("QMetaType::sizeOf")

	return int(C.QMetaType_QMetaType_SizeOf(C.int(ty)))
}

func (ptr *QMetaType) SizeOf2() int {
	defer qt.Recovering("QMetaType::sizeOf")

	if ptr.Pointer() != nil {
		return int(C.QMetaType_SizeOf2(ptr.Pointer()))
	}
	return 0
}

func QMetaType_Type2(typeName QByteArray_ITF) int {
	defer qt.Recovering("QMetaType::type")

	return int(C.QMetaType_QMetaType_Type2(PointerFromQByteArray(typeName)))
}

func QMetaType_Type(typeName string) int {
	defer qt.Recovering("QMetaType::type")

	return int(C.QMetaType_QMetaType_Type(C.CString(typeName)))
}

func QMetaType_TypeFlags(ty int) QMetaType__TypeFlag {
	defer qt.Recovering("QMetaType::typeFlags")

	return QMetaType__TypeFlag(C.QMetaType_QMetaType_TypeFlags(C.int(ty)))
}

func (ptr *QMetaType) DestroyQMetaType() {
	defer qt.Recovering("QMetaType::~QMetaType")

	if ptr.Pointer() != nil {
		C.QMetaType_DestroyQMetaType(ptr.Pointer())
	}
}
