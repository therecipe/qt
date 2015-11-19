#include "qgeoareamonitorsource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include <QGeoAreaMonitorInfo>
#include <QObject>
#include <QGeoAreaMonitorSource>
#include "_cgo_export.h"

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource {
public:
};

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources(){
	return QGeoAreaMonitorSource::availableSources().join("|").toUtf8().data();
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent){
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent){
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoAreaMonitorSource_Error(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->error();
}

void* QGeoAreaMonitorSource_PositionInfoSource(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->positionInfoSource();
}

int QGeoAreaMonitorSource_RequestUpdate(void* ptr, void* monitor, char* signal){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->requestUpdate(*static_cast<QGeoAreaMonitorInfo*>(monitor), const_cast<const char*>(signal));
}

void QGeoAreaMonitorSource_SetPositionInfoSource(void* ptr, void* newSource){
	static_cast<QGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

char* QGeoAreaMonitorSource_SourceName(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoAreaMonitorSource_StartMonitoring(void* ptr, void* monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->startMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_StopMonitoring(void* ptr, void* monitor){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->stopMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

int QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(void* ptr){
	return static_cast<QGeoAreaMonitorSource*>(ptr)->supportedAreaMonitorFeatures();
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(void* ptr){
	static_cast<QGeoAreaMonitorSource*>(ptr)->~QGeoAreaMonitorSource();
}

#include "qgeopositioninfosource.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include "_cgo_export.h"

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource {
public:
void Signal_UpdateTimeout(){callbackQGeoPositionInfoSourceUpdateTimeout(this->objectName().toUtf8().data());};
};

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

char* QGeoPositionInfoSource_SourceName(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8().data();
}

int QGeoPositionInfoSource_UpdateInterval(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->updateInterval();
}

char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources(){
	return QGeoPositionInfoSource::availableSources().join("|").toUtf8().data();
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent){
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoPositionInfoSource_Error(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->error();
}

int QGeoPositionInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QGeoPositionInfoSource_PreferredPositioningMethods(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->preferredPositioningMethods();
}

void QGeoPositionInfoSource_RequestUpdate(void* ptr, int timeout){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethods(void* ptr, int methods){
	static_cast<QGeoPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "startUpdates");
}

void QGeoPositionInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "stopUpdates");
}

int QGeoPositionInfoSource_SupportedPositioningMethods(void* ptr){
	return static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

void QGeoPositionInfoSource_ConnectUpdateTimeout(void* ptr){
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DisconnectUpdateTimeout(void* ptr){
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));;
}

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(void* ptr){
	static_cast<QGeoPositionInfoSource*>(ptr)->~QGeoPositionInfoSource();
}

#include "qgeoshape.h"
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoShape>
#include "_cgo_export.h"

class MyQGeoShape: public QGeoShape {
public:
};

void* QGeoShape_NewQGeoShape(){
	return new QGeoShape();
}

void* QGeoShape_NewQGeoShape2(void* other){
	return new QGeoShape(*static_cast<QGeoShape*>(other));
}

int QGeoShape_Contains(void* ptr, void* coordinate){
	return static_cast<QGeoShape*>(ptr)->contains(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoShape_ExtendShape(void* ptr, void* coordinate){
	static_cast<QGeoShape*>(ptr)->extendShape(*static_cast<QGeoCoordinate*>(coordinate));
}

int QGeoShape_IsEmpty(void* ptr){
	return static_cast<QGeoShape*>(ptr)->isEmpty();
}

int QGeoShape_IsValid(void* ptr){
	return static_cast<QGeoShape*>(ptr)->isValid();
}

char* QGeoShape_ToString(void* ptr){
	return static_cast<QGeoShape*>(ptr)->toString().toUtf8().data();
}

int QGeoShape_Type(void* ptr){
	return static_cast<QGeoShape*>(ptr)->type();
}

void QGeoShape_DestroyQGeoShape(void* ptr){
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

#include "qgeosatelliteinfo.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoSatelliteInfo>
#include "_cgo_export.h"

class MyQGeoSatelliteInfo: public QGeoSatelliteInfo {
public:
};

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo(){
	return new QGeoSatelliteInfo();
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo2(void* other){
	return new QGeoSatelliteInfo(*static_cast<QGeoSatelliteInfo*>(other));
}

double QGeoSatelliteInfo_Attribute(void* ptr, int attribute){
	return static_cast<double>(static_cast<QGeoSatelliteInfo*>(ptr)->attribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute)));
}

int QGeoSatelliteInfo_HasAttribute(void* ptr, int attribute){
	return static_cast<QGeoSatelliteInfo*>(ptr)->hasAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

void QGeoSatelliteInfo_RemoveAttribute(void* ptr, int attribute){
	static_cast<QGeoSatelliteInfo*>(ptr)->removeAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

int QGeoSatelliteInfo_SatelliteIdentifier(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteIdentifier();
}

int QGeoSatelliteInfo_SatelliteSystem(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteSystem();
}

void QGeoSatelliteInfo_SetAttribute(void* ptr, int attribute, double value){
	static_cast<QGeoSatelliteInfo*>(ptr)->setAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute), static_cast<qreal>(value));
}

void QGeoSatelliteInfo_SetSatelliteIdentifier(void* ptr, int satId){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteIdentifier(satId);
}

void QGeoSatelliteInfo_SetSatelliteSystem(void* ptr, int system){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteSystem(static_cast<QGeoSatelliteInfo::SatelliteSystem>(system));
}

void QGeoSatelliteInfo_SetSignalStrength(void* ptr, int signalStrength){
	static_cast<QGeoSatelliteInfo*>(ptr)->setSignalStrength(signalStrength);
}

int QGeoSatelliteInfo_SignalStrength(void* ptr){
	return static_cast<QGeoSatelliteInfo*>(ptr)->signalStrength();
}

void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(void* ptr){
	static_cast<QGeoSatelliteInfo*>(ptr)->~QGeoSatelliteInfo();
}

#include "qgeoareamonitorinfo.h"
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QDate>
#include <QGeoShape>
#include <QString>
#include <QVariant>
#include <QGeoAreaMonitorInfo>
#include "_cgo_export.h"

class MyQGeoAreaMonitorInfo: public QGeoAreaMonitorInfo {
public:
};

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(void* other){
	return new QGeoAreaMonitorInfo(*static_cast<QGeoAreaMonitorInfo*>(other));
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name){
	return new QGeoAreaMonitorInfo(QString(name));
}

void* QGeoAreaMonitorInfo_Expiration(void* ptr){
	return new QDateTime(static_cast<QGeoAreaMonitorInfo*>(ptr)->expiration());
}

char* QGeoAreaMonitorInfo_Identifier(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->identifier().toUtf8().data();
}

int QGeoAreaMonitorInfo_IsPersistent(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isPersistent();
}

int QGeoAreaMonitorInfo_IsValid(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isValid();
}

char* QGeoAreaMonitorInfo_Name(void* ptr){
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->name().toUtf8().data();
}

void QGeoAreaMonitorInfo_SetArea(void* ptr, void* newShape){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setArea(*static_cast<QGeoShape*>(newShape));
}

void QGeoAreaMonitorInfo_SetExpiration(void* ptr, void* expiry){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setExpiration(*static_cast<QDateTime*>(expiry));
}

void QGeoAreaMonitorInfo_SetName(void* ptr, char* name){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setName(QString(name));
}

void QGeoAreaMonitorInfo_SetPersistent(void* ptr, int isPersistent){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setPersistent(isPersistent != 0);
}

void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(void* ptr){
	static_cast<QGeoAreaMonitorInfo*>(ptr)->~QGeoAreaMonitorInfo();
}

#include "qgeocircle.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoShape>
#include <QGeoCircle>
#include "_cgo_export.h"

class MyQGeoCircle: public QGeoCircle {
public:
};

void* QGeoCircle_NewQGeoCircle(){
	return new QGeoCircle();
}

void* QGeoCircle_NewQGeoCircle3(void* other){
	return new QGeoCircle(*static_cast<QGeoCircle*>(other));
}

void* QGeoCircle_NewQGeoCircle2(void* center, double radius){
	return new QGeoCircle(*static_cast<QGeoCoordinate*>(center), static_cast<qreal>(radius));
}

void* QGeoCircle_NewQGeoCircle4(void* other){
	return new QGeoCircle(*static_cast<QGeoShape*>(other));
}

double QGeoCircle_Radius(void* ptr){
	return static_cast<double>(static_cast<QGeoCircle*>(ptr)->radius());
}

void QGeoCircle_SetCenter(void* ptr, void* center){
	static_cast<QGeoCircle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoCircle_SetRadius(void* ptr, double radius){
	static_cast<QGeoCircle*>(ptr)->setRadius(static_cast<qreal>(radius));
}

char* QGeoCircle_ToString(void* ptr){
	return static_cast<QGeoCircle*>(ptr)->toString().toUtf8().data();
}

void QGeoCircle_DestroyQGeoCircle(void* ptr){
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

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

#include "qgeorectangle.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoShape>
#include <QGeoRectangle>
#include "_cgo_export.h"

class MyQGeoRectangle: public QGeoRectangle {
public:
};

void* QGeoRectangle_NewQGeoRectangle(){
	return new QGeoRectangle();
}

void* QGeoRectangle_NewQGeoRectangle3(void* topLeft, void* bottomRight){
	return new QGeoRectangle(*static_cast<QGeoCoordinate*>(topLeft), *static_cast<QGeoCoordinate*>(bottomRight));
}

void* QGeoRectangle_NewQGeoRectangle5(void* other){
	return new QGeoRectangle(*static_cast<QGeoRectangle*>(other));
}

void* QGeoRectangle_NewQGeoRectangle6(void* other){
	return new QGeoRectangle(*static_cast<QGeoShape*>(other));
}

int QGeoRectangle_Contains(void* ptr, void* rectangle){
	return static_cast<QGeoRectangle*>(ptr)->contains(*static_cast<QGeoRectangle*>(rectangle));
}

int QGeoRectangle_Intersects(void* ptr, void* rectangle){
	return static_cast<QGeoRectangle*>(ptr)->intersects(*static_cast<QGeoRectangle*>(rectangle));
}

void QGeoRectangle_SetBottomLeft(void* ptr, void* bottomLeft){
	static_cast<QGeoRectangle*>(ptr)->setBottomLeft(*static_cast<QGeoCoordinate*>(bottomLeft));
}

void QGeoRectangle_SetBottomRight(void* ptr, void* bottomRight){
	static_cast<QGeoRectangle*>(ptr)->setBottomRight(*static_cast<QGeoCoordinate*>(bottomRight));
}

void QGeoRectangle_SetCenter(void* ptr, void* center){
	static_cast<QGeoRectangle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoRectangle_SetTopLeft(void* ptr, void* topLeft){
	static_cast<QGeoRectangle*>(ptr)->setTopLeft(*static_cast<QGeoCoordinate*>(topLeft));
}

void QGeoRectangle_SetTopRight(void* ptr, void* topRight){
	static_cast<QGeoRectangle*>(ptr)->setTopRight(*static_cast<QGeoCoordinate*>(topRight));
}

char* QGeoRectangle_ToString(void* ptr){
	return static_cast<QGeoRectangle*>(ptr)->toString().toUtf8().data();
}

void QGeoRectangle_DestroyQGeoRectangle(void* ptr){
	static_cast<QGeoRectangle*>(ptr)->~QGeoRectangle();
}

#include "qnmeapositioninfosource.h"
#include <QIODevice>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QNmeaPositionInfoSource>
#include "_cgo_export.h"

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource {
public:
};

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, void* parent){
	return new QNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
}

void* QNmeaPositionInfoSource_Device(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->device();
}

int QNmeaPositionInfoSource_Error(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->error();
}

int QNmeaPositionInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

void QNmeaPositionInfoSource_RequestUpdate(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_SetDevice(void* ptr, void* device){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QNmeaPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "stopUpdates");
}

int QNmeaPositionInfoSource_SupportedPositioningMethods(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

int QNmeaPositionInfoSource_UpdateMode(void* ptr){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->updateMode();
}

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(void* ptr){
	static_cast<QNmeaPositionInfoSource*>(ptr)->~QNmeaPositionInfoSource();
}

#include "qgeosatelliteinfosource.h"
#include <QGeoSatelliteInfo>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QMetaObject>
#include <QGeoSatelliteInfoSource>
#include "_cgo_export.h"

class MyQGeoSatelliteInfoSource: public QGeoSatelliteInfoSource {
public:
void Signal_RequestTimeout(){callbackQGeoSatelliteInfoSourceRequestTimeout(this->objectName().toUtf8().data());};
};

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

int QGeoSatelliteInfoSource_UpdateInterval(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->updateInterval();
}

char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources(){
	return QGeoSatelliteInfoSource::availableSources().join("|").toUtf8().data();
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent){
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoSatelliteInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

int QGeoSatelliteInfoSource_Error(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->error();
}

int QGeoSatelliteInfoSource_MinimumUpdateInterval(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->minimumUpdateInterval();
}

void QGeoSatelliteInfoSource_ConnectRequestTimeout(void* ptr){
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));;
}

void QGeoSatelliteInfoSource_DisconnectRequestTimeout(void* ptr){
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));;
}

void QGeoSatelliteInfoSource_RequestUpdate(void* ptr, int timeout){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

char* QGeoSatelliteInfoSource_SourceName(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->sourceName().toUtf8().data();
}

void QGeoSatelliteInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "startUpdates");
}

void QGeoSatelliteInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "stopUpdates");
}

void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(void* ptr){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->~QGeoSatelliteInfoSource();
}

#include "qgeocoordinate.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoCoordinate>
#include "_cgo_export.h"

class MyQGeoCoordinate: public QGeoCoordinate {
public:
};

void* QGeoCoordinate_NewQGeoCoordinate(){
	return new QGeoCoordinate();
}

void* QGeoCoordinate_NewQGeoCoordinate4(void* other){
	return new QGeoCoordinate(*static_cast<QGeoCoordinate*>(other));
}

double QGeoCoordinate_AzimuthTo(void* ptr, void* other){
	return static_cast<double>(static_cast<QGeoCoordinate*>(ptr)->azimuthTo(*static_cast<QGeoCoordinate*>(other)));
}

double QGeoCoordinate_DistanceTo(void* ptr, void* other){
	return static_cast<double>(static_cast<QGeoCoordinate*>(ptr)->distanceTo(*static_cast<QGeoCoordinate*>(other)));
}

int QGeoCoordinate_IsValid(void* ptr){
	return static_cast<QGeoCoordinate*>(ptr)->isValid();
}

char* QGeoCoordinate_ToString(void* ptr, int format){
	return static_cast<QGeoCoordinate*>(ptr)->toString(static_cast<QGeoCoordinate::CoordinateFormat>(format)).toUtf8().data();
}

int QGeoCoordinate_Type(void* ptr){
	return static_cast<QGeoCoordinate*>(ptr)->type();
}

void QGeoCoordinate_DestroyQGeoCoordinate(void* ptr){
	static_cast<QGeoCoordinate*>(ptr)->~QGeoCoordinate();
}

#include "qgeopositioninfosourcefactory.h"
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include <QString>
#include <QVariant>
#include <QGeoPositionInfoSourceFactory>
#include "_cgo_export.h"

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory {
public:
};

void* QGeoPositionInfoSourceFactory_AreaMonitor(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->areaMonitor(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_PositionInfoSource(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->positionInfoSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_SatelliteInfoSource(void* ptr, void* parent){
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->satelliteInfoSource(static_cast<QObject*>(parent));
}

void QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(void* ptr){
	static_cast<QGeoPositionInfoSourceFactory*>(ptr)->~QGeoPositionInfoSourceFactory();
}

#include "qgeopositioninfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDateTime>
#include <QDate>
#include <QGeoCoordinate>
#include <QGeoPositionInfo>
#include "_cgo_export.h"

class MyQGeoPositionInfo: public QGeoPositionInfo {
public:
};

void* QGeoPositionInfo_NewQGeoPositionInfo(){
	return new QGeoPositionInfo();
}

void* QGeoPositionInfo_NewQGeoPositionInfo2(void* coordinate, void* timestamp){
	return new QGeoPositionInfo(*static_cast<QGeoCoordinate*>(coordinate), *static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_NewQGeoPositionInfo3(void* other){
	return new QGeoPositionInfo(*static_cast<QGeoPositionInfo*>(other));
}

double QGeoPositionInfo_Attribute(void* ptr, int attribute){
	return static_cast<double>(static_cast<QGeoPositionInfo*>(ptr)->attribute(static_cast<QGeoPositionInfo::Attribute>(attribute)));
}

int QGeoPositionInfo_HasAttribute(void* ptr, int attribute){
	return static_cast<QGeoPositionInfo*>(ptr)->hasAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

int QGeoPositionInfo_IsValid(void* ptr){
	return static_cast<QGeoPositionInfo*>(ptr)->isValid();
}

void QGeoPositionInfo_RemoveAttribute(void* ptr, int attribute){
	static_cast<QGeoPositionInfo*>(ptr)->removeAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

void QGeoPositionInfo_SetAttribute(void* ptr, int attribute, double value){
	static_cast<QGeoPositionInfo*>(ptr)->setAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute), static_cast<qreal>(value));
}

void QGeoPositionInfo_SetCoordinate(void* ptr, void* coordinate){
	static_cast<QGeoPositionInfo*>(ptr)->setCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPositionInfo_SetTimestamp(void* ptr, void* timestamp){
	static_cast<QGeoPositionInfo*>(ptr)->setTimestamp(*static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_Timestamp(void* ptr){
	return new QDateTime(static_cast<QGeoPositionInfo*>(ptr)->timestamp());
}

void QGeoPositionInfo_DestroyQGeoPositionInfo(void* ptr){
	static_cast<QGeoPositionInfo*>(ptr)->~QGeoPositionInfo();
}

#include "qgeolocation.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include "_cgo_export.h"

