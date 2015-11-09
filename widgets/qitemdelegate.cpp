#include "qitemdelegate.h"
#include <QItemEditorFactory>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QAbstractItemModel>
#include <QObject>
#include <QPainter>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionViewItem>
#include <QItemDelegate>
#include "_cgo_export.h"

class MyQItemDelegate: public QItemDelegate {
public:
};

int QItemDelegate_HasClipping(void* ptr){
	return static_cast<QItemDelegate*>(ptr)->hasClipping();
}

void QItemDelegate_SetClipping(void* ptr, int clip){
	static_cast<QItemDelegate*>(ptr)->setClipping(clip != 0);
}

void* QItemDelegate_NewQItemDelegate(void* parent){
	return new QItemDelegate(static_cast<QObject*>(parent));
}

void* QItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void* QItemDelegate_ItemEditorFactory(void* ptr){
	return static_cast<QItemDelegate*>(ptr)->itemEditorFactory();
}

void QItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_SetItemEditorFactory(void* ptr, void* factory){
	static_cast<QItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QItemDelegate_DestroyQItemDelegate(void* ptr){
	static_cast<QItemDelegate*>(ptr)->~QItemDelegate();
}

