#include "qdbusobjectpath.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1String>
#include <QDBusObjectPath>
#include "_cgo_export.h"

class MyQDBusObjectPath: public QDBusObjectPath {
public:
};

QtObjectPtr QDBusObjectPath_NewQDBusObjectPath(){
	return new QDBusObjectPath();
}

QtObjectPtr QDBusObjectPath_NewQDBusObjectPath3(QtObjectPtr path){
	return new QDBusObjectPath(*static_cast<QLatin1String*>(path));
}

QtObjectPtr QDBusObjectPath_NewQDBusObjectPath4(char* path){
	return new QDBusObjectPath(QString(path));
}

QtObjectPtr QDBusObjectPath_NewQDBusObjectPath2(char* path){
	return new QDBusObjectPath(const_cast<const char*>(path));
}

char* QDBusObjectPath_Path(QtObjectPtr ptr){
	return static_cast<QDBusObjectPath*>(ptr)->path().toUtf8().data();
}

void QDBusObjectPath_SetPath(QtObjectPtr ptr, char* path){
	static_cast<QDBusObjectPath*>(ptr)->setPath(QString(path));
}

