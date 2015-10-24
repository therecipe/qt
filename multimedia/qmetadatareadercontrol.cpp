#include "qmetadatareadercontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaDataReaderControl>
#include "_cgo_export.h"

class MyQMetaDataReaderControl: public QMetaDataReaderControl {
public:
void Signal_MetaDataAvailableChanged(bool available){callbackQMetaDataReaderControlMetaDataAvailableChanged(this->objectName().toUtf8().data(), available);};
void Signal_MetaDataChanged(){callbackQMetaDataReaderControlMetaDataChanged(this->objectName().toUtf8().data());};
};

char* QMetaDataReaderControl_AvailableMetaData(QtObjectPtr ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->availableMetaData().join("|").toUtf8().data();
}

int QMetaDataReaderControl_IsMetaDataAvailable(QtObjectPtr ptr){
	return static_cast<QMetaDataReaderControl*>(ptr)->isMetaDataAvailable();
}

char* QMetaDataReaderControl_MetaData(QtObjectPtr ptr, char* key){
	return static_cast<QMetaDataReaderControl*>(ptr)->metaData(QString(key)).toString().toUtf8().data();
}

void QMetaDataReaderControl_ConnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataAvailableChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)(bool)>(&QMetaDataReaderControl::metaDataAvailableChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)(bool)>(&MyQMetaDataReaderControl::Signal_MetaDataAvailableChanged));;
}

void QMetaDataReaderControl_ConnectMetaDataChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DisconnectMetaDataChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QMetaDataReaderControl*>(ptr), static_cast<void (QMetaDataReaderControl::*)()>(&QMetaDataReaderControl::metaDataChanged), static_cast<MyQMetaDataReaderControl*>(ptr), static_cast<void (MyQMetaDataReaderControl::*)()>(&MyQMetaDataReaderControl::Signal_MetaDataChanged));;
}

void QMetaDataReaderControl_DestroyQMetaDataReaderControl(QtObjectPtr ptr){
	static_cast<QMetaDataReaderControl*>(ptr)->~QMetaDataReaderControl();
}

