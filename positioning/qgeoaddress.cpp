#include "qgeoaddress.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoAddress>
#include "_cgo_export.h"

class MyQGeoAddress: public QGeoAddress {
public:
};

QtObjectPtr QGeoAddress_NewQGeoAddress(){
	return new QGeoAddress();
}

QtObjectPtr QGeoAddress_NewQGeoAddress2(QtObjectPtr other){
	return new QGeoAddress(*static_cast<QGeoAddress*>(other));
}

char* QGeoAddress_City(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->city().toUtf8().data();
}

void QGeoAddress_Clear(QtObjectPtr ptr){
	static_cast<QGeoAddress*>(ptr)->clear();
}

char* QGeoAddress_Country(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->country().toUtf8().data();
}

char* QGeoAddress_CountryCode(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->countryCode().toUtf8().data();
}

char* QGeoAddress_County(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->county().toUtf8().data();
}

char* QGeoAddress_District(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->district().toUtf8().data();
}

int QGeoAddress_IsEmpty(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->isEmpty();
}

int QGeoAddress_IsTextGenerated(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->isTextGenerated();
}

char* QGeoAddress_PostalCode(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->postalCode().toUtf8().data();
}

void QGeoAddress_SetCity(QtObjectPtr ptr, char* city){
	static_cast<QGeoAddress*>(ptr)->setCity(QString(city));
}

void QGeoAddress_SetCountry(QtObjectPtr ptr, char* country){
	static_cast<QGeoAddress*>(ptr)->setCountry(QString(country));
}

void QGeoAddress_SetCountryCode(QtObjectPtr ptr, char* countryCode){
	static_cast<QGeoAddress*>(ptr)->setCountryCode(QString(countryCode));
}

void QGeoAddress_SetCounty(QtObjectPtr ptr, char* county){
	static_cast<QGeoAddress*>(ptr)->setCounty(QString(county));
}

void QGeoAddress_SetDistrict(QtObjectPtr ptr, char* district){
	static_cast<QGeoAddress*>(ptr)->setDistrict(QString(district));
}

void QGeoAddress_SetPostalCode(QtObjectPtr ptr, char* postalCode){
	static_cast<QGeoAddress*>(ptr)->setPostalCode(QString(postalCode));
}

void QGeoAddress_SetState(QtObjectPtr ptr, char* state){
	static_cast<QGeoAddress*>(ptr)->setState(QString(state));
}

void QGeoAddress_SetStreet(QtObjectPtr ptr, char* street){
	static_cast<QGeoAddress*>(ptr)->setStreet(QString(street));
}

void QGeoAddress_SetText(QtObjectPtr ptr, char* text){
	static_cast<QGeoAddress*>(ptr)->setText(QString(text));
}

char* QGeoAddress_Street(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->street().toUtf8().data();
}

char* QGeoAddress_Text(QtObjectPtr ptr){
	return static_cast<QGeoAddress*>(ptr)->text().toUtf8().data();
}

void QGeoAddress_DestroyQGeoAddress(QtObjectPtr ptr){
	static_cast<QGeoAddress*>(ptr)->~QGeoAddress();
}

