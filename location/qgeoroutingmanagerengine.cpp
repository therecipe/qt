#include "qgeoroutingmanagerengine.h"
#include <QModelIndex>
#include <QGeoRouteRequest>
#include <QGeoRouteReply>
#include <QLocale>
#include <QObject>
#include <QGeoRoutingManager>
#include <QGeoCoordinate>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGeoRoute>
#include <QGeoRoutingManagerEngine>
#include "_cgo_export.h"

class MyQGeoRoutingManagerEngine: public QGeoRoutingManagerEngine {
public:
void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString){callbackQGeoRoutingManagerEngineError(this->objectName().toUtf8().data(), reply, error, errorString.toUtf8().data());};
void Signal_Finished(QGeoRouteReply * reply){callbackQGeoRoutingManagerEngineFinished(this->objectName().toUtf8().data(), reply);};
};

QtObjectPtr QGeoRoutingManagerEngine_CalculateRoute(QtObjectPtr ptr, QtObjectPtr request){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManagerEngine_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));;
}

void QGeoRoutingManagerEngine_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManagerEngine::error), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManagerEngine::Signal_Error));;
}

void QGeoRoutingManagerEngine_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));;
}

void QGeoRoutingManagerEngine_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoRoutingManagerEngine*>(ptr), static_cast<void (QGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&QGeoRoutingManagerEngine::finished), static_cast<MyQGeoRoutingManagerEngine*>(ptr), static_cast<void (MyQGeoRoutingManagerEngine::*)(QGeoRouteReply *)>(&MyQGeoRoutingManagerEngine::Signal_Finished));;
}

char* QGeoRoutingManagerEngine_ManagerName(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerName().toUtf8().data();
}

int QGeoRoutingManagerEngine_ManagerVersion(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->managerVersion();
}

int QGeoRoutingManagerEngine_MeasurementSystem(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->measurementSystem();
}

void QGeoRoutingManagerEngine_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManagerEngine_SetMeasurementSystem(QtObjectPtr ptr, int system){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

int QGeoRoutingManagerEngine_SupportedFeatureTypes(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureTypes();
}

int QGeoRoutingManagerEngine_SupportedFeatureWeights(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedFeatureWeights();
}

int QGeoRoutingManagerEngine_SupportedManeuverDetails(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedManeuverDetails();
}

int QGeoRoutingManagerEngine_SupportedRouteOptimizations(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedRouteOptimizations();
}

int QGeoRoutingManagerEngine_SupportedSegmentDetails(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedSegmentDetails();
}

int QGeoRoutingManagerEngine_SupportedTravelModes(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->supportedTravelModes();
}

QtObjectPtr QGeoRoutingManagerEngine_UpdateRoute(QtObjectPtr ptr, QtObjectPtr route, QtObjectPtr position){
	return static_cast<QGeoRoutingManagerEngine*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManagerEngine_DestroyQGeoRoutingManagerEngine(QtObjectPtr ptr){
	static_cast<QGeoRoutingManagerEngine*>(ptr)->~QGeoRoutingManagerEngine();
}

