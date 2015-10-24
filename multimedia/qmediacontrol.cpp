#include "qmediacontrol.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QString>
#include <QMediaControl>
#include "_cgo_export.h"

class MyQMediaControl: public QMediaControl {
public:
};

void QMediaControl_DestroyQMediaControl(QtObjectPtr ptr){
	static_cast<QMediaControl*>(ptr)->~QMediaControl();
}

