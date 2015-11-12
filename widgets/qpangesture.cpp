#include "qpangesture.h"
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPanGesture>
#include "_cgo_export.h"

class MyQPanGesture: public QPanGesture {
public:
};

double QPanGesture_Acceleration(void* ptr){
	return static_cast<double>(static_cast<QPanGesture*>(ptr)->acceleration());
}

void QPanGesture_SetAcceleration(void* ptr, double value){
	static_cast<QPanGesture*>(ptr)->setAcceleration(static_cast<qreal>(value));
}

void QPanGesture_SetLastOffset(void* ptr, void* value){
	static_cast<QPanGesture*>(ptr)->setLastOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_SetOffset(void* ptr, void* value){
	static_cast<QPanGesture*>(ptr)->setOffset(*static_cast<QPointF*>(value));
}

void QPanGesture_DestroyQPanGesture(void* ptr){
	static_cast<QPanGesture*>(ptr)->~QPanGesture();
}

