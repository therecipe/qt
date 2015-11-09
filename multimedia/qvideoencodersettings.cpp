#include "qvideoencodersettings.h"
#include <QModelIndex>
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QVideoEncoderSettings>
#include "_cgo_export.h"

class MyQVideoEncoderSettings: public QVideoEncoderSettings {
public:
};

void QVideoEncoderSettings_SetFrameRate(void* ptr, double rate){
	static_cast<QVideoEncoderSettings*>(ptr)->setFrameRate(static_cast<qreal>(rate));
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings(){
	return new QVideoEncoderSettings();
}

void* QVideoEncoderSettings_NewQVideoEncoderSettings2(void* other){
	return new QVideoEncoderSettings(*static_cast<QVideoEncoderSettings*>(other));
}

int QVideoEncoderSettings_BitRate(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->bitRate();
}

char* QVideoEncoderSettings_Codec(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QVideoEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QVideoEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

double QVideoEncoderSettings_FrameRate(void* ptr){
	return static_cast<double>(static_cast<QVideoEncoderSettings*>(ptr)->frameRate());
}

int QVideoEncoderSettings_IsNull(void* ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->isNull();
}

void QVideoEncoderSettings_SetBitRate(void* ptr, int value){
	static_cast<QVideoEncoderSettings*>(ptr)->setBitRate(value);
}

void QVideoEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QVideoEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QVideoEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QVideoEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QVideoEncoderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QVideoEncoderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(width, height);
}

void QVideoEncoderSettings_DestroyQVideoEncoderSettings(void* ptr){
	static_cast<QVideoEncoderSettings*>(ptr)->~QVideoEncoderSettings();
}

