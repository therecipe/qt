#include "qhostaddress.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QHostAddress>
#include "_cgo_export.h"

class MyQHostAddress: public QHostAddress {
public:
};

void* QHostAddress_NewQHostAddress(){
	return new QHostAddress();
}

void* QHostAddress_NewQHostAddress9(int address){
	return new QHostAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

void* QHostAddress_NewQHostAddress8(void* address){
	return new QHostAddress(*static_cast<QHostAddress*>(address));
}

void* QHostAddress_NewQHostAddress7(char* address){
	return new QHostAddress(QString(address));
}

void QHostAddress_Clear(void* ptr){
	static_cast<QHostAddress*>(ptr)->clear();
}

int QHostAddress_IsInSubnet(void* ptr, void* subnet, int netmask){
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
}

int QHostAddress_IsLoopback(void* ptr){
	return static_cast<QHostAddress*>(ptr)->isLoopback();
}

int QHostAddress_IsNull(void* ptr){
	return static_cast<QHostAddress*>(ptr)->isNull();
}

int QHostAddress_Protocol(void* ptr){
	return static_cast<QHostAddress*>(ptr)->protocol();
}

char* QHostAddress_ScopeId(void* ptr){
	return static_cast<QHostAddress*>(ptr)->scopeId().toUtf8().data();
}

int QHostAddress_SetAddress5(void* ptr, char* address){
	return static_cast<QHostAddress*>(ptr)->setAddress(QString(address));
}

void QHostAddress_SetScopeId(void* ptr, char* id){
	static_cast<QHostAddress*>(ptr)->setScopeId(QString(id));
}

char* QHostAddress_ToString(void* ptr){
	return static_cast<QHostAddress*>(ptr)->toString().toUtf8().data();
}

void QHostAddress_DestroyQHostAddress(void* ptr){
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

