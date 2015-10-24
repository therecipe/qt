#include "qcamerazoomcontrol.h"
#include <QCamera>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCameraZoomControl>
#include "_cgo_export.h"

class MyQCameraZoomControl: public QCameraZoomControl {
public:
};

void QCameraZoomControl_DestroyQCameraZoomControl(QtObjectPtr ptr){
	static_cast<QCameraZoomControl*>(ptr)->~QCameraZoomControl();
}

