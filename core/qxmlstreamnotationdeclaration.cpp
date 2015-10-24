#include "qxmlstreamnotationdeclaration.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlStreamNotationDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamNotationDeclaration: public QXmlStreamNotationDeclaration {
public:
};

QtObjectPtr QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration(){
	return new QXmlStreamNotationDeclaration();
}

QtObjectPtr QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration2(QtObjectPtr other){
	return new QXmlStreamNotationDeclaration(*static_cast<QXmlStreamNotationDeclaration*>(other));
}

void QXmlStreamNotationDeclaration_DestroyQXmlStreamNotationDeclaration(QtObjectPtr ptr){
	static_cast<QXmlStreamNotationDeclaration*>(ptr)->~QXmlStreamNotationDeclaration();
}

