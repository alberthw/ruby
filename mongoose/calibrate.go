package mongoose

func SendStartCalibration() {
	sendMongooseCommand("cal.start")
}

func SendPrintCalibrationData() {
	sendMongooseCommand("cal.print")

}

func SendCalibrateSetIrms() {
	sendMongooseCommand("cal.setrms")
}

func SendCalibratedRMS(rms string) {
	sendMongooseCommand(rms)
}
