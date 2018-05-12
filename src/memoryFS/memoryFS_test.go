// Code generated by generate_unit_tests.go. DO NOT EDIT.
// This file contains a unit test for every functionality test (found in filesystem_tests.go).
// Generated at Fri May 11 8:41:46 PM.

package memoryFS

import "filesystem"
import "testing"

func TestMemoryFS_TestOpenAppend(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenAppend(t, &mfs)
}

func TestMemoryFS_TestOpenNotFound(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenNotFound(t, &mfs)
}

func TestMemoryFS_TestWriteReadBasic4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWriteReadBasic4(t, &mfs)
}

func TestMemoryFS_TestOpenBlockOnlyOne(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenBlockOnlyOne(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead512KBIter1MB(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead512KBIter1MB(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead128KBIter10MB(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead128KBIter10MB(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes64Kx160(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes64Kx160(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteMaxFD(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenCloseDeleteMaxFD(t, &mfs)
}

func TestMemoryFS_TestOpenBlockMultipleWaiting(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenBlockMultipleWaiting(t, &mfs)
}

func TestMemoryFS_TestCannotReadFromWriteOnly(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestCannotReadFromWriteOnly(t, &mfs)
}

func TestMemoryFS_TestOpenROClose64(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenROClose64(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteRootMax(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenCloseDeleteRootMax(t, &mfs)
}

func TestMemoryFS_TestSeekErrorBadOffsetOperation(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestSeekErrorBadOffsetOperation(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes256Kx40(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes256Kx40(t, &mfs)
}

func TestMemoryFS_TestReadClosedFile(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestReadClosedFile(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead8BytesIter8(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead8BytesIter8(t, &mfs)
}

func TestMemoryFS_TestCloseClosed(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestCloseClosed(t, &mfs)
}

func TestMemoryFS_TestWrite1KBytes(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite1KBytes(t, &mfs)
}

func TestMemoryFS_TestMkdirAlreadyExists(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestMkdirAlreadyExists(t, &mfs)
}

func TestMemoryFS_TestOpenBlockOneWaiting(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenBlockOneWaiting(t, &mfs)
}

func TestMemoryFS_TestCannotWriteToReadOnly(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestCannotWriteToReadOnly(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes1Mx10(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes1Mx10(t, &mfs)
}

func TestMemoryFS_TestMkdir(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestMkdir(t, &mfs)
}

func TestMemoryFS_TestOpenROClose(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenROClose(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead64BytesSimple(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead64BytesSimple(t, &mfs)
}

func TestMemoryFS_TestOpenAlreadyExists(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenAlreadyExists(t, &mfs)
}

func TestMemoryFS_TestOpenRWClose64(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenRWClose64(t, &mfs)
}

func TestMemoryFS_TestOpenOffsetEqualsZero(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenOffsetEqualsZero(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteRoot(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenCloseDeleteRoot(t, &mfs)
}

func TestMemoryFS_TestOpenBlockNoContention(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenBlockNoContention(t, &mfs)
}

func TestMemoryFS_TestWriteClosedFile(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWriteClosedFile(t, &mfs)
}

func TestMemoryFS_TestWrite1MBytes(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite1MBytes(t, &mfs)
}

func TestMemoryFS_TestMkdirTree(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestMkdirTree(t, &mfs)
}

func TestMemoryFS_TestOpenOpened(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenOpened(t, &mfs)
}

func TestMemoryFS_TestDeleteNotFound(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestDeleteNotFound(t, &mfs)
}

func TestMemoryFS_TestOpenTruncate(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenTruncate(t, &mfs)
}

func TestMemoryFS_TestWriteSomeButNotAll(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWriteSomeButNotAll(t, &mfs)
}

func TestMemoryFS_TestRndWriteReadVerfiyHoleExpansion(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteReadVerfiyHoleExpansion(t, &mfs)
}

func TestMemoryFS_TestBasicOpenClose(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestBasicOpenClose(t, &mfs)
}

func TestMemoryFS_TestWrite1Byte(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite1Byte(t, &mfs)
}

func TestMemoryFS_TestOpenCloseLeastFD(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenCloseLeastFD(t, &mfs)
}

func TestMemoryFS_TestWrite8Bytes(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite8Bytes(t, &mfs)
}

func TestMemoryFS_TestDeleteCannotDeleteRootDir(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestDeleteCannotDeleteRootDir(t, &mfs)
}

func TestMemoryFS_TestOpenRWClose4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenRWClose4(t, &mfs)
}

func TestMemoryFS_TestSeekOffEOF(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestSeekOffEOF(t, &mfs)
}

func TestMemoryFS_TestWriteReadBasic(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWriteReadBasic(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes512Kx20(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes512Kx20(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead8BytesIter64(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead8BytesIter64(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead6400BytesIter64K(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead6400BytesIter64K(t, &mfs)
}

func TestMemoryFS_TestSeekErrorBadFD(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestSeekErrorBadFD(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes128Kx80(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes128Kx80(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteAcrossDirectories(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenCloseDeleteAcrossDirectories(t, &mfs)
}

func TestMemoryFS_TestOpenRWClose(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenRWClose(t, &mfs)
}

func TestMemoryFS_TestMkdirNotFound(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestMkdirNotFound(t, &mfs)
}

func TestMemoryFS_TestWrite10MBytes10Mx1(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestWrite10MBytes10Mx1(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead1ByteSimple(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead1ByteSimple(t, &mfs)
}

func TestMemoryFS_TestRndWriteRead8BytesSimple(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestRndWriteRead8BytesSimple(t, &mfs)
}

func TestMemoryFS_TestOpenROClose4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        filesystem.TestOpenROClose4(t, &mfs)
}

