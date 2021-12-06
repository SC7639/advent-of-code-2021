<?php

function partOne(array $input): int
{
    $prevVal = null;
    $increased = 0;

    foreach ($input as $val) {
        if ($val > $prevVal && !is_null($prevVal)) {
            $increased++;
        }
        $prevVal = $val;
    }

    return $increased;
}

function partTwo(array $input): int
{
    $prevWindow = null;
    $increased = 0;


    for ($i=0; $i < count($input); $i++) {
        if ($i + 3 > count($input)) {
            break;
        }

        $window = array_sum(
            array_slice($input, $i, 3)
        );
        if ($window > $prevWindow && !is_null($prevWindow)) {
            $increased++;
        }

        $prevWindow = $window;
    }

    return $increased;
}

$input = file_get_contents(__DIR__ . "/sonar.txt");
$inputArr = explode("\n", $input);

$one = partOne($inputArr);
$two = partTwo($inputArr);

echo "There are $one measurements that are larger than previous measurement\n";
echo "There are $two window measurements that are lareger than previous windows\n";
