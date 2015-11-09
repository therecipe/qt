#include "qdomentityreference.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomEntity>
#include <QDomEntityReference>
#include "_cgo_export.h"

class MyQDomEntityReference: public QDomEntityReference {
public:
};

void* QDomEntityReference_NewQDomEntityReference(){
	return new QDomEntityReference();
}

void* QDomEntityReference_NewQDomEntityReference2(void* x){
	return new QDomEntityReference(*static_cast<QDomEntityReference*>(x));
}

int QDomEntityReference_NodeType(void* ptr){
	return static_cast<QDomEntityReference*>(ptr)->nodeType();
}

