#include "qtoolbox.h"
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QIcon>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QToolBox>
#include "_cgo_export.h"

class MyQToolBox: public QToolBox {
public:
void Signal_CurrentChanged(int index){callbackQToolBoxCurrentChanged(this->objectName().toUtf8().data(), index);};
};

int QToolBox_Count(void* ptr){
	return static_cast<QToolBox*>(ptr)->count();
}

int QToolBox_CurrentIndex(void* ptr){
	return static_cast<QToolBox*>(ptr)->currentIndex();
}

void QToolBox_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void* QToolBox_NewQToolBox(void* parent, int f){
	return new QToolBox(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QToolBox_AddItem2(void* ptr, void* w, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(w), QString(text));
}

int QToolBox_AddItem(void* ptr, void* widget, void* iconSet, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(widget), *static_cast<QIcon*>(iconSet), QString(text));
}

void QToolBox_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

void QToolBox_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

void* QToolBox_CurrentWidget(void* ptr){
	return static_cast<QToolBox*>(ptr)->currentWidget();
}

int QToolBox_IndexOf(void* ptr, void* widget){
	return static_cast<QToolBox*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QToolBox_InsertItem(void* ptr, int index, void* widget, void* icon, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), *static_cast<QIcon*>(icon), QString(text));
}

int QToolBox_InsertItem2(void* ptr, int index, void* widget, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), QString(text));
}

int QToolBox_IsItemEnabled(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->isItemEnabled(index);
}

char* QToolBox_ItemText(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemText(index).toUtf8().data();
}

char* QToolBox_ItemToolTip(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemToolTip(index).toUtf8().data();
}

void QToolBox_RemoveItem(void* ptr, int index){
	static_cast<QToolBox*>(ptr)->removeItem(index);
}

void QToolBox_SetCurrentWidget(void* ptr, void* widget){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QToolBox_SetItemEnabled(void* ptr, int index, int enabled){
	static_cast<QToolBox*>(ptr)->setItemEnabled(index, enabled != 0);
}

void QToolBox_SetItemIcon(void* ptr, int index, void* icon){
	static_cast<QToolBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QToolBox_SetItemText(void* ptr, int index, char* text){
	static_cast<QToolBox*>(ptr)->setItemText(index, QString(text));
}

void QToolBox_SetItemToolTip(void* ptr, int index, char* toolTip){
	static_cast<QToolBox*>(ptr)->setItemToolTip(index, QString(toolTip));
}

void* QToolBox_Widget(void* ptr, int index){
	return static_cast<QToolBox*>(ptr)->widget(index);
}

void QToolBox_DestroyQToolBox(void* ptr){
	static_cast<QToolBox*>(ptr)->~QToolBox();
}

