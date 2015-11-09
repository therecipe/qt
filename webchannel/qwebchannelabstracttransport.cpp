#include "qwebchannelabstracttransport.h"
#include <QModelIndex>
#include <QWebChannel>
#include <QMetaObject>
#include <QJsonObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QWebChannelAbstractTransport>
#include "_cgo_export.h"

class MyQWebChannelAbstractTransport: public QWebChannelAbstractTransport {
public:
void Signal_MessageReceived(const QJsonObject & message, QWebChannelAbstractTransport * transport){callbackQWebChannelAbstractTransportMessageReceived(this->objectName().toUtf8().data(), new QJsonObject(message), transport);};
};

void QWebChannelAbstractTransport_ConnectMessageReceived(void* ptr){
	QObject::connect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_DisconnectMessageReceived(void* ptr){
	QObject::disconnect(static_cast<QWebChannelAbstractTransport*>(ptr), static_cast<void (QWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&QWebChannelAbstractTransport::messageReceived), static_cast<MyQWebChannelAbstractTransport*>(ptr), static_cast<void (MyQWebChannelAbstractTransport::*)(const QJsonObject &, QWebChannelAbstractTransport *)>(&MyQWebChannelAbstractTransport::Signal_MessageReceived));;
}

void QWebChannelAbstractTransport_SendMessage(void* ptr, void* message){
	QMetaObject::invokeMethod(static_cast<QWebChannelAbstractTransport*>(ptr), "sendMessage", Q_ARG(QJsonObject, *static_cast<QJsonObject*>(message)));
}

void QWebChannelAbstractTransport_DestroyQWebChannelAbstractTransport(void* ptr){
	static_cast<QWebChannelAbstractTransport*>(ptr)->~QWebChannelAbstractTransport();
}

