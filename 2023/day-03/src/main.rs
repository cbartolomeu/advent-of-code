use std::fs::read_to_string;

fn main() {
    let lines = read_lines("input.txt");
    
    println!("{}", part_1(&lines));
    
    println!("{}", part_2(&lines));
}

fn part_1(lines: &Vec<String>) -> u32 {
    let mut result: u32 = 0; 

    for (i, line) in lines.iter().enumerate() {
        let mut num = String::from("");
        let mut has_adjacent = false;

        for (j, c) in line.chars().enumerate() {
            if is_char_digit(c) {
                num.push(c);

                if has_symbol_adjacent(&lines, i, j) {
                    has_adjacent = true;
                }

            } else {
                if has_adjacent {
                    let parsed_num: u32 = num.parse().unwrap();
                    result = result + parsed_num; 
                    has_adjacent = false;
                }
                num = String::from("");
            }
        }

        if has_adjacent {
            let parsed_num: u32 = num.parse().unwrap();
            result = result + parsed_num;
        }
    }

    result
}

fn part_2(lines: &Vec<String>) -> u32 {
    let mut result: u32 = 0; 

    for (i, line) in lines.iter().enumerate() {
        for (j, c) in line.chars().enumerate() {
            if c == '*' {
                let a = get_gear_ratio(&lines, i, j);
                let ratio = a;

                result = result + ratio;
            }
        }
    }

    result
}


fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}

fn is_char_digit(a: char) -> bool {
    a >= '0' && a <= '9'
}

fn is_adjacent(lines: &Vec<String>, i: usize, j: usize) -> bool {
    if i >= lines.len() {
        return false;
    }

    if j >= lines[0].len() {
        return false;
    }

    lines[i].chars().nth(j).unwrap() != '.'
}

fn is_adjacent_around(lines: &Vec<String>, i: usize, j: usize) -> bool {
    if i >= lines.len() {
        return false;
    }

    if j >= lines[0].len() {
        return false;
    }

    let c = lines[i].chars().nth(j).unwrap();
    !is_char_digit(c) && c != '.'
}

fn has_symbol_adjacent(lines: &Vec<String>, i: usize, j: usize) -> bool {
    let i_m1 = i.wrapping_sub(1);
    let j_m1 = j.wrapping_sub(1);

    // Up
    if is_adjacent(lines, i_m1, j_m1) || is_adjacent(lines, i_m1, j) || is_adjacent(lines, i_m1, j+1) {
        return true;
    }

    // Down
    if is_adjacent(lines, i+1, j_m1) || is_adjacent(lines, i+1, j) || is_adjacent(lines, i+1, j+1) {
        return true;
    }

    // Around
    if is_adjacent_around(lines, i, j_m1) || is_adjacent_around(lines, i, j+1) {
        return true;
    }

    false
}

fn get_adjacent(lines: &Vec<String>, i: usize, j: usize) -> char {
    if i >= lines.len() {
        return '.';
    }

    if j >= lines[0].len() {
        return '.';
    }

    lines[i].chars().nth(j).unwrap()
}

fn get_surround_number(lines: &Vec<String>, i: usize, j: usize) -> u32 {
    let mut num = String::from("");
    let mut start = j;

    loop {
        if start == 0 {
            break;
        }

        let start_m1 = start.wrapping_sub(1);
        let c = lines[i].chars().nth(start_m1).unwrap();
        if is_char_digit(c) {
            start = start_m1;
        } else {
            break;
        }
    };

    for (k, c) in lines[i].chars().enumerate() {
        if k >= start {
            if is_char_digit(c) {
                num.push(c);
            } else {
                break;
            }
        }
    }

    num.parse().unwrap()
}

fn get_gear_ratio(lines: &Vec<String>, i: usize, j: usize) -> u32 {
    let i_m1 = i.wrapping_sub(1);
    let j_m1 = j.wrapping_sub(1);
    let mut numbers = 0;
    let mut ratio = 1;
    let around_positions = [(i, j_m1), (i, j+1)];

    // Up Center
    if is_char_digit(get_adjacent(lines, i_m1, j)) {
        numbers = numbers + 1;
        ratio = ratio * get_surround_number(lines, i_m1, j);
    } else {
        // Check Up Left and Up Right
        if is_char_digit(get_adjacent(lines, i_m1, j_m1)) {
            numbers = numbers + 1;
            ratio = ratio * get_surround_number(lines, i_m1, j_m1);
        }

        if is_char_digit(get_adjacent(lines, i_m1, j+1)) {
            numbers = numbers + 1;
            ratio = ratio * get_surround_number(lines, i_m1, j+1);
        }
    }

    // Down Center
    if is_char_digit(get_adjacent(lines, i+1, j)) {
        numbers = numbers + 1;
        ratio = ratio * get_surround_number(lines, i+1, j);
    } else {
        // Check Down Left and Down Right
        if is_char_digit(get_adjacent(lines, i+1, j_m1)) {
            numbers = numbers + 1;
            ratio = ratio * get_surround_number(lines, i+1, j_m1);
        }

        if is_char_digit(get_adjacent(lines, i+1, j+1)) {
            numbers = numbers + 1;
            ratio = ratio * get_surround_number(lines, i+1, j+1);
        }
    }


    for pos in around_positions {
        if is_char_digit(get_adjacent(lines, pos.0, pos.1)) {
            numbers = numbers + 1;
            ratio = ratio * get_surround_number(lines, pos.0, pos.1);
        }
    }

    if numbers == 2 { ratio } else { 0 }
}
