#include "qmodelindex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QModelIndex>
#include "_cgo_export.h"

class MyQModelIndex: public QModelIndex {
public:
};

void* QModelIndex_NewQModelIndex(){
	return new QModelIndex();
}

void* QModelIndex_Child(void* ptr, int row, int column){
	return static_cast<QModelIndex*>(ptr)->child(row, column).internalPointer();
}

int QModelIndex_Column(void* ptr){
	return static_cast<QModelIndex*>(ptr)->column();
}

void* QModelIndex_Data(void* ptr, int role){
	return new QVariant(static_cast<QModelIndex*>(ptr)->data(role));
}

int QModelIndex_Flags(void* ptr){
	return static_cast<QModelIndex*>(ptr)->flags();
}

void* QModelIndex_InternalPointer(void* ptr){
	return static_cast<QModelIndex*>(ptr)->internalPointer();
}

int QModelIndex_IsValid(void* ptr){
	return static_cast<QModelIndex*>(ptr)->isValid();
}

void* QModelIndex_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QModelIndex*>(ptr)->model());
}

void* QModelIndex_Parent(void* ptr){
	return static_cast<QModelIndex*>(ptr)->parent().internalPointer();
}

int QModelIndex_Row(void* ptr){
	return static_cast<QModelIndex*>(ptr)->row();
}

void* QModelIndex_Sibling(void* ptr, int row, int column){
	return static_cast<QModelIndex*>(ptr)->sibling(row, column).internalPointer();
}

