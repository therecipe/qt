#include "qmediarecordercontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QMediaRecorder>
#include <QString>
#include <QMediaRecorderControl>
#include "_cgo_export.h"

class MyQMediaRecorderControl: public QMediaRecorderControl {
public:
void Signal_Error(int error, const QString & errorString){callbackQMediaRecorderControlError(this->objectName().toUtf8().data(), error, errorString.toUtf8().data());};
void Signal_MutedChanged(bool muted){callbackQMediaRecorderControlMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_StateChanged(QMediaRecorder::State state){callbackQMediaRecorderControlStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QMediaRecorder::Status status){callbackQMediaRecorderControlStatusChanged(this->objectName().toUtf8().data(), status);};
};

void QMediaRecorderControl_ApplySettings(void* ptr){
	static_cast<QMediaRecorderControl*>(ptr)->applySettings();
}

void QMediaRecorderControl_ConnectError(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

void QMediaRecorderControl_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(int, const QString &)>(&QMediaRecorderControl::error), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(int, const QString &)>(&MyQMediaRecorderControl::Signal_Error));;
}

int QMediaRecorderControl_IsMuted(void* ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->isMuted();
}

void QMediaRecorderControl_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

void QMediaRecorderControl_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(bool)>(&QMediaRecorderControl::mutedChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(bool)>(&MyQMediaRecorderControl::Signal_MutedChanged));;
}

void QMediaRecorderControl_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorderControl_SetOutputLocation(void* ptr, void* location){
	return static_cast<QMediaRecorderControl*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorderControl_SetState(void* ptr, int state){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setState", Q_ARG(QMediaRecorder::State, static_cast<QMediaRecorder::State>(state)));
}

void QMediaRecorderControl_SetVolume(void* ptr, double gain){
	QMetaObject::invokeMethod(static_cast<QMediaRecorderControl*>(ptr), "setVolume", Q_ARG(qreal, static_cast<qreal>(gain)));
}

void QMediaRecorderControl_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

void QMediaRecorderControl_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::State)>(&QMediaRecorderControl::stateChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::State)>(&MyQMediaRecorderControl::Signal_StateChanged));;
}

int QMediaRecorderControl_Status(void* ptr){
	return static_cast<QMediaRecorderControl*>(ptr)->status();
}

void QMediaRecorderControl_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

void QMediaRecorderControl_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorderControl*>(ptr), static_cast<void (QMediaRecorderControl::*)(QMediaRecorder::Status)>(&QMediaRecorderControl::statusChanged), static_cast<MyQMediaRecorderControl*>(ptr), static_cast<void (MyQMediaRecorderControl::*)(QMediaRecorder::Status)>(&MyQMediaRecorderControl::Signal_StatusChanged));;
}

double QMediaRecorderControl_Volume(void* ptr){
	return static_cast<double>(static_cast<QMediaRecorderControl*>(ptr)->volume());
}

void QMediaRecorderControl_DestroyQMediaRecorderControl(void* ptr){
	static_cast<QMediaRecorderControl*>(ptr)->~QMediaRecorderControl();
}

