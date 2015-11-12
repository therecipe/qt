#include "qxmllexicalhandler.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlLexicalHandler>
#include "_cgo_export.h"

class MyQXmlLexicalHandler: public QXmlLexicalHandler {
public:
};

int QXmlLexicalHandler_Comment(void* ptr, char* ch){
	return static_cast<QXmlLexicalHandler*>(ptr)->comment(QString(ch));
}

int QXmlLexicalHandler_EndCDATA(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endCDATA();
}

int QXmlLexicalHandler_EndDTD(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->endDTD();
}

int QXmlLexicalHandler_EndEntity(void* ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->endEntity(QString(name));
}

char* QXmlLexicalHandler_ErrorString(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->errorString().toUtf8().data();
}

int QXmlLexicalHandler_StartCDATA(void* ptr){
	return static_cast<QXmlLexicalHandler*>(ptr)->startCDATA();
}

int QXmlLexicalHandler_StartDTD(void* ptr, char* name, char* publicId, char* systemId){
	return static_cast<QXmlLexicalHandler*>(ptr)->startDTD(QString(name), QString(publicId), QString(systemId));
}

int QXmlLexicalHandler_StartEntity(void* ptr, char* name){
	return static_cast<QXmlLexicalHandler*>(ptr)->startEntity(QString(name));
}

void QXmlLexicalHandler_DestroyQXmlLexicalHandler(void* ptr){
	static_cast<QXmlLexicalHandler*>(ptr)->~QXmlLexicalHandler();
}

