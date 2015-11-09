#include "qpersistentmodelindex.h"
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPersistentModelIndex>
#include "_cgo_export.h"

class MyQPersistentModelIndex: public QPersistentModelIndex {
public:
};

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

int QPersistentModelIndex_Column(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->column();
}

int QPersistentModelIndex_IsValid(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->isValid();
}

int QPersistentModelIndex_Row(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->row();
}

void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_NewQPersistentModelIndex(void* index){
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

void* QPersistentModelIndex_Child(void* ptr, int row, int column){
	return static_cast<QPersistentModelIndex*>(ptr)->child(row, column).internalPointer();
}

void* QPersistentModelIndex_Data(void* ptr, int role){
	return new QVariant(static_cast<QPersistentModelIndex*>(ptr)->data(role));
}

int QPersistentModelIndex_Flags(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->flags();
}

void* QPersistentModelIndex_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QPersistentModelIndex*>(ptr)->model());
}

void* QPersistentModelIndex_Parent(void* ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->parent().internalPointer();
}

void* QPersistentModelIndex_Sibling(void* ptr, int row, int column){
	return static_cast<QPersistentModelIndex*>(ptr)->sibling(row, column).internalPointer();
}

void QPersistentModelIndex_Swap(void* ptr, void* other){
	static_cast<QPersistentModelIndex*>(ptr)->swap(*static_cast<QPersistentModelIndex*>(other));
}

