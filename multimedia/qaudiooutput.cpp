#include "qaudiooutput.h"
#include <QModelIndex>
#include <QIODevice>
#include <QAudioFormat>
#include <QAudioDeviceInfo>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioOutput>
#include "_cgo_export.h"

class MyQAudioOutput: public QAudioOutput {
public:
void Signal_Notify(){callbackQAudioOutputNotify(this->objectName().toUtf8().data());};
};

void* QAudioOutput_NewQAudioOutput2(void* audioDevice, void* format, void* parent){
	return new QAudioOutput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioOutput_NewQAudioOutput(void* format, void* parent){
	return new QAudioOutput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioOutput_BufferSize(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->bufferSize();
}

int QAudioOutput_BytesFree(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->bytesFree();
}

char* QAudioOutput_Category(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->category().toUtf8().data();
}

void QAudioOutput_ConnectNotify(void* ptr){
	QObject::connect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

void QAudioOutput_DisconnectNotify(void* ptr){
	QObject::disconnect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

int QAudioOutput_NotifyInterval(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->notifyInterval();
}

int QAudioOutput_PeriodSize(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->periodSize();
}

void QAudioOutput_Reset(void* ptr){
	static_cast<QAudioOutput*>(ptr)->reset();
}

void QAudioOutput_Resume(void* ptr){
	static_cast<QAudioOutput*>(ptr)->resume();
}

void QAudioOutput_SetBufferSize(void* ptr, int value){
	static_cast<QAudioOutput*>(ptr)->setBufferSize(value);
}

void QAudioOutput_SetCategory(void* ptr, char* category){
	static_cast<QAudioOutput*>(ptr)->setCategory(QString(category));
}

void QAudioOutput_SetNotifyInterval(void* ptr, int ms){
	static_cast<QAudioOutput*>(ptr)->setNotifyInterval(ms);
}

void QAudioOutput_SetVolume(void* ptr, double volume){
	static_cast<QAudioOutput*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void* QAudioOutput_Start2(void* ptr){
	return static_cast<QAudioOutput*>(ptr)->start();
}

void QAudioOutput_Start(void* ptr, void* device){
	static_cast<QAudioOutput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioOutput_Stop(void* ptr){
	static_cast<QAudioOutput*>(ptr)->stop();
}

void QAudioOutput_Suspend(void* ptr){
	static_cast<QAudioOutput*>(ptr)->suspend();
}

double QAudioOutput_Volume(void* ptr){
	return static_cast<double>(static_cast<QAudioOutput*>(ptr)->volume());
}

void QAudioOutput_DestroyQAudioOutput(void* ptr){
	static_cast<QAudioOutput*>(ptr)->~QAudioOutput();
}

