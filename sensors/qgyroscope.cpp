#include "qgyroscope.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QGyroscope>
#include "_cgo_export.h"

class MyQGyroscope: public QGyroscope {
public:
};

void* QGyroscope_Reading(void* ptr){
	return static_cast<QGyroscope*>(ptr)->reading();
}

void* QGyroscope_NewQGyroscope(void* parent){
	return new QGyroscope(static_cast<QObject*>(parent));
}

void QGyroscope_DestroyQGyroscope(void* ptr){
	static_cast<QGyroscope*>(ptr)->~QGyroscope();
}

