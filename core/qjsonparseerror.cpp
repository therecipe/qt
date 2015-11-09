#include "qjsonparseerror.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QJsonParseError>
#include "_cgo_export.h"

class MyQJsonParseError: public QJsonParseError {
public:
};

char* QJsonParseError_ErrorString(void* ptr){
	return static_cast<QJsonParseError*>(ptr)->errorString().toUtf8().data();
}

