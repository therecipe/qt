#include "qundostack.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QUndoCommand>
#include <QString>
#include <QUndoStack>
#include "_cgo_export.h"

class MyQUndoStack: public QUndoStack {
public:
void Signal_CanRedoChanged(bool canRedo){callbackQUndoStackCanRedoChanged(this->objectName().toUtf8().data(), canRedo);};
void Signal_CanUndoChanged(bool canUndo){callbackQUndoStackCanUndoChanged(this->objectName().toUtf8().data(), canUndo);};
void Signal_CleanChanged(bool clean){callbackQUndoStackCleanChanged(this->objectName().toUtf8().data(), clean);};
void Signal_IndexChanged(int idx){callbackQUndoStackIndexChanged(this->objectName().toUtf8().data(), idx);};
void Signal_RedoTextChanged(const QString & redoText){callbackQUndoStackRedoTextChanged(this->objectName().toUtf8().data(), redoText.toUtf8().data());};
void Signal_UndoTextChanged(const QString & undoText){callbackQUndoStackUndoTextChanged(this->objectName().toUtf8().data(), undoText.toUtf8().data());};
};

int QUndoStack_IsActive(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->isActive();
}

void QUndoStack_SetActive(QtObjectPtr ptr, int active){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setActive", Q_ARG(bool, active != 0));
}

void QUndoStack_SetUndoLimit(QtObjectPtr ptr, int limit){
	static_cast<QUndoStack*>(ptr)->setUndoLimit(limit);
}

int QUndoStack_UndoLimit(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->undoLimit();
}

QtObjectPtr QUndoStack_NewQUndoStack(QtObjectPtr parent){
	return new QUndoStack(static_cast<QObject*>(parent));
}

void QUndoStack_BeginMacro(QtObjectPtr ptr, char* text){
	static_cast<QUndoStack*>(ptr)->beginMacro(QString(text));
}

int QUndoStack_CanRedo(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->canRedo();
}

void QUndoStack_ConnectCanRedoChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

void QUndoStack_DisconnectCanRedoChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

int QUndoStack_CanUndo(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->canUndo();
}

void QUndoStack_ConnectCanUndoChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_DisconnectCanUndoChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_ConnectCleanChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

void QUndoStack_DisconnectCleanChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

int QUndoStack_CleanIndex(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->cleanIndex();
}

void QUndoStack_Clear(QtObjectPtr ptr){
	static_cast<QUndoStack*>(ptr)->clear();
}

QtObjectPtr QUndoStack_Command(QtObjectPtr ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoStack*>(ptr)->command(index));
}

int QUndoStack_Count(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->count();
}

QtObjectPtr QUndoStack_CreateRedoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createRedoAction(static_cast<QObject*>(parent), QString(prefix));
}

QtObjectPtr QUndoStack_CreateUndoAction(QtObjectPtr ptr, QtObjectPtr parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createUndoAction(static_cast<QObject*>(parent), QString(prefix));
}

void QUndoStack_EndMacro(QtObjectPtr ptr){
	static_cast<QUndoStack*>(ptr)->endMacro();
}

int QUndoStack_Index(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->index();
}

void QUndoStack_ConnectIndexChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

void QUndoStack_DisconnectIndexChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

int QUndoStack_IsClean(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->isClean();
}

void QUndoStack_Push(QtObjectPtr ptr, QtObjectPtr cmd){
	static_cast<QUndoStack*>(ptr)->push(static_cast<QUndoCommand*>(cmd));
}

void QUndoStack_Redo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "redo");
}

char* QUndoStack_RedoText(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->redoText().toUtf8().data();
}

void QUndoStack_ConnectRedoTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_DisconnectRedoTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_SetClean(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setClean");
}

void QUndoStack_SetIndex(QtObjectPtr ptr, int idx){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setIndex", Q_ARG(int, idx));
}

char* QUndoStack_Text(QtObjectPtr ptr, int idx){
	return static_cast<QUndoStack*>(ptr)->text(idx).toUtf8().data();
}

void QUndoStack_Undo(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "undo");
}

char* QUndoStack_UndoText(QtObjectPtr ptr){
	return static_cast<QUndoStack*>(ptr)->undoText().toUtf8().data();
}

void QUndoStack_ConnectUndoTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DisconnectUndoTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DestroyQUndoStack(QtObjectPtr ptr){
	static_cast<QUndoStack*>(ptr)->~QUndoStack();
}

