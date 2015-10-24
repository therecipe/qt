#include "qxmllexicalhandler.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QXmlLexicalHandler>
#include "_cgo_export.h"

class MyQXmlLexicalHandler: public QXmlLexicalHandler {
public:
};

int QXmlLexicalHandler_Comment(QtObjectPtr ptr, char* ch){
	return static_cast<QXmlLexicalHandler*>(ptr)->comment(QString(ch));
}

int QXmlLexicalHandler_EndCDATA(QtObjectPtr ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endCDATA();
}

int QXmlLexicalHandler_EndDTD(QtObjectPtr ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endDTD();
}

int QXmlLexicalHandler_EndEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->endEntity(QString(name));
}

char* QXmlLexicalHandler_ErrorString(QtObjectPtr ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlLexicalHandler_StartCDATA(QtObjectPtr ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->startCDATA();
}

int QXmlLexicalHandler_StartDTD(QtObjectPtr ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlLexicalHandler*>(ptr)->startDTD(QString(name), QString(publicId), QString(systemId));
}

int QXmlLexicalHandler_StartEntity(QtObjectPtr ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->startEntity(QString(name));
}

void QXmlLexicalHandler_DestroyQXmlLexicalHandler(QtObjectPtr ptr){
	static_cast<QXmlLexicalHandler*>(ptr)->~QXmlLexicalHandler();
}

