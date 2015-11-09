#include "qpinchgesture.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QPinchGesture>
#include "_cgo_export.h"

class MyQPinchGesture: public QPinchGesture {
public:
};

int QPinchGesture_ChangeFlags(void* ptr){
	return static_cast<QPinchGesture*>(ptr)->changeFlags();
}

double QPinchGesture_LastRotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->lastRotationAngle());
}

double QPinchGesture_LastScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->lastScaleFactor());
}

double QPinchGesture_RotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->rotationAngle());
}

double QPinchGesture_ScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->scaleFactor());
}

void QPinchGesture_SetCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetChangeFlags(void* ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

void QPinchGesture_SetLastCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setLastCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetLastRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setLastRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetLastScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setLastScaleFactor(static_cast<qreal>(value));
}

void QPinchGesture_SetRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setScaleFactor(static_cast<qreal>(value));
}

void QPinchGesture_SetStartCenterPoint(void* ptr, void* value){
	static_cast<QPinchGesture*>(ptr)->setStartCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetTotalChangeFlags(void* ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setTotalChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

void QPinchGesture_SetTotalRotationAngle(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setTotalRotationAngle(static_cast<qreal>(value));
}

void QPinchGesture_SetTotalScaleFactor(void* ptr, double value){
	static_cast<QPinchGesture*>(ptr)->setTotalScaleFactor(static_cast<qreal>(value));
}

int QPinchGesture_TotalChangeFlags(void* ptr){
	return static_cast<QPinchGesture*>(ptr)->totalChangeFlags();
}

double QPinchGesture_TotalRotationAngle(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->totalRotationAngle());
}

double QPinchGesture_TotalScaleFactor(void* ptr){
	return static_cast<double>(static_cast<QPinchGesture*>(ptr)->totalScaleFactor());
}

void QPinchGesture_DestroyQPinchGesture(void* ptr){
	static_cast<QPinchGesture*>(ptr)->~QPinchGesture();
}

