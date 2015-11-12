#include "qaudiodeviceinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioFormat>
#include <QAudioDeviceInfo>
#include "_cgo_export.h"

class MyQAudioDeviceInfo: public QAudioDeviceInfo {
public:
};

void* QAudioDeviceInfo_NewQAudioDeviceInfo(){
	return new QAudioDeviceInfo();
}

void* QAudioDeviceInfo_NewQAudioDeviceInfo2(void* other){
	return new QAudioDeviceInfo(*static_cast<QAudioDeviceInfo*>(other));
}

char* QAudioDeviceInfo_DeviceName(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->deviceName().toUtf8().data();
}

int QAudioDeviceInfo_IsFormatSupported(void* ptr, void* settings){
	return static_cast<QAudioDeviceInfo*>(ptr)->isFormatSupported(*static_cast<QAudioFormat*>(settings));
}

int QAudioDeviceInfo_IsNull(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->isNull();
}

char* QAudioDeviceInfo_SupportedCodecs(void* ptr){
	return static_cast<QAudioDeviceInfo*>(ptr)->supportedCodecs().join("|").toUtf8().data();
}

void QAudioDeviceInfo_DestroyQAudioDeviceInfo(void* ptr){
	static_cast<QAudioDeviceInfo*>(ptr)->~QAudioDeviceInfo();
}

