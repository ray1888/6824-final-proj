// Code generated by generate_unit_tests.go. DO NOT EDIT.
// This file contains a unit test for every combination of functionality test
// (found in filesystem_tests.go) and difficulty (found in difficulties.go).
// Generated at Wed May 9 8:12:27 PM.

package fsraft

import "filesystem"
import "testing"

func TestClerk_OneClerkThreeServersNoErrors_TestOpenOpened(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenOpened, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1KBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1KBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64BytesIter64K(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64BytesIter64K, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestHelpGenerateJenkinsPipeline(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestHelpGenerateJenkinsPipeline, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenNotFound, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenAlreadyExists(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenAlreadyExists, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite10MBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesSimple, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteReadVerfiyHoleExpansion(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteReadVerfiyHoleExpansion, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestCloseClosed(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestCloseClosed, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteMaxFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteMaxFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadOffsetOperation(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadOffsetOperation, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestReadClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestReadClosedFile, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRoot(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRoot, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64KBIter10MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter10MB, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestDeleteNotFound(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestDeleteNotFound, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose64(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenROClose64, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekErrorBadFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestSeekOffEOF(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestSeekOffEOF, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead1ByteSimple(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead1ByteSimple, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenRWClose4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseLeastFD(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseLeastFD, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRootMax(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestOpenCloseDeleteRootMax, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic4(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteReadBasic4, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite8Bytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite8Bytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite100MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite100MBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead1MBIter100MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead1MBIter100MB, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestBasicOpenClose(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestBasicOpenClose, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWriteClosedFile(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWriteClosedFile, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1Byte(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1Byte, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestWrite1MBytes(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestWrite1MBytes, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter8(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead8BytesIter8, OneClerkThreeServersNoErrors)
}

func TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64KBIter1MB(t *testing.T) {
	runFunctionalityTestWithDifficulty(t, filesystem.TestRndWriteRead64KBIter1MB, OneClerkThreeServersNoErrors)
}

