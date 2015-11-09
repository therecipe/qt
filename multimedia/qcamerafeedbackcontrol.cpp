#include "qcamerafeedbackcontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QString>
#include <QCameraFeedbackControl>
#include "_cgo_export.h"

class MyQCameraFeedbackControl: public QCameraFeedbackControl {
public:
};

int QCameraFeedbackControl_IsEventFeedbackEnabled(void* ptr, int event){
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event));
}

int QCameraFeedbackControl_IsEventFeedbackLocked(void* ptr, int event){
	return static_cast<QCameraFeedbackControl*>(ptr)->isEventFeedbackLocked(static_cast<QCameraFeedbackControl::EventType>(event));
}

void QCameraFeedbackControl_ResetEventFeedback(void* ptr, int event){
	static_cast<QCameraFeedbackControl*>(ptr)->resetEventFeedback(static_cast<QCameraFeedbackControl::EventType>(event));
}

int QCameraFeedbackControl_SetEventFeedbackEnabled(void* ptr, int event, int enabled){
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackEnabled(static_cast<QCameraFeedbackControl::EventType>(event), enabled != 0);
}

int QCameraFeedbackControl_SetEventFeedbackSound(void* ptr, int event, char* filePath){
	return static_cast<QCameraFeedbackControl*>(ptr)->setEventFeedbackSound(static_cast<QCameraFeedbackControl::EventType>(event), QString(filePath));
}

void QCameraFeedbackControl_DestroyQCameraFeedbackControl(void* ptr){
	static_cast<QCameraFeedbackControl*>(ptr)->~QCameraFeedbackControl();
}

