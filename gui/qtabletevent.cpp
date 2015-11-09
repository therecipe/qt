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

int QTabletEvent_Button(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->button();
}

int QTabletEvent_Buttons(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->buttons();
}

int QTabletEvent_Device(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->device();
}

int QTabletEvent_GlobalX(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->globalX();
}

int QTabletEvent_GlobalY(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->globalY();
}

double QTabletEvent_HiResGlobalX(void* ptr){
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->hiResGlobalX());
}

double QTabletEvent_HiResGlobalY(void* ptr){
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->hiResGlobalY());
}

int QTabletEvent_PointerType(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->pointerType();
}

double QTabletEvent_Pressure(void* ptr){
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->pressure());
}

double QTabletEvent_Rotation(void* ptr){
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->rotation());
}

double QTabletEvent_TangentialPressure(void* ptr){
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->tangentialPressure());
}

int QTabletEvent_X(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->x();
}

int QTabletEvent_XTilt(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->xTilt();
}

int QTabletEvent_Y(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->y();
}

int QTabletEvent_YTilt(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->yTilt();
}

int QTabletEvent_Z(void* ptr){
	return static_cast<QTabletEvent*>(ptr)->z();
}

