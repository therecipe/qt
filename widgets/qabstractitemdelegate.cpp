#include "qabstractitemdelegate.h"
#include <QStyle>
#include <QStyleOptionViewItem>
#include <QString>
#include <QModelIndex>
#include <QEvent>
#include <QStyleOption>
#include <QAbstractItemModel>
#include <QVariant>
#include <QPainter>
#include <QAbstractItemView>
#include <QWidget>
#include <QHelpEvent>
#include <QObject>
#include <QUrl>
#include <QAbstractItemDelegate>
#include "_cgo_export.h"

class MyQAbstractItemDelegate: public QAbstractItemDelegate {
public:
void Signal_CloseEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint){callbackQAbstractItemDelegateCloseEditor(this->objectName().toUtf8().data(), editor, hint);};
void Signal_CommitData(QWidget * editor){callbackQAbstractItemDelegateCommitData(this->objectName().toUtf8().data(), editor);};
void Signal_SizeHintChanged(const QModelIndex & index){callbackQAbstractItemDelegateSizeHintChanged(this->objectName().toUtf8().data(), index.internalPointer());};
};

void QAbstractItemDelegate_ConnectCloseEditor(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_DisconnectCloseEditor(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&QAbstractItemDelegate::closeEditor), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *, QAbstractItemDelegate::EndEditHint)>(&MyQAbstractItemDelegate::Signal_CloseEditor));;
}

void QAbstractItemDelegate_ConnectCommitData(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

void QAbstractItemDelegate_DisconnectCommitData(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(QWidget *)>(&QAbstractItemDelegate::commitData), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(QWidget *)>(&MyQAbstractItemDelegate::Signal_CommitData));;
}

void* QAbstractItemDelegate_CreateEditor(void* ptr, void* parent, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyEditor(void* ptr, void* editor, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

int QAbstractItemDelegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index){
	return static_cast<QAbstractItemDelegate*>(ptr)->helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_Paint(void* ptr, void* painter, void* option, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetEditorData(void* ptr, void* editor, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_SetModelData(void* ptr, void* editor, void* model, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_ConnectSizeHintChanged(void* ptr){
	QObject::connect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_DisconnectSizeHintChanged(void* ptr){
	QObject::disconnect(static_cast<QAbstractItemDelegate*>(ptr), static_cast<void (QAbstractItemDelegate::*)(const QModelIndex &)>(&QAbstractItemDelegate::sizeHintChanged), static_cast<MyQAbstractItemDelegate*>(ptr), static_cast<void (MyQAbstractItemDelegate::*)(const QModelIndex &)>(&MyQAbstractItemDelegate::Signal_SizeHintChanged));;
}

void QAbstractItemDelegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index){
	static_cast<QAbstractItemDelegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void QAbstractItemDelegate_DestroyQAbstractItemDelegate(void* ptr){
	static_cast<QAbstractItemDelegate*>(ptr)->~QAbstractItemDelegate();
}

