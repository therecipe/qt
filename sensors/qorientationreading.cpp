#include "qorientationreading.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QOrientationReading>
#include "_cgo_export.h"

class MyQOrientationReading: public QOrientationReading {
public:
};

int QOrientationReading_Orientation(void* ptr){
	return static_cast<QOrientationReading*>(ptr)->orientation();
}

void QOrientationReading_SetOrientation(void* ptr, int orientation){
	static_cast<QOrientationReading*>(ptr)->setOrientation(static_cast<QOrientationReading::Orientation>(orientation));
}

