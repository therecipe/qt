#include "qmediatimeinterval.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QMediaTimeInterval>
#include "_cgo_export.h"

class MyQMediaTimeInterval: public QMediaTimeInterval {
public:
};

QtObjectPtr QMediaTimeInterval_NewQMediaTimeInterval(){
	return new QMediaTimeInterval();
}

QtObjectPtr QMediaTimeInterval_NewQMediaTimeInterval3(QtObjectPtr other){
	return new QMediaTimeInterval(*static_cast<QMediaTimeInterval*>(other));
}

int QMediaTimeInterval_IsNormal(QtObjectPtr ptr){
	return static_cast<QMediaTimeInterval*>(ptr)->isNormal();
}

