#include "qaltimeter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QAltimeter>
#include "_cgo_export.h"

class MyQAltimeter: public QAltimeter {
public:
};

void* QAltimeter_Reading(void* ptr){
	return static_cast<QAltimeter*>(ptr)->reading();
}

void* QAltimeter_NewQAltimeter(void* parent){
	return new QAltimeter(static_cast<QObject*>(parent));
}

void QAltimeter_DestroyQAltimeter(void* ptr){
	static_cast<QAltimeter*>(ptr)->~QAltimeter();
}

