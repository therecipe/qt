// +build !minimal

#pragma once

#ifndef GO_QTWEBVIEW_H
#define GO_QTWEBVIEW_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtWebView_PackedString { char* data; long long len; };
struct QtWebView_PackedList { void* data; long long len; };
void QtWebView_QtWebView_Initialize();

#ifdef __cplusplus
}
#endif

#endif