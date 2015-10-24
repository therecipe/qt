#include "qpangesture.h"
#include <QPointF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPoint>
#include <QPanGesture>
#include "_cgo_export.h"

class MyQPanGesture: public QPanGesture {
public:
};

void QPanGesture_SetLastOffset(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QPanGesture*>(ptr)->setLastOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_SetOffset(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QPanGesture*>(ptr)->setOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_DestroyQPanGesture(QtObjectPtr ptr){
	static_cast<QPanGesture*>(ptr)->~QPanGesture();
}

