#include "qcoreapplication.h"
#include <QCoreApplication>
#include "cgoexport.h"



//Public Functions
QtObjectPtr QCoreApplication_New_Int_String(int argc, char* argv)
{
	return new QCoreApplication(argc, ((char**)(argv)));
}

void QCoreApplication_Destroy(QtObjectPtr ptr)
{
	((QCoreApplication*)(ptr))->~QCoreApplication();
}

//Static Public Members
void QCoreApplication_AddLibraryPath_String(char* path)
{
	QCoreApplication::addLibraryPath(QString(path));
}

char* QCoreApplication_ApplicationDirPath()
{
	return QCoreApplication::applicationDirPath().toUtf8().data();
}

char* QCoreApplication_ApplicationFilePath()
{
	return QCoreApplication::applicationFilePath().toUtf8().data();
}

char* QCoreApplication_ApplicationName()
{
	return QCoreApplication::applicationName().toUtf8().data();
}

char* QCoreApplication_ApplicationVersion()
{
	return QCoreApplication::applicationVersion().toUtf8().data();
}

int QCoreApplication_ClosingDown()
{
	return QCoreApplication::closingDown();
}

int QCoreApplication_Exec()
{
	return QCoreApplication::exec();
}

void QCoreApplication_Exit_Int(int returnCode)
{
	QCoreApplication::exit(returnCode);
}

QtObjectPtr QCoreApplication_Instance()
{
	return QCoreApplication::instance();
}

int QCoreApplication_IsQuitLockEnabled()
{
	return QCoreApplication::isQuitLockEnabled();
}

int QCoreApplication_IsSetuidAllowed()
{
	return QCoreApplication::isSetuidAllowed();
}

char* QCoreApplication_OrganizationDomain()
{
	return QCoreApplication::organizationDomain().toUtf8().data();
}

char* QCoreApplication_OrganizationName()
{
	return QCoreApplication::organizationName().toUtf8().data();
}

void QCoreApplication_RemoveLibraryPath_String(char* path)
{
	QCoreApplication::removeLibraryPath(QString(path));
}

void QCoreApplication_RemovePostedEvents_QObject_Int(QtObjectPtr receiver, int eventType)
{
	QCoreApplication::removePostedEvents(((QObject*)(receiver)), eventType);
}

void QCoreApplication_SendPostedEvents_QObject_Int(QtObjectPtr receiver, int event_type)
{
	QCoreApplication::sendPostedEvents(((QObject*)(receiver)), event_type);
}

void QCoreApplication_SetApplicationName_String(char* application)
{
	QCoreApplication::setApplicationName(QString(application));
}

void QCoreApplication_SetApplicationVersion_String(char* version)
{
	QCoreApplication::setApplicationVersion(QString(version));
}

void QCoreApplication_SetAttribute_ApplicationAttribute_Bool(int attribute, int on)
{
	QCoreApplication::setAttribute(((Qt::ApplicationAttribute)(attribute)), on != 0);
}

void QCoreApplication_SetOrganizationDomain_String(char* orgDomain)
{
	QCoreApplication::setOrganizationDomain(QString(orgDomain));
}

void QCoreApplication_SetOrganizationName_String(char* orgName)
{
	QCoreApplication::setOrganizationName(QString(orgName));
}

void QCoreApplication_SetQuitLockEnabled_Bool(int enabled)
{
	QCoreApplication::setQuitLockEnabled(enabled != 0);
}

void QCoreApplication_SetSetuidAllowed_Bool(int allow)
{
	QCoreApplication::setSetuidAllowed(allow != 0);
}

int QCoreApplication_StartingUp()
{
	return QCoreApplication::startingUp();
}

int QCoreApplication_TestAttribute_ApplicationAttribute(int attribute)
{
	return QCoreApplication::testAttribute(((Qt::ApplicationAttribute)(attribute)));
}

char* QCoreApplication_Translate_String_String_String_Int(char* context, char* sourceText, char* disambiguation, int n)
{
	return QCoreApplication::translate(context, sourceText, disambiguation, n).toUtf8().data();
}

