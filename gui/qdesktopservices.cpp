#include "qdesktopservices.h"
#include <QModelIndex>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QDesktopServices>
#include "_cgo_export.h"

class MyQDesktopServices: public QDesktopServices {
public:
};

int QDesktopServices_QDesktopServices_OpenUrl(char* url){
	return QDesktopServices::openUrl(QUrl(QString(url)));
}

void QDesktopServices_QDesktopServices_SetUrlHandler(char* scheme, QtObjectPtr receiver, char* method){
	QDesktopServices::setUrlHandler(QString(scheme), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

void QDesktopServices_QDesktopServices_UnsetUrlHandler(char* scheme){
	QDesktopServices::unsetUrlHandler(QString(scheme));
}

