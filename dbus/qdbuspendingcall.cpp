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

QtObjectPtr QDBusPendingCall_NewQDBusPendingCall(QtObjectPtr other){
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_DestroyQDBusPendingCall(QtObjectPtr ptr){
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

