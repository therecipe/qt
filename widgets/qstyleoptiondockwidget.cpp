#include "qstyleoptiondockwidget.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QStyleOption>
#include <QStyle>
#include <QStyleOptionDockWidget>
#include "_cgo_export.h"

class MyQStyleOptionDockWidget: public QStyleOptionDockWidget {
public:
};

QtObjectPtr QStyleOptionDockWidget_NewQStyleOptionDockWidget(){
	return new QStyleOptionDockWidget();
}

QtObjectPtr QStyleOptionDockWidget_NewQStyleOptionDockWidget2(QtObjectPtr other){
	return new QStyleOptionDockWidget(*static_cast<QStyleOptionDockWidget*>(other));
}

