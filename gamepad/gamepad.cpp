// +build !minimal

#define protected public
#define private public

#include "gamepad.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGamepad>
#include <QGamepadKeyNavigation>
#include <QGamepadManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QWidget>
#include <QWindow>

class MyQGamepad: public QGamepad
{
public:
	MyQGamepad(int deviceId = 0, QObject *parent = nullptr) : QGamepad(deviceId, parent) {QGamepad_QGamepad_QRegisterMetaType();};
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
	void Signal_Destroyed(QObject * obj) { callbackQGamepad_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGamepad_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGamepad_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGamepad_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGamepad_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGamepad_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGamepad*)

int QGamepad_QGamepad_QRegisterMetaType(){qRegisterMetaType<QGamepad*>(); return qRegisterMetaType<MyQGamepad*>();}

void* QGamepad_NewQGamepad(int deviceId, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGamepad(deviceId, static_cast<QWindow*>(parent));
	} else {
		return new MyQGamepad(deviceId, static_cast<QObject*>(parent));
	}
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

void QGamepad_NameChanged(void* ptr, struct QtGamepad_PackedString value)
{
	static_cast<QGamepad*>(ptr)->nameChanged(QString::fromUtf8(value.data, value.len));
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
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
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGamepad_EventDefault(void* ptr, void* e)
{
		return static_cast<QGamepad*>(ptr)->QGamepad::event(static_cast<QEvent*>(e));
}

char QGamepad_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGamepad*>(ptr)->QGamepad::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGamepad_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGamepad*>(ptr)->QGamepad::childEvent(static_cast<QChildEvent*>(event));
}

void QGamepad_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepad*>(ptr)->QGamepad::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepad_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGamepad*>(ptr)->QGamepad::customEvent(static_cast<QEvent*>(event));
}

void QGamepad_DeleteLaterDefault(void* ptr)
{
		static_cast<QGamepad*>(ptr)->QGamepad::deleteLater();
}

void QGamepad_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepad*>(ptr)->QGamepad::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepad_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGamepad*>(ptr)->QGamepad::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGamepad_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGamepad*>(ptr)->QGamepad::metaObject());
}

class MyQGamepadKeyNavigation: public QGamepadKeyNavigation
{
public:
	MyQGamepadKeyNavigation(QObject *parent = nullptr) : QGamepadKeyNavigation(parent) {QGamepadKeyNavigation_QGamepadKeyNavigation_QRegisterMetaType();};
	void Signal_ActiveChanged(bool isActive) { callbackQGamepadKeyNavigation_ActiveChanged(this, isActive); };
	void Signal_ButtonAKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonAKeyChanged(this, key); };
	void Signal_ButtonBKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonBKeyChanged(this, key); };
	void Signal_ButtonGuideKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonGuideKeyChanged(this, key); };
	void Signal_ButtonL1KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonL1KeyChanged(this, key); };
	void Signal_ButtonL2KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonL2KeyChanged(this, key); };
	void Signal_ButtonL3KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonL3KeyChanged(this, key); };
	void Signal_ButtonR1KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonR1KeyChanged(this, key); };
	void Signal_ButtonR2KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonR2KeyChanged(this, key); };
	void Signal_ButtonR3KeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonR3KeyChanged(this, key); };
	void Signal_ButtonSelectKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonSelectKeyChanged(this, key); };
	void Signal_ButtonStartKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonStartKeyChanged(this, key); };
	void Signal_ButtonXKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonXKeyChanged(this, key); };
	void Signal_ButtonYKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_ButtonYKeyChanged(this, key); };
	void Signal_DownKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_DownKeyChanged(this, key); };
	void Signal_GamepadChanged(QGamepad * gamepad) { callbackQGamepadKeyNavigation_GamepadChanged(this, gamepad); };
	void Signal_LeftKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_LeftKeyChanged(this, key); };
	void Signal_RightKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_RightKeyChanged(this, key); };
	void setActive(bool isActive) { callbackQGamepadKeyNavigation_SetActive(this, isActive); };
	void setButtonAKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonAKey(this, key); };
	void setButtonBKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonBKey(this, key); };
	void setButtonGuideKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonGuideKey(this, key); };
	void setButtonL1Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonL1Key(this, key); };
	void setButtonL2Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonL2Key(this, key); };
	void setButtonL3Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonL3Key(this, key); };
	void setButtonR1Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonR1Key(this, key); };
	void setButtonR2Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonR2Key(this, key); };
	void setButtonR3Key(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonR3Key(this, key); };
	void setButtonSelectKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonSelectKey(this, key); };
	void setButtonStartKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonStartKey(this, key); };
	void setButtonXKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonXKey(this, key); };
	void setButtonYKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetButtonYKey(this, key); };
	void setDownKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetDownKey(this, key); };
	void setGamepad(QGamepad * gamepad) { callbackQGamepadKeyNavigation_SetGamepad(this, gamepad); };
	void setLeftKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetLeftKey(this, key); };
	void setRightKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetRightKey(this, key); };
	void setUpKey(Qt::Key key) { callbackQGamepadKeyNavigation_SetUpKey(this, key); };
	void Signal_UpKeyChanged(Qt::Key key) { callbackQGamepadKeyNavigation_UpKeyChanged(this, key); };
	bool event(QEvent * e) { return callbackQGamepadKeyNavigation_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGamepadKeyNavigation_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGamepadKeyNavigation_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGamepadKeyNavigation_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGamepadKeyNavigation_CustomEvent(this, event); };
	void deleteLater() { callbackQGamepadKeyNavigation_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGamepadKeyNavigation_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGamepadKeyNavigation_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGamepad_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGamepadKeyNavigation_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGamepadKeyNavigation_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGamepadKeyNavigation_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGamepadKeyNavigation*)

int QGamepadKeyNavigation_QGamepadKeyNavigation_QRegisterMetaType(){qRegisterMetaType<QGamepadKeyNavigation*>(); return qRegisterMetaType<MyQGamepadKeyNavigation*>();}

void* QGamepadKeyNavigation_NewQGamepadKeyNavigation(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGamepadKeyNavigation(static_cast<QWindow*>(parent));
	} else {
		return new MyQGamepadKeyNavigation(static_cast<QObject*>(parent));
	}
}

void QGamepadKeyNavigation_ConnectActiveChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(bool)>(&QGamepadKeyNavigation::activeChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(bool)>(&MyQGamepadKeyNavigation::Signal_ActiveChanged));
}

void QGamepadKeyNavigation_DisconnectActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(bool)>(&QGamepadKeyNavigation::activeChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(bool)>(&MyQGamepadKeyNavigation::Signal_ActiveChanged));
}

void QGamepadKeyNavigation_ActiveChanged(void* ptr, char isActive)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->activeChanged(isActive != 0);
}

void QGamepadKeyNavigation_ConnectButtonAKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonAKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonAKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonAKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonAKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonAKeyChanged));
}

void QGamepadKeyNavigation_ButtonAKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonAKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonBKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonBKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonBKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonBKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonBKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonBKeyChanged));
}

void QGamepadKeyNavigation_ButtonBKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonBKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonGuideKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonGuideKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonGuideKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonGuideKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonGuideKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonGuideKeyChanged));
}

void QGamepadKeyNavigation_ButtonGuideKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonGuideKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonL1KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL1KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL1KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonL1KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL1KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL1KeyChanged));
}

void QGamepadKeyNavigation_ButtonL1KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonL1KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonL2KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL2KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL2KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonL2KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL2KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL2KeyChanged));
}

void QGamepadKeyNavigation_ButtonL2KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonL2KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonL3KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL3KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL3KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonL3KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonL3KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonL3KeyChanged));
}

void QGamepadKeyNavigation_ButtonL3KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonL3KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonR1KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR1KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR1KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonR1KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR1KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR1KeyChanged));
}

void QGamepadKeyNavigation_ButtonR1KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonR1KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonR2KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR2KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR2KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonR2KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR2KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR2KeyChanged));
}

void QGamepadKeyNavigation_ButtonR2KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonR2KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonR3KeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR3KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR3KeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonR3KeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonR3KeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonR3KeyChanged));
}

void QGamepadKeyNavigation_ButtonR3KeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonR3KeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonSelectKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonSelectKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonSelectKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonSelectKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonSelectKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonSelectKeyChanged));
}

void QGamepadKeyNavigation_ButtonSelectKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonSelectKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonStartKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonStartKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonStartKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonStartKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonStartKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonStartKeyChanged));
}

void QGamepadKeyNavigation_ButtonStartKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonStartKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonXKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonXKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonXKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonXKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonXKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonXKeyChanged));
}

void QGamepadKeyNavigation_ButtonXKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonXKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectButtonYKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonYKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonYKeyChanged));
}

void QGamepadKeyNavigation_DisconnectButtonYKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::buttonYKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_ButtonYKeyChanged));
}

void QGamepadKeyNavigation_ButtonYKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->buttonYKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectDownKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::downKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_DownKeyChanged));
}

void QGamepadKeyNavigation_DisconnectDownKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::downKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_DownKeyChanged));
}

void QGamepadKeyNavigation_DownKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->downKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectGamepadChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(QGamepad *)>(&QGamepadKeyNavigation::gamepadChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(QGamepad *)>(&MyQGamepadKeyNavigation::Signal_GamepadChanged));
}

void QGamepadKeyNavigation_DisconnectGamepadChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(QGamepad *)>(&QGamepadKeyNavigation::gamepadChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(QGamepad *)>(&MyQGamepadKeyNavigation::Signal_GamepadChanged));
}

void QGamepadKeyNavigation_GamepadChanged(void* ptr, void* gamepad)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->gamepadChanged(static_cast<QGamepad*>(gamepad));
}

void QGamepadKeyNavigation_ConnectLeftKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::leftKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_LeftKeyChanged));
}

void QGamepadKeyNavigation_DisconnectLeftKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::leftKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_LeftKeyChanged));
}

void QGamepadKeyNavigation_LeftKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->leftKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectRightKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::rightKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_RightKeyChanged));
}

void QGamepadKeyNavigation_DisconnectRightKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::rightKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_RightKeyChanged));
}

void QGamepadKeyNavigation_RightKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->rightKeyChanged(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetActive(void* ptr, char isActive)
{
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setActive", Q_ARG(bool, isActive != 0));
}

void QGamepadKeyNavigation_SetActiveDefault(void* ptr, char isActive)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setActive(isActive != 0);
}

void QGamepadKeyNavigation_SetButtonAKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonAKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonAKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonAKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonBKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonBKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonBKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonBKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonGuideKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonGuideKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonGuideKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonGuideKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonL1Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonL1Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonL1KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonL1Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonL2Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonL2Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonL2KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonL2Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonL3Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonL3Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonL3KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonL3Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonR1Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonR1Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonR1KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonR1Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonR2Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonR2Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonR2KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonR2Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonR3Key(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonR3Key", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonR3KeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonR3Key(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonSelectKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonSelectKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonSelectKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonSelectKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonStartKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonStartKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonStartKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonStartKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonXKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonXKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonXKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonXKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetButtonYKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setButtonYKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetButtonYKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setButtonYKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetDownKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setDownKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetDownKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setDownKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetGamepad(void* ptr, void* gamepad)
{
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setGamepad", Q_ARG(QGamepad*, static_cast<QGamepad*>(gamepad)));
}

void QGamepadKeyNavigation_SetGamepadDefault(void* ptr, void* gamepad)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setGamepad(static_cast<QGamepad*>(gamepad));
}

void QGamepadKeyNavigation_SetLeftKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setLeftKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetLeftKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setLeftKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetRightKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setRightKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetRightKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setRightKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_SetUpKey(void* ptr, long long key)
{
	qRegisterMetaType<Qt::Key>();
	QMetaObject::invokeMethod(static_cast<QGamepadKeyNavigation*>(ptr), "setUpKey", Q_ARG(Qt::Key, static_cast<Qt::Key>(key)));
}

void QGamepadKeyNavigation_SetUpKeyDefault(void* ptr, long long key)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::setUpKey(static_cast<Qt::Key>(key));
}

void QGamepadKeyNavigation_ConnectUpKeyChanged(void* ptr)
{
	qRegisterMetaType<Qt::Key>();
	QObject::connect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::upKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_UpKeyChanged));
}

void QGamepadKeyNavigation_DisconnectUpKeyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadKeyNavigation*>(ptr), static_cast<void (QGamepadKeyNavigation::*)(Qt::Key)>(&QGamepadKeyNavigation::upKeyChanged), static_cast<MyQGamepadKeyNavigation*>(ptr), static_cast<void (MyQGamepadKeyNavigation::*)(Qt::Key)>(&MyQGamepadKeyNavigation::Signal_UpKeyChanged));
}

void QGamepadKeyNavigation_UpKeyChanged(void* ptr, long long key)
{
	static_cast<QGamepadKeyNavigation*>(ptr)->upKeyChanged(static_cast<Qt::Key>(key));
}

void* QGamepadKeyNavigation_Gamepad(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->gamepad();
}

long long QGamepadKeyNavigation_ButtonAKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonAKey();
}

long long QGamepadKeyNavigation_ButtonBKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonBKey();
}

long long QGamepadKeyNavigation_ButtonGuideKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonGuideKey();
}

long long QGamepadKeyNavigation_ButtonL1Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonL1Key();
}

long long QGamepadKeyNavigation_ButtonL2Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonL2Key();
}

long long QGamepadKeyNavigation_ButtonL3Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonL3Key();
}

long long QGamepadKeyNavigation_ButtonR1Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonR1Key();
}

long long QGamepadKeyNavigation_ButtonR2Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonR2Key();
}

long long QGamepadKeyNavigation_ButtonR3Key(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonR3Key();
}

long long QGamepadKeyNavigation_ButtonSelectKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonSelectKey();
}

long long QGamepadKeyNavigation_ButtonStartKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonStartKey();
}

long long QGamepadKeyNavigation_ButtonXKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonXKey();
}

long long QGamepadKeyNavigation_ButtonYKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->buttonYKey();
}

long long QGamepadKeyNavigation_DownKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->downKey();
}

long long QGamepadKeyNavigation_LeftKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->leftKey();
}

long long QGamepadKeyNavigation_RightKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->rightKey();
}

long long QGamepadKeyNavigation_UpKey(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->upKey();
}

char QGamepadKeyNavigation_Active(void* ptr)
{
	return static_cast<QGamepadKeyNavigation*>(ptr)->active();
}

void* QGamepadKeyNavigation___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGamepadKeyNavigation___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGamepadKeyNavigation___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGamepadKeyNavigation___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadKeyNavigation___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadKeyNavigation___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadKeyNavigation___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadKeyNavigation___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadKeyNavigation___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadKeyNavigation___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadKeyNavigation___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadKeyNavigation___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadKeyNavigation___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGamepadKeyNavigation___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadKeyNavigation___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGamepadKeyNavigation_EventDefault(void* ptr, void* e)
{
		return static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::event(static_cast<QEvent*>(e));
}

char QGamepadKeyNavigation_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGamepadKeyNavigation_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::childEvent(static_cast<QChildEvent*>(event));
}

void QGamepadKeyNavigation_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepadKeyNavigation_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::customEvent(static_cast<QEvent*>(event));
}

void QGamepadKeyNavigation_DeleteLaterDefault(void* ptr)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::deleteLater();
}

void QGamepadKeyNavigation_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepadKeyNavigation_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGamepadKeyNavigation_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGamepadKeyNavigation*>(ptr)->QGamepadKeyNavigation::metaObject());
}

class MyQGamepadManager: public QGamepadManager
{
public:
	bool configureAxis(int deviceId, QGamepadManager::GamepadAxis axis) { return callbackQGamepadManager_ConfigureAxis(this, deviceId, axis) != 0; };
	bool configureButton(int deviceId, QGamepadManager::GamepadButton button) { return callbackQGamepadManager_ConfigureButton(this, deviceId, button) != 0; };
	bool setCancelConfigureButton(int deviceId, QGamepadManager::GamepadButton button) { return callbackQGamepadManager_SetCancelConfigureButton(this, deviceId, button) != 0; };
	void Signal_AxisConfigured(int deviceId, QGamepadManager::GamepadAxis axis) { callbackQGamepadManager_AxisConfigured(this, deviceId, axis); };
	void Signal_ButtonConfigured(int deviceId, QGamepadManager::GamepadButton button) { callbackQGamepadManager_ButtonConfigured(this, deviceId, button); };
	void Signal_ConfigurationCanceled(int deviceId) { callbackQGamepadManager_ConfigurationCanceled(this, deviceId); };
	void Signal_ConnectedGamepadsChanged() { callbackQGamepadManager_ConnectedGamepadsChanged(this); };
	void Signal_GamepadAxisEvent(int deviceId, QGamepadManager::GamepadAxis axis, double value) { callbackQGamepadManager_GamepadAxisEvent(this, deviceId, axis, value); };
	void Signal_GamepadButtonPressEvent(int deviceId, QGamepadManager::GamepadButton button, double value) { callbackQGamepadManager_GamepadButtonPressEvent(this, deviceId, button, value); };
	void Signal_GamepadButtonReleaseEvent(int deviceId, QGamepadManager::GamepadButton button) { callbackQGamepadManager_GamepadButtonReleaseEvent(this, deviceId, button); };
	void Signal_GamepadConnected(int deviceId) { callbackQGamepadManager_GamepadConnected(this, deviceId); };
	void Signal_GamepadDisconnected(int deviceId) { callbackQGamepadManager_GamepadDisconnected(this, deviceId); };
	void resetConfiguration(int deviceId) { callbackQGamepadManager_ResetConfiguration(this, deviceId); };
	void setSettingsFile(const QString & file) { QByteArray t971c41 = file.toUtf8(); QtGamepad_PackedString filePacked = { const_cast<char*>(t971c41.prepend("WHITESPACE").constData()+10), t971c41.size()-10 };callbackQGamepadManager_SetSettingsFile(this, filePacked); };
	bool isConfigurationNeeded(int deviceId) const { return callbackQGamepadManager_IsConfigurationNeeded(const_cast<void*>(static_cast<const void*>(this)), deviceId) != 0; };
	bool event(QEvent * e) { return callbackQGamepadManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGamepadManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGamepadManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGamepadManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGamepadManager_CustomEvent(this, event); };
	void deleteLater() { callbackQGamepadManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGamepadManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGamepadManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGamepad_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGamepadManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGamepadManager_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGamepadManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGamepadManager*)

int QGamepadManager_QGamepadManager_QRegisterMetaType(){qRegisterMetaType<QGamepadManager*>(); return qRegisterMetaType<MyQGamepadManager*>();}

void* QGamepadManager_QGamepadManager_Instance()
{
	return QGamepadManager::instance();
}

char QGamepadManager_ConfigureAxis(void* ptr, int deviceId, long long axis)
{
	qRegisterMetaType<QGamepadManager::GamepadAxis>();
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "configureAxis", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, deviceId), Q_ARG(QGamepadManager::GamepadAxis, static_cast<QGamepadManager::GamepadAxis>(axis)));
	return returnArg;
}

char QGamepadManager_ConfigureAxisDefault(void* ptr, int deviceId, long long axis)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::configureAxis(deviceId, static_cast<QGamepadManager::GamepadAxis>(axis));
}

char QGamepadManager_ConfigureButton(void* ptr, int deviceId, long long button)
{
	qRegisterMetaType<QGamepadManager::GamepadButton>();
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "configureButton", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, deviceId), Q_ARG(QGamepadManager::GamepadButton, static_cast<QGamepadManager::GamepadButton>(button)));
	return returnArg;
}

char QGamepadManager_ConfigureButtonDefault(void* ptr, int deviceId, long long button)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::configureButton(deviceId, static_cast<QGamepadManager::GamepadButton>(button));
}

char QGamepadManager_SetCancelConfigureButton(void* ptr, int deviceId, long long button)
{
	qRegisterMetaType<QGamepadManager::GamepadButton>();
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "setCancelConfigureButton", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, deviceId), Q_ARG(QGamepadManager::GamepadButton, static_cast<QGamepadManager::GamepadButton>(button)));
	return returnArg;
}

char QGamepadManager_SetCancelConfigureButtonDefault(void* ptr, int deviceId, long long button)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::setCancelConfigureButton(deviceId, static_cast<QGamepadManager::GamepadButton>(button));
}

void QGamepadManager_ConnectAxisConfigured(void* ptr)
{
	qRegisterMetaType<QGamepadManager::GamepadAxis>();
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadAxis)>(&QGamepadManager::axisConfigured), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadAxis)>(&MyQGamepadManager::Signal_AxisConfigured));
}

void QGamepadManager_DisconnectAxisConfigured(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadAxis)>(&QGamepadManager::axisConfigured), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadAxis)>(&MyQGamepadManager::Signal_AxisConfigured));
}

void QGamepadManager_AxisConfigured(void* ptr, int deviceId, long long axis)
{
	static_cast<QGamepadManager*>(ptr)->axisConfigured(deviceId, static_cast<QGamepadManager::GamepadAxis>(axis));
}

void QGamepadManager_ConnectButtonConfigured(void* ptr)
{
	qRegisterMetaType<QGamepadManager::GamepadButton>();
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&QGamepadManager::buttonConfigured), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&MyQGamepadManager::Signal_ButtonConfigured));
}

void QGamepadManager_DisconnectButtonConfigured(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&QGamepadManager::buttonConfigured), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&MyQGamepadManager::Signal_ButtonConfigured));
}

void QGamepadManager_ButtonConfigured(void* ptr, int deviceId, long long button)
{
	static_cast<QGamepadManager*>(ptr)->buttonConfigured(deviceId, static_cast<QGamepadManager::GamepadButton>(button));
}

void QGamepadManager_ConnectConfigurationCanceled(void* ptr)
{
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::configurationCanceled), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_ConfigurationCanceled));
}

void QGamepadManager_DisconnectConfigurationCanceled(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::configurationCanceled), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_ConfigurationCanceled));
}

void QGamepadManager_ConfigurationCanceled(void* ptr, int deviceId)
{
	static_cast<QGamepadManager*>(ptr)->configurationCanceled(deviceId);
}

void QGamepadManager_ConnectConnectedGamepadsChanged(void* ptr)
{
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)()>(&QGamepadManager::connectedGamepadsChanged), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)()>(&MyQGamepadManager::Signal_ConnectedGamepadsChanged));
}

void QGamepadManager_DisconnectConnectedGamepadsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)()>(&QGamepadManager::connectedGamepadsChanged), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)()>(&MyQGamepadManager::Signal_ConnectedGamepadsChanged));
}

void QGamepadManager_ConnectedGamepadsChanged(void* ptr)
{
	static_cast<QGamepadManager*>(ptr)->connectedGamepadsChanged();
}

void QGamepadManager_ConnectGamepadAxisEvent(void* ptr)
{
	qRegisterMetaType<QGamepadManager::GamepadAxis>();
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadAxis, double)>(&QGamepadManager::gamepadAxisEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadAxis, double)>(&MyQGamepadManager::Signal_GamepadAxisEvent));
}

void QGamepadManager_DisconnectGamepadAxisEvent(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadAxis, double)>(&QGamepadManager::gamepadAxisEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadAxis, double)>(&MyQGamepadManager::Signal_GamepadAxisEvent));
}

void QGamepadManager_GamepadAxisEvent(void* ptr, int deviceId, long long axis, double value)
{
	static_cast<QGamepadManager*>(ptr)->gamepadAxisEvent(deviceId, static_cast<QGamepadManager::GamepadAxis>(axis), value);
}

void QGamepadManager_ConnectGamepadButtonPressEvent(void* ptr)
{
	qRegisterMetaType<QGamepadManager::GamepadButton>();
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton, double)>(&QGamepadManager::gamepadButtonPressEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton, double)>(&MyQGamepadManager::Signal_GamepadButtonPressEvent));
}

void QGamepadManager_DisconnectGamepadButtonPressEvent(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton, double)>(&QGamepadManager::gamepadButtonPressEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton, double)>(&MyQGamepadManager::Signal_GamepadButtonPressEvent));
}

void QGamepadManager_GamepadButtonPressEvent(void* ptr, int deviceId, long long button, double value)
{
	static_cast<QGamepadManager*>(ptr)->gamepadButtonPressEvent(deviceId, static_cast<QGamepadManager::GamepadButton>(button), value);
}

void QGamepadManager_ConnectGamepadButtonReleaseEvent(void* ptr)
{
	qRegisterMetaType<QGamepadManager::GamepadButton>();
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&QGamepadManager::gamepadButtonReleaseEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&MyQGamepadManager::Signal_GamepadButtonReleaseEvent));
}

void QGamepadManager_DisconnectGamepadButtonReleaseEvent(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&QGamepadManager::gamepadButtonReleaseEvent), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int, QGamepadManager::GamepadButton)>(&MyQGamepadManager::Signal_GamepadButtonReleaseEvent));
}

void QGamepadManager_GamepadButtonReleaseEvent(void* ptr, int deviceId, long long button)
{
	static_cast<QGamepadManager*>(ptr)->gamepadButtonReleaseEvent(deviceId, static_cast<QGamepadManager::GamepadButton>(button));
}

void QGamepadManager_ConnectGamepadConnected(void* ptr)
{
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::gamepadConnected), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_GamepadConnected));
}

void QGamepadManager_DisconnectGamepadConnected(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::gamepadConnected), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_GamepadConnected));
}

void QGamepadManager_GamepadConnected(void* ptr, int deviceId)
{
	static_cast<QGamepadManager*>(ptr)->gamepadConnected(deviceId);
}

void QGamepadManager_ConnectGamepadDisconnected(void* ptr)
{
	QObject::connect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::gamepadDisconnected), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_GamepadDisconnected));
}

void QGamepadManager_DisconnectGamepadDisconnected(void* ptr)
{
	QObject::disconnect(static_cast<QGamepadManager*>(ptr), static_cast<void (QGamepadManager::*)(int)>(&QGamepadManager::gamepadDisconnected), static_cast<MyQGamepadManager*>(ptr), static_cast<void (MyQGamepadManager::*)(int)>(&MyQGamepadManager::Signal_GamepadDisconnected));
}

void QGamepadManager_GamepadDisconnected(void* ptr, int deviceId)
{
	static_cast<QGamepadManager*>(ptr)->gamepadDisconnected(deviceId);
}

void QGamepadManager_ResetConfiguration(void* ptr, int deviceId)
{
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "resetConfiguration", Q_ARG(int, deviceId));
}

void QGamepadManager_ResetConfigurationDefault(void* ptr, int deviceId)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::resetConfiguration(deviceId);
}

void QGamepadManager_SetSettingsFile(void* ptr, struct QtGamepad_PackedString file)
{
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "setSettingsFile", Q_ARG(const QString, QString::fromUtf8(file.data, file.len)));
}

void QGamepadManager_SetSettingsFileDefault(void* ptr, struct QtGamepad_PackedString file)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::setSettingsFile(QString::fromUtf8(file.data, file.len));
}

char QGamepadManager_IsConfigurationNeeded(void* ptr, int deviceId)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGamepadManager*>(ptr), "isConfigurationNeeded", Q_RETURN_ARG(bool, returnArg), Q_ARG(int, deviceId));
	return returnArg;
}

char QGamepadManager_IsConfigurationNeededDefault(void* ptr, int deviceId)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::isConfigurationNeeded(deviceId);
}

char QGamepadManager_IsGamepadConnected(void* ptr, int deviceId)
{
	return static_cast<QGamepadManager*>(ptr)->isGamepadConnected(deviceId);
}

struct QtGamepad_PackedList QGamepadManager_ConnectedGamepads(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QGamepadManager*>(ptr)->connectedGamepads()); QtGamepad_PackedList { tmpValue, tmpValue->size() }; });
}

int QGamepadManager___connectedGamepads_atList(void* ptr, int i)
{
	return static_cast<QList<int>*>(ptr)->at(i);
}

void QGamepadManager___connectedGamepads_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QGamepadManager___connectedGamepads_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>;
}

void* QGamepadManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGamepadManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGamepadManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGamepadManager___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadManager___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadManager___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGamepadManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGamepadManager___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGamepadManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGamepadManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGamepadManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::event(static_cast<QEvent*>(e));
}

char QGamepadManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGamepadManager*>(ptr)->QGamepadManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGamepadManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::childEvent(static_cast<QChildEvent*>(event));
}

void QGamepadManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepadManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::customEvent(static_cast<QEvent*>(event));
}

void QGamepadManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::deleteLater();
}

void QGamepadManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGamepadManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGamepadManager*>(ptr)->QGamepadManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGamepadManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGamepadManager*>(ptr)->QGamepadManager::metaObject());
}

