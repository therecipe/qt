#include "qwebsocketcorsauthenticator.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWebSocket>
#include <QWebSocketCorsAuthenticator>
#include "_cgo_export.h"

class MyQWebSocketCorsAuthenticator: public QWebSocketCorsAuthenticator {
public:
};

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(void* other){
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(char* origin){
	return new QWebSocketCorsAuthenticator(QString(origin));
}

void* QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(void* other){
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

int QWebSocketCorsAuthenticator_Allowed(void* ptr){
	return static_cast<QWebSocketCorsAuthenticator*>(ptr)->allowed();
}

char* QWebSocketCorsAuthenticator_Origin(void* ptr){
	return static_cast<QWebSocketCorsAuthenticator*>(ptr)->origin().toUtf8().data();
}

void QWebSocketCorsAuthenticator_SetAllowed(void* ptr, int allowed){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->setAllowed(allowed != 0);
}

void QWebSocketCorsAuthenticator_Swap(void* ptr, void* other){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->swap(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(void* ptr){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->~QWebSocketCorsAuthenticator();
}

