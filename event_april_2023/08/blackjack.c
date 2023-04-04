#include <stdlib.h>
#include <stdio.h>
#include <stdbool.h>

#define BLACKJACK		21
#define PICTURE_CARD	10
#define ACE_BASIC		11
#define ACE_SPECIAL		1

static bool	is_valid_args(int argc, char **argv)
{
	char	*str;

	if (argc != 2)
	{
		printf("Usage: ./blackjack [23456789TJDKA]\n");
		return (false);
	}
	str = argv[1];
	if (!str[0])
		return (false);
	return (true);
}

static unsigned long long	calc_ace_points(unsigned long long total)
{
	if (total + ACE_BASIC >= BLACKJACK)
		total += ACE_SPECIAL;
	else
		total += ACE_BASIC;
	return (total);
}

static bool	is_ace(char c)
{
	return (c == 'A');
}

static bool	is_alpha(char c)
{
	return (c == 'T' || c == 'J' || c == 'D' || c == 'K');
}

static bool	is_digit(char c)
{
	return ('2' <= c && c <= '9');
}

static void	print_result(unsigned long long total)
{
	if (total == BLACKJACK)
		printf("Blackjack!\n");
	else
		printf("%llu\n", total);
}

int	main(int argc, char **argv)
{
	char				*str;
	char				c;
	unsigned long long	total;
	size_t				i;

	if (!is_valid_args(argc, argv))
		return (EXIT_FAILURE);
	str = argv[1];
	total = 0;
	i = 0;
	while (str[i])
	{
		c = str[i];
		if (is_ace(c))
			total = calc_ace_points(total);
		else if (is_alpha(c))
			total += PICTURE_CARD;
		else if (is_digit(c))
			total += c - '0';
		else
			return (EXIT_FAILURE);
		i++;
	}
	print_result(total);
	return (EXIT_SUCCESS);
}
