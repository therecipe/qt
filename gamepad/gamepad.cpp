// +build !minimal

#define protected public
#define private public

#include "gamepad.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QEvent>
#include <QGamepad>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

class MyQGamepad: public QGamepad
{
public:
	MyQGamepad(int deviceId, QObject *parent) : QGamepad(deviceId, parent) {};
	void setDeviceId(int number) { callbackQGamepad_SetDeviceId(this, this->objectName().toUtf8().data(), number); };
	void Signal_ButtonAChanged(bool value) { callbackQGamepad_ButtonAChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonBChanged(bool value) { callbackQGamepad_ButtonBChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonCenterChanged(bool value) { callbackQGamepad_ButtonCenterChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonDownChanged(bool value) { callbackQGamepad_ButtonDownChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonGuideChanged(bool value) { callbackQGamepad_ButtonGuideChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonL1Changed(bool value) { callbackQGamepad_ButtonL1Changed(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonL3Changed(bool value) { callbackQGamepad_ButtonL3Changed(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonLeftChanged(bool value) { callbackQGamepad_ButtonLeftChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonR1Changed(bool value) { callbackQGamepad_ButtonR1Changed(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonR3Changed(bool value) { callbackQGamepad_ButtonR3Changed(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonRightChanged(bool value) { callbackQGamepad_ButtonRightChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonSelectChanged(bool value) { callbackQGamepad_ButtonSelectChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonStartChanged(bool value) { callbackQGamepad_ButtonStartChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonUpChanged(bool value) { callbackQGamepad_ButtonUpChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonXChanged(bool value) { callbackQGamepad_ButtonXChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ButtonYChanged(bool value) { callbackQGamepad_ButtonYChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_ConnectedChanged(bool value) { callbackQGamepad_ConnectedChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_DeviceIdChanged(int value) { callbackQGamepad_DeviceIdChanged(this, this->objectName().toUtf8().data(), value); };
	void Signal_NameChanged(QString value) { callbackQGamepad_NameChanged(this, this->objectName().toUtf8().data(), value.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQGamepad_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGamepad_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGamepad_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGamepad_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGamepad_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGamepad_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQGamepad_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGamepad_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGamepad_MetaObject(const_cast<MyQGamepad*>(this), this->objectName().toUtf8().data())); };
};

int QGamepad_ButtonA(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonA();
}

int QGamepad_ButtonB(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonB();
}

int QGamepad_ButtonCenter(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonCenter();
}

int QGamepad_ButtonDown(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonDown();
}

int QGamepad_ButtonGuide(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonGuide();
}

int QGamepad_ButtonL1(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonL1();
}

int QGamepad_ButtonL3(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonL3();
}

int QGamepad_ButtonLeft(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonLeft();
}

int QGamepad_ButtonR1(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonR1();
}

int QGamepad_ButtonR3(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonR3();
}

int QGamepad_ButtonRight(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonRight();
}

int QGamepad_ButtonSelect(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonSelect();
}

int QGamepad_ButtonStart(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonStart();
}

int QGamepad_ButtonUp(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonUp();
}

int QGamepad_ButtonX(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonX();
}

int QGamepad_ButtonY(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonY();
}

int QGamepad_DeviceId(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->deviceId();
}

int QGamepad_IsConnected(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->isConnected();
}

char* QGamepad_Name(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->name().toUtf8().data();
}

void QGamepad_SetDeviceId(void* ptr, int number)
{
	QMetaObject::invokeMethod(static_cast<QGamepad*>(ptr), "setDeviceId", Q_ARG(int, number));
}

void* QGamepad_NewQGamepad(int deviceId, void* parent)
{
	return new MyQGamepad(deviceId, static_cast<QObject*>(parent));
}

void QGamepad_ConnectButtonAChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonAChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonAChanged));
}

void QGamepad_DisconnectButtonAChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonAChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonAChanged));
}

void QGamepad_ButtonAChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonAChanged(value != 0);
}

void QGamepad_ConnectButtonBChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonBChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonBChanged));
}

void QGamepad_DisconnectButtonBChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonBChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonBChanged));
}

void QGamepad_ButtonBChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonBChanged(value != 0);
}

void QGamepad_ConnectButtonCenterChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonCenterChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonCenterChanged));
}

void QGamepad_DisconnectButtonCenterChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonCenterChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonCenterChanged));
}

void QGamepad_ButtonCenterChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonCenterChanged(value != 0);
}

void QGamepad_ConnectButtonDownChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonDownChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonDownChanged));
}

void QGamepad_DisconnectButtonDownChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonDownChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonDownChanged));
}

void QGamepad_ButtonDownChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonDownChanged(value != 0);
}

void QGamepad_ConnectButtonGuideChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonGuideChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonGuideChanged));
}

void QGamepad_DisconnectButtonGuideChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonGuideChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonGuideChanged));
}

void QGamepad_ButtonGuideChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonGuideChanged(value != 0);
}

void QGamepad_ConnectButtonL1Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL1Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL1Changed));
}

void QGamepad_DisconnectButtonL1Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL1Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL1Changed));
}

void QGamepad_ButtonL1Changed(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonL1Changed(value != 0);
}

void QGamepad_ConnectButtonL3Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL3Changed));
}

void QGamepad_DisconnectButtonL3Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL3Changed));
}

void QGamepad_ButtonL3Changed(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonL3Changed(value != 0);
}

void QGamepad_ConnectButtonLeftChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonLeftChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonLeftChanged));
}

void QGamepad_DisconnectButtonLeftChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonLeftChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonLeftChanged));
}

void QGamepad_ButtonLeftChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonLeftChanged(value != 0);
}

void QGamepad_ConnectButtonR1Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR1Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR1Changed));
}

void QGamepad_DisconnectButtonR1Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR1Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR1Changed));
}

void QGamepad_ButtonR1Changed(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonR1Changed(value != 0);
}

void QGamepad_ConnectButtonR3Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR3Changed));
}

void QGamepad_DisconnectButtonR3Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR3Changed));
}

void QGamepad_ButtonR3Changed(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonR3Changed(value != 0);
}

void QGamepad_ConnectButtonRightChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonRightChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonRightChanged));
}

void QGamepad_DisconnectButtonRightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonRightChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonRightChanged));
}

void QGamepad_ButtonRightChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonRightChanged(value != 0);
}

void QGamepad_ConnectButtonSelectChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonSelectChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonSelectChanged));
}

void QGamepad_DisconnectButtonSelectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonSelectChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonSelectChanged));
}

void QGamepad_ButtonSelectChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonSelectChanged(value != 0);
}

void QGamepad_ConnectButtonStartChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonStartChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonStartChanged));
}

void QGamepad_DisconnectButtonStartChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonStartChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonStartChanged));
}

void QGamepad_ButtonStartChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonStartChanged(value != 0);
}

void QGamepad_ConnectButtonUpChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonUpChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonUpChanged));
}

void QGamepad_DisconnectButtonUpChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonUpChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonUpChanged));
}

void QGamepad_ButtonUpChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonUpChanged(value != 0);
}

void QGamepad_ConnectButtonXChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonXChanged));
}

void QGamepad_DisconnectButtonXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonXChanged));
}

void QGamepad_ButtonXChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonXChanged(value != 0);
}

void QGamepad_ConnectButtonYChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonYChanged));
}

void QGamepad_DisconnectButtonYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonYChanged));
}

void QGamepad_ButtonYChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->buttonYChanged(value != 0);
}

void QGamepad_ConnectConnectedChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::connectedChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ConnectedChanged));
}

void QGamepad_DisconnectConnectedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::connectedChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ConnectedChanged));
}

void QGamepad_ConnectedChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->connectedChanged(value != 0);
}

void QGamepad_ConnectDeviceIdChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(int)>(&QGamepad::deviceIdChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(int)>(&MyQGamepad::Signal_DeviceIdChanged));
}

void QGamepad_DisconnectDeviceIdChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(int)>(&QGamepad::deviceIdChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(int)>(&MyQGamepad::Signal_DeviceIdChanged));
}

void QGamepad_DeviceIdChanged(void* ptr, int value)
{
	static_cast<QGamepad*>(ptr)->deviceIdChanged(value);
}

void QGamepad_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(QString)>(&QGamepad::nameChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(QString)>(&MyQGamepad::Signal_NameChanged));
}

void QGamepad_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(QString)>(&QGamepad::nameChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(QString)>(&MyQGamepad::Signal_NameChanged));
}

void QGamepad_NameChanged(void* ptr, char* value)
{
	static_cast<QGamepad*>(ptr)->nameChanged(QString(value));
}

void QGamepad_DestroyQGamepad(void* ptr)
{
	static_cast<QGamepad*>(ptr)->~QGamepad();
}

void QGamepad_TimerEvent(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGamepad_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->QGamepad::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGamepad_ChildEvent(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGamepad_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->QGamepad::childEvent(static_cast<QChildEvent*>(event));
}

void QGamepad_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGamepad*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepad_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGamepad*>(ptr)->QGamepad::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepad_CustomEvent(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGamepad_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->QGamepad::customEvent(static_cast<QEvent*>(event));
}

void QGamepad_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGamepad*>(ptr), "deleteLater");
}

void QGamepad_DeleteLaterDefault(void* ptr)
{
	static_cast<QGamepad*>(ptr)->QGamepad::deleteLater();
}

void QGamepad_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGamepad*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepad_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGamepad*>(ptr)->QGamepad::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGamepad_Event(void* ptr, void* e)
{
	return static_cast<QGamepad*>(ptr)->event(static_cast<QEvent*>(e));
}

int QGamepad_EventDefault(void* ptr, void* e)
{
	return static_cast<QGamepad*>(ptr)->QGamepad::event(static_cast<QEvent*>(e));
}

int QGamepad_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGamepad*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGamepad_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGamepad*>(ptr)->QGamepad::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGamepad_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGamepad*>(ptr)->metaObject());
}

void* QGamepad_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGamepad*>(ptr)->QGamepad::metaObject());
}

