#include "qdbusargument.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDBusArgument>
#include "_cgo_export.h"

class MyQDBusArgument: public QDBusArgument {
public:
};

QtObjectPtr QDBusArgument_NewQDBusArgument(){
	return new QDBusArgument();
}

QtObjectPtr QDBusArgument_NewQDBusArgument2(QtObjectPtr other){
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

char* QDBusArgument_AsVariant(QtObjectPtr ptr){
	return static_cast<QDBusArgument*>(ptr)->asVariant().toString().toUtf8().data();
}

int QDBusArgument_AtEnd(QtObjectPtr ptr){
	return static_cast<QDBusArgument*>(ptr)->atEnd();
}

void QDBusArgument_BeginArray(QtObjectPtr ptr, int id){
	static_cast<QDBusArgument*>(ptr)->beginArray(id);
}

void QDBusArgument_BeginArray2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginArray();
}

void QDBusArgument_BeginMap(QtObjectPtr ptr, int kid, int vid){
	static_cast<QDBusArgument*>(ptr)->beginMap(kid, vid);
}

void QDBusArgument_BeginMap2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginMap();
}

void QDBusArgument_BeginMapEntry(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginMapEntry2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginStructure(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

void QDBusArgument_BeginStructure2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

int QDBusArgument_CurrentType(QtObjectPtr ptr){
	return static_cast<QDBusArgument*>(ptr)->currentType();
}

void QDBusArgument_EndArray(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndArray2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndMap(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMap2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMapEntry(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndMapEntry2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndStructure(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_EndStructure2(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_DestroyQDBusArgument(QtObjectPtr ptr){
	static_cast<QDBusArgument*>(ptr)->~QDBusArgument();
}

