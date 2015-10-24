#include "qmarginsf.h"
#include <QModelIndex>
#include <QMargins>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QMarginsF>
#include "_cgo_export.h"

class MyQMarginsF: public QMarginsF {
public:
};

QtObjectPtr QMarginsF_NewQMarginsF(){
	return new QMarginsF();
}

QtObjectPtr QMarginsF_NewQMarginsF3(QtObjectPtr margins){
	return new QMarginsF(*static_cast<QMargins*>(margins));
}

int QMarginsF_IsNull(QtObjectPtr ptr){
	return static_cast<QMarginsF*>(ptr)->isNull();
}

