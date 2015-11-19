#include "qsensorgestureplugininterface.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorGesture>
#include <QSensor>
#include <QString>
#include <QSensorGesturePluginInterface>
#include "_cgo_export.h"

class MyQSensorGesturePluginInterface: public QSensorGesturePluginInterface {
public:
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

#include "qtapfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTapReading>
#include <QTapFilter>
#include "_cgo_export.h"

class MyQTapFilter: public QTapFilter {
public:
};

int QTapFilter_Filter(void* ptr, void* reading){
	return static_cast<QTapFilter*>(ptr)->filter(static_cast<QTapReading*>(reading));
}

#include "qlightsensor.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QLightSensor>
#include "_cgo_export.h"

class MyQLightSensor: public QLightSensor {
public:
};

double QLightSensor_FieldOfView(void* ptr){
	return static_cast<double>(static_cast<QLightSensor*>(ptr)->fieldOfView());
}

void* QLightSensor_Reading(void* ptr){
	return static_cast<QLightSensor*>(ptr)->reading();
}

void* QLightSensor_NewQLightSensor(void* parent){
	return new QLightSensor(static_cast<QObject*>(parent));
}

void QLightSensor_SetFieldOfView(void* ptr, double fieldOfView){
	static_cast<QLightSensor*>(ptr)->setFieldOfView(static_cast<qreal>(fieldOfView));
}

void QLightSensor_DestroyQLightSensor(void* ptr){
	static_cast<QLightSensor*>(ptr)->~QLightSensor();
}

#include "qambientlightreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightReading>
#include "_cgo_export.h"

class MyQAmbientLightReading: public QAmbientLightReading {
public:
};

int QAmbientLightReading_LightLevel(void* ptr){
	return static_cast<QAmbientLightReading*>(ptr)->lightLevel();
}

void QAmbientLightReading_SetLightLevel(void* ptr, int lightLevel){
	static_cast<QAmbientLightReading*>(ptr)->setLightLevel(static_cast<QAmbientLightReading::LightLevel>(lightLevel));
}

#include "qsensorgesturerecognizer.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorGesture>
#include <QString>
#include <QSensorGestureRecognizer>
#include "_cgo_export.h"

class MyQSensorGestureRecognizer: public QSensorGestureRecognizer {
public:
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

#include "qsensorbackendfactory.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorBackend>
#include <QSensorBackendFactory>
#include "_cgo_export.h"

class MyQSensorBackendFactory: public QSensorBackendFactory {
public:
};

void* QSensorBackendFactory_CreateBackend(void* ptr, void* sensor){
	return static_cast<QSensorBackendFactory*>(ptr)->createBackend(static_cast<QSensor*>(sensor));
}

#include "qdistancefilter.h"
#include <QModelIndex>
#include <QDistanceReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDistanceFilter>
#include "_cgo_export.h"

class MyQDistanceFilter: public QDistanceFilter {
public:
};

int QDistanceFilter_Filter(void* ptr, void* reading){
	return static_cast<QDistanceFilter*>(ptr)->filter(static_cast<QDistanceReading*>(reading));
}

#include "qaltimeter.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAltimeter>
#include "_cgo_export.h"

class MyQAltimeter: public QAltimeter {
public:
};

void* QAltimeter_Reading(void* ptr){
	return static_cast<QAltimeter*>(ptr)->reading();
}

void* QAltimeter_NewQAltimeter(void* parent){
	return new QAltimeter(static_cast<QObject*>(parent));
}

void QAltimeter_DestroyQAltimeter(void* ptr){
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

#include "qsensorchangesinterface.h"
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensorChangesInterface>
#include "_cgo_export.h"

class MyQSensorChangesInterface: public QSensorChangesInterface {
public:
};

void QSensorChangesInterface_SensorsChanged(void* ptr){
	static_cast<QSensorChangesInterface*>(ptr)->sensorsChanged();
}

#include "qpressurereading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPressureReading>
#include "_cgo_export.h"

class MyQPressureReading: public QPressureReading {
public:
};

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

#include "qsensormanager.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QSensor>
#include <QSensorBackendFactory>
#include <QSensorBackend>
#include <QSensorManager>
#include "_cgo_export.h"

class MyQSensorManager: public QSensorManager {
public:
};

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

#include "qcompass.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QCompass>
#include "_cgo_export.h"

class MyQCompass: public QCompass {
public:
};

void* QCompass_Reading(void* ptr){
	return static_cast<QCompass*>(ptr)->reading();
}

void* QCompass_NewQCompass(void* parent){
	return new QCompass(static_cast<QObject*>(parent));
}

void QCompass_DestroyQCompass(void* ptr){
	static_cast<QCompass*>(ptr)->~QCompass();
}

#include "qsensorreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorReading>
#include "_cgo_export.h"

class MyQSensorReading: public QSensorReading {
public:
};

void* QSensorReading_Value(void* ptr, int index){
	return new QVariant(static_cast<QSensorReading*>(ptr)->value(index));
}

int QSensorReading_ValueCount(void* ptr){
	return static_cast<QSensorReading*>(ptr)->valueCount();
}

#include "qaltimeterfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAltimeterReading>
#include <QAltimeter>
#include <QAltimeterFilter>
#include "_cgo_export.h"

class MyQAltimeterFilter: public QAltimeterFilter {
public:
};

int QAltimeterFilter_Filter(void* ptr, void* reading){
	return static_cast<QAltimeterFilter*>(ptr)->filter(static_cast<QAltimeterReading*>(reading));
}

#include "qorientationreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOrientationReading>
#include "_cgo_export.h"

class MyQOrientationReading: public QOrientationReading {
public:
};

int QOrientationReading_Orientation(void* ptr){
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

void QOrientationReading_SetOrientation(void* ptr, int orientation){
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

#include "qsensorbackend.h"
#include <QModelIndex>
#include <QSensor>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorBackend>
#include "_cgo_export.h"

class MyQSensorBackend: public QSensorBackend {
public:
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

#include "qtiltfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTiltReading>
#include <QTiltFilter>
#include "_cgo_export.h"

class MyQTiltFilter: public QTiltFilter {
public:
};

int QTiltFilter_Filter(void* ptr, void* reading){
	return static_cast<QTiltFilter*>(ptr)->filter(static_cast<QTiltReading*>(reading));
}

#include "qcompassreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompass>
#include <QCompassReading>
#include "_cgo_export.h"

class MyQCompassReading: public QCompassReading {
public:
};

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

#include "qmagnetometerreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMagnetometer>
#include <QMagnetometerReading>
#include "_cgo_export.h"

class MyQMagnetometerReading: public QMagnetometerReading {
public:
};

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

#include "qholsterfilter.h"
#include <QHolsterReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHolsterFilter>
#include "_cgo_export.h"

class MyQHolsterFilter: public QHolsterFilter {
public:
};

int QHolsterFilter_Filter(void* ptr, void* reading){
	return static_cast<QHolsterFilter*>(ptr)->filter(static_cast<QHolsterReading*>(reading));
}

#include "qtiltreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTiltReading>
#include "_cgo_export.h"

class MyQTiltReading: public QTiltReading {
public:
};

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

#include "qdistancereading.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QDistanceReading>
#include "_cgo_export.h"

class MyQDistanceReading: public QDistanceReading {
public:
};

double QDistanceReading_Distance(void* ptr){
	return static_cast<double>(static_cast<QDistanceReading*>(ptr)->distance());
}

void QDistanceReading_SetDistance(void* ptr, double distance){
	static_cast<QDistanceReading*>(ptr)->setDistance(static_cast<qreal>(distance));
}

#include "qirproximityreading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QIRProximityReading>
#include "_cgo_export.h"

class MyQIRProximityReading: public QIRProximityReading {
public:
};

double QIRProximityReading_Reflectance(void* ptr){
	return static_cast<double>(static_cast<QIRProximityReading*>(ptr)->reflectance());
}

void QIRProximityReading_SetReflectance(void* ptr, double reflectance){
	static_cast<QIRProximityReading*>(ptr)->setReflectance(static_cast<qreal>(reflectance));
}

#include "qaccelerometerreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometer>
#include <QString>
#include <QVariant>
#include <QAccelerometerReading>
#include "_cgo_export.h"

class MyQAccelerometerReading: public QAccelerometerReading {
public:
};

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

#include "qambientlightsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAmbientLightSensor>
#include "_cgo_export.h"

class MyQAmbientLightSensor: public QAmbientLightSensor {
public:
};

void* QAmbientLightSensor_Reading(void* ptr){
	return static_cast<QAmbientLightSensor*>(ptr)->reading();
}

void* QAmbientLightSensor_NewQAmbientLightSensor(void* parent){
	return new QAmbientLightSensor(static_cast<QObject*>(parent));
}

void QAmbientLightSensor_DestroyQAmbientLightSensor(void* ptr){
	static_cast<QAmbientLightSensor*>(ptr)->~QAmbientLightSensor();
}

#include "qgyroscope.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGyroscope>
#include "_cgo_export.h"

class MyQGyroscope: public QGyroscope {
public:
};

void* QGyroscope_Reading(void* ptr){
	return static_cast<QGyroscope*>(ptr)->reading();
}

void* QGyroscope_NewQGyroscope(void* parent){
	return new QGyroscope(static_cast<QObject*>(parent));
}

void QGyroscope_DestroyQGyroscope(void* ptr){
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

#include "qgyroscopereading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGyroscope>
#include <QGyroscopeReading>
#include "_cgo_export.h"

class MyQGyroscopeReading: public QGyroscopeReading {
public:
};

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

#include "qtapreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QTapReading>
#include "_cgo_export.h"

class MyQTapReading: public QTapReading {
public:
};

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

#include "qproximityfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProximityReading>
#include <QProximityFilter>
#include "_cgo_export.h"

class MyQProximityFilter: public QProximityFilter {
public:
};

int QProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QProximityFilter*>(ptr)->filter(static_cast<QProximityReading*>(reading));
}

#include "qsensorgesturemanager.h"
#include <QModelIndex>
#include <QSensorGesture>
#include <QSensor>
#include <QSensorGestureRecognizer>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensorGestureManager>
#include "_cgo_export.h"

class MyQSensorGestureManager: public QSensorGestureManager {
public:
void Signal_NewSensorGestureAvailable(){callbackQSensorGestureManagerNewSensorGestureAvailable(this->objectName().toUtf8().data());};
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

#include "qirproximitysensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIRProximitySensor>
#include "_cgo_export.h"

class MyQIRProximitySensor: public QIRProximitySensor {
public:
};

void* QIRProximitySensor_Reading(void* ptr){
	return static_cast<QIRProximitySensor*>(ptr)->reading();
}

void* QIRProximitySensor_NewQIRProximitySensor(void* parent){
	return new QIRProximitySensor(static_cast<QObject*>(parent));
}

void QIRProximitySensor_DestroyQIRProximitySensor(void* ptr){
	static_cast<QIRProximitySensor*>(ptr)->~QIRProximitySensor();
}

#include "qholstersensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QHolsterSensor>
#include "_cgo_export.h"

class MyQHolsterSensor: public QHolsterSensor {
public:
};

void* QHolsterSensor_Reading(void* ptr){
	return static_cast<QHolsterSensor*>(ptr)->reading();
}

void* QHolsterSensor_NewQHolsterSensor(void* parent){
	return new QHolsterSensor(static_cast<QObject*>(parent));
}

void QHolsterSensor_DestroyQHolsterSensor(void* ptr){
	static_cast<QHolsterSensor*>(ptr)->~QHolsterSensor();
}

#include "qirproximityfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIRProximityReading>
#include <QIRProximityFilter>
#include "_cgo_export.h"

class MyQIRProximityFilter: public QIRProximityFilter {
public:
};

int QIRProximityFilter_Filter(void* ptr, void* reading){
	return static_cast<QIRProximityFilter*>(ptr)->filter(static_cast<QIRProximityReading*>(reading));
}

#include "qmagnetometerfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMagnetometer>
#include <QMagnetometerReading>
#include <QMagnetometerFilter>
#include "_cgo_export.h"

class MyQMagnetometerFilter: public QMagnetometerFilter {
public:
};

int QMagnetometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QMagnetometerFilter*>(ptr)->filter(static_cast<QMagnetometerReading*>(reading));
}

#include "qpressurefilter.h"
#include <QPressureReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPressureFilter>
#include "_cgo_export.h"

class MyQPressureFilter: public QPressureFilter {
public:
};

int QPressureFilter_Filter(void* ptr, void* reading){
	return static_cast<QPressureFilter*>(ptr)->filter(static_cast<QPressureReading*>(reading));
}

#include "qlightfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QLightReading>
#include <QString>
#include <QVariant>
#include <QLightFilter>
#include "_cgo_export.h"

class MyQLightFilter: public QLightFilter {
public:
};

int QLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QLightFilter*>(ptr)->filter(static_cast<QLightReading*>(reading));
}

#include "qproximitysensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QProximitySensor>
#include "_cgo_export.h"

class MyQProximitySensor: public QProximitySensor {
public:
};

void* QProximitySensor_Reading(void* ptr){
	return static_cast<QProximitySensor*>(ptr)->reading();
}

void* QProximitySensor_NewQProximitySensor(void* parent){
	return new QProximitySensor(static_cast<QObject*>(parent));
}

void QProximitySensor_DestroyQProximitySensor(void* ptr){
	static_cast<QProximitySensor*>(ptr)->~QProximitySensor();
}

#include "qambienttemperaturefilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QAmbientTemperatureReading>
#include <QString>
#include <QVariant>
#include <QAmbientTemperatureFilter>
#include "_cgo_export.h"

class MyQAmbientTemperatureFilter: public QAmbientTemperatureFilter {
public:
};

int QAmbientTemperatureFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientTemperatureFilter*>(ptr)->filter(static_cast<QAmbientTemperatureReading*>(reading));
}

#include "qrotationsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QRotationSensor>
#include "_cgo_export.h"

class MyQRotationSensor: public QRotationSensor {
public:
void Signal_HasZChanged(bool hasZ){callbackQRotationSensorHasZChanged(this->objectName().toUtf8().data(), hasZ);};
};

int QRotationSensor_HasZ(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->hasZ();
}

void* QRotationSensor_Reading(void* ptr){
	return static_cast<QRotationSensor*>(ptr)->reading();
}

void* QRotationSensor_NewQRotationSensor(void* parent){
	return new QRotationSensor(static_cast<QObject*>(parent));
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

#include "qrotationfilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QRotationReading>
#include <QString>
#include <QVariant>
#include <QRotationFilter>
#include "_cgo_export.h"

class MyQRotationFilter: public QRotationFilter {
public:
};

int QRotationFilter_Filter(void* ptr, void* reading){
	return static_cast<QRotationFilter*>(ptr)->filter(static_cast<QRotationReading*>(reading));
}

#include "qpressuresensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPressureSensor>
#include "_cgo_export.h"

class MyQPressureSensor: public QPressureSensor {
public:
};

void* QPressureSensor_Reading(void* ptr){
	return static_cast<QPressureSensor*>(ptr)->reading();
}

void* QPressureSensor_NewQPressureSensor(void* parent){
	return new QPressureSensor(static_cast<QObject*>(parent));
}

void QPressureSensor_DestroyQPressureSensor(void* ptr){
	static_cast<QPressureSensor*>(ptr)->~QPressureSensor();
}

#include "qdistancesensor.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QDistanceSensor>
#include "_cgo_export.h"

class MyQDistanceSensor: public QDistanceSensor {
public:
};

void* QDistanceSensor_Reading(void* ptr){
	return static_cast<QDistanceSensor*>(ptr)->reading();
}

void* QDistanceSensor_NewQDistanceSensor(void* parent){
	return new QDistanceSensor(static_cast<QObject*>(parent));
}

void QDistanceSensor_DestroyQDistanceSensor(void* ptr){
	static_cast<QDistanceSensor*>(ptr)->~QDistanceSensor();
}

#include "qambientlightfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAmbientLightReading>
#include <QAmbientLightFilter>
#include "_cgo_export.h"

class MyQAmbientLightFilter: public QAmbientLightFilter {
public:
};

int QAmbientLightFilter_Filter(void* ptr, void* reading){
	return static_cast<QAmbientLightFilter*>(ptr)->filter(static_cast<QAmbientLightReading*>(reading));
}

#include "qtiltsensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QTiltSensor>
#include "_cgo_export.h"

class MyQTiltSensor: public QTiltSensor {
public:
};

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

#include "qlightreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLightReading>
#include "_cgo_export.h"

class MyQLightReading: public QLightReading {
public:
};

double QLightReading_Lux(void* ptr){
	return static_cast<double>(static_cast<QLightReading*>(ptr)->lux());
}

void QLightReading_SetLux(void* ptr, double lux){
	static_cast<QLightReading*>(ptr)->setLux(static_cast<qreal>(lux));
}

#include "qproximityreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QProximityReading>
#include "_cgo_export.h"

class MyQProximityReading: public QProximityReading {
public:
};

int QProximityReading_Close(void* ptr){
	return static_cast<QProximityReading*>(ptr)->close();
}

void QProximityReading_SetClose(void* ptr, int close){
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

#include "qambienttemperaturesensor.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAmbientTemperatureSensor>
#include "_cgo_export.h"

class MyQAmbientTemperatureSensor: public QAmbientTemperatureSensor {
public:
};

void* QAmbientTemperatureSensor_Reading(void* ptr){
	return static_cast<QAmbientTemperatureSensor*>(ptr)->reading();
}

void* QAmbientTemperatureSensor_NewQAmbientTemperatureSensor(void* parent){
	return new QAmbientTemperatureSensor(static_cast<QObject*>(parent));
}

void QAmbientTemperatureSensor_DestroyQAmbientTemperatureSensor(void* ptr){
	static_cast<QAmbientTemperatureSensor*>(ptr)->~QAmbientTemperatureSensor();
}

#include "qsensorgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QSensor>
#include <QSensorGesture>
#include "_cgo_export.h"

class MyQSensorGesture: public QSensorGesture {
public:
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

#include "qaccelerometerfilter.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAccelerometerReading>
#include <QAccelerometer>
#include <QString>
#include <QAccelerometerFilter>
#include "_cgo_export.h"

class MyQAccelerometerFilter: public QAccelerometerFilter {
public:
};

int QAccelerometerFilter_Filter(void* ptr, void* reading){
	return static_cast<QAccelerometerFilter*>(ptr)->filter(static_cast<QAccelerometerReading*>(reading));
}

#include "qambienttemperaturereading.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAmbientTemperatureReading>
#include "_cgo_export.h"

class MyQAmbientTemperatureReading: public QAmbientTemperatureReading {
public:
};

double QAmbientTemperatureReading_Temperature(void* ptr){
	return static_cast<double>(static_cast<QAmbientTemperatureReading*>(ptr)->temperature());
}

void QAmbientTemperatureReading_SetTemperature(void* ptr, double temperature){
	static_cast<QAmbientTemperatureReading*>(ptr)->setTemperature(static_cast<qreal>(temperature));
}

#include "qorientationsensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOrientationSensor>
#include "_cgo_export.h"

class MyQOrientationSensor: public QOrientationSensor {
public:
};

void* QOrientationSensor_Reading(void* ptr){
	return static_cast<QOrientationSensor*>(ptr)->reading();
}

void* QOrientationSensor_NewQOrientationSensor(void* parent){
	return new QOrientationSensor(static_cast<QObject*>(parent));
}

void QOrientationSensor_DestroyQOrientationSensor(void* ptr){
	static_cast<QOrientationSensor*>(ptr)->~QOrientationSensor();
}

#include "qmagnetometer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMagnetometer>
#include "_cgo_export.h"

class MyQMagnetometer: public QMagnetometer {
public:
void Signal_ReturnGeoValuesChanged(bool returnGeoValues){callbackQMagnetometerReturnGeoValuesChanged(this->objectName().toUtf8().data(), returnGeoValues);};
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
	return new QMagnetometer(static_cast<QObject*>(parent));
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

#include "qrotationreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRotationReading>
#include "_cgo_export.h"

class MyQRotationReading: public QRotationReading {
public:
};

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

#include "qtapsensor.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTapSensor>
#include "_cgo_export.h"

class MyQTapSensor: public QTapSensor {
public:
void Signal_ReturnDoubleTapEventsChanged(bool returnDoubleTapEvents){callbackQTapSensorReturnDoubleTapEventsChanged(this->objectName().toUtf8().data(), returnDoubleTapEvents);};
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
	return new QTapSensor(static_cast<QObject*>(parent));
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

#include "qsensorfilter.h"
#include <QSensorReading>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorFilter>
#include "_cgo_export.h"

class MyQSensorFilter: public QSensorFilter {
public:
};

int QSensorFilter_Filter(void* ptr, void* reading){
	return static_cast<QSensorFilter*>(ptr)->filter(static_cast<QSensorReading*>(reading));
}

#include "qholsterreading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QHolsterReading>
#include "_cgo_export.h"

class MyQHolsterReading: public QHolsterReading {
public:
};

int QHolsterReading_Holstered(void* ptr){
	return static_cast<QHolsterReading*>(ptr)->holstered();
}

void QHolsterReading_SetHolstered(void* ptr, int holstered){
	static_cast<QHolsterReading*>(ptr)->setHolstered(holstered != 0);
}

#include "qgyroscopefilter.h"
#include <QUrl>
#include <QModelIndex>
#include <QGyroscope>
#include <QGyroscopeReading>
#include <QString>
#include <QVariant>
#include <QGyroscopeFilter>
#include "_cgo_export.h"

class MyQGyroscopeFilter: public QGyroscopeFilter {
public:
};

int QGyroscopeFilter_Filter(void* ptr, void* reading){
	return static_cast<QGyroscopeFilter*>(ptr)->filter(static_cast<QGyroscopeReading*>(reading));
}

#include "qsensorplugininterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSensor>
#include <QSensorPluginInterface>
#include "_cgo_export.h"

class MyQSensorPluginInterface: public QSensorPluginInterface {
public:
};

void QSensorPluginInterface_RegisterSensors(void* ptr){
	static_cast<QSensorPluginInterface*>(ptr)->registerSensors();
}

#include "qaltimeterreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QAltimeter>
#include <QString>
#include <QVariant>
#include <QAltimeterReading>
#include "_cgo_export.h"

class MyQAltimeterReading: public QAltimeterReading {
public:
};

double QAltimeterReading_Altitude(void* ptr){
	return static_cast<double>(static_cast<QAltimeterReading*>(ptr)->altitude());
}

void QAltimeterReading_SetAltitude(void* ptr, double altitude){
	static_cast<QAltimeterReading*>(ptr)->setAltitude(static_cast<qreal>(altitude));
}

#include "qcompassfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCompass>
#include <QCompassReading>
#include <QCompassFilter>
#include "_cgo_export.h"

class MyQCompassFilter: public QCompassFilter {
public:
};

int QCompassFilter_Filter(void* ptr, void* reading){
	return static_cast<QCompassFilter*>(ptr)->filter(static_cast<QCompassReading*>(reading));
}

#include "qsensor.h"
#include <QModelIndex>
#include <QByteArray>
#include <QSensorReading>
#include <QMetaObject>
#include <QSensorFilter>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QSensor>
#include "_cgo_export.h"

class MyQSensor: public QSensor {
public:
void Signal_ActiveChanged(){callbackQSensorActiveChanged(this->objectName().toUtf8().data());};
void Signal_AlwaysOnChanged(){callbackQSensorAlwaysOnChanged(this->objectName().toUtf8().data());};
void Signal_AvailableSensorsChanged(){callbackQSensorAvailableSensorsChanged(this->objectName().toUtf8().data());};
void Signal_AxesOrientationModeChanged(QSensor::AxesOrientationMode axesOrientationMode){callbackQSensorAxesOrientationModeChanged(this->objectName().toUtf8().data(), axesOrientationMode);};
void Signal_BufferSizeChanged(int bufferSize){callbackQSensorBufferSizeChanged(this->objectName().toUtf8().data(), bufferSize);};
void Signal_BusyChanged(){callbackQSensorBusyChanged(this->objectName().toUtf8().data());};
void Signal_CurrentOrientationChanged(int currentOrientation){callbackQSensorCurrentOrientationChanged(this->objectName().toUtf8().data(), currentOrientation);};
void Signal_DataRateChanged(){callbackQSensorDataRateChanged(this->objectName().toUtf8().data());};
void Signal_EfficientBufferSizeChanged(int efficientBufferSize){callbackQSensorEfficientBufferSizeChanged(this->objectName().toUtf8().data(), efficientBufferSize);};
void Signal_MaxBufferSizeChanged(int maxBufferSize){callbackQSensorMaxBufferSizeChanged(this->objectName().toUtf8().data(), maxBufferSize);};
void Signal_ReadingChanged(){callbackQSensorReadingChanged(this->objectName().toUtf8().data());};
void Signal_SensorError(int error){callbackQSensorSensorError(this->objectName().toUtf8().data(), error);};
void Signal_SkipDuplicatesChanged(bool skipDuplicates){callbackQSensorSkipDuplicatesChanged(this->objectName().toUtf8().data(), skipDuplicates);};
void Signal_UserOrientationChanged(int userOrientation){callbackQSensorUserOrientationChanged(this->objectName().toUtf8().data(), userOrientation);};
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
	return new QSensor(*static_cast<QByteArray*>(ty), static_cast<QObject*>(parent));
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

#include "qorientationfilter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QOrientationReading>
#include <QOrientationFilter>
#include "_cgo_export.h"

class MyQOrientationFilter: public QOrientationFilter {
public:
};

int QOrientationFilter_Filter(void* ptr, void* reading){
	return static_cast<QOrientationFilter*>(ptr)->filter(static_cast<QOrientationReading*>(reading));
}

#include "qaccelerometer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAccelerometer>
#include "_cgo_export.h"

class MyQAccelerometer: public QAccelerometer {
public:
void Signal_AccelerationModeChanged(QAccelerometer::AccelerationMode accelerationMode){callbackQAccelerometerAccelerationModeChanged(this->objectName().toUtf8().data(), accelerationMode);};
};

int QAccelerometer_AccelerationMode(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->accelerationMode();
}

void* QAccelerometer_Reading(void* ptr){
	return static_cast<QAccelerometer*>(ptr)->reading();
}

void* QAccelerometer_NewQAccelerometer(void* parent){
	return new QAccelerometer(static_cast<QObject*>(parent));
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

