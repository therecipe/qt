#include "qstaticplugin.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStaticPlugin>
#include "_cgo_export.h"

class MyQStaticPlugin: public QStaticPlugin {
public:
};

QtObjectPtr QStaticPlugin_Instance(QtObjectPtr ptr){
	return static_cast<QStaticPlugin*>(ptr)->instance();
}

