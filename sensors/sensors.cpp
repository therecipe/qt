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
#include <QCompass>
#include <QCompassFilter>
#include <QCompassReading>
#include <QDistanceFilter>
#include <QDistanceReading>
#include <QDistanceSensor>
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
#include <QVariant>

class MyQAccelerometer: public QAccelerometer {
public:
	MyQAccelerometer(QObject *parent) : QAccelerometer(parent) {};
	void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode) { callbackQAccelerometerAccelerationModeChanged(this->objectName().toUtf8().data(), accelerationMode); };
protected:
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

void QAccelerometer_SetAccelerationMode(void* ptr, int accelerationMode){
	static_cast<QAccelerometer*>(ptr)->setAccelerationMode(static_cast<QAccelerometer::AccelerationMode>(accelerationMode));
}

void QAccelerometer_DestroyQAccelerometer(void* ptr){
	static_cast<QAccelerometer*>(ptr)->~QAccelerometer();
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
	static_cast<QAccelerometerReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QAccelerometerReading_SetY(void* ptr, double y){
	static_cast<QAccelerometerReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QAccelerometerReading_SetZ(void* ptr, double z){
	static_cast<QAccelerometerReading*>(ptr)->setZ(static_cast<qreal>(z));
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

int QAltimeterFilter_Filter(void* ptr, void* reading){
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

double QAltimeterReading_Altitude(void* ptr){
	return static_cast<double>(static_cast<QAltimeterReading*>(ptr)->altitude());
}

void QAltimeterReading_SetAltitude(void* ptr, double altitude){
	static_cast<QAltimeterReading*>(ptr)->setAltitude(static_cast<qreal>(altitude));
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

class MyQAmbientLightSensor: public QAmbientLightSensor {
public:
	MyQAmbientLightSensor(QObject *parent) : QAmbientLightSensor(parent) {};
protected:
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

int QAmbientTemperatureFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

double QAmbientTemperatureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QAmbientTemperatureReading*>(ptr)->temperature());
}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(static_cast<qreal>(temperature));
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

class MyQCompass: public QCompass {
public:
	MyQCompass(QObject *parent) : QCompass(parent) {};
protected:
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
	static_cast<QCompassReading*>(ptr)->setAzimuth(static_cast<qreal>(azimuth));
}

void QCompassReading_SetCalibrationLevel(void* ptr, double calibrationLevel){
	static_cast<QCompassReading*>(ptr)->setCalibrationLevel(static_cast<qreal>(calibrationLevel));
}

int QDistanceFilter_Filter(void* ptr, void* reading){
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

double QDistanceReading_Distance(void* ptr){
	return static_cast<double>(static_cast<QDistanceReading*>(ptr)->distance());
}

void QDistanceReading_SetDistance(void* ptr, double distance){
	static_cast<QDistanceReading*>(ptr)->setDistance(static_cast<qreal>(distance));
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

class MyQGyroscope: public QGyroscope {
public:
	MyQGyroscope(QObject *parent) : QGyroscope(parent) {};
protected:
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
	static_cast<QGyroscopeReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QGyroscopeReading_SetY(void* ptr, double y){
	static_cast<QGyroscopeReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QGyroscopeReading_SetZ(void* ptr, double z){
	static_cast<QGyroscopeReading*>(ptr)->setZ(static_cast<qreal>(z));
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

void* QHolsterSensor_Reading(void* ptr){
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

void* QHolsterSensor_NewQHolsterSensor(void* parent){
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr){
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

int QIRProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

double QIRProximityReading_Reflectance(void* ptr){
	return static_cast<double>(static_cast<QIRProximityReading*>(ptr)->reflectance());
}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance){
	static_cast<QIRProximityReading*>(ptr)->setReflectance(static_cast<qreal>(reflectance));
}

class MyQIRProximitySensor: public QIRProximitySensor {
public:
	MyQIRProximitySensor(QObject *parent) : QIRProximitySensor(parent) {};
protected:
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

int QLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

double QLightReading_Lux(void* ptr){
	return static_cast<double>(static_cast<QLightReading*>(ptr)->lux());
}

void QLightReading_SetLux(void* ptr, double lux){
	static_cast<QLightReading*>(ptr)->setLux(static_cast<qreal>(lux));
}

class MyQLightSensor: public QLightSensor {
public:
	MyQLightSensor(QObject *parent) : QLightSensor(parent) {};
protected:
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

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView){
	static_cast<QLightSensor*>(ptr)->setFieldOfView(static_cast<qreal>(fieldOfView));
}

void QLightSensor_DestroyQLightSensor(void* ptr){
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

class MyQMagnetometer: public QMagnetometer {
public:
	MyQMagnetometer(QObject *parent) : QMagnetometer(parent) {};
	void Signal_ReturnGeoValuesChanged(bool returnGeoValues) { callbackQMagnetometerReturnGeoValuesChanged(this->objectName().toUtf8().data(), returnGeoValues); };
protected:
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

void QMagnetometer_DestroyQMagnetometer(void* ptr){
	static_cast<QMagnetometer*>(ptr)->~QMagnetometer();
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
	static_cast<QMagnetometerReading*>(ptr)->setCalibrationLevel(static_cast<qreal>(calibrationLevel));
}

void QMagnetometerReading_SetX(void* ptr, double x){
	static_cast<QMagnetometerReading*>(ptr)->setX(static_cast<qreal>(x));
}

void QMagnetometerReading_SetY(void* ptr, double y){
	static_cast<QMagnetometerReading*>(ptr)->setY(static_cast<qreal>(y));
}

void QMagnetometerReading_SetZ(void* ptr, double z){
	static_cast<QMagnetometerReading*>(ptr)->setZ(static_cast<qreal>(z));
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

class MyQOrientationSensor: public QOrientationSensor {
public:
	MyQOrientationSensor(QObject *parent) : QOrientationSensor(parent) {};
protected:
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
	static_cast<QPressureReading*>(ptr)->setPressure(static_cast<qreal>(pressure));
}

void QPressureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QPressureReading*>(ptr)->setTemperature(static_cast<qreal>(temperature));
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

int QProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

int QProximityReading_Close(void* ptr){
	return static_cast<QProximityReading*>(ptr)->close();
}

void QProximityReading_SetClose(void* ptr, int close){
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

class MyQProximitySensor: public QProximitySensor {
public:
	MyQProximitySensor(QObject *parent) : QProximitySensor(parent) {};
protected:
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
	static_cast<QRotationReading*>(ptr)->setFromEuler(static_cast<qreal>(x), static_cast<qreal>(y), static_cast<qreal>(z));
}

class MyQRotationSensor: public QRotationSensor {
public:
	MyQRotationSensor(QObject *parent) : QRotationSensor(parent) {};
	void Signal_HasZChanged(bool hasZ) { callbackQRotationSensorHasZChanged(this->objectName().toUtf8().data(), hasZ); };
protected:
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

void QRotationSensor_SetHasZ(void* ptr, int hasZ){
	static_cast<QRotationSensor*>(ptr)->setHasZ(hasZ != 0);
}

void QRotationSensor_DestroyQRotationSensor(void* ptr){
	static_cast<QRotationSensor*>(ptr)->~QRotationSensor();
}

class MyQSensor: public QSensor {
public:
	MyQSensor(const QByteArray &type, QObject *parent) : QSensor(type, parent) {};
	void Signal_ActiveChanged() { callbackQSensorActiveChanged(this->objectName().toUtf8().data()); };
	void Signal_AlwaysOnChanged() { callbackQSensorAlwaysOnChanged(this->objectName().toUtf8().data()); };
	void Signal_AvailableSensorsChanged() { callbackQSensorAvailableSensorsChanged(this->objectName().toUtf8().data()); };
	void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode) { callbackQSensorAxesOrientationModeChanged(this->objectName().toUtf8().data(), axesOrientationMode); };
	void Signal_BufferSizeChanged(int bufferSize) { callbackQSensorBufferSizeChanged(this->objectName().toUtf8().data(), bufferSize); };
	void Signal_BusyChanged() { callbackQSensorBusyChanged(this->objectName().toUtf8().data()); };
	void Signal_CurrentOrientationChanged(int currentOrientation) { callbackQSensorCurrentOrientationChanged(this->objectName().toUtf8().data(), currentOrientation); };
	void Signal_DataRateChanged() { callbackQSensorDataRateChanged(this->objectName().toUtf8().data()); };
	void Signal_EfficientBufferSizeChanged(int efficientBufferSize) { callbackQSensorEfficientBufferSizeChanged(this->objectName().toUtf8().data(), efficientBufferSize); };
	void Signal_MaxBufferSizeChanged(int maxBufferSize) { callbackQSensorMaxBufferSizeChanged(this->objectName().toUtf8().data(), maxBufferSize); };
	void Signal_ReadingChanged() { callbackQSensorReadingChanged(this->objectName().toUtf8().data()); };
	void Signal_SensorError(int error) { callbackQSensorSensorError(this->objectName().toUtf8().data(), error); };
	void Signal_SkipDuplicatesChanged(bool skipDuplicates) { callbackQSensorSkipDuplicatesChanged(this->objectName().toUtf8().data(), skipDuplicates); };
	void Signal_UserOrientationChanged(int userOrientation) { callbackQSensorUserOrientationChanged(this->objectName().toUtf8().data(), userOrientation); };
protected:
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

void QSensor_AddFilter(void* ptr, void* filter){
	static_cast<QSensor*>(ptr)->addFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectAlwaysOnChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_DisconnectAlwaysOnChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_ConnectAvailableSensorsChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_DisconnectAvailableSensorsChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_ConnectAxesOrientationModeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_DisconnectAxesOrientationModeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_ConnectBufferSizeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_DisconnectBufferSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_ConnectBusyChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
}

void QSensor_DisconnectBusyChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
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

void QSensor_ConnectDataRateChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
}

void QSensor_DisconnectDataRateChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
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

int QSensor_IsFeatureSupported(void* ptr, int feature){
	return static_cast<QSensor*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensor_ConnectMaxBufferSizeChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_DisconnectMaxBufferSizeChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_ConnectReadingChanged(void* ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
}

void QSensor_DisconnectReadingChanged(void* ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
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

void QSensor_DestroyQSensor(void* ptr){
	static_cast<QSensor*>(ptr)->~QSensor();
}

class MyQSensorBackend: public QSensorBackend {
public:
protected:
};

void QSensorBackend_AddDataRate(void* ptr, double min, double max){
	static_cast<QSensorBackend*>(ptr)->addDataRate(static_cast<qreal>(min), static_cast<qreal>(max));
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
	static_cast<QSensorBackend*>(ptr)->addOutputRange(static_cast<qreal>(min), static_cast<qreal>(max), static_cast<qreal>(accuracy));
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
protected:
};

int QSensorFilter_Filter(void* ptr, void* reading){
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

char* QSensorFilter_ObjectNameAbs(void* ptr){
	return static_cast<MyQSensorFilter*>(ptr)->objectNameAbs().toUtf8().data();
}

void QSensorFilter_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQSensorFilter*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQSensorGesture: public QSensorGesture {
public:
protected:
};

void* QSensorGesture_NewQSensorGesture(char* ids, void* parent){
	return new QSensorGesture(QString(ids).split(",,,", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

char* QSensorGesture_GestureSignals(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->gestureSignals().join(",,,").toUtf8().data();
}

char* QSensorGesture_InvalidIds(void* ptr){
	return static_cast<QSensorGesture*>(ptr)->invalidIds().join(",,,").toUtf8().data();
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
	return static_cast<QSensorGesture*>(ptr)->validIds().join(",,,").toUtf8().data();
}

void QSensorGesture_DestroyQSensorGesture(void* ptr){
	static_cast<QSensorGesture*>(ptr)->~QSensorGesture();
}

class MyQSensorGestureManager: public QSensorGestureManager {
public:
	void Signal_NewSensorGestureAvailable() { callbackQSensorGestureManagerNewSensorGestureAvailable(this->objectName().toUtf8().data()); };
protected:
};

void* QSensorGestureManager_NewQSensorGestureManager(void* parent){
	return new QSensorGestureManager(static_cast<QObject*>(parent));
}

char* QSensorGestureManager_GestureIds(void* ptr){
	return static_cast<QSensorGestureManager*>(ptr)->gestureIds().join(",,,").toUtf8().data();
}

void QSensorGestureManager_ConnectNewSensorGestureAvailable(void* ptr){
	QObject::connect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

void QSensorGestureManager_DisconnectNewSensorGestureAvailable(void* ptr){
	QObject::disconnect(static_cast<QSensorGestureManager*>(ptr), static_cast<void (QSensorGestureManager::*)()>(&QSensorGestureManager::newSensorGestureAvailable), static_cast<MyQSensorGestureManager*>(ptr), static_cast<void (MyQSensorGestureManager::*)()>(&MyQSensorGestureManager::Signal_NewSensorGestureAvailable));;
}

char* QSensorGestureManager_RecognizerSignals(void* ptr, char* gestureId){
	return static_cast<QSensorGestureManager*>(ptr)->recognizerSignals(QString(gestureId)).join(",,,").toUtf8().data();
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

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
};

char* QSensorGesturePluginInterface_Name(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->name().toUtf8().data();
}

char* QSensorGesturePluginInterface_SupportedIds(void* ptr){
	return static_cast<QSensorGesturePluginInterface*>(ptr)->supportedIds().join(",,,").toUtf8().data();
}

void QSensorGesturePluginInterface_DestroyQSensorGesturePluginInterface(void* ptr){
	static_cast<QSensorGesturePluginInterface*>(ptr)->~QSensorGesturePluginInterface();
}

char* QSensorGesturePluginInterface_ObjectNameAbs(void* ptr){
	return static_cast<MyQSensorGesturePluginInterface*>(ptr)->objectNameAbs().toUtf8().data();
}

void QSensorGesturePluginInterface_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQSensorGesturePluginInterface*>(ptr)->setObjectNameAbs(QString(name));
}

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer {
public:
protected:
};

void QSensorGestureRecognizer_CreateBackend(void* ptr){
	static_cast<QSensorGestureRecognizer*>(ptr)->createBackend();
}

char* QSensorGestureRecognizer_GestureSignals(void* ptr){
	return static_cast<QSensorGestureRecognizer*>(ptr)->gestureSignals().join(",,,").toUtf8().data();
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

class MyQTapSensor: public QTapSensor {
public:
	MyQTapSensor(QObject *parent) : QTapSensor(parent) {};
	void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents) { callbackQTapSensorReturnDoubleTapEventsChanged(this->objectName().toUtf8().data(), returnDoubleTapEvents); };
protected:
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

void QTapSensor_DestroyQTapSensor(void* ptr){
	static_cast<QTapSensor*>(ptr)->~QTapSensor();
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
	static_cast<QTiltReading*>(ptr)->setXRotation(static_cast<qreal>(x));
}

void QTiltReading_SetYRotation(void* ptr, double y){
	static_cast<QTiltReading*>(ptr)->setYRotation(static_cast<qreal>(y));
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

