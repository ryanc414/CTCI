use rand::Rng;
use std::sync::{Arc, Mutex};
use std::thread;
use std::time::Duration;

const NUM_PHILSOPHERS: usize = 6;

fn main() {
    let chopsticks: [Mutex<()>; NUM_PHILSOPHERS] = Default::default();
    let chopsticks = Arc::new(chopsticks);

    for i in 0..NUM_PHILSOPHERS {
        let chopsticks = Arc::clone(&chopsticks);
        thread::spawn(move || {
            let mut rng = rand::thread_rng();
            loop {
                let (x, y) = (i, (i + 1) % NUM_PHILSOPHERS);
                let (x, y) = if x < y { (x, y) } else { (y, x) };

                {
                    let _left = chopsticks[x].lock().unwrap();
                    let _right = chopsticks[y].lock().unwrap();

                    println!("philosopher {} is eating", i);
                    thread::sleep(Duration::from_millis(rng.gen_range(10..100)));
                }

                thread::sleep(Duration::from_millis(rng.gen_range(10..100)));
            }
        });
    }

    thread::park();
}
