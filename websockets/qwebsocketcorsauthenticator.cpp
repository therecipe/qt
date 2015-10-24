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

QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator3(QtObjectPtr other){
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator(char* origin){
	return new QWebSocketCorsAuthenticator(QString(origin));
}

QtObjectPtr QWebSocketCorsAuthenticator_NewQWebSocketCorsAuthenticator2(QtObjectPtr other){
	return new QWebSocketCorsAuthenticator(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

int QWebSocketCorsAuthenticator_Allowed(QtObjectPtr ptr){
	return static_cast<QWebSocketCorsAuthenticator*>(ptr)->allowed();
}

char* QWebSocketCorsAuthenticator_Origin(QtObjectPtr ptr){
	return static_cast<QWebSocketCorsAuthenticator*>(ptr)->origin().toUtf8().data();
}

void QWebSocketCorsAuthenticator_SetAllowed(QtObjectPtr ptr, int allowed){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->setAllowed(allowed != 0);
}

void QWebSocketCorsAuthenticator_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->swap(*static_cast<QWebSocketCorsAuthenticator*>(other));
}

void QWebSocketCorsAuthenticator_DestroyQWebSocketCorsAuthenticator(QtObjectPtr ptr){
	static_cast<QWebSocketCorsAuthenticator*>(ptr)->~QWebSocketCorsAuthenticator();
}

