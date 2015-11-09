#include "qtextbrowser.h"
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QTextBrowser>
#include "_cgo_export.h"

class MyQTextBrowser: public QTextBrowser {
public:
void Signal_BackwardAvailable(bool available){callbackQTextBrowserBackwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_ForwardAvailable(bool available){callbackQTextBrowserForwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_HistoryChanged(){callbackQTextBrowserHistoryChanged(this->objectName().toUtf8().data());};
};

int QTextBrowser_OpenExternalLinks(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->openExternalLinks();
}

int QTextBrowser_OpenLinks(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->openLinks();
}

char* QTextBrowser_SearchPaths(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->searchPaths().join("|").toUtf8().data();
}

void QTextBrowser_SetOpenExternalLinks(void* ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenExternalLinks(open != 0);
}

void QTextBrowser_SetOpenLinks(void* ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenLinks(open != 0);
}

void QTextBrowser_SetSearchPaths(void* ptr, char* paths){
	static_cast<QTextBrowser*>(ptr)->setSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QTextBrowser_SetSource(void* ptr, void* name){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "setSource", Q_ARG(QUrl, *static_cast<QUrl*>(name)));
}

void* QTextBrowser_NewQTextBrowser(void* parent){
	return new QTextBrowser(static_cast<QWidget*>(parent));
}

void QTextBrowser_Backward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "backward");
}

void QTextBrowser_ConnectBackwardAvailable(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

void QTextBrowser_DisconnectBackwardAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

int QTextBrowser_BackwardHistoryCount(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->backwardHistoryCount();
}

void QTextBrowser_ClearHistory(void* ptr){
	static_cast<QTextBrowser*>(ptr)->clearHistory();
}

void QTextBrowser_Forward(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "forward");
}

void QTextBrowser_ConnectForwardAvailable(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

void QTextBrowser_DisconnectForwardAvailable(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

int QTextBrowser_ForwardHistoryCount(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->forwardHistoryCount();
}

void QTextBrowser_ConnectHistoryChanged(void* ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

void QTextBrowser_DisconnectHistoryChanged(void* ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

char* QTextBrowser_HistoryTitle(void* ptr, int i){
	return static_cast<QTextBrowser*>(ptr)->historyTitle(i).toUtf8().data();
}

void QTextBrowser_Home(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "home");
}

int QTextBrowser_IsBackwardAvailable(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->isBackwardAvailable();
}

int QTextBrowser_IsForwardAvailable(void* ptr){
	return static_cast<QTextBrowser*>(ptr)->isForwardAvailable();
}

void* QTextBrowser_LoadResource(void* ptr, int ty, void* name){
	return new QVariant(static_cast<QTextBrowser*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

void QTextBrowser_Reload(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "reload");
}

