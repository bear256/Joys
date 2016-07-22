<?php
$file = fopen("UCloud.txt", "r") or die("Unable to open file UCloud.txt!");
$total = 0;
$i = 0;
while(!feof($file)) {
    $text = fgets($file);
    $count = substr_count($text, "UCanUup");
    echo "Line ".$i." : ".$count."\n";
    $total += $count;
    $i++;
}
echo "Total: ".$total."\n";
fclose($file);
?>