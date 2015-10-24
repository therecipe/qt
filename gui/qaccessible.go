package gui

//#include "qaccessible.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessible struct {
	ptr unsafe.Pointer
}

type QAccessibleITF interface {
	QAccessiblePTR() *QAccessible
}

func (p *QAccessible) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAccessible) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAccessible(ptr QAccessibleITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessiblePTR().Pointer()
	}
	return nil
}

func QAccessibleFromPointer(ptr unsafe.Pointer) *QAccessible {
	var n = new(QAccessible)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessible) QAccessiblePTR() *QAccessible {
	return ptr
}

//QAccessible::Event
type QAccessible__Event int

var (
	QAccessible__SoundPlayed                     = QAccessible__Event(0x0001)
	QAccessible__Alert                           = QAccessible__Event(0x0002)
	QAccessible__ForegroundChanged               = QAccessible__Event(0x0003)
	QAccessible__MenuStart                       = QAccessible__Event(0x0004)
	QAccessible__MenuEnd                         = QAccessible__Event(0x0005)
	QAccessible__PopupMenuStart                  = QAccessible__Event(0x0006)
	QAccessible__PopupMenuEnd                    = QAccessible__Event(0x0007)
	QAccessible__ContextHelpStart                = QAccessible__Event(0x000C)
	QAccessible__ContextHelpEnd                  = QAccessible__Event(0x000D)
	QAccessible__DragDropStart                   = QAccessible__Event(0x000E)
	QAccessible__DragDropEnd                     = QAccessible__Event(0x000F)
	QAccessible__DialogStart                     = QAccessible__Event(0x0010)
	QAccessible__DialogEnd                       = QAccessible__Event(0x0011)
	QAccessible__ScrollingStart                  = QAccessible__Event(0x0012)
	QAccessible__ScrollingEnd                    = QAccessible__Event(0x0013)
	QAccessible__MenuCommand                     = QAccessible__Event(0x0018)
	QAccessible__ActionChanged                   = QAccessible__Event(0x0101)
	QAccessible__ActiveDescendantChanged         = QAccessible__Event(0x0102)
	QAccessible__AttributeChanged                = QAccessible__Event(0x0103)
	QAccessible__DocumentContentChanged          = QAccessible__Event(0x0104)
	QAccessible__DocumentLoadComplete            = QAccessible__Event(0x0105)
	QAccessible__DocumentLoadStopped             = QAccessible__Event(0x0106)
	QAccessible__DocumentReload                  = QAccessible__Event(0x0107)
	QAccessible__HyperlinkEndIndexChanged        = QAccessible__Event(0x0108)
	QAccessible__HyperlinkNumberOfAnchorsChanged = QAccessible__Event(0x0109)
	QAccessible__HyperlinkSelectedLinkChanged    = QAccessible__Event(0x010A)
	QAccessible__HypertextLinkActivated          = QAccessible__Event(0x010B)
	QAccessible__HypertextLinkSelected           = QAccessible__Event(0x010C)
	QAccessible__HyperlinkStartIndexChanged      = QAccessible__Event(0x010D)
	QAccessible__HypertextChanged                = QAccessible__Event(0x010E)
	QAccessible__HypertextNLinksChanged          = QAccessible__Event(0x010F)
	QAccessible__ObjectAttributeChanged          = QAccessible__Event(0x0110)
	QAccessible__PageChanged                     = QAccessible__Event(0x0111)
	QAccessible__SectionChanged                  = QAccessible__Event(0x0112)
	QAccessible__TableCaptionChanged             = QAccessible__Event(0x0113)
	QAccessible__TableColumnDescriptionChanged   = QAccessible__Event(0x0114)
	QAccessible__TableColumnHeaderChanged        = QAccessible__Event(0x0115)
	QAccessible__TableModelChanged               = QAccessible__Event(0x0116)
	QAccessible__TableRowDescriptionChanged      = QAccessible__Event(0x0117)
	QAccessible__TableRowHeaderChanged           = QAccessible__Event(0x0118)
	QAccessible__TableSummaryChanged             = QAccessible__Event(0x0119)
	QAccessible__TextAttributeChanged            = QAccessible__Event(0x011A)
	QAccessible__TextCaretMoved                  = QAccessible__Event(0x011B)
	QAccessible__TextColumnChanged               = QAccessible__Event(0x011D)
	QAccessible__TextInserted                    = QAccessible__Event(0x011E)
	QAccessible__TextRemoved                     = QAccessible__Event(0x011F)
	QAccessible__TextUpdated                     = QAccessible__Event(0x0120)
	QAccessible__TextSelectionChanged            = QAccessible__Event(0x0121)
	QAccessible__VisibleDataChanged              = QAccessible__Event(0x0122)
	QAccessible__ObjectCreated                   = QAccessible__Event(0x8000)
	QAccessible__ObjectDestroyed                 = QAccessible__Event(0x8001)
	QAccessible__ObjectShow                      = QAccessible__Event(0x8002)
	QAccessible__ObjectHide                      = QAccessible__Event(0x8003)
	QAccessible__ObjectReorder                   = QAccessible__Event(0x8004)
	QAccessible__Focus                           = QAccessible__Event(0x8005)
	QAccessible__Selection                       = QAccessible__Event(0x8006)
	QAccessible__SelectionAdd                    = QAccessible__Event(0x8007)
	QAccessible__SelectionRemove                 = QAccessible__Event(0x8008)
	QAccessible__SelectionWithin                 = QAccessible__Event(0x8009)
	QAccessible__StateChanged                    = QAccessible__Event(0x800A)
	QAccessible__LocationChanged                 = QAccessible__Event(0x800B)
	QAccessible__NameChanged                     = QAccessible__Event(0x800C)
	QAccessible__DescriptionChanged              = QAccessible__Event(0x800D)
	QAccessible__ValueChanged                    = QAccessible__Event(0x800E)
	QAccessible__ParentChanged                   = QAccessible__Event(0x800F)
	QAccessible__HelpChanged                     = QAccessible__Event(0x80A0)
	QAccessible__DefaultActionChanged            = QAccessible__Event(0x80B0)
	QAccessible__AcceleratorChanged              = QAccessible__Event(0x80C0)
	QAccessible__InvalidEvent                    = QAccessible__Event(C.QAccessible_InvalidEvent_Type())
)

//QAccessible::InterfaceType
type QAccessible__InterfaceType int

var (
	QAccessible__TextInterface         = QAccessible__InterfaceType(0)
	QAccessible__EditableTextInterface = QAccessible__InterfaceType(1)
	QAccessible__ValueInterface        = QAccessible__InterfaceType(2)
	QAccessible__ActionInterface       = QAccessible__InterfaceType(3)
	QAccessible__ImageInterface        = QAccessible__InterfaceType(4)
	QAccessible__TableInterface        = QAccessible__InterfaceType(5)
	QAccessible__TableCellInterface    = QAccessible__InterfaceType(6)
)

//QAccessible::RelationFlag
type QAccessible__RelationFlag int

var (
	QAccessible__Label        = QAccessible__RelationFlag(0x00000001)
	QAccessible__Labelled     = QAccessible__RelationFlag(0x00000002)
	QAccessible__Controller   = QAccessible__RelationFlag(0x00000004)
	QAccessible__Controlled   = QAccessible__RelationFlag(0x00000008)
	QAccessible__AllRelations = QAccessible__RelationFlag(0xffffffff)
)

//QAccessible::Role
type QAccessible__Role int

var (
	QAccessible__NoRole               = QAccessible__Role(0x00000000)
	QAccessible__TitleBar             = QAccessible__Role(0x00000001)
	QAccessible__MenuBar              = QAccessible__Role(0x00000002)
	QAccessible__ScrollBar            = QAccessible__Role(0x00000003)
	QAccessible__Grip                 = QAccessible__Role(0x00000004)
	QAccessible__Sound                = QAccessible__Role(0x00000005)
	QAccessible__Cursor               = QAccessible__Role(0x00000006)
	QAccessible__Caret                = QAccessible__Role(0x00000007)
	QAccessible__AlertMessage         = QAccessible__Role(0x00000008)
	QAccessible__Window               = QAccessible__Role(0x00000009)
	QAccessible__Client               = QAccessible__Role(0x0000000A)
	QAccessible__PopupMenu            = QAccessible__Role(0x0000000B)
	QAccessible__MenuItem             = QAccessible__Role(0x0000000C)
	QAccessible__ToolTip              = QAccessible__Role(0x0000000D)
	QAccessible__Application          = QAccessible__Role(0x0000000E)
	QAccessible__Document             = QAccessible__Role(0x0000000F)
	QAccessible__Pane                 = QAccessible__Role(0x00000010)
	QAccessible__Chart                = QAccessible__Role(0x00000011)
	QAccessible__Dialog               = QAccessible__Role(0x00000012)
	QAccessible__Border               = QAccessible__Role(0x00000013)
	QAccessible__Grouping             = QAccessible__Role(0x00000014)
	QAccessible__Separator            = QAccessible__Role(0x00000015)
	QAccessible__ToolBar              = QAccessible__Role(0x00000016)
	QAccessible__StatusBar            = QAccessible__Role(0x00000017)
	QAccessible__Table                = QAccessible__Role(0x00000018)
	QAccessible__ColumnHeader         = QAccessible__Role(0x00000019)
	QAccessible__RowHeader            = QAccessible__Role(0x0000001A)
	QAccessible__Column               = QAccessible__Role(0x0000001B)
	QAccessible__Row                  = QAccessible__Role(0x0000001C)
	QAccessible__Cell                 = QAccessible__Role(0x0000001D)
	QAccessible__Link                 = QAccessible__Role(0x0000001E)
	QAccessible__HelpBalloon          = QAccessible__Role(0x0000001F)
	QAccessible__Assistant            = QAccessible__Role(0x00000020)
	QAccessible__List                 = QAccessible__Role(0x00000021)
	QAccessible__ListItem             = QAccessible__Role(0x00000022)
	QAccessible__Tree                 = QAccessible__Role(0x00000023)
	QAccessible__TreeItem             = QAccessible__Role(0x00000024)
	QAccessible__PageTab              = QAccessible__Role(0x00000025)
	QAccessible__PropertyPage         = QAccessible__Role(0x00000026)
	QAccessible__Indicator            = QAccessible__Role(0x00000027)
	QAccessible__Graphic              = QAccessible__Role(0x00000028)
	QAccessible__StaticText           = QAccessible__Role(0x00000029)
	QAccessible__EditableText         = QAccessible__Role(0x0000002A)
	QAccessible__Button               = QAccessible__Role(0x0000002B)
	QAccessible__CheckBox             = QAccessible__Role(0x0000002C)
	QAccessible__RadioButton          = QAccessible__Role(0x0000002D)
	QAccessible__ComboBox             = QAccessible__Role(0x0000002E)
	QAccessible__ProgressBar          = QAccessible__Role(0x00000030)
	QAccessible__Dial                 = QAccessible__Role(0x00000031)
	QAccessible__HotkeyField          = QAccessible__Role(0x00000032)
	QAccessible__Slider               = QAccessible__Role(0x00000033)
	QAccessible__SpinBox              = QAccessible__Role(0x00000034)
	QAccessible__Canvas               = QAccessible__Role(0x00000035)
	QAccessible__Animation            = QAccessible__Role(0x00000036)
	QAccessible__Equation             = QAccessible__Role(0x00000037)
	QAccessible__ButtonDropDown       = QAccessible__Role(0x00000038)
	QAccessible__ButtonMenu           = QAccessible__Role(0x00000039)
	QAccessible__ButtonDropGrid       = QAccessible__Role(0x0000003A)
	QAccessible__Whitespace           = QAccessible__Role(0x0000003B)
	QAccessible__PageTabList          = QAccessible__Role(0x0000003C)
	QAccessible__Clock                = QAccessible__Role(0x0000003D)
	QAccessible__Splitter             = QAccessible__Role(0x0000003E)
	QAccessible__LayeredPane          = QAccessible__Role(0x00000080)
	QAccessible__Terminal             = QAccessible__Role(0x00000081)
	QAccessible__Desktop              = QAccessible__Role(0x00000082)
	QAccessible__Paragraph            = QAccessible__Role(0x00000083)
	QAccessible__WebDocument          = QAccessible__Role(0x00000084)
	QAccessible__Section              = QAccessible__Role(0x00000085)
	QAccessible__ColorChooser         = QAccessible__Role(0x404)
	QAccessible__Footer               = QAccessible__Role(0x40E)
	QAccessible__Form                 = QAccessible__Role(0x410)
	QAccessible__Heading              = QAccessible__Role(0x414)
	QAccessible__Note                 = QAccessible__Role(0x41B)
	QAccessible__ComplementaryContent = QAccessible__Role(0x42C)
	QAccessible__UserRole             = QAccessible__Role(0x0000ffff)
)

//QAccessible::Text
type QAccessible__Text int

var (
	QAccessible__Name             = QAccessible__Text(0)
	QAccessible__Description      = QAccessible__Text(1)
	QAccessible__Value            = QAccessible__Text(2)
	QAccessible__Help             = QAccessible__Text(3)
	QAccessible__Accelerator      = QAccessible__Text(4)
	QAccessible__DebugDescription = QAccessible__Text(5)
	QAccessible__UserText         = QAccessible__Text(0x0000ffff)
)

//QAccessible::TextBoundaryType
type QAccessible__TextBoundaryType int

var (
	QAccessible__CharBoundary      = QAccessible__TextBoundaryType(0)
	QAccessible__WordBoundary      = QAccessible__TextBoundaryType(1)
	QAccessible__SentenceBoundary  = QAccessible__TextBoundaryType(2)
	QAccessible__ParagraphBoundary = QAccessible__TextBoundaryType(3)
	QAccessible__LineBoundary      = QAccessible__TextBoundaryType(4)
	QAccessible__NoBoundary        = QAccessible__TextBoundaryType(5)
)

func QAccessible_IsActive() bool {
	return C.QAccessible_QAccessible_IsActive() != 0
}

func QAccessible_QueryAccessibleInterface(object core.QObjectITF) *QAccessibleInterface {
	return QAccessibleInterfaceFromPointer(unsafe.Pointer(C.QAccessible_QAccessible_QueryAccessibleInterface(C.QtObjectPtr(core.PointerFromQObject(object)))))
}

func QAccessible_SetRootObject(object core.QObjectITF) {
	C.QAccessible_QAccessible_SetRootObject(C.QtObjectPtr(core.PointerFromQObject(object)))
}

func QAccessible_UpdateAccessibility(event QAccessibleEventITF) {
	C.QAccessible_QAccessible_UpdateAccessibility(C.QtObjectPtr(PointerFromQAccessibleEvent(event)))
}
