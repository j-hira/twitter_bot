package bot

func exeFuncAtSpecifiedTime(f func(), hours []int, minutes []int) {
	if (len(hours) == 0 || contains(hours, currentTime.Hour())) && (len(minutes) == 0 || contains(minutes, currentTime.Minute())) {
		f()
	}
}
