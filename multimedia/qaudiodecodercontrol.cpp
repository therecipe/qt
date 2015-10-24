#include "qaudiodecodercontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAudioDecoder>
#include <QAudioFormat>
#include <QIODevice>
#include <QString>
#include <QVariant>
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

int QAudioDecoderControl_BufferAvailable(QtObjectPtr ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->bufferAvailable();
}

void QAudioDecoderControl_ConnectBufferAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_DisconnectBufferAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(bool)>(&QAudioDecoderControl::bufferAvailableChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(bool)>(&MyQAudioDecoderControl::Signal_BufferAvailableChanged));;
}

void QAudioDecoderControl_ConnectBufferReady(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_DisconnectBufferReady(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::bufferReady), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_BufferReady));;
}

void QAudioDecoderControl_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(int, const QString &)>(&QAudioDecoderControl::error), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(int, const QString &)>(&MyQAudioDecoderControl::Signal_Error));;
}

void QAudioDecoderControl_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::finished), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_Finished));;
}

void QAudioDecoderControl_SetAudioFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QAudioDecoderControl*>(ptr)->setAudioFormat(*static_cast<QAudioFormat*>(format));
}

void QAudioDecoderControl_SetSourceDevice(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceDevice(static_cast<QIODevice*>(device));
}

void QAudioDecoderControl_SetSourceFilename(QtObjectPtr ptr, char* fileName){
	static_cast<QAudioDecoderControl*>(ptr)->setSourceFilename(QString(fileName));
}

void QAudioDecoderControl_ConnectSourceChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

void QAudioDecoderControl_DisconnectSourceChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)()>(&QAudioDecoderControl::sourceChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)()>(&MyQAudioDecoderControl::Signal_SourceChanged));;
}

QtObjectPtr QAudioDecoderControl_SourceDevice(QtObjectPtr ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceDevice();
}

char* QAudioDecoderControl_SourceFilename(QtObjectPtr ptr){
	return static_cast<QAudioDecoderControl*>(ptr)->sourceFilename().toUtf8().data();
}

void QAudioDecoderControl_Start(QtObjectPtr ptr){
	static_cast<QAudioDecoderControl*>(ptr)->start();
}

void QAudioDecoderControl_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioDecoderControl*>(ptr), static_cast<void (QAudioDecoderControl::*)(QAudioDecoder::State)>(&QAudioDecoderControl::stateChanged), static_cast<MyQAudioDecoderControl*>(ptr), static_cast<void (MyQAudioDecoderControl::*)(QAudioDecoder::State)>(&MyQAudioDecoderControl::Signal_StateChanged));;
}

void QAudioDecoderControl_Stop(QtObjectPtr ptr){
	static_cast<QAudioDecoderControl*>(ptr)->stop();
}

void QAudioDecoderControl_DestroyQAudioDecoderControl(QtObjectPtr ptr){
	static_cast<QAudioDecoderControl*>(ptr)->~QAudioDecoderControl();
}

