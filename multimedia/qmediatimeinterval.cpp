#include "qmediatimeinterval.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaTimeInterval>
#include "_cgo_export.h"

class MyQMediaTimeInterval: public QMediaTimeInterval {
public:
};

void* QMediaTimeInterval_NewQMediaTimeInterval(){
	return new QMediaTimeInterval();
}

void* QMediaTimeInterval_NewQMediaTimeInterval3(void* other){
	return new QMediaTimeInterval(*static_cast<QMediaTimeInterval*>(other));
}

int QMediaTimeInterval_IsNormal(void* ptr){
	return static_cast<QMediaTimeInterval*>(ptr)->isNormal();
}

