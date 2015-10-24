#include "qscriptable.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QScriptable>
#include "_cgo_export.h"

class MyQScriptable: public QScriptable {
public:
};

int QScriptable_ArgumentCount(QtObjectPtr ptr){
	return static_cast<QScriptable*>(ptr)->argumentCount();
}

QtObjectPtr QScriptable_Context(QtObjectPtr ptr){
	return static_cast<QScriptable*>(ptr)->context();
}

QtObjectPtr QScriptable_Engine(QtObjectPtr ptr){
	return static_cast<QScriptable*>(ptr)->engine();
}

