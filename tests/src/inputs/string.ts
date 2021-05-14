export function main(a: String): String {
  // It appears not all standard JavaScript methods exists here, `toUpperCase` for example leads to a compilation error
  return a.slice(0, 5)
}
