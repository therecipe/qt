#include "qtapandholdgesture.h"
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTapAndHoldGesture>
#include "_cgo_export.h"

class MyQTapAndHoldGesture: public QTapAndHoldGesture {
public:
};

void QTapAndHoldGesture_SetPosition(void* ptr, void* pos){
	static_cast<QTapAndHoldGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(int msecs){
	QTapAndHoldGesture::setTimeout(msecs);
}

int QTapAndHoldGesture_QTapAndHoldGesture_Timeout(){
	return QTapAndHoldGesture::timeout();
}

void QTapAndHoldGesture_DestroyQTapAndHoldGesture(void* ptr){
	static_cast<QTapAndHoldGesture*>(ptr)->~QTapAndHoldGesture();
}

