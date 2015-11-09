#include "qundoview.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QUndoStack>
#include <QMetaObject>
#include <QWidget>
#include <QString>
#include <QUndoGroup>
#include <QIcon>
#include <QUndoView>
#include "_cgo_export.h"

class MyQUndoView: public QUndoView {
public:
};

char* QUndoView_EmptyLabel(void* ptr){
	return static_cast<QUndoView*>(ptr)->emptyLabel().toUtf8().data();
}

void QUndoView_SetCleanIcon(void* ptr, void* icon){
	static_cast<QUndoView*>(ptr)->setCleanIcon(*static_cast<QIcon*>(icon));
}

void QUndoView_SetEmptyLabel(void* ptr, char* label){
	static_cast<QUndoView*>(ptr)->setEmptyLabel(QString(label));
}

void* QUndoView_NewQUndoView3(void* group, void* parent){
	return new QUndoView(static_cast<QUndoGroup*>(group), static_cast<QWidget*>(parent));
}

void* QUndoView_NewQUndoView2(void* stack, void* parent){
	return new QUndoView(static_cast<QUndoStack*>(stack), static_cast<QWidget*>(parent));
}

void* QUndoView_NewQUndoView(void* parent){
	return new QUndoView(static_cast<QWidget*>(parent));
}

void* QUndoView_Group(void* ptr){
	return static_cast<QUndoView*>(ptr)->group();
}

void QUndoView_SetGroup(void* ptr, void* group){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setGroup", Q_ARG(QUndoGroup*, static_cast<QUndoGroup*>(group)));
}

void QUndoView_SetStack(void* ptr, void* stack){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setStack", Q_ARG(QUndoStack*, static_cast<QUndoStack*>(stack)));
}

void* QUndoView_Stack(void* ptr){
	return static_cast<QUndoView*>(ptr)->stack();
}

void QUndoView_DestroyQUndoView(void* ptr){
	static_cast<QUndoView*>(ptr)->~QUndoView();
}

