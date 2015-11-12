#include "qswipegesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSwipeGesture>
#include "_cgo_export.h"

class MyQSwipeGesture: public QSwipeGesture {
public:
};

int QSwipeGesture_HorizontalDirection(void* ptr){
	return static_cast<QSwipeGesture*>(ptr)->horizontalDirection();
}

void QSwipeGesture_SetSwipeAngle(void* ptr, double value){
	static_cast<QSwipeGesture*>(ptr)->setSwipeAngle(static_cast<qreal>(value));
}

double QSwipeGesture_SwipeAngle(void* ptr){
	return static_cast<double>(static_cast<QSwipeGesture*>(ptr)->swipeAngle());
}

int QSwipeGesture_VerticalDirection(void* ptr){
	return static_cast<QSwipeGesture*>(ptr)->verticalDirection();
}

void QSwipeGesture_DestroyQSwipeGesture(void* ptr){
	static_cast<QSwipeGesture*>(ptr)->~QSwipeGesture();
}

