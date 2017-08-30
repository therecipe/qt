// +build !minimal

#define protected public
#define private public

#include "location.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGeoCoordinate>
#include <QGeoManeuver>
#include <QGeoRectangle>
#include <QGeoRoute>
#include <QGeoRouteReply>
#include <QGeoRouteRequest>
#include <QGeoRouteSegment>
#include <QGeoRoutingManager>
#include <QGeoRoutingManagerEngine>
#include <QGeoServiceProvider>
#include <QGeoServiceProviderFactory>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLayout>
#include <QList>
#include <QLocale>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

void* QGeoManeuver_NewQGeoManeuver()
{
	return new QGeoManeuver();
}

void* QGeoManeuver_NewQGeoManeuver2(void* other)
{
	return new QGeoManeuver(*static_cast<QGeoManeuver*>(other));
}

void QGeoManeuver_SetDirection(void* ptr, long long direction)
{
	static_cast<QGeoManeuver*>(ptr)->setDirection(static_cast<QGeoManeuver::InstructionDirection>(direction));
}

void QGeoManeuver_SetDistanceToNextInstruction(void* ptr, double distance)
{
	static_cast<QGeoManeuver*>(ptr)->setDistanceToNextInstruction(distance);
}

void QGeoManeuver_SetInstructionText(void* ptr, struct QtLocation_PackedString instructionText)
{
	static_cast<QGeoManeuver*>(ptr)->setInstructionText(QString::fromUtf8(instructionText.data, instructionText.len));
}

void QGeoManeuver_SetPosition(void* ptr, void* position)
{
	static_cast<QGeoManeuver*>(ptr)->setPosition(*static_cast<QGeoCoordinate*>(position));
}

void QGeoManeuver_SetTimeToNextInstruction(void* ptr, int secs)
{
	static_cast<QGeoManeuver*>(ptr)->setTimeToNextInstruction(secs);
}

void QGeoManeuver_SetWaypoint(void* ptr, void* coordinate)
{
	static_cast<QGeoManeuver*>(ptr)->setWaypoint(*static_cast<QGeoCoordinate*>(coordinate));
}

void QGeoManeuver_DestroyQGeoManeuver(void* ptr)
{
	static_cast<QGeoManeuver*>(ptr)->~QGeoManeuver();
}

long long QGeoManeuver_Direction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->direction();
}

void* QGeoManeuver_Position(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->position());
}

void* QGeoManeuver_Waypoint(void* ptr)
{
	return new QGeoCoordinate(static_cast<QGeoManeuver*>(ptr)->waypoint());
}

struct QtLocation_PackedString QGeoManeuver_InstructionText(void* ptr)
{
	return ({ QByteArray t4cc4b3 = static_cast<QGeoManeuver*>(ptr)->instructionText().toUtf8(); QtLocation_PackedString { const_cast<char*>(t4cc4b3.prepend("WHITESPACE").constData()+10), t4cc4b3.size()-10 }; });
}

char QGeoManeuver_IsValid(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->isValid();
}

int QGeoManeuver_TimeToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->timeToNextInstruction();
}

double QGeoManeuver_DistanceToNextInstruction(void* ptr)
{
	return static_cast<QGeoManeuver*>(ptr)->distanceToNextInstruction();
}

void* QGeoRoute_NewQGeoRoute()
{
	return new QGeoRoute();
}

void* QGeoRoute_NewQGeoRoute2(void* other)
{
	return new QGeoRoute(*static_cast<QGeoRoute*>(other));
}

void QGeoRoute_SetBounds(void* ptr, void* bounds)
{
	static_cast<QGeoRoute*>(ptr)->setBounds(*static_cast<QGeoRectangle*>(bounds));
}

void QGeoRoute_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRoute*>(ptr)->setDistance(distance);
}

void QGeoRoute_SetFirstRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRoute*>(ptr)->setFirstRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRoute_SetPath(void* ptr, void* path)
{
	static_cast<QGeoRoute*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoRoute_SetRequest(void* ptr, void* request)
{
	static_cast<QGeoRoute*>(ptr)->setRequest(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoute_SetRouteId(void* ptr, struct QtLocation_PackedString id)
{
	static_cast<QGeoRoute*>(ptr)->setRouteId(QString::fromUtf8(id.data, id.len));
}

void QGeoRoute_SetTravelMode(void* ptr, long long mode)
{
	static_cast<QGeoRoute*>(ptr)->setTravelMode(static_cast<QGeoRouteRequest::TravelMode>(mode));
}

void QGeoRoute_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRoute*>(ptr)->setTravelTime(secs);
}

void QGeoRoute_DestroyQGeoRoute(void* ptr)
{
	static_cast<QGeoRoute*>(ptr)->~QGeoRoute();
}

void* QGeoRoute_Bounds(void* ptr)
{
	return new QGeoRectangle(static_cast<QGeoRoute*>(ptr)->bounds());
}

void* QGeoRoute_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRoute*>(ptr)->request());
}

long long QGeoRoute_TravelMode(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelMode();
}

void* QGeoRoute_FirstRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRoute*>(ptr)->firstRouteSegment());
}

struct QtLocation_PackedList QGeoRoute_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRoute*>(ptr)->path()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedString QGeoRoute_RouteId(void* ptr)
{
	return ({ QByteArray t7492be = static_cast<QGeoRoute*>(ptr)->routeId().toUtf8(); QtLocation_PackedString { const_cast<char*>(t7492be.prepend("WHITESPACE").constData()+10), t7492be.size()-10 }; });
}

int QGeoRoute_TravelTime(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->travelTime();
}

double QGeoRoute_Distance(void* ptr)
{
	return static_cast<QGeoRoute*>(ptr)->distance();
}

void* QGeoRoute___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRoute___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRoute___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

void* QGeoRoute___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRoute___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRoute___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

class MyQGeoRouteReply: public QGeoRouteReply
{
public:
	MyQGeoRouteReply(Error error, const QString &errorString, QObject *parent = Q_NULLPTR) : QGeoRouteReply(error, errorString, parent) {QGeoRouteReply_QGeoRouteReply_QRegisterMetaType();};
	MyQGeoRouteReply(const QGeoRouteRequest &request, QObject *parent = Q_NULLPTR) : QGeoRouteReply(request, parent) {QGeoRouteReply_QGeoRouteReply_QRegisterMetaType();};
	void abort() { callbackQGeoRouteReply_Abort(this); };
	void Signal_Aborted() { callbackQGeoRouteReply_Aborted(this); };
	void Signal_Error2(QGeoRouteReply::Error error, const QString & errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRouteReply_Error2(this, error, errorStringPacked); };
	void Signal_Finished() { callbackQGeoRouteReply_Finished(this); };
	 ~MyQGeoRouteReply() { callbackQGeoRouteReply_DestroyQGeoRouteReply(this); };
	bool event(QEvent * e) { return callbackQGeoRouteReply_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRouteReply_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoRouteReply_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRouteReply_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRouteReply_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRouteReply_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRouteReply_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoRouteReply_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRouteReply_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRouteReply_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGeoRouteReply*)

int QGeoRouteReply_QGeoRouteReply_QRegisterMetaType(){qRegisterMetaType<QGeoRouteReply*>(); return qRegisterMetaType<MyQGeoRouteReply*>();}

void* QGeoRouteReply_NewQGeoRouteReply(long long error, struct QtLocation_PackedString errorString, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRouteReply(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len), static_cast<QObject*>(parent));
	}
}

void* QGeoRouteReply_NewQGeoRouteReply2(void* request, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRouteReply(*static_cast<QGeoRouteRequest*>(request), static_cast<QObject*>(parent));
	}
}

void QGeoRouteReply_Abort(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->abort();
}

void QGeoRouteReply_AbortDefault(void* ptr)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::abort();
}

void QGeoRouteReply_ConnectAborted(void* ptr)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::aborted), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Aborted));
}

void QGeoRouteReply_DisconnectAborted(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::aborted), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Aborted));
}

void QGeoRouteReply_Aborted(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->aborted();
}

void QGeoRouteReply_AddRoutes(void* ptr, void* routes)
{
	static_cast<QGeoRouteReply*>(ptr)->addRoutes(*static_cast<QList<QGeoRoute>*>(routes));
}

void QGeoRouteReply_ConnectError2(void* ptr)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2));
}

void QGeoRouteReply_DisconnectError2(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&QGeoRouteReply::error), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)(QGeoRouteReply::Error, const QString &)>(&MyQGeoRouteReply::Signal_Error2));
}

void QGeoRouteReply_Error2(void* ptr, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->error(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRouteReply_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));
}

void QGeoRouteReply_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRouteReply*>(ptr), static_cast<void (QGeoRouteReply::*)()>(&QGeoRouteReply::finished), static_cast<MyQGeoRouteReply*>(ptr), static_cast<void (MyQGeoRouteReply::*)()>(&MyQGeoRouteReply::Signal_Finished));
}

void QGeoRouteReply_Finished(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->finished();
}

void QGeoRouteReply_SetError(void* ptr, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRouteReply*>(ptr)->setError(static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRouteReply_SetFinished(void* ptr, char finished)
{
	static_cast<QGeoRouteReply*>(ptr)->setFinished(finished != 0);
}

void QGeoRouteReply_SetRoutes(void* ptr, void* routes)
{
	static_cast<QGeoRouteReply*>(ptr)->setRoutes(*static_cast<QList<QGeoRoute>*>(routes));
}

void QGeoRouteReply_DestroyQGeoRouteReply(void* ptr)
{
	static_cast<QGeoRouteReply*>(ptr)->~QGeoRouteReply();
}

void QGeoRouteReply_DestroyQGeoRouteReplyDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QGeoRouteReply_Error(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->error();
}

void* QGeoRouteReply_Request(void* ptr)
{
	return new QGeoRouteRequest(static_cast<QGeoRouteReply*>(ptr)->request());
}

struct QtLocation_PackedList QGeoRouteReply_Routes(void* ptr)
{
	return ({ QList<QGeoRoute>* tmpValue = new QList<QGeoRoute>(static_cast<QGeoRouteReply*>(ptr)->routes()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedString QGeoRouteReply_ErrorString(void* ptr)
{
	return ({ QByteArray t834aee = static_cast<QGeoRouteReply*>(ptr)->errorString().toUtf8(); QtLocation_PackedString { const_cast<char*>(t834aee.prepend("WHITESPACE").constData()+10), t834aee.size()-10 }; });
}

char QGeoRouteReply_IsFinished(void* ptr)
{
	return static_cast<QGeoRouteReply*>(ptr)->isFinished();
}

void* QGeoRouteReply___addRoutes_routes_atList(void* ptr, int i)
{
	return new QGeoRoute(static_cast<QList<QGeoRoute>*>(ptr)->at(i));
}

void QGeoRouteReply___addRoutes_routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___addRoutes_routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>;
}

void* QGeoRouteReply___setRoutes_routes_atList(void* ptr, int i)
{
	return new QGeoRoute(static_cast<QList<QGeoRoute>*>(ptr)->at(i));
}

void QGeoRouteReply___setRoutes_routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___setRoutes_routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>;
}

void* QGeoRouteReply___routes_atList(void* ptr, int i)
{
	return new QGeoRoute(static_cast<QList<QGeoRoute>*>(ptr)->at(i));
}

void QGeoRouteReply___routes_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRoute>*>(ptr)->append(*static_cast<QGeoRoute*>(i));
}

void* QGeoRouteReply___routes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRoute>;
}

void* QGeoRouteReply___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoRouteReply___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRouteReply___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGeoRouteReply___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRouteReply___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRouteReply___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRouteReply___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRouteReply___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRouteReply___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRouteReply___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoRouteReply___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRouteReply___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGeoRouteReply_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::event(static_cast<QEvent*>(e));
}

char QGeoRouteReply_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoRouteReply_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRouteReply_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::customEvent(static_cast<QEvent*>(event));
}

void QGeoRouteReply_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::deleteLater();
}

void QGeoRouteReply_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRouteReply_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGeoRouteReply_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRouteReply*>(ptr)->QGeoRouteReply::metaObject());
}

void* QGeoRouteRequest_NewQGeoRouteRequest2(void* origin, void* destination)
{
	return new QGeoRouteRequest(*static_cast<QGeoCoordinate*>(origin), *static_cast<QGeoCoordinate*>(destination));
}

void* QGeoRouteRequest_NewQGeoRouteRequest3(void* other)
{
	return new QGeoRouteRequest(*static_cast<QGeoRouteRequest*>(other));
}

void* QGeoRouteRequest_NewQGeoRouteRequest(void* waypoints)
{
	return new QGeoRouteRequest(*static_cast<QList<QGeoCoordinate>*>(waypoints));
}

void QGeoRouteRequest_SetExcludeAreas(void* ptr, void* areas)
{
	static_cast<QGeoRouteRequest*>(ptr)->setExcludeAreas(*static_cast<QList<QGeoRectangle>*>(areas));
}

void QGeoRouteRequest_SetFeatureWeight(void* ptr, long long featureType, long long featureWeight)
{
	static_cast<QGeoRouteRequest*>(ptr)->setFeatureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType), static_cast<QGeoRouteRequest::FeatureWeight>(featureWeight));
}

void QGeoRouteRequest_SetManeuverDetail(void* ptr, long long maneuverDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setManeuverDetail(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetail));
}

void QGeoRouteRequest_SetNumberAlternativeRoutes(void* ptr, int alternatives)
{
	static_cast<QGeoRouteRequest*>(ptr)->setNumberAlternativeRoutes(alternatives);
}

void QGeoRouteRequest_SetRouteOptimization(void* ptr, long long optimization)
{
	static_cast<QGeoRouteRequest*>(ptr)->setRouteOptimization(static_cast<QGeoRouteRequest::RouteOptimization>(optimization));
}

void QGeoRouteRequest_SetSegmentDetail(void* ptr, long long segmentDetail)
{
	static_cast<QGeoRouteRequest*>(ptr)->setSegmentDetail(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetail));
}

void QGeoRouteRequest_SetTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRouteRequest*>(ptr)->setTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

void QGeoRouteRequest_SetWaypoints(void* ptr, void* waypoints)
{
	static_cast<QGeoRouteRequest*>(ptr)->setWaypoints(*static_cast<QList<QGeoCoordinate>*>(waypoints));
}

void QGeoRouteRequest_DestroyQGeoRouteRequest(void* ptr)
{
	static_cast<QGeoRouteRequest*>(ptr)->~QGeoRouteRequest();
}

long long QGeoRouteRequest_FeatureWeight(void* ptr, long long featureType)
{
	return static_cast<QGeoRouteRequest*>(ptr)->featureWeight(static_cast<QGeoRouteRequest::FeatureType>(featureType));
}

long long QGeoRouteRequest_ManeuverDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->maneuverDetail();
}

struct QtLocation_PackedList QGeoRouteRequest_Waypoints(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRouteRequest*>(ptr)->waypoints()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedList QGeoRouteRequest_ExcludeAreas(void* ptr)
{
	return ({ QList<QGeoRectangle>* tmpValue = new QList<QGeoRectangle>(static_cast<QGeoRouteRequest*>(ptr)->excludeAreas()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

long long QGeoRouteRequest_RouteOptimization(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->routeOptimization();
}

long long QGeoRouteRequest_SegmentDetail(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->segmentDetail();
}

long long QGeoRouteRequest_TravelModes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->travelModes();
}

int QGeoRouteRequest_NumberAlternativeRoutes(void* ptr)
{
	return static_cast<QGeoRouteRequest*>(ptr)->numberAlternativeRoutes();
}

void* QGeoRouteRequest___QGeoRouteRequest_waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRouteRequest___QGeoRouteRequest_waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___QGeoRouteRequest_waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

void* QGeoRouteRequest___setExcludeAreas_areas_atList(void* ptr, int i)
{
	return new QGeoRectangle(static_cast<QList<QGeoRectangle>*>(ptr)->at(i));
}

void QGeoRouteRequest___setExcludeAreas_areas_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRectangle>*>(ptr)->append(*static_cast<QGeoRectangle*>(i));
}

void* QGeoRouteRequest___setExcludeAreas_areas_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRectangle>;
}

void* QGeoRouteRequest___setWaypoints_waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRouteRequest___setWaypoints_waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___setWaypoints_waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

void* QGeoRouteRequest___waypoints_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRouteRequest___waypoints_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteRequest___waypoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

void* QGeoRouteRequest___excludeAreas_atList(void* ptr, int i)
{
	return new QGeoRectangle(static_cast<QList<QGeoRectangle>*>(ptr)->at(i));
}

void QGeoRouteRequest___excludeAreas_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoRectangle>*>(ptr)->append(*static_cast<QGeoRectangle*>(i));
}

void* QGeoRouteRequest___excludeAreas_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoRectangle>;
}

void* QGeoRouteSegment_NewQGeoRouteSegment()
{
	return new QGeoRouteSegment();
}

void* QGeoRouteSegment_NewQGeoRouteSegment2(void* other)
{
	return new QGeoRouteSegment(*static_cast<QGeoRouteSegment*>(other));
}

void QGeoRouteSegment_SetDistance(void* ptr, double distance)
{
	static_cast<QGeoRouteSegment*>(ptr)->setDistance(distance);
}

void QGeoRouteSegment_SetManeuver(void* ptr, void* maneuver)
{
	static_cast<QGeoRouteSegment*>(ptr)->setManeuver(*static_cast<QGeoManeuver*>(maneuver));
}

void QGeoRouteSegment_SetNextRouteSegment(void* ptr, void* routeSegment)
{
	static_cast<QGeoRouteSegment*>(ptr)->setNextRouteSegment(*static_cast<QGeoRouteSegment*>(routeSegment));
}

void QGeoRouteSegment_SetPath(void* ptr, void* path)
{
	static_cast<QGeoRouteSegment*>(ptr)->setPath(*static_cast<QList<QGeoCoordinate>*>(path));
}

void QGeoRouteSegment_SetTravelTime(void* ptr, int secs)
{
	static_cast<QGeoRouteSegment*>(ptr)->setTravelTime(secs);
}

void QGeoRouteSegment_DestroyQGeoRouteSegment(void* ptr)
{
	static_cast<QGeoRouteSegment*>(ptr)->~QGeoRouteSegment();
}

void* QGeoRouteSegment_Maneuver(void* ptr)
{
	return new QGeoManeuver(static_cast<QGeoRouteSegment*>(ptr)->maneuver());
}

void* QGeoRouteSegment_NextRouteSegment(void* ptr)
{
	return new QGeoRouteSegment(static_cast<QGeoRouteSegment*>(ptr)->nextRouteSegment());
}

struct QtLocation_PackedList QGeoRouteSegment_Path(void* ptr)
{
	return ({ QList<QGeoCoordinate>* tmpValue = new QList<QGeoCoordinate>(static_cast<QGeoRouteSegment*>(ptr)->path()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

char QGeoRouteSegment_IsValid(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->isValid();
}

int QGeoRouteSegment_TravelTime(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->travelTime();
}

double QGeoRouteSegment_Distance(void* ptr)
{
	return static_cast<QGeoRouteSegment*>(ptr)->distance();
}

void* QGeoRouteSegment___setPath_path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRouteSegment___setPath_path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteSegment___setPath_path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

void* QGeoRouteSegment___path_atList(void* ptr, int i)
{
	return new QGeoCoordinate(static_cast<QList<QGeoCoordinate>*>(ptr)->at(i));
}

void QGeoRouteSegment___path_setList(void* ptr, void* i)
{
	static_cast<QList<QGeoCoordinate>*>(ptr)->append(*static_cast<QGeoCoordinate*>(i));
}

void* QGeoRouteSegment___path_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QGeoCoordinate>;
}

class MyQGeoRoutingManager: public QGeoRoutingManager
{
public:
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRoutingManager_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManager_Finished(this, reply); };
	bool event(QEvent * e) { return callbackQGeoRoutingManager_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManager_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManager_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManager_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManager_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRoutingManager_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManager_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoRoutingManager_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManager_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManager_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGeoRoutingManager*)

int QGeoRoutingManager_QGeoRoutingManager_QRegisterMetaType(){qRegisterMetaType<QGeoRoutingManager*>(); return qRegisterMetaType<MyQGeoRoutingManager*>();}

void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManager*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManager*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManager_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));
}

void QGeoRoutingManager_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));
}

void QGeoRoutingManager_Error(void* ptr, void* reply, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRoutingManager*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRoutingManager_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));
}

void QGeoRoutingManager_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));
}

void QGeoRoutingManager_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManager*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void QGeoRoutingManager_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManager*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManager_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManager*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr)
{
	static_cast<QGeoRoutingManager*>(ptr)->~QGeoRoutingManager();
}

long long QGeoRoutingManager_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManager_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManager_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManager_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManager_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManager_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManager_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManager*>(ptr)->locale());
}

long long QGeoRoutingManager_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->measurementSystem();
}

struct QtLocation_PackedString QGeoRoutingManager_ManagerName(void* ptr)
{
	return ({ QByteArray t48ee82 = static_cast<QGeoRoutingManager*>(ptr)->managerName().toUtf8(); QtLocation_PackedString { const_cast<char*>(t48ee82.prepend("WHITESPACE").constData()+10), t48ee82.size()-10 }; });
}

int QGeoRoutingManager_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManager*>(ptr)->managerVersion();
}

void* QGeoRoutingManager___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoRoutingManager___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRoutingManager___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGeoRoutingManager___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManager___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManager___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManager___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManager___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManager___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManager___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoRoutingManager___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManager___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGeoRoutingManager_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoRoutingManager_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManager_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManager_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::deleteLater();
}

void QGeoRoutingManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManager_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGeoRoutingManager_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRoutingManager*>(ptr)->QGeoRoutingManager::metaObject());
}

class MyQGeoRoutingManagerEngine: public QGeoRoutingManagerEngine
{
public:
	MyQGeoRoutingManagerEngine(const QVariantMap &parameters, QObject *parent = Q_NULLPTR) : QGeoRoutingManagerEngine(parameters, parent) {QGeoRoutingManagerEngine_QGeoRoutingManagerEngine_QRegisterMetaType();};
	QGeoRouteReply * calculateRoute(const QGeoRouteRequest & request) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_CalculateRoute(this, const_cast<QGeoRouteRequest*>(&request))); };
	QGeoRouteReply * updateRoute(const QGeoRoute & route, const QGeoCoordinate & position) { return static_cast<QGeoRouteReply*>(callbackQGeoRoutingManagerEngine_UpdateRoute(this, const_cast<QGeoRoute*>(&route), const_cast<QGeoCoordinate*>(&position))); };
	void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString) { QByteArray tc8b6bd = errorString.toUtf8(); QtLocation_PackedString errorStringPacked = { const_cast<char*>(tc8b6bd.prepend("WHITESPACE").constData()+10), tc8b6bd.size()-10 };callbackQGeoRoutingManagerEngine_Error(this, reply, error, errorStringPacked); };
	void Signal_Finished(QGeoRouteReply * reply) { callbackQGeoRoutingManagerEngine_Finished(this, reply); };
	 ~MyQGeoRoutingManagerEngine() { callbackQGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(this); };
	bool event(QEvent * e) { return callbackQGeoRoutingManagerEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoRoutingManagerEngine_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoRoutingManagerEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoRoutingManagerEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoRoutingManagerEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoRoutingManagerEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoRoutingManagerEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoRoutingManagerEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoRoutingManagerEngine_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoRoutingManagerEngine_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGeoRoutingManagerEngine*)

int QGeoRoutingManagerEngine_QGeoRoutingManagerEngine_QRegisterMetaType(){qRegisterMetaType<QGeoRoutingManagerEngine*>(); return qRegisterMetaType<MyQGeoRoutingManagerEngine*>();}

void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void* QGeoRoutingManagerEngine_UpdateRouteDefault(void* ptr, void* route, void* position)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void* QGeoRoutingManagerEngine_NewQGeoRoutingManagerEngine(void* parameters, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QWindow*>(parent));
	} else {
		return new MyQGeoRoutingManagerEngine(*static_cast<QMap<QString, QVariant>*>(parameters), static_cast<QObject*>(parent));
	}
}

void QGeoRoutingManagerEngine_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));
}

void QGeoRoutingManagerEngine_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));
}

void QGeoRoutingManagerEngine_Error(void* ptr, void* reply, long long error, struct QtLocation_PackedString errorString)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->error(static_cast<QGeoRouteReply*>(reply), static_cast<QGeoRouteReply::Error>(error), QString::fromUtf8(errorString.data, errorString.len));
}

void QGeoRoutingManagerEngine_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));
}

void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));
}

void QGeoRoutingManagerEngine_Finished(void* ptr, void* reply)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->finished(static_cast<QGeoRouteReply*>(reply));
}

void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, long long system)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureTypes(void* ptr, long long featureTypes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureTypes(static_cast<QGeoRouteRequest::FeatureType>(featureTypes));
}

void QGeoRoutingManagerEngine_SetSupportedFeatureWeights(void* ptr, long long featureWeights)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedFeatureWeights(static_cast<QGeoRouteRequest::FeatureWeight>(featureWeights));
}

void QGeoRoutingManagerEngine_SetSupportedManeuverDetails(void* ptr, long long maneuverDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedManeuverDetails(static_cast<QGeoRouteRequest::ManeuverDetail>(maneuverDetails));
}

void QGeoRoutingManagerEngine_SetSupportedRouteOptimizations(void* ptr, long long optimizations)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedRouteOptimizations(static_cast<QGeoRouteRequest::RouteOptimization>(optimizations));
}

void QGeoRoutingManagerEngine_SetSupportedSegmentDetails(void* ptr, long long segmentDetails)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedSegmentDetails(static_cast<QGeoRouteRequest::SegmentDetail>(segmentDetails));
}

void QGeoRoutingManagerEngine_SetSupportedTravelModes(void* ptr, long long travelModes)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setSupportedTravelModes(static_cast<QGeoRouteRequest::TravelMode>(travelModes));
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr)
{
	static_cast<QGeoRoutingManagerEngine*>(ptr)->~QGeoRoutingManagerEngine();
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureTypes();
}

long long QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureWeights();
}

long long QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedManeuverDetails();
}

long long QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedRouteOptimizations();
}

long long QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedSegmentDetails();
}

long long QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManagerEngine_Locale(void* ptr)
{
	return new QLocale(static_cast<QGeoRoutingManagerEngine*>(ptr)->locale());
}

long long QGeoRoutingManagerEngine_MeasurementSystem(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->measurementSystem();
}

struct QtLocation_PackedString QGeoRoutingManagerEngine_ManagerName(void* ptr)
{
	return ({ QByteArray tce87b5 = static_cast<QGeoRoutingManagerEngine*>(ptr)->managerName().toUtf8(); QtLocation_PackedString { const_cast<char*>(tce87b5.prepend("WHITESPACE").constData()+10), tce87b5.size()-10 }; });
}

int QGeoRoutingManagerEngine_ManagerVersion(void* ptr)
{
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerVersion();
}

void* QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoRoutingManagerEngine___QGeoRoutingManagerEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedString QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoRoutingManagerEngine_____QGeoRoutingManagerEngine_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QGeoRoutingManagerEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoRoutingManagerEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoRoutingManagerEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGeoRoutingManagerEngine___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManagerEngine___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManagerEngine___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManagerEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManagerEngine___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoRoutingManagerEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoRoutingManagerEngine___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoRoutingManagerEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoRoutingManagerEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGeoRoutingManagerEngine_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::event(static_cast<QEvent*>(e));
}

char QGeoRoutingManagerEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoRoutingManagerEngine_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoRoutingManagerEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::customEvent(static_cast<QEvent*>(event));
}

void QGeoRoutingManagerEngine_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::deleteLater();
}

void QGeoRoutingManagerEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoRoutingManagerEngine_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGeoRoutingManagerEngine_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoRoutingManagerEngine*>(ptr)->QGeoRoutingManagerEngine::metaObject());
}

class MyQGeoServiceProvider: public QGeoServiceProvider
{
public:
	MyQGeoServiceProvider(const QString &providerName, const QVariantMap &parameters = QVariantMap(), bool allowExperimental = false) : QGeoServiceProvider(providerName, parameters, allowExperimental) {QGeoServiceProvider_QGeoServiceProvider_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQGeoServiceProvider_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGeoServiceProvider_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQGeoServiceProvider_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGeoServiceProvider_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGeoServiceProvider_CustomEvent(this, event); };
	void deleteLater() { callbackQGeoServiceProvider_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGeoServiceProvider_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGeoServiceProvider_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtLocation_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGeoServiceProvider_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGeoServiceProvider_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGeoServiceProvider_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQGeoServiceProvider*)

int QGeoServiceProvider_QGeoServiceProvider_QRegisterMetaType(){qRegisterMetaType<QGeoServiceProvider*>(); return qRegisterMetaType<MyQGeoServiceProvider*>();}

int QGeoServiceProvider_OnlineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OnlineGeocodingFeature;
}

int QGeoServiceProvider_OfflineGeocodingFeature_Type()
{
	return QGeoServiceProvider::OfflineGeocodingFeature;
}

int QGeoServiceProvider_ReverseGeocodingFeature_Type()
{
	return QGeoServiceProvider::ReverseGeocodingFeature;
}

int QGeoServiceProvider_LocalizedGeocodingFeature_Type()
{
	return QGeoServiceProvider::LocalizedGeocodingFeature;
}

int QGeoServiceProvider_AnyGeocodingFeatures_Type()
{
	return QGeoServiceProvider::AnyGeocodingFeatures;
}

int QGeoServiceProvider_OnlineMappingFeature_Type()
{
	return QGeoServiceProvider::OnlineMappingFeature;
}

int QGeoServiceProvider_OfflineMappingFeature_Type()
{
	return QGeoServiceProvider::OfflineMappingFeature;
}

int QGeoServiceProvider_LocalizedMappingFeature_Type()
{
	return QGeoServiceProvider::LocalizedMappingFeature;
}

int QGeoServiceProvider_AnyMappingFeatures_Type()
{
	return QGeoServiceProvider::AnyMappingFeatures;
}

int QGeoServiceProvider_OnlinePlacesFeature_Type()
{
	return QGeoServiceProvider::OnlinePlacesFeature;
}

int QGeoServiceProvider_OfflinePlacesFeature_Type()
{
	return QGeoServiceProvider::OfflinePlacesFeature;
}

int QGeoServiceProvider_SavePlaceFeature_Type()
{
	return QGeoServiceProvider::SavePlaceFeature;
}

int QGeoServiceProvider_RemovePlaceFeature_Type()
{
	return QGeoServiceProvider::RemovePlaceFeature;
}

int QGeoServiceProvider_SaveCategoryFeature_Type()
{
	return QGeoServiceProvider::SaveCategoryFeature;
}

int QGeoServiceProvider_RemoveCategoryFeature_Type()
{
	return QGeoServiceProvider::RemoveCategoryFeature;
}

int QGeoServiceProvider_PlaceRecommendationsFeature_Type()
{
	return QGeoServiceProvider::PlaceRecommendationsFeature;
}

int QGeoServiceProvider_SearchSuggestionsFeature_Type()
{
	return QGeoServiceProvider::SearchSuggestionsFeature;
}

int QGeoServiceProvider_LocalizedPlacesFeature_Type()
{
	return QGeoServiceProvider::LocalizedPlacesFeature;
}

int QGeoServiceProvider_NotificationsFeature_Type()
{
	return QGeoServiceProvider::NotificationsFeature;
}

int QGeoServiceProvider_PlaceMatchingFeature_Type()
{
	return QGeoServiceProvider::PlaceMatchingFeature;
}

int QGeoServiceProvider_AnyPlacesFeatures_Type()
{
	return QGeoServiceProvider::AnyPlacesFeatures;
}

int QGeoServiceProvider_OnlineRoutingFeature_Type()
{
	return QGeoServiceProvider::OnlineRoutingFeature;
}

int QGeoServiceProvider_OfflineRoutingFeature_Type()
{
	return QGeoServiceProvider::OfflineRoutingFeature;
}

int QGeoServiceProvider_LocalizedRoutingFeature_Type()
{
	return QGeoServiceProvider::LocalizedRoutingFeature;
}

int QGeoServiceProvider_RouteUpdatesFeature_Type()
{
	return QGeoServiceProvider::RouteUpdatesFeature;
}

int QGeoServiceProvider_AlternativeRoutesFeature_Type()
{
	return QGeoServiceProvider::AlternativeRoutesFeature;
}

int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type()
{
	return QGeoServiceProvider::ExcludeAreasRoutingFeature;
}

int QGeoServiceProvider_AnyRoutingFeatures_Type()
{
	return QGeoServiceProvider::AnyRoutingFeatures;
}

void* QGeoServiceProvider_NewQGeoServiceProvider(struct QtLocation_PackedString providerName, void* parameters, char allowExperimental)
{
	return new MyQGeoServiceProvider(QString::fromUtf8(providerName.data, providerName.len), *static_cast<QMap<QString, QVariant>*>(parameters), allowExperimental != 0);
}

struct QtLocation_PackedString QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders()
{
	return ({ QByteArray teec3ca = QGeoServiceProvider::availableServiceProviders().join("|").toUtf8(); QtLocation_PackedString { const_cast<char*>(teec3ca.prepend("WHITESPACE").constData()+10), teec3ca.size()-10 }; });
}

void QGeoServiceProvider_SetAllowExperimental(void* ptr, char allow)
{
	static_cast<QGeoServiceProvider*>(ptr)->setAllowExperimental(allow != 0);
}

void QGeoServiceProvider_SetLocale(void* ptr, void* locale)
{
	static_cast<QGeoServiceProvider*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoServiceProvider_SetParameters(void* ptr, void* parameters)
{
	static_cast<QGeoServiceProvider*>(ptr)->setParameters(*static_cast<QMap<QString, QVariant>*>(parameters));
}

void QGeoServiceProvider_DestroyQGeoServiceProvider(void* ptr)
{
	static_cast<QGeoServiceProvider*>(ptr)->~QGeoServiceProvider();
}

long long QGeoServiceProvider_Error(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->error();
}

long long QGeoServiceProvider_GeocodingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingFeatures();
}

long long QGeoServiceProvider_MappingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->mappingFeatures();
}

long long QGeoServiceProvider_PlacesFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placesFeatures();
}

void* QGeoServiceProvider_GeocodingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingManager();
}

void* QGeoServiceProvider_RoutingManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingManager();
}

void* QGeoServiceProvider_PlaceManager(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->placeManager();
}

struct QtLocation_PackedString QGeoServiceProvider_ErrorString(void* ptr)
{
	return ({ QByteArray t90b712 = static_cast<QGeoServiceProvider*>(ptr)->errorString().toUtf8(); QtLocation_PackedString { const_cast<char*>(t90b712.prepend("WHITESPACE").constData()+10), t90b712.size()-10 }; });
}

long long QGeoServiceProvider_RoutingFeatures(void* ptr)
{
	return static_cast<QGeoServiceProvider*>(ptr)->routingFeatures();
}

void* QGeoServiceProvider___QGeoServiceProvider_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProvider___QGeoServiceProvider_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProvider___QGeoServiceProvider_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProvider___QGeoServiceProvider_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoServiceProvider___setParameters_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProvider___setParameters_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProvider___setParameters_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProvider___setParameters_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedString QGeoServiceProvider_____QGeoServiceProvider_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProvider_____QGeoServiceProvider_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProvider_____QGeoServiceProvider_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtLocation_PackedString QGeoServiceProvider_____setParameters_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProvider_____setParameters_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProvider_____setParameters_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

void* QGeoServiceProvider___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QGeoServiceProvider___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGeoServiceProvider___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>;
}

void* QGeoServiceProvider___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoServiceProvider___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoServiceProvider___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoServiceProvider___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoServiceProvider___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QGeoServiceProvider___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>;
}

void* QGeoServiceProvider___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QGeoServiceProvider___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QGeoServiceProvider___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>;
}

char QGeoServiceProvider_EventDefault(void* ptr, void* e)
{
		return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::event(static_cast<QEvent*>(e));
}

char QGeoServiceProvider_EventFilterDefault(void* ptr, void* watched, void* event)
{
		return static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QGeoServiceProvider_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::childEvent(static_cast<QChildEvent*>(event));
}

void QGeoServiceProvider_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::customEvent(static_cast<QEvent*>(event));
}

void QGeoServiceProvider_DeleteLaterDefault(void* ptr)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::deleteLater();
}

void QGeoServiceProvider_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGeoServiceProvider_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QGeoServiceProvider_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QGeoServiceProvider*>(ptr)->QGeoServiceProvider::metaObject());
}

class MyQGeoServiceProviderFactory: public QGeoServiceProviderFactory
{
public:
	 ~MyQGeoServiceProviderFactory() { callbackQGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(this); };
};

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactory(void* ptr)
{
	static_cast<QGeoServiceProviderFactory*>(ptr)->~QGeoServiceProviderFactory();
}

void QGeoServiceProviderFactory_DestroyQGeoServiceProviderFactoryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createGeocodingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createGeocodingManagerEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoServiceProviderFactory___createMappingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProviderFactory___createMappingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createMappingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createMappingManagerEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createRoutingManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createRoutingManagerEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

void* QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_atList(void* ptr, struct QtLocation_PackedString i)
{
	return new QVariant(static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(i.data, i.len)));
}

void QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_setList(void* ptr, struct QtLocation_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QGeoServiceProviderFactory___createPlaceManagerEngine_parameters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>;
}

struct QtLocation_PackedList QGeoServiceProviderFactory___createPlaceManagerEngine_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtLocation_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createGeocodingManagerEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProviderFactory_____createGeocodingManagerEngine_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createGeocodingManagerEngine_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createMappingManagerEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProviderFactory_____createMappingManagerEngine_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createMappingManagerEngine_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createRoutingManagerEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProviderFactory_____createRoutingManagerEngine_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createRoutingManagerEngine_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

struct QtLocation_PackedString QGeoServiceProviderFactory_____createPlaceManagerEngine_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t29def6 = static_cast<QList<QString>*>(ptr)->at(i).toUtf8(); QtLocation_PackedString { const_cast<char*>(t29def6.prepend("WHITESPACE").constData()+10), t29def6.size()-10 }; });
}

void QGeoServiceProviderFactory_____createPlaceManagerEngine_keyList_setList(void* ptr, struct QtLocation_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QGeoServiceProviderFactory_____createPlaceManagerEngine_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>;
}

