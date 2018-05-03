// Code generated by generate_unit_tests.go. DO NOT EDIT.
// This file contains a unit test for every functionality test (found in filesystem_tests.go).
// Generated at Wed May 2 11:53:16 PM.

package memoryFS

import "testing"
import "fsraft"

func TestMemoryFS_TestOpenRWClose64(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenRWClose64(t, &mfs)
}

func TestMemoryFS_TestOpenRWClose4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenRWClose4(t, &mfs)
}

func TestMemoryFS_TestMkdirTree(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestMkdirTree(t, &mfs)
}

func TestMemoryFS_TestOpenCloseBasic(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenCloseBasic(t, &mfs)
}

func TestMemoryFS_TestOpenROClose4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenROClose4(t, &mfs)
}

func TestMemoryFS_TestReadWriteBasic4(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestReadWriteBasic4(t, &mfs)
}

func TestMemoryFS_TestMkdir(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestMkdir(t, &mfs)
}

func TestMemoryFS_TestSeekErrorBadFD(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestSeekErrorBadFD(t, &mfs)
}

func TestMemoryFS_TestSeekErrorBadOffsetOperation(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestSeekErrorBadOffsetOperation(t, &mfs)
}

func TestMemoryFS_TestOpenROClose(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenROClose(t, &mfs)
}

func TestMemoryFS_TestOpenRWClose(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenRWClose(t, &mfs)
}

func TestMemoryFS_TestOpenCloseLeastFD(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenCloseLeastFD(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteFD128(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenCloseDeleteFD128(t, &mfs)
}

func TestMemoryFS_TestOpenCloseDeleteAcrossDirectories(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestOpenCloseDeleteAcrossDirectories(t, &mfs)
}

func TestMemoryFS_TestReadWriteBasic(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestReadWriteBasic(t, &mfs)
}

func TestMemoryFS_TestSeekErrorBadOffset1(t *testing.T) {
	mfs := CreateEmptyMemoryFS()
        fsraft.TestSeekErrorBadOffset1(t, &mfs)
}

