//go:build !goexperiment.synctest

package parse

const globalUseSynctest bool = false

func synctestWait_LetAllOtherGoroFinish() {}

func bubbleOrNot(f func()) {
	f()
}
