#include "qmetadatawritercontrol.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMetaDataWriterControl>
#include "_cgo_export.h"

class MyQMetaDataWriterControl: public QMetaDataWriterControl {
public:
void Signal_MetaDataAvailableChanged(bool available){callbackQMetaDataWriterControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMetaDataWriterControlMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_WritableChanged(bool writable){callbackQMetaDataWriterControlWritableChanged(this->objectName().toUtf8().data(), writable);};
};

char* QMetaDataWriterControl_AvailableMetaData(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMetaDataWriterControl_IsMetaDataAvailable(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isMetaDataAvailable();
}

int QMetaDataWriterControl_IsWritable(void* ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isWritable();
}

void* QMetaDataWriterControl_MetaData(void* ptr, char* key){
	return new QVariant(static_cast<QMetaDataWriterControl*>(ptr)->metaData(QString(key)));
}

void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_ConnectMetaDataChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_SetMetaData(void* ptr, char* key, void* value){
	static_cast<QMetaDataWriterControl*>(ptr)->setMetaData(QString(key), *static_cast<QVariant*>(value));
}

void QMetaDataWriterControl_ConnectWritableChanged(void* ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DisconnectWritableChanged(void* ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DestroyQMetaDataWriterControl(void* ptr){
	static_cast<QMetaDataWriterControl*>(ptr)->~QMetaDataWriterControl();
}

