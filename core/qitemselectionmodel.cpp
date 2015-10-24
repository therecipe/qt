#include "qitemselectionmodel.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QItemSelection>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QItemSelectionModel>
#include "_cgo_export.h"

class MyQItemSelectionModel: public QItemSelectionModel {
public:
void Signal_CurrentChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_CurrentColumnChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentColumnChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_CurrentRowChanged(const QModelIndex & current, const QModelIndex & previous){callbackQItemSelectionModelCurrentRowChanged(this->objectName().toUtf8().data(), current.internalPointer(), previous.internalPointer());};
void Signal_ModelChanged(QAbstractItemModel * model){callbackQItemSelectionModelModelChanged(this->objectName().toUtf8().data(), model);};
};

QtObjectPtr QItemSelectionModel_NewQItemSelectionModel(QtObjectPtr model){
	return new QItemSelectionModel(static_cast<QAbstractItemModel*>(model));
}

QtObjectPtr QItemSelectionModel_NewQItemSelectionModel2(QtObjectPtr model, QtObjectPtr parent){
	return new QItemSelectionModel(static_cast<QAbstractItemModel*>(model), static_cast<QObject*>(parent));
}

void QItemSelectionModel_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clear");
}

void QItemSelectionModel_ClearCurrentIndex(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearCurrentIndex");
}

void QItemSelectionModel_ClearSelection(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "clearSelection");
}

int QItemSelectionModel_ColumnIntersectsSelection(QtObjectPtr ptr, int column, QtObjectPtr parent){
	return static_cast<QItemSelectionModel*>(ptr)->columnIntersectsSelection(column, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_ConnectCurrentChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_DisconnectCurrentChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentChanged));;
}

void QItemSelectionModel_ConnectCurrentColumnChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

void QItemSelectionModel_DisconnectCurrentColumnChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentColumnChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentColumnChanged));;
}

QtObjectPtr QItemSelectionModel_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QItemSelectionModel*>(ptr)->currentIndex().internalPointer();
}

void QItemSelectionModel_ConnectCurrentRowChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

void QItemSelectionModel_DisconnectCurrentRowChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&QItemSelectionModel::currentRowChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(const QModelIndex &, const QModelIndex &)>(&MyQItemSelectionModel::Signal_CurrentRowChanged));;
}

int QItemSelectionModel_HasSelection(QtObjectPtr ptr){
	return static_cast<QItemSelectionModel*>(ptr)->hasSelection();
}

int QItemSelectionModel_IsColumnSelected(QtObjectPtr ptr, int column, QtObjectPtr parent){
	return static_cast<QItemSelectionModel*>(ptr)->isColumnSelected(column, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsRowSelected(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QItemSelectionModel*>(ptr)->isRowSelected(row, *static_cast<QModelIndex*>(parent));
}

int QItemSelectionModel_IsSelected(QtObjectPtr ptr, QtObjectPtr index){
	return static_cast<QItemSelectionModel*>(ptr)->isSelected(*static_cast<QModelIndex*>(index));
}

QtObjectPtr QItemSelectionModel_Model2(QtObjectPtr ptr){
	return static_cast<QItemSelectionModel*>(ptr)->model();
}

QtObjectPtr QItemSelectionModel_Model(QtObjectPtr ptr){
	return const_cast<QAbstractItemModel*>(static_cast<QItemSelectionModel*>(ptr)->model());
}

void QItemSelectionModel_ConnectModelChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_DisconnectModelChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QItemSelectionModel*>(ptr), static_cast<void (QItemSelectionModel::*)(QAbstractItemModel *)>(&QItemSelectionModel::modelChanged), static_cast<MyQItemSelectionModel*>(ptr), static_cast<void (MyQItemSelectionModel::*)(QAbstractItemModel *)>(&MyQItemSelectionModel::Signal_ModelChanged));;
}

void QItemSelectionModel_Reset(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "reset");
}

int QItemSelectionModel_RowIntersectsSelection(QtObjectPtr ptr, int row, QtObjectPtr parent){
	return static_cast<QItemSelectionModel*>(ptr)->rowIntersectsSelection(row, *static_cast<QModelIndex*>(parent));
}

void QItemSelectionModel_Select2(QtObjectPtr ptr, QtObjectPtr selection, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "select", Q_ARG(QItemSelection, *static_cast<QItemSelection*>(selection)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_Select(QtObjectPtr ptr, QtObjectPtr index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "select", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetCurrentIndex(QtObjectPtr ptr, QtObjectPtr index, int command){
	QMetaObject::invokeMethod(static_cast<QItemSelectionModel*>(ptr), "setCurrentIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)), Q_ARG(QItemSelectionModel::SelectionFlag, static_cast<QItemSelectionModel::SelectionFlag>(command)));
}

void QItemSelectionModel_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QItemSelectionModel*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QItemSelectionModel_DestroyQItemSelectionModel(QtObjectPtr ptr){
	static_cast<QItemSelectionModel*>(ptr)->~QItemSelectionModel();
}

