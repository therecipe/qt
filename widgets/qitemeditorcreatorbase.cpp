#include "qitemeditorcreatorbase.h"
#include <QModelIndex>
#include <QItemEditorCreator>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QItemEditorCreatorBase>
#include "_cgo_export.h"

class MyQItemEditorCreatorBase: public QItemEditorCreatorBase {
public:
};

void QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(QtObjectPtr ptr){
	static_cast<QItemEditorCreatorBase*>(ptr)->~QItemEditorCreatorBase();
}

QtObjectPtr QItemEditorCreatorBase_CreateWidget(QtObjectPtr ptr, QtObjectPtr parent){
	return static_cast<QItemEditorCreatorBase*>(ptr)->createWidget(static_cast<QWidget*>(parent));
}

