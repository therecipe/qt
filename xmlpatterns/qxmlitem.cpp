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

void* QXmlItem_NewQXmlItem(){
	return new QXmlItem();
}

void* QXmlItem_NewQXmlItem4(void* atomicValue){
	return new QXmlItem(*static_cast<QVariant*>(atomicValue));
}

void* QXmlItem_NewQXmlItem2(void* other){
	return new QXmlItem(*static_cast<QXmlItem*>(other));
}

void* QXmlItem_NewQXmlItem3(void* node){
	return new QXmlItem(*static_cast<QXmlNodeModelIndex*>(node));
}

int QXmlItem_IsNode(void* ptr){
	return static_cast<QXmlItem*>(ptr)->isNode();
}

int QXmlItem_IsNull(void* ptr){
	return static_cast<QXmlItem*>(ptr)->isNull();
}

void QXmlItem_DestroyQXmlItem(void* ptr){
	static_cast<QXmlItem*>(ptr)->~QXmlItem();
}

