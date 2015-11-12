#include "qaudiooutputselectorcontrol.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioOutput>
#include <QAudioOutputSelectorControl>
#include "_cgo_export.h"

class MyQAudioOutputSelectorControl: public QAudioOutputSelectorControl {
public:
void Signal_ActiveOutputChanged(const QString & name){callbackQAudioOutputSelectorControlActiveOutputChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_AvailableOutputsChanged(){callbackQAudioOutputSelectorControlAvailableOutputsChanged(this->objectName().toUtf8().data());};
};

char* QAudioOutputSelectorControl_ActiveOutput(void* ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->activeOutput().toUtf8().data();
}

void QAudioOutputSelectorControl_ConnectActiveOutputChanged(void* ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(void* ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(void* ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

char* QAudioOutputSelectorControl_DefaultOutput(void* ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->defaultOutput().toUtf8().data();
}

char* QAudioOutputSelectorControl_OutputDescription(void* ptr, char* name){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->outputDescription(QString(name)).toUtf8().data();
}

void QAudioOutputSelectorControl_SetActiveOutput(void* ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioOutputSelectorControl*>(ptr), "setActiveOutput", Q_ARG(QString, QString(name)));
}

void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(void* ptr){
	static_cast<QAudioOutputSelectorControl*>(ptr)->~QAudioOutputSelectorControl();
}

