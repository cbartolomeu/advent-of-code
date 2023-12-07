use std::fs::read_to_string;

fn main() {
    let lines = read_lines("input.txt");
    let mut result = Vec::new();

    for line in lines {
        let mut first: char = 'a';
        let mut last: char = 'a';
        let mut string = String::from("");
     
        for c in line.chars() {
            string.push(c);

            let spelled_num = get_spelled_num(string.as_str());

            if is_char_digit(c) {
                if first == 'a' {
                    first = c;
                }
                last = c;

                string = String::from("");
            } else if is_char_digit(spelled_num) {
                if first == 'a' {
                    first = spelled_num;
                }
                last = spelled_num;

                string = String::from(c);
            }
        }

        result.push(chars_to_int(first, last));
    }

    let answer: u32 = result.iter().sum();
    println!("{answer}");
}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}

fn chars_to_int(a: char, b: char) -> u32 {
    let num = format!("{}{}", a, b);
    num.parse().unwrap()
}

fn is_char_digit(a: char) -> bool {
    a >= '0' && a <= '9'
}

fn get_spelled_num(string: &str) -> char {
    let string = String::from(string);

    match string {
        _ if string.contains("zero") => '0',
        _ if string.contains("one") => '1',
        _ if string.contains("two") => '2',
        _ if string.contains("three") => '3',
        _ if string.contains("four") => '4',
        _ if string.contains("five") => '5',
        _ if string.contains("six") => '6',
        _ if string.contains("seven") => '7',
        _ if string.contains("eight") => '8',
        _ if string.contains("nine") => '9',
        _ => 'a',
    }
}
