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

int QMediaObject_NotifyInterval(QtObjectPtr ptr){
	return static_cast<QMediaObject*>(ptr)->notifyInterval();
}

void QMediaObject_SetNotifyInterval(QtObjectPtr ptr, int milliSeconds){
	static_cast<QMediaObject*>(ptr)->setNotifyInterval(milliSeconds);
}

void QMediaObject_ConnectAvailabilityChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

void QMediaObject_DisconnectAvailabilityChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::availabilityChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_AvailabilityChanged));;
}

char* QMediaObject_AvailableMetaData(QtObjectPtr ptr){
	return static_cast<QMediaObject*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMediaObject_Bind(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QMediaObject*>(ptr)->bind(static_cast<QObject*>(object));
}

int QMediaObject_IsAvailable(QtObjectPtr ptr){
	return static_cast<QMediaObject*>(ptr)->isAvailable();
}

int QMediaObject_IsMetaDataAvailable(QtObjectPtr ptr){
	return static_cast<QMediaObject*>(ptr)->isMetaDataAvailable();
}

char* QMediaObject_MetaData(QtObjectPtr ptr, char* key){
	return static_cast<QMediaObject*>(ptr)->metaData(QString(key)).toString().toUtf8().data();
}

void QMediaObject_ConnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(bool)>(&QMediaObject::metaDataAvailableChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(bool)>(&MyQMediaObject::Signal_MetaDataAvailableChanged));;
}

void QMediaObject_ConnectMetaDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_DisconnectMetaDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)()>(&QMediaObject::metaDataChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)()>(&MyQMediaObject::Signal_MetaDataChanged));;
}

void QMediaObject_ConnectNotifyIntervalChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

void QMediaObject_DisconnectNotifyIntervalChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMediaObject*>(ptr), static_cast<void (QMediaObject::*)(int)>(&QMediaObject::notifyIntervalChanged), static_cast<MyQMediaObject*>(ptr), static_cast<void (MyQMediaObject::*)(int)>(&MyQMediaObject::Signal_NotifyIntervalChanged));;
}

QtObjectPtr QMediaObject_Service(QtObjectPtr ptr){
	return static_cast<QMediaObject*>(ptr)->service();
}

void QMediaObject_Unbind(QtObjectPtr ptr, QtObjectPtr object){
	static_cast<QMediaObject*>(ptr)->unbind(static_cast<QObject*>(object));
}

void QMediaObject_DestroyQMediaObject(QtObjectPtr ptr){
	static_cast<QMediaObject*>(ptr)->~QMediaObject();
}

