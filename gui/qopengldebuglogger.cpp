#include "qopengldebuglogger.h"
#include <QOpenGLDebugMessage>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QOpenGLDebugLogger>
#include "_cgo_export.h"

class MyQOpenGLDebugLogger: public QOpenGLDebugLogger {
public:
};

QtObjectPtr QOpenGLDebugLogger_NewQOpenGLDebugLogger(QtObjectPtr parent){
	return new QOpenGLDebugLogger(static_cast<QObject*>(parent));
}

void QOpenGLDebugLogger_DisableMessages(QtObjectPtr ptr, int sources, int types, int severities){
	static_cast<QOpenGLDebugLogger*>(ptr)->disableMessages(static_cast<QOpenGLDebugMessage::Source>(sources), static_cast<QOpenGLDebugMessage::Type>(types), static_cast<QOpenGLDebugMessage::Severity>(severities));
}

void QOpenGLDebugLogger_EnableMessages(QtObjectPtr ptr, int sources, int types, int severities){
	static_cast<QOpenGLDebugLogger*>(ptr)->enableMessages(static_cast<QOpenGLDebugMessage::Source>(sources), static_cast<QOpenGLDebugMessage::Type>(types), static_cast<QOpenGLDebugMessage::Severity>(severities));
}

int QOpenGLDebugLogger_Initialize(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugLogger*>(ptr)->initialize();
}

int QOpenGLDebugLogger_IsLogging(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugLogger*>(ptr)->isLogging();
}

void QOpenGLDebugLogger_LogMessage(QtObjectPtr ptr, QtObjectPtr debugMessage){
	QMetaObject::invokeMethod(static_cast<QOpenGLDebugLogger*>(ptr), "logMessage", Q_ARG(QOpenGLDebugMessage, *static_cast<QOpenGLDebugMessage*>(debugMessage)));
}

int QOpenGLDebugLogger_LoggingMode(QtObjectPtr ptr){
	return static_cast<QOpenGLDebugLogger*>(ptr)->loggingMode();
}

void QOpenGLDebugLogger_PopGroup(QtObjectPtr ptr){
	static_cast<QOpenGLDebugLogger*>(ptr)->popGroup();
}

void QOpenGLDebugLogger_StartLogging(QtObjectPtr ptr, int loggingMode){
	QMetaObject::invokeMethod(static_cast<QOpenGLDebugLogger*>(ptr), "startLogging", Q_ARG(QOpenGLDebugLogger::LoggingMode, static_cast<QOpenGLDebugLogger::LoggingMode>(loggingMode)));
}

void QOpenGLDebugLogger_StopLogging(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QOpenGLDebugLogger*>(ptr), "stopLogging");
}

void QOpenGLDebugLogger_DestroyQOpenGLDebugLogger(QtObjectPtr ptr){
	static_cast<QOpenGLDebugLogger*>(ptr)->~QOpenGLDebugLogger();
}

