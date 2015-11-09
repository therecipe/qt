#include "qnativegestureevent.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QNativeGestureEvent>
#include "_cgo_export.h"

class MyQNativeGestureEvent: public QNativeGestureEvent {
public:
};

int QNativeGestureEvent_GestureType(void* ptr){
	return static_cast<QNativeGestureEvent*>(ptr)->gestureType();
}

double QNativeGestureEvent_Value(void* ptr){
	return static_cast<double>(static_cast<QNativeGestureEvent*>(ptr)->value());
}

