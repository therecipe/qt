#include "qfileopenevent.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFile>
#include <QIODevice>
#include <QFileOpenEvent>
#include "_cgo_export.h"

class MyQFileOpenEvent: public QFileOpenEvent {
public:
};

int QFileOpenEvent_OpenFile(QtObjectPtr ptr, QtObjectPtr file, int flags){
	return static_cast<QFileOpenEvent*>(ptr)->openFile(*static_cast<QFile*>(file), static_cast<QIODevice::OpenModeFlag>(flags));
}

char* QFileOpenEvent_File(QtObjectPtr ptr){
	return static_cast<QFileOpenEvent*>(ptr)->file().toUtf8().data();
}

char* QFileOpenEvent_Url(QtObjectPtr ptr){
	return static_cast<QFileOpenEvent*>(ptr)->url().toString().toUtf8().data();
}

