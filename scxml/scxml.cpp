// +build !minimal

#define protected public
#define private public

#include "scxml.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QIODevice>
#include <QJSEngine>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QScxmlCppDataModel>
#include <QScxmlDataModel>
#include <QScxmlEcmaScriptDataModel>
#include <QScxmlError>
#include <QScxmlEvent>
#include <QScxmlEventFilter>
#include <QScxmlNullDataModel>
#include <QScxmlParser>
#include <QScxmlStateMachine>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QXmlStreamReader>

class MyQScxmlCppDataModel: public QScxmlCppDataModel
{
public:
	bool hasScxmlProperty(const QString & name) const { return callbackQScxmlCppDataModel_HasScxmlProperty(const_cast<MyQScxmlCppDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data()) != 0; };
	QVariant scxmlProperty(const QString & name) const { return *static_cast<QVariant*>(callbackQScxmlCppDataModel_ScxmlProperty(const_cast<MyQScxmlCppDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data())); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { return callbackQScxmlCppDataModel_SetScxmlProperty(this, this->objectName().toUtf8().data(), name.toUtf8().data(), const_cast<QVariant*>(&value), context.toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQScxmlCppDataModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScxmlCppDataModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlCppDataModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlCppDataModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScxmlCppDataModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlCppDataModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlCppDataModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlCppDataModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlCppDataModel_MetaObject(const_cast<MyQScxmlCppDataModel*>(this), this->objectName().toUtf8().data())); };
};

int QScxmlCppDataModel_In(void* ptr, char* stateName)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->In(QString(stateName));
}

int QScxmlCppDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

int QScxmlCppDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlCppDataModel_ScxmlEvent(void* ptr)
{
	return const_cast<QScxmlEvent*>(&static_cast<QScxmlCppDataModel*>(ptr)->scxmlEvent());
}

void* QScxmlCppDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlCppDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::scxmlProperty(QString(name)));
}

void QScxmlCppDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

int QScxmlCppDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

int QScxmlCppDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

void QScxmlCppDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlCppDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlCppDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlCppDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlCppDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlCppDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlCppDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlCppDataModel*>(ptr), "deleteLater");
}

void QScxmlCppDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::deleteLater();
}

void QScxmlCppDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlCppDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScxmlCppDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScxmlCppDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::event(static_cast<QEvent*>(e));
}

int QScxmlCppDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScxmlCppDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlCppDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlCppDataModel*>(ptr)->metaObject());
}

void* QScxmlCppDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlCppDataModel*>(ptr)->QScxmlCppDataModel::metaObject());
}

class MyQScxmlDataModel: public QScxmlDataModel
{
public:
	bool hasScxmlProperty(const QString & name) const { return callbackQScxmlDataModel_HasScxmlProperty(const_cast<MyQScxmlDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data()) != 0; };
	QVariant scxmlProperty(const QString & name) const { return *static_cast<QVariant*>(callbackQScxmlDataModel_ScxmlProperty(const_cast<MyQScxmlDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data())); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlDataModel_SetScxmlEvent(this, this->objectName().toUtf8().data(), const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { return callbackQScxmlDataModel_SetScxmlProperty(this, this->objectName().toUtf8().data(), name.toUtf8().data(), const_cast<QVariant*>(&value), context.toUtf8().data()) != 0; };
	void Signal_StateMachineChanged(QScxmlStateMachine * stateMachine) { callbackQScxmlDataModel_StateMachineChanged(this, this->objectName().toUtf8().data(), stateMachine); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlDataModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScxmlDataModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlDataModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScxmlDataModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlDataModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlDataModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlDataModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlDataModel_MetaObject(const_cast<MyQScxmlDataModel*>(this), this->objectName().toUtf8().data())); };
};

int QScxmlDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

void* QScxmlDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void QScxmlDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

int QScxmlDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

void QScxmlDataModel_SetStateMachine(void* ptr, void* stateMachine)
{
	static_cast<QScxmlDataModel*>(ptr)->setStateMachine(static_cast<QScxmlStateMachine*>(stateMachine));
}

void* QScxmlDataModel_StateMachine(void* ptr)
{
	return static_cast<QScxmlDataModel*>(ptr)->stateMachine();
}

void QScxmlDataModel_ConnectStateMachineChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlDataModel*>(ptr), static_cast<void (QScxmlDataModel::*)(QScxmlStateMachine *)>(&QScxmlDataModel::stateMachineChanged), static_cast<MyQScxmlDataModel*>(ptr), static_cast<void (MyQScxmlDataModel::*)(QScxmlStateMachine *)>(&MyQScxmlDataModel::Signal_StateMachineChanged));
}

void QScxmlDataModel_DisconnectStateMachineChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlDataModel*>(ptr), static_cast<void (QScxmlDataModel::*)(QScxmlStateMachine *)>(&QScxmlDataModel::stateMachineChanged), static_cast<MyQScxmlDataModel*>(ptr), static_cast<void (MyQScxmlDataModel::*)(QScxmlStateMachine *)>(&MyQScxmlDataModel::Signal_StateMachineChanged));
}

void QScxmlDataModel_StateMachineChanged(void* ptr, void* stateMachine)
{
	static_cast<QScxmlDataModel*>(ptr)->stateMachineChanged(static_cast<QScxmlStateMachine*>(stateMachine));
}

void QScxmlDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlDataModel*>(ptr), "deleteLater");
}

void QScxmlDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::deleteLater();
}

void QScxmlDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScxmlDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScxmlDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::event(static_cast<QEvent*>(e));
}

int QScxmlDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScxmlDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlDataModel*>(ptr)->metaObject());
}

void* QScxmlDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlDataModel*>(ptr)->QScxmlDataModel::metaObject());
}

class MyQScxmlEcmaScriptDataModel: public QScxmlEcmaScriptDataModel
{
public:
	MyQScxmlEcmaScriptDataModel(QObject *parent) : QScxmlEcmaScriptDataModel(parent) {};
	bool hasScxmlProperty(const QString & name) const { return callbackQScxmlEcmaScriptDataModel_HasScxmlProperty(const_cast<MyQScxmlEcmaScriptDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data()) != 0; };
	QVariant scxmlProperty(const QString & name) const { return *static_cast<QVariant*>(callbackQScxmlEcmaScriptDataModel_ScxmlProperty(const_cast<MyQScxmlEcmaScriptDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data())); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlEcmaScriptDataModel_SetScxmlEvent(this, this->objectName().toUtf8().data(), const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { return callbackQScxmlEcmaScriptDataModel_SetScxmlProperty(this, this->objectName().toUtf8().data(), name.toUtf8().data(), const_cast<QVariant*>(&value), context.toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQScxmlEcmaScriptDataModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScxmlEcmaScriptDataModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlEcmaScriptDataModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlEcmaScriptDataModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScxmlEcmaScriptDataModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlEcmaScriptDataModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlEcmaScriptDataModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlEcmaScriptDataModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlEcmaScriptDataModel_MetaObject(const_cast<MyQScxmlEcmaScriptDataModel*>(this), this->objectName().toUtf8().data())); };
};

void* QScxmlEcmaScriptDataModel_NewQScxmlEcmaScriptDataModel(void* parent)
{
	return new MyQScxmlEcmaScriptDataModel(static_cast<QObject*>(parent));
}

void* QScxmlEcmaScriptDataModel_Engine(void* ptr)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->engine();
}

int QScxmlEcmaScriptDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

int QScxmlEcmaScriptDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlEcmaScriptDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlEcmaScriptDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::scxmlProperty(QString(name)));
}

void QScxmlEcmaScriptDataModel_SetEngine(void* ptr, void* engine)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setEngine(static_cast<QJSEngine*>(engine));
}

void QScxmlEcmaScriptDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlEcmaScriptDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

int QScxmlEcmaScriptDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

int QScxmlEcmaScriptDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

void QScxmlEcmaScriptDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlEcmaScriptDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlEcmaScriptDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlEcmaScriptDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlEcmaScriptDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlEcmaScriptDataModel*>(ptr), "deleteLater");
}

void QScxmlEcmaScriptDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::deleteLater();
}

void QScxmlEcmaScriptDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlEcmaScriptDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScxmlEcmaScriptDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScxmlEcmaScriptDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::event(static_cast<QEvent*>(e));
}

int QScxmlEcmaScriptDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScxmlEcmaScriptDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlEcmaScriptDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->metaObject());
}

void* QScxmlEcmaScriptDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlEcmaScriptDataModel*>(ptr)->QScxmlEcmaScriptDataModel::metaObject());
}

void* QScxmlError_NewQScxmlError()
{
	return new QScxmlError();
}

void* QScxmlError_NewQScxmlError3(void* other)
{
	return new QScxmlError(*static_cast<QScxmlError*>(other));
}

void* QScxmlError_NewQScxmlError2(char* fileName, int line, int column, char* description)
{
	return new QScxmlError(QString(fileName), line, column, QString(description));
}

int QScxmlError_Column(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->column();
}

char* QScxmlError_Description(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->description().toUtf8().data();
}

char* QScxmlError_FileName(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->fileName().toUtf8().data();
}

int QScxmlError_IsValid(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->isValid();
}

int QScxmlError_Line(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->line();
}

char* QScxmlError_ToString(void* ptr)
{
	return static_cast<QScxmlError*>(ptr)->toString().toUtf8().data();
}

void QScxmlError_DestroyQScxmlError(void* ptr)
{
	static_cast<QScxmlError*>(ptr)->~QScxmlError();
}

void* QScxmlEvent_NewQScxmlEvent()
{
	return new QScxmlEvent();
}

void* QScxmlEvent_NewQScxmlEvent2(void* other)
{
	return new QScxmlEvent(*static_cast<QScxmlEvent*>(other));
}

void QScxmlEvent_Clear(void* ptr)
{
	static_cast<QScxmlEvent*>(ptr)->clear();
}

void* QScxmlEvent_Data(void* ptr)
{
	return new QVariant(static_cast<QScxmlEvent*>(ptr)->data());
}

int QScxmlEvent_Delay(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->delay();
}

char* QScxmlEvent_ErrorMessage(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->errorMessage().toUtf8().data();
}

int QScxmlEvent_EventType(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->eventType();
}

char* QScxmlEvent_InvokeId(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->invokeId().toUtf8().data();
}

int QScxmlEvent_IsErrorEvent(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->isErrorEvent();
}

char* QScxmlEvent_Name(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->name().toUtf8().data();
}

char* QScxmlEvent_Origin(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->origin().toUtf8().data();
}

char* QScxmlEvent_OriginType(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->originType().toUtf8().data();
}

char* QScxmlEvent_ScxmlType(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->scxmlType().toUtf8().data();
}

char* QScxmlEvent_SendId(void* ptr)
{
	return static_cast<QScxmlEvent*>(ptr)->sendId().toUtf8().data();
}

void QScxmlEvent_SetData(void* ptr, void* data)
{
	static_cast<QScxmlEvent*>(ptr)->setData(*static_cast<QVariant*>(data));
}

void QScxmlEvent_SetDelay(void* ptr, int delayInMiliSecs)
{
	static_cast<QScxmlEvent*>(ptr)->setDelay(delayInMiliSecs);
}

void QScxmlEvent_SetErrorMessage(void* ptr, char* message)
{
	static_cast<QScxmlEvent*>(ptr)->setErrorMessage(QString(message));
}

void QScxmlEvent_SetEventType(void* ptr, int ty)
{
	static_cast<QScxmlEvent*>(ptr)->setEventType(static_cast<QScxmlEvent::EventType>(ty));
}

void QScxmlEvent_SetInvokeId(void* ptr, char* invokeid)
{
	static_cast<QScxmlEvent*>(ptr)->setInvokeId(QString(invokeid));
}

void QScxmlEvent_SetName(void* ptr, char* name)
{
	static_cast<QScxmlEvent*>(ptr)->setName(QString(name));
}

void QScxmlEvent_SetOrigin(void* ptr, char* origin)
{
	static_cast<QScxmlEvent*>(ptr)->setOrigin(QString(origin));
}

void QScxmlEvent_SetOriginType(void* ptr, char* origintype)
{
	static_cast<QScxmlEvent*>(ptr)->setOriginType(QString(origintype));
}

void QScxmlEvent_SetSendId(void* ptr, char* sendid)
{
	static_cast<QScxmlEvent*>(ptr)->setSendId(QString(sendid));
}

void QScxmlEvent_DestroyQScxmlEvent(void* ptr)
{
	static_cast<QScxmlEvent*>(ptr)->~QScxmlEvent();
}

class MyQScxmlEventFilter: public QScxmlEventFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool handle(QScxmlEvent * event, QScxmlStateMachine * stateMachine) { return callbackQScxmlEventFilter_Handle(this, this->objectNameAbs().toUtf8().data(), event, stateMachine) != 0; };
};

int QScxmlEventFilter_Handle(void* ptr, void* event, void* stateMachine)
{
	return static_cast<QScxmlEventFilter*>(ptr)->handle(static_cast<QScxmlEvent*>(event), static_cast<QScxmlStateMachine*>(stateMachine));
}

void QScxmlEventFilter_DestroyQScxmlEventFilter(void* ptr)
{
	static_cast<QScxmlEventFilter*>(ptr)->~QScxmlEventFilter();
}

char* QScxmlEventFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQScxmlEventFilter*>(static_cast<QScxmlEventFilter*>(ptr))) {
		return static_cast<MyQScxmlEventFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QScxmlEventFilter_BASE").toUtf8().data();
}

void QScxmlEventFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQScxmlEventFilter*>(static_cast<QScxmlEventFilter*>(ptr))) {
		static_cast<MyQScxmlEventFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQScxmlNullDataModel: public QScxmlNullDataModel
{
public:
	MyQScxmlNullDataModel(QObject *parent) : QScxmlNullDataModel(parent) {};
	bool hasScxmlProperty(const QString & name) const { return callbackQScxmlNullDataModel_HasScxmlProperty(const_cast<MyQScxmlNullDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data()) != 0; };
	QVariant scxmlProperty(const QString & name) const { return *static_cast<QVariant*>(callbackQScxmlNullDataModel_ScxmlProperty(const_cast<MyQScxmlNullDataModel*>(this), this->objectName().toUtf8().data(), name.toUtf8().data())); };
	void setScxmlEvent(const QScxmlEvent & event) { callbackQScxmlNullDataModel_SetScxmlEvent(this, this->objectName().toUtf8().data(), const_cast<QScxmlEvent*>(&event)); };
	bool setScxmlProperty(const QString & name, const QVariant & value, const QString & context) { return callbackQScxmlNullDataModel_SetScxmlProperty(this, this->objectName().toUtf8().data(), name.toUtf8().data(), const_cast<QVariant*>(&value), context.toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQScxmlNullDataModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScxmlNullDataModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlNullDataModel_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlNullDataModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScxmlNullDataModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlNullDataModel_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlNullDataModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlNullDataModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlNullDataModel_MetaObject(const_cast<MyQScxmlNullDataModel*>(this), this->objectName().toUtf8().data())); };
};

void* QScxmlNullDataModel_NewQScxmlNullDataModel(void* parent)
{
	return new MyQScxmlNullDataModel(static_cast<QObject*>(parent));
}

int QScxmlNullDataModel_HasScxmlProperty(void* ptr, char* name)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->hasScxmlProperty(QString(name));
}

int QScxmlNullDataModel_HasScxmlPropertyDefault(void* ptr, char* name)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::hasScxmlProperty(QString(name));
}

void* QScxmlNullDataModel_ScxmlProperty(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->scxmlProperty(QString(name)));
}

void* QScxmlNullDataModel_ScxmlPropertyDefault(void* ptr, char* name)
{
	return new QVariant(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::scxmlProperty(QString(name)));
}

void QScxmlNullDataModel_SetScxmlEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

void QScxmlNullDataModel_SetScxmlEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlEvent(*static_cast<QScxmlEvent*>(event));
}

int QScxmlNullDataModel_SetScxmlProperty(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

int QScxmlNullDataModel_SetScxmlPropertyDefault(void* ptr, char* name, void* value, char* context)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::setScxmlProperty(QString(name), *static_cast<QVariant*>(value), QString(context));
}

void QScxmlNullDataModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlNullDataModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlNullDataModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlNullDataModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlNullDataModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlNullDataModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::customEvent(static_cast<QEvent*>(event));
}

void QScxmlNullDataModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlNullDataModel*>(ptr), "deleteLater");
}

void QScxmlNullDataModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::deleteLater();
}

void QScxmlNullDataModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlNullDataModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScxmlNullDataModel_Event(void* ptr, void* e)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScxmlNullDataModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::event(static_cast<QEvent*>(e));
}

int QScxmlNullDataModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScxmlNullDataModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlNullDataModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlNullDataModel*>(ptr)->metaObject());
}

void* QScxmlNullDataModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlNullDataModel*>(ptr)->QScxmlNullDataModel::metaObject());
}

void* QScxmlParser_NewQScxmlParser(void* reader)
{
	return new QScxmlParser(static_cast<QXmlStreamReader*>(reader));
}

void QScxmlParser_AddError(void* ptr, char* msg)
{
	static_cast<QScxmlParser*>(ptr)->addError(QString(msg));
}

char* QScxmlParser_FileName(void* ptr)
{
	return static_cast<QScxmlParser*>(ptr)->fileName().toUtf8().data();
}

void QScxmlParser_InstantiateDataModel(void* ptr, void* stateMachine)
{
	static_cast<QScxmlParser*>(ptr)->instantiateDataModel(static_cast<QScxmlStateMachine*>(stateMachine));
}

void* QScxmlParser_InstantiateStateMachine(void* ptr)
{
	return static_cast<QScxmlParser*>(ptr)->instantiateStateMachine();
}

void QScxmlParser_Parse(void* ptr)
{
	static_cast<QScxmlParser*>(ptr)->parse();
}

int QScxmlParser_QtMode(void* ptr)
{
	return static_cast<QScxmlParser*>(ptr)->qtMode();
}

void QScxmlParser_SetFileName(void* ptr, char* fileName)
{
	static_cast<QScxmlParser*>(ptr)->setFileName(QString(fileName));
}

void QScxmlParser_SetQtMode(void* ptr, int mode)
{
	static_cast<QScxmlParser*>(ptr)->setQtMode(static_cast<QScxmlParser::QtMode>(mode));
}

void QScxmlParser_DestroyQScxmlParser(void* ptr)
{
	static_cast<QScxmlParser*>(ptr)->~QScxmlParser();
}

class MyQScxmlStateMachine: public QScxmlStateMachine
{
public:
	void Signal_DataModelChanged(QScxmlDataModel * model) { callbackQScxmlStateMachine_DataModelChanged(this, this->objectName().toUtf8().data(), model); };
	void Signal_EventOccurred(const QScxmlEvent & event) { callbackQScxmlStateMachine_EventOccurred(this, this->objectName().toUtf8().data(), const_cast<QScxmlEvent*>(&event)); };
	void Signal_ExternalEventOccurred(const QScxmlEvent & event) { callbackQScxmlStateMachine_ExternalEventOccurred(this, this->objectName().toUtf8().data(), const_cast<QScxmlEvent*>(&event)); };
	void Signal_Finished() { callbackQScxmlStateMachine_Finished(this, this->objectName().toUtf8().data()); };
	bool init() { return callbackQScxmlStateMachine_Init(this, this->objectName().toUtf8().data()) != 0; };
	void Signal_InitializedChanged(bool initialized) { callbackQScxmlStateMachine_InitializedChanged(this, this->objectName().toUtf8().data(), initialized); };
	void Signal_Log(const QString & label, const QString & msg) { callbackQScxmlStateMachine_Log(this, this->objectName().toUtf8().data(), label.toUtf8().data(), msg.toUtf8().data()); };
	void Signal_ReachedStableState() { callbackQScxmlStateMachine_ReachedStableState(this, this->objectName().toUtf8().data()); };
	void Signal_RunningChanged(bool running) { callbackQScxmlStateMachine_RunningChanged(this, this->objectName().toUtf8().data(), running); };
	void start() { callbackQScxmlStateMachine_Start(this, this->objectName().toUtf8().data()); };
	void stop() { callbackQScxmlStateMachine_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQScxmlStateMachine_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScxmlStateMachine_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_ConnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScxmlStateMachine_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScxmlStateMachine_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScxmlStateMachine_DisconnectNotify(this, this->objectName().toUtf8().data(), const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScxmlStateMachine_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScxmlStateMachine_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScxmlStateMachine_MetaObject(const_cast<MyQScxmlStateMachine*>(this), this->objectName().toUtf8().data())); };
};

int QScxmlStateMachine_IsInitialized(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInitialized();
}

char* QScxmlStateMachine_ActiveStateNames(void* ptr, int compress)
{
	return static_cast<QScxmlStateMachine*>(ptr)->activeStateNames(compress != 0).join("|").toUtf8().data();
}

void QScxmlStateMachine_CancelDelayedEvent(void* ptr, char* sendId)
{
	static_cast<QScxmlStateMachine*>(ptr)->cancelDelayedEvent(QString(sendId));
}

int QScxmlStateMachine_DataBinding(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->dataBinding();
}

void* QScxmlStateMachine_DataModel(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->dataModel();
}

void QScxmlStateMachine_ConnectDataModelChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlDataModel *)>(&QScxmlStateMachine::dataModelChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlDataModel *)>(&MyQScxmlStateMachine::Signal_DataModelChanged));
}

void QScxmlStateMachine_DisconnectDataModelChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(QScxmlDataModel *)>(&QScxmlStateMachine::dataModelChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(QScxmlDataModel *)>(&MyQScxmlStateMachine::Signal_DataModelChanged));
}

void QScxmlStateMachine_DataModelChanged(void* ptr, void* model)
{
	static_cast<QScxmlStateMachine*>(ptr)->dataModelChanged(static_cast<QScxmlDataModel*>(model));
}

void QScxmlStateMachine_ConnectEventOccurred(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::eventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_EventOccurred));
}

void QScxmlStateMachine_DisconnectEventOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::eventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_EventOccurred));
}

void QScxmlStateMachine_EventOccurred(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->eventOccurred(*static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_ConnectExternalEventOccurred(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::externalEventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_ExternalEventOccurred));
}

void QScxmlStateMachine_DisconnectExternalEventOccurred(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QScxmlEvent &)>(&QScxmlStateMachine::externalEventOccurred), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QScxmlEvent &)>(&MyQScxmlStateMachine::Signal_ExternalEventOccurred));
}

void QScxmlStateMachine_ExternalEventOccurred(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->externalEventOccurred(*static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::finished), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_Finished));
}

void QScxmlStateMachine_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::finished), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_Finished));
}

void QScxmlStateMachine_Finished(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->finished();
}

void* QScxmlStateMachine_QScxmlStateMachine_FromData(void* data, char* fileName)
{
	return QScxmlStateMachine::fromData(static_cast<QIODevice*>(data), QString(fileName));
}

void* QScxmlStateMachine_QScxmlStateMachine_FromFile(char* fileName)
{
	return QScxmlStateMachine::fromFile(QString(fileName));
}

char* QScxmlStateMachine_QScxmlStateMachine_GenerateSessionId(char* prefix)
{
	return QScxmlStateMachine::generateSessionId(QString(prefix)).toUtf8().data();
}

int QScxmlStateMachine_Init(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "init", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

void QScxmlStateMachine_ConnectInitializedChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::initializedChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_InitializedChanged));
}

void QScxmlStateMachine_DisconnectInitializedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::initializedChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_InitializedChanged));
}

void QScxmlStateMachine_InitializedChanged(void* ptr, int initialized)
{
	static_cast<QScxmlStateMachine*>(ptr)->initializedChanged(initialized != 0);
}

int QScxmlStateMachine_IsActive(void* ptr, char* scxmlStateName)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isActive(QString(scxmlStateName));
}

int QScxmlStateMachine_IsDispatchableTarget(void* ptr, char* target)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isDispatchableTarget(QString(target));
}

int QScxmlStateMachine_IsInvoked(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isInvoked();
}

int QScxmlStateMachine_IsRunning(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->isRunning();
}

void QScxmlStateMachine_ConnectLog(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_DisconnectLog(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(const QString &, const QString &)>(&QScxmlStateMachine::log), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(const QString &, const QString &)>(&MyQScxmlStateMachine::Signal_Log));
}

void QScxmlStateMachine_Log(void* ptr, char* label, char* msg)
{
	static_cast<QScxmlStateMachine*>(ptr)->log(QString(label), QString(msg));
}

char* QScxmlStateMachine_Name(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->name().toUtf8().data();
}

void QScxmlStateMachine_ConnectReachedStableState(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::reachedStableState), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_ReachedStableState));
}

void QScxmlStateMachine_DisconnectReachedStableState(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)()>(&QScxmlStateMachine::reachedStableState), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)()>(&MyQScxmlStateMachine::Signal_ReachedStableState));
}

void QScxmlStateMachine_ReachedStableState(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->reachedStableState();
}

void QScxmlStateMachine_ConnectRunningChanged(void* ptr)
{
	QObject::connect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::runningChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_RunningChanged));
}

void QScxmlStateMachine_DisconnectRunningChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScxmlStateMachine*>(ptr), static_cast<void (QScxmlStateMachine::*)(bool)>(&QScxmlStateMachine::runningChanged), static_cast<MyQScxmlStateMachine*>(ptr), static_cast<void (MyQScxmlStateMachine::*)(bool)>(&MyQScxmlStateMachine::Signal_RunningChanged));
}

void QScxmlStateMachine_RunningChanged(void* ptr, int running)
{
	static_cast<QScxmlStateMachine*>(ptr)->runningChanged(running != 0);
}

void* QScxmlStateMachine_ScxmlEventFilter(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->scxmlEventFilter();
}

char* QScxmlStateMachine_SessionId(void* ptr)
{
	return static_cast<QScxmlStateMachine*>(ptr)->sessionId().toUtf8().data();
}

void QScxmlStateMachine_SetDataModel(void* ptr, void* model)
{
	static_cast<QScxmlStateMachine*>(ptr)->setDataModel(static_cast<QScxmlDataModel*>(model));
}

void QScxmlStateMachine_SetRunning(void* ptr, int running)
{
	static_cast<QScxmlStateMachine*>(ptr)->setRunning(running != 0);
}

void QScxmlStateMachine_SetScxmlEventFilter(void* ptr, void* newFilter)
{
	static_cast<QScxmlStateMachine*>(ptr)->setScxmlEventFilter(static_cast<QScxmlEventFilter*>(newFilter));
}

void QScxmlStateMachine_SetSessionId(void* ptr, char* id)
{
	static_cast<QScxmlStateMachine*>(ptr)->setSessionId(QString(id));
}

void QScxmlStateMachine_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "start");
}

char* QScxmlStateMachine_StateNames(void* ptr, int compress)
{
	return static_cast<QScxmlStateMachine*>(ptr)->stateNames(compress != 0).join("|").toUtf8().data();
}

void QScxmlStateMachine_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "stop");
}

void QScxmlStateMachine_SubmitEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(static_cast<QScxmlEvent*>(event));
}

void QScxmlStateMachine_SubmitEvent2(void* ptr, char* eventName)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString(eventName));
}

void QScxmlStateMachine_SubmitEvent3(void* ptr, char* eventName, void* data)
{
	static_cast<QScxmlStateMachine*>(ptr)->submitEvent(QString(eventName), *static_cast<QVariant*>(data));
}

void QScxmlStateMachine_TimerEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlStateMachine_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScxmlStateMachine_ChildEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlStateMachine_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::childEvent(static_cast<QChildEvent*>(event));
}

void QScxmlStateMachine_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_CustomEvent(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScxmlStateMachine_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::customEvent(static_cast<QEvent*>(event));
}

void QScxmlStateMachine_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScxmlStateMachine*>(ptr), "deleteLater");
}

void QScxmlStateMachine_DeleteLaterDefault(void* ptr)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::deleteLater();
}

void QScxmlStateMachine_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScxmlStateMachine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScxmlStateMachine_Event(void* ptr, void* e)
{
	return static_cast<QScxmlStateMachine*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScxmlStateMachine_EventDefault(void* ptr, void* e)
{
	return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::event(static_cast<QEvent*>(e));
}

int QScxmlStateMachine_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlStateMachine*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScxmlStateMachine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScxmlStateMachine_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlStateMachine*>(ptr)->metaObject());
}

void* QScxmlStateMachine_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScxmlStateMachine*>(ptr)->QScxmlStateMachine::metaObject());
}

