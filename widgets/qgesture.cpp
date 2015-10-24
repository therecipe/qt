#include "qgesture.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QPoint>
#include <QPointF>
#include <QString>
#include <QGesture>
#include "_cgo_export.h"

class MyQGesture: public QGesture {
public:
};

int QGesture_GestureCancelPolicy(QtObjectPtr ptr){
	return static_cast<QGesture*>(ptr)->gestureCancelPolicy();
}

int QGesture_GestureType(QtObjectPtr ptr){
	return static_cast<QGesture*>(ptr)->gestureType();
}

int QGesture_HasHotSpot(QtObjectPtr ptr){
	return static_cast<QGesture*>(ptr)->hasHotSpot();
}

void QGesture_SetGestureCancelPolicy(QtObjectPtr ptr, int policy){
	static_cast<QGesture*>(ptr)->setGestureCancelPolicy(static_cast<QGesture::GestureCancelPolicy>(policy));
}

void QGesture_SetHotSpot(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QGesture*>(ptr)->setHotSpot(*static_cast<QPointF*>(value));
}

void QGesture_UnsetHotSpot(QtObjectPtr ptr){
	static_cast<QGesture*>(ptr)->unsetHotSpot();
}

QtObjectPtr QGesture_NewQGesture(QtObjectPtr parent){
	return new QGesture(static_cast<QObject*>(parent));
}

void QGesture_DestroyQGesture(QtObjectPtr ptr){
	static_cast<QGesture*>(ptr)->~QGesture();
}

