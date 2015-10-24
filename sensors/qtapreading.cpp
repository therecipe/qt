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

int QTapReading_IsDoubleTap(QtObjectPtr ptr){
	return static_cast<QTapReading*>(ptr)->isDoubleTap();
}

int QTapReading_TapDirection(QtObjectPtr ptr){
	return static_cast<QTapReading*>(ptr)->tapDirection();
}

void QTapReading_SetDoubleTap(QtObjectPtr ptr, int doubleTap){
	static_cast<QTapReading*>(ptr)->setDoubleTap(doubleTap != 0);
}

void QTapReading_SetTapDirection(QtObjectPtr ptr, int tapDirection){
	static_cast<QTapReading*>(ptr)->setTapDirection(static_cast<QTapReading::TapDirection>(tapDirection));
}

