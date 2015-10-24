#include "qabstractlistmodel.h"
#include <QList>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QList>
#include <QMimeData>
#include <QAbstractListModel>
#include "_cgo_export.h"

class MyQAbstractListModel: public QAbstractListModel {
public:
};

QtObjectPtr QAbstractListModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractListModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QAbstractListModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractListModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractListModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QAbstractListModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QAbstractListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractListModel_DestroyQAbstractListModel(QtObjectPtr ptr){
	static_cast<QAbstractListModel*>(ptr)->~QAbstractListModel();
}

