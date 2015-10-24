#include "qxmlstreamattribute.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QXmlStreamAttribute>
#include "_cgo_export.h"

class MyQXmlStreamAttribute: public QXmlStreamAttribute {
public:
};

QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute(){
	return new QXmlStreamAttribute();
}

QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute3(char* namespaceUri, char* name, char* value){
	return new QXmlStreamAttribute(QString(namespaceUri), QString(name), QString(value));
}

QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute2(char* qualifiedName, char* value){
	return new QXmlStreamAttribute(QString(qualifiedName), QString(value));
}

QtObjectPtr QXmlStreamAttribute_NewQXmlStreamAttribute4(QtObjectPtr other){
	return new QXmlStreamAttribute(*static_cast<QXmlStreamAttribute*>(other));
}

int QXmlStreamAttribute_IsDefault(QtObjectPtr ptr){
	return static_cast<QXmlStreamAttribute*>(ptr)->isDefault();
}

void QXmlStreamAttribute_DestroyQXmlStreamAttribute(QtObjectPtr ptr){
	static_cast<QXmlStreamAttribute*>(ptr)->~QXmlStreamAttribute();
}

