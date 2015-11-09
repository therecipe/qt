#include "qnetworkproxyquery.h"
#include <QNetworkProxy>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkConfiguration>
#include <QNetworkProxyQuery>
#include "_cgo_export.h"

class MyQNetworkProxyQuery: public QNetworkProxyQuery {
public:
};

void* QNetworkProxyQuery_NewQNetworkProxyQuery(){
	return new QNetworkProxyQuery();
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery7(void* networkConfiguration, char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery6(void* networkConfiguration, void* requestUrl, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), *static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery5(void* other){
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

void* QNetworkProxyQuery_NewQNetworkProxyQuery2(void* requestUrl, int queryType){
	return new QNetworkProxyQuery(*static_cast<QUrl*>(requestUrl), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

int QNetworkProxyQuery_LocalPort(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

char* QNetworkProxyQuery_PeerHostName(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8().data();
}

int QNetworkProxyQuery_PeerPort(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

char* QNetworkProxyQuery_ProtocolTag(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8().data();
}

int QNetworkProxyQuery_QueryType(void* ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->queryType();
}

void QNetworkProxyQuery_SetLocalPort(void* ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setLocalPort(port);
}

void QNetworkProxyQuery_SetNetworkConfiguration(void* ptr, void* networkConfiguration){
	static_cast<QNetworkProxyQuery*>(ptr)->setNetworkConfiguration(*static_cast<QNetworkConfiguration*>(networkConfiguration));
}

void QNetworkProxyQuery_SetPeerHostName(void* ptr, char* hostname){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerHostName(QString(hostname));
}

void QNetworkProxyQuery_SetPeerPort(void* ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerPort(port);
}

void QNetworkProxyQuery_SetProtocolTag(void* ptr, char* protocolTag){
	static_cast<QNetworkProxyQuery*>(ptr)->setProtocolTag(QString(protocolTag));
}

void QNetworkProxyQuery_SetQueryType(void* ptr, int ty){
	static_cast<QNetworkProxyQuery*>(ptr)->setQueryType(static_cast<QNetworkProxyQuery::QueryType>(ty));
}

void QNetworkProxyQuery_SetUrl(void* ptr, void* url){
	static_cast<QNetworkProxyQuery*>(ptr)->setUrl(*static_cast<QUrl*>(url));
}

void QNetworkProxyQuery_Swap(void* ptr, void* other){
	static_cast<QNetworkProxyQuery*>(ptr)->swap(*static_cast<QNetworkProxyQuery*>(other));
}

void QNetworkProxyQuery_DestroyQNetworkProxyQuery(void* ptr){
	static_cast<QNetworkProxyQuery*>(ptr)->~QNetworkProxyQuery();
}

