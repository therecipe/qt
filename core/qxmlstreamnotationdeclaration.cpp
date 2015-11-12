#include "qxmlstreamnotationdeclaration.h"
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QString>
#include <QVariant>
#include <QXmlStreamNotationDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamNotationDeclaration: public QXmlStreamNotationDeclaration {
public:
};

void* QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration(){
	return new QXmlStreamNotationDeclaration();
}

void* QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration2(void* other){
	return new QXmlStreamNotationDeclaration(*static_cast<QXmlStreamNotationDeclaration*>(other));
}

void* QXmlStreamNotationDeclaration_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->name());
}

void* QXmlStreamNotationDeclaration_PublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->publicId());
}

void* QXmlStreamNotationDeclaration_SystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNotationDeclaration*>(ptr)->systemId());
}

void QXmlStreamNotationDeclaration_DestroyQXmlStreamNotationDeclaration(void* ptr){
	static_cast<QXmlStreamNotationDeclaration*>(ptr)->~QXmlStreamNotationDeclaration();
}

