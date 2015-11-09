#include "qmediacontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMediaControl>
#include "_cgo_export.h"

class MyQMediaControl: public QMediaControl {
public:
};

void QMediaControl_DestroyQMediaControl(void* ptr){
	static_cast<QMediaControl*>(ptr)->~QMediaControl();
}

