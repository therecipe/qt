package multimedia

//#include "multimedia.h"
import "C"
import ()

//QAudio::Error
type QAudio__Error int64

const (
	QAudio__NoError       = QAudio__Error(0)
	QAudio__OpenError     = QAudio__Error(1)
	QAudio__IOError       = QAudio__Error(2)
	QAudio__UnderrunError = QAudio__Error(3)
	QAudio__FatalError    = QAudio__Error(4)
)

//QAudio::Mode
type QAudio__Mode int64

const (
	QAudio__AudioInput  = QAudio__Mode(0)
	QAudio__AudioOutput = QAudio__Mode(1)
)

//QAudio::State
type QAudio__State int64

const (
	QAudio__ActiveState    = QAudio__State(0)
	QAudio__SuspendedState = QAudio__State(1)
	QAudio__StoppedState   = QAudio__State(2)
	QAudio__IdleState      = QAudio__State(3)
)
