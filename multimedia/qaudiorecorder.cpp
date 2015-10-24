#include "qaudiorecorder.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QAudioRecorder>
#include "_cgo_export.h"

class MyQAudioRecorder: public QAudioRecorder {
public:
void Signal_AudioInputChanged(const QString & name){callbackQAudioRecorderAudioInputChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_AvailableAudioInputsChanged(){callbackQAudioRecorderAvailableAudioInputsChanged(this->objectName().toUtf8().data());};
};

QtObjectPtr QAudioRecorder_NewQAudioRecorder(QtObjectPtr parent){
	return new QAudioRecorder(static_cast<QObject*>(parent));
}

char* QAudioRecorder_AudioInput(QtObjectPtr ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInput().toUtf8().data();
}

void QAudioRecorder_ConnectAudioInputChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

void QAudioRecorder_DisconnectAudioInputChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)(const QString &)>(&QAudioRecorder::audioInputChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)(const QString &)>(&MyQAudioRecorder::Signal_AudioInputChanged));;
}

char* QAudioRecorder_AudioInputDescription(QtObjectPtr ptr, char* name){
	return static_cast<QAudioRecorder*>(ptr)->audioInputDescription(QString(name)).toUtf8().data();
}

char* QAudioRecorder_AudioInputs(QtObjectPtr ptr){
	return static_cast<QAudioRecorder*>(ptr)->audioInputs().join("|").toUtf8().data();
}

void QAudioRecorder_ConnectAvailableAudioInputsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

void QAudioRecorder_DisconnectAvailableAudioInputsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioRecorder*>(ptr), static_cast<void (QAudioRecorder::*)()>(&QAudioRecorder::availableAudioInputsChanged), static_cast<MyQAudioRecorder*>(ptr), static_cast<void (MyQAudioRecorder::*)()>(&MyQAudioRecorder::Signal_AvailableAudioInputsChanged));;
}

char* QAudioRecorder_DefaultAudioInput(QtObjectPtr ptr){
	return static_cast<QAudioRecorder*>(ptr)->defaultAudioInput().toUtf8().data();
}

void QAudioRecorder_SetAudioInput(QtObjectPtr ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioRecorder*>(ptr), "setAudioInput", Q_ARG(QString, QString(name)));
}

void QAudioRecorder_DestroyQAudioRecorder(QtObjectPtr ptr){
	static_cast<QAudioRecorder*>(ptr)->~QAudioRecorder();
}

