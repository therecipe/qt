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

QtObjectPtr QXmlNodeModelIndex_NewQXmlNodeModelIndex(){
	return new QXmlNodeModelIndex();
}

QtObjectPtr QXmlNodeModelIndex_NewQXmlNodeModelIndex2(QtObjectPtr other){
	return new QXmlNodeModelIndex(*static_cast<QXmlNodeModelIndex*>(other));
}

void QXmlNodeModelIndex_InternalPointer(QtObjectPtr ptr){
	static_cast<QXmlNodeModelIndex*>(ptr)->internalPointer();
}

int QXmlNodeModelIndex_IsNull(QtObjectPtr ptr){
	return static_cast<QXmlNodeModelIndex*>(ptr)->isNull();
}

QtObjectPtr QXmlNodeModelIndex_Model(QtObjectPtr ptr){
	return const_cast<QAbstractXmlNodeModel*>(static_cast<QXmlNodeModelIndex*>(ptr)->model());
}

