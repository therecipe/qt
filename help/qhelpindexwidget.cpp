#include "qhelpindexwidget.h"
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpIndexWidget>
#include "_cgo_export.h"

class MyQHelpIndexWidget: public QHelpIndexWidget {
public:
};

void QHelpIndexWidget_ActivateCurrentItem(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(void* ptr, char* filter, char* wildcard){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

