#include "positioning.h"
#include "_cgo_export.h"

#include <QChildEvent>
#include <QDate>
#include <QDateTime>
#include <QEvent>
#include <QGeoAddress>
#include <QGeoAreaMonitorInfo>
#include <QGeoAreaMonitorSource>
#include <QGeoCircle>
#include <QGeoCoordinate>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include <QGeoPositionInfoSourceFactory>
#include <QGeoRectangle>
#include <QGeoSatelliteInfo>
#include <QGeoSatelliteInfoSource>
#include <QGeoShape>
#include <QIODevice>
#include <QMetaObject>
#include <QNmeaPositionInfoSource>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

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

char* QGeoAddress_State(void* ptr){
	return static_cast<QGeoAddress*>(ptr)->state().toUtf8().data();
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

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource {
public:
	void Signal_Error2(QGeoAreaMonitorSource::Error areaMonitoringError) { callbackQGeoAreaMonitorSourceError2(this->objectName().toUtf8().data(), areaMonitoringError); };
	void setPositionInfoSource(QGeoPositionInfoSource * newSource) { if (!callbackQGeoAreaMonitorSourceSetPositionInfoSource(this->objectName().toUtf8().data(), newSource)) { QGeoAreaMonitorSource::setPositionInfoSource(newSource); }; };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQGeoAreaMonitorSourceTimerEvent(this->objectName().toUtf8().data(), event)) { QGeoAreaMonitorSource::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQGeoAreaMonitorSourceChildEvent(this->objectName().toUtf8().data(), event)) { QGeoAreaMonitorSource::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQGeoAreaMonitorSourceCustomEvent(this->objectName().toUtf8().data(), event)) { QGeoAreaMonitorSource::customEvent(event); }; };
};

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources(){
	return QGeoAreaMonitorSource::availableSources().join(",,,").toUtf8().data();
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent){
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent){
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoAreaMonitorSource_ConnectError2(void* ptr){
	QObject::connect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&QGeoAreaMonitorSource::error), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&MyQGeoAreaMonitorSource::Signal_Error2));;
}

void QGeoAreaMonitorSource_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&QGeoAreaMonitorSource::error), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&MyQGeoAreaMonitorSource::Signal_Error2));;
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

void* QGeoCircle_NewQGeoCircle(){
	return new QGeoCircle();
}

void* QGeoCircle_NewQGeoCircle3(void* other){
	return new QGeoCircle(*static_cast<QGeoCircle*>(other));
}

void* QGeoCircle_NewQGeoCircle2(void* center, double radius){
	return new QGeoCircle(*static_cast<QGeoCoordinate*>(center), static_cast<double>(radius));
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
	static_cast<QGeoCircle*>(ptr)->setRadius(static_cast<double>(radius));
}

char* QGeoCircle_ToString(void* ptr){
	return static_cast<QGeoCircle*>(ptr)->toString().toUtf8().data();
}

void QGeoCircle_DestroyQGeoCircle(void* ptr){
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

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
	static_cast<QGeoPositionInfo*>(ptr)->setAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute), static_cast<double>(value));
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

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource {
public:
	void setUpdateInterval(int msec) { if (!callbackQGeoPositionInfoSourceSetUpdateInterval(this->objectName().toUtf8().data(), msec)) { QGeoPositionInfoSource::setUpdateInterval(msec); }; };
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSourceError2(this->objectName().toUtf8().data(), positioningError); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { if (!callbackQGeoPositionInfoSourceSetPreferredPositioningMethods(this->objectName().toUtf8().data(), methods)) { QGeoPositionInfoSource::setPreferredPositioningMethods(methods); }; };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSourceUpdateTimeout(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQGeoPositionInfoSourceTimerEvent(this->objectName().toUtf8().data(), event)) { QGeoPositionInfoSource::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQGeoPositionInfoSourceChildEvent(this->objectName().toUtf8().data(), event)) { QGeoPositionInfoSource::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQGeoPositionInfoSourceCustomEvent(this->objectName().toUtf8().data(), event)) { QGeoPositionInfoSource::customEvent(event); }; };
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
	return QGeoPositionInfoSource::availableSources().join(",,,").toUtf8().data();
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent){
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoPositionInfoSource_ConnectError2(void* ptr){
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&QGeoPositionInfoSource::error), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&MyQGeoPositionInfoSource::Signal_Error2));;
}

void QGeoPositionInfoSource_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&QGeoPositionInfoSource::error), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&MyQGeoPositionInfoSource::Signal_Error2));;
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

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
protected:
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

char* QGeoPositionInfoSourceFactory_ObjectNameAbs(void* ptr){
	return static_cast<MyQGeoPositionInfoSourceFactory*>(ptr)->objectNameAbs().toUtf8().data();
}

void QGeoPositionInfoSourceFactory_SetObjectNameAbs(void* ptr, char* name){
	static_cast<MyQGeoPositionInfoSourceFactory*>(ptr)->setObjectNameAbs(QString(name));
}

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
	static_cast<QGeoSatelliteInfo*>(ptr)->setAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute), static_cast<double>(value));
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

class MyQGeoSatelliteInfoSource: public QGeoSatelliteInfoSource {
public:
	void setUpdateInterval(int msec) { if (!callbackQGeoSatelliteInfoSourceSetUpdateInterval(this->objectName().toUtf8().data(), msec)) { QGeoSatelliteInfoSource::setUpdateInterval(msec); }; };
	void Signal_Error2(QGeoSatelliteInfoSource::Error satelliteError) { callbackQGeoSatelliteInfoSourceError2(this->objectName().toUtf8().data(), satelliteError); };
	void Signal_RequestTimeout() { callbackQGeoSatelliteInfoSourceRequestTimeout(this->objectName().toUtf8().data()); };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQGeoSatelliteInfoSourceTimerEvent(this->objectName().toUtf8().data(), event)) { QGeoSatelliteInfoSource::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQGeoSatelliteInfoSourceChildEvent(this->objectName().toUtf8().data(), event)) { QGeoSatelliteInfoSource::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQGeoSatelliteInfoSourceCustomEvent(this->objectName().toUtf8().data(), event)) { QGeoSatelliteInfoSource::customEvent(event); }; };
};

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

int QGeoSatelliteInfoSource_UpdateInterval(void* ptr){
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->updateInterval();
}

char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources(){
	return QGeoSatelliteInfoSource::availableSources().join(",,,").toUtf8().data();
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent){
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent){
	return QGeoSatelliteInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoSatelliteInfoSource_ConnectError2(void* ptr){
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&QGeoSatelliteInfoSource::error), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&MyQGeoSatelliteInfoSource::Signal_Error2));;
}

void QGeoSatelliteInfoSource_DisconnectError2(void* ptr){
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&QGeoSatelliteInfoSource::error), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&MyQGeoSatelliteInfoSource::Signal_Error2));;
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

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource {
public:
	MyQNmeaPositionInfoSource(UpdateMode updateMode, QObject *parent) : QNmeaPositionInfoSource(updateMode, parent) {};
	void requestUpdate(int msec) { if (!callbackQNmeaPositionInfoSourceRequestUpdate(this->objectName().toUtf8().data(), msec)) { QNmeaPositionInfoSource::requestUpdate(msec); }; };
	void setUpdateInterval(int msec) { if (!callbackQNmeaPositionInfoSourceSetUpdateInterval(this->objectName().toUtf8().data(), msec)) { QNmeaPositionInfoSource::setUpdateInterval(msec); }; };
	void startUpdates() { if (!callbackQNmeaPositionInfoSourceStartUpdates(this->objectName().toUtf8().data())) { QNmeaPositionInfoSource::startUpdates(); }; };
	void stopUpdates() { if (!callbackQNmeaPositionInfoSourceStopUpdates(this->objectName().toUtf8().data())) { QNmeaPositionInfoSource::stopUpdates(); }; };
	void setPreferredPositioningMethods(QNmeaPositionInfoSource::PositioningMethods methods) { if (!callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods(this->objectName().toUtf8().data(), methods)) { QNmeaPositionInfoSource::setPreferredPositioningMethods(methods); }; };
protected:
	void timerEvent(QTimerEvent * event) { if (!callbackQNmeaPositionInfoSourceTimerEvent(this->objectName().toUtf8().data(), event)) { QNmeaPositionInfoSource::timerEvent(event); }; };
	void childEvent(QChildEvent * event) { if (!callbackQNmeaPositionInfoSourceChildEvent(this->objectName().toUtf8().data(), event)) { QNmeaPositionInfoSource::childEvent(event); }; };
	void customEvent(QEvent * event) { if (!callbackQNmeaPositionInfoSourceCustomEvent(this->objectName().toUtf8().data(), event)) { QNmeaPositionInfoSource::customEvent(event); }; };
};

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(int updateMode, void* parent){
	return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
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

