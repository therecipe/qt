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

void* QDomAttr_NewQDomAttr(){
	return new QDomAttr();
}

void* QDomAttr_NewQDomAttr2(void* x){
	return new QDomAttr(*static_cast<QDomAttr*>(x));
}

char* QDomAttr_Name(void* ptr){
	return static_cast<QDomAttr*>(ptr)->name().toUtf8().data();
}

int QDomAttr_NodeType(void* ptr){
	return static_cast<QDomAttr*>(ptr)->nodeType();
}

void QDomAttr_SetValue(void* ptr, char* v){
	static_cast<QDomAttr*>(ptr)->setValue(QString(v));
}

int QDomAttr_Specified(void* ptr){
	return static_cast<QDomAttr*>(ptr)->specified();
}

char* QDomAttr_Value(void* ptr){
	return static_cast<QDomAttr*>(ptr)->value().toUtf8().data();
}

