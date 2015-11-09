#include "qdesktopservices.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QObject>
#include <QDesktopServices>
#include "_cgo_export.h"

class MyQDesktopServices: public QDesktopServices {
public:
};

int QDesktopServices_QDesktopServices_OpenUrl(void* url){
	return QDesktopServices::openUrl(*static_cast<QUrl*>(url));
}

void QDesktopServices_QDesktopServices_SetUrlHandler(char* scheme, void* receiver, char* method){
	QDesktopServices::setUrlHandler(QString(scheme), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

void QDesktopServices_QDesktopServices_UnsetUrlHandler(char* scheme){
	QDesktopServices::unsetUrlHandler(QString(scheme));
}

