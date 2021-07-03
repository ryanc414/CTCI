use std::sync::mpsc;
use std::sync::Arc;
use std::thread;

fn main() {
    println!("Hello, world!");

    let foo = Arc::new(Foo {});
    let mut handles = Vec::with_capacity(3);
    let (tx1, rx1) = mpsc::channel();
    let (tx2, rx2) = mpsc::channel();

    {
        let foo = Arc::clone(&foo);

        handles.push(thread::spawn(move || {
            foo.first();
            tx1.send(()).unwrap();
        }));
    }

    {
        let foo = Arc::clone(&foo);

        handles.push(thread::spawn(move || {
            rx1.recv().unwrap();
            foo.second();
            tx2.send(()).unwrap();
        }));
    }

    {
        let foo = Arc::clone(&foo);

        handles.push(thread::spawn(move || {
            rx2.recv().unwrap();
            foo.third();
        }));
    }

    for h in handles {
        h.join().unwrap();
    }
}

struct Foo {}

impl Foo {
    fn first(&self) {
        println!("first");
    }

    fn second(&self) {
        println!("second");
    }

    fn third(&self) {
        println!("third");
    }
}
