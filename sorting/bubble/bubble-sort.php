<?php
$inputArray = [5,2,6,8,1,3];
function printResult(array $inputArray) :void {
        $n = count($inputArray);
        for($i=0; $i< $n; $i++) {
                echo $inputArray[$i]." ";
        }
        echo(PHP_EOL);
}

function bubbleSort(array $input) :array {
	$n = count($input);
	for($i =0; $i < $n; $i++) {
		for($j = 0; $j < $n-$i-1; $j++) {
			if($input[$j] > $input[$j+1]) {
				$t = $input[$j];
				$input[$j] = $input[$j+1];
				$input[$j+1] = $t;
			}
		}
	}

	return $input;
}

print_r('Truoc khi sap xep'.PHP_EOL);
printResult($inputArray);
print_r('Sau khi sap xep'.PHP_EOL);
$result = bubbleSort($inputArray);
printResult($result);

