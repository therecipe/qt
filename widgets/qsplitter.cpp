#include "qsplitter.h"
#include <QByteArray>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QSplitter>
#include "_cgo_export.h"

class MyQSplitter: public QSplitter {
public:
void Signal_SplitterMoved(int pos, int index){callbackQSplitterSplitterMoved(this->objectName().toUtf8().data(), pos, index);};
};

int QSplitter_ChildrenCollapsible(void* ptr){
	return static_cast<QSplitter*>(ptr)->childrenCollapsible();
}

int QSplitter_HandleWidth(void* ptr){
	return static_cast<QSplitter*>(ptr)->handleWidth();
}

int QSplitter_IndexOf(void* ptr, void* widget){
	return static_cast<QSplitter*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QSplitter_OpaqueResize(void* ptr){
	return static_cast<QSplitter*>(ptr)->opaqueResize();
}

int QSplitter_Orientation(void* ptr){
	return static_cast<QSplitter*>(ptr)->orientation();
}

void QSplitter_SetChildrenCollapsible(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setChildrenCollapsible(v != 0);
}

void QSplitter_SetHandleWidth(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setHandleWidth(v);
}

void QSplitter_SetOpaqueResize(void* ptr, int opaque){
	static_cast<QSplitter*>(ptr)->setOpaqueResize(opaque != 0);
}

void QSplitter_SetOrientation(void* ptr, int v){
	static_cast<QSplitter*>(ptr)->setOrientation(static_cast<Qt::Orientation>(v));
}

void* QSplitter_NewQSplitter(void* parent){
	return new QSplitter(static_cast<QWidget*>(parent));
}

void* QSplitter_NewQSplitter2(int orientation, void* parent){
	return new QSplitter(static_cast<Qt::Orientation>(orientation), static_cast<QWidget*>(parent));
}

void QSplitter_AddWidget(void* ptr, void* widget){
	static_cast<QSplitter*>(ptr)->addWidget(static_cast<QWidget*>(widget));
}

int QSplitter_Count(void* ptr){
	return static_cast<QSplitter*>(ptr)->count();
}

void QSplitter_GetRange(void* ptr, int index, int min, int max){
	static_cast<QSplitter*>(ptr)->getRange(index, &min, &max);
}

void* QSplitter_Handle(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->handle(index);
}

void QSplitter_InsertWidget(void* ptr, int index, void* widget){
	static_cast<QSplitter*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget));
}

int QSplitter_IsCollapsible(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->isCollapsible(index);
}

void QSplitter_Refresh(void* ptr){
	static_cast<QSplitter*>(ptr)->refresh();
}

int QSplitter_RestoreState(void* ptr, void* state){
	return static_cast<QSplitter*>(ptr)->restoreState(*static_cast<QByteArray*>(state));
}

void* QSplitter_SaveState(void* ptr){
	return new QByteArray(static_cast<QSplitter*>(ptr)->saveState());
}

void QSplitter_SetCollapsible(void* ptr, int index, int collapse){
	static_cast<QSplitter*>(ptr)->setCollapsible(index, collapse != 0);
}

void QSplitter_SetStretchFactor(void* ptr, int index, int stretch){
	static_cast<QSplitter*>(ptr)->setStretchFactor(index, stretch);
}

void QSplitter_ConnectSplitterMoved(void* ptr){
	QObject::connect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

void QSplitter_DisconnectSplitterMoved(void* ptr){
	QObject::disconnect(static_cast<QSplitter*>(ptr), static_cast<void (QSplitter::*)(int, int)>(&QSplitter::splitterMoved), static_cast<MyQSplitter*>(ptr), static_cast<void (MyQSplitter::*)(int, int)>(&MyQSplitter::Signal_SplitterMoved));;
}

void* QSplitter_Widget(void* ptr, int index){
	return static_cast<QSplitter*>(ptr)->widget(index);
}

void QSplitter_DestroyQSplitter(void* ptr){
	static_cast<QSplitter*>(ptr)->~QSplitter();
}

