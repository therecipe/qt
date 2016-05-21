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
#include <QChildEvent>
#include <QCompass>
#include <QCompassFilter>
#include <QCompassReading>
#include <QDistanceFilter>
#include <QDistanceReading>
#include <QDistanceSensor>
#include <QEvent>
#include <QGyroscope>
#include <QGyroscopeFilter>
#include <QGyroscopeReading>
#include <QHolsterFilter>
#include <QHolsterReading>
#include <QHolsterSensor>
#include <QIRProximityFilter>
#include <QIRProximityReading>
#include <QIRProximitySensor>
#include <QLightFilter>
#include <QLightReading>
#include <QLightSensor>
#include <QMagnetometer>
#include <QMagnetometerFilter>
#include <QMagnetometerReading>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOrientationFilter>
#include <QOrientationReading>
#include <QOrientationSensor>
#include <QPressureFilter>
#include <QPressureReading>
#include <QPressureSensor>
#include <QProximityFilter>
#include <QProximityReading>
#include <QProximitySensor>
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

class MyQAccelerometer: public QAccelerometer
{
public:
	MyQAccelerometer(QObject *parent) : QAccelerometer(parent) {};
	void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode) { callbackQAccelerometer_AccelerationModeChanged(this, this->objectName().toUtf8().data(), accelerationMode); };
	bool start() { return callbackQAccelerometer_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQAccelerometer_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAccelerometer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAccelerometer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAccelerometer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAccelerometer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAccelerometer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAccelerometer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAccelerometer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAccelerometer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAccelerometer_MetaObject(const_cast<MyQAccelerometer*>(this), this->objectName().toUtf8().data())); };
};

int QAccelerometer_AccelerationMode(void* ptr)
{
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

void* QAccelerometer_Reading(void* ptr)
{
	return static_cast<QAccelerometer*>(ptr)->reading();
}

void* QAccelerometer_NewQAccelerometer(void* parent)
{
	return new MyQAccelerometer(static_cast<QObject*>(parent));
}

void QAccelerometer_ConnectAccelerationModeChanged(void* ptr)
{
	QObject::connect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));
}

void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));
}

void QAccelerometer_AccelerationModeChanged(void* ptr, int accelerationMode)
{
	static_cast<QAccelerometer*>(ptr)->accelerationModeChanged(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode)
{
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_DestroyQAccelerometer(void* ptr)
{
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
}

char* QAccelerometer_QAccelerometer_Type()
{
	return QString(QAccelerometer::type).toUtf8().data();
}

int QAccelerometer_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAccelerometer*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QAccelerometer_StartDefault(void* ptr)
{
	return static_cast<QAccelerometer*>(ptr)->QAccelerometer::start();
}

void QAccelerometer_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAccelerometer*>(ptr), "stop");
}

void QAccelerometer_StopDefault(void* ptr)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::stop();
}

void QAccelerometer_TimerEvent(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometer_ChildEvent(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAccelerometer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometer_CustomEvent(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAccelerometer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::customEvent(static_cast<QEvent*>(event));
}

void QAccelerometer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAccelerometer*>(ptr), "deleteLater");
}

void QAccelerometer_DeleteLaterDefault(void* ptr)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::deleteLater();
}

void QAccelerometer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAccelerometer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAccelerometer_Event(void* ptr, void* e)
{
	return static_cast<QAccelerometer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAccelerometer_EventDefault(void* ptr, void* e)
{
	return static_cast<QAccelerometer*>(ptr)->QAccelerometer::event(static_cast<QEvent*>(e));
}

int QAccelerometer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAccelerometer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAccelerometer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAccelerometer*>(ptr)->QAccelerometer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAccelerometer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccelerometer*>(ptr)->metaObject());
}

void* QAccelerometer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccelerometer*>(ptr)->QAccelerometer::metaObject());
}

class MyQAccelerometerFilter: public QAccelerometerFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QAccelerometerReading * reading) { return callbackQAccelerometerFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QAccelerometerFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

char* QAccelerometerFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccelerometerFilter*>(static_cast<QAccelerometerFilter*>(ptr))) {
		return static_cast<MyQAccelerometerFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccelerometerFilter_BASE").toUtf8().data();
}

void QAccelerometerFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccelerometerFilter*>(static_cast<QAccelerometerFilter*>(ptr))) {
		static_cast<MyQAccelerometerFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QAccelerometerReading_X(void* ptr)
{
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->x());
}

double QAccelerometerReading_Y(void* ptr)
{
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->y());
}

double QAccelerometerReading_Z(void* ptr)
{
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->z());
}

void QAccelerometerReading_SetX(void* ptr, double x)
{
	static_cast<QAccelerometerReading*>(ptr)->setX(static_cast<double>(x));
}

void QAccelerometerReading_SetY(void* ptr, double y)
{
	static_cast<QAccelerometerReading*>(ptr)->setY(static_cast<double>(y));
}

void QAccelerometerReading_SetZ(void* ptr, double z)
{
	static_cast<QAccelerometerReading*>(ptr)->setZ(static_cast<double>(z));
}

void QAccelerometerReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometerReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometerReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometerReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometerReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAccelerometerReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometerReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometerReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAccelerometerReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::customEvent(static_cast<QEvent*>(event));
}

void QAccelerometerReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAccelerometerReading*>(ptr), "deleteLater");
}

void QAccelerometerReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::deleteLater();
}

void QAccelerometerReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAccelerometerReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccelerometerReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAccelerometerReading_Event(void* ptr, void* e)
{
	return static_cast<QAccelerometerReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAccelerometerReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::event(static_cast<QEvent*>(e));
}

int QAccelerometerReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAccelerometerReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAccelerometerReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAccelerometerReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccelerometerReading*>(ptr)->metaObject());
}

void* QAccelerometerReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::metaObject());
}

void* QAltimeter_Reading(void* ptr)
{
	return static_cast<QAltimeter*>(ptr)->reading();
}

void* QAltimeter_NewQAltimeter(void* parent)
{
	return new QAltimeter(static_cast<QObject*>(parent));
}

void QAltimeter_DestroyQAltimeter(void* ptr)
{
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

char* QAltimeter_QAltimeter_Type()
{
	return QString(QAltimeter::type).toUtf8().data();
}

int QAltimeter_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAltimeter*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QAltimeter_StartDefault(void* ptr)
{
	return static_cast<QAltimeter*>(ptr)->QAltimeter::start();
}

void QAltimeter_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAltimeter*>(ptr), "stop");
}

void QAltimeter_StopDefault(void* ptr)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::stop();
}

void QAltimeter_TimerEvent(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeter_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeter_ChildEvent(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeter_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeter_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAltimeter*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeter_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeter_CustomEvent(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAltimeter_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::customEvent(static_cast<QEvent*>(event));
}

void QAltimeter_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAltimeter*>(ptr), "deleteLater");
}

void QAltimeter_DeleteLaterDefault(void* ptr)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::deleteLater();
}

void QAltimeter_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAltimeter*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeter_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAltimeter*>(ptr)->QAltimeter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAltimeter_Event(void* ptr, void* e)
{
	return static_cast<QAltimeter*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAltimeter_EventDefault(void* ptr, void* e)
{
	return static_cast<QAltimeter*>(ptr)->QAltimeter::event(static_cast<QEvent*>(e));
}

int QAltimeter_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAltimeter*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAltimeter_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAltimeter*>(ptr)->QAltimeter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAltimeter_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAltimeter*>(ptr)->metaObject());
}

void* QAltimeter_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAltimeter*>(ptr)->QAltimeter::metaObject());
}

class MyQAltimeterFilter: public QAltimeterFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QAltimeterReading * reading) { return callbackQAltimeterFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QAltimeterFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

char* QAltimeterFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAltimeterFilter*>(static_cast<QAltimeterFilter*>(ptr))) {
		return static_cast<MyQAltimeterFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAltimeterFilter_BASE").toUtf8().data();
}

void QAltimeterFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAltimeterFilter*>(static_cast<QAltimeterFilter*>(ptr))) {
		static_cast<MyQAltimeterFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QAltimeterReading_Altitude(void* ptr)
{
	return static_cast<double>(static_cast<QAltimeterReading*>(ptr)->altitude());
}

void QAltimeterReading_SetAltitude(void* ptr, double altitude)
{
	static_cast<QAltimeterReading*>(ptr)->setAltitude(static_cast<double>(altitude));
}

void QAltimeterReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeterReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeterReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeterReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeterReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAltimeterReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeterReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeterReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAltimeterReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::customEvent(static_cast<QEvent*>(event));
}

void QAltimeterReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAltimeterReading*>(ptr), "deleteLater");
}

void QAltimeterReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::deleteLater();
}

void QAltimeterReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAltimeterReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAltimeterReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAltimeterReading_Event(void* ptr, void* e)
{
	return static_cast<QAltimeterReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAltimeterReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::event(static_cast<QEvent*>(e));
}

int QAltimeterReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAltimeterReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAltimeterReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAltimeterReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAltimeterReading*>(ptr)->metaObject());
}

void* QAltimeterReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::metaObject());
}

class MyQAmbientLightFilter: public QAmbientLightFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QAmbientLightReading * reading) { return callbackQAmbientLightFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QAmbientLightFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

char* QAmbientLightFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAmbientLightFilter*>(static_cast<QAmbientLightFilter*>(ptr))) {
		return static_cast<MyQAmbientLightFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAmbientLightFilter_BASE").toUtf8().data();
}

void QAmbientLightFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAmbientLightFilter*>(static_cast<QAmbientLightFilter*>(ptr))) {
		static_cast<MyQAmbientLightFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QAmbientLightReading_LightLevel(void* ptr)
{
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

void QAmbientLightReading_SetLightLevel(void* ptr, int lightLevel)
{
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

void QAmbientLightReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientLightReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientLightReading*>(ptr), "deleteLater");
}

void QAmbientLightReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::deleteLater();
}

void QAmbientLightReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientLightReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAmbientLightReading_Event(void* ptr, void* e)
{
	return static_cast<QAmbientLightReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAmbientLightReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::event(static_cast<QEvent*>(e));
}

int QAmbientLightReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientLightReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAmbientLightReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAmbientLightReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientLightReading*>(ptr)->metaObject());
}

void* QAmbientLightReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::metaObject());
}

class MyQAmbientLightSensor: public QAmbientLightSensor
{
public:
	MyQAmbientLightSensor(QObject *parent) : QAmbientLightSensor(parent) {};
	bool start() { return callbackQAmbientLightSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQAmbientLightSensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQAmbientLightSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAmbientLightSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAmbientLightSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAmbientLightSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAmbientLightSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAmbientLightSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAmbientLightSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAmbientLightSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAmbientLightSensor_MetaObject(const_cast<MyQAmbientLightSensor*>(this), this->objectName().toUtf8().data())); };
};

void* QAmbientLightSensor_Reading(void* ptr)
{
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent)
{
	return new MyQAmbientLightSensor(static_cast<QObject*>(parent));
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr)
{
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

char* QAmbientLightSensor_QAmbientLightSensor_Type()
{
	return QString(QAmbientLightSensor::type).toUtf8().data();
}

int QAmbientLightSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAmbientLightSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QAmbientLightSensor_StartDefault(void* ptr)
{
	return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::start();
}

void QAmbientLightSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientLightSensor*>(ptr), "stop");
}

void QAmbientLightSensor_StopDefault(void* ptr)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::stop();
}

void QAmbientLightSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientLightSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientLightSensor*>(ptr), "deleteLater");
}

void QAmbientLightSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::deleteLater();
}

void QAmbientLightSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientLightSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientLightSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAmbientLightSensor_Event(void* ptr, void* e)
{
	return static_cast<QAmbientLightSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAmbientLightSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::event(static_cast<QEvent*>(e));
}

int QAmbientLightSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientLightSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAmbientLightSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAmbientLightSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientLightSensor*>(ptr)->metaObject());
}

void* QAmbientLightSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::metaObject());
}

class MyQAmbientTemperatureFilter: public QAmbientTemperatureFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QAmbientTemperatureReading * reading) { return callbackQAmbientTemperatureFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QAmbientTemperatureFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

char* QAmbientTemperatureFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAmbientTemperatureFilter*>(static_cast<QAmbientTemperatureFilter*>(ptr))) {
		return static_cast<MyQAmbientTemperatureFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAmbientTemperatureFilter_BASE").toUtf8().data();
}

void QAmbientTemperatureFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAmbientTemperatureFilter*>(static_cast<QAmbientTemperatureFilter*>(ptr))) {
		static_cast<MyQAmbientTemperatureFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QAmbientTemperatureReading_Temperature(void* ptr)
{
	return static_cast<double>(static_cast<QAmbientTemperatureReading*>(ptr)->temperature());
}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(static_cast<double>(temperature));
}

void QAmbientTemperatureReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientTemperatureReading*>(ptr), "deleteLater");
}

void QAmbientTemperatureReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::deleteLater();
}

void QAmbientTemperatureReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAmbientTemperatureReading_Event(void* ptr, void* e)
{
	return static_cast<QAmbientTemperatureReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAmbientTemperatureReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::event(static_cast<QEvent*>(e));
}

int QAmbientTemperatureReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientTemperatureReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAmbientTemperatureReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAmbientTemperatureReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureReading*>(ptr)->metaObject());
}

void* QAmbientTemperatureReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::metaObject());
}

void* QAmbientTemperatureSensor_Reading(void* ptr)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent)
{
	return new QAmbientTemperatureSensor(static_cast<QObject*>(parent));
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

char* QAmbientTemperatureSensor_QAmbientTemperatureSensor_Type()
{
	return QString(QAmbientTemperatureSensor::type).toUtf8().data();
}

int QAmbientTemperatureSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAmbientTemperatureSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QAmbientTemperatureSensor_StartDefault(void* ptr)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::start();
}

void QAmbientTemperatureSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientTemperatureSensor*>(ptr), "stop");
}

void QAmbientTemperatureSensor_StopDefault(void* ptr)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::stop();
}

void QAmbientTemperatureSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAmbientTemperatureSensor*>(ptr), "deleteLater");
}

void QAmbientTemperatureSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::deleteLater();
}

void QAmbientTemperatureSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAmbientTemperatureSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAmbientTemperatureSensor_Event(void* ptr, void* e)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAmbientTemperatureSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::event(static_cast<QEvent*>(e));
}

int QAmbientTemperatureSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAmbientTemperatureSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAmbientTemperatureSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureSensor*>(ptr)->metaObject());
}

void* QAmbientTemperatureSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::metaObject());
}

class MyQCompass: public QCompass
{
public:
	MyQCompass(QObject *parent) : QCompass(parent) {};
	bool start() { return callbackQCompass_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQCompass_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQCompass_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCompass_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQCompass_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQCompass_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQCompass_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQCompass_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQCompass_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQCompass_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQCompass_MetaObject(const_cast<MyQCompass*>(this), this->objectName().toUtf8().data())); };
};

void* QCompass_Reading(void* ptr)
{
	return static_cast<QCompass*>(ptr)->reading();
}

void* QCompass_NewQCompass(void* parent)
{
	return new MyQCompass(static_cast<QObject*>(parent));
}

void QCompass_DestroyQCompass(void* ptr)
{
	static_cast<QCompass*>(ptr)->~QCompass();
}

char* QCompass_QCompass_Type()
{
	return QString(QCompass::type).toUtf8().data();
}

int QCompass_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QCompass*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QCompass_StartDefault(void* ptr)
{
	return static_cast<QCompass*>(ptr)->QCompass::start();
}

void QCompass_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCompass*>(ptr), "stop");
}

void QCompass_StopDefault(void* ptr)
{
	static_cast<QCompass*>(ptr)->QCompass::stop();
}

void QCompass_TimerEvent(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompass_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->QCompass::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompass_ChildEvent(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCompass_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->QCompass::childEvent(static_cast<QChildEvent*>(event));
}

void QCompass_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCompass*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompass_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCompass*>(ptr)->QCompass::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompass_CustomEvent(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCompass_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCompass*>(ptr)->QCompass::customEvent(static_cast<QEvent*>(event));
}

void QCompass_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCompass*>(ptr), "deleteLater");
}

void QCompass_DeleteLaterDefault(void* ptr)
{
	static_cast<QCompass*>(ptr)->QCompass::deleteLater();
}

void QCompass_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCompass*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompass_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCompass*>(ptr)->QCompass::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QCompass_Event(void* ptr, void* e)
{
	return static_cast<QCompass*>(ptr)->event(static_cast<QEvent*>(e));
}

int QCompass_EventDefault(void* ptr, void* e)
{
	return static_cast<QCompass*>(ptr)->QCompass::event(static_cast<QEvent*>(e));
}

int QCompass_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCompass*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QCompass_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCompass*>(ptr)->QCompass::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCompass_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCompass*>(ptr)->metaObject());
}

void* QCompass_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCompass*>(ptr)->QCompass::metaObject());
}

class MyQCompassFilter: public QCompassFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QCompassReading * reading) { return callbackQCompassFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QCompassFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

char* QCompassFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQCompassFilter*>(static_cast<QCompassFilter*>(ptr))) {
		return static_cast<MyQCompassFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QCompassFilter_BASE").toUtf8().data();
}

void QCompassFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQCompassFilter*>(static_cast<QCompassFilter*>(ptr))) {
		static_cast<MyQCompassFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QCompassReading_Azimuth(void* ptr)
{
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->azimuth());
}

double QCompassReading_CalibrationLevel(void* ptr)
{
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->calibrationLevel());
}

void QCompassReading_SetAzimuth(void* ptr, double azimuth)
{
	static_cast<QCompassReading*>(ptr)->setAzimuth(static_cast<double>(azimuth));
}

void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel)
{
	static_cast<QCompassReading*>(ptr)->setCalibrationLevel(static_cast<double>(calibrationLevel));
}

void QCompassReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompassReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompassReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCompassReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::childEvent(static_cast<QChildEvent*>(event));
}

void QCompassReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QCompassReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompassReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompassReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCompassReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::customEvent(static_cast<QEvent*>(event));
}

void QCompassReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QCompassReading*>(ptr), "deleteLater");
}

void QCompassReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::deleteLater();
}

void QCompassReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QCompassReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QCompassReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QCompassReading*>(ptr)->QCompassReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QCompassReading_Event(void* ptr, void* e)
{
	return static_cast<QCompassReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QCompassReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QCompassReading*>(ptr)->QCompassReading::event(static_cast<QEvent*>(e));
}

int QCompassReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QCompassReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QCompassReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QCompassReading*>(ptr)->QCompassReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QCompassReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCompassReading*>(ptr)->metaObject());
}

void* QCompassReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QCompassReading*>(ptr)->QCompassReading::metaObject());
}

class MyQDistanceFilter: public QDistanceFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QDistanceReading * reading) { return callbackQDistanceFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QDistanceFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

char* QDistanceFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQDistanceFilter*>(static_cast<QDistanceFilter*>(ptr))) {
		return static_cast<MyQDistanceFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QDistanceFilter_BASE").toUtf8().data();
}

void QDistanceFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQDistanceFilter*>(static_cast<QDistanceFilter*>(ptr))) {
		static_cast<MyQDistanceFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QDistanceReading_Distance(void* ptr)
{
	return static_cast<double>(static_cast<QDistanceReading*>(ptr)->distance());
}

void QDistanceReading_SetDistance(void* ptr, double distance)
{
	static_cast<QDistanceReading*>(ptr)->setDistance(static_cast<double>(distance));
}

void QDistanceReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDistanceReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDistanceReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::customEvent(static_cast<QEvent*>(event));
}

void QDistanceReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDistanceReading*>(ptr), "deleteLater");
}

void QDistanceReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::deleteLater();
}

void QDistanceReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDistanceReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDistanceReading_Event(void* ptr, void* e)
{
	return static_cast<QDistanceReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDistanceReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QDistanceReading*>(ptr)->QDistanceReading::event(static_cast<QEvent*>(e));
}

int QDistanceReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDistanceReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDistanceReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDistanceReading*>(ptr)->QDistanceReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDistanceReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDistanceReading*>(ptr)->metaObject());
}

void* QDistanceReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDistanceReading*>(ptr)->QDistanceReading::metaObject());
}

void* QDistanceSensor_Reading(void* ptr)
{
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

void* QDistanceSensor_NewQDistanceSensor(void* parent)
{
	return new QDistanceSensor(static_cast<QObject*>(parent));
}

void QDistanceSensor_DestroyQDistanceSensor(void* ptr)
{
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

char* QDistanceSensor_QDistanceSensor_Type()
{
	return QString(QDistanceSensor::type).toUtf8().data();
}

int QDistanceSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QDistanceSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QDistanceSensor_StartDefault(void* ptr)
{
	return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::start();
}

void QDistanceSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDistanceSensor*>(ptr), "stop");
}

void QDistanceSensor_StopDefault(void* ptr)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::stop();
}

void QDistanceSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDistanceSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDistanceSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::customEvent(static_cast<QEvent*>(event));
}

void QDistanceSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDistanceSensor*>(ptr), "deleteLater");
}

void QDistanceSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::deleteLater();
}

void QDistanceSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDistanceSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDistanceSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDistanceSensor_Event(void* ptr, void* e)
{
	return static_cast<QDistanceSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDistanceSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::event(static_cast<QEvent*>(e));
}

int QDistanceSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDistanceSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDistanceSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDistanceSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDistanceSensor*>(ptr)->metaObject());
}

void* QDistanceSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::metaObject());
}

class MyQGyroscope: public QGyroscope
{
public:
	MyQGyroscope(QObject *parent) : QGyroscope(parent) {};
	bool start() { return callbackQGyroscope_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQGyroscope_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQGyroscope_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGyroscope_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGyroscope_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGyroscope_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGyroscope_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGyroscope_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQGyroscope_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGyroscope_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGyroscope_MetaObject(const_cast<MyQGyroscope*>(this), this->objectName().toUtf8().data())); };
};

void* QGyroscope_Reading(void* ptr)
{
	return static_cast<QGyroscope*>(ptr)->reading();
}

void* QGyroscope_NewQGyroscope(void* parent)
{
	return new MyQGyroscope(static_cast<QObject*>(parent));
}

void QGyroscope_DestroyQGyroscope(void* ptr)
{
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

char* QGyroscope_QGyroscope_Type()
{
	return QString(QGyroscope::type).toUtf8().data();
}

int QGyroscope_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QGyroscope*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QGyroscope_StartDefault(void* ptr)
{
	return static_cast<QGyroscope*>(ptr)->QGyroscope::start();
}

void QGyroscope_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGyroscope*>(ptr), "stop");
}

void QGyroscope_StopDefault(void* ptr)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::stop();
}

void QGyroscope_TimerEvent(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscope_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscope_ChildEvent(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscope_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscope_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGyroscope*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscope_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscope_CustomEvent(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGyroscope_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::customEvent(static_cast<QEvent*>(event));
}

void QGyroscope_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGyroscope*>(ptr), "deleteLater");
}

void QGyroscope_DeleteLaterDefault(void* ptr)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::deleteLater();
}

void QGyroscope_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGyroscope*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscope_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGyroscope*>(ptr)->QGyroscope::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGyroscope_Event(void* ptr, void* e)
{
	return static_cast<QGyroscope*>(ptr)->event(static_cast<QEvent*>(e));
}

int QGyroscope_EventDefault(void* ptr, void* e)
{
	return static_cast<QGyroscope*>(ptr)->QGyroscope::event(static_cast<QEvent*>(e));
}

int QGyroscope_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGyroscope*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGyroscope_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGyroscope*>(ptr)->QGyroscope::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGyroscope_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGyroscope*>(ptr)->metaObject());
}

void* QGyroscope_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGyroscope*>(ptr)->QGyroscope::metaObject());
}

class MyQGyroscopeFilter: public QGyroscopeFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QGyroscopeReading * reading) { return callbackQGyroscopeFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QGyroscopeFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

char* QGyroscopeFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQGyroscopeFilter*>(static_cast<QGyroscopeFilter*>(ptr))) {
		return static_cast<MyQGyroscopeFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QGyroscopeFilter_BASE").toUtf8().data();
}

void QGyroscopeFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQGyroscopeFilter*>(static_cast<QGyroscopeFilter*>(ptr))) {
		static_cast<MyQGyroscopeFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QGyroscopeReading_X(void* ptr)
{
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->x());
}

double QGyroscopeReading_Y(void* ptr)
{
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->y());
}

double QGyroscopeReading_Z(void* ptr)
{
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->z());
}

void QGyroscopeReading_SetX(void* ptr, double x)
{
	static_cast<QGyroscopeReading*>(ptr)->setX(static_cast<double>(x));
}

void QGyroscopeReading_SetY(void* ptr, double y)
{
	static_cast<QGyroscopeReading*>(ptr)->setY(static_cast<double>(y));
}

void QGyroscopeReading_SetZ(void* ptr, double z)
{
	static_cast<QGyroscopeReading*>(ptr)->setZ(static_cast<double>(z));
}

void QGyroscopeReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscopeReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscopeReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscopeReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscopeReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGyroscopeReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscopeReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscopeReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGyroscopeReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::customEvent(static_cast<QEvent*>(event));
}

void QGyroscopeReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGyroscopeReading*>(ptr), "deleteLater");
}

void QGyroscopeReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::deleteLater();
}

void QGyroscopeReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGyroscopeReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGyroscopeReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGyroscopeReading_Event(void* ptr, void* e)
{
	return static_cast<QGyroscopeReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QGyroscopeReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::event(static_cast<QEvent*>(e));
}

int QGyroscopeReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGyroscopeReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGyroscopeReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGyroscopeReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGyroscopeReading*>(ptr)->metaObject());
}

void* QGyroscopeReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::metaObject());
}

class MyQHolsterFilter: public QHolsterFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QHolsterReading * reading) { return callbackQHolsterFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QHolsterFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

char* QHolsterFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQHolsterFilter*>(static_cast<QHolsterFilter*>(ptr))) {
		return static_cast<MyQHolsterFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QHolsterFilter_BASE").toUtf8().data();
}

void QHolsterFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQHolsterFilter*>(static_cast<QHolsterFilter*>(ptr))) {
		static_cast<MyQHolsterFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QHolsterReading_Holstered(void* ptr)
{
	return static_cast<QHolsterReading*>(ptr)->holstered();
}

void QHolsterReading_SetHolstered(void* ptr, int holstered)
{
	static_cast<QHolsterReading*>(ptr)->setHolstered(holstered != 0);
}

void QHolsterReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHolsterReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHolsterReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::customEvent(static_cast<QEvent*>(event));
}

void QHolsterReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHolsterReading*>(ptr), "deleteLater");
}

void QHolsterReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::deleteLater();
}

void QHolsterReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHolsterReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHolsterReading_Event(void* ptr, void* e)
{
	return static_cast<QHolsterReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHolsterReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QHolsterReading*>(ptr)->QHolsterReading::event(static_cast<QEvent*>(e));
}

int QHolsterReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHolsterReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHolsterReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHolsterReading*>(ptr)->QHolsterReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHolsterReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHolsterReading*>(ptr)->metaObject());
}

void* QHolsterReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHolsterReading*>(ptr)->QHolsterReading::metaObject());
}

void* QHolsterSensor_Reading(void* ptr)
{
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

void* QHolsterSensor_NewQHolsterSensor(void* parent)
{
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr)
{
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

char* QHolsterSensor_QHolsterSensor_Type()
{
	return QString(QHolsterSensor::type).toUtf8().data();
}

int QHolsterSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QHolsterSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QHolsterSensor_StartDefault(void* ptr)
{
	return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::start();
}

void QHolsterSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHolsterSensor*>(ptr), "stop");
}

void QHolsterSensor_StopDefault(void* ptr)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::stop();
}

void QHolsterSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QHolsterSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHolsterSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::customEvent(static_cast<QEvent*>(event));
}

void QHolsterSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QHolsterSensor*>(ptr), "deleteLater");
}

void QHolsterSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::deleteLater();
}

void QHolsterSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QHolsterSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QHolsterSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QHolsterSensor_Event(void* ptr, void* e)
{
	return static_cast<QHolsterSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QHolsterSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::event(static_cast<QEvent*>(e));
}

int QHolsterSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QHolsterSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QHolsterSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QHolsterSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHolsterSensor*>(ptr)->metaObject());
}

void* QHolsterSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::metaObject());
}

class MyQIRProximityFilter: public QIRProximityFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QIRProximityReading * reading) { return callbackQIRProximityFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QIRProximityFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

char* QIRProximityFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQIRProximityFilter*>(static_cast<QIRProximityFilter*>(ptr))) {
		return static_cast<MyQIRProximityFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QIRProximityFilter_BASE").toUtf8().data();
}

void QIRProximityFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQIRProximityFilter*>(static_cast<QIRProximityFilter*>(ptr))) {
		static_cast<MyQIRProximityFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QIRProximityReading_Reflectance(void* ptr)
{
	return static_cast<double>(static_cast<QIRProximityReading*>(ptr)->reflectance());
}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance)
{
	static_cast<QIRProximityReading*>(ptr)->setReflectance(static_cast<double>(reflectance));
}

void QIRProximityReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximityReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximityReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximityReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximityReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QIRProximityReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximityReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximityReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIRProximityReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::customEvent(static_cast<QEvent*>(event));
}

void QIRProximityReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QIRProximityReading*>(ptr), "deleteLater");
}

void QIRProximityReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::deleteLater();
}

void QIRProximityReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QIRProximityReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximityReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QIRProximityReading_Event(void* ptr, void* e)
{
	return static_cast<QIRProximityReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QIRProximityReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::event(static_cast<QEvent*>(e));
}

int QIRProximityReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QIRProximityReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QIRProximityReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QIRProximityReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIRProximityReading*>(ptr)->metaObject());
}

void* QIRProximityReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::metaObject());
}

class MyQIRProximitySensor: public QIRProximitySensor
{
public:
	MyQIRProximitySensor(QObject *parent) : QIRProximitySensor(parent) {};
	bool start() { return callbackQIRProximitySensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQIRProximitySensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQIRProximitySensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIRProximitySensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQIRProximitySensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQIRProximitySensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQIRProximitySensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQIRProximitySensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQIRProximitySensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQIRProximitySensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQIRProximitySensor_MetaObject(const_cast<MyQIRProximitySensor*>(this), this->objectName().toUtf8().data())); };
};

void* QIRProximitySensor_Reading(void* ptr)
{
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

void* QIRProximitySensor_NewQIRProximitySensor(void* parent)
{
	return new MyQIRProximitySensor(static_cast<QObject*>(parent));
}

void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr)
{
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

char* QIRProximitySensor_QIRProximitySensor_Type()
{
	return QString(QIRProximitySensor::type).toUtf8().data();
}

int QIRProximitySensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QIRProximitySensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QIRProximitySensor_StartDefault(void* ptr)
{
	return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::start();
}

void QIRProximitySensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QIRProximitySensor*>(ptr), "stop");
}

void QIRProximitySensor_StopDefault(void* ptr)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::stop();
}

void QIRProximitySensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximitySensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximitySensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximitySensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximitySensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QIRProximitySensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximitySensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximitySensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIRProximitySensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::customEvent(static_cast<QEvent*>(event));
}

void QIRProximitySensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QIRProximitySensor*>(ptr), "deleteLater");
}

void QIRProximitySensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::deleteLater();
}

void QIRProximitySensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QIRProximitySensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIRProximitySensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QIRProximitySensor_Event(void* ptr, void* e)
{
	return static_cast<QIRProximitySensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QIRProximitySensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::event(static_cast<QEvent*>(e));
}

int QIRProximitySensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QIRProximitySensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QIRProximitySensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QIRProximitySensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIRProximitySensor*>(ptr)->metaObject());
}

void* QIRProximitySensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::metaObject());
}

class MyQLightFilter: public QLightFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QLightReading * reading) { return callbackQLightFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QLightFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

char* QLightFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQLightFilter*>(static_cast<QLightFilter*>(ptr))) {
		return static_cast<MyQLightFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QLightFilter_BASE").toUtf8().data();
}

void QLightFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQLightFilter*>(static_cast<QLightFilter*>(ptr))) {
		static_cast<MyQLightFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QLightReading_Lux(void* ptr)
{
	return static_cast<double>(static_cast<QLightReading*>(ptr)->lux());
}

void QLightReading_SetLux(void* ptr, double lux)
{
	static_cast<QLightReading*>(ptr)->setLux(static_cast<double>(lux));
}

void QLightReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->QLightReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLightReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->QLightReading::childEvent(static_cast<QChildEvent*>(event));
}

void QLightReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLightReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLightReading*>(ptr)->QLightReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLightReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLightReading*>(ptr)->QLightReading::customEvent(static_cast<QEvent*>(event));
}

void QLightReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLightReading*>(ptr), "deleteLater");
}

void QLightReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QLightReading*>(ptr)->QLightReading::deleteLater();
}

void QLightReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLightReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLightReading*>(ptr)->QLightReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QLightReading_Event(void* ptr, void* e)
{
	return static_cast<QLightReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QLightReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QLightReading*>(ptr)->QLightReading::event(static_cast<QEvent*>(e));
}

int QLightReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLightReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QLightReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLightReading*>(ptr)->QLightReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLightReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLightReading*>(ptr)->metaObject());
}

void* QLightReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLightReading*>(ptr)->QLightReading::metaObject());
}

class MyQLightSensor: public QLightSensor
{
public:
	MyQLightSensor(QObject *parent) : QLightSensor(parent) {};
	void Signal_FieldOfViewChanged(qreal fieldOfView) { callbackQLightSensor_FieldOfViewChanged(this, this->objectName().toUtf8().data(), static_cast<double>(fieldOfView)); };
	bool start() { return callbackQLightSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQLightSensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQLightSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLightSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQLightSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQLightSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQLightSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQLightSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQLightSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQLightSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQLightSensor_MetaObject(const_cast<MyQLightSensor*>(this), this->objectName().toUtf8().data())); };
};

double QLightSensor_FieldOfView(void* ptr)
{
	return static_cast<double>(static_cast<QLightSensor*>(ptr)->fieldOfView());
}

void* QLightSensor_Reading(void* ptr)
{
	return static_cast<QLightSensor*>(ptr)->reading();
}

void* QLightSensor_NewQLightSensor(void* parent)
{
	return new MyQLightSensor(static_cast<QObject*>(parent));
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
	static_cast<QLightSensor*>(ptr)->fieldOfViewChanged(static_cast<double>(fieldOfView));
}

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView)
{
	static_cast<QLightSensor*>(ptr)->setFieldOfView(static_cast<double>(fieldOfView));
}

void QLightSensor_DestroyQLightSensor(void* ptr)
{
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

char* QLightSensor_QLightSensor_Type()
{
	return QString(QLightSensor::type).toUtf8().data();
}

int QLightSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QLightSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QLightSensor_StartDefault(void* ptr)
{
	return static_cast<QLightSensor*>(ptr)->QLightSensor::start();
}

void QLightSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLightSensor*>(ptr), "stop");
}

void QLightSensor_StopDefault(void* ptr)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::stop();
}

void QLightSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLightSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QLightSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QLightSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLightSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::customEvent(static_cast<QEvent*>(event));
}

void QLightSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QLightSensor*>(ptr), "deleteLater");
}

void QLightSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::deleteLater();
}

void QLightSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QLightSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QLightSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QLightSensor*>(ptr)->QLightSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QLightSensor_Event(void* ptr, void* e)
{
	return static_cast<QLightSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QLightSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QLightSensor*>(ptr)->QLightSensor::event(static_cast<QEvent*>(e));
}

int QLightSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QLightSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QLightSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QLightSensor*>(ptr)->QLightSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QLightSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLightSensor*>(ptr)->metaObject());
}

void* QLightSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QLightSensor*>(ptr)->QLightSensor::metaObject());
}

class MyQMagnetometer: public QMagnetometer
{
public:
	MyQMagnetometer(QObject *parent) : QMagnetometer(parent) {};
	void Signal_ReturnGeoValuesChanged(bool returnGeoValues) { callbackQMagnetometer_ReturnGeoValuesChanged(this, this->objectName().toUtf8().data(), returnGeoValues); };
	bool start() { return callbackQMagnetometer_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQMagnetometer_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQMagnetometer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQMagnetometer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMagnetometer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQMagnetometer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQMagnetometer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMagnetometer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQMagnetometer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMagnetometer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMagnetometer_MetaObject(const_cast<MyQMagnetometer*>(this), this->objectName().toUtf8().data())); };
};

void* QMagnetometer_Reading(void* ptr)
{
	return static_cast<QMagnetometer*>(ptr)->reading();
}

int QMagnetometer_ReturnGeoValues(void* ptr)
{
	return static_cast<QMagnetometer*>(ptr)->returnGeoValues();
}

void QMagnetometer_SetReturnGeoValues(void* ptr, int returnGeoValues)
{
	static_cast<QMagnetometer*>(ptr)->setReturnGeoValues(returnGeoValues != 0);
}

void* QMagnetometer_NewQMagnetometer(void* parent)
{
	return new MyQMagnetometer(static_cast<QObject*>(parent));
}

void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr)
{
	QObject::connect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));
}

void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));
}

void QMagnetometer_ReturnGeoValuesChanged(void* ptr, int returnGeoValues)
{
	static_cast<QMagnetometer*>(ptr)->returnGeoValuesChanged(returnGeoValues != 0);
}

void QMagnetometer_DestroyQMagnetometer(void* ptr)
{
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
}

char* QMagnetometer_QMagnetometer_Type()
{
	return QString(QMagnetometer::type).toUtf8().data();
}

int QMagnetometer_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QMagnetometer*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QMagnetometer_StartDefault(void* ptr)
{
	return static_cast<QMagnetometer*>(ptr)->QMagnetometer::start();
}

void QMagnetometer_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMagnetometer*>(ptr), "stop");
}

void QMagnetometer_StopDefault(void* ptr)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::stop();
}

void QMagnetometer_TimerEvent(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometer_ChildEvent(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMagnetometer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometer_CustomEvent(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMagnetometer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::customEvent(static_cast<QEvent*>(event));
}

void QMagnetometer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMagnetometer*>(ptr), "deleteLater");
}

void QMagnetometer_DeleteLaterDefault(void* ptr)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::deleteLater();
}

void QMagnetometer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMagnetometer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QMagnetometer_Event(void* ptr, void* e)
{
	return static_cast<QMagnetometer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QMagnetometer_EventDefault(void* ptr, void* e)
{
	return static_cast<QMagnetometer*>(ptr)->QMagnetometer::event(static_cast<QEvent*>(e));
}

int QMagnetometer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMagnetometer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QMagnetometer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMagnetometer*>(ptr)->QMagnetometer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMagnetometer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMagnetometer*>(ptr)->metaObject());
}

void* QMagnetometer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMagnetometer*>(ptr)->QMagnetometer::metaObject());
}

class MyQMagnetometerFilter: public QMagnetometerFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QMagnetometerReading * reading) { return callbackQMagnetometerFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QMagnetometerFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

char* QMagnetometerFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQMagnetometerFilter*>(static_cast<QMagnetometerFilter*>(ptr))) {
		return static_cast<MyQMagnetometerFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QMagnetometerFilter_BASE").toUtf8().data();
}

void QMagnetometerFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQMagnetometerFilter*>(static_cast<QMagnetometerFilter*>(ptr))) {
		static_cast<MyQMagnetometerFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QMagnetometerReading_CalibrationLevel(void* ptr)
{
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->calibrationLevel());
}

double QMagnetometerReading_X(void* ptr)
{
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->x());
}

double QMagnetometerReading_Y(void* ptr)
{
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->y());
}

double QMagnetometerReading_Z(void* ptr)
{
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->z());
}

void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel)
{
	static_cast<QMagnetometerReading*>(ptr)->setCalibrationLevel(static_cast<double>(calibrationLevel));
}

void QMagnetometerReading_SetX(void* ptr, double x)
{
	static_cast<QMagnetometerReading*>(ptr)->setX(static_cast<double>(x));
}

void QMagnetometerReading_SetY(void* ptr, double y)
{
	static_cast<QMagnetometerReading*>(ptr)->setY(static_cast<double>(y));
}

void QMagnetometerReading_SetZ(void* ptr, double z)
{
	static_cast<QMagnetometerReading*>(ptr)->setZ(static_cast<double>(z));
}

void QMagnetometerReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometerReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometerReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometerReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometerReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMagnetometerReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometerReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometerReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMagnetometerReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::customEvent(static_cast<QEvent*>(event));
}

void QMagnetometerReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMagnetometerReading*>(ptr), "deleteLater");
}

void QMagnetometerReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::deleteLater();
}

void QMagnetometerReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMagnetometerReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMagnetometerReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QMagnetometerReading_Event(void* ptr, void* e)
{
	return static_cast<QMagnetometerReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QMagnetometerReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::event(static_cast<QEvent*>(e));
}

int QMagnetometerReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMagnetometerReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QMagnetometerReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMagnetometerReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMagnetometerReading*>(ptr)->metaObject());
}

void* QMagnetometerReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::metaObject());
}

class MyQOrientationFilter: public QOrientationFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QOrientationReading * reading) { return callbackQOrientationFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QOrientationFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QOrientationFilter*>(ptr)->filter(static_cast<QOrientationReading*>(reading));
}

char* QOrientationFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQOrientationFilter*>(static_cast<QOrientationFilter*>(ptr))) {
		return static_cast<MyQOrientationFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QOrientationFilter_BASE").toUtf8().data();
}

void QOrientationFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQOrientationFilter*>(static_cast<QOrientationFilter*>(ptr))) {
		static_cast<MyQOrientationFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QOrientationReading_Orientation(void* ptr)
{
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

void QOrientationReading_SetOrientation(void* ptr, int orientation)
{
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

void QOrientationReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QOrientationReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QOrientationReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::customEvent(static_cast<QEvent*>(event));
}

void QOrientationReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QOrientationReading*>(ptr), "deleteLater");
}

void QOrientationReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::deleteLater();
}

void QOrientationReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QOrientationReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QOrientationReading_Event(void* ptr, void* e)
{
	return static_cast<QOrientationReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QOrientationReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QOrientationReading*>(ptr)->QOrientationReading::event(static_cast<QEvent*>(e));
}

int QOrientationReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QOrientationReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QOrientationReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QOrientationReading*>(ptr)->QOrientationReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QOrientationReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOrientationReading*>(ptr)->metaObject());
}

void* QOrientationReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOrientationReading*>(ptr)->QOrientationReading::metaObject());
}

class MyQOrientationSensor: public QOrientationSensor
{
public:
	MyQOrientationSensor(QObject *parent) : QOrientationSensor(parent) {};
	bool start() { return callbackQOrientationSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQOrientationSensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQOrientationSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQOrientationSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQOrientationSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQOrientationSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQOrientationSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQOrientationSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQOrientationSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQOrientationSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQOrientationSensor_MetaObject(const_cast<MyQOrientationSensor*>(this), this->objectName().toUtf8().data())); };
};

void* QOrientationSensor_Reading(void* ptr)
{
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

void* QOrientationSensor_NewQOrientationSensor(void* parent)
{
	return new MyQOrientationSensor(static_cast<QObject*>(parent));
}

void QOrientationSensor_DestroyQOrientationSensor(void* ptr)
{
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

char* QOrientationSensor_QOrientationSensor_Type()
{
	return QString(QOrientationSensor::type).toUtf8().data();
}

int QOrientationSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QOrientationSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QOrientationSensor_StartDefault(void* ptr)
{
	return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::start();
}

void QOrientationSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QOrientationSensor*>(ptr), "stop");
}

void QOrientationSensor_StopDefault(void* ptr)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::stop();
}

void QOrientationSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QOrientationSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QOrientationSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::customEvent(static_cast<QEvent*>(event));
}

void QOrientationSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QOrientationSensor*>(ptr), "deleteLater");
}

void QOrientationSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::deleteLater();
}

void QOrientationSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QOrientationSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOrientationSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QOrientationSensor_Event(void* ptr, void* e)
{
	return static_cast<QOrientationSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QOrientationSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::event(static_cast<QEvent*>(e));
}

int QOrientationSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QOrientationSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QOrientationSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QOrientationSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOrientationSensor*>(ptr)->metaObject());
}

void* QOrientationSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::metaObject());
}

class MyQPressureFilter: public QPressureFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QPressureReading * reading) { return callbackQPressureFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QPressureFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QPressureFilter*>(ptr)->filter(static_cast<QPressureReading*>(reading));
}

char* QPressureFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQPressureFilter*>(static_cast<QPressureFilter*>(ptr))) {
		return static_cast<MyQPressureFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QPressureFilter_BASE").toUtf8().data();
}

void QPressureFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQPressureFilter*>(static_cast<QPressureFilter*>(ptr))) {
		static_cast<MyQPressureFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QPressureReading_Pressure(void* ptr)
{
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->pressure());
}

double QPressureReading_Temperature(void* ptr)
{
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->temperature());
}

void QPressureReading_SetPressure(void* ptr, double pressure)
{
	static_cast<QPressureReading*>(ptr)->setPressure(static_cast<double>(pressure));
}

void QPressureReading_SetTemperature(void* ptr, double temperature)
{
	static_cast<QPressureReading*>(ptr)->setTemperature(static_cast<double>(temperature));
}

void QPressureReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPressureReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::childEvent(static_cast<QChildEvent*>(event));
}

void QPressureReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPressureReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPressureReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::customEvent(static_cast<QEvent*>(event));
}

void QPressureReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPressureReading*>(ptr), "deleteLater");
}

void QPressureReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::deleteLater();
}

void QPressureReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPressureReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPressureReading*>(ptr)->QPressureReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QPressureReading_Event(void* ptr, void* e)
{
	return static_cast<QPressureReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QPressureReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QPressureReading*>(ptr)->QPressureReading::event(static_cast<QEvent*>(e));
}

int QPressureReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPressureReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QPressureReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPressureReading*>(ptr)->QPressureReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPressureReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPressureReading*>(ptr)->metaObject());
}

void* QPressureReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPressureReading*>(ptr)->QPressureReading::metaObject());
}

void* QPressureSensor_Reading(void* ptr)
{
	return static_cast<QPressureSensor*>(ptr)->reading();
}

void* QPressureSensor_NewQPressureSensor(void* parent)
{
	return new QPressureSensor(static_cast<QObject*>(parent));
}

void QPressureSensor_DestroyQPressureSensor(void* ptr)
{
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

char* QPressureSensor_QPressureSensor_Type()
{
	return QString(QPressureSensor::type).toUtf8().data();
}

int QPressureSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPressureSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QPressureSensor_StartDefault(void* ptr)
{
	return static_cast<QPressureSensor*>(ptr)->QPressureSensor::start();
}

void QPressureSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPressureSensor*>(ptr), "stop");
}

void QPressureSensor_StopDefault(void* ptr)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::stop();
}

void QPressureSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPressureSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QPressureSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPressureSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPressureSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::customEvent(static_cast<QEvent*>(event));
}

void QPressureSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPressureSensor*>(ptr), "deleteLater");
}

void QPressureSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::deleteLater();
}

void QPressureSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPressureSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPressureSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QPressureSensor_Event(void* ptr, void* e)
{
	return static_cast<QPressureSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QPressureSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QPressureSensor*>(ptr)->QPressureSensor::event(static_cast<QEvent*>(e));
}

int QPressureSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPressureSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QPressureSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPressureSensor*>(ptr)->QPressureSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPressureSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPressureSensor*>(ptr)->metaObject());
}

void* QPressureSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPressureSensor*>(ptr)->QPressureSensor::metaObject());
}

class MyQProximityFilter: public QProximityFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QProximityReading * reading) { return callbackQProximityFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QProximityFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

char* QProximityFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQProximityFilter*>(static_cast<QProximityFilter*>(ptr))) {
		return static_cast<MyQProximityFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QProximityFilter_BASE").toUtf8().data();
}

void QProximityFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQProximityFilter*>(static_cast<QProximityFilter*>(ptr))) {
		static_cast<MyQProximityFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QProximityReading_Close(void* ptr)
{
	return static_cast<QProximityReading*>(ptr)->close();
}

void QProximityReading_SetClose(void* ptr, int close)
{
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

void QProximityReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximityReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximityReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QProximityReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::childEvent(static_cast<QChildEvent*>(event));
}

void QProximityReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QProximityReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximityReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximityReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QProximityReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::customEvent(static_cast<QEvent*>(event));
}

void QProximityReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QProximityReading*>(ptr), "deleteLater");
}

void QProximityReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::deleteLater();
}

void QProximityReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QProximityReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximityReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QProximityReading*>(ptr)->QProximityReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QProximityReading_Event(void* ptr, void* e)
{
	return static_cast<QProximityReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QProximityReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QProximityReading*>(ptr)->QProximityReading::event(static_cast<QEvent*>(e));
}

int QProximityReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QProximityReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QProximityReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QProximityReading*>(ptr)->QProximityReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QProximityReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QProximityReading*>(ptr)->metaObject());
}

void* QProximityReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QProximityReading*>(ptr)->QProximityReading::metaObject());
}

class MyQProximitySensor: public QProximitySensor
{
public:
	MyQProximitySensor(QObject *parent) : QProximitySensor(parent) {};
	bool start() { return callbackQProximitySensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQProximitySensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQProximitySensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQProximitySensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQProximitySensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQProximitySensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQProximitySensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQProximitySensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQProximitySensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQProximitySensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQProximitySensor_MetaObject(const_cast<MyQProximitySensor*>(this), this->objectName().toUtf8().data())); };
};

void* QProximitySensor_Reading(void* ptr)
{
	return static_cast<QProximitySensor*>(ptr)->reading();
}

void* QProximitySensor_NewQProximitySensor(void* parent)
{
	return new MyQProximitySensor(static_cast<QObject*>(parent));
}

void QProximitySensor_DestroyQProximitySensor(void* ptr)
{
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

char* QProximitySensor_QProximitySensor_Type()
{
	return QString(QProximitySensor::type).toUtf8().data();
}

int QProximitySensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QProximitySensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QProximitySensor_StartDefault(void* ptr)
{
	return static_cast<QProximitySensor*>(ptr)->QProximitySensor::start();
}

void QProximitySensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QProximitySensor*>(ptr), "stop");
}

void QProximitySensor_StopDefault(void* ptr)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::stop();
}

void QProximitySensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximitySensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximitySensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QProximitySensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::childEvent(static_cast<QChildEvent*>(event));
}

void QProximitySensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QProximitySensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximitySensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximitySensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QProximitySensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::customEvent(static_cast<QEvent*>(event));
}

void QProximitySensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QProximitySensor*>(ptr), "deleteLater");
}

void QProximitySensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::deleteLater();
}

void QProximitySensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QProximitySensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QProximitySensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QProximitySensor_Event(void* ptr, void* e)
{
	return static_cast<QProximitySensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QProximitySensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QProximitySensor*>(ptr)->QProximitySensor::event(static_cast<QEvent*>(e));
}

int QProximitySensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QProximitySensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QProximitySensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QProximitySensor*>(ptr)->QProximitySensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QProximitySensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QProximitySensor*>(ptr)->metaObject());
}

void* QProximitySensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QProximitySensor*>(ptr)->QProximitySensor::metaObject());
}

class MyQRotationFilter: public QRotationFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QRotationReading * reading) { return callbackQRotationFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QRotationFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

char* QRotationFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQRotationFilter*>(static_cast<QRotationFilter*>(ptr))) {
		return static_cast<MyQRotationFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QRotationFilter_BASE").toUtf8().data();
}

void QRotationFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQRotationFilter*>(static_cast<QRotationFilter*>(ptr))) {
		static_cast<MyQRotationFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QRotationReading_X(void* ptr)
{
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->x());
}

double QRotationReading_Y(void* ptr)
{
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->y());
}

double QRotationReading_Z(void* ptr)
{
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->z());
}

void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z)
{
	static_cast<QRotationReading*>(ptr)->setFromEuler(static_cast<double>(x), static_cast<double>(y), static_cast<double>(z));
}

void QRotationReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRotationReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::childEvent(static_cast<QChildEvent*>(event));
}

void QRotationReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRotationReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRotationReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::customEvent(static_cast<QEvent*>(event));
}

void QRotationReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRotationReading*>(ptr), "deleteLater");
}

void QRotationReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::deleteLater();
}

void QRotationReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRotationReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRotationReading*>(ptr)->QRotationReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QRotationReading_Event(void* ptr, void* e)
{
	return static_cast<QRotationReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QRotationReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QRotationReading*>(ptr)->QRotationReading::event(static_cast<QEvent*>(e));
}

int QRotationReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRotationReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QRotationReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRotationReading*>(ptr)->QRotationReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRotationReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRotationReading*>(ptr)->metaObject());
}

void* QRotationReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRotationReading*>(ptr)->QRotationReading::metaObject());
}

class MyQRotationSensor: public QRotationSensor
{
public:
	MyQRotationSensor(QObject *parent) : QRotationSensor(parent) {};
	void Signal_HasZChanged(bool hasZ) { callbackQRotationSensor_HasZChanged(this, this->objectName().toUtf8().data(), hasZ); };
	bool start() { return callbackQRotationSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQRotationSensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQRotationSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQRotationSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRotationSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQRotationSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQRotationSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRotationSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQRotationSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRotationSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRotationSensor_MetaObject(const_cast<MyQRotationSensor*>(this), this->objectName().toUtf8().data())); };
};

int QRotationSensor_HasZ(void* ptr)
{
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

void* QRotationSensor_Reading(void* ptr)
{
	return static_cast<QRotationSensor*>(ptr)->reading();
}

void* QRotationSensor_NewQRotationSensor(void* parent)
{
	return new MyQRotationSensor(static_cast<QObject*>(parent));
}

void QRotationSensor_ConnectHasZChanged(void* ptr)
{
	QObject::connect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));
}

void QRotationSensor_DisconnectHasZChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));
}

void QRotationSensor_HasZChanged(void* ptr, int hasZ)
{
	static_cast<QRotationSensor*>(ptr)->hasZChanged(hasZ != 0);
}

void QRotationSensor_SetHasZ(void* ptr, int hasZ)
{
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(void* ptr)
{
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

char* QRotationSensor_QRotationSensor_Type()
{
	return QString(QRotationSensor::type).toUtf8().data();
}

int QRotationSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QRotationSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QRotationSensor_StartDefault(void* ptr)
{
	return static_cast<QRotationSensor*>(ptr)->QRotationSensor::start();
}

void QRotationSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRotationSensor*>(ptr), "stop");
}

void QRotationSensor_StopDefault(void* ptr)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::stop();
}

void QRotationSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRotationSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QRotationSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRotationSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRotationSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::customEvent(static_cast<QEvent*>(event));
}

void QRotationSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRotationSensor*>(ptr), "deleteLater");
}

void QRotationSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::deleteLater();
}

void QRotationSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRotationSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRotationSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QRotationSensor_Event(void* ptr, void* e)
{
	return static_cast<QRotationSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QRotationSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QRotationSensor*>(ptr)->QRotationSensor::event(static_cast<QEvent*>(e));
}

int QRotationSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRotationSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QRotationSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRotationSensor*>(ptr)->QRotationSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRotationSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRotationSensor*>(ptr)->metaObject());
}

void* QRotationSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRotationSensor*>(ptr)->QRotationSensor::metaObject());
}

class MyQSensor: public QSensor
{
public:
	MyQSensor(const QByteArray &type, QObject *parent) : QSensor(type, parent) {};
	void Signal_ActiveChanged() { callbackQSensor_ActiveChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AlwaysOnChanged() { callbackQSensor_AlwaysOnChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AvailableSensorsChanged() { callbackQSensor_AvailableSensorsChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensor_AxesOrientationModeChanged(this, this->objectName().toUtf8().data(), axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensor_BufferSizeChanged(this, this->objectName().toUtf8().data(), bufferSize); };
	void Signal_BusyChanged() { callbackQSensor_BusyChanged(this, this->objectName().toUtf8().data()); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensor_CurrentOrientationChanged(this, this->objectName().toUtf8().data(), currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensor_DataRateChanged(this, this->objectName().toUtf8().data()); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensor_EfficientBufferSizeChanged(this, this->objectName().toUtf8().data(), efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensor_MaxBufferSizeChanged(this, this->objectName().toUtf8().data(), maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensor_ReadingChanged(this, this->objectName().toUtf8().data()); };
	void Signal_SensorError(int error) { callbackQSensor_SensorError(this, this->objectName().toUtf8().data(), error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensor_SkipDuplicatesChanged(this, this->objectName().toUtf8().data(), skipDuplicates); };
	bool start() { return callbackQSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQSensor_Stop(this, this->objectName().toUtf8().data()); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensor_UserOrientationChanged(this, this->objectName().toUtf8().data(), userOrientation); };
	void timerEvent(QTimerEvent * event) { callbackQSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensor_MetaObject(const_cast<MyQSensor*>(this), this->objectName().toUtf8().data())); };
};

int QSensor_AxesOrientationMode(void* ptr)
{
	return static_cast<QSensor*>(ptr)->axesOrientationMode();
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

char* QSensor_Description(void* ptr)
{
	return static_cast<QSensor*>(ptr)->description().toUtf8().data();
}

int QSensor_EfficientBufferSize(void* ptr)
{
	return static_cast<QSensor*>(ptr)->efficientBufferSize();
}

int QSensor_Error(void* ptr)
{
	return static_cast<QSensor*>(ptr)->error();
}

char* QSensor_Identifier(void* ptr)
{
	return QString(static_cast<QSensor*>(ptr)->identifier()).toUtf8().data();
}

int QSensor_IsActive(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isActive();
}

int QSensor_IsAlwaysOn(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isAlwaysOn();
}

int QSensor_IsBusy(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isBusy();
}

int QSensor_IsConnectedToBackend(void* ptr)
{
	return static_cast<QSensor*>(ptr)->isConnectedToBackend();
}

int QSensor_MaxBufferSize(void* ptr)
{
	return static_cast<QSensor*>(ptr)->maxBufferSize();
}

int QSensor_OutputRange(void* ptr)
{
	return static_cast<QSensor*>(ptr)->outputRange();
}

void* QSensor_Reading(void* ptr)
{
	return static_cast<QSensor*>(ptr)->reading();
}

void QSensor_SetActive(void* ptr, int active)
{
	static_cast<QSensor*>(ptr)->setActive(active != 0);
}

void QSensor_SetAlwaysOn(void* ptr, int alwaysOn)
{
	static_cast<QSensor*>(ptr)->setAlwaysOn(alwaysOn != 0);
}

void QSensor_SetAxesOrientationMode(void* ptr, int axesOrientationMode)
{
	static_cast<QSensor*>(ptr)->setAxesOrientationMode(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_SetBufferSize(void* ptr, int bufferSize)
{
	static_cast<QSensor*>(ptr)->setBufferSize(bufferSize);
}

void QSensor_SetDataRate(void* ptr, int rate)
{
	static_cast<QSensor*>(ptr)->setDataRate(rate);
}

void QSensor_SetIdentifier(void* ptr, char* identifier)
{
	static_cast<QSensor*>(ptr)->setIdentifier(QByteArray(identifier));
}

void QSensor_SetOutputRange(void* ptr, int index)
{
	static_cast<QSensor*>(ptr)->setOutputRange(index);
}

void QSensor_SetUserOrientation(void* ptr, int userOrientation)
{
	static_cast<QSensor*>(ptr)->setUserOrientation(userOrientation);
}

int QSensor_SkipDuplicates(void* ptr)
{
	return static_cast<QSensor*>(ptr)->skipDuplicates();
}

char* QSensor_Type(void* ptr)
{
	return QString(static_cast<QSensor*>(ptr)->type()).toUtf8().data();
}

int QSensor_UserOrientation(void* ptr)
{
	return static_cast<QSensor*>(ptr)->userOrientation();
}

void* QSensor_NewQSensor(char* ty, void* parent)
{
	return new MyQSensor(QByteArray(ty), static_cast<QObject*>(parent));
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

void QSensor_AxesOrientationModeChanged(void* ptr, int axesOrientationMode)
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

int QSensor_ConnectToBackend(void* ptr)
{
	return static_cast<QSensor*>(ptr)->connectToBackend();
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

char* QSensor_QSensor_DefaultSensorForType(char* ty)
{
	return QString(QSensor::defaultSensorForType(QByteArray(ty))).toUtf8().data();
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

int QSensor_IsFeatureSupported(void* ptr, int feature)
{
	return static_cast<QSensor*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
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

void QSensor_SetCurrentOrientation(void* ptr, int currentOrientation)
{
	static_cast<QSensor*>(ptr)->setCurrentOrientation(currentOrientation);
}

void QSensor_SetEfficientBufferSize(void* ptr, int efficientBufferSize)
{
	static_cast<QSensor*>(ptr)->setEfficientBufferSize(efficientBufferSize);
}

void QSensor_SetMaxBufferSize(void* ptr, int maxBufferSize)
{
	static_cast<QSensor*>(ptr)->setMaxBufferSize(maxBufferSize);
}

void QSensor_SetSkipDuplicates(void* ptr, int skipDuplicates)
{
	static_cast<QSensor*>(ptr)->setSkipDuplicates(skipDuplicates != 0);
}

void QSensor_ConnectSkipDuplicatesChanged(void* ptr)
{
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));
}

void QSensor_DisconnectSkipDuplicatesChanged(void* ptr)
{
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));
}

void QSensor_SkipDuplicatesChanged(void* ptr, int skipDuplicates)
{
	static_cast<QSensor*>(ptr)->skipDuplicatesChanged(skipDuplicates != 0);
}

int QSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

void QSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "stop");
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

void QSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->QSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->QSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensor*>(ptr)->QSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensor*>(ptr)->QSensor::customEvent(static_cast<QEvent*>(event));
}

void QSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "deleteLater");
}

void QSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensor*>(ptr)->QSensor::deleteLater();
}

void QSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensor*>(ptr)->QSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensor_Event(void* ptr, void* e)
{
	return static_cast<QSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensor*>(ptr)->QSensor::event(static_cast<QEvent*>(e));
}

int QSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensor*>(ptr)->QSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensor*>(ptr)->metaObject());
}

void* QSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensor*>(ptr)->QSensor::metaObject());
}

class MyQSensorBackend: public QSensorBackend
{
public:
	bool isFeatureSupported(QSensor::Feature feature) const { return callbackQSensorBackend_IsFeatureSupported(const_cast<MyQSensorBackend*>(this), this->objectName().toUtf8().data(), feature) != 0; };
	void start() { callbackQSensorBackend_Start(this, this->objectName().toUtf8().data()); };
	void stop() { callbackQSensorBackend_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSensorBackend_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorBackend_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorBackend_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSensorBackend_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSensorBackend_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorBackend_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSensorBackend_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorBackend_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorBackend_MetaObject(const_cast<MyQSensorBackend*>(this), this->objectName().toUtf8().data())); };
};

void QSensorBackend_AddDataRate(void* ptr, double min, double max)
{
	static_cast<QSensorBackend*>(ptr)->addDataRate(static_cast<double>(min), static_cast<double>(max));
}

int QSensorBackend_IsFeatureSupported(void* ptr, int feature)
{
	return static_cast<QSensorBackend*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

int QSensorBackend_IsFeatureSupportedDefault(void* ptr, int feature)
{
	return static_cast<QSensorBackend*>(ptr)->QSensorBackend::isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensorBackend_SensorBusy(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->sensorBusy();
}

void QSensorBackend_SensorError(void* ptr, int error)
{
	static_cast<QSensorBackend*>(ptr)->sensorError(error);
}

void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy)
{
	static_cast<QSensorBackend*>(ptr)->addOutputRange(static_cast<double>(min), static_cast<double>(max), static_cast<double>(accuracy));
}

void QSensorBackend_NewReadingAvailable(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->newReadingAvailable();
}

void* QSensorBackend_Reading(void* ptr)
{
	return static_cast<QSensorBackend*>(ptr)->reading();
}

void* QSensorBackend_Sensor(void* ptr)
{
	return static_cast<QSensorBackend*>(ptr)->sensor();
}

void QSensorBackend_SensorStopped(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->sensorStopped();
}

void QSensorBackend_SetDataRates(void* ptr, void* otherSensor)
{
	static_cast<QSensorBackend*>(ptr)->setDataRates(static_cast<QSensor*>(otherSensor));
}

void QSensorBackend_SetDescription(void* ptr, char* description)
{
	static_cast<QSensorBackend*>(ptr)->setDescription(QString(description));
}

void QSensorBackend_Start(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->start();
}

void QSensorBackend_Stop(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->stop();
}

void QSensorBackend_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorBackend_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorBackend_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorBackend_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorBackend_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorBackend*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorBackend_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorBackend_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorBackend_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::customEvent(static_cast<QEvent*>(event));
}

void QSensorBackend_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensorBackend*>(ptr), "deleteLater");
}

void QSensorBackend_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::deleteLater();
}

void QSensorBackend_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorBackend*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorBackend_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensorBackend_Event(void* ptr, void* e)
{
	return static_cast<QSensorBackend*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensorBackend_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensorBackend*>(ptr)->QSensorBackend::event(static_cast<QEvent*>(e));
}

int QSensorBackend_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorBackend*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensorBackend_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorBackend*>(ptr)->QSensorBackend::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensorBackend_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorBackend*>(ptr)->metaObject());
}

void* QSensorBackend_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorBackend*>(ptr)->QSensorBackend::metaObject());
}

class MyQSensorBackendFactory: public QSensorBackendFactory
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QSensorBackend * createBackend(QSensor * sensor) { return static_cast<QSensorBackend*>(callbackQSensorBackendFactory_CreateBackend(this, this->objectNameAbs().toUtf8().data(), sensor)); };
};

void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor)
{
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

char* QSensorBackendFactory_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSensorBackendFactory*>(static_cast<QSensorBackendFactory*>(ptr))) {
		return static_cast<MyQSensorBackendFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorBackendFactory_BASE").toUtf8().data();
}

void QSensorBackendFactory_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSensorBackendFactory*>(static_cast<QSensorBackendFactory*>(ptr))) {
		static_cast<MyQSensorBackendFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorChangesInterface: public QSensorChangesInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void sensorsChanged() { callbackQSensorChangesInterface_SensorsChanged(this, this->objectNameAbs().toUtf8().data()); };
};

void QSensorChangesInterface_SensorsChanged(void* ptr)
{
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

char* QSensorChangesInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSensorChangesInterface*>(static_cast<QSensorChangesInterface*>(ptr))) {
		return static_cast<MyQSensorChangesInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorChangesInterface_BASE").toUtf8().data();
}

void QSensorChangesInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSensorChangesInterface*>(static_cast<QSensorChangesInterface*>(ptr))) {
		static_cast<MyQSensorChangesInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorFilter: public QSensorFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QSensorReading * reading) { return callbackQSensorFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QSensorFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

void QSensorFilter_DestroyQSensorFilter(void* ptr)
{
	static_cast<QSensorFilter*>(ptr)->~QSensorFilter();
}

void* QSensorFilter_M_sensor(void* ptr)
{
	return static_cast<QSensorFilter*>(ptr)->m_sensor;
}

void QSensorFilter_SetM_sensor(void* ptr, void* vqs)
{
	static_cast<QSensorFilter*>(ptr)->m_sensor = static_cast<QSensor*>(vqs);
}

char* QSensorFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSensorFilter*>(static_cast<QSensorFilter*>(ptr))) {
		return static_cast<MyQSensorFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorFilter_BASE").toUtf8().data();
}

void QSensorFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSensorFilter*>(static_cast<QSensorFilter*>(ptr))) {
		static_cast<MyQSensorFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorGesture: public QSensorGesture
{
public:
	MyQSensorGesture(const QStringList &ids, QObject *parent) : QSensorGesture(ids, parent) {};
	#ifdef Q_QDOC
		void Signal_Detected(QString gestureId) { callbackQSensorGesture_Detected(this, this->objectName().toUtf8().data(), gestureId.toUtf8().data()); };
	#endif
	void timerEvent(QTimerEvent * event) { callbackQSensorGesture_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGesture_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGesture_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSensorGesture_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSensorGesture_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGesture_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSensorGesture_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGesture_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGesture_MetaObject(const_cast<MyQSensorGesture*>(this), this->objectName().toUtf8().data())); };
};

void* QSensorGesture_NewQSensorGesture(char* ids, void* parent)
{
	return new MyQSensorGesture(QString(ids).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
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

void QSensorGesture_Detected(void* ptr, char* gestureId)
{
#ifdef Q_QDOC
	static_cast<QSensorGesture*>(ptr)->detected(QString(gestureId));
#endif
}

char* QSensorGesture_GestureSignals(void* ptr)
{
	return static_cast<QSensorGesture*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGesture_InvalidIds(void* ptr)
{
	return static_cast<QSensorGesture*>(ptr)->invalidIds().join("|").toUtf8().data();
}

int QSensorGesture_IsActive(void* ptr)
{
	return static_cast<QSensorGesture*>(ptr)->isActive();
}

void QSensorGesture_StartDetection(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->startDetection();
}

void QSensorGesture_StopDetection(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->stopDetection();
}

char* QSensorGesture_ValidIds(void* ptr)
{
	return static_cast<QSensorGesture*>(ptr)->validIds().join("|").toUtf8().data();
}

void QSensorGesture_DestroyQSensorGesture(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

void QSensorGesture_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGesture_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGesture_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGesture_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGesture_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGesture*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGesture_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGesture_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGesture_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::customEvent(static_cast<QEvent*>(event));
}

void QSensorGesture_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensorGesture*>(ptr), "deleteLater");
}

void QSensorGesture_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::deleteLater();
}

void QSensorGesture_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGesture*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGesture_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensorGesture_Event(void* ptr, void* e)
{
	return static_cast<QSensorGesture*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensorGesture_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensorGesture*>(ptr)->QSensorGesture::event(static_cast<QEvent*>(e));
}

int QSensorGesture_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGesture*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensorGesture_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGesture*>(ptr)->QSensorGesture::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensorGesture_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGesture*>(ptr)->metaObject());
}

void* QSensorGesture_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGesture*>(ptr)->QSensorGesture::metaObject());
}

class MyQSensorGestureManager: public QSensorGestureManager
{
public:
	MyQSensorGestureManager(QObject *parent) : QSensorGestureManager(parent) {};
	void Signal_NewSensorGestureAvailable() { callbackQSensorGestureManager_NewSensorGestureAvailable(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureManager_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGestureManager_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGestureManager_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSensorGestureManager_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSensorGestureManager_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGestureManager_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSensorGestureManager_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGestureManager_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGestureManager_MetaObject(const_cast<MyQSensorGestureManager*>(this), this->objectName().toUtf8().data())); };
};

void* QSensorGestureManager_NewQSensorGestureManager(void* parent)
{
	return new MyQSensorGestureManager(static_cast<QObject*>(parent));
}

char* QSensorGestureManager_GestureIds(void* ptr)
{
	return static_cast<QSensorGestureManager*>(ptr)->gestureIds().join("|").toUtf8().data();
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

char* QSensorGestureManager_RecognizerSignals(void* ptr, char* gestureId)
{
	return static_cast<QSensorGestureManager*>(ptr)->recognizerSignals(QString(gestureId)).join("|").toUtf8().data();
}

int QSensorGestureManager_RegisterSensorGestureRecognizer(void* ptr, void* recognizer)
{
	return static_cast<QSensorGestureManager*>(ptr)->registerSensorGestureRecognizer(static_cast<QSensorGestureRecognizer*>(recognizer));
}

void* QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id)
{
	return QSensorGestureManager::sensorGestureRecognizer(QString(id));
}

void QSensorGestureManager_DestroyQSensorGestureManager(void* ptr)
{
	static_cast<QSensorGestureManager*>(ptr)->~QSensorGestureManager();
}

void QSensorGestureManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGestureManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensorGestureManager*>(ptr), "deleteLater");
}

void QSensorGestureManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::deleteLater();
}

void QSensorGestureManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGestureManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensorGestureManager_Event(void* ptr, void* e)
{
	return static_cast<QSensorGestureManager*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensorGestureManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::event(static_cast<QEvent*>(e));
}

int QSensorGestureManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGestureManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensorGestureManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensorGestureManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGestureManager*>(ptr)->metaObject());
}

void* QSensorGestureManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::metaObject());
}

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QString name() const { return QString(callbackQSensorGesturePluginInterface_Name(const_cast<MyQSensorGesturePluginInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QStringList supportedIds() const { return QString(callbackQSensorGesturePluginInterface_SupportedIds(const_cast<MyQSensorGesturePluginInterface*>(this), this->objectNameAbs().toUtf8().data())).split("|", QString::SkipEmptyParts); };
};

char* QSensorGesturePluginInterface_Name(void* ptr)
{
	return static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8().data();
}

char* QSensorGesturePluginInterface_SupportedIds(void* ptr)
{
	return static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join("|").toUtf8().data();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr)
{
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

char* QSensorGesturePluginInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSensorGesturePluginInterface*>(static_cast<QSensorGesturePluginInterface*>(ptr))) {
		return static_cast<MyQSensorGesturePluginInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorGesturePluginInterface_BASE").toUtf8().data();
}

void QSensorGesturePluginInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSensorGesturePluginInterface*>(static_cast<QSensorGesturePluginInterface*>(ptr))) {
		static_cast<MyQSensorGesturePluginInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer
{
public:
	MyQSensorGestureRecognizer(QObject *parent) : QSensorGestureRecognizer(parent) {};
	void create() { callbackQSensorGestureRecognizer_Create(this, this->objectName().toUtf8().data()); };
	void Signal_Detected(const QString & gestureId) { callbackQSensorGestureRecognizer_Detected(this, this->objectName().toUtf8().data(), gestureId.toUtf8().data()); };
	QString id() const { return QString(callbackQSensorGestureRecognizer_Id(const_cast<MyQSensorGestureRecognizer*>(this), this->objectName().toUtf8().data())); };
	bool isActive() { return callbackQSensorGestureRecognizer_IsActive(this, this->objectName().toUtf8().data()) != 0; };
	bool start() { return callbackQSensorGestureRecognizer_Start(this, this->objectName().toUtf8().data()) != 0; };
	bool stop() { return callbackQSensorGestureRecognizer_Stop(this, this->objectName().toUtf8().data()) != 0; };
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureRecognizer_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGestureRecognizer_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSensorGestureRecognizer_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSensorGestureRecognizer_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSensorGestureRecognizer_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSensorGestureRecognizer_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSensorGestureRecognizer_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSensorGestureRecognizer_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSensorGestureRecognizer_MetaObject(const_cast<MyQSensorGestureRecognizer*>(this), this->objectName().toUtf8().data())); };
};

void* QSensorGestureRecognizer_NewQSensorGestureRecognizer(void* parent)
{
	return new MyQSensorGestureRecognizer(static_cast<QObject*>(parent));
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

void QSensorGestureRecognizer_Detected(void* ptr, char* gestureId)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->detected(QString(gestureId));
}

char* QSensorGestureRecognizer_GestureSignals(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGestureRecognizer_Id(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->id().toUtf8().data();
}

int QSensorGestureRecognizer_IsActive(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->isActive();
}

int QSensorGestureRecognizer_Start(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->start();
}

void QSensorGestureRecognizer_StartBackend(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->startBackend();
}

int QSensorGestureRecognizer_Stop(void* ptr)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->stop();
}

void QSensorGestureRecognizer_StopBackend(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->stopBackend();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->~QSensorGestureRecognizer();
}

void QSensorGestureRecognizer_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureRecognizer_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureRecognizer_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureRecognizer_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureRecognizer_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureRecognizer_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureRecognizer_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureRecognizer_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureRecognizer_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensorGestureRecognizer*>(ptr), "deleteLater");
}

void QSensorGestureRecognizer_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::deleteLater();
}

void QSensorGestureRecognizer_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorGestureRecognizer_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensorGestureRecognizer_Event(void* ptr, void* e)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensorGestureRecognizer_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::event(static_cast<QEvent*>(e));
}

int QSensorGestureRecognizer_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensorGestureRecognizer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensorGestureRecognizer_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGestureRecognizer*>(ptr)->metaObject());
}

void* QSensorGestureRecognizer_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::metaObject());
}

void* QSensorManager_QSensorManager_CreateBackend(void* sensor)
{
	return QSensorManager::createBackend(static_cast<QSensor*>(sensor));
}

int QSensorManager_QSensorManager_IsBackendRegistered(char* ty, char* identifier)
{
	return QSensorManager::isBackendRegistered(QByteArray(ty), QByteArray(identifier));
}

void QSensorManager_QSensorManager_RegisterBackend(char* ty, char* identifier, void* factory)
{
	QSensorManager::registerBackend(QByteArray(ty), QByteArray(identifier), static_cast<QSensorBackendFactory*>(factory));
}

void QSensorManager_QSensorManager_SetDefaultBackend(char* ty, char* identifier)
{
	QSensorManager::setDefaultBackend(QByteArray(ty), QByteArray(identifier));
}

void QSensorManager_QSensorManager_UnregisterBackend(char* ty, char* identifier)
{
	QSensorManager::unregisterBackend(QByteArray(ty), QByteArray(identifier));
}

class MyQSensorPluginInterface: public QSensorPluginInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void registerSensors() { callbackQSensorPluginInterface_RegisterSensors(this, this->objectNameAbs().toUtf8().data()); };
};

void QSensorPluginInterface_RegisterSensors(void* ptr)
{
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

char* QSensorPluginInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSensorPluginInterface*>(static_cast<QSensorPluginInterface*>(ptr))) {
		return static_cast<MyQSensorPluginInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorPluginInterface_BASE").toUtf8().data();
}

void QSensorPluginInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSensorPluginInterface*>(static_cast<QSensorPluginInterface*>(ptr))) {
		static_cast<MyQSensorPluginInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSensorReading_Value(void* ptr, int index)
{
	return new QVariant(static_cast<QSensorReading*>(ptr)->value(index));
}

int QSensorReading_ValueCount(void* ptr)
{
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

void QSensorReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::customEvent(static_cast<QEvent*>(event));
}

void QSensorReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSensorReading*>(ptr), "deleteLater");
}

void QSensorReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::deleteLater();
}

void QSensorReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSensorReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSensorReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSensorReading*>(ptr)->QSensorReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSensorReading_Event(void* ptr, void* e)
{
	return static_cast<QSensorReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSensorReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QSensorReading*>(ptr)->QSensorReading::event(static_cast<QEvent*>(e));
}

int QSensorReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSensorReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSensorReading*>(ptr)->QSensorReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSensorReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorReading*>(ptr)->metaObject());
}

void* QSensorReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSensorReading*>(ptr)->QSensorReading::metaObject());
}

class MyQTapFilter: public QTapFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QTapReading * reading) { return callbackQTapFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QTapFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

char* QTapFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQTapFilter*>(static_cast<QTapFilter*>(ptr))) {
		return static_cast<MyQTapFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTapFilter_BASE").toUtf8().data();
}

void QTapFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQTapFilter*>(static_cast<QTapFilter*>(ptr))) {
		static_cast<MyQTapFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QTapReading_IsDoubleTap(void* ptr)
{
	return static_cast<QTapReading*>(ptr)->isDoubleTap();
}

int QTapReading_TapDirection(void* ptr)
{
	return static_cast<QTapReading*>(ptr)->tapDirection();
}

void QTapReading_SetDoubleTap(void* ptr, int doubleTap)
{
	static_cast<QTapReading*>(ptr)->setDoubleTap(doubleTap != 0);
}

void QTapReading_SetTapDirection(void* ptr, int tapDirection)
{
	static_cast<QTapReading*>(ptr)->setTapDirection(static_cast<QTapReading::TapDirection>(tapDirection));
}

void QTapReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->QTapReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTapReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->QTapReading::childEvent(static_cast<QChildEvent*>(event));
}

void QTapReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTapReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTapReading*>(ptr)->QTapReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTapReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTapReading*>(ptr)->QTapReading::customEvent(static_cast<QEvent*>(event));
}

void QTapReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTapReading*>(ptr), "deleteLater");
}

void QTapReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QTapReading*>(ptr)->QTapReading::deleteLater();
}

void QTapReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTapReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTapReading*>(ptr)->QTapReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTapReading_Event(void* ptr, void* e)
{
	return static_cast<QTapReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTapReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QTapReading*>(ptr)->QTapReading::event(static_cast<QEvent*>(e));
}

int QTapReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTapReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTapReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTapReading*>(ptr)->QTapReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTapReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTapReading*>(ptr)->metaObject());
}

void* QTapReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTapReading*>(ptr)->QTapReading::metaObject());
}

class MyQTapSensor: public QTapSensor
{
public:
	MyQTapSensor(QObject *parent) : QTapSensor(parent) {};
	void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents) { callbackQTapSensor_ReturnDoubleTapEventsChanged(this, this->objectName().toUtf8().data(), returnDoubleTapEvents); };
	bool start() { return callbackQTapSensor_Start(this, this->objectName().toUtf8().data()) != 0; };
	void stop() { callbackQTapSensor_Stop(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQTapSensor_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTapSensor_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTapSensor_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQTapSensor_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQTapSensor_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTapSensor_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQTapSensor_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTapSensor_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTapSensor_MetaObject(const_cast<MyQTapSensor*>(this), this->objectName().toUtf8().data())); };
};

void* QTapSensor_Reading(void* ptr)
{
	return static_cast<QTapSensor*>(ptr)->reading();
}

int QTapSensor_ReturnDoubleTapEvents(void* ptr)
{
	return static_cast<QTapSensor*>(ptr)->returnDoubleTapEvents();
}

void QTapSensor_SetReturnDoubleTapEvents(void* ptr, int returnDoubleTapEvents)
{
	static_cast<QTapSensor*>(ptr)->setReturnDoubleTapEvents(returnDoubleTapEvents != 0);
}

void* QTapSensor_NewQTapSensor(void* parent)
{
	return new MyQTapSensor(static_cast<QObject*>(parent));
}

void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr)
{
	QObject::connect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));
}

void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));
}

void QTapSensor_ReturnDoubleTapEventsChanged(void* ptr, int returnDoubleTapEvents)
{
	static_cast<QTapSensor*>(ptr)->returnDoubleTapEventsChanged(returnDoubleTapEvents != 0);
}

void QTapSensor_DestroyQTapSensor(void* ptr)
{
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
}

char* QTapSensor_QTapSensor_Type()
{
	return QString(QTapSensor::type).toUtf8().data();
}

int QTapSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QTapSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QTapSensor_StartDefault(void* ptr)
{
	return static_cast<QTapSensor*>(ptr)->QTapSensor::start();
}

void QTapSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTapSensor*>(ptr), "stop");
}

void QTapSensor_StopDefault(void* ptr)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::stop();
}

void QTapSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTapSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QTapSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTapSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTapSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::customEvent(static_cast<QEvent*>(event));
}

void QTapSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTapSensor*>(ptr), "deleteLater");
}

void QTapSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::deleteLater();
}

void QTapSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTapSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTapSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTapSensor*>(ptr)->QTapSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTapSensor_Event(void* ptr, void* e)
{
	return static_cast<QTapSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTapSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QTapSensor*>(ptr)->QTapSensor::event(static_cast<QEvent*>(e));
}

int QTapSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTapSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTapSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTapSensor*>(ptr)->QTapSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTapSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTapSensor*>(ptr)->metaObject());
}

void* QTapSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTapSensor*>(ptr)->QTapSensor::metaObject());
}

class MyQTiltFilter: public QTiltFilter
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	bool filter(QTiltReading * reading) { return callbackQTiltFilter_Filter(this, this->objectNameAbs().toUtf8().data(), reading) != 0; };
};

int QTiltFilter_Filter(void* ptr, void* reading)
{
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

char* QTiltFilter_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQTiltFilter*>(static_cast<QTiltFilter*>(ptr))) {
		return static_cast<MyQTiltFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTiltFilter_BASE").toUtf8().data();
}

void QTiltFilter_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQTiltFilter*>(static_cast<QTiltFilter*>(ptr))) {
		static_cast<MyQTiltFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

double QTiltReading_XRotation(void* ptr)
{
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->xRotation());
}

double QTiltReading_YRotation(void* ptr)
{
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->yRotation());
}

void QTiltReading_SetXRotation(void* ptr, double x)
{
	static_cast<QTiltReading*>(ptr)->setXRotation(static_cast<double>(x));
}

void QTiltReading_SetYRotation(void* ptr, double y)
{
	static_cast<QTiltReading*>(ptr)->setYRotation(static_cast<double>(y));
}

void QTiltReading_TimerEvent(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltReading_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltReading_ChildEvent(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTiltReading_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::childEvent(static_cast<QChildEvent*>(event));
}

void QTiltReading_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTiltReading*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltReading_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltReading_CustomEvent(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTiltReading_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::customEvent(static_cast<QEvent*>(event));
}

void QTiltReading_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTiltReading*>(ptr), "deleteLater");
}

void QTiltReading_DeleteLaterDefault(void* ptr)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::deleteLater();
}

void QTiltReading_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTiltReading*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltReading_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTiltReading*>(ptr)->QTiltReading::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTiltReading_Event(void* ptr, void* e)
{
	return static_cast<QTiltReading*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTiltReading_EventDefault(void* ptr, void* e)
{
	return static_cast<QTiltReading*>(ptr)->QTiltReading::event(static_cast<QEvent*>(e));
}

int QTiltReading_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTiltReading*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTiltReading_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTiltReading*>(ptr)->QTiltReading::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTiltReading_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTiltReading*>(ptr)->metaObject());
}

void* QTiltReading_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTiltReading*>(ptr)->QTiltReading::metaObject());
}

void* QTiltSensor_NewQTiltSensor(void* parent)
{
	return new QTiltSensor(static_cast<QObject*>(parent));
}

void* QTiltSensor_Reading(void* ptr)
{
	return static_cast<QTiltSensor*>(ptr)->reading();
}

void QTiltSensor_DestroyQTiltSensor(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->~QTiltSensor();
}

void QTiltSensor_Calibrate(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->calibrate();
}

char* QTiltSensor_QTiltSensor_Type()
{
	return QString(QTiltSensor::type).toUtf8().data();
}

int QTiltSensor_Start(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QTiltSensor*>(ptr), "start", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QTiltSensor_StartDefault(void* ptr)
{
	return static_cast<QTiltSensor*>(ptr)->QTiltSensor::start();
}

void QTiltSensor_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTiltSensor*>(ptr), "stop");
}

void QTiltSensor_StopDefault(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::stop();
}

void QTiltSensor_TimerEvent(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltSensor_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltSensor_ChildEvent(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTiltSensor_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QTiltSensor_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTiltSensor*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltSensor_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltSensor_CustomEvent(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTiltSensor_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::customEvent(static_cast<QEvent*>(event));
}

void QTiltSensor_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTiltSensor*>(ptr), "deleteLater");
}

void QTiltSensor_DeleteLaterDefault(void* ptr)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::deleteLater();
}

void QTiltSensor_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTiltSensor*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTiltSensor_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTiltSensor_Event(void* ptr, void* e)
{
	return static_cast<QTiltSensor*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTiltSensor_EventDefault(void* ptr, void* e)
{
	return static_cast<QTiltSensor*>(ptr)->QTiltSensor::event(static_cast<QEvent*>(e));
}

int QTiltSensor_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTiltSensor*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTiltSensor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTiltSensor*>(ptr)->QTiltSensor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTiltSensor_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTiltSensor*>(ptr)->metaObject());
}

void* QTiltSensor_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTiltSensor*>(ptr)->QTiltSensor::metaObject());
}

