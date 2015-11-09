#include "qundocommand.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QUndoCommand>
#include "_cgo_export.h"

class MyQUndoCommand: public QUndoCommand {
public:
};

void* QUndoCommand_NewQUndoCommand(void* parent){
	return new QUndoCommand(static_cast<QUndoCommand*>(parent));
}

void* QUndoCommand_NewQUndoCommand2(char* text, void* parent){
	return new QUndoCommand(QString(text), static_cast<QUndoCommand*>(parent));
}

char* QUndoCommand_ActionText(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->actionText().toUtf8().data();
}

void* QUndoCommand_Child(void* ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoCommand*>(ptr)->child(index));
}

int QUndoCommand_ChildCount(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->childCount();
}

int QUndoCommand_Id(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->id();
}

int QUndoCommand_MergeWith(void* ptr, void* command){
	return static_cast<QUndoCommand*>(ptr)->mergeWith(static_cast<QUndoCommand*>(command));
}

void QUndoCommand_Redo(void* ptr){
	static_cast<QUndoCommand*>(ptr)->redo();
}

void QUndoCommand_SetText(void* ptr, char* text){
	static_cast<QUndoCommand*>(ptr)->setText(QString(text));
}

char* QUndoCommand_Text(void* ptr){
	return static_cast<QUndoCommand*>(ptr)->text().toUtf8().data();
}

void QUndoCommand_Undo(void* ptr){
	static_cast<QUndoCommand*>(ptr)->undo();
}

void QUndoCommand_DestroyQUndoCommand(void* ptr){
	static_cast<QUndoCommand*>(ptr)->~QUndoCommand();
}

