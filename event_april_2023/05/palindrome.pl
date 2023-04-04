use strict;
use warnings;

print "Enter a string: ";
my $str = <STDIN>;

my $str_len = length($str);

for (my $i = 0; $i < $str_len / 2; $i++) {
    my $left = substr($str, $i, 1);
    my $right = substr($str, $str_len - $i - 2, 1);
    if ($left ne $right) {
        print "The string is not a palindrome.\n";
        exit(1);
    }
}
print "The string is a palindrome!\n";