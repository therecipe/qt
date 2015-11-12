#include "qitemselectionmodel.h"
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QItemSelection>
#include <QAbstractItemModel>
#include <QString>
#include <QVariant>
#include <QItemSelectionModel>
#include "_cgo_export.h"

class MyQItemSelectionModel: public QItemSelectionModel {
public:
void Signal_CurrentChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_CurrentColumnChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentColumnChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_CurrentRowChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentRowChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_ModelChanged(QAbstractItemModel * model){callbackQItemSelectionModelModelChanged(this->objectName().toUtf8().data(), model);};
};

void* QItemSelectionModel_NewQItemSelectionModel(void* model){
	return new QItemSelectionModel(static_cast<QAbstractItemModel*>(model));
}

void* QItemSelectionModel_NewQItemSelectionModel2(void* model, void* parent){
	return new QItemSelectionModel(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

void QItemSelectionModel_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clear");
}

void QItemSelectionModel_ClearCurrentIndex(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearCurrentIndex");
}

void QItemSelectionModel_ClearSelection(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearSelection");
}

int QItemSelectionModel_ColumnIntersectsSelection(void* ptr, int column, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->columnIntersectsSelection(column, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_ConnectCurrentChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_DisconnectCurrentChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_ConnectCurrentColumnChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

void QItemSelectionModel_DisconnectCurrentColumnChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

void* QItemSelectionModel_CurrentIndex(void* ptr){
	return static_cast<QItemSelectionModel*>(ptr)->currentIndex().internalPointer();
}

void QItemSelectionModel_ConnectCurrentRowChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

void QItemSelectionModel_DisconnectCurrentRowChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

int QItemSelectionModel_HasSelection(void* ptr){
	return static_cast<QItemSelectionModel*>(ptr)->hasSelection();
}

int QItemSelectionModel_IsColumnSelected(void* ptr, int column, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->isColumnSelected(column, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsRowSelected(void* ptr, int row, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->isRowSelected(row, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsSelected(void* ptr, void* index){
	return static_cast<QItemSelectionModel*>(ptr)->isSelected(*static_cast<QModelIndex*>(index));
}

void* QItemSelectionModel_Model2(void* ptr){
	return static_cast<QItemSelectionModel*>(ptr)->model();
}

void* QItemSelectionModel_Model(void* ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionModel*>(ptr)->model());
}

void QItemSelectionModel_ConnectModelChanged(void* ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_DisconnectModelChanged(void* ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "reset");
}

int QItemSelectionModel_RowIntersectsSelection(void* ptr, int row, void* parent){
	return static_cast<QItemSelectionModel*>(ptr)->rowIntersectsSelection(row, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_Select2(void* ptr, void* selection, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "select", Q_ARG(QItemSelection, *static_cast<QItemSelection*>(selection)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_Select(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "select", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetCurrentIndex(void* ptr, void* index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetModel(void* ptr, void* model){
	static_cast<QItemSelectionModel*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QItemSelectionModel_DestroyQItemSelectionModel(void* ptr){
	static_cast<QItemSelectionModel*>(ptr)->~QItemSelectionModel();
}

