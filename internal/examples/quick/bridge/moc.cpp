#define protected public

#include "moc.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
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
	void Signal_SendToQml(QString source, QString action, QString data) { callbackQmlBridgeSendToQml(this, this->objectName().toUtf8().data(), source.toUtf8().data(), action.toUtf8().data(), data.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQmlBridgeTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQmlBridgeChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQmlBridgeCustomEvent(this, this->objectName().toUtf8().data(), event); };
signals:
	void sendToQml(QString source, QString action, QString data);
public slots:
	void sendToGo(QString source, QString action, QString data) { callbackQmlBridgeSendToGo(this, this->objectName().toUtf8().data(), source.toUtf8().data(), action.toUtf8().data(), data.toUtf8().data()); };
	void registerToGo(QObject* object) { callbackQmlBridgeRegisterToGo(this, this->objectName().toUtf8().data(), object); };
	void deregisterToGo(QString objectName) { callbackQmlBridgeDeregisterToGo(this, this->objectName().toUtf8().data(), objectName.toUtf8().data()); };
};

void QmlBridge_ConnectSendToQml(void* ptr){
	QObject::connect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::Signal_SendToQml));;
}

void QmlBridge_DisconnectSendToQml(void* ptr){
	QObject::disconnect(static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::sendToQml), static_cast<QmlBridge*>(ptr), static_cast<void (QmlBridge::*)(QString, QString, QString)>(&QmlBridge::Signal_SendToQml));;
}

void QmlBridge_SendToQml(void* ptr, char* source, char* action, char* data){
	static_cast<QmlBridge*>(ptr)->sendToQml(QString(source), QString(action), QString(data));
}

void QmlBridge_SendToGo(void* ptr, char* source, char* action, char* data){
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "sendToGo", Q_ARG(QString, QString(source)), Q_ARG(QString, QString(action)), Q_ARG(QString, QString(data)));
}

void QmlBridge_RegisterToGo(void* ptr, void* object){
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "registerToGo", Q_ARG(QObject*, static_cast<QObject*>(object)));
}

void QmlBridge_DeregisterToGo(void* ptr, char* objectName){
	QMetaObject::invokeMethod(static_cast<QmlBridge*>(ptr), "deregisterToGo", Q_ARG(QString, QString(objectName)));
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

void QmlBridge_ConnectNotify(void* ptr, void* sign){
	static_cast<QmlBridge*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_ConnectNotifyDefault(void* ptr, void* sign){
	static_cast<QObject*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_CustomEvent(void* ptr, void* event){
	static_cast<QmlBridge*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_CustomEventDefault(void* ptr, void* event){
	static_cast<QObject*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QmlBridge_DisconnectNotify(void* ptr, void* sign){
	static_cast<QmlBridge*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QmlBridge_DisconnectNotifyDefault(void* ptr, void* sign){
	static_cast<QObject*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

#include "moc_moc.h"