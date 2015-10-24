#include "qtextbrowser.h"
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QTextBrowser>
#include "_cgo_export.h"

class MyQTextBrowser: public QTextBrowser {
public:
void Signal_AnchorClicked(const QUrl & link){callbackQTextBrowserAnchorClicked(this->objectName().toUtf8().data(), link.toString().toUtf8().data());};
void Signal_BackwardAvailable(bool available){callbackQTextBrowserBackwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_ForwardAvailable(bool available){callbackQTextBrowserForwardAvailable(this->objectName().toUtf8().data(), available);};
void Signal_Highlighted(const QUrl & link){callbackQTextBrowserHighlighted(this->objectName().toUtf8().data(), link.toString().toUtf8().data());};
void Signal_HistoryChanged(){callbackQTextBrowserHistoryChanged(this->objectName().toUtf8().data());};
void Signal_SourceChanged(const QUrl & src){callbackQTextBrowserSourceChanged(this->objectName().toUtf8().data(), src.toString().toUtf8().data());};
};

int QTextBrowser_OpenExternalLinks(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->openExternalLinks();
}

int QTextBrowser_OpenLinks(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->openLinks();
}

char* QTextBrowser_SearchPaths(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->searchPaths().join("|").toUtf8().data();
}

void QTextBrowser_SetOpenExternalLinks(QtObjectPtr ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenExternalLinks(open != 0);
}

void QTextBrowser_SetOpenLinks(QtObjectPtr ptr, int open){
	static_cast<QTextBrowser*>(ptr)->setOpenLinks(open != 0);
}

void QTextBrowser_SetSearchPaths(QtObjectPtr ptr, char* paths){
	static_cast<QTextBrowser*>(ptr)->setSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QTextBrowser_SetSource(QtObjectPtr ptr, char* name){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "setSource", Q_ARG(QUrl, QUrl(QString(name))));
}

char* QTextBrowser_Source(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->source().toString().toUtf8().data();
}

QtObjectPtr QTextBrowser_NewQTextBrowser(QtObjectPtr parent){
	return new QTextBrowser(static_cast<QWidget*>(parent));
}

void QTextBrowser_ConnectAnchorClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::anchorClicked), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_AnchorClicked));;
}

void QTextBrowser_DisconnectAnchorClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::anchorClicked), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_AnchorClicked));;
}

void QTextBrowser_Backward(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "backward");
}

void QTextBrowser_ConnectBackwardAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

void QTextBrowser_DisconnectBackwardAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::backwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_BackwardAvailable));;
}

int QTextBrowser_BackwardHistoryCount(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->backwardHistoryCount();
}

void QTextBrowser_ClearHistory(QtObjectPtr ptr){
	static_cast<QTextBrowser*>(ptr)->clearHistory();
}

void QTextBrowser_Forward(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "forward");
}

void QTextBrowser_ConnectForwardAvailable(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

void QTextBrowser_DisconnectForwardAvailable(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(bool)>(&QTextBrowser::forwardAvailable), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(bool)>(&MyQTextBrowser::Signal_ForwardAvailable));;
}

int QTextBrowser_ForwardHistoryCount(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->forwardHistoryCount();
}

void QTextBrowser_ConnectHighlighted(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::highlighted), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_Highlighted));;
}

void QTextBrowser_DisconnectHighlighted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::highlighted), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_Highlighted));;
}

void QTextBrowser_ConnectHistoryChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

void QTextBrowser_DisconnectHistoryChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)()>(&QTextBrowser::historyChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)()>(&MyQTextBrowser::Signal_HistoryChanged));;
}

char* QTextBrowser_HistoryTitle(QtObjectPtr ptr, int i){
	return static_cast<QTextBrowser*>(ptr)->historyTitle(i).toUtf8().data();
}

char* QTextBrowser_HistoryUrl(QtObjectPtr ptr, int i){
	return static_cast<QTextBrowser*>(ptr)->historyUrl(i).toString().toUtf8().data();
}

void QTextBrowser_Home(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "home");
}

int QTextBrowser_IsBackwardAvailable(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->isBackwardAvailable();
}

int QTextBrowser_IsForwardAvailable(QtObjectPtr ptr){
	return static_cast<QTextBrowser*>(ptr)->isForwardAvailable();
}

char* QTextBrowser_LoadResource(QtObjectPtr ptr, int ty, char* name){
	return static_cast<QTextBrowser*>(ptr)->loadResource(ty, QUrl(QString(name))).toString().toUtf8().data();
}

void QTextBrowser_Reload(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTextBrowser*>(ptr), "reload");
}

void QTextBrowser_ConnectSourceChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::sourceChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_SourceChanged));;
}

void QTextBrowser_DisconnectSourceChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTextBrowser*>(ptr), static_cast<void (QTextBrowser::*)(const QUrl &)>(&QTextBrowser::sourceChanged), static_cast<MyQTextBrowser*>(ptr), static_cast<void (MyQTextBrowser::*)(const QUrl &)>(&MyQTextBrowser::Signal_SourceChanged));;
}

