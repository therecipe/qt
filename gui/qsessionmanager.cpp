#include "qsessionmanager.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QSessionManager>
#include "_cgo_export.h"

class MyQSessionManager: public QSessionManager {
public:
};

int QSessionManager_RestartHint(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->restartHint();
}

char* QSessionManager_SessionKey(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->sessionKey().toUtf8().data();
}

int QSessionManager_AllowsErrorInteraction(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->allowsErrorInteraction();
}

int QSessionManager_AllowsInteraction(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->allowsInteraction();
}

void QSessionManager_Cancel(QtObjectPtr ptr){
	static_cast<QSessionManager*>(ptr)->cancel();
}

char* QSessionManager_DiscardCommand(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->discardCommand().join("|").toUtf8().data();
}

int QSessionManager_IsPhase2(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->isPhase2();
}

void QSessionManager_Release(QtObjectPtr ptr){
	static_cast<QSessionManager*>(ptr)->release();
}

void QSessionManager_RequestPhase2(QtObjectPtr ptr){
	static_cast<QSessionManager*>(ptr)->requestPhase2();
}

char* QSessionManager_RestartCommand(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->restartCommand().join("|").toUtf8().data();
}

char* QSessionManager_SessionId(QtObjectPtr ptr){
	return static_cast<QSessionManager*>(ptr)->sessionId().toUtf8().data();
}

void QSessionManager_SetDiscardCommand(QtObjectPtr ptr, char* command){
	static_cast<QSessionManager*>(ptr)->setDiscardCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetManagerProperty2(QtObjectPtr ptr, char* name, char* value){
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value));
}

void QSessionManager_SetManagerProperty(QtObjectPtr ptr, char* name, char* value){
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartCommand(QtObjectPtr ptr, char* command){
	static_cast<QSessionManager*>(ptr)->setRestartCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartHint(QtObjectPtr ptr, int hint){
	static_cast<QSessionManager*>(ptr)->setRestartHint(static_cast<QSessionManager::RestartHint>(hint));
}

