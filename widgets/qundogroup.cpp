#include "qundogroup.h"
#include <QObject>
#include <QUndoStack>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QUndoGroup>
#include "_cgo_export.h"

class MyQUndoGroup: public QUndoGroup {
public:
void Signal_ActiveStackChanged(QUndoStack * stack){callbackQUndoGroupActiveStackChanged(this->objectName().toUtf8().data(), stack);};
void Signal_CanRedoChanged(bool canRedo){callbackQUndoGroupCanRedoChanged(this->objectName().toUtf8().data(), canRedo);};
void Signal_CanUndoChanged(bool canUndo){callbackQUndoGroupCanUndoChanged(this->objectName().toUtf8().data(), canUndo);};
void Signal_CleanChanged(bool clean){callbackQUndoGroupCleanChanged(this->objectName().toUtf8().data(), clean);};
void Signal_IndexChanged(int idx){callbackQUndoGroupIndexChanged(this->objectName().toUtf8().data(), idx);};
void Signal_RedoTextChanged(const QString & redoText){callbackQUndoGroupRedoTextChanged(this->objectName().toUtf8().data(), redoText.toUtf8().data());};
void Signal_UndoTextChanged(const QString & undoText){callbackQUndoGroupUndoTextChanged(this->objectName().toUtf8().data(), undoText.toUtf8().data());};
};

QtObjectPtr QUndoGroup_NewQUndoGroup(QtObjectPtr parent){
	return new QUndoGroup(static_cast<QObject*>(parent));
}

QtObjectPtr QUndoGroup_ActiveStack(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->activeStack();
}

void QUndoGroup_ConnectActiveStackChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(QUndoStack *)>(&QUndoGroup::activeStackChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(QUndoStack *)>(&MyQUndoGroup::Signal_ActiveStackChanged));;
}

void QUndoGroup_DisconnectActiveStackChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(QUndoStack *)>(&QUndoGroup::activeStackChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(QUndoStack *)>(&MyQUndoGroup::Signal_ActiveStackChanged));;
}

void QUndoGroup_AddStack(QtObjectPtr ptr, QtObjectPtr stack){
	static_cast<QUndoGroup*>(ptr)->addStack(static_cast<QUndoStack*>(stack));
}

int QUndoGroup_CanRedo(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->canRedo();
}

void QUndoGroup_ConnectCanRedoChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canRedoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanRedoChanged));;
}

void QUndoGroup_DisconnectCanRedoChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canRedoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanRedoChanged));;
}

int QUndoGroup_CanUndo(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->canUndo();
}

void QUndoGroup_ConnectCanUndoChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canUndoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanUndoChanged));;
}

void QUndoGroup_DisconnectCanUndoChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::canUndoChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CanUndoChanged));;
}

void QUndoGroup_ConnectCleanChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::cleanChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CleanChanged));;
}

void QUndoGroup_DisconnectCleanChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(bool)>(&QUndoGroup::cleanChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(bool)>(&MyQUndoGroup::Signal_CleanChanged));;
}

QtObjectPtr QUndoGroup_CreateRedoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix){
	return static_cast<QUndoGroup*>(ptr)->createRedoAction(static_cast<QObject*>(parent), QString(prefix));
}

QtObjectPtr QUndoGroup_CreateUndoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix){
	return static_cast<QUndoGroup*>(ptr)->createUndoAction(static_cast<QObject*>(parent), QString(prefix));
}

void QUndoGroup_ConnectIndexChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(int)>(&QUndoGroup::indexChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(int)>(&MyQUndoGroup::Signal_IndexChanged));;
}

void QUndoGroup_DisconnectIndexChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(int)>(&QUndoGroup::indexChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(int)>(&MyQUndoGroup::Signal_IndexChanged));;
}

int QUndoGroup_IsClean(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->isClean();
}

void QUndoGroup_Redo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "redo");
}

char* QUndoGroup_RedoText(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->redoText().toUtf8().data();
}

void QUndoGroup_ConnectRedoTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::redoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_RedoTextChanged));;
}

void QUndoGroup_DisconnectRedoTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::redoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_RedoTextChanged));;
}

void QUndoGroup_RemoveStack(QtObjectPtr ptr, QtObjectPtr stack){
	static_cast<QUndoGroup*>(ptr)->removeStack(static_cast<QUndoStack*>(stack));
}

void QUndoGroup_SetActiveStack(QtObjectPtr ptr, QtObjectPtr stack){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "setActiveStack", Q_ARG(QUndoStack*, static_cast<QUndoStack*>(stack)));
}

void QUndoGroup_Undo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QUndoGroup*>(ptr), "undo");
}

char* QUndoGroup_UndoText(QtObjectPtr ptr){
	return static_cast<QUndoGroup*>(ptr)->undoText().toUtf8().data();
}

void QUndoGroup_ConnectUndoTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::undoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_UndoTextChanged));;
}

void QUndoGroup_DisconnectUndoTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoGroup*>(ptr), static_cast<void (QUndoGroup::*)(const QString &)>(&QUndoGroup::undoTextChanged), static_cast<MyQUndoGroup*>(ptr), static_cast<void (MyQUndoGroup::*)(const QString &)>(&MyQUndoGroup::Signal_UndoTextChanged));;
}

void QUndoGroup_DestroyQUndoGroup(QtObjectPtr ptr){
	static_cast<QUndoGroup*>(ptr)->~QUndoGroup();
}

