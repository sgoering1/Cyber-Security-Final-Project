<?php
    $files = scandir('Cyber-Security-Final-Project/img');
    foreach($files as $file) {
        if($file !== "." && $file !== "..") {
            echo "<img src='$file' />";
        }
    }
?>