import sys

SIDE_EDGE = 2
BINARY_INT_MAX = 2147483647

# ------------------------------------------------------------
def is_valid_args():
    size = len(sys.argv)
    if size != 2:
        print("Usage: python3 main.py <integer>", file=sys.stderr)
        return False
    if not sys.argv[1].isdigit():
        print("Usage: python3 main.py <integer>", file=sys.stderr)
        return False
    if int(sys.argv[1]) > BINARY_INT_MAX:
        return False
    return True

# ------------------------------------------------------------
# first term, common difference, number of terms
def sigma_adn(a, d, n):
    return (2*a + (n - 1)*d) * n // 2

def calc_bottom_width(n):
    def calc_base_steps():
        total_heights = sigma_adn(3, 1, n)
        base_steps = 2 * (total_heights - 1)
        return base_steps

    def calc_additional_steps():
        additional_steps = 1
        if n >= 2:
            additional_steps += sigma_adn(4, 2, (n - 1) // 2) * 2
            if n % 2 == 0:
                additional_steps += 4 + 2 * ((n - 1) // 2)
        return additional_steps

    base_steps = calc_base_steps()
    additional_steps = calc_additional_steps()
    return base_steps + additional_steps

# ------------------------------------------------------------
#    n : 1 2 3 4 5 6 7 8 9 10
# door : 1 1 3 3 5 5 7 7 9  9
def is_door_row(n, step_i, row_i):
    if step_i != n:
        return False
    height = n + 2
    door_size = (n if n % 2 else n - 1)
    if height - row_i <= door_size:
        return True
    return False

def is_door_col(n, col_i, total_width):
    center = total_width // 2
    door_size = (n if n % 2 else n - 1)
    if (center - door_size // 2 <= col_i) and (col_i <= center + door_size // 2):
        return True
    return False

def is_door_knob(n, row_i, col_i, total_width):
    if n <= 4:
        return False
    # row check
    height = n + 2
    door_size = (n if n % 2 else n - 1)
    if height - row_i - 1 != door_size // 2:
        return False
    # col check
    center = total_width // 2
    if col_i != center + door_size // 2 - 1:
        return False
    return True

def print_door(n, row_i, col_i, total_width):
    if is_door_knob(n, row_i, col_i, total_width):
        print("$", end="")
    else:
        print("|", end="")

# ------------------------------------------------------------
def print_space(col_i, row_width, total_width):
    space = (total_width - row_width - SIDE_EDGE) // 2
    for _ in range(space):
        print(" ", end="")
        col_i += 1
    return col_i

def print_slash(col_i):
    print("/", end="")
    col_i += 1
    return col_i

def print_asterisk(n, step_i, row_i, col_i, row_width, total_width):
    for _ in range(row_width):
        if is_door_row(n, step_i, row_i) and is_door_col(n, col_i, total_width):
            print_door(n, row_i, col_i, total_width)
        else:
            print("*", end="")
        col_i += 1
    return col_i

def print_each_steps(n, step_i, top_width, total_width):
    row_width = top_width
    col_i = 0
    for row_i in range(step_i + 2):
        col_i = print_space(col_i, row_width, total_width)
        col_i = print_slash(col_i)
        col_i = print_asterisk(n, step_i, row_i, col_i, row_width, total_width)
        print("\\")
        col_i = 0
        row_width += 2

def print_pyramid(n, total_width):
    for step_i in range(1, n + 1):
        bottom_width = calc_bottom_width(step_i)
        top_width = bottom_width - 2 * (step_i + 1)
        print_each_steps(n, step_i, top_width, total_width)

# ------------------------------------------------------------
def main():
    if not is_valid_args():
        return
    n = int(sys.argv[1])
    bottom_width = calc_bottom_width(n)
    total_width = bottom_width + SIDE_EDGE
    # print("total_width:", total_width, file=sys.stderr)
    # print("=========================", file=sys.stderr)
    print_pyramid(n, total_width)

if __name__ == '__main__':
    main()
