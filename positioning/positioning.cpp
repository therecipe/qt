// +build !minimal

#define protected public
#define private public

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
#include <QMetaMethod>
#include <QMetaObject>
#include <QNmeaPositionInfoSource>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>

void* QGeoAddress_NewQGeoAddress()
{
	return new QGeoAddress();
}

void* QGeoAddress_NewQGeoAddress2(void* other)
{
	return new QGeoAddress(*static_cast<QGeoAddress*>(other));
}

char* QGeoAddress_City(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->city().toUtf8().constData());
}

void QGeoAddress_Clear(void* ptr)
{
	static_cast<QGeoAddress*>(ptr)->clear();
}

char* QGeoAddress_Country(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->country().toUtf8().constData());
}

char* QGeoAddress_CountryCode(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->countryCode().toUtf8().constData());
}

char* QGeoAddress_County(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->county().toUtf8().constData());
}

char* QGeoAddress_District(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->district().toUtf8().constData());
}

char QGeoAddress_IsEmpty(void* ptr)
{
	return static_cast<QGeoAddress*>(ptr)->isEmpty();
}

char QGeoAddress_IsTextGenerated(void* ptr)
{
	return static_cast<QGeoAddress*>(ptr)->isTextGenerated();
}

char* QGeoAddress_PostalCode(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->postalCode().toUtf8().constData());
}

void QGeoAddress_SetCity(void* ptr, char* city)
{
	static_cast<QGeoAddress*>(ptr)->setCity(QString(city));
}

void QGeoAddress_SetCountry(void* ptr, char* country)
{
	static_cast<QGeoAddress*>(ptr)->setCountry(QString(country));
}

void QGeoAddress_SetCountryCode(void* ptr, char* countryCode)
{
	static_cast<QGeoAddress*>(ptr)->setCountryCode(QString(countryCode));
}

void QGeoAddress_SetCounty(void* ptr, char* county)
{
	static_cast<QGeoAddress*>(ptr)->setCounty(QString(county));
}

void QGeoAddress_SetDistrict(void* ptr, char* district)
{
	static_cast<QGeoAddress*>(ptr)->setDistrict(QString(district));
}

void QGeoAddress_SetPostalCode(void* ptr, char* postalCode)
{
	static_cast<QGeoAddress*>(ptr)->setPostalCode(QString(postalCode));
}

void QGeoAddress_SetState(void* ptr, char* state)
{
	static_cast<QGeoAddress*>(ptr)->setState(QString(state));
}

void QGeoAddress_SetStreet(void* ptr, char* street)
{
	static_cast<QGeoAddress*>(ptr)->setStreet(QString(street));
}

void QGeoAddress_SetText(void* ptr, char* text)
{
	static_cast<QGeoAddress*>(ptr)->setText(QString(text));
}

char* QGeoAddress_State(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->state().toUtf8().constData());
}

char* QGeoAddress_Street(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->street().toUtf8().constData());
}

char* QGeoAddress_Text(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAddress*>(ptr)->text().toUtf8().constData());
}

void QGeoAddress_DestroyQGeoAddress(void* ptr)
{
	static_cast<QGeoAddress*>(ptr)->~QGeoAddress();
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(void* other)
{
	return new QGeoAreaMonitorInfo(*static_cast<QGeoAreaMonitorInfo*>(other));
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name)
{
	return new QGeoAreaMonitorInfo(QString(name));
}

void* QGeoAreaMonitorInfo_Area(void* ptr)
{
	return new QGeoShape(static_cast<QGeoAreaMonitorInfo*>(ptr)->area());
}

void* QGeoAreaMonitorInfo_Expiration(void* ptr)
{
	return new QDateTime(static_cast<QGeoAreaMonitorInfo*>(ptr)->expiration());
}

char* QGeoAreaMonitorInfo_Identifier(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAreaMonitorInfo*>(ptr)->identifier().toUtf8().constData());
}

char QGeoAreaMonitorInfo_IsPersistent(void* ptr)
{
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isPersistent();
}

char QGeoAreaMonitorInfo_IsValid(void* ptr)
{
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isValid();
}

char* QGeoAreaMonitorInfo_Name(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAreaMonitorInfo*>(ptr)->name().toUtf8().constData());
}

void QGeoAreaMonitorInfo_SetArea(void* ptr, void* newShape)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setArea(*static_cast<QGeoShape*>(newShape));
}

void QGeoAreaMonitorInfo_SetExpiration(void* ptr, void* expiry)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setExpiration(*static_cast<QDateTime*>(expiry));
}

void QGeoAreaMonitorInfo_SetName(void* ptr, char* name)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setName(QString(name));
}

void QGeoAreaMonitorInfo_SetPersistent(void* ptr, char isPersistent)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setPersistent(isPersistent != 0);
}

void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(void* ptr)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->~QGeoAreaMonitorInfo();
}

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource
{
public:
	void Signal_AreaEntered(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaEntered(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_AreaExited(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaExited(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_Error2(QGeoAreaMonitorSource::Error areaMonitoringError) { callbackQGeoAreaMonitorSource_Error2(this, areaMonitoringError); };
	Error error() const { return static_cast<QGeoAreaMonitorSource::Error>(callbackQGeoAreaMonitorSource_Error(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	void Signal_MonitorExpired(const QGeoAreaMonitorInfo & monitor) { callbackQGeoAreaMonitorSource_MonitorExpired(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)); };
	QGeoPositionInfoSource * positionInfoSource() const { return static_cast<QGeoPositionInfoSource*>(callbackQGeoAreaMonitorSource_PositionInfoSource(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	bool requestUpdate(const QGeoAreaMonitorInfo & monitor, const char * sign) { return callbackQGeoAreaMonitorSource_RequestUpdate(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<char*>(QString(sign).toUtf8().constData())) != 0; };
	void setPositionInfoSource(QGeoPositionInfoSource * newSource) { callbackQGeoAreaMonitorSource_SetPositionInfoSource(this, newSource); };
	bool startMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StartMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	bool stopMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StopMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	AreaMonitorFeatures supportedAreaMonitorFeatures() const { return static_cast<QGeoAreaMonitorSource::AreaMonitorFeature>(callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	void timerEvent(QTimerEvent * event) { callbackQGeoAreaMonitorSource_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoAreaMonitorSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoAreaMonitorSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoAreaMonitorSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoAreaMonitorSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoAreaMonitorSource_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoAreaMonitorSource_MetaObject(const_cast<MyQGeoAreaMonitorSource*>(this))); };
};

void QGeoAreaMonitorSource_ConnectAreaEntered(void* ptr)
{
	QObject::connect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&QGeoAreaMonitorSource::areaEntered), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&MyQGeoAreaMonitorSource::Signal_AreaEntered));
}

void QGeoAreaMonitorSource_DisconnectAreaEntered(void* ptr)
{
	QObject::disconnect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&QGeoAreaMonitorSource::areaEntered), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&MyQGeoAreaMonitorSource::Signal_AreaEntered));
}

void QGeoAreaMonitorSource_AreaEntered(void* ptr, void* monitor, void* update)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->areaEntered(*static_cast<QGeoAreaMonitorInfo*>(monitor), *static_cast<QGeoPositionInfo*>(update));
}

void QGeoAreaMonitorSource_ConnectAreaExited(void* ptr)
{
	QObject::connect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&QGeoAreaMonitorSource::areaExited), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&MyQGeoAreaMonitorSource::Signal_AreaExited));
}

void QGeoAreaMonitorSource_DisconnectAreaExited(void* ptr)
{
	QObject::disconnect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&QGeoAreaMonitorSource::areaExited), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &, const QGeoPositionInfo &)>(&MyQGeoAreaMonitorSource::Signal_AreaExited));
}

void QGeoAreaMonitorSource_AreaExited(void* ptr, void* monitor, void* update)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->areaExited(*static_cast<QGeoAreaMonitorInfo*>(monitor), *static_cast<QGeoPositionInfo*>(update));
}

char* QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()
{
	return const_cast<char*>(QGeoAreaMonitorSource::availableSources().join("|").toUtf8().constData());
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent)
{
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoAreaMonitorSource_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&QGeoAreaMonitorSource::error), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&MyQGeoAreaMonitorSource::Signal_Error2));
}

void QGeoAreaMonitorSource_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&QGeoAreaMonitorSource::error), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(QGeoAreaMonitorSource::Error)>(&MyQGeoAreaMonitorSource::Signal_Error2));
}

void QGeoAreaMonitorSource_Error2(void* ptr, long long areaMonitoringError)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->error(static_cast<QGeoAreaMonitorSource::Error>(areaMonitoringError));
}

long long QGeoAreaMonitorSource_Error(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->error();
}

void QGeoAreaMonitorSource_ConnectMonitorExpired(void* ptr)
{
	QObject::connect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &)>(&QGeoAreaMonitorSource::monitorExpired), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &)>(&MyQGeoAreaMonitorSource::Signal_MonitorExpired));
}

void QGeoAreaMonitorSource_DisconnectMonitorExpired(void* ptr)
{
	QObject::disconnect(static_cast<QGeoAreaMonitorSource*>(ptr), static_cast<void (QGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &)>(&QGeoAreaMonitorSource::monitorExpired), static_cast<MyQGeoAreaMonitorSource*>(ptr), static_cast<void (MyQGeoAreaMonitorSource::*)(const QGeoAreaMonitorInfo &)>(&MyQGeoAreaMonitorSource::Signal_MonitorExpired));
}

void QGeoAreaMonitorSource_MonitorExpired(void* ptr, void* monitor)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->monitorExpired(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

void* QGeoAreaMonitorSource_PositionInfoSource(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->positionInfoSource();
}

void* QGeoAreaMonitorSource_PositionInfoSourceDefault(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::positionInfoSource();
}

char QGeoAreaMonitorSource_RequestUpdate(void* ptr, void* monitor, char* sign)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->requestUpdate(*static_cast<QGeoAreaMonitorInfo*>(monitor), const_cast<const char*>(sign));
}

void QGeoAreaMonitorSource_SetPositionInfoSource(void* ptr, void* newSource)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

void QGeoAreaMonitorSource_SetPositionInfoSourceDefault(void* ptr, void* newSource)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

char* QGeoAreaMonitorSource_SourceName(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoAreaMonitorSource*>(ptr)->sourceName().toUtf8().constData());
}

char QGeoAreaMonitorSource_StartMonitoring(void* ptr, void* monitor)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->startMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

char QGeoAreaMonitorSource_StopMonitoring(void* ptr, void* monitor)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->stopMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

long long QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->supportedAreaMonitorFeatures();
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(void* ptr)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->~QGeoAreaMonitorSource();
}

void QGeoAreaMonitorSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoAreaMonitorSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoAreaMonitorSource_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoAreaMonitorSource_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoAreaMonitorSource_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoAreaMonitorSource_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoAreaMonitorSource_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoAreaMonitorSource_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::customEvent(static_cast<QEvent*>(event));
}

void QGeoAreaMonitorSource_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoAreaMonitorSource*>(ptr), "deleteLater");
}

void QGeoAreaMonitorSource_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::deleteLater();
}

void QGeoAreaMonitorSource_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoAreaMonitorSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoAreaMonitorSource_Event(void* ptr, void* e)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoAreaMonitorSource_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::event(static_cast<QEvent*>(e));
}

char QGeoAreaMonitorSource_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoAreaMonitorSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoAreaMonitorSource_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoAreaMonitorSource*>(ptr)->metaObject());
}

void* QGeoAreaMonitorSource_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::metaObject());
}

void* QGeoCircle_NewQGeoCircle()
{
	return new QGeoCircle();
}

void* QGeoCircle_NewQGeoCircle3(void* other)
{
	return new QGeoCircle(*static_cast<QGeoCircle*>(other));
}

void* QGeoCircle_NewQGeoCircle2(void* center, double radius)
{
	return new QGeoCircle(*static_cast<QGeoCoordinate*>(center), radius);
}

void* QGeoCircle_NewQGeoCircle4(void* other)
{
	return new QGeoCircle(*static_cast<QGeoShape*>(other));
}

void* QGeoCircle_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoCircle*>(ptr)->center());
}

double QGeoCircle_Radius(void* ptr)
{
	return static_cast<QGeoCircle*>(ptr)->radius();
}

void QGeoCircle_SetCenter(void* ptr, void* center)
{
	static_cast<QGeoCircle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoCircle_SetRadius(void* ptr, double radius)
{
	static_cast<QGeoCircle*>(ptr)->setRadius(radius);
}

char* QGeoCircle_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoCircle*>(ptr)->toString().toUtf8().constData());
}

void QGeoCircle_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoCircle*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void* QGeoCircle_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoCircle(static_cast<QGeoCircle*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

void QGeoCircle_DestroyQGeoCircle(void* ptr)
{
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

void* QGeoCoordinate_NewQGeoCoordinate()
{
	return new QGeoCoordinate();
}

void* QGeoCoordinate_NewQGeoCoordinate4(void* other)
{
	return new QGeoCoordinate(*static_cast<QGeoCoordinate*>(other));
}

void* QGeoCoordinate_NewQGeoCoordinate2(double latitude, double longitude)
{
	return new QGeoCoordinate(latitude, longitude);
}

void* QGeoCoordinate_NewQGeoCoordinate3(double latitude, double longitude, double altitude)
{
	return new QGeoCoordinate(latitude, longitude, altitude);
}

double QGeoCoordinate_Altitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->altitude();
}

void* QGeoCoordinate_AtDistanceAndAzimuth(void* ptr, double distance, double azimuth, double distanceUp)
{
	return new QGeoCoordinate(static_cast<QGeoCoordinate*>(ptr)->atDistanceAndAzimuth(distance, azimuth, distanceUp));
}

double QGeoCoordinate_AzimuthTo(void* ptr, void* other)
{
	return static_cast<QGeoCoordinate*>(ptr)->azimuthTo(*static_cast<QGeoCoordinate*>(other));
}

double QGeoCoordinate_DistanceTo(void* ptr, void* other)
{
	return static_cast<QGeoCoordinate*>(ptr)->distanceTo(*static_cast<QGeoCoordinate*>(other));
}

char QGeoCoordinate_IsValid(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->isValid();
}

double QGeoCoordinate_Latitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->latitude();
}

double QGeoCoordinate_Longitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->longitude();
}

void QGeoCoordinate_SetAltitude(void* ptr, double altitude)
{
	static_cast<QGeoCoordinate*>(ptr)->setAltitude(altitude);
}

void QGeoCoordinate_SetLatitude(void* ptr, double latitude)
{
	static_cast<QGeoCoordinate*>(ptr)->setLatitude(latitude);
}

void QGeoCoordinate_SetLongitude(void* ptr, double longitude)
{
	static_cast<QGeoCoordinate*>(ptr)->setLongitude(longitude);
}

char* QGeoCoordinate_ToString(void* ptr, long long format)
{
	return const_cast<char*>(static_cast<QGeoCoordinate*>(ptr)->toString(static_cast<QGeoCoordinate::CoordinateFormat>(format)).toUtf8().constData());
}

long long QGeoCoordinate_Type(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->type();
}

void QGeoCoordinate_DestroyQGeoCoordinate(void* ptr)
{
	static_cast<QGeoCoordinate*>(ptr)->~QGeoCoordinate();
}

void* QGeoPositionInfo_NewQGeoPositionInfo()
{
	return new QGeoPositionInfo();
}

void* QGeoPositionInfo_NewQGeoPositionInfo2(void* coordinate, void* timestamp)
{
	return new QGeoPositionInfo(*static_cast<QGeoCoordinate*>(coordinate), *static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_NewQGeoPositionInfo3(void* other)
{
	return new QGeoPositionInfo(*static_cast<QGeoPositionInfo*>(other));
}

double QGeoPositionInfo_Attribute(void* ptr, long long attribute)
{
	return static_cast<QGeoPositionInfo*>(ptr)->attribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

void* QGeoPositionInfo_Coordinate(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoPositionInfo*>(ptr)->coordinate());
}

char QGeoPositionInfo_HasAttribute(void* ptr, long long attribute)
{
	return static_cast<QGeoPositionInfo*>(ptr)->hasAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

char QGeoPositionInfo_IsValid(void* ptr)
{
	return static_cast<QGeoPositionInfo*>(ptr)->isValid();
}

void QGeoPositionInfo_RemoveAttribute(void* ptr, long long attribute)
{
	static_cast<QGeoPositionInfo*>(ptr)->removeAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

void QGeoPositionInfo_SetAttribute(void* ptr, long long attribute, double value)
{
	static_cast<QGeoPositionInfo*>(ptr)->setAttribute(static_cast<QGeoPositionInfo::Attribute>(attribute), value);
}

void QGeoPositionInfo_SetCoordinate(void* ptr, void* coordinate)
{
	static_cast<QGeoPositionInfo*>(ptr)->setCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPositionInfo_SetTimestamp(void* ptr, void* timestamp)
{
	static_cast<QGeoPositionInfo*>(ptr)->setTimestamp(*static_cast<QDateTime*>(timestamp));
}

void* QGeoPositionInfo_Timestamp(void* ptr)
{
	return new QDateTime(static_cast<QGeoPositionInfo*>(ptr)->timestamp());
}

void QGeoPositionInfo_DestroyQGeoPositionInfo(void* ptr)
{
	static_cast<QGeoPositionInfo*>(ptr)->~QGeoPositionInfo();
}

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource
{
public:
	MyQGeoPositionInfoSource(QObject *parent) : QGeoPositionInfoSource(parent) {};
	void setUpdateInterval(int msec) { callbackQGeoPositionInfoSource_SetUpdateInterval(this, msec); };
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSource_Error2(this, positioningError); };
	Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQGeoPositionInfoSource_Error(const_cast<MyQGeoPositionInfoSource*>(this))); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQGeoPositionInfoSource_LastKnownPosition(const_cast<MyQGeoPositionInfoSource*>(this), fromSatellitePositioningMethodsOnly)); };
	int minimumUpdateInterval() const { return callbackQGeoPositionInfoSource_MinimumUpdateInterval(const_cast<MyQGeoPositionInfoSource*>(this)); };
	void Signal_PositionUpdated(const QGeoPositionInfo & update) { callbackQGeoPositionInfoSource_PositionUpdated(this, const_cast<QGeoPositionInfo*>(&update)); };
	void requestUpdate(int timeout) { callbackQGeoPositionInfoSource_RequestUpdate(this, timeout); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	void startUpdates() { callbackQGeoPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoPositionInfoSource_StopUpdates(this); };
	PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQGeoPositionInfoSource_SupportedPositioningMethods(const_cast<MyQGeoPositionInfoSource*>(this))); };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSource_UpdateTimeout(this); };
	void timerEvent(QTimerEvent * event) { callbackQGeoPositionInfoSource_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoPositionInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoPositionInfoSource_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoPositionInfoSource_MetaObject(const_cast<MyQGeoPositionInfoSource*>(this))); };
};

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setUpdateInterval(msec);
}

char* QGeoPositionInfoSource_SourceName(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8().constData());
}

int QGeoPositionInfoSource_UpdateInterval(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->updateInterval();
}

void* QGeoPositionInfoSource_NewQGeoPositionInfoSource(void* parent)
{
	return new MyQGeoPositionInfoSource(static_cast<QObject*>(parent));
}

char* QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()
{
	return const_cast<char*>(QGeoPositionInfoSource::availableSources().join("|").toUtf8().constData());
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoPositionInfoSource_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&QGeoPositionInfoSource::error), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&MyQGeoPositionInfoSource::Signal_Error2));
}

void QGeoPositionInfoSource_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&QGeoPositionInfoSource::error), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(QGeoPositionInfoSource::Error)>(&MyQGeoPositionInfoSource::Signal_Error2));
}

void QGeoPositionInfoSource_Error2(void* ptr, long long positioningError)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->error(static_cast<QGeoPositionInfoSource::Error>(positioningError));
}

long long QGeoPositionInfoSource_Error(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->error();
}

void* QGeoPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QGeoPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

int QGeoPositionInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

void QGeoPositionInfoSource_ConnectPositionUpdated(void* ptr)
{
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(const QGeoPositionInfo &)>(&QGeoPositionInfoSource::positionUpdated), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(const QGeoPositionInfo &)>(&MyQGeoPositionInfoSource::Signal_PositionUpdated));
}

void QGeoPositionInfoSource_DisconnectPositionUpdated(void* ptr)
{
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)(const QGeoPositionInfo &)>(&QGeoPositionInfoSource::positionUpdated), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)(const QGeoPositionInfo &)>(&MyQGeoPositionInfoSource::Signal_PositionUpdated));
}

void QGeoPositionInfoSource_PositionUpdated(void* ptr, void* update)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->positionUpdated(*static_cast<QGeoPositionInfo*>(update));
}

long long QGeoPositionInfoSource_PreferredPositioningMethods(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->preferredPositioningMethods();
}

void QGeoPositionInfoSource_RequestUpdate(void* ptr, int timeout)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethods(void* ptr, long long methods)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_SetPreferredPositioningMethodsDefault(void* ptr, long long methods)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QGeoPositionInfoSource_StartUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "startUpdates");
}

void QGeoPositionInfoSource_StopUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "stopUpdates");
}

long long QGeoPositionInfoSource_SupportedPositioningMethods(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

void QGeoPositionInfoSource_ConnectUpdateTimeout(void* ptr)
{
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));
}

void QGeoPositionInfoSource_DisconnectUpdateTimeout(void* ptr)
{
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::updateTimeout), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_UpdateTimeout));
}

void QGeoPositionInfoSource_UpdateTimeout(void* ptr)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->updateTimeout();
}

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSource(void* ptr)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->~QGeoPositionInfoSource();
}

void QGeoPositionInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoPositionInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoPositionInfoSource_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoPositionInfoSource_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoPositionInfoSource_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoPositionInfoSource_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoPositionInfoSource_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoPositionInfoSource_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::customEvent(static_cast<QEvent*>(event));
}

void QGeoPositionInfoSource_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "deleteLater");
}

void QGeoPositionInfoSource_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::deleteLater();
}

void QGeoPositionInfoSource_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoPositionInfoSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoPositionInfoSource_Event(void* ptr, void* e)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoPositionInfoSource_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::event(static_cast<QEvent*>(e));
}

char QGeoPositionInfoSource_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoPositionInfoSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoPositionInfoSource_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoPositionInfoSource*>(ptr)->metaObject());
}

void* QGeoPositionInfoSource_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::metaObject());
}

class MyQGeoPositionInfoSourceFactory: public QGeoPositionInfoSourceFactory
{
public:
	QGeoAreaMonitorSource * areaMonitor(QObject * parent) { return static_cast<QGeoAreaMonitorSource*>(callbackQGeoPositionInfoSourceFactory_AreaMonitor(this, parent)); };
	QGeoPositionInfoSource * positionInfoSource(QObject * parent) { return static_cast<QGeoPositionInfoSource*>(callbackQGeoPositionInfoSourceFactory_PositionInfoSource(this, parent)); };
	QGeoSatelliteInfoSource * satelliteInfoSource(QObject * parent) { return static_cast<QGeoSatelliteInfoSource*>(callbackQGeoPositionInfoSourceFactory_SatelliteInfoSource(this, parent)); };
};

void* QGeoPositionInfoSourceFactory_AreaMonitor(void* ptr, void* parent)
{
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->areaMonitor(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_PositionInfoSource(void* ptr, void* parent)
{
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->positionInfoSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSourceFactory_SatelliteInfoSource(void* ptr, void* parent)
{
	return static_cast<QGeoPositionInfoSourceFactory*>(ptr)->satelliteInfoSource(static_cast<QObject*>(parent));
}

void QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(void* ptr)
{
	static_cast<QGeoPositionInfoSourceFactory*>(ptr)->~QGeoPositionInfoSourceFactory();
}

void* QGeoRectangle_NewQGeoRectangle()
{
	return new QGeoRectangle();
}

void* QGeoRectangle_NewQGeoRectangle2(void* center, double degreesWidth, double degreesHeight)
{
	return new QGeoRectangle(*static_cast<QGeoCoordinate*>(center), degreesWidth, degreesHeight);
}

void* QGeoRectangle_NewQGeoRectangle3(void* topLeft, void* bottomRight)
{
	return new QGeoRectangle(*static_cast<QGeoCoordinate*>(topLeft), *static_cast<QGeoCoordinate*>(bottomRight));
}

void* QGeoRectangle_NewQGeoRectangle5(void* other)
{
	return new QGeoRectangle(*static_cast<QGeoRectangle*>(other));
}

void* QGeoRectangle_NewQGeoRectangle6(void* other)
{
	return new QGeoRectangle(*static_cast<QGeoShape*>(other));
}

void* QGeoRectangle_BottomLeft(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->bottomLeft());
}

void* QGeoRectangle_BottomRight(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->bottomRight());
}

void* QGeoRectangle_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->center());
}

char QGeoRectangle_Contains(void* ptr, void* rectangle)
{
	return static_cast<QGeoRectangle*>(ptr)->contains(*static_cast<QGeoRectangle*>(rectangle));
}

double QGeoRectangle_Height(void* ptr)
{
	return static_cast<QGeoRectangle*>(ptr)->height();
}

char QGeoRectangle_Intersects(void* ptr, void* rectangle)
{
	return static_cast<QGeoRectangle*>(ptr)->intersects(*static_cast<QGeoRectangle*>(rectangle));
}

void QGeoRectangle_SetBottomLeft(void* ptr, void* bottomLeft)
{
	static_cast<QGeoRectangle*>(ptr)->setBottomLeft(*static_cast<QGeoCoordinate*>(bottomLeft));
}

void QGeoRectangle_SetBottomRight(void* ptr, void* bottomRight)
{
	static_cast<QGeoRectangle*>(ptr)->setBottomRight(*static_cast<QGeoCoordinate*>(bottomRight));
}

void QGeoRectangle_SetCenter(void* ptr, void* center)
{
	static_cast<QGeoRectangle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoRectangle_SetHeight(void* ptr, double degreesHeight)
{
	static_cast<QGeoRectangle*>(ptr)->setHeight(degreesHeight);
}

void QGeoRectangle_SetTopLeft(void* ptr, void* topLeft)
{
	static_cast<QGeoRectangle*>(ptr)->setTopLeft(*static_cast<QGeoCoordinate*>(topLeft));
}

void QGeoRectangle_SetTopRight(void* ptr, void* topRight)
{
	static_cast<QGeoRectangle*>(ptr)->setTopRight(*static_cast<QGeoCoordinate*>(topRight));
}

void QGeoRectangle_SetWidth(void* ptr, double degreesWidth)
{
	static_cast<QGeoRectangle*>(ptr)->setWidth(degreesWidth);
}

char* QGeoRectangle_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoRectangle*>(ptr)->toString().toUtf8().constData());
}

void* QGeoRectangle_TopLeft(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->topLeft());
}

void* QGeoRectangle_TopRight(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->topRight());
}

void QGeoRectangle_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoRectangle*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void* QGeoRectangle_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoRectangle(static_cast<QGeoRectangle*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

void* QGeoRectangle_United(void* ptr, void* rectangle)
{
	return new QGeoRectangle(static_cast<QGeoRectangle*>(ptr)->united(*static_cast<QGeoRectangle*>(rectangle)));
}

double QGeoRectangle_Width(void* ptr)
{
	return static_cast<QGeoRectangle*>(ptr)->width();
}

void QGeoRectangle_DestroyQGeoRectangle(void* ptr)
{
	static_cast<QGeoRectangle*>(ptr)->~QGeoRectangle();
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo()
{
	return new QGeoSatelliteInfo();
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo2(void* other)
{
	return new QGeoSatelliteInfo(*static_cast<QGeoSatelliteInfo*>(other));
}

double QGeoSatelliteInfo_Attribute(void* ptr, long long attribute)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->attribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

char QGeoSatelliteInfo_HasAttribute(void* ptr, long long attribute)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->hasAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

void QGeoSatelliteInfo_RemoveAttribute(void* ptr, long long attribute)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->removeAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

int QGeoSatelliteInfo_SatelliteIdentifier(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteIdentifier();
}

long long QGeoSatelliteInfo_SatelliteSystem(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteSystem();
}

void QGeoSatelliteInfo_SetAttribute(void* ptr, long long attribute, double value)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->setAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute), value);
}

void QGeoSatelliteInfo_SetSatelliteIdentifier(void* ptr, int satId)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteIdentifier(satId);
}

void QGeoSatelliteInfo_SetSatelliteSystem(void* ptr, long long system)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->setSatelliteSystem(static_cast<QGeoSatelliteInfo::SatelliteSystem>(system));
}

void QGeoSatelliteInfo_SetSignalStrength(void* ptr, int signalStrength)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->setSignalStrength(signalStrength);
}

int QGeoSatelliteInfo_SignalStrength(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->signalStrength();
}

void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(void* ptr)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->~QGeoSatelliteInfo();
}

class MyQGeoSatelliteInfoSource: public QGeoSatelliteInfoSource
{
public:
	MyQGeoSatelliteInfoSource(QObject *parent) : QGeoSatelliteInfoSource(parent) {};
	void setUpdateInterval(int msec) { callbackQGeoSatelliteInfoSource_SetUpdateInterval(this, msec); };
	void Signal_Error2(QGeoSatelliteInfoSource::Error satelliteError) { callbackQGeoSatelliteInfoSource_Error2(this, satelliteError); };
	Error error() const { return static_cast<QGeoSatelliteInfoSource::Error>(callbackQGeoSatelliteInfoSource_Error(const_cast<MyQGeoSatelliteInfoSource*>(this))); };
	int minimumUpdateInterval() const { return callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(const_cast<MyQGeoSatelliteInfoSource*>(this)); };
	void Signal_RequestTimeout() { callbackQGeoSatelliteInfoSource_RequestTimeout(this); };
	void requestUpdate(int timeout) { callbackQGeoSatelliteInfoSource_RequestUpdate(this, timeout); };
	void startUpdates() { callbackQGeoSatelliteInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoSatelliteInfoSource_StopUpdates(this); };
	void timerEvent(QTimerEvent * event) { callbackQGeoSatelliteInfoSource_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQGeoSatelliteInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoSatelliteInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoSatelliteInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQGeoSatelliteInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoSatelliteInfoSource_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoSatelliteInfoSource_MetaObject(const_cast<MyQGeoSatelliteInfoSource*>(this))); };
};

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoSatelliteInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::setUpdateInterval(msec);
}

int QGeoSatelliteInfoSource_UpdateInterval(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->updateInterval();
}

void* QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(void* parent)
{
	return new MyQGeoSatelliteInfoSource(static_cast<QObject*>(parent));
}

char* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()
{
	return const_cast<char*>(QGeoSatelliteInfoSource::availableSources().join("|").toUtf8().constData());
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoSatelliteInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void QGeoSatelliteInfoSource_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&QGeoSatelliteInfoSource::error), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&MyQGeoSatelliteInfoSource::Signal_Error2));
}

void QGeoSatelliteInfoSource_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&QGeoSatelliteInfoSource::error), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(QGeoSatelliteInfoSource::Error)>(&MyQGeoSatelliteInfoSource::Signal_Error2));
}

void QGeoSatelliteInfoSource_Error2(void* ptr, long long satelliteError)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->error(static_cast<QGeoSatelliteInfoSource::Error>(satelliteError));
}

long long QGeoSatelliteInfoSource_Error(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->error();
}

int QGeoSatelliteInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->minimumUpdateInterval();
}

void QGeoSatelliteInfoSource_ConnectRequestTimeout(void* ptr)
{
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));
}

void QGeoSatelliteInfoSource_DisconnectRequestTimeout(void* ptr)
{
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)()>(&QGeoSatelliteInfoSource::requestTimeout), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)()>(&MyQGeoSatelliteInfoSource::Signal_RequestTimeout));
}

void QGeoSatelliteInfoSource_RequestTimeout(void* ptr)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->requestTimeout();
}

void QGeoSatelliteInfoSource_RequestUpdate(void* ptr, int timeout)
{
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "requestUpdate", Q_ARG(int, timeout));
}

char* QGeoSatelliteInfoSource_SourceName(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoSatelliteInfoSource*>(ptr)->sourceName().toUtf8().constData());
}

void QGeoSatelliteInfoSource_StartUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "startUpdates");
}

void QGeoSatelliteInfoSource_StopUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "stopUpdates");
}

void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(void* ptr)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->~QGeoSatelliteInfoSource();
}

void QGeoSatelliteInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoSatelliteInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoSatelliteInfoSource_ChildEvent(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGeoSatelliteInfoSource_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoSatelliteInfoSource_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoSatelliteInfoSource_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoSatelliteInfoSource_CustomEvent(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGeoSatelliteInfoSource_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::customEvent(static_cast<QEvent*>(event));
}

void QGeoSatelliteInfoSource_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoSatelliteInfoSource*>(ptr), "deleteLater");
}

void QGeoSatelliteInfoSource_DeleteLaterDefault(void* ptr)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::deleteLater();
}

void QGeoSatelliteInfoSource_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoSatelliteInfoSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGeoSatelliteInfoSource_Event(void* ptr, void* e)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->event(static_cast<QEvent*>(e));
}

char QGeoSatelliteInfoSource_EventDefault(void* ptr, void* e)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::event(static_cast<QEvent*>(e));
}

char QGeoSatelliteInfoSource_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QGeoSatelliteInfoSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGeoSatelliteInfoSource_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoSatelliteInfoSource*>(ptr)->metaObject());
}

void* QGeoSatelliteInfoSource_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::metaObject());
}

void* QGeoShape_NewQGeoShape()
{
	return new QGeoShape();
}

void* QGeoShape_NewQGeoShape2(void* other)
{
	return new QGeoShape(*static_cast<QGeoShape*>(other));
}

void* QGeoShape_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoShape*>(ptr)->center());
}

char QGeoShape_Contains(void* ptr, void* coordinate)
{
	return static_cast<QGeoShape*>(ptr)->contains(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoShape_ExtendShape(void* ptr, void* coordinate)
{
	static_cast<QGeoShape*>(ptr)->extendShape(*static_cast<QGeoCoordinate*>(coordinate));
}

char QGeoShape_IsEmpty(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->isEmpty();
}

char QGeoShape_IsValid(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->isValid();
}

char* QGeoShape_ToString(void* ptr)
{
	return const_cast<char*>(static_cast<QGeoShape*>(ptr)->toString().toUtf8().constData());
}

long long QGeoShape_Type(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->type();
}

void QGeoShape_DestroyQGeoShape(void* ptr)
{
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource
{
public:
	MyQNmeaPositionInfoSource(UpdateMode updateMode, QObject *parent) : QNmeaPositionInfoSource(updateMode, parent) {};
	Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQNmeaPositionInfoSource_Error(const_cast<MyQNmeaPositionInfoSource*>(this))); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQNmeaPositionInfoSource_LastKnownPosition(const_cast<MyQNmeaPositionInfoSource*>(this), fromSatellitePositioningMethodsOnly)); };
	int minimumUpdateInterval() const { return callbackQNmeaPositionInfoSource_MinimumUpdateInterval(const_cast<MyQNmeaPositionInfoSource*>(this)); };
	bool parsePosInfoFromNmeaData(const char * data, int size, QGeoPositionInfo * posInfo, bool * hasFix) { return callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(this, const_cast<char*>(QString(data).toUtf8().constData()), size, posInfo, *hasFix) != 0; };
	void requestUpdate(int msec) { callbackQNmeaPositionInfoSource_RequestUpdate(this, msec); };
	void setUpdateInterval(int msec) { callbackQNmeaPositionInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQNmeaPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQNmeaPositionInfoSource_StopUpdates(this); };
	PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQNmeaPositionInfoSource_SupportedPositioningMethods(const_cast<MyQNmeaPositionInfoSource*>(this))); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQNmeaPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	void timerEvent(QTimerEvent * event) { callbackQNmeaPositionInfoSource_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackQNmeaPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNmeaPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNmeaPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQNmeaPositionInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNmeaPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQNmeaPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNmeaPositionInfoSource_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNmeaPositionInfoSource_MetaObject(const_cast<MyQNmeaPositionInfoSource*>(this))); };
};

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(long long updateMode, void* parent)
{
	return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
}

void* QNmeaPositionInfoSource_Device(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->device();
}

long long QNmeaPositionInfoSource_Error(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->error();
}

long long QNmeaPositionInfoSource_ErrorDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::error();
}

void* QNmeaPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

void* QNmeaPositionInfoSource_LastKnownPositionDefault(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

int QNmeaPositionInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QNmeaPositionInfoSource_MinimumUpdateIntervalDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::minimumUpdateInterval();
}

char QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(void* ptr, char* data, int size, void* posInfo, char hasFix)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->parsePosInfoFromNmeaData(const_cast<const char*>(data), size, static_cast<QGeoPositionInfo*>(posInfo), NULL);
}

char QNmeaPositionInfoSource_ParsePosInfoFromNmeaDataDefault(void* ptr, char* data, int size, void* posInfo, char hasFix)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::parsePosInfoFromNmeaData(const_cast<const char*>(data), size, static_cast<QGeoPositionInfo*>(posInfo), NULL);
}

void QNmeaPositionInfoSource_RequestUpdate(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "requestUpdate", Q_ARG(int, msec));
}

void QNmeaPositionInfoSource_RequestUpdateDefault(void* ptr, int msec)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::requestUpdate(msec);
}

void QNmeaPositionInfoSource_SetDevice(void* ptr, void* device)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QNmeaPositionInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setUpdateInterval(msec);
}

void QNmeaPositionInfoSource_SetUserEquivalentRangeError(void* ptr, double uere)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->setUserEquivalentRangeError(uere);
}

void QNmeaPositionInfoSource_StartUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "startUpdates");
}

void QNmeaPositionInfoSource_StartUpdatesDefault(void* ptr)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::startUpdates();
}

void QNmeaPositionInfoSource_StopUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "stopUpdates");
}

void QNmeaPositionInfoSource_StopUpdatesDefault(void* ptr)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::stopUpdates();
}

long long QNmeaPositionInfoSource_SupportedPositioningMethods(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

long long QNmeaPositionInfoSource_SupportedPositioningMethodsDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::supportedPositioningMethods();
}

long long QNmeaPositionInfoSource_UpdateMode(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->updateMode();
}

double QNmeaPositionInfoSource_UserEquivalentRangeError(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->userEquivalentRangeError();
}

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(void* ptr)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->~QNmeaPositionInfoSource();
}

void QNmeaPositionInfoSource_SetPreferredPositioningMethods(void* ptr, long long methods)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(void* ptr, long long methods)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QNmeaPositionInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNmeaPositionInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void QNmeaPositionInfoSource_ChildEvent(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QNmeaPositionInfoSource_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QNmeaPositionInfoSource_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNmeaPositionInfoSource_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNmeaPositionInfoSource_CustomEvent(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QNmeaPositionInfoSource_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::customEvent(static_cast<QEvent*>(event));
}

void QNmeaPositionInfoSource_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QNmeaPositionInfoSource*>(ptr), "deleteLater");
}

void QNmeaPositionInfoSource_DeleteLaterDefault(void* ptr)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::deleteLater();
}

void QNmeaPositionInfoSource_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QNmeaPositionInfoSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QNmeaPositionInfoSource_Event(void* ptr, void* e)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->event(static_cast<QEvent*>(e));
}

char QNmeaPositionInfoSource_EventDefault(void* ptr, void* e)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::event(static_cast<QEvent*>(e));
}

char QNmeaPositionInfoSource_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QNmeaPositionInfoSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QNmeaPositionInfoSource_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNmeaPositionInfoSource*>(ptr)->metaObject());
}

void* QNmeaPositionInfoSource_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::metaObject());
}

