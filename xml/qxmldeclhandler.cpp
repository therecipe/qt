#include "qxmldeclhandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlDeclHandler>
#include "_cgo_export.h"

class MyQXmlDeclHandler: public QXmlDeclHandler {
public:
};

int QXmlDeclHandler_AttributeDecl(QtObjectPtr ptr, char* eName, char* aName, char* ty, char* valueDefault, char* value){
	return static_cast<QXmlDeclHandler*>(ptr)->attributeDecl(QString(eName), QString(aName), QString(ty), QString(valueDefault), QString(value));
}

char* QXmlDeclHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlDeclHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDeclHandler_ExternalEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDeclHandler*>(ptr)->externalEntityDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDeclHandler_InternalEntityDecl(QtObjectPtr ptr, char* name, char* value){
	return static_cast<QXmlDeclHandler*>(ptr)->internalEntityDecl(QString(name), QString(value));
}

void QXmlDeclHandler_DestroyQXmlDeclHandler(QtObjectPtr ptr){
	static_cast<QXmlDeclHandler*>(ptr)->~QXmlDeclHandler();
}

