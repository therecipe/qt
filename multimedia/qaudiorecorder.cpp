#include "qaudiorecorder.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioRecorder>
#include "_cgo_export.h"

class MyQAudioRecorder: public QAudioRecorder {
public:
void Signal_AudioInputChanged(const QString & name){callbackQAudioRecorderAudioInputChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_AvailableAudioInputsChanged(){callbackQAudioRecorderAvailableAudioInputsChanged(this->objectName().toUtf8().data());};
};

void* QAudioRecorder_NewQAudioRecorder(void* parent){
	return new QAudioRecorder(static_cast<QObject*>(parent));
}

char* QAudioRecorder_AudioInput(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInput().toUtf8().data();
}

void QAudioRecorder_ConnectAudioInputChanged(void* ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

void QAudioRecorder_DisconnectAudioInputChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

char* QAudioRecorder_AudioInputDescription(void* ptr, char* name){
	return static_cast<QAudioRecorder*>(ptr)->audioInputDescription(QString(name)).toUtf8().data();
}

char* QAudioRecorder_AudioInputs(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInputs().join("|").toUtf8().data();
}

void QAudioRecorder_ConnectAvailableAudioInputsChanged(void* ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

void QAudioRecorder_DisconnectAvailableAudioInputsChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

char* QAudioRecorder_DefaultAudioInput(void* ptr){
	return static_cast<QAudioRecorder*>(ptr)->defaultAudioInput().toUtf8().data();
}

void QAudioRecorder_SetAudioInput(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setAudioInput", Q_ARG(QString, QString(name)));
}

void QAudioRecorder_DestroyQAudioRecorder(void* ptr){
	static_cast<QAudioRecorder*>(ptr)->~QAudioRecorder();
}

