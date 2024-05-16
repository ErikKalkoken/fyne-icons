package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type iconInfo struct {
	name string
	icon fyne.Resource
}

func main() {
	a := app.New()
	w := a.NewWindow("Fyne icons")

	b := loadIcons()

	page := container.NewVBox()
	for _, g := range b {
		row := container.NewGridWrap(fyne.Size{Width: 150, Height: 75})
		for _, i := range g {
			label := widget.NewLabel(i.name)
			label.Alignment = fyne.TextAlignCenter
			icon := canvas.NewImageFromResource(i.icon)
			icon.FillMode = canvas.ImageFillContain
			element := container.NewBorder(nil, label, nil, nil, icon)
			row.Add(element)
		}
		page.Add(row)
		page.Add(widget.NewLabel(""))
	}

	w.SetContent(container.NewVScroll(page))
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

func loadIcons() [][]iconInfo {
	return [][]iconInfo{
		{
			{"CancelIcon", theme.CancelIcon()},
			{"ConfirmIcon", theme.ConfirmIcon()},
			{"DeleteIcon", theme.DeleteIcon()},
			{"SearchIcon", theme.SearchIcon()},
			{"SearchReplaceIcon", theme.SearchReplaceIcon()},
		}, {

			{"CheckButtonIcon", theme.CheckButtonIcon()},
			{"CheckButtonCheckedIcon", theme.CheckButtonCheckedIcon()},
			{"RadioButtonIcon", theme.RadioButtonIcon()},
			{"RadioButtonCheckedIcon", theme.RadioButtonCheckedIcon()},
		}, {
			{"ColorAchromaticIcon", theme.ColorAchromaticIcon()},
			{"ColorChromaticIcon", theme.ColorChromaticIcon()},
			{"ColorPaletteIcon", theme.ColorPaletteIcon()},
		}, {
			{"ContentAddIcon", theme.ContentAddIcon()},
			{"ContentRemoveIcon", theme.ContentRemoveIcon()},
			{"ContentClearIcon", theme.ContentClearIcon()},
			{"ContentCutIcon", theme.ContentCutIcon()},
			{"ContentCopyIcon", theme.ContentCopyIcon()},
			{"ContentPasteIcon", theme.ContentPasteIcon()},
			{"ContentRedoIcon", theme.ContentRedoIcon()},
			{"ContentUndoIcon", theme.ContentUndoIcon()},
		}, {
			{"InfoIcon", theme.InfoIcon()},
			{"ErrorIcon", theme.ErrorIcon()},
			{"QuestionIcon", theme.QuestionIcon()},
			{"WarningIcon", theme.WarningIcon()},
			{"BrokenImageIcon", theme.BrokenImageIcon()},
		}, {
			{"DocumentIcon", theme.DocumentIcon()},
			{"DocumentCreateIcon", theme.DocumentCreateIcon()},
			{"DocumentPrintIcon", theme.DocumentPrintIcon()},
			{"DocumentSaveIcon", theme.DocumentSaveIcon()},
		}, {
			{"FileIcon", theme.FileIcon()},
			{"FileApplicationIcon", theme.FileApplicationIcon()},
			{"FileAudioIcon", theme.FileAudioIcon()},
			{"FileImageIcon", theme.FileImageIcon()},
			{"FileTextIcon", theme.FileTextIcon()},
			{"FileVideoIcon", theme.FileVideoIcon()},
			{"FolderIcon", theme.FolderIcon()},
			{"FolderNewIcon", theme.FolderNewIcon()},
			{"FolderOpenIcon", theme.FolderOpenIcon()},
			{"ComputerIcon", theme.ComputerIcon()},
			{"HomeIcon", theme.HomeIcon()},
			{"HelpIcon", theme.HelpIcon()},
			{"HistoryIcon", theme.HistoryIcon()},
			{"SettingsIcon", theme.SettingsIcon()},
			{"StorageIcon", theme.StorageIcon()},
			{"DownloadIcon", theme.DownloadIcon()},
			{"UploadIcon", theme.UploadIcon()},
		}, {
			{"ViewFullScreenIcon", theme.ViewFullScreenIcon()},
			{"ViewRestoreIcon", theme.ViewRestoreIcon()},
			{"ViewRefreshIcon", theme.ViewRefreshIcon()},
			{"VisibilityIcon", theme.VisibilityIcon()},
			{"VisibilityOffIcon", theme.VisibilityOffIcon()},
			{"ViewZoomFitIcon", theme.ZoomFitIcon()},
			{"ViewZoomInIcon", theme.ZoomInIcon()},
			{"ViewZoomOutIcon", theme.ZoomOutIcon()},
		}, {
			{"MoreHorizontalIcon", theme.MoreHorizontalIcon()},
			{"MoreVerticalIcon", theme.MoreVerticalIcon()},
		}, {
			{"MoveDownIcon", theme.MoveDownIcon()},
			{"MoveUpIcon", theme.MoveUpIcon()},
		}, {
			{"NavigateBackIcon", theme.NavigateBackIcon()},
			{"NavigateNextIcon", theme.NavigateNextIcon()},
		}, {
			{"Menu", theme.MenuIcon()},
			{"MenuExpand", theme.MenuExpandIcon()},
			{"MenuDropDown", theme.MenuDropDownIcon()},
			{"MenuDropUp", theme.MenuDropUpIcon()},
		}, {
			{"MailAttachmentIcon", theme.MailAttachmentIcon()},
			{"MailComposeIcon", theme.MailComposeIcon()},
			{"MailForwardIcon", theme.MailForwardIcon()},
			{"MailReplyIcon", theme.MailReplyIcon()},
			{"MailReplyAllIcon", theme.MailReplyAllIcon()},
			{"MailSendIcon", theme.MailSendIcon()},
		}, {
			{"MediaFastForward", theme.MediaFastForwardIcon()},
			{"MediaFastRewind", theme.MediaFastRewindIcon()},
			{"MediaPause", theme.MediaPauseIcon()},
			{"MediaPlay", theme.MediaPlayIcon()},
			{"MediaStop", theme.MediaStopIcon()},
			{"MediaRecord", theme.MediaRecordIcon()},
			{"MediaReplay", theme.MediaReplayIcon()},
			{"MediaSkipNext", theme.MediaSkipNextIcon()},
			{"MediaSkipPrevious", theme.MediaSkipPreviousIcon()},
		}, {
			{"VolumeDown", theme.VolumeDownIcon()},
			{"VolumeMute", theme.VolumeMuteIcon()},
			{"VolumeUp", theme.VolumeUpIcon()},
		}, {
			{"AccountIcon", theme.AccountIcon()},
			{"LoginIcon", theme.LoginIcon()},
			{"LogoutIcon", theme.LogoutIcon()},
		}, {
			{"ListIcon", theme.ListIcon()},
			{"GridIcon", theme.GridIcon()},
		}}
}
