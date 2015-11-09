#include "qabstractxmlnodemodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QXmlNodeModelIndex>
#include <QAbstractXmlNodeModel>
#include "_cgo_export.h"

class MyQAbstractXmlNodeModel: public QAbstractXmlNodeModel {
public:
};

int QAbstractXmlNodeModel_CompareOrder(void* ptr, void* ni1, void* ni2){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->compareOrder(*static_cast<QXmlNodeModelIndex*>(ni1), *static_cast<QXmlNodeModelIndex*>(ni2));
}

int QAbstractXmlNodeModel_Kind(void* ptr, void* ni){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->kind(*static_cast<QXmlNodeModelIndex*>(ni));
}

char* QAbstractXmlNodeModel_StringValue(void* ptr, void* n){
	return static_cast<QAbstractXmlNodeModel*>(ptr)->stringValue(*static_cast<QXmlNodeModelIndex*>(n)).toUtf8().data();
}

void* QAbstractXmlNodeModel_TypedValue(void* ptr, void* node){
	return new QVariant(static_cast<QAbstractXmlNodeModel*>(ptr)->typedValue(*static_cast<QXmlNodeModelIndex*>(node)));
}

void QAbstractXmlNodeModel_DestroyQAbstractXmlNodeModel(void* ptr){
	static_cast<QAbstractXmlNodeModel*>(ptr)->~QAbstractXmlNodeModel();
}

