#include "qabstractitemmodel.h"
#include <QMimeData>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QAbstractItemModel>
#include "_cgo_export.h"

class MyQAbstractItemModel: public QAbstractItemModel {
public:
void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelColumnsAboutToBeInserted(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn){callbackQAbstractItemModelColumnsAboutToBeMoved(this->objectName().toUtf8().data(), sourceParent.internalPointer(), sourceStart, sourceEnd, destinationParent.internalPointer(), destinationColumn);};
void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelColumnsAboutToBeRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelColumnsInserted(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column){callbackQAbstractItemModelColumnsMoved(this->objectName().toUtf8().data(), parent.internalPointer(), start, end, destination.internalPointer(), column);};
void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelColumnsRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last){callbackQAbstractItemModelHeaderDataChanged(this->objectName().toUtf8().data(), orientation, first, last);};
void Signal_ModelAboutToBeReset(){callbackQAbstractItemModelModelAboutToBeReset(this->objectName().toUtf8().data());};
void Signal_ModelReset(){callbackQAbstractItemModelModelReset(this->objectName().toUtf8().data());};
void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end){callbackQAbstractItemModelRowsAboutToBeInserted(this->objectName().toUtf8().data(), parent.internalPointer(), start, end);};
void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow){callbackQAbstractItemModelRowsAboutToBeMoved(this->objectName().toUtf8().data(), sourceParent.internalPointer(), sourceStart, sourceEnd, destinationParent.internalPointer(), destinationRow);};
void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelRowsAboutToBeRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_RowsInserted(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelRowsInserted(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row){callbackQAbstractItemModelRowsMoved(this->objectName().toUtf8().data(), parent.internalPointer(), start, end, destination.internalPointer(), row);};
void Signal_RowsRemoved(const QModelIndex & parent, int first, int last){callbackQAbstractItemModelRowsRemoved(this->objectName().toUtf8().data(), parent.internalPointer(), first, last);};
};

QtObjectPtr QAbstractItemModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr index){
	return static_cast<QAbstractItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(index)).internalPointer();
}

QtObjectPtr QAbstractItemModel_Buddy(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractItemModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QAbstractItemModel_CanDropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_CanFetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectColumnsAboutToBeInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted));;
}

void QAbstractItemModel_ConnectColumnsAboutToBeMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved));;
}

void QAbstractItemModel_ConnectColumnsAboutToBeRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved));;
}

void QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved));;
}

void QAbstractItemModel_ConnectColumnsInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted));;
}

void QAbstractItemModel_DisconnectColumnsInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted));;
}

void QAbstractItemModel_ConnectColumnsMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved));;
}

void QAbstractItemModel_DisconnectColumnsMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved));;
}

void QAbstractItemModel_ConnectColumnsRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved));;
}

void QAbstractItemModel_DisconnectColumnsRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved));;
}

char* QAbstractItemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QAbstractItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

int QAbstractItemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_FetchMore(QtObjectPtr ptr, QtObjectPtr parent){
	static_cast<QAbstractItemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QAbstractItemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_HasIndex(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->hasIndex(row, column, *static_cast<QModelIndex*>(parent));
}

char* QAbstractItemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QAbstractItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

void QAbstractItemModel_ConnectHeaderDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged));;
}

void QAbstractItemModel_DisconnectHeaderDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged));;
}

QtObjectPtr QAbstractItemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

int QAbstractItemModel_InsertColumn(QtObjectPtr ptr, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertRow(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char* QAbstractItemModel_MimeTypes(QtObjectPtr ptr){
	return static_cast<QAbstractItemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void QAbstractItemModel_ConnectModelAboutToBeReset(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset));;
}

void QAbstractItemModel_DisconnectModelAboutToBeReset(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset));;
}

void QAbstractItemModel_ConnectModelReset(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset));;
}

void QAbstractItemModel_DisconnectModelReset(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset));;
}

int QAbstractItemModel_MoveColumn(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceColumn, QtObjectPtr destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveColumn(*static_cast<QModelIndex*>(sourceParent), sourceColumn, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveColumns(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceColumn, int count, QtObjectPtr destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveRow(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceRow, QtObjectPtr destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveRow(*static_cast<QModelIndex*>(sourceParent), sourceRow, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QAbstractItemModel_MoveRows(QtObjectPtr ptr, QtObjectPtr sourceParent, int sourceRow, int count, QtObjectPtr destinationParent, int destinationChild){
	return static_cast<QAbstractItemModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

QtObjectPtr QAbstractItemModel_Parent(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QAbstractItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)).internalPointer();
}

int QAbstractItemModel_RemoveColumn(QtObjectPtr ptr, int column, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeColumn(column, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveRow(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeRow(row, *static_cast<QModelIndex*>(parent));
}

int QAbstractItemModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_Revert(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "revert");
}

int QAbstractItemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QAbstractItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectRowsAboutToBeInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted));;
}

void QAbstractItemModel_ConnectRowsAboutToBeMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved));;
}

void QAbstractItemModel_ConnectRowsAboutToBeRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved));;
}

void QAbstractItemModel_DisconnectRowsAboutToBeRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved));;
}

void QAbstractItemModel_ConnectRowsInserted(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted));;
}

void QAbstractItemModel_DisconnectRowsInserted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted));;
}

void QAbstractItemModel_ConnectRowsMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved));;
}

void QAbstractItemModel_DisconnectRowsMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved));;
}

void QAbstractItemModel_ConnectRowsRemoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved));;
}

void QAbstractItemModel_DisconnectRowsRemoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved));;
}

int QAbstractItemModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QAbstractItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

int QAbstractItemModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role){
	return static_cast<QAbstractItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), QVariant(value), role);
}

void QAbstractItemModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QAbstractItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QAbstractItemModel_Submit(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "submit");
}

int QAbstractItemModel_SupportedDragActions(QtObjectPtr ptr){
	return static_cast<QAbstractItemModel*>(ptr)->supportedDragActions();
}

int QAbstractItemModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QAbstractItemModel*>(ptr)->supportedDropActions();
}

void QAbstractItemModel_DestroyQAbstractItemModel(QtObjectPtr ptr){
	static_cast<QAbstractItemModel*>(ptr)->~QAbstractItemModel();
}

