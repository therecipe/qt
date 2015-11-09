#ifdef __cplusplus
extern "C" {
#endif

void* QPaintEngineState_BackgroundBrush(void* ptr);
int QPaintEngineState_BackgroundMode(void* ptr);
void* QPaintEngineState_Brush(void* ptr);
int QPaintEngineState_BrushNeedsResolving(void* ptr);
int QPaintEngineState_ClipOperation(void* ptr);
void* QPaintEngineState_ClipRegion(void* ptr);
int QPaintEngineState_CompositionMode(void* ptr);
int QPaintEngineState_IsClipEnabled(void* ptr);
double QPaintEngineState_Opacity(void* ptr);
void* QPaintEngineState_Painter(void* ptr);
int QPaintEngineState_PenNeedsResolving(void* ptr);
int QPaintEngineState_RenderHints(void* ptr);

#ifdef __cplusplus
}
#endif