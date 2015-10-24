#include "qsensor.h"
#include <QString>
#include <QVariant>
#include <QSensorFilter>
#include <QByteArray>
#include <QMetaObject>
#include <QUrl>
#include <QModelIndex>
#include <QSensorReading>
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

int QSensor_AxesOrientationMode(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->axesOrientationMode();
}

int QSensor_BufferSize(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->bufferSize();
}

int QSensor_CurrentOrientation(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->currentOrientation();
}

int QSensor_DataRate(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->dataRate();
}

char* QSensor_Description(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->description().toUtf8().data();
}

int QSensor_EfficientBufferSize(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->efficientBufferSize();
}

int QSensor_Error(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->error();
}

int QSensor_IsActive(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->isActive();
}

int QSensor_IsAlwaysOn(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->isAlwaysOn();
}

int QSensor_IsBusy(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->isBusy();
}

int QSensor_IsConnectedToBackend(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->isConnectedToBackend();
}

int QSensor_MaxBufferSize(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->maxBufferSize();
}

int QSensor_OutputRange(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->outputRange();
}

QtObjectPtr QSensor_Reading(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->reading();
}

void QSensor_SetActive(QtObjectPtr ptr, int active){
	static_cast<QSensor*>(ptr)->setActive(active != 0);
}

void QSensor_SetAlwaysOn(QtObjectPtr ptr, int alwaysOn){
	static_cast<QSensor*>(ptr)->setAlwaysOn(alwaysOn != 0);
}

void QSensor_SetAxesOrientationMode(QtObjectPtr ptr, int axesOrientationMode){
	static_cast<QSensor*>(ptr)->setAxesOrientationMode(static_cast<QSensor::AxesOrientationMode>(axesOrientationMode));
}

void QSensor_SetBufferSize(QtObjectPtr ptr, int bufferSize){
	static_cast<QSensor*>(ptr)->setBufferSize(bufferSize);
}

void QSensor_SetDataRate(QtObjectPtr ptr, int rate){
	static_cast<QSensor*>(ptr)->setDataRate(rate);
}

void QSensor_SetIdentifier(QtObjectPtr ptr, QtObjectPtr identifier){
	static_cast<QSensor*>(ptr)->setIdentifier(*static_cast<QByteArray*>(identifier));
}

void QSensor_SetOutputRange(QtObjectPtr ptr, int index){
	static_cast<QSensor*>(ptr)->setOutputRange(index);
}

void QSensor_SetUserOrientation(QtObjectPtr ptr, int userOrientation){
	static_cast<QSensor*>(ptr)->setUserOrientation(userOrientation);
}

int QSensor_SkipDuplicates(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->skipDuplicates();
}

int QSensor_UserOrientation(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->userOrientation();
}

QtObjectPtr QSensor_NewQSensor(QtObjectPtr ty, QtObjectPtr parent){
	return new QSensor(*static_cast<QByteArray*>(ty), static_cast<QObject*>(parent));
}

void QSensor_ConnectActiveChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));;
}

void QSensor_DisconnectActiveChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::activeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ActiveChanged));;
}

void QSensor_AddFilter(QtObjectPtr ptr, QtObjectPtr filter){
	static_cast<QSensor*>(ptr)->addFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectAlwaysOnChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_DisconnectAlwaysOnChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::alwaysOnChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AlwaysOnChanged));;
}

void QSensor_ConnectAvailableSensorsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_DisconnectAvailableSensorsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::availableSensorsChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_AvailableSensorsChanged));;
}

void QSensor_ConnectAxesOrientationModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_DisconnectAxesOrientationModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(QSensor::AxesOrientationMode)>(&QSensor::axesOrientationModeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(QSensor::AxesOrientationMode)>(&MyQSensor::Signal_AxesOrientationModeChanged));;
}

void QSensor_ConnectBufferSizeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_DisconnectBufferSizeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::bufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_BufferSizeChanged));;
}

void QSensor_ConnectBusyChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
}

void QSensor_DisconnectBusyChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::busyChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_BusyChanged));;
}

int QSensor_ConnectToBackend(QtObjectPtr ptr){
	return static_cast<QSensor*>(ptr)->connectToBackend();
}

void QSensor_ConnectCurrentOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));;
}

void QSensor_DisconnectCurrentOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::currentOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_CurrentOrientationChanged));;
}

void QSensor_ConnectDataRateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
}

void QSensor_DisconnectDataRateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::dataRateChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_DataRateChanged));;
}

void QSensor_ConnectEfficientBufferSizeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));;
}

void QSensor_DisconnectEfficientBufferSizeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::efficientBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_EfficientBufferSizeChanged));;
}

int QSensor_IsFeatureSupported(QtObjectPtr ptr, int feature){
	return static_cast<QSensor*>(ptr)->isFeatureSupported(static_cast<QSensor::Feature>(feature));
}

void QSensor_ConnectMaxBufferSizeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_DisconnectMaxBufferSizeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::maxBufferSizeChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_MaxBufferSizeChanged));;
}

void QSensor_ConnectReadingChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
}

void QSensor_DisconnectReadingChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)()>(&QSensor::readingChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)()>(&MyQSensor::Signal_ReadingChanged));;
}

void QSensor_RemoveFilter(QtObjectPtr ptr, QtObjectPtr filter){
	static_cast<QSensor*>(ptr)->removeFilter(static_cast<QSensorFilter*>(filter));
}

void QSensor_ConnectSensorError(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));;
}

void QSensor_DisconnectSensorError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::sensorError), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_SensorError));;
}

void QSensor_SetCurrentOrientation(QtObjectPtr ptr, int currentOrientation){
	static_cast<QSensor*>(ptr)->setCurrentOrientation(currentOrientation);
}

void QSensor_SetEfficientBufferSize(QtObjectPtr ptr, int efficientBufferSize){
	static_cast<QSensor*>(ptr)->setEfficientBufferSize(efficientBufferSize);
}

void QSensor_SetMaxBufferSize(QtObjectPtr ptr, int maxBufferSize){
	static_cast<QSensor*>(ptr)->setMaxBufferSize(maxBufferSize);
}

void QSensor_SetSkipDuplicates(QtObjectPtr ptr, int skipDuplicates){
	static_cast<QSensor*>(ptr)->setSkipDuplicates(skipDuplicates != 0);
}

void QSensor_ConnectSkipDuplicatesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));;
}

void QSensor_DisconnectSkipDuplicatesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(bool)>(&QSensor::skipDuplicatesChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(bool)>(&MyQSensor::Signal_SkipDuplicatesChanged));;
}

int QSensor_Start(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "start");
}

void QSensor_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSensor*>(ptr), "stop");
}

void QSensor_ConnectUserOrientationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));;
}

void QSensor_DisconnectUserOrientationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSensor*>(ptr), static_cast<void (QSensor::*)(int)>(&QSensor::userOrientationChanged), static_cast<MyQSensor*>(ptr), static_cast<void (MyQSensor::*)(int)>(&MyQSensor::Signal_UserOrientationChanged));;
}

void QSensor_DestroyQSensor(QtObjectPtr ptr){
	static_cast<QSensor*>(ptr)->~QSensor();
}

