#include "qdomdocument.h"
#include <QByteArray>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlInputSource>
#include <QString>
#include <QXmlReader>
#include <QDomDocumentType>
#include <QIODevice>
#include <QDomDocument>
#include "_cgo_export.h"

class MyQDomDocument: public QDomDocument {
public:
};

void* QDomDocument_NewQDomDocument(){
	return new QDomDocument();
}

void* QDomDocument_NewQDomDocument4(void* x){
	return new QDomDocument(*static_cast<QDomDocument*>(x));
}

void* QDomDocument_NewQDomDocument3(void* doctype){
	return new QDomDocument(*static_cast<QDomDocumentType*>(doctype));
}

void* QDomDocument_NewQDomDocument2(char* name){
	return new QDomDocument(QString(name));
}

int QDomDocument_NodeType(void* ptr){
	return static_cast<QDomDocument*>(ptr)->nodeType();
}

int QDomDocument_SetContent7(void* ptr, void* dev, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent3(void* ptr, void* dev, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent8(void* ptr, void* source, void* reader, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), static_cast<QXmlReader*>(reader), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent4(void* ptr, void* source, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent6(void* ptr, void* buffer, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(buffer), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent(void* ptr, void* data, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(data), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent5(void* ptr, char* text, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent2(void* ptr, char* text, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

void* QDomDocument_ToByteArray(void* ptr, int indent){
	return new QByteArray(static_cast<QDomDocument*>(ptr)->toByteArray(indent));
}

char* QDomDocument_ToString(void* ptr, int indent){
	return static_cast<QDomDocument*>(ptr)->toString(indent).toUtf8().data();
}

void QDomDocument_DestroyQDomDocument(void* ptr){
	static_cast<QDomDocument*>(ptr)->~QDomDocument();
}

