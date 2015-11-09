#include "qfileselector.h"
#include <QFile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QFileSelector>
#include "_cgo_export.h"

class MyQFileSelector: public QFileSelector {
public:
};

void* QFileSelector_NewQFileSelector(void* parent){
	return new QFileSelector(static_cast<QObject*>(parent));
}

char* QFileSelector_AllSelectors(void* ptr){
	return static_cast<QFileSelector*>(ptr)->allSelectors().join("|").toUtf8().data();
}

char* QFileSelector_ExtraSelectors(void* ptr){
	return static_cast<QFileSelector*>(ptr)->extraSelectors().join("|").toUtf8().data();
}

char* QFileSelector_Select(void* ptr, char* filePath){
	return static_cast<QFileSelector*>(ptr)->select(QString(filePath)).toUtf8().data();
}

void QFileSelector_SetExtraSelectors(void* ptr, char* list){
	static_cast<QFileSelector*>(ptr)->setExtraSelectors(QString(list).split("|", QString::SkipEmptyParts));
}

void QFileSelector_DestroyQFileSelector(void* ptr){
	static_cast<QFileSelector*>(ptr)->~QFileSelector();
}

