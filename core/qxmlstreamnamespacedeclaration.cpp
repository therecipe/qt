#include "qxmlstreamnamespacedeclaration.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlStreamNamespaceDeclaration>
#include "_cgo_export.h"

class MyQXmlStreamNamespaceDeclaration: public QXmlStreamNamespaceDeclaration {
public:
};

QtObjectPtr QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration(){
	return new QXmlStreamNamespaceDeclaration();
}

QtObjectPtr QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration3(char* prefix, char* namespaceUri){
	return new QXmlStreamNamespaceDeclaration(QString(prefix), QString(namespaceUri));
}

QtObjectPtr QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration2(QtObjectPtr other){
	return new QXmlStreamNamespaceDeclaration(*static_cast<QXmlStreamNamespaceDeclaration*>(other));
}

void QXmlStreamNamespaceDeclaration_DestroyQXmlStreamNamespaceDeclaration(QtObjectPtr ptr){
	static_cast<QXmlStreamNamespaceDeclaration*>(ptr)->~QXmlStreamNamespaceDeclaration();
}

