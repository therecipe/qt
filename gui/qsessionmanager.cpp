#include "qsessionmanager.h"
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QSessionManager>
#include "_cgo_export.h"

class MyQSessionManager: public QSessionManager {
public:
};

int QSessionManager_RestartHint(void* ptr){
	return static_cast<QSessionManager*>(ptr)->restartHint();
}

char* QSessionManager_SessionKey(void* ptr){
	return static_cast<QSessionManager*>(ptr)->sessionKey().toUtf8().data();
}

int QSessionManager_AllowsErrorInteraction(void* ptr){
	return static_cast<QSessionManager*>(ptr)->allowsErrorInteraction();
}

int QSessionManager_AllowsInteraction(void* ptr){
	return static_cast<QSessionManager*>(ptr)->allowsInteraction();
}

void QSessionManager_Cancel(void* ptr){
	static_cast<QSessionManager*>(ptr)->cancel();
}

char* QSessionManager_DiscardCommand(void* ptr){
	return static_cast<QSessionManager*>(ptr)->discardCommand().join("|").toUtf8().data();
}

int QSessionManager_IsPhase2(void* ptr){
	return static_cast<QSessionManager*>(ptr)->isPhase2();
}

void QSessionManager_Release(void* ptr){
	static_cast<QSessionManager*>(ptr)->release();
}

void QSessionManager_RequestPhase2(void* ptr){
	static_cast<QSessionManager*>(ptr)->requestPhase2();
}

char* QSessionManager_RestartCommand(void* ptr){
	return static_cast<QSessionManager*>(ptr)->restartCommand().join("|").toUtf8().data();
}

char* QSessionManager_SessionId(void* ptr){
	return static_cast<QSessionManager*>(ptr)->sessionId().toUtf8().data();
}

void QSessionManager_SetDiscardCommand(void* ptr, char* command){
	static_cast<QSessionManager*>(ptr)->setDiscardCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetManagerProperty2(void* ptr, char* name, char* value){
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value));
}

void QSessionManager_SetManagerProperty(void* ptr, char* name, char* value){
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartCommand(void* ptr, char* command){
	static_cast<QSessionManager*>(ptr)->setRestartCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartHint(void* ptr, int hint){
	static_cast<QSessionManager*>(ptr)->setRestartHint(static_cast<QSessionManager::RestartHint>(hint));
}

