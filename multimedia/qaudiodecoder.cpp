#include "qaudiodecoder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QAudioFormat>
#include <QIODevice>
#include <QAudioDecoder>
#include "_cgo_export.h"

class MyQAudioDecoder: public QAudioDecoder {
public:
void Signal_BufferAvailableChanged(bool available){callbackQAudioDecoderBufferAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_BufferReady(){callbackQAudioDecoderBufferReady(this->objectName().toUtf8().data());};
void Signal_Finished(){callbackQAudioDecoderFinished(this->objectName().toUtf8().data());};
void Signal_SourceChanged(){callbackQAudioDecoderSourceChanged(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAudioDecoder::State state){callbackQAudioDecoderStateChanged(this->objectName().toUtf8().data(), state);};
};

char* QAudioDecoder_ErrorString(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->errorString().toUtf8().data();
}

void* QAudioDecoder_NewQAudioDecoder(void* parent){
	return new QAudioDecoder(static_cast<QObject*>(parent));
}

int QAudioDecoder_BufferAvailable(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->bufferAvailable();
}

void QAudioDecoder_ConnectBufferAvailableChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_DisconnectBufferAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_ConnectBufferReady(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

void QAudioDecoder_DisconnectBufferReady(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

int QAudioDecoder_Error(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->error();
}

void QAudioDecoder_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_SetAudioFormat(void* ptr, void* format){
	static_cast<QAudioDecoder*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoder_SetSourceDevice(void* ptr, void* device){
	static_cast<QAudioDecoder*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoder_SetSourceFilename(void* ptr, char* fileName){
	static_cast<QAudioDecoder*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoder_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

void QAudioDecoder_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

void* QAudioDecoder_SourceDevice(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceDevice();
}

char* QAudioDecoder_SourceFilename(void* ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoder_Start(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "start");
}

void QAudioDecoder_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "stop");
}

void QAudioDecoder_DestroyQAudioDecoder(void* ptr){
	static_cast<QAudioDecoder*>(ptr)->~QAudioDecoder();
}

