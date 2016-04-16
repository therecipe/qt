#define protected public

#include "moc.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

class QmlBridge: public QObject {
Q_OBJECT
public:
	QmlBridge(QObject *parent) : QObject(parent) {};
	void Signal_SendToQml(QString data) { callbackQmlBridgeSendToQml(this, this->objectName().toUtf8().data(), data.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridgeTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQmlBridgeChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQmlBridgeCustomEvent(this, this->objectName().toUtf8().data(), event); };
signals:
	void sendToQml(QString data);
public slots:
	void sendToGo(QString data) { callbackQmlBridgeSendToGo(this, this->objectName().toUtf8().data(), data.toUtf8().data()); };
};

void QmlBridge_ConnectSendToQml(void* ptr){
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SendToQml));;
}

void QmlBridge_DisconnectSendToQml(void* ptr){
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString)>(&QmlBridge::Signal_SendToQml));;
}

void QmlBridge_SendToQml(void* ptr, char* data){
	static_cast<QmlBridge*>(ptr)->sendToQml(QString(data));
}

void QmlBridge_SendToGo(void* ptr, char* data){
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "sendToGo", Q_ARG(QString, QString(data)));
}

void* QmlBridge_NewQmlBridge(void* parent){
	return new QmlBridge(static_cast<QObject*>(parent));
}

void QmlBridge_DestroyQmlBridge(void* ptr){
	static_cast<QObject*>(ptr)->~QObject();
}

void QmlBridge_TimerEvent(void* ptr, void* event){
	static_cast<QmlBridge*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QmlBridge_TimerEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QmlBridge_ChildEvent(void* ptr, void* event){
	static_cast<QmlBridge*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge_ChildEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QmlBridge_CustomEvent(void* ptr, void* event){
	static_cast<QmlBridge*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_CustomEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

#include "moc_moc.h"