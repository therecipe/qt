#include "qxmlnamespacesupport.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNamespaceSupport>
#include "_cgo_export.h"

class MyQXmlNamespaceSupport: public QXmlNamespaceSupport {
public:
};

void* QXmlNamespaceSupport_NewQXmlNamespaceSupport(){
	return new QXmlNamespaceSupport();
}

void QXmlNamespaceSupport_PopContext(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->popContext();
}

char* QXmlNamespaceSupport_Prefix(void* ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefix(QString(uri)).toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes(void* ptr){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes().join("|").toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes2(void* ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes(QString(uri)).join("|").toUtf8().data();
}

void QXmlNamespaceSupport_PushContext(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->pushContext();
}

void QXmlNamespaceSupport_Reset(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->reset();
}

void QXmlNamespaceSupport_SetPrefix(void* ptr, char* pre, char* uri){
	static_cast<QXmlNamespaceSupport*>(ptr)->setPrefix(QString(pre), QString(uri));
}

char* QXmlNamespaceSupport_Uri(void* ptr, char* prefix){
	return static_cast<QXmlNamespaceSupport*>(ptr)->uri(QString(prefix)).toUtf8().data();
}

void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(void* ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->~QXmlNamespaceSupport();
}

