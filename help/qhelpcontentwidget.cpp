#include "qhelpcontentwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QHelpContentWidget>
#include "_cgo_export.h"

class MyQHelpContentWidget: public QHelpContentWidget {
public:
};

void* QHelpContentWidget_IndexOf(void* ptr, void* link){
	return static_cast<QHelpContentWidget*>(ptr)->indexOf(*static_cast<QUrl*>(link)).internalPointer();
}

