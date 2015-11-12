#include "qsensor.h"
#include <QString>
#include <QModelIndex>
#include <QSensorReading>
#include <QSensorFilter>
#include <QVariant>
#include <QUrl>
#include <QMetaObject>
#include <QByteArray>
#include <QObject>
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

