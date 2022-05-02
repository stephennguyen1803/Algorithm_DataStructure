<?php
//  64 25 12 22 11
$inputArray = [64,25,12,22,11];

function selectionSort(array $input) :array 
{
	$n = count($input)-1;
	for($i = 0; $i <= $n - 1; $i++) {
		for($j = $i; $j < $n; $j++) {
			if($input[$i] >= $input[$j+1]) {
				$t = $input[$i];
				$input[$i] = $input[$j+1];
				$input[$j+1] = $t;
			}
		}
	}

	return $input;
}

function printResult(array $inputArray) :void {
	$n = count($inputArray);
	for($i=0; $i< $n; $i++) {
		echo $inputArray[$i]." ";
	}
	echo(PHP_EOL);
}

print_r('Truoc khi sap xep'.PHP_EOL);
printResult($inputArray);
print_r('Sau khi sap xep'.PHP_EOL);
$result = selectionSort($inputArray);
printResult($result);


