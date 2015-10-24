#include "qsignalmapper.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QObject>
#include <QSignalMapper>
#include "_cgo_export.h"

class MyQSignalMapper: public QSignalMapper {
public:
void Signal_Mapped(int i){callbackQSignalMapperMapped(this->objectName().toUtf8().data(), i);};
};

QtObjectPtr QSignalMapper_NewQSignalMapper(QtObjectPtr parent){
	return new QSignalMapper(static_cast<QObject*>(parent));
}

void QSignalMapper_Map(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map");
}

void QSignalMapper_Map2(QtObjectPtr ptr, QtObjectPtr sender){
	QMetaObject::invokeMethod(static_cast<QSignalMapper*>(ptr), "map", Q_ARG(QObject*, static_cast<QObject*>(sender)));
}

void QSignalMapper_ConnectMapped(QtObjectPtr ptr){
	QObject::connect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

void QSignalMapper_DisconnectMapped(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QSignalMapper*>(ptr), static_cast<void (QSignalMapper::*)(int)>(&QSignalMapper::mapped), static_cast<MyQSignalMapper*>(ptr), static_cast<void (MyQSignalMapper::*)(int)>(&MyQSignalMapper::Signal_Mapped));;
}

QtObjectPtr QSignalMapper_Mapping4(QtObjectPtr ptr, QtObjectPtr object){
	return static_cast<QSignalMapper*>(ptr)->mapping(static_cast<QObject*>(object));
}

QtObjectPtr QSignalMapper_Mapping2(QtObjectPtr ptr, char* id){
	return static_cast<QSignalMapper*>(ptr)->mapping(QString(id));
}

QtObjectPtr QSignalMapper_Mapping(QtObjectPtr ptr, int id){
	return static_cast<QSignalMapper*>(ptr)->mapping(id);
}

void QSignalMapper_RemoveMappings(QtObjectPtr ptr, QtObjectPtr sender){
	static_cast<QSignalMapper*>(ptr)->removeMappings(static_cast<QObject*>(sender));
}

void QSignalMapper_SetMapping4(QtObjectPtr ptr, QtObjectPtr sender, QtObjectPtr object){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), static_cast<QObject*>(object));
}

void QSignalMapper_SetMapping2(QtObjectPtr ptr, QtObjectPtr sender, char* text){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), QString(text));
}

void QSignalMapper_SetMapping(QtObjectPtr ptr, QtObjectPtr sender, int id){
	static_cast<QSignalMapper*>(ptr)->setMapping(static_cast<QObject*>(sender), id);
}

void QSignalMapper_DestroyQSignalMapper(QtObjectPtr ptr){
	static_cast<QSignalMapper*>(ptr)->~QSignalMapper();
}

