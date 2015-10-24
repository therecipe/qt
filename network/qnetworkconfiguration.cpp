#include "qnetworkconfiguration.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QNetworkConfiguration>
#include "_cgo_export.h"

class MyQNetworkConfiguration: public QNetworkConfiguration {
public:
};

QtObjectPtr QNetworkConfiguration_NewQNetworkConfiguration(){
	return new QNetworkConfiguration();
}

QtObjectPtr QNetworkConfiguration_NewQNetworkConfiguration2(QtObjectPtr other){
	return new QNetworkConfiguration(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_BearerType(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

int QNetworkConfiguration_BearerTypeFamily(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

char* QNetworkConfiguration_BearerTypeName(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8().data();
}

char* QNetworkConfiguration_Identifier(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8().data();
}

int QNetworkConfiguration_IsRoamingAvailable(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

int QNetworkConfiguration_IsValid(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

char* QNetworkConfiguration_Name(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8().data();
}

int QNetworkConfiguration_Purpose(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

void QNetworkConfiguration_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_Type(QtObjectPtr ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(QtObjectPtr ptr){
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

