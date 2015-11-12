#include "qaudiodecodercontrol.h"
#include <QAudioDecoder>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAudioFormat>
#include <QIODevice>
#include <QAudioDecoderControl>
#include "_cgo_export.h"

class MyQAudioDecoderControl: public QAudioDecoderControl {
public:
void Signal_BufferAvailableChanged(bool available){callbackQAudioDecoderControlBufferAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_BufferReady(){callbackQAudioDecoderControlBufferReady(this->objectName().toUtf8().data());};
void Signal_Error(int error, const QString & errorString){callbackQAudioDecoderControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data());};
void Signal_Finished(){callbackQAudioDecoderControlFinished(this->objectName().toUtf8().data());};
void Signal_SourceChanged(){callbackQAudioDecoderControlSourceChanged(this->objectName().toUtf8().data());};
void Signal_StateChanged(QAudioDecoder::State state){callbackQAudioDecoderControlStateChanged(this->objectName().toUtf8().data(), state);};
};

int QAudioDecoderControl_BufferAvailable(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->bufferAvailable();
}

void QAudioDecoderControl_ConnectBufferAvailableChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_DisconnectBufferAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_ConnectBufferReady(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_DisconnectBufferReady(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_SetAudioFormat(void* ptr, void* format){
	static_cast<QAudioDecoderControl*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoderControl_SetSourceDevice(void* ptr, void* device){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoderControl_SetSourceFilename(void* ptr, char* fileName){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoderControl_ConnectSourceChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

void QAudioDecoderControl_DisconnectSourceChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

void* QAudioDecoderControl_SourceDevice(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceDevice();
}

char* QAudioDecoderControl_SourceFilename(void* ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoderControl_Start(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->start();
}

void QAudioDecoderControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_Stop(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->stop();
}

void QAudioDecoderControl_DestroyQAudioDecoderControl(void* ptr){
	static_cast<QAudioDecoderControl*>(ptr)->~QAudioDecoderControl();
}

