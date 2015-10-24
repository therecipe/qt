#include "qstandarditemmodel.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStandardItem>
#include <QObject>
#include <QMimeData>
#include <QString>
#include <QStandardItemModel>
#include "_cgo_export.h"

class MyQStandardItemModel: public QStandardItemModel {
public:
void Signal_ItemChanged(QStandardItem * item){callbackQStandardItemModelItemChanged(this->objectName().toUtf8().data(), item);};
};

void QStandardItemModel_SetSortRole(QtObjectPtr ptr, int role){
	static_cast<QStandardItemModel*>(ptr)->setSortRole(role);
}

int QStandardItemModel_SortRole(QtObjectPtr ptr){
	return static_cast<QStandardItemModel*>(ptr)->sortRole();
}

QtObjectPtr QStandardItemModel_NewQStandardItemModel(QtObjectPtr parent){
	return new QStandardItemModel(static_cast<QObject*>(parent));
}

QtObjectPtr QStandardItemModel_NewQStandardItemModel2(int rows, int columns, QtObjectPtr parent){
	return new QStandardItemModel(rows, columns, static_cast<QObject*>(parent));
}

void QStandardItemModel_AppendRow2(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_Clear(QtObjectPtr ptr){
	static_cast<QStandardItemModel*>(ptr)->clear();
}

int QStandardItemModel_ColumnCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

char* QStandardItemModel_Data(QtObjectPtr ptr, QtObjectPtr index, int role){
	return static_cast<QStandardItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role).toString().toUtf8().data();
}

int QStandardItemModel_DropMimeData(QtObjectPtr ptr, QtObjectPtr data, int action, int row, int column, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_Flags(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QStandardItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStandardItemModel_HasChildren(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char* QStandardItemModel_HeaderData(QtObjectPtr ptr, int section, int orientation, int role){
	return static_cast<QStandardItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role).toString().toUtf8().data();
}

QtObjectPtr QStandardItemModel_HorizontalHeaderItem(QtObjectPtr ptr, int column){
	return static_cast<QStandardItemModel*>(ptr)->horizontalHeaderItem(column);
}

QtObjectPtr QStandardItemModel_Index(QtObjectPtr ptr, int row, int column, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

QtObjectPtr QStandardItemModel_IndexFromItem(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QStandardItemModel*>(ptr)->indexFromItem(static_cast<QStandardItem*>(item)).internalPointer();
}

int QStandardItemModel_InsertColumn2(QtObjectPtr ptr, int column, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertRow2(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_InsertRow3(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

int QStandardItemModel_InsertRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

QtObjectPtr QStandardItemModel_InvisibleRootItem(QtObjectPtr ptr){
	return static_cast<QStandardItemModel*>(ptr)->invisibleRootItem();
}

QtObjectPtr QStandardItemModel_Item(QtObjectPtr ptr, int row, int column){
	return static_cast<QStandardItemModel*>(ptr)->item(row, column);
}

void QStandardItemModel_ConnectItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));;
}

void QStandardItemModel_DisconnectItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));;
}

QtObjectPtr QStandardItemModel_ItemFromIndex(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QStandardItemModel*>(ptr)->itemFromIndex(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QStandardItemModel_ItemPrototype(QtObjectPtr ptr){
	return const_cast<QStandardItem*>(static_cast<QStandardItemModel*>(ptr)->itemPrototype());
}

char* QStandardItemModel_MimeTypes(QtObjectPtr ptr){
	return static_cast<QStandardItemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

QtObjectPtr QStandardItemModel_Parent(QtObjectPtr ptr, QtObjectPtr child){
	return static_cast<QStandardItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QStandardItemModel_RemoveColumns(QtObjectPtr ptr, int column, int count, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RemoveRows(QtObjectPtr ptr, int row, int count, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCount(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QStandardItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_SetColumnCount(QtObjectPtr ptr, int columns){
	static_cast<QStandardItemModel*>(ptr)->setColumnCount(columns);
}

int QStandardItemModel_SetData(QtObjectPtr ptr, QtObjectPtr index, char* value, int role){
	return static_cast<QStandardItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), QVariant(value), role);
}

int QStandardItemModel_SetHeaderData(QtObjectPtr ptr, int section, int orientation, char* value, int role){
	return static_cast<QStandardItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), QVariant(value), role);
}

void QStandardItemModel_SetHorizontalHeaderItem(QtObjectPtr ptr, int column, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderItem(column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetHorizontalHeaderLabels(QtObjectPtr ptr, char* labels){
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QStandardItemModel_SetItem2(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->setItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItem(QtObjectPtr ptr, int row, int column, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->setItem(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItemPrototype(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->setItemPrototype(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetRowCount(QtObjectPtr ptr, int rows){
	static_cast<QStandardItemModel*>(ptr)->setRowCount(rows);
}

void QStandardItemModel_SetVerticalHeaderItem(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetVerticalHeaderLabels(QtObjectPtr ptr, char* labels){
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

QtObjectPtr QStandardItemModel_Sibling(QtObjectPtr ptr, int row, int column, QtObjectPtr idx){
	return static_cast<QStandardItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QStandardItemModel_Sort(QtObjectPtr ptr, int column, int order){
	static_cast<QStandardItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QStandardItemModel_SupportedDropActions(QtObjectPtr ptr){
	return static_cast<QStandardItemModel*>(ptr)->supportedDropActions();
}

QtObjectPtr QStandardItemModel_TakeHorizontalHeaderItem(QtObjectPtr ptr, int column){
	return static_cast<QStandardItemModel*>(ptr)->takeHorizontalHeaderItem(column);
}

QtObjectPtr QStandardItemModel_TakeItem(QtObjectPtr ptr, int row, int column){
	return static_cast<QStandardItemModel*>(ptr)->takeItem(row, column);
}

QtObjectPtr QStandardItemModel_TakeVerticalHeaderItem(QtObjectPtr ptr, int row){
	return static_cast<QStandardItemModel*>(ptr)->takeVerticalHeaderItem(row);
}

QtObjectPtr QStandardItemModel_VerticalHeaderItem(QtObjectPtr ptr, int row){
	return static_cast<QStandardItemModel*>(ptr)->verticalHeaderItem(row);
}

void QStandardItemModel_DestroyQStandardItemModel(QtObjectPtr ptr){
	static_cast<QStandardItemModel*>(ptr)->~QStandardItemModel();
}

