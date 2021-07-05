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
