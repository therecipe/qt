#include "qpinchgesture.h"
#include <QModelIndex>
#include <QPointF>
#include <QPoint>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPinchGesture>
#include "_cgo_export.h"

class MyQPinchGesture: public QPinchGesture {
public:
};

int QPinchGesture_ChangeFlags(QtObjectPtr ptr){
	return static_cast<QPinchGesture*>(ptr)->changeFlags();
}

void QPinchGesture_SetCenterPoint(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QPinchGesture*>(ptr)->setCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetChangeFlags(QtObjectPtr ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

void QPinchGesture_SetLastCenterPoint(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QPinchGesture*>(ptr)->setLastCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetStartCenterPoint(QtObjectPtr ptr, QtObjectPtr value){
	static_cast<QPinchGesture*>(ptr)->setStartCenterPoint(*static_cast<QPointF*>(value));
}

void QPinchGesture_SetTotalChangeFlags(QtObjectPtr ptr, int value){
	static_cast<QPinchGesture*>(ptr)->setTotalChangeFlags(static_cast<QPinchGesture::ChangeFlag>(value));
}

int QPinchGesture_TotalChangeFlags(QtObjectPtr ptr){
	return static_cast<QPinchGesture*>(ptr)->totalChangeFlags();
}

void QPinchGesture_DestroyQPinchGesture(QtObjectPtr ptr){
	static_cast<QPinchGesture*>(ptr)->~QPinchGesture();
}

