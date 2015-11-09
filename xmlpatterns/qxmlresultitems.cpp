#include "qxmlresultitems.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlResultItems>
#include "_cgo_export.h"

class MyQXmlResultItems: public QXmlResultItems {
public:
};

void* QXmlResultItems_NewQXmlResultItems(){
	return new QXmlResultItems();
}

int QXmlResultItems_HasError(void* ptr){
	return static_cast<QXmlResultItems*>(ptr)->hasError();
}

void QXmlResultItems_DestroyQXmlResultItems(void* ptr){
	static_cast<QXmlResultItems*>(ptr)->~QXmlResultItems();
}

