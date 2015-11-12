#include "qaccessibletableinterface.h"
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTableModelChangeEvent>
#include <QAccessibleTableInterface>
#include "_cgo_export.h"

class MyQAccessibleTableInterface: public QAccessibleTableInterface {
public:
};

void* QAccessibleTableInterface_Caption(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->caption();
}

void* QAccessibleTableInterface_CellAt(void* ptr, int row, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->cellAt(row, column);
}

int QAccessibleTableInterface_ColumnCount(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->columnCount();
}

char* QAccessibleTableInterface_ColumnDescription(void* ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->columnDescription(column).toUtf8().data();
}

int QAccessibleTableInterface_IsColumnSelected(void* ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->isColumnSelected(column);
}

int QAccessibleTableInterface_IsRowSelected(void* ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->isRowSelected(row);
}

void QAccessibleTableInterface_ModelChange(void* ptr, void* event){
	static_cast<QAccessibleTableInterface*>(ptr)->modelChange(static_cast<QAccessibleTableModelChangeEvent*>(event));
}

int QAccessibleTableInterface_RowCount(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->rowCount();
}

char* QAccessibleTableInterface_RowDescription(void* ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->rowDescription(row).toUtf8().data();
}

int QAccessibleTableInterface_SelectColumn(void* ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectColumn(column);
}

int QAccessibleTableInterface_SelectRow(void* ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectRow(row);
}

int QAccessibleTableInterface_SelectedCellCount(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedCellCount();
}

int QAccessibleTableInterface_SelectedColumnCount(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedColumnCount();
}

int QAccessibleTableInterface_SelectedRowCount(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedRowCount();
}

void* QAccessibleTableInterface_Summary(void* ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->summary();
}

int QAccessibleTableInterface_UnselectColumn(void* ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectColumn(column);
}

int QAccessibleTableInterface_UnselectRow(void* ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectRow(row);
}

void QAccessibleTableInterface_DestroyQAccessibleTableInterface(void* ptr){
	static_cast<QAccessibleTableInterface*>(ptr)->~QAccessibleTableInterface();
}

