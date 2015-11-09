#include "qhelpengine.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QHelpEngine>
#include "_cgo_export.h"

class MyQHelpEngine: public QHelpEngine {
public:
};

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent){
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

void* QHelpEngine_ContentModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_ContentWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_IndexModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

void* QHelpEngine_IndexWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

void* QHelpEngine_SearchEngine(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(void* ptr){
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

