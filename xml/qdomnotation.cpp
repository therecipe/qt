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

QtObjectPtr QDomNotation_NewQDomNotation(){
	return new QDomNotation();
}

QtObjectPtr QDomNotation_NewQDomNotation2(QtObjectPtr x){
	return new QDomNotation(*static_cast<QDomNotation*>(x));
}

int QDomNotation_NodeType(QtObjectPtr ptr){
	return static_cast<QDomNotation*>(ptr)->nodeType();
}

char* QDomNotation_PublicId(QtObjectPtr ptr){
	return static_cast<QDomNotation*>(ptr)->publicId().toUtf8().data();
}

char* QDomNotation_SystemId(QtObjectPtr ptr){
	return static_cast<QDomNotation*>(ptr)->systemId().toUtf8().data();
}

