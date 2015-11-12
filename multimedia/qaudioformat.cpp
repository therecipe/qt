#include "qaudioformat.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioFormat>
#include "_cgo_export.h"

class MyQAudioFormat: public QAudioFormat {
public:
};

void* QAudioFormat_NewQAudioFormat(){
	return new QAudioFormat();
}

void* QAudioFormat_NewQAudioFormat2(void* other){
	return new QAudioFormat(*static_cast<QAudioFormat*>(other));
}

int QAudioFormat_ByteOrder(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->byteOrder();
}

int QAudioFormat_BytesPerFrame(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->bytesPerFrame();
}

int QAudioFormat_ChannelCount(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->channelCount();
}

char* QAudioFormat_Codec(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->codec().toUtf8().data();
}

int QAudioFormat_IsValid(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->isValid();
}

int QAudioFormat_SampleRate(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleRate();
}

int QAudioFormat_SampleSize(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleSize();
}

int QAudioFormat_SampleType(void* ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleType();
}

void QAudioFormat_SetByteOrder(void* ptr, int byteOrder){
	static_cast<QAudioFormat*>(ptr)->setByteOrder(static_cast<QAudioFormat::Endian>(byteOrder));
}

void QAudioFormat_SetChannelCount(void* ptr, int channels){
	static_cast<QAudioFormat*>(ptr)->setChannelCount(channels);
}

void QAudioFormat_SetCodec(void* ptr, char* codec){
	static_cast<QAudioFormat*>(ptr)->setCodec(QString(codec));
}

void QAudioFormat_SetSampleRate(void* ptr, int samplerate){
	static_cast<QAudioFormat*>(ptr)->setSampleRate(samplerate);
}

void QAudioFormat_SetSampleSize(void* ptr, int sampleSize){
	static_cast<QAudioFormat*>(ptr)->setSampleSize(sampleSize);
}

void QAudioFormat_SetSampleType(void* ptr, int sampleType){
	static_cast<QAudioFormat*>(ptr)->setSampleType(static_cast<QAudioFormat::SampleType>(sampleType));
}

void QAudioFormat_DestroyQAudioFormat(void* ptr){
	static_cast<QAudioFormat*>(ptr)->~QAudioFormat();
}

