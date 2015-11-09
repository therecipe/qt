#include "qcompass.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QCompass>
#include "_cgo_export.h"

class MyQCompass: public QCompass {
public:
};

void* QCompass_Reading(void* ptr){
	return static_cast<QCompass*>(ptr)->reading();
}

void* QCompass_NewQCompass(void* parent){
	return new QCompass(static_cast<QObject*>(parent));
}

void QCompass_DestroyQCompass(void* ptr){
	static_cast<QCompass*>(ptr)->~QCompass();
}

