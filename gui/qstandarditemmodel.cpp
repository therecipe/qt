#include "qstandarditemmodel.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStandardItem>
#include <QMimeData>
#include <QStandardItemModel>
#include "_cgo_export.h"

class MyQStandardItemModel: public QStandardItemModel {
public:
void Signal_ItemChanged(QStandardItem * item){callbackQStandardItemModelItemChanged(this->objectName().toUtf8().data(), item);};
};

void QStandardItemModel_SetSortRole(void* ptr, int role){
	static_cast<QStandardItemModel*>(ptr)->setSortRole(role);
}

int QStandardItemModel_SortRole(void* ptr){
	return static_cast<QStandardItemModel*>(ptr)->sortRole();
}

void* QStandardItemModel_NewQStandardItemModel(void* parent){
	return new QStandardItemModel(static_cast<QObject*>(parent));
}

void* QStandardItemModel_NewQStandardItemModel2(int rows, int columns, void* parent){
	return new QStandardItemModel(rows, columns, static_cast<QObject*>(parent));
}

void QStandardItemModel_AppendRow2(void* ptr, void* item){
	static_cast<QStandardItemModel*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_Clear(void* ptr){
	static_cast<QStandardItemModel*>(ptr)->clear();
}

int QStandardItemModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QStandardItemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_Flags(void* ptr, void* index){
	return static_cast<QStandardItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStandardItemModel_HasChildren(void* ptr, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_HeaderData(void* ptr, int section, int orientation, int role){
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QStandardItemModel_HorizontalHeaderItem(void* ptr, int column){
	return static_cast<QStandardItemModel*>(ptr)->horizontalHeaderItem(column);
}

void* QStandardItemModel_Index(void* ptr, int row, int column, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)).internalPointer();
}

void* QStandardItemModel_IndexFromItem(void* ptr, void* item){
	return static_cast<QStandardItemModel*>(ptr)->indexFromItem(static_cast<QStandardItem*>(item)).internalPointer();
}

int QStandardItemModel_InsertColumn2(void* ptr, int column, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertRow2(void* ptr, int row, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_InsertRow3(void* ptr, int row, void* item){
	static_cast<QStandardItemModel*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

int QStandardItemModel_InsertRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_InvisibleRootItem(void* ptr){
	return static_cast<QStandardItemModel*>(ptr)->invisibleRootItem();
}

void* QStandardItemModel_Item(void* ptr, int row, int column){
	return static_cast<QStandardItemModel*>(ptr)->item(row, column);
}

void QStandardItemModel_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));;
}

void QStandardItemModel_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));;
}

void* QStandardItemModel_ItemFromIndex(void* ptr, void* index){
	return static_cast<QStandardItemModel*>(ptr)->itemFromIndex(*static_cast<QModelIndex*>(index));
}

void* QStandardItemModel_ItemPrototype(void* ptr){
	return const_cast<QStandardItem*>(static_cast<QStandardItemModel*>(ptr)->itemPrototype());
}

char* QStandardItemModel_MimeTypes(void* ptr){
	return static_cast<QStandardItemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

void* QStandardItemModel_Parent(void* ptr, void* child){
	return static_cast<QStandardItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)).internalPointer();
}

int QStandardItemModel_RemoveColumns(void* ptr, int column, int count, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RemoveRows(void* ptr, int row, int count, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCount(void* ptr, void* parent){
	return static_cast<QStandardItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_SetColumnCount(void* ptr, int columns){
	static_cast<QStandardItemModel*>(ptr)->setColumnCount(columns);
}

int QStandardItemModel_SetData(void* ptr, void* index, void* value, int role){
	return static_cast<QStandardItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QStandardItemModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role){
	return static_cast<QStandardItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QStandardItemModel_SetHorizontalHeaderItem(void* ptr, int column, void* item){
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderItem(column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetHorizontalHeaderLabels(void* ptr, char* labels){
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QStandardItemModel_SetItem2(void* ptr, int row, void* item){
	static_cast<QStandardItemModel*>(ptr)->setItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItem(void* ptr, int row, int column, void* item){
	static_cast<QStandardItemModel*>(ptr)->setItem(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItemPrototype(void* ptr, void* item){
	static_cast<QStandardItemModel*>(ptr)->setItemPrototype(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetRowCount(void* ptr, int rows){
	static_cast<QStandardItemModel*>(ptr)->setRowCount(rows);
}

void QStandardItemModel_SetVerticalHeaderItem(void* ptr, int row, void* item){
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetVerticalHeaderLabels(void* ptr, char* labels){
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void* QStandardItemModel_Sibling(void* ptr, int row, int column, void* idx){
	return static_cast<QStandardItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)).internalPointer();
}

void QStandardItemModel_Sort(void* ptr, int column, int order){
	static_cast<QStandardItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

int QStandardItemModel_SupportedDropActions(void* ptr){
	return static_cast<QStandardItemModel*>(ptr)->supportedDropActions();
}

void* QStandardItemModel_TakeHorizontalHeaderItem(void* ptr, int column){
	return static_cast<QStandardItemModel*>(ptr)->takeHorizontalHeaderItem(column);
}

void* QStandardItemModel_TakeItem(void* ptr, int row, int column){
	return static_cast<QStandardItemModel*>(ptr)->takeItem(row, column);
}

void* QStandardItemModel_TakeVerticalHeaderItem(void* ptr, int row){
	return static_cast<QStandardItemModel*>(ptr)->takeVerticalHeaderItem(row);
}

void* QStandardItemModel_VerticalHeaderItem(void* ptr, int row){
	return static_cast<QStandardItemModel*>(ptr)->verticalHeaderItem(row);
}

void QStandardItemModel_DestroyQStandardItemModel(void* ptr){
	static_cast<QStandardItemModel*>(ptr)->~QStandardItemModel();
}

