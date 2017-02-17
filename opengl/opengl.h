// +build !minimal

#pragma once

#ifndef GO_QTOPENGL_H
#define GO_QTOPENGL_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtOpenGL_PackedString { char* data; long long len; };
struct QtOpenGL_PackedList { void* data; long long len; };
int QGL_SingleBuffer_Type();
int QGL_NoDepthBuffer_Type();
int QGL_ColorIndex_Type();
int QGL_NoAlphaChannel_Type();
int QGL_NoAccumBuffer_Type();
int QGL_NoStencilBuffer_Type();
int QGL_NoStereoBuffers_Type();
int QGL_IndirectRendering_Type();
int QGL_NoOverlay_Type();
int QGL_NoSampleBuffers_Type();
int QGL_NoDeprecatedFunctions_Type();

#ifdef __cplusplus
}
#endif

#endif