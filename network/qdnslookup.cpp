#include "qdnslookup.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHostAddress>
#include <QMetaObject>
#include <QDnsLookup>
#include "_cgo_export.h"

class MyQDnsLookup: public QDnsLookup {
public:
void Signal_Finished(){callbackQDnsLookupFinished(this->objectName().toUtf8().data());};
void Signal_NameChanged(const QString & name){callbackQDnsLookupNameChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_TypeChanged(QDnsLookup::Type ty){callbackQDnsLookupTypeChanged(this->objectName().toUtf8().data(), ty);};
};

void* QDnsLookup_NewQDnsLookup3(int ty, char* name, void* nameserver, void* parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), *static_cast<QHostAddress*>(nameserver), static_cast<QObject*>(parent));
}

int QDnsLookup_Error(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->error();
}

char* QDnsLookup_ErrorString(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->errorString().toUtf8().data();
}

char* QDnsLookup_Name(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->name().toUtf8().data();
}

void QDnsLookup_SetName(void* ptr, char* name){
	static_cast<QDnsLookup*>(ptr)->setName(QString(name));
}

void QDnsLookup_SetNameserver(void* ptr, void* nameserver){
	static_cast<QDnsLookup*>(ptr)->setNameserver(*static_cast<QHostAddress*>(nameserver));
}

void QDnsLookup_SetType(void* ptr, int v){
	static_cast<QDnsLookup*>(ptr)->setType(static_cast<QDnsLookup::Type>(v));
}

int QDnsLookup_Type(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->type();
}

void* QDnsLookup_NewQDnsLookup(void* parent){
	return new QDnsLookup(static_cast<QObject*>(parent));
}

void* QDnsLookup_NewQDnsLookup2(int ty, char* name, void* parent){
	return new QDnsLookup(static_cast<QDnsLookup::Type>(ty), QString(name), static_cast<QObject*>(parent));
}

void QDnsLookup_Abort(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "abort");
}

void QDnsLookup_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

void QDnsLookup_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)()>(&QDnsLookup::finished), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)()>(&MyQDnsLookup::Signal_Finished));;
}

int QDnsLookup_IsFinished(void* ptr){
	return static_cast<QDnsLookup*>(ptr)->isFinished();
}

void QDnsLookup_Lookup(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDnsLookup*>(ptr), "lookup");
}

void QDnsLookup_ConnectNameChanged(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_DisconnectNameChanged(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(const QString &)>(&QDnsLookup::nameChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(const QString &)>(&MyQDnsLookup::Signal_NameChanged));;
}

void QDnsLookup_ConnectTypeChanged(void* ptr){
	QObject::connect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_DisconnectTypeChanged(void* ptr){
	QObject::disconnect(static_cast<QDnsLookup*>(ptr), static_cast<void (QDnsLookup::*)(QDnsLookup::Type)>(&QDnsLookup::typeChanged), static_cast<MyQDnsLookup*>(ptr), static_cast<void (MyQDnsLookup::*)(QDnsLookup::Type)>(&MyQDnsLookup::Signal_TypeChanged));;
}

void QDnsLookup_DestroyQDnsLookup(void* ptr){
	static_cast<QDnsLookup*>(ptr)->~QDnsLookup();
}

