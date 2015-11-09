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

void* QAbstractTableModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QAbstractTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QAbstractTableModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractTableModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractTableModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractTableModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QAbstractTableModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QAbstractTableModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractTableModel_DestroyQAbstractTableModel(void* ptr){
	static_cast<QAbstractTableModel*>(ptr)->~QAbstractTableModel();
}

