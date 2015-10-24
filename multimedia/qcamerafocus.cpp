#include "qcamerafocus.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QObject>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QCameraFocus>
#include "_cgo_export.h"

class MyQCameraFocus: public QCameraFocus {
public:
void Signal_FocusZonesChanged(){callbackQCameraFocusFocusZonesChanged(this->objectName().toUtf8().data());};
};

int QCameraFocus_FocusMode(QtObjectPtr ptr){
	return static_cast<QCameraFocus*>(ptr)->focusMode();
}

int QCameraFocus_FocusPointMode(QtObjectPtr ptr){
	return static_cast<QCameraFocus*>(ptr)->focusPointMode();
}

void QCameraFocus_SetCustomFocusPoint(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QCameraFocus*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocus_SetFocusMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocus_SetFocusPointMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraFocus*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocus_ConnectFocusZonesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

void QCameraFocus_DisconnectFocusZonesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraFocus*>(ptr), static_cast<void (QCameraFocus::*)()>(&QCameraFocus::focusZonesChanged), static_cast<MyQCameraFocus*>(ptr), static_cast<void (MyQCameraFocus::*)()>(&MyQCameraFocus::Signal_FocusZonesChanged));;
}

int QCameraFocus_IsAvailable(QtObjectPtr ptr){
	return static_cast<QCameraFocus*>(ptr)->isAvailable();
}

int QCameraFocus_IsFocusModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

int QCameraFocus_IsFocusPointModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraFocus*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

