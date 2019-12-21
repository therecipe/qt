// +build !minimal

#pragma once

#ifndef GO_QTQUICKCONTROLS2_H
#define GO_QTQUICKCONTROLS2_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtQuickControls2_PackedString { char* data; long long len; void* ptr; };
struct QtQuickControls2_PackedList { void* data; long long len; };
void QQuickStyle_QQuickStyle_AddStylePath(struct QtQuickControls2_PackedString path);
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_AvailableStyles();
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Name();
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Path();
void QQuickStyle_QQuickStyle_SetFallbackStyle(struct QtQuickControls2_PackedString style);
void QQuickStyle_QQuickStyle_SetStyle(struct QtQuickControls2_PackedString style);
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_StylePathList();

#ifdef __cplusplus
}
#endif

#endif