#include "qdesktopservices.h"
#include <QDesktopServices>
#include "cgoexport.h"


#include <QUrl>
//Static Public Members
int QDesktopServices_OpenUrl_String(char* url)
{
	return QDesktopServices::openUrl(QUrl(QString(url)));
}

void QDesktopServices_SetUrlHandler_String_QObject_String(char* scheme, QtObjectPtr receiver, char* method)
{
	QDesktopServices::setUrlHandler(QString(scheme), ((QObject*)(receiver)), method);
}

void QDesktopServices_UnsetUrlHandler_String(char* scheme)
{
	QDesktopServices::unsetUrlHandler(QString(scheme));
}

