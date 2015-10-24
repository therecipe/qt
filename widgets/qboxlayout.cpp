#include "qboxlayout.h"
#include <QString>
#include <QUrl>
#include <QSpacerItem>
#include <QWidget>
#include <QLayout>
#include <QVariant>
#include <QModelIndex>
#include <QLayoutItem>
#include <QRect>
#include <QBoxLayout>
#include "_cgo_export.h"

class MyQBoxLayout: public QBoxLayout {
public:
};

int QBoxLayout_Direction(QtObjectPtr ptr){
	return static_cast<QBoxLayout*>(ptr)->direction();
}

QtObjectPtr QBoxLayout_NewQBoxLayout(int dir, QtObjectPtr parent){
	return new QBoxLayout(static_cast<QBoxLayout::Direction>(dir), static_cast<QWidget*>(parent));
}

void QBoxLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QBoxLayout_AddLayout(QtObjectPtr ptr, QtObjectPtr layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_AddSpacerItem(QtObjectPtr ptr, QtObjectPtr spacerItem){
	static_cast<QBoxLayout*>(ptr)->addSpacerItem(static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_AddSpacing(QtObjectPtr ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addSpacing(size);
}

void QBoxLayout_AddStretch(QtObjectPtr ptr, int stretch){
	static_cast<QBoxLayout*>(ptr)->addStretch(stretch);
}

void QBoxLayout_AddStrut(QtObjectPtr ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addStrut(size);
}

void QBoxLayout_AddWidget(QtObjectPtr ptr, QtObjectPtr widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

int QBoxLayout_Count(QtObjectPtr ptr){
	return static_cast<QBoxLayout*>(ptr)->count();
}

int QBoxLayout_ExpandingDirections(QtObjectPtr ptr){
	return static_cast<QBoxLayout*>(ptr)->expandingDirections();
}

int QBoxLayout_HasHeightForWidth(QtObjectPtr ptr){
	return static_cast<QBoxLayout*>(ptr)->hasHeightForWidth();
}

int QBoxLayout_HeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->heightForWidth(w);
}

void QBoxLayout_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr item){
	static_cast<QBoxLayout*>(ptr)->insertItem(index, static_cast<QLayoutItem*>(item));
}

void QBoxLayout_InsertLayout(QtObjectPtr ptr, int index, QtObjectPtr layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertLayout(index, static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_InsertSpacerItem(QtObjectPtr ptr, int index, QtObjectPtr spacerItem){
	static_cast<QBoxLayout*>(ptr)->insertSpacerItem(index, static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_InsertSpacing(QtObjectPtr ptr, int index, int size){
	static_cast<QBoxLayout*>(ptr)->insertSpacing(index, size);
}

void QBoxLayout_InsertStretch(QtObjectPtr ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertStretch(index, stretch);
}

void QBoxLayout_InsertWidget(QtObjectPtr ptr, int index, QtObjectPtr widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

void QBoxLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QBoxLayout*>(ptr)->invalidate();
}

QtObjectPtr QBoxLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->itemAt(index);
}

int QBoxLayout_MinimumHeightForWidth(QtObjectPtr ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->minimumHeightForWidth(w);
}

void QBoxLayout_SetDirection(QtObjectPtr ptr, int direction){
	static_cast<QBoxLayout*>(ptr)->setDirection(static_cast<QBoxLayout::Direction>(direction));
}

void QBoxLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr r){
	static_cast<QBoxLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QBoxLayout_SetSpacing(QtObjectPtr ptr, int spacing){
	static_cast<QBoxLayout*>(ptr)->setSpacing(spacing);
}

void QBoxLayout_SetStretch(QtObjectPtr ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->setStretch(index, stretch);
}

int QBoxLayout_SetStretchFactor2(QtObjectPtr ptr, QtObjectPtr layout, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QLayout*>(layout), stretch);
}

int QBoxLayout_SetStretchFactor(QtObjectPtr ptr, QtObjectPtr widget, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QWidget*>(widget), stretch);
}

int QBoxLayout_Spacing(QtObjectPtr ptr){
	return static_cast<QBoxLayout*>(ptr)->spacing();
}

int QBoxLayout_Stretch(QtObjectPtr ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->stretch(index);
}

QtObjectPtr QBoxLayout_TakeAt(QtObjectPtr ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->takeAt(index);
}

void QBoxLayout_DestroyQBoxLayout(QtObjectPtr ptr){
	static_cast<QBoxLayout*>(ptr)->~QBoxLayout();
}

