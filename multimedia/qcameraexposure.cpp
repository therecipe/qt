#include "qcameraexposure.h"
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QPoint>
#include <QCamera>
#include <QUrl>
#include <QObject>
#include <QPointF>
#include <QCameraExposure>
#include "_cgo_export.h"

class MyQCameraExposure: public QCameraExposure {
public:
void Signal_ApertureRangeChanged(){callbackQCameraExposureApertureRangeChanged(this->objectName().toUtf8().data());};
void Signal_FlashReady(bool ready){callbackQCameraExposureFlashReady(this->objectName().toUtf8().data(), ready);};
void Signal_IsoSensitivityChanged(int value){callbackQCameraExposureIsoSensitivityChanged(this->objectName().toUtf8().data(), value);};
void Signal_ShutterSpeedRangeChanged(){callbackQCameraExposureShutterSpeedRangeChanged(this->objectName().toUtf8().data());};
};

int QCameraExposure_ExposureMode(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->exposureMode();
}

int QCameraExposure_FlashMode(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->flashMode();
}

int QCameraExposure_IsoSensitivity(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->isoSensitivity();
}

int QCameraExposure_MeteringMode(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->meteringMode();
}

void QCameraExposure_SetAutoAperture(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoAperture");
}

void QCameraExposure_SetAutoIsoSensitivity(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoIsoSensitivity");
}

void QCameraExposure_SetExposureMode(QtObjectPtr ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureMode", Q_ARG(QCameraExposure::ExposureMode, static_cast<QCameraExposure::ExposureMode>(mode)));
}

void QCameraExposure_SetFlashMode(QtObjectPtr ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setFlashMode", Q_ARG(QCameraExposure::FlashMode, static_cast<QCameraExposure::FlashMode>(mode)));
}

void QCameraExposure_SetManualIsoSensitivity(QtObjectPtr ptr, int iso){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualIsoSensitivity", Q_ARG(int, iso));
}

void QCameraExposure_SetMeteringMode(QtObjectPtr ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setMeteringMode", Q_ARG(QCameraExposure::MeteringMode, static_cast<QCameraExposure::MeteringMode>(mode)));
}

void QCameraExposure_SetSpotMeteringPoint(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QCameraExposure*>(ptr)->setSpotMeteringPoint(*static_cast<QPointF*>(point));
}

void QCameraExposure_ConnectApertureRangeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_DisconnectApertureRangeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_ConnectFlashReady(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

void QCameraExposure_DisconnectFlashReady(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

int QCameraExposure_IsAvailable(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->isAvailable();
}

int QCameraExposure_IsExposureModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isExposureModeSupported(static_cast<QCameraExposure::ExposureMode>(mode));
}

int QCameraExposure_IsFlashModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraExposure_IsFlashReady(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->isFlashReady();
}

int QCameraExposure_IsMeteringModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isMeteringModeSupported(static_cast<QCameraExposure::MeteringMode>(mode));
}

void QCameraExposure_ConnectIsoSensitivityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

void QCameraExposure_DisconnectIsoSensitivityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

int QCameraExposure_RequestedIsoSensitivity(QtObjectPtr ptr){
	return static_cast<QCameraExposure*>(ptr)->requestedIsoSensitivity();
}

void QCameraExposure_SetAutoShutterSpeed(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoShutterSpeed");
}

void QCameraExposure_ConnectShutterSpeedRangeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

void QCameraExposure_DisconnectShutterSpeedRangeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

