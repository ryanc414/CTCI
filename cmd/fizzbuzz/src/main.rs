use std::sync::mpsc;
use std::sync::mpsc::{Receiver, Sender};
use std::thread;

const N: u32 = 100;

fn main() {
    let (tx1, rx1) = mpsc::channel();
    let (tx2, rx2) = mpsc::channel();
    let (tx3, rx3) = mpsc::channel();
    let mut handles = Vec::with_capacity(3);

    handles.push(thread::spawn(move || {
        number_counter(rx3, tx1);
    }));

    handles.push(thread::spawn(move || {
        fizz_handler(rx1, tx2);
    }));

    handles.push(thread::spawn(move || {
        buzz_handler(rx2, tx3);
    }));

    for h in handles {
        h.join().unwrap();
    }
}

fn number_counter(rx: Receiver<u32>, tx: Sender<u32>) {
    print!("1 ");
    tx.send(1).unwrap();

    for i in 2_u32..N + 1 {
        rx.recv().unwrap();
        println!();

        if i % 3 != 0 && i % 5 != 0 {
            print!("{} ", i);
        }
        tx.send(i).unwrap();
    }

    rx.recv().unwrap();
    println!();
    tx.send(0).unwrap();
}

fn fizz_handler(rx: Receiver<u32>, tx: Sender<u32>) {
    loop {
        let i = rx.recv().unwrap();
        if i == 0 {
            tx.send(i).unwrap();
            return;
        }

        if i % 3 == 0 {
            print!("Fizz");
        }

        tx.send(i).unwrap();
    }
}

fn buzz_handler(rx: Receiver<u32>, tx: Sender<u32>) {
    loop {
        let i = rx.recv().unwrap();
        if i == 0 {
            return;
        }

        if i % 5 == 0 {
            print!("Buzz");
        }

        tx.send(i).unwrap();
    }
}
