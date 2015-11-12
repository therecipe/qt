#include "qdomdocumenttype.h"
#include <QModelIndex>
#include <QDomDocument>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDomDocumentType>
#include "_cgo_export.h"

class MyQDomDocumentType: public QDomDocumentType {
public:
};

void* QDomDocumentType_NewQDomDocumentType(){
	return new QDomDocumentType();
}

void* QDomDocumentType_NewQDomDocumentType2(void* n){
	return new QDomDocumentType(*static_cast<QDomDocumentType*>(n));
}

char* QDomDocumentType_InternalSubset(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->internalSubset().toUtf8().data();
}

char* QDomDocumentType_Name(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->name().toUtf8().data();
}

int QDomDocumentType_NodeType(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->nodeType();
}

char* QDomDocumentType_PublicId(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->publicId().toUtf8().data();
}

char* QDomDocumentType_SystemId(void* ptr){
	return static_cast<QDomDocumentType*>(ptr)->systemId().toUtf8().data();
}

