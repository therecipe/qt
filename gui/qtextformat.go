package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QTextFormat struct {
	ptr unsafe.Pointer
}

type QTextFormat_ITF interface {
	QTextFormat_PTR() *QTextFormat
}

func (p *QTextFormat) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTextFormat) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTextFormat(ptr QTextFormat_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTextFormat_PTR().Pointer()
	}
	return nil
}

func NewQTextFormatFromPointer(ptr unsafe.Pointer) *QTextFormat {
	var n = new(QTextFormat)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTextFormat) QTextFormat_PTR() *QTextFormat {
	return ptr
}

//QTextFormat::FormatType
type QTextFormat__FormatType int64

const (
	QTextFormat__InvalidFormat = QTextFormat__FormatType(-1)
	QTextFormat__BlockFormat   = QTextFormat__FormatType(1)
	QTextFormat__CharFormat    = QTextFormat__FormatType(2)
	QTextFormat__ListFormat    = QTextFormat__FormatType(3)
	QTextFormat__TableFormat   = QTextFormat__FormatType(4)
	QTextFormat__FrameFormat   = QTextFormat__FormatType(5)
	QTextFormat__UserFormat    = QTextFormat__FormatType(100)
)

//QTextFormat::ObjectTypes
type QTextFormat__ObjectTypes int64

const (
	QTextFormat__NoObject        = QTextFormat__ObjectTypes(0)
	QTextFormat__ImageObject     = QTextFormat__ObjectTypes(1)
	QTextFormat__TableObject     = QTextFormat__ObjectTypes(2)
	QTextFormat__TableCellObject = QTextFormat__ObjectTypes(3)
	QTextFormat__UserObject      = QTextFormat__ObjectTypes(0x1000)
)

//QTextFormat::PageBreakFlag
type QTextFormat__PageBreakFlag int64

const (
	QTextFormat__PageBreak_Auto         = QTextFormat__PageBreakFlag(0)
	QTextFormat__PageBreak_AlwaysBefore = QTextFormat__PageBreakFlag(0x001)
	QTextFormat__PageBreak_AlwaysAfter  = QTextFormat__PageBreakFlag(0x010)
)

//QTextFormat::Property
type QTextFormat__Property int64

const (
	QTextFormat__ObjectIndex                       = QTextFormat__Property(0x0)
	QTextFormat__CssFloat                          = QTextFormat__Property(0x0800)
	QTextFormat__LayoutDirection                   = QTextFormat__Property(0x0801)
	QTextFormat__OutlinePen                        = QTextFormat__Property(0x810)
	QTextFormat__BackgroundBrush                   = QTextFormat__Property(0x820)
	QTextFormat__ForegroundBrush                   = QTextFormat__Property(0x821)
	QTextFormat__BackgroundImageUrl                = QTextFormat__Property(0x823)
	QTextFormat__BlockAlignment                    = QTextFormat__Property(0x1010)
	QTextFormat__BlockTopMargin                    = QTextFormat__Property(0x1030)
	QTextFormat__BlockBottomMargin                 = QTextFormat__Property(0x1031)
	QTextFormat__BlockLeftMargin                   = QTextFormat__Property(0x1032)
	QTextFormat__BlockRightMargin                  = QTextFormat__Property(0x1033)
	QTextFormat__TextIndent                        = QTextFormat__Property(0x1034)
	QTextFormat__TabPositions                      = QTextFormat__Property(0x1035)
	QTextFormat__BlockIndent                       = QTextFormat__Property(0x1040)
	QTextFormat__LineHeight                        = QTextFormat__Property(0x1048)
	QTextFormat__LineHeightType                    = QTextFormat__Property(0x1049)
	QTextFormat__BlockNonBreakableLines            = QTextFormat__Property(0x1050)
	QTextFormat__BlockTrailingHorizontalRulerWidth = QTextFormat__Property(0x1060)
	QTextFormat__FirstFontProperty                 = QTextFormat__Property(0x1FE0)
	QTextFormat__FontCapitalization                = QTextFormat__Property(QTextFormat__FirstFontProperty)
	QTextFormat__FontLetterSpacingType             = QTextFormat__Property(0x2033)
	QTextFormat__FontLetterSpacing                 = QTextFormat__Property(0x1FE1)
	QTextFormat__FontWordSpacing                   = QTextFormat__Property(0x1FE2)
	QTextFormat__FontStretch                       = QTextFormat__Property(0x2034)
	QTextFormat__FontStyleHint                     = QTextFormat__Property(0x1FE3)
	QTextFormat__FontStyleStrategy                 = QTextFormat__Property(0x1FE4)
	QTextFormat__FontKerning                       = QTextFormat__Property(0x1FE5)
	QTextFormat__FontHintingPreference             = QTextFormat__Property(0x1FE6)
	QTextFormat__FontFamily                        = QTextFormat__Property(0x2000)
	QTextFormat__FontPointSize                     = QTextFormat__Property(0x2001)
	QTextFormat__FontSizeAdjustment                = QTextFormat__Property(0x2002)
	QTextFormat__FontSizeIncrement                 = QTextFormat__Property(QTextFormat__FontSizeAdjustment)
	QTextFormat__FontWeight                        = QTextFormat__Property(0x2003)
	QTextFormat__FontItalic                        = QTextFormat__Property(0x2004)
	QTextFormat__FontUnderline                     = QTextFormat__Property(0x2005)
	QTextFormat__FontOverline                      = QTextFormat__Property(0x2006)
	QTextFormat__FontStrikeOut                     = QTextFormat__Property(0x2007)
	QTextFormat__FontFixedPitch                    = QTextFormat__Property(0x2008)
	QTextFormat__FontPixelSize                     = QTextFormat__Property(0x2009)
	QTextFormat__LastFontProperty                  = QTextFormat__Property(QTextFormat__FontPixelSize)
	QTextFormat__TextUnderlineColor                = QTextFormat__Property(0x2010)
	QTextFormat__TextVerticalAlignment             = QTextFormat__Property(0x2021)
	QTextFormat__TextOutline                       = QTextFormat__Property(0x2022)
	QTextFormat__TextUnderlineStyle                = QTextFormat__Property(0x2023)
	QTextFormat__TextToolTip                       = QTextFormat__Property(0x2024)
	QTextFormat__IsAnchor                          = QTextFormat__Property(0x2030)
	QTextFormat__AnchorHref                        = QTextFormat__Property(0x2031)
	QTextFormat__AnchorName                        = QTextFormat__Property(0x2032)
	QTextFormat__ObjectType                        = QTextFormat__Property(0x2f00)
	QTextFormat__ListStyle                         = QTextFormat__Property(0x3000)
	QTextFormat__ListIndent                        = QTextFormat__Property(0x3001)
	QTextFormat__ListNumberPrefix                  = QTextFormat__Property(0x3002)
	QTextFormat__ListNumberSuffix                  = QTextFormat__Property(0x3003)
	QTextFormat__FrameBorder                       = QTextFormat__Property(0x4000)
	QTextFormat__FrameMargin                       = QTextFormat__Property(0x4001)
	QTextFormat__FramePadding                      = QTextFormat__Property(0x4002)
	QTextFormat__FrameWidth                        = QTextFormat__Property(0x4003)
	QTextFormat__FrameHeight                       = QTextFormat__Property(0x4004)
	QTextFormat__FrameTopMargin                    = QTextFormat__Property(0x4005)
	QTextFormat__FrameBottomMargin                 = QTextFormat__Property(0x4006)
	QTextFormat__FrameLeftMargin                   = QTextFormat__Property(0x4007)
	QTextFormat__FrameRightMargin                  = QTextFormat__Property(0x4008)
	QTextFormat__FrameBorderBrush                  = QTextFormat__Property(0x4009)
	QTextFormat__FrameBorderStyle                  = QTextFormat__Property(0x4010)
	QTextFormat__TableColumns                      = QTextFormat__Property(0x4100)
	QTextFormat__TableColumnWidthConstraints       = QTextFormat__Property(0x4101)
	QTextFormat__TableCellSpacing                  = QTextFormat__Property(0x4102)
	QTextFormat__TableCellPadding                  = QTextFormat__Property(0x4103)
	QTextFormat__TableHeaderRowCount               = QTextFormat__Property(0x4104)
	QTextFormat__TableCellRowSpan                  = QTextFormat__Property(0x4810)
	QTextFormat__TableCellColumnSpan               = QTextFormat__Property(0x4811)
	QTextFormat__TableCellTopPadding               = QTextFormat__Property(0x4812)
	QTextFormat__TableCellBottomPadding            = QTextFormat__Property(0x4813)
	QTextFormat__TableCellLeftPadding              = QTextFormat__Property(0x4814)
	QTextFormat__TableCellRightPadding             = QTextFormat__Property(0x4815)
	QTextFormat__ImageName                         = QTextFormat__Property(0x5000)
	QTextFormat__ImageWidth                        = QTextFormat__Property(0x5010)
	QTextFormat__ImageHeight                       = QTextFormat__Property(0x5011)
	QTextFormat__FullWidthSelection                = QTextFormat__Property(0x06000)
	QTextFormat__PageBreakPolicy                   = QTextFormat__Property(0x7000)
	QTextFormat__UserProperty                      = QTextFormat__Property(0x100000)
)

func NewQTextFormat3(other QTextFormat_ITF) *QTextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::QTextFormat")
		}
	}()

	return NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat3(PointerFromQTextFormat(other)))
}

func (ptr *QTextFormat) SetObjectIndex(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setObjectIndex")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetObjectIndex(ptr.Pointer(), C.int(index))
	}
}

func NewQTextFormat() *QTextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::QTextFormat")
		}
	}()

	return NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat())
}

func NewQTextFormat2(ty int) *QTextFormat {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::QTextFormat")
		}
	}()

	return NewQTextFormatFromPointer(C.QTextFormat_NewQTextFormat2(C.int(ty)))
}

func (ptr *QTextFormat) Background() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::background")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QTextFormat_Background(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFormat) BoolProperty(propertyId int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::boolProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_BoolProperty(ptr.Pointer(), C.int(propertyId)) != 0
	}
	return false
}

func (ptr *QTextFormat) BrushProperty(propertyId int) *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::brushProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QTextFormat_BrushProperty(ptr.Pointer(), C.int(propertyId)))
	}
	return nil
}

func (ptr *QTextFormat) ClearBackground() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::clearBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_ClearBackground(ptr.Pointer())
	}
}

func (ptr *QTextFormat) ClearForeground() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::clearForeground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_ClearForeground(ptr.Pointer())
	}
}

func (ptr *QTextFormat) ClearProperty(propertyId int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::clearProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_ClearProperty(ptr.Pointer(), C.int(propertyId))
	}
}

func (ptr *QTextFormat) ColorProperty(propertyId int) *QColor {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::colorProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQColorFromPointer(C.QTextFormat_ColorProperty(ptr.Pointer(), C.int(propertyId)))
	}
	return nil
}

func (ptr *QTextFormat) DoubleProperty(propertyId int) float64 {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::doubleProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return float64(C.QTextFormat_DoubleProperty(ptr.Pointer(), C.int(propertyId)))
	}
	return 0
}

func (ptr *QTextFormat) Foreground() *QBrush {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::foreground")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQBrushFromPointer(C.QTextFormat_Foreground(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTextFormat) HasProperty(propertyId int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::hasProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_HasProperty(ptr.Pointer(), C.int(propertyId)) != 0
	}
	return false
}

func (ptr *QTextFormat) IntProperty(propertyId int) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::intProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFormat_IntProperty(ptr.Pointer(), C.int(propertyId)))
	}
	return 0
}

func (ptr *QTextFormat) IsBlockFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isBlockFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsBlockFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsCharFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isCharFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsCharFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsEmpty() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isEmpty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsFrameFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isFrameFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsFrameFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsImageFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isImageFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsImageFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsListFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isListFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsListFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsTableCellFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isTableCellFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsTableCellFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsTableFormat() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isTableFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsTableFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QTextFormat_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QTextFormat) LayoutDirection() core.Qt__LayoutDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::layoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__LayoutDirection(C.QTextFormat_LayoutDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFormat) Merge(other QTextFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::merge")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_Merge(ptr.Pointer(), PointerFromQTextFormat(other))
	}
}

func (ptr *QTextFormat) ObjectIndex() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::objectIndex")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFormat_ObjectIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFormat) ObjectType() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::objectType")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFormat_ObjectType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFormat) Property(propertyId int) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::property")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QTextFormat_Property(ptr.Pointer(), C.int(propertyId)))
	}
	return nil
}

func (ptr *QTextFormat) PropertyCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::propertyCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFormat_PropertyCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFormat) SetBackground(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setBackground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetBackground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFormat) SetForeground(brush QBrush_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setForeground")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetForeground(ptr.Pointer(), PointerFromQBrush(brush))
	}
}

func (ptr *QTextFormat) SetLayoutDirection(direction core.Qt__LayoutDirection) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setLayoutDirection")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetLayoutDirection(ptr.Pointer(), C.int(direction))
	}
}

func (ptr *QTextFormat) SetObjectType(ty int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setObjectType")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetObjectType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QTextFormat) SetProperty(propertyId int, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::setProperty")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_SetProperty(ptr.Pointer(), C.int(propertyId), core.PointerFromQVariant(value))
	}
}

func (ptr *QTextFormat) StringProperty(propertyId int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::stringProperty")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QTextFormat_StringProperty(ptr.Pointer(), C.int(propertyId)))
	}
	return ""
}

func (ptr *QTextFormat) Swap(other QTextFormat_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_Swap(ptr.Pointer(), PointerFromQTextFormat(other))
	}
}

func (ptr *QTextFormat) Type() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::type")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QTextFormat_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTextFormat) DestroyQTextFormat() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QTextFormat::~QTextFormat")
		}
	}()

	if ptr.Pointer() != nil {
		C.QTextFormat_DestroyQTextFormat(ptr.Pointer())
	}
}
