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

QtObjectPtr QNetworkInterface_NewQNetworkInterface(){
	return new QNetworkInterface();
}

QtObjectPtr QNetworkInterface_NewQNetworkInterface2(QtObjectPtr other){
	return new QNetworkInterface(*static_cast<QNetworkInterface*>(other));
}

int QNetworkInterface_Flags(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->flags();
}

char* QNetworkInterface_HardwareAddress(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->hardwareAddress().toUtf8().data();
}

char* QNetworkInterface_HumanReadableName(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->humanReadableName().toUtf8().data();
}

int QNetworkInterface_Index(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->index();
}

int QNetworkInterface_IsValid(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->isValid();
}

char* QNetworkInterface_Name(QtObjectPtr ptr){
	return static_cast<QNetworkInterface*>(ptr)->name().toUtf8().data();
}

void QNetworkInterface_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkInterface*>(ptr)->swap(*static_cast<QNetworkInterface*>(other));
}

void QNetworkInterface_DestroyQNetworkInterface(QtObjectPtr ptr){
	static_cast<QNetworkInterface*>(ptr)->~QNetworkInterface();
}

