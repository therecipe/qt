#include "qaudioencodersettings.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioEncoderSettings>
#include "_cgo_export.h"

class MyQAudioEncoderSettings: public QAudioEncoderSettings {
public:
};

void* QAudioEncoderSettings_NewQAudioEncoderSettings(){
	return new QAudioEncoderSettings();
}

void* QAudioEncoderSettings_NewQAudioEncoderSettings2(void* other){
	return new QAudioEncoderSettings(*static_cast<QAudioEncoderSettings*>(other));
}

int QAudioEncoderSettings_BitRate(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->bitRate();
}

int QAudioEncoderSettings_ChannelCount(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->channelCount();
}

char* QAudioEncoderSettings_Codec(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->codec().toUtf8().data();
}

void* QAudioEncoderSettings_EncodingOption(void* ptr, char* option){
	return new QVariant(static_cast<QAudioEncoderSettings*>(ptr)->encodingOption(QString(option)));
}

int QAudioEncoderSettings_IsNull(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->isNull();
}

int QAudioEncoderSettings_SampleRate(void* ptr){
	return static_cast<QAudioEncoderSettings*>(ptr)->sampleRate();
}

void QAudioEncoderSettings_SetBitRate(void* ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setBitRate(rate);
}

void QAudioEncoderSettings_SetChannelCount(void* ptr, int channels){
	static_cast<QAudioEncoderSettings*>(ptr)->setChannelCount(channels);
}

void QAudioEncoderSettings_SetCodec(void* ptr, char* codec){
	static_cast<QAudioEncoderSettings*>(ptr)->setCodec(QString(codec));
}

void QAudioEncoderSettings_SetEncodingOption(void* ptr, char* option, void* value){
	static_cast<QAudioEncoderSettings*>(ptr)->setEncodingOption(QString(option), *static_cast<QVariant*>(value));
}

void QAudioEncoderSettings_SetSampleRate(void* ptr, int rate){
	static_cast<QAudioEncoderSettings*>(ptr)->setSampleRate(rate);
}

void QAudioEncoderSettings_DestroyQAudioEncoderSettings(void* ptr){
	static_cast<QAudioEncoderSettings*>(ptr)->~QAudioEncoderSettings();
}

