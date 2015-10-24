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

int QSwipeGesture_HorizontalDirection(QtObjectPtr ptr){
	return static_cast<QSwipeGesture*>(ptr)->horizontalDirection();
}

int QSwipeGesture_VerticalDirection(QtObjectPtr ptr){
	return static_cast<QSwipeGesture*>(ptr)->verticalDirection();
}

void QSwipeGesture_DestroyQSwipeGesture(QtObjectPtr ptr){
	static_cast<QSwipeGesture*>(ptr)->~QSwipeGesture();
}

