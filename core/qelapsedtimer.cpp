#include "qelapsedtimer.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QElapsedTimer>
#include "_cgo_export.h"

class MyQElapsedTimer: public QElapsedTimer {
public:
};

void* QElapsedTimer_NewQElapsedTimer(){
	return new QElapsedTimer();
}

void QElapsedTimer_Invalidate(void* ptr){
	static_cast<QElapsedTimer*>(ptr)->invalidate();
}

int QElapsedTimer_IsValid(void* ptr){
	return static_cast<QElapsedTimer*>(ptr)->isValid();
}

int QElapsedTimer_QElapsedTimer_ClockType(){
	return QElapsedTimer::clockType();
}

int QElapsedTimer_QElapsedTimer_IsMonotonic(){
	return QElapsedTimer::isMonotonic();
}

void QElapsedTimer_Start(void* ptr){
	static_cast<QElapsedTimer*>(ptr)->start();
}

