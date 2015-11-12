#include "qgeoroutingmanagerengine.h"
#include <QGeoRoute>
#include <QGeoRoutingManager>
#include <QVariant>
#include <QModelIndex>
#include <QGeoCoordinate>
#include <QGeoRouteRequest>
#include <QObject>
#include <QString>
#include <QUrl>
#include <QLocale>
#include <QGeoRouteReply>
#include <QGeoRoutingManagerEngine>
#include "_cgo_export.h"

class MyQGeoRoutingManagerEngine: public QGeoRoutingManagerEngine {
public:
void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString){callbackQGeoRoutingManagerEngineError(this->objectName().toUtf8().data(), reply, error, errorString.toUtf8().data());};
void Signal_Finished(QGeoRouteReply * reply){callbackQGeoRoutingManagerEngineFinished(this->objectName().toUtf8().data(), reply);};
};

void* QGeoRoutingManagerEngine_CalculateRoute(void* ptr, void* request){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManagerEngine_ConnectError(void* ptr){
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));;
}

void QGeoRoutingManagerEngine_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));;
}

void QGeoRoutingManagerEngine_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));;
}

void QGeoRoutingManagerEngine_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));;
}

char* QGeoRoutingManagerEngine_ManagerName(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerName().toUtf8().data();
}

int QGeoRoutingManagerEngine_ManagerVersion(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerVersion();
}

int QGeoRoutingManagerEngine_MeasurementSystem(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->measurementSystem();
}

void QGeoRoutingManagerEngine_SetLocale(void* ptr, void* locale){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManagerEngine_SetMeasurementSystem(void* ptr, int system){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

int QGeoRoutingManagerEngine_SupportedFeatureTypes(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureTypes();
}

int QGeoRoutingManagerEngine_SupportedFeatureWeights(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureWeights();
}

int QGeoRoutingManagerEngine_SupportedManeuverDetails(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedManeuverDetails();
}

int QGeoRoutingManagerEngine_SupportedRouteOptimizations(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedRouteOptimizations();
}

int QGeoRoutingManagerEngine_SupportedSegmentDetails(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedSegmentDetails();
}

int QGeoRoutingManagerEngine_SupportedTravelModes(void* ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManagerEngine_UpdateRoute(void* ptr, void* route, void* position){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(void* ptr){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->~QGeoRoutingManagerEngine();
}

