#include "qxmlstreamentitydeclaration.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QXmlStreamEntityDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamEntityDeclaration: public QXmlStreamEntityDeclaration {
public:
};

QtObjectPtr QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration(){
	return new QXmlStreamEntityDeclaration();
}

QtObjectPtr QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration2(QtObjectPtr other){
	return new QXmlStreamEntityDeclaration(*static_cast<QXmlStreamEntityDeclaration*>(other));
}

void QXmlStreamEntityDeclaration_DestroyQXmlStreamEntityDeclaration(QtObjectPtr ptr){
	static_cast<QXmlStreamEntityDeclaration*>(ptr)->~QXmlStreamEntityDeclaration();
}

