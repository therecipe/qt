#include "qmediaobject.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMediaObject>
#include "_cgo_export.h"

class MyQMediaObject: public QMediaObject {
public:
void Signal_AvailabilityChanged(bool available){callbackQMediaObjectAvailabilityChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataAvailableChanged(bool available){callbackQMediaObjectMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMediaObjectMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_NotifyIntervalChanged(int milliseconds){callbackQMediaObjectNotifyIntervalChanged(this->objectName().toUtf8().data(), milliseconds);};
};

int QMediaObject_NotifyInterval(void* ptr){
	return static_cast<QMediaObject*>(ptr)->notifyInterval();
}

void QMediaObject_SetNotifyInterval(void* ptr, int milliSeconds){
	static_cast<QMediaObject*>(ptr)->setNotifyInterval(milliSeconds);
}

void QMediaObject_ConnectAvailabilityChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

void QMediaObject_DisconnectAvailabilityChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

char* QMediaObject_AvailableMetaData(void* ptr){
	return static_cast<QMediaObject*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMediaObject_Bind(void* ptr, void* object){
	return static_cast<QMediaObject*>(ptr)->bind(static_cast<QObject*>(object));
}

int QMediaObject_IsAvailable(void* ptr){
	return static_cast<QMediaObject*>(ptr)->isAvailable();
}

int QMediaObject_IsMetaDataAvailable(void* ptr){
	return static_cast<QMediaObject*>(ptr)->isMetaDataAvailable();
}

void* QMediaObject_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMediaObject*>(ptr)->metaData(QString(key)));
}

void QMediaObject_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_ConnectNotifyIntervalChanged(void* ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

void QMediaObject_DisconnectNotifyIntervalChanged(void* ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

void* QMediaObject_Service(void* ptr){
	return static_cast<QMediaObject*>(ptr)->service();
}

void QMediaObject_Unbind(void* ptr, void* object){
	static_cast<QMediaObject*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QMediaObject_DestroyQMediaObject(void* ptr){
	static_cast<QMediaObject*>(ptr)->~QMediaObject();
}

