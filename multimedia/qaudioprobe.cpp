#include "qaudioprobe.h"
#include <QUrl>
#include <QModelIndex>
#include <QMediaObject>
#include <QMediaRecorder>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QAudioProbe>
#include "_cgo_export.h"

class MyQAudioProbe: public QAudioProbe {
public:
void Signal_Flush(){callbackQAudioProbeFlush(this->objectName().toUtf8().data());};
};

void* QAudioProbe_NewQAudioProbe(void* parent){
	return new QAudioProbe(static_cast<QObject*>(parent));
}

void QAudioProbe_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

void QAudioProbe_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

int QAudioProbe_IsActive(void* ptr){
	return static_cast<QAudioProbe*>(ptr)->isActive();
}

int QAudioProbe_SetSource(void* ptr, void* source){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QAudioProbe_SetSource2(void* ptr, void* mediaRecorder){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QAudioProbe_DestroyQAudioProbe(void* ptr){
	static_cast<QAudioProbe*>(ptr)->~QAudioProbe();
}

