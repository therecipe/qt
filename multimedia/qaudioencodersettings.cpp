#include "qaudioencodersettings.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QAudioEncoderSettings>
#include "_cgo_export.h"

class MyQAudioEncoderSettings: public QAudioEncoderSettings {
public:
};

QtObjectPtr QAudioEncoderSettings_NewQAudioEncoderSettings(){
	return new QAudioEncoderSettings();
}

QtObjectPtr QAudioEncoderSettings_NewQAudioEncoderSettings2(QtObjectPtr other){
	return new QAudioEncoderSettings(*static_cast<QAudioEncoderSettings*>(other));
}

int QAudioEncoderSettings_BitRate(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->bitRate();
}

int QAudioEncoderSettings_ChannelCount(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->channelCount();
}

char* QAudioEncoderSettings_Codec(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->codec().toUtf8().data();
}

char* QAudioEncoderSettings_EncodingOption(QtObjectPtr ptr, char* option){
	return static_cast<QAudioEncoderSettings*>(ptr)->encodingOption(QString(option)).toString().toUtf8().data();
}

int QAudioEncoderSettings_IsNull(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->isNull();
}

int QAudioEncoderSettings_SampleRate(QtObjectPtr ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->sampleRate();
}

void QAudioEncoderSettings_SetBitRate(QtObjectPtr ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setBitRate(rate);
}

void QAudioEncoderSettings_SetChannelCount(QtObjectPtr ptr, int channels){
	static_cast<QAudioEncoderSettings*>(ptr)->setChannelCount(channels);
}

void QAudioEncoderSettings_SetCodec(QtObjectPtr ptr, char* codec){
	static_cast<QAudioEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QAudioEncoderSettings_SetEncodingOption(QtObjectPtr ptr, char* option, char* value){
	static_cast<QAudioEncoderSettings*>(ptr)->setEncodingOption(QString(option), QVariant(value));
}

void QAudioEncoderSettings_SetSampleRate(QtObjectPtr ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setSampleRate(rate);
}

void QAudioEncoderSettings_DestroyQAudioEncoderSettings(QtObjectPtr ptr){
	static_cast<QAudioEncoderSettings*>(ptr)->~QAudioEncoderSettings();
}

