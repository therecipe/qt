#include "qhostinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHostInfo>
#include "_cgo_export.h"

class MyQHostInfo: public QHostInfo {
public:
};

void* QHostInfo_NewQHostInfo2(void* other){
	return new QHostInfo(*static_cast<QHostInfo*>(other));
}

void* QHostInfo_NewQHostInfo(int id){
	return new QHostInfo(id);
}

void QHostInfo_QHostInfo_AbortHostLookup(int id){
	QHostInfo::abortHostLookup(id);
}

int QHostInfo_Error(void* ptr){
	return static_cast<QHostInfo*>(ptr)->error();
}

char* QHostInfo_ErrorString(void* ptr){
	return static_cast<QHostInfo*>(ptr)->errorString().toUtf8().data();
}

char* QHostInfo_HostName(void* ptr){
	return static_cast<QHostInfo*>(ptr)->hostName().toUtf8().data();
}

int QHostInfo_QHostInfo_LookupHost(char* name, void* receiver, char* member){
	return QHostInfo::lookupHost(QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

int QHostInfo_LookupId(void* ptr){
	return static_cast<QHostInfo*>(ptr)->lookupId();
}

void QHostInfo_SetError(void* ptr, int error){
	static_cast<QHostInfo*>(ptr)->setError(static_cast<QHostInfo::HostInfoError>(error));
}

void QHostInfo_SetErrorString(void* ptr, char* str){
	static_cast<QHostInfo*>(ptr)->setErrorString(QString(str));
}

void QHostInfo_SetHostName(void* ptr, char* hostName){
	static_cast<QHostInfo*>(ptr)->setHostName(QString(hostName));
}

void QHostInfo_SetLookupId(void* ptr, int id){
	static_cast<QHostInfo*>(ptr)->setLookupId(id);
}

void QHostInfo_DestroyQHostInfo(void* ptr){
	static_cast<QHostInfo*>(ptr)->~QHostInfo();
}

char* QHostInfo_QHostInfo_LocalHostName(){
	return QHostInfo::localHostName().toUtf8().data();
}

char* QHostInfo_QHostInfo_LocalDomainName(){
	return QHostInfo::localDomainName().toUtf8().data();
}

