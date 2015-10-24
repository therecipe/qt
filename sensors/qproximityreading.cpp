#include "qproximityreading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QProximityReading>
#include "_cgo_export.h"

class MyQProximityReading: public QProximityReading {
public:
};

int QProximityReading_Close(QtObjectPtr ptr){
	return static_cast<QProximityReading*>(ptr)->close();
}

void QProximityReading_SetClose(QtObjectPtr ptr, int close){
	static_cast<QProximityReading*>(ptr)->setClose(close != 0);
}

