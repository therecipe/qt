#include "qfileopenevent.h"
#include <QIODevice>
#include <QFile>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFileOpenEvent>
#include "_cgo_export.h"

class MyQFileOpenEvent: public QFileOpenEvent {
public:
};

int QFileOpenEvent_OpenFile(void* ptr, void* file, int flags){
	return static_cast<QFileOpenEvent*>(ptr)->openFile(*static_cast<QFile*>(file), static_cast<QIODevice::OpenModeFlag>(flags));
}

char* QFileOpenEvent_File(void* ptr){
	return static_cast<QFileOpenEvent*>(ptr)->file().toUtf8().data();
}

