#include "qidentityproxymodel.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMimeData>
#include <QAbstractItemModel>
#include <QIdentityProxyModel>
#include "_cgo_export.h"

class MyQIdentityProxyModel: public QIdentityProxyModel {
public:
};

void* QIdentityProxyModel_NewQIdentityProxyModel(void* parent){
	return new QIdentityProxyModel(static_cast<QObject*>(parent));
}

int QIdentityProxyModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void* QIdentityProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QIdentityProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QIdentityProxyModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QIdentityProxyModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void* QIdentityProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return static_cast<QIdentityProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

void* QIdentityProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return static_cast<QIdentityProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

void* QIdentityProxyModel_Parent(void* ptr, void* child){
	return static_cast<QIdentityProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QIdentityProxyModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QIdentityProxyModel_RowCount(void* ptr, void* parent){
	return static_cast<QIdentityProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QIdentityProxyModel_SetSourceModel(void* ptr, void* newSourceModel){
	static_cast<QIdentityProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(newSourceModel));
}

void* QIdentityProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QIdentityProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QIdentityProxyModel_DestroyQIdentityProxyModel(void* ptr){
	static_cast<QIdentityProxyModel*>(ptr)->~QIdentityProxyModel();
}

