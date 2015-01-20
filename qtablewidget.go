package qt

//#include "qtablewidget.h"
import "C"

type qtablewidget struct {
	qtableview
}

type QTableWidget interface {
	QTableView
	CellWidget(row int, column int) QWidget
	ColumnCount() int
	CurrentColumn() int
	CurrentRow() int
	RemoveCellWidget(row int, column int)
	RowCount() int
	SetCellWidget(row int, column int, widget QWidget)
	SetColumnCount(columns int)
	SetCurrentCell(row int, column int)
	SetRowCount(rows int)
	SortItems(column int, order SortOrder)
	VisualColumn(logicalColumn int) int
	VisualRow(logicalRow int) int
	ConnectSlotClear()
	DisconnectSlotClear()
	SlotClear()
	ConnectSlotClearContents()
	DisconnectSlotClearContents()
	SlotClearContents()
	ConnectSlotInsertColumn()
	DisconnectSlotInsertColumn()
	SlotInsertColumn(column int)
	ConnectSlotInsertRow()
	DisconnectSlotInsertRow()
	SlotInsertRow(row int)
	ConnectSlotRemoveColumn()
	DisconnectSlotRemoveColumn()
	SlotRemoveColumn(column int)
	ConnectSlotRemoveRow()
	DisconnectSlotRemoveRow()
	SlotRemoveRow(row int)
	ConnectSignalCellActivated(f func())
	DisconnectSignalCellActivated()
	SignalCellActivated() func()
	ConnectSignalCellChanged(f func())
	DisconnectSignalCellChanged()
	SignalCellChanged() func()
	ConnectSignalCellClicked(f func())
	DisconnectSignalCellClicked()
	SignalCellClicked() func()
	ConnectSignalCellDoubleClicked(f func())
	DisconnectSignalCellDoubleClicked()
	SignalCellDoubleClicked() func()
	ConnectSignalCellEntered(f func())
	DisconnectSignalCellEntered()
	SignalCellEntered() func()
	ConnectSignalCellPressed(f func())
	DisconnectSignalCellPressed()
	SignalCellPressed() func()
	ConnectSignalCurrentCellChanged(f func())
	DisconnectSignalCurrentCellChanged()
	SignalCurrentCellChanged() func()
	ConnectSignalItemSelectionChanged(f func())
	DisconnectSignalItemSelectionChanged()
	SignalItemSelectionChanged() func()
}

func (p *qtablewidget) Pointer() (ptr C.QtObjectPtr) {
	return p.ptr
}

func (p *qtablewidget) SetPointer(ptr C.QtObjectPtr) {
	p.ptr = ptr
}

func NewQTableWidget(rows int, columns int, parent QWidget) QTableWidget {
	var parentPtr C.QtObjectPtr
	if parent != nil {
		parentPtr = parent.Pointer()
	}
	var qtablewidget = new(qtablewidget)
	qtablewidget.SetPointer(C.QTableWidget_New_Int_Int_QWidget(C.int(rows), C.int(columns), parentPtr))
	qtablewidget.SetObjectName("QTableWidget_" + randomIdentifier())
	return qtablewidget
}

func (p *qtablewidget) Destroy() {
	if p.Pointer() != nil {
		getSignal(p.ObjectName(), "destroyed")()
		C.QTableWidget_Destroy(p.Pointer())
		p.SetPointer(nil)
	}
}

func (p *qtablewidget) CellWidget(row int, column int) QWidget {
	if p.Pointer() == nil {
		return nil
	} else {
		var qwidget = new(qwidget)
		qwidget.SetPointer(C.QTableWidget_CellWidget_Int_Int(p.Pointer(), C.int(row), C.int(column)))
		if qwidget.ObjectName() == "" {
			qwidget.SetObjectName("QWidget_" + randomIdentifier())
		}
		return qwidget
	}
}

func (p *qtablewidget) ColumnCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_ColumnCount(p.Pointer()))
}

func (p *qtablewidget) CurrentColumn() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_CurrentColumn(p.Pointer()))
}

func (p *qtablewidget) CurrentRow() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_CurrentRow(p.Pointer()))
}

func (p *qtablewidget) RemoveCellWidget(row int, column int) {
	if p.Pointer() != nil {
		C.QTableWidget_RemoveCellWidget_Int_Int(p.Pointer(), C.int(row), C.int(column))
	}
}

func (p *qtablewidget) RowCount() int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_RowCount(p.Pointer()))
}

func (p *qtablewidget) SetCellWidget(row int, column int, widget QWidget) {
	if p.Pointer() != nil {
		var widgetPtr C.QtObjectPtr
		if widget != nil {
			widgetPtr = widget.Pointer()
		}
		C.QTableWidget_SetCellWidget_Int_Int_QWidget(p.Pointer(), C.int(row), C.int(column), widgetPtr)
	}
}

func (p *qtablewidget) SetColumnCount(columns int) {
	if p.Pointer() != nil {
		C.QTableWidget_SetColumnCount_Int(p.Pointer(), C.int(columns))
	}
}

func (p *qtablewidget) SetCurrentCell(row int, column int) {
	if p.Pointer() != nil {
		C.QTableWidget_SetCurrentCell_Int_Int(p.Pointer(), C.int(row), C.int(column))
	}
}

func (p *qtablewidget) SetRowCount(rows int) {
	if p.Pointer() != nil {
		C.QTableWidget_SetRowCount_Int(p.Pointer(), C.int(rows))
	}
}

func (p *qtablewidget) SortItems(column int, order SortOrder) {
	if p.Pointer() != nil {
		C.QTableWidget_SortItems_Int_SortOrder(p.Pointer(), C.int(column), C.int(order))
	}
}

func (p *qtablewidget) VisualColumn(logicalColumn int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_VisualColumn_Int(p.Pointer(), C.int(logicalColumn)))
}

func (p *qtablewidget) VisualRow(logicalRow int) int {
	if p.Pointer() == nil {
		return 0
	}
	return int(C.QTableWidget_VisualRow_Int(p.Pointer(), C.int(logicalRow)))
}

func (p *qtablewidget) ConnectSlotClear() {
	C.QTableWidget_ConnectSlotClear(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotClear() {
	C.QTableWidget_DisconnectSlotClear(p.Pointer())
}

func (p *qtablewidget) SlotClear() {
	if p.Pointer() != nil {
		C.QTableWidget_Clear(p.Pointer())
	}
}

func (p *qtablewidget) ConnectSlotClearContents() {
	C.QTableWidget_ConnectSlotClearContents(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotClearContents() {
	C.QTableWidget_DisconnectSlotClearContents(p.Pointer())
}

func (p *qtablewidget) SlotClearContents() {
	if p.Pointer() != nil {
		C.QTableWidget_ClearContents(p.Pointer())
	}
}

func (p *qtablewidget) ConnectSlotInsertColumn() {
	C.QTableWidget_ConnectSlotInsertColumn(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotInsertColumn() {
	C.QTableWidget_DisconnectSlotInsertColumn(p.Pointer())
}

func (p *qtablewidget) SlotInsertColumn(column int) {
	if p.Pointer() != nil {
		C.QTableWidget_InsertColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtablewidget) ConnectSlotInsertRow() {
	C.QTableWidget_ConnectSlotInsertRow(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotInsertRow() {
	C.QTableWidget_DisconnectSlotInsertRow(p.Pointer())
}

func (p *qtablewidget) SlotInsertRow(row int) {
	if p.Pointer() != nil {
		C.QTableWidget_InsertRow_Int(p.Pointer(), C.int(row))
	}
}

func (p *qtablewidget) ConnectSlotRemoveColumn() {
	C.QTableWidget_ConnectSlotRemoveColumn(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotRemoveColumn() {
	C.QTableWidget_DisconnectSlotRemoveColumn(p.Pointer())
}

func (p *qtablewidget) SlotRemoveColumn(column int) {
	if p.Pointer() != nil {
		C.QTableWidget_RemoveColumn_Int(p.Pointer(), C.int(column))
	}
}

func (p *qtablewidget) ConnectSlotRemoveRow() {
	C.QTableWidget_ConnectSlotRemoveRow(p.Pointer())
}

func (p *qtablewidget) DisconnectSlotRemoveRow() {
	C.QTableWidget_DisconnectSlotRemoveRow(p.Pointer())
}

func (p *qtablewidget) SlotRemoveRow(row int) {
	if p.Pointer() != nil {
		C.QTableWidget_RemoveRow_Int(p.Pointer(), C.int(row))
	}
}

func (p *qtablewidget) ConnectSignalCellActivated(f func()) {
	C.QTableWidget_ConnectSignalCellActivated(p.Pointer())
	connectSignal(p.ObjectName(), "cellActivated", f)
}

func (p *qtablewidget) DisconnectSignalCellActivated() {
	C.QTableWidget_DisconnectSignalCellActivated(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellActivated")
}

func (p *qtablewidget) SignalCellActivated() func() {
	return getSignal(p.ObjectName(), "cellActivated")
}

func (p *qtablewidget) ConnectSignalCellChanged(f func()) {
	C.QTableWidget_ConnectSignalCellChanged(p.Pointer())
	connectSignal(p.ObjectName(), "cellChanged", f)
}

func (p *qtablewidget) DisconnectSignalCellChanged() {
	C.QTableWidget_DisconnectSignalCellChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellChanged")
}

func (p *qtablewidget) SignalCellChanged() func() {
	return getSignal(p.ObjectName(), "cellChanged")
}

func (p *qtablewidget) ConnectSignalCellClicked(f func()) {
	C.QTableWidget_ConnectSignalCellClicked(p.Pointer())
	connectSignal(p.ObjectName(), "cellClicked", f)
}

func (p *qtablewidget) DisconnectSignalCellClicked() {
	C.QTableWidget_DisconnectSignalCellClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellClicked")
}

func (p *qtablewidget) SignalCellClicked() func() {
	return getSignal(p.ObjectName(), "cellClicked")
}

func (p *qtablewidget) ConnectSignalCellDoubleClicked(f func()) {
	C.QTableWidget_ConnectSignalCellDoubleClicked(p.Pointer())
	connectSignal(p.ObjectName(), "cellDoubleClicked", f)
}

func (p *qtablewidget) DisconnectSignalCellDoubleClicked() {
	C.QTableWidget_DisconnectSignalCellDoubleClicked(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellDoubleClicked")
}

func (p *qtablewidget) SignalCellDoubleClicked() func() {
	return getSignal(p.ObjectName(), "cellDoubleClicked")
}

func (p *qtablewidget) ConnectSignalCellEntered(f func()) {
	C.QTableWidget_ConnectSignalCellEntered(p.Pointer())
	connectSignal(p.ObjectName(), "cellEntered", f)
}

func (p *qtablewidget) DisconnectSignalCellEntered() {
	C.QTableWidget_DisconnectSignalCellEntered(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellEntered")
}

func (p *qtablewidget) SignalCellEntered() func() {
	return getSignal(p.ObjectName(), "cellEntered")
}

func (p *qtablewidget) ConnectSignalCellPressed(f func()) {
	C.QTableWidget_ConnectSignalCellPressed(p.Pointer())
	connectSignal(p.ObjectName(), "cellPressed", f)
}

func (p *qtablewidget) DisconnectSignalCellPressed() {
	C.QTableWidget_DisconnectSignalCellPressed(p.Pointer())
	disconnectSignal(p.ObjectName(), "cellPressed")
}

func (p *qtablewidget) SignalCellPressed() func() {
	return getSignal(p.ObjectName(), "cellPressed")
}

func (p *qtablewidget) ConnectSignalCurrentCellChanged(f func()) {
	C.QTableWidget_ConnectSignalCurrentCellChanged(p.Pointer())
	connectSignal(p.ObjectName(), "currentCellChanged", f)
}

func (p *qtablewidget) DisconnectSignalCurrentCellChanged() {
	C.QTableWidget_DisconnectSignalCurrentCellChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "currentCellChanged")
}

func (p *qtablewidget) SignalCurrentCellChanged() func() {
	return getSignal(p.ObjectName(), "currentCellChanged")
}

func (p *qtablewidget) ConnectSignalItemSelectionChanged(f func()) {
	C.QTableWidget_ConnectSignalItemSelectionChanged(p.Pointer())
	connectSignal(p.ObjectName(), "itemSelectionChanged", f)
}

func (p *qtablewidget) DisconnectSignalItemSelectionChanged() {
	C.QTableWidget_DisconnectSignalItemSelectionChanged(p.Pointer())
	disconnectSignal(p.ObjectName(), "itemSelectionChanged")
}

func (p *qtablewidget) SignalItemSelectionChanged() func() {
	return getSignal(p.ObjectName(), "itemSelectionChanged")
}

//export callbackQTableWidget
func callbackQTableWidget(ptr C.QtObjectPtr, msg *C.char) {
	var qtablewidget = new(qtablewidget)
	qtablewidget.SetPointer(ptr)
	getSignal(qtablewidget.ObjectName(), C.GoString(msg))()
}
