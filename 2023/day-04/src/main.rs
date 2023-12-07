use std::fs::read_to_string;

fn main() {
    let lines = read_lines("input.txt");
    let games = get_games(&lines);

    println!("{}", part_1(&games));

    println!("{}", part_2(&games));
}

fn part_1(games: &Vec<(Vec<u32>, Vec<u32>)>) -> u32 {
    let mut result: u32 = 0; 

    for g in games {
        let mut g_points = 0;

        for my_number in &g.0 {
            for winning_number in &g.1 {
                if *my_number == *winning_number {
                    g_points = if g_points == 0 { 1 } else { g_points * 2 };
                    break;
                }
            }
        }

        result += g_points;
    }

    result
}

fn part_2(games: &Vec<(Vec<u32>, Vec<u32>)>) -> u32 {
    let mut cards = vec![1;games.len()];

    for (i, g) in games.iter().enumerate() {
        let mut next_card = i + 1;
        for my_number in &g.0 {
            for winning_number in &g.1 {
                if *my_number == *winning_number {
                    cards[next_card] += cards[i];
                    next_card += 1;
                    break;
                }
            }
        }
    }

    cards.iter().sum()
}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}

fn get_games(lines: &Vec<String>) -> Vec<(Vec<u32>, Vec<u32>)> {
    let mut games: Vec<(Vec<u32>,Vec<u32>)> = Vec::new();

    for line in lines {
        let line: Vec<&str> = line.split(":").collect();
        let line: Vec<&str> = line[1].split("|").collect();

        games.push((get_numbers(line[0]), get_numbers(line[1])));
    }

    games
}

fn get_numbers(string: &str) -> Vec<u32> {
    string.to_string()
        .split(" ")
        .filter(|s| s.len() > 0)
        .map(|s| s.parse().unwrap())
        .collect()
}

