<?php

$data = fopen("data.txt", "r");
$containingPairs = 0;
if ($data) {
    while (($line = fgets($data)) !== false) {
        $elves = explode(',', str_replace("\r\n", "", $line));
        $elf1sections = explode('-', $elves[0]);
        $elf2sections = explode('-', $elves[1]);

        if($elf1sections[0] <= $elf2sections[0] && $elf1sections[1] >= $elf2sections[1]){
            $containingPairs++;
        }elseif($elf2sections[0] <= $elf1sections[0] && $elf2sections[1] >= $elf1sections[1]){
            $containingPairs++;
        }
    }

    print 'Total Count: ' . strval($containingPairs);

    fclose($data);
}