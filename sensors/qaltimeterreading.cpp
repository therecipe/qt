#include "qaltimeterreading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAltimeter>
#include <QString>
#include <QAltimeterReading>
#include "_cgo_export.h"

class MyQAltimeterReading: public QAltimeterReading {
public:
};

double QAltimeterReading_Altitude(void* ptr){
	return static_cast<double>(static_cast<QAltimeterReading*>(ptr)->altitude());
}

void QAltimeterReading_SetAltitude(void* ptr, double altitude){
	static_cast<QAltimeterReading*>(ptr)->setAltitude(static_cast<qreal>(altitude));
}

