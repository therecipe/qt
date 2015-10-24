#include "qplatformsurfaceevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPlatformSurfaceEvent>
#include "_cgo_export.h"

class MyQPlatformSurfaceEvent: public QPlatformSurfaceEvent {
public:
};

QtObjectPtr QPlatformSurfaceEvent_NewQPlatformSurfaceEvent(int surfaceEventType){
	return new QPlatformSurfaceEvent(static_cast<QPlatformSurfaceEvent::SurfaceEventType>(surfaceEventType));
}

int QPlatformSurfaceEvent_SurfaceEventType(QtObjectPtr ptr){
	return static_cast<QPlatformSurfaceEvent*>(ptr)->surfaceEventType();
}

