#include "qpaintenginestate.h"
#include <QModelIndex>
#include <QBrush>
#include <QRegion>
#include <QPaintEngine>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QPaintEngineState>
#include "_cgo_export.h"

class MyQPaintEngineState: public QPaintEngineState {
public:
};

void* QPaintEngineState_BackgroundBrush(void* ptr){
	return new QBrush(static_cast<QPaintEngineState*>(ptr)->backgroundBrush());
}

int QPaintEngineState_BackgroundMode(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->backgroundMode();
}

void* QPaintEngineState_Brush(void* ptr){
	return new QBrush(static_cast<QPaintEngineState*>(ptr)->brush());
}

int QPaintEngineState_BrushNeedsResolving(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->brushNeedsResolving();
}

int QPaintEngineState_ClipOperation(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->clipOperation();
}

void* QPaintEngineState_ClipRegion(void* ptr){
	return new QRegion(static_cast<QPaintEngineState*>(ptr)->clipRegion());
}

int QPaintEngineState_CompositionMode(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->compositionMode();
}

int QPaintEngineState_IsClipEnabled(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->isClipEnabled();
}

double QPaintEngineState_Opacity(void* ptr){
	return static_cast<double>(static_cast<QPaintEngineState*>(ptr)->opacity());
}

void* QPaintEngineState_Painter(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->painter();
}

int QPaintEngineState_PenNeedsResolving(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->penNeedsResolving();
}

int QPaintEngineState_RenderHints(void* ptr){
	return static_cast<QPaintEngineState*>(ptr)->renderHints();
}

