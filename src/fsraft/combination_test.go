// Code generated by generate_unit_tests.go. DO NOT EDIT.
// This file contains a unit test for every combination of functionality test
// (found in filesystem_tests.go) and difficulty (found in difficulties.go).
// Generated at Thu May 10 8:47:56 PM.

package fsraft

import "filesystem"
import "testing"

func TestClerk_OneClerkThreeServersNoErrors_TestSeekOffEOF(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekOffEOF, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestReadClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestReadClosedFile, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestDeleteNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestDeleteNotFound, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead64KBIter10MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter10MB, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenROClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead8BytesIter64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter64, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1Byte(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1Byte, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenNotFound, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead64BytesIter64K(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesIter64K, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirTree(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirTree, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWriteReadBasic(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteMaxFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteMaxFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteMaxFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteMaxFD, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestCannotReadFromWriteOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotReadFromWriteOnly, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWriteSomeButNotAll(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteSomeButNotAll, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersNoErrors_TestBasicOpenClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestBasicOpenClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesSimple, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestReadClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestReadClosedFile, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenROClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose4, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenCloseDeleteRoot(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRoot, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead64KBIter10MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter10MB, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenCloseDeleteAcrossDirectories(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteAcrossDirectories, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestDeleteNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestDeleteNotFound, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestMkdir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdir, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteAcrossDirectories(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteAcrossDirectories, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestDeleteNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestDeleteNotFound, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWrite8Bytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite8Bytes, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestMkdir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdir, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesSimple, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWriteReadBasic4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic4, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead8BytesIter8(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter8, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead64BytesIter64K(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesIter64K, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAlreadyExists, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead1ByteSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead1ByteSimple, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestCannotDeleteRootDir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotDeleteRootDir, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenRWClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose64, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead1ByteSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead1ByteSimple, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestCloseClosed(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCloseClosed, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64KBIter1MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter1MB, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestMkdir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdir, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestCannotWriteToReadOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotWriteToReadOnly, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestCannotDeleteRootDir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotDeleteRootDir, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1MBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestCloseClosed(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCloseClosed, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestBasicOpenClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestBasicOpenClose, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestCloseClosed(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCloseClosed, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1Byte(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1Byte, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead64KBIter1MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter1MB, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenCloseDeleteMaxFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteMaxFD, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenCloseDeleteRootMax(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRootMax, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWriteSomeButNotAll(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteSomeButNotAll, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite10MBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestBasicOpenClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestBasicOpenClose, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseLeastFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseLeastFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRoot(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRoot, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64KBIter10MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter10MB, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteReadVerfiyHoleExpansion(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteReadVerfiyHoleExpansion, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1KBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1KBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestCannotWriteToReadOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotWriteToReadOnly, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1MBytes, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestSeekErrorBadFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadFD, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose4, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenAppend(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAppend, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteRoot(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRoot, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWriteReadBasic(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead64BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesSimple, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead64KBIter1MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter1MB, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestMkdirAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirAlreadyExists, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestSeekErrorBadFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadFD, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWriteClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteClosedFile, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead64BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesSimple, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenCloseLeastFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseLeastFD, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWrite10MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite10MBytes, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestMkdirTree(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirTree, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAlreadyExists, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesIter8(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter8, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenOpened(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOpened, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestMkdirAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirAlreadyExists, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteReadVerfiyHoleExpansion(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteReadVerfiyHoleExpansion, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadOffsetOperation(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadOffsetOperation, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteClosedFile, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestCannotReadFromWriteOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotReadFromWriteOnly, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseLeastFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseLeastFD, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenOffsetEqualsZero(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOffsetEqualsZero, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWrite8Bytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite8Bytes, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenRWClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose4, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenOffsetEqualsZero(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOffsetEqualsZero, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestReadClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestReadClosedFile, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesSimple, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteRootMax(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRootMax, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWriteReadBasic4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic4, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestRndWriteRead8BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesSimple, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestCannotDeleteRootDir(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotDeleteRootDir, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWrite1Byte(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1Byte, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWrite1KBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1KBytes, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenOpened(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOpened, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose64, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteAcrossDirectories(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteAcrossDirectories, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirNotFound, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenNotFound, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestCannotWriteToReadOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotWriteToReadOnly, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestSeekErrorBadOffsetOperation(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadOffsetOperation, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestCannotReadFromWriteOnly(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCannotReadFromWriteOnly, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestSeekOffEOF(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekOffEOF, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenOpened(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOpened, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestMkdirNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirNotFound, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenTruncate(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenTruncate, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestSeekOffEOF(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekOffEOF, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenRWClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenTruncate(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenTruncate, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenAppend(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAppend, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite8Bytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite8Bytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter8(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter8, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1KBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1KBytes, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenROClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose64, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirAlreadyExists, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWrite1MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1MBytes, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAlreadyExists, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenTruncate(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenTruncate, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenAppend(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAppend, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteSomeButNotAll(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteSomeButNotAll, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteReadVerfiyHoleExpansion(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteReadVerfiyHoleExpansion, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersSnapshots_TestOpenNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenNotFound, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestWriteClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteClosedFile, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenOffsetEqualsZero(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOffsetEqualsZero, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRootMax(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRootMax, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64BytesIter64K(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesIter64K, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestMkdirTree(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirTree, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose4, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesIter64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter64, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead1ByteSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead1ByteSimple, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose64, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite10MBytes, OneClerkFiveServersUnreliableNet)
}

func TestClerk_OneClerkThreeServersSnapshots_TestSeekErrorBadOffsetOperation(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadOffsetOperation, OneClerkThreeServersSnapshots)
}

func TestClerk_OneClerkThreeServersSnapshots_TestMkdirNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestMkdirNotFound, OneClerkThreeServersSnapshots)
}

