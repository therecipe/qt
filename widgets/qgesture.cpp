#include "qgesture.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QObject>
#include <QPointF>
#include <QString>
#include <QGesture>
#include "_cgo_export.h"

class MyQGesture: public QGesture {
public:
};

int QGesture_GestureCancelPolicy(void* ptr){
	return static_cast<QGesture*>(ptr)->gestureCancelPolicy();
}

int QGesture_GestureType(void* ptr){
	return static_cast<QGesture*>(ptr)->gestureType();
}

int QGesture_HasHotSpot(void* ptr){
	return static_cast<QGesture*>(ptr)->hasHotSpot();
}

void QGesture_SetGestureCancelPolicy(void* ptr, int policy){
	static_cast<QGesture*>(ptr)->setGestureCancelPolicy(static_cast<QGesture::GestureCancelPolicy>(policy));
}

void QGesture_SetHotSpot(void* ptr, void* value){
	static_cast<QGesture*>(ptr)->setHotSpot(*static_cast<QPointF*>(value));
}

void QGesture_UnsetHotSpot(void* ptr){
	static_cast<QGesture*>(ptr)->unsetHotSpot();
}

void* QGesture_NewQGesture(void* parent){
	return new QGesture(static_cast<QObject*>(parent));
}

void QGesture_DestroyQGesture(void* ptr){
	static_cast<QGesture*>(ptr)->~QGesture();
}

