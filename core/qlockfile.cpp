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

QtObjectPtr QLockFile_NewQLockFile(char* fileName){
	return new QLockFile(QString(fileName));
}

int QLockFile_Error(QtObjectPtr ptr){
	return static_cast<QLockFile*>(ptr)->error();
}

int QLockFile_IsLocked(QtObjectPtr ptr){
	return static_cast<QLockFile*>(ptr)->isLocked();
}

int QLockFile_Lock(QtObjectPtr ptr){
	return static_cast<QLockFile*>(ptr)->lock();
}

int QLockFile_RemoveStaleLockFile(QtObjectPtr ptr){
	return static_cast<QLockFile*>(ptr)->removeStaleLockFile();
}

void QLockFile_SetStaleLockTime(QtObjectPtr ptr, int staleLockTime){
	static_cast<QLockFile*>(ptr)->setStaleLockTime(staleLockTime);
}

int QLockFile_StaleLockTime(QtObjectPtr ptr){
	return static_cast<QLockFile*>(ptr)->staleLockTime();
}

int QLockFile_TryLock(QtObjectPtr ptr, int timeout){
	return static_cast<QLockFile*>(ptr)->tryLock(timeout);
}

void QLockFile_DestroyQLockFile(QtObjectPtr ptr){
	static_cast<QLockFile*>(ptr)->~QLockFile();
}

void QLockFile_Unlock(QtObjectPtr ptr){
	static_cast<QLockFile*>(ptr)->unlock();
}

