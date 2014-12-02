#include "qtableview.h"
#include <QTableView>
#include "cgoexport.h"

//MyQTableView
class MyQTableView: public QTableView {
Q_OBJECT
public:

signals:
void Slot_HideColumn(int column);
void Slot_HideRow(int row);
void Slot_ResizeColumnToContents(int column);
void Slot_ResizeColumnsToContents();
void Slot_ResizeRowToContents(int row);
void Slot_ResizeRowsToContents();
void Slot_SelectColumn(int column);
void Slot_SelectRow(int row);
void Slot_SetShowGrid(bool show);
void Slot_ShowColumn(int column);
void Slot_ShowRow(int row);

};
#include "qtableview.moc"


//Public Functions
QtObjectPtr QTableView_New_QWidget(QtObjectPtr parent)
{
	return new QTableView(((QWidget*)(parent)));
}

void QTableView_Destroy(QtObjectPtr ptr)
{
	((QTableView*)(ptr))->~QTableView();
}

void QTableView_ClearSpans(QtObjectPtr ptr)
{
	((QTableView*)(ptr))->clearSpans();
}

int QTableView_ColumnAt_Int(QtObjectPtr ptr, int x)
{
	return ((QTableView*)(ptr))->columnAt(x);
}

int QTableView_ColumnSpan_Int_Int(QtObjectPtr ptr, int row, int column)
{
	return ((QTableView*)(ptr))->columnSpan(row, column);
}

int QTableView_ColumnViewportPosition_Int(QtObjectPtr ptr, int column)
{
	return ((QTableView*)(ptr))->columnViewportPosition(column);
}

int QTableView_ColumnWidth_Int(QtObjectPtr ptr, int column)
{
	return ((QTableView*)(ptr))->columnWidth(column);
}

int QTableView_GridStyle(QtObjectPtr ptr)
{
	return ((QTableView*)(ptr))->gridStyle();
}

int QTableView_IsColumnHidden_Int(QtObjectPtr ptr, int column)
{
	return ((QTableView*)(ptr))->isColumnHidden(column);
}

int QTableView_IsCornerButtonEnabled(QtObjectPtr ptr)
{
	return ((QTableView*)(ptr))->isCornerButtonEnabled();
}

int QTableView_IsRowHidden_Int(QtObjectPtr ptr, int row)
{
	return ((QTableView*)(ptr))->isRowHidden(row);
}

int QTableView_IsSortingEnabled(QtObjectPtr ptr)
{
	return ((QTableView*)(ptr))->isSortingEnabled();
}

int QTableView_RowAt_Int(QtObjectPtr ptr, int y)
{
	return ((QTableView*)(ptr))->rowAt(y);
}

int QTableView_RowHeight_Int(QtObjectPtr ptr, int row)
{
	return ((QTableView*)(ptr))->rowHeight(row);
}

int QTableView_RowSpan_Int_Int(QtObjectPtr ptr, int row, int column)
{
	return ((QTableView*)(ptr))->rowSpan(row, column);
}

int QTableView_RowViewportPosition_Int(QtObjectPtr ptr, int row)
{
	return ((QTableView*)(ptr))->rowViewportPosition(row);
}

void QTableView_SetColumnHidden_Int_Bool(QtObjectPtr ptr, int column, int hide)
{
	((QTableView*)(ptr))->setColumnHidden(column, hide != 0);
}

void QTableView_SetColumnWidth_Int_Int(QtObjectPtr ptr, int column, int width)
{
	((QTableView*)(ptr))->setColumnWidth(column, width);
}

void QTableView_SetCornerButtonEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QTableView*)(ptr))->setCornerButtonEnabled(enable != 0);
}

void QTableView_SetGridStyle_PenStyle(QtObjectPtr ptr, int style)
{
	((QTableView*)(ptr))->setGridStyle(((Qt::PenStyle)(style)));
}

void QTableView_SetRowHeight_Int_Int(QtObjectPtr ptr, int row, int height)
{
	((QTableView*)(ptr))->setRowHeight(row, height);
}

void QTableView_SetRowHidden_Int_Bool(QtObjectPtr ptr, int row, int hide)
{
	((QTableView*)(ptr))->setRowHidden(row, hide != 0);
}

void QTableView_SetSortingEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QTableView*)(ptr))->setSortingEnabled(enable != 0);
}

void QTableView_SetSpan_Int_Int_Int_Int(QtObjectPtr ptr, int row, int column, int rowSpanCount, int columnSpanCount)
{
	((QTableView*)(ptr))->setSpan(row, column, rowSpanCount, columnSpanCount);
}

void QTableView_SetWordWrap_Bool(QtObjectPtr ptr, int on)
{
	((QTableView*)(ptr))->setWordWrap(on != 0);
}

int QTableView_ShowGrid(QtObjectPtr ptr)
{
	return ((QTableView*)(ptr))->showGrid();
}

void QTableView_SortByColumn_Int_SortOrder(QtObjectPtr ptr, int column, int order)
{
	((QTableView*)(ptr))->sortByColumn(column, ((Qt::SortOrder)(order)));
}

int QTableView_WordWrap(QtObjectPtr ptr)
{
	return ((QTableView*)(ptr))->wordWrap();
}

//Public Slots
void QTableView_ConnectSlotHideColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_HideColumn, ((QTableView*)(ptr)), &QTableView::hideColumn, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotHideColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_HideColumn, ((QTableView*)(ptr)), &QTableView::hideColumn);
}

void QTableView_HideColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTableView*)(ptr))->Slot_HideColumn(column);
}

void QTableView_ConnectSlotHideRow(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_HideRow, ((QTableView*)(ptr)), &QTableView::hideRow, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotHideRow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_HideRow, ((QTableView*)(ptr)), &QTableView::hideRow);
}

void QTableView_HideRow_Int(QtObjectPtr ptr, int row)
{
	((MyQTableView*)(ptr))->Slot_HideRow(row);
}

void QTableView_ConnectSlotResizeColumnToContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeColumnToContents, ((QTableView*)(ptr)), &QTableView::resizeColumnToContents, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotResizeColumnToContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeColumnToContents, ((QTableView*)(ptr)), &QTableView::resizeColumnToContents);
}

void QTableView_ResizeColumnToContents_Int(QtObjectPtr ptr, int column)
{
	((MyQTableView*)(ptr))->Slot_ResizeColumnToContents(column);
}

void QTableView_ConnectSlotResizeColumnsToContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeColumnsToContents, ((QTableView*)(ptr)), &QTableView::resizeColumnsToContents, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotResizeColumnsToContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeColumnsToContents, ((QTableView*)(ptr)), &QTableView::resizeColumnsToContents);
}

void QTableView_ResizeColumnsToContents(QtObjectPtr ptr)
{
	((MyQTableView*)(ptr))->Slot_ResizeColumnsToContents();
}

void QTableView_ConnectSlotResizeRowToContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeRowToContents, ((QTableView*)(ptr)), &QTableView::resizeRowToContents, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotResizeRowToContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeRowToContents, ((QTableView*)(ptr)), &QTableView::resizeRowToContents);
}

void QTableView_ResizeRowToContents_Int(QtObjectPtr ptr, int row)
{
	((MyQTableView*)(ptr))->Slot_ResizeRowToContents(row);
}

void QTableView_ConnectSlotResizeRowsToContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeRowsToContents, ((QTableView*)(ptr)), &QTableView::resizeRowsToContents, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotResizeRowsToContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ResizeRowsToContents, ((QTableView*)(ptr)), &QTableView::resizeRowsToContents);
}

void QTableView_ResizeRowsToContents(QtObjectPtr ptr)
{
	((MyQTableView*)(ptr))->Slot_ResizeRowsToContents();
}

void QTableView_ConnectSlotSelectColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SelectColumn, ((QTableView*)(ptr)), &QTableView::selectColumn, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotSelectColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SelectColumn, ((QTableView*)(ptr)), &QTableView::selectColumn);
}

void QTableView_SelectColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTableView*)(ptr))->Slot_SelectColumn(column);
}

void QTableView_ConnectSlotSelectRow(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SelectRow, ((QTableView*)(ptr)), &QTableView::selectRow, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotSelectRow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SelectRow, ((QTableView*)(ptr)), &QTableView::selectRow);
}

void QTableView_SelectRow_Int(QtObjectPtr ptr, int row)
{
	((MyQTableView*)(ptr))->Slot_SelectRow(row);
}

void QTableView_ConnectSlotSetShowGrid(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SetShowGrid, ((QTableView*)(ptr)), &QTableView::setShowGrid, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotSetShowGrid(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_SetShowGrid, ((QTableView*)(ptr)), &QTableView::setShowGrid);
}

void QTableView_SetShowGrid_Bool(QtObjectPtr ptr, int show)
{
	((MyQTableView*)(ptr))->Slot_SetShowGrid(show != 0);
}

void QTableView_ConnectSlotShowColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ShowColumn, ((QTableView*)(ptr)), &QTableView::showColumn, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotShowColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ShowColumn, ((QTableView*)(ptr)), &QTableView::showColumn);
}

void QTableView_ShowColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTableView*)(ptr))->Slot_ShowColumn(column);
}

void QTableView_ConnectSlotShowRow(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ShowRow, ((QTableView*)(ptr)), &QTableView::showRow, Qt::QueuedConnection);
}

void QTableView_DisconnectSlotShowRow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableView*)(ptr)), &MyQTableView::Slot_ShowRow, ((QTableView*)(ptr)), &QTableView::showRow);
}

void QTableView_ShowRow_Int(QtObjectPtr ptr, int row)
{
	((MyQTableView*)(ptr))->Slot_ShowRow(row);
}

