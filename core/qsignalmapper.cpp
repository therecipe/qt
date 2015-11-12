#include "qsignalmapper.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QSignalMapper>
#include "_cgo_export.h"

class MyQSignalMapper: public QSignalMapper {
public:
void Signal_Mapped(int i){callbackQSignalMapperMapped(this->objectName().toUtf8().data(), i);};
};

void* QSignalMapper_NewQSignalMapper(void* parent){
	return new QSignalMapper(static_cast<QObject*>(parent));
}

void QSignalMapper_Map(void* ptr){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map");
}

void QSignalMapper_Map2(void* ptr, void* sender){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map", Q_ARG(QObject*, static_cast<QObject*>(sender)));
}

void QSignalMapper_ConnectMapped(void* ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

void QSignalMapper_DisconnectMapped(void* ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

void* QSignalMapper_Mapping4(void* ptr, void* object){
	return static_cast<QSignalMapper*>(ptr)->mapping(static_cast<QObject*>(object));
}

void* QSignalMapper_Mapping2(void* ptr, char* id){
	return static_cast<QSignalMapper*>(ptr)->mapping(QString(id));
}

void* QSignalMapper_Mapping(void* ptr, int id){
	return static_cast<QSignalMapper*>(ptr)->mapping(id);
}

void QSignalMapper_RemoveMappings(void* ptr, void* sender){
	static_cast<QSignalMapper*>(ptr)->removeMappings(static_cast<QObject*>(sender));
}

void QSignalMapper_SetMapping4(void* ptr, void* sender, void* object){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), static_cast<QObject*>(object));
}

void QSignalMapper_SetMapping2(void* ptr, void* sender, char* text){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), QString(text));
}

void QSignalMapper_SetMapping(void* ptr, void* sender, int id){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), id);
}

void QSignalMapper_DestroyQSignalMapper(void* ptr){
	static_cast<QSignalMapper*>(ptr)->~QSignalMapper();
}

