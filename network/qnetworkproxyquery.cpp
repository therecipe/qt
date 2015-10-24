#include "qnetworkproxyquery.h"
#include <QModelIndex>
#include <QNetworkProxy>
#include <QNetworkConfiguration>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QNetworkProxyQuery>
#include "_cgo_export.h"

class MyQNetworkProxyQuery: public QNetworkProxyQuery {
public:
};

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery(){
	return new QNetworkProxyQuery();
}

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery7(QtObjectPtr networkConfiguration, char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery6(QtObjectPtr networkConfiguration, char* requestUrl, int queryType){
	return new QNetworkProxyQuery(*static_cast<QNetworkConfiguration*>(networkConfiguration), QUrl(QString(requestUrl)), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery5(QtObjectPtr other){
	return new QNetworkProxyQuery(*static_cast<QNetworkProxyQuery*>(other));
}

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery3(char* hostname, int port, char* protocolTag, int queryType){
	return new QNetworkProxyQuery(QString(hostname), port, QString(protocolTag), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

QtObjectPtr QNetworkProxyQuery_NewQNetworkProxyQuery2(char* requestUrl, int queryType){
	return new QNetworkProxyQuery(QUrl(QString(requestUrl)), static_cast<QNetworkProxyQuery::QueryType>(queryType));
}

int QNetworkProxyQuery_LocalPort(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->localPort();
}

char* QNetworkProxyQuery_PeerHostName(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerHostName().toUtf8().data();
}

int QNetworkProxyQuery_PeerPort(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->peerPort();
}

char* QNetworkProxyQuery_ProtocolTag(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->protocolTag().toUtf8().data();
}

int QNetworkProxyQuery_QueryType(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->queryType();
}

void QNetworkProxyQuery_SetLocalPort(QtObjectPtr ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setLocalPort(port);
}

void QNetworkProxyQuery_SetNetworkConfiguration(QtObjectPtr ptr, QtObjectPtr networkConfiguration){
	static_cast<QNetworkProxyQuery*>(ptr)->setNetworkConfiguration(*static_cast<QNetworkConfiguration*>(networkConfiguration));
}

void QNetworkProxyQuery_SetPeerHostName(QtObjectPtr ptr, char* hostname){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerHostName(QString(hostname));
}

void QNetworkProxyQuery_SetPeerPort(QtObjectPtr ptr, int port){
	static_cast<QNetworkProxyQuery*>(ptr)->setPeerPort(port);
}

void QNetworkProxyQuery_SetProtocolTag(QtObjectPtr ptr, char* protocolTag){
	static_cast<QNetworkProxyQuery*>(ptr)->setProtocolTag(QString(protocolTag));
}

void QNetworkProxyQuery_SetQueryType(QtObjectPtr ptr, int ty){
	static_cast<QNetworkProxyQuery*>(ptr)->setQueryType(static_cast<QNetworkProxyQuery::QueryType>(ty));
}

void QNetworkProxyQuery_SetUrl(QtObjectPtr ptr, char* url){
	static_cast<QNetworkProxyQuery*>(ptr)->setUrl(QUrl(QString(url)));
}

void QNetworkProxyQuery_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkProxyQuery*>(ptr)->swap(*static_cast<QNetworkProxyQuery*>(other));
}

char* QNetworkProxyQuery_Url(QtObjectPtr ptr){
	return static_cast<QNetworkProxyQuery*>(ptr)->url().toString().toUtf8().data();
}

void QNetworkProxyQuery_DestroyQNetworkProxyQuery(QtObjectPtr ptr){
	static_cast<QNetworkProxyQuery*>(ptr)->~QNetworkProxyQuery();
}

