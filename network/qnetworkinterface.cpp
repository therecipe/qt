#include "qnetworkinterface.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QNetworkInterface>
#include "_cgo_export.h"

class MyQNetworkInterface: public QNetworkInterface {
public:
};

void* QNetworkInterface_NewQNetworkInterface(){
	return new QNetworkInterface();
}

void* QNetworkInterface_NewQNetworkInterface2(void* other){
	return new QNetworkInterface(*static_cast<QNetworkInterface*>(other));
}

int QNetworkInterface_Flags(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

char* QNetworkInterface_HardwareAddress(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8().data();
}

char* QNetworkInterface_HumanReadableName(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8().data();
}

int QNetworkInterface_Index(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->index();
}

int QNetworkInterface_IsValid(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

char* QNetworkInterface_Name(void* ptr){
	return static_cast<QNetworkInterface*>(ptr)->name().toUtf8().data();
}

void QNetworkInterface_Swap(void* ptr, void* other){
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(void* ptr){
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

