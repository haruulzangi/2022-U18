<?php
  require("flag.php");
  if (isset($_GET['source'])) {
    highlight_file(__FILE__);
    die();
  }
  if (isset($_GET['id'])) {
    $your_entered_string = $_GET['id'];
    $intermediate_string = 'HaRuUl18ZaNGi18U18';
    $final_string = preg_replace(
            "/$intermediate_string/", '', $your_entered_string);
    if ($final_string === $intermediate_string) {
      HaruulZangiU18();
    }
  }
?>
<!DOCTYPE HTML>
<html>
  <head>
    <title>pregpregreplacereplace</title>
  </head>
  <body>
    <p>Try to reach <code>HaruulZangiU18()</code></p>
    <a target="_blank" href="?source">See the source code</a>
  </body>
</html>