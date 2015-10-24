#include "qpaintenginestate.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintEngine>
#include <QPaintEngineState>
#include "_cgo_export.h"

class MyQPaintEngineState: public QPaintEngineState {
public:
};

int QPaintEngineState_BackgroundMode(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->backgroundMode();
}

int QPaintEngineState_BrushNeedsResolving(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->brushNeedsResolving();
}

int QPaintEngineState_ClipOperation(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->clipOperation();
}

int QPaintEngineState_CompositionMode(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->compositionMode();
}

int QPaintEngineState_IsClipEnabled(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->isClipEnabled();
}

QtObjectPtr QPaintEngineState_Painter(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->painter();
}

int QPaintEngineState_PenNeedsResolving(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->penNeedsResolving();
}

int QPaintEngineState_RenderHints(QtObjectPtr ptr){
	return static_cast<QPaintEngineState*>(ptr)->renderHints();
}

