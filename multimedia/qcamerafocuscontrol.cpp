#include "qcamerafocuscontrol.h"
#include <QVariant>
#include <QUrl>
#include <QCamera>
#include <QPointF>
#include <QString>
#include <QModelIndex>
#include <QCameraFocus>
#include <QObject>
#include <QPoint>
#include <QCameraFocusControl>
#include "_cgo_export.h"

class MyQCameraFocusControl: public QCameraFocusControl {
public:
void Signal_FocusModeChanged(QCameraFocus::FocusModes mode){callbackQCameraFocusControlFocusModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_FocusPointModeChanged(QCameraFocus::FocusPointMode mode){callbackQCameraFocusControlFocusPointModeChanged(this->objectName().toUtf8().data(), mode);};
void Signal_FocusZonesChanged(){callbackQCameraFocusControlFocusZonesChanged(this->objectName().toUtf8().data());};
};

int QCameraFocusControl_FocusMode(QtObjectPtr ptr){
	return static_cast<QCameraFocusControl*>(ptr)->focusMode();
}

void QCameraFocusControl_ConnectFocusModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));;
}

void QCameraFocusControl_DisconnectFocusModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusModes)>(&QCameraFocusControl::focusModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusModes)>(&MyQCameraFocusControl::Signal_FocusModeChanged));;
}

int QCameraFocusControl_FocusPointMode(QtObjectPtr ptr){
	return static_cast<QCameraFocusControl*>(ptr)->focusPointMode();
}

void QCameraFocusControl_ConnectFocusPointModeChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));;
}

void QCameraFocusControl_DisconnectFocusPointModeChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&QCameraFocusControl::focusPointModeChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)(QCameraFocus::FocusPointMode)>(&MyQCameraFocusControl::Signal_FocusPointModeChanged));;
}

void QCameraFocusControl_ConnectFocusZonesChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));;
}

void QCameraFocusControl_DisconnectFocusZonesChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCameraFocusControl*>(ptr), static_cast<void (QCameraFocusControl::*)()>(&QCameraFocusControl::focusZonesChanged), static_cast<MyQCameraFocusControl*>(ptr), static_cast<void (MyQCameraFocusControl::*)()>(&MyQCameraFocusControl::Signal_FocusZonesChanged));;
}

int QCameraFocusControl_IsFocusModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraFocusControl*>(ptr)->isFocusModeSupported(static_cast<QCameraFocus::FocusMode>(mode));
}

int QCameraFocusControl_IsFocusPointModeSupported(QtObjectPtr ptr, int mode){
	return static_cast<QCameraFocusControl*>(ptr)->isFocusPointModeSupported(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_SetCustomFocusPoint(QtObjectPtr ptr, QtObjectPtr point){
	static_cast<QCameraFocusControl*>(ptr)->setCustomFocusPoint(*static_cast<QPointF*>(point));
}

void QCameraFocusControl_SetFocusMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraFocusControl*>(ptr)->setFocusMode(static_cast<QCameraFocus::FocusMode>(mode));
}

void QCameraFocusControl_SetFocusPointMode(QtObjectPtr ptr, int mode){
	static_cast<QCameraFocusControl*>(ptr)->setFocusPointMode(static_cast<QCameraFocus::FocusPointMode>(mode));
}

void QCameraFocusControl_DestroyQCameraFocusControl(QtObjectPtr ptr){
	static_cast<QCameraFocusControl*>(ptr)->~QCameraFocusControl();
}

