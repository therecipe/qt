#include "qtreewidget.h"
#include <QItemSelectionModel>
#include <QWidget>
#include <QPoint>
#include <QItemSelection>
#include <QTreeWidgetItem>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemView>
#include <QMetaObject>
#include <QTreeWidget>
#include "_cgo_export.h"

class MyQTreeWidget: public QTreeWidget {
public:
void Signal_CurrentItemChanged(QTreeWidgetItem * current, QTreeWidgetItem * previous){callbackQTreeWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_ItemActivated(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemActivated(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemChanged(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemChanged(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemCollapsed(QTreeWidgetItem * item){callbackQTreeWidgetItemCollapsed(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemEntered(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemEntered(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemExpanded(QTreeWidgetItem * item){callbackQTreeWidgetItemExpanded(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QTreeWidgetItem * item, int column){callbackQTreeWidgetItemPressed(this->objectName().toUtf8().data(), item, column);};
void Signal_ItemSelectionChanged(){callbackQTreeWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

int QTreeWidget_ColumnCount(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->columnCount();
}

void QTreeWidget_SetColumnCount(QtObjectPtr ptr, int columns){
	static_cast<QTreeWidget*>(ptr)->setColumnCount(columns);
}

int QTreeWidget_TopLevelItemCount(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->topLevelItemCount();
}

QtObjectPtr QTreeWidget_NewQTreeWidget(QtObjectPtr parent){
	return new QTreeWidget(static_cast<QWidget*>(parent));
}

void QTreeWidget_AddTopLevelItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTreeWidget*>(ptr)->addTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "clear");
}

void QTreeWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item, int column){
	static_cast<QTreeWidget*>(ptr)->closePersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_CollapseItem(QtObjectPtr ptr, QtObjectPtr item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "collapseItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

int QTreeWidget_CurrentColumn(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->currentColumn();
}

QtObjectPtr QTreeWidget_CurrentItem(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->currentItem();
}

void QTreeWidget_ConnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&QTreeWidget::currentItemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, QTreeWidgetItem *)>(&MyQTreeWidget::Signal_CurrentItemChanged));;
}

void QTreeWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item, int column){
	static_cast<QTreeWidget*>(ptr)->editItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ExpandItem(QtObjectPtr ptr, QtObjectPtr item){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "expandItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)));
}

QtObjectPtr QTreeWidget_HeaderItem(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->headerItem();
}

int QTreeWidget_IndexOfTopLevelItem(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTreeWidget*>(ptr)->indexOfTopLevelItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_InsertTopLevelItem(QtObjectPtr ptr, int index, QtObjectPtr item){
	static_cast<QTreeWidget*>(ptr)->insertTopLevelItem(index, static_cast<QTreeWidgetItem*>(item));
}

QtObjectPtr QTreeWidget_InvisibleRootItem(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->invisibleRootItem();
}

int QTreeWidget_IsFirstItemColumnSpanned(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTreeWidget*>(ptr)->isFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item));
}

QtObjectPtr QTreeWidget_ItemAbove(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTreeWidget*>(ptr)->itemAbove(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

void QTreeWidget_DisconnectItemActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemActivated), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemActivated));;
}

QtObjectPtr QTreeWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr p){
	return static_cast<QTreeWidget*>(ptr)->itemAt(*static_cast<QPoint*>(p));
}

QtObjectPtr QTreeWidget_ItemAt2(QtObjectPtr ptr, int x, int y){
	return static_cast<QTreeWidget*>(ptr)->itemAt(x, y);
}

QtObjectPtr QTreeWidget_ItemBelow(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTreeWidget*>(ptr)->itemBelow(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_ConnectItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_DisconnectItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemChanged));;
}

void QTreeWidget_ConnectItemClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_DisconnectItemClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemClicked));;
}

void QTreeWidget_ConnectItemCollapsed(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_DisconnectItemCollapsed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemCollapsed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemCollapsed));;
}

void QTreeWidget_ConnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemDoubleClicked), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemDoubleClicked));;
}

void QTreeWidget_ConnectItemEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_DisconnectItemEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemEntered), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemEntered));;
}

void QTreeWidget_ConnectItemExpanded(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_DisconnectItemExpanded(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *)>(&QTreeWidget::itemExpanded), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *)>(&MyQTreeWidget::Signal_ItemExpanded));;
}

void QTreeWidget_ConnectItemPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_DisconnectItemPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)(QTreeWidgetItem *, int)>(&QTreeWidget::itemPressed), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)(QTreeWidgetItem *, int)>(&MyQTreeWidget::Signal_ItemPressed));;
}

void QTreeWidget_ConnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

void QTreeWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTreeWidget*>(ptr), static_cast<void (QTreeWidget::*)()>(&QTreeWidget::itemSelectionChanged), static_cast<MyQTreeWidget*>(ptr), static_cast<void (MyQTreeWidget::*)()>(&MyQTreeWidget::Signal_ItemSelectionChanged));;
}

QtObjectPtr QTreeWidget_ItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column){
	return static_cast<QTreeWidget*>(ptr)->itemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item, int column){
	static_cast<QTreeWidget*>(ptr)->openPersistentEditor(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_RemoveItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column){
	static_cast<QTreeWidget*>(ptr)->removeItemWidget(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint){
	QMetaObject::invokeMethod(static_cast<QTreeWidget*>(ptr), "scrollToItem", Q_ARG(QTreeWidgetItem*, static_cast<QTreeWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTreeWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int column){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column);
}

void QTreeWidget_SetCurrentItem3(QtObjectPtr ptr, QtObjectPtr item, int column, int command){
	static_cast<QTreeWidget*>(ptr)->setCurrentItem(static_cast<QTreeWidgetItem*>(item), column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTreeWidget_SetFirstItemColumnSpanned(QtObjectPtr ptr, QtObjectPtr item, int span){
	static_cast<QTreeWidget*>(ptr)->setFirstItemColumnSpanned(static_cast<QTreeWidgetItem*>(item), span != 0);
}

void QTreeWidget_SetHeaderItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTreeWidget*>(ptr)->setHeaderItem(static_cast<QTreeWidgetItem*>(item));
}

void QTreeWidget_SetHeaderLabel(QtObjectPtr ptr, char* label){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabel(QString(label));
}

void QTreeWidget_SetHeaderLabels(QtObjectPtr ptr, char* labels){
	static_cast<QTreeWidget*>(ptr)->setHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTreeWidget_SetItemWidget(QtObjectPtr ptr, QtObjectPtr item, int column, QtObjectPtr widget){
	static_cast<QTreeWidget*>(ptr)->setItemWidget(static_cast<QTreeWidgetItem*>(item), column, static_cast<QWidget*>(widget));
}

void QTreeWidget_SetSelectionModel(QtObjectPtr ptr, QtObjectPtr selectionModel){
	static_cast<QTreeWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

int QTreeWidget_SortColumn(QtObjectPtr ptr){
	return static_cast<QTreeWidget*>(ptr)->sortColumn();
}

void QTreeWidget_SortItems(QtObjectPtr ptr, int column, int order){
	static_cast<QTreeWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

QtObjectPtr QTreeWidget_TakeTopLevelItem(QtObjectPtr ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->takeTopLevelItem(index);
}

QtObjectPtr QTreeWidget_TopLevelItem(QtObjectPtr ptr, int index){
	return static_cast<QTreeWidget*>(ptr)->topLevelItem(index);
}

void QTreeWidget_DestroyQTreeWidget(QtObjectPtr ptr){
	static_cast<QTreeWidget*>(ptr)->~QTreeWidget();
}

