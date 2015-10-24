#include "qdomdocument.h"
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlInputSource>
#include <QByteArray>
#include <QXmlReader>
#include <QDomDocumentType>
#include <QDomDocument>
#include "_cgo_export.h"

class MyQDomDocument: public QDomDocument {
public:
};

QtObjectPtr QDomDocument_NewQDomDocument(){
	return new QDomDocument();
}

QtObjectPtr QDomDocument_NewQDomDocument4(QtObjectPtr x){
	return new QDomDocument(*static_cast<QDomDocument*>(x));
}

QtObjectPtr QDomDocument_NewQDomDocument3(QtObjectPtr doctype){
	return new QDomDocument(*static_cast<QDomDocumentType*>(doctype));
}

QtObjectPtr QDomDocument_NewQDomDocument2(char* name){
	return new QDomDocument(QString(name));
}

int QDomDocument_NodeType(QtObjectPtr ptr){
	return static_cast<QDomDocument*>(ptr)->nodeType();
}

int QDomDocument_SetContent7(QtObjectPtr ptr, QtObjectPtr dev, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent3(QtObjectPtr ptr, QtObjectPtr dev, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QIODevice*>(dev), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent8(QtObjectPtr ptr, QtObjectPtr source, QtObjectPtr reader, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), static_cast<QXmlReader*>(reader), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent4(QtObjectPtr ptr, QtObjectPtr source, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(static_cast<QXmlInputSource*>(source), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent6(QtObjectPtr ptr, QtObjectPtr buffer, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(buffer), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent(QtObjectPtr ptr, QtObjectPtr data, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(*static_cast<QByteArray*>(data), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent5(QtObjectPtr ptr, char* text, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), new QString(errorMsg), &errorLine, &errorColumn);
}

int QDomDocument_SetContent2(QtObjectPtr ptr, char* text, int namespaceProcessing, char* errorMsg, int errorLine, int errorColumn){
	return static_cast<QDomDocument*>(ptr)->setContent(QString(text), namespaceProcessing != 0, new QString(errorMsg), &errorLine, &errorColumn);
}

char* QDomDocument_ToString(QtObjectPtr ptr, int indent){
	return static_cast<QDomDocument*>(ptr)->toString(indent).toUtf8().data();
}

void QDomDocument_DestroyQDomDocument(QtObjectPtr ptr){
	static_cast<QDomDocument*>(ptr)->~QDomDocument();
}

