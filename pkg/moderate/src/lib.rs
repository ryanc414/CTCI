#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_find_intersection() {
        let l1 = Line {
            0: Point { x: 2.0, y: 6.0 },
            1: Point { x: 3.0, y: 9.0 },
        };
        let l2 = Line {
            0: Point { x: 4.0, y: -1.0 },
            1: Point { x: 5.0, y: 2.0 },
        };

        let ix = find_intersection(&l1, &l2);
        assert_eq!(ix, None);

        let l3 = Line {
            0: Point { x: 1.0, y: 10.0 },
            1: Point { x: 5.0, y: 11.0 },
        };

        let ix = find_intersection(&l1, &l3);
        assert_eq!(
            ix,
            Some(Point {
                x: 3.5454545454545454,
                y: 10.636363636363637,
            })
        )
    }

    #[test]
    fn test_find_winner() {
        let grid = TicTacToeGrid::from_string("XX \nO X\n  O").unwrap();
        let winner = grid.find_winner();
        assert_eq!(winner, None);

        let grid = TicTacToeGrid::from_string("XXO\nOOX\nO X").unwrap();
        let winner = grid.find_winner();
        assert_eq!(winner, Some(TicTacToeSymbol::Circle));
    }

    #[test]
    fn test_smallest_diff() {
        let a = vec![1, 3, 15, 11, 2];
        let b = vec![23, 127, 235, 19, 8];

        let diff = smallest_difference(&a, &b);
        assert_eq!(diff, Some(3));
    }

    #[test]
    fn test_number_max() {
        assert_eq!(number_max(3, 5), 5);
        assert_eq!(number_max(1, 1), 1);
        assert_eq!(number_max(44, -12), 44);
        assert_eq!(number_max(-99, -121), -99);
    }

    #[test]
    fn test_english_int() {
        assert_eq!(english_int(0).unwrap().as_str(), "zero");
        assert_eq!(english_int(7).unwrap().as_str(), "seven");
        assert_eq!(english_int(13).unwrap().as_str(), "thirteen");
        assert_eq!(english_int(42).unwrap().as_str(), "forty two");
        assert_eq!(english_int(256).unwrap().as_str(), "two hundred fifty six");
        assert_eq!(
            english_int(1234).unwrap().as_str(),
            "one thousand two hundred thirty four"
        );
        assert_eq!(english_int(1000000).unwrap().as_str(), "one million",);
        assert_eq!(
            english_int(3250000000).unwrap().as_str(),
            "three billion two hundred fifty million"
        );
    }
}

#[derive(Debug, PartialEq)]
pub struct Point {
    x: f64,
    y: f64,
}

#[derive(Debug, PartialEq)]
pub struct Line(Point, Point);

#[derive(Debug, PartialEq)]
struct Coeffs {
    m: f64,
    c: f64,
}

impl Line {
    fn coeffs(&self) -> Coeffs {
        let m = (self.1.y - self.0.y) / (self.1.x - self.0.x);
        let c = self.0.y - (m * self.0.x);

        Coeffs { m, c }
    }
}

pub fn find_intersection(l1: &Line, l2: &Line) -> Option<Point> {
    let Coeffs { m: m1, c: c1 } = l1.coeffs();
    let Coeffs { m: m2, c: c2 } = l2.coeffs();

    if (m1 - m2).abs() < f64::EPSILON {
        return None;
    }

    let x = (c2 - c1) / (m1 - m2);
    let y = (m1 * x) + c1;

    Some(Point { x, y })
}

#[derive(Debug, Copy, Clone, PartialEq)]
pub enum TicTacToeSymbol {
    Circle,
    Cross,
}

#[derive(Debug)]
pub struct ParseErr;

impl TicTacToeSymbol {
    pub fn from(c: char) -> Result<Option<TicTacToeSymbol>, ParseErr> {
        match c {
            'o' | 'O' => Ok(Some(TicTacToeSymbol::Circle)),
            'x' | 'X' => Ok(Some(TicTacToeSymbol::Cross)),
            ' ' => Ok(None),
            _ => Err(ParseErr),
        }
    }
}

pub struct TicTacToeGrid([[Option<TicTacToeSymbol>; 3]; 3]);

impl TicTacToeGrid {
    pub fn from_string(from: &str) -> Result<TicTacToeGrid, ParseErr> {
        let rows: Vec<Vec<Option<TicTacToeSymbol>>> = from
            .split('\n')
            .map(|line| {
                line.chars()
                    .map(|c| TicTacToeSymbol::from(c).unwrap())
                    .collect()
            })
            .collect();

        if rows.len() != 3 {
            return Err(ParseErr);
        }

        for r in rows.iter() {
            if r.len() != 3 {
                return Err(ParseErr);
            }
        }

        let grid: [[Option<TicTacToeSymbol>; 3]; 3] = [
            [rows[0][0], rows[0][1], rows[0][2]],
            [rows[1][0], rows[1][1], rows[1][2]],
            [rows[2][0], rows[2][1], rows[2][2]],
        ];

        Ok(TicTacToeGrid(grid))
    }

    pub fn find_winner(&self) -> Option<TicTacToeSymbol> {
        // rows
        for row in self.0.iter() {
            if let Some(winner) = Self::row_winner(row) {
                return Some(winner);
            }
        }

        // cols
        for i in 0..3 {
            let col: [Option<TicTacToeSymbol>; 3] = [self.0[0][i], self.0[1][i], self.0[2][i]];
            if let Some(winner) = Self::row_winner(&col) {
                return Some(winner);
            }
        }

        // diags
        let diag = [self.0[0][0], self.0[1][1], self.0[2][2]];
        if let Some(winner) = Self::row_winner(&diag) {
            return Some(winner);
        }

        let diag = [self.0[0][2], self.0[1][1], self.0[2][0]];
        if let Some(winner) = Self::row_winner(&diag) {
            return Some(winner);
        }

        None
    }

    fn row_winner(row: &[Option<TicTacToeSymbol>; 3]) -> Option<TicTacToeSymbol> {
        if row == &[Some(TicTacToeSymbol::Circle); 3] {
            return Some(TicTacToeSymbol::Circle);
        }

        if row == &[Some(TicTacToeSymbol::Cross); 3] {
            return Some(TicTacToeSymbol::Cross);
        }

        None
    }
}

pub fn smallest_difference(a: &[i64], b: &[i64]) -> Option<i64> {
    if a.is_empty() || b.is_empty() {
        return None;
    }

    let mut a = a.to_owned();
    a.sort_unstable();
    let mut b = b.to_owned();
    b.sort_unstable();

    let mut i: usize = 0;
    let mut smallest_diff: Option<i64> = None;

    for n in a {
        let (diff, j) = smallest_diff_from(n, &b[i..]).unwrap();
        if smallest_diff.is_none() || diff < smallest_diff.unwrap() {
            smallest_diff = Some(diff);
        }
        i += j;
    }

    smallest_diff
}

fn smallest_diff_from(n: i64, b: &[i64]) -> Option<(i64, usize)> {
    let mut smallest_diff_vals: Option<(i64, usize)> = None;

    for (i, &m) in b.iter().enumerate() {
        let diff = if m > n { m - n } else { n - m };

        match smallest_diff_vals {
            None => smallest_diff_vals = Some((diff, i)),
            Some((smol_diff, _)) => {
                if diff < smol_diff {
                    smallest_diff_vals = Some((diff, i));
                } else {
                    return smallest_diff_vals;
                }
            }
        }
    }

    smallest_diff_vals
}

pub fn number_max(x: i64, y: i64) -> i64 {
    let diff = (y - x).abs();
    (x + y + diff) / 2
}

const DIGITS: [&str; 10] = [
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
];

const TEENS: [&str; 10] = [
    "ten",
    "eleven",
    "twelve",
    "thirteen",
    "fourteen",
    "fifteen",
    "sixteen",
    "seventeen",
    "eighteen",
    "nineteen",
];

const TENS: [&str; 8] = [
    "twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety",
];

const HUNDRED: &str = "hundred";

const THOUSANDS: [&str; 4] = ["thousand", "million", "billion", "trillion"];

const MAX: i64 = 10000000000000;

pub fn english_int(x: i64) -> Result<String, ()> {
    if x > MAX || x < 0 {
        return Err(());
    }

    if x == 0 {
        return Ok(String::from(DIGITS[0]));
    }

    let mut s = String::new();

    for i in 0..4 {
        let limit: i64 = 10i64.pow(3 * (4 - i));
        if x < limit {
            continue;
        }

        let base = (x / limit) % 1000;
        if base == 0 {
            continue;
        }

        let text = english_int_base(base).unwrap();
        s.push_str(&text);
        s.push(' ');
        s.push_str(THOUSANDS[(4 - i - 1) as usize]);
        s.push(' ');
    }

    let rest = x % 1000;
    if rest > 0 {
        let text = english_int_base(rest).unwrap();
        s.push_str(&text);
    }

    s = s.trim_end().to_owned();

    Ok(s)
}

fn english_int_base(x: i64) -> Result<String, ()> {
    if x >= 1000 || x < 0 {
        return Err(());
    }

    let mut s = String::new();

    if x > 100 {
        let hundreds = x / 100;
        s.push_str(DIGITS[hundreds as usize]);
        s.push(' ');
        s.push_str(HUNDRED);
        s.push(' ');
    }

    let tens = x % 100;
    if tens >= 20 {
        let i = tens / 10;
        s.push_str(TENS[(i - 2) as usize]);
        if tens % 10 > 0 {
            s.push(' ');
            s.push_str(DIGITS[(tens % 10) as usize]);
        }
    } else if tens >= 10 {
        s.push_str(TEENS[(tens - 10) as usize]);
    } else if tens > 0 {
        s.push_str(DIGITS[tens as usize]);
    }

    Ok(s)
}
