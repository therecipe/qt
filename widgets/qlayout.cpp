#include "qlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRect>
#include <QLayoutItem>
#include <QMargins>
#include <QWidget>
#include <QLayout>
#include "_cgo_export.h"

class MyQLayout: public QLayout {
public:
};

void QLayout_SetSizeConstraint(void* ptr, int v){
	static_cast<QLayout*>(ptr)->setSizeConstraint(static_cast<QLayout::SizeConstraint>(v));
}

void QLayout_SetSpacing(void* ptr, int v){
	static_cast<QLayout*>(ptr)->setSpacing(v);
}

int QLayout_SizeConstraint(void* ptr){
	return static_cast<QLayout*>(ptr)->sizeConstraint();
}

int QLayout_Spacing(void* ptr){
	return static_cast<QLayout*>(ptr)->spacing();
}

int QLayout_Activate(void* ptr){
	return static_cast<QLayout*>(ptr)->activate();
}

void QLayout_AddItem(void* ptr, void* item){
	static_cast<QLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QLayout_AddWidget(void* ptr, void* w){
	static_cast<QLayout*>(ptr)->addWidget(static_cast<QWidget*>(w));
}

int QLayout_ControlTypes(void* ptr){
	return static_cast<QLayout*>(ptr)->controlTypes();
}

int QLayout_Count(void* ptr){
	return static_cast<QLayout*>(ptr)->count();
}

int QLayout_ExpandingDirections(void* ptr){
	return static_cast<QLayout*>(ptr)->expandingDirections();
}

void QLayout_GetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->getContentsMargins(&left, &top, &right, &bottom);
}

int QLayout_IndexOf(void* ptr, void* widget){
	return static_cast<QLayout*>(ptr)->indexOf(static_cast<QWidget*>(widget));
}

void QLayout_Invalidate(void* ptr){
	static_cast<QLayout*>(ptr)->invalidate();
}

int QLayout_IsEmpty(void* ptr){
	return static_cast<QLayout*>(ptr)->isEmpty();
}

int QLayout_IsEnabled(void* ptr){
	return static_cast<QLayout*>(ptr)->isEnabled();
}

void* QLayout_ItemAt(void* ptr, int index){
	return static_cast<QLayout*>(ptr)->itemAt(index);
}

void* QLayout_Layout(void* ptr){
	return static_cast<QLayout*>(ptr)->layout();
}

void* QLayout_MenuBar(void* ptr){
	return static_cast<QLayout*>(ptr)->menuBar();
}

void* QLayout_ParentWidget(void* ptr){
	return static_cast<QLayout*>(ptr)->parentWidget();
}

void QLayout_RemoveItem(void* ptr, void* item){
	static_cast<QLayout*>(ptr)->removeItem(static_cast<QLayoutItem*>(item));
}

void QLayout_RemoveWidget(void* ptr, void* widget){
	static_cast<QLayout*>(ptr)->removeWidget(static_cast<QWidget*>(widget));
}

void* QLayout_ReplaceWidget(void* ptr, void* from, void* to, int options){
	return static_cast<QLayout*>(ptr)->replaceWidget(static_cast<QWidget*>(from), static_cast<QWidget*>(to), static_cast<Qt::FindChildOption>(options));
}

int QLayout_SetAlignment2(void* ptr, void* l, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QLayout*>(l), static_cast<Qt::AlignmentFlag>(alignment));
}

int QLayout_SetAlignment(void* ptr, void* w, int alignment){
	return static_cast<QLayout*>(ptr)->setAlignment(static_cast<QWidget*>(w), static_cast<Qt::AlignmentFlag>(alignment));
}

void QLayout_SetContentsMargins2(void* ptr, void* margins){
	static_cast<QLayout*>(ptr)->setContentsMargins(*static_cast<QMargins*>(margins));
}

void QLayout_SetContentsMargins(void* ptr, int left, int top, int right, int bottom){
	static_cast<QLayout*>(ptr)->setContentsMargins(left, top, right, bottom);
}

void QLayout_SetEnabled(void* ptr, int enable){
	static_cast<QLayout*>(ptr)->setEnabled(enable != 0);
}

void QLayout_SetGeometry(void* ptr, void* r){
	static_cast<QLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QLayout_SetMenuBar(void* ptr, void* widget){
	static_cast<QLayout*>(ptr)->setMenuBar(static_cast<QWidget*>(widget));
}

void* QLayout_TakeAt(void* ptr, int index){
	return static_cast<QLayout*>(ptr)->takeAt(index);
}

void QLayout_Update(void* ptr){
	static_cast<QLayout*>(ptr)->update();
}

