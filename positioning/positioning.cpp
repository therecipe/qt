#define protected public

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
	void Signal_Error2(QGeoAreaMonitorSource::Error areaMonitoringError) { callbackQGeoAreaMonitorSourceError2(this, this->objectName().toUtf8().data(), areaMonitoringError); };
	void setPositionInfoSource(QGeoPositionInfoSource * newSource) { callbackQGeoAreaMonitorSourceSetPositionInfoSource(this, this->objectName().toUtf8().data(), newSource); };
	void timerEvent(QTimerEvent * event) { callbackQGeoAreaMonitorSourceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGeoAreaMonitorSourceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGeoAreaMonitorSourceCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

void QGeoAreaMonitorSource_Error2(void* ptr, int areaMonitoringError){
	static_cast<QGeoAreaMonitorSource*>(ptr)->error(static_cast<QGeoAreaMonitorSource::Error>(areaMonitoringError));
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
	static_cast<MyQGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

void QGeoAreaMonitorSource_SetPositionInfoSourceDefault(void* ptr, void* newSource){
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
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

void QGeoAreaMonitorSource_TimerEvent(void* ptr, void* event){
	static_cast<MyQGeoAreaMonitorSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoAreaMonitorSource_TimerEventDefault(void* ptr, void* event){
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoAreaMonitorSource_ChildEvent(void* ptr, void* event){
	static_cast<MyQGeoAreaMonitorSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoAreaMonitorSource_ChildEventDefault(void* ptr, void* event){
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoAreaMonitorSource_CustomEvent(void* ptr, void* event){
	static_cast<MyQGeoAreaMonitorSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoAreaMonitorSource_CustomEventDefault(void* ptr, void* event){
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::customEvent(static_cast<QEvent*>(event));
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
	void setUpdateInterval(int msec) { callbackQGeoPositionInfoSourceSetUpdateInterval(this, this->objectName().toUtf8().data(), msec); };
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSourceError2(this, this->objectName().toUtf8().data(), positioningError); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQGeoPositionInfoSourceSetPreferredPositioningMethods(this, this->objectName().toUtf8().data(), methods); };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSourceUpdateTimeout(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQGeoPositionInfoSourceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGeoPositionInfoSourceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGeoPositionInfoSourceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<MyQGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec){
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setUpdateInterval(msec);
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

void QGeoPositionInfoSource_Error2(void* ptr, int positioningError){
	static_cast<QGeoPositionInfoSource*>(ptr)->error(static_cast<QGeoPositionInfoSource::Error>(positioningError));
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
	static_cast<MyQGeoPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethodsDefault(void* ptr, int methods){
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
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

void QGeoPositionInfoSource_UpdateTimeout(void* ptr){
	static_cast<QGeoPositionInfoSource*>(ptr)->updateTimeout();
}

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(void* ptr){
	static_cast<QGeoPositionInfoSource*>(ptr)->~QGeoPositionInfoSource();
}

void QGeoPositionInfoSource_TimerEvent(void* ptr, void* event){
	static_cast<MyQGeoPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoPositionInfoSource_TimerEventDefault(void* ptr, void* event){
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoPositionInfoSource_ChildEvent(void* ptr, void* event){
	static_cast<MyQGeoPositionInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoPositionInfoSource_ChildEventDefault(void* ptr, void* event){
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoPositionInfoSource_CustomEvent(void* ptr, void* event){
	static_cast<MyQGeoPositionInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoPositionInfoSource_CustomEventDefault(void* ptr, void* event){
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::customEvent(static_cast<QEvent*>(event));
}

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory {
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
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
	if (dynamic_cast<MyQGeoPositionInfoSourceFactory*>(static_cast<QGeoPositionInfoSourceFactory*>(ptr))) {
		return static_cast<MyQGeoPositionInfoSourceFactory*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QGeoPositionInfoSourceFactory_BASE").toUtf8().data();
}

void QGeoPositionInfoSourceFactory_SetObjectNameAbs(void* ptr, char* name){
	if (dynamic_cast<MyQGeoPositionInfoSourceFactory*>(static_cast<QGeoPositionInfoSourceFactory*>(ptr))) {
		static_cast<MyQGeoPositionInfoSourceFactory*>(ptr)->setObjectNameAbs(QString(name));
	}
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
	void setUpdateInterval(int msec) { callbackQGeoSatelliteInfoSourceSetUpdateInterval(this, this->objectName().toUtf8().data(), msec); };
	void Signal_Error2(QGeoSatelliteInfoSource::Error satelliteError) { callbackQGeoSatelliteInfoSourceError2(this, this->objectName().toUtf8().data(), satelliteError); };
	void Signal_RequestTimeout() { callbackQGeoSatelliteInfoSourceRequestTimeout(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQGeoSatelliteInfoSourceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGeoSatelliteInfoSourceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQGeoSatelliteInfoSourceCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<MyQGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoSatelliteInfoSource_SetUpdateIntervalDefault(void* ptr, int msec){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::setUpdateInterval(msec);
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

void QGeoSatelliteInfoSource_Error2(void* ptr, int satelliteError){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->error(static_cast<QGeoSatelliteInfoSource::Error>(satelliteError));
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

void QGeoSatelliteInfoSource_RequestTimeout(void* ptr){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->requestTimeout();
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

void QGeoSatelliteInfoSource_TimerEvent(void* ptr, void* event){
	static_cast<MyQGeoSatelliteInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoSatelliteInfoSource_TimerEventDefault(void* ptr, void* event){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoSatelliteInfoSource_ChildEvent(void* ptr, void* event){
	static_cast<MyQGeoSatelliteInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoSatelliteInfoSource_ChildEventDefault(void* ptr, void* event){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoSatelliteInfoSource_CustomEvent(void* ptr, void* event){
	static_cast<MyQGeoSatelliteInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoSatelliteInfoSource_CustomEventDefault(void* ptr, void* event){
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::customEvent(static_cast<QEvent*>(event));
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
	void requestUpdate(int msec) { callbackQNmeaPositionInfoSourceRequestUpdate(this, this->objectName().toUtf8().data(), msec); };
	void setUpdateInterval(int msec) { callbackQNmeaPositionInfoSourceSetUpdateInterval(this, this->objectName().toUtf8().data(), msec); };
	void startUpdates() { callbackQNmeaPositionInfoSourceStartUpdates(this, this->objectName().toUtf8().data()); };
	void stopUpdates() { callbackQNmeaPositionInfoSourceStopUpdates(this, this->objectName().toUtf8().data()); };
	void setPreferredPositioningMethods(QNmeaPositionInfoSource::PositioningMethods methods) { callbackQNmeaPositionInfoSourceSetPreferredPositioningMethods(this, this->objectName().toUtf8().data(), methods); };
	void timerEvent(QTimerEvent * event) { callbackQNmeaPositionInfoSourceTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQNmeaPositionInfoSourceChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQNmeaPositionInfoSourceCustomEvent(this, this->objectName().toUtf8().data(), event); };
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

int QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(void* ptr, char* data, int size, void* posInfo, int hasFix){
	return static_cast<QNmeaPositionInfoSource*>(ptr)->parsePosInfoFromNmeaData(const_cast<const char*>(data), size, static_cast<QGeoPositionInfo*>(posInfo), NULL);
}

void QNmeaPositionInfoSource_RequestUpdate(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<MyQNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_RequestUpdateDefault(void* ptr, int msec){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_SetDevice(void* ptr, void* device){
	static_cast<QNmeaPositionInfoSource*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QNmeaPositionInfoSource_SetUpdateInterval(void* ptr, int msec){
	static_cast<MyQNmeaPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec){
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_StartUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StartUpdatesDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StopUpdates(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQNmeaPositionInfoSource*>(ptr), "stopUpdates");
}

void QNmeaPositionInfoSource_StopUpdatesDefault(void* ptr){
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

void QNmeaPositionInfoSource_SetPreferredPositioningMethods(void* ptr, int methods){
	static_cast<MyQNmeaPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QNmeaPositionInfoSource::PositioningMethod>(methods));
}

void QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(void* ptr, int methods){
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setPreferredPositioningMethods(static_cast<QNmeaPositionInfoSource::PositioningMethod>(methods));
}

void QNmeaPositionInfoSource_TimerEvent(void* ptr, void* event){
	static_cast<MyQNmeaPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNmeaPositionInfoSource_TimerEventDefault(void* ptr, void* event){
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNmeaPositionInfoSource_ChildEvent(void* ptr, void* event){
	static_cast<MyQNmeaPositionInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNmeaPositionInfoSource_ChildEventDefault(void* ptr, void* event){
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QNmeaPositionInfoSource_CustomEvent(void* ptr, void* event){
	static_cast<MyQNmeaPositionInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNmeaPositionInfoSource_CustomEventDefault(void* ptr, void* event){
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::customEvent(static_cast<QEvent*>(event));
}

