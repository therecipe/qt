#include "qstyleplugin.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyle>
#include <QStylePlugin>
#include "_cgo_export.h"

class MyQStylePlugin: public QStylePlugin {
public:
};

QtObjectPtr QStylePlugin_Create(QtObjectPtr ptr, char* key){
	return static_cast<QStylePlugin*>(ptr)->create(QString(key));
}

void QStylePlugin_DestroyQStylePlugin(QtObjectPtr ptr){
	static_cast<QStylePlugin*>(ptr)->~QStylePlugin();
}

