#include "qidentityproxymodel.h"
#include <QObject>
#include <QMimeData>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIdentityProxyModel>
#include "_cgo_export.h"

class MyQIdentityProxyModel: public QIdentityProxyModel {
public:
};

QtObjectPtr QIdentityProxyModel_NewQIdentityProxyModel(QtObjectPtr parent){
	return new QIdentityProxyModel(static_cast<QObject*>(parent));
}

int QIdentityProxyModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char* QIdentityProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QIdentityProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

QtObjectPtr QIdentityProxyModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QIdentityProxyModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

QtObjectPtr QIdentityProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex){
	return static_cast<QIdentityProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

QtObjectPtr QIdentityProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex){
	return static_cast<QIdentityProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

QtObjectPtr QIdentityProxyModel_Parent(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QIdentityProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QIdentityProxyModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QIdentityProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QIdentityProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr newSourceModel){
	static_cast<QIdentityProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(newSourceModel));
}

QtObjectPtr QIdentityProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QIdentityProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QIdentityProxyModel_DestroyQIdentityProxyModel(QtObjectPtr ptr){
	static_cast<QIdentityProxyModel*>(ptr)->~QIdentityProxyModel();
}

