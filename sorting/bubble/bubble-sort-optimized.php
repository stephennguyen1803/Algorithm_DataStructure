<?php
$inputArray = [1,8,3,5,6];
function printResult(array $inputArray) :void {
        $n = count($inputArray);
        for($i=0; $i< $n; $i++) {
                echo $inputArray[$i]." ";
        }
        echo(PHP_EOL);
}

function bubbleSortOptimize(array $input) :array {
	$n = count($input);
	$flagCheck = true;
	for($i = 0; $i< $n; $i++) {
		if ($flagCheck === false) {
			return $input;
		}
		for($j = 0; $j < $n - $i - 1; $j++) {
			$flagCheck = false;
			if($input[$j] > $input[$j+1]) {
				$flagCheck = true;
				$t = $input[$j];
				$input[$j] = $input[$j+1];
				$input[$j+1] = $t;
			}
		}
	}
}

print_r('Truoc khi sap xep'.PHP_EOL);
printResult($inputArray);
print_r('Sau khi sap xep'.PHP_EOL);
$result = bubbleSortOptimize($inputArray);
printResult($result);




