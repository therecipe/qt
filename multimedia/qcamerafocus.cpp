#include "qcamerafocus.h"
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QObject>
#include <QPointF>
#include <QCameraFocus>
#include "_cgo_export.h"

class MyQCameraFocus: public QCameraFocus {
public:
void Signal_FocusZonesChanged(){callbackQCameraFocusFocusZonesChanged(this->objectName().toUtf8().data());};
};

double QCameraFocus_DigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->digitalZoom());
}

int QCameraFocus_FocusMode(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->focusMode();
}

int QCameraFocus_FocusPointMode(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->focusPointMode();
}

double QCameraFocus_OpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->opticalZoom());
}

void QCameraFocus_SetCustomFocusPoint(void* ptr, void* point){
	static_cast<QCameraFocus*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocus_SetFocusMode(void* ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocus_SetFocusPointMode(void* ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocus_ConnectFocusZonesChanged(void* ptr){
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

void QCameraFocus_DisconnectFocusZonesChanged(void* ptr){
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

int QCameraFocus_IsAvailable(void* ptr){
	return static_cast<QCameraFocus*>(ptr)->isAvailable();
}

int QCameraFocus_IsFocusModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

int QCameraFocus_IsFocusPointModeSupported(void* ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

double QCameraFocus_MaximumDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->maximumDigitalZoom());
}

double QCameraFocus_MaximumOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraFocus*>(ptr)->maximumOpticalZoom());
}

void QCameraFocus_ZoomTo(void* ptr, double optical, double digital){
	static_cast<QCameraFocus*>(ptr)->zoomTo(static_cast<qreal>(optical), static_cast<qreal>(digital));
}

