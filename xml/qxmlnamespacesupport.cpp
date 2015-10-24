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

QtObjectPtr QXmlNamespaceSupport_NewQXmlNamespaceSupport(){
	return new QXmlNamespaceSupport();
}

void QXmlNamespaceSupport_PopContext(QtObjectPtr ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->popContext();
}

char* QXmlNamespaceSupport_Prefix(QtObjectPtr ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefix(QString(uri)).toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes(QtObjectPtr ptr){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes().join("|").toUtf8().data();
}

char* QXmlNamespaceSupport_Prefixes2(QtObjectPtr ptr, char* uri){
	return static_cast<QXmlNamespaceSupport*>(ptr)->prefixes(QString(uri)).join("|").toUtf8().data();
}

void QXmlNamespaceSupport_PushContext(QtObjectPtr ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->pushContext();
}

void QXmlNamespaceSupport_Reset(QtObjectPtr ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->reset();
}

void QXmlNamespaceSupport_SetPrefix(QtObjectPtr ptr, char* pre, char* uri){
	static_cast<QXmlNamespaceSupport*>(ptr)->setPrefix(QString(pre), QString(uri));
}

char* QXmlNamespaceSupport_Uri(QtObjectPtr ptr, char* prefix){
	return static_cast<QXmlNamespaceSupport*>(ptr)->uri(QString(prefix)).toUtf8().data();
}

void QXmlNamespaceSupport_DestroyQXmlNamespaceSupport(QtObjectPtr ptr){
	static_cast<QXmlNamespaceSupport*>(ptr)->~QXmlNamespaceSupport();
}

