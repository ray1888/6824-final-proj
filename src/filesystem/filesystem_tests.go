package filesystem

import (
	"ad"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Functionality tests for a FileSystem go here.
// Functions in this file are NOT "unit tests" because they ill not be run by "go test" because
// this file ends in "_testS", plural and they have more than one argument.
// Instead, these functionality tests can run against any class that implements the FileSystem interface
// by creating a unit test suite for your implementation class that calls these tests.

// @tests
// Whenever you add a new functionality test, be sure to add it to this list.
// This list is used in test_setup.go to run every functionality test on every difficulty.
var FunctionalityTests = []func(t *testing.T, fs FileSystem){
	TestBasicOpenClose,
	TestDeleteNotFound,
	TestCloseClosed,
	TestOpenOpened,
	TestOpenNotFound,
	TestOpenAlreadyExists,
	TestOpenROClose,
	TestOpenROClose4,
	TestOpenROClose64,
	TestOpenRWClose,
	TestOpenRWClose4,
	TestOpenRWClose64,
	TestOpenCloseLeastFD,
	TestOpenCloseDeleteMaxFD,
	TestOpenOffsetEqualsZero,
	TestOpenTruncate,
	TestOpenAppend,
	TestOpenCloseDeleteRoot,
	TestOpenCloseDeleteRootMax,
	TestOpenBlockNoContention,
	TestOpenBlockOneWaiting,
	TestOpenBlockMultipleWaiting,
	TestOpenBlockOnlyOne,
	TestSeekErrorBadFD,
	TestSeekErrorBadOffsetOperation,
	TestSeekOffEOF,
	TestWriteClosedFile,
	TestWriteReadBasic,
	TestWriteReadBasic4,
	TestCannotReadFromWriteOnly,
	TestCannotWriteToReadOnly,
	TestWriteSomeButNotAll,
	TestWrite1Byte,
	TestWrite8Bytes,
	TestWrite1KBytes,
	TestWrite1MBytes,
	TestWrite10MBytes64Kx160,
	TestWrite10MBytes128Kx80,
	TestWrite10MBytes256Kx40,
	TestWrite10MBytes128Kx80,
	TestWrite10MBytes256Kx40,
	TestWrite10MBytes512Kx20,
	TestWrite10MBytes1Mx10,
	TestWrite10MBytes10Mx1,
	TestReadClosedFile,
	TestRndWriteRead1ByteSimple,
	TestRndWriteRead8BytesSimple,
	TestRndWriteRead8BytesIter8,
	TestRndWriteRead8BytesIter64,
	TestRndWriteRead64BytesSimple,
	TestRndWriteRead6400BytesIter64K,
	TestRndWriteRead512KBIter1MB,
	TestRndWriteRead128KBIter10MB,
	TestMkdir,
	TestMkdirTree,
	TestOpenCloseDeleteAcrossDirectories,
	TestMkdirNotFound,
	TestMkdirAlreadyExists,
	TestRndWriteReadVerfiyHoleExpansion,
	TestDeleteCannotDeleteRootDir,
}

var testNames = []string{
	"TestBasicOpenClose",
	"TestDeleteNotFound",
	"TestCloseClosed",
	"TestOpenOpened",
	"TestOpenNotFound",
	"TestOpenAlreadyExists",
	"TestOpenROClose",
	"TestOpenROClose4",
	"TestOpenROClose64",
	"TestOpenRWClose",
	"TestOpenRWClose4",
	"TestOpenRWClose64",
	"TestOpenCloseLeastFD",
	"TestOpenCloseDeleteMaxFD",
	"TestOpenCloseDeleteRoot",
	"TestOpenCloseDeleteRootMax",
	"TestSeekErrorBadFD",
	"TestSeekErrorBadOffsetOperation",
	"TestSeekOffEOF",
	"TestWriteClosedFile",
	"TestWriteReadBasic",
	"TestWriteReadBasic4",
	"TestWrite1Byte",
	"TestWrite8Bytes",
	"TestWrite1KBytes",
	"TestWrite1MBytes",
	"TestWrite10MBytes64Kx160",
	"TestWrite10MBytes128Kx80",
	"TestWrite10MBytes256Kx40",
	"TestWrite10MBytes128Kx80",
	"TestWrite10MBytes256Kx40",
	"TestWrite10MBytes512Kx20",
	"TestWrite10MBytes1Mx10",
	"TestWrite10MBytes10Mx1",
	"TestReadClosedFile",
	"TestRndWriteRead1ByteSimple",
	"TestRndWriteRead8BytesSimple",
	"TestRndWriteRead8BytesIter8",
	"TestRndWriteRead8BytesIter64",
	"TestRndWriteRead6400BytesIter64K",
	"TestRndWriteRead512KBIter1MB",
	"TestRndWriteRead128KBIter10MB",
	"TestRndWriteReadVerfiyHoleExpansion",
}

func HelpVerifyBytes(t *testing.T, a []byte, b []byte, msg string) {
	ad.AssertExplainT(t, len(a) == len(b), "len(a) %d != len(b) %d", len(a), len(b))
	for i := 0; i < len(a); i++ {
		ad.AssertExplainT(t, a[i] == b[i], "%s ; i=%d (%d vs %d)",
			msg, i, a[i], b[i])
	}
}

func HelpMakeZeros(t *testing.T, n int) []byte {
	zeros := make([]byte, n)
	for i := 0; i < len(zeros); i++ {
		zeros[i] = 0
	}
	return zeros
}

// ===== BEGIN OPEN CLOSE DELETE HELPERS =====

func HelpDelete(t *testing.T, fs FileSystem, pathname string) {
	success, err := fs.Delete(pathname)
	ad.AssertExplainT(t, success && err == nil, "err %s deleting %s", err, pathname)
}

func HelpOpen(t *testing.T, fs FileSystem,
	path string, mode OpenMode, flags OpenFlags) int {
	fd, err := fs.Open(path, mode, flags)
	ad.AssertExplainT(t, err == nil, "Got error %s when trying to Open(%s, %s, %s)", err, path, mode, flags)
	ad.AssertExplainT(t, fd > 0, "Got negative fd %d when trying to Open(%s, %s, %s)", fd, path, mode, flags)
	return fd
}

func HelpTestOpenNotFound(t *testing.T, fs FileSystem,
	mode OpenMode, flags OpenFlags) {
	fd, err := fs.Open("/dirs/dont/exist/file", mode, flags)
	ad.AssertExplainT(t, err == NotFound, "1st case didnt produce error")
	ad.AssertExplainT(t, fd == -1, "fd should be negative on error")

	fd, err = fs.Open("f_wo_root_slash", mode, flags)
	ad.AssertExplainT(t, err == NotFound, "2nd case didnt produce error")
	ad.AssertExplainT(t, fd == -1, "fd should be negative on error")
}

func HelpClose(t *testing.T, fs FileSystem, fd int) {
	success, err := fs.Close(fd)
	ad.AssertExplainT(t, success && err == nil, "err closing fd %d", fd)
}

func HelpOpenClose(t *testing.T, fs FileSystem,
	path string, mode OpenMode, flags OpenFlags) {
	HelpClose(t, fs, HelpOpen(t, fs, path, mode, flags))
}

func HelpBatchOpen(t *testing.T, fs FileSystem,
	nFiles int, fmtStr string, mode OpenMode, flags OpenFlags) []int {
	fds := make([]int, nFiles)
	for ix := 0; ix < nFiles; ix++ {
		fds[ix] = HelpOpen(t, fs, fmt.Sprintf(fmtStr, ix), mode, flags)
	}
	return fds
}

func HelpBatchClose(t *testing.T, fs FileSystem, fds []int) {
	for ix := 0; ix < len(fds); ix++ {
		HelpClose(t, fs, fds[ix])
	}
}

func HelpBatchOpenClose(t *testing.T, fs FileSystem,
	nFiles int, fmtStr string, mode OpenMode, flags OpenFlags) {
	fds := HelpBatchOpen(t, fs, nFiles, fmtStr, mode, flags)
	HelpBatchClose(t, fs, fds)
}

func HelpBatchDelete(t *testing.T, fs FileSystem,
	nFiles int, fmtStr string) {
	for ix := 0; ix < nFiles; ix++ {
		HelpDelete(t, fs, fmt.Sprintf(fmtStr, ix))
	}
}

// ====== END OPEN CLOSE DELETE HELPERS =====

// ===== BEGIN MKDIR HELPERS =====

func HelpMkdir(t *testing.T, fs FileSystem, path string) {
	success, err := fs.Mkdir(path)
	ad.AssertNoErrorT(t, err)
	ad.AssertT(t, success)
	ad.AssertExplainT(t, success, "mkdir fail on %s", path)
}

// ===== END MKDIR HELPERS =====

// ===== BEGIN READ WRITE SEEK HELPERS =====

func HelpMakeRndBytes(t *testing.T, n int) []byte {
	rndBytes := make([]byte, n)
	num, err := rand.Read(rndBytes)
	ad.AssertExplainT(t, num == n, "mkbyte %d instead of %d", num, n)
	ad.AssertExplainT(t, err == nil, "mkbyte err %s", err)
	return rndBytes
}

func HelpSeek(t *testing.T, fs FileSystem,
	fd int, offset int, mode SeekMode) int {
	newPosition, err := fs.Seek(fd, offset, mode)
	ad.AssertNoErrorT(t, err)
	if mode == FromBeginning {
		ad.AssertEqualsT(t, offset, newPosition)
	}
	ad.AssertT(t, newPosition >= 0)
	return newPosition
}

func HelpRead(t *testing.T, fs FileSystem, fd int, numBytes int) (int, []byte) {
	numRead, data, err := fs.Read(fd, numBytes)
	ad.AssertNoErrorT(t, err)
	ad.AssertEqualsT(t, numBytes, numRead)
	return numRead, data
}

func HelpWriteBytes(t *testing.T, fs FileSystem, fd int, bytes []byte) int {
	numBytes := len(bytes)
	numWritten, err := fs.Write(fd, numBytes, bytes)
	ad.AssertNoErrorT(t, err)
	ad.AssertEqualsT(t, numBytes, numWritten)
	return numWritten
}

func HelpWriteString(t *testing.T, fs FileSystem, fd int, contents string) int {
	return HelpWriteBytes(t, fs, fd, []byte(contents))
}

func HelpReadWrite(t *testing.T, fs FileSystem, path string, contents string) int {
	fd := HelpOpen(t, fs, path, ReadWrite, Create)
	HelpSeek(t, fs, fd, 0, FromBeginning)
	nBytes := HelpWriteString(t, fs, fd, contents)
	ad.AssertExplainT(t, nBytes == len(contents), "%d bytes written vs %d", nBytes, len(contents))
	HelpSeek(t, fs, fd, 0, FromBeginning) //rewind to start reading
	nBytes, data := HelpRead(t, fs, fd, len(contents))
	for bite := 0; bite < len(contents); bite++ { //'byte' is reserved
		ad.AssertExplainT(t, data[bite] == contents[bite],
			"read data %s vs %s", data[bite], contents[bite])
	}
	return nBytes
}

// ===== BEGIN OPEN CLOSE TESTS ======

func TestBasicOpenClose(t *testing.T, fs FileSystem) {
	HelpOpenClose(t, fs, "/foo.txt", ReadWrite, Create)
}

func TestDeleteNotFound(t *testing.T, fs FileSystem) {
	success, err := fs.Delete("/this-file-doesnt-exist")
	ad.AssertExplainT(t, !success, "delete on missing file was successful")
	ad.AssertExplainT(t, err == NotFound, "err was not NotFound")
}

func TestCloseClosed(t *testing.T, fs FileSystem) {
	success, err := fs.Close(5) //arbirtary closed FD
	ad.AssertExplainT(t, !success, "close on closed FD was successful")
	ad.AssertExplainT(t, err == InactiveFD, "error needs to show issue with FD")
}

func TestOpenOpened(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/file-open-successfully", ReadWrite, Create)
	fd2, err := fs.Open("/file-open-successfully", ReadWrite, Create)
	ad.AssertExplainT(t, err == AlreadyOpen, "opened file returned err %s", err)
	ad.AssertExplainT(t, fd2 == -1, "-1 needed on open error")
	HelpClose(t, fs, fd)
	HelpDelete(t, fs, "/file-open-successfully")
}

func TestOpenNotFound(t *testing.T, fs FileSystem) {
	HelpTestOpenNotFound(t, fs, ReadWrite, Append)
	HelpTestOpenNotFound(t, fs, ReadWrite, Create)
	HelpTestOpenNotFound(t, fs, ReadWrite, Truncate)
	HelpTestOpenNotFound(t, fs, ReadOnly, Append)
	HelpTestOpenNotFound(t, fs, ReadOnly, Create)
	HelpTestOpenNotFound(t, fs, ReadOnly, Truncate)
}

func TestOpenAlreadyExists(t *testing.T, fs FileSystem) {
	_ = HelpOpen(t, fs, "/first_file", ReadWrite, Create)
	fd, err := fs.Open("/first_file", ReadWrite, Create)
	ad.AssertExplainT(t, err == AlreadyOpen, "wanted AlreadyOpen but err %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")
}

func TestOpenROClose(t *testing.T, fs FileSystem) {
	HelpOpenClose(t, fs, "/fooRO.txt", ReadOnly, Create)
}

func TestOpenROClose4(t *testing.T, fs FileSystem) {
	HelpBatchOpenClose(t, fs, 4, "/a-str-with-a-%d", ReadOnly, Create)
}

func TestOpenROClose64(t *testing.T, fs FileSystem) {
	HelpBatchOpenClose(t, fs, 64, "/str-2-with-a-%d", ReadOnly, Create)
}

func TestOpenRWClose(t *testing.T, fs FileSystem) {
	HelpOpenClose(t, fs, "/fooRO.txt", ReadOnly, Create)
}

func TestOpenRWClose4(t *testing.T, fs FileSystem) {
	HelpBatchOpenClose(t, fs, 4, "/str-3-with-a-%d", ReadWrite, Create)
}

func TestOpenRWClose64(t *testing.T, fs FileSystem) {
	HelpBatchOpenClose(t, fs, 64, "/str-4-with-a-%d", ReadWrite, Create)
}

func TestOpenCloseLeastFD(t *testing.T, fs FileSystem) {
	fd3A := HelpOpen(t, fs, "/A.txt", ReadWrite, Create)
	// Should be 3 because that's the lowest non-reserved non-active FD.
	ad.AssertEqualsT(t, 3, fd3A)
	HelpClose(t, fs, fd3A)

	fd3B := HelpOpen(t, fs, "/B.txt", ReadWrite, Create)
	// Should be 3 again because A.txt was closed, so FD=3 is now non-active again.
	ad.AssertEqualsT(t, 3, fd3B)
	// we're not closing it just yet

	fd4 := HelpOpen(t, fs, "/C.txt", ReadWrite, Create)
	// Should be 4 because 0-2 are reserved, 3 is taken, and 4 is next.
	ad.AssertEqualsT(t, 4, fd4)

	HelpClose(t, fs, fd3B)

	fd3C := HelpOpen(t, fs, "/D.txt", ReadWrite, Create)
	// B.txt was closed, so FD=3 is now non-active again.
	ad.AssertEqualsT(t, 3, fd3C)

	HelpClose(t, fs, fd3C)
	HelpClose(t, fs, fd4)
}

func TestOpenOffsetEqualsZero(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/foo.txt", ReadOnly, Create)
	// Seek to 0 from current just gets us the current position
	pos := HelpSeek(t, fs, fd, 0, FromCurrent)
	ad.AssertExplainT(t, pos == 0, "Newly created file had non-zero offset %d", pos)
}

func TestOpenTruncate(t *testing.T, fs FileSystem) {
	contentString := "Some arbitrary text here"
	fd := HelpOpen(t, fs, "/foo.txt", WriteOnly, Create)
	HelpWriteString(t, fs, fd, contentString)
	HelpClose(t, fs, fd)

	fd = HelpOpen(t, fs, "/foo.txt", ReadOnly, Truncate)
	offset := HelpSeek(t, fs, fd, 0, FromCurrent)
	ad.AssertExplainT(t, offset == 0, "Opened a file with Truncate, but got non-zero offset %d.", offset)
	bytesRead, _, err := fs.Read(fd, len([]byte(contentString)))
	ad.AssertExplainT(t, bytesRead == 0, "Able to read %d bytes from a file opened with Truncate, expected 0.", bytesRead)
	ad.AssertExplainT(t, err == nil, "Got non-nil error when trying to read fro mempty file: %s", err)

}

func TestOpenAppend(t *testing.T, fs FileSystem) {
	contentString := "Some arbitrary text here"
	contentLengthBytes := len([]byte(contentString))
	fd := HelpOpen(t, fs, "/foo.txt", WriteOnly, Create)
	HelpWriteString(t, fs, fd, contentString)
	HelpClose(t, fs, fd)

	fd = HelpOpen(t, fs, "/foo.txt", ReadOnly, Append)
	// Seek to 0 from current to get the current position.
	// Make sure it is the right distance from the beginning of the file.
	pos := HelpSeek(t, fs, fd, 0, FromCurrent)
	ad.AssertExplainT(t, pos == contentLengthBytes,
		"Opened a file for append and expected "+
			"the offset to be at the end of the file (position %d), but it was actually at position %d.")

	// Now make sure that position was actually at the end of the file.
	newPos := HelpSeek(t, fs, fd, 0, FromEnd)
	ad.AssertEqualsT(t, pos, newPos)
}

// open and close files checking all FDs open correctly up to limit,
// open a few past the limit, confirm we get errors, then close and delete all.
func TestOpenCloseDeleteMaxFD(t *testing.T, fs FileSystem) {
	maxFDCount := 128       //filesystem.MaxActiveFDs.
	maxFD := maxFDCount + 2 //max is offby1, & stdin, out, err...
	fmtStr := "/max-fd-%d.txt"
	prevFD := 0
	fds := make([]int, maxFDCount)
	for ix := 0; ix < maxFDCount; ix++ {
		fds[ix] = HelpOpen(t, fs, fmt.Sprintf(fmtStr, ix),
			ReadWrite, Create)
		ad.AssertExplainT(t, fds[ix] > prevFD, "%d -> ? %d", prevFD, fds[ix])
		prevFD = fds[ix]
	}
	ad.AssertExplainT(t, prevFD == maxFD,
		"wanted max FD %d but ended with %d", maxFD, prevFD)

	fd, err := fs.Open("/max-fd-one-more1.txt", ReadWrite, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "RW 1 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")
	fd, err = fs.Open("/max-fd-one-more2.txt", ReadWrite, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "RW 2 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")
	fd, err = fs.Open("/max-fd-one-more3.txt", ReadWrite, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "RW 3 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")

	fd, err = fs.Open("/max-fd-one-more1-ro.txt", ReadOnly, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "R0 4 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")
	fd, err = fs.Open("/max-fd-one-more2-ro.txt", ReadOnly, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "R0 5 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")
	fd, err = fs.Open("/max-fd-one-more3-ro.txt", ReadOnly, Create)
	ad.AssertExplainT(t, err == TooManyFDsOpen, "R0 6 past max opened err: %s", err)
	ad.AssertExplainT(t, fd == -1, "-1 needed on open error")

	HelpBatchClose(t, fs, fds)
	HelpBatchDelete(t, fs, maxFDCount, fmtStr)
}

func TestOpenCloseDeleteRoot(t *testing.T, fs FileSystem) {
	fd1 := HelpOpen(t, fs, "/foo", ReadWrite, Create)
	fd2 := HelpOpen(t, fs, "/bar", ReadWrite, Create)
	fd3 := HelpOpen(t, fs, "/baz", ReadWrite, Create)
	HelpClose(t, fs, fd1)
	HelpClose(t, fs, fd2)
	HelpClose(t, fs, fd3)
	HelpDelete(t, fs, "/foo")
	HelpDelete(t, fs, "/bar")
	HelpDelete(t, fs, "/baz")
}

//somewhat redundant with TestOpenCloseDeleteMaxFD
func TestOpenCloseDeleteRootMax(t *testing.T, fs FileSystem) {
	maxFD := 128
	fds := HelpBatchOpen(t, fs, maxFD, "/max-root-opens-%d", ReadWrite, Create)
	HelpBatchClose(t, fs, fds)
	HelpBatchDelete(t, fs, maxFD, "/max-root-opens-%d")
}

// How long to wait for things to happen before deciding
// that something was going to take forever.
const WaitTime = 3 * time.Second

func TestOpenBlockNoContention(t *testing.T, fs FileSystem) {
	// Make sure that if no one has a file open, an Open with the Block flag
	// succeeds immediately.

	// Multithreading is needed here to make sure that if it actually
	// does take forever, this test fails instead of hanging.
	openDone := make(chan int)
	go func() {
		fd := HelpOpen(t, fs, "/file.txt", ReadOnly, Create|Block)
		openDone <- fd
	}()

	select {
	case fd := <-openDone:
		// the open succeeded, we pass.
		HelpClose(t, fs, fd)
	case <-time.After(WaitTime):
		t.Fatalf("Blocking open with no contention took too long!")
	}
}

func TestOpenBlockOneWaiting(t *testing.T, fs FileSystem) {
	// Make sure that if one thread is waiting on a blocking open,
	// they get it when it is closed (and not before).
	fd := HelpOpen(t, fs, "/file.txt", ReadOnly, Create)

	// Start a blocking open in another thread.
	otherOpenFinished := make(chan int)
	go func() {
		fd2, _ := fs.Open("/file.txt", ReadOnly, Block)
		otherOpenFinished <- fd2
	}()

	// The file has not been closed, so the blocking open should not complete.
	select {
	case <-otherOpenFinished:
		t.Fatalf("Blocking open completed, but it shouldn't have!")
	case <-time.After(WaitTime):
		// Do nothing and proceed with the test.
	}

	HelpClose(t, fs, fd)

	// Now, blocking open should have finished.
	fd2 := <-otherOpenFinished
	HelpClose(t, fs, fd2)
}

func TestOpenBlockMultipleWaiting(t *testing.T, fs FileSystem) {
	// Make sure that if multiple people are waiting on a blocking open,
	// exactly on of them gets it.
	fd := HelpOpen(t, fs, "/file.txt", ReadOnly, Create)

	// Start a blocking open in two other threads.
	// This thread is later referred to as "thread #1" and these two are threads #2 and #3 respectively.
	secondOpenFinished := make(chan int)
	go func() {
		fd2 := HelpOpen(t, fs, "/file.txt", ReadOnly, Block)
		secondOpenFinished <- fd2
	}()
	thirdOpenFinished := make(chan int)
	go func() {
		fd3 := HelpOpen(t, fs, "/file.txt", ReadOnly, Block)
		thirdOpenFinished <- fd3
	}()

	// The file has not been closed, so the blocking open should not complete.
	select {
	case <-secondOpenFinished:
		t.Fatalf("Blocking open in thread #2 completed, but it shouldn't have!")
	case <-thirdOpenFinished:
		t.Fatalf("Blocking open in thread #3 completed, but it shouldn't have!")
	case <-time.After(WaitTime):
		// Do nothing and proceed with the test.
	}

	HelpClose(t, fs, fd)

	// Now, exactly one blocking open should have finished.
	var fdToClose int
	var whichSucceeded int
	select {
	case fdToClose = <-secondOpenFinished:
		whichSucceeded = 2
		select {
		case <-thirdOpenFinished:
			t.Fatalf("Both blocking opens completed!")
		case <-time.After(WaitTime):
			// #2 completed, #3 did not. This is okay.
		}
	case fdToClose = <-thirdOpenFinished:
		whichSucceeded = 3
		select {
		case <-secondOpenFinished:
			t.Fatalf("Both blocking opens completed!")
		case <-time.After(WaitTime):
			// #3 completed, #2 did not. This is okay.
		}
	case <-time.After(WaitTime):
		t.Fatalf("File was closed, but neither waiting block completed!")
	}

	// Now close whichever one succeeded and verify that the last one gets the file.
	HelpClose(t, fs, fdToClose)
	switch whichSucceeded {
	case 2:
		// now the third one should be done
		select {
		case fd3 := <-thirdOpenFinished:
			// Success!
			HelpClose(t, fs, fd3)
		case <-time.After(WaitTime):
			t.Fatalf("Threads #1 and #2 opened and closed a file, but #3 didn't open it!")
		}
	case 3:
		// now the second one should be done
		select {
		case fd2 := <-secondOpenFinished:
			// Success!
			HelpClose(t, fs, fd2)
		case <-time.After(WaitTime):
			t.Fatalf("Threads #1 and #3 opened and closed a file, but #2 didn't open it!")
		}
	default:
		panic("Something went wrong with the test, should never get here")
	}
}

func TestOpenBlockOnlyOne(t *testing.T, fs FileSystem) {
	// Tests another potential bug that could let a fle be opened in multiple places.
	// Makes sure that if a file is closed twice, you can't have two people opening it.
	filename := "/file.txt"
	fd := HelpOpen(t, fs, filename, ReadOnly, Create)
	HelpClose(t, fs, fd)
	fd = HelpOpen(t, fs, filename, ReadOnly, 0)
	HelpClose(t, fs, fd)

	// Should succeed.
	fd = HelpOpen(t, fs, filename, ReadOnly, Block)
	// But now another attempt to open should block.

	otherOpenFinished := make(chan int)
	go func() {
		fd2, _ := fs.Open(filename, ReadOnly, Block)
		otherOpenFinished <- fd2
	}()
	select {
	case <-otherOpenFinished:
		t.Fatalf("Blocking open completed, but it shouldn't have!")
	case <-time.After(WaitTime):
		// We passed
	}

	HelpClose(t, fs, fd)
}

// ===== END OPEN CLOSE TESTS =====

// ===== BEGIN OPEN CLOSE SEEK & DELETE TESTS =====

func TestSeekErrorBadFD(t *testing.T, fs FileSystem) {
	// must open an invalid FD
	_, err := fs.Seek(123456, 0, FromBeginning)
	ad.AssertEqualsT(t, err, InactiveFD)
}

func TestSeekErrorBadOffsetOperation(t *testing.T, fs FileSystem) {
	filename := "/bad-offset-operation.txt"
	fd := HelpOpen(t, fs, filename, ReadWrite, Create)

	_, err := fs.Seek(fd, -1, FromBeginning)
	ad.AssertExplainT(t, err == IllegalArgument,
		"Attempted to Seek to before the beginning of a file. "+
			"Expected IllegalArgument error, got %v", err)
}

func TestSeekOffEOF(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/seek-eof.txt", ReadWrite, Create)
	_, err := fs.Seek(fd, -1, FromBeginning) // can't be negative
	ad.AssertExplainT(t, err == IllegalArgument, "illegal offset err %s", err)

	HelpSeek(t, fs, fd, 0, FromEnd)     // valid - at byte 0
	HelpSeek(t, fs, fd, 0, FromCurrent) // valid - at byte 0

	HelpSeek(t, fs, fd, 1, FromBeginning) // valid - off end of file at byte 1
	HelpSeek(t, fs, fd, 2, FromBeginning) // valid - off end of file at byte 2
	HelpSeek(t, fs, fd, 3, FromBeginning) // valid - off end of file at byte 3

	HelpSeek(t, fs, fd, 3, FromBeginning) // valid - off end of file at byte 3
	HelpSeek(t, fs, fd, 2, FromBeginning) // valid - off end of file at byte 2
	HelpSeek(t, fs, fd, 1, FromBeginning) // valid - off end of file at byte 1
	HelpSeek(t, fs, fd, 0, FromBeginning) // valid - at byte 0

	HelpDelete(t, fs, "/seek-eof.txt")
}

// ===== BEGIN ITERATIVE WRITE CHUNK TESTS EXPANDING FILES =====

func TestWriteClosedFile(t *testing.T, fs FileSystem) {
	n, err := fs.Write(555, 5, HelpMakeRndBytes(t, 5)) //must be uninit
	ad.AssertExplainT(t, err == InactiveFD, "err %s", err)
	ad.AssertExplainT(t, n == -1, "wr %d", n)
}

func TestWriteNBytesIter(t *testing.T, fs FileSystem, fileName string, nBytes int, iters int) {
	fd := HelpOpen(t, fs, fileName, ReadWrite, Create)
	data := make([]byte, 0)
	for i := 0; i < iters; i++ {
		data = HelpMakeRndBytes(t, nBytes)
		ad.AssertExplainT(t, len(data) == nBytes, "made %d len array", len(data))
		n, err := fs.Write(fd, nBytes, data)
		ad.AssertExplainT(t, err == nil, "err %s", err)
		ad.AssertExplainT(t, n == nBytes, "wr %d", n)
	}
	HelpClose(t, fs, fd)
	HelpDelete(t, fs, fileName)
}

func TestWrite1Byte(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-1.txt", 1, 5)
}

func TestWrite8Bytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-8.txt", 8, 5)
}

func TestWrite1KBytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-1k.txt", 1000, 1)
}

func TestWrite1MBytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-1m.txt", 128000, 8)
}

func TestWrite10MBytes64Kx160(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-64Kx160.txt", 64*1000, 160)
}

func TestWrite10MBytes128Kx80(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-128Kx80.txt", 128*1000, 80)
}

func TestWrite10MBytes256Kx40(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-256Kx40.txt", 256*1000, 40)
}

func TestWrite10MBytes512Kx20(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-512Kx20.txt", 512*1000, 20)
}

func TestWrite10MBytes1Mx10(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-1Mx10.txt", 1000*1000, 10)
}

func TestWrite10MBytes10Mx1(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m-10Mx1.txt", 10*1000*1000, 1)
}

// ===== BEGIN ITERATIVE WRITE & READ CHUNK TESTS EXPANDING FILES =====

func TestReadClosedFile(t *testing.T, fs FileSystem) {
	n, data, err := fs.Read(555, 5) //must be uninit
	ad.AssertExplainT(t, err == InactiveFD, "err %s", err)
	ad.AssertExplainT(t, n == -1, "wr %d", n)
	ad.AssertExplainT(t, len(data) == 0, "no data should have been read")
}

func TestRndWriteReadNBytesIter(t *testing.T, fs FileSystem,
	fileName string, nBytes int, iters int) {
	fd := HelpOpen(t, fs, fileName, ReadWrite, Create)
	dataIn := make([]byte, 0)
	for i := 0; i < iters; i++ {
		//@dedup
		dataIn = HelpMakeRndBytes(t, nBytes)
		ad.AssertExplainT(t, len(dataIn) == nBytes, "made %d len array", len(dataIn))
		nWr, err := fs.Write(fd, nBytes, dataIn)
		ad.AssertExplainT(t, err == nil, "err %s", err)
		ad.AssertExplainT(t, nWr == nBytes, "wr %d", nWr)
		// seek back to beginning of write for the current chunk
		HelpSeek(t, fs, fd, 0+nBytes*i, FromBeginning)
		nRd, dataOut, err := fs.Read(fd, nBytes) //should put seek back at end of write block
		ad.AssertExplainT(t, err == nil, "err %s", err)
		ad.AssertExplainT(t, nRd == nBytes, "rd %d", nRd)
		for i := 0; i < len(dataIn); i++ {
			ad.AssertExplainT(t, dataIn[i] == dataOut[i],
				"data was corrupted at i=%d (%d vs %d)", i, dataIn[i], dataOut[i])
		}
	}
	HelpClose(t, fs, fd)
	HelpDelete(t, fs, fileName)
}

func TestWriteReadBasic(t *testing.T, fs FileSystem) {
	HelpReadWrite(t, fs, "/foo.txt", "bar")
}

func TestWriteReadBasic4(t *testing.T, fs FileSystem) {
	HelpReadWrite(t, fs, "/foo1.txt", "bar1")
	HelpReadWrite(t, fs, "/foo2.txt", "bar2")
	HelpReadWrite(t, fs, "/foo3.txt", "bar3")
	HelpReadWrite(t, fs, "/foo4.txt", "bar4")
}

func TestCannotReadFromWriteOnly(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/foo.txt", WriteOnly, Create)
	bytesRead, _, err := fs.Read(fd, 1) // Arbitrarily, as long as it's more than 0
	ad.AssertExplainT(t, err == WrongMode, "Expected a WrongMode error when attempting to read from a WriteOnly file. "+
		"Instead, got err=%v", err)
	ad.AssertEqualsT(t, -1, bytesRead)
	// if err is non-nil, data is unspecified, which is why we don't check it here.
}

func TestCannotWriteToReadOnly(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/foo.txt", ReadOnly, Create)
	content := []byte("arbitrary content here")
	bytesWritten, err := fs.Write(fd, len(content), content)
	ad.AssertExplainT(t, err == WrongMode, "Expected a WrongMode error when attempting to write to a ReadOnly file. "+
		"Instead, got err=%v", err)
	ad.AssertEqualsT(t, -1, bytesWritten)
}

func TestWriteSomeButNotAll(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/foo.txt", ReadWrite, Create)
	data := []byte("arbitrary content here")
	numBytesToWrite := 5                   // arbitrarily, but must be < len(data)
	ad.Assert(numBytesToWrite < len(data)) // or else the test won't test what it's supposed to test
	bytesWritten, err := fs.Write(fd, numBytesToWrite, data)
	ad.AssertExplainT(t, err == nil, "When writing %d bytes of a %d-byte data, got non-nil err %v.",
		numBytesToWrite, len(data), err)
	ad.AssertExplainT(t, bytesWritten == numBytesToWrite, "When writing %d bytes of a %d-byte data, actually wrote %d bytes",
		numBytesToWrite, len(data), bytesWritten)
}

// ====== BYTE LEVEL WRITE & READ CHUNK TESTS =====

func TestRndWriteRead1ByteSimple(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-8.txt", 1, 1)
}

func TestRndWriteRead8BytesSimple(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-8.txt", 8, 1)
}

func TestRndWriteRead8BytesIter8(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-8-iter-8.txt", 8, 8)
}

func TestRndWriteRead8BytesIter64(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-8-iter-64.txt", 8, 64)
}

func TestRndWriteRead64BytesSimple(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64-iter-1.txt", 64, 1)
}

func TestRndWriteRead6400BytesIter64K(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-1K.txt", 6400, 10)
}

func TestRndWriteRead512KBIter1MB(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-1M.txt", 512000, 2)
}

func TestRndWriteRead128KBIter10MB(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-10M.txt", 128000, 80)
}

func TestRndWriteReadVerfiyHoleExpansion(t *testing.T, fs FileSystem) {
	fd := HelpOpen(t, fs, "/few-holey-bytes.txt", ReadWrite, Create)
	nBytes := 64
	shore := HelpMakeRndBytes(t, nBytes)      // 64 random bytes
	HelpWriteBytes(t, fs, fd, shore)          // offset now 64
	HelpSeek(t, fs, fd, 0, FromBeginning)     // offset now 0
	_, shoreRd := HelpRead(t, fs, fd, nBytes) // offset now 64
	HelpVerifyBytes(t, shore, shoreRd, "shore integrity")

	// seek past EOF leaving 64 byte hole
	HelpSeek(t, fs, fd, 128, FromBeginning) // offset now 128 leaving 64 byte gap
	// write another few bytes to make an island
	nBytes = 128
	island := HelpMakeRndBytes(t, nBytes)
	fs.Write(fd, nBytes, island) // offset now 256 leaving 128 rnd byte island

	HelpSeek(t, fs, fd, 0, FromBeginning) // offset now 0
	nBytes = 64
	_, shoreRd = HelpRead(t, fs, fd, nBytes) // offset now 64 (beginning of hole)
	HelpVerifyBytes(t, shore, shoreRd, "shore deviation")

	// verify the zero filled hole appeared
	_, holeRd := HelpRead(t, fs, fd, nBytes) // offset now 128 (end of hole / start island)
	HelpVerifyBytes(t, HelpMakeZeros(t, 64), holeRd, "hole deviation")

	// verify island
	_, islandRd := HelpRead(t, fs, fd, 128) // offset now 256
	HelpVerifyBytes(t, island, islandRd, "island deviation")

	// verify read past EOF results in 0 bytes read
	nRd, data, err := fs.Read(fd, 1) // seek past EOF at 257
	ad.AssertExplainT(t, nRd == 0, "EOF read %d", nRd)
	ad.AssertExplainT(t, len(data) == 0, "EOF len(data) %d", len(data))
	ad.AssertExplainT(t, err == nil, "EOF read err %s", err)

	fs.Close(fd)
	fs.Delete("/few-holey-bytes.txt")
}

// ===== BEGIN WRITE READ HOLE TESTS =====

// ===== BEGIN MKDIR TESTS =====

func TestMkdir(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/a-dir-1")
	HelpMkdir(t, fs, "/a-dir-2")
	HelpMkdir(t, fs, "/a-dir-3")
	HelpMkdir(t, fs, "/a-dir-4")
}

func TestMkdirTree(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/grandparent")
	HelpMkdir(t, fs, "/grandparent/parent")
	HelpMkdir(t, fs, "/grandparent/parent/child")
	HelpMkdir(t, fs, "/grandparent/auntOrUncle")
	HelpMkdir(t, fs, "/grandparent/auntOrUncle/cousin")
}

func TestMkdirNotFound(t *testing.T, fs FileSystem) {
	success, err := fs.Mkdir("/foo/bar/baz")
	ad.AssertEqualsT(t, NotFound, err)
	ad.AssertT(t, !success)
}

func TestMkdirAlreadyExists(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/foo")
	success, err := fs.Mkdir("/foo")
	ad.AssertEqualsT(t, AlreadyExists, err)
	ad.AssertT(t, !success)
}

func TestOpenCloseDeleteAcrossDirectories(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/dir1")
	HelpMkdir(t, fs, "/dir2")
	HelpMkdir(t, fs, "/dir3")
	fd1 := HelpOpen(t, fs, "/dir1/foo", ReadWrite, Create)
	fd2 := HelpOpen(t, fs, "/dir2/bar", ReadWrite, Create)
	fd3 := HelpOpen(t, fs, "/dir3/baz", ReadWrite, Create)
	HelpClose(t, fs, fd1)
	HelpClose(t, fs, fd2)
	HelpClose(t, fs, fd3)
	HelpDelete(t, fs, "/dir1/foo")
	HelpDelete(t, fs, "/dir2/bar")
	HelpDelete(t, fs, "/dir3/baz")
	HelpDelete(t, fs, "/dir1")
	HelpDelete(t, fs, "/dir2")
	HelpDelete(t, fs, "/dir3")
}

func TestDeleteCannotDeleteRootDir(t *testing.T, fs FileSystem) {
	success, err := fs.Delete("/")
	ad.AssertExplainT(t, !success, "Attempted to delete the root directory of a filesystem, expected success=false, "+
		"got success=true.")
	ad.AssertExplainT(t, err == IllegalArgument, "Attempted to delete the root directory of a filesystem, expected "+
		"IllegalArgument error, got %s", err)
}
