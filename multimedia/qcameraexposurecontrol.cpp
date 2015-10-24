#include "qcameraexposurecontrol.h"
#include <QCameraExposure>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraExposureControl>
#include "_cgo_export.h"

class MyQCameraExposureControl: public QCameraExposureControl {
public:
void Signal_ActualValueChanged(int parameter){callbackQCameraExposureControlActualValueChanged(this->objectName().toUtf8().data(), parameter);};
void Signal_ParameterRangeChanged(int parameter){callbackQCameraExposureControlParameterRangeChanged(this->objectName().toUtf8().data(), parameter);};
void Signal_RequestedValueChanged(int parameter){callbackQCameraExposureControlRequestedValueChanged(this->objectName().toUtf8().data(), parameter);};
};

char* QCameraExposureControl_ActualValue(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraExposureControl*>(ptr)->actualValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)).toString().toUtf8().data();
}

void QCameraExposureControl_ConnectActualValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

void QCameraExposureControl_DisconnectActualValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::actualValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ActualValueChanged));;
}

int QCameraExposureControl_IsParameterSupported(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraExposureControl*>(ptr)->isParameterSupported(static_cast<QCameraExposureControl::ExposureParameter>(parameter));
}

void QCameraExposureControl_ConnectParameterRangeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

void QCameraExposureControl_DisconnectParameterRangeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::parameterRangeChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_ParameterRangeChanged));;
}

char* QCameraExposureControl_RequestedValue(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraExposureControl*>(ptr)->requestedValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter)).toString().toUtf8().data();
}

void QCameraExposureControl_ConnectRequestedValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

void QCameraExposureControl_DisconnectRequestedValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposureControl*>(ptr), static_cast<void (QCameraExposureControl::*)(int)>(&QCameraExposureControl::requestedValueChanged), static_cast<MyQCameraExposureControl*>(ptr), static_cast<void (MyQCameraExposureControl::*)(int)>(&MyQCameraExposureControl::Signal_RequestedValueChanged));;
}

int QCameraExposureControl_SetValue(QtObjectPtr ptr, int parameter, char* value){
	return static_cast<QCameraExposureControl*>(ptr)->setValue(static_cast<QCameraExposureControl::ExposureParameter>(parameter), QVariant(value));
}

void QCameraExposureControl_DestroyQCameraExposureControl(QtObjectPtr ptr){
	static_cast<QCameraExposureControl*>(ptr)->~QCameraExposureControl();
}

