#include "qundostack.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QUndoCommand>
#include <QObject>
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

int QUndoStack_IsActive(void* ptr){
	return static_cast<QUndoStack*>(ptr)->isActive();
}

void QUndoStack_SetActive(void* ptr, int active){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setActive", Q_ARG(bool, active != 0));
}

void QUndoStack_SetUndoLimit(void* ptr, int limit){
	static_cast<QUndoStack*>(ptr)->setUndoLimit(limit);
}

int QUndoStack_UndoLimit(void* ptr){
	return static_cast<QUndoStack*>(ptr)->undoLimit();
}

void* QUndoStack_NewQUndoStack(void* parent){
	return new QUndoStack(static_cast<QObject*>(parent));
}

void QUndoStack_BeginMacro(void* ptr, char* text){
	static_cast<QUndoStack*>(ptr)->beginMacro(QString(text));
}

int QUndoStack_CanRedo(void* ptr){
	return static_cast<QUndoStack*>(ptr)->canRedo();
}

void QUndoStack_ConnectCanRedoChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

void QUndoStack_DisconnectCanRedoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canRedoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanRedoChanged));;
}

int QUndoStack_CanUndo(void* ptr){
	return static_cast<QUndoStack*>(ptr)->canUndo();
}

void QUndoStack_ConnectCanUndoChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_DisconnectCanUndoChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::canUndoChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CanUndoChanged));;
}

void QUndoStack_ConnectCleanChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

void QUndoStack_DisconnectCleanChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(bool)>(&QUndoStack::cleanChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(bool)>(&MyQUndoStack::Signal_CleanChanged));;
}

int QUndoStack_CleanIndex(void* ptr){
	return static_cast<QUndoStack*>(ptr)->cleanIndex();
}

void QUndoStack_Clear(void* ptr){
	static_cast<QUndoStack*>(ptr)->clear();
}

void* QUndoStack_Command(void* ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoStack*>(ptr)->command(index));
}

int QUndoStack_Count(void* ptr){
	return static_cast<QUndoStack*>(ptr)->count();
}

void* QUndoStack_CreateRedoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createRedoAction(static_cast<QObject*>(parent), QString(prefix));
}

void* QUndoStack_CreateUndoAction(void* ptr, void* parent, char* prefix){
	return static_cast<QUndoStack*>(ptr)->createUndoAction(static_cast<QObject*>(parent), QString(prefix));
}

void QUndoStack_EndMacro(void* ptr){
	static_cast<QUndoStack*>(ptr)->endMacro();
}

int QUndoStack_Index(void* ptr){
	return static_cast<QUndoStack*>(ptr)->index();
}

void QUndoStack_ConnectIndexChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

void QUndoStack_DisconnectIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(int)>(&QUndoStack::indexChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(int)>(&MyQUndoStack::Signal_IndexChanged));;
}

int QUndoStack_IsClean(void* ptr){
	return static_cast<QUndoStack*>(ptr)->isClean();
}

void QUndoStack_Push(void* ptr, void* cmd){
	static_cast<QUndoStack*>(ptr)->push(static_cast<QUndoCommand*>(cmd));
}

void QUndoStack_Redo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "redo");
}

char* QUndoStack_RedoText(void* ptr){
	return static_cast<QUndoStack*>(ptr)->redoText().toUtf8().data();
}

void QUndoStack_ConnectRedoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_DisconnectRedoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::redoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_RedoTextChanged));;
}

void QUndoStack_SetClean(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setClean");
}

void QUndoStack_SetIndex(void* ptr, int idx){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "setIndex", Q_ARG(int, idx));
}

char* QUndoStack_Text(void* ptr, int idx){
	return static_cast<QUndoStack*>(ptr)->text(idx).toUtf8().data();
}

void QUndoStack_Undo(void* ptr){
	QMetaObject::invokeMethod(static_cast<QUndoStack*>(ptr), "undo");
}

char* QUndoStack_UndoText(void* ptr){
	return static_cast<QUndoStack*>(ptr)->undoText().toUtf8().data();
}

void QUndoStack_ConnectUndoTextChanged(void* ptr){
	QObject::connect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DisconnectUndoTextChanged(void* ptr){
	QObject::disconnect(static_cast<QUndoStack*>(ptr), static_cast<void (QUndoStack::*)(const QString &)>(&QUndoStack::undoTextChanged), static_cast<MyQUndoStack*>(ptr), static_cast<void (MyQUndoStack::*)(const QString &)>(&MyQUndoStack::Signal_UndoTextChanged));;
}

void QUndoStack_DestroyQUndoStack(void* ptr){
	static_cast<QUndoStack*>(ptr)->~QUndoStack();
}

