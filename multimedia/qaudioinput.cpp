#include "qaudioinput.h"
#include <QAudioFormat>
#include <QAudioDeviceInfo>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QIODevice>
#include <QAudioInput>
#include "_cgo_export.h"

class MyQAudioInput: public QAudioInput {
public:
void Signal_Notify(){callbackQAudioInputNotify(this->objectName().toUtf8().data());};
};

void* QAudioInput_NewQAudioInput2(void* audioDevice, void* format, void* parent){
	return new QAudioInput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

void* QAudioInput_NewQAudioInput(void* format, void* parent){
	return new QAudioInput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioInput_BufferSize(void* ptr){
	return static_cast<QAudioInput*>(ptr)->bufferSize();
}

int QAudioInput_BytesReady(void* ptr){
	return static_cast<QAudioInput*>(ptr)->bytesReady();
}

void QAudioInput_ConnectNotify(void* ptr){
	QObject::connect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

void QAudioInput_DisconnectNotify(void* ptr){
	QObject::disconnect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

int QAudioInput_NotifyInterval(void* ptr){
	return static_cast<QAudioInput*>(ptr)->notifyInterval();
}

int QAudioInput_PeriodSize(void* ptr){
	return static_cast<QAudioInput*>(ptr)->periodSize();
}

void QAudioInput_Reset(void* ptr){
	static_cast<QAudioInput*>(ptr)->reset();
}

void QAudioInput_Resume(void* ptr){
	static_cast<QAudioInput*>(ptr)->resume();
}

void QAudioInput_SetBufferSize(void* ptr, int value){
	static_cast<QAudioInput*>(ptr)->setBufferSize(value);
}

void QAudioInput_SetNotifyInterval(void* ptr, int ms){
	static_cast<QAudioInput*>(ptr)->setNotifyInterval(ms);
}

void QAudioInput_SetVolume(void* ptr, double volume){
	static_cast<QAudioInput*>(ptr)->setVolume(static_cast<qreal>(volume));
}

void* QAudioInput_Start2(void* ptr){
	return static_cast<QAudioInput*>(ptr)->start();
}

void QAudioInput_Start(void* ptr, void* device){
	static_cast<QAudioInput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioInput_Stop(void* ptr){
	static_cast<QAudioInput*>(ptr)->stop();
}

void QAudioInput_Suspend(void* ptr){
	static_cast<QAudioInput*>(ptr)->suspend();
}

double QAudioInput_Volume(void* ptr){
	return static_cast<double>(static_cast<QAudioInput*>(ptr)->volume());
}

void QAudioInput_DestroyQAudioInput(void* ptr){
	static_cast<QAudioInput*>(ptr)->~QAudioInput();
}

