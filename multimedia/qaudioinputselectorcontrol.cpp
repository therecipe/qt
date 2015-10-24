#include "qaudioinputselectorcontrol.h"
#include <QModelIndex>
#include <QAudioInput>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAudioInputSelectorControl>
#include "_cgo_export.h"

class MyQAudioInputSelectorControl: public QAudioInputSelectorControl {
public:
void Signal_ActiveInputChanged(const QString & name){callbackQAudioInputSelectorControlActiveInputChanged(this->objectName().toUtf8().data(), name.toUtf8().data());};
void Signal_AvailableInputsChanged(){callbackQAudioInputSelectorControlAvailableInputsChanged(this->objectName().toUtf8().data());};
};

char* QAudioInputSelectorControl_ActiveInput(QtObjectPtr ptr){
	return static_cast<QAudioInputSelectorControl*>(ptr)->activeInput().toUtf8().data();
}

void QAudioInputSelectorControl_ConnectActiveInputChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));;
}

void QAudioInputSelectorControl_DisconnectActiveInputChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)(const QString &)>(&QAudioInputSelectorControl::activeInputChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)(const QString &)>(&MyQAudioInputSelectorControl::Signal_ActiveInputChanged));;
}

void QAudioInputSelectorControl_ConnectAvailableInputsChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));;
}

void QAudioInputSelectorControl_DisconnectAvailableInputsChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioInputSelectorControl*>(ptr), static_cast<void (QAudioInputSelectorControl::*)()>(&QAudioInputSelectorControl::availableInputsChanged), static_cast<MyQAudioInputSelectorControl*>(ptr), static_cast<void (MyQAudioInputSelectorControl::*)()>(&MyQAudioInputSelectorControl::Signal_AvailableInputsChanged));;
}

char* QAudioInputSelectorControl_DefaultInput(QtObjectPtr ptr){
	return static_cast<QAudioInputSelectorControl*>(ptr)->defaultInput().toUtf8().data();
}

char* QAudioInputSelectorControl_InputDescription(QtObjectPtr ptr, char* name){
	return static_cast<QAudioInputSelectorControl*>(ptr)->inputDescription(QString(name)).toUtf8().data();
}

void QAudioInputSelectorControl_SetActiveInput(QtObjectPtr ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QAudioInputSelectorControl*>(ptr), "setActiveInput", Q_ARG(QString, QString(name)));
}

void QAudioInputSelectorControl_DestroyQAudioInputSelectorControl(QtObjectPtr ptr){
	static_cast<QAudioInputSelectorControl*>(ptr)->~QAudioInputSelectorControl();
}

