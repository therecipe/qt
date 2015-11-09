#include "qtapreading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTapReading>
#include "_cgo_export.h"

class MyQTapReading: public QTapReading {
public:
};

int QTapReading_IsDoubleTap(void* ptr){
	return static_cast<QTapReading*>(ptr)->isDoubleTap();
}

int QTapReading_TapDirection(void* ptr){
	return static_cast<QTapReading*>(ptr)->tapDirection();
}

void QTapReading_SetDoubleTap(void* ptr, int doubleTap){
	static_cast<QTapReading*>(ptr)->setDoubleTap(doubleTap != 0);
}

void QTapReading_SetTapDirection(void* ptr, int tapDirection){
	static_cast<QTapReading*>(ptr)->setTapDirection(static_cast<QTapReading::TapDirection>(tapDirection));
}

