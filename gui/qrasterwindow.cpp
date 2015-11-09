#include "qrasterwindow.h"
#include <QWindow>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QRasterWindow>
#include "_cgo_export.h"

class MyQRasterWindow: public QRasterWindow {
public:
};

void* QRasterWindow_NewQRasterWindow(void* parent){
	return new QRasterWindow(static_cast<QWindow*>(parent));
}

