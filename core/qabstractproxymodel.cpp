#include "qabstractproxymodel.h"
#include <QObject>
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QAbstractProxyModel>
#include "_cgo_export.h"

class MyQAbstractProxyModel: public QAbstractProxyModel {
public:
void Signal_SourceModelChanged(){callbackQAbstractProxyModelSourceModelChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QAbstractProxyModel_Buddy(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QAbstractProxyModel_CanDropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

char* QAbstractProxyModel_Data(QtObjectPtr ptr, QtObjectPtr proxyIndex, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(proxyIndex), role).toString().toUtf8().data();
}

int QAbstractProxyModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractProxyModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QAbstractProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QAbstractProxyModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char* QAbstractProxyModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

QtObjectPtr QAbstractProxyModel_MapFromSource(QtObjectPtr ptr, QtObjectPtr sourceIndex){
	return static_cast<QAbstractProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

QtObjectPtr QAbstractProxyModel_MapToSource(QtObjectPtr ptr, QtObjectPtr proxyIndex){
	return static_cast<QAbstractProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

char* QAbstractProxyModel_MimeTypes(QtObjectPtr ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void QAbstractProxyModel_Revert(QtObjectPtr ptr){
	static_cast<QAbstractProxyModel*>(ptr)->revert();
}

int QAbstractProxyModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

int QAbstractProxyModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), QVariant(value), role);
}

void QAbstractProxyModel_SetSourceModel(QtObjectPtr ptr, QtObjectPtr sourceModel){
	static_cast<QAbstractProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

QtObjectPtr QAbstractProxyModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QAbstractProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractProxyModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QAbstractProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

QtObjectPtr QAbstractProxyModel_SourceModel(QtObjectPtr ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->sourceModel();
}

void QAbstractProxyModel_ConnectSourceModelChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

void QAbstractProxyModel_DisconnectSourceModelChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

int QAbstractProxyModel_Submit(QtObjectPtr ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->submit();
}

int QAbstractProxyModel_SupportedDragActions(QtObjectPtr ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDragActions();
}

int QAbstractProxyModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDropActions();
}

void QAbstractProxyModel_DestroyQAbstractProxyModel(QtObjectPtr ptr){
	static_cast<QAbstractProxyModel*>(ptr)->~QAbstractProxyModel();
}

