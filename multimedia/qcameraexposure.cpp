#include "qcameraexposure.h"
#include <QObject>
#include <QCamera>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QPointF>
#include <QPoint>
#include <QVariant>
#include <QCameraExposure>
#include "_cgo_export.h"

class MyQCameraExposure: public QCameraExposure {
public:
void Signal_ApertureRangeChanged(){callbackQCameraExposureApertureRangeChanged(this->objectName().toUtf8().data());};
void Signal_FlashReady(bool ready){callbackQCameraExposureFlashReady(this->objectName().toUtf8().data(), ready);};
void Signal_IsoSensitivityChanged(int value){callbackQCameraExposureIsoSensitivityChanged(this->objectName().toUtf8().data(), value);};
void Signal_ShutterSpeedRangeChanged(){callbackQCameraExposureShutterSpeedRangeChanged(this->objectName().toUtf8().data());};
};

double QCameraExposure_Aperture(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->aperture());
}

double QCameraExposure_ExposureCompensation(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->exposureCompensation());
}

int QCameraExposure_ExposureMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->exposureMode();
}

int QCameraExposure_FlashMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->flashMode();
}

int QCameraExposure_IsoSensitivity(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isoSensitivity();
}

int QCameraExposure_MeteringMode(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->meteringMode();
}

void QCameraExposure_SetAutoAperture(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoAperture");
}

void QCameraExposure_SetAutoIsoSensitivity(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoIsoSensitivity");
}

void QCameraExposure_SetExposureCompensation(void* ptr, double ev){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureCompensation", Q_ARG(qreal, static_cast<qreal>(ev)));
}

void QCameraExposure_SetExposureMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setExposureMode", Q_ARG(QCameraExposure::ExposureMode, static_cast<QCameraExposure::ExposureMode>(mode)));
}

void QCameraExposure_SetFlashMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setFlashMode", Q_ARG(QCameraExposure::FlashMode, static_cast<QCameraExposure::FlashMode>(mode)));
}

void QCameraExposure_SetManualAperture(void* ptr, double aperture){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualAperture", Q_ARG(qreal, static_cast<qreal>(aperture)));
}

void QCameraExposure_SetManualIsoSensitivity(void* ptr, int iso){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualIsoSensitivity", Q_ARG(int, iso));
}

void QCameraExposure_SetMeteringMode(void* ptr, int mode){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setMeteringMode", Q_ARG(QCameraExposure::MeteringMode, static_cast<QCameraExposure::MeteringMode>(mode)));
}

void QCameraExposure_SetSpotMeteringPoint(void* ptr, void* point){
	static_cast<QCameraExposure*>(ptr)->setSpotMeteringPoint(*static_cast<QPointF*>(point));
}

void QCameraExposure_ConnectApertureRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_DisconnectApertureRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::apertureRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ApertureRangeChanged));;
}

void QCameraExposure_ConnectFlashReady(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

void QCameraExposure_DisconnectFlashReady(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(bool)>(&QCameraExposure::flashReady), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(bool)>(&MyQCameraExposure::Signal_FlashReady));;
}

int QCameraExposure_IsAvailable(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isAvailable();
}

int QCameraExposure_IsExposureModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isExposureModeSupported(static_cast<QCameraExposure::ExposureMode>(mode));
}

int QCameraExposure_IsFlashModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isFlashModeSupported(static_cast<QCameraExposure::FlashMode>(mode));
}

int QCameraExposure_IsFlashReady(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->isFlashReady();
}

int QCameraExposure_IsMeteringModeSupported(void* ptr, int mode){
	return static_cast<QCameraExposure*>(ptr)->isMeteringModeSupported(static_cast<QCameraExposure::MeteringMode>(mode));
}

void QCameraExposure_ConnectIsoSensitivityChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

void QCameraExposure_DisconnectIsoSensitivityChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)(int)>(&QCameraExposure::isoSensitivityChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)(int)>(&MyQCameraExposure::Signal_IsoSensitivityChanged));;
}

double QCameraExposure_RequestedAperture(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->requestedAperture());
}

int QCameraExposure_RequestedIsoSensitivity(void* ptr){
	return static_cast<QCameraExposure*>(ptr)->requestedIsoSensitivity();
}

double QCameraExposure_RequestedShutterSpeed(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->requestedShutterSpeed());
}

void QCameraExposure_SetAutoShutterSpeed(void* ptr){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setAutoShutterSpeed");
}

void QCameraExposure_SetManualShutterSpeed(void* ptr, double seconds){
	QMetaObject::invokeMethod(static_cast<QCameraExposure*>(ptr), "setManualShutterSpeed", Q_ARG(qreal, static_cast<qreal>(seconds)));
}

double QCameraExposure_ShutterSpeed(void* ptr){
	return static_cast<double>(static_cast<QCameraExposure*>(ptr)->shutterSpeed());
}

void QCameraExposure_ConnectShutterSpeedRangeChanged(void* ptr){
	QObject::connect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

void QCameraExposure_DisconnectShutterSpeedRangeChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraExposure*>(ptr), static_cast<void (QCameraExposure::*)()>(&QCameraExposure::shutterSpeedRangeChanged), static_cast<MyQCameraExposure*>(ptr), static_cast<void (MyQCameraExposure::*)()>(&MyQCameraExposure::Signal_ShutterSpeedRangeChanged));;
}

