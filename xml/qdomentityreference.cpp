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

QtObjectPtr QDomEntityReference_NewQDomEntityReference(){
	return new QDomEntityReference();
}

QtObjectPtr QDomEntityReference_NewQDomEntityReference2(QtObjectPtr x){
	return new QDomEntityReference(*static_cast<QDomEntityReference*>(x));
}

int QDomEntityReference_NodeType(QtObjectPtr ptr){
	return static_cast<QDomEntityReference*>(ptr)->nodeType();
}

