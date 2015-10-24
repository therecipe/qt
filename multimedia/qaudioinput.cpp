#include "qaudioinput.h"
#include <QModelIndex>
#include <QIODevice>
#include <QAudioDeviceInfo>
#include <QObject>
#include <QAudioFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioInput>
#include "_cgo_export.h"

class MyQAudioInput: public QAudioInput {
public:
void Signal_Notify(){callbackQAudioInputNotify(this->objectName().toUtf8().data());};
};

QtObjectPtr QAudioInput_NewQAudioInput2(QtObjectPtr audioDevice, QtObjectPtr format, QtObjectPtr parent){
	return new QAudioInput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

QtObjectPtr QAudioInput_NewQAudioInput(QtObjectPtr format, QtObjectPtr parent){
	return new QAudioInput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioInput_BufferSize(QtObjectPtr ptr){
	return static_cast<QAudioInput*>(ptr)->bufferSize();
}

int QAudioInput_BytesReady(QtObjectPtr ptr){
	return static_cast<QAudioInput*>(ptr)->bytesReady();
}

void QAudioInput_ConnectNotify(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

void QAudioInput_DisconnectNotify(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioInput*>(ptr), static_cast<void (QAudioInput::*)()>(&QAudioInput::notify), static_cast<MyQAudioInput*>(ptr), static_cast<void (MyQAudioInput::*)()>(&MyQAudioInput::Signal_Notify));;
}

int QAudioInput_NotifyInterval(QtObjectPtr ptr){
	return static_cast<QAudioInput*>(ptr)->notifyInterval();
}

int QAudioInput_PeriodSize(QtObjectPtr ptr){
	return static_cast<QAudioInput*>(ptr)->periodSize();
}

void QAudioInput_Reset(QtObjectPtr ptr){
	static_cast<QAudioInput*>(ptr)->reset();
}

void QAudioInput_Resume(QtObjectPtr ptr){
	static_cast<QAudioInput*>(ptr)->resume();
}

void QAudioInput_SetBufferSize(QtObjectPtr ptr, int value){
	static_cast<QAudioInput*>(ptr)->setBufferSize(value);
}

void QAudioInput_SetNotifyInterval(QtObjectPtr ptr, int ms){
	static_cast<QAudioInput*>(ptr)->setNotifyInterval(ms);
}

QtObjectPtr QAudioInput_Start2(QtObjectPtr ptr){
	return static_cast<QAudioInput*>(ptr)->start();
}

void QAudioInput_Start(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QAudioInput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioInput_Stop(QtObjectPtr ptr){
	static_cast<QAudioInput*>(ptr)->stop();
}

void QAudioInput_Suspend(QtObjectPtr ptr){
	static_cast<QAudioInput*>(ptr)->suspend();
}

void QAudioInput_DestroyQAudioInput(QtObjectPtr ptr){
	static_cast<QAudioInput*>(ptr)->~QAudioInput();
}

