#include "qdomnotation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDomNotation>
#include "_cgo_export.h"

class MyQDomNotation: public QDomNotation {
public:
};

void* QDomNotation_NewQDomNotation(){
	return new QDomNotation();
}

void* QDomNotation_NewQDomNotation2(void* x){
	return new QDomNotation(*static_cast<QDomNotation*>(x));
}

int QDomNotation_NodeType(void* ptr){
	return static_cast<QDomNotation*>(ptr)->nodeType();
}

char* QDomNotation_PublicId(void* ptr){
	return static_cast<QDomNotation*>(ptr)->publicId().toUtf8().data();
}

char* QDomNotation_SystemId(void* ptr){
	return static_cast<QDomNotation*>(ptr)->systemId().toUtf8().data();
}

