#include "qcameraviewfindersettingscontrol.h"
#include <QCameraViewfinderSettings>
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraViewfinderSettingsControl>
#include "_cgo_export.h"

class MyQCameraViewfinderSettingsControl: public QCameraViewfinderSettingsControl {
public:
};

int QCameraViewfinderSettingsControl_IsViewfinderParameterSupported(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->isViewfinderParameterSupported(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter));
}

void QCameraViewfinderSettingsControl_SetViewfinderParameter(QtObjectPtr ptr, int parameter, char* value){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->setViewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter), QVariant(value));
}

char* QCameraViewfinderSettingsControl_ViewfinderParameter(QtObjectPtr ptr, int parameter){
	return static_cast<QCameraViewfinderSettingsControl*>(ptr)->viewfinderParameter(static_cast<QCameraViewfinderSettingsControl::ViewfinderParameter>(parameter)).toString().toUtf8().data();
}

void QCameraViewfinderSettingsControl_DestroyQCameraViewfinderSettingsControl(QtObjectPtr ptr){
	static_cast<QCameraViewfinderSettingsControl*>(ptr)->~QCameraViewfinderSettingsControl();
}

