#include "qprocessenvironment.h"
#include <QProcess>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProcessEnvironment>
#include "_cgo_export.h"

class MyQProcessEnvironment: public QProcessEnvironment {
public:
};

void* QProcessEnvironment_NewQProcessEnvironment(){
	return new QProcessEnvironment();
}

void* QProcessEnvironment_NewQProcessEnvironment2(void* other){
	return new QProcessEnvironment(*static_cast<QProcessEnvironment*>(other));
}

void QProcessEnvironment_Clear(void* ptr){
	static_cast<QProcessEnvironment*>(ptr)->clear();
}

int QProcessEnvironment_Contains(void* ptr, char* name){
	return static_cast<QProcessEnvironment*>(ptr)->contains(QString(name));
}

int QProcessEnvironment_IsEmpty(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->isEmpty();
}

char* QProcessEnvironment_Keys(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->keys().join("|").toUtf8().data();
}

void QProcessEnvironment_Swap(void* ptr, void* other){
	static_cast<QProcessEnvironment*>(ptr)->swap(*static_cast<QProcessEnvironment*>(other));
}

char* QProcessEnvironment_ToStringList(void* ptr){
	return static_cast<QProcessEnvironment*>(ptr)->toStringList().join("|").toUtf8().data();
}

char* QProcessEnvironment_Value(void* ptr, char* name, char* defaultValue){
	return static_cast<QProcessEnvironment*>(ptr)->value(QString(name), QString(defaultValue)).toUtf8().data();
}

void QProcessEnvironment_DestroyQProcessEnvironment(void* ptr){
	static_cast<QProcessEnvironment*>(ptr)->~QProcessEnvironment();
}

