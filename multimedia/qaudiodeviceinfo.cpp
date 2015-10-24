#include "qaudiodeviceinfo.h"
#include <QModelIndex>
#include <QAudioFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioDeviceInfo>
#include "_cgo_export.h"

class MyQAudioDeviceInfo: public QAudioDeviceInfo {
public:
};

QtObjectPtr QAudioDeviceInfo_NewQAudioDeviceInfo(){
	return new QAudioDeviceInfo();
}

QtObjectPtr QAudioDeviceInfo_NewQAudioDeviceInfo2(QtObjectPtr other){
	return new QAudioDeviceInfo(*static_cast<QAudioDeviceInfo*>(other));
}

char* QAudioDeviceInfo_DeviceName(QtObjectPtr ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->deviceName().toUtf8().data();
}

int QAudioDeviceInfo_IsFormatSupported(QtObjectPtr ptr, QtObjectPtr settings){
	return static_cast<QAudioDeviceInfo*>(ptr)->isFormatSupported(*static_cast<QAudioFormat*>(settings));
}

int QAudioDeviceInfo_IsNull(QtObjectPtr ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->isNull();
}

char* QAudioDeviceInfo_SupportedCodecs(QtObjectPtr ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->supportedCodecs().join("|").toUtf8().data();
}

void QAudioDeviceInfo_DestroyQAudioDeviceInfo(QtObjectPtr ptr){
	static_cast<QAudioDeviceInfo*>(ptr)->~QAudioDeviceInfo();
}

