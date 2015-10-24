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

QtObjectPtr QMediaTimeRange_NewQMediaTimeRange(){
	return new QMediaTimeRange();
}

QtObjectPtr QMediaTimeRange_NewQMediaTimeRange3(QtObjectPtr interval){
	return new QMediaTimeRange(*static_cast<QMediaTimeInterval*>(interval));
}

QtObjectPtr QMediaTimeRange_NewQMediaTimeRange4(QtObjectPtr ran){
	return new QMediaTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_AddInterval(QtObjectPtr ptr, QtObjectPtr interval){
	static_cast<QMediaTimeRange*>(ptr)->addInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_AddTimeRange(QtObjectPtr ptr, QtObjectPtr ran){
	static_cast<QMediaTimeRange*>(ptr)->addTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_Clear(QtObjectPtr ptr){
	static_cast<QMediaTimeRange*>(ptr)->clear();
}

int QMediaTimeRange_IsContinuous(QtObjectPtr ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isContinuous();
}

int QMediaTimeRange_IsEmpty(QtObjectPtr ptr){
	return static_cast<QMediaTimeRange*>(ptr)->isEmpty();
}

void QMediaTimeRange_RemoveInterval(QtObjectPtr ptr, QtObjectPtr interval){
	static_cast<QMediaTimeRange*>(ptr)->removeInterval(*static_cast<QMediaTimeInterval*>(interval));
}

void QMediaTimeRange_RemoveTimeRange(QtObjectPtr ptr, QtObjectPtr ran){
	static_cast<QMediaTimeRange*>(ptr)->removeTimeRange(*static_cast<QMediaTimeRange*>(ran));
}

void QMediaTimeRange_DestroyQMediaTimeRange(QtObjectPtr ptr){
	static_cast<QMediaTimeRange*>(ptr)->~QMediaTimeRange();
}

