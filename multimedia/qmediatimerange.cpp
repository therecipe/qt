#include "qmediatimerange.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaTimeInterval>
#include <QMediaTimeRange>
#include "_cgo_export.h"

class MyQMediaTimeRange: public QMediaTimeRange {
public:
};

void* QMediaTimeRange_NewQMediaTimeRange(){
	return new QMediaTimeRange();
}

void* QMediaTimeRange_NewQMediaTimeRange3(void* interval){
	return new QMediaTimeRange(*static_cast<QMediaTimeInterval*>(interval));
}

void* QMediaTimeRange_NewQMediaTimeRange4(void* ran){
	return new QMediaTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_AddInterval(void* ptr, void* interval){
	static_cast<QMediaTimeRange*>(ptr)->addInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_AddTimeRange(void* ptr, void* ran){
	static_cast<QMediaTimeRange*>(ptr)->addTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_Clear(void* ptr){
	static_cast<QMediaTimeRange*>(ptr)->clear();
}

int QMediaTimeRange_IsContinuous(void* ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isContinuous();
}

int QMediaTimeRange_IsEmpty(void* ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isEmpty();
}

void QMediaTimeRange_RemoveInterval(void* ptr, void* interval){
	static_cast<QMediaTimeRange*>(ptr)->removeInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_RemoveTimeRange(void* ptr, void* ran){
	static_cast<QMediaTimeRange*>(ptr)->removeTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_DestroyQMediaTimeRange(void* ptr){
	static_cast<QMediaTimeRange*>(ptr)->~QMediaTimeRange();
}

