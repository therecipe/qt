#include "qmetadatawritercontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaDataWriterControl>
#include "_cgo_export.h"

class MyQMetaDataWriterControl: public QMetaDataWriterControl {
public:
void Signal_MetaDataAvailableChanged(bool available){callbackQMetaDataWriterControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMetaDataWriterControlMetaDataChanged(this->objectName().toUtf8().data());};
void Signal_WritableChanged(bool writable){callbackQMetaDataWriterControlWritableChanged(this->objectName().toUtf8().data(), writable);};
};

char* QMetaDataWriterControl_AvailableMetaData(QtObjectPtr ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMetaDataWriterControl_IsMetaDataAvailable(QtObjectPtr ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isMetaDataAvailable();
}

int QMetaDataWriterControl_IsWritable(QtObjectPtr ptr){
	return static_cast<QMetaDataWriterControl*>(ptr)->isWritable();
}

char* QMetaDataWriterControl_MetaData(QtObjectPtr ptr, char* key){
	return static_cast<QMetaDataWriterControl*>(ptr)->metaData(QString(key)).toString().toUtf8().data();
}

void QMetaDataWriterControl_ConnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::metaDataAvailableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataWriterControl_ConnectMetaDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_DisconnectMetaDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)()>(&QMetaDataWriterControl::metaDataChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)()>(&MyQMetaDataWriterControl::Signal_MetaDataChanged));;
}

void QMetaDataWriterControl_SetMetaData(QtObjectPtr ptr, char* key, char* value){
	static_cast<QMetaDataWriterControl*>(ptr)->setMetaData(QString(key), QVariant(value));
}

void QMetaDataWriterControl_ConnectWritableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DisconnectWritableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMetaDataWriterControl*>(ptr), static_cast<void (QMetaDataWriterControl::*)(bool)>(&QMetaDataWriterControl::writableChanged), static_cast<MyQMetaDataWriterControl*>(ptr), static_cast<void (MyQMetaDataWriterControl::*)(bool)>(&MyQMetaDataWriterControl::Signal_WritableChanged));;
}

void QMetaDataWriterControl_DestroyQMetaDataWriterControl(QtObjectPtr ptr){
	static_cast<QMetaDataWriterControl*>(ptr)->~QMetaDataWriterControl();
}

