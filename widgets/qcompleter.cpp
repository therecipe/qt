#include "qcompleter.h"
#include <QUrl>
#include <QAbstractItemView>
#include <QWidget>
#include <QRect>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QCompleter>
#include "_cgo_export.h"

class MyQCompleter: public QCompleter {
public:
void Signal_Activated(const QString & text){callbackQCompleterActivated(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_Highlighted(const QString & text){callbackQCompleterHighlighted(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

int QCompleter_CaseSensitivity(void* ptr){
	return static_cast<QCompleter*>(ptr)->caseSensitivity();
}

int QCompleter_CompletionColumn(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionColumn();
}

int QCompleter_CompletionMode(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionMode();
}

char* QCompleter_CompletionPrefix(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionPrefix().toUtf8().data();
}

int QCompleter_CompletionRole(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionRole();
}

int QCompleter_FilterMode(void* ptr){
	return static_cast<QCompleter*>(ptr)->filterMode();
}

int QCompleter_MaxVisibleItems(void* ptr){
	return static_cast<QCompleter*>(ptr)->maxVisibleItems();
}

int QCompleter_ModelSorting(void* ptr){
	return static_cast<QCompleter*>(ptr)->modelSorting();
}

void QCompleter_SetCaseSensitivity(void* ptr, int caseSensitivity){
	static_cast<QCompleter*>(ptr)->setCaseSensitivity(static_cast<Qt::CaseSensitivity>(caseSensitivity));
}

void QCompleter_SetCompletionColumn(void* ptr, int column){
	static_cast<QCompleter*>(ptr)->setCompletionColumn(column);
}

void QCompleter_SetCompletionMode(void* ptr, int mode){
	static_cast<QCompleter*>(ptr)->setCompletionMode(static_cast<QCompleter::CompletionMode>(mode));
}

void QCompleter_SetCompletionPrefix(void* ptr, char* prefix){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setCompletionPrefix", Q_ARG(QString, QString(prefix)));
}

void QCompleter_SetCompletionRole(void* ptr, int role){
	static_cast<QCompleter*>(ptr)->setCompletionRole(role);
}

void QCompleter_SetFilterMode(void* ptr, int filterMode){
	static_cast<QCompleter*>(ptr)->setFilterMode(static_cast<Qt::MatchFlag>(filterMode));
}

void QCompleter_SetMaxVisibleItems(void* ptr, int maxItems){
	static_cast<QCompleter*>(ptr)->setMaxVisibleItems(maxItems);
}

void QCompleter_SetModelSorting(void* ptr, int sorting){
	static_cast<QCompleter*>(ptr)->setModelSorting(static_cast<QCompleter::ModelSorting>(sorting));
}

void QCompleter_SetWrapAround(void* ptr, int wrap){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "setWrapAround", Q_ARG(bool, wrap != 0));
}

int QCompleter_WrapAround(void* ptr){
	return static_cast<QCompleter*>(ptr)->wrapAround();
}

void* QCompleter_NewQCompleter2(void* model, void* parent){
	return new QCompleter(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

void* QCompleter_NewQCompleter(void* parent){
	return new QCompleter(static_cast<QObject*>(parent));
}

void* QCompleter_NewQCompleter3(char* list, void* parent){
	return new QCompleter(QString(list).split("|", QString::SkipEmptyParts), static_cast<QObject*>(parent));
}

void QCompleter_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::activated), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Activated));;
}

void QCompleter_Complete(void* ptr, void* rect){
	QMetaObject::invokeMethod(static_cast<QCompleter*>(ptr), "complete", Q_ARG(QRect, *static_cast<QRect*>(rect)));
}

int QCompleter_CompletionCount(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionCount();
}

void* QCompleter_CompletionModel(void* ptr){
	return static_cast<QCompleter*>(ptr)->completionModel();
}

char* QCompleter_CurrentCompletion(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentCompletion().toUtf8().data();
}

void* QCompleter_CurrentIndex(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentIndex().internalPointer();
}

int QCompleter_CurrentRow(void* ptr){
	return static_cast<QCompleter*>(ptr)->currentRow();
}

void QCompleter_ConnectHighlighted(void* ptr){
	QObject::connect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

void QCompleter_DisconnectHighlighted(void* ptr){
	QObject::disconnect(static_cast<QCompleter*>(ptr), static_cast<void (QCompleter::*)(const QString &)>(&QCompleter::highlighted), static_cast<MyQCompleter*>(ptr), static_cast<void (MyQCompleter::*)(const QString &)>(&MyQCompleter::Signal_Highlighted));;
}

void* QCompleter_Model(void* ptr){
	return static_cast<QCompleter*>(ptr)->model();
}

char* QCompleter_PathFromIndex(void* ptr, void* index){
	return static_cast<QCompleter*>(ptr)->pathFromIndex(*static_cast<QModelIndex*>(index)).toUtf8().data();
}

void* QCompleter_Popup(void* ptr){
	return static_cast<QCompleter*>(ptr)->popup();
}

int QCompleter_SetCurrentRow(void* ptr, int row){
	return static_cast<QCompleter*>(ptr)->setCurrentRow(row);
}

void QCompleter_SetModel(void* ptr, void* model){
	static_cast<QCompleter*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QCompleter_SetPopup(void* ptr, void* popup){
	static_cast<QCompleter*>(ptr)->setPopup(static_cast<QAbstractItemView*>(popup));
}

void QCompleter_SetWidget(void* ptr, void* widget){
	static_cast<QCompleter*>(ptr)->setWidget(static_cast<QWidget*>(widget));
}

char* QCompleter_SplitPath(void* ptr, char* path){
	return static_cast<QCompleter*>(ptr)->splitPath(QString(path)).join("|").toUtf8().data();
}

void* QCompleter_Widget(void* ptr){
	return static_cast<QCompleter*>(ptr)->widget();
}

void QCompleter_DestroyQCompleter(void* ptr){
	static_cast<QCompleter*>(ptr)->~QCompleter();
}

