#include "qxmlnodemodelindex.h"
#include <QUrl>
#include <QModelIndex>
#include <QAbstractXmlNodeModel>
#include <QString>
#include <QVariant>
#include <QXmlNodeModelIndex>
#include "_cgo_export.h"

class MyQXmlNodeModelIndex: public QXmlNodeModelIndex {
public:
};

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex(){
	return new QXmlNodeModelIndex();
}

void* QXmlNodeModelIndex_NewQXmlNodeModelIndex2(void* other){
	return new QXmlNodeModelIndex(*static_cast<QXmlNodeModelIndex*>(other));
}

void* QXmlNodeModelIndex_InternalPointer(void* ptr){
	return static_cast<QXmlNodeModelIndex*>(ptr)->internalPointer();
}

int QXmlNodeModelIndex_IsNull(void* ptr){
	return static_cast<QXmlNodeModelIndex*>(ptr)->isNull();
}

void* QXmlNodeModelIndex_Model(void* ptr){
	return const_cast<QAbstractXmlNodeModel*>(static_cast<QXmlNodeModelIndex*>(ptr)->model());
}

