#include "qvideoprobe.h"
#include <QModelIndex>
#include <QMediaObject>
#include <QMediaRecorder>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QVideoProbe>
#include "_cgo_export.h"

class MyQVideoProbe: public QVideoProbe {
public:
void Signal_Flush(){callbackQVideoProbeFlush(this->objectName().toUtf8().data());};
};

QtObjectPtr QVideoProbe_NewQVideoProbe(QtObjectPtr parent){
	return new QVideoProbe(static_cast<QObject*>(parent));
}

void QVideoProbe_ConnectFlush(QtObjectPtr ptr){
	QObject::connect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

void QVideoProbe_DisconnectFlush(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QVideoProbe*>(ptr), static_cast<void (QVideoProbe::*)()>(&QVideoProbe::flush), static_cast<MyQVideoProbe*>(ptr), static_cast<void (MyQVideoProbe::*)()>(&MyQVideoProbe::Signal_Flush));;
}

int QVideoProbe_IsActive(QtObjectPtr ptr){
	return static_cast<QVideoProbe*>(ptr)->isActive();
}

int QVideoProbe_SetSource(QtObjectPtr ptr, QtObjectPtr source){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaObject*>(source));
}

int QVideoProbe_SetSource2(QtObjectPtr ptr, QtObjectPtr mediaRecorder){
	return static_cast<QVideoProbe*>(ptr)->setSource(static_cast<QMediaRecorder*>(mediaRecorder));
}

void QVideoProbe_DestroyQVideoProbe(QtObjectPtr ptr){
	static_cast<QVideoProbe*>(ptr)->~QVideoProbe();
}

