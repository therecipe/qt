// +build !minimal

#define protected public
#define private public

#include "sensors.h"
#include "_cgo_export.h"

#include <QAccelerometer>
#include <QAccelerometerFilter>
#include <QAccelerometerReading>
#include <QAltimeter>
#include <QAltimeterFilter>
#include <QAltimeterReading>
#include <QAmbientLightFilter>
#include <QAmbientLightReading>
#include <QAmbientLightSensor>
#include <QAmbientTemperatureFilter>
#include <QAmbientTemperatureReading>
#include <QAmbientTemperatureSensor>
#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QCompass>
#include <QCompassFilter>
#include <QCompassReading>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDistanceFilter>
#include <QDistanceReading>
#include <QDistanceSensor>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QGyroscope>
#include <QGyroscopeFilter>
#include <QGyroscopeReading>
#include <QHolsterFilter>
#include <QHolsterReading>
#include <QHolsterSensor>
#include <QHumidityFilter>
#include <QHumidityReading>
#include <QHumiditySensor>
#include <QIRProximityFilter>
#include <QIRProximityReading>
#include <QIRProximitySensor>
#include <QLayout>
#include <QLidFilter>
#include <QLidReading>
#include <QLidSensor>
#include <QLightFilter>
#include <QLightReading>
#include <QLightSensor>
#include <QList>
#include <QMagnetometer>
#include <QMagnetometerFilter>
#include <QMagnetometerReading>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QOrientationFilter>
#include <QOrientationReading>
#include <QOrientationSensor>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPressureFilter>
#include <QPressureReading>
#include <QPressureSensor>
#include <QProximityFilter>
#include <QProximityReading>
#include <QProximitySensor>
#include <QQuickItem>
#include <QRadioData>
#include <QRotationFilter>
#include <QRotationReading>
#include <QRotationSensor>
#include <QSensor>
#include <QSensorBackend>
#include <QSensorBackendFactory>
#include <QSensorChangesInterface>
#include <QSensorFilter>
#include <QSensorGesture>
#include <QSensorGestureManager>
#include <QSensorGesturePluginInterface>
#include <QSensorGestureRecognizer>
#include <QSensorManager>
#include <QSensorPluginInterface>
#include <QSensorReading>
#include <QSignalSpy>
#include <QString>
#include <QStringList>
#include <QTapFilter>
#include <QTapReading>
#include <QTapSensor>
#include <QTiltFilter>
#include <QTiltReading>
#include <QTiltSensor>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

class MyQAccelerometer: public QAccelerometer
{
public:
	MyQAccelerometer(QObject *parent = Q_NULLPTR) : QAccelerometer(parent) {QAccelerometer_QAccelerometer_QRegisterMetaType();};
	void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode) { callbackQAccelerometer_AccelerationModeChanged(this, accelerationMode); };
	 ~MyQAccelerometer() { callbackQAccelerometer_DestroyQAccelerometer(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAccelerometer*)

int QAccelerometer_QAccelerometer_QRegisterMetaType(){qRegisterMetaType<QAccelerometer*>(); return qRegisterMetaType<MyQAccelerometer*>();}

void QAccelerometer_ConnectAccelerationModeChanged(void* ptr)
{
	QObject::connect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));
}

void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));
}

void QAccelerometer_AccelerationModeChanged(void* ptr, long long accelerationMode)
{
	static_cast<QAccelerometer*>(ptr)->accelerationModeChanged(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_SetAccelerationMode(void* ptr, long long accelerationMode)
{
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

long long QAccelerometer_AccelerationMode(void* ptr)
{
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

void* QAccelerometer_NewQAccelerometer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAccelerometer(static_cast<QWindow*>(parent));
	} else {
		return new MyQAccelerometer(static_cast<QObject*>(parent));
	}
}

void QAccelerometer_DestroyQAccelerometer(void* ptr)
{
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
}

void QAccelerometer_DestroyQAccelerometerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAccelerometer_Reading(void* ptr)
{
	return static_cast<QAccelerometer*>(ptr)->reading();
}

struct QtSensors_PackedString QAccelerometer_QAccelerometer_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QAccelerometer::type), -1 };
}

class MyQAccelerometerFilter: public QAccelerometerFilter
{
public:
	bool filter(QAccelerometerReading * reading) { return callbackQAccelerometerFilter_Filter(this, reading) != 0; };
};

char QAccelerometerFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

class MyQAccelerometerReading: public QAccelerometerReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAccelerometerReading*)

int QAccelerometerReading_QAccelerometerReading_QRegisterMetaType(){qRegisterMetaType<QAccelerometerReading*>(); return qRegisterMetaType<MyQAccelerometerReading*>();}

void QAccelerometerReading_SetX(void* ptr, double x)
{
	static_cast<QAccelerometerReading*>(ptr)->setX(x);
}

void QAccelerometerReading_SetY(void* ptr, double y)
{
	static_cast<QAccelerometerReading*>(ptr)->setY(y);
}

void QAccelerometerReading_SetZ(void* ptr, double z)
{
	static_cast<QAccelerometerReading*>(ptr)->setZ(z);
}

double QAccelerometerReading_X(void* ptr)
{
	return static_cast<QAccelerometerReading*>(ptr)->x();
}

double QAccelerometerReading_Y(void* ptr)
{
	return static_cast<QAccelerometerReading*>(ptr)->y();
}

double QAccelerometerReading_Z(void* ptr)
{
	return static_cast<QAccelerometerReading*>(ptr)->z();
}

class MyQAltimeter: public QAltimeter
{
public:
	MyQAltimeter(QObject *parent = Q_NULLPTR) : QAltimeter(parent) {QAltimeter_QAltimeter_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAltimeter*)

int QAltimeter_QAltimeter_QRegisterMetaType(){qRegisterMetaType<QAltimeter*>(); return qRegisterMetaType<MyQAltimeter*>();}

void* QAltimeter_NewQAltimeter(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAltimeter(static_cast<QWindow*>(parent));
	} else {
		return new MyQAltimeter(static_cast<QObject*>(parent));
	}
}

void QAltimeter_DestroyQAltimeter(void* ptr)
{
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

void* QAltimeter_Reading(void* ptr)
{
	return static_cast<QAltimeter*>(ptr)->reading();
}

struct QtSensors_PackedString QAltimeter_QAltimeter_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QAltimeter::type), -1 };
}

class MyQAltimeterFilter: public QAltimeterFilter
{
public:
	bool filter(QAltimeterReading * reading) { return callbackQAltimeterFilter_Filter(this, reading) != 0; };
};

char QAltimeterFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

class MyQAltimeterReading: public QAltimeterReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAltimeterReading*)

int QAltimeterReading_QAltimeterReading_QRegisterMetaType(){qRegisterMetaType<QAltimeterReading*>(); return qRegisterMetaType<MyQAltimeterReading*>();}

void QAltimeterReading_SetAltitude(void* ptr, double altitude)
{
	static_cast<QAltimeterReading*>(ptr)->setAltitude(altitude);
}

double QAltimeterReading_Altitude(void* ptr)
{
	return static_cast<QAltimeterReading*>(ptr)->altitude();
}

class MyQAmbientLightFilter: public QAmbientLightFilter
{
public:
	bool filter(QAmbientLightReading * reading) { return callbackQAmbientLightFilter_Filter(this, reading) != 0; };
};

char QAmbientLightFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

class MyQAmbientLightReading: public QAmbientLightReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAmbientLightReading*)

int QAmbientLightReading_QAmbientLightReading_QRegisterMetaType(){qRegisterMetaType<QAmbientLightReading*>(); return qRegisterMetaType<MyQAmbientLightReading*>();}

void QAmbientLightReading_SetLightLevel(void* ptr, long long lightLevel)
{
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

long long QAmbientLightReading_LightLevel(void* ptr)
{
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

class MyQAmbientLightSensor: public QAmbientLightSensor
{
public:
	MyQAmbientLightSensor(QObject *parent = Q_NULLPTR) : QAmbientLightSensor(parent) {QAmbientLightSensor_QAmbientLightSensor_QRegisterMetaType();};
	 ~MyQAmbientLightSensor() { callbackQAmbientLightSensor_DestroyQAmbientLightSensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAmbientLightSensor*)

int QAmbientLightSensor_QAmbientLightSensor_QRegisterMetaType(){qRegisterMetaType<QAmbientLightSensor*>(); return qRegisterMetaType<MyQAmbientLightSensor*>();}

void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientLightSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQAmbientLightSensor(static_cast<QObject*>(parent));
	}
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr)
{
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

void QAmbientLightSensor_DestroyQAmbientLightSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAmbientLightSensor_Reading(void* ptr)
{
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QAmbientLightSensor_QAmbientLightSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QAmbientLightSensor::type), -1 };
}

class MyQAmbientTemperatureFilter: public QAmbientTemperatureFilter
{
public:
	bool filter(QAmbientTemperatureReading * reading) { return callbackQAmbientTemperatureFilter_Filter(this, reading) != 0; };
};

char QAmbientTemperatureFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

class MyQAmbientTemperatureReading: public QAmbientTemperatureReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAmbientTemperatureReading*)

int QAmbientTemperatureReading_QAmbientTemperatureReading_QRegisterMetaType(){qRegisterMetaType<QAmbientTemperatureReading*>(); return qRegisterMetaType<MyQAmbientTemperatureReading*>();}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(temperature);
}

double QAmbientTemperatureReading_Temperature(void* ptr)
{
	return static_cast<QAmbientTemperatureReading*>(ptr)->temperature();
}

class MyQAmbientTemperatureSensor: public QAmbientTemperatureSensor
{
public:
	MyQAmbientTemperatureSensor(QObject *parent = Q_NULLPTR) : QAmbientTemperatureSensor(parent) {QAmbientTemperatureSensor_QAmbientTemperatureSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQAmbientTemperatureSensor*)

int QAmbientTemperatureSensor_QAmbientTemperatureSensor_QRegisterMetaType(){qRegisterMetaType<QAmbientTemperatureSensor*>(); return qRegisterMetaType<MyQAmbientTemperatureSensor*>();}

void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAmbientTemperatureSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQAmbientTemperatureSensor(static_cast<QObject*>(parent));
	}
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

void* QAmbientTemperatureSensor_Reading(void* ptr)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QAmbientTemperatureSensor_QAmbientTemperatureSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QAmbientTemperatureSensor::type), -1 };
}

class MyQCompass: public QCompass
{
public:
	MyQCompass(QObject *parent = Q_NULLPTR) : QCompass(parent) {QCompass_QCompass_QRegisterMetaType();};
	 ~MyQCompass() { callbackQCompass_DestroyQCompass(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQCompass*)

int QCompass_QCompass_QRegisterMetaType(){qRegisterMetaType<QCompass*>(); return qRegisterMetaType<MyQCompass*>();}

void* QCompass_NewQCompass(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQCompass(static_cast<QWindow*>(parent));
	} else {
		return new MyQCompass(static_cast<QObject*>(parent));
	}
}

void QCompass_DestroyQCompass(void* ptr)
{
	static_cast<QCompass*>(ptr)->~QCompass();
}

void QCompass_DestroyQCompassDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QCompass_Reading(void* ptr)
{
	return static_cast<QCompass*>(ptr)->reading();
}

struct QtSensors_PackedString QCompass_QCompass_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QCompass::type), -1 };
}

class MyQCompassFilter: public QCompassFilter
{
public:
	bool filter(QCompassReading * reading) { return callbackQCompassFilter_Filter(this, reading) != 0; };
};

char QCompassFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

class MyQCompassReading: public QCompassReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQCompassReading*)

int QCompassReading_QCompassReading_QRegisterMetaType(){qRegisterMetaType<QCompassReading*>(); return qRegisterMetaType<MyQCompassReading*>();}

void QCompassReading_SetAzimuth(void* ptr, double azimuth)
{
	static_cast<QCompassReading*>(ptr)->setAzimuth(azimuth);
}

void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel)
{
	static_cast<QCompassReading*>(ptr)->setCalibrationLevel(calibrationLevel);
}

double QCompassReading_Azimuth(void* ptr)
{
	return static_cast<QCompassReading*>(ptr)->azimuth();
}

double QCompassReading_CalibrationLevel(void* ptr)
{
	return static_cast<QCompassReading*>(ptr)->calibrationLevel();
}

class MyQDistanceFilter: public QDistanceFilter
{
public:
	bool filter(QDistanceReading * reading) { return callbackQDistanceFilter_Filter(this, reading) != 0; };
};

char QDistanceFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

class MyQDistanceReading: public QDistanceReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDistanceReading*)

int QDistanceReading_QDistanceReading_QRegisterMetaType(){qRegisterMetaType<QDistanceReading*>(); return qRegisterMetaType<MyQDistanceReading*>();}

void QDistanceReading_SetDistance(void* ptr, double distance)
{
	static_cast<QDistanceReading*>(ptr)->setDistance(distance);
}

double QDistanceReading_Distance(void* ptr)
{
	return static_cast<QDistanceReading*>(ptr)->distance();
}

class MyQDistanceSensor: public QDistanceSensor
{
public:
	MyQDistanceSensor(QObject *parent = Q_NULLPTR) : QDistanceSensor(parent) {QDistanceSensor_QDistanceSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQDistanceSensor*)

int QDistanceSensor_QDistanceSensor_QRegisterMetaType(){qRegisterMetaType<QDistanceSensor*>(); return qRegisterMetaType<MyQDistanceSensor*>();}

void* QDistanceSensor_NewQDistanceSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDistanceSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQDistanceSensor(static_cast<QObject*>(parent));
	}
}

void QDistanceSensor_DestroyQDistanceSensor(void* ptr)
{
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

void* QDistanceSensor_Reading(void* ptr)
{
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QDistanceSensor_QDistanceSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QDistanceSensor::type), -1 };
}

class MyQGyroscope: public QGyroscope
{
public:
	MyQGyroscope(QObject *parent = Q_NULLPTR) : QGyroscope(parent) {QGyroscope_QGyroscope_QRegisterMetaType();};
	 ~MyQGyroscope() { callbackQGyroscope_DestroyQGyroscope(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGyroscope*)

int QGyroscope_QGyroscope_QRegisterMetaType(){qRegisterMetaType<QGyroscope*>(); return qRegisterMetaType<MyQGyroscope*>();}

void* QGyroscope_NewQGyroscope(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGyroscope(static_cast<QWindow*>(parent));
	} else {
		return new MyQGyroscope(static_cast<QObject*>(parent));
	}
}

void QGyroscope_DestroyQGyroscope(void* ptr)
{
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

void QGyroscope_DestroyQGyroscopeDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGyroscope_Reading(void* ptr)
{
	return static_cast<QGyroscope*>(ptr)->reading();
}

struct QtSensors_PackedString QGyroscope_QGyroscope_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QGyroscope::type), -1 };
}

class MyQGyroscopeFilter: public QGyroscopeFilter
{
public:
	bool filter(QGyroscopeReading * reading) { return callbackQGyroscopeFilter_Filter(this, reading) != 0; };
};

char QGyroscopeFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

class MyQGyroscopeReading: public QGyroscopeReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGyroscopeReading*)

int QGyroscopeReading_QGyroscopeReading_QRegisterMetaType(){qRegisterMetaType<QGyroscopeReading*>(); return qRegisterMetaType<MyQGyroscopeReading*>();}

void QGyroscopeReading_SetX(void* ptr, double x)
{
	static_cast<QGyroscopeReading*>(ptr)->setX(x);
}

void QGyroscopeReading_SetY(void* ptr, double y)
{
	static_cast<QGyroscopeReading*>(ptr)->setY(y);
}

void QGyroscopeReading_SetZ(void* ptr, double z)
{
	static_cast<QGyroscopeReading*>(ptr)->setZ(z);
}

double QGyroscopeReading_X(void* ptr)
{
	return static_cast<QGyroscopeReading*>(ptr)->x();
}

double QGyroscopeReading_Y(void* ptr)
{
	return static_cast<QGyroscopeReading*>(ptr)->y();
}

double QGyroscopeReading_Z(void* ptr)
{
	return static_cast<QGyroscopeReading*>(ptr)->z();
}

class MyQHolsterFilter: public QHolsterFilter
{
public:
	bool filter(QHolsterReading * reading) { return callbackQHolsterFilter_Filter(this, reading) != 0; };
};

char QHolsterFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

class MyQHolsterReading: public QHolsterReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHolsterReading*)

int QHolsterReading_QHolsterReading_QRegisterMetaType(){qRegisterMetaType<QHolsterReading*>(); return qRegisterMetaType<MyQHolsterReading*>();}

void QHolsterReading_SetHolstered(void* ptr, char holstered)
{
	static_cast<QHolsterReading*>(ptr)->setHolstered(holstered != 0);
}

char QHolsterReading_Holstered(void* ptr)
{
	return static_cast<QHolsterReading*>(ptr)->holstered();
}

class MyQHolsterSensor: public QHolsterSensor
{
public:
	MyQHolsterSensor(QObject *parent = Q_NULLPTR) : QHolsterSensor(parent) {QHolsterSensor_QHolsterSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHolsterSensor*)

int QHolsterSensor_QHolsterSensor_QRegisterMetaType(){qRegisterMetaType<QHolsterSensor*>(); return qRegisterMetaType<MyQHolsterSensor*>();}

void* QHolsterSensor_NewQHolsterSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHolsterSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQHolsterSensor(static_cast<QObject*>(parent));
	}
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr)
{
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

void* QHolsterSensor_Reading(void* ptr)
{
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QHolsterSensor_QHolsterSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QHolsterSensor::type), -1 };
}

class MyQHumidityFilter: public QHumidityFilter
{
public:
	bool filter(QHumidityReading * reading) { return callbackQHumidityFilter_Filter(this, reading) != 0; };
};

char QHumidityFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QHumidityFilter*>(ptr)->filter(static_cast<QHumidityReading*>(reading));
}

class MyQHumidityReading: public QHumidityReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHumidityReading*)

int QHumidityReading_QHumidityReading_QRegisterMetaType(){qRegisterMetaType<QHumidityReading*>(); return qRegisterMetaType<MyQHumidityReading*>();}

void QHumidityReading_SetAbsoluteHumidity(void* ptr, double value)
{
	static_cast<QHumidityReading*>(ptr)->setAbsoluteHumidity(value);
}

void QHumidityReading_SetRelativeHumidity(void* ptr, double humidity)
{
	static_cast<QHumidityReading*>(ptr)->setRelativeHumidity(humidity);
}

double QHumidityReading_AbsoluteHumidity(void* ptr)
{
	return static_cast<QHumidityReading*>(ptr)->absoluteHumidity();
}

double QHumidityReading_RelativeHumidity(void* ptr)
{
	return static_cast<QHumidityReading*>(ptr)->relativeHumidity();
}

class MyQHumiditySensor: public QHumiditySensor
{
public:
	MyQHumiditySensor(QObject *parent = Q_NULLPTR) : QHumiditySensor(parent) {QHumiditySensor_QHumiditySensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQHumiditySensor*)

int QHumiditySensor_QHumiditySensor_QRegisterMetaType(){qRegisterMetaType<QHumiditySensor*>(); return qRegisterMetaType<MyQHumiditySensor*>();}

void* QHumiditySensor_NewQHumiditySensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQHumiditySensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQHumiditySensor(static_cast<QObject*>(parent));
	}
}

void QHumiditySensor_DestroyQHumiditySensor(void* ptr)
{
	static_cast<QHumiditySensor*>(ptr)->~QHumiditySensor();
}

void* QHumiditySensor_Reading(void* ptr)
{
	return static_cast<QHumiditySensor*>(ptr)->reading();
}

struct QtSensors_PackedString QHumiditySensor_QHumiditySensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QHumiditySensor::type), -1 };
}

class MyQIRProximityFilter: public QIRProximityFilter
{
public:
	bool filter(QIRProximityReading * reading) { return callbackQIRProximityFilter_Filter(this, reading) != 0; };
};

char QIRProximityFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

class MyQIRProximityReading: public QIRProximityReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQIRProximityReading*)

int QIRProximityReading_QIRProximityReading_QRegisterMetaType(){qRegisterMetaType<QIRProximityReading*>(); return qRegisterMetaType<MyQIRProximityReading*>();}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance)
{
	static_cast<QIRProximityReading*>(ptr)->setReflectance(reflectance);
}

double QIRProximityReading_Reflectance(void* ptr)
{
	return static_cast<QIRProximityReading*>(ptr)->reflectance();
}

class MyQIRProximitySensor: public QIRProximitySensor
{
public:
	MyQIRProximitySensor(QObject *parent = Q_NULLPTR) : QIRProximitySensor(parent) {QIRProximitySensor_QIRProximitySensor_QRegisterMetaType();};
	 ~MyQIRProximitySensor() { callbackQIRProximitySensor_DestroyQIRProximitySensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQIRProximitySensor*)

int QIRProximitySensor_QIRProximitySensor_QRegisterMetaType(){qRegisterMetaType<QIRProximitySensor*>(); return qRegisterMetaType<MyQIRProximitySensor*>();}

void* QIRProximitySensor_NewQIRProximitySensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIRProximitySensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQIRProximitySensor(static_cast<QObject*>(parent));
	}
}

void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr)
{
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

void QIRProximitySensor_DestroyQIRProximitySensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QIRProximitySensor_Reading(void* ptr)
{
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

struct QtSensors_PackedString QIRProximitySensor_QIRProximitySensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QIRProximitySensor::type), -1 };
}

class MyQLidFilter: public QLidFilter
{
public:
	bool filter(QLidReading * reading) { return callbackQLidFilter_Filter(this, reading) != 0; };
};

char QLidFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QLidFilter*>(ptr)->filter(static_cast<QLidReading*>(reading));
}

class MyQLidReading: public QLidReading
{
public:
	void Signal_BackLidChanged(bool closed) { callbackQLidReading_BackLidChanged(this, closed); };
	void Signal_FrontLidChanged(bool closed) { callbackQLidReading_FrontLidChanged(this, closed); };
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQLidReading*)

int QLidReading_QLidReading_QRegisterMetaType(){qRegisterMetaType<QLidReading*>(); return qRegisterMetaType<MyQLidReading*>();}

void QLidReading_ConnectBackLidChanged(void* ptr)
{
	QObject::connect(static_cast<QLidReading*>(ptr), static_cast<void (QLidReading::*)(bool)>(&QLidReading::backLidChanged), static_cast<MyQLidReading*>(ptr), static_cast<void (MyQLidReading::*)(bool)>(&MyQLidReading::Signal_BackLidChanged));
}

void QLidReading_DisconnectBackLidChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLidReading*>(ptr), static_cast<void (QLidReading::*)(bool)>(&QLidReading::backLidChanged), static_cast<MyQLidReading*>(ptr), static_cast<void (MyQLidReading::*)(bool)>(&MyQLidReading::Signal_BackLidChanged));
}

void QLidReading_BackLidChanged(void* ptr, char closed)
{
	static_cast<QLidReading*>(ptr)->backLidChanged(closed != 0);
}

void QLidReading_ConnectFrontLidChanged(void* ptr)
{
	QObject::connect(static_cast<QLidReading*>(ptr), static_cast<void (QLidReading::*)(bool)>(&QLidReading::frontLidChanged), static_cast<MyQLidReading*>(ptr), static_cast<void (MyQLidReading::*)(bool)>(&MyQLidReading::Signal_FrontLidChanged));
}

void QLidReading_DisconnectFrontLidChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLidReading*>(ptr), static_cast<void (QLidReading::*)(bool)>(&QLidReading::frontLidChanged), static_cast<MyQLidReading*>(ptr), static_cast<void (MyQLidReading::*)(bool)>(&MyQLidReading::Signal_FrontLidChanged));
}

void QLidReading_FrontLidChanged(void* ptr, char closed)
{
	static_cast<QLidReading*>(ptr)->frontLidChanged(closed != 0);
}

void QLidReading_SetBackLidClosed(void* ptr, char closed)
{
	static_cast<QLidReading*>(ptr)->setBackLidClosed(closed != 0);
}

void QLidReading_SetFrontLidClosed(void* ptr, char closed)
{
	static_cast<QLidReading*>(ptr)->setFrontLidClosed(closed != 0);
}

char QLidReading_BackLidClosed(void* ptr)
{
	return static_cast<QLidReading*>(ptr)->backLidClosed();
}

char QLidReading_FrontLidClosed(void* ptr)
{
	return static_cast<QLidReading*>(ptr)->frontLidClosed();
}

class MyQLidSensor: public QLidSensor
{
public:
	MyQLidSensor(QObject *parent = Q_NULLPTR) : QLidSensor(parent) {QLidSensor_QLidSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQLidSensor*)

int QLidSensor_QLidSensor_QRegisterMetaType(){qRegisterMetaType<QLidSensor*>(); return qRegisterMetaType<MyQLidSensor*>();}

void* QLidSensor_NewQLidSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLidSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQLidSensor(static_cast<QObject*>(parent));
	}
}

void QLidSensor_DestroyQLidSensor(void* ptr)
{
	static_cast<QLidSensor*>(ptr)->~QLidSensor();
}

void* QLidSensor_Reading(void* ptr)
{
	return static_cast<QLidSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QLidSensor_QLidSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QLidSensor::type), -1 };
}

class MyQLightFilter: public QLightFilter
{
public:
	bool filter(QLightReading * reading) { return callbackQLightFilter_Filter(this, reading) != 0; };
};

char QLightFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

class MyQLightReading: public QLightReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQLightReading*)

int QLightReading_QLightReading_QRegisterMetaType(){qRegisterMetaType<QLightReading*>(); return qRegisterMetaType<MyQLightReading*>();}

void QLightReading_SetLux(void* ptr, double lux)
{
	static_cast<QLightReading*>(ptr)->setLux(lux);
}

double QLightReading_Lux(void* ptr)
{
	return static_cast<QLightReading*>(ptr)->lux();
}

class MyQLightSensor: public QLightSensor
{
public:
	MyQLightSensor(QObject *parent = Q_NULLPTR) : QLightSensor(parent) {QLightSensor_QLightSensor_QRegisterMetaType();};
	void Signal_FieldOfViewChanged(qreal fieldOfView) { callbackQLightSensor_FieldOfViewChanged(this, fieldOfView); };
	 ~MyQLightSensor() { callbackQLightSensor_DestroyQLightSensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQLightSensor*)

int QLightSensor_QLightSensor_QRegisterMetaType(){qRegisterMetaType<QLightSensor*>(); return qRegisterMetaType<MyQLightSensor*>();}

void* QLightSensor_NewQLightSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQLightSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQLightSensor(static_cast<QObject*>(parent));
	}
}

void QLightSensor_ConnectFieldOfViewChanged(void* ptr)
{
	QObject::connect(static_cast<QLightSensor*>(ptr), static_cast<void (QLightSensor::*)(qreal)>(&QLightSensor::fieldOfViewChanged), static_cast<MyQLightSensor*>(ptr), static_cast<void (MyQLightSensor::*)(qreal)>(&MyQLightSensor::Signal_FieldOfViewChanged));
}

void QLightSensor_DisconnectFieldOfViewChanged(void* ptr)
{
	QObject::disconnect(static_cast<QLightSensor*>(ptr), static_cast<void (QLightSensor::*)(qreal)>(&QLightSensor::fieldOfViewChanged), static_cast<MyQLightSensor*>(ptr), static_cast<void (MyQLightSensor::*)(qreal)>(&MyQLightSensor::Signal_FieldOfViewChanged));
}

void QLightSensor_FieldOfViewChanged(void* ptr, double fieldOfView)
{
	static_cast<QLightSensor*>(ptr)->fieldOfViewChanged(fieldOfView);
}

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView)
{
	static_cast<QLightSensor*>(ptr)->setFieldOfView(fieldOfView);
}

void QLightSensor_DestroyQLightSensor(void* ptr)
{
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

void QLightSensor_DestroyQLightSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QLightSensor_Reading(void* ptr)
{
	return static_cast<QLightSensor*>(ptr)->reading();
}

double QLightSensor_FieldOfView(void* ptr)
{
	return static_cast<QLightSensor*>(ptr)->fieldOfView();
}

struct QtSensors_PackedString QLightSensor_QLightSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QLightSensor::type), -1 };
}

class MyQMagnetometer: public QMagnetometer
{
public:
	MyQMagnetometer(QObject *parent = Q_NULLPTR) : QMagnetometer(parent) {QMagnetometer_QMagnetometer_QRegisterMetaType();};
	void Signal_ReturnGeoValuesChanged(bool returnGeoValues) { callbackQMagnetometer_ReturnGeoValuesChanged(this, returnGeoValues); };
	 ~MyQMagnetometer() { callbackQMagnetometer_DestroyQMagnetometer(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQMagnetometer*)

int QMagnetometer_QMagnetometer_QRegisterMetaType(){qRegisterMetaType<QMagnetometer*>(); return qRegisterMetaType<MyQMagnetometer*>();}

void* QMagnetometer_NewQMagnetometer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQMagnetometer(static_cast<QWindow*>(parent));
	} else {
		return new MyQMagnetometer(static_cast<QObject*>(parent));
	}
}

void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr)
{
	QObject::connect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));
}

void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));
}

void QMagnetometer_ReturnGeoValuesChanged(void* ptr, char returnGeoValues)
{
	static_cast<QMagnetometer*>(ptr)->returnGeoValuesChanged(returnGeoValues != 0);
}

void QMagnetometer_SetReturnGeoValues(void* ptr, char returnGeoValues)
{
	static_cast<QMagnetometer*>(ptr)->setReturnGeoValues(returnGeoValues != 0);
}

void QMagnetometer_DestroyQMagnetometer(void* ptr)
{
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
}

void QMagnetometer_DestroyQMagnetometerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMagnetometer_Reading(void* ptr)
{
	return static_cast<QMagnetometer*>(ptr)->reading();
}

char QMagnetometer_ReturnGeoValues(void* ptr)
{
	return static_cast<QMagnetometer*>(ptr)->returnGeoValues();
}

struct QtSensors_PackedString QMagnetometer_QMagnetometer_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QMagnetometer::type), -1 };
}

class MyQMagnetometerFilter: public QMagnetometerFilter
{
public:
	bool filter(QMagnetometerReading * reading) { return callbackQMagnetometerFilter_Filter(this, reading) != 0; };
};

char QMagnetometerFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

class MyQMagnetometerReading: public QMagnetometerReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQMagnetometerReading*)

int QMagnetometerReading_QMagnetometerReading_QRegisterMetaType(){qRegisterMetaType<QMagnetometerReading*>(); return qRegisterMetaType<MyQMagnetometerReading*>();}

void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel)
{
	static_cast<QMagnetometerReading*>(ptr)->setCalibrationLevel(calibrationLevel);
}

void QMagnetometerReading_SetX(void* ptr, double x)
{
	static_cast<QMagnetometerReading*>(ptr)->setX(x);
}

void QMagnetometerReading_SetY(void* ptr, double y)
{
	static_cast<QMagnetometerReading*>(ptr)->setY(y);
}

void QMagnetometerReading_SetZ(void* ptr, double z)
{
	static_cast<QMagnetometerReading*>(ptr)->setZ(z);
}

double QMagnetometerReading_CalibrationLevel(void* ptr)
{
	return static_cast<QMagnetometerReading*>(ptr)->calibrationLevel();
}

double QMagnetometerReading_X(void* ptr)
{
	return static_cast<QMagnetometerReading*>(ptr)->x();
}

double QMagnetometerReading_Y(void* ptr)
{
	return static_cast<QMagnetometerReading*>(ptr)->y();
}

double QMagnetometerReading_Z(void* ptr)
{
	return static_cast<QMagnetometerReading*>(ptr)->z();
}

class MyQOrientationFilter: public QOrientationFilter
{
public:
	bool filter(QOrientationReading * reading) { return callbackQOrientationFilter_Filter(this, reading) != 0; };
};

char QOrientationFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QOrientationFilter*>(ptr)->filter(static_cast<QOrientationReading*>(reading));
}

class MyQOrientationReading: public QOrientationReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQOrientationReading*)

int QOrientationReading_QOrientationReading_QRegisterMetaType(){qRegisterMetaType<QOrientationReading*>(); return qRegisterMetaType<MyQOrientationReading*>();}

void QOrientationReading_SetOrientation(void* ptr, long long orientation)
{
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

long long QOrientationReading_Orientation(void* ptr)
{
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

class MyQOrientationSensor: public QOrientationSensor
{
public:
	MyQOrientationSensor(QObject *parent = Q_NULLPTR) : QOrientationSensor(parent) {QOrientationSensor_QOrientationSensor_QRegisterMetaType();};
	 ~MyQOrientationSensor() { callbackQOrientationSensor_DestroyQOrientationSensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQOrientationSensor*)

int QOrientationSensor_QOrientationSensor_QRegisterMetaType(){qRegisterMetaType<QOrientationSensor*>(); return qRegisterMetaType<MyQOrientationSensor*>();}

void* QOrientationSensor_NewQOrientationSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQOrientationSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQOrientationSensor(static_cast<QObject*>(parent));
	}
}

void QOrientationSensor_DestroyQOrientationSensor(void* ptr)
{
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

void QOrientationSensor_DestroyQOrientationSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QOrientationSensor_Reading(void* ptr)
{
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QOrientationSensor_QOrientationSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QOrientationSensor::type), -1 };
}

class MyQPressureFilter: public QPressureFilter
{
public:
	bool filter(QPressureReading * reading) { return callbackQPressureFilter_Filter(this, reading) != 0; };
};

char QPressureFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QPressureFilter*>(ptr)->filter(static_cast<QPressureReading*>(reading));
}

class MyQPressureReading: public QPressureReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPressureReading*)

int QPressureReading_QPressureReading_QRegisterMetaType(){qRegisterMetaType<QPressureReading*>(); return qRegisterMetaType<MyQPressureReading*>();}

void QPressureReading_SetPressure(void* ptr, double pressure)
{
	static_cast<QPressureReading*>(ptr)->setPressure(pressure);
}

void QPressureReading_SetTemperature(void* ptr, double temperature)
{
	static_cast<QPressureReading*>(ptr)->setTemperature(temperature);
}

double QPressureReading_Pressure(void* ptr)
{
	return static_cast<QPressureReading*>(ptr)->pressure();
}

double QPressureReading_Temperature(void* ptr)
{
	return static_cast<QPressureReading*>(ptr)->temperature();
}

class MyQPressureSensor: public QPressureSensor
{
public:
	MyQPressureSensor(QObject *parent = Q_NULLPTR) : QPressureSensor(parent) {QPressureSensor_QPressureSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQPressureSensor*)

int QPressureSensor_QPressureSensor_QRegisterMetaType(){qRegisterMetaType<QPressureSensor*>(); return qRegisterMetaType<MyQPressureSensor*>();}

void* QPressureSensor_NewQPressureSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQPressureSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQPressureSensor(static_cast<QObject*>(parent));
	}
}

void QPressureSensor_DestroyQPressureSensor(void* ptr)
{
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

void* QPressureSensor_Reading(void* ptr)
{
	return static_cast<QPressureSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QPressureSensor_QPressureSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QPressureSensor::type), -1 };
}

class MyQProximityFilter: public QProximityFilter
{
public:
	bool filter(QProximityReading * reading) { return callbackQProximityFilter_Filter(this, reading) != 0; };
};

char QProximityFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

class MyQProximityReading: public QProximityReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQProximityReading*)

int QProximityReading_QProximityReading_QRegisterMetaType(){qRegisterMetaType<QProximityReading*>(); return qRegisterMetaType<MyQProximityReading*>();}

void QProximityReading_SetClose(void* ptr, char close)
{
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

char QProximityReading_Close(void* ptr)
{
	return static_cast<QProximityReading*>(ptr)->close();
}

class MyQProximitySensor: public QProximitySensor
{
public:
	MyQProximitySensor(QObject *parent = Q_NULLPTR) : QProximitySensor(parent) {QProximitySensor_QProximitySensor_QRegisterMetaType();};
	 ~MyQProximitySensor() { callbackQProximitySensor_DestroyQProximitySensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQProximitySensor*)

int QProximitySensor_QProximitySensor_QRegisterMetaType(){qRegisterMetaType<QProximitySensor*>(); return qRegisterMetaType<MyQProximitySensor*>();}

void* QProximitySensor_NewQProximitySensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQProximitySensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQProximitySensor(static_cast<QObject*>(parent));
	}
}

void QProximitySensor_DestroyQProximitySensor(void* ptr)
{
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

void QProximitySensor_DestroyQProximitySensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QProximitySensor_Reading(void* ptr)
{
	return static_cast<QProximitySensor*>(ptr)->reading();
}

struct QtSensors_PackedString QProximitySensor_QProximitySensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QProximitySensor::type), -1 };
}

class MyQRotationFilter: public QRotationFilter
{
public:
	bool filter(QRotationReading * reading) { return callbackQRotationFilter_Filter(this, reading) != 0; };
};

char QRotationFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

class MyQRotationReading: public QRotationReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRotationReading*)

int QRotationReading_QRotationReading_QRegisterMetaType(){qRegisterMetaType<QRotationReading*>(); return qRegisterMetaType<MyQRotationReading*>();}

void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z)
{
	static_cast<QRotationReading*>(ptr)->setFromEuler(x, y, z);
}

double QRotationReading_X(void* ptr)
{
	return static_cast<QRotationReading*>(ptr)->x();
}

double QRotationReading_Y(void* ptr)
{
	return static_cast<QRotationReading*>(ptr)->y();
}

double QRotationReading_Z(void* ptr)
{
	return static_cast<QRotationReading*>(ptr)->z();
}

class MyQRotationSensor: public QRotationSensor
{
public:
	MyQRotationSensor(QObject *parent = Q_NULLPTR) : QRotationSensor(parent) {QRotationSensor_QRotationSensor_QRegisterMetaType();};
	void Signal_HasZChanged(bool hasZ) { callbackQRotationSensor_HasZChanged(this, hasZ); };
	 ~MyQRotationSensor() { callbackQRotationSensor_DestroyQRotationSensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQRotationSensor*)

int QRotationSensor_QRotationSensor_QRegisterMetaType(){qRegisterMetaType<QRotationSensor*>(); return qRegisterMetaType<MyQRotationSensor*>();}

void* QRotationSensor_NewQRotationSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRotationSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQRotationSensor(static_cast<QObject*>(parent));
	}
}

void QRotationSensor_ConnectHasZChanged(void* ptr)
{
	QObject::connect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));
}

void QRotationSensor_DisconnectHasZChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));
}

void QRotationSensor_HasZChanged(void* ptr, char hasZ)
{
	static_cast<QRotationSensor*>(ptr)->hasZChanged(hasZ != 0);
}

void QRotationSensor_SetHasZ(void* ptr, char hasZ)
{
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(void* ptr)
{
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

void QRotationSensor_DestroyQRotationSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QRotationSensor_Reading(void* ptr)
{
	return static_cast<QRotationSensor*>(ptr)->reading();
}

char QRotationSensor_HasZ(void* ptr)
{
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

struct QtSensors_PackedString QRotationSensor_QRotationSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QRotationSensor::type), -1 };
}

class MyQSensor: public QSensor
{
public:
	MyQSensor(const QByteArray &ty, QObject *parent = Q_NULLPTR) : QSensor(ty, parent) {QSensor_QSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	 ~MyQSensor() { callbackQSensor_DestroyQSensor(this); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensor*)

int QSensor_QSensor_QRegisterMetaType(){qRegisterMetaType<QSensor*>(); return qRegisterMetaType<MyQSensor*>();}

void* QSensor_QSensor_DefaultSensorForType(void* ty)
{
	return new QByteArray(QSensor::defaultSensorForType(*static_cast<QByteArray*>(ty)));
}

struct QtSensors_PackedList QSensor_QSensor_SensorTypes()
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(QSensor::sensorTypes()); QtSensors_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtSensors_PackedList QSensor_QSensor_SensorsForType(void* ty)
{
	return ({ QList<QByteArray>* tmpValue = new QList<QByteArray>(QSensor::sensorsForType(*static_cast<QByteArray*>(ty))); QtSensors_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSensor_NewQSensor(void* ty, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QWindow*>(parent));
	} else {
		return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QObject*>(parent));
	}
}

char QSensor_ConnectToBackend(void* ptr)
{
	return static_cast<QSensor*>(ptr)->connectToBackend();
}

char QSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QSensor_StartDefault(void* ptr)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTiltSensor*>(ptr)->QTiltSensor::start();
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTapSensor*>(ptr)->QTapSensor::start();
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRotationSensor*>(ptr)->QRotationSensor::start();
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QProximitySensor*>(ptr)->QProximitySensor::start();
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPressureSensor*>(ptr)->QPressureSensor::start();
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::start();
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMagnetometer*>(ptr)->QMagnetometer::start();
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLightSensor*>(ptr)->QLightSensor::start();
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLidSensor*>(ptr)->QLidSensor::start();
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::start();
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::start();
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::start();
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGyroscope*>(ptr)->QGyroscope::start();
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::start();
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCompass*>(ptr)->QCompass::start();
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::start();
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::start();
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAltimeter*>(ptr)->QAltimeter::start();
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAccelerometer*>(ptr)->QAccelerometer::start();
	} else {
		return static_cast<QSensor*>(ptr)->QSensor::start();
	}
}

void QSensor_ConnectActiveChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));
}

void QSensor_DisconnectActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));
}

void QSensor_ActiveChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->activeChanged();
}

void QSensor_AddFilter(void* ptr, void* filter)
{
	static_cast<QSensor*>(ptr)->addFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectAlwaysOnChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));
}

void QSensor_DisconnectAlwaysOnChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));
}

void QSensor_AlwaysOnChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->alwaysOnChanged();
}

void QSensor_ConnectAvailableSensorsChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));
}

void QSensor_DisconnectAvailableSensorsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));
}

void QSensor_AvailableSensorsChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->availableSensorsChanged();
}

void QSensor_ConnectAxesOrientationModeChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));
}

void QSensor_DisconnectAxesOrientationModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));
}

void QSensor_AxesOrientationModeChanged(void* ptr, long long axesOrientationMode)
{
	static_cast<QSensor*>(ptr)->axesOrientationModeChanged(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_ConnectBufferSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));
}

void QSensor_DisconnectBufferSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));
}

void QSensor_BufferSizeChanged(void* ptr, int bufferSize)
{
	static_cast<QSensor*>(ptr)->bufferSizeChanged(bufferSize);
}

void QSensor_ConnectBusyChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));
}

void QSensor_DisconnectBusyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));
}

void QSensor_BusyChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->busyChanged();
}

void QSensor_ConnectCurrentOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));
}

void QSensor_DisconnectCurrentOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));
}

void QSensor_CurrentOrientationChanged(void* ptr, int currentOrientation)
{
	static_cast<QSensor*>(ptr)->currentOrientationChanged(currentOrientation);
}

void QSensor_ConnectDataRateChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));
}

void QSensor_DisconnectDataRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));
}

void QSensor_DataRateChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->dataRateChanged();
}

void QSensor_ConnectEfficientBufferSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));
}

void QSensor_DisconnectEfficientBufferSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));
}

void QSensor_EfficientBufferSizeChanged(void* ptr, int efficientBufferSize)
{
	static_cast<QSensor*>(ptr)->efficientBufferSizeChanged(efficientBufferSize);
}

void QSensor_ConnectMaxBufferSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));
}

void QSensor_DisconnectMaxBufferSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));
}

void QSensor_MaxBufferSizeChanged(void* ptr, int maxBufferSize)
{
	static_cast<QSensor*>(ptr)->maxBufferSizeChanged(maxBufferSize);
}

void QSensor_ConnectReadingChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));
}

void QSensor_DisconnectReadingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));
}

void QSensor_ReadingChanged(void* ptr)
{
	static_cast<QSensor*>(ptr)->readingChanged();
}

void QSensor_RemoveFilter(void* ptr, void* filter)
{
	static_cast<QSensor*>(ptr)->removeFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectSensorError(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));
}

void QSensor_DisconnectSensorError(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));
}

void QSensor_SensorError(void* ptr, int error)
{
	static_cast<QSensor*>(ptr)->sensorError(error);
}

void QSensor_SetActive(void* ptr, char active)
{
	static_cast<QSensor*>(ptr)->setActive(active != 0);
}

void QSensor_SetAlwaysOn(void* ptr, char alwaysOn)
{
	static_cast<QSensor*>(ptr)->setAlwaysOn(alwaysOn != 0);
}

void QSensor_SetAxesOrientationMode(void* ptr, long long axesOrientationMode)
{
	static_cast<QSensor*>(ptr)->setAxesOrientationMode(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_SetBufferSize(void* ptr, int bufferSize)
{
	static_cast<QSensor*>(ptr)->setBufferSize(bufferSize);
}

void QSensor_SetCurrentOrientation(void* ptr, int currentOrientation)
{
	static_cast<QSensor*>(ptr)->setCurrentOrientation(currentOrientation);
}

void QSensor_SetDataRate(void* ptr, int rate)
{
	static_cast<QSensor*>(ptr)->setDataRate(rate);
}

void QSensor_SetEfficientBufferSize(void* ptr, int efficientBufferSize)
{
	static_cast<QSensor*>(ptr)->setEfficientBufferSize(efficientBufferSize);
}

void QSensor_SetIdentifier(void* ptr, void* identifier)
{
	static_cast<QSensor*>(ptr)->setIdentifier(*static_cast<QByteArray*>(identifier));
}

void QSensor_SetMaxBufferSize(void* ptr, int maxBufferSize)
{
	static_cast<QSensor*>(ptr)->setMaxBufferSize(maxBufferSize);
}

void QSensor_SetOutputRange(void* ptr, int index)
{
	static_cast<QSensor*>(ptr)->setOutputRange(index);
}

void QSensor_SetSkipDuplicates(void* ptr, char skipDuplicates)
{
	static_cast<QSensor*>(ptr)->setSkipDuplicates(skipDuplicates != 0);
}

void QSensor_SetUserOrientation(void* ptr, int userOrientation)
{
	static_cast<QSensor*>(ptr)->setUserOrientation(userOrientation);
}

void QSensor_ConnectSkipDuplicatesChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));
}

void QSensor_DisconnectSkipDuplicatesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));
}

void QSensor_SkipDuplicatesChanged(void* ptr, char skipDuplicates)
{
	static_cast<QSensor*>(ptr)->skipDuplicatesChanged(skipDuplicates != 0);
}

void QSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "stop");
}

void QSensor_StopDefault(void* ptr)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::stop();
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::stop();
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::stop();
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::stop();
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::stop();
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::stop();
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::stop();
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::stop();
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::stop();
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::stop();
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::stop();
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::stop();
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::stop();
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::stop();
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::stop();
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::stop();
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::stop();
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::stop();
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::stop();
	} else {
		static_cast<QSensor*>(ptr)->QSensor::stop();
	}
}

void QSensor_ConnectUserOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));
}

void QSensor_DisconnectUserOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));
}

void QSensor_UserOrientationChanged(void* ptr, int userOrientation)
{
	static_cast<QSensor*>(ptr)->userOrientationChanged(userOrientation);
}

void QSensor_DestroyQSensor(void* ptr)
{
	static_cast<QSensor*>(ptr)->~QSensor();
}

void QSensor_DestroyQSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QSensor_AxesOrientationMode(void* ptr)
{
	return static_cast<QSensor*>(ptr)->axesOrientationMode();
}

void* QSensor_Identifier(void* ptr)
{
	return new QByteArray(static_cast<QSensor*>(ptr)->identifier());
}

void* QSensor_Type(void* ptr)
{
	return new QByteArray(static_cast<QSensor*>(ptr)->type());
}

struct QtSensors_PackedList QSensor_Filters(void* ptr)
{
	return ({ QList<QSensorFilter *>* tmpValue = new QList<QSensorFilter *>(static_cast<QSensor*>(ptr)->filters()); QtSensors_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSensor_Reading(void* ptr)
{
	return static_cast<QSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QSensor_Description(void* ptr)
{
	return ({ QByteArray te03b11 = static_cast<QSensor*>(ptr)->description().toUtf8(); QtSensors_PackedString { const_cast<char*>(te03b11.prepend("WHITESPACE").constData()+10), te03b11.size()-10 }; });
}

char QSensor_IsActive(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isActive();
}

char QSensor_IsAlwaysOn(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isAlwaysOn();
}

char QSensor_IsBusy(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isBusy();
}

char QSensor_IsConnectedToBackend(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isConnectedToBackend();
}

char QSensor_IsFeatureSupported(void* ptr, long long feature)
{
	return static_cast<QSensor*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

char QSensor_SkipDuplicates(void* ptr)
{
	return static_cast<QSensor*>(ptr)->skipDuplicates();
}

int QSensor_BufferSize(void* ptr)
{
	return static_cast<QSensor*>(ptr)->bufferSize();
}

int QSensor_CurrentOrientation(void* ptr)
{
	return static_cast<QSensor*>(ptr)->currentOrientation();
}

int QSensor_DataRate(void* ptr)
{
	return static_cast<QSensor*>(ptr)->dataRate();
}

int QSensor_EfficientBufferSize(void* ptr)
{
	return static_cast<QSensor*>(ptr)->efficientBufferSize();
}

int QSensor_Error(void* ptr)
{
	return static_cast<QSensor*>(ptr)->error();
}

int QSensor_MaxBufferSize(void* ptr)
{
	return static_cast<QSensor*>(ptr)->maxBufferSize();
}

int QSensor_OutputRange(void* ptr)
{
	return static_cast<QSensor*>(ptr)->outputRange();
}

int QSensor_UserOrientation(void* ptr)
{
	return static_cast<QSensor*>(ptr)->userOrientation();
}

void* QSensor___sensorTypes_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensor___sensorTypes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensor___sensorTypes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensor___sensorsForType_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensor___sensorsForType_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensor___sensorsForType_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensor___filters_atList(void* ptr, int i)
{
	return const_cast<QSensorFilter*>(static_cast<QList<QSensorFilter *>*>(ptr)->at(i));
}

void QSensor___filters_setList(void* ptr, void* i)
{
	static_cast<QList<QSensorFilter *>*>(ptr)->append(static_cast<QSensorFilter*>(i));
}

void* QSensor___filters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSensorFilter *>;
}

void* QSensor___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensor___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensor___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensor___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensor___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensor___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensor___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensor___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensor___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensor___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensor___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensor___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensor___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensor___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensor___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensor_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTiltSensor*>(ptr)->QTiltSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTapSensor*>(ptr)->QTapSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRotationSensor*>(ptr)->QRotationSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QProximitySensor*>(ptr)->QProximitySensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPressureSensor*>(ptr)->QPressureSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMagnetometer*>(ptr)->QMagnetometer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLightSensor*>(ptr)->QLightSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLidSensor*>(ptr)->QLidSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGyroscope*>(ptr)->QGyroscope::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCompass*>(ptr)->QCompass::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAltimeter*>(ptr)->QAltimeter::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAccelerometer*>(ptr)->QAccelerometer::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QSensor*>(ptr)->QSensor::event(static_cast<QEvent*>(e));
	}
}

char QSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTiltSensor*>(ptr)->QTiltSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTapSensor*>(ptr)->QTapSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRotationSensor*>(ptr)->QRotationSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QProximitySensor*>(ptr)->QProximitySensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPressureSensor*>(ptr)->QPressureSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMagnetometer*>(ptr)->QMagnetometer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLightSensor*>(ptr)->QLightSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLidSensor*>(ptr)->QLidSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGyroscope*>(ptr)->QGyroscope::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCompass*>(ptr)->QCompass::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAltimeter*>(ptr)->QAltimeter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAccelerometer*>(ptr)->QAccelerometer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSensor*>(ptr)->QSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QSensor_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QSensor*>(ptr)->QSensor::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSensor*>(ptr)->QSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSensor_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QSensor*>(ptr)->QSensor::customEvent(static_cast<QEvent*>(event));
	}
}

void QSensor_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::deleteLater();
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::deleteLater();
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::deleteLater();
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::deleteLater();
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::deleteLater();
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::deleteLater();
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::deleteLater();
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::deleteLater();
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::deleteLater();
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::deleteLater();
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::deleteLater();
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::deleteLater();
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::deleteLater();
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::deleteLater();
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::deleteLater();
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::deleteLater();
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::deleteLater();
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::deleteLater();
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::deleteLater();
	} else {
		static_cast<QSensor*>(ptr)->QSensor::deleteLater();
	}
}

void QSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSensor*>(ptr)->QSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSensor_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltSensor*>(ptr)->QTiltSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapSensor*>(ptr)->QTapSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationSensor*>(ptr)->QRotationSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximitySensor*>(ptr)->QProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureSensor*>(ptr)->QPressureSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometer*>(ptr)->QMagnetometer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightSensor*>(ptr)->QLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidSensor*>(ptr)->QLidSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscope*>(ptr)->QGyroscope::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompass*>(ptr)->QCompass::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeter*>(ptr)->QAltimeter::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometer*>(ptr)->QAccelerometer::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QSensor*>(ptr)->QSensor::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QSensor_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QTiltSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QTiltSensor*>(ptr)->QTiltSensor::metaObject());
	} else if (dynamic_cast<QTapSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QTapSensor*>(ptr)->QTapSensor::metaObject());
	} else if (dynamic_cast<QRotationSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRotationSensor*>(ptr)->QRotationSensor::metaObject());
	} else if (dynamic_cast<QProximitySensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QProximitySensor*>(ptr)->QProximitySensor::metaObject());
	} else if (dynamic_cast<QPressureSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPressureSensor*>(ptr)->QPressureSensor::metaObject());
	} else if (dynamic_cast<QOrientationSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::metaObject());
	} else if (dynamic_cast<QMagnetometer*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QMagnetometer*>(ptr)->QMagnetometer::metaObject());
	} else if (dynamic_cast<QLightSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLightSensor*>(ptr)->QLightSensor::metaObject());
	} else if (dynamic_cast<QLidSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLidSensor*>(ptr)->QLidSensor::metaObject());
	} else if (dynamic_cast<QIRProximitySensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::metaObject());
	} else if (dynamic_cast<QHumiditySensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHumiditySensor*>(ptr)->QHumiditySensor::metaObject());
	} else if (dynamic_cast<QHolsterSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::metaObject());
	} else if (dynamic_cast<QGyroscope*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QGyroscope*>(ptr)->QGyroscope::metaObject());
	} else if (dynamic_cast<QDistanceSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::metaObject());
	} else if (dynamic_cast<QCompass*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QCompass*>(ptr)->QCompass::metaObject());
	} else if (dynamic_cast<QAmbientTemperatureSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::metaObject());
	} else if (dynamic_cast<QAmbientLightSensor*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::metaObject());
	} else if (dynamic_cast<QAltimeter*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAltimeter*>(ptr)->QAltimeter::metaObject());
	} else if (dynamic_cast<QAccelerometer*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAccelerometer*>(ptr)->QAccelerometer::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QSensor*>(ptr)->QSensor::metaObject());
	}
}

class MyQSensorBackend: public QSensorBackend
{
public:
	void start() { callbackQSensorBackend_Start(this); };
	void stop() { callbackQSensorBackend_Stop(this); };
	bool isFeatureSupported(QSensor::Feature feature) const { return callbackQSensorBackend_IsFeatureSupported(const_cast<void*>(static_cast<const void*>(this)), feature) != 0; };
	bool event(QEvent * e) { return callbackQSensorBackend_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorBackend_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorBackend_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorBackend_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorBackend_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorBackend_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorBackend_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorBackend_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorBackend_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorBackend_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorBackend_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensorBackend*)

int QSensorBackend_QSensorBackend_QRegisterMetaType(){qRegisterMetaType<QSensorBackend*>(); return qRegisterMetaType<MyQSensorBackend*>();}

void QSensorBackend_AddDataRate(void* ptr, double min, double max)
{
	static_cast<QSensorBackend*>(ptr)->addDataRate(min, max);
}

void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy)
{
	static_cast<QSensorBackend*>(ptr)->addOutputRange(min, max, accuracy);
}

void QSensorBackend_NewReadingAvailable(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->newReadingAvailable();
}

void QSensorBackend_SensorBusy(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->sensorBusy();
}

void QSensorBackend_SensorError(void* ptr, int error)
{
	static_cast<QSensorBackend*>(ptr)->sensorError(error);
}

void QSensorBackend_SensorStopped(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->sensorStopped();
}

void QSensorBackend_SetDataRates(void* ptr, void* otherSensor)
{
	static_cast<QSensorBackend*>(ptr)->setDataRates(static_cast<QSensor*>(otherSensor));
}

void QSensorBackend_SetDescription(void* ptr, struct QtSensors_PackedString description)
{
	static_cast<QSensorBackend*>(ptr)->setDescription(QString::fromUtf8(description.data, description.len));
}

void QSensorBackend_Start(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->start();
}

void QSensorBackend_Stop(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->stop();
}

void* QSensorBackend_Sensor(void* ptr)
{
	return static_cast<QSensorBackend*>(ptr)->sensor();
}

void* QSensorBackend_Reading(void* ptr)
{
	return static_cast<QSensorBackend*>(ptr)->reading();
}

char QSensorBackend_IsFeatureSupported(void* ptr, long long feature)
{
	return static_cast<QSensorBackend*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

char QSensorBackend_IsFeatureSupportedDefault(void* ptr, long long feature)
{
		return static_cast<QSensorBackend*>(ptr)->QSensorBackend::isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void* QSensorBackend___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensorBackend___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensorBackend___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensorBackend___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorBackend___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorBackend___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorBackend___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorBackend___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorBackend___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorBackend___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorBackend___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorBackend___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorBackend___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensorBackend___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorBackend___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensorBackend_EventDefault(void* ptr, void* e)
{
		return static_cast<QSensorBackend*>(ptr)->QSensorBackend::event(static_cast<QEvent*>(e));
}

char QSensorBackend_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSensorBackend*>(ptr)->QSensorBackend::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSensorBackend_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorBackend_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorBackend_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::customEvent(static_cast<QEvent*>(event));
}

void QSensorBackend_DeleteLaterDefault(void* ptr)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::deleteLater();
}

void QSensorBackend_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorBackend_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSensorBackend*>(ptr)->QSensorBackend::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSensorBackend_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSensorBackend*>(ptr)->QSensorBackend::metaObject());
}

class MyQSensorBackendFactory: public QSensorBackendFactory
{
public:
	QSensorBackend * createBackend(QSensor * sensor) { return static_cast<QSensorBackend*>(callbackQSensorBackendFactory_CreateBackend(this, sensor)); };
};

void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor)
{
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

class MyQSensorChangesInterface: public QSensorChangesInterface
{
public:
	void sensorsChanged() { callbackQSensorChangesInterface_SensorsChanged(this); };
};

void QSensorChangesInterface_SensorsChanged(void* ptr)
{
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

class MyQSensorFilter: public QSensorFilter
{
public:
	bool filter(QSensorReading * reading) { return callbackQSensorFilter_Filter(this, reading) != 0; };
	 ~MyQSensorFilter() { callbackQSensorFilter_DestroyQSensorFilter(this); };
};

char QSensorFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

void QSensorFilter_DestroyQSensorFilter(void* ptr)
{
	static_cast<QSensorFilter*>(ptr)->~QSensorFilter();
}

void QSensorFilter_DestroyQSensorFilterDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSensorFilter_M_sensor(void* ptr)
{
	return static_cast<QSensorFilter*>(ptr)->m_sensor;
}

void QSensorFilter_SetM_sensor(void* ptr, void* vqs)
{
	static_cast<QSensorFilter*>(ptr)->m_sensor = static_cast<QSensor*>(vqs);
}

class MyQSensorGesture: public QSensorGesture
{
public:
	MyQSensorGesture(const QStringList &ids, QObject *parent = Q_NULLPTR) : QSensorGesture(ids, parent) {QSensorGesture_QSensorGesture_QRegisterMetaType();};
	#ifdef Q_QDOC
		void Signal_Detected(QString gestureId) { QByteArray t7bc790 = gestureId.toUtf8(); QtSensors_PackedString gestureIdPacked = { const_cast<char*>(t7bc790.prepend("WHITESPACE").constData()+10), t7bc790.size()-10 };callbackQSensorGesture_Detected(this, gestureIdPacked); };
	#endif
	bool event(QEvent * e) { return callbackQSensorGesture_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGesture_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorGesture_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGesture_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorGesture_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorGesture_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorGesture_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGesture_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorGesture_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorGesture_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGesture_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensorGesture*)

int QSensorGesture_QSensorGesture_QRegisterMetaType(){qRegisterMetaType<QSensorGesture*>(); return qRegisterMetaType<MyQSensorGesture*>();}

void* QSensorGesture_NewQSensorGesture(struct QtSensors_PackedString ids, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QWindow*>(parent));
	} else {
		return new MyQSensorGesture(QString::fromUtf8(ids.data, ids.len).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
	}
}

char QSensorGesture_IsActive(void* ptr)
{
	return static_cast<QSensorGesture*>(ptr)->isActive();
}

void QSensorGesture_ConnectDetected(void* ptr)
{
#ifdef Q_QDOC
	QObject::connect(static_cast<QSensorGesture*>(ptr), static_cast<void (QSensorGesture::*)(QString)>(&QSensorGesture::detected), static_cast<MyQSensorGesture*>(ptr), static_cast<void (MyQSensorGesture::*)(QString)>(&MyQSensorGesture::Signal_Detected));
#endif
}

void QSensorGesture_DisconnectDetected(void* ptr)
{
#ifdef Q_QDOC
	QObject::disconnect(static_cast<QSensorGesture*>(ptr), static_cast<void (QSensorGesture::*)(QString)>(&QSensorGesture::detected), static_cast<MyQSensorGesture*>(ptr), static_cast<void (MyQSensorGesture::*)(QString)>(&MyQSensorGesture::Signal_Detected));
#endif
}

void QSensorGesture_Detected(void* ptr, struct QtSensors_PackedString gestureId)
{
#ifdef Q_QDOC
	static_cast<QSensorGesture*>(ptr)->detected(QString::fromUtf8(gestureId.data, gestureId.len));
#endif
}

void QSensorGesture_StartDetection(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->startDetection();
}

void QSensorGesture_StopDetection(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->stopDetection();
}

void QSensorGesture_DestroyQSensorGesture(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

struct QtSensors_PackedString QSensorGesture_GestureSignals(void* ptr)
{
	return ({ QByteArray t7a3c3d = static_cast<QSensorGesture*>(ptr)->gestureSignals().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(t7a3c3d.prepend("WHITESPACE").constData()+10), t7a3c3d.size()-10 }; });
}

struct QtSensors_PackedString QSensorGesture_InvalidIds(void* ptr)
{
	return ({ QByteArray ta7952e = static_cast<QSensorGesture*>(ptr)->invalidIds().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(ta7952e.prepend("WHITESPACE").constData()+10), ta7952e.size()-10 }; });
}

struct QtSensors_PackedString QSensorGesture_ValidIds(void* ptr)
{
	return ({ QByteArray t98eddb = static_cast<QSensorGesture*>(ptr)->validIds().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(t98eddb.prepend("WHITESPACE").constData()+10), t98eddb.size()-10 }; });
}

void* QSensorGesture___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensorGesture___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensorGesture___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensorGesture___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGesture___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGesture___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGesture___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGesture___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGesture___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGesture___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGesture___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGesture___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGesture___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensorGesture___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGesture___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensorGesture_EventDefault(void* ptr, void* e)
{
		return static_cast<QSensorGesture*>(ptr)->QSensorGesture::event(static_cast<QEvent*>(e));
}

char QSensorGesture_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSensorGesture*>(ptr)->QSensorGesture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSensorGesture_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGesture_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGesture_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::customEvent(static_cast<QEvent*>(event));
}

void QSensorGesture_DeleteLaterDefault(void* ptr)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::deleteLater();
}

void QSensorGesture_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGesture_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGesture*>(ptr)->QSensorGesture::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSensorGesture_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSensorGesture*>(ptr)->QSensorGesture::metaObject());
}

class MyQSensorGestureManager: public QSensorGestureManager
{
public:
	MyQSensorGestureManager(QObject *parent = Q_NULLPTR) : QSensorGestureManager(parent) {QSensorGestureManager_QSensorGestureManager_QRegisterMetaType();};
	void Signal_NewSensorGestureAvailable() { callbackQSensorGestureManager_NewSensorGestureAvailable(this); };
	bool event(QEvent * e) { return callbackQSensorGestureManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGestureManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorGestureManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGestureManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorGestureManager_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorGestureManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorGestureManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGestureManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorGestureManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureManager_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGestureManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensorGestureManager*)

int QSensorGestureManager_QSensorGestureManager_QRegisterMetaType(){qRegisterMetaType<QSensorGestureManager*>(); return qRegisterMetaType<MyQSensorGestureManager*>();}

void* QSensorGestureManager_NewQSensorGestureManager(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureManager(static_cast<QWindow*>(parent));
	} else {
		return new MyQSensorGestureManager(static_cast<QObject*>(parent));
	}
}

void* QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(struct QtSensors_PackedString id)
{
	return QSensorGestureManager::sensorGestureRecognizer(QString::fromUtf8(id.data, id.len));
}

char QSensorGestureManager_RegisterSensorGestureRecognizer(void* ptr, void* recognizer)
{
	return static_cast<QSensorGestureManager*>(ptr)->registerSensorGestureRecognizer(static_cast<QSensorGestureRecognizer*>(recognizer));
}

void QSensorGestureManager_ConnectNewSensorGestureAvailable(void* ptr)
{
	QObject::connect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));
}

void QSensorGestureManager_DisconnectNewSensorGestureAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));
}

void QSensorGestureManager_NewSensorGestureAvailable(void* ptr)
{
	static_cast<QSensorGestureManager*>(ptr)->newSensorGestureAvailable();
}

void QSensorGestureManager_DestroyQSensorGestureManager(void* ptr)
{
	static_cast<QSensorGestureManager*>(ptr)->~QSensorGestureManager();
}

struct QtSensors_PackedString QSensorGestureManager_GestureIds(void* ptr)
{
	return ({ QByteArray t5f71c3 = static_cast<QSensorGestureManager*>(ptr)->gestureIds().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(t5f71c3.prepend("WHITESPACE").constData()+10), t5f71c3.size()-10 }; });
}

struct QtSensors_PackedString QSensorGestureManager_RecognizerSignals(void* ptr, struct QtSensors_PackedString gestureId)
{
	return ({ QByteArray t0f00fe = static_cast<QSensorGestureManager*>(ptr)->recognizerSignals(QString::fromUtf8(gestureId.data, gestureId.len)).join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(t0f00fe.prepend("WHITESPACE").constData()+10), t0f00fe.size()-10 }; });
}

void* QSensorGestureManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensorGestureManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensorGestureManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensorGestureManager___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureManager___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureManager___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureManager___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensorGestureManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensorGestureManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::event(static_cast<QEvent*>(e));
}

char QSensorGestureManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSensorGestureManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::deleteLater();
}

void QSensorGestureManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSensorGestureManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::metaObject());
}

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface
{
public:
	MyQSensorGesturePluginInterface() : QSensorGesturePluginInterface() {};
	QList<QSensorGestureRecognizer *> createRecognizers() { return *static_cast<QList<QSensorGestureRecognizer *>*>(callbackQSensorGesturePluginInterface_CreateRecognizers(this)); };
	 ~MyQSensorGesturePluginInterface() { callbackQSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(this); };
	QString name() const { return ({ QtSensors_PackedString tempVal = callbackQSensorGesturePluginInterface_Name(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	QStringList supportedIds() const { return ({ QtSensors_PackedString tempVal = callbackQSensorGesturePluginInterface_SupportedIds(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("|", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
};

struct QtSensors_PackedList QSensorGesturePluginInterface_CreateRecognizers(void* ptr)
{
	return ({ QList<QSensorGestureRecognizer *>* tmpValue = new QList<QSensorGestureRecognizer *>(static_cast<QSensorGesturePluginInterface*>(ptr)->createRecognizers()); QtSensors_PackedList { tmpValue, tmpValue->size() }; });
}

void* QSensorGesturePluginInterface_NewQSensorGesturePluginInterface()
{
	return new MyQSensorGesturePluginInterface();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr)
{
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtSensors_PackedString QSensorGesturePluginInterface_Name(void* ptr)
{
	return ({ QByteArray t5bfb97 = static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8(); QtSensors_PackedString { const_cast<char*>(t5bfb97.prepend("WHITESPACE").constData()+10), t5bfb97.size()-10 }; });
}

struct QtSensors_PackedString QSensorGesturePluginInterface_SupportedIds(void* ptr)
{
	return ({ QByteArray tab1a26 = static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(tab1a26.prepend("WHITESPACE").constData()+10), tab1a26.size()-10 }; });
}

void* QSensorGesturePluginInterface___createRecognizers_atList(void* ptr, int i)
{
	return const_cast<QSensorGestureRecognizer*>(static_cast<QList<QSensorGestureRecognizer *>*>(ptr)->at(i));
}

void QSensorGesturePluginInterface___createRecognizers_setList(void* ptr, void* i)
{
	static_cast<QList<QSensorGestureRecognizer *>*>(ptr)->append(static_cast<QSensorGestureRecognizer*>(i));
}

void* QSensorGesturePluginInterface___createRecognizers_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSensorGestureRecognizer *>;
}

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer
{
public:
	MyQSensorGestureRecognizer(QObject *parent = Q_NULLPTR) : QSensorGestureRecognizer(parent) {QSensorGestureRecognizer_QSensorGestureRecognizer_QRegisterMetaType();};
	bool isActive() { return callbackQSensorGestureRecognizer_IsActive(this) != 0; };
	bool start() { return callbackQSensorGestureRecognizer_Start(this) != 0; };
	bool stop() { return callbackQSensorGestureRecognizer_Stop(this) != 0; };
	void create() { callbackQSensorGestureRecognizer_Create(this); };
	void Signal_Detected(const QString & gestureId) { QByteArray t7bc790 = gestureId.toUtf8(); QtSensors_PackedString gestureIdPacked = { const_cast<char*>(t7bc790.prepend("WHITESPACE").constData()+10), t7bc790.size()-10 };callbackQSensorGestureRecognizer_Detected(this, gestureIdPacked); };
	 ~MyQSensorGestureRecognizer() { callbackQSensorGestureRecognizer_DestroyQSensorGestureRecognizer(this); };
	QString id() const { return ({ QtSensors_PackedString tempVal = callbackQSensorGestureRecognizer_Id(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool event(QEvent * e) { return callbackQSensorGestureRecognizer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGestureRecognizer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorGestureRecognizer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGestureRecognizer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorGestureRecognizer_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorGestureRecognizer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorGestureRecognizer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGestureRecognizer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorGestureRecognizer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureRecognizer_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGestureRecognizer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensorGestureRecognizer*)

int QSensorGestureRecognizer_QSensorGestureRecognizer_QRegisterMetaType(){qRegisterMetaType<QSensorGestureRecognizer*>(); return qRegisterMetaType<MyQSensorGestureRecognizer*>();}

void* QSensorGestureRecognizer_NewQSensorGestureRecognizer(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSensorGestureRecognizer(static_cast<QWindow*>(parent));
	} else {
		return new MyQSensorGestureRecognizer(static_cast<QObject*>(parent));
	}
}

char QSensorGestureRecognizer_IsActive(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->isActive();
}

char QSensorGestureRecognizer_Start(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->start();
}

char QSensorGestureRecognizer_Stop(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->stop();
}

void QSensorGestureRecognizer_Create(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->create();
}

void QSensorGestureRecognizer_CreateBackend(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->createBackend();
}

void QSensorGestureRecognizer_ConnectDetected(void* ptr)
{
	QObject::connect(static_cast<QSensorGestureRecognizer*>(ptr), static_cast<void (QSensorGestureRecognizer::*)(const QString &)>(&QSensorGestureRecognizer::detected), static_cast<MyQSensorGestureRecognizer*>(ptr), static_cast<void (MyQSensorGestureRecognizer::*)(const QString &)>(&MyQSensorGestureRecognizer::Signal_Detected));
}

void QSensorGestureRecognizer_DisconnectDetected(void* ptr)
{
	QObject::disconnect(static_cast<QSensorGestureRecognizer*>(ptr), static_cast<void (QSensorGestureRecognizer::*)(const QString &)>(&QSensorGestureRecognizer::detected), static_cast<MyQSensorGestureRecognizer*>(ptr), static_cast<void (MyQSensorGestureRecognizer::*)(const QString &)>(&MyQSensorGestureRecognizer::Signal_Detected));
}

void QSensorGestureRecognizer_Detected(void* ptr, struct QtSensors_PackedString gestureId)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->detected(QString::fromUtf8(gestureId.data, gestureId.len));
}

void QSensorGestureRecognizer_StartBackend(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->startBackend();
}

void QSensorGestureRecognizer_StopBackend(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->stopBackend();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->~QSensorGestureRecognizer();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

struct QtSensors_PackedString QSensorGestureRecognizer_Id(void* ptr)
{
	return ({ QByteArray t1336bf = static_cast<QSensorGestureRecognizer*>(ptr)->id().toUtf8(); QtSensors_PackedString { const_cast<char*>(t1336bf.prepend("WHITESPACE").constData()+10), t1336bf.size()-10 }; });
}

struct QtSensors_PackedString QSensorGestureRecognizer_GestureSignals(void* ptr)
{
	return ({ QByteArray t79f8ee = static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join("|").toUtf8(); QtSensors_PackedString { const_cast<char*>(t79f8ee.prepend("WHITESPACE").constData()+10), t79f8ee.size()-10 }; });
}

void* QSensorGestureRecognizer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensorGestureRecognizer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensorGestureRecognizer___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensorGestureRecognizer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureRecognizer___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureRecognizer___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureRecognizer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureRecognizer___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureRecognizer___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureRecognizer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorGestureRecognizer___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureRecognizer___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorGestureRecognizer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensorGestureRecognizer___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorGestureRecognizer___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensorGestureRecognizer_EventDefault(void* ptr, void* e)
{
		return static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::event(static_cast<QEvent*>(e));
}

char QSensorGestureRecognizer_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QSensorGestureRecognizer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureRecognizer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureRecognizer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureRecognizer_DeleteLaterDefault(void* ptr)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::deleteLater();
}

void QSensorGestureRecognizer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureRecognizer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QSensorGestureRecognizer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::metaObject());
}

void* QSensorManager_QSensorManager_CreateBackend(void* sensor)
{
	return QSensorManager::createBackend(static_cast<QSensor*>(sensor));
}

char QSensorManager_QSensorManager_IsBackendRegistered(void* ty, void* identifier)
{
	return QSensorManager::isBackendRegistered(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_RegisterBackend(void* ty, void* identifier, void* factory)
{
	QSensorManager::registerBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier), static_cast<QSensorBackendFactory*>(factory));
}

void QSensorManager_QSensorManager_SetDefaultBackend(void* ty, void* identifier)
{
	QSensorManager::setDefaultBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_UnregisterBackend(void* ty, void* identifier)
{
	QSensorManager::unregisterBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

class MyQSensorPluginInterface: public QSensorPluginInterface
{
public:
	void registerSensors() { callbackQSensorPluginInterface_RegisterSensors(this); };
};

void QSensorPluginInterface_RegisterSensors(void* ptr)
{
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

class MyQSensorReading: public QSensorReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQSensorReading*)

int QSensorReading_QSensorReading_QRegisterMetaType(){qRegisterMetaType<QSensorReading*>(); return qRegisterMetaType<MyQSensorReading*>();}

void QSensorReading_SetTimestamp(void* ptr, unsigned long long timestamp)
{
	static_cast<QSensorReading*>(ptr)->setTimestamp(timestamp);
}

void* QSensorReading_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSensorReading*>(ptr)->value(index));
}

int QSensorReading_ValueCount(void* ptr)
{
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

unsigned long long QSensorReading_Timestamp(void* ptr)
{
	return static_cast<QSensorReading*>(ptr)->timestamp();
}

void* QSensorReading___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QSensorReading___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QSensorReading___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QSensorReading___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorReading___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorReading___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorReading___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorReading___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorReading___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorReading___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QSensorReading___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorReading___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QSensorReading___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QSensorReading___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QSensorReading___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QSensorReading_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTiltReading*>(ptr)->QTiltReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTapReading*>(ptr)->QTapReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRotationReading*>(ptr)->QRotationReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QProximityReading*>(ptr)->QProximityReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPressureReading*>(ptr)->QPressureReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QOrientationReading*>(ptr)->QOrientationReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLightReading*>(ptr)->QLightReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLidReading*>(ptr)->QLidReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHumidityReading*>(ptr)->QHumidityReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHolsterReading*>(ptr)->QHolsterReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDistanceReading*>(ptr)->QDistanceReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCompassReading*>(ptr)->QCompassReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QSensorReading*>(ptr)->QSensorReading::event(static_cast<QEvent*>(e));
	}
}

char QSensorReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTiltReading*>(ptr)->QTiltReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTapReading*>(ptr)->QTapReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRotationReading*>(ptr)->QRotationReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QProximityReading*>(ptr)->QProximityReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QPressureReading*>(ptr)->QPressureReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QOrientationReading*>(ptr)->QOrientationReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLightReading*>(ptr)->QLightReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QLidReading*>(ptr)->QLidReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHumidityReading*>(ptr)->QHumidityReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QHolsterReading*>(ptr)->QHolsterReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDistanceReading*>(ptr)->QDistanceReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCompassReading*>(ptr)->QCompassReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QSensorReading*>(ptr)->QSensorReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QSensorReading_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QSensorReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSensorReading_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::customEvent(static_cast<QEvent*>(event));
	}
}

void QSensorReading_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::deleteLater();
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::deleteLater();
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::deleteLater();
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::deleteLater();
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::deleteLater();
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::deleteLater();
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::deleteLater();
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::deleteLater();
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::deleteLater();
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::deleteLater();
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::deleteLater();
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::deleteLater();
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::deleteLater();
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::deleteLater();
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::deleteLater();
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::deleteLater();
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::deleteLater();
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::deleteLater();
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::deleteLater();
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::deleteLater();
	}
}

void QSensorReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QSensorReading_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTiltReading*>(ptr)->QTiltReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QTapReading*>(ptr)->QTapReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QRotationReading*>(ptr)->QRotationReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QProximityReading*>(ptr)->QProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QPressureReading*>(ptr)->QPressureReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QOrientationReading*>(ptr)->QOrientationReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLightReading*>(ptr)->QLightReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QLidReading*>(ptr)->QLidReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHumidityReading*>(ptr)->QHumidityReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QHolsterReading*>(ptr)->QHolsterReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QDistanceReading*>(ptr)->QDistanceReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QCompassReading*>(ptr)->QCompassReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QSensorReading*>(ptr)->QSensorReading::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QSensorReading_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QTiltReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QTiltReading*>(ptr)->QTiltReading::metaObject());
	} else if (dynamic_cast<QTapReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QTapReading*>(ptr)->QTapReading::metaObject());
	} else if (dynamic_cast<QRotationReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRotationReading*>(ptr)->QRotationReading::metaObject());
	} else if (dynamic_cast<QProximityReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QProximityReading*>(ptr)->QProximityReading::metaObject());
	} else if (dynamic_cast<QPressureReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QPressureReading*>(ptr)->QPressureReading::metaObject());
	} else if (dynamic_cast<QOrientationReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QOrientationReading*>(ptr)->QOrientationReading::metaObject());
	} else if (dynamic_cast<QMagnetometerReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::metaObject());
	} else if (dynamic_cast<QLightReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLightReading*>(ptr)->QLightReading::metaObject());
	} else if (dynamic_cast<QLidReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QLidReading*>(ptr)->QLidReading::metaObject());
	} else if (dynamic_cast<QIRProximityReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::metaObject());
	} else if (dynamic_cast<QHumidityReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHumidityReading*>(ptr)->QHumidityReading::metaObject());
	} else if (dynamic_cast<QHolsterReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QHolsterReading*>(ptr)->QHolsterReading::metaObject());
	} else if (dynamic_cast<QGyroscopeReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::metaObject());
	} else if (dynamic_cast<QDistanceReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QDistanceReading*>(ptr)->QDistanceReading::metaObject());
	} else if (dynamic_cast<QCompassReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QCompassReading*>(ptr)->QCompassReading::metaObject());
	} else if (dynamic_cast<QAmbientTemperatureReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::metaObject());
	} else if (dynamic_cast<QAmbientLightReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::metaObject());
	} else if (dynamic_cast<QAltimeterReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::metaObject());
	} else if (dynamic_cast<QAccelerometerReading*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QSensorReading*>(ptr)->QSensorReading::metaObject());
	}
}

class MyQTapFilter: public QTapFilter
{
public:
	bool filter(QTapReading * reading) { return callbackQTapFilter_Filter(this, reading) != 0; };
};

char QTapFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

class MyQTapReading: public QTapReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQTapReading*)

int QTapReading_QTapReading_QRegisterMetaType(){qRegisterMetaType<QTapReading*>(); return qRegisterMetaType<MyQTapReading*>();}

void QTapReading_SetDoubleTap(void* ptr, char doubleTap)
{
	static_cast<QTapReading*>(ptr)->setDoubleTap(doubleTap != 0);
}

void QTapReading_SetTapDirection(void* ptr, long long tapDirection)
{
	static_cast<QTapReading*>(ptr)->setTapDirection(static_cast<QTapReading::TapDirection>(tapDirection));
}

long long QTapReading_TapDirection(void* ptr)
{
	return static_cast<QTapReading*>(ptr)->tapDirection();
}

char QTapReading_IsDoubleTap(void* ptr)
{
	return static_cast<QTapReading*>(ptr)->isDoubleTap();
}

class MyQTapSensor: public QTapSensor
{
public:
	MyQTapSensor(QObject *parent = Q_NULLPTR) : QTapSensor(parent) {QTapSensor_QTapSensor_QRegisterMetaType();};
	void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents) { callbackQTapSensor_ReturnDoubleTapEventsChanged(this, returnDoubleTapEvents); };
	 ~MyQTapSensor() { callbackQTapSensor_DestroyQTapSensor(this); };
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQTapSensor*)

int QTapSensor_QTapSensor_QRegisterMetaType(){qRegisterMetaType<QTapSensor*>(); return qRegisterMetaType<MyQTapSensor*>();}

void* QTapSensor_NewQTapSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTapSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQTapSensor(static_cast<QObject*>(parent));
	}
}

void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr)
{
	QObject::connect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));
}

void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));
}

void QTapSensor_ReturnDoubleTapEventsChanged(void* ptr, char returnDoubleTapEvents)
{
	static_cast<QTapSensor*>(ptr)->returnDoubleTapEventsChanged(returnDoubleTapEvents != 0);
}

void QTapSensor_SetReturnDoubleTapEvents(void* ptr, char returnDoubleTapEvents)
{
	static_cast<QTapSensor*>(ptr)->setReturnDoubleTapEvents(returnDoubleTapEvents != 0);
}

void QTapSensor_DestroyQTapSensor(void* ptr)
{
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
}

void QTapSensor_DestroyQTapSensorDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTapSensor_Reading(void* ptr)
{
	return static_cast<QTapSensor*>(ptr)->reading();
}

char QTapSensor_ReturnDoubleTapEvents(void* ptr)
{
	return static_cast<QTapSensor*>(ptr)->returnDoubleTapEvents();
}

struct QtSensors_PackedString QTapSensor_QTapSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QTapSensor::type), -1 };
}

class MyQTiltFilter: public QTiltFilter
{
public:
	bool filter(QTiltReading * reading) { return callbackQTiltFilter_Filter(this, reading) != 0; };
};

char QTiltFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

class MyQTiltReading: public QTiltReading
{
public:
	bool event(QEvent * e) { return callbackQSensorReading_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorReading_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensorReading_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorReading_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensorReading_CustomEvent(this, event); };
	void deleteLater() { callbackQSensorReading_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensorReading_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorReading_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensorReading_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensorReading_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorReading_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQTiltReading*)

int QTiltReading_QTiltReading_QRegisterMetaType(){qRegisterMetaType<QTiltReading*>(); return qRegisterMetaType<MyQTiltReading*>();}

void QTiltReading_SetXRotation(void* ptr, double x)
{
	static_cast<QTiltReading*>(ptr)->setXRotation(x);
}

void QTiltReading_SetYRotation(void* ptr, double y)
{
	static_cast<QTiltReading*>(ptr)->setYRotation(y);
}

double QTiltReading_XRotation(void* ptr)
{
	return static_cast<QTiltReading*>(ptr)->xRotation();
}

double QTiltReading_YRotation(void* ptr)
{
	return static_cast<QTiltReading*>(ptr)->yRotation();
}

class MyQTiltSensor: public QTiltSensor
{
public:
	MyQTiltSensor(QObject *parent = Q_NULLPTR) : QTiltSensor(parent) {QTiltSensor_QTiltSensor_QRegisterMetaType();};
	bool start() { return callbackQSensor_Start(this) != 0; };
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, skipDuplicates); };
	void stop() { callbackQSensor_Stop(this); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, userOrientation); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, event); };
	void deleteLater() { callbackQSensor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQSensor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtSensors_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQSensor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQTiltSensor*)

int QTiltSensor_QTiltSensor_QRegisterMetaType(){qRegisterMetaType<QTiltSensor*>(); return qRegisterMetaType<MyQTiltSensor*>();}

void* QTiltSensor_NewQTiltSensor(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTiltSensor(static_cast<QWindow*>(parent));
	} else {
		return new MyQTiltSensor(static_cast<QObject*>(parent));
	}
}

void QTiltSensor_Calibrate(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->calibrate();
}

void QTiltSensor_DestroyQTiltSensor(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->~QTiltSensor();
}

void* QTiltSensor_Reading(void* ptr)
{
	return static_cast<QTiltSensor*>(ptr)->reading();
}

struct QtSensors_PackedString QTiltSensor_QTiltSensor_Type()
{
	return QtSensors_PackedString { const_cast<char*>(QTiltSensor::type), -1 };
}

