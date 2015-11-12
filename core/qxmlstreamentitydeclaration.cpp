#include "qxmlstreamentitydeclaration.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QXmlStreamEntityDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamEntityDeclaration: public QXmlStreamEntityDeclaration {
public:
};

void* QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration(){
	return new QXmlStreamEntityDeclaration();
}

void* QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration2(void* other){
	return new QXmlStreamEntityDeclaration(*static_cast<QXmlStreamEntityDeclaration*>(other));
}

void* QXmlStreamEntityDeclaration_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->name());
}

void* QXmlStreamEntityDeclaration_NotationName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->notationName());
}

void* QXmlStreamEntityDeclaration_PublicId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->publicId());
}

void* QXmlStreamEntityDeclaration_SystemId(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->systemId());
}

void* QXmlStreamEntityDeclaration_Value(void* ptr){
	return new QStringRef(static_cast<QXmlStreamEntityDeclaration*>(ptr)->value());
}

void QXmlStreamEntityDeclaration_DestroyQXmlStreamEntityDeclaration(void* ptr){
	static_cast<QXmlStreamEntityDeclaration*>(ptr)->~QXmlStreamEntityDeclaration();
}

