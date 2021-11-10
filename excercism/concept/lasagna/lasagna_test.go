func TestOvenTime (t *testing.T) {
  
  tt := lasagnaTests{
                  name: "Calculates how many minutes the lasagna should be in the oven",
    layers: 0,
    time: 40,
    expected: 40,
  }
  
  if got := OvenTime; got != tt.expected {
    t.Errorf("OvenTime(%d) = %d; want %d", tt.expected, got, tt.expected)
  }

}
