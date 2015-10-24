#include "qtableview.h"
#include <QUrl>
#include <QMetaObject>
#include <QPoint>
#include <QItemSelectionModel>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QHeaderView>
#include <QItemSelection>
#include <QTableView>
#include "_cgo_export.h"

class MyQTableView: public QTableView {
public:
};

int QTableView_GridStyle(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->gridStyle();
}

int QTableView_IsCornerButtonEnabled(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->isCornerButtonEnabled();
}

int QTableView_IsSortingEnabled(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->isSortingEnabled();
}

void QTableView_SetCornerButtonEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTableView*>(ptr)->setCornerButtonEnabled(enable != 0);
}

void QTableView_SetGridStyle(QtObjectPtr ptr, int style){
	static_cast<QTableView*>(ptr)->setGridStyle(static_cast<Qt::PenStyle>(style));
}

void QTableView_SetShowGrid(QtObjectPtr ptr, int show){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "setShowGrid", Q_ARG(bool, show != 0));
}

void QTableView_SetSpan(QtObjectPtr ptr, int row, int column, int rowSpanCount, int columnSpanCount){
	static_cast<QTableView*>(ptr)->setSpan(row, column, rowSpanCount, columnSpanCount);
}

void QTableView_SetWordWrap(QtObjectPtr ptr, int on){
	static_cast<QTableView*>(ptr)->setWordWrap(on != 0);
}

int QTableView_ShowGrid(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->showGrid();
}

int QTableView_WordWrap(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->wordWrap();
}

void QTableView_ClearSpans(QtObjectPtr ptr){
	static_cast<QTableView*>(ptr)->clearSpans();
}

int QTableView_ColumnAt(QtObjectPtr ptr, int x){
	return static_cast<QTableView*>(ptr)->columnAt(x);
}

int QTableView_ColumnSpan(QtObjectPtr ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->columnSpan(row, column);
}

int QTableView_ColumnViewportPosition(QtObjectPtr ptr, int column){
	return static_cast<QTableView*>(ptr)->columnViewportPosition(column);
}

int QTableView_ColumnWidth(QtObjectPtr ptr, int column){
	return static_cast<QTableView*>(ptr)->columnWidth(column);
}

void QTableView_HideColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideColumn", Q_ARG(int, column));
}

void QTableView_HideRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "hideRow", Q_ARG(int, row));
}

QtObjectPtr QTableView_HorizontalHeader(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->horizontalHeader();
}

QtObjectPtr QTableView_IndexAt(QtObjectPtr ptr, QtObjectPtr pos){
	return static_cast<QTableView*>(ptr)->indexAt(*static_cast<QPoint*>(pos)).internalPointer();
}

int QTableView_IsColumnHidden(QtObjectPtr ptr, int column){
	return static_cast<QTableView*>(ptr)->isColumnHidden(column);
}

int QTableView_IsRowHidden(QtObjectPtr ptr, int row){
	return static_cast<QTableView*>(ptr)->isRowHidden(row);
}

void QTableView_ResizeColumnToContents(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnToContents", Q_ARG(int, column));
}

void QTableView_ResizeColumnsToContents(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeColumnsToContents");
}

void QTableView_ResizeRowToContents(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowToContents", Q_ARG(int, row));
}

void QTableView_ResizeRowsToContents(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "resizeRowsToContents");
}

int QTableView_RowAt(QtObjectPtr ptr, int y){
	return static_cast<QTableView*>(ptr)->rowAt(y);
}

int QTableView_RowHeight(QtObjectPtr ptr, int row){
	return static_cast<QTableView*>(ptr)->rowHeight(row);
}

int QTableView_RowSpan(QtObjectPtr ptr, int row, int column){
	return static_cast<QTableView*>(ptr)->rowSpan(row, column);
}

int QTableView_RowViewportPosition(QtObjectPtr ptr, int row){
	return static_cast<QTableView*>(ptr)->rowViewportPosition(row);
}

void QTableView_SelectColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectColumn", Q_ARG(int, column));
}

void QTableView_SelectRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "selectRow", Q_ARG(int, row));
}

void QTableView_SetColumnHidden(QtObjectPtr ptr, int column, int hide){
	static_cast<QTableView*>(ptr)->setColumnHidden(column, hide != 0);
}

void QTableView_SetColumnWidth(QtObjectPtr ptr, int column, int width){
	static_cast<QTableView*>(ptr)->setColumnWidth(column, width);
}

void QTableView_SetHorizontalHeader(QtObjectPtr ptr, QtObjectPtr header){
	static_cast<QTableView*>(ptr)->setHorizontalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QTableView*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QTableView_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QTableView*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

void QTableView_SetRowHeight(QtObjectPtr ptr, int row, int height){
	static_cast<QTableView*>(ptr)->setRowHeight(row, height);
}

void QTableView_SetRowHidden(QtObjectPtr ptr, int row, int hide){
	static_cast<QTableView*>(ptr)->setRowHidden(row, hide != 0);
}

void QTableView_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel){
	static_cast<QTableView*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QTableView_SetSortingEnabled(QtObjectPtr ptr, int enable){
	static_cast<QTableView*>(ptr)->setSortingEnabled(enable != 0);
}

void QTableView_SetVerticalHeader(QtObjectPtr ptr, QtObjectPtr header){
	static_cast<QTableView*>(ptr)->setVerticalHeader(static_cast<QHeaderView*>(header));
}

void QTableView_ShowColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showColumn", Q_ARG(int, column));
}

void QTableView_ShowRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableView*>(ptr), "showRow", Q_ARG(int, row));
}

void QTableView_SortByColumn(QtObjectPtr ptr, int column, int order){
	static_cast<QTableView*>(ptr)->sortByColumn(column, static_cast<Qt::SortOrder>(order));
}

QtObjectPtr QTableView_VerticalHeader(QtObjectPtr ptr){
	return static_cast<QTableView*>(ptr)->verticalHeader();
}

void QTableView_DestroyQTableView(QtObjectPtr ptr){
	static_cast<QTableView*>(ptr)->~QTableView();
}

