#include "qvideoprobe.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaRecorder>
#include <QMediaObject>
#include <QObject>
#include <QString>
#include <QVideoProbe>
#include "_cgo_export.h"

class MyQVideoProbe: public QVideoProbe {
public:
void Signal_Flush(){callbackQVideoProbeFlush(this->objectName().toUtf8().data());};
};

void* QVideoProbe_NewQVideoProbe(void* parent){
	return new QVideoProbe(static_cast<QObject*>(parent));
}

void QVideoProbe_ConnectFlush(void* ptr){
	QObject::connect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

void QVideoProbe_DisconnectFlush(void* ptr){
	QObject::disconnect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

int QVideoProbe_IsActive(void* ptr){
	return static_cast<QVideoProbe*>(ptr)->isActive();
}

int QVideoProbe_SetSource(void* ptr, void* source){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QVideoProbe_SetSource2(void* ptr, void* mediaRecorder){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QVideoProbe_DestroyQVideoProbe(void* ptr){
	static_cast<QVideoProbe*>(ptr)->~QVideoProbe();
}

