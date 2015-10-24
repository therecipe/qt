#include "qgenericargument.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QGenericArgument>
#include "_cgo_export.h"

class MyQGenericArgument: public QGenericArgument {
public:
};

void QGenericArgument_Data(QtObjectPtr ptr){
	static_cast<QGenericArgument*>(ptr)->data();
}

