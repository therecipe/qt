#include "qgeoroutingmanager.h"
#include <QString>
#include <QUrl>
#include <QGeoRouteReply>
#include <QGeoRoute>
#include <QObject>
#include <QVariant>
#include <QModelIndex>
#include <QLocale>
#include <QGeoCoordinate>
#include <QGeoRouteRequest>
#include <QGeoRoutingManager>
#include "_cgo_export.h"

class MyQGeoRoutingManager: public QGeoRoutingManager {
public:
void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString){callbackQGeoRoutingManagerError(this->objectName().toUtf8().data(), reply, error, errorString.toUtf8().data());};
void Signal_Finished(QGeoRouteReply * reply){callbackQGeoRoutingManagerFinished(this->objectName().toUtf8().data(), reply);};
};

void* QGeoRoutingManager_CalculateRoute(void* ptr, void* request){
	return static_cast<QGeoRoutingManager*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManager_ConnectError(void* ptr){
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));;
}

void QGeoRoutingManager_DisconnectError(void* ptr){
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));;
}

void QGeoRoutingManager_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));;
}

void QGeoRoutingManager_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));;
}

char* QGeoRoutingManager_ManagerName(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->managerName().toUtf8().data();
}

int QGeoRoutingManager_ManagerVersion(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->managerVersion();
}

int QGeoRoutingManager_MeasurementSystem(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->measurementSystem();
}

void QGeoRoutingManager_SetLocale(void* ptr, void* locale){
	static_cast<QGeoRoutingManager*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManager_SetMeasurementSystem(void* ptr, int system){
	static_cast<QGeoRoutingManager*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

int QGeoRoutingManager_SupportedFeatureTypes(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureTypes();
}

int QGeoRoutingManager_SupportedFeatureWeights(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureWeights();
}

int QGeoRoutingManager_SupportedManeuverDetails(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedManeuverDetails();
}

int QGeoRoutingManager_SupportedRouteOptimizations(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedRouteOptimizations();
}

int QGeoRoutingManager_SupportedSegmentDetails(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedSegmentDetails();
}

int QGeoRoutingManager_SupportedTravelModes(void* ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedTravelModes();
}

void* QGeoRoutingManager_UpdateRoute(void* ptr, void* route, void* position){
	return static_cast<QGeoRoutingManager*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManager_DestroyQGeoRoutingManager(void* ptr){
	static_cast<QGeoRoutingManager*>(ptr)->~QGeoRoutingManager();
}

