#include "qabstractlistmodel.h"
#include <QList>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QList>
#include <QString>
#include <QVariant>
#include <QAbstractListModel>
#include "_cgo_export.h"

class MyQAbstractListModel: public QAbstractListModel {
public:
};

void* QAbstractListModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QAbstractListModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QAbstractListModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractListModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractListModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractListModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

void* QAbstractListModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QAbstractListModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractListModel_DestroyQAbstractListModel(void* ptr){
	static_cast<QAbstractListModel*>(ptr)->~QAbstractListModel();
}

