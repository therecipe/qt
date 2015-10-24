#include "qsimplexmlnodemodel.h"
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QString>
#include <QVariant>
#include <QSimpleXmlNodeModel>
#include "_cgo_export.h"

class MyQSimpleXmlNodeModel: public QSimpleXmlNodeModel {
public:
};

char* QSimpleXmlNodeModel_BaseUri(QtObjectPtr ptr, QtObjectPtr node){
	return static_cast<QSimpleXmlNodeModel*>(ptr)->baseUri(*static_cast<QXmlNodeModelIndex*>(node)).toString().toUtf8().data();
}

char* QSimpleXmlNodeModel_StringValue(QtObjectPtr ptr, QtObjectPtr node){
	return static_cast<QSimpleXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8().data();
}

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(QtObjectPtr ptr){
	static_cast<QSimpleXmlNodeModel*>(ptr)->~QSimpleXmlNodeModel();
}

