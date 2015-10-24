#include "qabstractxmlnodemodel.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QString>
#include <QVariant>
#include <QAbstractXmlNodeModel>
#include "_cgo_export.h"

class MyQAbstractXmlNodeModel: public QAbstractXmlNodeModel {
public:
};

char* QAbstractXmlNodeModel_BaseUri(QtObjectPtr ptr, QtObjectPtr n){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(n)).toString().toUtf8().data();
}

int QAbstractXmlNodeModel_CompareOrder(QtObjectPtr ptr, QtObjectPtr ni1, QtObjectPtr ni2){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

char* QAbstractXmlNodeModel_DocumentUri(QtObjectPtr ptr, QtObjectPtr n){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->documentUri(*static_cast<QXmlNodeModelIndex*>(n)).toString().toUtf8().data();
}

int QAbstractXmlNodeModel_Kind(QtObjectPtr ptr, QtObjectPtr ni){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

char* QAbstractXmlNodeModel_StringValue(QtObjectPtr ptr, QtObjectPtr n){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(n)).toUtf8().data();
}

char* QAbstractXmlNodeModel_TypedValue(QtObjectPtr ptr, QtObjectPtr node){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->typedValue(*static_cast<QXmlNodeModelIndex*>(node)).toString().toUtf8().data();
}

void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(QtObjectPtr ptr){
	static_cast<QAbstractXmlNodeModel*>(ptr)->~QAbstractXmlNodeModel();
}

