#include "qrasterwindow.h"
#include <QModelIndex>
#include <QWindow>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QRasterWindow>
#include "_cgo_export.h"

class MyQRasterWindow: public QRasterWindow {
public:
};

QtObjectPtr QRasterWindow_NewQRasterWindow(QtObjectPtr parent){
	return new QRasterWindow(static_cast<QWindow*>(parent));
}

