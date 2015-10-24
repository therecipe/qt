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

QtObjectPtr QMediaResource_NewQMediaResource(){
	return new QMediaResource();
}

QtObjectPtr QMediaResource_NewQMediaResource4(QtObjectPtr other){
	return new QMediaResource(*static_cast<QMediaResource*>(other));
}

QtObjectPtr QMediaResource_NewQMediaResource3(QtObjectPtr request, char* mimeType){
	return new QMediaResource(*static_cast<QNetworkRequest*>(request), QString(mimeType));
}

QtObjectPtr QMediaResource_NewQMediaResource2(char* url, char* mimeType){
	return new QMediaResource(QUrl(QString(url)), QString(mimeType));
}

int QMediaResource_AudioBitRate(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->audioBitRate();
}

char* QMediaResource_AudioCodec(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->audioCodec().toUtf8().data();
}

int QMediaResource_ChannelCount(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->channelCount();
}

int QMediaResource_IsNull(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->isNull();
}

char* QMediaResource_Language(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->language().toUtf8().data();
}

char* QMediaResource_MimeType(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->mimeType().toUtf8().data();
}

int QMediaResource_SampleRate(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->sampleRate();
}

void QMediaResource_SetAudioBitRate(QtObjectPtr ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setAudioBitRate(rate);
}

void QMediaResource_SetAudioCodec(QtObjectPtr ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setAudioCodec(QString(codec));
}

void QMediaResource_SetChannelCount(QtObjectPtr ptr, int channels){
	static_cast<QMediaResource*>(ptr)->setChannelCount(channels);
}

void QMediaResource_SetLanguage(QtObjectPtr ptr, char* language){
	static_cast<QMediaResource*>(ptr)->setLanguage(QString(language));
}

void QMediaResource_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution){
	static_cast<QMediaResource*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QMediaResource_SetResolution2(QtObjectPtr ptr, int width, int height){
	static_cast<QMediaResource*>(ptr)->setResolution(width, height);
}

void QMediaResource_SetSampleRate(QtObjectPtr ptr, int sampleRate){
	static_cast<QMediaResource*>(ptr)->setSampleRate(sampleRate);
}

void QMediaResource_SetVideoBitRate(QtObjectPtr ptr, int rate){
	static_cast<QMediaResource*>(ptr)->setVideoBitRate(rate);
}

void QMediaResource_SetVideoCodec(QtObjectPtr ptr, char* codec){
	static_cast<QMediaResource*>(ptr)->setVideoCodec(QString(codec));
}

char* QMediaResource_Url(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->url().toString().toUtf8().data();
}

int QMediaResource_VideoBitRate(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->videoBitRate();
}

char* QMediaResource_VideoCodec(QtObjectPtr ptr){
	return static_cast<QMediaResource*>(ptr)->videoCodec().toUtf8().data();
}

void QMediaResource_DestroyQMediaResource(QtObjectPtr ptr){
	static_cast<QMediaResource*>(ptr)->~QMediaResource();
}

