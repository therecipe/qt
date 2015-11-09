#include "qxmlstreamattribute.h"
#include <QStringRef>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlStreamAttribute>
#include "_cgo_export.h"

class MyQXmlStreamAttribute: public QXmlStreamAttribute {
public:
};

void* QXmlStreamAttribute_NewQXmlStreamAttribute(){
	return new QXmlStreamAttribute();
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute3(char* namespaceUri, char* name, char* value){
	return new QXmlStreamAttribute(QString(namespaceUri), QString(name), QString(value));
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute2(char* qualifiedName, char* value){
	return new QXmlStreamAttribute(QString(qualifiedName), QString(value));
}

void* QXmlStreamAttribute_NewQXmlStreamAttribute4(void* other){
	return new QXmlStreamAttribute(*static_cast<QXmlStreamAttribute*>(other));
}

int QXmlStreamAttribute_IsDefault(void* ptr){
	return static_cast<QXmlStreamAttribute*>(ptr)->isDefault();
}

void* QXmlStreamAttribute_Name(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->name());
}

void* QXmlStreamAttribute_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->namespaceUri());
}

void* QXmlStreamAttribute_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->prefix());
}

void* QXmlStreamAttribute_QualifiedName(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->qualifiedName());
}

void* QXmlStreamAttribute_Value(void* ptr){
	return new QStringRef(static_cast<QXmlStreamAttribute*>(ptr)->value());
}

void QXmlStreamAttribute_DestroyQXmlStreamAttribute(void* ptr){
	static_cast<QXmlStreamAttribute*>(ptr)->~QXmlStreamAttribute();
}

