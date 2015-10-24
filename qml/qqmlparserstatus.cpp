#include "qqmlparserstatus.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QQmlParserStatus>
#include "_cgo_export.h"

class MyQQmlParserStatus: public QQmlParserStatus {
public:
};

void QQmlParserStatus_ClassBegin(QtObjectPtr ptr){
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(QtObjectPtr ptr){
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

