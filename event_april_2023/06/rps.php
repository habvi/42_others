<?php

echo "Choose rock, paper, or scissors: ";
$input = fgets(STDIN);

$to_number = array(
    "rock" => 0,
    "paper" => 1,
    "scissors" => 2,
);

$to_choice = array(
    0 => "rock",
    1 => "paper",
    2 => "scissors",
);

if ($input[-1] == "\n") {
    $input = substr($input, 0, -1);
}

$pc = rand(0, 2);
$pc_choice = $to_choice[$pc];
if (array_key_exists($input, $to_number)) {
    $user = $to_number[$input];
    if ($user == $pc) {
        $result = 0;
    } elseif ($user == 0 and $pc == 2 or $user == 1 and $pc == 0 or $user == 2 and $pc == 1) {
        $result = 1;
    } else {
        $result = 2;
    }
    if ($result == 1) {
        echo "Congratulations! You won! The computer chose $pc_choice.";
    } elseif ($result == 2) {
        echo "Sorry, you lost. The computer chose $pc_choice.\n";
    } else {
        echo "Draw. The computer chose $pc_choice.\n";
    }
} else {
    fputs(STDERR, "Error: invalid input");
}

?>
