#include "qaudioprobe.h"
#include <QMediaRecorder>
#include <QObject>
#include <QMediaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAudioProbe>
#include "_cgo_export.h"

class MyQAudioProbe: public QAudioProbe {
public:
void Signal_Flush(){callbackQAudioProbeFlush(this->objectName().toUtf8().data());};
};

QtObjectPtr QAudioProbe_NewQAudioProbe(QtObjectPtr parent){
	return new QAudioProbe(static_cast<QObject*>(parent));
}

void QAudioProbe_ConnectFlush(QtObjectPtr ptr){
	QObject::connect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

void QAudioProbe_DisconnectFlush(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAudioProbe*>(ptr), static_cast<void (QAudioProbe::*)()>(&QAudioProbe::flush), static_cast<MyQAudioProbe*>(ptr), static_cast<void (MyQAudioProbe::*)()>(&MyQAudioProbe::Signal_Flush));;
}

int QAudioProbe_IsActive(QtObjectPtr ptr){
	return static_cast<QAudioProbe*>(ptr)->isActive();
}

int QAudioProbe_SetSource(QtObjectPtr ptr, QtObjectPtr source){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QAudioProbe_SetSource2(QtObjectPtr ptr, QtObjectPtr mediaRecorder){
	return static_cast<QAudioProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QAudioProbe_DestroyQAudioProbe(QtObjectPtr ptr){
	static_cast<QAudioProbe*>(ptr)->~QAudioProbe();
}

