#include "qundoview.h"
#include <QUndoStack>
#include <QWidget>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QUrl>
#include <QIcon>
#include <QUndoGroup>
#include <QUndoView>
#include "_cgo_export.h"

class MyQUndoView: public QUndoView {
public:
};

char* QUndoView_EmptyLabel(QtObjectPtr ptr){
	return static_cast<QUndoView*>(ptr)->emptyLabel().toUtf8().data();
}

void QUndoView_SetCleanIcon(QtObjectPtr ptr, QtObjectPtr icon){
	static_cast<QUndoView*>(ptr)->setCleanIcon(*static_cast<QIcon*>(icon));
}

void QUndoView_SetEmptyLabel(QtObjectPtr ptr, char* label){
	static_cast<QUndoView*>(ptr)->setEmptyLabel(QString(label));
}

QtObjectPtr QUndoView_NewQUndoView3(QtObjectPtr group, QtObjectPtr parent){
	return new QUndoView(static_cast<QUndoGroup*>(group), static_cast<QWidget*>(parent));
}

QtObjectPtr QUndoView_NewQUndoView2(QtObjectPtr stack, QtObjectPtr parent){
	return new QUndoView(static_cast<QUndoStack*>(stack), static_cast<QWidget*>(parent));
}

QtObjectPtr QUndoView_NewQUndoView(QtObjectPtr parent){
	return new QUndoView(static_cast<QWidget*>(parent));
}

QtObjectPtr QUndoView_Group(QtObjectPtr ptr){
	return static_cast<QUndoView*>(ptr)->group();
}

void QUndoView_SetGroup(QtObjectPtr ptr, QtObjectPtr group){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setGroup", Q_ARG(QUndoGroup*, static_cast<QUndoGroup*>(group)));
}

void QUndoView_SetStack(QtObjectPtr ptr, QtObjectPtr stack){
	QMetaObject::invokeMethod(static_cast<QUndoView*>(ptr), "setStack", Q_ARG(QUndoStack*, static_cast<QUndoStack*>(stack)));
}

QtObjectPtr QUndoView_Stack(QtObjectPtr ptr){
	return static_cast<QUndoView*>(ptr)->stack();
}

void QUndoView_DestroyQUndoView(QtObjectPtr ptr){
	static_cast<QUndoView*>(ptr)->~QUndoView();
}

