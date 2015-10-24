#include "qmediarecordercontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaRecorder>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QMediaRecorderControl>
#include "_cgo_export.h"

class MyQMediaRecorderControl: public QMediaRecorderControl {
public:
void Signal_ActualLocationChanged(const QUrl & location){callbackQMediaRecorderControlActualLocationChanged(this->objectName().toUtf8().data(), location.toString().toUtf8().data());};
void Signal_Error(int error, const QString & errorString){callbackQMediaRecorderControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data());};
void Signal_MutedChanged(bool muted){callbackQMediaRecorderControlMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_StateChanged(QMediaRecorder::State state){callbackQMediaRecorderControlStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QMediaRecorder::Status status){callbackQMediaRecorderControlStatusChanged(this->objectName().toUtf8().data(), status);};
};

void QMediaRecorderControl_ConnectActualLocationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(const QUrl &)>(&QMediaRecorderControl::actualLocationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(const QUrl &)>(&MyQMediaRecorderControl::Signal_ActualLocationChanged));;
}

void QMediaRecorderControl_DisconnectActualLocationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(const QUrl &)>(&QMediaRecorderControl::actualLocationChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(const QUrl &)>(&MyQMediaRecorderControl::Signal_ActualLocationChanged));;
}

void QMediaRecorderControl_ApplySettings(QtObjectPtr ptr){
	static_cast<QMediaRecorderControl*>(ptr)->applySettings();
}

void QMediaRecorderControl_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

void QMediaRecorderControl_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

int QMediaRecorderControl_IsMuted(QtObjectPtr ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->isMuted();
}

void QMediaRecorderControl_ConnectMutedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

void QMediaRecorderControl_DisconnectMutedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

char* QMediaRecorderControl_OutputLocation(QtObjectPtr ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->outputLocation().toString().toUtf8().data();
}

void QMediaRecorderControl_SetMuted(QtObjectPtr ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorderControl_SetOutputLocation(QtObjectPtr ptr, char* location){
	return static_cast<QMediaRecorderControl*>(ptr)->setOutputLocation(QUrl(QString(location)));
}

void QMediaRecorderControl_SetState(QtObjectPtr ptr, int state){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setState", Q_ARG(QMediaRecorder::State, static_cast<QMediaRecorder::State>(state)));
}

void QMediaRecorderControl_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

void QMediaRecorderControl_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

int QMediaRecorderControl_Status(QtObjectPtr ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->status();
}

void QMediaRecorderControl_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

void QMediaRecorderControl_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

void QMediaRecorderControl_DestroyQMediaRecorderControl(QtObjectPtr ptr){
	static_cast<QMediaRecorderControl*>(ptr)->~QMediaRecorderControl();
}

