#include "qcompleter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QAbstractItemModel>
#include <QObject>
#include <QRect>
#include <QAbstractItemView>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QCompleter>
#include "_cgo_export.h"

class MyQCompleter: public QCompleter {
public:
void Signal_Activated(const QString & text){callbackQCompleterActivated(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_Highlighted(const QString & text){callbackQCompleterHighlighted(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

int QCompleter_CaseSensitivity(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->caseSensitivity();
}

int QCompleter_CompletionColumn(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionColumn();
}

int QCompleter_CompletionMode(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionMode();
}

char* QCompleter_CompletionPrefix(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionPrefix().toUtf8().data();
}

int QCompleter_CompletionRole(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionRole();
}

int QCompleter_FilterMode(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->filterMode();
}

int QCompleter_MaxVisibleItems(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->maxVisibleItems();
}

int QCompleter_ModelSorting(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->modelSorting();
}

void QCompleter_SetCaseSensitivity(QtObjectPtr ptr, int caseSensitivity){
	static_cast<QCompleter*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(caseSensitivity));
}

void QCompleter_SetCompletionColumn(QtObjectPtr ptr, int column){
	static_cast<QCompleter*>(ptr)->setCompletionColumn(column);
}

void QCompleter_SetCompletionMode(QtObjectPtr ptr, int mode){
	static_cast<QCompleter*>(ptr)->setCompletionMode(static_cast<QCompleter::CompletionMode>(mode));
}

void QCompleter_SetCompletionPrefix(QtObjectPtr ptr, char* prefix){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setCompletionPrefix", Q_ARG(QString, QString(prefix)));
}

void QCompleter_SetCompletionRole(QtObjectPtr ptr, int role){
	static_cast<QCompleter*>(ptr)->setCompletionRole(role);
}

void QCompleter_SetFilterMode(QtObjectPtr ptr, int filterMode){
	static_cast<QCompleter*>(ptr)->setFilterMode(static_cast<Qt::MatchFlag>(filterMode));
}

void QCompleter_SetMaxVisibleItems(QtObjectPtr ptr, int maxItems){
	static_cast<QCompleter*>(ptr)->setMaxVisibleItems(maxItems);
}

void QCompleter_SetModelSorting(QtObjectPtr ptr, int sorting){
	static_cast<QCompleter*>(ptr)->setModelSorting(static_cast<QCompleter::ModelSorting>(sorting));
}

void QCompleter_SetWrapAround(QtObjectPtr ptr, int wrap){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setWrapAround", Q_ARG(bool, wrap != 0));
}

int QCompleter_WrapAround(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->wrapAround();
}

QtObjectPtr QCompleter_NewQCompleter2(QtObjectPtr model, QtObjectPtr parent){
	return new QCompleter(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

QtObjectPtr QCompleter_NewQCompleter(QtObjectPtr parent){
	return new QCompleter(static_cast<QObject*>(parent));
}

QtObjectPtr QCompleter_NewQCompleter3(char* list, QtObjectPtr parent){
	return new QCompleter(QString(list).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

void QCompleter_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_Complete(QtObjectPtr ptr, QtObjectPtr rect){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "complete", Q_ARG(QRect, *static_cast<QRect*>(rect)));
}

int QCompleter_CompletionCount(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionCount();
}

QtObjectPtr QCompleter_CompletionModel(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->completionModel();
}

char* QCompleter_CurrentCompletion(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->currentCompletion().toUtf8().data();
}

QtObjectPtr QCompleter_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->currentIndex().internalPointer();
}

int QCompleter_CurrentRow(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->currentRow();
}

void QCompleter_ConnectHighlighted(QtObjectPtr ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

void QCompleter_DisconnectHighlighted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

QtObjectPtr QCompleter_Model(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->model();
}

char* QCompleter_PathFromIndex(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QCompleter*>(ptr)->pathFromIndex(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

QtObjectPtr QCompleter_Popup(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->popup();
}

int QCompleter_SetCurrentRow(QtObjectPtr ptr, int row){
	return static_cast<QCompleter*>(ptr)->setCurrentRow(row);
}

void QCompleter_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QCompleter*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QCompleter_SetPopup(QtObjectPtr ptr, QtObjectPtr popup){
	static_cast<QCompleter*>(ptr)->setPopup(static_cast<QAbstractItemView*>(popup));
}

void QCompleter_SetWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QCompleter*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

char* QCompleter_SplitPath(QtObjectPtr ptr, char* path){
	return static_cast<QCompleter*>(ptr)->splitPath(QString(path)).join("|").toUtf8().data();
}

QtObjectPtr QCompleter_Widget(QtObjectPtr ptr){
	return static_cast<QCompleter*>(ptr)->widget();
}

void QCompleter_DestroyQCompleter(QtObjectPtr ptr){
	static_cast<QCompleter*>(ptr)->~QCompleter();
}

