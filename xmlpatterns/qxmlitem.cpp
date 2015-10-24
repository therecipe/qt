#include "qxmlitem.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QXmlItem>
#include "_cgo_export.h"

class MyQXmlItem: public QXmlItem {
public:
};

QtObjectPtr QXmlItem_NewQXmlItem(){
	return new QXmlItem();
}

QtObjectPtr QXmlItem_NewQXmlItem4(char* atomicValue){
	return new QXmlItem(QVariant(atomicValue));
}

QtObjectPtr QXmlItem_NewQXmlItem2(QtObjectPtr other){
	return new QXmlItem(*static_cast<QXmlItem*>(other));
}

QtObjectPtr QXmlItem_NewQXmlItem3(QtObjectPtr node){
	return new QXmlItem(*static_cast<QXmlNodeModelIndex*>(node));
}

int QXmlItem_IsNode(QtObjectPtr ptr){
	return static_cast<QXmlItem*>(ptr)->isNode();
}

int QXmlItem_IsNull(QtObjectPtr ptr){
	return static_cast<QXmlItem*>(ptr)->isNull();
}

void QXmlItem_DestroyQXmlItem(QtObjectPtr ptr){
	static_cast<QXmlItem*>(ptr)->~QXmlItem();
}

