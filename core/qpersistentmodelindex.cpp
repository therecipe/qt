#include "qpersistentmodelindex.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QPersistentModelIndex>
#include "_cgo_export.h"

class MyQPersistentModelIndex: public QPersistentModelIndex {
public:
};

QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex3(QtObjectPtr other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

int QPersistentModelIndex_Column(QtObjectPtr ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->column();
}

int QPersistentModelIndex_IsValid(QtObjectPtr ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->isValid();
}

int QPersistentModelIndex_Row(QtObjectPtr ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->row();
}

QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex4(QtObjectPtr other){
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

QtObjectPtr QPersistentModelIndex_NewQPersistentModelIndex(QtObjectPtr index){
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QPersistentModelIndex_Child(QtObjectPtr ptr, int row, int column){
	return static_cast<QPersistentModelIndex*>(ptr)->child(row, column).internalPointer();
}

char* QPersistentModelIndex_Data(QtObjectPtr ptr, int role){
	return static_cast<QPersistentModelIndex*>(ptr)->data(role).toString().toUtf8().data();
}

int QPersistentModelIndex_Flags(QtObjectPtr ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->flags();
}

QtObjectPtr QPersistentModelIndex_Model(QtObjectPtr ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QPersistentModelIndex*>(ptr)->model());
}

QtObjectPtr QPersistentModelIndex_Parent(QtObjectPtr ptr){
	return static_cast<QPersistentModelIndex*>(ptr)->parent().internalPointer();
}

QtObjectPtr QPersistentModelIndex_Sibling(QtObjectPtr ptr, int row, int column){
	return static_cast<QPersistentModelIndex*>(ptr)->sibling(row, column).internalPointer();
}

void QPersistentModelIndex_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QPersistentModelIndex*>(ptr)->swap(*static_cast<QPersistentModelIndex*>(other));
}

