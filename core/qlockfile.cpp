#include "qlockfile.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLockFile>
#include "_cgo_export.h"

class MyQLockFile: public QLockFile {
public:
};

void* QLockFile_NewQLockFile(char* fileName){
	return new QLockFile(QString(fileName));
}

int QLockFile_Error(void* ptr){
	return static_cast<QLockFile*>(ptr)->error();
}

int QLockFile_IsLocked(void* ptr){
	return static_cast<QLockFile*>(ptr)->isLocked();
}

int QLockFile_Lock(void* ptr){
	return static_cast<QLockFile*>(ptr)->lock();
}

int QLockFile_RemoveStaleLockFile(void* ptr){
	return static_cast<QLockFile*>(ptr)->removeStaleLockFile();
}

void QLockFile_SetStaleLockTime(void* ptr, int staleLockTime){
	static_cast<QLockFile*>(ptr)->setStaleLockTime(staleLockTime);
}

int QLockFile_StaleLockTime(void* ptr){
	return static_cast<QLockFile*>(ptr)->staleLockTime();
}

int QLockFile_TryLock(void* ptr, int timeout){
	return static_cast<QLockFile*>(ptr)->tryLock(timeout);
}

void QLockFile_DestroyQLockFile(void* ptr){
	static_cast<QLockFile*>(ptr)->~QLockFile();
}

void QLockFile_Unlock(void* ptr){
	static_cast<QLockFile*>(ptr)->unlock();
}

