#include "qaccessibletablecellinterface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccessible>
#include <QString>
#include <QAccessibleTableCellInterface>
#include "_cgo_export.h"

class MyQAccessibleTableCellInterface: public QAccessibleTableCellInterface {
public:
};

int QAccessibleTableCellInterface_ColumnExtent(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnExtent();
}

int QAccessibleTableCellInterface_ColumnIndex(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnIndex();
}

int QAccessibleTableCellInterface_IsSelected(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->isSelected();
}

int QAccessibleTableCellInterface_RowExtent(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowExtent();
}

int QAccessibleTableCellInterface_RowIndex(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowIndex();
}

QtObjectPtr QAccessibleTableCellInterface_Table(QtObjectPtr ptr){
	return static_cast<QAccessibleTableCellInterface*>(ptr)->table();
}

void QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(QtObjectPtr ptr){
	static_cast<QAccessibleTableCellInterface*>(ptr)->~QAccessibleTableCellInterface();
}

