#include "qaudiodecoder.h"
#include <QMetaObject>
#include <QObject>
#include <QAudioFormat>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
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

char* QAudioDecoder_ErrorString(QtObjectPtr ptr){
	return static_cast<QAudioDecoder*>(ptr)->errorString().toUtf8().data();
}

QtObjectPtr QAudioDecoder_NewQAudioDecoder(QtObjectPtr parent){
	return new QAudioDecoder(static_cast<QObject*>(parent));
}

int QAudioDecoder_BufferAvailable(QtObjectPtr ptr){
	return static_cast<QAudioDecoder*>(ptr)->bufferAvailable();
}

void QAudioDecoder_ConnectBufferAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_DisconnectBufferAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(bool)>(&QAudioDecoder::bufferAvailableChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(bool)>(&MyQAudioDecoder::Signal_BufferAvailableChanged));;
}

void QAudioDecoder_ConnectBufferReady(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

void QAudioDecoder_DisconnectBufferReady(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::bufferReady), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_BufferReady));;
}

int QAudioDecoder_Error(QtObjectPtr ptr){
	return static_cast<QAudioDecoder*>(ptr)->error();
}

void QAudioDecoder_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::finished), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_Finished));;
}

void QAudioDecoder_SetAudioFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QAudioDecoder*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoder_SetSourceDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QAudioDecoder*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoder_SetSourceFilename(QtObjectPtr ptr, char* fileName){
	static_cast<QAudioDecoder*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoder_ConnectSourceChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

void QAudioDecoder_DisconnectSourceChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)()>(&QAudioDecoder::sourceChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)()>(&MyQAudioDecoder::Signal_SourceChanged));;
}

QtObjectPtr QAudioDecoder_SourceDevice(QtObjectPtr ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceDevice();
}

char* QAudioDecoder_SourceFilename(QtObjectPtr ptr){
	return static_cast<QAudioDecoder*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoder_Start(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "start");
}

void QAudioDecoder_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoder*>(ptr), static_cast<void (QAudioDecoder::*)(QAudioDecoder::State)>(&QAudioDecoder::stateChanged), static_cast<MyQAudioDecoder*>(ptr), static_cast<void (MyQAudioDecoder::*)(QAudioDecoder::State)>(&MyQAudioDecoder::Signal_StateChanged));;
}

void QAudioDecoder_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAudioDecoder*>(ptr), "stop");
}

void QAudioDecoder_DestroyQAudioDecoder(QtObjectPtr ptr){
	static_cast<QAudioDecoder*>(ptr)->~QAudioDecoder();
}

