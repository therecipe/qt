#include "qgeoroutingmanager.h"
#include <QObject>
#include <QGeoRouteReply>
#include <QVariant>
#include <QModelIndex>
#include <QGeoRouteRequest>
#include <QLocale>
#include <QGeoCoordinate>
#include <QString>
#include <QUrl>
#include <QGeoRoute>
#include <QGeoRoutingManager>
#include "_cgo_export.h"

class MyQGeoRoutingManager: public QGeoRoutingManager {
public:
void Signal_Error(QGeoRouteReply * reply, QGeoRouteReply::Error error, QString errorString){callbackQGeoRoutingManagerError(this->objectName().toUtf8().data(), reply, error, errorString.toUtf8().data());};
void Signal_Finished(QGeoRouteReply * reply){callbackQGeoRoutingManagerFinished(this->objectName().toUtf8().data(), reply);};
};

QtObjectPtr QGeoRoutingManager_CalculateRoute(QtObjectPtr ptr, QtObjectPtr request){
	return static_cast<QGeoRoutingManager*>(ptr)->calculateRoute(*static_cast<QGeoRouteRequest*>(request));
}

void QGeoRoutingManager_ConnectError(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));;
}

void QGeoRoutingManager_DisconnectError(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&QGeoRoutingManager::error), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *, QGeoRouteReply::Error, QString)>(&MyQGeoRoutingManager::Signal_Error));;
}

void QGeoRoutingManager_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));;
}

void QGeoRoutingManager_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QGeoRoutingManager*>(ptr), static_cast<void (QGeoRoutingManager::*)(QGeoRouteReply *)>(&QGeoRoutingManager::finished), static_cast<MyQGeoRoutingManager*>(ptr), static_cast<void (MyQGeoRoutingManager::*)(QGeoRouteReply *)>(&MyQGeoRoutingManager::Signal_Finished));;
}

char* QGeoRoutingManager_ManagerName(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->managerName().toUtf8().data();
}

int QGeoRoutingManager_ManagerVersion(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->managerVersion();
}

int QGeoRoutingManager_MeasurementSystem(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->measurementSystem();
}

void QGeoRoutingManager_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QGeoRoutingManager*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoRoutingManager_SetMeasurementSystem(QtObjectPtr ptr, int system){
	static_cast<QGeoRoutingManager*>(ptr)->setMeasurementSystem(static_cast<QLocale::MeasurementSystem>(system));
}

int QGeoRoutingManager_SupportedFeatureTypes(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureTypes();
}

int QGeoRoutingManager_SupportedFeatureWeights(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedFeatureWeights();
}

int QGeoRoutingManager_SupportedManeuverDetails(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedManeuverDetails();
}

int QGeoRoutingManager_SupportedRouteOptimizations(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedRouteOptimizations();
}

int QGeoRoutingManager_SupportedSegmentDetails(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedSegmentDetails();
}

int QGeoRoutingManager_SupportedTravelModes(QtObjectPtr ptr){
	return static_cast<QGeoRoutingManager*>(ptr)->supportedTravelModes();
}

QtObjectPtr QGeoRoutingManager_UpdateRoute(QtObjectPtr ptr, QtObjectPtr route, QtObjectPtr position){
	return static_cast<QGeoRoutingManager*>(ptr)->updateRoute(*static_cast<QGeoRoute*>(route), *static_cast<QGeoCoordinate*>(position));
}

void QGeoRoutingManager_DestroyQGeoRoutingManager(QtObjectPtr ptr){
	static_cast<QGeoRoutingManager*>(ptr)->~QGeoRoutingManager();
}

