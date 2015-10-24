#include "qdomelement.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomElement>
#include "_cgo_export.h"

class MyQDomElement: public QDomElement {
public:
};

QtObjectPtr QDomElement_NewQDomElement(){
	return new QDomElement();
}

QtObjectPtr QDomElement_NewQDomElement2(QtObjectPtr x){
	return new QDomElement(*static_cast<QDomElement*>(x));
}

char* QDomElement_Attribute(QtObjectPtr ptr, char* name, char* defValue){
	return static_cast<QDomElement*>(ptr)->attribute(QString(name), QString(defValue)).toUtf8().data();
}

char* QDomElement_AttributeNS(QtObjectPtr ptr, char* nsURI, char* localName, char* defValue){
	return static_cast<QDomElement*>(ptr)->attributeNS(QString(nsURI), QString(localName), QString(defValue)).toUtf8().data();
}

int QDomElement_HasAttribute(QtObjectPtr ptr, char* name){
	return static_cast<QDomElement*>(ptr)->hasAttribute(QString(name));
}

int QDomElement_HasAttributeNS(QtObjectPtr ptr, char* nsURI, char* localName){
	return static_cast<QDomElement*>(ptr)->hasAttributeNS(QString(nsURI), QString(localName));
}

int QDomElement_NodeType(QtObjectPtr ptr){
	return static_cast<QDomElement*>(ptr)->nodeType();
}

void QDomElement_RemoveAttribute(QtObjectPtr ptr, char* name){
	static_cast<QDomElement*>(ptr)->removeAttribute(QString(name));
}

void QDomElement_RemoveAttributeNS(QtObjectPtr ptr, char* nsURI, char* localName){
	static_cast<QDomElement*>(ptr)->removeAttributeNS(QString(nsURI), QString(localName));
}

void QDomElement_SetAttribute(QtObjectPtr ptr, char* name, char* value){
	static_cast<QDomElement*>(ptr)->setAttribute(QString(name), QString(value));
}

void QDomElement_SetAttribute2(QtObjectPtr ptr, char* name, int value){
	static_cast<QDomElement*>(ptr)->setAttribute(QString(name), value);
}

void QDomElement_SetAttributeNS(QtObjectPtr ptr, char* nsURI, char* qName, char* value){
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString(nsURI), QString(qName), QString(value));
}

void QDomElement_SetAttributeNS2(QtObjectPtr ptr, char* nsURI, char* qName, int value){
	static_cast<QDomElement*>(ptr)->setAttributeNS(QString(nsURI), QString(qName), value);
}

void QDomElement_SetTagName(QtObjectPtr ptr, char* name){
	static_cast<QDomElement*>(ptr)->setTagName(QString(name));
}

char* QDomElement_TagName(QtObjectPtr ptr){
	return static_cast<QDomElement*>(ptr)->tagName().toUtf8().data();
}

char* QDomElement_Text(QtObjectPtr ptr){
	return static_cast<QDomElement*>(ptr)->text().toUtf8().data();
}

