#include "qcameraimageprocessingcontrol.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QCamera>
#include <QCameraImageProcessing>
#include <QCameraImageProcessingControl>
#include "_cgo_export.h"

class MyQCameraImageProcessingControl: public QCameraImageProcessingControl {
public:
};

int QCameraImageProcessingControl_IsParameterSupported(void* ptr, int parameter){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter));
}

int QCameraImageProcessingControl_IsParameterValueSupported(void* ptr, int parameter, void* value){
	return static_cast<QCameraImageProcessingControl*>(ptr)->isParameterValueSupported(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void* QCameraImageProcessingControl_Parameter(void* ptr, int parameter){
	return new QVariant(static_cast<QCameraImageProcessingControl*>(ptr)->parameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter)));
}

void QCameraImageProcessingControl_SetParameter(void* ptr, int parameter, void* value){
	static_cast<QCameraImageProcessingControl*>(ptr)->setParameter(static_cast<QCameraImageProcessingControl::ProcessingParameter>(parameter), *static_cast<QVariant*>(value));
}

void QCameraImageProcessingControl_DestroyQCameraImageProcessingControl(void* ptr){
	static_cast<QCameraImageProcessingControl*>(ptr)->~QCameraImageProcessingControl();
}

