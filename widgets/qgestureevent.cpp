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

void QGestureEvent_Accept(void* ptr, void* gesture){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Accept2(void* ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->accept(static_cast<Qt::GestureType>(gestureType));
}

void* QGestureEvent_Gesture(void* ptr, int ty){
	return static_cast<QGestureEvent*>(ptr)->gesture(static_cast<Qt::GestureType>(ty));
}

void QGestureEvent_Ignore(void* ptr, void* gesture){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<QGesture*>(gesture));
}

void QGestureEvent_Ignore2(void* ptr, int gestureType){
	static_cast<QGestureEvent*>(ptr)->ignore(static_cast<Qt::GestureType>(gestureType));
}

int QGestureEvent_IsAccepted(void* ptr, void* gesture){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<QGesture*>(gesture));
}

int QGestureEvent_IsAccepted2(void* ptr, int gestureType){
	return static_cast<QGestureEvent*>(ptr)->isAccepted(static_cast<Qt::GestureType>(gestureType));
}

void QGestureEvent_SetAccepted(void* ptr, void* gesture, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<QGesture*>(gesture), value != 0);
}

void QGestureEvent_SetAccepted2(void* ptr, int gestureType, int value){
	static_cast<QGestureEvent*>(ptr)->setAccepted(static_cast<Qt::GestureType>(gestureType), value != 0);
}

void* QGestureEvent_Widget(void* ptr){
	return static_cast<QGestureEvent*>(ptr)->widget();
}

void QGestureEvent_DestroyQGestureEvent(void* ptr){
	static_cast<QGestureEvent*>(ptr)->~QGestureEvent();
}

