#include "qaltimeter.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAltimeter>
#include "_cgo_export.h"

class MyQAltimeter: public QAltimeter {
public:
};

QtObjectPtr QAltimeter_Reading(QtObjectPtr ptr){
	return static_cast<QAltimeter*>(ptr)->reading();
}

QtObjectPtr QAltimeter_NewQAltimeter(QtObjectPtr parent){
	return new QAltimeter(static_cast<QObject*>(parent));
}

void QAltimeter_DestroyQAltimeter(QtObjectPtr ptr){
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

