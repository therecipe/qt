#include "qdbuspendingcall.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QDBusPendingCall>
#include "_cgo_export.h"

class MyQDBusPendingCall: public QDBusPendingCall {
public:
};

void* QDBusPendingCall_NewQDBusPendingCall(void* other){
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_Swap(void* ptr, void* other){
	static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr){
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

