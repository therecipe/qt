#include "qtoolbox.h"
#include <QUrl>
#include <QModelIndex>
#include <QIcon>
#include <QWidget>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QToolBox>
#include "_cgo_export.h"

class MyQToolBox: public QToolBox {
public:
void Signal_CurrentChanged(int index){callbackQToolBoxCurrentChanged(this->objectName().toUtf8().data(), index);};
};

int QToolBox_Count(QtObjectPtr ptr){
	return static_cast<QToolBox*>(ptr)->count();
}

int QToolBox_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QToolBox*>(ptr)->currentIndex();
}

void QToolBox_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

QtObjectPtr QToolBox_NewQToolBox(QtObjectPtr parent, int f){
	return new QToolBox(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

int QToolBox_AddItem2(QtObjectPtr ptr, QtObjectPtr w, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(w), QString(text));
}

int QToolBox_AddItem(QtObjectPtr ptr, QtObjectPtr widget, QtObjectPtr iconSet, char* text){
	return static_cast<QToolBox*>(ptr)->addItem(static_cast<QWidget*>(widget), *static_cast<QIcon*>(iconSet), QString(text));
}

void QToolBox_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

void QToolBox_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QToolBox*>(ptr), static_cast<void (QToolBox::*)(int)>(&QToolBox::currentChanged), static_cast<MyQToolBox*>(ptr), static_cast<void (MyQToolBox::*)(int)>(&MyQToolBox::Signal_CurrentChanged));;
}

QtObjectPtr QToolBox_CurrentWidget(QtObjectPtr ptr){
	return static_cast<QToolBox*>(ptr)->currentWidget();
}

int QToolBox_IndexOf(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QToolBox*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

int QToolBox_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr widget, QtObjectPtr icon, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), *static_cast<QIcon*>(icon), QString(text));
}

int QToolBox_InsertItem2(QtObjectPtr ptr, int index, QtObjectPtr widget, char* text){
	return static_cast<QToolBox*>(ptr)->insertItem(index, static_cast<QWidget*>(widget), QString(text));
}

int QToolBox_IsItemEnabled(QtObjectPtr ptr, int index){
	return static_cast<QToolBox*>(ptr)->isItemEnabled(index);
}

char* QToolBox_ItemText(QtObjectPtr ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemText(index).toUtf8().data();
}

char* QToolBox_ItemToolTip(QtObjectPtr ptr, int index){
	return static_cast<QToolBox*>(ptr)->itemToolTip(index).toUtf8().data();
}

void QToolBox_RemoveItem(QtObjectPtr ptr, int index){
	static_cast<QToolBox*>(ptr)->removeItem(index);
}

void QToolBox_SetCurrentWidget(QtObjectPtr ptr, QtObjectPtr widget){
	QMetaObject::invokeMethod(static_cast<QToolBox*>(ptr), "setCurrentWidget", Q_ARG(QWidget*, static_cast<QWidget*>(widget)));
}

void QToolBox_SetItemEnabled(QtObjectPtr ptr, int index, int enabled){
	static_cast<QToolBox*>(ptr)->setItemEnabled(index, enabled != 0);
}

void QToolBox_SetItemIcon(QtObjectPtr ptr, int index, QtObjectPtr icon){
	static_cast<QToolBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QToolBox_SetItemText(QtObjectPtr ptr, int index, char* text){
	static_cast<QToolBox*>(ptr)->setItemText(index, QString(text));
}

void QToolBox_SetItemToolTip(QtObjectPtr ptr, int index, char* toolTip){
	static_cast<QToolBox*>(ptr)->setItemToolTip(index, QString(toolTip));
}

QtObjectPtr QToolBox_Widget(QtObjectPtr ptr, int index){
	return static_cast<QToolBox*>(ptr)->widget(index);
}

void QToolBox_DestroyQToolBox(QtObjectPtr ptr){
	static_cast<QToolBox*>(ptr)->~QToolBox();
}

