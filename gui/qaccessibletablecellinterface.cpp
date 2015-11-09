#include "qaccessibletablecellinterface.h"
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QVariant>
#include <QAccessibleTableCellInterface>
#include "_cgo_export.h"

class MyQAccessibleTableCellInterface: public QAccessibleTableCellInterface {
public:
};

int QAccessibleTableCellInterface_ColumnExtent(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnExtent();
}

int QAccessibleTableCellInterface_ColumnIndex(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnIndex();
}

int QAccessibleTableCellInterface_IsSelected(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->isSelected();
}

int QAccessibleTableCellInterface_RowExtent(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowExtent();
}

int QAccessibleTableCellInterface_RowIndex(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowIndex();
}

void* QAccessibleTableCellInterface_Table(void* ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->table();
}

void QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(void* ptr){
	static_cast<QAccessibleTableCellInterface*>(ptr)->~QAccessibleTableCellInterface();
}

