package help

//#include "help.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/widgets"
	"log"
	"unsafe"
)

type QHelpSearchQueryWidget struct {
	widgets.QWidget
}

type QHelpSearchQueryWidget_ITF interface {
	widgets.QWidget_ITF
	QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget
}

func PointerFromQHelpSearchQueryWidget(ptr QHelpSearchQueryWidget_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHelpSearchQueryWidget_PTR().Pointer()
	}
	return nil
}

func NewQHelpSearchQueryWidgetFromPointer(ptr unsafe.Pointer) *QHelpSearchQueryWidget {
	var n = new(QHelpSearchQueryWidget)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QHelpSearchQueryWidget_") {
		n.SetObjectName("QHelpSearchQueryWidget_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHelpSearchQueryWidget) QHelpSearchQueryWidget_PTR() *QHelpSearchQueryWidget {
	return ptr
}

func (ptr *QHelpSearchQueryWidget) IsCompactMode() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::isCompactMode")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHelpSearchQueryWidget_IsCompactMode(ptr.Pointer()) != 0
	}
	return false
}

func NewQHelpSearchQueryWidget(parent widgets.QWidget_ITF) *QHelpSearchQueryWidget {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::QHelpSearchQueryWidget")
		}
	}()

	return NewQHelpSearchQueryWidgetFromPointer(C.QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(widgets.PointerFromQWidget(parent)))
}

func (ptr *QHelpSearchQueryWidget) CollapseExtendedSearch() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::collapseExtendedSearch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_CollapseExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ExpandExtendedSearch() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::expandExtendedSearch")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ExpandExtendedSearch(ptr.Pointer())
	}
}

func (ptr *QHelpSearchQueryWidget) ConnectSearch(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::search")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_ConnectSearch(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "search", f)
	}
}

func (ptr *QHelpSearchQueryWidget) DisconnectSearch() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::search")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DisconnectSearch(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "search")
	}
}

//export callbackQHelpSearchQueryWidgetSearch
func callbackQHelpSearchQueryWidgetSearch(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::search")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "search").(func())()
}

func (ptr *QHelpSearchQueryWidget) DestroyQHelpSearchQueryWidget() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHelpSearchQueryWidget::~QHelpSearchQueryWidget")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
