#include "qlayout.h"
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMargins>
#include <QLayoutItem>
#include <QLayout>
#include "_cgo_export.h"

class MyQLayout: public QLayout {
public:
};

void QLayout_SetSizeConstraint(QtObjectPtr ptr, int v){
	static_cast<QLayout*>(ptr)->setSizeConstraint(static_cast<QLayout::SizeConstraint>(v));
}

void QLayout_SetSpacing(QtObjectPtr ptr, int v){
	static_cast<QLayout*>(ptr)->setSpacing(v);
}

int QLayout_SizeConstraint(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->sizeConstraint();
}

int QLayout_Spacing(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->spacing();
}

int QLayout_Activate(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->activate();
}

void QLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr w){
	static_cast<QLayout*>(ptr)->addWidget(static_cast<QWidget*>(w));
}

int QLayout_ControlTypes(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->controlTypes();
}

int QLayout_Count(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->count();
}

int QLayout_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->expandingDirections();
}

void QLayout_GetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

int QLayout_IndexOf(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QLayout*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

void QLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QLayout*>(ptr)->invalidate();
}

int QLayout_IsEmpty(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->isEmpty();
}

int QLayout_IsEnabled(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->isEnabled();
}

QtObjectPtr QLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QLayout*>(ptr)->itemAt(index);
}

QtObjectPtr QLayout_Layout(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->layout();
}

QtObjectPtr QLayout_MenuBar(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->menuBar();
}

QtObjectPtr QLayout_ParentWidget(QtObjectPtr ptr){
	return static_cast<QLayout*>(ptr)->parentWidget();
}

void QLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QLayout*>(ptr)->removeItem(static_cast<QLayoutItem*>(item));
}

void QLayout_RemoveWidget(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QLayout*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

QtObjectPtr QLayout_ReplaceWidget(QtObjectPtr ptr, QtObjectPtr from, QtObjectPtr to, int options){
	return static_cast<QLayout*>(ptr)->replaceWidget(static_cast<QWidget*>(from), static_cast<QWidget*>(to), static_cast<Qt::FindChildOption>(options));
}

int QLayout_SetAlignment2(QtObjectPtr ptr, QtObjectPtr l, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QLayout*>(l), static_cast<Qt::AlignmentFlag>(alignment));
}

int QLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr w, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QWidget*>(w), static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayout_SetContentsMargins2(QtObjectPtr ptr, QtObjectPtr margins){
	static_cast<QLayout*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QLayout_SetContentsMargins(QtObjectPtr ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QLayout_SetEnabled(QtObjectPtr ptr, int enable){
	static_cast<QLayout*>(ptr)->setEnabled(enable != 0);
}

void QLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QLayout_SetMenuBar(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QLayout*>(ptr)->setMenuBar(static_cast<QWidget*>(widget));
}

QtObjectPtr QLayout_TakeAt(QtObjectPtr ptr, int index){
	return static_cast<QLayout*>(ptr)->takeAt(index);
}

void QLayout_Update(QtObjectPtr ptr){
	static_cast<QLayout*>(ptr)->update();
}

