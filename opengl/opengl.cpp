// +build !minimal

#define protected public
#define private public

#include "opengl.h"
#include "_cgo_export.h"

#include <QGL>

int QGL_SingleBuffer_Type()
{
	return QGL::SingleBuffer;
}

int QGL_NoDepthBuffer_Type()
{
	return QGL::NoDepthBuffer;
}

int QGL_ColorIndex_Type()
{
	return QGL::ColorIndex;
}

int QGL_NoAlphaChannel_Type()
{
	return QGL::NoAlphaChannel;
}

int QGL_NoAccumBuffer_Type()
{
	return QGL::NoAccumBuffer;
}

int QGL_NoStencilBuffer_Type()
{
	return QGL::NoStencilBuffer;
}

int QGL_NoStereoBuffers_Type()
{
	return QGL::NoStereoBuffers;
}

int QGL_IndirectRendering_Type()
{
	return QGL::IndirectRendering;
}

int QGL_NoOverlay_Type()
{
	return QGL::NoOverlay;
}

int QGL_NoSampleBuffers_Type()
{
	return QGL::NoSampleBuffers;
}

int QGL_NoDeprecatedFunctions_Type()
{
	return QGL::NoDeprecatedFunctions;
}

