#include "qaudiooutput.h"
#include <QModelIndex>
#include <QAudioDeviceInfo>
#include <QObject>
#include <QAudioFormat>
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioOutput>
#include "_cgo_export.h"

class MyQAudioOutput: public QAudioOutput {
public:
void Signal_Notify(){callbackQAudioOutputNotify(this->objectName().toUtf8().data());};
};

QtObjectPtr QAudioOutput_NewQAudioOutput2(QtObjectPtr audioDevice, QtObjectPtr format, QtObjectPtr parent){
	return new QAudioOutput(*static_cast<QAudioDeviceInfo*>(audioDevice), *static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

QtObjectPtr QAudioOutput_NewQAudioOutput(QtObjectPtr format, QtObjectPtr parent){
	return new QAudioOutput(*static_cast<QAudioFormat*>(format), static_cast<QObject*>(parent));
}

int QAudioOutput_BufferSize(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->bufferSize();
}

int QAudioOutput_BytesFree(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->bytesFree();
}

char* QAudioOutput_Category(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->category().toUtf8().data();
}

void QAudioOutput_ConnectNotify(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

void QAudioOutput_DisconnectNotify(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioOutput*>(ptr), static_cast<void (QAudioOutput::*)()>(&QAudioOutput::notify), static_cast<MyQAudioOutput*>(ptr), static_cast<void (MyQAudioOutput::*)()>(&MyQAudioOutput::Signal_Notify));;
}

int QAudioOutput_NotifyInterval(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->notifyInterval();
}

int QAudioOutput_PeriodSize(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->periodSize();
}

void QAudioOutput_Reset(QtObjectPtr ptr){
	static_cast<QAudioOutput*>(ptr)->reset();
}

void QAudioOutput_Resume(QtObjectPtr ptr){
	static_cast<QAudioOutput*>(ptr)->resume();
}

void QAudioOutput_SetBufferSize(QtObjectPtr ptr, int value){
	static_cast<QAudioOutput*>(ptr)->setBufferSize(value);
}

void QAudioOutput_SetCategory(QtObjectPtr ptr, char* category){
	static_cast<QAudioOutput*>(ptr)->setCategory(QString(category));
}

void QAudioOutput_SetNotifyInterval(QtObjectPtr ptr, int ms){
	static_cast<QAudioOutput*>(ptr)->setNotifyInterval(ms);
}

QtObjectPtr QAudioOutput_Start2(QtObjectPtr ptr){
	return static_cast<QAudioOutput*>(ptr)->start();
}

void QAudioOutput_Start(QtObjectPtr ptr, QtObjectPtr device){
	static_cast<QAudioOutput*>(ptr)->start(static_cast<QIODevice*>(device));
}

void QAudioOutput_Stop(QtObjectPtr ptr){
	static_cast<QAudioOutput*>(ptr)->stop();
}

void QAudioOutput_Suspend(QtObjectPtr ptr){
	static_cast<QAudioOutput*>(ptr)->suspend();
}

void QAudioOutput_DestroyQAudioOutput(QtObjectPtr ptr){
	static_cast<QAudioOutput*>(ptr)->~QAudioOutput();
}

