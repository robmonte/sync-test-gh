#!/bin/bash

RESULTS_FILENAME="codespace-results.txt"



echo ""
echo ""
echo "Beginning test case checks:"



if [[ -z "$STOREGH_TEST_TYPE" ]]
then
	echo "missing required \"STOREGH_TEST_TYPE\" field to determine test type"
	echo -n "FAIL" > $RESULTS_FILENAME
	exit 1
fi
echo "running \"$STOREGH_TEST_TYPE\" test type"



if [[ "$STOREGH_TEST_TYPE" == "WRITE_TEST_TOKEN" ]]
then
	var="STOREGH_TEST_TOKEN_WRITE_KEY_1_BAD_SYMBOLS________${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret one!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi

	var="STOREGH_TEST_TOKEN_WRITE_KEY_2${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "I am secret two!" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi

	var="STOREGH_TEST_TOKEN_WRITE_KEY_3${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "1234" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi

	var="STOREGH_TEST_TOKEN_WRITE_KEY_4${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "{\"0\":\"my\",\"1\":\"map\",\"2\":\"value\"}" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi

	var="STOREGH_TEST_TOKEN_WRITE_KEY_5${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "{\"number\":1,\"true\":false}" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi
fi



if [[ "$STOREGH_TEST_TYPE" == "WRITE_TEST_APP" ]]
then
	var="STOREGH_TEST_APP_WRITE_KEY_1${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "12345" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi
fi



if [[ "$STOREGH_TEST_TYPE" == "UPDATE_TEST_TOKEN" ]]
then
	var="STOREGH_TEST_TOKEN_UPDATE_KEY_1${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "value-1-updated" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value!"
	fi

	var="STOREGH_TEST_TOKEN_UPDATE_KEY_2${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "{\"0\":\"value-2-updated\"}" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi
fi



if [[ "$STOREGH_TEST_TYPE" == "UPDATE_TEST_APP" ]]
then
	var="STOREGH_TEST_APP_UPDATE_KEY_1${STOREGH_TEST_UUID}"
	if [[ "${!var}" != "value-1-updated" ]]
	then
		echo "received \"${var}\" value but did not match expected value"
		echo -n "FAIL" > $RESULTS_FILENAME
		exit 1
	else
		echo "the value of \"${var}\" matched the expected value"
	fi
fi



echo "Success!"
echo -n "SUCCESS" > $RESULTS_FILENAME
echo ""
echo ""
sleep 2



git add "codespace-results.txt"
git commit --allow-empty -m "update codespace-results.txt with test status"
git push
