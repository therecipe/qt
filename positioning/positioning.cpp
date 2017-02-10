// +build !minimal

#define protected public
#define private public

#include "positioning.h"
#include "_cgo_export.h"

#include <QByteArray>
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
#include <QList>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNmeaPositionInfoSource>
#include <QObject>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>

void* QGeoAddress_NewQGeoAddress()
{
	return new QGeoAddress();
}

void* QGeoAddress_NewQGeoAddress2(void* other)
{
	return new QGeoAddress(*static_cast<QGeoAddress*>(other));
}

void QGeoAddress_Clear(void* ptr)
{
	static_cast<QGeoAddress*>(ptr)->clear();
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

void QGeoAddress_DestroyQGeoAddress(void* ptr)
{
	static_cast<QGeoAddress*>(ptr)->~QGeoAddress();
}

struct QtPositioning_PackedString QGeoAddress_City(void* ptr)
{
	return ({ QByteArray t1ae2a5 = static_cast<QGeoAddress*>(ptr)->city().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t1ae2a5.prepend("WHITESPACE").constData()+10), t1ae2a5.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_Country(void* ptr)
{
	return ({ QByteArray te96f62 = static_cast<QGeoAddress*>(ptr)->country().toUtf8(); QtPositioning_PackedString { const_cast<char*>(te96f62.prepend("WHITESPACE").constData()+10), te96f62.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_CountryCode(void* ptr)
{
	return ({ QByteArray tbdca51 = static_cast<QGeoAddress*>(ptr)->countryCode().toUtf8(); QtPositioning_PackedString { const_cast<char*>(tbdca51.prepend("WHITESPACE").constData()+10), tbdca51.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_County(void* ptr)
{
	return ({ QByteArray t5bb60c = static_cast<QGeoAddress*>(ptr)->county().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t5bb60c.prepend("WHITESPACE").constData()+10), t5bb60c.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_District(void* ptr)
{
	return ({ QByteArray t09161b = static_cast<QGeoAddress*>(ptr)->district().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t09161b.prepend("WHITESPACE").constData()+10), t09161b.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_PostalCode(void* ptr)
{
	return ({ QByteArray t3693ff = static_cast<QGeoAddress*>(ptr)->postalCode().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t3693ff.prepend("WHITESPACE").constData()+10), t3693ff.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_State(void* ptr)
{
	return ({ QByteArray t2cd605 = static_cast<QGeoAddress*>(ptr)->state().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t2cd605.prepend("WHITESPACE").constData()+10), t2cd605.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_Street(void* ptr)
{
	return ({ QByteArray tdff817 = static_cast<QGeoAddress*>(ptr)->street().toUtf8(); QtPositioning_PackedString { const_cast<char*>(tdff817.prepend("WHITESPACE").constData()+10), tdff817.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAddress_Text(void* ptr)
{
	return ({ QByteArray tba5c7b = static_cast<QGeoAddress*>(ptr)->text().toUtf8(); QtPositioning_PackedString { const_cast<char*>(tba5c7b.prepend("WHITESPACE").constData()+10), tba5c7b.size()-10 }; });
}

char QGeoAddress_IsEmpty(void* ptr)
{
	return static_cast<QGeoAddress*>(ptr)->isEmpty();
}

char QGeoAddress_IsTextGenerated(void* ptr)
{
	return static_cast<QGeoAddress*>(ptr)->isTextGenerated();
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo2(void* other)
{
	return new QGeoAreaMonitorInfo(*static_cast<QGeoAreaMonitorInfo*>(other));
}

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(char* name)
{
	return new QGeoAreaMonitorInfo(QString(name));
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

void QGeoAreaMonitorInfo_SetNotificationParameters(void* ptr, void* parameters)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setNotificationParameters(*static_cast<QMap<QString, QVariant>*>(parameters));
}

void QGeoAreaMonitorInfo_SetPersistent(void* ptr, char isPersistent)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setPersistent(isPersistent != 0);
}

void QGeoAreaMonitorInfo_DestroyQGeoAreaMonitorInfo(void* ptr)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->~QGeoAreaMonitorInfo();
}

void* QGeoAreaMonitorInfo_Expiration(void* ptr)
{
	return new QDateTime(static_cast<QGeoAreaMonitorInfo*>(ptr)->expiration());
}

void* QGeoAreaMonitorInfo_Area(void* ptr)
{
	return new QGeoShape(static_cast<QGeoAreaMonitorInfo*>(ptr)->area());
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_Identifier(void* ptr)
{
	return ({ QByteArray tba0e32 = static_cast<QGeoAreaMonitorInfo*>(ptr)->identifier().toUtf8(); QtPositioning_PackedString { const_cast<char*>(tba0e32.prepend("WHITESPACE").constData()+10), tba0e32.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_Name(void* ptr)
{
	return ({ QByteArray ta2df77 = static_cast<QGeoAreaMonitorInfo*>(ptr)->name().toUtf8(); QtPositioning_PackedString { const_cast<char*>(ta2df77.prepend("WHITESPACE").constData()+10), ta2df77.size()-10 }; });
}

struct QtPositioning_PackedList QGeoAreaMonitorInfo_NotificationParameters(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue = new QMap<QString, QVariant>(static_cast<QGeoAreaMonitorInfo*>(ptr)->notificationParameters()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

char QGeoAreaMonitorInfo_IsPersistent(void* ptr)
{
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isPersistent();
}

char QGeoAreaMonitorInfo_IsValid(void* ptr)
{
	return static_cast<QGeoAreaMonitorInfo*>(ptr)->isValid();
}

void* QGeoAreaMonitorInfo___setNotificationParameters_parameters_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QGeoAreaMonitorInfo___setNotificationParameters_parameters_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QGeoAreaMonitorInfo___setNotificationParameters_parameters_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtPositioning_PackedList QGeoAreaMonitorInfo___setNotificationParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoAreaMonitorInfo___notificationParameters_atList(void* ptr, char* i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString(i)));
}

void QGeoAreaMonitorInfo___notificationParameters_setList(void* ptr, char* key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString(key), *static_cast<QVariant*>(i));
}

void* QGeoAreaMonitorInfo___notificationParameters_newList(void* ptr)
{
	return new QMap<QString, QVariant>;
}

struct QtPositioning_PackedList QGeoAreaMonitorInfo___notificationParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_____setNotificationParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoAreaMonitorInfo_____setNotificationParameters_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QGeoAreaMonitorInfo_____setNotificationParameters_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_____notificationParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoAreaMonitorInfo_____notificationParameters_keyList_setList(void* ptr, char* i)
{
	static_cast<QList<QString>*>(ptr)->append(QString(i));
}

void* QGeoAreaMonitorInfo_____notificationParameters_keyList_newList(void* ptr)
{
	return new QList<QString>;
}

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource
{
public:
	MyQGeoAreaMonitorSource(QObject *parent) : QGeoAreaMonitorSource(parent) {};
	bool requestUpdate(const QGeoAreaMonitorInfo & monitor, const char * sign) { QtPositioning_PackedString signPacked = { const_cast<char*>(sign), -1 };return callbackQGeoAreaMonitorSource_RequestUpdate(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), signPacked) != 0; };
	bool startMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StartMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	bool stopMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StopMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	void Signal_AreaEntered(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaEntered(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_AreaExited(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaExited(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_Error2(QGeoAreaMonitorSource::Error areaMonitoringError) { callbackQGeoAreaMonitorSource_Error2(this, areaMonitoringError); };
	void Signal_MonitorExpired(const QGeoAreaMonitorInfo & monitor) { callbackQGeoAreaMonitorSource_MonitorExpired(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)); };
	void setPositionInfoSource(QGeoPositionInfoSource * newSource) { callbackQGeoAreaMonitorSource_SetPositionInfoSource(this, newSource); };
	 ~MyQGeoAreaMonitorSource() { callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(this); };
	AreaMonitorFeatures supportedAreaMonitorFeatures() const { return static_cast<QGeoAreaMonitorSource::AreaMonitorFeature>(callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	Error error() const { return static_cast<QGeoAreaMonitorSource::Error>(callbackQGeoAreaMonitorSource_Error(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	QGeoPositionInfoSource * positionInfoSource() const { return static_cast<QGeoPositionInfoSource*>(callbackQGeoAreaMonitorSource_PositionInfoSource(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	QList<QGeoAreaMonitorInfo> activeMonitors() const { return *static_cast<QList<QGeoAreaMonitorInfo>*>(callbackQGeoAreaMonitorSource_ActiveMonitors(const_cast<MyQGeoAreaMonitorSource*>(this))); };
	QList<QGeoAreaMonitorInfo> activeMonitors(const QGeoShape & lookupArea) const { return *static_cast<QList<QGeoAreaMonitorInfo>*>(callbackQGeoAreaMonitorSource_ActiveMonitors2(const_cast<MyQGeoAreaMonitorSource*>(this), const_cast<QGeoShape*>(&lookupArea))); };
	bool event(QEvent * e) { return callbackQGeoAreaMonitorSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoAreaMonitorSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoAreaMonitorSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoAreaMonitorSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoAreaMonitorSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void timerEvent(QTimerEvent * event) { callbackQGeoAreaMonitorSource_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoAreaMonitorSource_MetaObject(const_cast<MyQGeoAreaMonitorSource*>(this))); };
};

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent)
{
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoAreaMonitorSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_NewQGeoAreaMonitorSource(void* parent)
{
	return new MyQGeoAreaMonitorSource(static_cast<QObject*>(parent));
}

struct QtPositioning_PackedString QGeoAreaMonitorSource_QGeoAreaMonitorSource_AvailableSources()
{
	return ({ QByteArray t3a05b8 = QGeoAreaMonitorSource::availableSources().join("|").toUtf8(); QtPositioning_PackedString { const_cast<char*>(t3a05b8.prepend("WHITESPACE").constData()+10), t3a05b8.size()-10 }; });
}

char QGeoAreaMonitorSource_RequestUpdate(void* ptr, void* monitor, char* sign)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->requestUpdate(*static_cast<QGeoAreaMonitorInfo*>(monitor), const_cast<const char*>(sign));
}

char QGeoAreaMonitorSource_StartMonitoring(void* ptr, void* monitor)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->startMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

char QGeoAreaMonitorSource_StopMonitoring(void* ptr, void* monitor)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->stopMonitoring(*static_cast<QGeoAreaMonitorInfo*>(monitor));
}

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

void QGeoAreaMonitorSource_SetPositionInfoSource(void* ptr, void* newSource)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

void QGeoAreaMonitorSource_SetPositionInfoSourceDefault(void* ptr, void* newSource)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::setPositionInfoSource(static_cast<QGeoPositionInfoSource*>(newSource));
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(void* ptr)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->~QGeoAreaMonitorSource();
}

void QGeoAreaMonitorSource_DestroyQGeoAreaMonitorSourceDefault(void* ptr)
{

}

long long QGeoAreaMonitorSource_SupportedAreaMonitorFeatures(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->supportedAreaMonitorFeatures();
}

long long QGeoAreaMonitorSource_Error(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->error();
}

void* QGeoAreaMonitorSource_PositionInfoSource(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->positionInfoSource();
}

void* QGeoAreaMonitorSource_PositionInfoSourceDefault(void* ptr)
{
	return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::positionInfoSource();
}

struct QtPositioning_PackedList QGeoAreaMonitorSource_ActiveMonitors(void* ptr)
{
	return ({ QList<QGeoAreaMonitorInfo>* tmpValue = new QList<QGeoAreaMonitorInfo>(static_cast<QGeoAreaMonitorSource*>(ptr)->activeMonitors()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedList QGeoAreaMonitorSource_ActiveMonitors2(void* ptr, void* lookupArea)
{
	return ({ QList<QGeoAreaMonitorInfo>* tmpValue = new QList<QGeoAreaMonitorInfo>(static_cast<QGeoAreaMonitorSource*>(ptr)->activeMonitors(*static_cast<QGeoShape*>(lookupArea))); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedString QGeoAreaMonitorSource_SourceName(void* ptr)
{
	return ({ QByteArray t77e733 = static_cast<QGeoAreaMonitorSource*>(ptr)->sourceName().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t77e733.prepend("WHITESPACE").constData()+10), t77e733.size()-10 }; });
}

void* QGeoAreaMonitorSource___activeMonitors_atList(void* ptr, int i)
{
	return new QGeoAreaMonitorInfo(static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___activeMonitors_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->append(*static_cast<QGeoAreaMonitorInfo*>(i));
}

void* QGeoAreaMonitorSource___activeMonitors_newList(void* ptr)
{
	return new QList<QGeoAreaMonitorInfo>;
}

void* QGeoAreaMonitorSource___activeMonitors_atList2(void* ptr, int i)
{
	return new QGeoAreaMonitorInfo(static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___activeMonitors_setList2(void* ptr, void* i)
{
	static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->append(*static_cast<QGeoAreaMonitorInfo*>(i));
}

void* QGeoAreaMonitorSource___activeMonitors_newList2(void* ptr)
{
	return new QList<QGeoAreaMonitorInfo>;
}

void* QGeoAreaMonitorSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoAreaMonitorSource___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QGeoAreaMonitorSource___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoAreaMonitorSource___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoAreaMonitorSource___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoAreaMonitorSource___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoAreaMonitorSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___children_newList(void* ptr)
{
	return new QList<QObject *>;
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

void QGeoAreaMonitorSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoAreaMonitorSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::timerEvent(static_cast<QTimerEvent*>(event));
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

void QGeoCircle_SetCenter(void* ptr, void* center)
{
	static_cast<QGeoCircle*>(ptr)->setCenter(*static_cast<QGeoCoordinate*>(center));
}

void QGeoCircle_SetRadius(void* ptr, double radius)
{
	static_cast<QGeoCircle*>(ptr)->setRadius(radius);
}

void QGeoCircle_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoCircle*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void QGeoCircle_DestroyQGeoCircle(void* ptr)
{
	static_cast<QGeoCircle*>(ptr)->~QGeoCircle();
}

void* QGeoCircle_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoCircle(static_cast<QGeoCircle*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

void* QGeoCircle_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoCircle*>(ptr)->center());
}

struct QtPositioning_PackedString QGeoCircle_ToString(void* ptr)
{
	return ({ QByteArray t8fcc7c = static_cast<QGeoCircle*>(ptr)->toString().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t8fcc7c.prepend("WHITESPACE").constData()+10), t8fcc7c.size()-10 }; });
}

double QGeoCircle_Radius(void* ptr)
{
	return static_cast<QGeoCircle*>(ptr)->radius();
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

void QGeoCoordinate_DestroyQGeoCoordinate(void* ptr)
{
	static_cast<QGeoCoordinate*>(ptr)->~QGeoCoordinate();
}

long long QGeoCoordinate_Type(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->type();
}

void* QGeoCoordinate_AtDistanceAndAzimuth(void* ptr, double distance, double azimuth, double distanceUp)
{
	return new QGeoCoordinate(static_cast<QGeoCoordinate*>(ptr)->atDistanceAndAzimuth(distance, azimuth, distanceUp));
}

struct QtPositioning_PackedString QGeoCoordinate_ToString(void* ptr, long long format)
{
	return ({ QByteArray t4ceea1 = static_cast<QGeoCoordinate*>(ptr)->toString(static_cast<QGeoCoordinate::CoordinateFormat>(format)).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t4ceea1.prepend("WHITESPACE").constData()+10), t4ceea1.size()-10 }; });
}

char QGeoCoordinate_IsValid(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->isValid();
}

double QGeoCoordinate_Altitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->altitude();
}

double QGeoCoordinate_Latitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->latitude();
}

double QGeoCoordinate_Longitude(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->longitude();
}

double QGeoCoordinate_AzimuthTo(void* ptr, void* other)
{
	return static_cast<QGeoCoordinate*>(ptr)->azimuthTo(*static_cast<QGeoCoordinate*>(other));
}

double QGeoCoordinate_DistanceTo(void* ptr, void* other)
{
	return static_cast<QGeoCoordinate*>(ptr)->distanceTo(*static_cast<QGeoCoordinate*>(other));
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

void QGeoPositionInfo_DestroyQGeoPositionInfo(void* ptr)
{
	static_cast<QGeoPositionInfo*>(ptr)->~QGeoPositionInfo();
}

void* QGeoPositionInfo_Timestamp(void* ptr)
{
	return new QDateTime(static_cast<QGeoPositionInfo*>(ptr)->timestamp());
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

double QGeoPositionInfo_Attribute(void* ptr, long long attribute)
{
	return static_cast<QGeoPositionInfo*>(ptr)->attribute(static_cast<QGeoPositionInfo::Attribute>(attribute));
}

class MyQGeoPositionInfoSource: public QGeoPositionInfoSource
{
public:
	MyQGeoPositionInfoSource(QObject *parent) : QGeoPositionInfoSource(parent) {};
	void Signal_PositionUpdated(const QGeoPositionInfo & update) { callbackQGeoPositionInfoSource_PositionUpdated(this, const_cast<QGeoPositionInfo*>(&update)); };
	void setUpdateInterval(int msec) { callbackQGeoPositionInfoSource_SetUpdateInterval(this, msec); };
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSource_Error2(this, positioningError); };
	void requestUpdate(int timeout) { callbackQGeoPositionInfoSource_RequestUpdate(this, timeout); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	void startUpdates() { callbackQGeoPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoPositionInfoSource_StopUpdates(this); };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSource_UpdateTimeout(this); };
	 ~MyQGeoPositionInfoSource() { callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource(this); };
	Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQGeoPositionInfoSource_Error(const_cast<MyQGeoPositionInfoSource*>(this))); };
	PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQGeoPositionInfoSource_SupportedPositioningMethods(const_cast<MyQGeoPositionInfoSource*>(this))); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQGeoPositionInfoSource_LastKnownPosition(const_cast<MyQGeoPositionInfoSource*>(this), fromSatellitePositioningMethodsOnly)); };
	int minimumUpdateInterval() const { return callbackQGeoPositionInfoSource_MinimumUpdateInterval(const_cast<MyQGeoPositionInfoSource*>(this)); };
	bool event(QEvent * e) { return callbackQGeoPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoPositionInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoPositionInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void timerEvent(QTimerEvent * event) { callbackQGeoPositionInfoSource_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoPositionInfoSource_MetaObject(const_cast<MyQGeoPositionInfoSource*>(this))); };
};

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

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setUpdateInterval(msec);
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoPositionInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_NewQGeoPositionInfoSource(void* parent)
{
	return new MyQGeoPositionInfoSource(static_cast<QObject*>(parent));
}

struct QtPositioning_PackedString QGeoPositionInfoSource_QGeoPositionInfoSource_AvailableSources()
{
	return ({ QByteArray t547e22 = QGeoPositionInfoSource::availableSources().join("|").toUtf8(); QtPositioning_PackedString { const_cast<char*>(t547e22.prepend("WHITESPACE").constData()+10), t547e22.size()-10 }; });
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

void QGeoPositionInfoSource_DestroyQGeoPositionInfoSourceDefault(void* ptr)
{

}

long long QGeoPositionInfoSource_Error(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->error();
}

long long QGeoPositionInfoSource_PreferredPositioningMethods(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->preferredPositioningMethods();
}

long long QGeoPositionInfoSource_SupportedPositioningMethods(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

void* QGeoPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QGeoPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

struct QtPositioning_PackedString QGeoPositionInfoSource_SourceName(void* ptr)
{
	return ({ QByteArray ta59e39 = static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8(); QtPositioning_PackedString { const_cast<char*>(ta59e39.prepend("WHITESPACE").constData()+10), ta59e39.size()-10 }; });
}

int QGeoPositionInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QGeoPositionInfoSource_UpdateInterval(void* ptr)
{
	return static_cast<QGeoPositionInfoSource*>(ptr)->updateInterval();
}

void* QGeoPositionInfoSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoPositionInfoSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoPositionInfoSource___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QGeoPositionInfoSource___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoPositionInfoSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoPositionInfoSource___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoPositionInfoSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoPositionInfoSource___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoPositionInfoSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoPositionInfoSource___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoPositionInfoSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___children_newList(void* ptr)
{
	return new QList<QObject *>;
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

void QGeoPositionInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoPositionInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
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
	 ~MyQGeoPositionInfoSourceFactory() { callbackQGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactory(this); };
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

void QGeoPositionInfoSourceFactory_DestroyQGeoPositionInfoSourceFactoryDefault(void* ptr)
{

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

void* QGeoRectangle_NewQGeoRectangle4(void* coordinates)
{
	return new QGeoRectangle(*static_cast<QList<QGeoCoordinate>*>(coordinates));
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

void QGeoRectangle_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoRectangle*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void QGeoRectangle_DestroyQGeoRectangle(void* ptr)
{
	static_cast<QGeoRectangle*>(ptr)->~QGeoRectangle();
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

void* QGeoRectangle_TopLeft(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->topLeft());
}

void* QGeoRectangle_TopRight(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoRectangle*>(ptr)->topRight());
}

void* QGeoRectangle_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoRectangle(static_cast<QGeoRectangle*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

void* QGeoRectangle_United(void* ptr, void* rectangle)
{
	return new QGeoRectangle(static_cast<QGeoRectangle*>(ptr)->united(*static_cast<QGeoRectangle*>(rectangle)));
}

struct QtPositioning_PackedString QGeoRectangle_ToString(void* ptr)
{
	return ({ QByteArray tf2b3c6 = static_cast<QGeoRectangle*>(ptr)->toString().toUtf8(); QtPositioning_PackedString { const_cast<char*>(tf2b3c6.prepend("WHITESPACE").constData()+10), tf2b3c6.size()-10 }; });
}

char QGeoRectangle_Contains(void* ptr, void* rectangle)
{
	return static_cast<QGeoRectangle*>(ptr)->contains(*static_cast<QGeoRectangle*>(rectangle));
}

char QGeoRectangle_Intersects(void* ptr, void* rectangle)
{
	return static_cast<QGeoRectangle*>(ptr)->intersects(*static_cast<QGeoRectangle*>(rectangle));
}

double QGeoRectangle_Height(void* ptr)
{
	return static_cast<QGeoRectangle*>(ptr)->height();
}

double QGeoRectangle_Width(void* ptr)
{
	return static_cast<QGeoRectangle*>(ptr)->width();
}

void* QGeoRectangle___QGeoRectangle_coordinates_atList4(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRectangle___QGeoRectangle_coordinates_setList4(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRectangle___QGeoRectangle_coordinates_newList4(void* ptr)
{
	return new QList<QGeoCoordinate>;
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo()
{
	return new QGeoSatelliteInfo();
}

void* QGeoSatelliteInfo_NewQGeoSatelliteInfo2(void* other)
{
	return new QGeoSatelliteInfo(*static_cast<QGeoSatelliteInfo*>(other));
}

void QGeoSatelliteInfo_RemoveAttribute(void* ptr, long long attribute)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->removeAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
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

void QGeoSatelliteInfo_DestroyQGeoSatelliteInfo(void* ptr)
{
	static_cast<QGeoSatelliteInfo*>(ptr)->~QGeoSatelliteInfo();
}

long long QGeoSatelliteInfo_SatelliteSystem(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteSystem();
}

char QGeoSatelliteInfo_HasAttribute(void* ptr, long long attribute)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->hasAttribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

int QGeoSatelliteInfo_SatelliteIdentifier(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->satelliteIdentifier();
}

int QGeoSatelliteInfo_SignalStrength(void* ptr)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->signalStrength();
}

double QGeoSatelliteInfo_Attribute(void* ptr, long long attribute)
{
	return static_cast<QGeoSatelliteInfo*>(ptr)->attribute(static_cast<QGeoSatelliteInfo::Attribute>(attribute));
}

class MyQGeoSatelliteInfoSource: public QGeoSatelliteInfoSource
{
public:
	MyQGeoSatelliteInfoSource(QObject *parent) : QGeoSatelliteInfoSource(parent) {};
	void Signal_Error2(QGeoSatelliteInfoSource::Error satelliteError) { callbackQGeoSatelliteInfoSource_Error2(this, satelliteError); };
	void Signal_RequestTimeout() { callbackQGeoSatelliteInfoSource_RequestTimeout(this); };
	void requestUpdate(int timeout) { callbackQGeoSatelliteInfoSource_RequestUpdate(this, timeout); };
	void Signal_SatellitesInUseUpdated(const QList<QGeoSatelliteInfo> & satellites) { callbackQGeoSatelliteInfoSource_SatellitesInUseUpdated(this, ({ QList<QGeoSatelliteInfo>* tmpValue = const_cast<QList<QGeoSatelliteInfo>*>(&satellites); QtPositioning_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SatellitesInViewUpdated(const QList<QGeoSatelliteInfo> & satellites) { callbackQGeoSatelliteInfoSource_SatellitesInViewUpdated(this, ({ QList<QGeoSatelliteInfo>* tmpValue = const_cast<QList<QGeoSatelliteInfo>*>(&satellites); QtPositioning_PackedList { tmpValue, tmpValue->size() }; })); };
	void setUpdateInterval(int msec) { callbackQGeoSatelliteInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQGeoSatelliteInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoSatelliteInfoSource_StopUpdates(this); };
	 ~MyQGeoSatelliteInfoSource() { callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(this); };
	Error error() const { return static_cast<QGeoSatelliteInfoSource::Error>(callbackQGeoSatelliteInfoSource_Error(const_cast<MyQGeoSatelliteInfoSource*>(this))); };
	int minimumUpdateInterval() const { return callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(const_cast<MyQGeoSatelliteInfoSource*>(this)); };
	bool event(QEvent * e) { return callbackQGeoSatelliteInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoSatelliteInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoSatelliteInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoSatelliteInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoSatelliteInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void timerEvent(QTimerEvent * event) { callbackQGeoSatelliteInfoSource_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoSatelliteInfoSource_MetaObject(const_cast<MyQGeoSatelliteInfoSource*>(this))); };
};

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(char* sourceName, void* parent)
{
	return QGeoSatelliteInfoSource::createSource(QString(sourceName), static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(void* parent)
{
	return new MyQGeoSatelliteInfoSource(static_cast<QObject*>(parent));
}

struct QtPositioning_PackedString QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_AvailableSources()
{
	return ({ QByteArray tab00fe = QGeoSatelliteInfoSource::availableSources().join("|").toUtf8(); QtPositioning_PackedString { const_cast<char*>(tab00fe.prepend("WHITESPACE").constData()+10), tab00fe.size()-10 }; });
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

void QGeoSatelliteInfoSource_ConnectSatellitesInUseUpdated(void* ptr)
{
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&QGeoSatelliteInfoSource::satellitesInUseUpdated), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&MyQGeoSatelliteInfoSource::Signal_SatellitesInUseUpdated));
}

void QGeoSatelliteInfoSource_DisconnectSatellitesInUseUpdated(void* ptr)
{
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&QGeoSatelliteInfoSource::satellitesInUseUpdated), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&MyQGeoSatelliteInfoSource::Signal_SatellitesInUseUpdated));
}

void QGeoSatelliteInfoSource_SatellitesInUseUpdated(void* ptr, void* satellites)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->satellitesInUseUpdated(*static_cast<QList<QGeoSatelliteInfo>*>(satellites));
}

void QGeoSatelliteInfoSource_ConnectSatellitesInViewUpdated(void* ptr)
{
	QObject::connect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&QGeoSatelliteInfoSource::satellitesInViewUpdated), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&MyQGeoSatelliteInfoSource::Signal_SatellitesInViewUpdated));
}

void QGeoSatelliteInfoSource_DisconnectSatellitesInViewUpdated(void* ptr)
{
	QObject::disconnect(static_cast<QGeoSatelliteInfoSource*>(ptr), static_cast<void (QGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&QGeoSatelliteInfoSource::satellitesInViewUpdated), static_cast<MyQGeoSatelliteInfoSource*>(ptr), static_cast<void (MyQGeoSatelliteInfoSource::*)(const QList<QGeoSatelliteInfo> &)>(&MyQGeoSatelliteInfoSource::Signal_SatellitesInViewUpdated));
}

void QGeoSatelliteInfoSource_SatellitesInViewUpdated(void* ptr, void* satellites)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->satellitesInViewUpdated(*static_cast<QList<QGeoSatelliteInfo>*>(satellites));
}

void QGeoSatelliteInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoSatelliteInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::setUpdateInterval(msec);
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

void QGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSourceDefault(void* ptr)
{

}

long long QGeoSatelliteInfoSource_Error(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->error();
}

struct QtPositioning_PackedString QGeoSatelliteInfoSource_SourceName(void* ptr)
{
	return ({ QByteArray t6fe95a = static_cast<QGeoSatelliteInfoSource*>(ptr)->sourceName().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t6fe95a.prepend("WHITESPACE").constData()+10), t6fe95a.size()-10 }; });
}

int QGeoSatelliteInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->minimumUpdateInterval();
}

int QGeoSatelliteInfoSource_UpdateInterval(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->updateInterval();
}

void* QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_atList(void* ptr, int i)
{
	return new QGeoSatelliteInfo(static_cast<QList<QGeoSatelliteInfo>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoSatelliteInfo>*>(ptr)->append(*static_cast<QGeoSatelliteInfo*>(i));
}

void* QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_newList(void* ptr)
{
	return new QList<QGeoSatelliteInfo>;
}

void* QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_atList(void* ptr, int i)
{
	return new QGeoSatelliteInfo(static_cast<QList<QGeoSatelliteInfo>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoSatelliteInfo>*>(ptr)->append(*static_cast<QGeoSatelliteInfo*>(i));
}

void* QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_newList(void* ptr)
{
	return new QList<QGeoSatelliteInfo>;
}

void* QGeoSatelliteInfoSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoSatelliteInfoSource___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QGeoSatelliteInfoSource___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoSatelliteInfoSource___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoSatelliteInfoSource___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QGeoSatelliteInfoSource___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoSatelliteInfoSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___children_newList(void* ptr)
{
	return new QList<QObject *>;
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

void QGeoSatelliteInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGeoSatelliteInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
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

void QGeoShape_ExtendShape(void* ptr, void* coordinate)
{
	static_cast<QGeoShape*>(ptr)->extendShape(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoShape_DestroyQGeoShape(void* ptr)
{
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

void* QGeoShape_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoShape*>(ptr)->center());
}

struct QtPositioning_PackedString QGeoShape_ToString(void* ptr)
{
	return ({ QByteArray teed09f = static_cast<QGeoShape*>(ptr)->toString().toUtf8(); QtPositioning_PackedString { const_cast<char*>(teed09f.prepend("WHITESPACE").constData()+10), teed09f.size()-10 }; });
}

long long QGeoShape_Type(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->type();
}

char QGeoShape_Contains(void* ptr, void* coordinate)
{
	return static_cast<QGeoShape*>(ptr)->contains(*static_cast<QGeoCoordinate*>(coordinate));
}

char QGeoShape_IsEmpty(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->isEmpty();
}

char QGeoShape_IsValid(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->isValid();
}

class MyQNmeaPositionInfoSource: public QNmeaPositionInfoSource
{
public:
	MyQNmeaPositionInfoSource(UpdateMode updateMode, QObject *parent) : QNmeaPositionInfoSource(updateMode, parent) {};
	bool parsePosInfoFromNmeaData(const char * data, int size, QGeoPositionInfo * posInfo, bool * hasFix) { QtPositioning_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(this, dataPacked, size, posInfo, *hasFix) != 0; };
	void requestUpdate(int msec) { callbackQNmeaPositionInfoSource_RequestUpdate(this, msec); };
	void setUpdateInterval(int msec) { callbackQNmeaPositionInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQNmeaPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQNmeaPositionInfoSource_StopUpdates(this); };
	Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQNmeaPositionInfoSource_Error(const_cast<MyQNmeaPositionInfoSource*>(this))); };
	PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQNmeaPositionInfoSource_SupportedPositioningMethods(const_cast<MyQNmeaPositionInfoSource*>(this))); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQNmeaPositionInfoSource_LastKnownPosition(const_cast<MyQNmeaPositionInfoSource*>(this), fromSatellitePositioningMethodsOnly)); };
	int minimumUpdateInterval() const { return callbackQNmeaPositionInfoSource_MinimumUpdateInterval(const_cast<MyQNmeaPositionInfoSource*>(this)); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQNmeaPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	bool event(QEvent * e) { return callbackQNmeaPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQNmeaPositionInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQNmeaPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQNmeaPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQNmeaPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQNmeaPositionInfoSource_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQNmeaPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void timerEvent(QTimerEvent * event) { callbackQNmeaPositionInfoSource_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQNmeaPositionInfoSource_MetaObject(const_cast<MyQNmeaPositionInfoSource*>(this))); };
};

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(long long updateMode, void* parent)
{
	return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
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

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(void* ptr)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->~QNmeaPositionInfoSource();
}

long long QNmeaPositionInfoSource_Error(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->error();
}

long long QNmeaPositionInfoSource_ErrorDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::error();
}

long long QNmeaPositionInfoSource_SupportedPositioningMethods(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->supportedPositioningMethods();
}

long long QNmeaPositionInfoSource_SupportedPositioningMethodsDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::supportedPositioningMethods();
}

void* QNmeaPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

void* QNmeaPositionInfoSource_LastKnownPositionDefault(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

void* QNmeaPositionInfoSource_Device(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->device();
}

long long QNmeaPositionInfoSource_UpdateMode(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->updateMode();
}

double QNmeaPositionInfoSource_UserEquivalentRangeError(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->userEquivalentRangeError();
}

int QNmeaPositionInfoSource_MinimumUpdateInterval(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->minimumUpdateInterval();
}

int QNmeaPositionInfoSource_MinimumUpdateIntervalDefault(void* ptr)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::minimumUpdateInterval();
}

void* QNmeaPositionInfoSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QNmeaPositionInfoSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QNmeaPositionInfoSource___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QNmeaPositionInfoSource___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QNmeaPositionInfoSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNmeaPositionInfoSource___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QNmeaPositionInfoSource___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QNmeaPositionInfoSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNmeaPositionInfoSource___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QNmeaPositionInfoSource___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QNmeaPositionInfoSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNmeaPositionInfoSource___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QNmeaPositionInfoSource___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QNmeaPositionInfoSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QNmeaPositionInfoSource___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

void QNmeaPositionInfoSource_SetPreferredPositioningMethods(void* ptr, long long methods)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
}

void QNmeaPositionInfoSource_SetPreferredPositioningMethodsDefault(void* ptr, long long methods)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
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

void QNmeaPositionInfoSource_TimerEvent(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QNmeaPositionInfoSource_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QNmeaPositionInfoSource_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNmeaPositionInfoSource*>(ptr)->metaObject());
}

void* QNmeaPositionInfoSource_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::metaObject());
}

