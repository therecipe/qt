#include "qstandardpaths.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QStandardPaths>
#include "_cgo_export.h"

class MyQStandardPaths: public QStandardPaths {
public:
};

void QStandardPaths_QStandardPaths_SetTestModeEnabled(int testMode){
	QStandardPaths::setTestModeEnabled(testMode != 0);
}

char* QStandardPaths_QStandardPaths_FindExecutable(char* executableName, char* paths){
	return QStandardPaths::findExecutable(QString(executableName), QString(paths).split("|", QString::SkipEmptyParts)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_Locate(int ty, char* fileName, int options){
	return QStandardPaths::locate(static_cast<QStandardPaths::StandardLocation>(ty), QString(fileName), static_cast<QStandardPaths::LocateOption>(options)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_LocateAll(int ty, char* fileName, int options){
	return QStandardPaths::locateAll(static_cast<QStandardPaths::StandardLocation>(ty), QString(fileName), static_cast<QStandardPaths::LocateOption>(options)).join("|").toUtf8().data();
}

char* QStandardPaths_QStandardPaths_DisplayName(int ty){
	return QStandardPaths::displayName(static_cast<QStandardPaths::StandardLocation>(ty)).toUtf8().data();
}

char* QStandardPaths_QStandardPaths_StandardLocations(int ty){
	return QStandardPaths::standardLocations(static_cast<QStandardPaths::StandardLocation>(ty)).join("|").toUtf8().data();
}

char* QStandardPaths_QStandardPaths_WritableLocation(int ty){
	return QStandardPaths::writableLocation(static_cast<QStandardPaths::StandardLocation>(ty)).toUtf8().data();
}

