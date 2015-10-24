#include "qdomentity.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomEntity>
#include "_cgo_export.h"

class MyQDomEntity: public QDomEntity {
public:
};

QtObjectPtr QDomEntity_NewQDomEntity(){
	return new QDomEntity();
}

QtObjectPtr QDomEntity_NewQDomEntity2(QtObjectPtr x){
	return new QDomEntity(*static_cast<QDomEntity*>(x));
}

int QDomEntity_NodeType(QtObjectPtr ptr){
	return static_cast<QDomEntity*>(ptr)->nodeType();
}

char* QDomEntity_NotationName(QtObjectPtr ptr){
	return static_cast<QDomEntity*>(ptr)->notationName().toUtf8().data();
}

char* QDomEntity_PublicId(QtObjectPtr ptr){
	return static_cast<QDomEntity*>(ptr)->publicId().toUtf8().data();
}

char* QDomEntity_SystemId(QtObjectPtr ptr){
	return static_cast<QDomEntity*>(ptr)->systemId().toUtf8().data();
}

