#include "qtablewidget.h"
#include <QVariant>
#include <QUrl>
#include <QItemSelectionModel>
#include <QObject>
#include <QPoint>
#include <QItemSelection>
#include <QString>
#include <QModelIndex>
#include <QTableWidgetSelectionRange>
#include <QAbstractItemView>
#include <QWidget>
#include <QMetaObject>
#include <QTableWidgetItem>
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

QtObjectPtr QTableWidget_ItemAt(QtObjectPtr ptr, QtObjectPtr point){
	return static_cast<QTableWidget*>(ptr)->itemAt(*static_cast<QPoint*>(point));
}

void QTableWidget_ConnectCellActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_DisconnectCellActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellActivated));;
}

void QTableWidget_ConnectCellChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_DisconnectCellChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellChanged));;
}

void QTableWidget_ConnectCellClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_DisconnectCellClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellClicked));;
}

void QTableWidget_ConnectCellDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_DisconnectCellDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellDoubleClicked));;
}

void QTableWidget_ConnectCellEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_DisconnectCellEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellEntered));;
}

void QTableWidget_ConnectCellPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

void QTableWidget_DisconnectCellPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int)>(&QTableWidget::cellPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int)>(&MyQTableWidget::Signal_CellPressed));;
}

QtObjectPtr QTableWidget_CellWidget(QtObjectPtr ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->cellWidget(row, column);
}

void QTableWidget_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clear");
}

void QTableWidget_ClearContents(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "clearContents");
}

void QTableWidget_ClosePersistentEditor(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->closePersistentEditor(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_Column(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTableWidget*>(ptr)->column(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_ColumnCount(QtObjectPtr ptr){
	return static_cast<QTableWidget*>(ptr)->columnCount();
}

void QTableWidget_ConnectCurrentCellChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

void QTableWidget_DisconnectCurrentCellChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(int, int, int, int)>(&QTableWidget::currentCellChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(int, int, int, int)>(&MyQTableWidget::Signal_CurrentCellChanged));;
}

int QTableWidget_CurrentColumn(QtObjectPtr ptr){
	return static_cast<QTableWidget*>(ptr)->currentColumn();
}

QtObjectPtr QTableWidget_CurrentItem(QtObjectPtr ptr){
	return static_cast<QTableWidget*>(ptr)->currentItem();
}

void QTableWidget_ConnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

void QTableWidget_DisconnectCurrentItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&QTableWidget::currentItemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *, QTableWidgetItem *)>(&MyQTableWidget::Signal_CurrentItemChanged));;
}

int QTableWidget_CurrentRow(QtObjectPtr ptr){
	return static_cast<QTableWidget*>(ptr)->currentRow();
}

void QTableWidget_EditItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->editItem(static_cast<QTableWidgetItem*>(item));
}

QtObjectPtr QTableWidget_HorizontalHeaderItem(QtObjectPtr ptr, int column){
	return static_cast<QTableWidget*>(ptr)->horizontalHeaderItem(column);
}

void QTableWidget_InsertColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertColumn", Q_ARG(int, column));
}

void QTableWidget_InsertRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "insertRow", Q_ARG(int, row));
}

QtObjectPtr QTableWidget_Item(QtObjectPtr ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->item(row, column);
}

void QTableWidget_ConnectItemActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

void QTableWidget_DisconnectItemActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemActivated), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemActivated));;
}

QtObjectPtr QTableWidget_ItemAt2(QtObjectPtr ptr, int ax, int ay){
	return static_cast<QTableWidget*>(ptr)->itemAt(ax, ay);
}

void QTableWidget_ConnectItemChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_DisconnectItemChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemChanged));;
}

void QTableWidget_ConnectItemClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_DisconnectItemClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemClicked));;
}

void QTableWidget_ConnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_DisconnectItemDoubleClicked(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemDoubleClicked), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemDoubleClicked));;
}

void QTableWidget_ConnectItemEntered(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_DisconnectItemEntered(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemEntered), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemEntered));;
}

void QTableWidget_ConnectItemPressed(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

void QTableWidget_DisconnectItemPressed(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)(QTableWidgetItem *)>(&QTableWidget::itemPressed), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)(QTableWidgetItem *)>(&MyQTableWidget::Signal_ItemPressed));;
}

QtObjectPtr QTableWidget_ItemPrototype(QtObjectPtr ptr){
	return const_cast<QTableWidgetItem*>(static_cast<QTableWidget*>(ptr)->itemPrototype());
}

void QTableWidget_ConnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_DisconnectItemSelectionChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QTableWidget*>(ptr), static_cast<void (QTableWidget::*)()>(&QTableWidget::itemSelectionChanged), static_cast<MyQTableWidget*>(ptr), static_cast<void (MyQTableWidget::*)()>(&MyQTableWidget::Signal_ItemSelectionChanged));;
}

void QTableWidget_OpenPersistentEditor(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->openPersistentEditor(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_RemoveCellWidget(QtObjectPtr ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->removeCellWidget(row, column);
}

void QTableWidget_RemoveColumn(QtObjectPtr ptr, int column){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeColumn", Q_ARG(int, column));
}

void QTableWidget_RemoveRow(QtObjectPtr ptr, int row){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "removeRow", Q_ARG(int, row));
}

int QTableWidget_Row(QtObjectPtr ptr, QtObjectPtr item){
	return static_cast<QTableWidget*>(ptr)->row(static_cast<QTableWidgetItem*>(item));
}

int QTableWidget_RowCount(QtObjectPtr ptr){
	return static_cast<QTableWidget*>(ptr)->rowCount();
}

void QTableWidget_ScrollToItem(QtObjectPtr ptr, QtObjectPtr item, int hint){
	QMetaObject::invokeMethod(static_cast<QTableWidget*>(ptr), "scrollToItem", Q_ARG(QTableWidgetItem*, static_cast<QTableWidgetItem*>(item)), Q_ARG(QAbstractItemView::ScrollHint, static_cast<QAbstractItemView::ScrollHint>(hint)));
}

void QTableWidget_SetCellWidget(QtObjectPtr ptr, int row, int column, QtObjectPtr widget){
	static_cast<QTableWidget*>(ptr)->setCellWidget(row, column, static_cast<QWidget*>(widget));
}

void QTableWidget_SetColumnCount(QtObjectPtr ptr, int columns){
	static_cast<QTableWidget*>(ptr)->setColumnCount(columns);
}

void QTableWidget_SetCurrentCell(QtObjectPtr ptr, int row, int column){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column);
}

void QTableWidget_SetCurrentCell2(QtObjectPtr ptr, int row, int column, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentCell(row, column, static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetCurrentItem(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetCurrentItem2(QtObjectPtr ptr, QtObjectPtr item, int command){
	static_cast<QTableWidget*>(ptr)->setCurrentItem(static_cast<QTableWidgetItem*>(item), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QTableWidget_SetHorizontalHeaderItem(QtObjectPtr ptr, int column, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderItem(column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetHorizontalHeaderLabels(QtObjectPtr ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SetItem(QtObjectPtr ptr, int row, int column, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->setItem(row, column, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetItemPrototype(QtObjectPtr ptr, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->setItemPrototype(static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetRangeSelected(QtObjectPtr ptr, QtObjectPtr ran, int sele){
	static_cast<QTableWidget*>(ptr)->setRangeSelected(*static_cast<QTableWidgetSelectionRange*>(ran), sele != 0);
}

void QTableWidget_SetRowCount(QtObjectPtr ptr, int rows){
	static_cast<QTableWidget*>(ptr)->setRowCount(rows);
}

void QTableWidget_SetVerticalHeaderItem(QtObjectPtr ptr, int row, QtObjectPtr item){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderItem(row, static_cast<QTableWidgetItem*>(item));
}

void QTableWidget_SetVerticalHeaderLabels(QtObjectPtr ptr, char* labels){
	static_cast<QTableWidget*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QTableWidget_SortItems(QtObjectPtr ptr, int column, int order){
	static_cast<QTableWidget*>(ptr)->sortItems(column, static_cast<Qt::SortOrder>(order));
}

QtObjectPtr QTableWidget_TakeHorizontalHeaderItem(QtObjectPtr ptr, int column){
	return static_cast<QTableWidget*>(ptr)->takeHorizontalHeaderItem(column);
}

QtObjectPtr QTableWidget_TakeItem(QtObjectPtr ptr, int row, int column){
	return static_cast<QTableWidget*>(ptr)->takeItem(row, column);
}

QtObjectPtr QTableWidget_TakeVerticalHeaderItem(QtObjectPtr ptr, int row){
	return static_cast<QTableWidget*>(ptr)->takeVerticalHeaderItem(row);
}

QtObjectPtr QTableWidget_VerticalHeaderItem(QtObjectPtr ptr, int row){
	return static_cast<QTableWidget*>(ptr)->verticalHeaderItem(row);
}

int QTableWidget_VisualColumn(QtObjectPtr ptr, int logicalColumn){
	return static_cast<QTableWidget*>(ptr)->visualColumn(logicalColumn);
}

int QTableWidget_VisualRow(QtObjectPtr ptr, int logicalRow){
	return static_cast<QTableWidget*>(ptr)->visualRow(logicalRow);
}

void QTableWidget_DestroyQTableWidget(QtObjectPtr ptr){
	static_cast<QTableWidget*>(ptr)->~QTableWidget();
}

