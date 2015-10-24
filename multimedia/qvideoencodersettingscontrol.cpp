#include "qvideoencodersettingscontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoEncoderSettings>
#include <QVideoEncoderSettingsControl>
#include "_cgo_export.h"

class MyQVideoEncoderSettingsControl: public QVideoEncoderSettingsControl {
public:
};

void QVideoEncoderSettingsControl_SetVideoSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QVideoEncoderSettingsControl*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

char* QVideoEncoderSettingsControl_SupportedVideoCodecs(QtObjectPtr ptr){
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->supportedVideoCodecs().join("|").toUtf8().data();
}

char* QVideoEncoderSettingsControl_VideoCodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QVideoEncoderSettingsControl*>(ptr)->videoCodecDescription(QString(codec)).toUtf8().data();
}

void QVideoEncoderSettingsControl_DestroyQVideoEncoderSettingsControl(QtObjectPtr ptr){
	static_cast<QVideoEncoderSettingsControl*>(ptr)->~QVideoEncoderSettingsControl();
}

