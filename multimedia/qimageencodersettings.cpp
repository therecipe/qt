#include "qimageencodersettings.h"
#include <QUrl>
#include <QModelIndex>
#include <QSize>
#include <QImage>
#include <QString>
#include <QVariant>
#include <QImageEncoderSettings>
#include "_cgo_export.h"

class MyQImageEncoderSettings: public QImageEncoderSettings {
public:
};

QtObjectPtr QImageEncoderSettings_NewQImageEncoderSettings(){
	return new QImageEncoderSettings();
}

QtObjectPtr QImageEncoderSettings_NewQImageEncoderSettings2(QtObjectPtr other){
	return new QImageEncoderSettings(*static_cast<QImageEncoderSettings*>(other));
}

char* QImageEncoderSettings_Codec(QtObjectPtr ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->codec().toUtf8().data();
}

char* QImageEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option){
	return static_cast<QImageEncoderSettings*>(ptr)->encodingOption(QString(option)).toString().toUtf8().data();
}

int QImageEncoderSettings_IsNull(QtObjectPtr ptr){
	return static_cast<QImageEncoderSettings*>(ptr)->isNull();
}

void QImageEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec){
	static_cast<QImageEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QImageEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value){
	static_cast<QImageEncoderSettings*>(ptr)->setEncodingOption(QString(option), QVariant(value));
}

void QImageEncoderSettings_SetResolution(QtObjectPtr ptr, QtObjectPtr resolution){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(*static_cast<QSize*>(resolution));
}

void QImageEncoderSettings_SetResolution2(QtObjectPtr ptr, int width, int height){
	static_cast<QImageEncoderSettings*>(ptr)->setResolution(width, height);
}

void QImageEncoderSettings_DestroyQImageEncoderSettings(QtObjectPtr ptr){
	static_cast<QImageEncoderSettings*>(ptr)->~QImageEncoderSettings();
}

