#include "qsimplexmlnodemodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QSimpleXmlNodeModel>
#include "_cgo_export.h"

class MyQSimpleXmlNodeModel: public QSimpleXmlNodeModel {
public:
};

char* QSimpleXmlNodeModel_StringValue(void* ptr, void* node){
	return static_cast<QSimpleXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(node)).toUtf8().data();
}

void QSimpleXmlNodeModel_DestroyQSimpleXmlNodeModel(void* ptr){
	static_cast<QSimpleXmlNodeModel*>(ptr)->~QSimpleXmlNodeModel();
}

