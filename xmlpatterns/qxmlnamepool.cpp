#include "qxmlnamepool.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QXmlNamePool>
#include "_cgo_export.h"

class MyQXmlNamePool: public QXmlNamePool {
public:
};

void* QXmlNamePool_NewQXmlNamePool(){
	return new QXmlNamePool();
}

void* QXmlNamePool_NewQXmlNamePool2(void* other){
	return new QXmlNamePool(*static_cast<QXmlNamePool*>(other));
}

void QXmlNamePool_DestroyQXmlNamePool(void* ptr){
	static_cast<QXmlNamePool*>(ptr)->~QXmlNamePool();
}

