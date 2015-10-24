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

int QNativeGestureEvent_GestureType(QtObjectPtr ptr){
	return static_cast<QNativeGestureEvent*>(ptr)->gestureType();
}

