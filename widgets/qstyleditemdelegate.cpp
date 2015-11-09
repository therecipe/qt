#include "qstyleditemdelegate.h"
#include <QVariant>
#include <QUrl>
#include <QPainter>
#include <QStyleOption>
#include <QString>
#include <QStyleOptionViewItem>
#include <QWidget>
#include <QItemEditorFactory>
#include <QAbstractItemModel>
#include <QObject>
#include <QLocale>
#include <QStyle>
#include <QModelIndex>
#include <QStyledItemDelegate>
#include "_cgo_export.h"

class MyQStyledItemDelegate: public QStyledItemDelegate {
public:
};

void* QStyledItemDelegate_NewQStyledItemDelegate(void* parent){
	return new QStyledItemDelegate(static_cast<QObject*>(parent));
}

void* QStyledItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QStyledItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char* QStyledItemDelegate_DisplayText(void* ptr, void* value, void* locale){
	return static_cast<QStyledItemDelegate*>(ptr)->displayText(*static_cast<QVariant*>(value), *static_cast<QLocale*>(locale)).toUtf8().data();
}

void* QStyledItemDelegate_ItemEditorFactory(void* ptr){
	return static_cast<QStyledItemDelegate*>(ptr)->itemEditorFactory();
}

void QStyledItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_SetItemEditorFactory(void* ptr, void* factory){
	static_cast<QStyledItemDelegate*>(ptr)->setItemEditorFactory(static_cast<QItemEditorFactory*>(factory));
}

void QStyledItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QStyledItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QStyledItemDelegate_DestroyQStyledItemDelegate(void* ptr){
	static_cast<QStyledItemDelegate*>(ptr)->~QStyledItemDelegate();
}

