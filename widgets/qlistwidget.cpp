#include "qlistwidget.h"
#include <QList>
#include <QDropEvent>
#include <QList>
#include <QListWidgetItem>
#include <QModelIndex>
#include <QItemSelection>
#include <QWidget>
#include <QString>
#include <QPoint>
#include <QAbstractItemView>
#include <QUrl>
#include <QMetaObject>
#include <QObject>
#include <QVariant>
#include <QItemSelectionModel>
#include <QListWidget>
#include "_cgo_export.h"

class MyQListWidget: public QListWidget {
public:
void Signal_CurrentItemChanged(QListWidgetItem * current, QListWidgetItem * previous){callbackQListWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_CurrentRowChanged(int currentRow){callbackQListWidgetCurrentRowChanged(this->objectName().toUtf8().data(), currentRow);};
void Signal_CurrentTextChanged(const QString & currentText){callbackQListWidgetCurrentTextChanged(this->objectName().toUtf8().data(), currentText.toUtf8().data());};
void Signal_ItemActivated(QListWidgetItem * item){callbackQListWidgetItemActivated(this->objectName().toUtf8().data(), item);};
void Signal_ItemChanged(QListWidgetItem * item){callbackQListWidgetItemChanged(this->objectName().toUtf8().data(), item);};
void Signal_ItemClicked(QListWidgetItem * item){callbackQListWidgetItemClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QListWidgetItem * item){callbackQListWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemEntered(QListWidgetItem * item){callbackQListWidgetItemEntered(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QListWidgetItem * item){callbackQListWidgetItemPressed(this->objectName().toUtf8().data(), item);};
void Signal_ItemSelectionChanged(){callbackQListWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QListWidget_Count(void* ptr){
	return static_cast<QListWidget*>(ptr)->count();
}

int QListWidget_CurrentRow(void* ptr){
	return static_cast<QListWidget*>(ptr)->currentRow();
}

int QListWidget_IsSortingEnabled(void* ptr){
	return static_cast<QListWidget*>(ptr)->isSortingEnabled();
}

void QListWidget_SetCurrentRow(void* ptr, int row){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row);
}

void QListWidget_SetSortingEnabled(void* ptr, int enable){
	static_cast<QListWidget*>(ptr)->setSortingEnabled(enable != 0);
}

void* QListWidget_NewQListWidget(void* parent){
	return new QListWidget(static_cast<QWidget*>(parent));
}

void QListWidget_AddItem2(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->addItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_AddItem(void* ptr, char* label){
	static_cast<QListWidget*>(ptr)->addItem(QString(label));
}

void QListWidget_AddItems(void* ptr, char* labels){
	static_cast<QListWidget*>(ptr)->addItems(QString(labels).split("|", QString::SkipEmptyParts));
}

void QListWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "clear");
}

void QListWidget_ClosePersistentEditor(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->closePersistentEditor(static_cast<QListWidgetItem*>(item));
}

void* QListWidget_CurrentItem(void* ptr){
	return static_cast<QListWidget*>(ptr)->currentItem();
}

void QListWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&QListWidget::currentItemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *, QListWidgetItem *)>(&MyQListWidget::Signal_CurrentItemChanged));;
}

void QListWidget_ConnectCurrentRowChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_DisconnectCurrentRowChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(int)>(&QListWidget::currentRowChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(int)>(&MyQListWidget::Signal_CurrentRowChanged));;
}

void QListWidget_ConnectCurrentTextChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DisconnectCurrentTextChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(const QString &)>(&QListWidget::currentTextChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(const QString &)>(&MyQListWidget::Signal_CurrentTextChanged));;
}

void QListWidget_DropEvent(void* ptr, void* event){
	static_cast<QListWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QListWidget_EditItem(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->editItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem(void* ptr, int row, void* item){
	static_cast<QListWidget*>(ptr)->insertItem(row, static_cast<QListWidgetItem*>(item));
}

void QListWidget_InsertItem2(void* ptr, int row, char* label){
	static_cast<QListWidget*>(ptr)->insertItem(row, QString(label));
}

void QListWidget_InsertItems(void* ptr, int row, char* labels){
	static_cast<QListWidget*>(ptr)->insertItems(row, QString(labels).split("|", QString::SkipEmptyParts));
}

void* QListWidget_Item(void* ptr, int row){
	return static_cast<QListWidget*>(ptr)->item(row);
}

void QListWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

void QListWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemActivated), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemActivated));;
}

void* QListWidget_ItemAt(void* ptr, void* p){
	return static_cast<QListWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

void* QListWidget_ItemAt2(void* ptr, int x, int y){
	return static_cast<QListWidget*>(ptr)->itemAt(x, y);
}

void QListWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemChanged));;
}

void QListWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemClicked));;
}

void QListWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemDoubleClicked), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemDoubleClicked));;
}

void QListWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemEntered), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemEntered));;
}

void QListWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)(QListWidgetItem *)>(&QListWidget::itemPressed), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)(QListWidgetItem *)>(&MyQListWidget::Signal_ItemPressed));;
}

void QListWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

void QListWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QListWidget*>(ptr), static_cast<void (QListWidget::*)()>(&QListWidget::itemSelectionChanged), static_cast<MyQListWidget*>(ptr), static_cast<void (MyQListWidget::*)()>(&MyQListWidget::Signal_ItemSelectionChanged));;
}

void* QListWidget_ItemWidget(void* ptr, void* item){
	return static_cast<QListWidget*>(ptr)->itemWidget(static_cast<QListWidgetItem*>(item));
}

void QListWidget_OpenPersistentEditor(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->openPersistentEditor(static_cast<QListWidgetItem*>(item));
}

void QListWidget_RemoveItemWidget(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->removeItemWidget(static_cast<QListWidgetItem*>(item));
}

int QListWidget_Row(void* ptr, void* item){
	return static_cast<QListWidget*>(ptr)->row(static_cast<QListWidgetItem*>(item));
}

void QListWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QListWidget*>(ptr), "scrollToItem", Q_ARG(QListWidgetItem*, static_cast<QListWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QListWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item));
}

void QListWidget_SetCurrentItem2(void* ptr, void* item, int command){
	static_cast<QListWidget*>(ptr)->setCurrentItem(static_cast<QListWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetCurrentRow2(void* ptr, int row, int command){
	static_cast<QListWidget*>(ptr)->setCurrentRow(row, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QListWidget_SetItemWidget(void* ptr, void* item, void* widget){
	static_cast<QListWidget*>(ptr)->setItemWidget(static_cast<QListWidgetItem*>(item), static_cast<QWidget*>(widget));
}

void QListWidget_SortItems(void* ptr, int order){
	static_cast<QListWidget*>(ptr)->sortItems(static_cast<Qt::SortOrder>(order));
}

void* QListWidget_TakeItem(void* ptr, int row){
	return static_cast<QListWidget*>(ptr)->takeItem(row);
}

void QListWidget_DestroyQListWidget(void* ptr){
	static_cast<QListWidget*>(ptr)->~QListWidget();
}

