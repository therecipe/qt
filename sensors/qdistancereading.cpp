#include "qdistancereading.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDistanceReading>
#include "_cgo_export.h"

class MyQDistanceReading: public QDistanceReading {
public:
};

double QDistanceReading_Distance(void* ptr){
	return static_cast<double>(static_cast<QDistanceReading*>(ptr)->distance());
}

void QDistanceReading_SetDistance(void* ptr, double distance){
	static_cast<QDistanceReading*>(ptr)->setDistance(static_cast<qreal>(distance));
}

