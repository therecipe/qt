#include "qabstracttablemodel.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QString>
#include <QAbstractTableModel>
#include "_cgo_export.h"

class MyQAbstractTableModel: public QAbstractTableModel {
public:
};

QtObjectPtr QAbstractTableModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QAbstractTableModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractTableModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractTableModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QAbstractTableModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QAbstractTableModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractTableModel_DestroyQAbstractTableModel(QtObjectPtr ptr){
	static_cast<QAbstractTableModel*>(ptr)->~QAbstractTableModel();
}

