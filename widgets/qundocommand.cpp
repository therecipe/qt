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

QtObjectPtr QUndoCommand_NewQUndoCommand(QtObjectPtr parent){
	return new QUndoCommand(static_cast<QUndoCommand*>(parent));
}

QtObjectPtr QUndoCommand_NewQUndoCommand2(char* text, QtObjectPtr parent){
	return new QUndoCommand(QString(text), static_cast<QUndoCommand*>(parent));
}

char* QUndoCommand_ActionText(QtObjectPtr ptr){
	return static_cast<QUndoCommand*>(ptr)->actionText().toUtf8().data();
}

QtObjectPtr QUndoCommand_Child(QtObjectPtr ptr, int index){
	return const_cast<QUndoCommand*>(static_cast<QUndoCommand*>(ptr)->child(index));
}

int QUndoCommand_ChildCount(QtObjectPtr ptr){
	return static_cast<QUndoCommand*>(ptr)->childCount();
}

int QUndoCommand_Id(QtObjectPtr ptr){
	return static_cast<QUndoCommand*>(ptr)->id();
}

int QUndoCommand_MergeWith(QtObjectPtr ptr, QtObjectPtr command){
	return static_cast<QUndoCommand*>(ptr)->mergeWith(static_cast<QUndoCommand*>(command));
}

void QUndoCommand_Redo(QtObjectPtr ptr){
	static_cast<QUndoCommand*>(ptr)->redo();
}

void QUndoCommand_SetText(QtObjectPtr ptr, char* text){
	static_cast<QUndoCommand*>(ptr)->setText(QString(text));
}

char* QUndoCommand_Text(QtObjectPtr ptr){
	return static_cast<QUndoCommand*>(ptr)->text().toUtf8().data();
}

void QUndoCommand_Undo(QtObjectPtr ptr){
	static_cast<QUndoCommand*>(ptr)->undo();
}

void QUndoCommand_DestroyQUndoCommand(QtObjectPtr ptr){
	static_cast<QUndoCommand*>(ptr)->~QUndoCommand();
}

