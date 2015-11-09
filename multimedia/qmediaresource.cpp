#include "qmediaresource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QNetworkRequest>
#include <QMediaResource>
#include "_cgo_export.h"

class MyQMediaResource: public QMediaResource {
public:
};

void* QMediaResource_NewQMediaResource(){
	return new QMediaResource();
}

void* QMediaResource_NewQMediaResource4(void* other){
	return new QMediaResource(*static_cast<QMediaResource*>(other));
}

void* QMediaResource_NewQMediaResource3(void* request, char* mimeType){
	return new QMediaResource(*static_cast<QNetworkRequest*>(request), QString(mimeType));
}

void* QMediaResource_NewQMediaResource2(void* url, char* mimeType){
	return new QMediaResource(*static_cast<QUrl*>(url), QString(mimeType));
}

int QMediaResource_AudioBitRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->audioBitRate();
}

char* QMediaResource_AudioCodec(void* ptr){
	return static_cast<QMediaResource*>(ptr)->audioCodec().toUtf8().data();
}

int QMediaResource_ChannelCount(void* ptr){
	return static_cast<QMediaResource*>(ptr)->channelCount();
}

int QMediaResource_IsNull(void* ptr){
	return static_cast<QMediaResource*>(ptr)->isNull();
}

char* QMediaResource_Language(void* ptr){
	return static_cast<QMediaResource*>(ptr)->language().toUtf8().data();
}

char* QMediaResource_MimeType(void* ptr){
	return static_cast<QMediaResource*>(ptr)->mimeType().toUtf8().data();
}

int QMediaResource_SampleRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->sampleRate();
}

void QMediaResource_SetAudioBitRate(void* ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setAudioBitRate(rate);
}

void QMediaResource_SetAudioCodec(void* ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setAudioCodec(QString(codec));
}

void QMediaResource_SetChannelCount(void* ptr, int channels){
	static_cast<QMediaResource*>(ptr)->setChannelCount(channels);
}

void QMediaResource_SetLanguage(void* ptr, char* language){
	static_cast<QMediaResource*>(ptr)->setLanguage(QString(language));
}

void QMediaResource_SetResolution(void* ptr, void* resolution){
	static_cast<QMediaResource*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QMediaResource_SetResolution2(void* ptr, int width, int height){
	static_cast<QMediaResource*>(ptr)->setResolution(width, height);
}

void QMediaResource_SetSampleRate(void* ptr, int sampleRate){
	static_cast<QMediaResource*>(ptr)->setSampleRate(sampleRate);
}

void QMediaResource_SetVideoBitRate(void* ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setVideoBitRate(rate);
}

void QMediaResource_SetVideoCodec(void* ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setVideoCodec(QString(codec));
}

int QMediaResource_VideoBitRate(void* ptr){
	return static_cast<QMediaResource*>(ptr)->videoBitRate();
}

char* QMediaResource_VideoCodec(void* ptr){
	return static_cast<QMediaResource*>(ptr)->videoCodec().toUtf8().data();
}

void QMediaResource_DestroyQMediaResource(void* ptr){
	static_cast<QMediaResource*>(ptr)->~QMediaResource();
}

