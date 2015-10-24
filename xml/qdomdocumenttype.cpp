#include "qdomdocumenttype.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomDocument>
#include <QDomDocumentType>
#include "_cgo_export.h"

class MyQDomDocumentType: public QDomDocumentType {
public:
};

QtObjectPtr QDomDocumentType_NewQDomDocumentType(){
	return new QDomDocumentType();
}

QtObjectPtr QDomDocumentType_NewQDomDocumentType2(QtObjectPtr n){
	return new QDomDocumentType(*static_cast<QDomDocumentType*>(n));
}

char* QDomDocumentType_InternalSubset(QtObjectPtr ptr){
	return static_cast<QDomDocumentType*>(ptr)->internalSubset().toUtf8().data();
}

char* QDomDocumentType_Name(QtObjectPtr ptr){
	return static_cast<QDomDocumentType*>(ptr)->name().toUtf8().data();
}

int QDomDocumentType_NodeType(QtObjectPtr ptr){
	return static_cast<QDomDocumentType*>(ptr)->nodeType();
}

char* QDomDocumentType_PublicId(QtObjectPtr ptr){
	return static_cast<QDomDocumentType*>(ptr)->publicId().toUtf8().data();
}

char* QDomDocumentType_SystemId(QtObjectPtr ptr){
	return static_cast<QDomDocumentType*>(ptr)->systemId().toUtf8().data();
}

