#include "qcamerazoomcontrol.h"
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QCameraZoomControl>
#include "_cgo_export.h"

class MyQCameraZoomControl: public QCameraZoomControl {
public:
};

double QCameraZoomControl_CurrentDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->currentDigitalZoom());
}

double QCameraZoomControl_CurrentOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->currentOpticalZoom());
}

double QCameraZoomControl_MaximumDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->maximumDigitalZoom());
}

double QCameraZoomControl_MaximumOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->maximumOpticalZoom());
}

double QCameraZoomControl_RequestedDigitalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->requestedDigitalZoom());
}

double QCameraZoomControl_RequestedOpticalZoom(void* ptr){
	return static_cast<double>(static_cast<QCameraZoomControl*>(ptr)->requestedOpticalZoom());
}

void QCameraZoomControl_ZoomTo(void* ptr, double optical, double digital){
	static_cast<QCameraZoomControl*>(ptr)->zoomTo(static_cast<qreal>(optical), static_cast<qreal>(digital));
}

void QCameraZoomControl_DestroyQCameraZoomControl(void* ptr){
	static_cast<QCameraZoomControl*>(ptr)->~QCameraZoomControl();
}

