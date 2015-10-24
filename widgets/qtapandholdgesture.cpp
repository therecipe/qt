#include "qtapandholdgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPointF>
#include <QTapAndHoldGesture>
#include "_cgo_export.h"

class MyQTapAndHoldGesture: public QTapAndHoldGesture {
public:
};

void QTapAndHoldGesture_SetPosition(QtObjectPtr ptr, QtObjectPtr pos){
	static_cast<QTapAndHoldGesture*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

void QTapAndHoldGesture_QTapAndHoldGesture_SetTimeout(int msecs){
	QTapAndHoldGesture::setTimeout(msecs);
}

int QTapAndHoldGesture_QTapAndHoldGesture_Timeout(){
	return QTapAndHoldGesture::timeout();
}

void QTapAndHoldGesture_DestroyQTapAndHoldGesture(QtObjectPtr ptr){
	static_cast<QTapAndHoldGesture*>(ptr)->~QTapAndHoldGesture();
}

