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
