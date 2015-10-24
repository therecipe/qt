#include "qxmlnamepool.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlName>
#include <QString>
#include <QVariant>
#include <QXmlNamePool>
#include "_cgo_export.h"

class MyQXmlNamePool: public QXmlNamePool {
public:
};

QtObjectPtr QXmlNamePool_NewQXmlNamePool(){
	return new QXmlNamePool();
}

QtObjectPtr QXmlNamePool_NewQXmlNamePool2(QtObjectPtr other){
	return new QXmlNamePool(*static_cast<QXmlNamePool*>(other));
}

void QXmlNamePool_DestroyQXmlNamePool(QtObjectPtr ptr){
	static_cast<QXmlNamePool*>(ptr)->~QXmlNamePool();
}

