// +build !minimal

#define protected public
#define private public

#include "webview.h"
#include "_cgo_export.h"

#include <QtWebView>

void QtWebView_QtWebView_Initialize()
{
	QtWebView::initialize();
}

