#include "qtouchevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTouchEvent>
#include "_cgo_export.h"

class MyQTouchEvent: public QTouchEvent {
public:
};

QtObjectPtr QTouchEvent_Device(QtObjectPtr ptr){
	return static_cast<QTouchEvent*>(ptr)->device();
}

QtObjectPtr QTouchEvent_Target(QtObjectPtr ptr){
	return static_cast<QTouchEvent*>(ptr)->target();
}

int QTouchEvent_TouchPointStates(QtObjectPtr ptr){
	return static_cast<QTouchEvent*>(ptr)->touchPointStates();
}

QtObjectPtr QTouchEvent_Window(QtObjectPtr ptr){
	return static_cast<QTouchEvent*>(ptr)->window();
}

void QTouchEvent_DestroyQTouchEvent(QtObjectPtr ptr){
	static_cast<QTouchEvent*>(ptr)->~QTouchEvent();
}

