#include "qsplitter.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QObject>
#include <QWidget>
#include <QSplitter>
#include "_cgo_export.h"

class MyQSplitter: public QSplitter {
public:
void Signal_SplitterMoved(int pos, int index){callbackQSplitterSplitterMoved(this->objectName().toUtf8().data(), pos, index);};
};

int QSplitter_ChildrenCollapsible(QtObjectPtr ptr){
	return static_cast<QSplitter*>(ptr)->childrenCollapsible();
}

int QSplitter_HandleWidth(QtObjectPtr ptr){
	return static_cast<QSplitter*>(ptr)->handleWidth();
}

int QSplitter_IndexOf(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QSplitter*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QSplitter_OpaqueResize(QtObjectPtr ptr){
	return static_cast<QSplitter*>(ptr)->opaqueResize();
}

int QSplitter_Orientation(QtObjectPtr ptr){
	return static_cast<QSplitter*>(ptr)->orientation();
}

void QSplitter_SetChildrenCollapsible(QtObjectPtr ptr, int v){
	static_cast<QSplitter*>(ptr)->setChildrenCollapsible(v != 0);
}

void QSplitter_SetHandleWidth(QtObjectPtr ptr, int v){
	static_cast<QSplitter*>(ptr)->setHandleWidth(v);
}

void QSplitter_SetOpaqueResize(QtObjectPtr ptr, int opaque){
	static_cast<QSplitter*>(ptr)->setOpaqueResize(opaque != 0);
}

void QSplitter_SetOrientation(QtObjectPtr ptr, int v){
	static_cast<QSplitter*>(ptr)->setOrientation(static_cast<Qt::Orientation>(v));
}

QtObjectPtr QSplitter_NewQSplitter(QtObjectPtr parent){
	return new QSplitter(static_cast<QWidget*>(parent));
}

QtObjectPtr QSplitter_NewQSplitter2(int orientation, QtObjectPtr parent){
	return new QSplitter(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QSplitter_AddWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QSplitter*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

int QSplitter_Count(QtObjectPtr ptr){
	return static_cast<QSplitter*>(ptr)->count();
}

void QSplitter_GetRange(QtObjectPtr ptr, int index, int min, int max){
	static_cast<QSplitter*>(ptr)->getRange(index, &min, &max);
}

QtObjectPtr QSplitter_Handle(QtObjectPtr ptr, int index){
	return static_cast<QSplitter*>(ptr)->handle(index);
}

void QSplitter_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget){
	static_cast<QSplitter*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

int QSplitter_IsCollapsible(QtObjectPtr ptr, int index){
	return static_cast<QSplitter*>(ptr)->isCollapsible(index);
}

void QSplitter_Refresh(QtObjectPtr ptr){
	static_cast<QSplitter*>(ptr)->refresh();
}

int QSplitter_RestoreState(QtObjectPtr ptr, QtObjectPtr state){
	return static_cast<QSplitter*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void QSplitter_SetCollapsible(QtObjectPtr ptr, int index, int collapse){
	static_cast<QSplitter*>(ptr)->setCollapsible(index, collapse != 0);
}

void QSplitter_SetStretchFactor(QtObjectPtr ptr, int index, int stretch){
	static_cast<QSplitter*>(ptr)->setStretchFactor(index, stretch);
}

void QSplitter_ConnectSplitterMoved(QtObjectPtr ptr){
	QObject::connect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

void QSplitter_DisconnectSplitterMoved(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

QtObjectPtr QSplitter_Widget(QtObjectPtr ptr, int index){
	return static_cast<QSplitter*>(ptr)->widget(index);
}

void QSplitter_DestroyQSplitter(QtObjectPtr ptr){
	static_cast<QSplitter*>(ptr)->~QSplitter();
}

