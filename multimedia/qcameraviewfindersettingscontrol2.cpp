#include "qcameraviewfindersettingscontrol2.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraViewfinderSettingsControl>
#include <QCameraViewfinderSettings>
#include <QString>
#include <QVariant>
#include <QCameraViewfinderSettingsControl2>
#include "_cgo_export.h"

class MyQCameraViewfinderSettingsControl2: public QCameraViewfinderSettingsControl2 {
public:
};

void QCameraViewfinderSettingsControl2_SetViewfinderSettings(void* ptr, void* settings){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->setViewfinderSettings(*static_cast<QCameraViewfinderSettings*>(settings));
}

void QCameraViewfinderSettingsControl2_DestroyQCameraViewfinderSettingsControl2(void* ptr){
	static_cast<QCameraViewfinderSettingsControl2*>(ptr)->~QCameraViewfinderSettingsControl2();
}

