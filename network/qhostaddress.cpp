#include "qhostaddress.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHostAddress>
#include "_cgo_export.h"

class MyQHostAddress: public QHostAddress {
public:
};

QtObjectPtr QHostAddress_NewQHostAddress(){
	return new QHostAddress();
}

QtObjectPtr QHostAddress_NewQHostAddress9(int address){
	return new QHostAddress(static_cast<QHostAddress::SpecialAddress>(address));
}

QtObjectPtr QHostAddress_NewQHostAddress8(QtObjectPtr address){
	return new QHostAddress(*static_cast<QHostAddress*>(address));
}

QtObjectPtr QHostAddress_NewQHostAddress7(char* address){
	return new QHostAddress(QString(address));
}

void QHostAddress_Clear(QtObjectPtr ptr){
	static_cast<QHostAddress*>(ptr)->clear();
}

int QHostAddress_IsInSubnet(QtObjectPtr ptr, QtObjectPtr subnet, int netmask){
	return static_cast<QHostAddress*>(ptr)->isInSubnet(*static_cast<QHostAddress*>(subnet), netmask);
}

int QHostAddress_IsLoopback(QtObjectPtr ptr){
	return static_cast<QHostAddress*>(ptr)->isLoopback();
}

int QHostAddress_IsNull(QtObjectPtr ptr){
	return static_cast<QHostAddress*>(ptr)->isNull();
}

int QHostAddress_Protocol(QtObjectPtr ptr){
	return static_cast<QHostAddress*>(ptr)->protocol();
}

char* QHostAddress_ScopeId(QtObjectPtr ptr){
	return static_cast<QHostAddress*>(ptr)->scopeId().toUtf8().data();
}

int QHostAddress_SetAddress5(QtObjectPtr ptr, char* address){
	return static_cast<QHostAddress*>(ptr)->setAddress(QString(address));
}

void QHostAddress_SetScopeId(QtObjectPtr ptr, char* id){
	static_cast<QHostAddress*>(ptr)->setScopeId(QString(id));
}

char* QHostAddress_ToString(QtObjectPtr ptr){
	return static_cast<QHostAddress*>(ptr)->toString().toUtf8().data();
}

void QHostAddress_DestroyQHostAddress(QtObjectPtr ptr){
	static_cast<QHostAddress*>(ptr)->~QHostAddress();
}

