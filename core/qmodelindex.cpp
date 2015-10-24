#include "qmodelindex.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemModel>
#include <QString>
#include <QModelIndex>
#include "_cgo_export.h"

class MyQModelIndex: public QModelIndex {
public:
};

QtObjectPtr QModelIndex_NewQModelIndex(){
	return new QModelIndex();
}

QtObjectPtr QModelIndex_Child(QtObjectPtr ptr, int row, int column){
	return static_cast<QModelIndex*>(ptr)->child(row, column).internalPointer();
}

int QModelIndex_Column(QtObjectPtr ptr){
	return static_cast<QModelIndex*>(ptr)->column();
}

char* QModelIndex_Data(QtObjectPtr ptr, int role){
	return static_cast<QModelIndex*>(ptr)->data(role).toString().toUtf8().data();
}

int QModelIndex_Flags(QtObjectPtr ptr){
	return static_cast<QModelIndex*>(ptr)->flags();
}

void QModelIndex_InternalPointer(QtObjectPtr ptr){
	static_cast<QModelIndex*>(ptr)->internalPointer();
}

int QModelIndex_IsValid(QtObjectPtr ptr){
	return static_cast<QModelIndex*>(ptr)->isValid();
}

QtObjectPtr QModelIndex_Model(QtObjectPtr ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QModelIndex*>(ptr)->model());
}

QtObjectPtr QModelIndex_Parent(QtObjectPtr ptr){
	return static_cast<QModelIndex*>(ptr)->parent().internalPointer();
}

int QModelIndex_Row(QtObjectPtr ptr){
	return static_cast<QModelIndex*>(ptr)->row();
}

QtObjectPtr QModelIndex_Sibling(QtObjectPtr ptr, int row, int column){
	return static_cast<QModelIndex*>(ptr)->sibling(row, column).internalPointer();
}

