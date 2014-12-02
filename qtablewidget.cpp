#include "qtablewidget.h"
#include <QTableWidget>
#include "cgoexport.h"

//MyQTableWidget
class MyQTableWidget: public QTableWidget {
Q_OBJECT
public:
void Signal_CellActivated() { callbackQTableWidget(this, QString("cellActivated").toUtf8().data()); };
void Signal_CellChanged() { callbackQTableWidget(this, QString("cellChanged").toUtf8().data()); };
void Signal_CellClicked() { callbackQTableWidget(this, QString("cellClicked").toUtf8().data()); };
void Signal_CellDoubleClicked() { callbackQTableWidget(this, QString("cellDoubleClicked").toUtf8().data()); };
void Signal_CellEntered() { callbackQTableWidget(this, QString("cellEntered").toUtf8().data()); };
void Signal_CellPressed() { callbackQTableWidget(this, QString("cellPressed").toUtf8().data()); };
void Signal_CurrentCellChanged() { callbackQTableWidget(this, QString("currentCellChanged").toUtf8().data()); };
void Signal_ItemSelectionChanged() { callbackQTableWidget(this, QString("itemSelectionChanged").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_ClearContents();
void Slot_InsertColumn(int column);
void Slot_InsertRow(int row);
void Slot_RemoveColumn(int column);
void Slot_RemoveRow(int row);

};
#include "qtablewidget.moc"


//Public Functions
QtObjectPtr QTableWidget_New_Int_Int_QWidget(int rows, int columns, QtObjectPtr parent)
{
	return new QTableWidget(rows, columns, ((QWidget*)(parent)));
}

void QTableWidget_Destroy(QtObjectPtr ptr)
{
	((QTableWidget*)(ptr))->~QTableWidget();
}

QtObjectPtr QTableWidget_CellWidget_Int_Int(QtObjectPtr ptr, int row, int column)
{
	return ((QTableWidget*)(ptr))->cellWidget(row, column);
}

int QTableWidget_ColumnCount(QtObjectPtr ptr)
{
	return ((QTableWidget*)(ptr))->columnCount();
}

int QTableWidget_CurrentColumn(QtObjectPtr ptr)
{
	return ((QTableWidget*)(ptr))->currentColumn();
}

int QTableWidget_CurrentRow(QtObjectPtr ptr)
{
	return ((QTableWidget*)(ptr))->currentRow();
}

void QTableWidget_RemoveCellWidget_Int_Int(QtObjectPtr ptr, int row, int column)
{
	((QTableWidget*)(ptr))->removeCellWidget(row, column);
}

int QTableWidget_RowCount(QtObjectPtr ptr)
{
	return ((QTableWidget*)(ptr))->rowCount();
}

void QTableWidget_SetCellWidget_Int_Int_QWidget(QtObjectPtr ptr, int row, int column, QtObjectPtr widget)
{
	((QTableWidget*)(ptr))->setCellWidget(row, column, ((QWidget*)(widget)));
}

void QTableWidget_SetColumnCount_Int(QtObjectPtr ptr, int columns)
{
	((QTableWidget*)(ptr))->setColumnCount(columns);
}

void QTableWidget_SetCurrentCell_Int_Int(QtObjectPtr ptr, int row, int column)
{
	((QTableWidget*)(ptr))->setCurrentCell(row, column);
}

void QTableWidget_SetRowCount_Int(QtObjectPtr ptr, int rows)
{
	((QTableWidget*)(ptr))->setRowCount(rows);
}

void QTableWidget_SortItems_Int_SortOrder(QtObjectPtr ptr, int column, int order)
{
	((QTableWidget*)(ptr))->sortItems(column, ((Qt::SortOrder)(order)));
}

int QTableWidget_VisualColumn_Int(QtObjectPtr ptr, int logicalColumn)
{
	return ((QTableWidget*)(ptr))->visualColumn(logicalColumn);
}

int QTableWidget_VisualRow_Int(QtObjectPtr ptr, int logicalRow)
{
	return ((QTableWidget*)(ptr))->visualRow(logicalRow);
}

//Public Slots
void QTableWidget_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_Clear, ((QTableWidget*)(ptr)), &QTableWidget::clear, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_Clear, ((QTableWidget*)(ptr)), &QTableWidget::clear);
}

void QTableWidget_Clear(QtObjectPtr ptr)
{
	((MyQTableWidget*)(ptr))->Slot_Clear();
}

void QTableWidget_ConnectSlotClearContents(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_ClearContents, ((QTableWidget*)(ptr)), &QTableWidget::clearContents, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotClearContents(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_ClearContents, ((QTableWidget*)(ptr)), &QTableWidget::clearContents);
}

void QTableWidget_ClearContents(QtObjectPtr ptr)
{
	((MyQTableWidget*)(ptr))->Slot_ClearContents();
}

void QTableWidget_ConnectSlotInsertColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_InsertColumn, ((QTableWidget*)(ptr)), &QTableWidget::insertColumn, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotInsertColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_InsertColumn, ((QTableWidget*)(ptr)), &QTableWidget::insertColumn);
}

void QTableWidget_InsertColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTableWidget*)(ptr))->Slot_InsertColumn(column);
}

void QTableWidget_ConnectSlotInsertRow(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_InsertRow, ((QTableWidget*)(ptr)), &QTableWidget::insertRow, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotInsertRow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_InsertRow, ((QTableWidget*)(ptr)), &QTableWidget::insertRow);
}

void QTableWidget_InsertRow_Int(QtObjectPtr ptr, int row)
{
	((MyQTableWidget*)(ptr))->Slot_InsertRow(row);
}

void QTableWidget_ConnectSlotRemoveColumn(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_RemoveColumn, ((QTableWidget*)(ptr)), &QTableWidget::removeColumn, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotRemoveColumn(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_RemoveColumn, ((QTableWidget*)(ptr)), &QTableWidget::removeColumn);
}

void QTableWidget_RemoveColumn_Int(QtObjectPtr ptr, int column)
{
	((MyQTableWidget*)(ptr))->Slot_RemoveColumn(column);
}

void QTableWidget_ConnectSlotRemoveRow(QtObjectPtr ptr)
{
	QObject::connect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_RemoveRow, ((QTableWidget*)(ptr)), &QTableWidget::removeRow, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSlotRemoveRow(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQTableWidget*)(ptr)), &MyQTableWidget::Slot_RemoveRow, ((QTableWidget*)(ptr)), &QTableWidget::removeRow);
}

void QTableWidget_RemoveRow_Int(QtObjectPtr ptr, int row)
{
	((MyQTableWidget*)(ptr))->Slot_RemoveRow(row);
}

//Signals
void QTableWidget_ConnectSignalCellActivated(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellActivated, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellActivated, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellActivated(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellActivated, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellActivated);
}

void QTableWidget_ConnectSignalCellChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellChanged, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellChanged);
}

void QTableWidget_ConnectSignalCellClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellClicked, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellClicked, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellClicked, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellClicked);
}

void QTableWidget_ConnectSignalCellDoubleClicked(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellDoubleClicked, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellDoubleClicked, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellDoubleClicked(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellDoubleClicked, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellDoubleClicked);
}

void QTableWidget_ConnectSignalCellEntered(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellEntered, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellEntered, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellEntered(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellEntered, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellEntered);
}

void QTableWidget_ConnectSignalCellPressed(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::cellPressed, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellPressed, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCellPressed(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::cellPressed, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CellPressed);
}

void QTableWidget_ConnectSignalCurrentCellChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::currentCellChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CurrentCellChanged, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalCurrentCellChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::currentCellChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_CurrentCellChanged);
}

void QTableWidget_ConnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::connect(((QTableWidget*)(ptr)), &QTableWidget::itemSelectionChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_ItemSelectionChanged, Qt::QueuedConnection);
}

void QTableWidget_DisconnectSignalItemSelectionChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QTableWidget*)(ptr)), &QTableWidget::itemSelectionChanged, ((MyQTableWidget*)(ptr)), &MyQTableWidget::Signal_ItemSelectionChanged);
}

