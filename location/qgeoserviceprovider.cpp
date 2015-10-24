#include "qgeoserviceprovider.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLocale>
#include <QGeoServiceProvider>
#include "_cgo_export.h"

class MyQGeoServiceProvider: public QGeoServiceProvider {
public:
};

int QGeoServiceProvider_OnlineGeocodingFeature_Type(){
	return QGeoServiceProvider::OnlineGeocodingFeature;
}

int QGeoServiceProvider_OfflineGeocodingFeature_Type(){
	return QGeoServiceProvider::OfflineGeocodingFeature;
}

int QGeoServiceProvider_ReverseGeocodingFeature_Type(){
	return QGeoServiceProvider::ReverseGeocodingFeature;
}

int QGeoServiceProvider_LocalizedGeocodingFeature_Type(){
	return QGeoServiceProvider::LocalizedGeocodingFeature;
}

int QGeoServiceProvider_AnyGeocodingFeatures_Type(){
	return QGeoServiceProvider::AnyGeocodingFeatures;
}

int QGeoServiceProvider_OnlineMappingFeature_Type(){
	return QGeoServiceProvider::OnlineMappingFeature;
}

int QGeoServiceProvider_OfflineMappingFeature_Type(){
	return QGeoServiceProvider::OfflineMappingFeature;
}

int QGeoServiceProvider_LocalizedMappingFeature_Type(){
	return QGeoServiceProvider::LocalizedMappingFeature;
}

int QGeoServiceProvider_AnyMappingFeatures_Type(){
	return QGeoServiceProvider::AnyMappingFeatures;
}

int QGeoServiceProvider_OnlinePlacesFeature_Type(){
	return QGeoServiceProvider::OnlinePlacesFeature;
}

int QGeoServiceProvider_OfflinePlacesFeature_Type(){
	return QGeoServiceProvider::OfflinePlacesFeature;
}

int QGeoServiceProvider_SavePlaceFeature_Type(){
	return QGeoServiceProvider::SavePlaceFeature;
}

int QGeoServiceProvider_RemovePlaceFeature_Type(){
	return QGeoServiceProvider::RemovePlaceFeature;
}

int QGeoServiceProvider_SaveCategoryFeature_Type(){
	return QGeoServiceProvider::SaveCategoryFeature;
}

int QGeoServiceProvider_RemoveCategoryFeature_Type(){
	return QGeoServiceProvider::RemoveCategoryFeature;
}

int QGeoServiceProvider_PlaceRecommendationsFeature_Type(){
	return QGeoServiceProvider::PlaceRecommendationsFeature;
}

int QGeoServiceProvider_SearchSuggestionsFeature_Type(){
	return QGeoServiceProvider::SearchSuggestionsFeature;
}

int QGeoServiceProvider_LocalizedPlacesFeature_Type(){
	return QGeoServiceProvider::LocalizedPlacesFeature;
}

int QGeoServiceProvider_NotificationsFeature_Type(){
	return QGeoServiceProvider::NotificationsFeature;
}

int QGeoServiceProvider_PlaceMatchingFeature_Type(){
	return QGeoServiceProvider::PlaceMatchingFeature;
}

int QGeoServiceProvider_AnyPlacesFeatures_Type(){
	return QGeoServiceProvider::AnyPlacesFeatures;
}

int QGeoServiceProvider_OnlineRoutingFeature_Type(){
	return QGeoServiceProvider::OnlineRoutingFeature;
}

int QGeoServiceProvider_OfflineRoutingFeature_Type(){
	return QGeoServiceProvider::OfflineRoutingFeature;
}

int QGeoServiceProvider_LocalizedRoutingFeature_Type(){
	return QGeoServiceProvider::LocalizedRoutingFeature;
}

int QGeoServiceProvider_RouteUpdatesFeature_Type(){
	return QGeoServiceProvider::RouteUpdatesFeature;
}

int QGeoServiceProvider_AlternativeRoutesFeature_Type(){
	return QGeoServiceProvider::AlternativeRoutesFeature;
}

int QGeoServiceProvider_ExcludeAreasRoutingFeature_Type(){
	return QGeoServiceProvider::ExcludeAreasRoutingFeature;
}

int QGeoServiceProvider_AnyRoutingFeatures_Type(){
	return QGeoServiceProvider::AnyRoutingFeatures;
}

char* QGeoServiceProvider_QGeoServiceProvider_AvailableServiceProviders(){
	return QGeoServiceProvider::availableServiceProviders().join("|").toUtf8().data();
}

int QGeoServiceProvider_Error(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->error();
}

char* QGeoServiceProvider_ErrorString(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->errorString().toUtf8().data();
}

int QGeoServiceProvider_GeocodingFeatures(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingFeatures();
}

QtObjectPtr QGeoServiceProvider_GeocodingManager(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->geocodingManager();
}

int QGeoServiceProvider_MappingFeatures(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->mappingFeatures();
}

QtObjectPtr QGeoServiceProvider_PlaceManager(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->placeManager();
}

int QGeoServiceProvider_PlacesFeatures(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->placesFeatures();
}

int QGeoServiceProvider_RoutingFeatures(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->routingFeatures();
}

QtObjectPtr QGeoServiceProvider_RoutingManager(QtObjectPtr ptr){
	return static_cast<QGeoServiceProvider*>(ptr)->routingManager();
}

void QGeoServiceProvider_SetAllowExperimental(QtObjectPtr ptr, int allow){
	static_cast<QGeoServiceProvider*>(ptr)->setAllowExperimental(allow != 0);
}

void QGeoServiceProvider_SetLocale(QtObjectPtr ptr, QtObjectPtr locale){
	static_cast<QGeoServiceProvider*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}

void QGeoServiceProvider_DestroyQGeoServiceProvider(QtObjectPtr ptr){
	static_cast<QGeoServiceProvider*>(ptr)->~QGeoServiceProvider();
}

