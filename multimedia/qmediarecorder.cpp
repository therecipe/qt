#include "qmediarecorder.h"
#include <QMediaObject>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QVariant>
#include <QVideoEncoderSettings>
#include <QAudioEncoderSettings>
#include <QMediaRecorder>
#include "_cgo_export.h"

class MyQMediaRecorder: public QMediaRecorder {
public:
void Signal_ActualLocationChanged(const QUrl & location){callbackQMediaRecorderActualLocationChanged(this->objectName().toUtf8().data(), location.toString().toUtf8().data());};
void Signal_AvailabilityChanged(bool available){callbackQMediaRecorderAvailabilityChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataAvailableChanged(bool available){callbackQMediaRecorderMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMediaRecorderMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_MetaDataWritableChanged(bool writable){callbackQMediaRecorderMetaDataWritableChanged(this->objectName().toUtf8().data(), writable);};
void Signal_MutedChanged(bool muted){callbackQMediaRecorderMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_StateChanged(QMediaRecorder::State state){callbackQMediaRecorderStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QMediaRecorder::Status status){callbackQMediaRecorderStatusChanged(this->objectName().toUtf8().data(), status);};
};

char* QMediaRecorder_ActualLocation(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->actualLocation().toString().toUtf8().data();
}

int QMediaRecorder_IsMetaDataAvailable(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataAvailable();
}

int QMediaRecorder_IsMetaDataWritable(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataWritable();
}

int QMediaRecorder_IsMuted(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMuted();
}

char* QMediaRecorder_OutputLocation(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->outputLocation().toString().toUtf8().data();
}

void QMediaRecorder_SetMuted(QtObjectPtr ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorder_SetOutputLocation(QtObjectPtr ptr, char* location){
	return static_cast<QMediaRecorder*>(ptr)->setOutputLocation(QUrl(QString(location)));
}

QtObjectPtr QMediaRecorder_NewQMediaRecorder(QtObjectPtr mediaObject, QtObjectPtr parent){
	return new QMediaRecorder(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

void QMediaRecorder_ConnectActualLocationChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QUrl &)>(&QMediaRecorder::actualLocationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QUrl &)>(&MyQMediaRecorder::Signal_ActualLocationChanged));;
}

void QMediaRecorder_DisconnectActualLocationChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(const QUrl &)>(&QMediaRecorder::actualLocationChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(const QUrl &)>(&MyQMediaRecorder::Signal_ActualLocationChanged));;
}

char* QMediaRecorder_AudioCodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->audioCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_ConnectAvailabilityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

void QMediaRecorder_DisconnectAvailabilityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

char* QMediaRecorder_AvailableMetaData(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

char* QMediaRecorder_ContainerDescription(QtObjectPtr ptr, char* format){
	return static_cast<QMediaRecorder*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaRecorder_ContainerFormat(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->containerFormat().toUtf8().data();
}

int QMediaRecorder_Error(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->error();
}

char* QMediaRecorder_ErrorString(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->errorString().toUtf8().data();
}

int QMediaRecorder_IsAvailable(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->isAvailable();
}

QtObjectPtr QMediaRecorder_MediaObject(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->mediaObject();
}

char* QMediaRecorder_MetaData(QtObjectPtr ptr, char* key){
	return static_cast<QMediaRecorder*>(ptr)->metaData(QString(key)).toString().toUtf8().data();
}

void QMediaRecorder_ConnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_ConnectMetaDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_DisconnectMetaDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_ConnectMetaDataWritableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_DisconnectMetaDataWritableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_ConnectMutedChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_DisconnectMutedChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_Pause(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "pause");
}

void QMediaRecorder_Record(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "record");
}

void QMediaRecorder_SetAudioSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QMediaRecorder*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

void QMediaRecorder_SetContainerFormat(QtObjectPtr ptr, char* container){
	static_cast<QMediaRecorder*>(ptr)->setContainerFormat(QString(container));
}

void QMediaRecorder_SetEncodingSettings(QtObjectPtr ptr, QtObjectPtr audio, QtObjectPtr video, char* container){
	static_cast<QMediaRecorder*>(ptr)->setEncodingSettings(*static_cast<QAudioEncoderSettings*>(audio), *static_cast<QVideoEncoderSettings*>(video), QString(container));
}

void QMediaRecorder_SetMetaData(QtObjectPtr ptr, char* key, char* value){
	static_cast<QMediaRecorder*>(ptr)->setMetaData(QString(key), QVariant(value));
}

void QMediaRecorder_SetVideoSettings(QtObjectPtr ptr, QtObjectPtr settings){
	static_cast<QMediaRecorder*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

void QMediaRecorder_ConnectStateChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

void QMediaRecorder_DisconnectStateChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

int QMediaRecorder_Status(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->status();
}

void QMediaRecorder_ConnectStatusChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_DisconnectStatusChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_Stop(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "stop");
}

char* QMediaRecorder_SupportedAudioCodecs(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedAudioCodecs().join("|").toUtf8().data();
}

char* QMediaRecorder_SupportedContainers(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedContainers().join("|").toUtf8().data();
}

char* QMediaRecorder_SupportedVideoCodecs(QtObjectPtr ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedVideoCodecs().join("|").toUtf8().data();
}

char* QMediaRecorder_VideoCodecDescription(QtObjectPtr ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->videoCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_DestroyQMediaRecorder(QtObjectPtr ptr){
	static_cast<QMediaRecorder*>(ptr)->~QMediaRecorder();
}

