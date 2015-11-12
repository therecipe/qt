#include "qimageencodersettings.h"
#include <QSize>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QImage>
#include <QImageEncoderSettings>
#include "_cgo_export.h"

class MyQImageEncoderSettings: public QImageEncoderSettings {
public:
};

void* QImageEncoderSettings_NewQImageEncoderSettings(){
	return new QImageEncoderSettings();
}

void* QImageEncoderSettings_NewQImageEncoderSettings2(void* other){
	return new QImageEncoderSettings(*static_cast<QImageEncoderSettings*>(other));
}

char* QImageEncoderSettings_Codec(void* ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QImageEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QImageEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

int QImageEncoderSettings_IsNull(void* ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->isNull();
}

void QImageEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QImageEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QImageEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QImageEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QImageEncoderSettings_SetResolution(void* ptr, void* resolution){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QImageEncoderSettings_SetResolution2(void* ptr, int width, int height){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(width, height);
}

void QImageEncoderSettings_DestroyQImageEncoderSettings(void* ptr){
	static_cast<QImageEncoderSettings*>(ptr)->~QImageEncoderSettings();
}

