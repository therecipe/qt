#include "qcoreapplication.h"
#include <QVariant>
#include <QEvent>
#include <QObject>
#include <QTranslator>
#include <QAbstractEventDispatcher>
#include <QAbstractNativeEventFilter>
#include <QString>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QEventLoop>
#include <QCoreApplication>
#include "_cgo_export.h"

class MyQCoreApplication: public QCoreApplication {
public:
void Signal_AboutToQuit(){callbackQCoreApplicationAboutToQuit(this->objectName().toUtf8().data());};
};

char* QCoreApplication_QCoreApplication_ApplicationName(){
	return QCoreApplication::applicationName().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_ApplicationVersion(){
	return QCoreApplication::applicationVersion().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_OrganizationDomain(){
	return QCoreApplication::organizationDomain().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_OrganizationName(){
	return QCoreApplication::organizationName().toUtf8().data();
}

void QCoreApplication_QCoreApplication_SetApplicationName(char* application){
	QCoreApplication::setApplicationName(QString(application));
}

void QCoreApplication_QCoreApplication_SetApplicationVersion(char* version){
	QCoreApplication::setApplicationVersion(QString(version));
}

void QCoreApplication_QCoreApplication_SetOrganizationDomain(char* orgDomain){
	QCoreApplication::setOrganizationDomain(QString(orgDomain));
}

void QCoreApplication_QCoreApplication_SetOrganizationName(char* orgName){
	QCoreApplication::setOrganizationName(QString(orgName));
}

QtObjectPtr QCoreApplication_NewQCoreApplication(int argc, char* argv){
	return new QCoreApplication(argc, &argv);
}

void QCoreApplication_ConnectAboutToQuit(QtObjectPtr ptr){
	QObject::connect(static_cast<QCoreApplication*>(ptr), &QCoreApplication::aboutToQuit, static_cast<MyQCoreApplication*>(ptr), static_cast<void (MyQCoreApplication::*)()>(&MyQCoreApplication::Signal_AboutToQuit));;
}

void QCoreApplication_DisconnectAboutToQuit(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QCoreApplication*>(ptr), &QCoreApplication::aboutToQuit, static_cast<MyQCoreApplication*>(ptr), static_cast<void (MyQCoreApplication::*)()>(&MyQCoreApplication::Signal_AboutToQuit));;
}

void QCoreApplication_QCoreApplication_AddLibraryPath(char* path){
	QCoreApplication::addLibraryPath(QString(path));
}

char* QCoreApplication_QCoreApplication_ApplicationDirPath(){
	return QCoreApplication::applicationDirPath().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_ApplicationFilePath(){
	return QCoreApplication::applicationFilePath().toUtf8().data();
}

char* QCoreApplication_QCoreApplication_Arguments(){
	return QCoreApplication::arguments().join("|").toUtf8().data();
}

int QCoreApplication_QCoreApplication_ClosingDown(){
	return QCoreApplication::closingDown();
}

QtObjectPtr QCoreApplication_QCoreApplication_EventDispatcher(){
	return QCoreApplication::eventDispatcher();
}

int QCoreApplication_QCoreApplication_Exec(){
	return QCoreApplication::exec();
}

void QCoreApplication_QCoreApplication_Exit(int returnCode){
	QCoreApplication::exit(returnCode);
}

void QCoreApplication_QCoreApplication_Flush(){
	QCoreApplication::flush();
}

void QCoreApplication_InstallNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObj){
	static_cast<QCoreApplication*>(ptr)->installNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObj));
}

int QCoreApplication_QCoreApplication_InstallTranslator(QtObjectPtr translationFile){
	return QCoreApplication::installTranslator(static_cast<QTranslator*>(translationFile));
}

QtObjectPtr QCoreApplication_QCoreApplication_Instance(){
	return QCoreApplication::instance();
}

int QCoreApplication_QCoreApplication_IsQuitLockEnabled(){
	return QCoreApplication::isQuitLockEnabled();
}

int QCoreApplication_QCoreApplication_IsSetuidAllowed(){
	return QCoreApplication::isSetuidAllowed();
}

char* QCoreApplication_QCoreApplication_LibraryPaths(){
	return QCoreApplication::libraryPaths().join("|").toUtf8().data();
}

int QCoreApplication_Notify(QtObjectPtr ptr, QtObjectPtr receiver, QtObjectPtr event){
	return static_cast<QCoreApplication*>(ptr)->notify(static_cast<QObject*>(receiver), static_cast<QEvent*>(event));
}

void QCoreApplication_QCoreApplication_PostEvent(QtObjectPtr receiver, QtObjectPtr event, int priority){
	QCoreApplication::postEvent(static_cast<QObject*>(receiver), static_cast<QEvent*>(event), priority);
}

void QCoreApplication_QCoreApplication_ProcessEvents(int flags){
	QCoreApplication::processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags));
}

void QCoreApplication_QCoreApplication_ProcessEvents2(int flags, int maxtime){
	QCoreApplication::processEvents(static_cast<QEventLoop::ProcessEventsFlag>(flags), maxtime);
}

void QCoreApplication_QCoreApplication_Quit(){
	QMetaObject::invokeMethod(QCoreApplication::instance(), "quit");
}

void QCoreApplication_QCoreApplication_RemoveLibraryPath(char* path){
	QCoreApplication::removeLibraryPath(QString(path));
}

void QCoreApplication_RemoveNativeEventFilter(QtObjectPtr ptr, QtObjectPtr filterObject){
	static_cast<QCoreApplication*>(ptr)->removeNativeEventFilter(static_cast<QAbstractNativeEventFilter*>(filterObject));
}

void QCoreApplication_QCoreApplication_RemovePostedEvents(QtObjectPtr receiver, int eventType){
	QCoreApplication::removePostedEvents(static_cast<QObject*>(receiver), eventType);
}

int QCoreApplication_QCoreApplication_RemoveTranslator(QtObjectPtr translationFile){
	return QCoreApplication::removeTranslator(static_cast<QTranslator*>(translationFile));
}

int QCoreApplication_QCoreApplication_SendEvent(QtObjectPtr receiver, QtObjectPtr event){
	return QCoreApplication::sendEvent(static_cast<QObject*>(receiver), static_cast<QEvent*>(event));
}

void QCoreApplication_QCoreApplication_SendPostedEvents(QtObjectPtr receiver, int event_type){
	QCoreApplication::sendPostedEvents(static_cast<QObject*>(receiver), event_type);
}

void QCoreApplication_QCoreApplication_SetAttribute(int attribute, int on){
	QCoreApplication::setAttribute(static_cast<Qt::ApplicationAttribute>(attribute), on != 0);
}

void QCoreApplication_QCoreApplication_SetEventDispatcher(QtObjectPtr eventDispatcher){
	QCoreApplication::setEventDispatcher(static_cast<QAbstractEventDispatcher*>(eventDispatcher));
}

void QCoreApplication_QCoreApplication_SetLibraryPaths(char* paths){
	QCoreApplication::setLibraryPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QCoreApplication_QCoreApplication_SetQuitLockEnabled(int enabled){
	QCoreApplication::setQuitLockEnabled(enabled != 0);
}

void QCoreApplication_QCoreApplication_SetSetuidAllowed(int allow){
	QCoreApplication::setSetuidAllowed(allow != 0);
}

int QCoreApplication_QCoreApplication_StartingUp(){
	return QCoreApplication::startingUp();
}

int QCoreApplication_QCoreApplication_TestAttribute(int attribute){
	return QCoreApplication::testAttribute(static_cast<Qt::ApplicationAttribute>(attribute));
}

char* QCoreApplication_QCoreApplication_Translate(char* context, char* sourceText, char* disambiguation, int n){
	return QCoreApplication::translate(const_cast<const char*>(context), const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8().data();
}

void QCoreApplication_DestroyQCoreApplication(QtObjectPtr ptr){
	static_cast<QCoreApplication*>(ptr)->~QCoreApplication();
}

