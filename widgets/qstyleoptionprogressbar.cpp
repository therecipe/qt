#include "qstyleoptionprogressbar.h"
#include <QModelIndex>
#include <QStyle>
#include <QStyleOption>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QStyleOptionProgressBar>
#include "_cgo_export.h"

class MyQStyleOptionProgressBar: public QStyleOptionProgressBar {
public:
};

QtObjectPtr QStyleOptionProgressBar_NewQStyleOptionProgressBar(){
	return new QStyleOptionProgressBar();
}

QtObjectPtr QStyleOptionProgressBar_NewQStyleOptionProgressBar2(QtObjectPtr other){
	return new QStyleOptionProgressBar(*static_cast<QStyleOptionProgressBar*>(other));
}

