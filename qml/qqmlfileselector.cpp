#include "qqmlfileselector.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFileSelector>
#include <QQmlEngine>
#include <QObject>
#include <QFile>
#include <QQmlFileSelector>
#include "_cgo_export.h"

class MyQQmlFileSelector: public QQmlFileSelector {
public:
};

void* QQmlFileSelector_NewQQmlFileSelector(void* engine, void* parent){
	return new QQmlFileSelector(static_cast<QQmlEngine*>(engine), static_cast<QObject*>(parent));
}

void* QQmlFileSelector_QQmlFileSelector_Get(void* engine){
	return QQmlFileSelector::get(static_cast<QQmlEngine*>(engine));
}

void QQmlFileSelector_SetExtraSelectors(void* ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetExtraSelectors2(void* ptr, char* strin){
	static_cast<QQmlFileSelector*>(ptr)->setExtraSelectors(QString(strin).split("|", QString::SkipEmptyParts));
}

void QQmlFileSelector_SetSelector(void* ptr, void* selector){
	static_cast<QQmlFileSelector*>(ptr)->setSelector(static_cast<QFileSelector*>(selector));
}

void QQmlFileSelector_DestroyQQmlFileSelector(void* ptr){
	static_cast<QQmlFileSelector*>(ptr)->~QQmlFileSelector();
}

