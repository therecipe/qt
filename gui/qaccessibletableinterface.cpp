#include "qaccessibletableinterface.h"
#include <QUrl>
#include <QModelIndex>
#include <QAccessibleTableModelChangeEvent>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QAccessibleTableInterface>
#include "_cgo_export.h"

class MyQAccessibleTableInterface: public QAccessibleTableInterface {
public:
};

QtObjectPtr QAccessibleTableInterface_Caption(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->caption();
}

QtObjectPtr QAccessibleTableInterface_CellAt(QtObjectPtr ptr, int row, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->cellAt(row, column);
}

int QAccessibleTableInterface_ColumnCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->columnCount();
}

char* QAccessibleTableInterface_ColumnDescription(QtObjectPtr ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->columnDescription(column).toUtf8().data();
}

int QAccessibleTableInterface_IsColumnSelected(QtObjectPtr ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->isColumnSelected(column);
}

int QAccessibleTableInterface_IsRowSelected(QtObjectPtr ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->isRowSelected(row);
}

void QAccessibleTableInterface_ModelChange(QtObjectPtr ptr, QtObjectPtr event){
	static_cast<QAccessibleTableInterface*>(ptr)->modelChange(static_cast<QAccessibleTableModelChangeEvent*>(event));
}

int QAccessibleTableInterface_RowCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->rowCount();
}

char* QAccessibleTableInterface_RowDescription(QtObjectPtr ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->rowDescription(row).toUtf8().data();
}

int QAccessibleTableInterface_SelectColumn(QtObjectPtr ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectColumn(column);
}

int QAccessibleTableInterface_SelectRow(QtObjectPtr ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectRow(row);
}

int QAccessibleTableInterface_SelectedCellCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedCellCount();
}

int QAccessibleTableInterface_SelectedColumnCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedColumnCount();
}

int QAccessibleTableInterface_SelectedRowCount(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedRowCount();
}

QtObjectPtr QAccessibleTableInterface_Summary(QtObjectPtr ptr){
	return static_cast<QAccessibleTableInterface*>(ptr)->summary();
}

int QAccessibleTableInterface_UnselectColumn(QtObjectPtr ptr, int column){
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectColumn(column);
}

int QAccessibleTableInterface_UnselectRow(QtObjectPtr ptr, int row){
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectRow(row);
}

void QAccessibleTableInterface_DestroyQAccessibleTableInterface(QtObjectPtr ptr){
	static_cast<QAccessibleTableInterface*>(ptr)->~QAccessibleTableInterface();
}

