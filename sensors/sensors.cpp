#define protected public

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

class MyQAccelerometer: public QAccelerometer {
public:
	MyQAccelerometer(QObject *parent) : QAccelerometer(parent) {};
	void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode) { callbackQAccelerometerAccelerationModeChanged(this, this->objectName().toUtf8().data(), accelerationMode); };
	void timerEvent(QTimerEvent * event) { callbackQAccelerometerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAccelerometerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAccelerometerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QAccelerometer_AccelerationMode(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

void* QAccelerometer_Reading(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->reading();
}

void* QAccelerometer_NewQAccelerometer(void* parent){
	return new MyQAccelerometer(static_cast<QObject*>(parent));
}

void QAccelerometer_ConnectAccelerationModeChanged(void* ptr){
	QObject::connect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_DisconnectAccelerationModeChanged(void* ptr){
	QObject::disconnect(static_cast<QAccelerometer*>(ptr), static_cast<void (QAccelerometer::*)(QAccelerometer::AccelerationMode)>(&QAccelerometer::accelerationModeChanged), static_cast<MyQAccelerometer*>(ptr), static_cast<void (MyQAccelerometer::*)(QAccelerometer::AccelerationMode)>(&MyQAccelerometer::Signal_AccelerationModeChanged));;
}

void QAccelerometer_AccelerationModeChanged(void* ptr, int accelerationMode){
	static_cast<QAccelerometer*>(ptr)->accelerationModeChanged(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode){
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_DestroyQAccelerometer(void* ptr){
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
}

void QAccelerometer_TimerEvent(void* ptr, void* event){
	static_cast<MyQAccelerometer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometer_TimerEventDefault(void* ptr, void* event){
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometer_ChildEvent(void* ptr, void* event){
	static_cast<MyQAccelerometer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometer_ChildEventDefault(void* ptr, void* event){
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometer_CustomEvent(void* ptr, void* event){
	static_cast<MyQAccelerometer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAccelerometer_CustomEventDefault(void* ptr, void* event){
	static_cast<QAccelerometer*>(ptr)->QAccelerometer::customEvent(static_cast<QEvent*>(event));
}

int QAccelerometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

double QAccelerometerReading_X(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->x());
}

double QAccelerometerReading_Y(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->y());
}

double QAccelerometerReading_Z(void* ptr){
	return static_cast<double>(static_cast<QAccelerometerReading*>(ptr)->z());
}

void QAccelerometerReading_SetX(void* ptr, double x){
	static_cast<QAccelerometerReading*>(ptr)->setX(static_cast<double>(x));
}

void QAccelerometerReading_SetY(void* ptr, double y){
	static_cast<QAccelerometerReading*>(ptr)->setY(static_cast<double>(y));
}

void QAccelerometerReading_SetZ(void* ptr, double z){
	static_cast<QAccelerometerReading*>(ptr)->setZ(static_cast<double>(z));
}

void QAccelerometerReading_TimerEvent(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometerReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccelerometerReading_ChildEvent(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometerReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAccelerometerReading_CustomEvent(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAccelerometerReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QAccelerometerReading*>(ptr)->QAccelerometerReading::customEvent(static_cast<QEvent*>(event));
}

void* QAltimeter_Reading(void* ptr){
	return static_cast<QAltimeter*>(ptr)->reading();
}

void* QAltimeter_NewQAltimeter(void* parent){
	return new QAltimeter(static_cast<QObject*>(parent));
}

void QAltimeter_DestroyQAltimeter(void* ptr){
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

void QAltimeter_TimerEvent(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeter_TimerEventDefault(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->QAltimeter::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeter_ChildEvent(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeter_ChildEventDefault(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->QAltimeter::childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeter_CustomEvent(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAltimeter_CustomEventDefault(void* ptr, void* event){
	static_cast<QAltimeter*>(ptr)->QAltimeter::customEvent(static_cast<QEvent*>(event));
}

int QAltimeterFilter_Filter(void* ptr, void* reading){
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

double QAltimeterReading_Altitude(void* ptr){
	return static_cast<double>(static_cast<QAltimeterReading*>(ptr)->altitude());
}

void QAltimeterReading_SetAltitude(void* ptr, double altitude){
	static_cast<QAltimeterReading*>(ptr)->setAltitude(static_cast<double>(altitude));
}

void QAltimeterReading_TimerEvent(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeterReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAltimeterReading_ChildEvent(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeterReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAltimeterReading_CustomEvent(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAltimeterReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QAltimeterReading*>(ptr)->QAltimeterReading::customEvent(static_cast<QEvent*>(event));
}

int QAmbientLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

int QAmbientLightReading_LightLevel(void* ptr){
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

void QAmbientLightReading_SetLightLevel(void* ptr, int lightLevel){
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

void QAmbientLightReading_TimerEvent(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightReading_ChildEvent(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightReading_CustomEvent(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightReading*>(ptr)->QAmbientLightReading::customEvent(static_cast<QEvent*>(event));
}

class MyQAmbientLightSensor: public QAmbientLightSensor {
public:
	MyQAmbientLightSensor(QObject *parent) : QAmbientLightSensor(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQAmbientLightSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAmbientLightSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQAmbientLightSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QAmbientLightSensor_Reading(void* ptr){
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent){
	return new MyQAmbientLightSensor(static_cast<QObject*>(parent));
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr){
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

void QAmbientLightSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQAmbientLightSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientLightSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQAmbientLightSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientLightSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQAmbientLightSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientLightSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QAmbientLightSensor*>(ptr)->QAmbientLightSensor::customEvent(static_cast<QEvent*>(event));
}

int QAmbientTemperatureFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

double QAmbientTemperatureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QAmbientTemperatureReading*>(ptr)->temperature());
}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(static_cast<double>(temperature));
}

void QAmbientTemperatureReading_TimerEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureReading_ChildEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureReading_CustomEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureReading*>(ptr)->QAmbientTemperatureReading::customEvent(static_cast<QEvent*>(event));
}

void* QAmbientTemperatureSensor_Reading(void* ptr){
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent){
	return new QAmbientTemperatureSensor(static_cast<QObject*>(parent));
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr){
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

void QAmbientTemperatureSensor_TimerEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAmbientTemperatureSensor_ChildEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QAmbientTemperatureSensor_CustomEvent(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAmbientTemperatureSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QAmbientTemperatureSensor*>(ptr)->QAmbientTemperatureSensor::customEvent(static_cast<QEvent*>(event));
}

class MyQCompass: public QCompass {
public:
	MyQCompass(QObject *parent) : QCompass(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQCompassTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQCompassChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQCompassCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QCompass_Reading(void* ptr){
	return static_cast<QCompass*>(ptr)->reading();
}

void* QCompass_NewQCompass(void* parent){
	return new MyQCompass(static_cast<QObject*>(parent));
}

void QCompass_DestroyQCompass(void* ptr){
	static_cast<QCompass*>(ptr)->~QCompass();
}

void QCompass_TimerEvent(void* ptr, void* event){
	static_cast<MyQCompass*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompass_TimerEventDefault(void* ptr, void* event){
	static_cast<QCompass*>(ptr)->QCompass::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompass_ChildEvent(void* ptr, void* event){
	static_cast<MyQCompass*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCompass_ChildEventDefault(void* ptr, void* event){
	static_cast<QCompass*>(ptr)->QCompass::childEvent(static_cast<QChildEvent*>(event));
}

void QCompass_CustomEvent(void* ptr, void* event){
	static_cast<MyQCompass*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCompass_CustomEventDefault(void* ptr, void* event){
	static_cast<QCompass*>(ptr)->QCompass::customEvent(static_cast<QEvent*>(event));
}

int QCompassFilter_Filter(void* ptr, void* reading){
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

double QCompassReading_Azimuth(void* ptr){
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->azimuth());
}

double QCompassReading_CalibrationLevel(void* ptr){
	return static_cast<double>(static_cast<QCompassReading*>(ptr)->calibrationLevel());
}

void QCompassReading_SetAzimuth(void* ptr, double azimuth){
	static_cast<QCompassReading*>(ptr)->setAzimuth(static_cast<double>(azimuth));
}

void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel){
	static_cast<QCompassReading*>(ptr)->setCalibrationLevel(static_cast<double>(calibrationLevel));
}

void QCompassReading_TimerEvent(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompassReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->QCompassReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QCompassReading_ChildEvent(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QCompassReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->QCompassReading::childEvent(static_cast<QChildEvent*>(event));
}

void QCompassReading_CustomEvent(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QCompassReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QCompassReading*>(ptr)->QCompassReading::customEvent(static_cast<QEvent*>(event));
}

int QDistanceFilter_Filter(void* ptr, void* reading){
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

double QDistanceReading_Distance(void* ptr){
	return static_cast<double>(static_cast<QDistanceReading*>(ptr)->distance());
}

void QDistanceReading_SetDistance(void* ptr, double distance){
	static_cast<QDistanceReading*>(ptr)->setDistance(static_cast<double>(distance));
}

void QDistanceReading_TimerEvent(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceReading_ChildEvent(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceReading_CustomEvent(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDistanceReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QDistanceReading*>(ptr)->QDistanceReading::customEvent(static_cast<QEvent*>(event));
}

void* QDistanceSensor_Reading(void* ptr){
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

void* QDistanceSensor_NewQDistanceSensor(void* parent){
	return new QDistanceSensor(static_cast<QObject*>(parent));
}

void QDistanceSensor_DestroyQDistanceSensor(void* ptr){
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

void QDistanceSensor_TimerEvent(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDistanceSensor_ChildEvent(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QDistanceSensor_CustomEvent(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDistanceSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QDistanceSensor*>(ptr)->QDistanceSensor::customEvent(static_cast<QEvent*>(event));
}

class MyQGyroscope: public QGyroscope {
public:
	MyQGyroscope(QObject *parent) : QGyroscope(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQGyroscopeTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGyroscopeChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGyroscopeCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QGyroscope_Reading(void* ptr){
	return static_cast<QGyroscope*>(ptr)->reading();
}

void* QGyroscope_NewQGyroscope(void* parent){
	return new MyQGyroscope(static_cast<QObject*>(parent));
}

void QGyroscope_DestroyQGyroscope(void* ptr){
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

void QGyroscope_TimerEvent(void* ptr, void* event){
	static_cast<MyQGyroscope*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscope_TimerEventDefault(void* ptr, void* event){
	static_cast<QGyroscope*>(ptr)->QGyroscope::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscope_ChildEvent(void* ptr, void* event){
	static_cast<MyQGyroscope*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscope_ChildEventDefault(void* ptr, void* event){
	static_cast<QGyroscope*>(ptr)->QGyroscope::childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscope_CustomEvent(void* ptr, void* event){
	static_cast<MyQGyroscope*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGyroscope_CustomEventDefault(void* ptr, void* event){
	static_cast<QGyroscope*>(ptr)->QGyroscope::customEvent(static_cast<QEvent*>(event));
}

int QGyroscopeFilter_Filter(void* ptr, void* reading){
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

double QGyroscopeReading_X(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->x());
}

double QGyroscopeReading_Y(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->y());
}

double QGyroscopeReading_Z(void* ptr){
	return static_cast<double>(static_cast<QGyroscopeReading*>(ptr)->z());
}

void QGyroscopeReading_SetX(void* ptr, double x){
	static_cast<QGyroscopeReading*>(ptr)->setX(static_cast<double>(x));
}

void QGyroscopeReading_SetY(void* ptr, double y){
	static_cast<QGyroscopeReading*>(ptr)->setY(static_cast<double>(y));
}

void QGyroscopeReading_SetZ(void* ptr, double z){
	static_cast<QGyroscopeReading*>(ptr)->setZ(static_cast<double>(z));
}

void QGyroscopeReading_TimerEvent(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscopeReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGyroscopeReading_ChildEvent(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscopeReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::childEvent(static_cast<QChildEvent*>(event));
}

void QGyroscopeReading_CustomEvent(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGyroscopeReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QGyroscopeReading*>(ptr)->QGyroscopeReading::customEvent(static_cast<QEvent*>(event));
}

int QHolsterFilter_Filter(void* ptr, void* reading){
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

int QHolsterReading_Holstered(void* ptr){
	return static_cast<QHolsterReading*>(ptr)->holstered();
}

void QHolsterReading_SetHolstered(void* ptr, int holstered){
	static_cast<QHolsterReading*>(ptr)->setHolstered(holstered != 0);
}

void QHolsterReading_TimerEvent(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterReading_ChildEvent(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterReading_CustomEvent(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHolsterReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QHolsterReading*>(ptr)->QHolsterReading::customEvent(static_cast<QEvent*>(event));
}

void* QHolsterSensor_Reading(void* ptr){
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

void* QHolsterSensor_NewQHolsterSensor(void* parent){
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr){
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

void QHolsterSensor_TimerEvent(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHolsterSensor_ChildEvent(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QHolsterSensor_CustomEvent(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHolsterSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QHolsterSensor*>(ptr)->QHolsterSensor::customEvent(static_cast<QEvent*>(event));
}

int QIRProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

double QIRProximityReading_Reflectance(void* ptr){
	return static_cast<double>(static_cast<QIRProximityReading*>(ptr)->reflectance());
}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance){
	static_cast<QIRProximityReading*>(ptr)->setReflectance(static_cast<double>(reflectance));
}

void QIRProximityReading_TimerEvent(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximityReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximityReading_ChildEvent(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximityReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximityReading_CustomEvent(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIRProximityReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QIRProximityReading*>(ptr)->QIRProximityReading::customEvent(static_cast<QEvent*>(event));
}

class MyQIRProximitySensor: public QIRProximitySensor {
public:
	MyQIRProximitySensor(QObject *parent) : QIRProximitySensor(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQIRProximitySensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIRProximitySensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQIRProximitySensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QIRProximitySensor_Reading(void* ptr){
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

void* QIRProximitySensor_NewQIRProximitySensor(void* parent){
	return new MyQIRProximitySensor(static_cast<QObject*>(parent));
}

void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr){
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

void QIRProximitySensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQIRProximitySensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximitySensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIRProximitySensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQIRProximitySensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximitySensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::childEvent(static_cast<QChildEvent*>(event));
}

void QIRProximitySensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQIRProximitySensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIRProximitySensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QIRProximitySensor*>(ptr)->QIRProximitySensor::customEvent(static_cast<QEvent*>(event));
}

int QLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

double QLightReading_Lux(void* ptr){
	return static_cast<double>(static_cast<QLightReading*>(ptr)->lux());
}

void QLightReading_SetLux(void* ptr, double lux){
	static_cast<QLightReading*>(ptr)->setLux(static_cast<double>(lux));
}

void QLightReading_TimerEvent(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->QLightReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightReading_ChildEvent(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLightReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->QLightReading::childEvent(static_cast<QChildEvent*>(event));
}

void QLightReading_CustomEvent(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLightReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QLightReading*>(ptr)->QLightReading::customEvent(static_cast<QEvent*>(event));
}

class MyQLightSensor: public QLightSensor {
public:
	MyQLightSensor(QObject *parent) : QLightSensor(parent) {};
	void Signal_FieldOfViewChanged(qreal fieldOfView) { callbackQLightSensorFieldOfViewChanged(this, this->objectName().toUtf8().data(), static_cast<double>(fieldOfView)); };
	void timerEvent(QTimerEvent * event) { callbackQLightSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQLightSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQLightSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

double QLightSensor_FieldOfView(void* ptr){
	return static_cast<double>(static_cast<QLightSensor*>(ptr)->fieldOfView());
}

void* QLightSensor_Reading(void* ptr){
	return static_cast<QLightSensor*>(ptr)->reading();
}

void* QLightSensor_NewQLightSensor(void* parent){
	return new MyQLightSensor(static_cast<QObject*>(parent));
}

void QLightSensor_ConnectFieldOfViewChanged(void* ptr){
	QObject::connect(static_cast<QLightSensor*>(ptr), static_cast<void (QLightSensor::*)(qreal)>(&QLightSensor::fieldOfViewChanged), static_cast<MyQLightSensor*>(ptr), static_cast<void (MyQLightSensor::*)(qreal)>(&MyQLightSensor::Signal_FieldOfViewChanged));;
}

void QLightSensor_DisconnectFieldOfViewChanged(void* ptr){
	QObject::disconnect(static_cast<QLightSensor*>(ptr), static_cast<void (QLightSensor::*)(qreal)>(&QLightSensor::fieldOfViewChanged), static_cast<MyQLightSensor*>(ptr), static_cast<void (MyQLightSensor::*)(qreal)>(&MyQLightSensor::Signal_FieldOfViewChanged));;
}

void QLightSensor_FieldOfViewChanged(void* ptr, double fieldOfView){
	static_cast<QLightSensor*>(ptr)->fieldOfViewChanged(static_cast<double>(fieldOfView));
}

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView){
	static_cast<QLightSensor*>(ptr)->setFieldOfView(static_cast<double>(fieldOfView));
}

void QLightSensor_DestroyQLightSensor(void* ptr){
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

void QLightSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQLightSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QLightSensor*>(ptr)->QLightSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QLightSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQLightSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QLightSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QLightSensor*>(ptr)->QLightSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QLightSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQLightSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QLightSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QLightSensor*>(ptr)->QLightSensor::customEvent(static_cast<QEvent*>(event));
}

class MyQMagnetometer: public QMagnetometer {
public:
	MyQMagnetometer(QObject *parent) : QMagnetometer(parent) {};
	void Signal_ReturnGeoValuesChanged(bool returnGeoValues) { callbackQMagnetometerReturnGeoValuesChanged(this, this->objectName().toUtf8().data(), returnGeoValues); };
	void timerEvent(QTimerEvent * event) { callbackQMagnetometerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQMagnetometerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQMagnetometerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QMagnetometer_Reading(void* ptr){
	return static_cast<QMagnetometer*>(ptr)->reading();
}

int QMagnetometer_ReturnGeoValues(void* ptr){
	return static_cast<QMagnetometer*>(ptr)->returnGeoValues();
}

void QMagnetometer_SetReturnGeoValues(void* ptr, int returnGeoValues){
	static_cast<QMagnetometer*>(ptr)->setReturnGeoValues(returnGeoValues != 0);
}

void* QMagnetometer_NewQMagnetometer(void* parent){
	return new MyQMagnetometer(static_cast<QObject*>(parent));
}

void QMagnetometer_ConnectReturnGeoValuesChanged(void* ptr){
	QObject::connect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_DisconnectReturnGeoValuesChanged(void* ptr){
	QObject::disconnect(static_cast<QMagnetometer*>(ptr), static_cast<void (QMagnetometer::*)(bool)>(&QMagnetometer::returnGeoValuesChanged), static_cast<MyQMagnetometer*>(ptr), static_cast<void (MyQMagnetometer::*)(bool)>(&MyQMagnetometer::Signal_ReturnGeoValuesChanged));;
}

void QMagnetometer_ReturnGeoValuesChanged(void* ptr, int returnGeoValues){
	static_cast<QMagnetometer*>(ptr)->returnGeoValuesChanged(returnGeoValues != 0);
}

void QMagnetometer_DestroyQMagnetometer(void* ptr){
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
}

void QMagnetometer_TimerEvent(void* ptr, void* event){
	static_cast<MyQMagnetometer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometer_TimerEventDefault(void* ptr, void* event){
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometer_ChildEvent(void* ptr, void* event){
	static_cast<MyQMagnetometer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometer_ChildEventDefault(void* ptr, void* event){
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometer_CustomEvent(void* ptr, void* event){
	static_cast<MyQMagnetometer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMagnetometer_CustomEventDefault(void* ptr, void* event){
	static_cast<QMagnetometer*>(ptr)->QMagnetometer::customEvent(static_cast<QEvent*>(event));
}

int QMagnetometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

double QMagnetometerReading_CalibrationLevel(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->calibrationLevel());
}

double QMagnetometerReading_X(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->x());
}

double QMagnetometerReading_Y(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->y());
}

double QMagnetometerReading_Z(void* ptr){
	return static_cast<double>(static_cast<QMagnetometerReading*>(ptr)->z());
}

void QMagnetometerReading_SetCalibrationLevel(void* ptr, double calibrationLevel){
	static_cast<QMagnetometerReading*>(ptr)->setCalibrationLevel(static_cast<double>(calibrationLevel));
}

void QMagnetometerReading_SetX(void* ptr, double x){
	static_cast<QMagnetometerReading*>(ptr)->setX(static_cast<double>(x));
}

void QMagnetometerReading_SetY(void* ptr, double y){
	static_cast<QMagnetometerReading*>(ptr)->setY(static_cast<double>(y));
}

void QMagnetometerReading_SetZ(void* ptr, double z){
	static_cast<QMagnetometerReading*>(ptr)->setZ(static_cast<double>(z));
}

void QMagnetometerReading_TimerEvent(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometerReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMagnetometerReading_ChildEvent(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometerReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::childEvent(static_cast<QChildEvent*>(event));
}

void QMagnetometerReading_CustomEvent(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMagnetometerReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QMagnetometerReading*>(ptr)->QMagnetometerReading::customEvent(static_cast<QEvent*>(event));
}

int QOrientationFilter_Filter(void* ptr, void* reading){
	return static_cast<QOrientationFilter*>(ptr)->filter(static_cast<QOrientationReading*>(reading));
}

int QOrientationReading_Orientation(void* ptr){
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

void QOrientationReading_SetOrientation(void* ptr, int orientation){
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

void QOrientationReading_TimerEvent(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationReading_ChildEvent(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationReading_CustomEvent(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QOrientationReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QOrientationReading*>(ptr)->QOrientationReading::customEvent(static_cast<QEvent*>(event));
}

class MyQOrientationSensor: public QOrientationSensor {
public:
	MyQOrientationSensor(QObject *parent) : QOrientationSensor(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQOrientationSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQOrientationSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQOrientationSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QOrientationSensor_Reading(void* ptr){
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

void* QOrientationSensor_NewQOrientationSensor(void* parent){
	return new MyQOrientationSensor(static_cast<QObject*>(parent));
}

void QOrientationSensor_DestroyQOrientationSensor(void* ptr){
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

void QOrientationSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQOrientationSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QOrientationSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQOrientationSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QOrientationSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQOrientationSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QOrientationSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QOrientationSensor*>(ptr)->QOrientationSensor::customEvent(static_cast<QEvent*>(event));
}

int QPressureFilter_Filter(void* ptr, void* reading){
	return static_cast<QPressureFilter*>(ptr)->filter(static_cast<QPressureReading*>(reading));
}

double QPressureReading_Pressure(void* ptr){
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->pressure());
}

double QPressureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QPressureReading*>(ptr)->temperature());
}

void QPressureReading_SetPressure(void* ptr, double pressure){
	static_cast<QPressureReading*>(ptr)->setPressure(static_cast<double>(pressure));
}

void QPressureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QPressureReading*>(ptr)->setTemperature(static_cast<double>(temperature));
}

void QPressureReading_TimerEvent(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->QPressureReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureReading_ChildEvent(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPressureReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->QPressureReading::childEvent(static_cast<QChildEvent*>(event));
}

void QPressureReading_CustomEvent(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPressureReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QPressureReading*>(ptr)->QPressureReading::customEvent(static_cast<QEvent*>(event));
}

void* QPressureSensor_Reading(void* ptr){
	return static_cast<QPressureSensor*>(ptr)->reading();
}

void* QPressureSensor_NewQPressureSensor(void* parent){
	return new QPressureSensor(static_cast<QObject*>(parent));
}

void QPressureSensor_DestroyQPressureSensor(void* ptr){
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

void QPressureSensor_TimerEvent(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPressureSensor_ChildEvent(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPressureSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QPressureSensor_CustomEvent(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPressureSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QPressureSensor*>(ptr)->QPressureSensor::customEvent(static_cast<QEvent*>(event));
}

int QProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

int QProximityReading_Close(void* ptr){
	return static_cast<QProximityReading*>(ptr)->close();
}

void QProximityReading_SetClose(void* ptr, int close){
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

void QProximityReading_TimerEvent(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximityReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->QProximityReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximityReading_ChildEvent(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QProximityReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->QProximityReading::childEvent(static_cast<QChildEvent*>(event));
}

void QProximityReading_CustomEvent(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QProximityReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QProximityReading*>(ptr)->QProximityReading::customEvent(static_cast<QEvent*>(event));
}

class MyQProximitySensor: public QProximitySensor {
public:
	MyQProximitySensor(QObject *parent) : QProximitySensor(parent) {};
	void timerEvent(QTimerEvent * event) { callbackQProximitySensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQProximitySensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQProximitySensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QProximitySensor_Reading(void* ptr){
	return static_cast<QProximitySensor*>(ptr)->reading();
}

void* QProximitySensor_NewQProximitySensor(void* parent){
	return new MyQProximitySensor(static_cast<QObject*>(parent));
}

void QProximitySensor_DestroyQProximitySensor(void* ptr){
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

void QProximitySensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQProximitySensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximitySensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QProximitySensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQProximitySensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QProximitySensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::childEvent(static_cast<QChildEvent*>(event));
}

void QProximitySensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQProximitySensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QProximitySensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QProximitySensor*>(ptr)->QProximitySensor::customEvent(static_cast<QEvent*>(event));
}

int QRotationFilter_Filter(void* ptr, void* reading){
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

double QRotationReading_X(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->x());
}

double QRotationReading_Y(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->y());
}

double QRotationReading_Z(void* ptr){
	return static_cast<double>(static_cast<QRotationReading*>(ptr)->z());
}

void QRotationReading_SetFromEuler(void* ptr, double x, double y, double z){
	static_cast<QRotationReading*>(ptr)->setFromEuler(static_cast<double>(x), static_cast<double>(y), static_cast<double>(z));
}

void QRotationReading_TimerEvent(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->QRotationReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationReading_ChildEvent(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRotationReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->QRotationReading::childEvent(static_cast<QChildEvent*>(event));
}

void QRotationReading_CustomEvent(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRotationReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QRotationReading*>(ptr)->QRotationReading::customEvent(static_cast<QEvent*>(event));
}

class MyQRotationSensor: public QRotationSensor {
public:
	MyQRotationSensor(QObject *parent) : QRotationSensor(parent) {};
	void Signal_HasZChanged(bool hasZ) { callbackQRotationSensorHasZChanged(this, this->objectName().toUtf8().data(), hasZ); };
	void timerEvent(QTimerEvent * event) { callbackQRotationSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQRotationSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQRotationSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QRotationSensor_HasZ(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

void* QRotationSensor_Reading(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->reading();
}

void* QRotationSensor_NewQRotationSensor(void* parent){
	return new MyQRotationSensor(static_cast<QObject*>(parent));
}

void QRotationSensor_ConnectHasZChanged(void* ptr){
	QObject::connect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_DisconnectHasZChanged(void* ptr){
	QObject::disconnect(static_cast<QRotationSensor*>(ptr), static_cast<void (QRotationSensor::*)(bool)>(&QRotationSensor::hasZChanged), static_cast<MyQRotationSensor*>(ptr), static_cast<void (MyQRotationSensor::*)(bool)>(&MyQRotationSensor::Signal_HasZChanged));;
}

void QRotationSensor_HasZChanged(void* ptr, int hasZ){
	static_cast<QRotationSensor*>(ptr)->hasZChanged(hasZ != 0);
}

void QRotationSensor_SetHasZ(void* ptr, int hasZ){
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(void* ptr){
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

void QRotationSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQRotationSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRotationSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQRotationSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRotationSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QRotationSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQRotationSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRotationSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QRotationSensor*>(ptr)->QRotationSensor::customEvent(static_cast<QEvent*>(event));
}

class MyQSensor: public QSensor {
public:
	MyQSensor(const QByteArray &type, QObject *parent) : QSensor(type, parent) {};
	void Signal_ActiveChanged() { callbackQSensorActiveChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AlwaysOnChanged() { callbackQSensorAlwaysOnChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AvailableSensorsChanged() { callbackQSensorAvailableSensorsChanged(this, this->objectName().toUtf8().data()); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensorAxesOrientationModeChanged(this, this->objectName().toUtf8().data(), axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensorBufferSizeChanged(this, this->objectName().toUtf8().data(), bufferSize); };
	void Signal_BusyChanged() { callbackQSensorBusyChanged(this, this->objectName().toUtf8().data()); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensorCurrentOrientationChanged(this, this->objectName().toUtf8().data(), currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensorDataRateChanged(this, this->objectName().toUtf8().data()); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensorEfficientBufferSizeChanged(this, this->objectName().toUtf8().data(), efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensorMaxBufferSizeChanged(this, this->objectName().toUtf8().data(), maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensorReadingChanged(this, this->objectName().toUtf8().data()); };
	void Signal_SensorError(int error) { callbackQSensorSensorError(this, this->objectName().toUtf8().data(), error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensorSkipDuplicatesChanged(this, this->objectName().toUtf8().data(), skipDuplicates); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensorUserOrientationChanged(this, this->objectName().toUtf8().data(), userOrientation); };
	void timerEvent(QTimerEvent * event) { callbackQSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QSensor_AxesOrientationMode(void* ptr){
	return static_cast<QSensor*>(ptr)->axesOrientationMode();
}

int QSensor_BufferSize(void* ptr){
	return static_cast<QSensor*>(ptr)->bufferSize();
}

int QSensor_CurrentOrientation(void* ptr){
	return static_cast<QSensor*>(ptr)->currentOrientation();
}

int QSensor_DataRate(void* ptr){
	return static_cast<QSensor*>(ptr)->dataRate();
}

char* QSensor_Description(void* ptr){
	return static_cast<QSensor*>(ptr)->description().toUtf8().data();
}

int QSensor_EfficientBufferSize(void* ptr){
	return static_cast<QSensor*>(ptr)->efficientBufferSize();
}

int QSensor_Error(void* ptr){
	return static_cast<QSensor*>(ptr)->error();
}

void* QSensor_Identifier(void* ptr){
	return new QByteArray(static_cast<QSensor*>(ptr)->identifier());
}

int QSensor_IsActive(void* ptr){
	return static_cast<QSensor*>(ptr)->isActive();
}

int QSensor_IsAlwaysOn(void* ptr){
	return static_cast<QSensor*>(ptr)->isAlwaysOn();
}

int QSensor_IsBusy(void* ptr){
	return static_cast<QSensor*>(ptr)->isBusy();
}

int QSensor_IsConnectedToBackend(void* ptr){
	return static_cast<QSensor*>(ptr)->isConnectedToBackend();
}

int QSensor_MaxBufferSize(void* ptr){
	return static_cast<QSensor*>(ptr)->maxBufferSize();
}

int QSensor_OutputRange(void* ptr){
	return static_cast<QSensor*>(ptr)->outputRange();
}

void* QSensor_Reading(void* ptr){
	return static_cast<QSensor*>(ptr)->reading();
}

void QSensor_SetActive(void* ptr, int active){
	static_cast<QSensor*>(ptr)->setActive(active != 0);
}

void QSensor_SetAlwaysOn(void* ptr, int alwaysOn){
	static_cast<QSensor*>(ptr)->setAlwaysOn(alwaysOn != 0);
}

void QSensor_SetAxesOrientationMode(void* ptr, int axesOrientationMode){
	static_cast<QSensor*>(ptr)->setAxesOrientationMode(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_SetBufferSize(void* ptr, int bufferSize){
	static_cast<QSensor*>(ptr)->setBufferSize(bufferSize);
}

void QSensor_SetDataRate(void* ptr, int rate){
	static_cast<QSensor*>(ptr)->setDataRate(rate);
}

void QSensor_SetIdentifier(void* ptr, void* identifier){
	static_cast<QSensor*>(ptr)->setIdentifier(*static_cast<QByteArray*>(identifier));
}

void QSensor_SetOutputRange(void* ptr, int index){
	static_cast<QSensor*>(ptr)->setOutputRange(index);
}

void QSensor_SetUserOrientation(void* ptr, int userOrientation){
	static_cast<QSensor*>(ptr)->setUserOrientation(userOrientation);
}

int QSensor_SkipDuplicates(void* ptr){
	return static_cast<QSensor*>(ptr)->skipDuplicates();
}

void* QSensor_Type(void* ptr){
	return new QByteArray(static_cast<QSensor*>(ptr)->type());
}

int QSensor_UserOrientation(void* ptr){
	return static_cast<QSensor*>(ptr)->userOrientation();
}

void* QSensor_NewQSensor(void* ty, void* parent){
	return new MyQSensor(*static_cast<QByteArray*>(ty), static_cast<QObject*>(parent));
}

void QSensor_ConnectActiveChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));;
}

void QSensor_DisconnectActiveChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));;
}

void QSensor_ActiveChanged(void* ptr){
	static_cast<QSensor*>(ptr)->activeChanged();
}

void QSensor_AddFilter(void* ptr, void* filter){
	static_cast<QSensor*>(ptr)->addFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectAlwaysOnChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_DisconnectAlwaysOnChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_AlwaysOnChanged(void* ptr){
	static_cast<QSensor*>(ptr)->alwaysOnChanged();
}

void QSensor_ConnectAvailableSensorsChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_DisconnectAvailableSensorsChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_AvailableSensorsChanged(void* ptr){
	static_cast<QSensor*>(ptr)->availableSensorsChanged();
}

void QSensor_ConnectAxesOrientationModeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_DisconnectAxesOrientationModeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_AxesOrientationModeChanged(void* ptr, int axesOrientationMode){
	static_cast<QSensor*>(ptr)->axesOrientationModeChanged(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_ConnectBufferSizeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_DisconnectBufferSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_BufferSizeChanged(void* ptr, int bufferSize){
	static_cast<QSensor*>(ptr)->bufferSizeChanged(bufferSize);
}

void QSensor_ConnectBusyChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
}

void QSensor_DisconnectBusyChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
}

void QSensor_BusyChanged(void* ptr){
	static_cast<QSensor*>(ptr)->busyChanged();
}

int QSensor_ConnectToBackend(void* ptr){
	return static_cast<QSensor*>(ptr)->connectToBackend();
}

void QSensor_ConnectCurrentOrientationChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));;
}

void QSensor_DisconnectCurrentOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));;
}

void QSensor_CurrentOrientationChanged(void* ptr, int currentOrientation){
	static_cast<QSensor*>(ptr)->currentOrientationChanged(currentOrientation);
}

void QSensor_ConnectDataRateChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
}

void QSensor_DisconnectDataRateChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
}

void QSensor_DataRateChanged(void* ptr){
	static_cast<QSensor*>(ptr)->dataRateChanged();
}

void* QSensor_QSensor_DefaultSensorForType(void* ty){
	return new QByteArray(QSensor::defaultSensorForType(*static_cast<QByteArray*>(ty)));
}

void QSensor_ConnectEfficientBufferSizeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));;
}

void QSensor_DisconnectEfficientBufferSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));;
}

void QSensor_EfficientBufferSizeChanged(void* ptr, int efficientBufferSize){
	static_cast<QSensor*>(ptr)->efficientBufferSizeChanged(efficientBufferSize);
}

int QSensor_IsFeatureSupported(void* ptr, int feature){
	return static_cast<QSensor*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensor_ConnectMaxBufferSizeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_DisconnectMaxBufferSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_MaxBufferSizeChanged(void* ptr, int maxBufferSize){
	static_cast<QSensor*>(ptr)->maxBufferSizeChanged(maxBufferSize);
}

void QSensor_ConnectReadingChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
}

void QSensor_DisconnectReadingChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
}

void QSensor_ReadingChanged(void* ptr){
	static_cast<QSensor*>(ptr)->readingChanged();
}

void QSensor_RemoveFilter(void* ptr, void* filter){
	static_cast<QSensor*>(ptr)->removeFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectSensorError(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));;
}

void QSensor_DisconnectSensorError(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));;
}

void QSensor_SensorError(void* ptr, int error){
	static_cast<QSensor*>(ptr)->sensorError(error);
}

void QSensor_SetCurrentOrientation(void* ptr, int currentOrientation){
	static_cast<QSensor*>(ptr)->setCurrentOrientation(currentOrientation);
}

void QSensor_SetEfficientBufferSize(void* ptr, int efficientBufferSize){
	static_cast<QSensor*>(ptr)->setEfficientBufferSize(efficientBufferSize);
}

void QSensor_SetMaxBufferSize(void* ptr, int maxBufferSize){
	static_cast<QSensor*>(ptr)->setMaxBufferSize(maxBufferSize);
}

void QSensor_SetSkipDuplicates(void* ptr, int skipDuplicates){
	static_cast<QSensor*>(ptr)->setSkipDuplicates(skipDuplicates != 0);
}

void QSensor_ConnectSkipDuplicatesChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));;
}

void QSensor_DisconnectSkipDuplicatesChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));;
}

void QSensor_SkipDuplicatesChanged(void* ptr, int skipDuplicates){
	static_cast<QSensor*>(ptr)->skipDuplicatesChanged(skipDuplicates != 0);
}

int QSensor_Start(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "start");
}

void QSensor_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "stop");
}

void QSensor_ConnectUserOrientationChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));;
}

void QSensor_DisconnectUserOrientationChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));;
}

void QSensor_UserOrientationChanged(void* ptr, int userOrientation){
	static_cast<QSensor*>(ptr)->userOrientationChanged(userOrientation);
}

void QSensor_DestroyQSensor(void* ptr){
	static_cast<QSensor*>(ptr)->~QSensor();
}

void QSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensor*>(ptr)->QSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensor*>(ptr)->QSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensor*>(ptr)->QSensor::customEvent(static_cast<QEvent*>(event));
}

class MyQSensorBackend: public QSensorBackend {
public:
	void timerEvent(QTimerEvent * event) { callbackQSensorBackendTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorBackendChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSensorBackendCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSensorBackend_AddDataRate(void* ptr, double min, double max){
	static_cast<QSensorBackend*>(ptr)->addDataRate(static_cast<double>(min), static_cast<double>(max));
}

int QSensorBackend_IsFeatureSupported(void* ptr, int feature){
	return static_cast<QSensorBackend*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensorBackend_SensorBusy(void* ptr){
	static_cast<QSensorBackend*>(ptr)->sensorBusy();
}

void QSensorBackend_SensorError(void* ptr, int error){
	static_cast<QSensorBackend*>(ptr)->sensorError(error);
}

void QSensorBackend_AddOutputRange(void* ptr, double min, double max, double accuracy){
	static_cast<QSensorBackend*>(ptr)->addOutputRange(static_cast<double>(min), static_cast<double>(max), static_cast<double>(accuracy));
}

void QSensorBackend_NewReadingAvailable(void* ptr){
	static_cast<QSensorBackend*>(ptr)->newReadingAvailable();
}

void* QSensorBackend_Reading(void* ptr){
	return static_cast<QSensorBackend*>(ptr)->reading();
}

void* QSensorBackend_Sensor(void* ptr){
	return static_cast<QSensorBackend*>(ptr)->sensor();
}

void QSensorBackend_SensorStopped(void* ptr){
	static_cast<QSensorBackend*>(ptr)->sensorStopped();
}

void QSensorBackend_SetDataRates(void* ptr, void* otherSensor){
	static_cast<QSensorBackend*>(ptr)->setDataRates(static_cast<QSensor*>(otherSensor));
}

void QSensorBackend_SetDescription(void* ptr, char* description){
	static_cast<QSensorBackend*>(ptr)->setDescription(QString(description));
}

void QSensorBackend_Start(void* ptr){
	static_cast<QSensorBackend*>(ptr)->start();
}

void QSensorBackend_Stop(void* ptr){
	static_cast<QSensorBackend*>(ptr)->stop();
}

void QSensorBackend_TimerEvent(void* ptr, void* event){
	static_cast<MyQSensorBackend*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorBackend_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorBackend_ChildEvent(void* ptr, void* event){
	static_cast<MyQSensorBackend*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorBackend_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorBackend_CustomEvent(void* ptr, void* event){
	static_cast<MyQSensorBackend*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorBackend_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensorBackend*>(ptr)->QSensorBackend::customEvent(static_cast<QEvent*>(event));
}

void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor){
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

void QSensorChangesInterface_SensorsChanged(void* ptr){
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

class MyQSensorFilter: public QSensorFilter {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

int QSensorFilter_Filter(void* ptr, void* reading){
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

void QSensorFilter_DestroyQSensorFilter(void* ptr){
	static_cast<QSensorFilter*>(ptr)->~QSensorFilter();
}

char* QSensorFilter_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSensorFilter*>(static_cast<QSensorFilter*>(ptr))) {
		return static_cast<MyQSensorFilter*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorFilter_BASE").toUtf8().data();
}

void QSensorFilter_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSensorFilter*>(static_cast<QSensorFilter*>(ptr))) {
		static_cast<MyQSensorFilter*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorGesture: public QSensorGesture {
public:
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGestureChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSensorGestureCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSensorGesture_NewQSensorGesture(char* ids, void* parent){
	return new QSensorGesture(QString(ids).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

char* QSensorGesture_GestureSignals(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGesture_InvalidIds(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->invalidIds().join("|").toUtf8().data();
}

int QSensorGesture_IsActive(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->isActive();
}

void QSensorGesture_StartDetection(void* ptr){
	static_cast<QSensorGesture*>(ptr)->startDetection();
}

void QSensorGesture_StopDetection(void* ptr){
	static_cast<QSensorGesture*>(ptr)->stopDetection();
}

char* QSensorGesture_ValidIds(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->validIds().join("|").toUtf8().data();
}

void QSensorGesture_DestroyQSensorGesture(void* ptr){
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

void QSensorGesture_TimerEvent(void* ptr, void* event){
	static_cast<MyQSensorGesture*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGesture_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGesture_ChildEvent(void* ptr, void* event){
	static_cast<MyQSensorGesture*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGesture_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGesture_CustomEvent(void* ptr, void* event){
	static_cast<MyQSensorGesture*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGesture_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensorGesture*>(ptr)->QSensorGesture::customEvent(static_cast<QEvent*>(event));
}

class MyQSensorGestureManager: public QSensorGestureManager {
public:
	void Signal_NewSensorGestureAvailable() { callbackQSensorGestureManagerNewSensorGestureAvailable(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureManagerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGestureManagerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSensorGestureManagerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QSensorGestureManager_NewQSensorGestureManager(void* parent){
	return new QSensorGestureManager(static_cast<QObject*>(parent));
}

char* QSensorGestureManager_GestureIds(void* ptr){
	return static_cast<QSensorGestureManager*>(ptr)->gestureIds().join("|").toUtf8().data();
}

void QSensorGestureManager_ConnectNewSensorGestureAvailable(void* ptr){
	QObject::connect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

void QSensorGestureManager_DisconnectNewSensorGestureAvailable(void* ptr){
	QObject::disconnect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

void QSensorGestureManager_NewSensorGestureAvailable(void* ptr){
	static_cast<QSensorGestureManager*>(ptr)->newSensorGestureAvailable();
}

char* QSensorGestureManager_RecognizerSignals(void* ptr, char* gestureId){
	return static_cast<QSensorGestureManager*>(ptr)->recognizerSignals(QString(gestureId)).join("|").toUtf8().data();
}

int QSensorGestureManager_RegisterSensorGestureRecognizer(void* ptr, void* recognizer){
	return static_cast<QSensorGestureManager*>(ptr)->registerSensorGestureRecognizer(static_cast<QSensorGestureRecognizer*>(recognizer));
}

void* QSensorGestureManager_QSensorGestureManager_SensorGestureRecognizer(char* id){
	return QSensorGestureManager::sensorGestureRecognizer(QString(id));
}

void QSensorGestureManager_DestroyQSensorGestureManager(void* ptr){
	static_cast<QSensorGestureManager*>(ptr)->~QSensorGestureManager();
}

void QSensorGestureManager_TimerEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureManager_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureManager_ChildEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureManager_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureManager_CustomEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureManager_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureManager*>(ptr)->QSensorGestureManager::customEvent(static_cast<QEvent*>(event));
}

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

char* QSensorGesturePluginInterface_Name(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8().data();
}

char* QSensorGesturePluginInterface_SupportedIds(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join("|").toUtf8().data();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr){
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

char* QSensorGesturePluginInterface_ObjectNameAbs(void* ptr){
	if (dynamic_cast<MyQSensorGesturePluginInterface*>(static_cast<QSensorGesturePluginInterface*>(ptr))) {
		return static_cast<MyQSensorGesturePluginInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSensorGesturePluginInterface_BASE").toUtf8().data();
}

void QSensorGesturePluginInterface_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQSensorGesturePluginInterface*>(static_cast<QSensorGesturePluginInterface*>(ptr))) {
		static_cast<MyQSensorGesturePluginInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer {
public:
	void timerEvent(QTimerEvent * event) { callbackQSensorGestureRecognizerTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSensorGestureRecognizerChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQSensorGestureRecognizerCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QSensorGestureRecognizer_CreateBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->createBackend();
}

char* QSensorGestureRecognizer_GestureSignals(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join("|").toUtf8().data();
}

char* QSensorGestureRecognizer_Id(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->id().toUtf8().data();
}

int QSensorGestureRecognizer_IsActive(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->isActive();
}

void QSensorGestureRecognizer_StartBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->startBackend();
}

void QSensorGestureRecognizer_StopBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->stopBackend();
}

void QSensorGestureRecognizer_DestroyQSensorGestureRecognizer(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->~QSensorGestureRecognizer();
}

void QSensorGestureRecognizer_TimerEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureRecognizer*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureRecognizer_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorGestureRecognizer_ChildEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureRecognizer*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureRecognizer_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorGestureRecognizer_CustomEvent(void* ptr, void* event){
	static_cast<MyQSensorGestureRecognizer*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorGestureRecognizer_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensorGestureRecognizer*>(ptr)->QSensorGestureRecognizer::customEvent(static_cast<QEvent*>(event));
}

void* QSensorManager_QSensorManager_CreateBackend(void* sensor){
	return QSensorManager::createBackend(static_cast<QSensor*>(sensor));
}

int QSensorManager_QSensorManager_IsBackendRegistered(void* ty, void* identifier){
	return QSensorManager::isBackendRegistered(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_RegisterBackend(void* ty, void* identifier, void* factory){
	QSensorManager::registerBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier), static_cast<QSensorBackendFactory*>(factory));
}

void QSensorManager_QSensorManager_SetDefaultBackend(void* ty, void* identifier){
	QSensorManager::setDefaultBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorManager_QSensorManager_UnregisterBackend(void* ty, void* identifier){
	QSensorManager::unregisterBackend(*static_cast<QByteArray*>(ty), *static_cast<QByteArray*>(identifier));
}

void QSensorPluginInterface_RegisterSensors(void* ptr){
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

void* QSensorReading_Value(void* ptr, int index){
	return new QVariant(static_cast<QSensorReading*>(ptr)->value(index));
}

int QSensorReading_ValueCount(void* ptr){
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

void QSensorReading_TimerEvent(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->QSensorReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSensorReading_ChildEvent(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSensorReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->QSensorReading::childEvent(static_cast<QChildEvent*>(event));
}

void QSensorReading_CustomEvent(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSensorReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QSensorReading*>(ptr)->QSensorReading::customEvent(static_cast<QEvent*>(event));
}

int QTapFilter_Filter(void* ptr, void* reading){
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

int QTapReading_IsDoubleTap(void* ptr){
	return static_cast<QTapReading*>(ptr)->isDoubleTap();
}

int QTapReading_TapDirection(void* ptr){
	return static_cast<QTapReading*>(ptr)->tapDirection();
}

void QTapReading_SetDoubleTap(void* ptr, int doubleTap){
	static_cast<QTapReading*>(ptr)->setDoubleTap(doubleTap != 0);
}

void QTapReading_SetTapDirection(void* ptr, int tapDirection){
	static_cast<QTapReading*>(ptr)->setTapDirection(static_cast<QTapReading::TapDirection>(tapDirection));
}

void QTapReading_TimerEvent(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->QTapReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapReading_ChildEvent(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTapReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->QTapReading::childEvent(static_cast<QChildEvent*>(event));
}

void QTapReading_CustomEvent(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTapReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QTapReading*>(ptr)->QTapReading::customEvent(static_cast<QEvent*>(event));
}

class MyQTapSensor: public QTapSensor {
public:
	MyQTapSensor(QObject *parent) : QTapSensor(parent) {};
	void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents) { callbackQTapSensorReturnDoubleTapEventsChanged(this, this->objectName().toUtf8().data(), returnDoubleTapEvents); };
	void timerEvent(QTimerEvent * event) { callbackQTapSensorTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTapSensorChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQTapSensorCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QTapSensor_Reading(void* ptr){
	return static_cast<QTapSensor*>(ptr)->reading();
}

int QTapSensor_ReturnDoubleTapEvents(void* ptr){
	return static_cast<QTapSensor*>(ptr)->returnDoubleTapEvents();
}

void QTapSensor_SetReturnDoubleTapEvents(void* ptr, int returnDoubleTapEvents){
	static_cast<QTapSensor*>(ptr)->setReturnDoubleTapEvents(returnDoubleTapEvents != 0);
}

void* QTapSensor_NewQTapSensor(void* parent){
	return new MyQTapSensor(static_cast<QObject*>(parent));
}

void QTapSensor_ConnectReturnDoubleTapEventsChanged(void* ptr){
	QObject::connect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_DisconnectReturnDoubleTapEventsChanged(void* ptr){
	QObject::disconnect(static_cast<QTapSensor*>(ptr), static_cast<void (QTapSensor::*)(bool)>(&QTapSensor::returnDoubleTapEventsChanged), static_cast<MyQTapSensor*>(ptr), static_cast<void (MyQTapSensor::*)(bool)>(&MyQTapSensor::Signal_ReturnDoubleTapEventsChanged));;
}

void QTapSensor_ReturnDoubleTapEventsChanged(void* ptr, int returnDoubleTapEvents){
	static_cast<QTapSensor*>(ptr)->returnDoubleTapEventsChanged(returnDoubleTapEvents != 0);
}

void QTapSensor_DestroyQTapSensor(void* ptr){
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
}

void QTapSensor_TimerEvent(void* ptr, void* event){
	static_cast<MyQTapSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QTapSensor*>(ptr)->QTapSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTapSensor_ChildEvent(void* ptr, void* event){
	static_cast<MyQTapSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTapSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QTapSensor*>(ptr)->QTapSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QTapSensor_CustomEvent(void* ptr, void* event){
	static_cast<MyQTapSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTapSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QTapSensor*>(ptr)->QTapSensor::customEvent(static_cast<QEvent*>(event));
}

int QTiltFilter_Filter(void* ptr, void* reading){
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

double QTiltReading_XRotation(void* ptr){
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->xRotation());
}

double QTiltReading_YRotation(void* ptr){
	return static_cast<double>(static_cast<QTiltReading*>(ptr)->yRotation());
}

void QTiltReading_SetXRotation(void* ptr, double x){
	static_cast<QTiltReading*>(ptr)->setXRotation(static_cast<double>(x));
}

void QTiltReading_SetYRotation(void* ptr, double y){
	static_cast<QTiltReading*>(ptr)->setYRotation(static_cast<double>(y));
}

void QTiltReading_TimerEvent(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltReading_TimerEventDefault(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->QTiltReading::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltReading_ChildEvent(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTiltReading_ChildEventDefault(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->QTiltReading::childEvent(static_cast<QChildEvent*>(event));
}

void QTiltReading_CustomEvent(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTiltReading_CustomEventDefault(void* ptr, void* event){
	static_cast<QTiltReading*>(ptr)->QTiltReading::customEvent(static_cast<QEvent*>(event));
}

void* QTiltSensor_NewQTiltSensor(void* parent){
	return new QTiltSensor(static_cast<QObject*>(parent));
}

void* QTiltSensor_Reading(void* ptr){
	return static_cast<QTiltSensor*>(ptr)->reading();
}

void QTiltSensor_DestroyQTiltSensor(void* ptr){
	static_cast<QTiltSensor*>(ptr)->~QTiltSensor();
}

void QTiltSensor_Calibrate(void* ptr){
	static_cast<QTiltSensor*>(ptr)->calibrate();
}

void QTiltSensor_TimerEvent(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltSensor_TimerEventDefault(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTiltSensor_ChildEvent(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTiltSensor_ChildEventDefault(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::childEvent(static_cast<QChildEvent*>(event));
}

void QTiltSensor_CustomEvent(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTiltSensor_CustomEventDefault(void* ptr, void* event){
	static_cast<QTiltSensor*>(ptr)->QTiltSensor::customEvent(static_cast<QEvent*>(event));
}

