#include "qgraphicslinearlayout.h"
#include <QRect>
#include <QRectF>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGraphicsLayout>
#include <QGraphicsLayoutItem>
#include <QGraphicsLinearLayout>
#include "_cgo_export.h"

class MyQGraphicsLinearLayout: public QGraphicsLinearLayout {
public:
};

QtObjectPtr QGraphicsLinearLayout_NewQGraphicsLinearLayout(QtObjectPtr parent){
	return new QGraphicsLinearLayout(static_cast<QGraphicsLayoutItem*>(parent));
}

QtObjectPtr QGraphicsLinearLayout_NewQGraphicsLinearLayout2(int orientation, QtObjectPtr parent){
	return new QGraphicsLinearLayout(static_cast<Qt::Orientation>(orientation), static_cast<QGraphicsLayoutItem*>(parent));
}

void QGraphicsLinearLayout_AddItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsLinearLayout*>(ptr)->addItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_AddStretch(QtObjectPtr ptr, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->addStretch(stretch);
}

int QGraphicsLinearLayout_Alignment(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->alignment(static_cast<QGraphicsLayoutItem*>(item));
}

int QGraphicsLinearLayout_Count(QtObjectPtr ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->count();
}

void QGraphicsLinearLayout_InsertItem(QtObjectPtr ptr, int index, QtObjectPtr item){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertItem(index, static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_InsertStretch(QtObjectPtr ptr, int index, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->insertStretch(index, stretch);
}

void QGraphicsLinearLayout_Invalidate(QtObjectPtr ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->invalidate();
}

QtObjectPtr QGraphicsLinearLayout_ItemAt(QtObjectPtr ptr, int index){
	return static_cast<QGraphicsLinearLayout*>(ptr)->itemAt(index);
}

int QGraphicsLinearLayout_Orientation(QtObjectPtr ptr){
	return static_cast<QGraphicsLinearLayout*>(ptr)->orientation();
}

void QGraphicsLinearLayout_RemoveAt(QtObjectPtr ptr, int index){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeAt(index);
}

void QGraphicsLinearLayout_RemoveItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QGraphicsLinearLayout*>(ptr)->removeItem(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_SetAlignment(QtObjectPtr ptr, QtObjectPtr item, int alignment){
	static_cast<QGraphicsLinearLayout*>(ptr)->setAlignment(static_cast<QGraphicsLayoutItem*>(item), static_cast<Qt::AlignmentFlag>(alignment));
}

void QGraphicsLinearLayout_SetGeometry(QtObjectPtr ptr, QtObjectPtr rect){
	static_cast<QGraphicsLinearLayout*>(ptr)->setGeometry(*static_cast<QRectF*>(rect));
}

void QGraphicsLinearLayout_SetOrientation(QtObjectPtr ptr, int orientation){
	static_cast<QGraphicsLinearLayout*>(ptr)->setOrientation(static_cast<Qt::Orientation>(orientation));
}

void QGraphicsLinearLayout_SetStretchFactor(QtObjectPtr ptr, QtObjectPtr item, int stretch){
	static_cast<QGraphicsLinearLayout*>(ptr)->setStretchFactor(static_cast<QGraphicsLayoutItem*>(item), stretch);
}

int QGraphicsLinearLayout_StretchFactor(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QGraphicsLinearLayout*>(ptr)->stretchFactor(static_cast<QGraphicsLayoutItem*>(item));
}

void QGraphicsLinearLayout_DestroyQGraphicsLinearLayout(QtObjectPtr ptr){
	static_cast<QGraphicsLinearLayout*>(ptr)->~QGraphicsLinearLayout();
}

