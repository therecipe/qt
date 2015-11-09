#include "qwebsocketserver.h"
#include <QVariant>
#include <QUrl>
#include <QAbstractSocket>
#include <QObject>
#include <QWebSocketCorsAuthenticator>
#include <QString>
#include <QModelIndex>
#include <QSslConfiguration>
#include <QNetworkProxy>
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

void* QWebSocketServer_NewQWebSocketServer(char* serverName, int secureMode, void* parent){
	return new QWebSocketServer(QString(serverName), static_cast<QWebSocketServer::SslMode>(secureMode), static_cast<QObject*>(parent));
}

void QWebSocketServer_ConnectAcceptError(void* ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));;
}

void QWebSocketServer_DisconnectAcceptError(void* ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QAbstractSocket::SocketError)>(&QWebSocketServer::acceptError), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QAbstractSocket::SocketError)>(&MyQWebSocketServer::Signal_AcceptError));;
}

void QWebSocketServer_Close(void* ptr){
	static_cast<QWebSocketServer*>(ptr)->close();
}

void QWebSocketServer_ConnectClosed(void* ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));;
}

void QWebSocketServer_DisconnectClosed(void* ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::closed), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_Closed));;
}

char* QWebSocketServer_ErrorString(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->errorString().toUtf8().data();
}

int QWebSocketServer_HasPendingConnections(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->hasPendingConnections();
}

int QWebSocketServer_IsListening(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->isListening();
}

int QWebSocketServer_MaxPendingConnections(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->maxPendingConnections();
}

void QWebSocketServer_ConnectNewConnection(void* ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));;
}

void QWebSocketServer_DisconnectNewConnection(void* ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)()>(&QWebSocketServer::newConnection), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)()>(&MyQWebSocketServer::Signal_NewConnection));;
}

void* QWebSocketServer_NextPendingConnection(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->nextPendingConnection();
}

void QWebSocketServer_ConnectOriginAuthenticationRequired(void* ptr){
	QObject::connect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));;
}

void QWebSocketServer_DisconnectOriginAuthenticationRequired(void* ptr){
	QObject::disconnect(static_cast<QWebSocketServer*>(ptr), static_cast<void (QWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&QWebSocketServer::originAuthenticationRequired), static_cast<MyQWebSocketServer*>(ptr), static_cast<void (MyQWebSocketServer::*)(QWebSocketCorsAuthenticator *)>(&MyQWebSocketServer::Signal_OriginAuthenticationRequired));;
}

void QWebSocketServer_PauseAccepting(void* ptr){
	static_cast<QWebSocketServer*>(ptr)->pauseAccepting();
}

void QWebSocketServer_ResumeAccepting(void* ptr){
	static_cast<QWebSocketServer*>(ptr)->resumeAccepting();
}

int QWebSocketServer_SecureMode(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->secureMode();
}

char* QWebSocketServer_ServerName(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->serverName().toUtf8().data();
}

void QWebSocketServer_SetMaxPendingConnections(void* ptr, int numConnections){
	static_cast<QWebSocketServer*>(ptr)->setMaxPendingConnections(numConnections);
}

void QWebSocketServer_SetProxy(void* ptr, void* networkProxy){
	static_cast<QWebSocketServer*>(ptr)->setProxy(*static_cast<QNetworkProxy*>(networkProxy));
}

void QWebSocketServer_SetServerName(void* ptr, char* serverName){
	static_cast<QWebSocketServer*>(ptr)->setServerName(QString(serverName));
}

int QWebSocketServer_SetSocketDescriptor(void* ptr, int socketDescriptor){
	return static_cast<QWebSocketServer*>(ptr)->setSocketDescriptor(socketDescriptor);
}

void QWebSocketServer_SetSslConfiguration(void* ptr, void* sslConfiguration){
	static_cast<QWebSocketServer*>(ptr)->setSslConfiguration(*static_cast<QSslConfiguration*>(sslConfiguration));
}

int QWebSocketServer_SocketDescriptor(void* ptr){
	return static_cast<QWebSocketServer*>(ptr)->socketDescriptor();
}

void QWebSocketServer_DestroyQWebSocketServer(void* ptr){
	static_cast<QWebSocketServer*>(ptr)->~QWebSocketServer();
}

