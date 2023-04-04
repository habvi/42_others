use std::{env, num::ParseIntError};

// ---------------------------------------------------------------
const USAGE: &str = "Usage: python3 main.py <unsigned intger>";
const SIDE_EDGE: usize = 2;

fn parse_args(str: &String) -> Result<isize, ParseIntError> {
    let num: isize = match str.trim().parse() {
        Ok(num) => num,
        Err(e) => return Err(e),
    };
    Ok(num)
}

fn is_valid_args(argv: &Vec<String>) -> bool {
    let size = argv.len();
    if size != 2 {
        eprintln!("{}", USAGE);
        return false;
    }
    let number: isize = match parse_args(&argv[1]) {
        Ok(ret) => ret,
        Err(_) => {
            eprintln!("{}", USAGE);
            return false;
        },
    };
    if number <= 0 || number > 2147483647 {
        eprintln!("{}", USAGE);
        return false;
    }
    true
}

// ---------------------------------------------------------------
// first term, common difference, number of terms
fn sigma_adn(a: isize, d: isize, n: isize) -> isize {
    return (2 * a + (n - 1) * d) * n / 2;
}

fn calc_base_steps(total_steps: isize) -> usize {
    let total_heights = sigma_adn(3, 1, total_steps);
    let base_steps = 2 * (total_heights - 1);
    base_steps as usize
}

fn calc_additional_steps(total_steps: usize) -> usize {
    let mut additional_steps = 1;
    if total_steps >= 2 {
        additional_steps += sigma_adn(4, 2, (total_steps as isize - 1) / 2) * 2;
        if total_steps % 2 == 0 {
            additional_steps += 4 + 2 * ((total_steps as isize - 1) / 2);
        }
    }
    additional_steps as usize
}

fn calc_bottom_width(total_steps: usize) -> usize {
    let base_steps = calc_base_steps(total_steps.try_into().unwrap());
    let additional_steps = calc_additional_steps(total_steps);
    base_steps + additional_steps
}

// ---------------------------------------------------------------
fn get_door_size(total_steps: usize) -> usize {
    match total_steps % 2 {
        0 => return total_steps - 1,
        1 => return total_steps,
        _ => unreachable!()
    };
}

//    n : 1 2 3 4 5 6 7 8 9 10
// door : 1 1 3 3 5 5 7 7 9  9
fn is_door_row(pyra: &mut Pyramid, row_i: usize) -> bool {
    if pyra.step_i != pyra.total_steps {
        return false;
    }
    let height = pyra.total_steps + 2;
    let door_size = get_door_size(pyra.total_steps);
    if height - row_i <= door_size {
        return true;
    }
    false
}

fn is_door_col(pyra: &mut Pyramid, col_i: &mut usize) -> bool {
    let center = pyra.total_width / 2;
    let door_size = get_door_size(pyra.total_steps);
    if (center - door_size / 2 <= *col_i) && (*col_i <= center + door_size / 2) {
        return true;
    }
    false
}

fn is_door_knob(pyra: &mut Pyramid, row_i: usize, col_i: usize) -> bool {
    if pyra.total_steps <= 4 {
        return false;
    }
    // row check
    let height = pyra.total_steps + 2;
    let door_size = get_door_size(pyra.total_steps);
    if height - row_i - 1 != door_size / 2 {
        return false;
    }
    // col check
    let center = pyra.total_width / 2;
    if col_i != center + door_size / 2 - 1 {
        return false;
    }
    true
}

fn print_door(pyra: &mut Pyramid, row_i: usize, col_i: usize) -> () {
    if is_door_knob(pyra, row_i, col_i) {
        print!("$");
    } else {
        print!("|");
    }
}

// ---------------------------------------------------------------
fn print_space(pyra: &mut Pyramid, row_width: usize, col_i: &mut usize) -> () {
    let space = (pyra.total_width - row_width - SIDE_EDGE) / 2;
    for _ in 0..space {
        print!(" ");
        *col_i += 1;
    }
}

fn print_slash(col_i: &mut usize) -> () {
    print!("/");
    *col_i += 1;
}

fn print_asterisk(pyra: &mut Pyramid, row_i: usize, row_width: usize, col_i: &mut usize) -> () {
    for _ in 0..row_width {
        if is_door_row(pyra, row_i) && is_door_col(pyra, col_i) {
            print_door(pyra, row_i, *col_i);
        } else {
            print!("*");
        }
        *col_i += 1;
    }
}

fn print_each_steps(pyra: &mut Pyramid) -> () {
    let mut row_width = pyra.top_width;
    let mut col_i = 0;
    for row_i in 0..pyra.step_i + 2 {
        print_space(pyra, row_width, &mut col_i);
        print_slash(&mut col_i);
        print_asterisk(pyra, row_i, row_width, &mut col_i);
        println!("\\");
        col_i = 0;
        row_width += 2;
    }
}

fn print_pyramid(pyra: &mut Pyramid) -> () {
    for step_i in 1..=pyra.total_steps {
        let bottom_width = calc_bottom_width(step_i);
        let top_width = bottom_width - 2 * (step_i + 1);
        pyra.step_i = step_i;
        pyra.bottom_width = bottom_width;
        pyra.top_width = top_width;
        print_each_steps(pyra);
    }
}

// ---------------------------------------------------------------
struct Pyramid {
    step_i: usize,
    total_steps: usize,
    top_width: usize,
    bottom_width: usize,
    total_width: usize,
}

fn main() {
    let argv: Vec<String> = env::args().collect();
    if !is_valid_args(&argv) {
        return;
    }
    let total_steps: usize = parse_args(&argv[1]).ok().unwrap() as usize;
    let bottom_width = calc_bottom_width(total_steps);
    let total_width = bottom_width + SIDE_EDGE;
    let mut pyra = Pyramid {
        step_i: 0,
        total_steps,
        top_width: 0,
        bottom_width: 0,
        total_width,
    };
    // eprintln!("total_width: {}", total_width);
    // eprintln!("=========================");
    print_pyramid(&mut pyra);
}
