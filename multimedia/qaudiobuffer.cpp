#include "qaudiobuffer.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QAudioBuffer>
#include "_cgo_export.h"

class MyQAudioBuffer: public QAudioBuffer {
public:
};

void* QAudioBuffer_NewQAudioBuffer(){
	return new QAudioBuffer();
}

void* QAudioBuffer_NewQAudioBuffer3(void* other){
	return new QAudioBuffer(*static_cast<QAudioBuffer*>(other));
}

int QAudioBuffer_ByteCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->byteCount();
}

void* QAudioBuffer_ConstData(void* ptr){
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->constData());
}

void* QAudioBuffer_Data2(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->data();
}

void* QAudioBuffer_Data(void* ptr){
	return const_cast<void*>(static_cast<QAudioBuffer*>(ptr)->data());
}

int QAudioBuffer_FrameCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->frameCount();
}

int QAudioBuffer_IsValid(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->isValid();
}

int QAudioBuffer_SampleCount(void* ptr){
	return static_cast<QAudioBuffer*>(ptr)->sampleCount();
}

void QAudioBuffer_DestroyQAudioBuffer(void* ptr){
	static_cast<QAudioBuffer*>(ptr)->~QAudioBuffer();
}

