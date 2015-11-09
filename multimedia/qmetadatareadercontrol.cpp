#include "qmetadatareadercontrol.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaDataReaderControl>
#include "_cgo_export.h"

class MyQMetaDataReaderControl: public QMetaDataReaderControl {
public:
void Signal_MetaDataAvailableChanged(bool available){callbackQMetaDataReaderControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMetaDataReaderControlMetaDataChanged(this->objectName().toUtf8().data());};
};

char* QMetaDataReaderControl_AvailableMetaData(void* ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMetaDataReaderControl_IsMetaDataAvailable(void* ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->isMetaDataAvailable();
}

void* QMetaDataReaderControl_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMetaDataReaderControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DestroyQMetaDataReaderControl(void* ptr){
	static_cast<QMetaDataReaderControl*>(ptr)->~QMetaDataReaderControl();
}

