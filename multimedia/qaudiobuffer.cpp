#include "qaudiobuffer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioBuffer>
#include "_cgo_export.h"

class MyQAudioBuffer: public QAudioBuffer {
public:
};

QtObjectPtr QAudioBuffer_NewQAudioBuffer(){
	return new QAudioBuffer();
}

QtObjectPtr QAudioBuffer_NewQAudioBuffer3(QtObjectPtr other){
	return new QAudioBuffer(*static_cast<QAudioBuffer*>(other));
}

int QAudioBuffer_ByteCount(QtObjectPtr ptr){
	return static_cast<QAudioBuffer*>(ptr)->byteCount();
}

void QAudioBuffer_ConstData(QtObjectPtr ptr){
	static_cast<QAudioBuffer*>(ptr)->constData();
}

void QAudioBuffer_Data2(QtObjectPtr ptr){
	static_cast<QAudioBuffer*>(ptr)->data();
}

void QAudioBuffer_Data(QtObjectPtr ptr){
	static_cast<QAudioBuffer*>(ptr)->data();
}

int QAudioBuffer_FrameCount(QtObjectPtr ptr){
	return static_cast<QAudioBuffer*>(ptr)->frameCount();
}

int QAudioBuffer_IsValid(QtObjectPtr ptr){
	return static_cast<QAudioBuffer*>(ptr)->isValid();
}

int QAudioBuffer_SampleCount(QtObjectPtr ptr){
	return static_cast<QAudioBuffer*>(ptr)->sampleCount();
}

void QAudioBuffer_DestroyQAudioBuffer(QtObjectPtr ptr){
	static_cast<QAudioBuffer*>(ptr)->~QAudioBuffer();
}

