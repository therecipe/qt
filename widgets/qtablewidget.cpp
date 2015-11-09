#include "qtablewidget.h"
#include <QTableWidgetItem>
#include <QTableWidgetSelectionRange>
#include <QMetaObject>
#include <QItemSelectionModel>
#include <QAbstractItemView>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QItemSelection>
#include <QVariant>
#include <QWidget>
#include <QPoint>
#include <QTableWidget>
#include "_cgo_export.h"

class MyQTableWidget: public QTableWidget {
public:
void Signal_CellActivated(int row, int column){callbackQTableWidgetCellActivated(this->objectName().toUtf8().data(), row, column);};
void Signal_CellChanged(int row, int column){callbackQTableWidgetCellChanged(this->objectName().toUtf8().data(), row, column);};
void Signal_CellClicked(int row, int column){callbackQTableWidgetCellClicked(this->objectName().toUtf8().data(), row, column);};
void Signal_CellDoubleClicked(int row, int column){callbackQTableWidgetCellDoubleClicked(this->objectName().toUtf8().data(), row, column);};
void Signal_CellEntered(int row, int column){callbackQTableWidgetCellEntered(this->objectName().toUtf8().data(), row, column);};
void Signal_CellPressed(int row, int column){callbackQTableWidgetCellPressed(this->objectName().toUtf8().data(), row, column);};
void Signal_CurrentCellChanged(int currentRow, int currentColumn, int previousRow, int previousColumn){callbackQTableWidgetCurrentCellChanged(this->objectName().toUtf8().data(), currentRow, currentColumn, previousRow, previousColumn);};
void Signal_CurrentItemChanged(QTableWidgetItem * current, QTableWidgetItem * previous){callbackQTableWidgetCurrentItemChanged(this->objectName().toUtf8().data(), current, previous);};
void Signal_ItemActivated(QTableWidgetItem * item){callbackQTableWidgetItemActivated(this->objectName().toUtf8().data(), item);};
void Signal_ItemChanged(QTableWidgetItem * item){callbackQTableWidgetItemChanged(this->objectName().toUtf8().data(), item);};
void Signal_ItemClicked(QTableWidgetItem * item){callbackQTableWidgetItemClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemDoubleClicked(QTableWidgetItem * item){callbackQTableWidgetItemDoubleClicked(this->objectName().toUtf8().data(), item);};
void Signal_ItemEntered(QTableWidgetItem * item){callbackQTableWidgetItemEntered(this->objectName().toUtf8().data(), item);};
void Signal_ItemPressed(QTableWidgetItem * item){callbackQTableWidgetItemPressed(this->objectName().toUtf8().data(), item);};
void Signal_ItemSelectionChanged(){callbackQTableWidgetItemSelectionChanged(this->objectName().toUtf8().data());};
};

void* QTableWidget_ItemAt(void* ptr, void* point){
	return static_cast<QTableWidget*>(ptr)->itemAt(*static_cast<QPoint*>(point));
}

void QTableWidget_ConnectCellActivated(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_DisconnectCellActivated(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_ConnectCellChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_DisconnectCellChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_ConnectCellClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_DisconnectCellClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_ConnectCellDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_DisconnectCellDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_ConnectCellEntered(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_DisconnectCellEntered(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_ConnectCellPressed(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

void QTableWidget_DisconnectCellPressed(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

void* QTableWidget_CellWidget(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->cellWidget(row, column);
}

void QTableWidget_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clear");
}

void QTableWidget_ClearContents(void* ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clearContents");
}

void QTableWidget_ClosePersistentEditor(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->closePersistentEditor(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_Column(void* ptr, void* item){
	return static_cast<QTableWidget*>(ptr)->column(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_ColumnCount(void* ptr){
	return static_cast<QTableWidget*>(ptr)->columnCount();
}

void QTableWidget_ConnectCurrentCellChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

void QTableWidget_DisconnectCurrentCellChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

int QTableWidget_CurrentColumn(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentColumn();
}

void* QTableWidget_CurrentItem(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentItem();
}

void QTableWidget_ConnectCurrentItemChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

void QTableWidget_DisconnectCurrentItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

int QTableWidget_CurrentRow(void* ptr){
	return static_cast<QTableWidget*>(ptr)->currentRow();
}

void QTableWidget_EditItem(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->editItem(static_cast<QTableWidgetItem*>(item));
}

void* QTableWidget_HorizontalHeaderItem(void* ptr, int column){
	return static_cast<QTableWidget*>(ptr)->horizontalHeaderItem(column);
}

void QTableWidget_InsertColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertColumn", Q_ARG(int, column));
}

void QTableWidget_InsertRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertRow", Q_ARG(int, row));
}

void* QTableWidget_Item(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->item(row, column);
}

void QTableWidget_ConnectItemActivated(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

void QTableWidget_DisconnectItemActivated(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

void* QTableWidget_ItemAt2(void* ptr, int ax, int ay){
	return static_cast<QTableWidget*>(ptr)->itemAt(ax, ay);
}

void QTableWidget_ConnectItemChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_DisconnectItemChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_ConnectItemClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_DisconnectItemClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_ConnectItemDoubleClicked(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_DisconnectItemDoubleClicked(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_ConnectItemEntered(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_DisconnectItemEntered(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_ConnectItemPressed(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

void QTableWidget_DisconnectItemPressed(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

void* QTableWidget_ItemPrototype(void* ptr){
	return const_cast<QTableWidgetItem*>(static_cast<QTableWidget*>(ptr)->itemPrototype());
}

void QTableWidget_ConnectItemSelectionChanged(void* ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_DisconnectItemSelectionChanged(void* ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_OpenPersistentEditor(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->openPersistentEditor(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_RemoveCellWidget(void* ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->removeCellWidget(row, column);
}

void QTableWidget_RemoveColumn(void* ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeColumn", Q_ARG(int, column));
}

void QTableWidget_RemoveRow(void* ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeRow", Q_ARG(int, row));
}

int QTableWidget_Row(void* ptr, void* item){
	return static_cast<QTableWidget*>(ptr)->row(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_RowCount(void* ptr){
	return static_cast<QTableWidget*>(ptr)->rowCount();
}

void QTableWidget_ScrollToItem(void* ptr, void* item, int hint){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "scrollToItem", Q_ARG(QTableWidgetItem*, static_cast<QTableWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTableWidget_SetCellWidget(void* ptr, int row, int column, void* widget){
	static_cast<QTableWidget*>(ptr)->setCellWidget(row, column, static_cast<QWidget*>(widget));
}

void QTableWidget_SetColumnCount(void* ptr, int columns){
	static_cast<QTableWidget*>(ptr)->setColumnCount(columns);
}

void QTableWidget_SetCurrentCell(void* ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column);
}

void QTableWidget_SetCurrentCell2(void* ptr, int row, int column, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetCurrentItem(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetCurrentItem2(void* ptr, void* item, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetHorizontalHeaderItem(void* ptr, int column, void* item){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderItem(column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetHorizontalHeaderLabels(void* ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SetItem(void* ptr, int row, int column, void* item){
	static_cast<QTableWidget*>(ptr)->setItem(row, column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetItemPrototype(void* ptr, void* item){
	static_cast<QTableWidget*>(ptr)->setItemPrototype(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetRangeSelected(void* ptr, void* ran, int sele){
	static_cast<QTableWidget*>(ptr)->setRangeSelected(*static_cast<QTableWidgetSelectionRange*>(ran), sele != 0);
}

void QTableWidget_SetRowCount(void* ptr, int rows){
	static_cast<QTableWidget*>(ptr)->setRowCount(rows);
}

void QTableWidget_SetVerticalHeaderItem(void* ptr, int row, void* item){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderItem(row, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetVerticalHeaderLabels(void* ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SortItems(void* ptr, int column, int order){
	static_cast<QTableWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

void* QTableWidget_TakeHorizontalHeaderItem(void* ptr, int column){
	return static_cast<QTableWidget*>(ptr)->takeHorizontalHeaderItem(column);
}

void* QTableWidget_TakeItem(void* ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->takeItem(row, column);
}

void* QTableWidget_TakeVerticalHeaderItem(void* ptr, int row){
	return static_cast<QTableWidget*>(ptr)->takeVerticalHeaderItem(row);
}

void* QTableWidget_VerticalHeaderItem(void* ptr, int row){
	return static_cast<QTableWidget*>(ptr)->verticalHeaderItem(row);
}

int QTableWidget_VisualColumn(void* ptr, int logicalColumn){
	return static_cast<QTableWidget*>(ptr)->visualColumn(logicalColumn);
}

int QTableWidget_VisualRow(void* ptr, int logicalRow){
	return static_cast<QTableWidget*>(ptr)->visualRow(logicalRow);
}

void QTableWidget_DestroyQTableWidget(void* ptr){
	static_cast<QTableWidget*>(ptr)->~QTableWidget();
}

