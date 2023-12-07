use std::fs::read_to_string;

#[derive(Debug)]
struct Set {
    red: u32,
    green: u32,
    blue: u32,
}

#[derive(Debug)]
struct Game {
    id: u32,
    sets: Vec<Set>,
}

impl Game {
    fn from_string(string: String) -> Game {
        let string: Vec<&str> = string.split(":").collect();
        let game_id: Vec<&str> = string[0].split(" ").collect();
        let mut sets: Vec<Set> = Vec::new();

        for set in string[1].split(";").collect::<Vec<&str>>() {
            sets.push(Set::from_string(String::from(set)));
        }

        Game {
            id: game_id[1].parse().unwrap(),
            sets,
        }
    }
}

impl Set {
    fn from_string(string: String) -> Set {
        let mut red = 0;
        let mut green = 0;
        let mut blue = 0;

        for color in string.split(",").collect::<Vec<&str>>() {
            let color = String::from(color);
            let set = color.trim().split(" ").collect::<Vec<&str>>();
            let cubes: u32 = set[0].parse().unwrap();

            match set[1] {
                "red" => {
                    red = cubes;
                },
                "green" => {
                    green = cubes;
                },
                _ => {
                    blue = cubes;
                },
            };
        }

        Set {
            red,
            green,
            blue,
        }
    }
}

fn main() {
    let lines = read_lines("input.txt");
    let games = parse_games(lines);

    println!("{}", part_1(&games));
    println!("{}", part_2(&games));
}

fn part_1(games: &Vec<Game>) -> u32 {
    let mut possible_games = 0;
    let red = 12;
    let green = 13;
    let blue = 14;

    for game in games {
        let mut possible_game = true;
        for set in &game.sets {
            if set.green > green || set.blue > blue || set.red > red {
                possible_game = false;
            }
        }
        possible_games = if possible_game { possible_games + game.id } else { possible_games };
    }

    possible_games
}

fn part_2(games: &Vec<Game>) -> u32 {
    let mut power = 0;
    for game in games {
        let mut red = 0;
        let mut green = 0;
        let mut blue = 0;

        for set in &game.sets {
            if set.green > green {
                green = set.green;
            }

            if set.blue > blue {
                blue = set.blue;
            }

            if set.red > red {
                red = set.red;
            }
        }
        power = power + red * green * blue;
    }

    power
}

fn read_lines(filename: &str) -> Vec<String> {
    let mut result = Vec::new();

    for line in read_to_string(filename).unwrap().lines() {
        result.push(line.to_string())
    }

    result
}

fn parse_games(lines: Vec<String>) -> Vec<Game> {
    let mut games: Vec<Game> = Vec::new();

    for line in lines {
        games.push(Game::from_string(line));
    }

    games
}
