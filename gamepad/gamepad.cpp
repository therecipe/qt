// +build !minimal

#define protected public
#define private public

#include "gamepad.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGamepad>
#include <QList>
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
	void Signal_AxisLeftXChanged(double value) { callbackQGamepad_AxisLeftXChanged(this, value); };
	void Signal_AxisLeftYChanged(double value) { callbackQGamepad_AxisLeftYChanged(this, value); };
	void Signal_AxisRightXChanged(double value) { callbackQGamepad_AxisRightXChanged(this, value); };
	void Signal_AxisRightYChanged(double value) { callbackQGamepad_AxisRightYChanged(this, value); };
	void Signal_ButtonAChanged(bool value) { callbackQGamepad_ButtonAChanged(this, value); };
	void Signal_ButtonBChanged(bool value) { callbackQGamepad_ButtonBChanged(this, value); };
	void Signal_ButtonCenterChanged(bool value) { callbackQGamepad_ButtonCenterChanged(this, value); };
	void Signal_ButtonDownChanged(bool value) { callbackQGamepad_ButtonDownChanged(this, value); };
	void Signal_ButtonGuideChanged(bool value) { callbackQGamepad_ButtonGuideChanged(this, value); };
	void Signal_ButtonL1Changed(bool value) { callbackQGamepad_ButtonL1Changed(this, value); };
	void Signal_ButtonL2Changed(double value) { callbackQGamepad_ButtonL2Changed(this, value); };
	void Signal_ButtonL3Changed(bool value) { callbackQGamepad_ButtonL3Changed(this, value); };
	void Signal_ButtonLeftChanged(bool value) { callbackQGamepad_ButtonLeftChanged(this, value); };
	void Signal_ButtonR1Changed(bool value) { callbackQGamepad_ButtonR1Changed(this, value); };
	void Signal_ButtonR2Changed(double value) { callbackQGamepad_ButtonR2Changed(this, value); };
	void Signal_ButtonR3Changed(bool value) { callbackQGamepad_ButtonR3Changed(this, value); };
	void Signal_ButtonRightChanged(bool value) { callbackQGamepad_ButtonRightChanged(this, value); };
	void Signal_ButtonSelectChanged(bool value) { callbackQGamepad_ButtonSelectChanged(this, value); };
	void Signal_ButtonStartChanged(bool value) { callbackQGamepad_ButtonStartChanged(this, value); };
	void Signal_ButtonUpChanged(bool value) { callbackQGamepad_ButtonUpChanged(this, value); };
	void Signal_ButtonXChanged(bool value) { callbackQGamepad_ButtonXChanged(this, value); };
	void Signal_ButtonYChanged(bool value) { callbackQGamepad_ButtonYChanged(this, value); };
	void Signal_ConnectedChanged(bool value) { callbackQGamepad_ConnectedChanged(this, value); };
	void Signal_DeviceIdChanged(int value) { callbackQGamepad_DeviceIdChanged(this, value); };
	void Signal_NameChanged(QString value) { QByteArray tf32b67 = value.toUtf8(); QtGamepad_PackedString valuePacked = { const_cast<char*>(tf32b67.prepend("WHITESPACE").constData()+10), tf32b67.size()-10 };callbackQGamepad_NameChanged(this, valuePacked); };
	void setDeviceId(int number) { callbackQGamepad_SetDeviceId(this, number); };
	bool event(QEvent * e) { return callbackQGamepad_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGamepad_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGamepad_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGamepad_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGamepad_CustomEvent(this, event); };
	void deleteLater() { callbackQGamepad_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGamepad_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void timerEvent(QTimerEvent * event) { callbackQGamepad_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGamepad_MetaObject(const_cast<MyQGamepad*>(this))); };
};

void* QGamepad_NewQGamepad(int deviceId, void* parent)
{
	return new MyQGamepad(deviceId, static_cast<QObject*>(parent));
}

void QGamepad_ConnectAxisLeftXChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisLeftXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisLeftXChanged));
}

void QGamepad_DisconnectAxisLeftXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisLeftXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisLeftXChanged));
}

void QGamepad_AxisLeftXChanged(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->axisLeftXChanged(value);
}

void QGamepad_ConnectAxisLeftYChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisLeftYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisLeftYChanged));
}

void QGamepad_DisconnectAxisLeftYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisLeftYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisLeftYChanged));
}

void QGamepad_AxisLeftYChanged(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->axisLeftYChanged(value);
}

void QGamepad_ConnectAxisRightXChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisRightXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisRightXChanged));
}

void QGamepad_DisconnectAxisRightXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisRightXChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisRightXChanged));
}

void QGamepad_AxisRightXChanged(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->axisRightXChanged(value);
}

void QGamepad_ConnectAxisRightYChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisRightYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisRightYChanged));
}

void QGamepad_DisconnectAxisRightYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::axisRightYChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_AxisRightYChanged));
}

void QGamepad_AxisRightYChanged(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->axisRightYChanged(value);
}

void QGamepad_ConnectButtonAChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonAChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonAChanged));
}

void QGamepad_DisconnectButtonAChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonAChanged), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonAChanged));
}

void QGamepad_ButtonAChanged(void* ptr, char value)
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

void QGamepad_ButtonBChanged(void* ptr, char value)
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

void QGamepad_ButtonCenterChanged(void* ptr, char value)
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

void QGamepad_ButtonDownChanged(void* ptr, char value)
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

void QGamepad_ButtonGuideChanged(void* ptr, char value)
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

void QGamepad_ButtonL1Changed(void* ptr, char value)
{
	static_cast<QGamepad*>(ptr)->buttonL1Changed(value != 0);
}

void QGamepad_ConnectButtonL2Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::buttonL2Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_ButtonL2Changed));
}

void QGamepad_DisconnectButtonL2Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::buttonL2Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_ButtonL2Changed));
}

void QGamepad_ButtonL2Changed(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->buttonL2Changed(value);
}

void QGamepad_ConnectButtonL3Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL3Changed));
}

void QGamepad_DisconnectButtonL3Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonL3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonL3Changed));
}

void QGamepad_ButtonL3Changed(void* ptr, char value)
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

void QGamepad_ButtonLeftChanged(void* ptr, char value)
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

void QGamepad_ButtonR1Changed(void* ptr, char value)
{
	static_cast<QGamepad*>(ptr)->buttonR1Changed(value != 0);
}

void QGamepad_ConnectButtonR2Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::buttonR2Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_ButtonR2Changed));
}

void QGamepad_DisconnectButtonR2Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(double)>(&QGamepad::buttonR2Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(double)>(&MyQGamepad::Signal_ButtonR2Changed));
}

void QGamepad_ButtonR2Changed(void* ptr, double value)
{
	static_cast<QGamepad*>(ptr)->buttonR2Changed(value);
}

void QGamepad_ConnectButtonR3Changed(void* ptr)
{
	QObject::connect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR3Changed));
}

void QGamepad_DisconnectButtonR3Changed(void* ptr)
{
	QObject::disconnect(static_cast<QGamepad*>(ptr), static_cast<void (QGamepad::*)(bool)>(&QGamepad::buttonR3Changed), static_cast<MyQGamepad*>(ptr), static_cast<void (MyQGamepad::*)(bool)>(&MyQGamepad::Signal_ButtonR3Changed));
}

void QGamepad_ButtonR3Changed(void* ptr, char value)
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

void QGamepad_ButtonRightChanged(void* ptr, char value)
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

void QGamepad_ButtonSelectChanged(void* ptr, char value)
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

void QGamepad_ButtonStartChanged(void* ptr, char value)
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

void QGamepad_ButtonUpChanged(void* ptr, char value)
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

void QGamepad_ButtonXChanged(void* ptr, char value)
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

void QGamepad_ButtonYChanged(void* ptr, char value)
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

void QGamepad_ConnectedChanged(void* ptr, char value)
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

void QGamepad_SetDeviceId(void* ptr, int number)
{
	QMetaObject::invokeMethod(static_cast<QGamepad*>(ptr), "setDeviceId", Q_ARG(int, number));
}

void QGamepad_SetDeviceIdDefault(void* ptr, int number)
{
	static_cast<QGamepad*>(ptr)->QGamepad::setDeviceId(number);
}

void QGamepad_DestroyQGamepad(void* ptr)
{
	static_cast<QGamepad*>(ptr)->~QGamepad();
}

struct QtGamepad_PackedString QGamepad_Name(void* ptr)
{
	return ({ QByteArray t4fe02c = static_cast<QGamepad*>(ptr)->name().toUtf8(); QtGamepad_PackedString { const_cast<char*>(t4fe02c.prepend("WHITESPACE").constData()+10), t4fe02c.size()-10 }; });
}

char QGamepad_ButtonA(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonA();
}

char QGamepad_ButtonB(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonB();
}

char QGamepad_ButtonCenter(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonCenter();
}

char QGamepad_ButtonDown(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonDown();
}

char QGamepad_ButtonGuide(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonGuide();
}

char QGamepad_ButtonL1(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonL1();
}

char QGamepad_ButtonL3(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonL3();
}

char QGamepad_ButtonLeft(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonLeft();
}

char QGamepad_ButtonR1(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonR1();
}

char QGamepad_ButtonR3(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonR3();
}

char QGamepad_ButtonRight(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonRight();
}

char QGamepad_ButtonSelect(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonSelect();
}

char QGamepad_ButtonStart(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonStart();
}

char QGamepad_ButtonUp(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonUp();
}

char QGamepad_ButtonX(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonX();
}

char QGamepad_ButtonY(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonY();
}

char QGamepad_IsConnected(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->isConnected();
}

double QGamepad_AxisLeftX(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->axisLeftX();
}

double QGamepad_AxisLeftY(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->axisLeftY();
}

double QGamepad_AxisRightX(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->axisRightX();
}

double QGamepad_AxisRightY(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->axisRightY();
}

double QGamepad_ButtonL2(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonL2();
}

double QGamepad_ButtonR2(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->buttonR2();
}

int QGamepad_DeviceId(void* ptr)
{
	return static_cast<QGamepad*>(ptr)->deviceId();
}

void* QGamepad___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGamepad___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGamepad___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QGamepad___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepad___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepad___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QGamepad___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepad___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepad___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QGamepad___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepad___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepad___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QGamepad___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGamepad___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepad___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QGamepad_Event(void* ptr, void* e)
{
	return static_cast<QGamepad*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGamepad_EventDefault(void* ptr, void* e)
{
	return static_cast<QGamepad*>(ptr)->QGamepad::event(static_cast<QEvent*>(e));
}

char QGamepad_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGamepad*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGamepad_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGamepad*>(ptr)->QGamepad::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
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

void QGamepad_TimerEvent(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGamepad_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGamepad*>(ptr)->QGamepad::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGamepad_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGamepad*>(ptr)->metaObject());
}

void* QGamepad_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGamepad*>(ptr)->QGamepad::metaObject());
}

