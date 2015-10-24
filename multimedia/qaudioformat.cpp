#include "qaudioformat.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QAudioFormat>
#include "_cgo_export.h"

class MyQAudioFormat: public QAudioFormat {
public:
};

QtObjectPtr QAudioFormat_NewQAudioFormat(){
	return new QAudioFormat();
}

QtObjectPtr QAudioFormat_NewQAudioFormat2(QtObjectPtr other){
	return new QAudioFormat(*static_cast<QAudioFormat*>(other));
}

int QAudioFormat_ByteOrder(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->byteOrder();
}

int QAudioFormat_BytesPerFrame(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->bytesPerFrame();
}

int QAudioFormat_ChannelCount(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->channelCount();
}

char* QAudioFormat_Codec(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->codec().toUtf8().data();
}

int QAudioFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->isValid();
}

int QAudioFormat_SampleRate(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleRate();
}

int QAudioFormat_SampleSize(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleSize();
}

int QAudioFormat_SampleType(QtObjectPtr ptr){
	return static_cast<QAudioFormat*>(ptr)->sampleType();
}

void QAudioFormat_SetByteOrder(QtObjectPtr ptr, int byteOrder){
	static_cast<QAudioFormat*>(ptr)->setByteOrder(static_cast<QAudioFormat::Endian>(byteOrder));
}

void QAudioFormat_SetChannelCount(QtObjectPtr ptr, int channels){
	static_cast<QAudioFormat*>(ptr)->setChannelCount(channels);
}

void QAudioFormat_SetCodec(QtObjectPtr ptr, char* codec){
	static_cast<QAudioFormat*>(ptr)->setCodec(QString(codec));
}

void QAudioFormat_SetSampleRate(QtObjectPtr ptr, int samplerate){
	static_cast<QAudioFormat*>(ptr)->setSampleRate(samplerate);
}

void QAudioFormat_SetSampleSize(QtObjectPtr ptr, int sampleSize){
	static_cast<QAudioFormat*>(ptr)->setSampleSize(sampleSize);
}

void QAudioFormat_SetSampleType(QtObjectPtr ptr, int sampleType){
	static_cast<QAudioFormat*>(ptr)->setSampleType(static_cast<QAudioFormat::SampleType>(sampleType));
}

void QAudioFormat_DestroyQAudioFormat(QtObjectPtr ptr){
	static_cast<QAudioFormat*>(ptr)->~QAudioFormat();
}

