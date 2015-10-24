#include "qlibraryinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLibrary>
#include <QLibraryInfo>
#include "_cgo_export.h"

class MyQLibraryInfo: public QLibraryInfo {
public:
};

int QLibraryInfo_QLibraryInfo_IsDebugBuild(){
	return QLibraryInfo::isDebugBuild();
}

char* QLibraryInfo_QLibraryInfo_LicensedProducts(){
	return QLibraryInfo::licensedProducts().toUtf8().data();
}

char* QLibraryInfo_QLibraryInfo_Licensee(){
	return QLibraryInfo::licensee().toUtf8().data();
}

char* QLibraryInfo_QLibraryInfo_Location(int loc){
	return QLibraryInfo::location(static_cast<QLibraryInfo::LibraryLocation>(loc)).toUtf8().data();
}

