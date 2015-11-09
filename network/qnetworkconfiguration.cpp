#include "qnetworkconfiguration.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QNetworkConfiguration>
#include "_cgo_export.h"

class MyQNetworkConfiguration: public QNetworkConfiguration {
public:
};

void* QNetworkConfiguration_NewQNetworkConfiguration(){
	return new QNetworkConfiguration();
}

void* QNetworkConfiguration_NewQNetworkConfiguration2(void* other){
	return new QNetworkConfiguration(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_BearerType(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerType();
}

int QNetworkConfiguration_BearerTypeFamily(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeFamily();
}

char* QNetworkConfiguration_BearerTypeName(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->bearerTypeName().toUtf8().data();
}

char* QNetworkConfiguration_Identifier(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->identifier().toUtf8().data();
}

int QNetworkConfiguration_IsRoamingAvailable(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isRoamingAvailable();
}

int QNetworkConfiguration_IsValid(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->isValid();
}

char* QNetworkConfiguration_Name(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->name().toUtf8().data();
}

int QNetworkConfiguration_Purpose(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->purpose();
}

void QNetworkConfiguration_Swap(void* ptr, void* other){
	static_cast<QNetworkConfiguration*>(ptr)->swap(*static_cast<QNetworkConfiguration*>(other));
}

int QNetworkConfiguration_Type(void* ptr){
	return static_cast<QNetworkConfiguration*>(ptr)->type();
}

void QNetworkConfiguration_DestroyQNetworkConfiguration(void* ptr){
	static_cast<QNetworkConfiguration*>(ptr)->~QNetworkConfiguration();
}

