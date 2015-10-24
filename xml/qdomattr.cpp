#include "qdomattr.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QDomAttr>
#include "_cgo_export.h"

class MyQDomAttr: public QDomAttr {
public:
};

QtObjectPtr QDomAttr_NewQDomAttr(){
	return new QDomAttr();
}

QtObjectPtr QDomAttr_NewQDomAttr2(QtObjectPtr x){
	return new QDomAttr(*static_cast<QDomAttr*>(x));
}

char* QDomAttr_Name(QtObjectPtr ptr){
	return static_cast<QDomAttr*>(ptr)->name().toUtf8().data();
}

int QDomAttr_NodeType(QtObjectPtr ptr){
	return static_cast<QDomAttr*>(ptr)->nodeType();
}

void QDomAttr_SetValue(QtObjectPtr ptr, char* v){
	static_cast<QDomAttr*>(ptr)->setValue(QString(v));
}

int QDomAttr_Specified(QtObjectPtr ptr){
	return static_cast<QDomAttr*>(ptr)->specified();
}

char* QDomAttr_Value(QtObjectPtr ptr){
	return static_cast<QDomAttr*>(ptr)->value().toUtf8().data();
}

