#ifdef __cplusplus
extern "C" {
#endif
#include "../cgoutil.h"

void QPaintDevice_DestroyQPaintDevice(QtObjectPtr ptr);
int QPaintDevice_ColorCount(QtObjectPtr ptr);
int QPaintDevice_Depth(QtObjectPtr ptr);
int QPaintDevice_DevicePixelRatio(QtObjectPtr ptr);
int QPaintDevice_Height(QtObjectPtr ptr);
int QPaintDevice_HeightMM(QtObjectPtr ptr);
int QPaintDevice_LogicalDpiX(QtObjectPtr ptr);
int QPaintDevice_LogicalDpiY(QtObjectPtr ptr);
QtObjectPtr QPaintDevice_PaintEngine(QtObjectPtr ptr);
int QPaintDevice_PaintingActive(QtObjectPtr ptr);
int QPaintDevice_PhysicalDpiX(QtObjectPtr ptr);
int QPaintDevice_PhysicalDpiY(QtObjectPtr ptr);
int QPaintDevice_Width(QtObjectPtr ptr);
int QPaintDevice_WidthMM(QtObjectPtr ptr);

#ifdef __cplusplus
}
#endif