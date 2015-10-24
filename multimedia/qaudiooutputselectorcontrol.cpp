#include "qaudiooutputselectorcontrol.h"
#include <QObject>
#include <QAudioOutput>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QAudioOutputSelectorControl>
#include "_cgo_export.h"

class MyQAudioOutputSelectorControl: public QAudioOutputSelectorControl {
public:
void Signal_ActiveOutputChanged(const QString & name){callbackQAudioOutputSelectorControlActiveOutputChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_AvailableOutputsChanged(){callbackQAudioOutputSelectorControlAvailableOutputsChanged(this->objectName().toUtf8().data());};
};

char* QAudioOutputSelectorControl_ActiveOutput(QtObjectPtr ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->activeOutput().toUtf8().data();
}

void QAudioOutputSelectorControl_ConnectActiveOutputChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_DisconnectActiveOutputChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)(const QString &)>(&QAudioOutputSelectorControl::activeOutputChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)(const QString &)>(&MyQAudioOutputSelectorControl::Signal_ActiveOutputChanged));;
}

void QAudioOutputSelectorControl_ConnectAvailableOutputsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

void QAudioOutputSelectorControl_DisconnectAvailableOutputsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioOutputSelectorControl*>(ptr), static_cast<void (QAudioOutputSelectorControl::*)()>(&QAudioOutputSelectorControl::availableOutputsChanged), static_cast<MyQAudioOutputSelectorControl*>(ptr), static_cast<void (MyQAudioOutputSelectorControl::*)()>(&MyQAudioOutputSelectorControl::Signal_AvailableOutputsChanged));;
}

char* QAudioOutputSelectorControl_DefaultOutput(QtObjectPtr ptr){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->defaultOutput().toUtf8().data();
}

char* QAudioOutputSelectorControl_OutputDescription(QtObjectPtr ptr, char* name){
	return static_cast<QAudioOutputSelectorControl*>(ptr)->outputDescription(QString(name)).toUtf8().data();
}

void QAudioOutputSelectorControl_SetActiveOutput(QtObjectPtr ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioOutputSelectorControl*>(ptr), "setActiveOutput", Q_ARG(QString, QString(name)));
}

void QAudioOutputSelectorControl_DestroyQAudioOutputSelectorControl(QtObjectPtr ptr){
	static_cast<QAudioOutputSelectorControl*>(ptr)->~QAudioOutputSelectorControl();
}

