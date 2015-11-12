#include "qcameraviewfindersettingscontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QCameraViewfinderSettings>
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QCameraViewfinderSettingsControl>
#include "_cgo_export.h"

class MyQCameraViewfinderSettingsControl: public QCameraViewfinderSettingsControl {
public:
};

int QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->isViewfinderParameterSupported(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter));
}

void QCameraViewfinderSettingsControl_SetViewfinderParameter(void* ptr, int parameter, void* value){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->setViewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraViewfinderSettingsControl_ViewfinderParameter(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraViewfinderSettingsControl*>(ptr)->viewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter)));
}

void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(void* ptr){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->~QCameraViewfinderSettingsControl();
}

