#include "qwebchannel.h"
#include <QWebChannelAbstractTransport>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWebChannel>
#include "_cgo_export.h"

class MyQWebChannel: public QWebChannel {
public:
void Signal_BlockUpdatesChanged(bool block){callbackQWebChannelBlockUpdatesChanged(this->objectName().toUtf8().data(), block);};
};

int QWebChannel_BlockUpdates(QtObjectPtr ptr){
	return static_cast<QWebChannel*>(ptr)->blockUpdates();
}

void QWebChannel_SetBlockUpdates(QtObjectPtr ptr, int block){
	static_cast<QWebChannel*>(ptr)->setBlockUpdates(block != 0);
}

QtObjectPtr QWebChannel_NewQWebChannel(QtObjectPtr parent){
	return new QWebChannel(static_cast<QObject*>(parent));
}

void QWebChannel_ConnectBlockUpdatesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));;
}

void QWebChannel_DisconnectBlockUpdatesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QWebChannel*>(ptr), static_cast<void (QWebChannel::*)(bool)>(&QWebChannel::blockUpdatesChanged), static_cast<MyQWebChannel*>(ptr), static_cast<void (MyQWebChannel::*)(bool)>(&MyQWebChannel::Signal_BlockUpdatesChanged));;
}

void QWebChannel_ConnectTo(QtObjectPtr ptr, QtObjectPtr transport){
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "connectTo", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_DeregisterObject(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QWebChannel*>(ptr)->deregisterObject(static_cast<QObject*>(object));
}

void QWebChannel_DisconnectFrom(QtObjectPtr ptr, QtObjectPtr transport){
	QMetaObject::invokeMethod(static_cast<QWebChannel*>(ptr), "disconnectFrom", Q_ARG(QWebChannelAbstractTransport*, static_cast<QWebChannelAbstractTransport*>(transport)));
}

void QWebChannel_RegisterObject(QtObjectPtr ptr, char* id, QtObjectPtr object){
	static_cast<QWebChannel*>(ptr)->registerObject(QString(id), static_cast<QObject*>(object));
}

void QWebChannel_DestroyQWebChannel(QtObjectPtr ptr){
	static_cast<QWebChannel*>(ptr)->~QWebChannel();
}

