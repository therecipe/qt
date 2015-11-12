#include "qhelpsearchengine.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QHelpEngineCore>
#include <QHelpEngine>
#include <QString>
#include <QVariant>
#include <QHelpSearchEngine>
#include "_cgo_export.h"

class MyQHelpSearchEngine: public QHelpSearchEngine {
public:
void Signal_IndexingFinished(){callbackQHelpSearchEngineIndexingFinished(this->objectName().toUtf8().data());};
void Signal_IndexingStarted(){callbackQHelpSearchEngineIndexingStarted(this->objectName().toUtf8().data());};
void Signal_SearchingFinished(int hits){callbackQHelpSearchEngineSearchingFinished(this->objectName().toUtf8().data(), hits);};
void Signal_SearchingStarted(){callbackQHelpSearchEngineSearchingStarted(this->objectName().toUtf8().data());};
};

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent){
	return new QHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_ConnectIndexingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void* QHelpSearchEngine_QueryWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void* QHelpSearchEngine_ResultWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_ConnectSearchingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

