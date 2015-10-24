#include "qwebsocketserver.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAbstractSocket>
#include <QWebSocketCorsAuthenticator>
#include <QString>
#include <QVariant>
#include <QNetworkProxy>
#include <QSslConfiguration>
#include <QWebSocket>
#include <QWebSocketServer>
#include "_cgo_export.h"

class MyQWebSocketServer: public QWebSocketServer {
public:
void Signal_AcceptError(QAbstractSocket::SocketError socketError){callbackQWebSocketServerAcceptError(this->objectName().toUtf8().data(), socketError);};
void Signal_Closed(){callbackQWebSocketServerClosed(this->objectName().toUtf8().data());};
void Signal_NewConnection(){callbackQWebSocketServerNewConnection(this->objectName().toUtf8().data());};
void Signal_OriginAuthenticationRequired(QWebSocketCorsAuthenticator * authenticator){callbackQWebSocketServerOriginAuthenticationRequired(this->objectName().toUtf8().data(), authenticator);};
};

QtObjectPtr QWebSocketServer_NewQWebSocketServer(char* serverName, int secureMode, QtObjectPtr parent){
	return new QWebSocketServer(QString(serverName), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QObject*>(parent));
}

void QWebSocketServer_ConnectAcceptError(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));;
}

void QWebSocketServer_DisconnectAcceptError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));;
}

void QWebSocketServer_Close(QtObjectPtr ptr){
	static_cast<QWebSocketServer*>(ptr)->close();
}

void QWebSocketServer_ConnectClosed(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));;
}

void QWebSocketServer_DisconnectClosed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));;
}

char* QWebSocketServer_ErrorString(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->errorString().toUtf8().data();
}

int QWebSocketServer_HasPendingConnections(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->hasPendingConnections();
}

int QWebSocketServer_IsListening(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->isListening();
}

int QWebSocketServer_MaxPendingConnections(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->maxPendingConnections();
}

void QWebSocketServer_ConnectNewConnection(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));;
}

void QWebSocketServer_DisconnectNewConnection(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));;
}

QtObjectPtr QWebSocketServer_NextPendingConnection(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->nextPendingConnection();
}

void QWebSocketServer_ConnectOriginAuthenticationRequired(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));;
}

void QWebSocketServer_DisconnectOriginAuthenticationRequired(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));;
}

void QWebSocketServer_PauseAccepting(QtObjectPtr ptr){
	static_cast<QWebSocketServer*>(ptr)->pauseAccepting();
}

void QWebSocketServer_ResumeAccepting(QtObjectPtr ptr){
	static_cast<QWebSocketServer*>(ptr)->resumeAccepting();
}

int QWebSocketServer_SecureMode(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->secureMode();
}

char* QWebSocketServer_ServerName(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->serverName().toUtf8().data();
}

char* QWebSocketServer_ServerUrl(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->serverUrl().toString().toUtf8().data();
}

void QWebSocketServer_SetMaxPendingConnections(QtObjectPtr ptr, int numConnections){
	static_cast<QWebSocketServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QWebSocketServer_SetProxy(QtObjectPtr ptr, QtObjectPtr networkProxy){
	static_cast<QWebSocketServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocketServer_SetServerName(QtObjectPtr ptr, char* serverName){
	static_cast<QWebSocketServer*>(ptr)->setServerName(QString(serverName));
}

int QWebSocketServer_SetSocketDescriptor(QtObjectPtr ptr, int socketDescriptor){
	return static_cast<QWebSocketServer*>(ptr)->setSocketDescriptor(socketDescriptor);
}

void QWebSocketServer_SetSslConfiguration(QtObjectPtr ptr, QtObjectPtr sslConfiguration){
	static_cast<QWebSocketServer*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

int QWebSocketServer_SocketDescriptor(QtObjectPtr ptr){
	return static_cast<QWebSocketServer*>(ptr)->socketDescriptor();
}

void QWebSocketServer_DestroyQWebSocketServer(QtObjectPtr ptr){
	static_cast<QWebSocketServer*>(ptr)->~QWebSocketServer();
}

