#include "qdatawidgetmapper.h"
#include <QObject>
#include <QByteArray>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QAbstractItemModel>
#include <QMetaObject>
#include <QUrl>
#include <QModelIndex>
#include <QAbstractItemDelegate>
#include <QDataWidgetMapper>
#include "_cgo_export.h"

class MyQDataWidgetMapper: public QDataWidgetMapper {
public:
void Signal_CurrentIndexChanged(int index){callbackQDataWidgetMapperCurrentIndexChanged(this->objectName().toUtf8().data(), index);};
};

int QDataWidgetMapper_CurrentIndex(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->currentIndex();
}

int QDataWidgetMapper_Orientation(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->orientation();
}

void QDataWidgetMapper_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QDataWidgetMapper_SetOrientation(void* ptr, int aOrientation){
	static_cast<QDataWidgetMapper*>(ptr)->setOrientation(static_cast<Qt::Orientation>(aOrientation));
}

void QDataWidgetMapper_SetSubmitPolicy(void* ptr, int policy){
	static_cast<QDataWidgetMapper*>(ptr)->setSubmitPolicy(static_cast<QDataWidgetMapper::SubmitPolicy>(policy));
}

int QDataWidgetMapper_SubmitPolicy(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->submitPolicy();
}

void* QDataWidgetMapper_NewQDataWidgetMapper(void* parent){
	return new QDataWidgetMapper(static_cast<QObject*>(parent));
}

void QDataWidgetMapper_AddMapping(void* ptr, void* widget, int section){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section);
}

void QDataWidgetMapper_AddMapping2(void* ptr, void* widget, int section, void* propertyName){
	static_cast<QDataWidgetMapper*>(ptr)->addMapping(static_cast<QWidget*>(widget), section, *static_cast<QByteArray*>(propertyName));
}

void QDataWidgetMapper_ClearMapping(void* ptr){
	static_cast<QDataWidgetMapper*>(ptr)->clearMapping();
}

void QDataWidgetMapper_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

void QDataWidgetMapper_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QDataWidgetMapper*>(ptr), static_cast<void (QDataWidgetMapper::*)(int)>(&QDataWidgetMapper::currentIndexChanged), static_cast<MyQDataWidgetMapper*>(ptr), static_cast<void (MyQDataWidgetMapper::*)(int)>(&MyQDataWidgetMapper::Signal_CurrentIndexChanged));;
}

void* QDataWidgetMapper_ItemDelegate(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->itemDelegate();
}

void* QDataWidgetMapper_MappedPropertyName(void* ptr, void* widget){
	return new QByteArray(static_cast<QDataWidgetMapper*>(ptr)->mappedPropertyName(static_cast<QWidget*>(widget)));
}

int QDataWidgetMapper_MappedSection(void* ptr, void* widget){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedSection(static_cast<QWidget*>(widget));
}

void* QDataWidgetMapper_MappedWidgetAt(void* ptr, int section){
	return static_cast<QDataWidgetMapper*>(ptr)->mappedWidgetAt(section);
}

void* QDataWidgetMapper_Model(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->model();
}

void QDataWidgetMapper_RemoveMapping(void* ptr, void* widget){
	static_cast<QDataWidgetMapper*>(ptr)->removeMapping(static_cast<QWidget*>(widget));
}

void QDataWidgetMapper_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "revert");
}

void* QDataWidgetMapper_RootIndex(void* ptr){
	return static_cast<QDataWidgetMapper*>(ptr)->rootIndex().internalPointer();
}

void QDataWidgetMapper_SetCurrentModelIndex(void* ptr, void* index){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "setCurrentModelIndex", Q_ARG(QModelIndex, *static_cast<QModelIndex*>(index)));
}

void QDataWidgetMapper_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QDataWidgetMapper*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QDataWidgetMapper_SetModel(void* ptr, void* model){
	static_cast<QDataWidgetMapper*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QDataWidgetMapper_SetRootIndex(void* ptr, void* index){
	static_cast<QDataWidgetMapper*>(ptr)->setRootIndex(*static_cast<QModelIndex*>(index));
}

int QDataWidgetMapper_Submit(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "submit");
}

void QDataWidgetMapper_ToFirst(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toFirst");
}

void QDataWidgetMapper_ToLast(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toLast");
}

void QDataWidgetMapper_ToNext(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toNext");
}

void QDataWidgetMapper_ToPrevious(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDataWidgetMapper*>(ptr), "toPrevious");
}

void QDataWidgetMapper_DestroyQDataWidgetMapper(void* ptr){
	static_cast<QDataWidgetMapper*>(ptr)->~QDataWidgetMapper();
}

