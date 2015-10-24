#include "qdatawidgetmapper.h"
#include <QVariant>
#include <QAbstractItemDelegate>
#include <QObject>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QByteArray>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QDataWidgetMapper>
#include "_cgo_export.h"

class MyQDataWidgetMapper: public QDataWidgetMapper {
public:
void Signal_CurrentIndexChanged(int index){callbackQDataWidgetMapperCurrentIndexChanged(this->objectName().toUtf8().data(), index);};
};

int QDataWidgetMapper_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->currentIndex();
}

int QDataWidgetMapper_Orientation(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->orientation();
}

void QDataWidgetMapper_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QDataWidgetMapper_SetOrientation(QtObjectPtr ptr, int aOrientation){
	static_cast<QDataWidgetMapper*>(ptr)->setOrientation(static_cast<Qt::Orientation>(aOrientation));
}

void QDataWidgetMapper_SetSubmitPolicy(QtObjectPtr ptr, int policy){
	static_cast<QDataWidgetMapper*>(ptr)->setSubmitPolicy(static_cast<QDataWidgetMapper::SubmitPolicy>(policy));
}

int QDataWidgetMapper_SubmitPolicy(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->submitPolicy();
}

QtObjectPtr QDataWidgetMapper_NewQDataWidgetMapper(QtObjectPtr parent){
	return new QDataWidgetMapper(static_cast<QObject*>(parent));
}

void QDataWidgetMapper_AddMapping(QtObjectPtr ptr, QtObjectPtr widget, int section){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section);
}

void QDataWidgetMapper_AddMapping2(QtObjectPtr ptr, QtObjectPtr widget, int section, QtObjectPtr propertyName){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section, *static_cast<QByteArray*>(propertyName));
}

void QDataWidgetMapper_ClearMapping(QtObjectPtr ptr){
	static_cast<QDataWidgetMapper*>(ptr)->clearMapping();
}

void QDataWidgetMapper_ConnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

void QDataWidgetMapper_DisconnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

QtObjectPtr QDataWidgetMapper_ItemDelegate(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->itemDelegate();
}

int QDataWidgetMapper_MappedSection(QtObjectPtr ptr, QtObjectPtr widget){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedSection(static_cast<QWidget*>(widget));
}

QtObjectPtr QDataWidgetMapper_MappedWidgetAt(QtObjectPtr ptr, int section){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedWidgetAt(section);
}

QtObjectPtr QDataWidgetMapper_Model(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->model();
}

void QDataWidgetMapper_RemoveMapping(QtObjectPtr ptr, QtObjectPtr widget){
	static_cast<QDataWidgetMapper*>(ptr)->removeMapping(static_cast<QWidget*>(widget));
}

void QDataWidgetMapper_Revert(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "revert");
}

QtObjectPtr QDataWidgetMapper_RootIndex(QtObjectPtr ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->rootIndex().internalPointer();
}

void QDataWidgetMapper_SetCurrentModelIndex(QtObjectPtr ptr, QtObjectPtr index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentModelIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QDataWidgetMapper_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate){
	static_cast<QDataWidgetMapper*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QDataWidgetMapper_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QDataWidgetMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QDataWidgetMapper_SetRootIndex(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QDataWidgetMapper*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

int QDataWidgetMapper_Submit(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "submit");
}

void QDataWidgetMapper_ToFirst(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toFirst");
}

void QDataWidgetMapper_ToLast(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toLast");
}

void QDataWidgetMapper_ToNext(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toNext");
}

void QDataWidgetMapper_ToPrevious(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toPrevious");
}

void QDataWidgetMapper_DestroyQDataWidgetMapper(QtObjectPtr ptr){
	static_cast<QDataWidgetMapper*>(ptr)->~QDataWidgetMapper();
}

