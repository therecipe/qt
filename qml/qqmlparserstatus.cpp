#include "qqmlparserstatus.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QQmlParserStatus>
#include "_cgo_export.h"

class MyQQmlParserStatus: public QQmlParserStatus {
public:
};

void QQmlParserStatus_ClassBegin(void* ptr){
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(void* ptr){
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

