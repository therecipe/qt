#include "qboxlayout.h"
#include <QLayout>
#include <QUrl>
#include <QSpacerItem>
#include <QLayoutItem>
#include <QRect>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QBoxLayout>
#include "_cgo_export.h"

class MyQBoxLayout: public QBoxLayout {
public:
};

int QBoxLayout_Direction(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->direction();
}

void* QBoxLayout_NewQBoxLayout(int dir, void* parent){
	return new QBoxLayout(static_cast<QBoxLayout::Direction>(dir), static_cast<QWidget*>(parent));
}

void QBoxLayout_AddItem(void* ptr, void* item){
	static_cast<QBoxLayout*>(ptr)->addItem(static_cast<QLayoutItem*>(item));
}

void QBoxLayout_AddLayout(void* ptr, void* layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->addLayout(static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_AddSpacerItem(void* ptr, void* spacerItem){
	static_cast<QBoxLayout*>(ptr)->addSpacerItem(static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_AddSpacing(void* ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addSpacing(size);
}

void QBoxLayout_AddStretch(void* ptr, int stretch){
	static_cast<QBoxLayout*>(ptr)->addStretch(stretch);
}

void QBoxLayout_AddStrut(void* ptr, int size){
	static_cast<QBoxLayout*>(ptr)->addStrut(size);
}

void QBoxLayout_AddWidget(void* ptr, void* widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->addWidget(static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

int QBoxLayout_Count(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->count();
}

int QBoxLayout_ExpandingDirections(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->expandingDirections();
}

int QBoxLayout_HasHeightForWidth(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->hasHeightForWidth();
}

int QBoxLayout_HeightForWidth(void* ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->heightForWidth(w);
}

void QBoxLayout_InsertItem(void* ptr, int index, void* item){
	static_cast<QBoxLayout*>(ptr)->insertItem(index, static_cast<QLayoutItem*>(item));
}

void QBoxLayout_InsertLayout(void* ptr, int index, void* layout, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertLayout(index, static_cast<QLayout*>(layout), stretch);
}

void QBoxLayout_InsertSpacerItem(void* ptr, int index, void* spacerItem){
	static_cast<QBoxLayout*>(ptr)->insertSpacerItem(index, static_cast<QSpacerItem*>(spacerItem));
}

void QBoxLayout_InsertSpacing(void* ptr, int index, int size){
	static_cast<QBoxLayout*>(ptr)->insertSpacing(index, size);
}

void QBoxLayout_InsertStretch(void* ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->insertStretch(index, stretch);
}

void QBoxLayout_InsertWidget(void* ptr, int index, void* widget, int stretch, int alignment){
	static_cast<QBoxLayout*>(ptr)->insertWidget(index, static_cast<QWidget*>(widget), stretch, static_cast<Qt::AlignmentFlag>(alignment));
}

void QBoxLayout_Invalidate(void* ptr){
	static_cast<QBoxLayout*>(ptr)->invalidate();
}

void* QBoxLayout_ItemAt(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->itemAt(index);
}

int QBoxLayout_MinimumHeightForWidth(void* ptr, int w){
	return static_cast<QBoxLayout*>(ptr)->minimumHeightForWidth(w);
}

void QBoxLayout_SetDirection(void* ptr, int direction){
	static_cast<QBoxLayout*>(ptr)->setDirection(static_cast<QBoxLayout::Direction>(direction));
}

void QBoxLayout_SetGeometry(void* ptr, void* r){
	static_cast<QBoxLayout*>(ptr)->setGeometry(*static_cast<QRect*>(r));
}

void QBoxLayout_SetSpacing(void* ptr, int spacing){
	static_cast<QBoxLayout*>(ptr)->setSpacing(spacing);
}

void QBoxLayout_SetStretch(void* ptr, int index, int stretch){
	static_cast<QBoxLayout*>(ptr)->setStretch(index, stretch);
}

int QBoxLayout_SetStretchFactor2(void* ptr, void* layout, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QLayout*>(layout), stretch);
}

int QBoxLayout_SetStretchFactor(void* ptr, void* widget, int stretch){
	return static_cast<QBoxLayout*>(ptr)->setStretchFactor(static_cast<QWidget*>(widget), stretch);
}

int QBoxLayout_Spacing(void* ptr){
	return static_cast<QBoxLayout*>(ptr)->spacing();
}

int QBoxLayout_Stretch(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->stretch(index);
}

void* QBoxLayout_TakeAt(void* ptr, int index){
	return static_cast<QBoxLayout*>(ptr)->takeAt(index);
}

void QBoxLayout_DestroyQBoxLayout(void* ptr){
	static_cast<QBoxLayout*>(ptr)->~QBoxLayout();
}

