import gleam/int as i
import gleam/list as l

pub fn square_of_sum(n: Int) -> Int {
  let sum =
    l.range(1, n)
    |> i.sum

  sum * sum
}

pub fn sum_of_squares(n: Int) -> Int {
  l.range(from: 1, to: n)
  |> l.map(fn(a) { a * a })
  |> i.sum
}

pub fn difference(n: Int) -> Int {
  square_of_sum(n) - sum_of_squares(n)
}
