#include "qtouchevent.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QTouchEvent>
#include "_cgo_export.h"

class MyQTouchEvent: public QTouchEvent {
public:
};

void* QTouchEvent_Device(void* ptr){
	return static_cast<QTouchEvent*>(ptr)->device();
}

void* QTouchEvent_Target(void* ptr){
	return static_cast<QTouchEvent*>(ptr)->target();
}

int QTouchEvent_TouchPointStates(void* ptr){
	return static_cast<QTouchEvent*>(ptr)->touchPointStates();
}

void* QTouchEvent_Window(void* ptr){
	return static_cast<QTouchEvent*>(ptr)->window();
}

void QTouchEvent_DestroyQTouchEvent(void* ptr){
	static_cast<QTouchEvent*>(ptr)->~QTouchEvent();
}

