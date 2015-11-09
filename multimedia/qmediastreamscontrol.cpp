#include "qmediastreamscontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QMediaStreamsControl>
#include "_cgo_export.h"

class MyQMediaStreamsControl: public QMediaStreamsControl {
public:
void Signal_ActiveStreamsChanged(){callbackQMediaStreamsControlActiveStreamsChanged(this->objectName().toUtf8().data());};
void Signal_StreamsChanged(){callbackQMediaStreamsControlStreamsChanged(this->objectName().toUtf8().data());};
};

void QMediaStreamsControl_ConnectActiveStreamsChanged(void* ptr){
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));;
}

void QMediaStreamsControl_DisconnectActiveStreamsChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::activeStreamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_ActiveStreamsChanged));;
}

int QMediaStreamsControl_IsActive(void* ptr, int stream){
	return static_cast<QMediaStreamsControl*>(ptr)->isActive(stream);
}

void* QMediaStreamsControl_MetaData(void* ptr, int stream, char* key){
	return new QVariant(static_cast<QMediaStreamsControl*>(ptr)->metaData(stream, QString(key)));
}

void QMediaStreamsControl_SetActive(void* ptr, int stream, int state){
	static_cast<QMediaStreamsControl*>(ptr)->setActive(stream, state != 0);
}

int QMediaStreamsControl_StreamCount(void* ptr){
	return static_cast<QMediaStreamsControl*>(ptr)->streamCount();
}

int QMediaStreamsControl_StreamType(void* ptr, int stream){
	return static_cast<QMediaStreamsControl*>(ptr)->streamType(stream);
}

void QMediaStreamsControl_ConnectStreamsChanged(void* ptr){
	QObject::connect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));;
}

void QMediaStreamsControl_DisconnectStreamsChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaStreamsControl*>(ptr), static_cast<void (QMediaStreamsControl::*)()>(&QMediaStreamsControl::streamsChanged), static_cast<MyQMediaStreamsControl*>(ptr), static_cast<void (MyQMediaStreamsControl::*)()>(&MyQMediaStreamsControl::Signal_StreamsChanged));;
}

void QMediaStreamsControl_DestroyQMediaStreamsControl(void* ptr){
	static_cast<QMediaStreamsControl*>(ptr)->~QMediaStreamsControl();
}

