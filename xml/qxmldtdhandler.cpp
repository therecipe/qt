#include "qxmldtdhandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlDTDHandler>
#include "_cgo_export.h"

class MyQXmlDTDHandler: public QXmlDTDHandler {
public:
};

char* QXmlDTDHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlDTDHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlDTDHandler_NotationDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlDTDHandler*>(ptr)->notationDecl(QString(name), QString(publicId), QString(systemId));
}

int QXmlDTDHandler_UnparsedEntityDecl(QtObjectPtr ptr, char* name, char* publicId, char* systemId, char* notationName){
	return static_cast<QXmlDTDHandler*>(ptr)->unparsedEntityDecl(QString(name), QString(publicId), QString(systemId), QString(notationName));
}

void QXmlDTDHandler_DestroyQXmlDTDHandler(QtObjectPtr ptr){
	static_cast<QXmlDTDHandler*>(ptr)->~QXmlDTDHandler();
}

