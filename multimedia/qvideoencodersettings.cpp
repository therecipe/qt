#include "qvideoencodersettings.h"
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QVideoEncoderSettings>
#include "_cgo_export.h"

class MyQVideoEncoderSettings: public QVideoEncoderSettings {
public:
};

QtObjectPtr QVideoEncoderSettings_NewQVideoEncoderSettings(){
	return new QVideoEncoderSettings();
}

QtObjectPtr QVideoEncoderSettings_NewQVideoEncoderSettings2(QtObjectPtr other){
	return new QVideoEncoderSettings(*static_cast<QVideoEncoderSettings*>(other));
}

int QVideoEncoderSettings_BitRate(QtObjectPtr ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->bitRate();
}

char* QVideoEncoderSettings_Codec(QtObjectPtr ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->codec().toUtf8().data();
}

char* QVideoEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option){
	return static_cast<QVideoEncoderSettings*>(ptr)->encodingOption(QString(option)).toString().toUtf8().data();
}

int QVideoEncoderSettings_IsNull(QtObjectPtr ptr){
	return static_cast<QVideoEncoderSettings*>(ptr)->isNull();
}

void QVideoEncoderSettings_SetBitRate(QtObjectPtr ptr, int value){
	static_cast<QVideoEncoderSettings*>(ptr)->setBitRate(value);
}

void QVideoEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec){
	static_cast<QVideoEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QVideoEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value){
	static_cast<QVideoEncoderSettings*>(ptr)->setEncodingOption(QString(option), QVariant(value));
}

void QVideoEncoderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QVideoEncoderSettings_SetResolution2(QtObjectPtr ptr, int width, int height){
	static_cast<QVideoEncoderSettings*>(ptr)->setResolution(width, height);
}

void QVideoEncoderSettings_DestroyQVideoEncoderSettings(QtObjectPtr ptr){
	static_cast<QVideoEncoderSettings*>(ptr)->~QVideoEncoderSettings();
}

