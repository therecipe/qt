#ifdef __cplusplus
extern "C" {
#endif

void QPaintDevice_DestroyQPaintDevice(void* ptr);
int QPaintDevice_ColorCount(void* ptr);
int QPaintDevice_Depth(void* ptr);
int QPaintDevice_DevicePixelRatio(void* ptr);
int QPaintDevice_Height(void* ptr);
int QPaintDevice_HeightMM(void* ptr);
int QPaintDevice_LogicalDpiX(void* ptr);
int QPaintDevice_LogicalDpiY(void* ptr);
void* QPaintDevice_PaintEngine(void* ptr);
int QPaintDevice_PaintingActive(void* ptr);
int QPaintDevice_PhysicalDpiX(void* ptr);
int QPaintDevice_PhysicalDpiY(void* ptr);
int QPaintDevice_Width(void* ptr);
int QPaintDevice_WidthMM(void* ptr);

#ifdef __cplusplus
}
#endif