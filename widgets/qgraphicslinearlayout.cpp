#include "qgraphicslinearlayout.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QRectF>
#include <QRect>
#include <QGraphicsLinearLayout>
#include "_cgo_export.h"

class MyQGraphicsLinearLayout: public QGraphicsLinearLayout {
public:
};

void* QGraphicsLinearLayout_NewQGraphicsLinearLayout(void* parent){
	return new QGraphicsLinearLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

void* QGraphicsLinearLayout_NewQGraphicsLinearLayout2(int orientation, void* parent){
	return new QGraphicsLinearLayout(static_cast<Qt::Orientation>(orientation), static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsLinearLayout_AddItem(void* ptr, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_AddStretch(void* ptr, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->addStretch(stretch);
}

int QGraphicsLinearLayout_Alignment(void* ptr, void* item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsLinearLayout_Count(void* ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->count();
}

void QGraphicsLinearLayout_InsertItem(void* ptr, int index, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertItem(index, static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_InsertStretch(void* ptr, int index, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertStretch(index, stretch);
}

void QGraphicsLinearLayout_Invalidate(void* ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->invalidate();
}

void* QGraphicsLinearLayout_ItemAt(void* ptr, int index){
	return static_cast<QGraphicsLinearLayout*>(ptr)->itemAt(index);
}

double QGraphicsLinearLayout_ItemSpacing(void* ptr, int index){
	return static_cast<double>(static_cast<QGraphicsLinearLayout*>(ptr)->itemSpacing(index));
}

int QGraphicsLinearLayout_Orientation(void* ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->orientation();
}

void QGraphicsLinearLayout_RemoveAt(void* ptr, int index){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeAt(index);
}

void QGraphicsLinearLayout_RemoveItem(void* ptr, void* item){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_SetAlignment(void* ptr, void* item, int alignment){
	static_cast<QGraphicsLinearLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsLinearLayout_SetGeometry(void* ptr, void* rect){
	static_cast<QGraphicsLinearLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsLinearLayout_SetItemSpacing(void* ptr, int index, double spacing){
	static_cast<QGraphicsLinearLayout*>(ptr)->setItemSpacing(index, static_cast<qreal>(spacing));
}

void QGraphicsLinearLayout_SetOrientation(void* ptr, int orientation){
	static_cast<QGraphicsLinearLayout*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QGraphicsLinearLayout_SetSpacing(void* ptr, double spacing){
	static_cast<QGraphicsLinearLayout*>(ptr)->setSpacing(static_cast<qreal>(spacing));
}

void QGraphicsLinearLayout_SetStretchFactor(void* ptr, void* item, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->setStretchFactor(static_cast<QGraphicsLayoutItem*>(item), stretch);
}

double QGraphicsLinearLayout_Spacing(void* ptr){
	return static_cast<double>(static_cast<QGraphicsLinearLayout*>(ptr)->spacing());
}

int QGraphicsLinearLayout_StretchFactor(void* ptr, void* item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->stretchFactor(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(void* ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->~QGraphicsLinearLayout();
}

