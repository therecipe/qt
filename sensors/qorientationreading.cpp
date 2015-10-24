#include "qorientationreading.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QOrientationReading>
#include "_cgo_export.h"

class MyQOrientationReading: public QOrientationReading {
public:
};

int QOrientationReading_Orientation(QtObjectPtr ptr){
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

void QOrientationReading_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

