use std::io;

pub fn run() -> Game {
    println!("Welcome to CLI poker! Please specify your name:");
    let player_name: String = set_player_name("".to_string());
    println!("---\nWelcome, {}! How many people should be playing with you? (1 to 9)", player_name);
    let player_amount: u8 = set_player_amount(0);
    println!("---\n{} players are ready to start the game. What size should the big blinds be?", player_amount);
    println!("Hint: This also determines the overall pot size (big blind * amount of players * 100).");
    let big_blind: u32 = set_big_blind(0);
    let small_blind: u32 = big_blind / 2;
    println!("---\nBig blinds are {}$, small blinds are {}$.", big_blind, small_blind);
    let real_player_amount = player_amount + 1;
    let pot_size: i32 = (big_blind as i32) * (real_player_amount as i32) * 100;
    println!("And the pot size is: {}$\n---", pot_size);
    let new_game = Game {
        players: real_player_amount,
        big_blind_size: big_blind,
        small_blind_size: small_blind,
        full_pot_size: pot_size,
        human_player: player_name
    };
    return new_game;
}

fn set_player_name(input: String) -> String {
    let mut local_player_name = input;

    io::stdin()
        .read_line(&mut local_player_name)
        .expect("Failed to read line");

    let local_player_name: String = match local_player_name.trim().parse() {
        Ok(string) => string,
        Err(error) => panic!("Problem with player name: {:?}", error),
    };

    if local_player_name.len() < 25 && local_player_name.len() > 0 {
        return local_player_name.to_string();
    } else {
        println!("Please enter a player name between 1 and 25 characters long.");
        set_player_name(local_player_name.to_string())
    }
}

fn set_player_amount(count: u8) -> u8 {
    let mut local_player_count = count.to_string();

    io::stdin()
        .read_line(&mut local_player_count)
        .expect("failed to read from stdin");

    let sanitized_player_count: u8 = match local_player_count.trim().parse() {
        Ok(i) => i,
        Err(error) => panic!("this was not an integer: {:?}", error),
    };

    if sanitized_player_count < 10 && sanitized_player_count > 0 {
        return sanitized_player_count;
    } else {
        println!("Please enter a number between 1 and 9.");
        set_player_amount(0)
    }
}

fn set_big_blind(size: u32) -> u32 {
    let mut big_blind_size = size.to_string();

    io::stdin()
        .read_line(&mut big_blind_size)
        .expect("failed to read from stdin");

    let sanitized_big_blind: u32 = match big_blind_size.trim().parse() {
        Ok(i) => i,
        Err(error) => panic!("this was not an integer: {:?}", error),
    };

    if sanitized_big_blind < 2000 && sanitized_big_blind > 1 {
        return sanitized_big_blind;
    } else {
        println!("Please enter a number between 1 and 2000.");
        set_big_blind(0)
    }
}

pub struct Game{
    pub players: u8,
    pub big_blind_size: u32,
    pub small_blind_size: u32,
    pub full_pot_size: i32,
    pub human_player: String
}
