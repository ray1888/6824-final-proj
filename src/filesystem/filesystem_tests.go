package filesystem

import (
	"ad"
	"fmt"
	"math/rand"
	"testing"
)

// Functionality tests for a FileSystem go here.
// Functions in this file are NOT "unit tests" because they ill not be run by "go test" because
// this file ends in "_testS", plural and they have more than one argument.
// Instead, these functionality tests can run against any class that implements the FileSystem interface
// by creating a unit test suite for your implementation class that calls these tests.

func HelpVerifyBytes(t *testing.T, a []byte, b []byte, msg string) {
   ad.AssertExplainT(t, len(a) == len(b), "len(a) %d != len(b) %d", len(a), len(b))
   for i := 0; i < len(a); i++ {
	   ad.AssertExplainT(t, a[i] == b[i], msg, i, a[i], b[i])
   }
}
// ===== BEGIN OPEN CLOSE DELETE HELPERS =====

func HelpDelete(t *testing.T, fs FileSystem,
	pathname string) {
	success, err := fs.Delete(pathname)
	ad.AssertExplainT(t, success && err == nil, "err %s deleting %s", err, pathname)
}

func HelpOpen(t *testing.T, fs FileSystem,
	path string, mode OpenMode, flags OpenFlags) int {
	fd, err := fs.Open(path, mode, flags)
	ad.AssertExplainT(t, fd > 0 && err == nil, "err %s opening %s", err, path)
	return fd
}

func HelpTestOpenNotFound(t *testing.T, fs FileSystem,
	mode OpenMode, flags OpenFlags) {
	fd, err := fs.Open("/dirs/dont/exist/file", mode, flags)
	ad.AssertExplainT(t, err == NotFound, "1st case didnt produce error")
	ad.AssertExplainT(t, fd == -1, "fd should be negative on error")

	fd, err = fs.Open("f_wo_root_slash", mode, flags) // should we handle this?
	ad.AssertExplainT(t, err == NotFound, "2nd case didnt produce error")
	ad.AssertExplainT(t, fd == -1, "fd should be negative on error")
}

func HelpClose(t *testing.T, fs FileSystem,
	fd int) {
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

func HelpMkdir(t *testing.T, fs FileSystem,
	path string) {
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

// error checked helper
func HelpSeek(t *testing.T, fs FileSystem,
	fd int, offset int, mode SeekMode) int {
	newPosition, err := fs.Seek(fd, offset, mode)
	ad.AssertNoErrorT(t, err)
	if mode == FromBeginning {
		ad.AssertEqualsT(t, offset, newPosition)
	} // can we auto-check more seek behavior...
	return newPosition
}

// error checked helper
func HelpRead(t *testing.T, fs FileSystem, fd int, numBytes int) (int, []byte) {
	numRead, data, err := fs.Read(fd, numBytes)
	ad.AssertNoErrorT(t, err)
	ad.AssertEqualsT(t, numBytes, numRead)
	return numRead, data
}

// error checked helper
func HelpWrite(t *testing.T, fs FileSystem,
	fd int, nBytes int, bytes []byte) int {
	nWr, err := fs.Write(fd, nBytes, bytes)
	ad.AssertNoErrorT(t, err)
	ad.AssertEqualsT(t, nBytes, nWr)
	return nWr
}

// error checked helper
func HelpReadWrite(t *testing.T, fs FileSystem,
	path string, strData string) int {
   contents := []byte(strData)
	fd := HelpOpen(t, fs, path, ReadWrite, Create)
	HelpSeek(t, fs, fd, 0, FromBeginning)
	nBytes := HelpWrite(t, fs, fd, len(contents), contents)
	ad.AssertExplainT(t, nBytes == len(contents),
		"%d bytes written vs %d", nBytes, len(contents))
	HelpSeek(t, fs, fd, 0, FromBeginning) //rewind to start reading
	nBytes, data := HelpRead(t, fs, fd, len(contents))
	for bite := 0; bite < len(contents); bite++ { //byte is reserved
		ad.AssertExplainT(t, data[bite] == contents[bite],
			"read data %s vs %s", data[bite], contents[bite])
	}
	return nBytes
}

// @tests
// Whenever you add a new functionality test, be sure to add it to this list.
// This list is used in test_setup.go to run every functionality test on every difficulty.
var FunctionalityTests = []func(t *testing.T, fs FileSystem){
	TestHelpGenerateJenkinsPipeline,
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
	TestOpenCloseDeleteRoot,
	TestOpenCloseDeleteRootMax,
	TestSeekErrorBadFD,
	TestSeekErrorBadOffsetOperation,
	TestSeekOffEOF,
	TestWriteClosedFile,
	TestWriteReadBasic,
	TestWriteReadBasic4,
	TestWrite1Byte,
	TestWrite8Bytes,
	TestWrite1KBytes,
	TestWrite1MBytes,
	TestWrite10MBytes,
	TestWrite100MBytes,
	TestReadClosedFile,
	TestRndWriteRead1ByteSimple,
	TestRndWriteRead8BytesSimple,
	TestRndWriteRead8BytesIter8,
	TestRndWriteRead8BytesIter64,
	TestRndWriteRead64BytesIter64K,
	TestRndWriteRead64KBIter1MB,
	TestRndWriteRead64KBIter10MB,
	TestRndWriteRead1MBIter100MB,
   TestRndWriteReadFewHoleyBytes,
	// ========= the line in the sand =======
	//TestMkdir,
	//TestMkdirTree,
	//TestOpenCloseDeleteAcrossDirectories,
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
	"TestWrite10MBytes",
	"TestWrite100MBytes",
	"TestReadClosedFile",
	"TestRndWriteRead1ByteSimple",
	"TestRndWriteRead8BytesSimple",
	"TestRndWriteRead8BytesIter8",
	"TestRndWriteRead8BytesIter64",
	"TestRndWriteRead64BytesIter64K",
	"TestRndWriteRead64KBIter1MB",
	"TestRndWriteRead64KBIter10MB",
	"TestRndWriteRead1MBIter100MB",
	"TestMkdir",
	"TestMkdirTree",
	"TestOpenCloseDeleteAcrossDirectories",
}

// should save countless hours messing with Jenkinsfile.
// eventually could automate the generation and push of the entire Jenkinsfile.
func TestHelpGenerateJenkinsPipeline(t *testing.T, fs FileSystem) {
	moduleStr := "MemoryFS"
	//TODO timestamp the generation ... share with D's raft test case generation code eventually
	for i := 0; i < len(testNames); i++ {
		print(fmt.Sprintf(" stage('DB3 TestMemoryFS_%s') {\n", testNames[i]))
		print("\t\tsteps {\n")
		print("\t\t\tscript {\n")
		print(fmt.Sprintf("\t\t\t\tsh 'THA_GO_DEBUG=3 DFS_DEFAULT_DEBUG_LEVEL=3 "+ //ugly ... break string into local vars eventually
			"GO_TEST_PKG=Test%s_%s /volumes/babtin-volume/babtin/babtin/jenkins_dse_debug_optimized.sh 1 1 1 1'\n", moduleStr, testNames[i]))
		print("\t\t\t}\n")
		print("\t\t}\n")
		print("\t}\n")
	}
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

// open and close files checking all FDs open correctly up to limit,
// open a few past the limit, confirm we get errors, then close and delete all.
func TestOpenCloseDeleteMaxFD(t *testing.T, fs FileSystem) {
	maxFDCount := MaxActiveFDs
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

func TestOpenCloseDeleteRootMax(t *testing.T, fs FileSystem) {
	maxFD := 64 //XXX update once we set it!!
	fds := HelpBatchOpen(t, fs, 64, "/max-root-opens-%d", ReadWrite, Create)
	HelpBatchClose(t, fs, fds)
	HelpBatchDelete(t, fs, maxFD, "/max-root-opens-%d")
}

// TODO next is same set of tests involving subdirs

//  ================== the line in the sand ====================
//  keeps moving down as tests begin passing and stay passing!

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
	// Enforce only one option
	_, err := fs.Seek(fd, 0, -1)
	ad.AssertExplainT(t, err == IllegalArgument, "illegal seek mode wrong err")
	_, err = fs.Seek(fd, 0, 3)
	ad.AssertExplainT(t, err == IllegalArgument, "illegal seek mode wrong err")
	_, err = fs.Seek(fd, 0, 0)
	ad.AssertExplainT(t, err == nil, "illegal seek mode err")
	_, err = fs.Seek(fd, 0, 1)
	ad.AssertExplainT(t, err == nil, "illegal seek mode err")
	_, err = fs.Seek(fd, 0, 2)
	ad.AssertExplainT(t, err == nil, "illegal seek mode err")
	HelpClose(t, fs, fd)
	HelpDelete(t, fs, filename)
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

// TODO need a debug interface to simulate the test datastore runs out of space...
// TODO need a debug interface to simulate the test datastore has an IO error...

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
	TestWriteNBytesIter(t, fs, "/wr-1k.txt", 1000, 5)
}

func TestWrite1MBytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-1m.txt", 1000000, 5)
}

func TestWrite10MBytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-10m.txt", 10000000, 5)
}

func TestWrite100MBytes(t *testing.T, fs FileSystem) {
	TestWriteNBytesIter(t, fs, "/wr-100m.txt", 100000000, 3)
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

func TestRndWriteRead64BytesIter64K(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-1K.txt", 64, 1000)
}

func TestRndWriteRead64KBIter1MB(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-10M.txt", 6400, 160)
}

func TestRndWriteRead64KBIter10MB(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-64k-iter-100M.txt", 6400, 1600)
}

func TestRndWriteRead1MBIter100MB(t *testing.T, fs FileSystem) {
	TestRndWriteReadNBytesIter(t, fs, "/r-1MB-iter-100M.txt", 1000000, 10)
}


func TestRndWriteReadFewHoleyBytes(t *testing.T, fs FileSystem) {
   fd := HelpOpen(t, fs, "/few-holey-bytes.txt", ReadWrite, Create)
   nBytes := 64
   shore := HelpMakeRndBytes(t, nBytes)                                         // 64 random bytes
   HelpWrite(t, fs, fd, nBytes, shore)                                          // offset now 64
   HelpSeek(t, fs, fd, 0, FromBeginning)                                        // offset now 0
   _, shoreRd := HelpRead(t, fs, fd, nBytes)                                    // offset now 64
   HelpVerifyBytes(t, shore, shoreRd, "shore integrity; i=%d (%d vs %d)")

                                                                                //                 64 byte hole 128                      256
   // seek past EOF                                                             // ----- 64 bytes |              | island                 |
   nBytes = 128
   HelpSeek(t, fs, fd, 64, FromBeginning)                                       // offset now 128 leaving 64 byte gap
   // write another few bytes to make an island
   island := HelpMakeRndBytes(t, nBytes)
   fs.Write(fd, nBytes, island)                                                 // offset now 256 leaving 128 rnd byte island
   HelpSeek(t, fs, fd, 0, FromBeginning)                                        // offset now 0   

   nBytes = 64
   _, shoreRd = HelpRead(t, fs, fd, nBytes)                                     // offset now 64 (beginning of hole)
   HelpVerifyBytes(t, shore, island, "1st island deviation")
   // verify the zero filled hole appeared
   _, holeRd := HelpRead(t, fs, fd, nBytes)                                     // offset now 128 (end of hole / beginning of island)
   HelpVerifyBytes(t, make([]byte, nBytes), holeRd, "hole deviation")

   // verify island
   _, islandRd := HelpRead(t, fs, fd, 64)                                       // offset now 128 
   HelpVerifyBytes(t, island, islandRd, "2nd island deviation")

   // PUT IN OWN TEST EVENTUALLY
   // verify read past EOF results in 0 bytes read
   nRd, data := HelpRead(t, fs, fd, 1)                                          // seek past EOF at 257
   ad.AssertExplainT(t, nRd == 0, "EOF read %d", nRd)
   ad.AssertExplainT(t, len(data) == 0, "EOF len(data) %d", len(data))

   fs.Close(fd)
   fs.Delete("/few-holey-bytes.txt")
}


// ===== THE LINE IN THE SAND ==========
// TOMORROW: Wed 5/9/18

// ===== BEGIN WRITE READ HOLE TESTS =====

// TODO these will start writing larger amounts of data
// TODO the next set of tests will create holes, seek around, fill files with data

// ===== BEGIN MKDIR TESTS =====
// TODO subdirectories next...

func TestMkdir(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/a-dir-1")
	HelpMkdir(t, fs, "/a-dir-2")
	HelpMkdir(t, fs, "/a-dir-3")
	HelpMkdir(t, fs, "/a-dir-4")
	HelpDelete(t, fs, "/a-dir-1")
	HelpDelete(t, fs, "/a-dir-2")
	HelpDelete(t, fs, "/a-dir-3")
	HelpDelete(t, fs, "/a-dir-4")
}

func TestMkdirTree(t *testing.T, fs FileSystem) {
	HelpMkdir(t, fs, "/a-dir-1")
	HelpMkdir(t, fs, "/a-dir-2")
	HelpMkdir(t, fs, "/a-dir-3")
	HelpMkdir(t, fs, "/a-dir-4")
	HelpMkdir(t, fs, "/a-dir-1/sub1")
	HelpMkdir(t, fs, "/a-dir-2/sub2/sub3")
	HelpMkdir(t, fs, "/a-dir-3/sub2/sub3/sub4/sub5")
	HelpMkdir(t, fs, "/a-dir-4/sub2/sub3/sub4/sub5/sub6")
	HelpDelete(t, fs, "/a-dir-1")
	HelpDelete(t, fs, "/a-dir-2")
	HelpDelete(t, fs, "/a-dir-3")
	HelpDelete(t, fs, "/a-dir-4")
}

//TODO larger trees coming soon

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
