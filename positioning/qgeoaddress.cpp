#include "qgeoaddress.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoAddress>
#include "_cgo_export.h"

class MyQGeoAddress: public QGeoAddress {
public:
};

void* QGeoAddress_NewQGeoAddress(){
	return new QGeoAddress();
}

void* QGeoAddress_NewQGeoAddress2(void* other){
	return new QGeoAddress(*static_cast<QGeoAddress*>(other));
}

char* QGeoAddress_City(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->city().toUtf8().data();
}

void QGeoAddress_Clear(void* ptr){
	static_cast<QGeoAddress*>(ptr)->clear();
}

char* QGeoAddress_Country(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->country().toUtf8().data();
}

char* QGeoAddress_CountryCode(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->countryCode().toUtf8().data();
}

char* QGeoAddress_County(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->county().toUtf8().data();
}

char* QGeoAddress_District(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->district().toUtf8().data();
}

int QGeoAddress_IsEmpty(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->isEmpty();
}

int QGeoAddress_IsTextGenerated(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->isTextGenerated();
}

char* QGeoAddress_PostalCode(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->postalCode().toUtf8().data();
}

void QGeoAddress_SetCity(void* ptr, char* city){
	static_cast<QGeoAddress*>(ptr)->setCity(QString(city));
}

void QGeoAddress_SetCountry(void* ptr, char* country){
	static_cast<QGeoAddress*>(ptr)->setCountry(QString(country));
}

void QGeoAddress_SetCountryCode(void* ptr, char* countryCode){
	static_cast<QGeoAddress*>(ptr)->setCountryCode(QString(countryCode));
}

void QGeoAddress_SetCounty(void* ptr, char* county){
	static_cast<QGeoAddress*>(ptr)->setCounty(QString(county));
}

void QGeoAddress_SetDistrict(void* ptr, char* district){
	static_cast<QGeoAddress*>(ptr)->setDistrict(QString(district));
}

void QGeoAddress_SetPostalCode(void* ptr, char* postalCode){
	static_cast<QGeoAddress*>(ptr)->setPostalCode(QString(postalCode));
}

void QGeoAddress_SetState(void* ptr, char* state){
	static_cast<QGeoAddress*>(ptr)->setState(QString(state));
}

void QGeoAddress_SetStreet(void* ptr, char* street){
	static_cast<QGeoAddress*>(ptr)->setStreet(QString(street));
}

void QGeoAddress_SetText(void* ptr, char* text){
	static_cast<QGeoAddress*>(ptr)->setText(QString(text));
}

char* QGeoAddress_Street(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->street().toUtf8().data();
}

char* QGeoAddress_Text(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->text().toUtf8().data();
}

void QGeoAddress_DestroyQGeoAddress(void* ptr){
	static_cast<QGeoAddress*>(ptr)->~QGeoAddress();
}

