//source: http://doc.qt.io/qt-5/qtmultimediawidgets-videowidget-example.html

package main

import (
	"fmt"
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/multimedia"
	"github.com/therecipe/qt/widgets"
)

var (
	videoPlayer     *widgets.QWidget
	mediaPlayer     *multimedia.QMediaPlayer
	playButton      *widgets.QPushButton
	positionsSlider *widgets.QSlider
	errorLabel      *widgets.QLabel
)

func main() {
	widgets.NewQApplication(len(os.Args), os.Args)

	var player = newVideoPlayer()
	player.Resize2(320, 240)
	player.Show()

	widgets.QApplication_Exec()
}

func newVideoPlayer() *widgets.QWidget {
	videoPlayer = widgets.NewQWidget(nil, 0)
	mediaPlayer = multimedia.NewQMediaPlayer(nil, multimedia.QMediaPlayer__VideoSurface)

	var (
		videoWidget = multimedia.NewQVideoWidget(nil)
		openButton  = widgets.NewQPushButton2("Open...", nil)
	)
	openButton.ConnectClicked(func(_ bool) { openFile() })

	playButton = widgets.NewQPushButton(nil)
	playButton.SetEnabled(false)
	playButton.SetIcon(videoPlayer.Style().StandardIcon(widgets.QStyle__SP_MediaPlay, nil, nil))
	playButton.ConnectClicked(func(_ bool) { play() })

	positionsSlider = widgets.NewQSlider2(core.Qt__Horizontal, nil)
	positionsSlider.SetRange(0, 0)
	positionsSlider.ConnectSliderMoved(setPosition)

	errorLabel = widgets.NewQLabel(nil, 0)
	errorLabel.SetTextFormat(core.Qt__RichText)
	errorLabel.SetSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Maximum)

	var controlLayout = widgets.NewQHBoxLayout()
	controlLayout.SetContentsMargins(0, 0, 0, 0)
	controlLayout.AddWidget(openButton, 0, 0)
	controlLayout.AddWidget(playButton, 0, 0)
	controlLayout.AddWidget(positionsSlider, 0, 0)

	var layout = widgets.NewQVBoxLayout()
	layout.AddWidget(videoWidget, 0, 0)
	layout.AddLayout(controlLayout, 0)
	layout.AddWidget(errorLabel, 0, 0)

	videoPlayer.SetLayout(layout)

	mediaPlayer.SetVideoOutput(videoWidget)
	mediaPlayer.ConnectStateChanged(mediaStateChanged)
	mediaPlayer.ConnectPositionChanged(positionChanged)
	mediaPlayer.ConnectDurationChanged(durationChanged)
	mediaPlayer.ConnectError2(handleError)

	return videoPlayer
}

func openFile() {
	errorLabel.SetText("")

	var fileName = widgets.QFileDialog_GetOpenFileName(nil, "Open Movie", core.QDir_HomePath(), "", "", 0)
	if fileName != "" {
		mediaPlayer.SetMedia(multimedia.NewQMediaContent2(core.QUrl_FromLocalFile(fileName)), nil)
		playButton.SetEnabled(true)
	}
}

func play() {
	switch mediaPlayer.State() {
	case multimedia.QMediaPlayer__PlayingState:
		{
			mediaPlayer.Pause()
		}

	default:
		{
			mediaPlayer.Play()
		}
	}
}

func mediaStateChanged(state multimedia.QMediaPlayer__State) {
	switch state {
	case multimedia.QMediaPlayer__PlayingState:
		{
			playButton.SetIcon(videoPlayer.Style().StandardIcon(widgets.QStyle__SP_MediaPause, nil, nil))
		}

	default:
		{
			playButton.SetIcon(videoPlayer.Style().StandardIcon(widgets.QStyle__SP_MediaPlay, nil, nil))
		}
	}
}

func positionChanged(position int64) {
	positionsSlider.SetValue(int(position))
}

func durationChanged(duration int64) {
	positionsSlider.SetRange(0, int(duration))
}

func setPosition(position int) {
	mediaPlayer.SetPosition(int64(position))
}

func handleError(err multimedia.QMediaPlayer__Error) {
	playButton.SetEnabled(false)

	var errString = mediaPlayer.ErrorString()

	if errString == "" {
		switch err {
		case multimedia.QMediaPlayer__NoError:
			{
				errString = "QMediaPlayer::NoError"
			}

		case multimedia.QMediaPlayer__ResourceError:
			{
				errString = "QMediaPlayer::ResourceError"
			}

		case multimedia.QMediaPlayer__FormatError:
			{
				errString = "QMediaPlayer::FormatError"
			}

		case multimedia.QMediaPlayer__NetworkError:
			{
				errString = "QMediaPlayer::NetworkError"
			}

		case multimedia.QMediaPlayer__AccessDeniedError:
			{
				errString = "QMediaPlayer::AccessDeniedError"
			}

		case multimedia.QMediaPlayer__ServiceMissingError:
			{
				errString = "QMediaPlayer::ServiceMissingError"
			}
		}
	}

	errorLabel.SetText(fmt.Sprintf("File: %v<br>Error: %v<br>Supported MIME-Types: %v %v %v %v %v %v",
		mediaPlayer.CurrentMedia().CanonicalUrl().ToString(0),
		errString,
		hasSupport("video/avi"),
		hasSupport("video/mp4"),
		hasSupport("video/x-flv"),
		hasSupport("video/3gpp"),
		hasSupport("video/x-ms-wmv"),
		hasSupport("video/x-matroska")))
}

func hasSupport(mimeType string) string {
	if multimedia.QMediaPlayer_HasSupport(mimeType, make([]string, 0), 0) >= multimedia.QMultimedia__MaybeSupported {
		return mimeType + "=<font color=\"green\">true</font>"
	}
	return mimeType + "=<font color=\"red\">false</font>"
}
