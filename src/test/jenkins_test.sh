#!/usr/bin/env bash

export JENKINS_FAIL=0
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
mkdir -p $SCRIPT_DIR/outfiles
if [ -z $JENKINS ]; then
  rm $SCRIPT_DIR/outfiles/*
fi
touch $SCRIPT_DIR/outfiles/log

iter() {
        testname="$1" 
        index="$2"
        outfile="$SCRIPT_DIR/outfiles/${testname}.${index}.out"
        touch "$outfile"
        #echo "running ${testname}.${index}"
        if [ -z $JENKINS ]; then
           go test -run "$testname" > "$outfile" 2>&1
           exit_code=$?
        else
           $GOBIN test -run "$testname" > "$outfile" 2>&1
           exit_code=$?
        fi
        if [ $exit_code == 0 ]; then
                echo "${testname} success ${index}" | tee -a $SCRIPT_DIR/outfiles/log
                rm "$outfile"
        else
                echo "${testname} FAIL ${index}!" | tee -a $SCRIPT_DIR/outfiles/log
                export JENKINS_FAIL=1
        fi
}

run_test() {
        testname="$1"
        quantity="$2"
        let n=0 # 
        #echo "running $testname $quantity times"
        while [ $n -lt ${quantity} ]; do
                let n++
                time iter $testname $n
        done
}

cd $SCRIPT_DIR/../memoryFS
echo Begin Core MemoryFS Tests
 run_test "TestMemoryFS_TestWrite1KBytes" 1
 run_test "TestMemoryFS_TestReadClosedFile" 1
 run_test "TestMemoryFS_TestRndWriteRead1ByteSimple" 1
 run_test "TestMemoryFS_TestBasicOpenClose" 1
 run_test "TestMemoryFS_TestOpenOpened" 1
 run_test "TestMemoryFS_TestSeekOffEOF" 1
 run_test "TestMemoryFS_TestWriteReadBasic4" 1
 run_test "TestMemoryFS_TestCannotWriteToReadOnly" 1
 run_test "TestMemoryFS_TestRndWriteRead8BytesSimple" 1
 run_test "TestMemoryFS_TestOpenROClose64" 1
 run_test "TestMemoryFS_TestOpenOffsetEqualsZero" 1
 run_test "TestMemoryFS_TestSeekErrorBadFD" 1
 run_test "TestMemoryFS_TestCannotReadFromWriteOnly" 1
 run_test "TestMemoryFS_TestWrite10MBytes512Kx20" 1
 run_test "TestMemoryFS_TestWrite1Byte" 1
 run_test "TestMemoryFS_TestWrite10MBytes128Kx80" 1
 run_test "TestMemoryFS_TestMkdirTree" 1
 run_test "TestMemoryFS_TestRndWriteReadVerfiyHoleExpansion" 1
 run_test "TestMemoryFS_TestDeleteCannotDeleteRootDir" 1
 run_test "TestMemoryFS_TestDeleteNotFound" 1
 run_test "TestMemoryFS_TestOpenAlreadyExists" 1
 run_test "TestMemoryFS_TestOpenCloseLeastFD" 1
 run_test "TestMemoryFS_TestWrite10MBytes256Kx40" 1
 run_test "TestMemoryFS_TestRndWriteRead8BytesIter64" 1
 run_test "TestMemoryFS_TestWriteSomeButNotAll" 1
 run_test "TestMemoryFS_TestOpenBlockNoContention" 0
 run_test "TestMemoryFS_TestRndWriteRead64BytesSimple" 1
 run_test "TestMemoryFS_TestMkdirAlreadyExists" 1
 run_test "TestMemoryFS_TestOpenNotFound" 1
 run_test "TestMemoryFS_TestWrite8Bytes" 1
 run_test "TestMemoryFS_TestRndWriteRead6400BytesIter64K" 1
 run_test "TestMemoryFS_TestOpenRWClose" 1
 run_test "TestMemoryFS_TestOpenCloseDeleteMaxFD" 1
 run_test "TestMemoryFS_TestOpenBlockMultipleWaiting" 0
 run_test "TestMemoryFS_TestOpenRWClose64" 1
 run_test "TestMemoryFS_TestOpenTruncate" 1
 run_test "TestMemoryFS_TestSeekErrorBadOffsetOperation" 1
 run_test "TestMemoryFS_TestCloseClosed" 1
 run_test "TestMemoryFS_TestOpenROClose" 1
 run_test "TestMemoryFS_TestWrite10MBytes64Kx160" 1
 run_test "TestMemoryFS_TestWrite10MBytes1Mx10" 1
 run_test "TestMemoryFS_TestOpenCloseDeleteAcrossDirectories" 1
 run_test "TestMemoryFS_TestOpenROClose4" 1
 run_test "TestMemoryFS_TestOpenBlockOnlyOne" 0
 run_test "TestMemoryFS_TestRndWriteRead512KBIter1MB" 1
 run_test "TestMemoryFS_TestRndWriteRead128KBIter10MB" 1
 run_test "TestMemoryFS_TestOpenCloseDeleteRootMax" 1
 run_test "TestMemoryFS_TestWriteClosedFile" 1
 run_test "TestMemoryFS_TestMkdirNotFound" 1
 run_test "TestMemoryFS_TestRndWriteRead8BytesIter8" 1
 run_test "TestMemoryFS_TestMkdir" 1
 run_test "TestMemoryFS_TestOpenRWClose4" 1
 run_test "TestMemoryFS_TestOpenAppend" 1
 run_test "TestMemoryFS_TestOpenBlockOneWaiting" 0
 run_test "TestMemoryFS_TestWriteReadBasic" 1
 run_test "TestMemoryFS_TestWrite1MBytes" 1
 run_test "TestMemoryFS_TestOpenCloseDeleteRoot" 1
 run_test "TestMemoryFS_TestWrite10MBytes10Mx1" 1
cd $SCRIPT_DIR/../fsraft
echo Begin Raft Difficulty 1 Tests - Reliable Network - Clerk_OneClerkThreeServersNoErrors Tests
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseLeastFD" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes256Kx40" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter64" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestDeleteNotFound" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenAlreadyExists" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWriteSomeButNotAll" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestMkdirAlreadyExists" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenBlockNoContention" 0
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead64BytesSimple" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead6400BytesIter64K" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenNotFound" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite8Bytes" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenBlockMultipleWaiting" 0
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteMaxFD" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadOffsetOperation" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose64" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenTruncate" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes64Kx160" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes1Mx10" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteAcrossDirectories" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestCloseClosed" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead512KBIter1MB" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead128KBIter10MB" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose4" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenBlockOnlyOne" 0
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestMkdirNotFound" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRootMax" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWriteClosedFile" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenBlockOneWaiting" 0
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite1MBytes" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesIter8" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestMkdir" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenRWClose4" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenAppend" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenCloseDeleteRoot" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes10Mx1" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestSeekOffEOF" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWriteReadBasic4" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestCannotWriteToReadOnly" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite1KBytes" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestReadClosedFile" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead1ByteSimple" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestBasicOpenClose" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenOpened" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteRead8BytesSimple" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestSeekErrorBadFD" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestCannotReadFromWriteOnly" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes512Kx20" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenROClose64" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestOpenOffsetEqualsZero" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestMkdirTree" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestRndWriteReadVerfiyHoleExpansion" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestDeleteCannotDeleteRootDir" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite1Byte" 1
 run_test "TestClerk_OneClerkThreeServersNoErrors_TestWrite10MBytes128Kx80" 1
echo Begin Raft Difficulty 2 Tests - Lossy Network - Clerk_OneClerkFiveServersUnreliableNet Tests
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose64" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenTruncate" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestSeekErrorBadOffsetOperation" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteAcrossDirectories" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestCloseClosed" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes64Kx160" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes1Mx10" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose4" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenBlockOnlyOne" 0
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead512KBIter1MB" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead128KBIter10MB" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteRootMax" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWriteClosedFile" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirNotFound" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1MBytes" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesIter8" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestMkdir" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose4" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenAppend" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenBlockOneWaiting" 0
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWriteReadBasic" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteRoot" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes10Mx1" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestCannotWriteToReadOnly" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1KBytes" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestReadClosedFile" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead1ByteSimple" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestBasicOpenClose" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenOpened" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestSeekOffEOF" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWriteReadBasic4" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesSimple" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes512Kx20" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenROClose64" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenOffsetEqualsZero" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestSeekErrorBadFD" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestCannotReadFromWriteOnly" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestDeleteCannotDeleteRootDir" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite1Byte" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes128Kx80" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirTree" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteReadVerfiyHoleExpansion" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead8BytesIter64" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestDeleteNotFound" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenAlreadyExists" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseLeastFD" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite10MBytes256Kx40" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWriteSomeButNotAll" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenBlockNoContention" 0
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead64BytesSimple" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestMkdirAlreadyExists" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenNotFound" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestWrite8Bytes" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestRndWriteRead6400BytesIter64K" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenRWClose" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenCloseDeleteMaxFD" 1
 run_test "TestClerk_OneClerkFiveServersUnreliableNet_TestOpenBlockMultipleWaiting" 0
 echo "Begin Singleton Tests"
 run_test "TestOneClerkFiveServersPartition" 1
 run_test "TestKVBasic" 1
if [ ! -z $JENKINS ]; then
  exit $JENKINS_FAIL
fi
