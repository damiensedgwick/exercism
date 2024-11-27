import gleam/float as f
import gleam/int as i
import gleam/string as s

pub fn pence_to_pounds(pence: Int) -> Float {
  i.to_float(pence) /. 100.0
}

pub fn pounds_to_string(pounds: Float) -> String {
  let symbol: String = "£"
  let pounds_str: String = f.to_string(pounds)
  s.concat([symbol, pounds_str])
}
