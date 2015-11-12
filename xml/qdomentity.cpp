#include "qdomentity.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QDomEntity>
#include "_cgo_export.h"

class MyQDomEntity: public QDomEntity {
public:
};

void* QDomEntity_NewQDomEntity(){
	return new QDomEntity();
}

void* QDomEntity_NewQDomEntity2(void* x){
	return new QDomEntity(*static_cast<QDomEntity*>(x));
}

int QDomEntity_NodeType(void* ptr){
	return static_cast<QDomEntity*>(ptr)->nodeType();
}

char* QDomEntity_NotationName(void* ptr){
	return static_cast<QDomEntity*>(ptr)->notationName().toUtf8().data();
}

char* QDomEntity_PublicId(void* ptr){
	return static_cast<QDomEntity*>(ptr)->publicId().toUtf8().data();
}

char* QDomEntity_SystemId(void* ptr){
	return static_cast<QDomEntity*>(ptr)->systemId().toUtf8().data();
}

