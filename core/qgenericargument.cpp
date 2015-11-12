#include "qgenericargument.h"
#include <QModelIndex>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QGenericArgument>
#include "_cgo_export.h"

class MyQGenericArgument: public QGenericArgument {
public:
};

void* QGenericArgument_Data(void* ptr){
	return static_cast<QGenericArgument*>(ptr)->data();
}

