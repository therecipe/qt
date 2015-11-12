#include "qdbusobjectpath.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QString>
#include <QDBusObjectPath>
#include "_cgo_export.h"

class MyQDBusObjectPath: public QDBusObjectPath {
public:
};

void* QDBusObjectPath_NewQDBusObjectPath(){
	return new QDBusObjectPath();
}

void* QDBusObjectPath_NewQDBusObjectPath3(void* path){
	return new QDBusObjectPath(*static_cast<QLatin1String*>(path));
}

void* QDBusObjectPath_NewQDBusObjectPath4(char* path){
	return new QDBusObjectPath(QString(path));
}

void* QDBusObjectPath_NewQDBusObjectPath2(char* path){
	return new QDBusObjectPath(const_cast<const char*>(path));
}

char* QDBusObjectPath_Path(void* ptr){
	return static_cast<QDBusObjectPath*>(ptr)->path().toUtf8().data();
}

void QDBusObjectPath_SetPath(void* ptr, char* path){
	static_cast<QDBusObjectPath*>(ptr)->setPath(QString(path));
}

