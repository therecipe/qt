#include "qtableview.h"
#include <QString>
#include <QUrl>
#include <QHeaderView>
#include <QMetaObject>
#include <QItemSelectionModel>
#include <QVariant>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QPoint>
#include <QItemSelection>
#include <QTableView>
#include "_cgo_export.h"

class MyQTableView: public QTableView {
public:
};

int QTableView_GridStyle(void* ptr){
	return static_cast<QTableView*>(ptr)->gridStyle();
}

int QTableView_IsCornerButtonEnabled(void* ptr){
	return static_cast<QTableView*>(ptr)->isCornerButtonEnabled();
}

int QTableView_IsSortingEnabled(void* ptr){
	return static_cast<QTableView*>(ptr)->isSortingEnabled();
}

void QTableView_SetCornerButtonEnabled(void* ptr, int enable){
	static_cast<QTableView*>(ptr)->setCornerButtonEnabled(enable != 0);
}

void QTableView_SetGridStyle(void* ptr, int style){
	static_cast<QTableView*>(ptr)->setGridStyle(static_cast<Qt::PenStyle>(style));
}

void QTableView_SetShowGrid(void* ptr, int show){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "setShowGrid", Q_ARG(bool, show != 0));
}

void QTableView_SetSpan(void* ptr, int row, int column, int rowSpanCount, int columnSpanCount){
	static_cast<QTableView*>(ptr)->setSpan(row, column, rowSpanCount, columnSpanCount);
}

void QTableView_SetWordWrap(void* ptr, int on){
	static_cast<QTableView*>(ptr)->setWordWrap(on != 0);
}

int QTableView_ShowGrid(void* ptr){
	return static_cast<QTableView*>(ptr)->showGrid();
}

int QTableView_WordWrap(void* ptr){
	return static_cast<QTableView*>(ptr)->wordWrap();
}

void QTableView_ClearSpans(void* ptr){
	static_cast<QTableView*>(ptr)->clearSpans();
}

int QTableView_ColumnAt(void* ptr, int x){
	return static_cast<QTableView*>(ptr)->columnAt(x);
}

int QTableView_ColumnSpan(void* ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->columnSpan(row, column);
}

int QTableView_ColumnViewportPosition(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->columnViewportPosition(column);
}

int QTableView_ColumnWidth(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->columnWidth(column);
}

void QTableView_HideColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideColumn", Q_ARG(int, column));
}

void QTableView_HideRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideRow", Q_ARG(int, row));
}

void* QTableView_HorizontalHeader(void* ptr){
	return static_cast<QTableView*>(ptr)->horizontalHeader();
}

void* QTableView_IndexAt(void* ptr, void* pos){
	return static_cast<QTableView*>(ptr)->indexAt(*static_cast<QPoint*>(pos)).internalPointer();
}

int QTableView_IsColumnHidden(void* ptr, int column){
	return static_cast<QTableView*>(ptr)->isColumnHidden(column);
}

int QTableView_IsRowHidden(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->isRowHidden(row);
}

void QTableView_ResizeColumnToContents(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTableView_ResizeColumnsToContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnsToContents");
}

void QTableView_ResizeRowToContents(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowToContents", Q_ARG(int, row));
}

void QTableView_ResizeRowsToContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowsToContents");
}

int QTableView_RowAt(void* ptr, int y){
	return static_cast<QTableView*>(ptr)->rowAt(y);
}

int QTableView_RowHeight(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->rowHeight(row);
}

int QTableView_RowSpan(void* ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->rowSpan(row, column);
}

int QTableView_RowViewportPosition(void* ptr, int row){
	return static_cast<QTableView*>(ptr)->rowViewportPosition(row);
}

void QTableView_SelectColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectColumn", Q_ARG(int, column));
}

void QTableView_SelectRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectRow", Q_ARG(int, row));
}

void QTableView_SetColumnHidden(void* ptr, int column, int hide){
	static_cast<QTableView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTableView_SetColumnWidth(void* ptr, int column, int width){
	static_cast<QTableView*>(ptr)->setColumnWidth(column, width);
}

void QTableView_SetHorizontalHeader(void* ptr, void* header){
	static_cast<QTableView*>(ptr)->setHorizontalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_SetModel(void* ptr, void* model){
	static_cast<QTableView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTableView_SetRootIndex(void* ptr, void* index){
	static_cast<QTableView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTableView_SetRowHeight(void* ptr, int row, int height){
	static_cast<QTableView*>(ptr)->setRowHeight(row, height);
}

void QTableView_SetRowHidden(void* ptr, int row, int hide){
	static_cast<QTableView*>(ptr)->setRowHidden(row, hide != 0);
}

void QTableView_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<QTableView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTableView_SetSortingEnabled(void* ptr, int enable){
	static_cast<QTableView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTableView_SetVerticalHeader(void* ptr, void* header){
	static_cast<QTableView*>(ptr)->setVerticalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_ShowColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTableView_ShowRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showRow", Q_ARG(int, row));
}

void QTableView_SortByColumn(void* ptr, int column, int order){
	static_cast<QTableView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

void* QTableView_VerticalHeader(void* ptr){
	return static_cast<QTableView*>(ptr)->verticalHeader();
}

void QTableView_DestroyQTableView(void* ptr){
	static_cast<QTableView*>(ptr)->~QTableView();
}

