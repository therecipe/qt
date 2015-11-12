#include "qmediarecorder.h"
#include <QAudioEncoderSettings>
#include <QMetaObject>
#include <QString>
#include <QUrl>
#include <QMediaObject>
#include <QVideoEncoderSettings>
#include <QObject>
#include <QVariant>
#include <QModelIndex>
#include <QMediaRecorder>
#include "_cgo_export.h"

class MyQMediaRecorder: public QMediaRecorder {
public:
void Signal_AvailabilityChanged(bool available){callbackQMediaRecorderAvailabilityChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataAvailableChanged(bool available){callbackQMediaRecorderMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMediaRecorderMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_MetaDataWritableChanged(bool writable){callbackQMediaRecorderMetaDataWritableChanged(this->objectName().toUtf8().data(), writable);};
void Signal_MutedChanged(bool muted){callbackQMediaRecorderMutedChanged(this->objectName().toUtf8().data(), muted);};
void Signal_StateChanged(QMediaRecorder::State state){callbackQMediaRecorderStateChanged(this->objectName().toUtf8().data(), state);};
void Signal_StatusChanged(QMediaRecorder::Status status){callbackQMediaRecorderStatusChanged(this->objectName().toUtf8().data(), status);};
};

int QMediaRecorder_IsMetaDataAvailable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataAvailable();
}

int QMediaRecorder_IsMetaDataWritable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMetaDataWritable();
}

int QMediaRecorder_IsMuted(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isMuted();
}

void QMediaRecorder_SetMuted(void* ptr, int muted){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setMuted", Q_ARG(bool, muted != 0));
}

int QMediaRecorder_SetOutputLocation(void* ptr, void* location){
	return static_cast<QMediaRecorder*>(ptr)->setOutputLocation(*static_cast<QUrl*>(location));
}

void QMediaRecorder_SetVolume(void* ptr, double volume){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "setVolume", Q_ARG(qreal, static_cast<qreal>(volume)));
}

double QMediaRecorder_Volume(void* ptr){
	return static_cast<double>(static_cast<QMediaRecorder*>(ptr)->volume());
}

void* QMediaRecorder_NewQMediaRecorder(void* mediaObject, void* parent){
	return new QMediaRecorder(static_cast<QMediaObject*>(mediaObject), static_cast<QObject*>(parent));
}

char* QMediaRecorder_AudioCodecDescription(void* ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->audioCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_ConnectAvailabilityChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

void QMediaRecorder_DisconnectAvailabilityChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::availabilityChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_AvailabilityChanged));;
}

char* QMediaRecorder_AvailableMetaData(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

char* QMediaRecorder_ContainerDescription(void* ptr, char* format){
	return static_cast<QMediaRecorder*>(ptr)->containerDescription(QString(format)).toUtf8().data();
}

char* QMediaRecorder_ContainerFormat(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->containerFormat().toUtf8().data();
}

int QMediaRecorder_Error(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->error();
}

char* QMediaRecorder_ErrorString(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->errorString().toUtf8().data();
}

int QMediaRecorder_IsAvailable(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->isAvailable();
}

void* QMediaRecorder_MediaObject(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->mediaObject();
}

void* QMediaRecorder_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMediaRecorder*>(ptr)->metaData(QString(key)));
}

void QMediaRecorder_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataAvailableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataAvailableChanged));;
}

void QMediaRecorder_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)()>(&QMediaRecorder::metaDataChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)()>(&MyQMediaRecorder::Signal_MetaDataChanged));;
}

void QMediaRecorder_ConnectMetaDataWritableChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_DisconnectMetaDataWritableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::metaDataWritableChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MetaDataWritableChanged));;
}

void QMediaRecorder_ConnectMutedChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_DisconnectMutedChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(bool)>(&QMediaRecorder::mutedChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(bool)>(&MyQMediaRecorder::Signal_MutedChanged));;
}

void QMediaRecorder_Pause(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "pause");
}

void QMediaRecorder_Record(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "record");
}

void QMediaRecorder_SetAudioSettings(void* ptr, void* settings){
	static_cast<QMediaRecorder*>(ptr)->setAudioSettings(*static_cast<QAudioEncoderSettings*>(settings));
}

void QMediaRecorder_SetContainerFormat(void* ptr, char* container){
	static_cast<QMediaRecorder*>(ptr)->setContainerFormat(QString(container));
}

void QMediaRecorder_SetEncodingSettings(void* ptr, void* audio, void* video, char* container){
	static_cast<QMediaRecorder*>(ptr)->setEncodingSettings(*static_cast<QAudioEncoderSettings*>(audio), *static_cast<QVideoEncoderSettings*>(video), QString(container));
}

void QMediaRecorder_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QMediaRecorder*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMediaRecorder_SetVideoSettings(void* ptr, void* settings){
	static_cast<QMediaRecorder*>(ptr)->setVideoSettings(*static_cast<QVideoEncoderSettings*>(settings));
}

void QMediaRecorder_ConnectStateChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

void QMediaRecorder_DisconnectStateChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::State)>(&QMediaRecorder::stateChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::State)>(&MyQMediaRecorder::Signal_StateChanged));;
}

int QMediaRecorder_Status(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->status();
}

void QMediaRecorder_ConnectStatusChanged(void* ptr){
	QObject::connect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_DisconnectStatusChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaRecorder*>(ptr), static_cast<void (QMediaRecorder::*)(QMediaRecorder::Status)>(&QMediaRecorder::statusChanged), static_cast<MyQMediaRecorder*>(ptr), static_cast<void (MyQMediaRecorder::*)(QMediaRecorder::Status)>(&MyQMediaRecorder::Signal_StatusChanged));;
}

void QMediaRecorder_Stop(void* ptr){
	QMetaObject::invokeMethod(static_cast<QMediaRecorder*>(ptr), "stop");
}

char* QMediaRecorder_SupportedAudioCodecs(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedAudioCodecs().join("|").toUtf8().data();
}

char* QMediaRecorder_SupportedContainers(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedContainers().join("|").toUtf8().data();
}

char* QMediaRecorder_SupportedVideoCodecs(void* ptr){
	return static_cast<QMediaRecorder*>(ptr)->supportedVideoCodecs().join("|").toUtf8().data();
}

char* QMediaRecorder_VideoCodecDescription(void* ptr, char* codec){
	return static_cast<QMediaRecorder*>(ptr)->videoCodecDescription(QString(codec)).toUtf8().data();
}

void QMediaRecorder_DestroyQMediaRecorder(void* ptr){
	static_cast<QMediaRecorder*>(ptr)->~QMediaRecorder();
}

