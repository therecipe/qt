package widgets

//#include "qgraphicsitem.h"
import "C"
import (
	"unsafe"
)

type QGraphicsItem struct {
	ptr unsafe.Pointer
}

type QGraphicsItem_ITF interface {
	QGraphicsItem_PTR() *QGraphicsItem
}

func (p *QGraphicsItem) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QGraphicsItem) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQGraphicsItem(ptr QGraphicsItem_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsItem_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsItemFromPointer(ptr unsafe.Pointer) *QGraphicsItem {
	var n = new(QGraphicsItem)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsItem) QGraphicsItem_PTR() *QGraphicsItem {
	return ptr
}

//QGraphicsItem::CacheMode
type QGraphicsItem__CacheMode int64

const (
	QGraphicsItem__NoCache               = QGraphicsItem__CacheMode(0)
	QGraphicsItem__ItemCoordinateCache   = QGraphicsItem__CacheMode(1)
	QGraphicsItem__DeviceCoordinateCache = QGraphicsItem__CacheMode(2)
)

//QGraphicsItem::GraphicsItemChange
type QGraphicsItem__GraphicsItemChange int64

const (
	QGraphicsItem__ItemPositionChange                 = QGraphicsItem__GraphicsItemChange(0)
	QGraphicsItem__ItemMatrixChange                   = QGraphicsItem__GraphicsItemChange(1)
	QGraphicsItem__ItemVisibleChange                  = QGraphicsItem__GraphicsItemChange(2)
	QGraphicsItem__ItemEnabledChange                  = QGraphicsItem__GraphicsItemChange(3)
	QGraphicsItem__ItemSelectedChange                 = QGraphicsItem__GraphicsItemChange(4)
	QGraphicsItem__ItemParentChange                   = QGraphicsItem__GraphicsItemChange(5)
	QGraphicsItem__ItemChildAddedChange               = QGraphicsItem__GraphicsItemChange(6)
	QGraphicsItem__ItemChildRemovedChange             = QGraphicsItem__GraphicsItemChange(7)
	QGraphicsItem__ItemTransformChange                = QGraphicsItem__GraphicsItemChange(8)
	QGraphicsItem__ItemPositionHasChanged             = QGraphicsItem__GraphicsItemChange(9)
	QGraphicsItem__ItemTransformHasChanged            = QGraphicsItem__GraphicsItemChange(10)
	QGraphicsItem__ItemSceneChange                    = QGraphicsItem__GraphicsItemChange(11)
	QGraphicsItem__ItemVisibleHasChanged              = QGraphicsItem__GraphicsItemChange(12)
	QGraphicsItem__ItemEnabledHasChanged              = QGraphicsItem__GraphicsItemChange(13)
	QGraphicsItem__ItemSelectedHasChanged             = QGraphicsItem__GraphicsItemChange(14)
	QGraphicsItem__ItemParentHasChanged               = QGraphicsItem__GraphicsItemChange(15)
	QGraphicsItem__ItemSceneHasChanged                = QGraphicsItem__GraphicsItemChange(16)
	QGraphicsItem__ItemCursorChange                   = QGraphicsItem__GraphicsItemChange(17)
	QGraphicsItem__ItemCursorHasChanged               = QGraphicsItem__GraphicsItemChange(18)
	QGraphicsItem__ItemToolTipChange                  = QGraphicsItem__GraphicsItemChange(19)
	QGraphicsItem__ItemToolTipHasChanged              = QGraphicsItem__GraphicsItemChange(20)
	QGraphicsItem__ItemFlagsChange                    = QGraphicsItem__GraphicsItemChange(21)
	QGraphicsItem__ItemFlagsHaveChanged               = QGraphicsItem__GraphicsItemChange(22)
	QGraphicsItem__ItemZValueChange                   = QGraphicsItem__GraphicsItemChange(23)
	QGraphicsItem__ItemZValueHasChanged               = QGraphicsItem__GraphicsItemChange(24)
	QGraphicsItem__ItemOpacityChange                  = QGraphicsItem__GraphicsItemChange(25)
	QGraphicsItem__ItemOpacityHasChanged              = QGraphicsItem__GraphicsItemChange(26)
	QGraphicsItem__ItemScenePositionHasChanged        = QGraphicsItem__GraphicsItemChange(27)
	QGraphicsItem__ItemRotationChange                 = QGraphicsItem__GraphicsItemChange(28)
	QGraphicsItem__ItemRotationHasChanged             = QGraphicsItem__GraphicsItemChange(29)
	QGraphicsItem__ItemScaleChange                    = QGraphicsItem__GraphicsItemChange(30)
	QGraphicsItem__ItemScaleHasChanged                = QGraphicsItem__GraphicsItemChange(31)
	QGraphicsItem__ItemTransformOriginPointChange     = QGraphicsItem__GraphicsItemChange(32)
	QGraphicsItem__ItemTransformOriginPointHasChanged = QGraphicsItem__GraphicsItemChange(33)
)

//QGraphicsItem::GraphicsItemFlag
type QGraphicsItem__GraphicsItemFlag int64

const (
	QGraphicsItem__ItemIsMovable                        = QGraphicsItem__GraphicsItemFlag(0x1)
	QGraphicsItem__ItemIsSelectable                     = QGraphicsItem__GraphicsItemFlag(0x2)
	QGraphicsItem__ItemIsFocusable                      = QGraphicsItem__GraphicsItemFlag(0x4)
	QGraphicsItem__ItemClipsToShape                     = QGraphicsItem__GraphicsItemFlag(0x8)
	QGraphicsItem__ItemClipsChildrenToShape             = QGraphicsItem__GraphicsItemFlag(0x10)
	QGraphicsItem__ItemIgnoresTransformations           = QGraphicsItem__GraphicsItemFlag(0x20)
	QGraphicsItem__ItemIgnoresParentOpacity             = QGraphicsItem__GraphicsItemFlag(0x40)
	QGraphicsItem__ItemDoesntPropagateOpacityToChildren = QGraphicsItem__GraphicsItemFlag(0x80)
	QGraphicsItem__ItemStacksBehindParent               = QGraphicsItem__GraphicsItemFlag(0x100)
	QGraphicsItem__ItemUsesExtendedStyleOption          = QGraphicsItem__GraphicsItemFlag(0x200)
	QGraphicsItem__ItemHasNoContents                    = QGraphicsItem__GraphicsItemFlag(0x400)
	QGraphicsItem__ItemSendsGeometryChanges             = QGraphicsItem__GraphicsItemFlag(0x800)
	QGraphicsItem__ItemAcceptsInputMethod               = QGraphicsItem__GraphicsItemFlag(0x1000)
	QGraphicsItem__ItemNegativeZStacksBehindParent      = QGraphicsItem__GraphicsItemFlag(0x2000)
	QGraphicsItem__ItemIsPanel                          = QGraphicsItem__GraphicsItemFlag(0x4000)
	QGraphicsItem__ItemIsFocusScope                     = QGraphicsItem__GraphicsItemFlag(0x8000)
	QGraphicsItem__ItemSendsScenePositionChanges        = QGraphicsItem__GraphicsItemFlag(0x10000)
	QGraphicsItem__ItemStopsClickFocusPropagation       = QGraphicsItem__GraphicsItemFlag(0x20000)
	QGraphicsItem__ItemStopsFocusHandling               = QGraphicsItem__GraphicsItemFlag(0x40000)
	QGraphicsItem__ItemContainsChildrenInShape          = QGraphicsItem__GraphicsItemFlag(0x80000)
)

//QGraphicsItem::PanelModality
type QGraphicsItem__PanelModality int64

const (
	QGraphicsItem__NonModal   = QGraphicsItem__PanelModality(0)
	QGraphicsItem__PanelModal = QGraphicsItem__PanelModality(1)
	QGraphicsItem__SceneModal = QGraphicsItem__PanelModality(2)
)
