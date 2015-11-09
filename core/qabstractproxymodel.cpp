#include "qabstractproxymodel.h"
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QObject>
#include <QAbstractProxyModel>
#include "_cgo_export.h"

class MyQAbstractProxyModel: public QAbstractProxyModel {
public:
void Signal_SourceModelChanged(){callbackQAbstractProxyModelSourceModelChanged(this->objectName().toUtf8().data());};
};

void* QAbstractProxyModel_Buddy(void* ptr, void* index){
	return static_cast<QAbstractProxyModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QAbstractProxyModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_CanFetchMore(void* ptr, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

void* QAbstractProxyModel_Data(void* ptr, void* proxyIndex, int role){
	return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(proxyIndex), role));
}

int QAbstractProxyModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractProxyModel_FetchMore(void* ptr, void* parent){
	static_cast<QAbstractProxyModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_Flags(void* ptr, void* index){
	return static_cast<QAbstractProxyModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QAbstractProxyModel_HasChildren(void* ptr, void* parent){
	return static_cast<QAbstractProxyModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QAbstractProxyModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QAbstractProxyModel_MapFromSource(void* ptr, void* sourceIndex){
	return static_cast<QAbstractProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)).internalPointer();
}

void* QAbstractProxyModel_MapToSource(void* ptr, void* proxyIndex){
	return static_cast<QAbstractProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)).internalPointer();
}

char* QAbstractProxyModel_MimeTypes(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void QAbstractProxyModel_Revert(void* ptr){
	static_cast<QAbstractProxyModel*>(ptr)->revert();
}

int QAbstractProxyModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QAbstractProxyModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QAbstractProxyModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QAbstractProxyModel_SetSourceModel(void* ptr, void* sourceModel){
	static_cast<QAbstractProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void* QAbstractProxyModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QAbstractProxyModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QAbstractProxyModel_Sort(void* ptr, int column, int order){
	static_cast<QAbstractProxyModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void* QAbstractProxyModel_SourceModel(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->sourceModel();
}

void QAbstractProxyModel_ConnectSourceModelChanged(void* ptr){
	QObject::connect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

void QAbstractProxyModel_DisconnectSourceModelChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractProxyModel*>(ptr), &QAbstractProxyModel::sourceModelChanged, static_cast<MyQAbstractProxyModel*>(ptr), static_cast<void (MyQAbstractProxyModel::*)()>(&MyQAbstractProxyModel::Signal_SourceModelChanged));;
}

int QAbstractProxyModel_Submit(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->submit();
}

int QAbstractProxyModel_SupportedDragActions(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDragActions();
}

int QAbstractProxyModel_SupportedDropActions(void* ptr){
	return static_cast<QAbstractProxyModel*>(ptr)->supportedDropActions();
}

void QAbstractProxyModel_DestroyQAbstractProxyModel(void* ptr){
	static_cast<QAbstractProxyModel*>(ptr)->~QAbstractProxyModel();
}

