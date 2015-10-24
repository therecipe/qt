#include "qcameraviewfindersettingscontrol2.h"
#include <QCameraViewfinderSettings>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraViewfinderSettingsControl>
#include <QCameraViewfinderSettingsControl2>
#include "_cgo_export.h"

class MyQCameraViewfinderSettingsControl2: public QCameraViewfinderSettingsControl2 {
public:
};

void QCameraViewfinderSettingsControl2_SetViewfinderSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(QtObjectPtr ptr){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->~QCameraViewfinderSettingsControl2();
}

