#include "qtabletevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTabletEvent>
#include "_cgo_export.h"

class MyQTabletEvent: public QTabletEvent {
public:
};

int QTabletEvent_Button(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->button();
}

int QTabletEvent_Buttons(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->buttons();
}

int QTabletEvent_Device(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->device();
}

int QTabletEvent_GlobalX(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->globalX();
}

int QTabletEvent_GlobalY(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->globalY();
}

int QTabletEvent_PointerType(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->pointerType();
}

int QTabletEvent_X(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->x();
}

int QTabletEvent_XTilt(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->xTilt();
}

int QTabletEvent_Y(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->y();
}

int QTabletEvent_YTilt(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->yTilt();
}

int QTabletEvent_Z(QtObjectPtr ptr){
	return static_cast<QTabletEvent*>(ptr)->z();
}

