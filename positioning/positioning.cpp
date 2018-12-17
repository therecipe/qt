// +build !minimal

#define protected public
#define private public

#include "positioning.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QDateTime>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGeoAddress>
#include <QGeoAreaMonitorInfo>
#include <QGeoAreaMonitorSource>
#include <QGeoCircle>
#include <QGeoCoordinate>
#include <QGeoPath>
#include <QGeoPolygon>
#include <QGeoPositionInfo>
#include <QGeoPositionInfoSource>
#include <QGeoPositionInfoSourceFactory>
#include <QGeoRectangle>
#include <QGeoSatelliteInfo>
#include <QGeoSatelliteInfoSource>
#include <QGeoShape>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QIODevice>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QNmeaPositionInfoSource>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

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

void QGeoAddress_SetCity(void* ptr, struct QtPositioning_PackedString city)
{
	static_cast<QGeoAddress*>(ptr)->setCity(QString::fromUtf8(city.data, city.len));
}

void QGeoAddress_SetCountry(void* ptr, struct QtPositioning_PackedString country)
{
	static_cast<QGeoAddress*>(ptr)->setCountry(QString::fromUtf8(country.data, country.len));
}

void QGeoAddress_SetCountryCode(void* ptr, struct QtPositioning_PackedString countryCode)
{
	static_cast<QGeoAddress*>(ptr)->setCountryCode(QString::fromUtf8(countryCode.data, countryCode.len));
}

void QGeoAddress_SetCounty(void* ptr, struct QtPositioning_PackedString county)
{
	static_cast<QGeoAddress*>(ptr)->setCounty(QString::fromUtf8(county.data, county.len));
}

void QGeoAddress_SetDistrict(void* ptr, struct QtPositioning_PackedString district)
{
	static_cast<QGeoAddress*>(ptr)->setDistrict(QString::fromUtf8(district.data, district.len));
}

void QGeoAddress_SetPostalCode(void* ptr, struct QtPositioning_PackedString postalCode)
{
	static_cast<QGeoAddress*>(ptr)->setPostalCode(QString::fromUtf8(postalCode.data, postalCode.len));
}

void QGeoAddress_SetState(void* ptr, struct QtPositioning_PackedString state)
{
	static_cast<QGeoAddress*>(ptr)->setState(QString::fromUtf8(state.data, state.len));
}

void QGeoAddress_SetStreet(void* ptr, struct QtPositioning_PackedString street)
{
	static_cast<QGeoAddress*>(ptr)->setStreet(QString::fromUtf8(street.data, street.len));
}

void QGeoAddress_SetText(void* ptr, struct QtPositioning_PackedString text)
{
	static_cast<QGeoAddress*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
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

void* QGeoAreaMonitorInfo_NewQGeoAreaMonitorInfo(struct QtPositioning_PackedString name)
{
	return new QGeoAreaMonitorInfo(QString::fromUtf8(name.data, name.len));
}

void QGeoAreaMonitorInfo_SetArea(void* ptr, void* newShape)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setArea(*static_cast<QGeoShape*>(newShape));
}

void QGeoAreaMonitorInfo_SetExpiration(void* ptr, void* expiry)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setExpiration(*static_cast<QDateTime*>(expiry));
}

void QGeoAreaMonitorInfo_SetName(void* ptr, struct QtPositioning_PackedString name)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QGeoAreaMonitorInfo_SetNotificationParameters(void* ptr, void* parameters)
{
	static_cast<QGeoAreaMonitorInfo*>(ptr)->setNotificationParameters(({ QMap<QString, QVariant>* tmpP = static_cast<QMap<QString, QVariant>*>(parameters); QMap<QString, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
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

void* QGeoAreaMonitorInfo___setNotificationParameters_parameters_atList(void* ptr, struct QtPositioning_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoAreaMonitorInfo___setNotificationParameters_parameters_setList(void* ptr, struct QtPositioning_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoAreaMonitorInfo___setNotificationParameters_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtPositioning_PackedList QGeoAreaMonitorInfo___setNotificationParameters_parameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoAreaMonitorInfo___notificationParameters_atList(void* ptr, struct QtPositioning_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QGeoAreaMonitorInfo___notificationParameters_setList(void* ptr, struct QtPositioning_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoAreaMonitorInfo___notificationParameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtPositioning_PackedList QGeoAreaMonitorInfo___notificationParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_setList(void* ptr, struct QtPositioning_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoAreaMonitorInfo_____setNotificationParameters_parameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtPositioning_PackedString QGeoAreaMonitorInfo_____notificationParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QGeoAreaMonitorInfo_____notificationParameters_keyList_setList(void* ptr, struct QtPositioning_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoAreaMonitorInfo_____notificationParameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

class MyQGeoAreaMonitorSource: public QGeoAreaMonitorSource
{
public:
	MyQGeoAreaMonitorSource(QObject *parent) : QGeoAreaMonitorSource(parent) {QGeoAreaMonitorSource_QGeoAreaMonitorSource_QRegisterMetaType();};
	bool requestUpdate(const QGeoAreaMonitorInfo & monitor, const char * sign) { QtPositioning_PackedString signPacked = { const_cast<char*>(sign), -1 };return callbackQGeoAreaMonitorSource_RequestUpdate(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), signPacked) != 0; };
	bool startMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StartMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	bool stopMonitoring(const QGeoAreaMonitorInfo & monitor) { return callbackQGeoAreaMonitorSource_StopMonitoring(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)) != 0; };
	void Signal_AreaEntered(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaEntered(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_AreaExited(const QGeoAreaMonitorInfo & monitor, const QGeoPositionInfo & update) { callbackQGeoAreaMonitorSource_AreaExited(this, const_cast<QGeoAreaMonitorInfo*>(&monitor), const_cast<QGeoPositionInfo*>(&update)); };
	void Signal_Error2(QGeoAreaMonitorSource::Error areaMonitoringError) { callbackQGeoAreaMonitorSource_Error2(this, areaMonitoringError); };
	void Signal_MonitorExpired(const QGeoAreaMonitorInfo & monitor) { callbackQGeoAreaMonitorSource_MonitorExpired(this, const_cast<QGeoAreaMonitorInfo*>(&monitor)); };
	void setPositionInfoSource(QGeoPositionInfoSource * newSource) { callbackQGeoAreaMonitorSource_SetPositionInfoSource(this, newSource); };
	 ~MyQGeoAreaMonitorSource() { callbackQGeoAreaMonitorSource_DestroyQGeoAreaMonitorSource(this); };
	QGeoAreaMonitorSource::AreaMonitorFeatures supportedAreaMonitorFeatures() const { return static_cast<QGeoAreaMonitorSource::AreaMonitorFeature>(callbackQGeoAreaMonitorSource_SupportedAreaMonitorFeatures(const_cast<void*>(static_cast<const void*>(this)))); };
	QGeoAreaMonitorSource::Error error() const { return static_cast<QGeoAreaMonitorSource::Error>(callbackQGeoAreaMonitorSource_Error(const_cast<void*>(static_cast<const void*>(this)))); };
	QGeoPositionInfoSource * positionInfoSource() const { return static_cast<QGeoPositionInfoSource*>(callbackQGeoAreaMonitorSource_PositionInfoSource(const_cast<void*>(static_cast<const void*>(this)))); };
	QList<QGeoAreaMonitorInfo> activeMonitors() const { return ({ QList<QGeoAreaMonitorInfo>* tmpP = static_cast<QList<QGeoAreaMonitorInfo>*>(callbackQGeoAreaMonitorSource_ActiveMonitors(const_cast<void*>(static_cast<const void*>(this)))); QList<QGeoAreaMonitorInfo> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QList<QGeoAreaMonitorInfo> activeMonitors(const QGeoShape & lookupArea) const { return ({ QList<QGeoAreaMonitorInfo>* tmpP = static_cast<QList<QGeoAreaMonitorInfo>*>(callbackQGeoAreaMonitorSource_ActiveMonitors2(const_cast<void*>(static_cast<const void*>(this)), const_cast<QGeoShape*>(&lookupArea))); QList<QGeoAreaMonitorInfo> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoAreaMonitorSource_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	bool event(QEvent * e) { return callbackQGeoAreaMonitorSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoAreaMonitorSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoAreaMonitorSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoAreaMonitorSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoAreaMonitorSource_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoAreaMonitorSource_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoAreaMonitorSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPositioning_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoAreaMonitorSource_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoAreaMonitorSource_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQGeoAreaMonitorSource*)

int QGeoAreaMonitorSource_QGeoAreaMonitorSource_QRegisterMetaType(){qRegisterMetaType<QGeoAreaMonitorSource*>(); return qRegisterMetaType<MyQGeoAreaMonitorSource*>();}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateDefaultSource(void* parent)
{
	return QGeoAreaMonitorSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_QGeoAreaMonitorSource_CreateSource(struct QtPositioning_PackedString sourceName, void* parent)
{
	return QGeoAreaMonitorSource::createSource(QString::fromUtf8(sourceName.data, sourceName.len), static_cast<QObject*>(parent));
}

void* QGeoAreaMonitorSource_NewQGeoAreaMonitorSource(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoAreaMonitorSource(static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoAreaMonitorSource(static_cast<QObject*>(parent));
	}
}

struct QtPositioning_PackedString QGeoAreaMonitorSource_QGeoAreaMonitorSource_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t0d0aae = QGeoAreaMonitorSource::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t0d0aae.prepend("WHITESPACE").constData()+10), t0d0aae.size()-10 }; });
}

struct QtPositioning_PackedString QGeoAreaMonitorSource_QGeoAreaMonitorSource_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t9898d6 = QGeoAreaMonitorSource::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t9898d6.prepend("WHITESPACE").constData()+10), t9898d6.size()-10 }; });
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
	Q_UNUSED(ptr);

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

void* QGeoAreaMonitorSource_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::metaObject());
}

void* QGeoAreaMonitorSource___activeMonitors_atList(void* ptr, int i)
{
	return new QGeoAreaMonitorInfo(({QGeoAreaMonitorInfo tmp = static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->at(i); if (i == static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->size()-1) { static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoAreaMonitorSource___activeMonitors_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->append(*static_cast<QGeoAreaMonitorInfo*>(i));
}

void* QGeoAreaMonitorSource___activeMonitors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoAreaMonitorInfo>();
}

void* QGeoAreaMonitorSource___activeMonitors_atList2(void* ptr, int i)
{
	return new QGeoAreaMonitorInfo(({QGeoAreaMonitorInfo tmp = static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->at(i); if (i == static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->size()-1) { static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoAreaMonitorSource___activeMonitors_setList2(void* ptr, void* i)
{
	static_cast<QList<QGeoAreaMonitorInfo>*>(ptr)->append(*static_cast<QGeoAreaMonitorInfo*>(i));
}

void* QGeoAreaMonitorSource___activeMonitors_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoAreaMonitorInfo>();
}

void* QGeoAreaMonitorSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoAreaMonitorSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoAreaMonitorSource___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoAreaMonitorSource___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoAreaMonitorSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoAreaMonitorSource___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoAreaMonitorSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoAreaMonitorSource___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoAreaMonitorSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoAreaMonitorSource___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoAreaMonitorSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoAreaMonitorSource___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QGeoAreaMonitorSource_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::event(static_cast<QEvent*>(e));
}

char QGeoAreaMonitorSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoAreaMonitorSource_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoAreaMonitorSource_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoAreaMonitorSource_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::customEvent(static_cast<QEvent*>(event));
}

void QGeoAreaMonitorSource_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::deleteLater();
}

void QGeoAreaMonitorSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoAreaMonitorSource_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoAreaMonitorSource*>(ptr)->QGeoAreaMonitorSource::timerEvent(static_cast<QTimerEvent*>(event));
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

void QGeoCircle_ExtendCircle(void* ptr, void* coordinate)
{
	static_cast<QGeoCircle*>(ptr)->extendCircle(*static_cast<QGeoCoordinate*>(coordinate));
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

void* QGeoCoordinate_AtDistanceAndAzimuth(void* ptr, double distance, double azimuth, double distanceUp)
{
	return new QGeoCoordinate(static_cast<QGeoCoordinate*>(ptr)->atDistanceAndAzimuth(distance, azimuth, distanceUp));
}

long long QGeoCoordinate_Type(void* ptr)
{
	return static_cast<QGeoCoordinate*>(ptr)->type();
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

void* QGeoPath_NewQGeoPath()
{
	return new QGeoPath();
}

void* QGeoPath_NewQGeoPath3(void* other)
{
	return new QGeoPath(*static_cast<QGeoPath*>(other));
}

void* QGeoPath_NewQGeoPath4(void* other)
{
	return new QGeoPath(*static_cast<QGeoShape*>(other));
}

void* QGeoPath_NewQGeoPath2(void* path, double width)
{
	return new QGeoPath(*static_cast<QList<QGeoCoordinate>*>(path), width);
}

void QGeoPath_AddCoordinate(void* ptr, void* coordinate)
{
	static_cast<QGeoPath*>(ptr)->addCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPath_ClearPath(void* ptr)
{
	static_cast<QGeoPath*>(ptr)->clearPath();
}

void QGeoPath_InsertCoordinate(void* ptr, int index, void* coordinate)
{
	static_cast<QGeoPath*>(ptr)->insertCoordinate(index, *static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPath_RemoveCoordinate(void* ptr, void* coordinate)
{
	static_cast<QGeoPath*>(ptr)->removeCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPath_RemoveCoordinate2(void* ptr, int index)
{
	static_cast<QGeoPath*>(ptr)->removeCoordinate(index);
}

void QGeoPath_ReplaceCoordinate(void* ptr, int index, void* coordinate)
{
	static_cast<QGeoPath*>(ptr)->replaceCoordinate(index, *static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPath_SetPath(void* ptr, void* path)
{
	static_cast<QGeoPath*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoPath_SetWidth(void* ptr, double width)
{
	static_cast<QGeoPath*>(ptr)->setWidth(width);
}

void QGeoPath_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoPath*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void QGeoPath_DestroyQGeoPath(void* ptr)
{
	static_cast<QGeoPath*>(ptr)->~QGeoPath();
}

void* QGeoPath_CoordinateAt(void* ptr, int index)
{
	return new QGeoCoordinate(static_cast<QGeoPath*>(ptr)->coordinateAt(index));
}

void* QGeoPath_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoPath(static_cast<QGeoPath*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

char QGeoPath_ContainsCoordinate(void* ptr, void* coordinate)
{
	return static_cast<QGeoPath*>(ptr)->containsCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

struct QtPositioning_PackedList QGeoPath_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = const_cast<QList<QGeoCoordinate>*>(&static_cast<QGeoPath*>(ptr)->path()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

double QGeoPath_Length(void* ptr, int indexFrom, int indexTo)
{
	return static_cast<QGeoPath*>(ptr)->length(indexFrom, indexTo);
}

int QGeoPath_Size(void* ptr)
{
	return static_cast<QGeoPath*>(ptr)->size();
}

double QGeoPath_Width(void* ptr)
{
	return static_cast<QGeoPath*>(ptr)->width();
}

void* QGeoPath___QGeoPath_path_atList2(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPath___QGeoPath_path_setList2(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPath___QGeoPath_path_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPath___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPath___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPath___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPath___setVariantPath_path_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPath___setVariantPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QGeoPath___setVariantPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QGeoPath___variantPath_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPath___variantPath_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QGeoPath___variantPath_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QGeoPath___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPath___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPath___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon_NewQGeoPolygon()
{
	return new QGeoPolygon();
}

void* QGeoPolygon_NewQGeoPolygon3(void* other)
{
	return new QGeoPolygon(*static_cast<QGeoPolygon*>(other));
}

void* QGeoPolygon_NewQGeoPolygon4(void* other)
{
	return new QGeoPolygon(*static_cast<QGeoShape*>(other));
}

void* QGeoPolygon_NewQGeoPolygon2(void* path)
{
	return new QGeoPolygon(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoPolygon_AddCoordinate(void* ptr, void* coordinate)
{
	static_cast<QGeoPolygon*>(ptr)->addCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPolygon_AddHole2(void* ptr, void* holePath)
{
	static_cast<QGeoPolygon*>(ptr)->addHole(*static_cast<QList<QGeoCoordinate>*>(holePath));
}

void QGeoPolygon_AddHole(void* ptr, void* holePath)
{
	static_cast<QGeoPolygon*>(ptr)->addHole(*static_cast<QVariant*>(holePath));
}

void QGeoPolygon_InsertCoordinate(void* ptr, int index, void* coordinate)
{
	static_cast<QGeoPolygon*>(ptr)->insertCoordinate(index, *static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPolygon_RemoveCoordinate(void* ptr, void* coordinate)
{
	static_cast<QGeoPolygon*>(ptr)->removeCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPolygon_RemoveCoordinate2(void* ptr, int index)
{
	static_cast<QGeoPolygon*>(ptr)->removeCoordinate(index);
}

void QGeoPolygon_RemoveHole(void* ptr, int index)
{
	static_cast<QGeoPolygon*>(ptr)->removeHole(index);
}

void QGeoPolygon_ReplaceCoordinate(void* ptr, int index, void* coordinate)
{
	static_cast<QGeoPolygon*>(ptr)->replaceCoordinate(index, *static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoPolygon_SetPath(void* ptr, void* path)
{
	static_cast<QGeoPolygon*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoPolygon_SetPerimeter(void* ptr, void* path)
{
	static_cast<QGeoPolygon*>(ptr)->setPerimeter(({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(path); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QGeoPolygon_Translate(void* ptr, double degreesLatitude, double degreesLongitude)
{
	static_cast<QGeoPolygon*>(ptr)->translate(degreesLatitude, degreesLongitude);
}

void QGeoPolygon_DestroyQGeoPolygon(void* ptr)
{
	static_cast<QGeoPolygon*>(ptr)->~QGeoPolygon();
}

void* QGeoPolygon_CoordinateAt(void* ptr, int index)
{
	return new QGeoCoordinate(static_cast<QGeoPolygon*>(ptr)->coordinateAt(index));
}

void* QGeoPolygon_Translated(void* ptr, double degreesLatitude, double degreesLongitude)
{
	return new QGeoPolygon(static_cast<QGeoPolygon*>(ptr)->translated(degreesLatitude, degreesLongitude));
}

struct QtPositioning_PackedList QGeoPolygon_Perimeter(void* ptr)
{
	return ({ QList<QVariant>* tmpValue = new QList<QVariant>(static_cast<QGeoPolygon*>(ptr)->perimeter()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

char QGeoPolygon_ContainsCoordinate(void* ptr, void* coordinate)
{
	return static_cast<QGeoPolygon*>(ptr)->containsCoordinate(*static_cast<QGeoCoordinate*>(coordinate));
}

struct QtPositioning_PackedList QGeoPolygon_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = const_cast<QList<QGeoCoordinate>*>(&static_cast<QGeoPolygon*>(ptr)->path()); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedList QGeoPolygon_HolePath(void* ptr, int index)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoPolygon*>(ptr)->holePath(index)); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtPositioning_PackedList QGeoPolygon_Hole(void* ptr, int index)
{
	return ({ QList<QVariant>* tmpValue = new QList<QVariant>(static_cast<QGeoPolygon*>(ptr)->hole(index)); QtPositioning_PackedList { tmpValue, tmpValue->size() }; });
}

double QGeoPolygon_Length(void* ptr, int indexFrom, int indexTo)
{
	return static_cast<QGeoPolygon*>(ptr)->length(indexFrom, indexTo);
}

int QGeoPolygon_HolesCount(void* ptr)
{
	return static_cast<QGeoPolygon*>(ptr)->holesCount();
}

int QGeoPolygon_Size(void* ptr)
{
	return static_cast<QGeoPolygon*>(ptr)->size();
}

void* QGeoPolygon___QGeoPolygon_path_atList2(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___QGeoPolygon_path_setList2(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPolygon___QGeoPolygon_path_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon___addHole_holePath_atList2(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___addHole_holePath_setList2(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPolygon___addHole_holePath_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPolygon___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon___setPerimeter_path_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___setPerimeter_path_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QGeoPolygon___setPerimeter_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QGeoPolygon___perimeter_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___perimeter_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QGeoPolygon___perimeter_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QGeoPolygon___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPolygon___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon___holePath_atList(void* ptr, int i)
{
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___holePath_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoPolygon___holePath_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
}

void* QGeoPolygon___hole_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPolygon___hole_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QGeoPolygon___hole_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
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
	MyQGeoPositionInfoSource(QObject *parent) : QGeoPositionInfoSource(parent) {QGeoPositionInfoSource_QGeoPositionInfoSource_QRegisterMetaType();};
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSource_Error2(this, positioningError); };
	void Signal_PositionUpdated(const QGeoPositionInfo & update) { callbackQGeoPositionInfoSource_PositionUpdated(this, const_cast<QGeoPositionInfo*>(&update)); };
	void requestUpdate(int timeout) { callbackQGeoPositionInfoSource_RequestUpdate(this, timeout); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	void setUpdateInterval(int msec) { callbackQGeoPositionInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQGeoPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoPositionInfoSource_StopUpdates(this); };
	void Signal_SupportedPositioningMethodsChanged() { callbackQGeoPositionInfoSource_SupportedPositioningMethodsChanged(this); };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSource_UpdateTimeout(this); };
	 ~MyQGeoPositionInfoSource() { callbackQGeoPositionInfoSource_DestroyQGeoPositionInfoSource(this); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQGeoPositionInfoSource_LastKnownPosition(const_cast<void*>(static_cast<const void*>(this)), fromSatellitePositioningMethodsOnly)); };
	QGeoPositionInfoSource::Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQGeoPositionInfoSource_Error(const_cast<void*>(static_cast<const void*>(this)))); };
	QGeoPositionInfoSource::PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQGeoPositionInfoSource_SupportedPositioningMethods(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoPositionInfoSource_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumUpdateInterval() const { return callbackQGeoPositionInfoSource_MinimumUpdateInterval(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQGeoPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoPositionInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoPositionInfoSource_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoPositionInfoSource_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPositioning_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoPositionInfoSource_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoPositionInfoSource_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQGeoPositionInfoSource*)

int QGeoPositionInfoSource_QGeoPositionInfoSource_QRegisterMetaType(){qRegisterMetaType<QGeoPositionInfoSource*>(); return qRegisterMetaType<MyQGeoPositionInfoSource*>();}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoPositionInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_QGeoPositionInfoSource_CreateSource(struct QtPositioning_PackedString sourceName, void* parent)
{
	return QGeoPositionInfoSource::createSource(QString::fromUtf8(sourceName.data, sourceName.len), static_cast<QObject*>(parent));
}

void* QGeoPositionInfoSource_NewQGeoPositionInfoSource(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoPositionInfoSource(static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoPositionInfoSource(static_cast<QObject*>(parent));
	}
}

struct QtPositioning_PackedString QGeoPositionInfoSource_QGeoPositionInfoSource_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t87fe70 = QGeoPositionInfoSource::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t87fe70.prepend("WHITESPACE").constData()+10), t87fe70.size()-10 }; });
}

struct QtPositioning_PackedString QGeoPositionInfoSource_QGeoPositionInfoSource_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray tcd97a8 = QGeoPositionInfoSource::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(tcd97a8.prepend("WHITESPACE").constData()+10), tcd97a8.size()-10 }; });
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
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setPreferredPositioningMethods(static_cast<QGeoPositionInfoSource::PositioningMethod>(methods));
	}
}

void QGeoPositionInfoSource_SetUpdateInterval(void* ptr, int msec)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->setUpdateInterval(msec);
}

void QGeoPositionInfoSource_SetUpdateIntervalDefault(void* ptr, int msec)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::setUpdateInterval(msec);
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::setUpdateInterval(msec);
	}
}

void QGeoPositionInfoSource_StartUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "startUpdates");
}

void QGeoPositionInfoSource_StopUpdates(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGeoPositionInfoSource*>(ptr), "stopUpdates");
}

void QGeoPositionInfoSource_ConnectSupportedPositioningMethodsChanged(void* ptr)
{
	QObject::connect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::supportedPositioningMethodsChanged), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_SupportedPositioningMethodsChanged));
}

void QGeoPositionInfoSource_DisconnectSupportedPositioningMethodsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGeoPositionInfoSource*>(ptr), static_cast<void (QGeoPositionInfoSource::*)()>(&QGeoPositionInfoSource::supportedPositioningMethodsChanged), static_cast<MyQGeoPositionInfoSource*>(ptr), static_cast<void (MyQGeoPositionInfoSource::*)()>(&MyQGeoPositionInfoSource::Signal_SupportedPositioningMethodsChanged));
}

void QGeoPositionInfoSource_SupportedPositioningMethodsChanged(void* ptr)
{
	static_cast<QGeoPositionInfoSource*>(ptr)->supportedPositioningMethodsChanged();
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
	Q_UNUSED(ptr);

}

void* QGeoPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QGeoPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
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

struct QtPositioning_PackedString QGeoPositionInfoSource_SourceName(void* ptr)
{
	return ({ QByteArray ta59e39 = static_cast<QGeoPositionInfoSource*>(ptr)->sourceName().toUtf8(); QtPositioning_PackedString { const_cast<char*>(ta59e39.prepend("WHITESPACE").constData()+10), ta59e39.size()-10 }; });
}

void* QGeoPositionInfoSource_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::metaObject());
	}
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
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoPositionInfoSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoPositionInfoSource___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoPositionInfoSource___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoPositionInfoSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoPositionInfoSource___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoPositionInfoSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoPositionInfoSource___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoPositionInfoSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoPositionInfoSource___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoPositionInfoSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoPositionInfoSource___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QGeoPositionInfoSource_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::event(static_cast<QEvent*>(e));
	}
}

char QGeoPositionInfoSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QGeoPositionInfoSource_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QGeoPositionInfoSource_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QGeoPositionInfoSource_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::customEvent(static_cast<QEvent*>(event));
	}
}

void QGeoPositionInfoSource_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::deleteLater();
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::deleteLater();
	}
}

void QGeoPositionInfoSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QGeoPositionInfoSource_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QNmeaPositionInfoSource*>(static_cast<QObject*>(ptr))) {
		static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QGeoPositionInfoSource*>(ptr)->QGeoPositionInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
	}
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
	Q_UNUSED(ptr);

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

void QGeoRectangle_ExtendRectangle(void* ptr, void* coordinate)
{
	static_cast<QGeoRectangle*>(ptr)->extendRectangle(*static_cast<QGeoCoordinate*>(coordinate));
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
	return new QGeoCoordinate(({QGeoCoordinate tmp = static_cast<QList<QGeoCoordinate>*>(ptr)->at(i); if (i == static_cast<QList<QGeoCoordinate>*>(ptr)->size()-1) { static_cast<QList<QGeoCoordinate>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoRectangle___QGeoRectangle_coordinates_setList4(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRectangle___QGeoRectangle_coordinates_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>();
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
	MyQGeoSatelliteInfoSource(QObject *parent) : QGeoSatelliteInfoSource(parent) {QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_QRegisterMetaType();};
	void Signal_Error2(QGeoSatelliteInfoSource::Error satelliteError) { callbackQGeoSatelliteInfoSource_Error2(this, satelliteError); };
	void Signal_RequestTimeout() { callbackQGeoSatelliteInfoSource_RequestTimeout(this); };
	void requestUpdate(int timeout) { callbackQGeoSatelliteInfoSource_RequestUpdate(this, timeout); };
	void Signal_SatellitesInUseUpdated(const QList<QGeoSatelliteInfo> & satellites) { callbackQGeoSatelliteInfoSource_SatellitesInUseUpdated(this, ({ QList<QGeoSatelliteInfo>* tmpValue = const_cast<QList<QGeoSatelliteInfo>*>(&satellites); QtPositioning_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_SatellitesInViewUpdated(const QList<QGeoSatelliteInfo> & satellites) { callbackQGeoSatelliteInfoSource_SatellitesInViewUpdated(this, ({ QList<QGeoSatelliteInfo>* tmpValue = const_cast<QList<QGeoSatelliteInfo>*>(&satellites); QtPositioning_PackedList { tmpValue, tmpValue->size() }; })); };
	void setUpdateInterval(int msec) { callbackQGeoSatelliteInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQGeoSatelliteInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQGeoSatelliteInfoSource_StopUpdates(this); };
	 ~MyQGeoSatelliteInfoSource() { callbackQGeoSatelliteInfoSource_DestroyQGeoSatelliteInfoSource(this); };
	QGeoSatelliteInfoSource::Error error() const { return static_cast<QGeoSatelliteInfoSource::Error>(callbackQGeoSatelliteInfoSource_Error(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoSatelliteInfoSource_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumUpdateInterval() const { return callbackQGeoSatelliteInfoSource_MinimumUpdateInterval(const_cast<void*>(static_cast<const void*>(this))); };
	bool event(QEvent * e) { return callbackQGeoSatelliteInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoSatelliteInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoSatelliteInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoSatelliteInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoSatelliteInfoSource_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoSatelliteInfoSource_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoSatelliteInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPositioning_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoSatelliteInfoSource_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoSatelliteInfoSource_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQGeoSatelliteInfoSource*)

int QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_QRegisterMetaType(){qRegisterMetaType<QGeoSatelliteInfoSource*>(); return qRegisterMetaType<MyQGeoSatelliteInfoSource*>();}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateDefaultSource(void* parent)
{
	return QGeoSatelliteInfoSource::createDefaultSource(static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_CreateSource(struct QtPositioning_PackedString sourceName, void* parent)
{
	return QGeoSatelliteInfoSource::createSource(QString::fromUtf8(sourceName.data, sourceName.len), static_cast<QObject*>(parent));
}

void* QGeoSatelliteInfoSource_NewQGeoSatelliteInfoSource(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoSatelliteInfoSource(static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoSatelliteInfoSource(static_cast<QObject*>(parent));
	}
}

struct QtPositioning_PackedString QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_Tr(char* s, char* c, int n)
{
	return ({ QByteArray t7435c8 = QGeoSatelliteInfoSource::tr(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t7435c8.prepend("WHITESPACE").constData()+10), t7435c8.size()-10 }; });
}

struct QtPositioning_PackedString QGeoSatelliteInfoSource_QGeoSatelliteInfoSource_TrUtf8(char* s, char* c, int n)
{
	return ({ QByteArray t66bd78 = QGeoSatelliteInfoSource::trUtf8(const_cast<const char*>(s), const_cast<const char*>(c), n).toUtf8(); QtPositioning_PackedString { const_cast<char*>(t66bd78.prepend("WHITESPACE").constData()+10), t66bd78.size()-10 }; });
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
	Q_UNUSED(ptr);

}

long long QGeoSatelliteInfoSource_Error(void* ptr)
{
	return static_cast<QGeoSatelliteInfoSource*>(ptr)->error();
}

struct QtPositioning_PackedString QGeoSatelliteInfoSource_SourceName(void* ptr)
{
	return ({ QByteArray t6fe95a = static_cast<QGeoSatelliteInfoSource*>(ptr)->sourceName().toUtf8(); QtPositioning_PackedString { const_cast<char*>(t6fe95a.prepend("WHITESPACE").constData()+10), t6fe95a.size()-10 }; });
}

void* QGeoSatelliteInfoSource_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::metaObject());
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
	return new QGeoSatelliteInfo(({QGeoSatelliteInfo tmp = static_cast<QList<QGeoSatelliteInfo>*>(ptr)->at(i); if (i == static_cast<QList<QGeoSatelliteInfo>*>(ptr)->size()-1) { static_cast<QList<QGeoSatelliteInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoSatelliteInfo>*>(ptr)->append(*static_cast<QGeoSatelliteInfo*>(i));
}

void* QGeoSatelliteInfoSource___satellitesInUseUpdated_satellites_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoSatelliteInfo>();
}

void* QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_atList(void* ptr, int i)
{
	return new QGeoSatelliteInfo(({QGeoSatelliteInfo tmp = static_cast<QList<QGeoSatelliteInfo>*>(ptr)->at(i); if (i == static_cast<QList<QGeoSatelliteInfo>*>(ptr)->size()-1) { static_cast<QList<QGeoSatelliteInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoSatelliteInfo>*>(ptr)->append(*static_cast<QGeoSatelliteInfo*>(i));
}

void* QGeoSatelliteInfoSource___satellitesInViewUpdated_satellites_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoSatelliteInfo>();
}

void* QGeoSatelliteInfoSource___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGeoSatelliteInfoSource___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoSatelliteInfoSource___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGeoSatelliteInfoSource___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoSatelliteInfoSource___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoSatelliteInfoSource___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoSatelliteInfoSource___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoSatelliteInfoSource___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoSatelliteInfoSource___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGeoSatelliteInfoSource___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGeoSatelliteInfoSource___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoSatelliteInfoSource___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

char QGeoSatelliteInfoSource_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::event(static_cast<QEvent*>(e));
}

char QGeoSatelliteInfoSource_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoSatelliteInfoSource_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoSatelliteInfoSource_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoSatelliteInfoSource_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::customEvent(static_cast<QEvent*>(event));
}

void QGeoSatelliteInfoSource_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::deleteLater();
}

void QGeoSatelliteInfoSource_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoSatelliteInfoSource_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoSatelliteInfoSource*>(ptr)->QGeoSatelliteInfoSource::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGeoShape_NewQGeoShape()
{
	return new QGeoShape();
}

void* QGeoShape_NewQGeoShape2(void* other)
{
	return new QGeoShape(*static_cast<QGeoShape*>(other));
}

void QGeoShape_DestroyQGeoShape(void* ptr)
{
	static_cast<QGeoShape*>(ptr)->~QGeoShape();
}

void* QGeoShape_Center(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoShape*>(ptr)->center());
}

void* QGeoShape_BoundingGeoRectangle(void* ptr)
{
	return new QGeoRectangle(static_cast<QGeoShape*>(ptr)->boundingGeoRectangle());
}

long long QGeoShape_Type(void* ptr)
{
	return static_cast<QGeoShape*>(ptr)->type();
}

struct QtPositioning_PackedString QGeoShape_ToString(void* ptr)
{
	return ({ QByteArray teed09f = static_cast<QGeoShape*>(ptr)->toString().toUtf8(); QtPositioning_PackedString { const_cast<char*>(teed09f.prepend("WHITESPACE").constData()+10), teed09f.size()-10 }; });
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
	MyQNmeaPositionInfoSource(QNmeaPositionInfoSource::UpdateMode updateMode, QObject *parent = Q_NULLPTR) : QNmeaPositionInfoSource(updateMode, parent) {QNmeaPositionInfoSource_QNmeaPositionInfoSource_QRegisterMetaType();};
	bool parsePosInfoFromNmeaData(const char * data, int size, QGeoPositionInfo * posInfo, bool * hasFix) { QtPositioning_PackedString dataPacked = { const_cast<char*>(data), size };return callbackQNmeaPositionInfoSource_ParsePosInfoFromNmeaData(this, dataPacked, size, posInfo, reinterpret_cast<char*>(hasFix)) != 0; };
	void requestUpdate(int msec) { callbackQNmeaPositionInfoSource_RequestUpdate(this, msec); };
	void setUpdateInterval(int msec) { callbackQGeoPositionInfoSource_SetUpdateInterval(this, msec); };
	void startUpdates() { callbackQNmeaPositionInfoSource_StartUpdates(this); };
	void stopUpdates() { callbackQNmeaPositionInfoSource_StopUpdates(this); };
	 ~MyQNmeaPositionInfoSource() { callbackQNmeaPositionInfoSource_DestroyQNmeaPositionInfoSource(this); };
	QGeoPositionInfo lastKnownPosition(bool fromSatellitePositioningMethodsOnly) const { return *static_cast<QGeoPositionInfo*>(callbackQNmeaPositionInfoSource_LastKnownPosition(const_cast<void*>(static_cast<const void*>(this)), fromSatellitePositioningMethodsOnly)); };
	QGeoPositionInfoSource::Error error() const { return static_cast<QGeoPositionInfoSource::Error>(callbackQNmeaPositionInfoSource_Error(const_cast<void*>(static_cast<const void*>(this)))); };
	QGeoPositionInfoSource::PositioningMethods supportedPositioningMethods() const { return static_cast<QGeoPositionInfoSource::PositioningMethod>(callbackQNmeaPositionInfoSource_SupportedPositioningMethods(const_cast<void*>(static_cast<const void*>(this)))); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoPositionInfoSource_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	int minimumUpdateInterval() const { return callbackQNmeaPositionInfoSource_MinimumUpdateInterval(const_cast<void*>(static_cast<const void*>(this))); };
	void Signal_Error2(QGeoPositionInfoSource::Error positioningError) { callbackQGeoPositionInfoSource_Error2(this, positioningError); };
	void Signal_PositionUpdated(const QGeoPositionInfo & update) { callbackQGeoPositionInfoSource_PositionUpdated(this, const_cast<QGeoPositionInfo*>(&update)); };
	void setPreferredPositioningMethods(QGeoPositionInfoSource::PositioningMethods methods) { callbackQGeoPositionInfoSource_SetPreferredPositioningMethods(this, methods); };
	void Signal_SupportedPositioningMethodsChanged() { callbackQGeoPositionInfoSource_SupportedPositioningMethodsChanged(this); };
	void Signal_UpdateTimeout() { callbackQGeoPositionInfoSource_UpdateTimeout(this); };
	bool event(QEvent * e) { return callbackQGeoPositionInfoSource_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoPositionInfoSource_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoPositionInfoSource_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoPositionInfoSource_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoPositionInfoSource_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoPositionInfoSource_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoPositionInfoSource_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtPositioning_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoPositionInfoSource_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoPositionInfoSource_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQNmeaPositionInfoSource*)

int QNmeaPositionInfoSource_QNmeaPositionInfoSource_QRegisterMetaType(){qRegisterMetaType<QNmeaPositionInfoSource*>(); return qRegisterMetaType<MyQNmeaPositionInfoSource*>();}

void* QNmeaPositionInfoSource_NewQNmeaPositionInfoSource(long long updateMode, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QWindow*>(parent));
	} else {
		return new MyQNmeaPositionInfoSource(static_cast<QNmeaPositionInfoSource::UpdateMode>(updateMode), static_cast<QObject*>(parent));
	}
}

char QNmeaPositionInfoSource_ParsePosInfoFromNmeaData(void* ptr, char* data, int size, void* posInfo, char* hasFix)
{
	return static_cast<QNmeaPositionInfoSource*>(ptr)->parsePosInfoFromNmeaData(const_cast<const char*>(data), size, static_cast<QGeoPositionInfo*>(posInfo), reinterpret_cast<bool*>(hasFix));
}

char QNmeaPositionInfoSource_ParsePosInfoFromNmeaDataDefault(void* ptr, char* data, int size, void* posInfo, char* hasFix)
{
		return static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::parsePosInfoFromNmeaData(const_cast<const char*>(data), size, static_cast<QGeoPositionInfo*>(posInfo), reinterpret_cast<bool*>(hasFix));
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

void QNmeaPositionInfoSource_DestroyQNmeaPositionInfoSourceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QNmeaPositionInfoSource_LastKnownPosition(void* ptr, char fromSatellitePositioningMethodsOnly)
{
	return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
}

void* QNmeaPositionInfoSource_LastKnownPositionDefault(void* ptr, char fromSatellitePositioningMethodsOnly)
{
		return new QGeoPositionInfo(static_cast<QNmeaPositionInfoSource*>(ptr)->QNmeaPositionInfoSource::lastKnownPosition(fromSatellitePositioningMethodsOnly != 0));
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

