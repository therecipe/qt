#include "qdnslookup.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QHostAddress>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDnsLookup>
#include "_cgo_export.h"

class MyQDnsLookup: public QDnsLookup {
public:
void Signal_Finished(){callbackQDnsLookupFinished(this->objectName().toUtf8().data());};
void Signal_NameChanged(const QString & name){callbackQDnsLookupNameChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_TypeChanged(QDnsLookup::Type ty){callbackQDnsLookupTypeChanged(this->objectName().toUtf8().data(), ty);};
};

QtObjectPtr QDnsLookup_NewQDnsLookup3(int ty, char* name, QtObjectPtr nameserver, QtObjectPtr parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
}

int QDnsLookup_Error(QtObjectPtr ptr){
	return static_cast<QDnsLookup*>(ptr)->error();
}

char* QDnsLookup_ErrorString(QtObjectPtr ptr){
	return static_cast<QDnsLookup*>(ptr)->errorString().toUtf8().data();
}

char* QDnsLookup_Name(QtObjectPtr ptr){
	return static_cast<QDnsLookup*>(ptr)->name().toUtf8().data();
}

void QDnsLookup_SetName(QtObjectPtr ptr, char* name){
	static_cast<QDnsLookup*>(ptr)->setName(QString(name));
}

void QDnsLookup_SetNameserver(QtObjectPtr ptr, QtObjectPtr nameserver){
	static_cast<QDnsLookup*>(ptr)->setNameserver(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetType(QtObjectPtr ptr, int v){
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(v));
}

int QDnsLookup_Type(QtObjectPtr ptr){
	return static_cast<QDnsLookup*>(ptr)->type();
}

QtObjectPtr QDnsLookup_NewQDnsLookup(QtObjectPtr parent){
	return new QDnsLookup(static_cast<QObject*>(parent));
}

QtObjectPtr QDnsLookup_NewQDnsLookup2(int ty, char* name, QtObjectPtr parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), static_cast<QObject*>(parent));
}

void QDnsLookup_Abort(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
}

void QDnsLookup_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

void QDnsLookup_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

int QDnsLookup_IsFinished(QtObjectPtr ptr){
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void QDnsLookup_Lookup(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

void QDnsLookup_ConnectNameChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_DisconnectNameChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_ConnectTypeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_DisconnectTypeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_DestroyQDnsLookup(QtObjectPtr ptr){
	static_cast<QDnsLookup*>(ptr)->~QDnsLookup();
}

