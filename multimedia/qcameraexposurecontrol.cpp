#include "qcameraexposurecontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QObject>
#include <QCameraExposure>
#include <QString>
#include <QVariant>
#include <QCameraExposureControl>
#include "_cgo_export.h"

class MyQCameraExposureControl: public QCameraExposureControl {
public:
void Signal_ActualValueChanged(int parameter){callbackQCameraExposureControlActualValueChanged(this->objectName().toUtf8().data(), parameter);};
void Signal_ParameterRangeChanged(int parameter){callbackQCameraExposureControlParameterRangeChanged(this->objectName().toUtf8().data(), parameter);};
void Signal_RequestedValueChanged(int parameter){callbackQCameraExposureControlRequestedValueChanged(this->objectName().toUtf8().data(), parameter);};
};

void* QCameraExposureControl_ActualValue(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->actualValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectActualValueChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

void QCameraExposureControl_DisconnectActualValueChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

int QCameraExposureControl_IsParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraExposureControl*>(ptr)->isParameterSupported(static_cast<QCameraExposureControl::ExposureParameter>(parameter));
}

void QCameraExposureControl_ConnectParameterRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

void QCameraExposureControl_DisconnectParameterRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

void* QCameraExposureControl_RequestedValue(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraExposureControl*>(ptr)->requestedValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)));
}

void QCameraExposureControl_ConnectRequestedValueChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

void QCameraExposureControl_DisconnectRequestedValueChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

int QCameraExposureControl_SetValue(void* ptr, int parameter, void* value){
	return static_cast<QCameraExposureControl*>(ptr)->setValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraExposureControl_DestroyQCameraExposureControl(void* ptr){
	static_cast<QCameraExposureControl*>(ptr)->~QCameraExposureControl();
}

