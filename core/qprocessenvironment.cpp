#include "qprocessenvironment.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProcess>
#include <QProcessEnvironment>
#include "_cgo_export.h"

class MyQProcessEnvironment: public QProcessEnvironment {
public:
};

QtObjectPtr QProcessEnvironment_NewQProcessEnvironment(){
	return new QProcessEnvironment();
}

QtObjectPtr QProcessEnvironment_NewQProcessEnvironment2(QtObjectPtr other){
	return new QProcessEnvironment(*static_cast<QProcessEnvironment*>(other));
}

void QProcessEnvironment_Clear(QtObjectPtr ptr){
	static_cast<QProcessEnvironment*>(ptr)->clear();
}

int QProcessEnvironment_Contains(QtObjectPtr ptr, char* name){
	return static_cast<QProcessEnvironment*>(ptr)->contains(QString(name));
}

int QProcessEnvironment_IsEmpty(QtObjectPtr ptr){
	return static_cast<QProcessEnvironment*>(ptr)->isEmpty();
}

char* QProcessEnvironment_Keys(QtObjectPtr ptr){
	return static_cast<QProcessEnvironment*>(ptr)->keys().join("|").toUtf8().data();
}

void QProcessEnvironment_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QProcessEnvironment*>(ptr)->swap(*static_cast<QProcessEnvironment*>(other));
}

char* QProcessEnvironment_ToStringList(QtObjectPtr ptr){
	return static_cast<QProcessEnvironment*>(ptr)->toStringList().join("|").toUtf8().data();
}

char* QProcessEnvironment_Value(QtObjectPtr ptr, char* name, char* defaultValue){
	return static_cast<QProcessEnvironment*>(ptr)->value(QString(name), QString(defaultValue)).toUtf8().data();
}

void QProcessEnvironment_DestroyQProcessEnvironment(QtObjectPtr ptr){
	static_cast<QProcessEnvironment*>(ptr)->~QProcessEnvironment();
}

