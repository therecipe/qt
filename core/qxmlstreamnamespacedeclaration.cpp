#include "qxmlstreamnamespacedeclaration.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStringRef>
#include <QXmlStreamNamespaceDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamNamespaceDeclaration: public QXmlStreamNamespaceDeclaration {
public:
};

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration(){
	return new QXmlStreamNamespaceDeclaration();
}

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration3(char* prefix, char* namespaceUri){
	return new QXmlStreamNamespaceDeclaration(QString(prefix), QString(namespaceUri));
}

void* QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration2(void* other){
	return new QXmlStreamNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(other));
}

void* QXmlStreamNamespaceDeclaration_NamespaceUri(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->namespaceUri());
}

void* QXmlStreamNamespaceDeclaration_Prefix(void* ptr){
	return new QStringRef(static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->prefix());
}

void QXmlStreamNamespaceDeclaration_DestroyQXmlStreamNamespaceDeclaration(void* ptr){
	static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->~QXmlStreamNamespaceDeclaration();
}

