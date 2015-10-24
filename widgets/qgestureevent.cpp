#include "qgestureevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGesture>
#include <QGestureEvent>
#include "_cgo_export.h"

class MyQGestureEvent: public QGestureEvent {
public:
};

void QGestureEvent_Accept(QtObjectPtr ptr, QtObjectPtr gesture){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Accept2(QtObjectPtr ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<Qt::GestureType>(gestureType));
}

QtObjectPtr QGestureEvent_Gesture(QtObjectPtr ptr, int ty){
	return static_cast<QGestureEvent*>(ptr)->gesture(static_cast<Qt::GestureType>(ty));
}

void QGestureEvent_Ignore(QtObjectPtr ptr, QtObjectPtr gesture){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Ignore2(QtObjectPtr ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<Qt::GestureType>(gestureType));
}

int QGestureEvent_IsAccepted(QtObjectPtr ptr, QtObjectPtr gesture){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<QGesture*>(gesture));
}

int QGestureEvent_IsAccepted2(QtObjectPtr ptr, int gestureType){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<Qt::GestureType>(gestureType));
}

void QGestureEvent_SetAccepted(QtObjectPtr ptr, QtObjectPtr gesture, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<QGesture*>(gesture), value != 0);
}

void QGestureEvent_SetAccepted2(QtObjectPtr ptr, int gestureType, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<Qt::GestureType>(gestureType), value != 0);
}

QtObjectPtr QGestureEvent_Widget(QtObjectPtr ptr){
	return static_cast<QGestureEvent*>(ptr)->widget();
}

void QGestureEvent_DestroyQGestureEvent(QtObjectPtr ptr){
	static_cast<QGestureEvent*>(ptr)->~QGestureEvent();
}

